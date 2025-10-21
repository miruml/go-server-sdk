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

func TestReleaseGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Releases.Get(
		context.TODO(),
		"rls_123",
		miruserver.ReleaseGetParams{
			Expand: []string{"config_schemas"},
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

func TestReleaseListWithOptionalParams(t *testing.T) {
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
	_, err := client.Releases.List(context.TODO(), miruserver.ReleaseListParams{
		ID:      miruserver.String("rls_123"),
		Expand:  []string{"total_count"},
		Limit:   miruserver.Int(1),
		Offset:  miruserver.Int(0),
		OrderBy: miruserver.ReleaseListParamsOrderByIDAsc,
		Version: miruserver.String("v1.0.0"),
	})
	if err != nil {
		var apierr *miruserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
