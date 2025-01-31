package sandbox

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/khulnasoft/infra/packages/shared/pkg/consts"
)

const maxRetries = 120

func (s *Sandbox) syncOldEnvd(ctx context.Context) error {
	address := fmt.Sprintf("http://%s:%d/sync", s.Slot.HostIP(), consts.OldEnvdServerPort)

	var response *http.Response
	for i := 0; i < maxRetries; i++ {
		reqCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
		request, err := http.NewRequestWithContext(reqCtx, "POST", address, nil)
		if err != nil {
			cancel()
			return err
		}

		response, err = httpClient.Do(request)
		if err == nil {
			cancel()
			break
		}

		cancel()
		time.Sleep(5 * time.Millisecond)
	}

	if response == nil {
		return fmt.Errorf("failed to sync envd")
	}

	_, err := io.Copy(io.Discard, response.Body)
	if err != nil {
		return err
	}

	err = response.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

type PostInitJSONBody struct {
	EnvVars *map[string]string `json:"envVars"`
}

func (s *Sandbox) initEnvd(ctx context.Context, tracer trace.Tracer, envVars map[string]string) error {
	childCtx, childSpan := tracer.Start(ctx, "envd-init")
	defer childSpan.End()

	address := fmt.Sprintf("http://%s:%d/init", s.Slot.HostIP(), consts.DefaultEnvdServerPort)

	jsonBody := &PostInitJSONBody{
		EnvVars: &envVars,
	}

	envVarsJSON, err := json.Marshal(jsonBody)
	if err != nil {
		return err
	}

	var response *http.Response
	for i := 0; i < maxRetries; i++ {
		reqCtx, cancel := context.WithTimeout(childCtx, 50*time.Millisecond)
		request, err := http.NewRequestWithContext(reqCtx, "POST", address, bytes.NewReader(envVarsJSON))
		if err != nil {
			cancel()
			return err
		}

		response, err = httpClient.Do(request)
		if err == nil {
			cancel()
			break
		}

		cancel()
		time.Sleep(5 * time.Millisecond)
	}

	if response == nil {
		return fmt.Errorf("failed to init envd")
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	_, err = io.Copy(io.Discard, response.Body)
	if err != nil {
		return err
	}

	return nil
}
