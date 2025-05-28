package jpid

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var _ MappedNullable = &JtokenReq{}

type JtokenReq struct {
	GrantType string `json:"grant_type"`

	ClientId string `json:"client_id"`

	SecretKey string `json:"secret_key"`
}

type _JtokenReq JtokenReq

func NewJtokenReq(grantType string, clientId string, secretKey string) *JtokenReq {
	this := JtokenReq{}
	this.GrantType = grantType
	this.ClientId = clientId
	this.SecretKey = secretKey
	return &this
}

func NewJtokenReqWithDefaults() *JtokenReq {
	this := JtokenReq{}
	return &this
}

func (o *JtokenReq) GetGrantType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantType
}

func (o *JtokenReq) GetGrantTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

func (o *JtokenReq) SetGrantType(v string) {
	o.GrantType = v
}

func (o *JtokenReq) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

func (o *JtokenReq) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

func (o *JtokenReq) SetClientId(v string) {
	o.ClientId = v
}

func (o *JtokenReq) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

func (o *JtokenReq) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretKey, true
}

func (o *JtokenReq) SetSecretKey(v string) {
	o.SecretKey = v
}

func (o JtokenReq) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JtokenReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grant_type"] = o.GrantType
	toSerialize["client_id"] = o.ClientId
	toSerialize["secret_key"] = o.SecretKey
	return toSerialize, nil
}

func (o *JtokenReq) UnmarshalJSON(data []byte) (err error) {

	requiredProperties := []string{
		"grant_type",
		"client_id",
		"secret_key",
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

	varJtokenReq := _JtokenReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJtokenReq)

	if err != nil {
		return err
	}

	*o = JtokenReq(varJtokenReq)

	return err
}

type NullableJtokenReq struct {
	value *JtokenReq
	isSet bool
}

func (v NullableJtokenReq) Get() *JtokenReq {
	return v.value
}

func (v *NullableJtokenReq) Set(val *JtokenReq) {
	v.value = val
	v.isSet = true
}

func (v NullableJtokenReq) IsSet() bool {
	return v.isSet
}

func (v *NullableJtokenReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJtokenReq(val *JtokenReq) *NullableJtokenReq {
	return &NullableJtokenReq{value: val, isSet: true}
}

func (v NullableJtokenReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJtokenReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
