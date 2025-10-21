// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package miruserver

import (
	"time"

	"github.com/stainless-sdks/miru-server-go/internal/apijson"
	"github.com/stainless-sdks/miru-server-go/option"
	"github.com/stainless-sdks/miru-server-go/packages/respjson"
)

// WebhookService contains methods and other services that help with interacting
// with the miru API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

func (r *WebhookService) Unwrap(payload []byte, opts ...option.RequestOption) (*UnwrapWebhookEvent, error) {
	res := &UnwrapWebhookEvent{}
	err := res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

type DeploymentValidateWebhookEvent struct {
	// The data associated with the event
	Data DeploymentValidateWebhookEventData `json:"data,required"`
	// The object that occurred
	//
	// Any of "event".
	Object DeploymentValidateWebhookEventObject `json:"object,required"`
	// The timestamp of the event
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The type of event
	//
	// Any of "deployment.validate".
	Type DeploymentValidateWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentValidateWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *DeploymentValidateWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The data associated with the event
type DeploymentValidateWebhookEventData struct {
	Deployment DeploymentValidateWebhookEventDataDeployment `json:"deployment,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deployment  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentValidateWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *DeploymentValidateWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentValidateWebhookEventDataDeployment struct {
	// ID of the deployment
	ID string `json:"id,required"`
	// Timestamp of when the device release was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// ID of the device
	DeviceID string `json:"device_id,required"`
	// Any of "deployment".
	Object string `json:"object,required"`
	// The version of the release
	ReleaseID string `json:"release_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DeviceID    respjson.Field
		Object      respjson.Field
		ReleaseID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentValidateWebhookEventDataDeployment) RawJSON() string { return r.JSON.raw }
func (r *DeploymentValidateWebhookEventDataDeployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object that occurred
type DeploymentValidateWebhookEventObject string

const (
	DeploymentValidateWebhookEventObjectEvent DeploymentValidateWebhookEventObject = "event"
)

// The type of event
type DeploymentValidateWebhookEventType string

const (
	DeploymentValidateWebhookEventTypeDeploymentValidate DeploymentValidateWebhookEventType = "deployment.validate"
)

type UnwrapWebhookEvent struct {
	// The data associated with the event
	Data UnwrapWebhookEventData `json:"data,required"`
	// The object that occurred
	//
	// Any of "event".
	Object UnwrapWebhookEventObject `json:"object,required"`
	// The timestamp of the event
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The type of event
	//
	// Any of "deployment.validate".
	Type UnwrapWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The data associated with the event
type UnwrapWebhookEventData struct {
	Deployment UnwrapWebhookEventDataDeployment `json:"deployment,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deployment  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnwrapWebhookEventDataDeployment struct {
	// ID of the deployment
	ID string `json:"id,required"`
	// Timestamp of when the device release was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// ID of the device
	DeviceID string `json:"device_id,required"`
	// Any of "deployment".
	Object string `json:"object,required"`
	// The version of the release
	ReleaseID string `json:"release_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DeviceID    respjson.Field
		Object      respjson.Field
		ReleaseID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataDeployment) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataDeployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object that occurred
type UnwrapWebhookEventObject string

const (
	UnwrapWebhookEventObjectEvent UnwrapWebhookEventObject = "event"
)

// The type of event
type UnwrapWebhookEventType string

const (
	UnwrapWebhookEventTypeDeploymentValidate UnwrapWebhookEventType = "deployment.validate"
)
