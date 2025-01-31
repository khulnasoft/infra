package orchestrator

import (
	"context"
	"fmt"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc"

	"github.com/khulnasoft/infra/packages/api/internal/api"
	"github.com/khulnasoft/infra/packages/api/internal/node"
	khulnasoftgrpc "github.com/khulnasoft/infra/packages/shared/pkg/grpc"
	"github.com/khulnasoft/infra/packages/shared/pkg/grpc/orchestrator"
	"github.com/khulnasoft/infra/packages/shared/pkg/smap"
)

type GRPCClient struct {
	Sandbox    orchestrator.SandboxServiceClient
	connection khulnasoftgrpc.ClientConnInterface
}

func NewClient(host string) (*GRPCClient, error) {
	conn, err := khulnasoftgrpc.GetConnection(host, false, grpc.WithStatsHandler(otelgrpc.NewClientHandler()), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		return nil, fmt.Errorf("failed to establish GRPC connection: %w", err)
	}

	client := orchestrator.NewSandboxServiceClient(conn)

	return &GRPCClient{Sandbox: client, connection: conn}, nil
}

func (a *GRPCClient) Close() error {
	err := a.connection.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	return nil
}

func (o *Orchestrator) connectToNode(ctx context.Context, node *node.NodeInfo) error {
	ctx, childSpan := o.tracer.Start(ctx, "connect-to-node")
	childSpan.SetAttributes(attribute.String("node.id", node.ID))

	defer childSpan.End()

	client, err := NewClient(node.OrchestratorAddress)
	if err != nil {
		return err
	}

	buildCache := ttlcache.New[string, interface{}]()
	go buildCache.Start()

	n := &Node{
		Client:         client,
		buildCache:     buildCache,
		sbxsInProgress: smap.New[*sbxInProgress](),
		status:         api.NodeStatusReady,
		Info:           node,
	}

	o.nodes.Insert(n.Info.ID, n)

	return nil
}

func (o *Orchestrator) GetClient(nodeID string) (*GRPCClient, error) {
	n := o.GetNode(nodeID)
	if n == nil {
		return nil, fmt.Errorf("node '%s' not found", nodeID)
	}

	return n.Client, nil
}
