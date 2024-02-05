/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the ConsumptionTypeCountsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsumptionTypeCountsDTO{}

// ConsumptionTypeCountsDTO Represents virtual / physical entitlement consumption counts
type ConsumptionTypeCountsDTO struct {
	Physical *int32 `json:"physical,omitempty"`
	Guest *int32 `json:"guest,omitempty"`
}

// NewConsumptionTypeCountsDTO instantiates a new ConsumptionTypeCountsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumptionTypeCountsDTO() *ConsumptionTypeCountsDTO {
	this := ConsumptionTypeCountsDTO{}
	return &this
}

// NewConsumptionTypeCountsDTOWithDefaults instantiates a new ConsumptionTypeCountsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumptionTypeCountsDTOWithDefaults() *ConsumptionTypeCountsDTO {
	this := ConsumptionTypeCountsDTO{}
	return &this
}

// GetPhysical returns the Physical field value if set, zero value otherwise.
func (o *ConsumptionTypeCountsDTO) GetPhysical() int32 {
	if o == nil || IsNil(o.Physical) {
		var ret int32
		return ret
	}
	return *o.Physical
}

// GetPhysicalOk returns a tuple with the Physical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionTypeCountsDTO) GetPhysicalOk() (*int32, bool) {
	if o == nil || IsNil(o.Physical) {
		return nil, false
	}
	return o.Physical, true
}

// HasPhysical returns a boolean if a field has been set.
func (o *ConsumptionTypeCountsDTO) HasPhysical() bool {
	if o != nil && !IsNil(o.Physical) {
		return true
	}

	return false
}

// SetPhysical gets a reference to the given int32 and assigns it to the Physical field.
func (o *ConsumptionTypeCountsDTO) SetPhysical(v int32) {
	o.Physical = &v
}

// GetGuest returns the Guest field value if set, zero value otherwise.
func (o *ConsumptionTypeCountsDTO) GetGuest() int32 {
	if o == nil || IsNil(o.Guest) {
		var ret int32
		return ret
	}
	return *o.Guest
}

// GetGuestOk returns a tuple with the Guest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionTypeCountsDTO) GetGuestOk() (*int32, bool) {
	if o == nil || IsNil(o.Guest) {
		return nil, false
	}
	return o.Guest, true
}

// HasGuest returns a boolean if a field has been set.
func (o *ConsumptionTypeCountsDTO) HasGuest() bool {
	if o != nil && !IsNil(o.Guest) {
		return true
	}

	return false
}

// SetGuest gets a reference to the given int32 and assigns it to the Guest field.
func (o *ConsumptionTypeCountsDTO) SetGuest(v int32) {
	o.Guest = &v
}

func (o ConsumptionTypeCountsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsumptionTypeCountsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Physical) {
		toSerialize["physical"] = o.Physical
	}
	if !IsNil(o.Guest) {
		toSerialize["guest"] = o.Guest
	}
	return toSerialize, nil
}

type NullableConsumptionTypeCountsDTO struct {
	value *ConsumptionTypeCountsDTO
	isSet bool
}

func (v NullableConsumptionTypeCountsDTO) Get() *ConsumptionTypeCountsDTO {
	return v.value
}

func (v *NullableConsumptionTypeCountsDTO) Set(val *ConsumptionTypeCountsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumptionTypeCountsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumptionTypeCountsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumptionTypeCountsDTO(val *ConsumptionTypeCountsDTO) *NullableConsumptionTypeCountsDTO {
	return &NullableConsumptionTypeCountsDTO{value: val, isSet: true}
}

func (v NullableConsumptionTypeCountsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumptionTypeCountsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


