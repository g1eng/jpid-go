package jpid

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var _ MappedNullable = &Forbidden{}

type Forbidden struct {
	RequestId string `json:"request_id"`

	ErrorCode string `json:"error_code"`

	Message string `json:"message"`
}

type _Forbidden Forbidden

func NewForbidden(requestId string, errorCode string, message string) *Forbidden {
	this := Forbidden{}
	this.RequestId = requestId
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

func NewForbiddenWithDefaults() *Forbidden {
	this := Forbidden{}
	return &this
}

func (o *Forbidden) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

func (o *Forbidden) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

func (o *Forbidden) SetRequestId(v string) {
	o.RequestId = v
}

func (o *Forbidden) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

func (o *Forbidden) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

func (o *Forbidden) SetErrorCode(v string) {
	o.ErrorCode = v
}

func (o *Forbidden) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

func (o *Forbidden) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

func (o *Forbidden) SetMessage(v string) {
	o.Message = v
}

func (o Forbidden) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Forbidden) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["error_code"] = o.ErrorCode
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *Forbidden) UnmarshalJSON(data []byte) (err error) {

	requiredProperties := []string{
		"request_id",
		"error_code",
		"message",
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

	varForbidden := _Forbidden{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varForbidden)

	if err != nil {
		return err
	}

	*o = Forbidden(varForbidden)

	return err
}

type NullableForbidden struct {
	value *Forbidden
	isSet bool
}

func (v NullableForbidden) Get() *Forbidden {
	return v.value
}

func (v *NullableForbidden) Set(val *Forbidden) {
	v.value = val
	v.isSet = true
}

func (v NullableForbidden) IsSet() bool {
	return v.isSet
}

func (v *NullableForbidden) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForbidden(val *Forbidden) *NullableForbidden {
	return &NullableForbidden{value: val, isSet: true}
}

func (v NullableForbidden) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForbidden) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
