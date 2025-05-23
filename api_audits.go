/*
Cloud DNS

This is the Rackspace Cloud DNS v2 contract.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// AuditsAPIService AuditsAPI service
type AuditsAPIService service

type ApiGetAuditsRequest struct {
	ctx context.Context
	ApiService *AuditsAPIService
	tenantId string
	zoneId string
	month *string
	year *string
}

// archived audits for specific month. Must be between 01 and 12
func (r ApiGetAuditsRequest) Month(month string) ApiGetAuditsRequest {
	r.month = &month
	return r
}

// archived audits for specific year. Required if month is included. Must be current or previous year value in CCYY format
func (r ApiGetAuditsRequest) Year(year string) ApiGetAuditsRequest {
	r.year = &year
	return r
}

func (r ApiGetAuditsRequest) Execute() (*AuditCollection, *http.Response, error) {
	return r.ApiService.GetAuditsExecute(r)
}

/*
GetAudits Get all audits for a zone.

Returns a list of audits.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantId ID of the tenant
 @param zoneId ID of the zone
 @return ApiGetAuditsRequest
*/
func (a *AuditsAPIService) GetAudits(ctx context.Context, tenantId string, zoneId string) ApiGetAuditsRequest {
	return ApiGetAuditsRequest{
		ApiService: a,
		ctx: ctx,
		tenantId: tenantId,
		zoneId: zoneId,
	}
}

// Execute executes the request
//  @return AuditCollection
func (a *AuditsAPIService) GetAuditsExecute(r ApiGetAuditsRequest) (*AuditCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditsAPIService.GetAudits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{tenantId}/zones/{zoneId}/audits"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantId"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"zoneId"+"}", url.PathEscape(parameterValueToString(r.zoneId, "zoneId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.month != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "month", r.month, "form", "")
	}
	if r.year != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "year", r.year, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["AuthToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
