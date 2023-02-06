/*
Metric Ruleset API

Metric ruleset API 

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metric_ruleset

import (
	"encoding/json"
)

// PropertyFilter Single filter to include or exclude metric specified in the  matcher 
type PropertyFilter struct {
	// Custom property or dimension for which you want to filter 
	Property *string `json:"property,omitempty"`
	// Custom property or dimension values to filter 
	PropertyValue []string `json:"propertyValue,omitempty"`
	// Indicates whether you want the property to be included or excluded from the filter 
	NOT *bool `json:"NOT,omitempty"`
}

// NewPropertyFilter instantiates a new PropertyFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyFilter() *PropertyFilter {
	this := PropertyFilter{}
	return &this
}

// NewPropertyFilterWithDefaults instantiates a new PropertyFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyFilterWithDefaults() *PropertyFilter {
	this := PropertyFilter{}
	return &this
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *PropertyFilter) GetProperty() string {
	if o == nil || isNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyFilter) GetPropertyOk() (*string, bool) {
	if o == nil || isNil(o.Property) {
    return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *PropertyFilter) HasProperty() bool {
	if o != nil && !isNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *PropertyFilter) SetProperty(v string) {
	o.Property = &v
}

// GetPropertyValue returns the PropertyValue field value if set, zero value otherwise.
func (o *PropertyFilter) GetPropertyValue() []string {
	if o == nil || isNil(o.PropertyValue) {
		var ret []string
		return ret
	}
	return o.PropertyValue
}

// GetPropertyValueOk returns a tuple with the PropertyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyFilter) GetPropertyValueOk() ([]string, bool) {
	if o == nil || isNil(o.PropertyValue) {
    return nil, false
	}
	return o.PropertyValue, true
}

// HasPropertyValue returns a boolean if a field has been set.
func (o *PropertyFilter) HasPropertyValue() bool {
	if o != nil && !isNil(o.PropertyValue) {
		return true
	}

	return false
}

// SetPropertyValue gets a reference to the given []string and assigns it to the PropertyValue field.
func (o *PropertyFilter) SetPropertyValue(v []string) {
	o.PropertyValue = v
}

// GetNOT returns the NOT field value if set, zero value otherwise.
func (o *PropertyFilter) GetNOT() bool {
	if o == nil || isNil(o.NOT) {
		var ret bool
		return ret
	}
	return *o.NOT
}

// GetNOTOk returns a tuple with the NOT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyFilter) GetNOTOk() (*bool, bool) {
	if o == nil || isNil(o.NOT) {
    return nil, false
	}
	return o.NOT, true
}

// HasNOT returns a boolean if a field has been set.
func (o *PropertyFilter) HasNOT() bool {
	if o != nil && !isNil(o.NOT) {
		return true
	}

	return false
}

// SetNOT gets a reference to the given bool and assigns it to the NOT field.
func (o *PropertyFilter) SetNOT(v bool) {
	o.NOT = &v
}

func (o PropertyFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	if !isNil(o.PropertyValue) {
		toSerialize["propertyValue"] = o.PropertyValue
	}
	if !isNil(o.NOT) {
		toSerialize["NOT"] = o.NOT
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyFilter struct {
	value *PropertyFilter
	isSet bool
}

func (v NullablePropertyFilter) Get() *PropertyFilter {
	return v.value
}

func (v *NullablePropertyFilter) Set(val *PropertyFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyFilter(val *PropertyFilter) *NullablePropertyFilter {
	return &NullablePropertyFilter{value: val, isSet: true}
}

func (v NullablePropertyFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


