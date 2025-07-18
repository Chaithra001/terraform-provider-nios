/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type RecordNsec3paramAPI interface {
	/*
		List Retrieve record:nsec3param objects

		Returns a list of record:nsec3param objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RecordNsec3paramAPIListRequest
	*/
	List(ctx context.Context) RecordNsec3paramAPIListRequest

	// ListExecute executes the request
	//  @return ListRecordNsec3paramResponse
	ListExecute(r RecordNsec3paramAPIListRequest) (*ListRecordNsec3paramResponse, *http.Response, error)
	/*
		Read Get a specific record:nsec3param object

		Returns a specific record:nsec3param object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the record:nsec3param object
		@return RecordNsec3paramAPIReadRequest
	*/
	Read(ctx context.Context, reference string) RecordNsec3paramAPIReadRequest

	// ReadExecute executes the request
	//  @return GetRecordNsec3paramResponse
	ReadExecute(r RecordNsec3paramAPIReadRequest) (*GetRecordNsec3paramResponse, *http.Response, error)
}

// RecordNsec3paramAPIService RecordNsec3paramAPI service
type RecordNsec3paramAPIService internal.Service

type RecordNsec3paramAPIListRequest struct {
	ctx              context.Context
	ApiService       RecordNsec3paramAPI
	returnFields     *string
	returnFieldsPlus *string
	maxResults       *int32
	returnAsObject   *int32
	paging           *int32
	pageId           *string
	filters          *map[string]interface{}
	extattrfilter    *map[string]interface{}
}

// Enter the field names followed by comma
func (r RecordNsec3paramAPIListRequest) ReturnFields(returnFields string) RecordNsec3paramAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordNsec3paramAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordNsec3paramAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r RecordNsec3paramAPIListRequest) MaxResults(maxResults int32) RecordNsec3paramAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r RecordNsec3paramAPIListRequest) ReturnAsObject(returnAsObject int32) RecordNsec3paramAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r RecordNsec3paramAPIListRequest) Paging(paging int32) RecordNsec3paramAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r RecordNsec3paramAPIListRequest) PageId(pageId string) RecordNsec3paramAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r RecordNsec3paramAPIListRequest) Filters(filters map[string]interface{}) RecordNsec3paramAPIListRequest {
	r.filters = &filters
	return r
}

func (r RecordNsec3paramAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) RecordNsec3paramAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r RecordNsec3paramAPIListRequest) Execute() (*ListRecordNsec3paramResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve record:nsec3param objects

Returns a list of record:nsec3param objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecordNsec3paramAPIListRequest
*/
func (a *RecordNsec3paramAPIService) List(ctx context.Context) RecordNsec3paramAPIListRequest {
	return RecordNsec3paramAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListRecordNsec3paramResponse
func (a *RecordNsec3paramAPIService) ListExecute(r RecordNsec3paramAPIListRequest) (*ListRecordNsec3paramResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListRecordNsec3paramResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordNsec3paramAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:nsec3param"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "form", "")
	}
	if r.returnFieldsPlus != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFieldsPlus, "form", "")
	}
	if r.maxResults != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_max_results", r.maxResults, "form", "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "form", "")
	}
	if r.paging != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_paging", r.paging, "form", "")
	}
	if r.pageId != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_id", r.pageId, "form", "")
	}
	if r.filters != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "filters", r.filters, "form", "")
	}
	if r.extattrfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "extattrfilter", r.extattrfilter, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordNsec3paramAPIReadRequest struct {
	ctx              context.Context
	ApiService       RecordNsec3paramAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r RecordNsec3paramAPIReadRequest) ReturnFields(returnFields string) RecordNsec3paramAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordNsec3paramAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordNsec3paramAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r RecordNsec3paramAPIReadRequest) ReturnAsObject(returnAsObject int32) RecordNsec3paramAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r RecordNsec3paramAPIReadRequest) Execute() (*GetRecordNsec3paramResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific record:nsec3param object

Returns a specific record:nsec3param object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the record:nsec3param object
	@return RecordNsec3paramAPIReadRequest
*/
func (a *RecordNsec3paramAPIService) Read(ctx context.Context, reference string) RecordNsec3paramAPIReadRequest {
	return RecordNsec3paramAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetRecordNsec3paramResponse
func (a *RecordNsec3paramAPIService) ReadExecute(r RecordNsec3paramAPIReadRequest) (*GetRecordNsec3paramResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetRecordNsec3paramResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordNsec3paramAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:nsec3param/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "form", "")
	}
	if r.returnFieldsPlus != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFieldsPlus, "form", "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
