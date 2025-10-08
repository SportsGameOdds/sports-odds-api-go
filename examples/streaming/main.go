package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/SportsGameOdds/sports-odds-api-go"
	"github.com/SportsGameOdds/sports-odds-api-go/option"
)

func main() {
	apiKey := os.Getenv("SPORTS_ODDS_API_KEY_HEADER")
	if apiKey == "" {
		fmt.Println("Error: SPORTS_ODDS_API_KEY_HEADER environment variable not set")
		fmt.Println("Usage: export SPORTS_ODDS_API_KEY_HEADER='your-api-key-here'")
		os.Exit(1)
	}

	client := sportsoddsapi.NewClient(
		option.WithAPIKeyParam(apiKey),
		option.WithRequestTimeout(30),
		option.WithMaxRetries(2),
	)

	ctx := context.Background()

	fmt.Println("Sports Odds API Go SDK - Streaming Example")
	fmt.Println("Note: Streaming requires an AllStar plan subscription")

	streamFeed := "events:live" // Options: events:upcoming, events:byid, events:live
	fmt.Println("=== Setting up Event Stream ===")
	fmt.Printf("Feed: %s\n", streamFeed)

	events := make(map[string]sportsoddsapi.Event)

	fmt.Println("Fetching stream info and initial data...")
	streamInfo, err := client.Stream.Events(ctx, sportsoddsapi.StreamEventsParams{
		Feed: sportsoddsapi.String(streamFeed),
	})
	if err != nil {
		if apiErr, ok := err.(*sportsoddsapi.Error); ok && apiErr.StatusCode == 403 {
			fmt.Println("✗ Error: Streaming requires an AllStar plan subscription")
			fmt.Println("Visit https://sportsgameodds.com/pricing to upgrade your plan")
			os.Exit(1)
		}
		fmt.Printf("✗ API Error: %v\n", err)
		os.Exit(1)
	}

	for _, event := range streamInfo.Data {
		events[event.EventID] = event
	}

	fmt.Printf("✓ Loaded %d initial events\n", len(events))
	fmt.Println("✓ Stream configuration retrieved")

	fmt.Println("\n=== Pusher Configuration ===")
	fmt.Printf("Key: %s\nCluster: %s\nChannel: %s\nWS Host: %s\nWSS Port: %d\n",
		streamInfo.PusherKey,
		streamInfo.PusherOptions.Cluster,
		streamInfo.Channel,
		streamInfo.PusherOptions.WsHost,
		streamInfo.PusherOptions.WssPort)

	fmt.Println("Note: To establish a WebSocket connection, you'll need a Pusher client library.")
	fmt.Println("For Go, you can use: github.com/pusher/pusher-websocket-go")
	fmt.Println("Or implement a custom WebSocket client using the configuration above.")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	fmt.Println("Monitoring for updates (demo mode - no actual WebSocket connection)...")

	handleUpdate := func(changedEventIDs []string) {
		fmt.Printf("\n[%s] Received update for %d event(s)\n",
			time.Now().Format("15:04:05"), len(changedEventIDs))

		eventIDs := strings.Join(changedEventIDs, ",")
		updatedPage, err := client.Events.Get(ctx, sportsoddsapi.EventGetParams{
			EventID: sportsoddsapi.String(eventIDs),
		})
		if err != nil {
			fmt.Printf("  Error fetching updated events: %v\n", err)
			return
		}

		for _, event := range updatedPage.Data {
			events[event.EventID] = event
			fmt.Printf("  Updated: %s\n    %s @ %s\n    Status: %s\n",
				event.EventID, event.Teams.Away.Names.Short, event.Teams.Home.Names.Short,
				event.Status.DisplayShort)
		}
	}

	for {
		select {
		case <-sigChan:
			fmt.Println("\nShutting down...")
			fmt.Println("✓ Stream example completed")
			os.Exit(0)
		case <-ticker.C:
			// This is a demo; no actual WebSocket updates
			_ = handleUpdate
			fmt.Printf("[%s] Waiting for updates (demo)...\n", time.Now().Format("15:04:05"))
		}
	}
}
