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

func TestDeviceNew(t *testing.T) {
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
	_, err := client.Devices.New(context.TODO(), miruserver.DeviceNewParams{
		Name: "Robot 1",
	})
	if err != nil {
		var apierr *miruserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceGet(t *testing.T) {
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
	_, err := client.Devices.Get(context.TODO(), "dvc_123")
	if err != nil {
		var apierr *miruserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Devices.Update(
		context.TODO(),
		"dvc_123",
		miruserver.DeviceUpdateParams{
			Name: miruserver.String("Robot 1"),
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

func TestDeviceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Devices.List(context.TODO(), miruserver.DeviceListParams{
		ID:               miruserver.String("dev_123"),
		AgentVersion:     miruserver.String("v1.0.0"),
		CurrentReleaseID: miruserver.String("rls_123"),
		Expand:           []string{"total_count"},
		Limit:            miruserver.Int(1),
		Name:             miruserver.String("My Device"),
		Offset:           miruserver.Int(0),
		OrderBy:          miruserver.DeviceListParamsOrderByIDAsc,
	})
	if err != nil {
		var apierr *miruserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceDelete(t *testing.T) {
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
	_, err := client.Devices.Delete(context.TODO(), "dvc_123")
	if err != nil {
		var apierr *miruserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceNewActivationTokenWithOptionalParams(t *testing.T) {
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
	_, err := client.Devices.NewActivationToken(
		context.TODO(),
		"dvc_123",
		miruserver.DeviceNewActivationTokenParams{
			AllowReactivation: miruserver.Bool(false),
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
