package uptime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"armada-node/api"
	"armada-node/model"
)

type ContentNodeProvider struct {
	m  model.Client
	hc *http.Client
}

func NewContentNodeProvider(m model.Client, hc *http.Client) *ContentNodeProvider {
	return &ContentNodeProvider{
		m:  m,
		hc: hc,
	}
}

func (p *ContentNodeProvider) AllNodes(ctx context.Context, projectID model.ID) ([]Node, error) {
	modelNodes, err := p.m.ContentNodes(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("fetching nodes: %v", err)
	}

	var nodes []Node
	for _, node := range modelNodes {
		nodes = append(nodes, newContentNode(node, p.hc))
	}
	return nodes, nil
}

type contentNode struct {
	node *model.Node
	hc   *http.Client
}

func newContentNode(node *model.Node, hc *http.Client) *contentNode {
	return &contentNode{
		node: node,
		hc:   hc,
	}
}

func (n *contentNode) Host() string {
	return n.node.Host
}

func (n *contentNode) Uptime(ctx context.Context, resCh chan<- fetchResult, startTime, endTime time.Time) {
	var result fetchResult
	defer func() {
		resCh <- result
	}()

	url := fmt.Sprintf("https://%s/v1/uptime", n.node.Host)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		result.err = fmt.Errorf("creating http request: %v", err)
		return
	}

	q := req.URL.Query()
	q.Add("start", fmt.Sprintf("%d", startTime.Unix()))
	q.Add("end", fmt.Sprintf("%d", endTime.Unix()))
	req.URL.RawQuery = q.Encode()

	resp, err := n.hc.Do(req)
	if err != nil {
		result.err = fmt.Errorf("fetching /v1/uptime: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		result.err = fmt.Errorf("bad http status code: %d", resp.StatusCode)
		return
	}

	var uptimeResp api.UptimeResponse
	if err := json.NewDecoder(resp.Body).Decode(&uptimeResp); err != nil {
		result.err = fmt.Errorf("decoding response body: %v", err)
		return
	}

	result.resp = &uptimeResp
}
