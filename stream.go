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

// StreamService contains methods and other services that help with interacting
// with the SportsGameOdds API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStreamService] method instead.
type StreamService struct {
	Options []option.RequestOption
}

// NewStreamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStreamService(opts ...option.RequestOption) (r StreamService) {
	r = StreamService{}
	r.Options = opts
	return
}

// Setup streamed (WebSocket) connection
func (r *StreamService) Events(ctx context.Context, query StreamEventsParams, opts ...option.RequestOption) (res *StreamEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "stream/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type StreamEventsResponse struct {
	Channel       string                            `json:"channel"`
	Data          []Event                           `json:"data"`
	PusherKey     string                            `json:"pusherKey"`
	PusherOptions StreamEventsResponsePusherOptions `json:"pusherOptions"`
	Success       bool                              `json:"success"`
	User          string                            `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel       respjson.Field
		Data          respjson.Field
		PusherKey     respjson.Field
		PusherOptions respjson.Field
		Success       respjson.Field
		User          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *StreamEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamEventsResponsePusherOptions struct {
	ChannelAuthorization StreamEventsResponsePusherOptionsChannelAuthorization `json:"channelAuthorization"`
	Cluster              string                                                `json:"cluster"`
	HTTPHost             string                                                `json:"httpHost"`
	HTTPPort             int64                                                 `json:"httpPort"`
	HTTPSPort            int64                                                 `json:"httpsPort"`
	WsHost               string                                                `json:"wsHost"`
	WsPort               int64                                                 `json:"wsPort"`
	WssPort              int64                                                 `json:"wssPort"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelAuthorization respjson.Field
		Cluster              respjson.Field
		HTTPHost             respjson.Field
		HTTPPort             respjson.Field
		HTTPSPort            respjson.Field
		WsHost               respjson.Field
		WsPort               respjson.Field
		WssPort              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamEventsResponsePusherOptions) RawJSON() string { return r.JSON.raw }
func (r *StreamEventsResponsePusherOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamEventsResponsePusherOptionsChannelAuthorization struct {
	Endpoint string            `json:"endpoint" format:"uri"`
	Headers  map[string]string `json:"headers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Endpoint    respjson.Field
		Headers     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamEventsResponsePusherOptionsChannelAuthorization) RawJSON() string { return r.JSON.raw }
func (r *StreamEventsResponsePusherOptionsChannelAuthorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamEventsParams struct {
	// An eventID to stream events for
	EventID param.Opt[string] `query:"eventID,omitzero" json:"-"`
	// The feed you would like to subscribe to
	Feed param.Opt[string] `query:"feed,omitzero" json:"-"`
	// A leagueID to stream events for
	LeagueID param.Opt[string] `query:"leagueID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StreamEventsParams]'s query parameters as `url.Values`.
func (r StreamEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
