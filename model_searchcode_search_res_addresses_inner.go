package jpid

import (
	"encoding/json"
)

var _ MappedNullable = &SearchcodeSearchResAddressesInner{}

type SearchcodeSearchResAddressesInner struct {
	Dgacode   NullableString  `json:"dgacode,omitempty"`
	ZipCode   *int32          `json:"zip_code,omitempty"`
	PrefCode  *string         `json:"pref_code,omitempty"`
	PrefName  *string         `json:"pref_name,omitempty"`
	PrefKana  NullableString  `json:"pref_kana,omitempty"`
	PrefRoma  NullableString  `json:"pref_roma,omitempty"`
	CityCode  *int32          `json:"city_code,omitempty"`
	CityName  *string         `json:"city_name,omitempty"`
	CityKana  NullableString  `json:"city_kana,omitempty"`
	CityRoma  NullableString  `json:"city_roma,omitempty"`
	TownName  *string         `json:"town_name,omitempty"`
	TownKana  NullableString  `json:"town_kana,omitempty"`
	TownRoma  NullableString  `json:"town_roma,omitempty"`
	BizName   *string         `json:"biz_name,omitempty"`
	BizKana   *string         `json:"biz_kana,omitempty"`
	BizRoma   NullableString  `json:"biz_roma,omitempty"`
	BlockName *string         `json:"block_name,omitempty"`
	OtherName NullableString  `json:"other_name,omitempty"`
	Address   NullableString  `json:"address,omitempty"`
	Longitude NullableFloat32 `json:"longitude,omitempty"`
	Latitude  NullableFloat32 `json:"latitude,omitempty"`
}

func NewSearchcodeSearchResAddressesInner() *SearchcodeSearchResAddressesInner {
	this := SearchcodeSearchResAddressesInner{}
	return &this
}

func NewSearchcodeSearchResAddressesInnerWithDefaults() *SearchcodeSearchResAddressesInner {
	this := SearchcodeSearchResAddressesInner{}
	return &this
}

func (o *SearchcodeSearchResAddressesInner) GetDgacode() string {
	if o == nil || IsNil(o.Dgacode.Get()) {
		var ret string
		return ret
	}
	return *o.Dgacode.Get()
}

func (o *SearchcodeSearchResAddressesInner) GetDgacodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dgacode.Get(), o.Dgacode.IsSet()
}

