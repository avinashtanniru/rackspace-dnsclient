# AuditCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audits** | Pointer to [**[]Audit**](Audit.md) |  | [optional] 
**Links** | Pointer to [**[]ZoneCollectionLinksInner**](ZoneCollectionLinksInner.md) |  | [optional] 

## Methods

### NewAuditCollection

`func NewAuditCollection() *AuditCollection`

NewAuditCollection instantiates a new AuditCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditCollectionWithDefaults

`func NewAuditCollectionWithDefaults() *AuditCollection`

NewAuditCollectionWithDefaults instantiates a new AuditCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudits

`func (o *AuditCollection) GetAudits() []Audit`

GetAudits returns the Audits field if non-nil, zero value otherwise.

### GetAuditsOk

`func (o *AuditCollection) GetAuditsOk() (*[]Audit, bool)`

GetAuditsOk returns a tuple with the Audits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudits

`func (o *AuditCollection) SetAudits(v []Audit)`

SetAudits sets Audits field to given value.

### HasAudits

`func (o *AuditCollection) HasAudits() bool`

HasAudits returns a boolean if a field has been set.

### GetLinks

`func (o *AuditCollection) GetLinks() []ZoneCollectionLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuditCollection) GetLinksOk() (*[]ZoneCollectionLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuditCollection) SetLinks(v []ZoneCollectionLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuditCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


