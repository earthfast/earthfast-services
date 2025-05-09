package uptime

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sort"
	"time"

	"armada-node/api"
	"armada-node/model"
)

type fetchResult struct {
	resp *api.UptimeResponse
	err  error
}

type Node interface {
	Host() string
	Uptime(ctx context.Context, resCh chan<- fetchResult, startTime, endTime time.Time)
}

type NodeProvider interface {
	AllNodes(context.Context, model.ID) ([]Node, error)
}

type UptimeReport struct {
	StartTime    time.Time
	EndTime      time.Time
	UptimeRatios map[string]float64
}

type ReportOptions struct {
	// StartTime is the beginning of the report window.
	StartTime time.Time

	// EndTime is the end of the report window.
	EndTime time.Time

	// RequestCountPercentile is the percentile to use when selecting the consensus
	// value for per-node request counts throughout the report window.
	RequestCountPercentile int

	// ProbeFailureCountPercentile is the percentile to use when selecting the consensus
	// value for per-node probe failure counts throughout the report window.
	ProbeFailureCountPercentile int

	// NominalUptimeRatioPercentile is the percentile to use when selecting an uptime
	// ratio that will be considered to be nominal performance. All other uptime ratios
	// will be scaled relative to the chosen ratio.
	NominalUptimeRatioPercentile int
}

type Reporter struct {
	np   NodeProvider
	opts ReportOptions
}

func NewReporter(np NodeProvider, opts ReportOptions) (*Reporter, error) {
	if opts.StartTime.IsZero() {
		return nil, errors.New("StartTime is required")
	}
	if opts.EndTime.IsZero() {
		return nil, errors.New("EndTime is required")
	}
	if opts.RequestCountPercentile < 0 || opts.RequestCountPercentile > 100 {
		return nil, fmt.Errorf("RequestCountPercentile is out of bounds [0, 100]: %d", opts.RequestCountPercentile)
	}
	if opts.ProbeFailureCountPercentile < 0 || opts.ProbeFailureCountPercentile > 100 {
		return nil, fmt.Errorf("ProbeFailureCountPercentile is out of bounds [0, 100]: %d", opts.ProbeFailureCountPercentile)
	}
	if opts.NominalUptimeRatioPercentile < 0 || opts.NominalUptimeRatioPercentile > 100 {
		return nil, fmt.Errorf("NominalUptimeRatioPercentile is out of bounds [0, 100]: %d", opts.NominalUptimeRatioPercentile)
	}
	return &Reporter{
		np:   np,
		opts: opts,
	}, nil
}

func (r *Reporter) Run(ctx context.Context, projectID model.ID) (UptimeReport, error) {
	nodes, err := r.np.AllNodes(ctx, projectID)
	if err != nil {
		return UptimeReport{}, fmt.Errorf("fetching nodes: %v", err)
	}
	if len(nodes) == 0 {
		return UptimeReport{}, errors.New("no nodes available")
	}

	results, err := r.fetch(ctx, nodes)
	if err != nil {
		return UptimeReport{}, fmt.Errorf("fetching uptime reports: %v", err)
	}

	return r.generate(nodes, results)
}

func (r *Reporter) fetch(ctx context.Context, nodes []Node) ([]fetchResult, error) {
	// Request all uptime reports in parallel.
	resCh := make(chan fetchResult, len(nodes))
	for _, node := range nodes {
		go func(n Node) {
			n.Uptime(ctx, resCh, r.opts.StartTime, r.opts.EndTime)
		}(node)
	}

	// Wait for every uptime report, exiting early if any of them failed.
	var results []fetchResult
	for _, node := range nodes {
		select {
		case res := <-resCh:
			if res.err != nil {
				return nil, fmt.Errorf("retrieving uptime report from node %s: %v", node.Host(), res.err)
			}
			results = append(results, res)
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	return results, nil
}

func (r *Reporter) generate(nodes []Node, results []fetchResult) (UptimeReport, error) {
	// Combine the individual uptime reports into a collection of slices, maintaining
	// all of the original data points for the time being.
	combined := struct {
		requestCounts     []int
		failedProbeCounts map[string][]int
	}{
		failedProbeCounts: make(map[string][]int),
	}
	for _, res := range results {
		combined.requestCounts = append(combined.requestCounts, res.resp.RequestCount)
		for host, counts := range res.resp.ProbeResults {
			combined.failedProbeCounts[host] = append(combined.failedProbeCounts[host], counts.Failure)
		}
	}

	// Sort all the combined slices so we'll be able to select percentile values.
	sort.Ints(combined.requestCounts)
	for _, failures := range combined.failedProbeCounts {
		sort.Ints(failures)
	}

	// Choose an authoritative value for:
	//   1. The non-retry end user request count that each node should have served.
	//   2. The number of failed probe requests on a node as observed by a single peer node.
	requestCount, err := percentileInt(r.opts.RequestCountPercentile, combined.requestCounts)
	if err != nil {
		return UptimeReport{}, fmt.Errorf("taking percentile of request counts: %v", err)
	}
	failureCounts := make(map[string]int)
	for host, counts := range combined.failedProbeCounts {
		val, err := percentileInt(r.opts.ProbeFailureCountPercentile, counts)
		if err != nil {
			return UptimeReport{}, fmt.Errorf("taking percentile of failure counts for %s: %v", host, err)
		}
		failureCounts[host] = val
	}

	// Using the authoritative values from above, compute the uptime ratio for each node.
	numPeers := len(nodes) - 1
	rawRatios := make(map[string]float64)
	var allRatios []float64
	for _, node := range nodes {
		host := node.Host()
		scaledFailureCount := failureCounts[host] * numPeers
		if requestCount == 0 {
			rawRatios[host] = 1
		} else if scaledFailureCount > requestCount {
			rawRatios[host] = 0
		} else {
			rawRatios[host] = 1 - (float64(scaledFailureCount) / float64(requestCount))
		}
		allRatios = append(allRatios, rawRatios[host])
	}

	// Choose the target uptime ratio.
	sort.Float64s(allRatios)
	targetRatio, err := percentileFloat64(r.opts.NominalUptimeRatioPercentile, allRatios)
	if err != nil {
		return UptimeReport{}, fmt.Errorf("taking percentile of uptime ratios: %v", err)
	}

	// Compile the final report.
	report := UptimeReport{
		StartTime:    r.opts.StartTime,
		EndTime:      r.opts.EndTime,
		UptimeRatios: make(map[string]float64),
	}
	for host, ratio := range rawRatios {
		report.UptimeRatios[host] = math.Min(1.0, ratio/targetRatio)
	}
	return report, nil
}

func percentileInt(pct int, vals []int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("no values")
	}
	if pct < 0 || pct > 100 {
		return 0, fmt.Errorf("percentile out of bounds [0, 100]: %d", pct)
	}
	if !sort.IntsAreSorted(vals) {
		return 0, fmt.Errorf("values must be sorted")
	}
	if pct == 100 {
		return vals[len(vals)-1], nil
	}
	return vals[int(float64(len(vals))*float64(pct)/100.0)], nil
}

func percentileFloat64(pct int, vals []float64) (float64, error) {
	if len(vals) == 0 {
		return 0, errors.New("no values")
	}
	if pct < 0 || pct > 100 {
		return 0, fmt.Errorf("percentile out of bounds [0, 100]: %d", pct)
	}
	if !sort.Float64sAreSorted(vals) {
		return 0, fmt.Errorf("values must be sorted")
	}
	if pct == 100 {
		return vals[len(vals)-1], nil
	}
	return vals[int(float64(len(vals))*float64(pct)/100.0)], nil
}
