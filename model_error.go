package jpid

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var _ MappedNullable = &Error{}

type Error struct {
	RequestId string `json:"request_id"`

	ErrorCode string `json:"error_code"`

	Message string `json:"message"`
}

type _Error Error

func NewError(requestId string, errorCode string, message string) *Error {
	this := Error{}
	this.RequestId = requestId
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

func (o *Error) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

func (o *Error) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

func (o *Error) SetRequestId(v string) {
	o.RequestId = v
}

func (o *Error) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

func (o *Error) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

func (o *Error) SetErrorCode(v string) {
	o.ErrorCode = v
}

func (o *Error) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

func (o *Error) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

func (o *Error) SetMessage(v string) {
	o.Message = v
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["error_code"] = o.ErrorCode
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *Error) UnmarshalJSON(data []byte) (err error) {

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

	varError := _Error{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varError)

	if err != nil {
		return err
	}

	*o = Error(varError)

	return err
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
