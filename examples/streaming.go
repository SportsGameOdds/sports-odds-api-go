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
	// Get your API key from https://sportsgameodds.com/pricing
	// Note: Streaming requires an AllStar plan subscription
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

	fmt.Println("Sports Odds API Go SDK - Streaming Example")
	fmt.Println("Note: Streaming requires an AllStar plan subscription\n")

	streamFeed := "events:live" // Options: events:upcoming, events:byid, events:live

	fmt.Println("=== Setting up Event Stream ===")
	fmt.Printf("Feed: %s\n\n", streamFeed)

	// Initialize a data structure where we'll save the event data
	events := make(map[string]sportsoddsapi.Event)

	// Call this endpoint to get initial data and connection parameters
	fmt.Println("Fetching stream info and initial data...")
	streamInfo, err := client.Stream.Events(ctx, sportsoddsapi.StreamEventsParams{
		Feed: sportsoddsapi.String(streamFeed),
	})

	if err != nil {
		// Check for PermissionDeniedError
		if apiErr, ok := err.(*sportsoddsapi.Error); ok {
			if apiErr.StatusCode == 403 {
				fmt.Println("✗ Error: Streaming requires an AllStar plan subscription")
				fmt.Println("Visit https://sportsgameodds.com/pricing to upgrade your plan")
				os.Exit(1)
			}
			fmt.Printf("✗ API Error: Status %d\n", apiErr.StatusCode)
		} else {
			fmt.Printf("✗ Error: %v\n", err)
		}
		os.Exit(1)
	}

	// Seed initial data
	for _, event := range streamInfo.Data {
		events[event.EventID] = event
	}

	fmt.Printf("✓ Loaded %d initial events\n", len(events))
	fmt.Println("✓ Stream configuration retrieved")

	// Print Pusher configuration for WebSocket connection
	fmt.Println("\n=== Pusher Configuration ===")
	fmt.Printf("Key: %s\n", streamInfo.PusherKey)
	fmt.Printf("Cluster: %s\n", streamInfo.PusherOptions.Cluster)
	fmt.Printf("Channel: %s\n", streamInfo.Channel)
	fmt.Printf("WS Host: %s\n", streamInfo.PusherOptions.WsHost)
	fmt.Printf("WSS Port: %d\n", streamInfo.PusherOptions.WssPort)

	fmt.Println("\nNote: To establish a WebSocket connection, you'll need a Pusher client library.")
	fmt.Println("For Go, you can use: github.com/pusher/pusher-websocket-go")
	fmt.Println("Or implement a custom WebSocket client using the configuration above.")
	fmt.Println("\nPress Ctrl+C to stop\n")

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	// Simulate receiving updates (in a real implementation, you'd connect to the WebSocket)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	fmt.Println("Note: This example shows the structure for streaming.")
	fmt.Println("For full WebSocket support, integrate a Pusher WebSocket client library.\n")

	// Example function showing how to handle updates when they arrive via WebSocket
	handleUpdate := func(changedEventIDs []string) {
		fmt.Printf("\n[%s] Received update for %d event(s)\n",
			time.Now().Format("15:04:05"), len(changedEventIDs))

		// Get the eventIDs that changed
		eventIDs := strings.Join(changedEventIDs, ",")

		// Get the full event data for the changed events
		updatedPage, err := client.Events.Get(ctx, sportsoddsapi.EventGetParams{
			EventID: sportsoddsapi.String(eventIDs),
		})

		if err != nil {
			fmt.Printf("  Error fetching updated events: %v\n", err)
			return
		}

		for _, event := range updatedPage.Data {
			// Update our data with the full event data
			events[event.EventID] = event

			fmt.Printf("  Updated: %s\n", event.EventID)
			fmt.Printf("    %s @ %s\n",
				event.Teams.Away.Names.Short,
				event.Teams.Home.Names.Short)
			fmt.Printf("    Status: %s\n", event.Status.DisplayShort)
		}
	}

	fmt.Println("\n=== How to Use This Function ===")

	fmt.Println("When you receive a 'data' event from the WebSocket, the payload will contain")
	fmt.Println("an array of changed events. Extract the eventIDs and call handleUpdate().")
	fmt.Println("\nExample WebSocket event handler:")
	fmt.Println("  channel.Bind(\"data\", func(data []map[string]string) {")
	fmt.Println("    eventIDs := []string{}")
	fmt.Println("    for _, evt := range data {")
	fmt.Println("      eventIDs = append(eventIDs, evt[\"eventID\"])")
	fmt.Println("    }")
	fmt.Println("    handleUpdate(eventIDs)")
	fmt.Println("  })")

	// Keep the program running to demonstrate structure
	fmt.Println("\nMonitoring for updates (demo mode - no actual WebSocket connection)...")
	for {
		select {
		case <-sigChan:
			fmt.Println("\n\nShutting down...")
			fmt.Println("✓ Stream example completed")
			os.Exit(0)
		case <-ticker.C:
			// In a real implementation, this is where you would receive
			// WebSocket events and call handleUpdate with the changed event IDs
			_ = handleUpdate
			fmt.Printf("[%s] Waiting for updates (demo)...\n", time.Now().Format("15:04:05"))
		}
	}
}
