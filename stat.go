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
	"github.com/SportsGameOdds/sports-odds-api-go/packages/param"
	"github.com/SportsGameOdds/sports-odds-api-go/packages/respjson"
)

// StatService contains methods and other services that help with interacting with
// the SportsGameOdds API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatService] method instead.
type StatService struct {
	Options []option.RequestOption
}

// NewStatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStatService(opts ...option.RequestOption) (r StatService) {
	r = StatService{}
	r.Options = opts
	return
}

// Get a list of StatIDs
func (r *StatService) Get(ctx context.Context, query StatGetParams, opts ...option.RequestOption) (res *[]Stat, err error) {
	var env StatGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "stats/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

type Stat struct {
	Description     string              `json:"description"`
	Displays        StatDisplays        `json:"displays"`
	IsScoreStat     bool                `json:"isScoreStat"`
	StatID          string              `json:"statID"`
	SupportedLevels StatSupportedLevels `json:"supportedLevels"`
	SupportedSports map[string]any      `json:"supportedSports"`
	Units           StatUnits           `json:"units"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description     respjson.Field
		Displays        respjson.Field
		IsScoreStat     respjson.Field
		StatID          respjson.Field
		SupportedLevels respjson.Field
		SupportedSports respjson.Field
		Units           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Stat) RawJSON() string { return r.JSON.raw }
func (r *Stat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatDisplays struct {
	Long  string `json:"long"`
	Short string `json:"short"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Long        respjson.Field
		Short       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatDisplays) RawJSON() string { return r.JSON.raw }
func (r *StatDisplays) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatSupportedLevels struct {
	All    bool `json:"all"`
	Player bool `json:"player"`
	Team   bool `json:"team"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		All         respjson.Field
		Player      respjson.Field
		Team        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatSupportedLevels) RawJSON() string { return r.JSON.raw }
func (r *StatSupportedLevels) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatUnits struct {
	Long  StatUnitsLong  `json:"long"`
	Short StatUnitsShort `json:"short"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Long        respjson.Field
		Short       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatUnits) RawJSON() string { return r.JSON.raw }
func (r *StatUnits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatUnitsLong struct {
	Plural   string `json:"plural"`
	Singular string `json:"singular"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Plural      respjson.Field
		Singular    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatUnitsLong) RawJSON() string { return r.JSON.raw }
func (r *StatUnitsLong) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatUnitsShort struct {
	Plural   string `json:"plural"`
	Singular string `json:"singular"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Plural      respjson.Field
		Singular    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatUnitsShort) RawJSON() string { return r.JSON.raw }
func (r *StatUnitsShort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatGetParams struct {
	// SportID to get StatIDs for
	SportID param.Opt[string] `query:"sportID,omitzero" json:"-"`
	// StatID to get data for
	StatID param.Opt[string] `query:"statID,omitzero" json:"-"`
	// Level of the stat, must be used in combination with sportID. Must be one of all,
	// player, or team. Shows stats that are applicable to that specified entity,
	// defaults to all.
	StatLevel param.Opt[string] `query:"statLevel,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatGetParams]'s query parameters as `url.Values`.
func (r StatGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StatGetResponseEnvelope struct {
	Data []Stat `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *StatGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
