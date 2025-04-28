# Zone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **int32** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**MinimumTtl** | Pointer to **int32** |  | [optional] 
**Expiry** | Pointer to **int32** |  | [optional] 
**Refresh** | Pointer to **int32** |  | [optional] 
**UpdateRetry** | Pointer to **int32** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**[]ZoneLinksInner**](ZoneLinksInner.md) |  | [optional] 

## Methods

### NewZone

`func NewZone() *Zone`

NewZone instantiates a new Zone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneWithDefaults

`func NewZoneWithDefaults() *Zone`

NewZoneWithDefaults instantiates a new Zone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Zone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Zone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Zone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Zone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFqdn

`func (o *Zone) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Zone) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Zone) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *Zone) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetEmail

`func (o *Zone) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Zone) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Zone) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Zone) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTenantId

`func (o *Zone) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Zone) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Zone) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Zone) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Zone) GetSerialNumber() int32`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Zone) GetSerialNumberOk() (*int32, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Zone) SetSerialNumber(v int32)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Zone) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetTtl

`func (o *Zone) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Zone) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Zone) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Zone) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetMinimumTtl

`func (o *Zone) GetMinimumTtl() int32`

GetMinimumTtl returns the MinimumTtl field if non-nil, zero value otherwise.

### GetMinimumTtlOk

`func (o *Zone) GetMinimumTtlOk() (*int32, bool)`

GetMinimumTtlOk returns a tuple with the MinimumTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTtl

`func (o *Zone) SetMinimumTtl(v int32)`

SetMinimumTtl sets MinimumTtl field to given value.

### HasMinimumTtl

`func (o *Zone) HasMinimumTtl() bool`

HasMinimumTtl returns a boolean if a field has been set.

### GetExpiry

`func (o *Zone) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *Zone) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *Zone) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *Zone) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetRefresh

`func (o *Zone) GetRefresh() int32`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *Zone) GetRefreshOk() (*int32, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *Zone) SetRefresh(v int32)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *Zone) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetUpdateRetry

`func (o *Zone) GetUpdateRetry() int32`

GetUpdateRetry returns the UpdateRetry field if non-nil, zero value otherwise.

### GetUpdateRetryOk

`func (o *Zone) GetUpdateRetryOk() (*int32, bool)`

GetUpdateRetryOk returns a tuple with the UpdateRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRetry

`func (o *Zone) SetUpdateRetry(v int32)`

SetUpdateRetry sets UpdateRetry field to given value.

### HasUpdateRetry

`func (o *Zone) HasUpdateRetry() bool`

HasUpdateRetry returns a boolean if a field has been set.

### GetCreated

`func (o *Zone) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Zone) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Zone) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Zone) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Zone) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Zone) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Zone) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Zone) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdated

`func (o *Zone) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Zone) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Zone) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Zone) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Zone) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Zone) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Zone) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Zone) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetComment

`func (o *Zone) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Zone) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Zone) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Zone) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLinks

`func (o *Zone) GetLinks() []ZoneLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Zone) GetLinksOk() (*[]ZoneLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Zone) SetLinks(v []ZoneLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Zone) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


