# \PtrAPI

All URIs are relative to *https://dns.api.rackspacecloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePtrRecord**](PtrAPI.md#CreatePtrRecord) | **Post** /{tenantId}/ptr | Create a PTR record.
[**DeletePtrRecord**](PtrAPI.md#DeletePtrRecord) | **Delete** /{tenantId}/ptr/{ptrId} | Delete a PTR record.
[**GetPtrRecord**](PtrAPI.md#GetPtrRecord) | **Get** /{tenantId}/ptr/{ptrId} | Get a PTR record by ID.
[**GetPtrRecords**](PtrAPI.md#GetPtrRecords) | **Get** /{tenantId}/ptr | Get all PTR records for a tenant.
[**UpdatePtrRecord**](PtrAPI.md#UpdatePtrRecord) | **Patch** /{tenantId}/ptr/{ptrId} | Update a PTR record.



## CreatePtrRecord

> PtrRecord CreatePtrRecord(ctx, tenantId).PtrCreate(ptrCreate).Execute()

Create a PTR record.



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
	ptrCreate := *openapiclient.NewPtrCreate("example.com", "1.2.3.4") // PtrCreate | A JSON object for the creation of a PTR record.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PtrAPI.CreatePtrRecord(context.Background(), tenantId).PtrCreate(ptrCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PtrAPI.CreatePtrRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePtrRecord`: PtrRecord
	fmt.Fprintf(os.Stdout, "Response from `PtrAPI.CreatePtrRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePtrRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ptrCreate** | [**PtrCreate**](PtrCreate.md) | A JSON object for the creation of a PTR record. | 

### Return type

[**PtrRecord**](PtrRecord.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePtrRecord

> DeletePtrRecord(ctx, tenantId, ptrId).Execute()

Delete a PTR record.



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
	ptrId := "ptrId_example" // string | ID of the PTR record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PtrAPI.DeletePtrRecord(context.Background(), tenantId, ptrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PtrAPI.DeletePtrRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**ptrId** | **string** | ID of the PTR record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePtrRecordRequest struct via the builder pattern


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


## GetPtrRecord

> PtrRecord GetPtrRecord(ctx, tenantId, ptrId).Execute()

Get a PTR record by ID.



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
	ptrId := "ptrId_example" // string | ID of the PTR record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PtrAPI.GetPtrRecord(context.Background(), tenantId, ptrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PtrAPI.GetPtrRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPtrRecord`: PtrRecord
	fmt.Fprintf(os.Stdout, "Response from `PtrAPI.GetPtrRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**ptrId** | **string** | ID of the PTR record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPtrRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PtrRecord**](PtrRecord.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPtrRecords

> PtrRecordCollection GetPtrRecords(ctx, tenantId).IpAddress(ipAddress).Execute()

Get all PTR records for a tenant.



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
	ipAddress := "ipAddress_example" // string | PTR record IP address (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PtrAPI.GetPtrRecords(context.Background(), tenantId).IpAddress(ipAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PtrAPI.GetPtrRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPtrRecords`: PtrRecordCollection
	fmt.Fprintf(os.Stdout, "Response from `PtrAPI.GetPtrRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPtrRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipAddress** | **string** | PTR record IP address | 

### Return type

[**PtrRecordCollection**](PtrRecordCollection.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePtrRecord

> PtrRecord UpdatePtrRecord(ctx, tenantId, ptrId).PtrUpdate(ptrUpdate).Execute()

Update a PTR record.



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
	ptrId := "ptrId_example" // string | ID of the PTR record
	ptrUpdate := *openapiclient.NewPtrUpdate() // PtrUpdate | A JSON object for the partial update, or idempotent create or update, of PTR record data.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PtrAPI.UpdatePtrRecord(context.Background(), tenantId, ptrId).PtrUpdate(ptrUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PtrAPI.UpdatePtrRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePtrRecord`: PtrRecord
	fmt.Fprintf(os.Stdout, "Response from `PtrAPI.UpdatePtrRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**ptrId** | **string** | ID of the PTR record | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePtrRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ptrUpdate** | [**PtrUpdate**](PtrUpdate.md) | A JSON object for the partial update, or idempotent create or update, of PTR record data. | 

### Return type

[**PtrRecord**](PtrRecord.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

