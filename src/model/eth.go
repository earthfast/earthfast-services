package model

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"armada-node/contracts/nodes"
	"armada-node/contracts/projects"
	"armada-node/contracts/registry"
	"armada-node/contracts/reservations"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/zap"
)

var (
	zeroID     = [32]byte{0}
	zeroBigInt = big.NewInt(0)

	DefaultConfirmationBlocks = 1
	DefaultPageSize           = big.NewInt(1000)
)

type EthereumClient interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

type RegistryCaller interface {
	GetNodes(*bind.CallOpts) (common.Address, error)
	GetProjects(*bind.CallOpts) (common.Address, error)
	GetReservations(*bind.CallOpts) (common.Address, error)
}

type NodesCaller interface {
	GetNode(*bind.CallOpts, [32]byte) (nodes.EarthfastNode, error)
	GetNodes(*bind.CallOpts, [32]byte, *big.Int, *big.Int) ([]nodes.EarthfastNode, error)
}

type ProjectsCaller interface {
	GetProject(*bind.CallOpts, [32]byte) (projects.EarthfastProject, error)
}

type ReservationsCaller interface {
	GetReservations(*bind.CallOpts, [32]byte, *big.Int, *big.Int) ([]reservations.EarthfastNode, error)
}

type EthClient struct {
	client EthereumClient
	opts   EthClientOptions

	nodes        NodesCaller
	projects     ProjectsCaller
	reservations ReservationsCaller

	logger *zap.Logger
}

type EthClientArgs struct {
	// Endpoint is the URL to connect to for making Ethereum RPC API calls.
	Endpoint string

	// Address is the hex representation of the Registry contract address.
	Address string

	// Logger is used for logging.
	Logger *zap.Logger
}

type EthClientOptions struct {
	// ConfirmationBlocks specifies the age of the data to fetch, in blocks. If equal to 1, data
	// from the head block will be used. If not set, DefaultConfirmationBlocks will be used.
	ConfirmationBlocks int

	// PageSize specifies how many entities to request when making paginated API calls.
	PageSize *big.Int
}

func NewEthClient(ctx context.Context, args EthClientArgs, opts EthClientOptions) (*EthClient, error) {
	if args.Endpoint == "" {
		return nil, errors.New("Endpoint is required")
	}
	if args.Address == "" {
		return nil, errors.New("Address is required")
	}
	if args.Logger == nil {
		return nil, errors.New("Logger is required")
	}
	if opts.ConfirmationBlocks == 0 {
		opts.ConfirmationBlocks = DefaultConfirmationBlocks
	}
	if opts.PageSize == nil || opts.PageSize.Int64() == 0 {
		opts.PageSize = DefaultPageSize
	}
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	rpcClient, err := rpc.DialOptions(ctx, args.Endpoint, rpc.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("dialing rpc client: %v", err)
	}

	conn := ethclient.NewClient(rpcClient)

	registryCaller, err := registry.NewRegistryCaller(common.HexToAddress(args.Address), conn)
	if err != nil {
		return nil, fmt.Errorf("creating registry caller: %v", err)
	}

	callOpts := &bind.CallOpts{Context: ctx}
	nodesAddr, err := registryCaller.GetNodes(callOpts)
	if err != nil {
		return nil, fmt.Errorf("getting nodes contract address: %v", err)
	}
	projectsAddr, err := registryCaller.GetProjects(callOpts)
	if err != nil {
		return nil, fmt.Errorf("getting projects contract address: %v", err)
	}
	reservationsAddr, err := registryCaller.GetReservations(callOpts)
	if err != nil {
		return nil, fmt.Errorf("getting reservations contract address: %v", err)
	}

	nodesCaller, err := nodes.NewNodesCaller(nodesAddr, conn)
	if err != nil {
		return nil, fmt.Errorf("creating nodes caller: %v", err)
	}
	projectsCaller, err := projects.NewProjectsCaller(projectsAddr, conn)
	if err != nil {
		return nil, fmt.Errorf("creating projects caller: %v", err)
	}
	reservationsCaller, err := reservations.NewReservationsCaller(reservationsAddr, conn)
	if err != nil {
		return nil, fmt.Errorf("creating reservations caller: %v", err)
	}

	return &EthClient{
		client: conn,
		opts:   opts,

		nodes:        nodesCaller,
		projects:     projectsCaller,
		reservations: reservationsCaller,

		logger: args.Logger,
	}, nil
}

