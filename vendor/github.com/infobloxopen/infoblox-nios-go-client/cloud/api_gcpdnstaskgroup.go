/*
Infoblox CLOUD API

OpenAPI specification for Infoblox NIOS WAPI CLOUD objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloud

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type GcpdnstaskgroupAPI interface {
	/*
		Create Create a gcpdnstaskgroup object

		Creates a new gcpdnstaskgroup object

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return GcpdnstaskgroupAPICreateRequest
	*/
	Create(ctx context.Context) GcpdnstaskgroupAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateGcpdnstaskgroupResponse
	CreateExecute(r GcpdnstaskgroupAPICreateRequest) (*CreateGcpdnstaskgroupResponse, *http.Response, error)
	/*
		Delete Delete a gcpdnstaskgroup object

		Deletes a specific gcpdnstaskgroup object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the gcpdnstaskgroup object
		@return GcpdnstaskgroupAPIDeleteRequest
	*/
	Delete(ctx context.Context, reference string) GcpdnstaskgroupAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r GcpdnstaskgroupAPIDeleteRequest) (*http.Response, error)
	/*
		List Retrieve gcpdnstaskgroup objects

		Returns a list of gcpdnstaskgroup objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return GcpdnstaskgroupAPIListRequest
	*/
	List(ctx context.Context) GcpdnstaskgroupAPIListRequest

	// ListExecute executes the request
	//  @return ListGcpdnstaskgroupResponse
	ListExecute(r GcpdnstaskgroupAPIListRequest) (*ListGcpdnstaskgroupResponse, *http.Response, error)
	/*
		Read Get a specific gcpdnstaskgroup object

		Returns a specific gcpdnstaskgroup object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the gcpdnstaskgroup object
		@return GcpdnstaskgroupAPIReadRequest
	*/
	Read(ctx context.Context, reference string) GcpdnstaskgroupAPIReadRequest

	// ReadExecute executes the request
	//  @return GetGcpdnstaskgroupResponse
	ReadExecute(r GcpdnstaskgroupAPIReadRequest) (*GetGcpdnstaskgroupResponse, *http.Response, error)
	/*
		Update Update a gcpdnstaskgroup object

		Updates a specific gcpdnstaskgroup object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the gcpdnstaskgroup object
		@return GcpdnstaskgroupAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) GcpdnstaskgroupAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateGcpdnstaskgroupResponse
	UpdateExecute(r GcpdnstaskgroupAPIUpdateRequest) (*UpdateGcpdnstaskgroupResponse, *http.Response, error)
}

// GcpdnstaskgroupAPIService GcpdnstaskgroupAPI service
type GcpdnstaskgroupAPIService internal.Service

type GcpdnstaskgroupAPICreateRequest struct {
	ctx              context.Context
	ApiService       GcpdnstaskgroupAPI
	gcpdnstaskgroup  *Gcpdnstaskgroup
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to create
func (r GcpdnstaskgroupAPICreateRequest) Gcpdnstaskgroup(gcpdnstaskgroup Gcpdnstaskgroup) GcpdnstaskgroupAPICreateRequest {
	r.gcpdnstaskgroup = &gcpdnstaskgroup
	return r
}

// Enter the field names followed by comma
func (r GcpdnstaskgroupAPICreateRequest) ReturnFields(returnFields string) GcpdnstaskgroupAPICreateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r GcpdnstaskgroupAPICreateRequest) ReturnFieldsPlus(returnFieldsPlus string) GcpdnstaskgroupAPICreateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r GcpdnstaskgroupAPICreateRequest) ReturnAsObject(returnAsObject int32) GcpdnstaskgroupAPICreateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r GcpdnstaskgroupAPICreateRequest) Execute() (*CreateGcpdnstaskgroupResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a gcpdnstaskgroup object

Creates a new gcpdnstaskgroup object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GcpdnstaskgroupAPICreateRequest
*/
func (a *GcpdnstaskgroupAPIService) Create(ctx context.Context) GcpdnstaskgroupAPICreateRequest {
	return GcpdnstaskgroupAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateGcpdnstaskgroupResponse
func (a *GcpdnstaskgroupAPIService) CreateExecute(r GcpdnstaskgroupAPICreateRequest) (*CreateGcpdnstaskgroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateGcpdnstaskgroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GcpdnstaskgroupAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/gcpdnstaskgroup"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gcpdnstaskgroup == nil {
		return localVarReturnValue, nil, internal.ReportError("gcpdnstaskgroup is required and must be specified")
	}

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
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.gcpdnstaskgroup
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

type GcpdnstaskgroupAPIDeleteRequest struct {
	ctx        context.Context
	ApiService GcpdnstaskgroupAPI
	reference  string
}

func (r GcpdnstaskgroupAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a gcpdnstaskgroup object

Deletes a specific gcpdnstaskgroup object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the gcpdnstaskgroup object
	@return GcpdnstaskgroupAPIDeleteRequest
*/
func (a *GcpdnstaskgroupAPIService) Delete(ctx context.Context, reference string) GcpdnstaskgroupAPIDeleteRequest {
	return GcpdnstaskgroupAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
func (a *GcpdnstaskgroupAPIService) DeleteExecute(r GcpdnstaskgroupAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GcpdnstaskgroupAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/gcpdnstaskgroup/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type GcpdnstaskgroupAPIListRequest struct {
	ctx              context.Context
	ApiService       GcpdnstaskgroupAPI
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
func (r GcpdnstaskgroupAPIListRequest) ReturnFields(returnFields string) GcpdnstaskgroupAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r GcpdnstaskgroupAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) GcpdnstaskgroupAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r GcpdnstaskgroupAPIListRequest) MaxResults(maxResults int32) GcpdnstaskgroupAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r GcpdnstaskgroupAPIListRequest) ReturnAsObject(returnAsObject int32) GcpdnstaskgroupAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r GcpdnstaskgroupAPIListRequest) Paging(paging int32) GcpdnstaskgroupAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r GcpdnstaskgroupAPIListRequest) PageId(pageId string) GcpdnstaskgroupAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r GcpdnstaskgroupAPIListRequest) Filters(filters map[string]interface{}) GcpdnstaskgroupAPIListRequest {
	r.filters = &filters
	return r
}

