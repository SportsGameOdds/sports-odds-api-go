// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sportsoddsapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/sports-odds-api-go"
	"github.com/stainless-sdks/sports-odds-api-go/internal/testutil"
	"github.com/stainless-sdks/sports-odds-api-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sportsoddsapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKeyHeader("My API Key Header"),
	)
	page, err := client.Events.Get(context.TODO(), sportsoddsapi.EventGetParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", page)
}
