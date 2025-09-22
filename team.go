// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sportsoddsapi

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/sports-odds-api-go/internal/apijson"
	"github.com/stainless-sdks/sports-odds-api-go/internal/apiquery"
	"github.com/stainless-sdks/sports-odds-api-go/internal/requestconfig"
	"github.com/stainless-sdks/sports-odds-api-go/option"
	"github.com/stainless-sdks/sports-odds-api-go/packages/pagination"
	"github.com/stainless-sdks/sports-odds-api-go/packages/param"
	"github.com/stainless-sdks/sports-odds-api-go/packages/respjson"
)

// TeamService contains methods and other services that help with interacting with
// the SportsGameOdds API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamService] method instead.
type TeamService struct {
	Options []option.RequestOption
}

// NewTeamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTeamService(opts ...option.RequestOption) (r TeamService) {
	r = TeamService{}
	r.Options = opts
	return
}

// Get a list of Teams by ID or league
func (r *TeamService) Get(ctx context.Context, query TeamGetParams, opts ...option.RequestOption) (res *pagination.NextCursorPage[Team], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "teams/"
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

// Get a list of Teams by ID or league
func (r *TeamService) GetAutoPaging(ctx context.Context, query TeamGetParams, opts ...option.RequestOption) *pagination.NextCursorPageAutoPager[Team] {
	return pagination.NewNextCursorPageAutoPager(r.Get(ctx, query, opts...))
}

type Team struct {
	Colors    TeamColors    `json:"colors"`
	LeagueID  string        `json:"leagueID"`
	Logo      string        `json:"logo"`
	Lookups   TeamLookups   `json:"lookups"`
	Names     TeamNames     `json:"names"`
	SportID   string        `json:"sportID"`
	Standings TeamStandings `json:"standings"`
	TeamID    string        `json:"teamID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Colors      respjson.Field
		LeagueID    respjson.Field
		Logo        respjson.Field
		Lookups     respjson.Field
		Names       respjson.Field
		SportID     respjson.Field
		Standings   respjson.Field
		TeamID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Team) RawJSON() string { return r.JSON.raw }
func (r *Team) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamColors struct {
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
func (r TeamColors) RawJSON() string { return r.JSON.raw }
func (r *TeamColors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamLookups struct {
	TeamName []string `json:"teamName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TeamName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamLookups) RawJSON() string { return r.JSON.raw }
func (r *TeamLookups) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamNames struct {
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
func (r TeamNames) RawJSON() string { return r.JSON.raw }
func (r *TeamNames) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamStandings struct {
	Losses   float64 `json:"losses"`
	Played   float64 `json:"played"`
	Position string  `json:"position"`
	Record   string  `json:"record"`
	Ties     float64 `json:"ties"`
	Wins     float64 `json:"wins"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Losses      respjson.Field
		Played      respjson.Field
		Position    respjson.Field
		Record      respjson.Field
		Ties        respjson.Field
		Wins        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamStandings) RawJSON() string { return r.JSON.raw }
func (r *TeamStandings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamGetParams struct {
	// The cursor for the request. Used to get the next group of Teams. This should be
	// the nextCursor from the prior response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// A single leagueID or comma-separated list of leagueIDs to get Teams for
	LeagueID param.Opt[string] `query:"leagueID,omitzero" json:"-"`
	// The maximum number of Teams to return
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// A single sportID or comma-separated list of sportIDs to get Teams for
	SportID param.Opt[string] `query:"sportID,omitzero" json:"-"`
	// A single teamID or comma-separated list of teamIDs to get data for
	TeamID param.Opt[string] `query:"teamID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TeamGetParams]'s query parameters as `url.Values`.
func (r TeamGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
