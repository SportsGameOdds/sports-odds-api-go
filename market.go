// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sportsoddsapi

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/SportsGameOdds/sports-odds-api-go/internal/apijson"
	"github.com/SportsGameOdds/sports-odds-api-go/internal/apiquery"
	"github.com/SportsGameOdds/sports-odds-api-go/internal/requestconfig"
	"github.com/SportsGameOdds/sports-odds-api-go/option"
	"github.com/SportsGameOdds/sports-odds-api-go/packages/pagination"
	"github.com/SportsGameOdds/sports-odds-api-go/packages/param"
	"github.com/SportsGameOdds/sports-odds-api-go/packages/respjson"
)

// Get metadata on supported Markets
//
// MarketService contains methods and other services that help with interacting
// with the SportsGameOdds API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketService] method instead.
type MarketService struct {
	Options []option.RequestOption
}

// NewMarketService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMarketService(opts ...option.RequestOption) (r MarketService) {
	r = MarketService{}
	r.Options = opts
	return
}

// Get a list of Markets
func (r *MarketService) Get(ctx context.Context, query MarketGetParams, opts ...option.RequestOption) (res *pagination.NextCursorPage[Market], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "markets/"
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

// Get a list of Markets
func (r *MarketService) GetAutoPaging(ctx context.Context, query MarketGetParams, opts ...option.RequestOption) *pagination.NextCursorPageAutoPager[Market] {
	return pagination.NewNextCursorPageAutoPager(r.Get(ctx, query, opts...))
}

type Market struct {
	// The number of unique active events with available odds for this market across
	// all supported league and bookmaker combinations.
	ActiveEvents float64 `json:"activeEvents"`
	// The type of bet
	BetTypeID string `json:"betTypeID"`
	// True if this is a sub-period of a main market
	IsMainDerivative bool `json:"isMainDerivative"`
	// True if this is a main market
	IsMainMarket bool `json:"isMainMarket"`
	// True if this is a prop bet
	IsProp bool `json:"isProp"`
	// True if this market is for a sub-period
	IsSubPeriod bool `json:"isSubPeriod"`
	// True if this market is supported by at least one league/bookmaker.
	IsSupported bool `json:"isSupported"`
	// The unique identifier for the group (all sides of the market) this market
	// belongs to
	MarketGroupID string `json:"marketGroupID"`
	// The primary display name for this market's group
	MarketGroupName string `json:"marketGroupName"`
	// An alternative display name for this market's group
	MarketGroupNameAlias string `json:"marketGroupNameAlias"`
	// Sport-specific market group names when they differ from the primary name
	MarketGroupNameBySport map[string]string `json:"marketGroupNameBySport"`
	// The unique identifier for this market
	OddID string `json:"oddID"`
	// The period of the event this market applies to
	PeriodID string `json:"periodID"`
	// Set to a player's unique playerID if it's a player prop
	PlayerID string `json:"playerID"`
	// The type of prop bet
	//
	// Any of "game_prop", "team_prop", "player_prop", "other_prop".
	PropType MarketPropType `json:"propType"`
	// The side of the bet
	SideID string `json:"sideID"`
	// The statEntityID represents whose performance on the stat is being evaluated
	StatEntityID string `json:"statEntityID"`
	// The statistic which is being evaluated as a part of this market
	StatID string `json:"statID"`
	// Nested object showing which leagues and bookmakers support this market.
	Support map[string]map[string]MarketSupport `json:"support"`
	// Set to team's unique teamID if it's a team prop for a tournament type event
	TeamID string `json:"teamID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveEvents           respjson.Field
		BetTypeID              respjson.Field
		IsMainDerivative       respjson.Field
		IsMainMarket           respjson.Field
		IsProp                 respjson.Field
		IsSubPeriod            respjson.Field
		IsSupported            respjson.Field
		MarketGroupID          respjson.Field
		MarketGroupName        respjson.Field
		MarketGroupNameAlias   respjson.Field
		MarketGroupNameBySport respjson.Field
		OddID                  respjson.Field
		PeriodID               respjson.Field
		PlayerID               respjson.Field
		PropType               respjson.Field
		SideID                 respjson.Field
		StatEntityID           respjson.Field
		StatID                 respjson.Field
		Support                respjson.Field
		TeamID                 respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Market) RawJSON() string { return r.JSON.raw }
func (r *Market) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of prop bet
type MarketPropType string

const (
	MarketPropTypeGameProp   MarketPropType = "game_prop"
	MarketPropTypeTeamProp   MarketPropType = "team_prop"
	MarketPropTypePlayerProp MarketPropType = "player_prop"
	MarketPropTypeOtherProp  MarketPropType = "other_prop"
)

type MarketSupport struct {
	// Whether this market is supported for the given league and bookmaker combination.
	Supported bool `json:"supported"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Supported   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketSupport) RawJSON() string { return r.JSON.raw }
func (r *MarketSupport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketGetParams struct {
	// A single betTypeID or comma-separated list of betTypeIDs to filter Markets by
	BetTypeID param.Opt[string] `query:"betTypeID,omitzero" json:"-"`
	// A single bookmakerID or comma-separated list of bookmakerIDs to filter Markets
	// by
	BookmakerID param.Opt[string] `query:"bookmakerID,omitzero" json:"-"`
	// The cursor for pagination. Use nextCursor from prior response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter to only include main markets (main period moneyline, spread, and
	// over/under)
	IsMainMarket param.Opt[bool] `query:"isMainMarket,omitzero" json:"-"`
	// Filter by whether it is any type of prop bet market
	IsProp param.Opt[bool] `query:"isProp,omitzero" json:"-"`
	// Filter by whether it tracks a sub/non-main period
	IsSubPeriod param.Opt[bool] `query:"isSubPeriod,omitzero" json:"-"`
	// Filter whether this market is fully supported by at least 1 bookmaker in at
	// least 1 league. Defaults to true if not specified.
	IsSupported param.Opt[bool] `query:"isSupported,omitzero" json:"-"`
	// A single leagueID or comma-separated list of leagueIDs to filter Markets by
	LeagueID param.Opt[string] `query:"leagueID,omitzero" json:"-"`
	// The maximum number of Markets to return (default: 100, max: 10000)
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// A single oddID or comma-separated list of oddIDs. Used to specify specific
	// Markets to return.
	OddID param.Opt[string] `query:"oddID,omitzero" json:"-"`
	// A single periodID or comma-separated list of periodIDs to filter Markets by
	PeriodID param.Opt[string] `query:"periodID,omitzero" json:"-"`
	// Filter by prop type (game_prop, team_prop, player_prop, other_prop)
	PropType param.Opt[string] `query:"propType,omitzero" json:"-"`
	// A single sideID or comma-separated list of sideIDs to filter Markets by
	SideID param.Opt[string] `query:"sideID,omitzero" json:"-"`
	// A single sportID or comma-separated list of sportIDs to filter Markets by
	SportID param.Opt[string] `query:"sportID,omitzero" json:"-"`
	// A single statEntityID or comma-separated list of statEntityIDs to filter Markets
	// by
	StatEntityID param.Opt[string] `query:"statEntityID,omitzero" json:"-"`
	// A single statID or comma-separated list of statIDs to filter Markets by
	StatID param.Opt[string] `query:"statID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MarketGetParams]'s query parameters as `url.Values`.
func (r MarketGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
