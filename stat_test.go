// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sportsoddsapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/SportsGameOdds/sports-odds-api-go"
	"github.com/SportsGameOdds/sports-odds-api-go/internal/testutil"
	"github.com/SportsGameOdds/sports-odds-api-go/option"
)

func TestStatGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Stats.Get(context.TODO(), sportsoddsapi.StatGetParams{
		SportID:   sportsoddsapi.String("sportID"),
		StatID:    sportsoddsapi.String("statID"),
		StatLevel: sportsoddsapi.String("statLevel"),
	})
	if err != nil {
		var apierr *sportsoddsapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
