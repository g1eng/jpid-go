package jpid

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var _ MappedNullable = &AddressRes{}

type AddressRes struct {
	Level int32 `json:"level"`

	Page int32 `json:"page"`

	Limit int32 `json:"limit"`

	Count     int32                      `json:"count"`
	Addresses []AddressResAddressesInner `json:"addresses"`
}

type _AddressRes AddressRes

func NewAddressRes(level int32, page int32, limit int32, count int32, addresses []AddressResAddressesInner) *AddressRes {
	this := AddressRes{}
	this.Level = level
	this.Page = page
	this.Limit = limit
	this.Count = count
	this.Addresses = addresses
	return &this
}

func NewAddressResWithDefaults() *AddressRes {
	this := AddressRes{}
	return &this
}

func (o *AddressRes) GetLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Level
}

func (o *AddressRes) GetLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

func (o *AddressRes) SetLevel(v int32) {
	o.Level = v
}

func (o *AddressRes) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

func (o *AddressRes) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

func (o *AddressRes) SetPage(v int32) {
	o.Page = v
}

func (o *AddressRes) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

func (o *AddressRes) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

func (o *AddressRes) SetLimit(v int32) {
	o.Limit = v
}

func (o *AddressRes) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

func (o *AddressRes) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

func (o *AddressRes) SetCount(v int32) {
	o.Count = v
}

func (o *AddressRes) GetAddresses() []AddressResAddressesInner {
	if o == nil {
		var ret []AddressResAddressesInner
		return ret
	}

	return o.Addresses
}

func (o *AddressRes) GetAddressesOk() ([]AddressResAddressesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

func (o *AddressRes) SetAddresses(v []AddressResAddressesInner) {
	o.Addresses = v
}

func (o AddressRes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	toSerialize["page"] = o.Page
	toSerialize["limit"] = o.Limit
	toSerialize["count"] = o.Count
	toSerialize["addresses"] = o.Addresses
	return toSerialize, nil
}

func (o *AddressRes) UnmarshalJSON(data []byte) (err error) {

	requiredProperties := []string{
		"level",
		"page",
		"limit",
		"count",
		"addresses",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddressRes := _AddressRes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressRes)

	if err != nil {
		return err
	}

	*o = AddressRes(varAddressRes)

	return err
}

type NullableAddressRes struct {
	value *AddressRes
	isSet bool
}

func (v NullableAddressRes) Get() *AddressRes {
	return v.value
}

func (v *NullableAddressRes) Set(val *AddressRes) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressRes) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressRes(val *AddressRes) *NullableAddressRes {
	return &NullableAddressRes{value: val, isSet: true}
}

func (v NullableAddressRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
