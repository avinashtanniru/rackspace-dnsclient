# PtrCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | **string** |  | 
**IpAddress** | **string** |  | 
**Ttl** | Pointer to **int32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**CloudUrl** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 

## Methods

### NewPtrCreate

`func NewPtrCreate(fqdn string, ipAddress string, ) *PtrCreate`

NewPtrCreate instantiates a new PtrCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPtrCreateWithDefaults

`func NewPtrCreateWithDefaults() *PtrCreate`

NewPtrCreateWithDefaults instantiates a new PtrCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *PtrCreate) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *PtrCreate) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *PtrCreate) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetIpAddress

`func (o *PtrCreate) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PtrCreate) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PtrCreate) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetTtl

`func (o *PtrCreate) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PtrCreate) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PtrCreate) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PtrCreate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetComment

`func (o *PtrCreate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PtrCreate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PtrCreate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PtrCreate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCloudUrl

`func (o *PtrCreate) GetCloudUrl() string`

GetCloudUrl returns the CloudUrl field if non-nil, zero value otherwise.

### GetCloudUrlOk

`func (o *PtrCreate) GetCloudUrlOk() (*string, bool)`

GetCloudUrlOk returns a tuple with the CloudUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudUrl

`func (o *PtrCreate) SetCloudUrl(v string)`

SetCloudUrl sets CloudUrl field to given value.

### HasCloudUrl

`func (o *PtrCreate) HasCloudUrl() bool`

HasCloudUrl returns a boolean if a field has been set.

### GetServiceName

`func (o *PtrCreate) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *PtrCreate) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *PtrCreate) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *PtrCreate) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


