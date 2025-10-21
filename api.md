# ConfigInstances

Response Types:

- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ConfigInstance">ConfigInstance</a>
- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ConfigSchema">ConfigSchema</a>
- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ConfigType">ConfigType</a>
- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#PaginatedList">PaginatedList</a>
- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ConfigInstanceListResponse">ConfigInstanceListResponse</a>

Methods:

- <code title="get /config_instances/{config_instance_id}">client.ConfigInstances.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ConfigInstanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, configInstanceID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ConfigInstanceGetParams">ConfigInstanceGetParams</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ConfigInstance">ConfigInstance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /config_instances">client.ConfigInstances.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ConfigInstanceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ConfigInstanceListParams">ConfigInstanceListParams</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ConfigInstanceListResponse">ConfigInstanceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Deployments

Response Types:

- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#Deployment">Deployment</a>
- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentListResponse">DeploymentListResponse</a>
- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentValidateResponse">DeploymentValidateResponse</a>

Methods:

- <code title="post /deployments">client.Deployments.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentNewParams">DeploymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#Deployment">Deployment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /deployments/{deployment_id}">client.Deployments.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deploymentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentGetParams">DeploymentGetParams</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#Deployment">Deployment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /deployments">client.Deployments.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentListParams">DeploymentListParams</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentListResponse">DeploymentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /deployments/{deployment_id}/validate">client.Deployments.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deploymentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentValidateParams">DeploymentValidateParams</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentValidateResponse">DeploymentValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Devices

Response Types:

- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#Device">Device</a>
- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceListResponse">DeviceListResponse</a>
- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceDeleteResponse">DeviceDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceNewActivationTokenResponse">DeviceNewActivationTokenResponse</a>

Methods:

- <code title="post /devices">client.Devices.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceNewParams">DeviceNewParams</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#Device">Device</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /devices/{device_id}">client.Devices.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deviceID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#Device">Device</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /devices/{device_id}">client.Devices.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deviceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceUpdateParams">DeviceUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#Device">Device</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /devices">client.Devices.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceListParams">DeviceListParams</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceListResponse">DeviceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /devices/{device_id}">client.Devices.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deviceID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceDeleteResponse">DeviceDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /devices/{device_id}/activation_token">client.Devices.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceService.NewActivationToken">NewActivationToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deviceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceNewActivationTokenParams">DeviceNewActivationTokenParams</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeviceNewActivationTokenResponse">DeviceNewActivationTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Releases

Response Types:

- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#Release">Release</a>
- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ReleaseListResponse">ReleaseListResponse</a>

Methods:

- <code title="get /releases/{release_id}">client.Releases.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ReleaseService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, releaseID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ReleaseGetParams">ReleaseGetParams</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#Release">Release</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /releases">client.Releases.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ReleaseService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ReleaseListParams">ReleaseListParams</a>) (<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#ReleaseListResponse">ReleaseListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#DeploymentValidateWebhookEvent">DeploymentValidateWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/miruml/go-server-sdk">miruserver</a>.<a href="https://pkg.go.dev/github.com/miruml/go-server-sdk#UnwrapWebhookEvent">UnwrapWebhookEvent</a>
