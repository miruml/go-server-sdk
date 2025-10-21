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

	"github.com/miruml/go-server-sdk/internal/apijson"
	"github.com/miruml/go-server-sdk/internal/apiquery"
	"github.com/miruml/go-server-sdk/internal/requestconfig"
	"github.com/miruml/go-server-sdk/option"
	"github.com/miruml/go-server-sdk/packages/param"
	"github.com/miruml/go-server-sdk/packages/respjson"
)

// ConfigInstanceService contains methods and other services that help with
// interacting with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigInstanceService] method instead.
type ConfigInstanceService struct {
	Options []option.RequestOption
}

// NewConfigInstanceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigInstanceService(opts ...option.RequestOption) (r ConfigInstanceService) {
	r = ConfigInstanceService{}
	r.Options = opts
	return
}

// Retrieve a config instance by ID.
func (r *ConfigInstanceService) Get(ctx context.Context, configInstanceID string, query ConfigInstanceGetParams, opts ...option.RequestOption) (res *ConfigInstance, err error) {
	opts = slices.Concat(r.Options, opts)
	if configInstanceID == "" {
		err = errors.New("missing required config_instance_id parameter")
		return
	}
	path := fmt.Sprintf("config_instances/%s", configInstanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List config instances.
func (r *ConfigInstanceService) List(ctx context.Context, query ConfigInstanceListParams, opts ...option.RequestOption) (res *ConfigInstanceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config_instances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ConfigInstance struct {
	// ID of the config instance.
	ID string `json:"id,required"`
	// Last known activity state of the config instance.
	//
	//   - Created: config instance has been created and can be deployed in the future
	//   - Queued: config instance is waiting to be received by the device; will be
	//     deployed as soon as the device is online
	//   - Deployed: config instance is currently available for consumption on the device
	//   - Removed: the config instance is available for historical reference but cannot
	//     be deployed and is not active on the device
	//
	// Any of "created", "queued", "deployed", "removed".
	ActivityStatus ConfigInstanceActivityStatus `json:"activity_status,required"`
	// Expand the config schema using 'expand[]=config_schema' in the query string.
	ConfigSchema ConfigSchema `json:"config_schema,required"`
	// ID of the config schema which the config instance must adhere to.
	ConfigSchemaID string `json:"config_schema_id,required"`
	// ID of the config type which the config instance (and its schema) is a part of.
	ConfigTypeID string `json:"config_type_id,required"`
	// The configuration values associated with the config instance. Expand the content
	// using 'expand[]=content' in the query string.
	Content any `json:"content,required"`
	// The timestamp of when the config instance was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Device    Device    `json:"device,required"`
	// ID of the device which the config instance is deployed to.
	DeviceID string `json:"device_id,required"`
	// Last known error state of the config instance deployment.
	//
	//   - None: there are no errors with the config instance deployment
	//   - Retrying: an error has been encountered and the agent is attempting to try
	//     again to reach the target status
	//   - Failed: a fatal error has been encountered; the config instance is archived
	//     and (if deployed) removed from the device
	//
	// Any of "none", "failed", "retrying".
	ErrorStatus ConfigInstanceErrorStatus `json:"error_status,required"`
	// Any of "config_instance".
	Object ConfigInstanceObject `json:"object,required"`
	// The file path to deploy the config instance relative to
	// `/srv/miru/config_instances`. `v1/motion-control.json` would deploy to
	// `/srv/miru/config_instances/v1/motion-control.json`.
	RelativeFilepath string `json:"relative_filepath,required"`
	// This status merges the 'activity_status' and 'error_status' fields, with error
	// states taking precedence over activity states when errors are present. For
	// example, if the activity status is 'deployed' but the error status is 'failed',
	// the status is 'failed'. However, if the error status is 'none' and the activity
	// status is 'deployed', the status is 'deployed'.
	//
	// Any of "created", "queued", "deployed", "removed", "failed", "retrying".
	Status ConfigInstanceStatus `json:"status,required"`
	// Desired state of the config instance.
	//
	//   - Created: config instance is created and can be deployed in the future
	//   - Deployed: config instance is available for consumption on the device
	//   - Removed: config instance is available for historical reference but cannot be
	//     deployed and is not active on the device
	//
	// Any of "created", "deployed", "removed".
	TargetStatus ConfigInstanceTargetStatus `json:"target_status,required"`
	// The timestamp of when the config instance was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Expand the config type using 'expand[]=config_type' in the query string.
	ConfigType ConfigType `json:"config_type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ActivityStatus   respjson.Field
		ConfigSchema     respjson.Field
		ConfigSchemaID   respjson.Field
		ConfigTypeID     respjson.Field
		Content          respjson.Field
		CreatedAt        respjson.Field
		Device           respjson.Field
		DeviceID         respjson.Field
		ErrorStatus      respjson.Field
		Object           respjson.Field
		RelativeFilepath respjson.Field
		Status           respjson.Field
		TargetStatus     respjson.Field
		UpdatedAt        respjson.Field
		ConfigType       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigInstance) RawJSON() string { return r.JSON.raw }
func (r *ConfigInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Last known activity state of the config instance.
//
//   - Created: config instance has been created and can be deployed in the future
//   - Queued: config instance is waiting to be received by the device; will be
//     deployed as soon as the device is online
//   - Deployed: config instance is currently available for consumption on the device
//   - Removed: the config instance is available for historical reference but cannot
//     be deployed and is not active on the device
type ConfigInstanceActivityStatus string

const (
	ConfigInstanceActivityStatusCreated  ConfigInstanceActivityStatus = "created"
	ConfigInstanceActivityStatusQueued   ConfigInstanceActivityStatus = "queued"
	ConfigInstanceActivityStatusDeployed ConfigInstanceActivityStatus = "deployed"
	ConfigInstanceActivityStatusRemoved  ConfigInstanceActivityStatus = "removed"
)

// Last known error state of the config instance deployment.
//
//   - None: there are no errors with the config instance deployment
//   - Retrying: an error has been encountered and the agent is attempting to try
//     again to reach the target status
//   - Failed: a fatal error has been encountered; the config instance is archived
//     and (if deployed) removed from the device
type ConfigInstanceErrorStatus string

const (
	ConfigInstanceErrorStatusNone     ConfigInstanceErrorStatus = "none"
	ConfigInstanceErrorStatusFailed   ConfigInstanceErrorStatus = "failed"
	ConfigInstanceErrorStatusRetrying ConfigInstanceErrorStatus = "retrying"
)

type ConfigInstanceObject string

const (
	ConfigInstanceObjectConfigInstance ConfigInstanceObject = "config_instance"
)

// This status merges the 'activity_status' and 'error_status' fields, with error
// states taking precedence over activity states when errors are present. For
// example, if the activity status is 'deployed' but the error status is 'failed',
// the status is 'failed'. However, if the error status is 'none' and the activity
// status is 'deployed', the status is 'deployed'.
type ConfigInstanceStatus string

const (
	ConfigInstanceStatusCreated  ConfigInstanceStatus = "created"
	ConfigInstanceStatusQueued   ConfigInstanceStatus = "queued"
	ConfigInstanceStatusDeployed ConfigInstanceStatus = "deployed"
	ConfigInstanceStatusRemoved  ConfigInstanceStatus = "removed"
	ConfigInstanceStatusFailed   ConfigInstanceStatus = "failed"
	ConfigInstanceStatusRetrying ConfigInstanceStatus = "retrying"
)

// Desired state of the config instance.
//
//   - Created: config instance is created and can be deployed in the future
//   - Deployed: config instance is available for consumption on the device
//   - Removed: config instance is available for historical reference but cannot be
//     deployed and is not active on the device
type ConfigInstanceTargetStatus string

const (
	ConfigInstanceTargetStatusCreated  ConfigInstanceTargetStatus = "created"
	ConfigInstanceTargetStatusDeployed ConfigInstanceTargetStatus = "deployed"
	ConfigInstanceTargetStatusRemoved  ConfigInstanceTargetStatus = "removed"
)

type ConfigSchema struct {
	// ID of the config schema.
	ID string `json:"id,required"`
	// Expand the config type using 'expand[]=config_type' in the query string.
	ConfigType ConfigType `json:"config_type,required"`
	// ID of the config type.
	ConfigTypeID string `json:"config_type_id,required"`
	// The config schema's JSON Schema definition.
	Content any `json:"content,required"`
	// Timestamp of when the config schema was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Hash of the config schema contents.
	Digest string `json:"digest,required"`
	// Any of "config_schema".
	Object ConfigSchemaObject `json:"object,required"`
	// The file path to deploy the config instance relative to
	// `/srv/miru/config_instances`. `v1/motion-control.json` would deploy to
	// `/srv/miru/config_instances/v1/motion-control.json`.
	RelativeFilepath string `json:"relative_filepath,required"`
	// Timestamp of when the config schema was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Config schema version for the config type.
	Version int64 `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ConfigType       respjson.Field
		ConfigTypeID     respjson.Field
		Content          respjson.Field
		CreatedAt        respjson.Field
		Digest           respjson.Field
		Object           respjson.Field
		RelativeFilepath respjson.Field
		UpdatedAt        respjson.Field
		Version          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigSchema) RawJSON() string { return r.JSON.raw }
func (r *ConfigSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigSchemaObject string

const (
	ConfigSchemaObjectConfigSchema ConfigSchemaObject = "config_schema"
)

type ConfigType struct {
	// ID of the config type.
	ID string `json:"id,required"`
	// Timestamp of when the config type was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Name of the config type.
	Name string `json:"name,required"`
	// Any of "config_type".
	Object ConfigTypeObject `json:"object,required"`
	// An immutable, code-friendly name for the config type.
	Slug string `json:"slug,required"`
	// Timestamp of when the config type was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Slug        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigType) RawJSON() string { return r.JSON.raw }
func (r *ConfigType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigTypeObject string

const (
	ConfigTypeObjectConfigType ConfigTypeObject = "config_type"
)

type PaginatedList struct {
	// True if there are more items in the list to return. False if there are no more
	// items to return.
	HasMore bool `json:"has_more,required"`
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit int64 `json:"limit,required"`
	// Any of "list".
	Object PaginatedListObject `json:"object,required"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset int64 `json:"offset,required"`
	// The total number of items in the list. By default the total count is not
	// returned. The total count must be expanded (using expand[]=total_count) to get
	// the total number of items in the list.
	TotalCount int64 `json:"total_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Limit       respjson.Field
		Object      respjson.Field
		Offset      respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginatedList) RawJSON() string { return r.JSON.raw }
func (r *PaginatedList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaginatedListObject string

const (
	PaginatedListObjectList PaginatedListObject = "list"
)

type ConfigInstanceListResponse struct {
	Data []ConfigInstance `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	PaginatedList
}

// Returns the unmodified JSON received from the API
func (r ConfigInstanceListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConfigInstanceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigInstanceGetParams struct {
	// The fields to expand in the config instance.
	//
	// Any of "content", "config_schema", "device", "config_type".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConfigInstanceGetParams]'s query parameters as
// `url.Values`.
func (r ConfigInstanceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigInstanceListParams struct {
	// The config instance ID to filter by.
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// The config schema ID to filter by.
	ConfigSchemaID param.Opt[string] `query:"config_schema_id,omitzero" json:"-"`
	// The config type ID to filter by.
	ConfigTypeID param.Opt[string] `query:"config_type_id,omitzero" json:"-"`
	// The device ID to filter by.
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The config instance activity status to filter by.
	//
	// Any of "created", "queued", "deployed", "removed".
	ActivityStatus ConfigInstanceListParamsActivityStatus `query:"activity_status,omitzero" json:"-"`
	// The config instance error status to filter by.
	//
	// Any of "none", "failed", "retrying".
	ErrorStatus ConfigInstanceListParamsErrorStatus `query:"error_status,omitzero" json:"-"`
	// The fields to expand in the config instance list.
	//
	// Any of "total_count", "content", "config_schema", "device", "config_type".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Any of "id:asc", "id:desc", "created_at:desc", "created_at:asc".
	OrderBy ConfigInstanceListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// The config instance target status to filter by.
	//
	// Any of "created", "deployed", "removed".
	TargetStatus ConfigInstanceListParamsTargetStatus `query:"target_status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConfigInstanceListParams]'s query parameters as
// `url.Values`.
func (r ConfigInstanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The config instance activity status to filter by.
type ConfigInstanceListParamsActivityStatus string

const (
	ConfigInstanceListParamsActivityStatusCreated  ConfigInstanceListParamsActivityStatus = "created"
	ConfigInstanceListParamsActivityStatusQueued   ConfigInstanceListParamsActivityStatus = "queued"
	ConfigInstanceListParamsActivityStatusDeployed ConfigInstanceListParamsActivityStatus = "deployed"
	ConfigInstanceListParamsActivityStatusRemoved  ConfigInstanceListParamsActivityStatus = "removed"
)

// The config instance error status to filter by.
type ConfigInstanceListParamsErrorStatus string

const (
	ConfigInstanceListParamsErrorStatusNone     ConfigInstanceListParamsErrorStatus = "none"
	ConfigInstanceListParamsErrorStatusFailed   ConfigInstanceListParamsErrorStatus = "failed"
	ConfigInstanceListParamsErrorStatusRetrying ConfigInstanceListParamsErrorStatus = "retrying"
)

type ConfigInstanceListParamsOrderBy string

const (
	ConfigInstanceListParamsOrderByIDAsc         ConfigInstanceListParamsOrderBy = "id:asc"
	ConfigInstanceListParamsOrderByIDDesc        ConfigInstanceListParamsOrderBy = "id:desc"
	ConfigInstanceListParamsOrderByCreatedAtDesc ConfigInstanceListParamsOrderBy = "created_at:desc"
	ConfigInstanceListParamsOrderByCreatedAtAsc  ConfigInstanceListParamsOrderBy = "created_at:asc"
)

// The config instance target status to filter by.
type ConfigInstanceListParamsTargetStatus string

const (
	ConfigInstanceListParamsTargetStatusCreated  ConfigInstanceListParamsTargetStatus = "created"
	ConfigInstanceListParamsTargetStatusDeployed ConfigInstanceListParamsTargetStatus = "deployed"
	ConfigInstanceListParamsTargetStatusRemoved  ConfigInstanceListParamsTargetStatus = "removed"
)
