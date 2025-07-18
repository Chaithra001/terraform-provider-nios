/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type MemberThreatprotectionAPI interface {
	/*
		List Retrieve member:threatprotection objects

		Returns a list of member:threatprotection objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return MemberThreatprotectionAPIListRequest
	*/
	List(ctx context.Context) MemberThreatprotectionAPIListRequest

	// ListExecute executes the request
	//  @return ListMemberThreatprotectionResponse
	ListExecute(r MemberThreatprotectionAPIListRequest) (*ListMemberThreatprotectionResponse, *http.Response, error)
	/*
		Read Get a specific member:threatprotection object

		Returns a specific member:threatprotection object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the member:threatprotection object
		@return MemberThreatprotectionAPIReadRequest
	*/
	Read(ctx context.Context, reference string) MemberThreatprotectionAPIReadRequest

	// ReadExecute executes the request
	//  @return GetMemberThreatprotectionResponse
	ReadExecute(r MemberThreatprotectionAPIReadRequest) (*GetMemberThreatprotectionResponse, *http.Response, error)
	/*
		Update Update a member:threatprotection object

		Updates a specific member:threatprotection object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the member:threatprotection object
		@return MemberThreatprotectionAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) MemberThreatprotectionAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateMemberThreatprotectionResponse
	UpdateExecute(r MemberThreatprotectionAPIUpdateRequest) (*UpdateMemberThreatprotectionResponse, *http.Response, error)
}

// MemberThreatprotectionAPIService MemberThreatprotectionAPI service
type MemberThreatprotectionAPIService internal.Service

type MemberThreatprotectionAPIListRequest struct {
	ctx              context.Context
	ApiService       MemberThreatprotectionAPI
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
func (r MemberThreatprotectionAPIListRequest) ReturnFields(returnFields string) MemberThreatprotectionAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r MemberThreatprotectionAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) MemberThreatprotectionAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r MemberThreatprotectionAPIListRequest) MaxResults(maxResults int32) MemberThreatprotectionAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r MemberThreatprotectionAPIListRequest) ReturnAsObject(returnAsObject int32) MemberThreatprotectionAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r MemberThreatprotectionAPIListRequest) Paging(paging int32) MemberThreatprotectionAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r MemberThreatprotectionAPIListRequest) PageId(pageId string) MemberThreatprotectionAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r MemberThreatprotectionAPIListRequest) Filters(filters map[string]interface{}) MemberThreatprotectionAPIListRequest {
	r.filters = &filters
	return r
}

func (r MemberThreatprotectionAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) MemberThreatprotectionAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r MemberThreatprotectionAPIListRequest) Execute() (*ListMemberThreatprotectionResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve member:threatprotection objects

Returns a list of member:threatprotection objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MemberThreatprotectionAPIListRequest
*/
func (a *MemberThreatprotectionAPIService) List(ctx context.Context) MemberThreatprotectionAPIListRequest {
	return MemberThreatprotectionAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListMemberThreatprotectionResponse
func (a *MemberThreatprotectionAPIService) ListExecute(r MemberThreatprotectionAPIListRequest) (*ListMemberThreatprotectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListMemberThreatprotectionResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "MemberThreatprotectionAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/member:threatprotection"

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

type MemberThreatprotectionAPIReadRequest struct {
	ctx              context.Context
	ApiService       MemberThreatprotectionAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r MemberThreatprotectionAPIReadRequest) ReturnFields(returnFields string) MemberThreatprotectionAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r MemberThreatprotectionAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) MemberThreatprotectionAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r MemberThreatprotectionAPIReadRequest) ReturnAsObject(returnAsObject int32) MemberThreatprotectionAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r MemberThreatprotectionAPIReadRequest) Execute() (*GetMemberThreatprotectionResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific member:threatprotection object

Returns a specific member:threatprotection object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the member:threatprotection object
	@return MemberThreatprotectionAPIReadRequest
*/
func (a *MemberThreatprotectionAPIService) Read(ctx context.Context, reference string) MemberThreatprotectionAPIReadRequest {
	return MemberThreatprotectionAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetMemberThreatprotectionResponse
func (a *MemberThreatprotectionAPIService) ReadExecute(r MemberThreatprotectionAPIReadRequest) (*GetMemberThreatprotectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetMemberThreatprotectionResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "MemberThreatprotectionAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/member:threatprotection/{reference}"
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

type MemberThreatprotectionAPIUpdateRequest struct {
	ctx                    context.Context
	ApiService             MemberThreatprotectionAPI
	reference              string
	memberThreatprotection *MemberThreatprotection
	returnFields           *string
	returnFieldsPlus       *string
	returnAsObject         *int32
}

// Object data to update
func (r MemberThreatprotectionAPIUpdateRequest) MemberThreatprotection(memberThreatprotection MemberThreatprotection) MemberThreatprotectionAPIUpdateRequest {
	r.memberThreatprotection = &memberThreatprotection
	return r
}

// Enter the field names followed by comma
func (r MemberThreatprotectionAPIUpdateRequest) ReturnFields(returnFields string) MemberThreatprotectionAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r MemberThreatprotectionAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) MemberThreatprotectionAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r MemberThreatprotectionAPIUpdateRequest) ReturnAsObject(returnAsObject int32) MemberThreatprotectionAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r MemberThreatprotectionAPIUpdateRequest) Execute() (*UpdateMemberThreatprotectionResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a member:threatprotection object

Updates a specific member:threatprotection object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the member:threatprotection object
	@return MemberThreatprotectionAPIUpdateRequest
*/
func (a *MemberThreatprotectionAPIService) Update(ctx context.Context, reference string) MemberThreatprotectionAPIUpdateRequest {
	return MemberThreatprotectionAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateMemberThreatprotectionResponse
func (a *MemberThreatprotectionAPIService) UpdateExecute(r MemberThreatprotectionAPIUpdateRequest) (*UpdateMemberThreatprotectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateMemberThreatprotectionResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "MemberThreatprotectionAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/member:threatprotection/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.memberThreatprotection == nil {
		return localVarReturnValue, nil, internal.ReportError("memberThreatprotection is required and must be specified")
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
	localVarPostBody = r.memberThreatprotection
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
