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

// PlayerService contains methods and other services that help with interacting
// with the SportsGameOdds API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlayerService] method instead.
type PlayerService struct {
	Options []option.RequestOption
}

// NewPlayerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPlayerService(opts ...option.RequestOption) (r PlayerService) {
	r = PlayerService{}
	r.Options = opts
	return
}

// Get a list of Players for a specific Team or Event
func (r *PlayerService) Get(ctx context.Context, query PlayerGetParams, opts ...option.RequestOption) (res *pagination.NextCursorPage[Player], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "players/"
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

// Get a list of Players for a specific Team or Event
func (r *PlayerService) GetAutoPaging(ctx context.Context, query PlayerGetParams, opts ...option.RequestOption) *pagination.NextCursorPageAutoPager[Player] {
	return pagination.NewNextCursorPageAutoPager(r.Get(ctx, query, opts...))
}

type Player struct {
	Aliases      []string                    `json:"aliases"`
	JerseyNumber float64                     `json:"jerseyNumber"`
	LeagueID     string                      `json:"leagueID"`
	Lookups      PlayerLookups               `json:"lookups"`
	Names        PlayerNames                 `json:"names"`
	PlayerID     string                      `json:"playerID"`
	PlayerTeams  map[string]PlayerPlayerTeam `json:"playerTeams"`
	Position     string                      `json:"position"`
	SportID      string                      `json:"sportID"`
	TeamID       string                      `json:"teamID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aliases      respjson.Field
		JerseyNumber respjson.Field
		LeagueID     respjson.Field
		Lookups      respjson.Field
		Names        respjson.Field
		PlayerID     respjson.Field
		PlayerTeams  respjson.Field
		Position     respjson.Field
		SportID      respjson.Field
		TeamID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Player) RawJSON() string { return r.JSON.raw }
func (r *Player) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlayerLookups struct {
	AnyName  []string `json:"anyName"`
	FullName []string `json:"fullName"`
	Initials []string `json:"initials"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnyName     respjson.Field
		FullName    respjson.Field
		Initials    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlayerLookups) RawJSON() string { return r.JSON.raw }
func (r *PlayerLookups) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlayerNames struct {
	Display   string `json:"display"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Display     respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlayerNames) RawJSON() string { return r.JSON.raw }
func (r *PlayerNames) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlayerPlayerTeam struct {
	TeamID string `json:"teamID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TeamID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlayerPlayerTeam) RawJSON() string { return r.JSON.raw }
func (r *PlayerPlayerTeam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlayerGetParams struct {
	// The cursor for the request. Used to get the next group of Players. This should
	// be the nextCursor from the prior response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// EventID to get Players data for
	EventID param.Opt[string] `query:"eventID,omitzero" json:"-"`
	// The maximum number of Players to return
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// PlayerID to get data for
	PlayerID param.Opt[string] `query:"playerID,omitzero" json:"-"`
	// TeamID to get Players data for
	TeamID param.Opt[string] `query:"teamID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PlayerGetParams]'s query parameters as `url.Values`.
func (r PlayerGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
