# ZoneCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewZoneCreate

`func NewZoneCreate() *ZoneCreate`

NewZoneCreate instantiates a new ZoneCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCreateWithDefaults

`func NewZoneCreateWithDefaults() *ZoneCreate`

NewZoneCreateWithDefaults instantiates a new ZoneCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *ZoneCreate) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ZoneCreate) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ZoneCreate) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ZoneCreate) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetEmail

`func (o *ZoneCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ZoneCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ZoneCreate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ZoneCreate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTtl

`func (o *ZoneCreate) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ZoneCreate) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ZoneCreate) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ZoneCreate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetComment

`func (o *ZoneCreate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ZoneCreate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ZoneCreate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ZoneCreate) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


