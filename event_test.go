// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sportsoddsapi_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/SportsGameOdds/sports-odds-api-go"
	"github.com/SportsGameOdds/sports-odds-api-go/internal/testutil"
	"github.com/SportsGameOdds/sports-odds-api-go/option"
)

func TestEventGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Events.Get(context.TODO(), sportsoddsapi.EventGetParams{
		BookmakerID:         sportsoddsapi.String("bookmakerID"),
		Cancelled:           sportsoddsapi.Bool(true),
		Cursor:              sportsoddsapi.String("cursor"),
		Ended:               sportsoddsapi.Bool(true),
		EventID:             sportsoddsapi.String("eventID"),
		EventIDs:            sportsoddsapi.String("eventIDs"),
		Finalized:           sportsoddsapi.Bool(true),
		IncludeAltLines:     sportsoddsapi.Bool(true),
		IncludeOpposingOdds: sportsoddsapi.Bool(true),
		LeagueID:            sportsoddsapi.String("leagueID"),
		Limit:               sportsoddsapi.Float(0),
		Live:                sportsoddsapi.Bool(true),
		OddID:               sportsoddsapi.String("oddID"),
		OddsAvailable:       sportsoddsapi.Bool(true),
		OddsPresent:         sportsoddsapi.Bool(true),
		PlayerID:            sportsoddsapi.String("playerID"),
		SportID:             sportsoddsapi.String("sportID"),
		Started:             sportsoddsapi.Bool(true),
		StartsAfter:         sportsoddsapi.Time(time.Now()),
		StartsBefore:        sportsoddsapi.Time(time.Now()),
		TeamID:              sportsoddsapi.String("teamID"),
		Type:                sportsoddsapi.String("type"),
	})
	if err != nil {
		var apierr *sportsoddsapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
