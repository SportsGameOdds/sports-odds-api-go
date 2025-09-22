// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sportsoddsapi

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/sports-odds-api-go/internal/apijson"
	"github.com/stainless-sdks/sports-odds-api-go/internal/requestconfig"
	"github.com/stainless-sdks/sports-odds-api-go/option"
	"github.com/stainless-sdks/sports-odds-api-go/packages/respjson"
)

// SportService contains methods and other services that help with interacting with
// the SportsGameOdds API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSportService] method instead.
type SportService struct {
	Options []option.RequestOption
}

// NewSportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSportService(opts ...option.RequestOption) (r SportService) {
	r = SportService{}
	r.Options = opts
	return
}

// Get a list of sports
func (r *SportService) Get(ctx context.Context, opts ...option.RequestOption) (res *[]Sport, err error) {
	var env SportGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "sports/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

type Sport struct {
	BackgroundImage        string         `json:"backgroundImage"`
	BasePeriods            []string       `json:"basePeriods"`
	ClockType              string         `json:"clockType"`
	DefaultPopularityScore float64        `json:"defaultPopularityScore"`
	Enabled                bool           `json:"enabled"`
	EventWord              SportEventWord `json:"eventWord"`
	ExtraPeriods           []string       `json:"extraPeriods"`
	HasMeaningfulHomeAway  bool           `json:"hasMeaningfulHomeAway"`
	ImageIcon              string         `json:"imageIcon"`
	Name                   string         `json:"name"`
	PointWord              SportPointWord `json:"pointWord"`
	ShortName              string         `json:"shortName"`
	SportID                string         `json:"sportID"`
	SquareImage            string         `json:"squareImage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BackgroundImage        respjson.Field
		BasePeriods            respjson.Field
		ClockType              respjson.Field
		DefaultPopularityScore respjson.Field
		Enabled                respjson.Field
		EventWord              respjson.Field
		ExtraPeriods           respjson.Field
		HasMeaningfulHomeAway  respjson.Field
		ImageIcon              respjson.Field
		Name                   respjson.Field
		PointWord              respjson.Field
		ShortName              respjson.Field
		SportID                respjson.Field
		SquareImage            respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Sport) RawJSON() string { return r.JSON.raw }
func (r *Sport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SportEventWord struct {
	Long  SportEventWordLong  `json:"long"`
	Short SportEventWordShort `json:"short"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Long        respjson.Field
		Short       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportEventWord) RawJSON() string { return r.JSON.raw }
func (r *SportEventWord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SportEventWordLong struct {
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
func (r SportEventWordLong) RawJSON() string { return r.JSON.raw }
func (r *SportEventWordLong) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SportEventWordShort struct {
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
func (r SportEventWordShort) RawJSON() string { return r.JSON.raw }
func (r *SportEventWordShort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SportPointWord struct {
	Long  SportPointWordLong  `json:"long"`
	Short SportPointWordShort `json:"short"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Long        respjson.Field
		Short       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportPointWord) RawJSON() string { return r.JSON.raw }
func (r *SportPointWord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SportPointWordLong struct {
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
func (r SportPointWordLong) RawJSON() string { return r.JSON.raw }
func (r *SportPointWordLong) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SportPointWordShort struct {
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
func (r SportPointWordShort) RawJSON() string { return r.JSON.raw }
func (r *SportPointWordShort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SportGetResponseEnvelope struct {
	Data []Sport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SportGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *SportGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
