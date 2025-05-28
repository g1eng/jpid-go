package jpid

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type SearchcodeAPIService service

type ApiApiV1SearchcodeSearchCodeGetRequest struct {
	ctx        context.Context
	ApiService *SearchcodeAPIService
	searchCode string
	page       *float32
	limit      *float32
	ecUid      *string
	choikitype *float32
	searchtype *float32
}

func (r ApiApiV1SearchcodeSearchCodeGetRequest) Page(page float32) ApiApiV1SearchcodeSearchCodeGetRequest {
	r.page = &page
	return r
}

func (r ApiApiV1SearchcodeSearchCodeGetRequest) Limit(limit float32) ApiApiV1SearchcodeSearchCodeGetRequest {
	r.limit = &limit
	return r
}

func (r ApiApiV1SearchcodeSearchCodeGetRequest) EcUid(ecUid string) ApiApiV1SearchcodeSearchCodeGetRequest {
	r.ecUid = &ecUid
	return r
}

func (r ApiApiV1SearchcodeSearchCodeGetRequest) Choikitype(choikitype float32) ApiApiV1SearchcodeSearchCodeGetRequest {
	r.choikitype = &choikitype
	return r
}

func (r ApiApiV1SearchcodeSearchCodeGetRequest) Searchtype(searchtype float32) ApiApiV1SearchcodeSearchCodeGetRequest {
	r.searchtype = &searchtype
	return r
}

func (r ApiApiV1SearchcodeSearchCodeGetRequest) Execute() (*SearchcodeSearchRes, *http.Response, error) {
	return r.ApiService.ApiV1SearchcodeSearchCodeGetExecute(r)
}

func (a *SearchcodeAPIService) ApiV1SearchcodeSearchCodeGet(ctx context.Context, searchCode string) ApiApiV1SearchcodeSearchCodeGetRequest {
	return ApiApiV1SearchcodeSearchCodeGetRequest{
		ApiService: a,
		ctx:        ctx,
		searchCode: searchCode,
	}
}

func (a *SearchcodeAPIService) ApiV1SearchcodeSearchCodeGetExecute(r ApiApiV1SearchcodeSearchCodeGetRequest) (*SearchcodeSearchRes, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchcodeSearchRes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchcodeAPIService.ApiV1SearchcodeSearchCodeGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/searchcode/{search_code}"
	localVarPath = strings.Replace(localVarPath, "{"+"search_code"+"}", url.PathEscape(parameterValueToString(r.searchCode, "searchCode")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.ecUid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ec_uid", r.ecUid, "form", "")
	}
	if r.choikitype != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "choikitype", r.choikitype, "form", "")
	}
	if r.searchtype != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchtype", r.searchtype, "form", "")
	}

	localVarHTTPContentTypes := []string{}

	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	localVarHTTPHeaderAccepts := []string{"application/json"}

	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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
			var v BadRequest
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
