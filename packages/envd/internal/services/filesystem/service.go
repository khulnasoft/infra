package filesystem

import (
	"connectrpc.com/connect"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"

	"github.com/khulnasoft/infra/packages/envd/internal/logs"
	spec "github.com/khulnasoft/infra/packages/envd/internal/services/spec/filesystem/filesystemconnect"
	"github.com/khulnasoft/infra/packages/envd/internal/utils"
)

type Service struct {
	logger   *zerolog.Logger
	watchers *utils.Map[string, *FileWatcher]
}

func Handle(server *chi.Mux, l *zerolog.Logger) {
	service := Service{
		logger:   l,
		watchers: utils.NewMap[string, *FileWatcher](),
	}

	interceptors := connect.WithInterceptors(logs.NewUnaryLogInterceptor(l))

	path, handler := spec.NewFilesystemHandler(service, interceptors)

	server.Mount(path, handler)
}
