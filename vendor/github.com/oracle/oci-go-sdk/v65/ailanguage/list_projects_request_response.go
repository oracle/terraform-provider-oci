// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ailanguage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProjectsRequest wrapper for the ListProjects operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/ListProjects.go.html to see an example of how to use ListProjectsRequest.
type ListProjectsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ProjectLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The ID of the project for which to list the objects.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListProjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown
	// in descending order. When you sort by `displayName`, the results are
	// shown in ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListProjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProjectLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetProjectLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProjectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProjectsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProjectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProjectsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProjectsResponse wrapper for the ListProjects operation
type ListProjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProjectCollection instances
	ProjectCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProjectsSortOrderEnum Enum with underlying type: string
type ListProjectsSortOrderEnum string

// Set of constants representing the allowable values for ListProjectsSortOrderEnum
const (
	ListProjectsSortOrderAsc  ListProjectsSortOrderEnum = "ASC"
	ListProjectsSortOrderDesc ListProjectsSortOrderEnum = "DESC"
)

var mappingListProjectsSortOrderEnum = map[string]ListProjectsSortOrderEnum{
	"ASC":  ListProjectsSortOrderAsc,
	"DESC": ListProjectsSortOrderDesc,
}

var mappingListProjectsSortOrderEnumLowerCase = map[string]ListProjectsSortOrderEnum{
	"asc":  ListProjectsSortOrderAsc,
	"desc": ListProjectsSortOrderDesc,
}

// GetListProjectsSortOrderEnumValues Enumerates the set of values for ListProjectsSortOrderEnum
func GetListProjectsSortOrderEnumValues() []ListProjectsSortOrderEnum {
	values := make([]ListProjectsSortOrderEnum, 0)
	for _, v := range mappingListProjectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProjectsSortOrderEnumStringValues Enumerates the set of values in String for ListProjectsSortOrderEnum
func GetListProjectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProjectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProjectsSortOrderEnum(val string) (ListProjectsSortOrderEnum, bool) {
	enum, ok := mappingListProjectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProjectsSortByEnum Enum with underlying type: string
type ListProjectsSortByEnum string

// Set of constants representing the allowable values for ListProjectsSortByEnum
const (
	ListProjectsSortByTimecreated ListProjectsSortByEnum = "timeCreated"
	ListProjectsSortByDisplayname ListProjectsSortByEnum = "displayName"
)

var mappingListProjectsSortByEnum = map[string]ListProjectsSortByEnum{
	"timeCreated": ListProjectsSortByTimecreated,
	"displayName": ListProjectsSortByDisplayname,
}

var mappingListProjectsSortByEnumLowerCase = map[string]ListProjectsSortByEnum{
	"timecreated": ListProjectsSortByTimecreated,
	"displayname": ListProjectsSortByDisplayname,
}

// GetListProjectsSortByEnumValues Enumerates the set of values for ListProjectsSortByEnum
func GetListProjectsSortByEnumValues() []ListProjectsSortByEnum {
	values := make([]ListProjectsSortByEnum, 0)
	for _, v := range mappingListProjectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProjectsSortByEnumStringValues Enumerates the set of values in String for ListProjectsSortByEnum
func GetListProjectsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListProjectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProjectsSortByEnum(val string) (ListProjectsSortByEnum, bool) {
	enum, ok := mappingListProjectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
