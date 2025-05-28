package jpid

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var _ MappedNullable = &Unauthorized{}

type Unauthorized struct {
	RequestId string `json:"request_id"`

	ErrorCode string `json:"error_code"`

	Message string `json:"message"`
}

type _Unauthorized Unauthorized

func NewUnauthorized(requestId string, errorCode string, message string) *Unauthorized {
	this := Unauthorized{}
	this.RequestId = requestId
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

func NewUnauthorizedWithDefaults() *Unauthorized {
	this := Unauthorized{}
	return &this
}

func (o *Unauthorized) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

func (o *Unauthorized) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

func (o *Unauthorized) SetRequestId(v string) {
	o.RequestId = v
}

func (o *Unauthorized) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

func (o *Unauthorized) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

func (o *Unauthorized) SetErrorCode(v string) {
	o.ErrorCode = v
}

func (o *Unauthorized) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

func (o *Unauthorized) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

func (o *Unauthorized) SetMessage(v string) {
	o.Message = v
}

func (o Unauthorized) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Unauthorized) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["error_code"] = o.ErrorCode
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *Unauthorized) UnmarshalJSON(data []byte) (err error) {

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

	varUnauthorized := _Unauthorized{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnauthorized)

	if err != nil {
		return err
	}

	*o = Unauthorized(varUnauthorized)

	return err
}

type NullableUnauthorized struct {
	value *Unauthorized
	isSet bool
}

func (v NullableUnauthorized) Get() *Unauthorized {
	return v.value
}

func (v *NullableUnauthorized) Set(val *Unauthorized) {
	v.value = val
	v.isSet = true
}

func (v NullableUnauthorized) IsSet() bool {
	return v.isSet
}

func (v *NullableUnauthorized) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnauthorized(val *Unauthorized) *NullableUnauthorized {
	return &NullableUnauthorized{value: val, isSet: true}
}

func (v NullableUnauthorized) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnauthorized) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
