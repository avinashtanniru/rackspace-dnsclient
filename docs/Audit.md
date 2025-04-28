# Audit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**After** | Pointer to **string** |  | [optional] 
**AuditId** | Pointer to **string** |  | [optional] 
**Before** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**RecordId** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **float32** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **string** |  | [optional] 

## Methods

### NewAudit

`func NewAudit() *Audit`

NewAudit instantiates a new Audit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditWithDefaults

`func NewAuditWithDefaults() *Audit`

NewAuditWithDefaults instantiates a new Audit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Audit) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Audit) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Audit) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Audit) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAfter

`func (o *Audit) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *Audit) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *Audit) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *Audit) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetAuditId

`func (o *Audit) GetAuditId() string`

GetAuditId returns the AuditId field if non-nil, zero value otherwise.

### GetAuditIdOk

`func (o *Audit) GetAuditIdOk() (*string, bool)`

GetAuditIdOk returns a tuple with the AuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditId

`func (o *Audit) SetAuditId(v string)`

SetAuditId sets AuditId field to given value.

### HasAuditId

`func (o *Audit) HasAuditId() bool`

HasAuditId returns a boolean if a field has been set.

### GetBefore

`func (o *Audit) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *Audit) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *Audit) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *Audit) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Audit) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Audit) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Audit) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Audit) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Audit) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Audit) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Audit) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Audit) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetFqdn

`func (o *Audit) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Audit) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Audit) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *Audit) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetRecordId

`func (o *Audit) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *Audit) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *Audit) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.

### HasRecordId

`func (o *Audit) HasRecordId() bool`

HasRecordId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Audit) GetSerialNumber() float32`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Audit) GetSerialNumberOk() (*float32, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Audit) SetSerialNumber(v float32)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Audit) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetTenantId

`func (o *Audit) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Audit) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Audit) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Audit) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetType

`func (o *Audit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Audit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Audit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Audit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetZoneId

`func (o *Audit) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *Audit) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *Audit) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *Audit) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


