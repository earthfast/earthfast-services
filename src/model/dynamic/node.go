package dynamic

import (
	"context"
	"fmt"
	"sync"
	"time"

	"armada-node/model"

	"go.uber.org/zap"
)

const (
	DefaultRefreshTimeout    = 10 * time.Second
	DefaultRefreshInterval   = time.Minute
	DefaultBootstrapInterval = 5 * time.Second
)

// NodeListener is a callback function that gets called when the node state changes
type NodeListener func(*model.Node)

type Options struct {
	RefreshTimeout    time.Duration
	RefreshInterval   time.Duration
	BootstrapInterval time.Duration
}

// Node is a periodically self-updating model.Node.
type Node struct {
	m    model.Client
	id   model.ID
	opts Options

	mu     *sync.RWMutex
	latest *model.Node

	stopCh chan struct{}
	stopWg *sync.WaitGroup

	logger *zap.Logger

	// Listeners for node changes
	listeners  []NodeListener
	listenerMu sync.Mutex
}

func NewNode(logger *zap.Logger, m model.Client, id model.ID, opts Options) *Node {
	if opts.RefreshTimeout == 0 {
		opts.RefreshTimeout = DefaultRefreshTimeout
	}
	if opts.RefreshInterval == 0 {
		opts.RefreshInterval = DefaultRefreshInterval
	}
	if opts.BootstrapInterval == 0 {
		opts.BootstrapInterval = DefaultBootstrapInterval
	}
	return &Node{
		m:    m,
		id:   id,
		opts: opts,

		mu: &sync.RWMutex{},

		stopCh: make(chan struct{}),
		stopWg: &sync.WaitGroup{},

		logger:    logger,
		listeners: make([]NodeListener, 0),
	}
}

// Start instructs Node to start polling the model for updates.
func (n *Node) Start() {
	n.stopWg.Add(1)
	go func() {
		defer n.stopWg.Done()
		n.run()
	}()
}

// Stop ceases polling the model.
func (n *Node) Stop() {
	close(n.stopCh)
	n.stopWg.Wait()
}

// AddListener registers a function to be called when the node state changes
func (n *Node) AddListener(listener NodeListener) {
	n.listenerMu.Lock()
	n.listeners = append(n.listeners, listener)
	n.listenerMu.Unlock()

	// Call immediately with current state if available
	n.mu.RLock()
	currentNode := n.latest
	n.mu.RUnlock()

	if currentNode != nil {
		listener(currentNode)
	}
}

// Get returns the current node state
func (n *Node) Get() *model.Node {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.latest
}

// Bootstrap attempts to fetch the latest copy of the Node data from the model,
// retrying until either it succeeds or the context is cancelled.
func (n *Node) Bootstrap(ctx context.Context) error {
	parentCtx := ctx
	for {
		ctx, cancel := context.WithTimeout(parentCtx, n.opts.RefreshTimeout)
		err := n.refresh(ctx)
		cancel()
		if err == nil {
			return nil
		}
		n.logger.Error("Bootstrap failed, will retry...", zap.Error(err))

		select {
		case <-parentCtx.Done():
			return parentCtx.Err()
		case <-time.After(n.opts.BootstrapInterval):
		}
	}
}

func (n *Node) ID() model.ID {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.latest.ID
}

func (n *Node) ProjectID() model.ID {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.latest.ProjectID
}

func (n *Node) Host() string {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.latest.Host
}

func (n *Node) run() {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), n.opts.RefreshTimeout)
		if err := n.refresh(ctx); err != nil {
			n.logger.Error("Failed to refresh Node", zap.Error(err))
		}
		cancel()

		select {
		case <-n.stopCh:
			return
		case <-time.After(n.opts.RefreshInterval):
		}
	}
}

func (n *Node) refresh(ctx context.Context) error {
	node, err := n.m.GetNode(ctx, n.id)
	if err != nil {
		return fmt.Errorf("fetching node :%v", err)
	}
	if node == nil {
		return fmt.Errorf("node not found: id=%s", n.id.Hex())
	}

	n.mu.RLock()
	oldNode := n.latest
	dirty := oldNode == nil || !oldNode.Equal(node)
	n.mu.RUnlock()

	if !dirty {
		return nil
	}

	n.mu.Lock()
	n.latest = node
	n.mu.Unlock()

	n.logger.Debug("Updated dynamic.Node",
		zap.String("nodeID", n.latest.ID.Hex()),
		zap.String("host", n.latest.Host),
		zap.String("projectID", n.latest.ProjectID.Hex()),
	)

	// Notify listeners of the change
	if dirty {
		n.notifyListeners(node)
	}

	return nil
}

// notifyListeners calls all registered listeners with the current node state
func (n *Node) notifyListeners(node *model.Node) {
	n.listenerMu.Lock()
	listeners := make([]NodeListener, len(n.listeners))
	copy(listeners, n.listeners)
	n.listenerMu.Unlock()

	for _, listener := range listeners {
		listener(node)
	}
}
