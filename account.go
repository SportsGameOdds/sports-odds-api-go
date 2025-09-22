// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sportsoddsapi

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/SportsGameOdds/sports-odds-api-go/internal/apijson"
	"github.com/SportsGameOdds/sports-odds-api-go/internal/requestconfig"
	"github.com/SportsGameOdds/sports-odds-api-go/option"
	"github.com/SportsGameOdds/sports-odds-api-go/packages/respjson"
)

// AccountService contains methods and other services that help with interacting
// with the SportsGameOdds API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.Options = opts
	return
}

// Get rate-limits and usage data about your API key
func (r *AccountService) GetUsage(ctx context.Context, opts ...option.RequestOption) (res *AccountUsage, err error) {
	var env AccountGetUsageResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "account/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

type AccountUsage struct {
	// The Stripe customer ID for the account
	CustomerID string `json:"customerID"`
	// The email address associated with the account
	Email string `json:"email"`
	// Whether the API key is active
	IsActive bool `json:"isActive"`
	// The hashed identifier for the API key
	KeyID      string                 `json:"keyID"`
	RateLimits AccountUsageRateLimits `json:"rateLimits"`
	// The current subscription tier
	Tier string `json:"tier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID  respjson.Field
		Email       respjson.Field
		IsActive    respjson.Field
		KeyID       respjson.Field
		RateLimits  respjson.Field
		Tier        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountUsage) RawJSON() string { return r.JSON.raw }
func (r *AccountUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUsageRateLimits struct {
	PerDay    RateLimitInterval `json:"per-day"`
	PerHour   RateLimitInterval `json:"per-hour"`
	PerMinute RateLimitInterval `json:"per-minute"`
	PerMonth  RateLimitInterval `json:"per-month"`
	PerSecond RateLimitInterval `json:"per-second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PerDay      respjson.Field
		PerHour     respjson.Field
		PerMinute   respjson.Field
		PerMonth    respjson.Field
		PerSecond   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountUsageRateLimits) RawJSON() string { return r.JSON.raw }
func (r *AccountUsageRateLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitInterval struct {
	// Current number of entities accessed in the interval
	CurrentEntities int64 `json:"current-entities"`
	// Current number of requests made in the interval
	CurrentRequests int64 `json:"current-requests"`
	// Maximum allowed entity accesses in the interval
	MaxEntities RateLimitIntervalMaxEntitiesUnion `json:"max-entities"`
	// Maximum allowed requests in the interval
	MaxRequests RateLimitIntervalMaxRequestsUnion `json:"max-requests"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentEntities respjson.Field
		CurrentRequests respjson.Field
		MaxEntities     respjson.Field
		MaxRequests     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RateLimitInterval) RawJSON() string { return r.JSON.raw }
func (r *RateLimitInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RateLimitIntervalMaxEntitiesUnion contains all possible properties and values
// from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfRateLimitIntervalMaxEntitiesString OfInt]
type RateLimitIntervalMaxEntitiesUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfRateLimitIntervalMaxEntitiesString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfRateLimitIntervalMaxEntitiesString respjson.Field
		OfInt                                respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u RateLimitIntervalMaxEntitiesUnion) AsRateLimitIntervalMaxEntitiesString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RateLimitIntervalMaxEntitiesUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RateLimitIntervalMaxEntitiesUnion) RawJSON() string { return u.JSON.raw }

func (r *RateLimitIntervalMaxEntitiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitIntervalMaxEntitiesString string

const (
	RateLimitIntervalMaxEntitiesStringUnlimited RateLimitIntervalMaxEntitiesString = "unlimited"
)

// RateLimitIntervalMaxRequestsUnion contains all possible properties and values
// from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfRateLimitIntervalMaxRequestsString OfInt]
type RateLimitIntervalMaxRequestsUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfRateLimitIntervalMaxRequestsString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfRateLimitIntervalMaxRequestsString respjson.Field
		OfInt                                respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u RateLimitIntervalMaxRequestsUnion) AsRateLimitIntervalMaxRequestsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RateLimitIntervalMaxRequestsUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RateLimitIntervalMaxRequestsUnion) RawJSON() string { return u.JSON.raw }

func (r *RateLimitIntervalMaxRequestsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitIntervalMaxRequestsString string

const (
	RateLimitIntervalMaxRequestsStringUnlimited RateLimitIntervalMaxRequestsString = "unlimited"
)

type AccountGetUsageResponseEnvelope struct {
	Data    AccountUsage `json:"data"`
	Success bool         `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountGetUsageResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AccountGetUsageResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
