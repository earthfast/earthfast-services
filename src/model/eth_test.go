package model

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"armada-node/contracts/nodes"
	"armada-node/contracts/projects"
	"armada-node/contracts/reservations"
	"armada-node/geo"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap/zaptest"
)

type fakeClient struct{}

func newFakeClient() *fakeClient {
	return &fakeClient{}
}

func (c *fakeClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return &types.Header{Number: big.NewInt(100)}, nil
}

type fakeContract struct {
	nodes             []nodes.EarthfastNode
	reservedNodes     []reservations.EarthfastNode
	projects          []projects.EarthfastProject
	callOptsValidator func(*bind.CallOpts)
}

func newFakeContract(nodes []nodes.EarthfastNode, reservedNodes []reservations.EarthfastNode, projects []projects.EarthfastProject) *fakeContract {
	return &fakeContract{
		nodes:             nodes,
		reservedNodes:     reservedNodes,
		projects:          projects,
		callOptsValidator: func(*bind.CallOpts) {},
	}
}

func (c *fakeContract) GetNode(opts *bind.CallOpts, id [32]byte) (nodes.EarthfastNode, error) {
	c.callOptsValidator(opts)

	for _, n := range c.nodes {
		if n.Id == id {
			return n, nil
		}
	}

	return nodes.EarthfastNode{}, nil
}

func (c *fakeContract) GetNodes(opts *bind.CallOpts, operatorIdOrZero [32]byte, skipBig *big.Int, sizeBig *big.Int) ([]nodes.EarthfastNode, error) {
	c.callOptsValidator(opts)

	var nodes []nodes.EarthfastNode
	for _, n := range c.nodes {
		nodes = append(nodes, n)
	}

	skip := int(skipBig.Int64())
	nodes = nodes[min(skip, len(nodes)):]

	size := int(sizeBig.Int64())
	nodes = nodes[:min(size, len(nodes))]

	return nodes, nil
}

func (c *fakeContract) GetProject(opts *bind.CallOpts, id [32]byte) (projects.EarthfastProject, error) {
	c.callOptsValidator(opts)
	for _, p := range c.projects {
		if p.Id == id {
			return p, nil
		}
	}
	return projects.EarthfastProject{}, nil
}

func (c *fakeContract) GetReservations(opts *bind.CallOpts, id [32]byte, skipBig *big.Int, sizeBig *big.Int) ([]reservations.EarthfastNode, error) {
	c.callOptsValidator(opts)

	var nodes []reservations.EarthfastNode
	for _, n := range c.reservedNodes {
		if n.ProjectIds[0] == id {
			nodes = append(nodes, n)
		}
	}

	skip := int(skipBig.Int64())
	nodes = nodes[min(skip, len(nodes)):]

	size := int(sizeBig.Int64())
	nodes = nodes[:min(size, len(nodes))]

	return nodes, nil
}

func newEthClient(t *testing.T, contract *fakeContract, opts EthClientOptions) *EthClient {
	if opts.ConfirmationBlocks == 0 {
		opts.ConfirmationBlocks = DefaultConfirmationBlocks
	}
	if opts.PageSize == nil || opts.PageSize.Int64() == 0 {
		opts.PageSize = DefaultPageSize
	}
	return &EthClient{
		client: newFakeClient(),
		opts:   opts,

		nodes:        contract,
		projects:     contract,
		reservations: contract,

		logger: zaptest.NewLogger(t),
	}
}

func TestEthClient_GetNode(t *testing.T) {
	content0 := nodes.EarthfastNode{
		Id:     randomID(t),
		Host:   "content0.armadanetwork.com",
		Region: geo.NorthAmerica.ID,
	}
	content1 := nodes.EarthfastNode{
		Id:     randomID(t),
		Host:   "content1.armadanetwork.com",
		Region: geo.SouthAmerica.ID,
	}

	cases := []struct {
		name string
		id   ID
		want *Node
	}{
		{
			name: "Not found",
			id:   randomID(t),
			want: nil,
		},
		{
			name: "Success",
			id:   content0.Id,
			want: &Node{ID: content0.Id, Host: content0.Host, Region: geo.NorthAmerica.ID},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			contract := newFakeContract(
				[]nodes.EarthfastNode{content0, content1},
				nil,
				nil,
			)
			c := newEthClient(t, contract, EthClientOptions{})

			got, err := c.GetNode(ctx, tc.id)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("GetNode(%x) mismatch (-want +got):\n%s", tc.id, diff)
			}
		})
	}
}

