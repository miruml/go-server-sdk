// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package miruserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/miru-server-go/internal/apijson"
	"github.com/stainless-sdks/miru-server-go/internal/apiquery"
	"github.com/stainless-sdks/miru-server-go/internal/requestconfig"
	"github.com/stainless-sdks/miru-server-go/option"
	"github.com/stainless-sdks/miru-server-go/packages/param"
	"github.com/stainless-sdks/miru-server-go/packages/respjson"
)

// ReleaseService contains methods and other services that help with interacting
// with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReleaseService] method instead.
type ReleaseService struct {
	Options []option.RequestOption
}

// NewReleaseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewReleaseService(opts ...option.RequestOption) (r ReleaseService) {
	r = ReleaseService{}
	r.Options = opts
	return
}

// Retrieve a release by ID.
func (r *ReleaseService) Get(ctx context.Context, releaseID string, query ReleaseGetParams, opts ...option.RequestOption) (res *Release, err error) {
	opts = slices.Concat(r.Options, opts)
	if releaseID == "" {
		err = errors.New("missing required release_id parameter")
		return
	}
	path := fmt.Sprintf("releases/%s", releaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List releases.
func (r *ReleaseService) List(ctx context.Context, query ReleaseListParams, opts ...option.RequestOption) (res *ReleaseListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "releases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Release struct {
	// ID of the release.
	ID string `json:"id,required"`
	// Expand the config schemas using 'expand[]=config_schemas' in the query string.
	ConfigSchemas []ConfigSchema `json:"config_schemas,required"`
	// Timestamp of when the release was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Any of "release".
	Object ReleaseObject `json:"object,required"`
	// Timestamp of when the release was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The version of the release.
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ConfigSchemas respjson.Field
		CreatedAt     respjson.Field
		Object        respjson.Field
		UpdatedAt     respjson.Field
		Version       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Release) RawJSON() string { return r.JSON.raw }
func (r *Release) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReleaseObject string

const (
	ReleaseObjectRelease ReleaseObject = "release"
)

type ReleaseListResponse struct {
	Data []Release `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	PaginatedList
}

// Returns the unmodified JSON received from the API
func (r ReleaseListResponse) RawJSON() string { return r.JSON.raw }
func (r *ReleaseListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReleaseGetParams struct {
	// The fields to expand in the releases.
	//
	// Any of "config_schemas".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReleaseGetParams]'s query parameters as `url.Values`.
func (r ReleaseGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReleaseListParams struct {
	// The release ID to filter by.
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The release version to filter by.
	Version param.Opt[string] `query:"version,omitzero" json:"-"`
	// The fields to expand in the releases list.
	//
	// Any of "total_count", "config_schemas".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Any of "id:asc", "id:desc", "created_at:desc", "created_at:asc".
	OrderBy ReleaseListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReleaseListParams]'s query parameters as `url.Values`.
func (r ReleaseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReleaseListParamsOrderBy string

const (
	ReleaseListParamsOrderByIDAsc         ReleaseListParamsOrderBy = "id:asc"
	ReleaseListParamsOrderByIDDesc        ReleaseListParamsOrderBy = "id:desc"
	ReleaseListParamsOrderByCreatedAtDesc ReleaseListParamsOrderBy = "created_at:desc"
	ReleaseListParamsOrderByCreatedAtAsc  ReleaseListParamsOrderBy = "created_at:asc"
)
