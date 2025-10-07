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

	fmt.Println("Sports Odds API Go SDK - Basic Usage Examples\n")

	// Example 1: Fetch recent events
	fmt.Println("=== Fetching Events ===")
	page, err := client.Events.Get(ctx, sportsoddsapi.EventGetParams{
		Limit: sportsoddsapi.Float(10),
	})

	if err != nil {
		fmt.Printf("Error fetching events: %v\n", err)
		os.Exit(1)
	}

	if len(page.Data) == 0 {
		fmt.Println("No events found")
	} else {
		fmt.Printf("Found %d events:\n", len(page.Data))
		for i := 0; i < 3 && i < len(page.Data); i++ {
			event := page.Data[i]
			fmt.Printf("  - %s: %s @ %s\n",
				event.EventID,
				event.Teams.Away.Names.Short,
				event.Teams.Home.Names.Short)
		}
	}

	// Example 2: Auto-pagination
	fmt.Println("\n=== Auto-Pagination Example ===")
	iter := client.Events.GetAutoPaging(ctx, sportsoddsapi.EventGetParams{
		Limit: sportsoddsapi.Float(5),
	})

	count := 0
	maxEvents := 15 // Limit for demo purposes

	for iter.Next() {
		event := iter.Current()
		count++
		if count <= 10 {
			fmt.Printf("  Event %d: %s\n", count, event.EventID)
		}
		if count >= maxEvents {
			break
		}
	}

	if err := iter.Err(); err != nil {
		fmt.Printf("Error during auto-pagination: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Processed %d events across multiple pages\n", count)

	// Example 3: Error handling
	fmt.Println("\n=== Error Handling Example ===")
	_, err = client.Events.Get(ctx, sportsoddsapi.EventGetParams{
		EventID: sportsoddsapi.String("invalid-id"),
	})

	if err != nil {
		// Check if it's an API error
		if apiErr, ok := err.(*sportsoddsapi.Error); ok {
			fmt.Printf("Caught API Error: Status %d\n", apiErr.StatusCode)
		} else {
			fmt.Printf("Caught Error: %v\n", err)
		}
	}

	fmt.Println("\nExamples completed successfully!")
}