func TestEthClient_GetProject(t *testing.T) {
	project0 := projects.EarthfastProject{
		Id:       randomID(t),
		Name:     "hello-world",
		Content:  "http://example.com/site.tar.gz",
		Checksum: [32]byte{1, 2},
	}

	cases := []struct {
		name string
		id   ID
		want *Project
	}{
		{
			name: "Not found",
			id:   randomID(t),
			want: nil,
		},
		{
			name: "Success",
			id:   project0.Id,
			want: &Project{
				ID:       project0.Id,
				Name:     project0.Name,
				Content:  project0.Content,
				Checksum: fmt.Sprintf("%x", project0.Checksum),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			contract := newFakeContract(nil, nil, []projects.EarthfastProject{project0})
			c := newEthClient(t, contract, EthClientOptions{})

			got, err := c.GetProject(ctx, tc.id)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("GetProject(%x) mismatch (-want +got):\n%s", tc.id, diff)
			}
		})
	}
}

func TestEthClient_ContentNodes(t *testing.T) {
	project0 := projects.EarthfastProject{
		Id:       randomID(t),
		Name:     "hello-world",
		Content:  "http://example.com/site.tar.gz",
		Checksum: [32]byte{1, 2},
	}

	content0 := reservations.EarthfastNode{
		Id:         randomID(t),
		Host:       "content0.armadanetwork.com",
		ProjectIds: [2][32]byte{project0.Id, zeroID},
		Region:     geo.NorthAmerica.ID,
	}
	content1 := reservations.EarthfastNode{
		Id:     randomID(t),
		Host:   "content1.armadanetwork.com",
		Region: geo.SouthAmerica.ID,
	}
	content2 := reservations.EarthfastNode{
		Id:         randomID(t),
		Host:       "content2.armadanetwork.com",
		ProjectIds: [2][32]byte{project0.Id, zeroID},
		Region:     geo.Asia.ID,
	}

	cases := []struct {
		name string
		id   ID
		want []*Node
	}{
		{
			name: "Not found",
			id:   randomID(t),
			want: nil,
		},
		{
			name: "Success",
			id:   project0.Id,
			want: []*Node{
				&Node{ID: content0.Id, Host: content0.Host, ProjectID: project0.Id, Region: geo.NorthAmerica.ID},
				&Node{ID: content2.Id, Host: content2.Host, ProjectID: project0.Id, Region: geo.Asia.ID},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			contract := newFakeContract(
				nil,
				[]reservations.EarthfastNode{content0, content1, content2},
				nil,
			)
			c := newEthClient(t, contract, EthClientOptions{})

			got, err := c.ContentNodes(ctx, tc.id)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("ContentNodes(%x) mismatch (-want +got):\n%s", tc.id, diff)
			}
		})
	}
}

func TestEthClient_Confirmations(t *testing.T) {
	cases := []struct {
		name               string
		blockConfirmations int
		wantBlockNumber    *big.Int
	}{
		{
			name:               "latest block",
			blockConfirmations: 1,
			wantBlockNumber:    nil,
		},
		{
			name:               "older block",
			blockConfirmations: 2,
			wantBlockNumber:    big.NewInt(99),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			contract := newFakeContract(nil, nil, nil)
			contract.callOptsValidator = func(opts *bind.CallOpts) {
				if diff := cmp.Diff(tc.wantBlockNumber.String(), opts.BlockNumber.String()); diff != "" {
					t.Errorf("BlockNumber mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
