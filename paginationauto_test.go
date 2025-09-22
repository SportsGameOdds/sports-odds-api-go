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

func TestAutoPagination(t *testing.T) {
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
	iter := client.Events.GetAutoPaging(context.TODO(), sportsoddsapi.EventGetParams{
		Limit: sportsoddsapi.Float(30),
	})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		event := iter.Current()
		t.Logf("%+v\n", event.Activity)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
