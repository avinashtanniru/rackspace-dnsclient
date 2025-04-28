# ZoneCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zones** | Pointer to [**[]Zone**](Zone.md) |  | [optional] 
**Links** | Pointer to [**[]ZoneCollectionLinksInner**](ZoneCollectionLinksInner.md) |  | [optional] 

## Methods

### NewZoneCollection

`func NewZoneCollection() *ZoneCollection`

NewZoneCollection instantiates a new ZoneCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCollectionWithDefaults

`func NewZoneCollectionWithDefaults() *ZoneCollection`

NewZoneCollectionWithDefaults instantiates a new ZoneCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZones

`func (o *ZoneCollection) GetZones() []Zone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ZoneCollection) GetZonesOk() (*[]Zone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ZoneCollection) SetZones(v []Zone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ZoneCollection) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetLinks

`func (o *ZoneCollection) GetLinks() []ZoneCollectionLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ZoneCollection) GetLinksOk() (*[]ZoneCollectionLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ZoneCollection) SetLinks(v []ZoneCollectionLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ZoneCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


