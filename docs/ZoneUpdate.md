# ZoneUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewZoneUpdate

`func NewZoneUpdate() *ZoneUpdate`

NewZoneUpdate instantiates a new ZoneUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneUpdateWithDefaults

`func NewZoneUpdateWithDefaults() *ZoneUpdate`

NewZoneUpdateWithDefaults instantiates a new ZoneUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ZoneUpdate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ZoneUpdate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ZoneUpdate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ZoneUpdate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTtl

`func (o *ZoneUpdate) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ZoneUpdate) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ZoneUpdate) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ZoneUpdate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetComment

`func (o *ZoneUpdate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ZoneUpdate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ZoneUpdate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ZoneUpdate) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


