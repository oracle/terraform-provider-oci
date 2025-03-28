// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOpaInstancesRequest wrapper for the ListOpaInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opa/ListOpaInstances.go.html to see an example of how to use ListOpaInstancesRequest.
type ListOpaInstancesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState OpaInstanceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique OpaInstance identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOpaInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListOpaInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOpaInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOpaInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOpaInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOpaInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOpaInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOpaInstanceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOpaInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOpaInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOpaInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOpaInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOpaInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOpaInstancesResponse wrapper for the ListOpaInstances operation
type ListOpaInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OpaInstanceCollection instances
	OpaInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOpaInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOpaInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOpaInstancesSortOrderEnum Enum with underlying type: string
type ListOpaInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListOpaInstancesSortOrderEnum
const (
	ListOpaInstancesSortOrderAsc  ListOpaInstancesSortOrderEnum = "ASC"
	ListOpaInstancesSortOrderDesc ListOpaInstancesSortOrderEnum = "DESC"
)

var mappingListOpaInstancesSortOrderEnum = map[string]ListOpaInstancesSortOrderEnum{
	"ASC":  ListOpaInstancesSortOrderAsc,
	"DESC": ListOpaInstancesSortOrderDesc,
}

var mappingListOpaInstancesSortOrderEnumLowerCase = map[string]ListOpaInstancesSortOrderEnum{
	"asc":  ListOpaInstancesSortOrderAsc,
	"desc": ListOpaInstancesSortOrderDesc,
}

// GetListOpaInstancesSortOrderEnumValues Enumerates the set of values for ListOpaInstancesSortOrderEnum
func GetListOpaInstancesSortOrderEnumValues() []ListOpaInstancesSortOrderEnum {
	values := make([]ListOpaInstancesSortOrderEnum, 0)
	for _, v := range mappingListOpaInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpaInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListOpaInstancesSortOrderEnum
func GetListOpaInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOpaInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpaInstancesSortOrderEnum(val string) (ListOpaInstancesSortOrderEnum, bool) {
	enum, ok := mappingListOpaInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOpaInstancesSortByEnum Enum with underlying type: string
type ListOpaInstancesSortByEnum string

// Set of constants representing the allowable values for ListOpaInstancesSortByEnum
const (
	ListOpaInstancesSortByTimecreated ListOpaInstancesSortByEnum = "timeCreated"
	ListOpaInstancesSortByDisplayname ListOpaInstancesSortByEnum = "displayName"
)

var mappingListOpaInstancesSortByEnum = map[string]ListOpaInstancesSortByEnum{
	"timeCreated": ListOpaInstancesSortByTimecreated,
	"displayName": ListOpaInstancesSortByDisplayname,
}

var mappingListOpaInstancesSortByEnumLowerCase = map[string]ListOpaInstancesSortByEnum{
	"timecreated": ListOpaInstancesSortByTimecreated,
	"displayname": ListOpaInstancesSortByDisplayname,
}

// GetListOpaInstancesSortByEnumValues Enumerates the set of values for ListOpaInstancesSortByEnum
func GetListOpaInstancesSortByEnumValues() []ListOpaInstancesSortByEnum {
	values := make([]ListOpaInstancesSortByEnum, 0)
	for _, v := range mappingListOpaInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpaInstancesSortByEnumStringValues Enumerates the set of values in String for ListOpaInstancesSortByEnum
func GetListOpaInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOpaInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpaInstancesSortByEnum(val string) (ListOpaInstancesSortByEnum, bool) {
	enum, ok := mappingListOpaInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