func (o *SearchcodeSearchResAddressesInner) HasDgacode() bool {
	if o != nil && o.Dgacode.IsSet() {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetDgacode(v string) {
	o.Dgacode.Set(&v)
}

func (o *SearchcodeSearchResAddressesInner) SetDgacodeNil() {
	o.Dgacode.Set(nil)
}

func (o *SearchcodeSearchResAddressesInner) UnsetDgacode() {
	o.Dgacode.Unset()
}

func (o *SearchcodeSearchResAddressesInner) GetZipCode() int32 {
	if o == nil || IsNil(o.ZipCode) {
		var ret int32
		return ret
	}
	return *o.ZipCode
}

func (o *SearchcodeSearchResAddressesInner) GetZipCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return o.ZipCode, true
}

func (o *SearchcodeSearchResAddressesInner) HasZipCode() bool {
	if o != nil && !IsNil(o.ZipCode) {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetZipCode(v int32) {
	o.ZipCode = &v
}

func (o *SearchcodeSearchResAddressesInner) GetPrefCode() string {
	if o == nil || IsNil(o.PrefCode) {
		var ret string
		return ret
	}
	return *o.PrefCode
}

func (o *SearchcodeSearchResAddressesInner) GetPrefCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PrefCode) {
		return nil, false
	}
	return o.PrefCode, true
}

func (o *SearchcodeSearchResAddressesInner) HasPrefCode() bool {
	if o != nil && !IsNil(o.PrefCode) {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetPrefCode(v string) {
	o.PrefCode = &v
}

func (o *SearchcodeSearchResAddressesInner) GetPrefName() string {
	if o == nil || IsNil(o.PrefName) {
		var ret string
		return ret
	}
	return *o.PrefName
}

func (o *SearchcodeSearchResAddressesInner) GetPrefNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrefName) {
		return nil, false
	}
	return o.PrefName, true
}

func (o *SearchcodeSearchResAddressesInner) HasPrefName() bool {
	if o != nil && !IsNil(o.PrefName) {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetPrefName(v string) {
	o.PrefName = &v
}

func (o *SearchcodeSearchResAddressesInner) GetPrefKana() string {
	if o == nil || IsNil(o.PrefKana.Get()) {
		var ret string
		return ret
	}
	return *o.PrefKana.Get()
}

func (o *SearchcodeSearchResAddressesInner) GetPrefKanaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrefKana.Get(), o.PrefKana.IsSet()
}

func (o *SearchcodeSearchResAddressesInner) HasPrefKana() bool {
	if o != nil && o.PrefKana.IsSet() {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetPrefKana(v string) {
	o.PrefKana.Set(&v)
}

func (o *SearchcodeSearchResAddressesInner) SetPrefKanaNil() {
	o.PrefKana.Set(nil)
}

func (o *SearchcodeSearchResAddressesInner) UnsetPrefKana() {
	o.PrefKana.Unset()
}

func (o *SearchcodeSearchResAddressesInner) GetPrefRoma() string {
	if o == nil || IsNil(o.PrefRoma.Get()) {
		var ret string
		return ret
	}
	return *o.PrefRoma.Get()
}

func (o *SearchcodeSearchResAddressesInner) GetPrefRomaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrefRoma.Get(), o.PrefRoma.IsSet()
}

func (o *SearchcodeSearchResAddressesInner) HasPrefRoma() bool {
	if o != nil && o.PrefRoma.IsSet() {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetPrefRoma(v string) {
	o.PrefRoma.Set(&v)
}

func (o *SearchcodeSearchResAddressesInner) SetPrefRomaNil() {
	o.PrefRoma.Set(nil)
}

func (o *SearchcodeSearchResAddressesInner) UnsetPrefRoma() {
	o.PrefRoma.Unset()
}

func (o *SearchcodeSearchResAddressesInner) GetCityCode() int32 {
	if o == nil || IsNil(o.CityCode) {
		var ret int32
		return ret
	}
	return *o.CityCode
}

func (o *SearchcodeSearchResAddressesInner) GetCityCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.CityCode) {
		return nil, false
	}
	return o.CityCode, true
}

func (o *SearchcodeSearchResAddressesInner) HasCityCode() bool {
	if o != nil && !IsNil(o.CityCode) {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetCityCode(v int32) {
	o.CityCode = &v
}

func (o *SearchcodeSearchResAddressesInner) GetCityName() string {
	if o == nil || IsNil(o.CityName) {
		var ret string
		return ret
	}
	return *o.CityName
}

func (o *SearchcodeSearchResAddressesInner) GetCityNameOk() (*string, bool) {
	if o == nil || IsNil(o.CityName) {
		return nil, false
	}
	return o.CityName, true
}

func (o *SearchcodeSearchResAddressesInner) HasCityName() bool {
	if o != nil && !IsNil(o.CityName) {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetCityName(v string) {
	o.CityName = &v
}

func (o *SearchcodeSearchResAddressesInner) GetCityKana() string {
	if o == nil || IsNil(o.CityKana.Get()) {
		var ret string
		return ret
	}
	return *o.CityKana.Get()
}

func (o *SearchcodeSearchResAddressesInner) GetCityKanaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CityKana.Get(), o.CityKana.IsSet()
}

func (o *SearchcodeSearchResAddressesInner) HasCityKana() bool {
	if o != nil && o.CityKana.IsSet() {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetCityKana(v string) {
	o.CityKana.Set(&v)
}

func (o *SearchcodeSearchResAddressesInner) SetCityKanaNil() {
	o.CityKana.Set(nil)
}

func (o *SearchcodeSearchResAddressesInner) UnsetCityKana() {
	o.CityKana.Unset()
}

func (o *SearchcodeSearchResAddressesInner) GetCityRoma() string {
	if o == nil || IsNil(o.CityRoma.Get()) {
		var ret string
		return ret
	}
	return *o.CityRoma.Get()
}

func (o *SearchcodeSearchResAddressesInner) GetCityRomaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CityRoma.Get(), o.CityRoma.IsSet()
}

func (o *SearchcodeSearchResAddressesInner) HasCityRoma() bool {
	if o != nil && o.CityRoma.IsSet() {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetCityRoma(v string) {
	o.CityRoma.Set(&v)
}

func (o *SearchcodeSearchResAddressesInner) SetCityRomaNil() {
	o.CityRoma.Set(nil)
}

func (o *SearchcodeSearchResAddressesInner) UnsetCityRoma() {
	o.CityRoma.Unset()
}

func (o *SearchcodeSearchResAddressesInner) GetTownName() string {
	if o == nil || IsNil(o.TownName) {
		var ret string
		return ret
	}
	return *o.TownName
}

func (o *SearchcodeSearchResAddressesInner) GetTownNameOk() (*string, bool) {
	if o == nil || IsNil(o.TownName) {
		return nil, false
	}
	return o.TownName, true
}

func (o *SearchcodeSearchResAddressesInner) HasTownName() bool {
	if o != nil && !IsNil(o.TownName) {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetTownName(v string) {
	o.TownName = &v
}

func (o *SearchcodeSearchResAddressesInner) GetTownKana() string {
	if o == nil || IsNil(o.TownKana.Get()) {
		var ret string
		return ret
	}
	return *o.TownKana.Get()
}

func (o *SearchcodeSearchResAddressesInner) GetTownKanaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TownKana.Get(), o.TownKana.IsSet()
}

func (o *SearchcodeSearchResAddressesInner) HasTownKana() bool {
	if o != nil && o.TownKana.IsSet() {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetTownKana(v string) {
	o.TownKana.Set(&v)
}

func (o *SearchcodeSearchResAddressesInner) SetTownKanaNil() {
	o.TownKana.Set(nil)
}

func (o *SearchcodeSearchResAddressesInner) UnsetTownKana() {
	o.TownKana.Unset()
}

func (o *SearchcodeSearchResAddressesInner) GetTownRoma() string {
	if o == nil || IsNil(o.TownRoma.Get()) {
		var ret string
		return ret
	}
	return *o.TownRoma.Get()
}

func (o *SearchcodeSearchResAddressesInner) GetTownRomaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TownRoma.Get(), o.TownRoma.IsSet()
}

func (o *SearchcodeSearchResAddressesInner) HasTownRoma() bool {
	if o != nil && o.TownRoma.IsSet() {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetTownRoma(v string) {
	o.TownRoma.Set(&v)
}

func (o *SearchcodeSearchResAddressesInner) SetTownRomaNil() {
	o.TownRoma.Set(nil)
}

func (o *SearchcodeSearchResAddressesInner) UnsetTownRoma() {
	o.TownRoma.Unset()
}

func (o *SearchcodeSearchResAddressesInner) GetBizName() string {
	if o == nil || IsNil(o.BizName) {
		var ret string
		return ret
	}
	return *o.BizName
}

func (o *SearchcodeSearchResAddressesInner) GetBizNameOk() (*string, bool) {
	if o == nil || IsNil(o.BizName) {
		return nil, false
	}
	return o.BizName, true
}

func (o *SearchcodeSearchResAddressesInner) HasBizName() bool {
	if o != nil && !IsNil(o.BizName) {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetBizName(v string) {
	o.BizName = &v
}

func (o *SearchcodeSearchResAddressesInner) GetBizKana() string {
	if o == nil || IsNil(o.BizKana) {
		var ret string
		return ret
	}
	return *o.BizKana
}

func (o *SearchcodeSearchResAddressesInner) GetBizKanaOk() (*string, bool) {
	if o == nil || IsNil(o.BizKana) {
		return nil, false
	}
	return o.BizKana, true
}

func (o *SearchcodeSearchResAddressesInner) HasBizKana() bool {
	if o != nil && !IsNil(o.BizKana) {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetBizKana(v string) {
	o.BizKana = &v
}

func (o *SearchcodeSearchResAddressesInner) GetBizRoma() string {
	if o == nil || IsNil(o.BizRoma.Get()) {
		var ret string
		return ret
	}
	return *o.BizRoma.Get()
}

func (o *SearchcodeSearchResAddressesInner) GetBizRomaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BizRoma.Get(), o.BizRoma.IsSet()
}

func (o *SearchcodeSearchResAddressesInner) HasBizRoma() bool {
	if o != nil && o.BizRoma.IsSet() {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetBizRoma(v string) {
	o.BizRoma.Set(&v)
}

func (o *SearchcodeSearchResAddressesInner) SetBizRomaNil() {
	o.BizRoma.Set(nil)
}

func (o *SearchcodeSearchResAddressesInner) UnsetBizRoma() {
	o.BizRoma.Unset()
}

func (o *SearchcodeSearchResAddressesInner) GetBlockName() string {
	if o == nil || IsNil(o.BlockName) {
		var ret string
		return ret
	}
	return *o.BlockName
}

func (o *SearchcodeSearchResAddressesInner) GetBlockNameOk() (*string, bool) {
	if o == nil || IsNil(o.BlockName) {
		return nil, false
	}
	return o.BlockName, true
}

func (o *SearchcodeSearchResAddressesInner) HasBlockName() bool {
	if o != nil && !IsNil(o.BlockName) {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetBlockName(v string) {
	o.BlockName = &v
}

func (o *SearchcodeSearchResAddressesInner) GetOtherName() string {
	if o == nil || IsNil(o.OtherName.Get()) {
		var ret string
		return ret
	}
	return *o.OtherName.Get()
}

func (o *SearchcodeSearchResAddressesInner) GetOtherNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OtherName.Get(), o.OtherName.IsSet()
}

func (o *SearchcodeSearchResAddressesInner) HasOtherName() bool {
	if o != nil && o.OtherName.IsSet() {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetOtherName(v string) {
	o.OtherName.Set(&v)
}

func (o *SearchcodeSearchResAddressesInner) SetOtherNameNil() {
	o.OtherName.Set(nil)
}

func (o *SearchcodeSearchResAddressesInner) UnsetOtherName() {
	o.OtherName.Unset()
}

func (o *SearchcodeSearchResAddressesInner) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

func (o *SearchcodeSearchResAddressesInner) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

func (o *SearchcodeSearchResAddressesInner) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetAddress(v string) {
	o.Address.Set(&v)
}

func (o *SearchcodeSearchResAddressesInner) SetAddressNil() {
	o.Address.Set(nil)
}

func (o *SearchcodeSearchResAddressesInner) UnsetAddress() {
	o.Address.Unset()
}

func (o *SearchcodeSearchResAddressesInner) GetLongitude() float32 {
	if o == nil || IsNil(o.Longitude.Get()) {
		var ret float32
		return ret
	}
	return *o.Longitude.Get()
}

func (o *SearchcodeSearchResAddressesInner) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

func (o *SearchcodeSearchResAddressesInner) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetLongitude(v float32) {
	o.Longitude.Set(&v)
}

func (o *SearchcodeSearchResAddressesInner) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

func (o *SearchcodeSearchResAddressesInner) UnsetLongitude() {
	o.Longitude.Unset()
}

func (o *SearchcodeSearchResAddressesInner) GetLatitude() float32 {
	if o == nil || IsNil(o.Latitude.Get()) {
		var ret float32
		return ret
	}
	return *o.Latitude.Get()
}

func (o *SearchcodeSearchResAddressesInner) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

func (o *SearchcodeSearchResAddressesInner) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

func (o *SearchcodeSearchResAddressesInner) SetLatitude(v float32) {
	o.Latitude.Set(&v)
}

func (o *SearchcodeSearchResAddressesInner) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

func (o *SearchcodeSearchResAddressesInner) UnsetLatitude() {
	o.Latitude.Unset()
}

func (o SearchcodeSearchResAddressesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchcodeSearchResAddressesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Dgacode.IsSet() {
		toSerialize["dgacode"] = o.Dgacode.Get()
	}
	if !IsNil(o.ZipCode) {
		toSerialize["zip_code"] = o.ZipCode
	}
	if !IsNil(o.PrefCode) {
		toSerialize["pref_code"] = o.PrefCode
	}
	if !IsNil(o.PrefName) {
		toSerialize["pref_name"] = o.PrefName
	}
	if o.PrefKana.IsSet() {
		toSerialize["pref_kana"] = o.PrefKana.Get()
	}
	if o.PrefRoma.IsSet() {
		toSerialize["pref_roma"] = o.PrefRoma.Get()
	}
	if !IsNil(o.CityCode) {
		toSerialize["city_code"] = o.CityCode
	}
	if !IsNil(o.CityName) {
		toSerialize["city_name"] = o.CityName
	}
	if o.CityKana.IsSet() {
		toSerialize["city_kana"] = o.CityKana.Get()
	}
	if o.CityRoma.IsSet() {
		toSerialize["city_roma"] = o.CityRoma.Get()
	}
	if !IsNil(o.TownName) {
		toSerialize["town_name"] = o.TownName
	}
	if o.TownKana.IsSet() {
		toSerialize["town_kana"] = o.TownKana.Get()
	}
	if o.TownRoma.IsSet() {
		toSerialize["town_roma"] = o.TownRoma.Get()
	}
	if !IsNil(o.BizName) {
		toSerialize["biz_name"] = o.BizName
	}
	if !IsNil(o.BizKana) {
		toSerialize["biz_kana"] = o.BizKana
	}
	if o.BizRoma.IsSet() {
		toSerialize["biz_roma"] = o.BizRoma.Get()
	}
	if !IsNil(o.BlockName) {
		toSerialize["block_name"] = o.BlockName
	}
	if o.OtherName.IsSet() {
		toSerialize["other_name"] = o.OtherName.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.Longitude.IsSet() {
		toSerialize["longitude"] = o.Longitude.Get()
	}
	if o.Latitude.IsSet() {
		toSerialize["latitude"] = o.Latitude.Get()
	}
	return toSerialize, nil
}

type NullableSearchcodeSearchResAddressesInner struct {
	value *SearchcodeSearchResAddressesInner
	isSet bool
}

func (v NullableSearchcodeSearchResAddressesInner) Get() *SearchcodeSearchResAddressesInner {
	return v.value
}

func (v *NullableSearchcodeSearchResAddressesInner) Set(val *SearchcodeSearchResAddressesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchcodeSearchResAddressesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchcodeSearchResAddressesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchcodeSearchResAddressesInner(val *SearchcodeSearchResAddressesInner) *NullableSearchcodeSearchResAddressesInner {
	return &NullableSearchcodeSearchResAddressesInner{value: val, isSet: true}
}

func (v NullableSearchcodeSearchResAddressesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchcodeSearchResAddressesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
