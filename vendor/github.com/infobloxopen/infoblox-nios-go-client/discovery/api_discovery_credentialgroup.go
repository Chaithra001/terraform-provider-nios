/*
Infoblox DISCOVERY API

OpenAPI specification for Infoblox NIOS WAPI DISCOVERY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discovery

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type DiscoveryCredentialgroupAPI interface {
	/*
		Create Create a discovery:credentialgroup object

		Creates a new discovery:credentialgroup object

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return DiscoveryCredentialgroupAPICreateRequest
	*/
	Create(ctx context.Context) DiscoveryCredentialgroupAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateDiscoveryCredentialgroupResponse
	CreateExecute(r DiscoveryCredentialgroupAPICreateRequest) (*CreateDiscoveryCredentialgroupResponse, *http.Response, error)
	/*
		Delete Delete a discovery:credentialgroup object

		Deletes a specific discovery:credentialgroup object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the discovery:credentialgroup object
		@return DiscoveryCredentialgroupAPIDeleteRequest
	*/
	Delete(ctx context.Context, reference string) DiscoveryCredentialgroupAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r DiscoveryCredentialgroupAPIDeleteRequest) (*http.Response, error)
	/*
		List Retrieve discovery:credentialgroup objects

		Returns a list of discovery:credentialgroup objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return DiscoveryCredentialgroupAPIListRequest
	*/
	List(ctx context.Context) DiscoveryCredentialgroupAPIListRequest

	// ListExecute executes the request
	//  @return ListDiscoveryCredentialgroupResponse
	ListExecute(r DiscoveryCredentialgroupAPIListRequest) (*ListDiscoveryCredentialgroupResponse, *http.Response, error)
	/*
		Read Get a specific discovery:credentialgroup object

		Returns a specific discovery:credentialgroup object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the discovery:credentialgroup object
		@return DiscoveryCredentialgroupAPIReadRequest
	*/
	Read(ctx context.Context, reference string) DiscoveryCredentialgroupAPIReadRequest

	// ReadExecute executes the request
	//  @return GetDiscoveryCredentialgroupResponse
	ReadExecute(r DiscoveryCredentialgroupAPIReadRequest) (*GetDiscoveryCredentialgroupResponse, *http.Response, error)
	/*
		Update Update a discovery:credentialgroup object

		Updates a specific discovery:credentialgroup object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the discovery:credentialgroup object
		@return DiscoveryCredentialgroupAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) DiscoveryCredentialgroupAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateDiscoveryCredentialgroupResponse
	UpdateExecute(r DiscoveryCredentialgroupAPIUpdateRequest) (*UpdateDiscoveryCredentialgroupResponse, *http.Response, error)
}

// DiscoveryCredentialgroupAPIService DiscoveryCredentialgroupAPI service
type DiscoveryCredentialgroupAPIService internal.Service

type DiscoveryCredentialgroupAPICreateRequest struct {
	ctx                      context.Context
	ApiService               DiscoveryCredentialgroupAPI
	discoveryCredentialgroup *DiscoveryCredentialgroup
	returnFields             *string
	returnFieldsPlus         *string
	returnAsObject           *int32
}

// Object data to create
func (r DiscoveryCredentialgroupAPICreateRequest) DiscoveryCredentialgroup(discoveryCredentialgroup DiscoveryCredentialgroup) DiscoveryCredentialgroupAPICreateRequest {
	r.discoveryCredentialgroup = &discoveryCredentialgroup
	return r
}

// Enter the field names followed by comma
func (r DiscoveryCredentialgroupAPICreateRequest) ReturnFields(returnFields string) DiscoveryCredentialgroupAPICreateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DiscoveryCredentialgroupAPICreateRequest) ReturnFieldsPlus(returnFieldsPlus string) DiscoveryCredentialgroupAPICreateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DiscoveryCredentialgroupAPICreateRequest) ReturnAsObject(returnAsObject int32) DiscoveryCredentialgroupAPICreateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DiscoveryCredentialgroupAPICreateRequest) Execute() (*CreateDiscoveryCredentialgroupResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a discovery:credentialgroup object

Creates a new discovery:credentialgroup object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DiscoveryCredentialgroupAPICreateRequest
*/
func (a *DiscoveryCredentialgroupAPIService) Create(ctx context.Context) DiscoveryCredentialgroupAPICreateRequest {
	return DiscoveryCredentialgroupAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateDiscoveryCredentialgroupResponse
func (a *DiscoveryCredentialgroupAPIService) CreateExecute(r DiscoveryCredentialgroupAPICreateRequest) (*CreateDiscoveryCredentialgroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateDiscoveryCredentialgroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DiscoveryCredentialgroupAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/discovery:credentialgroup"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.discoveryCredentialgroup == nil {
		return localVarReturnValue, nil, internal.ReportError("discoveryCredentialgroup is required and must be specified")
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
	localVarPostBody = r.discoveryCredentialgroup
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

type DiscoveryCredentialgroupAPIDeleteRequest struct {
	ctx        context.Context
	ApiService DiscoveryCredentialgroupAPI
	reference  string
}

func (r DiscoveryCredentialgroupAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a discovery:credentialgroup object

Deletes a specific discovery:credentialgroup object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the discovery:credentialgroup object
	@return DiscoveryCredentialgroupAPIDeleteRequest
*/
func (a *DiscoveryCredentialgroupAPIService) Delete(ctx context.Context, reference string) DiscoveryCredentialgroupAPIDeleteRequest {
	return DiscoveryCredentialgroupAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
func (a *DiscoveryCredentialgroupAPIService) DeleteExecute(r DiscoveryCredentialgroupAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DiscoveryCredentialgroupAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/discovery:credentialgroup/{reference}"
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

type DiscoveryCredentialgroupAPIListRequest struct {
	ctx              context.Context
	ApiService       DiscoveryCredentialgroupAPI
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
func (r DiscoveryCredentialgroupAPIListRequest) ReturnFields(returnFields string) DiscoveryCredentialgroupAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DiscoveryCredentialgroupAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) DiscoveryCredentialgroupAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r DiscoveryCredentialgroupAPIListRequest) MaxResults(maxResults int32) DiscoveryCredentialgroupAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r DiscoveryCredentialgroupAPIListRequest) ReturnAsObject(returnAsObject int32) DiscoveryCredentialgroupAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r DiscoveryCredentialgroupAPIListRequest) Paging(paging int32) DiscoveryCredentialgroupAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r DiscoveryCredentialgroupAPIListRequest) PageId(pageId string) DiscoveryCredentialgroupAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r DiscoveryCredentialgroupAPIListRequest) Filters(filters map[string]interface{}) DiscoveryCredentialgroupAPIListRequest {
	r.filters = &filters
	return r
}

func (r DiscoveryCredentialgroupAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) DiscoveryCredentialgroupAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r DiscoveryCredentialgroupAPIListRequest) Execute() (*ListDiscoveryCredentialgroupResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve discovery:credentialgroup objects

Returns a list of discovery:credentialgroup objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DiscoveryCredentialgroupAPIListRequest
*/
func (a *DiscoveryCredentialgroupAPIService) List(ctx context.Context) DiscoveryCredentialgroupAPIListRequest {
	return DiscoveryCredentialgroupAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListDiscoveryCredentialgroupResponse
func (a *DiscoveryCredentialgroupAPIService) ListExecute(r DiscoveryCredentialgroupAPIListRequest) (*ListDiscoveryCredentialgroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListDiscoveryCredentialgroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DiscoveryCredentialgroupAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/discovery:credentialgroup"

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

type DiscoveryCredentialgroupAPIReadRequest struct {
	ctx              context.Context
	ApiService       DiscoveryCredentialgroupAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r DiscoveryCredentialgroupAPIReadRequest) ReturnFields(returnFields string) DiscoveryCredentialgroupAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DiscoveryCredentialgroupAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) DiscoveryCredentialgroupAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DiscoveryCredentialgroupAPIReadRequest) ReturnAsObject(returnAsObject int32) DiscoveryCredentialgroupAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DiscoveryCredentialgroupAPIReadRequest) Execute() (*GetDiscoveryCredentialgroupResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific discovery:credentialgroup object

Returns a specific discovery:credentialgroup object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the discovery:credentialgroup object
	@return DiscoveryCredentialgroupAPIReadRequest
*/
func (a *DiscoveryCredentialgroupAPIService) Read(ctx context.Context, reference string) DiscoveryCredentialgroupAPIReadRequest {
	return DiscoveryCredentialgroupAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetDiscoveryCredentialgroupResponse
func (a *DiscoveryCredentialgroupAPIService) ReadExecute(r DiscoveryCredentialgroupAPIReadRequest) (*GetDiscoveryCredentialgroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetDiscoveryCredentialgroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DiscoveryCredentialgroupAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/discovery:credentialgroup/{reference}"
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

type DiscoveryCredentialgroupAPIUpdateRequest struct {
	ctx                      context.Context
	ApiService               DiscoveryCredentialgroupAPI
	reference                string
	discoveryCredentialgroup *DiscoveryCredentialgroup
	returnFields             *string
	returnFieldsPlus         *string
	returnAsObject           *int32
}

// Object data to update
func (r DiscoveryCredentialgroupAPIUpdateRequest) DiscoveryCredentialgroup(discoveryCredentialgroup DiscoveryCredentialgroup) DiscoveryCredentialgroupAPIUpdateRequest {
	r.discoveryCredentialgroup = &discoveryCredentialgroup
	return r
}

// Enter the field names followed by comma
func (r DiscoveryCredentialgroupAPIUpdateRequest) ReturnFields(returnFields string) DiscoveryCredentialgroupAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DiscoveryCredentialgroupAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) DiscoveryCredentialgroupAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DiscoveryCredentialgroupAPIUpdateRequest) ReturnAsObject(returnAsObject int32) DiscoveryCredentialgroupAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DiscoveryCredentialgroupAPIUpdateRequest) Execute() (*UpdateDiscoveryCredentialgroupResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a discovery:credentialgroup object

Updates a specific discovery:credentialgroup object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the discovery:credentialgroup object
	@return DiscoveryCredentialgroupAPIUpdateRequest
*/
func (a *DiscoveryCredentialgroupAPIService) Update(ctx context.Context, reference string) DiscoveryCredentialgroupAPIUpdateRequest {
	return DiscoveryCredentialgroupAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateDiscoveryCredentialgroupResponse
func (a *DiscoveryCredentialgroupAPIService) UpdateExecute(r DiscoveryCredentialgroupAPIUpdateRequest) (*UpdateDiscoveryCredentialgroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateDiscoveryCredentialgroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DiscoveryCredentialgroupAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/discovery:credentialgroup/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.discoveryCredentialgroup == nil {
		return localVarReturnValue, nil, internal.ReportError("discoveryCredentialgroup is required and must be specified")
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
	localVarPostBody = r.discoveryCredentialgroup
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
