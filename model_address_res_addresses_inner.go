package jpid

import (
	"encoding/json"
)

var _ MappedNullable = &AddressResAddressesInner{}

type AddressResAddressesInner struct {
	ZipCode  *string `json:"zip_code,omitempty"`
	PrefCode *string `json:"pref_code,omitempty"`
	PrefName *string `json:"pref_name,omitempty"`
	PrefKana *string `json:"pref_kana,omitempty"`
	PrefRoma *string `json:"pref_roma,omitempty"`
	CityCode *string `json:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty"`
	CityKana *string `json:"city_kana,omitempty"`
	CityRoma *string `json:"city_roma,omitempty"`
	TownName *string `json:"town_name,omitempty"`
	TownKana *string `json:"town_kana,omitempty"`
	TownRoma *string `json:"town_roma,omitempty"`
}

func NewAddressResAddressesInner() *AddressResAddressesInner {
	this := AddressResAddressesInner{}
	return &this
}

func NewAddressResAddressesInnerWithDefaults() *AddressResAddressesInner {
	this := AddressResAddressesInner{}
	return &this
}

func (o *AddressResAddressesInner) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode) {
		var ret string
		return ret
	}
	return *o.ZipCode
}

func (o *AddressResAddressesInner) GetZipCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return o.ZipCode, true
}

func (o *AddressResAddressesInner) HasZipCode() bool {
	if o != nil && !IsNil(o.ZipCode) {
		return true
	}

	return false
}

func (o *AddressResAddressesInner) SetZipCode(v string) {
	o.ZipCode = &v
}

func (o *AddressResAddressesInner) GetPrefCode() string {
	if o == nil || IsNil(o.PrefCode) {
		var ret string
		return ret
	}
	return *o.PrefCode
}

func (o *AddressResAddressesInner) GetPrefCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PrefCode) {
		return nil, false
	}
	return o.PrefCode, true
}

func (o *AddressResAddressesInner) HasPrefCode() bool {
	if o != nil && !IsNil(o.PrefCode) {
		return true
	}

	return false
}

func (o *AddressResAddressesInner) SetPrefCode(v string) {
	o.PrefCode = &v
}

func (o *AddressResAddressesInner) GetPrefName() string {
	if o == nil || IsNil(o.PrefName) {
		var ret string
		return ret
	}
	return *o.PrefName
}

func (o *AddressResAddressesInner) GetPrefNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrefName) {
		return nil, false
	}
	return o.PrefName, true
}

func (o *AddressResAddressesInner) HasPrefName() bool {
	if o != nil && !IsNil(o.PrefName) {
		return true
	}

	return false
}

func (o *AddressResAddressesInner) SetPrefName(v string) {
	o.PrefName = &v
}

func (o *AddressResAddressesInner) GetPrefKana() string {
	if o == nil || IsNil(o.PrefKana) {
		var ret string
		return ret
	}
	return *o.PrefKana
}

func (o *AddressResAddressesInner) GetPrefKanaOk() (*string, bool) {
	if o == nil || IsNil(o.PrefKana) {
		return nil, false
	}
	return o.PrefKana, true
}

func (o *AddressResAddressesInner) HasPrefKana() bool {
	if o != nil && !IsNil(o.PrefKana) {
		return true
	}

	return false
}

func (o *AddressResAddressesInner) SetPrefKana(v string) {
	o.PrefKana = &v
}

func (o *AddressResAddressesInner) GetPrefRoma() string {
	if o == nil || IsNil(o.PrefRoma) {
		var ret string
		return ret
	}
	return *o.PrefRoma
}

func (o *AddressResAddressesInner) GetPrefRomaOk() (*string, bool) {
	if o == nil || IsNil(o.PrefRoma) {
		return nil, false
	}
	return o.PrefRoma, true
}

func (o *AddressResAddressesInner) HasPrefRoma() bool {
	if o != nil && !IsNil(o.PrefRoma) {
		return true
	}

	return false
}

func (o *AddressResAddressesInner) SetPrefRoma(v string) {
	o.PrefRoma = &v
}

func (o *AddressResAddressesInner) GetCityCode() string {
	if o == nil || IsNil(o.CityCode) {
		var ret string
		return ret
	}
	return *o.CityCode
}

func (o *AddressResAddressesInner) GetCityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CityCode) {
		return nil, false
	}
	return o.CityCode, true
}

func (o *AddressResAddressesInner) HasCityCode() bool {
	if o != nil && !IsNil(o.CityCode) {
		return true
	}

	return false
}

func (o *AddressResAddressesInner) SetCityCode(v string) {
	o.CityCode = &v
}

func (o *AddressResAddressesInner) GetCityName() string {
	if o == nil || IsNil(o.CityName) {
		var ret string
		return ret
	}
	return *o.CityName
}

func (o *AddressResAddressesInner) GetCityNameOk() (*string, bool) {
	if o == nil || IsNil(o.CityName) {
		return nil, false
	}
	return o.CityName, true
}

func (o *AddressResAddressesInner) HasCityName() bool {
	if o != nil && !IsNil(o.CityName) {
		return true
	}

	return false
}

