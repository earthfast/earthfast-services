package tarballfs

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
)

const (
	defaultReadTimeout   = 15 * time.Second
	defaultRetryDuration = 30 * time.Second
)

var (
	defaultClient = &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   3 * time.Minute,
	}
)

type Options struct {
	MountDir  string
	SourceURL string
	Checksum  string

	Client        *http.Client
	ReadTimeout   time.Duration
	RetryDuration time.Duration

	Logger *zap.Logger
}

type FS struct {
	mountDir  string
	sourceURL *url.URL
	checksum  string

	client        *http.Client
	readTimeout   time.Duration
	retryDuration time.Duration

	readyCh   chan struct{}
	loadErr   error
	loadErrCh chan struct{}

	stopCh chan struct{}
	stopWg sync.WaitGroup

	logger *zap.Logger
}

func New(opts Options) (*FS, error) {
	if opts.MountDir == "" {
		return nil, errors.New("MountDir is required")
	}
	if opts.SourceURL == "" {
		return nil, errors.New("SourceURL is required")
	}
	sourceURL, err := url.Parse(opts.SourceURL)
	if err != nil {
		return nil, fmt.Errorf("parsing SourceURL: %v", err)
	}
	if opts.Checksum == "" {
		return nil, errors.New("Checksum is required")
	}
	if opts.Logger == nil {
		return nil, errors.New("Logger is required")
	}
	if opts.Client == nil {
		opts.Client = defaultClient
	}
	if opts.ReadTimeout == 0 {
		opts.ReadTimeout = defaultReadTimeout
	}
	if opts.RetryDuration == 0 {
		opts.RetryDuration = defaultRetryDuration
	}

	return &FS{
		mountDir:  filepath.Clean(opts.MountDir),
		sourceURL: sourceURL,
		checksum:  opts.Checksum,

		client:        opts.Client,
		readTimeout:   opts.ReadTimeout,
		retryDuration: opts.RetryDuration,

		readyCh:   make(chan struct{}),
		loadErrCh: make(chan struct{}),
		stopCh:    make(chan struct{}),

		logger: opts.Logger.Named("tarballfs").With(zap.String("url", opts.SourceURL)),
	}, nil
}

func (tbfs *FS) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-tbfs.readyCh:
		case <-tbfs.stopCh:
		}
		cancel()
	}()

	tbfs.stopWg.Add(1)
	go func() {
		defer tbfs.stopWg.Done()
		tbfs.run(ctx)
	}()
}

func (tbfs *FS) Stop() {
	close(tbfs.stopCh)
	tbfs.stopWg.Wait()
}

func (tbfs *FS) Delete() error {
	if err := os.RemoveAll(tbfs.mountDir); err != nil {
		return fmt.Errorf("removing mount directory: %v", err)
	}
	return nil
}

func (tbfs *FS) Open(name string) (fs.File, error) {
	select {
	case <-tbfs.readyCh:
	default:
		select {
		case <-tbfs.readyCh:
		case <-tbfs.loadErrCh:
			return nil, tbfs.loadErr
		case <-tbfs.stopCh:
			return nil, errors.New("TarballFS stopping")
		case <-time.After(tbfs.readTimeout):
			return nil, errors.New("TarballFS read timeout")
		}
	}
	return os.Open(filepath.Join(tbfs.mountDir, name))
}

func (tbfs *FS) run(ctx context.Context) {
	attempt := 0
	for {
		attempt++
		err := tbfs.load(ctx)

		// Success, we're done
		if err == nil {
			close(tbfs.readyCh)
			return
		}

		tbfs.logger.Warn("failed to load",
			zap.Int("attempt", attempt),
			zap.Error(err),
		)
		tbfs.loadErr = err
		if attempt == 1 {
			close(tbfs.loadErrCh)
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(tbfs.retryDuration):
		}
	}
}

func (tbfs *FS) load(ctx context.Context) error {
	// If the destination directory already exists, there's nothing to do.
	if _, err := os.Stat(tbfs.mountDir); err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("getting fileinfo for %s: %v", tbfs.mountDir, err)
	}

	// Create a temporary directory to use as scratch space that automatically
	// gets erased at the end of this method call. Since we know that we have
	// write permissions for the parent of the mount directory, we choose to
	// put the staging directory there as well (i.e. as a sibling to the target
	// directory).
	stageDir, err := os.MkdirTemp(
		filepath.Dir(tbfs.mountDir),
		fmt.Sprintf("%s-staging-", filepath.Base(tbfs.mountDir)),
	)
	if err != nil {
		return fmt.Errorf("creating staging directory: %v", err)
	}
	defer os.RemoveAll(stageDir)

	// Create an empty file to hold the downloaded tarball
	tarball, err := os.Create(filepath.Join(stageDir, "source.tar.gz"))
	if err != nil {
		return fmt.Errorf("creating tarball destination file: %v", err)
	}
	defer tarball.Close()

	// Fetch the remote tarball
	tbfs.logger.Debug("fetching tarball")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tbfs.sourceURL.String(), nil)
	if err != nil {
		return fmt.Errorf("creating http request: %v", err)
	}
	resp, err := tbfs.client.Do(req)
	if err != nil {
		return fmt.Errorf("fetching tarball: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("http %d", resp.StatusCode)
	}

	// Write tarball data to the local file while simultaneously computing its checksum
	hash := sha256.New()
	_, err = io.Copy(tarball, io.TeeReader(resp.Body, hash))
	if err != nil {
		return fmt.Errorf("writing tarball data: %v", err)
	}
	if stats, err := tarball.Stat(); err == nil {
		tbfs.logger.Debug("retrieved tarball",
			zap.Int64("bytes", stats.Size()),
		)
	}

	// Validate checksum
	if got := fmt.Sprintf("%x", hash.Sum(nil)); got != tbfs.checksum {
		return fmt.Errorf("checksum mismatch: got %s, want %s", got, tbfs.checksum)
	}

	// Extract the tarball into the staging directory
	outDir := path.Join(stageDir, "contents")
	if err := os.Mkdir(outDir, 0755); err != nil {
		return fmt.Errorf("creating tarball contents directory: %v", err)
	}
	tarball.Seek(0, 0)
	if err := extract(tbfs.logger, outDir, tarball, 1); err != nil {
		return fmt.Errorf("extracting tarball: %v", err)
	}

	// Everything succeeded, move the output into its final position
	if err := os.Rename(outDir, tbfs.mountDir); err != nil {
		return fmt.Errorf("finalizing staging directory: %v", err)
	}
	return nil
}
