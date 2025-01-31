package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"

	"github.com/khulnasoft/infra/packages/shared/pkg/utils"
)

var (
	consulToken = utils.RequiredEnv("CONSUL_TOKEN", "Consul token for authenticating requests to the Consul API")
	Client      = utils.Must(newClient())
)

func newClient() (*api.Client, error) {
	config := api.DefaultConfig()
	config.Token = consulToken

	consulClient, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Consul client: %w", err)
	}

	return consulClient, nil
}
