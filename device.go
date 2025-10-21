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

// DeviceService contains methods and other services that help with interacting
// with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceService] method instead.
type DeviceService struct {
	Options []option.RequestOption
}

// NewDeviceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDeviceService(opts ...option.RequestOption) (r DeviceService) {
	r = DeviceService{}
	r.Options = opts
	return
}

// Create a new device.
func (r *DeviceService) New(ctx context.Context, body DeviceNewParams, opts ...option.RequestOption) (res *Device, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a device by ID.
func (r *DeviceService) Get(ctx context.Context, deviceID string, opts ...option.RequestOption) (res *Device, err error) {
	opts = slices.Concat(r.Options, opts)
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("devices/%s", deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a device by ID.
func (r *DeviceService) Update(ctx context.Context, deviceID string, body DeviceUpdateParams, opts ...option.RequestOption) (res *Device, err error) {
	opts = slices.Concat(r.Options, opts)
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("devices/%s", deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List devices.
func (r *DeviceService) List(ctx context.Context, query DeviceListParams, opts ...option.RequestOption) (res *DeviceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a device by ID.
func (r *DeviceService) Delete(ctx context.Context, deviceID string, opts ...option.RequestOption) (res *DeviceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("devices/%s", deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new device activation token.
func (r *DeviceService) NewActivationToken(ctx context.Context, deviceID string, body DeviceNewActivationTokenParams, opts ...option.RequestOption) (res *DeviceNewActivationTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("devices/%s/activation_token", deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Device struct {
	// ID of the device.
	ID string `json:"id,required"`
	// Timestamp of when the device was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Timestamp of when the device was last made an initial connection (this is not
	// the same as the last time the device was seen).
	LastConnectedAt time.Time `json:"last_connected_at,required" format:"date-time"`
	// Timestamp of when the device was last disconnected (this is not the same as the
	// last time the device was seen).
	LastDisconnectedAt time.Time `json:"last_disconnected_at,required" format:"date-time"`
	// Name of the device.
	Name string `json:"name,required"`
	// Any of "device".
	Object DeviceObject `json:"object,required"`
	// The status of the device.
	//
	//   - Inactive: The miru agent has not yet been installed / authenticated
	//   - Activating: The miru agent is currently being installed / authenticated
	//     (should only last for a few seconds)
	//   - Online: The miru agent has successfully pinged the server within the last 60
	//     seconds.
	//   - Offline: The miru agent has not successfully pinged the server within the last
	//     60 seconds (e.g. network issues, device is powered off, etc.)
	//
	// Any of "inactive", "activating", "online", "offline".
	Status DeviceStatus `json:"status,required"`
	// Timestamp of when the device was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		LastConnectedAt    respjson.Field
		LastDisconnectedAt respjson.Field
		Name               respjson.Field
		Object             respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Device) RawJSON() string { return r.JSON.raw }
func (r *Device) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceObject string

const (
	DeviceObjectDevice DeviceObject = "device"
)

// The status of the device.
//
//   - Inactive: The miru agent has not yet been installed / authenticated
//   - Activating: The miru agent is currently being installed / authenticated
//     (should only last for a few seconds)
//   - Online: The miru agent has successfully pinged the server within the last 60
//     seconds.
//   - Offline: The miru agent has not successfully pinged the server within the last
//     60 seconds (e.g. network issues, device is powered off, etc.)
type DeviceStatus string

const (
	DeviceStatusInactive   DeviceStatus = "inactive"
	DeviceStatusActivating DeviceStatus = "activating"
	DeviceStatusOnline     DeviceStatus = "online"
	DeviceStatusOffline    DeviceStatus = "offline"
)

type DeviceListResponse struct {
	Data []Device `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	PaginatedList
}

// Returns the unmodified JSON received from the API
func (r DeviceListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDeleteResponse struct {
	// The ID of the device.
	ID string `json:"id,required"`
	// Whether the device was deleted.
	Deleted bool `json:"deleted,required"`
	// Any of "device".
	Object DeviceDeleteResponseObject `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceDeleteResponseObject string

const (
	DeviceDeleteResponseObjectDevice DeviceDeleteResponseObject = "device"
)

type DeviceNewActivationTokenResponse struct {
	// The token.
	Token string `json:"token,required"`
	// The expiration date and time of the token.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExpiresAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceNewActivationTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceNewActivationTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNewParams struct {
	// The name of the device.
	Name string `json:"name,required"`
	paramObj
}

func (r DeviceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceUpdateParams struct {
	// The new name of the device. If excluded from the request, the device name is not
	// updated.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r DeviceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceListParams struct {
	// The device ID to filter by.
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// The agent version to filter by.
	AgentVersion param.Opt[string] `query:"agent_version,omitzero" json:"-"`
	// The current release ID to filter by.
	CurrentReleaseID param.Opt[string] `query:"current_release_id,omitzero" json:"-"`
	// The maximum number of items to return. A limit of 15 with an offset of 0 returns
	// items 1-15.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The device name to filter by.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// The offset of the items to return. An offset of 10 with a limit of 10 returns
	// items 11-20.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// The fields to expand in the device list.
	//
	// Any of "total_count".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Any of "id:asc", "id:desc", "created_at:desc", "created_at:asc".
	OrderBy DeviceListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceListParams]'s query parameters as `url.Values`.
func (r DeviceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceListParamsOrderBy string

const (
	DeviceListParamsOrderByIDAsc         DeviceListParamsOrderBy = "id:asc"
	DeviceListParamsOrderByIDDesc        DeviceListParamsOrderBy = "id:desc"
	DeviceListParamsOrderByCreatedAtDesc DeviceListParamsOrderBy = "created_at:desc"
	DeviceListParamsOrderByCreatedAtAsc  DeviceListParamsOrderBy = "created_at:asc"
)

type DeviceNewActivationTokenParams struct {
	// Whether this token can reactivate already activated devices. If false, the token
	// is unable to activate devices which are already activated.
	AllowReactivation param.Opt[bool] `json:"allow_reactivation,omitzero"`
	paramObj
}

func (r DeviceNewActivationTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceNewActivationTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceNewActivationTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
