# RecordCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]Record**](Record.md) |  | [optional] 
**Links** | Pointer to [**[]ZoneCollectionLinksInner**](ZoneCollectionLinksInner.md) |  | [optional] 

## Methods

### NewRecordCollection

`func NewRecordCollection() *RecordCollection`

NewRecordCollection instantiates a new RecordCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordCollectionWithDefaults

`func NewRecordCollectionWithDefaults() *RecordCollection`

NewRecordCollectionWithDefaults instantiates a new RecordCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *RecordCollection) GetRecords() []Record`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *RecordCollection) GetRecordsOk() (*[]Record, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *RecordCollection) SetRecords(v []Record)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *RecordCollection) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetLinks

`func (o *RecordCollection) GetLinks() []ZoneCollectionLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RecordCollection) GetLinksOk() (*[]ZoneCollectionLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RecordCollection) SetLinks(v []ZoneCollectionLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RecordCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


