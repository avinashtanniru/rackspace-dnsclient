/*
Cloud DNS

This is the Rackspace Cloud DNS v2 contract.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsclient

import (
	"encoding/json"
)

// checks if the PtrRecordCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PtrRecordCollection{}

// PtrRecordCollection struct for PtrRecordCollection
type PtrRecordCollection struct {
	PtrRecords []PtrRecord `json:"ptrRecords,omitempty"`
	Links []ZoneCollectionLinksInner `json:"links,omitempty"`
}

// NewPtrRecordCollection instantiates a new PtrRecordCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPtrRecordCollection() *PtrRecordCollection {
	this := PtrRecordCollection{}
	return &this
}

// NewPtrRecordCollectionWithDefaults instantiates a new PtrRecordCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPtrRecordCollectionWithDefaults() *PtrRecordCollection {
	this := PtrRecordCollection{}
	return &this
}

// GetPtrRecords returns the PtrRecords field value if set, zero value otherwise.
func (o *PtrRecordCollection) GetPtrRecords() []PtrRecord {
	if o == nil || IsNil(o.PtrRecords) {
		var ret []PtrRecord
		return ret
	}
	return o.PtrRecords
}

// GetPtrRecordsOk returns a tuple with the PtrRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtrRecordCollection) GetPtrRecordsOk() ([]PtrRecord, bool) {
	if o == nil || IsNil(o.PtrRecords) {
		return nil, false
	}
	return o.PtrRecords, true
}

// HasPtrRecords returns a boolean if a field has been set.
func (o *PtrRecordCollection) HasPtrRecords() bool {
	if o != nil && !IsNil(o.PtrRecords) {
		return true
	}

	return false
}

// SetPtrRecords gets a reference to the given []PtrRecord and assigns it to the PtrRecords field.
func (o *PtrRecordCollection) SetPtrRecords(v []PtrRecord) {
	o.PtrRecords = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PtrRecordCollection) GetLinks() []ZoneCollectionLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []ZoneCollectionLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtrRecordCollection) GetLinksOk() ([]ZoneCollectionLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PtrRecordCollection) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ZoneCollectionLinksInner and assigns it to the Links field.
func (o *PtrRecordCollection) SetLinks(v []ZoneCollectionLinksInner) {
	o.Links = v
}

func (o PtrRecordCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PtrRecordCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PtrRecords) {
		toSerialize["ptrRecords"] = o.PtrRecords
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullablePtrRecordCollection struct {
	value *PtrRecordCollection
	isSet bool
}

func (v NullablePtrRecordCollection) Get() *PtrRecordCollection {
	return v.value
}

func (v *NullablePtrRecordCollection) Set(val *PtrRecordCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePtrRecordCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePtrRecordCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePtrRecordCollection(val *PtrRecordCollection) *NullablePtrRecordCollection {
	return &NullablePtrRecordCollection{value: val, isSet: true}
}

func (v NullablePtrRecordCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePtrRecordCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


