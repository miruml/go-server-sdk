// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package miruserver_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/miru-server-go"
	"github.com/stainless-sdks/miru-server-go/internal/testutil"
	"github.com/stainless-sdks/miru-server-go/option"
)

func TestDeploymentNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := miruserver.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Deployments.New(context.TODO(), miruserver.DeploymentNewParams{
		Description: "Deployment for the motion control config instance",
		DeviceID:    "dvc_123",
		NewConfigInstances: []miruserver.DeploymentNewParamsNewConfigInstance{{
			ConfigSchemaID: "cfg_sch_123",
			Content: map[string]interface{}{
				"direction": "forward",
				"speed":     100,
				"duration":  10,
			},
			RelativeFilepath: "/v1/motion-control.json",
		}},
		ReleaseID:     "rls_123",
		TargetStatus:  miruserver.DeploymentNewParamsTargetStatusStaged,
		Expand:        []string{"device"},
		PatchSourceID: miruserver.String("dpl_123"),
	})
	if err != nil {
		var apierr *miruserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeploymentGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := miruserver.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Deployments.Get(
		context.TODO(),
		"dpl_123",
		miruserver.DeploymentGetParams{
			Expand: []string{"device"},
		},
	)
	if err != nil {
		var apierr *miruserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeploymentListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := miruserver.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Deployments.List(context.TODO(), miruserver.DeploymentListParams{
		ID:             miruserver.String("dpl_123"),
		ActivityStatus: miruserver.DeploymentListParamsActivityStatusValidating,
		DeviceID:       miruserver.String("dvc_123"),
		ErrorStatus:    miruserver.DeploymentListParamsErrorStatusNone,
		Expand:         []string{"total_count"},
		Limit:          miruserver.Int(1),
		Offset:         miruserver.Int(0),
		OrderBy:        miruserver.DeploymentListParamsOrderByIDAsc,
		ReleaseID:      miruserver.String("rls_123"),
		TargetStatus:   miruserver.DeploymentListParamsTargetStatusStaged,
	})
	if err != nil {
		var apierr *miruserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeploymentValidate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := miruserver.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Deployments.Validate(
		context.TODO(),
		"dpl_123",
		miruserver.DeploymentValidateParams{
			ConfigInstances: []miruserver.DeploymentValidateParamsConfigInstance{{
				ID:      "cfg_inst_123",
				Message: "The config instance contains additional parameters that are not supported by the config schema.",
				Parameters: []miruserver.DeploymentValidateParamsConfigInstanceParameter{{
					Message: "The value must be between 10 and 60.",
					Path:    []string{"content", "telemetry", "upload_interval_sec"},
				}},
			}},
			IsValid: false,
			Message: "Validation failed!",
		},
	)
	if err != nil {
		var apierr *miruserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
