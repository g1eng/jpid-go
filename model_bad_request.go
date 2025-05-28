package jpid

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var _ MappedNullable = &BadRequest{}

type BadRequest struct {
	RequestId string `json:"request_id"`

	ErrorCode string `json:"error_code"`

	Message string `json:"message"`
}

type _BadRequest BadRequest

func NewBadRequest(requestId string, errorCode string, message string) *BadRequest {
	this := BadRequest{}
	this.RequestId = requestId
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

func NewBadRequestWithDefaults() *BadRequest {
	this := BadRequest{}
	return &this
}

func (o *BadRequest) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

func (o *BadRequest) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

func (o *BadRequest) SetRequestId(v string) {
	o.RequestId = v
}

func (o *BadRequest) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

func (o *BadRequest) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

func (o *BadRequest) SetErrorCode(v string) {
	o.ErrorCode = v
}

func (o *BadRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

func (o *BadRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

func (o *BadRequest) SetMessage(v string) {
	o.Message = v
}

func (o BadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["error_code"] = o.ErrorCode
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *BadRequest) UnmarshalJSON(data []byte) (err error) {

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

	varBadRequest := _BadRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBadRequest)

	if err != nil {
		return err
	}

	*o = BadRequest(varBadRequest)

	return err
}

type NullableBadRequest struct {
	value *BadRequest
	isSet bool
}

func (v NullableBadRequest) Get() *BadRequest {
	return v.value
}

func (v *NullableBadRequest) Set(val *BadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequest(val *BadRequest) *NullableBadRequest {
	return &NullableBadRequest{value: val, isSet: true}
}

func (v NullableBadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
