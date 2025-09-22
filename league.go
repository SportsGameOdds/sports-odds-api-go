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
	"github.com/stainless-sdks/sports-odds-api-go/packages/param"
	"github.com/stainless-sdks/sports-odds-api-go/packages/respjson"
)

// LeagueService contains methods and other services that help with interacting
// with the SportsGameOdds API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLeagueService] method instead.
type LeagueService struct {
	Options []option.RequestOption
}

// NewLeagueService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLeagueService(opts ...option.RequestOption) (r LeagueService) {
	r = LeagueService{}
	r.Options = opts
	return
}

// Get a list of Leagues
func (r *LeagueService) Get(ctx context.Context, query LeagueGetParams, opts ...option.RequestOption) (res *[]League, err error) {
	var env LeagueGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "leagues/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

type League struct {
	Enabled   bool   `json:"enabled"`
	LeagueID  string `json:"leagueID"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	SportID   string `json:"sportID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		LeagueID    respjson.Field
		Name        respjson.Field
		ShortName   respjson.Field
		SportID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r League) RawJSON() string { return r.JSON.raw }
func (r *League) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LeagueGetParams struct {
	// The league to get data for
	LeagueID param.Opt[string] `query:"leagueID,omitzero" json:"-"`
	// The sport to get leagues for
	SportID param.Opt[string] `query:"sportID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LeagueGetParams]'s query parameters as `url.Values`.
func (r LeagueGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LeagueGetResponseEnvelope struct {
	Data []League `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LeagueGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *LeagueGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
