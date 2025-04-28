# PtrUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewPtrUpdate

`func NewPtrUpdate() *PtrUpdate`

NewPtrUpdate instantiates a new PtrUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPtrUpdateWithDefaults

`func NewPtrUpdateWithDefaults() *PtrUpdate`

NewPtrUpdateWithDefaults instantiates a new PtrUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *PtrUpdate) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *PtrUpdate) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *PtrUpdate) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *PtrUpdate) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetTtl

`func (o *PtrUpdate) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PtrUpdate) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PtrUpdate) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PtrUpdate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetComment

`func (o *PtrUpdate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PtrUpdate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PtrUpdate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PtrUpdate) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


