package uptime

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"armada-node/model"

	"go.uber.org/zap"
)

const (
	DefaultIntervalDuration = 5 * time.Minute

	DefaultWorkersPerHost    = 10
	DefaultBufferSizePerHost = 100

	DefaultNodeRefreshInterval = 1 * time.Minute
	DefaultNodeRefreshTimeout  = 15 * time.Second

	DefaultExportIntervalTimeout = 15 * time.Second
)

var (
	DefaultHTTPClient = &http.Client{Timeout: 15 * time.Second}
)

type ProbeCounts struct {
	Success uint32
	Failure uint32
}

type IntervalData struct {
	StartTime time.Time
	EndTime   time.Time

	// RequestCount is the number of first-try content requests served by this
	// node between StartTime and EndTime. Since this only tracks "first-try"
	// requests, retry requests and probe requests are excluded.
	RequestCount uint32

	// ProbeResults contains success/failure counts for probes that were sent
	// to peers from this node. The map key is the peer's hostname.
	ProbeResults map[string]ProbeCounts
}

type IntervalStore interface {
	Put(context.Context, IntervalData) error
	Range(ctx context.Context, start, end time.Time, fn func(IntervalData) error) error
}

type ReadOnlyNode interface {
	ID() model.ID
	Host() string
	ProjectID() model.ID
}

// MeterArgs are required arguments for initializing an uptime Meter.
type MeterArgs struct {
	Model  model.Client
	Self   ReadOnlyNode
	Store  IntervalStore
	Logger *zap.Logger
}

// MeterOptions are optional arguments for configuring an uptime Meter.
type MeterOptions struct {
	// IntervalDuration specifies the duration of an uptime interval.
	// This is intended to be globally defined by the network.
	IntervalDuration time.Duration

	// WorkersPerHost controls the concurrency of the underlying probing system.
	// Meter will run up to WorkersPerHost probes in parallel, per target host.
	WorkersPerHost int

	// BufferSizePerHost specifies the queue length for probe requests. If the
	// buffer is full, any additional probe requests will be dropped.
	BufferSizePerHost int

	// HTTPClient to use for making probe requests.
	HTTPClient *http.Client

	// ModelRefreshInterval specifies how often to check for new content nodes.
	NodeRefreshInterval time.Duration

	// NodeRefreshTimeout specifies the maximum amount of time that's allowed
	// for updating the current list of content nodes.
	NodeRefreshTimeout time.Duration

	// ExportIntervalTimeout specifies the maximum amount of time that's allowed
	// for saving the most recent interval data.
	ExportIntervalTimeout time.Duration

	// Custom clock for testing.
	clock clock
}

// A Meter measures the uptime of content nodes by issuing probes and tracking
// their outcomes. It divides time into intervals (as dictated by the IntervalDuration
// option), reporting on each node's uptime throughout a given interval. Uptime is
// simply defined as the ratio of probes that were successful in an interval.
type Meter struct {
	m     model.Client
	self  ReadOnlyNode
	store IntervalStore
	opts  MeterOptions

	interval     int64
	requestCount uint32

	probers     map[string]*prober
	probersLock *sync.RWMutex

	stopCh chan struct{}
	stopWg *sync.WaitGroup

	logger *zap.Logger

	// For testing only.
	clock clock
}

func NewMeter(args MeterArgs, opts MeterOptions) (*Meter, error) {
	// Require all arguments.
	if args.Model == nil {
		return nil, errors.New("Model is required")
	}
	if args.Self == nil {
		return nil, errors.New("Self is required")
	}
	if args.Store == nil {
		return nil, errors.New("Store is required")
	}
	if args.Logger == nil {
		return nil, errors.New("Logger is required")
	}

	// Apply defaults to all options.
	if opts.IntervalDuration == 0 {
		opts.IntervalDuration = DefaultIntervalDuration
	}
	if opts.WorkersPerHost == 0 {
		opts.WorkersPerHost = DefaultWorkersPerHost
	}
	if opts.BufferSizePerHost == 0 {
		opts.BufferSizePerHost = DefaultBufferSizePerHost
	}
	if opts.HTTPClient == nil {
		opts.HTTPClient = DefaultHTTPClient
	}
	if opts.NodeRefreshInterval == 0 {
		opts.NodeRefreshInterval = DefaultNodeRefreshInterval
	}
	if opts.NodeRefreshTimeout == 0 {
		opts.NodeRefreshTimeout = DefaultNodeRefreshTimeout
	}
	if opts.ExportIntervalTimeout == 0 {
		opts.ExportIntervalTimeout = DefaultExportIntervalTimeout
	}
	if opts.clock == nil {
		opts.clock = realClock{}
	}

	return &Meter{
		m:     args.Model,
		self:  args.Self,
		store: args.Store,
		opts:  opts,

		probers:     make(map[string]*prober),
		probersLock: &sync.RWMutex{},

		stopCh: make(chan struct{}),
		stopWg: &sync.WaitGroup{},

		logger: args.Logger.Named("uptime"),
		clock:  opts.clock,
	}, nil
}

func (m *Meter) Start() {
	m.stopWg.Add(1)
	go func() {
		defer m.stopWg.Done()
		m.runProberManagement()
	}()

	m.stopWg.Add(1)
	go func() {
		defer m.stopWg.Done()
		m.runIntervalManagement()
	}()
}

func (m *Meter) Stop() {
	close(m.stopCh)
	m.stopWg.Wait()

	m.probersLock.Lock()
	defer m.probersLock.Unlock()
	for _, prb := range m.probers {
		prb.Stop(true)
	}
}

