package uptime

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"

	"go.uber.org/zap"
)

type probeRequest struct {
	url      *url.URL
	checksum string
	resultCh chan<- bool
}

type prober struct {
	httpClient *http.Client

	successCount uint32
	failureCount uint32
	requestCh    chan probeRequest

	stopCh chan struct{}
	stopWg *sync.WaitGroup

	logger *zap.Logger
}

func newProber(logger *zap.Logger, bufferSize int, httpClient *http.Client) *prober {
	return &prober{
		httpClient: httpClient,

		requestCh: make(chan probeRequest, bufferSize),

		stopCh: make(chan struct{}),
		stopWg: &sync.WaitGroup{},

		logger: logger,
	}
}

func (p *prober) Start(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		p.stopWg.Add(1)
		go func() {
			defer p.stopWg.Done()
			p.run()
		}()
	}
}

func (p *prober) Stop(blocking bool) {
	close(p.stopCh)
	if blocking {
		p.stopWg.Wait()
	}
}

// Enqueue requests that a new probe be run for the given URL, returning true if it was
// successfully enqueued, false otherwise (i.e. the buffer was full). If the probe was
// successfully enqueued and resultCh is provided, the final result of the probe attempt
// will be sent on resultCh once it completes.
func (p *prober) Enqueue(probeURL *url.URL, checksum string, resultCh chan<- bool) bool {
	req := probeRequest{
		url:      probeURL,
		checksum: checksum,
		resultCh: resultCh,
	}
	select {
	case p.requestCh <- req:
		return true
	default:
		p.logger.Warn("Dropping probe request, buffer is full",
			zap.Stringer("url", probeURL),
			zap.String("checksum", checksum),
		)
		return false
	}
}

// ReportAndReset returns the success and failure counts that have accrued since the
// last time ReportAndReset was called.
func (p *prober) ReportAndReset() (successCount, failureCount uint32) {
	return atomic.SwapUint32(&p.successCount, 0), atomic.SwapUint32(&p.failureCount, 0)
}

// run repeatedly consumes work in the form of probeRequests until stopped.
func (p *prober) run() {
	for {
		select {
		case <-p.stopCh:
			return
		case req := <-p.requestCh:
			ok := p.handleRequest(req)
			if req.resultCh != nil {
				req.resultCh <- ok
			}
		}
	}
}

// handleRequest executes a single probe and records the result.
func (p *prober) handleRequest(req probeRequest) bool {
	if err := p.doProbe(req); err != nil {
		p.logger.Warn("Probe failed", zap.Error(err))
		atomic.AddUint32(&p.failureCount, 1)
		return false
	}

	atomic.AddUint32(&p.successCount, 1)
	return true
}

// doProbe performs an HTTP request and validates the checksum of the response body.
func (p *prober) doProbe(req probeRequest) error {
	resp, err := p.httpClient.Get(req.url.String())
	if err != nil {
		return fmt.Errorf("making request: %v", err)
	}
	defer resp.Body.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, resp.Body); err != nil {
		return fmt.Errorf("reading body: %v", err)
	}
	if got := fmt.Sprintf("%x", hash.Sum(nil)); got != req.checksum {
		return fmt.Errorf("incorrect checksum: got %s, want %s", got, req.checksum)
	}

	return nil
}
