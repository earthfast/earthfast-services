package uptime

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"go.uber.org/zap/zaptest"
)

func newNodeServer(t *testing.T, body string) (cleanup func(), tsURL *url.URL, checksum string) {
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, body)
	}))
	cleanup = ts.Close

	var err error
	if tsURL, err = url.Parse(ts.URL); err != nil {
		t.Fatal(err)
	}

	hash := sha1.New()
	hash.Write([]byte(body))
	checksum = fmt.Sprintf("%x", hash.Sum(nil))

	return
}

func TestProber(t *testing.T) {
	logger := zaptest.NewLogger(t)

	cleanup, probeURL, wantChecksum := newNodeServer(t, "Hello, world!")
	defer cleanup()

	p := newProber(logger, 10, http.DefaultClient)
	p.Start(5)
	defer p.Stop(true)

	resultCh := make(chan bool)
	p.Enqueue(probeURL, wantChecksum, resultCh)
	p.Enqueue(probeURL, "bad checksum", resultCh)
	p.Enqueue(probeURL, wantChecksum, resultCh)
	for i := 0; i < 3; i++ {
		select {
		case <-resultCh:
		case <-time.After(3 * time.Second):
			t.Fatal("Timed out waiting for probe results")
		}
	}

	gotSuccess, gotFail := p.ReportAndReset()
	wantSuccess, wantFail := uint32(2), uint32(1)
	if gotSuccess != wantSuccess || gotFail != wantFail {
		t.Errorf("Incorrect result counts: got (%d, %d), want (%d, %d)", gotSuccess, gotFail, wantSuccess, wantFail)
	}

	gotSuccess, gotFail = p.ReportAndReset()
	wantSuccess, wantFail = uint32(0), uint32(0)
	if gotSuccess != wantSuccess || gotFail != wantFail {
		t.Errorf("Incorrect result counts: got (%d, %d), want (%d, %d)", gotSuccess, gotFail, wantSuccess, wantFail)
	}
}

func TestProber_FullBuffer(t *testing.T) {
	logger := zaptest.NewLogger(t)

	p := newProber(logger, 2, http.DefaultClient)

	// Note: purposefully not calling p.Start() here so the buffer fills up.

	if ok := p.Enqueue(nil, "checksum", nil); !ok {
		t.Fatal("Enqueue with empty buffer failed, wanted success")
	}
	if ok := p.Enqueue(nil, "checksum", nil); !ok {
		t.Fatal("Enqueue with available buffer failed, wanted success")
	}
	if ok := p.Enqueue(nil, "checksum", nil); ok {
		t.Fatal("Enqueue with full buffer succeeded, wanted failure")
	}
}
