# \RecordsAPI

All URIs are relative to *https://dns.api.rackspacecloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRecord**](RecordsAPI.md#CreateRecord) | **Post** /{tenantId}/zones/{zoneId}/records | Create a record.
[**DeleteRecord**](RecordsAPI.md#DeleteRecord) | **Delete** /{tenantId}/zones/{zoneId}/records/{recordId} | Delete a record.
[**GetRecord**](RecordsAPI.md#GetRecord) | **Get** /{tenantId}/zones/{zoneId}/records/{recordId} | Get a record by ID.
[**GetRecords**](RecordsAPI.md#GetRecords) | **Get** /{tenantId}/zones/{zoneId}/records | Get all records for a zone.
[**UpdateRecord**](RecordsAPI.md#UpdateRecord) | **Patch** /{tenantId}/zones/{zoneId}/records/{recordId} | Update a record.



## CreateRecord

> Record CreateRecord(ctx, tenantId, zoneId).RecordCreate(recordCreate).Execute()

Create a record.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/avinashtanniru/rackspace-dnsclient"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	zoneId := "zoneId_example" // string | ID of the zone
	recordCreate := *openapiclient.NewRecordCreate() // RecordCreate | A JSON object containing record data.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.CreateRecord(context.Background(), tenantId, zoneId).RecordCreate(recordCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.CreateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRecord`: Record
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.CreateRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**zoneId** | **string** | ID of the zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recordCreate** | [**RecordCreate**](RecordCreate.md) | A JSON object containing record data. | 

### Return type

[**Record**](Record.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecord

> DeleteRecord(ctx, tenantId, zoneId, recordId).Execute()

Delete a record.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/avinashtanniru/rackspace-dnsclient"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	zoneId := "zoneId_example" // string | ID of the zone
	recordId := "recordId_example" // string | ID of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecordsAPI.DeleteRecord(context.Background(), tenantId, zoneId, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.DeleteRecord``: %v\n", err)
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
**recordId** | **string** | ID of the record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordRequest struct via the builder pattern


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


## GetRecord

> Record GetRecord(ctx, tenantId, zoneId, recordId).Execute()

Get a record by ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/avinashtanniru/rackspace-dnsclient"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	zoneId := "zoneId_example" // string | ID of the zone
	recordId := "recordId_example" // string | ID of the record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.GetRecord(context.Background(), tenantId, zoneId, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecord`: Record
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**zoneId** | **string** | ID of the zone | 
**recordId** | **string** | ID of the record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Record**](Record.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecords

> RecordCollection GetRecords(ctx, tenantId, zoneId).Type_(type_).Execute()

Get all records for a zone.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/avinashtanniru/rackspace-dnsclient"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	zoneId := "zoneId_example" // string | ID of the zone
	type_ := "type__example" // string | record type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.GetRecords(context.Background(), tenantId, zoneId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecords`: RecordCollection
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**zoneId** | **string** | ID of the zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** | record type | 

### Return type

[**RecordCollection**](RecordCollection.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecord

> Record UpdateRecord(ctx, tenantId, zoneId, recordId).RecordUpdate(recordUpdate).Execute()

Update a record.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/avinashtanniru/rackspace-dnsclient"
)

func main() {
	tenantId := "tenantId_example" // string | ID of the tenant
	zoneId := "zoneId_example" // string | ID of the zone
	recordId := "recordId_example" // string | ID of the record
	recordUpdate := *openapiclient.NewRecordUpdate() // RecordUpdate | A JSON object containing record data.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.UpdateRecord(context.Background(), tenantId, zoneId, recordId).RecordUpdate(recordUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.UpdateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRecord`: Record
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.UpdateRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**zoneId** | **string** | ID of the zone | 
**recordId** | **string** | ID of the record | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **recordUpdate** | [**RecordUpdate**](RecordUpdate.md) | A JSON object containing record data. | 

### Return type

[**Record**](Record.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

