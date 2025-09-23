// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sportsoddsapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/SportsGameOdds/sports-odds-api-go"
	"github.com/SportsGameOdds/sports-odds-api-go/internal/testutil"
	"github.com/SportsGameOdds/sports-odds-api-go/option"
)

func TestManualPagination(t *testing.T) {
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
	page, err := client.Events.Get(context.TODO(), sportsoddsapi.EventGetParams{
		Limit: sportsoddsapi.Float(30),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, event := range page.Data {
		t.Logf("%+v\n", event.Activity)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, event := range page.Data {
			t.Logf("%+v\n", event.Activity)
		}
	}
}
