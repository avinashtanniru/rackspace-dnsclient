# RecordUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Flag** | Pointer to **int32** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewRecordUpdate

`func NewRecordUpdate() *RecordUpdate`

NewRecordUpdate instantiates a new RecordUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordUpdateWithDefaults

`func NewRecordUpdateWithDefaults() *RecordUpdate`

NewRecordUpdateWithDefaults instantiates a new RecordUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *RecordUpdate) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *RecordUpdate) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *RecordUpdate) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *RecordUpdate) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetDestination

`func (o *RecordUpdate) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RecordUpdate) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RecordUpdate) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *RecordUpdate) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetTtl

`func (o *RecordUpdate) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordUpdate) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordUpdate) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordUpdate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetService

`func (o *RecordUpdate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *RecordUpdate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *RecordUpdate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *RecordUpdate) HasService() bool`

HasService returns a boolean if a field has been set.

### GetProtocol

`func (o *RecordUpdate) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RecordUpdate) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RecordUpdate) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *RecordUpdate) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPort

`func (o *RecordUpdate) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RecordUpdate) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RecordUpdate) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RecordUpdate) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetWeight

`func (o *RecordUpdate) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RecordUpdate) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RecordUpdate) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *RecordUpdate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPriority

`func (o *RecordUpdate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RecordUpdate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RecordUpdate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RecordUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetText

`func (o *RecordUpdate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RecordUpdate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RecordUpdate) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *RecordUpdate) HasText() bool`

HasText returns a boolean if a field has been set.

### GetFlag

`func (o *RecordUpdate) GetFlag() int32`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *RecordUpdate) GetFlagOk() (*int32, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *RecordUpdate) SetFlag(v int32)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *RecordUpdate) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTag

`func (o *RecordUpdate) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RecordUpdate) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RecordUpdate) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RecordUpdate) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetComment

`func (o *RecordUpdate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordUpdate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordUpdate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordUpdate) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


