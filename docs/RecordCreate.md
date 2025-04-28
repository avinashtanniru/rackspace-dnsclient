# RecordCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
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

### NewRecordCreate

`func NewRecordCreate() *RecordCreate`

NewRecordCreate instantiates a new RecordCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordCreateWithDefaults

`func NewRecordCreateWithDefaults() *RecordCreate`

NewRecordCreateWithDefaults instantiates a new RecordCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecordCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RecordCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFqdn

`func (o *RecordCreate) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *RecordCreate) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *RecordCreate) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *RecordCreate) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetDestination

`func (o *RecordCreate) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RecordCreate) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RecordCreate) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *RecordCreate) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetTtl

`func (o *RecordCreate) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordCreate) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordCreate) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordCreate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetService

`func (o *RecordCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *RecordCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *RecordCreate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *RecordCreate) HasService() bool`

HasService returns a boolean if a field has been set.

### GetProtocol

`func (o *RecordCreate) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RecordCreate) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RecordCreate) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *RecordCreate) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPort

`func (o *RecordCreate) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RecordCreate) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RecordCreate) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RecordCreate) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetWeight

`func (o *RecordCreate) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RecordCreate) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RecordCreate) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *RecordCreate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPriority

`func (o *RecordCreate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RecordCreate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RecordCreate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RecordCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetText

`func (o *RecordCreate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RecordCreate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RecordCreate) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *RecordCreate) HasText() bool`

HasText returns a boolean if a field has been set.

### GetFlag

`func (o *RecordCreate) GetFlag() int32`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *RecordCreate) GetFlagOk() (*int32, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *RecordCreate) SetFlag(v int32)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *RecordCreate) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTag

`func (o *RecordCreate) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RecordCreate) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RecordCreate) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RecordCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetComment

`func (o *RecordCreate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordCreate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordCreate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordCreate) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