func (r GcpdnstaskgroupAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) GcpdnstaskgroupAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r GcpdnstaskgroupAPIListRequest) Execute() (*ListGcpdnstaskgroupResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve gcpdnstaskgroup objects

Returns a list of gcpdnstaskgroup objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GcpdnstaskgroupAPIListRequest
*/
func (a *GcpdnstaskgroupAPIService) List(ctx context.Context) GcpdnstaskgroupAPIListRequest {
	return GcpdnstaskgroupAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListGcpdnstaskgroupResponse
func (a *GcpdnstaskgroupAPIService) ListExecute(r GcpdnstaskgroupAPIListRequest) (*ListGcpdnstaskgroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListGcpdnstaskgroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GcpdnstaskgroupAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/gcpdnstaskgroup"

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

type GcpdnstaskgroupAPIReadRequest struct {
	ctx              context.Context
	ApiService       GcpdnstaskgroupAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r GcpdnstaskgroupAPIReadRequest) ReturnFields(returnFields string) GcpdnstaskgroupAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r GcpdnstaskgroupAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) GcpdnstaskgroupAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r GcpdnstaskgroupAPIReadRequest) ReturnAsObject(returnAsObject int32) GcpdnstaskgroupAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r GcpdnstaskgroupAPIReadRequest) Execute() (*GetGcpdnstaskgroupResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific gcpdnstaskgroup object

Returns a specific gcpdnstaskgroup object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the gcpdnstaskgroup object
	@return GcpdnstaskgroupAPIReadRequest
*/
func (a *GcpdnstaskgroupAPIService) Read(ctx context.Context, reference string) GcpdnstaskgroupAPIReadRequest {
	return GcpdnstaskgroupAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetGcpdnstaskgroupResponse
func (a *GcpdnstaskgroupAPIService) ReadExecute(r GcpdnstaskgroupAPIReadRequest) (*GetGcpdnstaskgroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetGcpdnstaskgroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GcpdnstaskgroupAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/gcpdnstaskgroup/{reference}"
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

type GcpdnstaskgroupAPIUpdateRequest struct {
	ctx              context.Context
	ApiService       GcpdnstaskgroupAPI
	reference        string
	gcpdnstaskgroup  *Gcpdnstaskgroup
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to update
func (r GcpdnstaskgroupAPIUpdateRequest) Gcpdnstaskgroup(gcpdnstaskgroup Gcpdnstaskgroup) GcpdnstaskgroupAPIUpdateRequest {
	r.gcpdnstaskgroup = &gcpdnstaskgroup
	return r
}

// Enter the field names followed by comma
func (r GcpdnstaskgroupAPIUpdateRequest) ReturnFields(returnFields string) GcpdnstaskgroupAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r GcpdnstaskgroupAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) GcpdnstaskgroupAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r GcpdnstaskgroupAPIUpdateRequest) ReturnAsObject(returnAsObject int32) GcpdnstaskgroupAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r GcpdnstaskgroupAPIUpdateRequest) Execute() (*UpdateGcpdnstaskgroupResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a gcpdnstaskgroup object

Updates a specific gcpdnstaskgroup object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the gcpdnstaskgroup object
	@return GcpdnstaskgroupAPIUpdateRequest
*/
func (a *GcpdnstaskgroupAPIService) Update(ctx context.Context, reference string) GcpdnstaskgroupAPIUpdateRequest {
	return GcpdnstaskgroupAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateGcpdnstaskgroupResponse
func (a *GcpdnstaskgroupAPIService) UpdateExecute(r GcpdnstaskgroupAPIUpdateRequest) (*UpdateGcpdnstaskgroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateGcpdnstaskgroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GcpdnstaskgroupAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/gcpdnstaskgroup/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gcpdnstaskgroup == nil {
		return localVarReturnValue, nil, internal.ReportError("gcpdnstaskgroup is required and must be specified")
	}

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
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.gcpdnstaskgroup
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
