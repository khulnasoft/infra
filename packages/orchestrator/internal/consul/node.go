package consul

import (
	"github.com/khulnasoft/infra/packages/shared/pkg/utils"
)

const shortNodeIDLength = 8

var (
	nodeID = utils.RequiredEnv("NODE_ID", "Nomad ID of the instance node")
	// Node ID must be at least 8 characters long.
	ClientID = nodeID[:shortNodeIDLength]
)
