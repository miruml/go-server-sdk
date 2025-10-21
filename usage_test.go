// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package miruserver_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/miru-server-go"
	"github.com/stainless-sdks/miru-server-go/internal/testutil"
	"github.com/stainless-sdks/miru-server-go/option"
)

func TestUsage(t *testing.T) {
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
	configInstance, err := client.ConfigInstances.Get(
		context.TODO(),
		"cfg_inst_123",
		miruserver.ConfigInstanceGetParams{},
	)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", configInstance.ID)
}
