package jpid

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var _ MappedNullable = &SearchcodeSearchRes{}

type SearchcodeSearchRes struct {
	Page int32 `json:"page"`

	Limit int32 `json:"limit"`

	Count int32 `json:"count"`

	Searchtype string                              `json:"searchtype"`
	Addresses  []SearchcodeSearchResAddressesInner `json:"addresses"`
}

type _SearchcodeSearchRes SearchcodeSearchRes

func NewSearchcodeSearchRes(page int32, limit int32, count int32, searchtype string, addresses []SearchcodeSearchResAddressesInner) *SearchcodeSearchRes {
	this := SearchcodeSearchRes{}
	this.Page = page
	this.Limit = limit
	this.Count = count
	this.Searchtype = searchtype
	this.Addresses = addresses
	return &this
}

func NewSearchcodeSearchResWithDefaults() *SearchcodeSearchRes {
	this := SearchcodeSearchRes{}
	return &this
}

func (o *SearchcodeSearchRes) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

func (o *SearchcodeSearchRes) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

func (o *SearchcodeSearchRes) SetPage(v int32) {
	o.Page = v
}

func (o *SearchcodeSearchRes) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

func (o *SearchcodeSearchRes) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

func (o *SearchcodeSearchRes) SetLimit(v int32) {
	o.Limit = v
}

func (o *SearchcodeSearchRes) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

func (o *SearchcodeSearchRes) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

func (o *SearchcodeSearchRes) SetCount(v int32) {
	o.Count = v
}

func (o *SearchcodeSearchRes) GetSearchtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Searchtype
}

func (o *SearchcodeSearchRes) GetSearchtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Searchtype, true
}

func (o *SearchcodeSearchRes) SetSearchtype(v string) {
	o.Searchtype = v
}

func (o *SearchcodeSearchRes) GetAddresses() []SearchcodeSearchResAddressesInner {
	if o == nil {
		var ret []SearchcodeSearchResAddressesInner
		return ret
	}

	return o.Addresses
}

func (o *SearchcodeSearchRes) GetAddressesOk() ([]SearchcodeSearchResAddressesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

func (o *SearchcodeSearchRes) SetAddresses(v []SearchcodeSearchResAddressesInner) {
	o.Addresses = v
}

func (o SearchcodeSearchRes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchcodeSearchRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["page"] = o.Page
	toSerialize["limit"] = o.Limit
	toSerialize["count"] = o.Count
	toSerialize["searchtype"] = o.Searchtype
	toSerialize["addresses"] = o.Addresses
	return toSerialize, nil
}

func (o *SearchcodeSearchRes) UnmarshalJSON(data []byte) (err error) {

	requiredProperties := []string{
		"page",
		"limit",
		"count",
		"searchtype",
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

	varSearchcodeSearchRes := _SearchcodeSearchRes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchcodeSearchRes)

	if err != nil {
		return err
	}

	*o = SearchcodeSearchRes(varSearchcodeSearchRes)

	return err
}

type NullableSearchcodeSearchRes struct {
	value *SearchcodeSearchRes
	isSet bool
}

func (v NullableSearchcodeSearchRes) Get() *SearchcodeSearchRes {
	return v.value
}

func (v *NullableSearchcodeSearchRes) Set(val *SearchcodeSearchRes) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchcodeSearchRes) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchcodeSearchRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchcodeSearchRes(val *SearchcodeSearchRes) *NullableSearchcodeSearchRes {
	return &NullableSearchcodeSearchRes{value: val, isSet: true}
}

func (v NullableSearchcodeSearchRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchcodeSearchRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