func (o *AddressResAddressesInner) SetCityName(v string) {
	o.CityName = &v
}

func (o *AddressResAddressesInner) GetCityKana() string {
	if o == nil || IsNil(o.CityKana) {
		var ret string
		return ret
	}
	return *o.CityKana
}

func (o *AddressResAddressesInner) GetCityKanaOk() (*string, bool) {
	if o == nil || IsNil(o.CityKana) {
		return nil, false
	}
	return o.CityKana, true
}

func (o *AddressResAddressesInner) HasCityKana() bool {
	if o != nil && !IsNil(o.CityKana) {
		return true
	}

	return false
}

func (o *AddressResAddressesInner) SetCityKana(v string) {
	o.CityKana = &v
}

func (o *AddressResAddressesInner) GetCityRoma() string {
	if o == nil || IsNil(o.CityRoma) {
		var ret string
		return ret
	}
	return *o.CityRoma
}

func (o *AddressResAddressesInner) GetCityRomaOk() (*string, bool) {
	if o == nil || IsNil(o.CityRoma) {
		return nil, false
	}
	return o.CityRoma, true
}

func (o *AddressResAddressesInner) HasCityRoma() bool {
	if o != nil && !IsNil(o.CityRoma) {
		return true
	}

	return false
}

func (o *AddressResAddressesInner) SetCityRoma(v string) {
	o.CityRoma = &v
}

func (o *AddressResAddressesInner) GetTownName() string {
	if o == nil || IsNil(o.TownName) {
		var ret string
		return ret
	}
	return *o.TownName
}

func (o *AddressResAddressesInner) GetTownNameOk() (*string, bool) {
	if o == nil || IsNil(o.TownName) {
		return nil, false
	}
	return o.TownName, true
}

func (o *AddressResAddressesInner) HasTownName() bool {
	if o != nil && !IsNil(o.TownName) {
		return true
	}

	return false
}

func (o *AddressResAddressesInner) SetTownName(v string) {
	o.TownName = &v
}

func (o *AddressResAddressesInner) GetTownKana() string {
	if o == nil || IsNil(o.TownKana) {
		var ret string
		return ret
	}
	return *o.TownKana
}

func (o *AddressResAddressesInner) GetTownKanaOk() (*string, bool) {
	if o == nil || IsNil(o.TownKana) {
		return nil, false
	}
	return o.TownKana, true
}

func (o *AddressResAddressesInner) HasTownKana() bool {
	if o != nil && !IsNil(o.TownKana) {
		return true
	}

	return false
}

func (o *AddressResAddressesInner) SetTownKana(v string) {
	o.TownKana = &v
}

func (o *AddressResAddressesInner) GetTownRoma() string {
	if o == nil || IsNil(o.TownRoma) {
		var ret string
		return ret
	}
	return *o.TownRoma
}

func (o *AddressResAddressesInner) GetTownRomaOk() (*string, bool) {
	if o == nil || IsNil(o.TownRoma) {
		return nil, false
	}
	return o.TownRoma, true
}

func (o *AddressResAddressesInner) HasTownRoma() bool {
	if o != nil && !IsNil(o.TownRoma) {
		return true
	}

	return false
}

func (o *AddressResAddressesInner) SetTownRoma(v string) {
	o.TownRoma = &v
}

func (o AddressResAddressesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressResAddressesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ZipCode) {
		toSerialize["zip_code"] = o.ZipCode
	}
	if !IsNil(o.PrefCode) {
		toSerialize["pref_code"] = o.PrefCode
	}
	if !IsNil(o.PrefName) {
		toSerialize["pref_name"] = o.PrefName
	}
	if !IsNil(o.PrefKana) {
		toSerialize["pref_kana"] = o.PrefKana
	}
	if !IsNil(o.PrefRoma) {
		toSerialize["pref_roma"] = o.PrefRoma
	}
	if !IsNil(o.CityCode) {
		toSerialize["city_code"] = o.CityCode
	}
	if !IsNil(o.CityName) {
		toSerialize["city_name"] = o.CityName
	}
	if !IsNil(o.CityKana) {
		toSerialize["city_kana"] = o.CityKana
	}
	if !IsNil(o.CityRoma) {
		toSerialize["city_roma"] = o.CityRoma
	}
	if !IsNil(o.TownName) {
		toSerialize["town_name"] = o.TownName
	}
	if !IsNil(o.TownKana) {
		toSerialize["town_kana"] = o.TownKana
	}
	if !IsNil(o.TownRoma) {
		toSerialize["town_roma"] = o.TownRoma
	}
	return toSerialize, nil
}

type NullableAddressResAddressesInner struct {
	value *AddressResAddressesInner
	isSet bool
}

func (v NullableAddressResAddressesInner) Get() *AddressResAddressesInner {
	return v.value
}

func (v *NullableAddressResAddressesInner) Set(val *AddressResAddressesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressResAddressesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressResAddressesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressResAddressesInner(val *AddressResAddressesInner) *NullableAddressResAddressesInner {
	return &NullableAddressResAddressesInner{value: val, isSet: true}
}

func (v NullableAddressResAddressesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressResAddressesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
