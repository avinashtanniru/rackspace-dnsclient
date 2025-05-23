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

// checks if the RecordUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordUpdate{}

// RecordUpdate struct for RecordUpdate
type RecordUpdate struct {
	Fqdn *string `json:"fqdn,omitempty"`
	Destination *string `json:"destination,omitempty"`
	Ttl *int32 `json:"ttl,omitempty"`
	Service *string `json:"service,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	Text *string `json:"text,omitempty"`
	Flag *int32 `json:"flag,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Comment *string `json:"comment,omitempty"`
}

// NewRecordUpdate instantiates a new RecordUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordUpdate() *RecordUpdate {
	this := RecordUpdate{}
	return &this
}

// NewRecordUpdateWithDefaults instantiates a new RecordUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordUpdateWithDefaults() *RecordUpdate {
	this := RecordUpdate{}
	return &this
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *RecordUpdate) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordUpdate) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *RecordUpdate) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *RecordUpdate) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *RecordUpdate) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordUpdate) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *RecordUpdate) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *RecordUpdate) SetDestination(v string) {
	o.Destination = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *RecordUpdate) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordUpdate) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *RecordUpdate) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *RecordUpdate) SetTtl(v int32) {
	o.Ttl = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *RecordUpdate) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordUpdate) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *RecordUpdate) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *RecordUpdate) SetService(v string) {
	o.Service = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *RecordUpdate) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordUpdate) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *RecordUpdate) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *RecordUpdate) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *RecordUpdate) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordUpdate) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *RecordUpdate) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *RecordUpdate) SetPort(v int32) {
	o.Port = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *RecordUpdate) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordUpdate) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *RecordUpdate) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *RecordUpdate) SetWeight(v int32) {
	o.Weight = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *RecordUpdate) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordUpdate) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *RecordUpdate) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *RecordUpdate) SetPriority(v int32) {
	o.Priority = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *RecordUpdate) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordUpdate) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *RecordUpdate) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *RecordUpdate) SetText(v string) {
	o.Text = &v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *RecordUpdate) GetFlag() int32 {
	if o == nil || IsNil(o.Flag) {
		var ret int32
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordUpdate) GetFlagOk() (*int32, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *RecordUpdate) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given int32 and assigns it to the Flag field.
func (o *RecordUpdate) SetFlag(v int32) {
	o.Flag = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *RecordUpdate) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordUpdate) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *RecordUpdate) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *RecordUpdate) SetTag(v string) {
	o.Tag = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RecordUpdate) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordUpdate) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RecordUpdate) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RecordUpdate) SetComment(v string) {
	o.Comment = &v
}

func (o RecordUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Flag) {
		toSerialize["flag"] = o.Flag
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableRecordUpdate struct {
	value *RecordUpdate
	isSet bool
}

func (v NullableRecordUpdate) Get() *RecordUpdate {
	return v.value
}

func (v *NullableRecordUpdate) Set(val *RecordUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordUpdate(val *RecordUpdate) *NullableRecordUpdate {
	return &NullableRecordUpdate{value: val, isSet: true}
}

func (v NullableRecordUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


