# Sports Odds API - Live Sports Data & Sportsbook Betting Odds - Powered by SportsGameOdds Go API Library

Get live betting odds, spreads, and totals for NFL, NBA, MLB, and 50 additional sports and leagues. Production-ready Go SDK with WebSocket support, 99.9% uptime, and sub-minute updates during live games. Perfect for developers building sportsbook platforms, odds comparison tools, positive EV models, and anything else that requires fast, accurate sports data.

<a href="https://pkg.go.dev/github.com/SportsGameOdds/sports-odds-api-go"><img src="https://pkg.go.dev/badge/github.com/SportsGameOdds/sports-odds-api-go.svg" alt="Go Reference"></a>

This library provides convenient access to the Sports Game Odds REST API from applications written in Go.

The REST API documentation can be found on [sportsgameodds.com](https://sportsgameodds.com/docs/). The full API of this library can be found in [api.md](api.md).

It is generated with [Stainless](https://www.stainless.com/).

## Features

**For developers building the next generation of sports stats and/or betting applications:**

- 📈 **3k+ odds markets** including moneylines, spreads, over/unders, team props, player props & more
- 🏈 **50+ leagues covered** including NFL, NBA, MLB, NHL, NCAAF, NCAAB, EPL, UCL, UFC, PGA, ATP & more
- 📊 **80+ sportsbooks** with unified odds formats, alt lines & deeplinks
- 📺 **Live scores & stats** coverage on all games, teams, and players
- ⚡ **Sub-100ms response times** and sub-minute updates for fast data
- 🔧 **Typed requests & responses** leveraging Go structs and JSON tags
- 💰 **Developer-friendly pricing** with a generous free tier
- ⏱️ **5-minute setup** with copy-paste examples

## Installation

```go
import (
    "github.com/SportsGameOdds/sports-odds-api-go" // imported as sportsoddsapi
)
```

Or to pin the version:

```sh
go get -u 'github.com/SportsGameOdds/sports-odds-api-go@v0.0.1'
```

## Obtain an API Key

Get a free API key from [sportsgameodds.com](https://sportsgameodds.com/pricing).

Unlike enterprise-only solutions, the Sports Game Odds API offers a developer-friendly experience, transparent pricing, comprehensive documentation, and a generous free tier.

## Requirements

This library requires Go 1.22+.

## Usage

The full API of this library can be found in [api.md](api.md).

```go
package main

import (
    "context"
    "fmt"

    "github.com/SportsGameOdds/sports-odds-api-go"
    "github.com/SportsGameOdds/sports-odds-api-go/option"
)

func main() {
    client := sportsoddsapi.NewClient(
        option.WithAPIKeyParam("My API Key Param"), // defaults to os.LookupEnv("SPORTS_ODDS_API_KEY_HEADER")
    )
    page, err := client.Events.Get(context.TODO(), sportsoddsapi.EventGetParams{})
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("%+v\n", page)
}
```

# Real-Time Event Streaming API

This API endpoint is only available to **AllStar** and **custom plan** subscribers. It is not included with basic subscription tiers. [Contact support](mailto:api@sportsgameodds.com) to get access.

This streaming API is currently in **beta**. API call patterns, response formats, and functionality may change. Fully managed streaming via SDK may be available in future releases.

Our Streaming API provides real-time updates for Event objects through WebSocket connections. Instead of polling our REST endpoints, you can maintain a persistent connection to receive instant notifications when events change. This is ideal for applications that need immediate updates with minimal delay.

We use [Pusher Protocol](https://pusher.com/docs/channels/library_auth_reference/pusher-websockets-protocol/) for WebSocket communication. While you can connect using any WebSocket library, we strongly recommend using any [Pusher Client Library](https://pusher.com/docs/channels/library_auth_reference/pusher-client-libraries) (ex: [Go](https://github.com/pusher/pusher-http-go))

## How It Works

The streaming process involves two steps:

1. **Get Connection Details**: Make a request using `client.Stream.Events()` to receive:
    - WebSocket authentication credentials
    - WebSocket URL/channel info
    - Initial snapshot of current data

2. **Connect and Stream**: Use the provided details to connect via Pusher (or another WebSocket library) and receive real-time `eventID` notifications for changed events

Your API key will have limits on concurrent streams.

## Available Feeds

Subscribe to different feeds using the `feed` query parameter:

| Feed              | Description                                                                 | Required Parameters |
| ----------------- | --------------------------------------------------------------------------- | ------------------- |
| `events:live`     | All events currently in progress (started but not finished)                | None                |
| `events:upcoming` | Upcoming events with available odds for a specific league                  | `leagueID`          |
| `events:byid`     | Updates for a single specific event                                         | `eventID`           |

The number of supported feeds will increase over time. Please reach out if you have a use case which can't be covered by these feeds.

## Quick Start Example

Here's the minimal code to connect to live events:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/SportsGameOdds/sports-odds-api-go"
    "github.com/SportsGameOdds/sports-odds-api-go/option"
    "github.com/pusher/pusher-http-go"
)

func main() {
    const STREAM_FEED = "events:live" // ex: events:upcoming, events:byid, events:live
    const API_KEY = "YOUR API KEY"

    client := sportsoddsapi.NewClient(option.WithAPIKeyParam(API_KEY))

    // Call this endpoint to get initial data and connection parameters
    streamInfo, err := client.Stream.Events(context.TODO(), sportsoddsapi.StreamEventParams{Feed: sportsoddsapi.String(STREAM_FEED)})
    if err != nil {
        log.Fatal(err)
    }

    // Seed initial data
    events := make(map[string]any)
    for _, event := range streamInfo.Data {
        events[event.EventID] = event
    }

    // Connect to WebSocket server
    pusherClient := pusher.Client{
        AppID: streamInfo.PusherKey,
        Key:   streamInfo.PusherKey,
        Secret: "your_secret",
    }

    channel := pusherClient.Channel(streamInfo.Channel)
    channel.Bind("data", func(changedEvents []map[string]any) {
        for _, changed := range changedEvents {
            eventID := changed["eventID"].(string)
            page, _ := client.Events.Get(context.TODO(), sportsoddsapi.EventGetParams{})
            for _, event := range page.Data {
                if event.EventID == eventID {
                    events[eventID] = event
                }
            }
        }
    })

    fmt.Println("Streaming started...")
}
```

### Request & Response types

This library includes Go type definitions for all request params and response fields.  
Responses are returned as Go structs, providing full type safety and IDE autocomplete.

## Handling errors

When the library is unable to connect to the API,
or if the API returns a non-success status code (i.e., 4xx or 5xx response),
an error of type `*sportsoddsapi.Error` will be returned:

```go
_, err := client.Events.Get(context.TODO(), sportsoddsapi.EventGetParams{})
if err != nil {
    var apierr *sportsoddsapi.Error
    if errors.As(err, &apierr) {
        fmt.Println(apierr.StatusCode)
    } else {
        panic(err)
    }
}
```

Error codes are as follows:

| Status Code | Error Type                 |
| ----------- | -------------------------- |
| 400         | `BadRequestError`          |
| 401         | `AuthenticationError`      |
| 403         | `PermissionDeniedError`    |
| 404         | `NotFoundError`            |
| 422         | `UnprocessableEntityError` |
| 429         | `RateLimitError`           |
| >=500       | `InternalServerError`      |
| N/A         | `APIConnectionError`       |

### Retries

Certain errors will be automatically retried 2 times by default, with exponential backoff.  
You can configure retries with the `option.WithMaxRetries` option:

```go
client := sportsoddsapi.NewClient(option.WithMaxRetries(0)) // default is 2

client.Events.Get(
    context.TODO(),
    sportsoddsapi.EventGetParams{},
    option.WithMaxRetries(5),
)
```

### Timeouts

Requests do not time out by default; use context to configure a timeout:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

client.Events.Get(ctx, sportsoddsapi.EventGetParams{}, option.WithRequestTimeout(20*time.Second))
```

## Auto-pagination

```go
iter := client.Events.GetAutoPaging(context.TODO(), sportsoddsapi.EventGetParams{
    Limit: sportsoddsapi.Float(30),
})
for iter.Next() {
    event := iter.Current()
    fmt.Printf("%+v\n", event)
}
if err := iter.Err(); err != nil {
    panic(err.Error())
}
```

Or with `.GetNextPage()`:

```go
page, err := client.Events.Get(context.TODO(), sportsoddsapi.EventGetParams{Limit: sportsoddsapi.Float(30)})
for page != nil {
    for _, event := range page.Data {
        fmt.Printf("%+v\n", event)
    }
    page, err = page.GetNextPage()
}
```

## Advanced Usage

### Accessing raw response data (e.g., headers)

```go
var response *http.Response
page, err := client.Events.Get(
    context.TODO(),
    sportsoddsapi.EventGetParams{},
    option.WithResponseInto(&response),
)
fmt.Println(response.StatusCode)
fmt.Println(response.Header)
```

### Making custom/undocumented requests

```go
var result *http.Response
params := map[string]any{"my_param": true}
err := client.Post(context.Background(), "/foo", params, &result)
if err != nil {
    panic(err)
}
```

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/SportsGameOdds/sports-odds-api-go/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
