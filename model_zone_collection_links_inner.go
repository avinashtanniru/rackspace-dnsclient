/*
Cloud DNS

This is the Rackspace Cloud DNS v2 contract.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsclient

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// ZoneCollectionLinksInner - struct for ZoneCollectionLinksInner
type ZoneCollectionLinksInner struct {
	PaginationLink *PaginationLink
}

// PaginationLinkAsZoneCollectionLinksInner is a convenience function that returns PaginationLink wrapped in ZoneCollectionLinksInner
func PaginationLinkAsZoneCollectionLinksInner(v *PaginationLink) ZoneCollectionLinksInner {
	return ZoneCollectionLinksInner{
		PaginationLink: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ZoneCollectionLinksInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PaginationLink
	err = newStrictDecoder(data).Decode(&dst.PaginationLink)
	if err == nil {
		jsonPaginationLink, _ := json.Marshal(dst.PaginationLink)
		if string(jsonPaginationLink) == "{}" { // empty struct
			dst.PaginationLink = nil
		} else {
			if err = validator.Validate(dst.PaginationLink); err != nil {
				dst.PaginationLink = nil
			} else {
				match++
			}
		}
	} else {
		dst.PaginationLink = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PaginationLink = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ZoneCollectionLinksInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ZoneCollectionLinksInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ZoneCollectionLinksInner) MarshalJSON() ([]byte, error) {
	if src.PaginationLink != nil {
		return json.Marshal(&src.PaginationLink)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ZoneCollectionLinksInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PaginationLink != nil {
		return obj.PaginationLink
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ZoneCollectionLinksInner) GetActualInstanceValue() (interface{}) {
	if obj.PaginationLink != nil {
		return *obj.PaginationLink
	}

	// all schemas are nil
	return nil
}

type NullableZoneCollectionLinksInner struct {
	value *ZoneCollectionLinksInner
	isSet bool
}

func (v NullableZoneCollectionLinksInner) Get() *ZoneCollectionLinksInner {
	return v.value
}

func (v *NullableZoneCollectionLinksInner) Set(val *ZoneCollectionLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneCollectionLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneCollectionLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneCollectionLinksInner(val *ZoneCollectionLinksInner) *NullableZoneCollectionLinksInner {
	return &NullableZoneCollectionLinksInner{value: val, isSet: true}
}

func (v NullableZoneCollectionLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneCollectionLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


