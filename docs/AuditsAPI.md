# \AuditsAPI

All URIs are relative to *https://dns.api.rackspacecloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAudits**](AuditsAPI.md#GetAudits) | **Get** /{tenantId}/zones/{zoneId}/audits | Get all audits for a zone.



## GetAudits

> AuditCollection GetAudits(ctx, tenantId, zoneId).Month(month).Year(year).Execute()

Get all audits for a zone.



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
	month := "month_example" // string | archived audits for specific month. Must be between 01 and 12 (optional)
	year := "2020" // string | archived audits for specific year. Required if month is included. Must be current or previous year value in CCYY format (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditsAPI.GetAudits(context.Background(), tenantId, zoneId).Month(month).Year(year).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditsAPI.GetAudits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudits`: AuditCollection
	fmt.Fprintf(os.Stdout, "Response from `AuditsAPI.GetAudits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | ID of the tenant | 
**zoneId** | **string** | ID of the zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **month** | **string** | archived audits for specific month. Must be between 01 and 12 | 
 **year** | **string** | archived audits for specific year. Required if month is included. Must be current or previous year value in CCYY format | 

### Return type

[**AuditCollection**](AuditCollection.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

