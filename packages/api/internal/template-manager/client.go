package template_manager

import (
	"fmt"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	khulnasoftgrpc "github.com/khulnasoft/infra/packages/shared/pkg/grpc"
	template_manager "github.com/khulnasoft/infra/packages/shared/pkg/grpc/template-manager"
)

var (
	host = os.Getenv("TEMPLATE_MANAGER_ADDRESS")
)

type GRPCClient struct {
	Client     template_manager.TemplateServiceClient
	connection khulnasoftgrpc.ClientConnInterface
}

func NewClient() (*GRPCClient, error) {
	conn, err := khulnasoftgrpc.GetConnection(host, false, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		return nil, fmt.Errorf("failed to establish GRPC connection: %w", err)
	}

	client := template_manager.NewTemplateServiceClient(conn)

	return &GRPCClient{Client: client, connection: conn}, nil
}

func (a *GRPCClient) Close() error {
	err := a.connection.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	return nil
}
