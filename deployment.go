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

// DeploymentService contains methods and other services that help with interacting
// with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeploymentService] method instead.
type DeploymentService struct {
	Options []option.RequestOption
}

// NewDeploymentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeploymentService(opts ...option.RequestOption) (r DeploymentService) {
	r = DeploymentService{}
	r.Options = opts
	return
}

// Stage or deploy a new deployment.
func (r *DeploymentService) New(ctx context.Context, params DeploymentNewParams, opts ...option.RequestOption) (res *Deployment, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get
func (r *DeploymentService) Get(ctx context.Context, deploymentID string, query DeploymentGetParams, opts ...option.RequestOption) (res *Deployment, err error) {
	opts = slices.Concat(r.Options, opts)
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("deployments/%s", deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List
func (r *DeploymentService) List(ctx context.Context, query DeploymentListParams, opts ...option.RequestOption) (res *DeploymentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Validate a deployment with your custom validation.
func (r *DeploymentService) Validate(ctx context.Context, deploymentID string, body DeploymentValidateParams, opts ...option.RequestOption) (res *DeploymentValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("deployments/%s/validate", deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Deployment struct {
	// ID of the deployment.
	ID string `json:"id,required"`
	// Last known activity state of the deployment.
	//
	//   - Validating: the deployment's config instances are being validated with user's
	//     custom validation
	//   - Needs review: deployment needs to be reviewed before it can be deployed
	//   - Staged: is ready to be deployed
	//   - Queued: the deployment's config instances are waiting to be received by the
	//     device; will be deployed as soon as the device is online
	//   - Deployed: the deployment's config instances are currently available for
	//     consumption on the device
	//   - Removing: the deployment's config instances are being removed from the device
	//   - Archived: the deployment is available for historical reference but cannot be
	//     deployed and is not active on the device
	//
	// Any of "validating", "needs_review", "staged", "queued", "deployed", "removing",
	// "archived".
	ActivityStatus DeploymentActivityStatus `json:"activity_status,required"`
	// Expand the config instances using 'expand[]=config_instances' in the query
	// string.
	ConfigInstances []ConfigInstance `json:"config_instances,required"`
	// Timestamp of when the device release was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description of the deployment.
	Description string `json:"description,required"`
	// Expand the device using 'expand[]=device' in the query string.
	Device Device `json:"device,required"`
	// ID of the device.
	DeviceID string `json:"device_id,required"`
	// Last known error state of the deployment.
	//
	//   - None: no errors
	//   - Retrying: an error has been encountered and the agent is retrying to reach the
	//     target status
	//   - Failed: a fatal error has been encountered; the deployment is archived and (if
	//     deployed) removed from the device
	//
	// Any of "none", "failed", "retrying".
	ErrorStatus DeploymentErrorStatus `json:"error_status,required"`
	// Any of "deployment".
	Object DeploymentObject `json:"object,required"`
	// Expand the release using 'expand[]=release' in the query string.
	Release Release `json:"release,required"`
	// The version of the release.
	ReleaseID string `json:"release_id,required"`
	// This status merges the 'activity_status' and 'error_status' fields, with error
	// states taking precedence over activity states when errors are present. For
	// example, if the activity status is 'deployed' but the error status is 'failed',
	// the status is 'failed'. However, if the error status is 'none' and the activity
	// status is 'deployed', the status is 'deployed'.
	//
	// Any of "validating", "needs_review", "staged", "queued", "deployed", "removing",
	// "archived", "failed", "retrying".
	Status DeploymentStatus `json:"status,required"`
	// Desired state of the deployment.
	//
	//   - Staged: is ready to be deployed
	//   - Deployed: all config instances part of the deployment are available for
	//     consumption on the device
	//   - Archived: the deployment is available for historical reference but cannot be
	//     deployed and is not active on the device
	//
	// Any of "staged", "deployed", "archived".
	TargetStatus DeploymentTargetStatus `json:"target_status,required"`
	// Timestamp of when the device release was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ActivityStatus  respjson.Field
		ConfigInstances respjson.Field
		CreatedAt       respjson.Field
		Description     respjson.Field
		Device          respjson.Field
		DeviceID        respjson.Field
		ErrorStatus     respjson.Field
		Object          respjson.Field
		Release         respjson.Field
		ReleaseID       respjson.Field
		Status          respjson.Field
		TargetStatus    respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Deployment) RawJSON() string { return r.JSON.raw }
func (r *Deployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Last known activity state of the deployment.
//
//   - Validating: the deployment's config instances are being validated with user's
//     custom validation
//   - Needs review: deployment needs to be reviewed before it can be deployed
//   - Staged: is ready to be deployed
//   - Queued: the deployment's config instances are waiting to be received by the
//     device; will be deployed as soon as the device is online
//   - Deployed: the deployment's config instances are currently available for
//     consumption on the device
//   - Removing: the deployment's config instances are being removed from the device
//   - Archived: the deployment is available for historical reference but cannot be
//     deployed and is not active on the device
type DeploymentActivityStatus string

const (
	DeploymentActivityStatusValidating  DeploymentActivityStatus = "validating"
	DeploymentActivityStatusNeedsReview DeploymentActivityStatus = "needs_review"
	DeploymentActivityStatusStaged      DeploymentActivityStatus = "staged"
	DeploymentActivityStatusQueued      DeploymentActivityStatus = "queued"
	DeploymentActivityStatusDeployed    DeploymentActivityStatus = "deployed"
	DeploymentActivityStatusRemoving    DeploymentActivityStatus = "removing"
	DeploymentActivityStatusArchived    DeploymentActivityStatus = "archived"
)

// Last known error state of the deployment.
//
//   - None: no errors
//   - Retrying: an error has been encountered and the agent is retrying to reach the
//     target status
//   - Failed: a fatal error has been encountered; the deployment is archived and (if
//     deployed) removed from the device
type DeploymentErrorStatus string

const (
	DeploymentErrorStatusNone     DeploymentErrorStatus = "none"
	DeploymentErrorStatusFailed   DeploymentErrorStatus = "failed"
	DeploymentErrorStatusRetrying DeploymentErrorStatus = "retrying"
)

type DeploymentObject string

const (
	DeploymentObjectDeployment DeploymentObject = "deployment"
)

// This status merges the 'activity_status' and 'error_status' fields, with error
// states taking precedence over activity states when errors are present. For
// example, if the activity status is 'deployed' but the error status is 'failed',
// the status is 'failed'. However, if the error status is 'none' and the activity
// status is 'deployed', the status is 'deployed'.
type DeploymentStatus string

const (
	DeploymentStatusValidating  DeploymentStatus = "validating"
	DeploymentStatusNeedsReview DeploymentStatus = "needs_review"
	DeploymentStatusStaged      DeploymentStatus = "staged"
	DeploymentStatusQueued      DeploymentStatus = "queued"
	DeploymentStatusDeployed    DeploymentStatus = "deployed"
	DeploymentStatusRemoving    DeploymentStatus = "removing"
	DeploymentStatusArchived    DeploymentStatus = "archived"
	DeploymentStatusFailed      DeploymentStatus = "failed"
	DeploymentStatusRetrying    DeploymentStatus = "retrying"
)

// Desired state of the deployment.
//
//   - Staged: is ready to be deployed
//   - Deployed: all config instances part of the deployment are available for
//     consumption on the device
//   - Archived: the deployment is available for historical reference but cannot be
//     deployed and is not active on the device
type DeploymentTargetStatus string

const (
	DeploymentTargetStatusStaged   DeploymentTargetStatus = "staged"
	DeploymentTargetStatusDeployed DeploymentTargetStatus = "deployed"
	DeploymentTargetStatusArchived DeploymentTargetStatus = "archived"
)

type DeploymentListResponse struct {
	Data []Deployment `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	PaginatedList
}

// Returns the unmodified JSON received from the API
func (r DeploymentListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeploymentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentValidateResponse struct {
	Deployment Deployment `json:"deployment,required"`
	// The effect of the validation.
	//
	// Any of "none", "stage", "deploy", "reject", "void".
	Effect DeploymentValidateResponseEffect `json:"effect,required"`
	// A message explaining the validation effect.
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deployment  respjson.Field
		Effect      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *DeploymentValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The effect of the validation.
type DeploymentValidateResponseEffect string

const (
	DeploymentValidateResponseEffectNone   DeploymentValidateResponseEffect = "none"
	DeploymentValidateResponseEffectStage  DeploymentValidateResponseEffect = "stage"
	DeploymentValidateResponseEffectDeploy DeploymentValidateResponseEffect = "deploy"
	DeploymentValidateResponseEffectReject DeploymentValidateResponseEffect = "reject"
	DeploymentValidateResponseEffectVoid   DeploymentValidateResponseEffect = "void"
)

type DeploymentNewParams struct {
	// The description of the deployment.
	Description string `json:"description,required"`
	// The ID of the device that the deployment is being created for.
	DeviceID string `json:"device_id,required"`
	// The _new_ config instances to create for this deployment. A deployment must have
	// exactly one config instance for each config schema in the deployment's release.
	// If less config instances are provided than the number of schemas, the deployment
	// will 'transfer' config instances from the deployment it is patched from.
	// Archived config instances cannot be transferred.
	NewConfigInstances []DeploymentNewParamsNewConfigInstance `json:"new_config_instances,omitzero,required"`
	// The release ID which this deployment adheres to.
	ReleaseID string `json:"release_id,required"`
	// Desired state of the deployment.
	//
	//   - Staged: ready for deployment. Deployments can only be staged if their release
	//     is not the current release for the device.
	//   - Deployed: deployed to the device. Deployments can only be deployed if their
	//     release is the device's current release.
	//
	// If custom validation is enabled for the release, the deployment must pass
	// validation before fulfilling the target status.
	//
	// Any of "staged", "deployed".
	TargetStatus DeploymentNewParamsTargetStatus `json:"target_status,omitzero,required"`
	// The ID of the deployment that this deployment was patched from.
	PatchSourceID param.Opt[string] `json:"patch_source_id,omitzero"`
	// The fields to expand in the deployment.
	//
	// Any of "device", "release", "config_instances".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

func (r DeploymentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DeploymentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeploymentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [DeploymentNewParams]'s query parameters as `url.Values`.
func (r DeploymentNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The properties ConfigSchemaID, Content, RelativeFilepath are required.
type DeploymentNewParamsNewConfigInstance struct {
	// The ID of the config schema which this config instance must adhere to. This
	// schema must exist in the deployment's release.
	ConfigSchemaID string `json:"config_schema_id,required"`
	// The configuration data.
	Content any `json:"content,omitzero,required"`
	// The file path to deploy the config instance relative to
	// `/srv/miru/config_instances`. `v1/motion-control.json` would deploy to
	// `/srv/miru/config_instances/v1/motion-control.json`.
	RelativeFilepath string `json:"relative_filepath,required"`
	paramObj
}

func (r DeploymentNewParamsNewConfigInstance) MarshalJSON() (data []byte, err error) {
	type shadow DeploymentNewParamsNewConfigInstance
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeploymentNewParamsNewConfigInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Desired state of the deployment.
//
//   - Staged: ready for deployment. Deployments can only be staged if their release
//     is not the current release for the device.
//   - Deployed: deployed to the device. Deployments can only be deployed if their
//     release is the device's current release.
//
// If custom validation is enabled for the release, the deployment must pass
// validation before fulfilling the target status.
type DeploymentNewParamsTargetStatus string

const (
	DeploymentNewParamsTargetStatusStaged   DeploymentNewParamsTargetStatus = "staged"
	DeploymentNewParamsTargetStatusDeployed DeploymentNewParamsTargetStatus = "deployed"
)

type DeploymentGetParams struct {
	// The fields to expand in the deployment.
	//
	// Any of "device", "release", "config_instances".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeploymentGetParams]'s query parameters as `url.Values`.
func (r DeploymentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeploymentListParams struct {
	// The deployment ID to filter by.
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// The deployment device ID to filter by.
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The deployment release ID to filter by.
	ReleaseID param.Opt[string] `query:"release_id,omitzero" json:"-"`
	// The deployment activity status to filter by.
	//
	// Any of "validating", "needs_review", "staged", "queued", "deployed", "removing",
	// "archived".
	ActivityStatus DeploymentListParamsActivityStatus `query:"activity_status,omitzero" json:"-"`
	// The deployment error status to filter by.
	//
	// Any of "none", "failed", "retrying".
	ErrorStatus DeploymentListParamsErrorStatus `query:"error_status,omitzero" json:"-"`
	// The fields to expand in the deployments list.
	//
	// Any of "total_count", "device", "release", "config_instances".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Any of "id:asc", "id:desc", "created_at:desc", "created_at:asc".
	OrderBy DeploymentListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// The deployment target status to filter by.
	//
	// Any of "staged", "deployed", "archived".
	TargetStatus DeploymentListParamsTargetStatus `query:"target_status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeploymentListParams]'s query parameters as `url.Values`.
func (r DeploymentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The deployment activity status to filter by.
type DeploymentListParamsActivityStatus string

const (
	DeploymentListParamsActivityStatusValidating  DeploymentListParamsActivityStatus = "validating"
	DeploymentListParamsActivityStatusNeedsReview DeploymentListParamsActivityStatus = "needs_review"
	DeploymentListParamsActivityStatusStaged      DeploymentListParamsActivityStatus = "staged"
	DeploymentListParamsActivityStatusQueued      DeploymentListParamsActivityStatus = "queued"
	DeploymentListParamsActivityStatusDeployed    DeploymentListParamsActivityStatus = "deployed"
	DeploymentListParamsActivityStatusRemoving    DeploymentListParamsActivityStatus = "removing"
	DeploymentListParamsActivityStatusArchived    DeploymentListParamsActivityStatus = "archived"
)

// The deployment error status to filter by.
type DeploymentListParamsErrorStatus string

const (
	DeploymentListParamsErrorStatusNone     DeploymentListParamsErrorStatus = "none"
	DeploymentListParamsErrorStatusFailed   DeploymentListParamsErrorStatus = "failed"
	DeploymentListParamsErrorStatusRetrying DeploymentListParamsErrorStatus = "retrying"
)

type DeploymentListParamsOrderBy string

const (
	DeploymentListParamsOrderByIDAsc         DeploymentListParamsOrderBy = "id:asc"
	DeploymentListParamsOrderByIDDesc        DeploymentListParamsOrderBy = "id:desc"
	DeploymentListParamsOrderByCreatedAtDesc DeploymentListParamsOrderBy = "created_at:desc"
	DeploymentListParamsOrderByCreatedAtAsc  DeploymentListParamsOrderBy = "created_at:asc"
)

// The deployment target status to filter by.
type DeploymentListParamsTargetStatus string

const (
	DeploymentListParamsTargetStatusStaged   DeploymentListParamsTargetStatus = "staged"
	DeploymentListParamsTargetStatusDeployed DeploymentListParamsTargetStatus = "deployed"
	DeploymentListParamsTargetStatusArchived DeploymentListParamsTargetStatus = "archived"
)

type DeploymentValidateParams struct {
	// The config instance errors for this deployment.
	ConfigInstances []DeploymentValidateParamsConfigInstance `json:"config_instances,omitzero,required"`
	// Whether the deployment is valid. If invalid, the deployment is immediately
	// archived and marked as 'failed'.
	IsValid bool `json:"is_valid,required"`
	// A message displayed on the deployment level in the UI.
	Message string `json:"message,required"`
	paramObj
}

func (r DeploymentValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow DeploymentValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeploymentValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The validation errors(s) for a specific config instance in the deployment.
//
// The properties ID, Message, Parameters are required.
type DeploymentValidateParamsConfigInstance struct {
	// ID of the config instance.
	ID string `json:"id,required"`
	// A message displayed on the config instance level in the UI.
	Message string `json:"message,required"`
	// The parameter errors for this config instance.
	Parameters []DeploymentValidateParamsConfigInstanceParameter `json:"parameters,omitzero,required"`
	paramObj
}

func (r DeploymentValidateParamsConfigInstance) MarshalJSON() (data []byte, err error) {
	type shadow DeploymentValidateParamsConfigInstance
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeploymentValidateParamsConfigInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The validation error for a parameter in the config instance.
//
// The properties Message, Path are required.
type DeploymentValidateParamsConfigInstanceParameter struct {
	// An error message displayed for an individual parameter.
	Message string `json:"message,required"`
	// The path to the parameter that is invalid.
	Path []string `json:"path,omitzero,required"`
	paramObj
}

func (r DeploymentValidateParamsConfigInstanceParameter) MarshalJSON() (data []byte, err error) {
	type shadow DeploymentValidateParamsConfigInstanceParameter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeploymentValidateParamsConfigInstanceParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
