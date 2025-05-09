package site

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"armada-node/model"

	"go.uber.org/zap"
)

var (
	staleVersionDuration = 10 * time.Second
)

// Version is an http.Handler that is capable of serving a single, specific
// instance of a Site. Therefore, a Site can (and should) have many Versions
// throughout its lifetime.
//
// Since the resources necessary to serve a Version aren't always immediately
// available, some lifecycle management is necessary. See Start() and Stop().
type Version interface {
	fmt.Stringer
	http.Handler

	// Start begins any asyncronous work necessary to serve the Version. Start
	// must be called before the Version's http.Handler is invoked.
	Start()

	// Stop terminates any asynchronous work that was created by Start().
	Stop()

	// IsProject returns true if the provided Project accurately reflects the
	// Version's content, false otherwise.
	IsProject(*model.Project) bool

	// Delete causes the Version to remove any persisted data it created.
	// Since this is destructive, it must only be called when the Version
	// is in a stopped state.
	Delete() error
}

type VersionProvider interface {
	VersionForProject(*model.Project) (Version, error)
}

type Handler struct {
	m         model.Client
	vp        VersionProvider
	projectID model.ID

	live          Version
	liveLock      *sync.RWMutex
	liveStaleTime time.Time

	stopped bool
	logger  *zap.Logger
}

func NewHandler(logger *zap.Logger, m model.Client, vp VersionProvider, projectID model.ID) *Handler {
	return &Handler{
		m:         m,
		vp:        vp,
		projectID: projectID,

		liveLock: &sync.RWMutex{},

		logger: logger.With(zap.String("projectID", fmt.Sprintf("%x", projectID))),
	}
}

func (h *Handler) Stop() {
	h.liveLock.Lock()
	defer h.liveLock.Unlock()

	h.setLiveVersionLocked(nil)
	h.stopped = true
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.liveLock.RLock()
	version := h.getLiveVersionLocked()
	h.liveLock.RUnlock()

	if version == nil {
		var err error
		if version, err = h.updateLiveVersion(r.Context()); err != nil {
			h.logger.Error("failed to get live version", zap.Error(err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	version.ServeHTTP(w, r)
}

func (h *Handler) getLiveVersionLocked() Version {
	if h.live == nil || time.Now().After(h.liveStaleTime) {
		return nil
	}
	return h.live
}

func (h *Handler) setLiveVersionLocked(ver Version) {
	// Stop the currently serving version.
	if h.live != nil {
		h.live.Stop()

		// Delete the old version if it's being replaced by a new version.
		// We don't just always delete because we want to leave the data in
		// place in the case of a shutdown.
		if ver != nil {
			if err := h.live.Delete(); err != nil {
				h.logger.Error("failed to delete version", zap.Error(err))
			}
		}
	}

	h.live = ver
	h.liveStaleTime = time.Now().Add(staleVersionDuration)

	if h.live != nil {
		h.live.Start()
		h.logger.Debug("set live version", zap.Stringer("version", h.live))
	}
}

func (h *Handler) updateLiveVersion(ctx context.Context) (Version, error) {
	h.liveLock.Lock()
	defer h.liveLock.Unlock()

	// We don't want to allow new Versions to be created if the Handler has
	// been stopped (the server is shutting down, most likely).
	if h.stopped {
		return nil, errors.New("stopped")
	}

	// If there's already a valid live version, then there's nothing to do.
	// This can happen if another goroutine updated the live version while we
	// were waiting for the lock.
	if live := h.getLiveVersionLocked(); live != nil {
		return live, nil
	}

	// Fetch the latest copy of the Project.
	project, err := h.m.GetProject(ctx, h.projectID)
	if err != nil {
		return nil, fmt.Errorf("getting project: %v", err)
	}

	// If the latest Project data matches the live version then the site is still
	// up-to-date and we only need to reset the stale time.
	if h.live != nil && h.live.IsProject(project) {
		h.logger.Debug("site is up-to-date", zap.Stringer("version", h.live))
		h.liveStaleTime = time.Now().Add(staleVersionDuration)
		return h.live, nil
	}

	// Create the new version and set it as live.
	ver, err := h.vp.VersionForProject(project)
	if err != nil {
		return nil, fmt.Errorf("getting site version: %v", err)
	}
	h.setLiveVersionLocked(ver)

	return ver, nil
}
