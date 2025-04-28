# \ZonesAPI

All URIs are relative to *https://dns.api.rackspacecloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallImport**](ZonesAPI.md#CallImport) | **Post** /{tenantId}/zones/import | Import a zone file&#39;s contents.
[**CreateZone**](ZonesAPI.md#CreateZone) | **Post** /{tenantId}/zones | Create a zone.
[**Export**](ZonesAPI.md#Export) | **Get** /{tenantId}/zones/{zoneId}/export | Exports a single zone
[**GetZone**](ZonesAPI.md#GetZone) | **Get** /{tenantId}/zones/{zoneId} | Get a zone by ID.
[**GetZones**](ZonesAPI.md#GetZones) | **Get** /{tenantId}/zones | Get all zones for a tenant.
[**Migrate**](ZonesAPI.md#Migrate) | **Post** /{tenantId}/zones/{zoneId}/migrate | Migrates a single zone to a different tenant
[**TenantIdZonesZoneIdDelete**](ZonesAPI.md#TenantIdZonesZoneIdDelete) | **Delete** /{tenantId}/zones/{zoneId} | Delete a zone.
[**UpdateZone**](ZonesAPI.md#UpdateZone) | **Patch** /{tenantId}/zones/{zoneId} | Update a zone.



## CallImport

> Job CallImport(ctx, tenantId).Body(body).Execute()

Import a zone file's contents.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	body := "body_example" // string | A zone file in text format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.CallImport(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.CallImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallImport`: Job
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.CallImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 

### Other Parameters

Other parameters are passed through a pointer to a apiCallImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | A zone file in text format. | 

### Return type

[**Job**](Job.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: text/vnd.bind9
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateZone

> Zone CreateZone(ctx, tenantId).ZoneCreate(zoneCreate).Execute()

Create a zone.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	zoneCreate := *openapiclient.NewZoneCreate() // ZoneCreate | A JSON object containing zone data.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.CreateZone(context.Background(), tenantId).ZoneCreate(zoneCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.CreateZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateZone`: Zone
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.CreateZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zoneCreate** | [**ZoneCreate**](ZoneCreate.md) | A JSON object containing zone data. | 

### Return type

[**Zone**](Zone.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Export

> string Export(ctx, tenantId, zoneId).Execute()

Exports a single zone



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	zoneId := "zoneId_example" // string | ID of the zone

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.Export(context.Background(), tenantId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.Export``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Export`: string
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.Export`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**zoneId** | **string** | ID of the zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/vnd.bind9

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZone

> Zone GetZone(ctx, tenantId, zoneId).Execute()

Get a zone by ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	zoneId := "zoneId_example" // string | ID of the zone

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.GetZone(context.Background(), tenantId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZone`: Zone
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**zoneId** | **string** | ID of the zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Zone**](Zone.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZones

> ZoneCollection GetZones(ctx, tenantId).Fqdn(fqdn).Execute()

Get all zones for a tenant.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	fqdn := "fqdn_example" // string | filter on fqdn (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.GetZones(context.Background(), tenantId).Fqdn(fqdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZones`: ZoneCollection
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fqdn** | **string** | filter on fqdn | 

### Return type

[**ZoneCollection**](ZoneCollection.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Migrate

> Job Migrate(ctx, tenantId, zoneId).ZoneMigrate(zoneMigrate).Execute()

Migrates a single zone to a different tenant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	zoneId := "zoneId_example" // string | ID of the zone
	zoneMigrate := *openapiclient.NewZoneMigrate() // ZoneMigrate | A JSON object containing new tenant ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.Migrate(context.Background(), tenantId, zoneId).ZoneMigrate(zoneMigrate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.Migrate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Migrate`: Job
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.Migrate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**zoneId** | **string** | ID of the zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneMigrate** | [**ZoneMigrate**](ZoneMigrate.md) | A JSON object containing new tenant ID. | 

### Return type

[**Job**](Job.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantIdZonesZoneIdDelete

> TenantIdZonesZoneIdDelete(ctx, tenantId, zoneId).Execute()

Delete a zone.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	zoneId := "zoneId_example" // string | ID of the zone

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZonesAPI.TenantIdZonesZoneIdDelete(context.Background(), tenantId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.TenantIdZonesZoneIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**zoneId** | **string** | ID of the zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantIdZonesZoneIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateZone

> Zone UpdateZone(ctx, tenantId, zoneId).ZoneUpdate(zoneUpdate).Execute()

Update a zone.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	zoneId := "zoneId_example" // string | ID of the zone
	zoneUpdate := *openapiclient.NewZoneUpdate() // ZoneUpdate | A JSON object containing zone data.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.UpdateZone(context.Background(), tenantId, zoneId).ZoneUpdate(zoneUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.UpdateZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateZone`: Zone
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.UpdateZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**zoneId** | **string** | ID of the zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneUpdate** | [**ZoneUpdate**](ZoneUpdate.md) | A JSON object containing zone data. | 

### Return type

[**Zone**](Zone.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

