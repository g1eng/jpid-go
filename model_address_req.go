package jpid

import (
	"encoding/json"
)

var _ MappedNullable = &AddressReq{}

type AddressReq struct {
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

	Freeword *string `json:"freeword,omitempty"`

	FlgGetcity *float32 `json:"flg_getcity,omitempty"`

	FlgGetpref *float32 `json:"flg_getpref,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func NewAddressReq() *AddressReq {
	this := AddressReq{}
	return &this
}

func NewAddressReqWithDefaults() *AddressReq {
	this := AddressReq{}
	return &this
}

func (o *AddressReq) GetPrefCode() string {
	if o == nil || IsNil(o.PrefCode) {
		var ret string
		return ret
	}
	return *o.PrefCode
}

func (o *AddressReq) GetPrefCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PrefCode) {
		return nil, false
	}
	return o.PrefCode, true
}

func (o *AddressReq) HasPrefCode() bool {
	if o != nil && !IsNil(o.PrefCode) {
		return true
	}

	return false
}

func (o *AddressReq) SetPrefCode(v string) {
	o.PrefCode = &v
}

func (o *AddressReq) GetPrefName() string {
	if o == nil || IsNil(o.PrefName) {
		var ret string
		return ret
	}
	return *o.PrefName
}

func (o *AddressReq) GetPrefNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrefName) {
		return nil, false
	}
	return o.PrefName, true
}

func (o *AddressReq) HasPrefName() bool {
	if o != nil && !IsNil(o.PrefName) {
		return true
	}

	return false
}

func (o *AddressReq) SetPrefName(v string) {
	o.PrefName = &v
}

func (o *AddressReq) GetPrefKana() string {
	if o == nil || IsNil(o.PrefKana) {
		var ret string
		return ret
	}
	return *o.PrefKana
}

func (o *AddressReq) GetPrefKanaOk() (*string, bool) {
	if o == nil || IsNil(o.PrefKana) {
		return nil, false
	}
	return o.PrefKana, true
}

func (o *AddressReq) HasPrefKana() bool {
	if o != nil && !IsNil(o.PrefKana) {
		return true
	}

	return false
}

func (o *AddressReq) SetPrefKana(v string) {
	o.PrefKana = &v
}

func (o *AddressReq) GetPrefRoma() string {
	if o == nil || IsNil(o.PrefRoma) {
		var ret string
		return ret
	}
	return *o.PrefRoma
}

func (o *AddressReq) GetPrefRomaOk() (*string, bool) {
	if o == nil || IsNil(o.PrefRoma) {
		return nil, false
	}
	return o.PrefRoma, true
}

func (o *AddressReq) HasPrefRoma() bool {
	if o != nil && !IsNil(o.PrefRoma) {
		return true
	}

	return false
}

func (o *AddressReq) SetPrefRoma(v string) {
	o.PrefRoma = &v
}

func (o *AddressReq) GetCityCode() string {
	if o == nil || IsNil(o.CityCode) {
		var ret string
		return ret
	}
	return *o.CityCode
}

func (o *AddressReq) GetCityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CityCode) {
		return nil, false
	}
	return o.CityCode, true
}

func (o *AddressReq) HasCityCode() bool {
	if o != nil && !IsNil(o.CityCode) {
		return true
	}

	return false
}

func (o *AddressReq) SetCityCode(v string) {
	o.CityCode = &v
}

func (o *AddressReq) GetCityName() string {
	if o == nil || IsNil(o.CityName) {
		var ret string
		return ret
	}
	return *o.CityName
}

func (o *AddressReq) GetCityNameOk() (*string, bool) {
	if o == nil || IsNil(o.CityName) {
		return nil, false
	}
	return o.CityName, true
}

func (o *AddressReq) HasCityName() bool {
	if o != nil && !IsNil(o.CityName) {
		return true
	}

	return false
}

func (o *AddressReq) SetCityName(v string) {
	o.CityName = &v
}

func (o *AddressReq) GetCityKana() string {
	if o == nil || IsNil(o.CityKana) {
		var ret string
		return ret
	}
	return *o.CityKana
}

func (o *AddressReq) GetCityKanaOk() (*string, bool) {
	if o == nil || IsNil(o.CityKana) {
		return nil, false
	}
	return o.CityKana, true
}

func (o *AddressReq) HasCityKana() bool {
	if o != nil && !IsNil(o.CityKana) {
		return true
	}

	return false
}

func (o *AddressReq) SetCityKana(v string) {
	o.CityKana = &v
}

func (o *AddressReq) GetCityRoma() string {
	if o == nil || IsNil(o.CityRoma) {
		var ret string
		return ret
	}
	return *o.CityRoma
}

func (o *AddressReq) GetCityRomaOk() (*string, bool) {
	if o == nil || IsNil(o.CityRoma) {
		return nil, false
	}
	return o.CityRoma, true
}

func (o *AddressReq) HasCityRoma() bool {
	if o != nil && !IsNil(o.CityRoma) {
		return true
	}

	return false
}

func (o *AddressReq) SetCityRoma(v string) {
	o.CityRoma = &v
}

func (o *AddressReq) GetTownName() string {
	if o == nil || IsNil(o.TownName) {
		var ret string
		return ret
	}
	return *o.TownName
}

func (o *AddressReq) GetTownNameOk() (*string, bool) {
	if o == nil || IsNil(o.TownName) {
		return nil, false
	}
	return o.TownName, true
}

func (o *AddressReq) HasTownName() bool {
	if o != nil && !IsNil(o.TownName) {
		return true
	}

	return false
}

func (o *AddressReq) SetTownName(v string) {
	o.TownName = &v
}

func (o *AddressReq) GetTownKana() string {
	if o == nil || IsNil(o.TownKana) {
		var ret string
		return ret
	}
	return *o.TownKana
}

func (o *AddressReq) GetTownKanaOk() (*string, bool) {
	if o == nil || IsNil(o.TownKana) {
		return nil, false
	}
	return o.TownKana, true
}

func (o *AddressReq) HasTownKana() bool {
	if o != nil && !IsNil(o.TownKana) {
		return true
	}

	return false
}

func (o *AddressReq) SetTownKana(v string) {
	o.TownKana = &v
}

func (o *AddressReq) GetTownRoma() string {
	if o == nil || IsNil(o.TownRoma) {
		var ret string
		return ret
	}
	return *o.TownRoma
}

func (o *AddressReq) GetTownRomaOk() (*string, bool) {
	if o == nil || IsNil(o.TownRoma) {
		return nil, false
	}
	return o.TownRoma, true
}

func (o *AddressReq) HasTownRoma() bool {
	if o != nil && !IsNil(o.TownRoma) {
		return true
	}

	return false
}

func (o *AddressReq) SetTownRoma(v string) {
	o.TownRoma = &v
}

func (o *AddressReq) GetFreeword() string {
	if o == nil || IsNil(o.Freeword) {
		var ret string
		return ret
	}
	return *o.Freeword
}

func (o *AddressReq) GetFreewordOk() (*string, bool) {
	if o == nil || IsNil(o.Freeword) {
		return nil, false
	}
	return o.Freeword, true
}

func (o *AddressReq) HasFreeword() bool {
	if o != nil && !IsNil(o.Freeword) {
		return true
	}

	return false
}

func (o *AddressReq) SetFreeword(v string) {
	o.Freeword = &v
}

func (o *AddressReq) GetFlgGetcity() float32 {
	if o == nil || IsNil(o.FlgGetcity) {
		var ret float32
		return ret
	}
	return *o.FlgGetcity
}

func (o *AddressReq) GetFlgGetcityOk() (*float32, bool) {
	if o == nil || IsNil(o.FlgGetcity) {
		return nil, false
	}
	return o.FlgGetcity, true
}

func (o *AddressReq) HasFlgGetcity() bool {
	if o != nil && !IsNil(o.FlgGetcity) {
		return true
	}

	return false
}

func (o *AddressReq) SetFlgGetcity(v float32) {
	o.FlgGetcity = &v
}

func (o *AddressReq) GetFlgGetpref() float32 {
	if o == nil || IsNil(o.FlgGetpref) {
		var ret float32
		return ret
	}
	return *o.FlgGetpref
}

func (o *AddressReq) GetFlgGetprefOk() (*float32, bool) {
	if o == nil || IsNil(o.FlgGetpref) {
		return nil, false
	}
	return o.FlgGetpref, true
}

func (o *AddressReq) HasFlgGetpref() bool {
	if o != nil && !IsNil(o.FlgGetpref) {
		return true
	}

	return false
}

func (o *AddressReq) SetFlgGetpref(v float32) {
	o.FlgGetpref = &v
}

func (o *AddressReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

func (o *AddressReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

func (o *AddressReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

func (o *AddressReq) SetPage(v int32) {
	o.Page = &v
}

func (o *AddressReq) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

func (o *AddressReq) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

func (o *AddressReq) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

func (o *AddressReq) SetLimit(v int32) {
	o.Limit = &v
}

func (o AddressReq) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.Freeword) {
		toSerialize["freeword"] = o.Freeword
	}
	if !IsNil(o.FlgGetcity) {
		toSerialize["flg_getcity"] = o.FlgGetcity
	}
	if !IsNil(o.FlgGetpref) {
		toSerialize["flg_getpref"] = o.FlgGetpref
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

type NullableAddressReq struct {
	value *AddressReq
	isSet bool
}

func (v NullableAddressReq) Get() *AddressReq {
	return v.value
}

func (v *NullableAddressReq) Set(val *AddressReq) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressReq) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressReq(val *AddressReq) *NullableAddressReq {
	return &NullableAddressReq{value: val, isSet: true}
}

func (v NullableAddressReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
