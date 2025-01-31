package server

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/khulnasoft/infra/packages/orchestrator/internal/dns"
	"github.com/khulnasoft/infra/packages/orchestrator/internal/sandbox"
	"github.com/khulnasoft/infra/packages/orchestrator/internal/sandbox/network"
	"github.com/khulnasoft/infra/packages/orchestrator/internal/sandbox/template"
	khulnasoftgrpc "github.com/khulnasoft/infra/packages/shared/pkg/grpc"
	"github.com/khulnasoft/infra/packages/shared/pkg/grpc/orchestrator"
	"github.com/khulnasoft/infra/packages/shared/pkg/smap"
)

const ServiceName = "orchestrator"

type server struct {
	orchestrator.UnimplementedSandboxServiceServer
	sandboxes     *smap.Map[*sandbox.Sandbox]
	dns           *dns.DNS
	tracer        trace.Tracer
	networkPool   *network.Pool
	templateCache *template.Cache

	pauseMu sync.Mutex
}

func New() (*grpc.Server, error) {
	ctx := context.Background()

	dnsServer := dns.New()
	go func() {
		log.Printf("Starting DNS server")

		err := dnsServer.Start("127.0.0.4", 53)
		if err != nil {
			log.Fatalf("Failed running DNS server: %s\n", err.Error())
		}
	}()

	templateCache, err := template.NewCache(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create template cache: %w", err)
	}

	networkPool, err := network.NewPool(ctx, network.NewSlotsPoolSize, network.ReusedSlotsPoolSize)
	if err != nil {
		return nil, fmt.Errorf("failed to create network pool: %w", err)
	}

	s := grpc.NewServer(
		grpc.StatsHandler(khulnasoftgrpc.NewStatsWrapper(otelgrpc.NewServerHandler())),
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(),
		),
	)

	orchestrator.RegisterSandboxServiceServer(s, &server{
		tracer:        otel.Tracer(ServiceName),
		dns:           dnsServer,
		sandboxes:     smap.New[*sandbox.Sandbox](),
		networkPool:   networkPool,
		templateCache: templateCache,
	})

	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	return s, nil
}