func (c *EthClient) getConfirmedBlock(ctx context.Context) (*big.Int, error) {
	// One block was already mined
	wait := c.opts.ConfirmationBlocks - 1
	if wait <= 0 {
		return nil, nil
	}

	header, err := c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("fetching block header: %v", err)
	}

	block := new(big.Int).Sub(header.Number, big.NewInt(int64(wait)))
	return block, nil
}

func (c *EthClient) GetNode(ctx context.Context, id ID) (*Node, error) {
	block, err := c.getConfirmedBlock(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching block number: %v", err)
	}
	opts := &bind.CallOpts{Context: ctx, BlockNumber: block}

	ethNode, err := c.nodes.GetNode(opts, id)
	if err != nil {
		return nil, fmt.Errorf("fetching node: %v", err)
	}
	if ethNode.Id == zeroID {
		return nil, nil
	}
	return &Node{
		ID:        ethNode.Id,
		Host:      ethNode.Host,
		ProjectID: ethNode.ProjectIds[0],
		Region:    ethNode.Region,
	}, nil
}

func (c *EthClient) GetProject(ctx context.Context, id ID) (*Project, error) {
	block, err := c.getConfirmedBlock(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching block number: %v", err)
	}
	opts := &bind.CallOpts{Context: ctx, BlockNumber: block}

	ethProject, err := c.projects.GetProject(opts, id)
	if err != nil {
		return nil, fmt.Errorf("fetching project: %v", err)
	}
	if ethProject.Id == zeroID {
		return nil, nil
	}

	c.logger.Info("Retrieved project from blockchain",
		zap.String("id", fmt.Sprintf("%x", ethProject.Id)),
		zap.String("name", ethProject.Name),
		zap.String("metadata", ethProject.Metadata))

	return &Project{
		ID:       ID(ethProject.Id),
		Name:     ethProject.Name,
		Content:  ethProject.Content,
		Checksum: fmt.Sprintf("%x", ethProject.Checksum),
		Metadata: ethProject.Metadata,
	}, nil
}

func (c *EthClient) ContentNodes(ctx context.Context, projectID ID) ([]*Node, error) {
	block, err := c.getConfirmedBlock(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching block number: %v", err)
	}
	opts := &bind.CallOpts{Context: ctx, BlockNumber: block}

	// Fetch "all" of the project's content nodes. "All" is in quotes because we're actually
	// just asking the API to return way more nodes than we expect will exist. We're doing
	// so as a shortcut to dealing with any potential inconsistencies that could arise by
	// making separate, paginated calls.
	ethNodes, err := c.reservations.GetReservations(opts, projectID, zeroBigInt, c.opts.PageSize)
	if err != nil {
		return nil, fmt.Errorf("fetching reservations: %v", err)
	}
	if n := int64(len(ethNodes)); n > c.opts.PageSize.Int64()/2 {
		c.logger.Warn(
			"The number of reserved content nodes for this project is nearing the per-API-call limit",
			zap.String("projectID", projectID.Hex()),
			zap.Int64("length", n),
			zap.Int64("limit", c.opts.PageSize.Int64()),
		)
	}

	var nodes []*Node
	for _, n := range ethNodes {
		nodes = append(nodes, &Node{
			ID:        n.Id,
			Host:      n.Host,
			ProjectID: n.ProjectIds[0],
			Region:    n.Region,
		})
	}
	return nodes, nil
}
