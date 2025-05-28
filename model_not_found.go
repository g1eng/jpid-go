package jpid

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var _ MappedNullable = &NotFound{}

type NotFound struct {
	RequestId string `json:"request_id"`

	ErrorCode string `json:"error_code"`

	Message string `json:"message"`
}

type _NotFound NotFound

func NewNotFound(requestId string, errorCode string, message string) *NotFound {
	this := NotFound{}
	this.RequestId = requestId
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

func NewNotFoundWithDefaults() *NotFound {
	this := NotFound{}
	return &this
}

func (o *NotFound) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

func (o *NotFound) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

func (o *NotFound) SetRequestId(v string) {
	o.RequestId = v
}

func (o *NotFound) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

func (o *NotFound) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

func (o *NotFound) SetErrorCode(v string) {
	o.ErrorCode = v
}

func (o *NotFound) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

func (o *NotFound) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

func (o *NotFound) SetMessage(v string) {
	o.Message = v
}

func (o NotFound) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotFound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["error_code"] = o.ErrorCode
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *NotFound) UnmarshalJSON(data []byte) (err error) {

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

	varNotFound := _NotFound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotFound)

	if err != nil {
		return err
	}

	*o = NotFound(varNotFound)

	return err
}

type NullableNotFound struct {
	value *NotFound
	isSet bool
}

func (v NullableNotFound) Get() *NotFound {
	return v.value
}

func (v *NullableNotFound) Set(val *NotFound) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFound) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFound(val *NotFound) *NullableNotFound {
	return &NullableNotFound{value: val, isSet: true}
}

func (v NullableNotFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
