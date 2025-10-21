// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package miruserver_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/miruml/go-server-sdk"
	"github.com/miruml/go-server-sdk/internal/testutil"
	"github.com/miruml/go-server-sdk/option"
)

func TestConfigInstanceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ConfigInstances.Get(
		context.TODO(),
		"cfg_inst_123",
		miruserver.ConfigInstanceGetParams{
			Expand: []string{"content"},
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

func TestConfigInstanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.ConfigInstances.List(context.TODO(), miruserver.ConfigInstanceListParams{
		ID:             miruserver.String("cfg_inst_123"),
		ActivityStatus: miruserver.ConfigInstanceListParamsActivityStatusCreated,
		ConfigSchemaID: miruserver.String("cfg_sch_123"),
		ConfigTypeID:   miruserver.String("cfg_typ_123"),
		DeviceID:       miruserver.String("dvc_123"),
		ErrorStatus:    miruserver.ConfigInstanceListParamsErrorStatusNone,
		Expand:         []string{"total_count"},
		Limit:          miruserver.Int(1),
		Offset:         miruserver.Int(0),
		OrderBy:        miruserver.ConfigInstanceListParamsOrderByIDAsc,
		TargetStatus:   miruserver.ConfigInstanceListParamsTargetStatusCreated,
	})
	if err != nil {
		var apierr *miruserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
