package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"

	"github.com/khulnasoft/infra/packages/api/internal/auth"
	authcache "github.com/khulnasoft/infra/packages/api/internal/cache/auth"
	"github.com/khulnasoft/infra/packages/api/internal/utils"
	"github.com/khulnasoft/infra/packages/shared/pkg/telemetry"
)

func (a *APIStore) DeleteSandboxesSandboxID(
	c *gin.Context,
	sandboxID string,
) {
	ctx := c.Request.Context()
	sandboxID = utils.ShortID(sandboxID)

	teamID := c.Value(auth.TeamContextKey).(authcache.AuthTeamInfo).Team.ID

	telemetry.SetAttributes(ctx,
		attribute.String("instance.id", sandboxID),
		attribute.String("team.id", teamID.String()),
	)

	sbx, err := a.orchestrator.GetSandbox(sandboxID)
	if err != nil {
		a.sendAPIStoreError(c, http.StatusNotFound, fmt.Sprintf("Error killing sandbox - sandbox '%s' was not found", sandboxID))

		return
	}

	sbx.Logger.Debugf("Sandbox kill request received")

	if *sbx.TeamID != teamID {
		errMsg := fmt.Errorf("sandbox '%s' does not belong to team '%s'", sandboxID, teamID.String())
		telemetry.ReportCriticalError(ctx, errMsg)

		a.sendAPIStoreError(c, http.StatusUnauthorized, fmt.Sprintf("Error killing sandbox - sandbox '%s' does not belong to your team '%s'", sandboxID, teamID.String()))

		return
	}

	found := a.orchestrator.DeleteInstance(ctx, sandboxID)
	if !found {
		a.sendAPIStoreError(c, http.StatusNotFound, fmt.Sprintf("Error killing sandbox - sandbox '%s' was not found", sandboxID))

		return
	}

	c.Status(http.StatusNoContent)
}
