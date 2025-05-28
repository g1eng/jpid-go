package jpid

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type AddresszipAPIService service

type ApiApiV1AddresszipPostRequest struct {
	ctx        context.Context
	ApiService *AddresszipAPIService
	addressReq *AddressReq
	ecUid      *string
}

func (r ApiApiV1AddresszipPostRequest) AddressReq(addressReq AddressReq) ApiApiV1AddresszipPostRequest {
	r.addressReq = &addressReq
	return r
}

func (r ApiApiV1AddresszipPostRequest) EcUid(ecUid string) ApiApiV1AddresszipPostRequest {
	r.ecUid = &ecUid
	return r
}

func (r ApiApiV1AddresszipPostRequest) Execute() (*AddressRes, *http.Response, error) {
	return r.ApiService.ApiV1AddresszipPostExecute(r)
}

func (a *AddresszipAPIService) ApiV1AddresszipPost(ctx context.Context) ApiApiV1AddresszipPostRequest {
	return ApiApiV1AddresszipPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *AddresszipAPIService) ApiV1AddresszipPostExecute(r ApiApiV1AddresszipPostRequest) (*AddressRes, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddressRes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddresszipAPIService.ApiV1AddresszipPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/addresszip"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addressReq == nil {
		return localVarReturnValue, nil, reportError("addressReq is required and must be specified")
	}

	if r.ecUid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ec_uid", r.ecUid, "form", "")
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	localVarHTTPHeaderAccepts := []string{"application/json"}

	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	localVarPostBody = r.addressReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Unauthorized
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v NotFound
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
