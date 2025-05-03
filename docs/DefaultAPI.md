# \DefaultAPI

All URIs are relative to *https://dns.api.rackspacecloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InfoGet**](DefaultAPI.md#InfoGet) | **Get** /info | Get metadata info



## InfoGet

> AppInfo InfoGet(ctx).Execute()

Get metadata info



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.InfoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InfoGet`: AppInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.InfoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInfoGetRequest struct via the builder pattern


### Return type

[**AppInfo**](AppInfo.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