// Probe instructs the Meter to probe the given URL, expecting the response body
// content to match the provided checksum. Returns true if the probe was enqueued
// for asynchronous execution, false if it was discarded due to a full buffer.
// If the probe was successfully enqueued and resultCh is provided, the final result
// of the probe attempt will be sent on resultCh once it completes.
func (m *Meter) Probe(probeURL *url.URL, checksum string, resultCh chan<- bool) bool {
	m.probersLock.RLock()
	defer m.probersLock.RUnlock()

	prb, ok := m.probers[probeURL.Host]
	if !ok {
		m.logger.Warn("Ignoring probe request, unknown host", zap.String("host", probeURL.Host))
		return false
	}
	return prb.Enqueue(probeURL, checksum, resultCh)
}

func (m *Meter) Results(ctx context.Context, start, end time.Time, fn func(IntervalData) error) error {
	return m.store.Range(ctx, start, end, fn)
}

func (m *Meter) IncrementRequestCount(delta uint32) {
	atomic.AddUint32(&m.requestCount, delta)
}

func (m *Meter) runProberManagement() {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), m.opts.NodeRefreshTimeout)
		if err := m.manageProbers(ctx); err != nil {
			m.logger.Error("Failed to manage probers", zap.Error(err))
		}
		cancel()

		select {
		case <-m.stopCh:
			return
		case <-time.After(m.opts.NodeRefreshInterval):
		}
	}
}

func (m *Meter) manageProbers(ctx context.Context) error {
	// Fetch an up-to-date set of peer content nodes and construct a lookup map for later.
	if m.self.ProjectID().IsZero() {
		return nil
	}
	nodes, err := m.m.ContentNodes(ctx, m.self.ProjectID())
	if err != nil {
		return fmt.Errorf("fetching content nodes: %v", err)
	}
	nodeHosts := make(map[string]struct{})
	for _, node := range nodes {
		nodeHosts[node.Host] = struct{}{}
	}

	// Prevent probing ourselves.
	delete(nodeHosts, m.self.Host())

	// Determine if any new probers are needed or if any of the existing
	// probers are now obsolete.
	var addHosts []string
	var removeHosts []string
	m.probersLock.RLock()
	for host := range nodeHosts {
		if _, ok := m.probers[host]; !ok {
			addHosts = append(addHosts, host)
		}
	}
	for host := range m.probers {
		if _, ok := nodeHosts[host]; !ok {
			removeHosts = append(removeHosts, host)
		}
	}
	m.probersLock.RUnlock()

	// No changes necessary.
	if len(addHosts) == 0 && len(removeHosts) == 0 {
		return nil
	}

	// We need to modify probers somehow, so grab the exclusive lock.
	m.probersLock.Lock()
	defer m.probersLock.Unlock()

	// Add missing probers.
	for _, host := range addHosts {
		m.logger.Debug("Adding prober", zap.String("host", host))
		m.probers[host] = newProber(m.logger, m.opts.BufferSizePerHost, m.opts.HTTPClient)
		m.probers[host].Start(m.opts.WorkersPerHost)
	}

	// Tear down obsolete probers.
	for _, host := range removeHosts {
		m.logger.Debug("Removing prober", zap.String("host", host))
		if prb, ok := m.probers[host]; ok {
			prb.Stop(false)
			delete(m.probers, host)
		}
	}

	return nil
}

func (m *Meter) runIntervalManagement() {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), m.opts.ExportIntervalTimeout)
		if err := m.manageInterval(ctx); err != nil {
			m.logger.Error("Failed to manage interval", zap.Error(err))
		}
		cancel()

		// Sleep until the next interval starts.
		nextStart := time.Unix((m.interval+1)*int64(m.opts.IntervalDuration.Seconds()), 0)

		select {
		case <-m.stopCh:
			return
		case <-time.After(time.Until(nextStart)):
		}
	}
}

func (m *Meter) manageInterval(ctx context.Context) error {
	// Determine if we should be advancing to a new interval.
	newInterval := m.clock.Now().Unix() / int64(m.opts.IntervalDuration.Seconds())
	if newInterval == m.interval {
		return nil
	}
	if newInterval < m.interval {
		return fmt.Errorf("decreasing interval: current=%d new=%d", m.interval, newInterval)
	}

	// Export the current interval. If there's an error we log it but just continue
	// advancing to the new interval because the probe counters will have reset anyway.
	if err := m.persistCurrentInterval(ctx); err != nil {
		m.logger.Error("Failed to export interval", zap.Error(err))
	}

	// Start the new interval.
	m.interval = newInterval
	atomic.StoreUint32(&m.requestCount, 0)

	return nil
}

func (m *Meter) persistCurrentInterval(ctx context.Context) error {
	if m.interval == 0 {
		return nil
	}

	data := IntervalData{
		StartTime:    time.Unix(m.interval*int64(m.opts.IntervalDuration.Seconds()), 0),
		EndTime:      time.Unix((m.interval+1)*int64(m.opts.IntervalDuration.Seconds()), 0),
		RequestCount: atomic.LoadUint32(&m.requestCount),
		ProbeResults: make(map[string]ProbeCounts),
	}

	m.probersLock.RLock()
	for host, prb := range m.probers {
		counts := ProbeCounts{}
		counts.Success, counts.Failure = prb.ReportAndReset()
		if counts.Success == 0 && counts.Failure == 0 {
			continue
		}
		data.ProbeResults[host] = counts
	}
	m.probersLock.RUnlock()

	// Only persist intervals that have non-zero data.
	if data.RequestCount == 0 && len(data.ProbeResults) == 0 {
		return nil
	}

	return m.store.Put(ctx, data)
}
