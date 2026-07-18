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

func TestMarketGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIKeyParam("My API Key Param"),
	)
	_, err := client.Markets.Get(context.TODO(), sportsoddsapi.MarketGetParams{
		BetTypeID:    sportsoddsapi.String("betTypeID"),
		BookmakerID:  sportsoddsapi.String("bookmakerID"),
		Cursor:       sportsoddsapi.String("cursor"),
		IsMainMarket: sportsoddsapi.Bool(true),
		IsProp:       sportsoddsapi.Bool(true),
		IsSubPeriod:  sportsoddsapi.Bool(true),
		IsSupported:  sportsoddsapi.Bool(true),
		LeagueID:     sportsoddsapi.String("leagueID"),
		Limit:        sportsoddsapi.Float(0),
		OddID:        sportsoddsapi.String("oddID"),
		PeriodID:     sportsoddsapi.String("periodID"),
		PropType:     sportsoddsapi.String("propType"),
		SideID:       sportsoddsapi.String("sideID"),
		SportID:      sportsoddsapi.String("sportID"),
		StatEntityID: sportsoddsapi.String("statEntityID"),
		StatID:       sportsoddsapi.String("statID"),
	})
	if err != nil {
		var apierr *sportsoddsapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
