package jpid

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var _ MappedNullable = &JtokenRes{}

type JtokenRes struct {
	Token string `json:"token"`

	TokenType string `json:"token_type"`

	ExpiresIn int32 `json:"expires_in"`

	Scope string `json:"scope"`
}

type _JtokenRes JtokenRes

func NewJtokenRes(token string, tokenType string, expiresIn int32, scope string) *JtokenRes {
	this := JtokenRes{}
	this.Token = token
	this.TokenType = tokenType
	this.ExpiresIn = expiresIn
	this.Scope = scope
	return &this
}

func NewJtokenResWithDefaults() *JtokenRes {
	this := JtokenRes{}
	return &this
}

func (o *JtokenRes) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

func (o *JtokenRes) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

func (o *JtokenRes) SetToken(v string) {
	o.Token = v
}

func (o *JtokenRes) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

func (o *JtokenRes) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

func (o *JtokenRes) SetTokenType(v string) {
	o.TokenType = v
}

func (o *JtokenRes) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

func (o *JtokenRes) GetExpiresInOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

func (o *JtokenRes) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

func (o *JtokenRes) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

func (o *JtokenRes) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

func (o *JtokenRes) SetScope(v string) {
	o.Scope = v
}

func (o JtokenRes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JtokenRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	toSerialize["token_type"] = o.TokenType
	toSerialize["expires_in"] = o.ExpiresIn
	toSerialize["scope"] = o.Scope
	return toSerialize, nil
}

func (o *JtokenRes) UnmarshalJSON(data []byte) (err error) {

	requiredProperties := []string{
		"token",
		"token_type",
		"expires_in",
		"scope",
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

	varJtokenRes := _JtokenRes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJtokenRes)

	if err != nil {
		return err
	}

	*o = JtokenRes(varJtokenRes)

	return err
}

type NullableJtokenRes struct {
	value *JtokenRes
	isSet bool
}

func (v NullableJtokenRes) Get() *JtokenRes {
	return v.value
}

func (v *NullableJtokenRes) Set(val *JtokenRes) {
	v.value = val
	v.isSet = true
}

func (v NullableJtokenRes) IsSet() bool {
	return v.isSet
}

func (v *NullableJtokenRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJtokenRes(val *JtokenRes) *NullableJtokenRes {
	return &NullableJtokenRes{value: val, isSet: true}
}

func (v NullableJtokenRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJtokenRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
