/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
	"time"
)

// checks if the EnvironmentDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentDTO{}

// EnvironmentDTO Represents an environment within an org used to enable/disable/promote content in specific places.
type EnvironmentDTO struct {
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	ContentPrefix *string `json:"contentPrefix,omitempty"`
	Owner *NestedOwnerDTO `json:"owner,omitempty"`
	EnvironmentContent []EnvironmentContentDTO `json:"environmentContent,omitempty"`
}

// NewEnvironmentDTO instantiates a new EnvironmentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentDTO() *EnvironmentDTO {
	this := EnvironmentDTO{}
	return &this
}

// NewEnvironmentDTOWithDefaults instantiates a new EnvironmentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentDTOWithDefaults() *EnvironmentDTO {
	this := EnvironmentDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *EnvironmentDTO) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDTO) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *EnvironmentDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *EnvironmentDTO) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *EnvironmentDTO) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDTO) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *EnvironmentDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *EnvironmentDTO) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentDTO) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EnvironmentDTO) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDTO) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnvironmentDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EnvironmentDTO) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnvironmentDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnvironmentDTO) SetDescription(v string) {
	o.Description = &v
}

// GetContentPrefix returns the ContentPrefix field value if set, zero value otherwise.
func (o *EnvironmentDTO) GetContentPrefix() string {
	if o == nil || IsNil(o.ContentPrefix) {
		var ret string
		return ret
	}
	return *o.ContentPrefix
}

// GetContentPrefixOk returns a tuple with the ContentPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDTO) GetContentPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.ContentPrefix) {
		return nil, false
	}
	return o.ContentPrefix, true
}

// HasContentPrefix returns a boolean if a field has been set.
func (o *EnvironmentDTO) HasContentPrefix() bool {
	if o != nil && !IsNil(o.ContentPrefix) {
		return true
	}

	return false
}

// SetContentPrefix gets a reference to the given string and assigns it to the ContentPrefix field.
func (o *EnvironmentDTO) SetContentPrefix(v string) {
	o.ContentPrefix = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *EnvironmentDTO) GetOwner() NestedOwnerDTO {
	if o == nil || IsNil(o.Owner) {
		var ret NestedOwnerDTO
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDTO) GetOwnerOk() (*NestedOwnerDTO, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *EnvironmentDTO) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NestedOwnerDTO and assigns it to the Owner field.
func (o *EnvironmentDTO) SetOwner(v NestedOwnerDTO) {
	o.Owner = &v
}

// GetEnvironmentContent returns the EnvironmentContent field value if set, zero value otherwise.
func (o *EnvironmentDTO) GetEnvironmentContent() []EnvironmentContentDTO {
	if o == nil || IsNil(o.EnvironmentContent) {
		var ret []EnvironmentContentDTO
		return ret
	}
	return o.EnvironmentContent
}

// GetEnvironmentContentOk returns a tuple with the EnvironmentContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDTO) GetEnvironmentContentOk() ([]EnvironmentContentDTO, bool) {
	if o == nil || IsNil(o.EnvironmentContent) {
		return nil, false
	}
	return o.EnvironmentContent, true
}

// HasEnvironmentContent returns a boolean if a field has been set.
func (o *EnvironmentDTO) HasEnvironmentContent() bool {
	if o != nil && !IsNil(o.EnvironmentContent) {
		return true
	}

	return false
}

// SetEnvironmentContent gets a reference to the given []EnvironmentContentDTO and assigns it to the EnvironmentContent field.
func (o *EnvironmentDTO) SetEnvironmentContent(v []EnvironmentContentDTO) {
	o.EnvironmentContent = v
}

func (o EnvironmentDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ContentPrefix) {
		toSerialize["contentPrefix"] = o.ContentPrefix
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.EnvironmentContent) {
		toSerialize["environmentContent"] = o.EnvironmentContent
	}
	return toSerialize, nil
}

type NullableEnvironmentDTO struct {
	value *EnvironmentDTO
	isSet bool
}

func (v NullableEnvironmentDTO) Get() *EnvironmentDTO {
	return v.value
}

func (v *NullableEnvironmentDTO) Set(val *EnvironmentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentDTO(val *EnvironmentDTO) *NullableEnvironmentDTO {
	return &NullableEnvironmentDTO{value: val, isSet: true}
}

func (v NullableEnvironmentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


