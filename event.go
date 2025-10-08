// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sportsoddsapi

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/SportsGameOdds/sports-odds-api-go/internal/apijson"
	"github.com/SportsGameOdds/sports-odds-api-go/internal/apiquery"
	"github.com/SportsGameOdds/sports-odds-api-go/internal/requestconfig"
	"github.com/SportsGameOdds/sports-odds-api-go/option"
	"github.com/SportsGameOdds/sports-odds-api-go/packages/pagination"
	"github.com/SportsGameOdds/sports-odds-api-go/packages/param"
	"github.com/SportsGameOdds/sports-odds-api-go/packages/respjson"
)

// EventService contains methods and other services that help with interacting with
// the SportsGameOdds API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.Options = opts
	return
}

// Get a list of Events
func (r *EventService) Get(ctx context.Context, query EventGetParams, opts ...option.RequestOption) (res *pagination.NextCursorPage[Event], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "events/"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of Events
func (r *EventService) GetAutoPaging(ctx context.Context, query EventGetParams, opts ...option.RequestOption) *pagination.NextCursorPageAutoPager[Event] {
	return pagination.NewNextCursorPageAutoPager(r.Get(ctx, query, opts...))
}

type Event struct {
	Activity EventActivity                            `json:"activity"`
	EventID  string                                   `json:"eventID"`
	Info     EventInfo                                `json:"info"`
	LeagueID string                                   `json:"leagueID"`
	Manual   bool                                     `json:"manual"`
	Odds     map[string]EventOdd                      `json:"odds"`
	Players  map[string]EventPlayer                   `json:"players"`
	Results  map[string]map[string]map[string]float64 `json:"results"`
	SportID  string                                   `json:"sportID"`
	Status   EventStatus                              `json:"status"`
	Teams    EventTeams                               `json:"teams"`
	Type     string                                   `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Activity    respjson.Field
		EventID     respjson.Field
		Info        respjson.Field
		LeagueID    respjson.Field
		Manual      respjson.Field
		Odds        respjson.Field
		Players     respjson.Field
		Results     respjson.Field
		SportID     respjson.Field
		Status      respjson.Field
		Teams       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Event) RawJSON() string { return r.JSON.raw }
func (r *Event) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventActivity struct {
	Count float64 `json:"count"`
	Score float64 `json:"score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventActivity) RawJSON() string { return r.JSON.raw }
func (r *EventActivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventInfo struct {
	SeasonWeek string `json:"seasonWeek"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SeasonWeek  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventInfo) RawJSON() string { return r.JSON.raw }
func (r *EventInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventOdd struct {
	BetTypeID         string                         `json:"betTypeID"`
	BookOdds          string                         `json:"bookOdds"`
	BookOddsAvailable bool                           `json:"bookOddsAvailable"`
	BookOverUnder     string                         `json:"bookOverUnder"`
	BookSpread        string                         `json:"bookSpread"`
	ByBookmaker       map[string]EventOddByBookmaker `json:"byBookmaker"`
	Cancelled         bool                           `json:"cancelled"`
	Ended             bool                           `json:"ended"`
	FairOdds          string                         `json:"fairOdds"`
	FairOddsAvailable bool                           `json:"fairOddsAvailable"`
	FairOverUnder     string                         `json:"fairOverUnder"`
	FairSpread        string                         `json:"fairSpread"`
	MarketName        string                         `json:"marketName"`
	OddID             string                         `json:"oddID"`
	OpposingOddID     string                         `json:"opposingOddID"`
	PeriodID          string                         `json:"periodID"`
	PlayerID          string                         `json:"playerID"`
	Score             float64                        `json:"score"`
	ScoringSupported  bool                           `json:"scoringSupported"`
	SideID            string                         `json:"sideID"`
	Started           bool                           `json:"started"`
	StatEntityID      string                         `json:"statEntityID"`
	StatID            string                         `json:"statID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BetTypeID         respjson.Field
		BookOdds          respjson.Field
		BookOddsAvailable respjson.Field
		BookOverUnder     respjson.Field
		BookSpread        respjson.Field
		ByBookmaker       respjson.Field
		Cancelled         respjson.Field
		Ended             respjson.Field
		FairOdds          respjson.Field
		FairOddsAvailable respjson.Field
		FairOverUnder     respjson.Field
		FairSpread        respjson.Field
		MarketName        respjson.Field
		OddID             respjson.Field
		OpposingOddID     respjson.Field
		PeriodID          respjson.Field
		PlayerID          respjson.Field
		Score             respjson.Field
		ScoringSupported  respjson.Field
		SideID            respjson.Field
		Started           respjson.Field
		StatEntityID      respjson.Field
		StatID            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventOdd) RawJSON() string { return r.JSON.raw }
func (r *EventOdd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventOddByBookmaker struct {
	Available     bool      `json:"available"`
	BookmakerID   string    `json:"bookmakerID"`
	IsMainLine    bool      `json:"isMainLine"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt" format:"date-time"`
	Odds          string    `json:"odds"`
	OverUnder     string    `json:"overUnder"`
	Spread        string    `json:"spread"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Available     respjson.Field
		BookmakerID   respjson.Field
		IsMainLine    respjson.Field
		LastUpdatedAt respjson.Field
		Odds          respjson.Field
		OverUnder     respjson.Field
		Spread        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventOddByBookmaker) RawJSON() string { return r.JSON.raw }
func (r *EventOddByBookmaker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventPlayer struct {
	Alias     string `json:"alias"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Name      string `json:"name"`
	Photo     string `json:"photo"`
	PlayerID  string `json:"playerID"`
	TeamID    string `json:"teamID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alias       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		Name        respjson.Field
		Photo       respjson.Field
		PlayerID    respjson.Field
		TeamID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventPlayer) RawJSON() string { return r.JSON.raw }
func (r *EventPlayer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventStatus struct {
	Cancelled        bool               `json:"cancelled"`
	Completed        bool               `json:"completed"`
	CurrentPeriodID  string             `json:"currentPeriodID"`
	Delayed          bool               `json:"delayed"`
	DisplayLong      string             `json:"displayLong"`
	DisplayShort     string             `json:"displayShort"`
	Ended            bool               `json:"ended"`
	Finalized        bool               `json:"finalized"`
	HardStart        bool               `json:"hardStart"`
	Live             bool               `json:"live"`
	OddsAvailable    bool               `json:"oddsAvailable"`
	OddsPresent      bool               `json:"oddsPresent"`
	Periods          EventStatusPeriods `json:"periods"`
	PreviousPeriodID string             `json:"previousPeriodID"`
	ReGrade          bool               `json:"reGrade"`
	Started          bool               `json:"started"`
	StartsAt         time.Time          `json:"startsAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled        respjson.Field
		Completed        respjson.Field
		CurrentPeriodID  respjson.Field
		Delayed          respjson.Field
		DisplayLong      respjson.Field
		DisplayShort     respjson.Field
		Ended            respjson.Field
		Finalized        respjson.Field
		HardStart        respjson.Field
		Live             respjson.Field
		OddsAvailable    respjson.Field
		OddsPresent      respjson.Field
		Periods          respjson.Field
		PreviousPeriodID respjson.Field
		ReGrade          respjson.Field
		Started          respjson.Field
		StartsAt         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventStatus) RawJSON() string { return r.JSON.raw }
func (r *EventStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventStatusPeriods struct {
	Ended   []string `json:"ended"`
	Started []string `json:"started"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ended       respjson.Field
		Started     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventStatusPeriods) RawJSON() string { return r.JSON.raw }
func (r *EventStatusPeriods) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTeams struct {
	Away EventTeamsAway `json:"away"`
	Home EventTeamsHome `json:"home"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Away        respjson.Field
		Home        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventTeams) RawJSON() string { return r.JSON.raw }
func (r *EventTeams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTeamsAway struct {
	Colors       EventTeamsAwayColors `json:"colors"`
	Logo         string               `json:"logo"`
	Names        EventTeamsAwayNames  `json:"names"`
	Score        float64              `json:"score"`
	StatEntityID string               `json:"statEntityID"`
	TeamID       string               `json:"teamID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Colors       respjson.Field
		Logo         respjson.Field
		Names        respjson.Field
		Score        respjson.Field
		StatEntityID respjson.Field
		TeamID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventTeamsAway) RawJSON() string { return r.JSON.raw }
func (r *EventTeamsAway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTeamsAwayColors struct {
	Primary           string `json:"primary"`
	PrimaryContrast   string `json:"primaryContrast"`
	Secondary         string `json:"secondary"`
	SecondaryContrast string `json:"secondaryContrast"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary           respjson.Field
		PrimaryContrast   respjson.Field
		Secondary         respjson.Field
		SecondaryContrast respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventTeamsAwayColors) RawJSON() string { return r.JSON.raw }
func (r *EventTeamsAwayColors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTeamsAwayNames struct {
	Long   string `json:"long"`
	Medium string `json:"medium"`
	Short  string `json:"short"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Long        respjson.Field
		Medium      respjson.Field
		Short       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventTeamsAwayNames) RawJSON() string { return r.JSON.raw }
func (r *EventTeamsAwayNames) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTeamsHome struct {
	Colors       EventTeamsHomeColors `json:"colors"`
	Logo         string               `json:"logo"`
	Names        EventTeamsHomeNames  `json:"names"`
	Score        float64              `json:"score"`
	StatEntityID string               `json:"statEntityID"`
	TeamID       string               `json:"teamID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Colors       respjson.Field
		Logo         respjson.Field
		Names        respjson.Field
		Score        respjson.Field
		StatEntityID respjson.Field
		TeamID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventTeamsHome) RawJSON() string { return r.JSON.raw }
func (r *EventTeamsHome) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTeamsHomeColors struct {
	Primary           string `json:"primary"`
	PrimaryContrast   string `json:"primaryContrast"`
	Secondary         string `json:"secondary"`
	SecondaryContrast string `json:"secondaryContrast"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary           respjson.Field
		PrimaryContrast   respjson.Field
		Secondary         respjson.Field
		SecondaryContrast respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventTeamsHomeColors) RawJSON() string { return r.JSON.raw }
func (r *EventTeamsHomeColors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTeamsHomeNames struct {
	Long   string `json:"long"`
	Medium string `json:"medium"`
	Short  string `json:"short"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Long        respjson.Field
		Medium      respjson.Field
		Short       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventTeamsHomeNames) RawJSON() string { return r.JSON.raw }
func (r *EventTeamsHomeNames) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventGetParams struct {
	// A bookmakerID or comma-separated list of bookmakerIDs to include odds for
	BookmakerID param.Opt[string] `query:"bookmakerID,omitzero" json:"-"`
	// Only include cancelled Events (true), only non-cancelled Events (false) or all
	// Events (omit)
	Cancelled param.Opt[bool] `query:"cancelled,omitzero" json:"-"`
	// The cursor for the request. Used to get the next group of Events. This should be
	// the nextCursor from the prior response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only include Events which have have ended (true), only Events which have not
	// ended (false) or all Events (omit)
	Ended param.Opt[bool] `query:"ended,omitzero" json:"-"`
	// An eventID to get Event data for
	EventID param.Opt[string] `query:"eventID,omitzero" json:"-"`
	// A comma separated list of eventIDs to get Event data for
	EventIDs param.Opt[string] `query:"eventIDs,omitzero" json:"-"`
	// Only include finalized Events (true), exclude unfinalized Events (false) or all
	// Events (omit)
	Finalized param.Opt[bool] `query:"finalized,omitzero" json:"-"`
	// Whether to include alternate lines in the odds byBookmaker data
	IncludeAltLines param.Opt[bool] `query:"includeAltLines,omitzero" json:"-"`
	// Whether to include opposing odds for each included oddID
	IncludeOpposingOdds param.Opt[bool] `query:"includeOpposingOdds,omitzero" json:"-"`
	// A leagueID or comma-separated list of leagueIDs to get Events for
	LeagueID param.Opt[string] `query:"leagueID,omitzero" json:"-"`
	// The maximum number of Events to return
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Only include live Events (true), only non-live Events (false) or all Events
	// (omit)
	Live param.Opt[bool] `query:"live,omitzero" json:"-"`
	// An oddID or comma-separated list of oddIDs to include odds for
	OddID param.Opt[string] `query:"oddID,omitzero" json:"-"`
	// Whether you want only Events which do (true) or do not (false) have odds markets
	// which are currently available (open for wagering)
	OddsAvailable param.Opt[bool] `query:"oddsAvailable,omitzero" json:"-"`
	// Whether you want only Events which do (true) or do not (false) have any
	// associated odds markets regardless of whether those odds markets are currently
	// available (open for wagering)
	OddsPresent param.Opt[bool] `query:"oddsPresent,omitzero" json:"-"`
	// A playerID or comma-separated list of playerIDs to include Events (and
	// associated odds) for
	PlayerID param.Opt[string] `query:"playerID,omitzero" json:"-"`
	// A sportID or comma-separated list of sportIDs to get Events for
	SportID param.Opt[string] `query:"sportID,omitzero" json:"-"`
	// Only include Events which have have previously started (true), only Events which
	// have not previously started (false) or all Events (omit)
	Started param.Opt[bool] `query:"started,omitzero" json:"-"`
	// Get Events that start after this date
	StartsAfter param.Opt[time.Time] `query:"startsAfter,omitzero" format:"date-time" json:"-"`
	// Get Events that start before this date
	StartsBefore param.Opt[time.Time] `query:"startsBefore,omitzero" format:"date-time" json:"-"`
	// A teamID or comma-separated list of teamIDs to include Events for
	TeamID param.Opt[string] `query:"teamID,omitzero" json:"-"`
	// Only include Events of the specified type
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventGetParams]'s query parameters as `url.Values`.
func (r EventGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
