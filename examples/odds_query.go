package main

import (
	"context"
	"fmt"
	"os"

	"github.com/SportsGameOdds/sports-odds-api-go"
	"github.com/SportsGameOdds/sports-odds-api-go/option"
)

func main() {
	// Get your API key from https://sportsgameodds.com/pricing
	apiKey := os.Getenv("SPORTS_ODDS_API_KEY_HEADER")
	if apiKey == "" {
		fmt.Println("Error: SPORTS_ODDS_API_KEY_HEADER environment variable not set")
		fmt.Println("Usage: export SPORTS_ODDS_API_KEY_HEADER='your-api-key-here'")
		os.Exit(1)
	}

	// Initialize the client
	client := sportsoddsapi.NewClient(
		option.WithAPIKeyParam(apiKey),
		option.WithRequestTimeout(30),
		option.WithMaxRetries(2),
	)

	ctx := context.Background()

	fmt.Println("Sports Odds API Go SDK - Odds Query Example\n")

	// Query for NFL events that are not finalized and have odds available
	fmt.Println("=== Querying NFL Events with Odds ===")
	fmt.Println("Filters: leagueID=NFL, finalized=false, oddsAvailable=true\n")

	page, err := client.Events.Get(ctx, sportsoddsapi.EventGetParams{
		LeagueID:      sportsoddsapi.String("NFL"),
		Finalized:     sportsoddsapi.Bool(false),
		OddsAvailable: sportsoddsapi.Bool(true),
		Limit:         sportsoddsapi.Float(10),
	})

	if err != nil {
		fmt.Printf("Error fetching events: %v\n", err)
		os.Exit(1)
	}

	if len(page.Data) == 0 {
		fmt.Println("No NFL events with odds found")
		os.Exit(0)
	}

	fmt.Printf("Found %d NFL events with odds\n\n", len(page.Data))

	// Parse all odds markets into a map
	// Structure: map[eventID]map[betTypeID][]odd
	oddsMap := make(map[string]map[string][]sportsoddsapi.EventOdd)

	for _, event := range page.Data {
		eventID := event.EventID
		oddsMap[eventID] = make(map[string][]sportsoddsapi.EventOdd)

		fmt.Printf("Event: %s\n", eventID)
		fmt.Printf("  %s @ %s\n", event.Teams.Away.Names.Short, event.Teams.Home.Names.Short)

		// Check if odds exist
		if len(event.Odds) == 0 {
			fmt.Println("  No odds markets available")
			fmt.Println()
			continue
		}

		// Group odds by betTypeID
		// Note: event.Odds is a map keyed by oddID, not a slice
		for _, odd := range event.Odds {
			betTypeID := odd.BetTypeID
			oddsMap[eventID][betTypeID] = append(oddsMap[eventID][betTypeID], odd)
		}

		// Display summary of odds markets for this event
		if len(oddsMap[eventID]) > 0 {
			fmt.Println("  Odds Markets:")
			for betTypeID, markets := range oddsMap[eventID] {
				fmt.Printf("    betTypeID %s: %d markets\n", betTypeID, len(markets))
			}
		} else {
			fmt.Println("  No odds markets available")
		}

		fmt.Println()
	}

	// Display summary
	fmt.Println("\n=== Summary ===")
	totalEvents := len(oddsMap)
	totalBetTypes := 0
	totalMarkets := 0

	for _, markets := range oddsMap {
		totalBetTypes += len(markets)
		for _, odds := range markets {
			totalMarkets += len(odds)
		}
	}

	fmt.Printf("Total events processed: %d\n", totalEvents)
	fmt.Printf("Total unique bet types: %d\n", totalBetTypes)
	fmt.Printf("Total odds markets: %d\n", totalMarkets)

	// Show example of accessing the odds map
	if len(oddsMap) > 0 {
		var firstEventID string
		for k := range oddsMap {
			firstEventID = k
			break
		}

		fmt.Printf("\nExample - Accessing odds for event %s:\n", firstEventID)
		for betTypeID, markets := range oddsMap[firstEventID] {
			fmt.Printf("  betTypeID %s: %d markets\n", betTypeID, len(markets))
			if len(markets) > 0 {
				firstMarket := markets[0]
				fmt.Printf("    Sample market: oddID=%s, betTypeID=%s\n",
					firstMarket.OddID, firstMarket.BetTypeID)
			}
		}
	}

	fmt.Println("\nOdds query example completed successfully!")
}
