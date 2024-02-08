// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWorkspacesRequest wrapper for the ListWorkspaces operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListWorkspaces.go.html to see an example of how to use ListWorkspacesRequest.
type ListWorkspacesRequest struct {

	// The OCID of the compartment containing the resources you want to list.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The lifecycle state of a resource. When specified, the operation only returns resources that match the given lifecycle state. When not specified, all lifecycle states are processed as a match.
	LifecycleState WorkspaceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListWorkspacesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListWorkspacesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkspacesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkspacesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkspacesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkspacesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWorkspacesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkspaceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetWorkspaceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkspacesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWorkspacesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkspacesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWorkspacesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWorkspacesResponse wrapper for the ListWorkspaces operation
type ListWorkspacesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WorkspaceSummary instances
	Items []WorkspaceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkspacesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkspacesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkspacesSortOrderEnum Enum with underlying type: string
type ListWorkspacesSortOrderEnum string

// Set of constants representing the allowable values for ListWorkspacesSortOrderEnum
const (
	ListWorkspacesSortOrderAsc  ListWorkspacesSortOrderEnum = "ASC"
	ListWorkspacesSortOrderDesc ListWorkspacesSortOrderEnum = "DESC"
)

var mappingListWorkspacesSortOrderEnum = map[string]ListWorkspacesSortOrderEnum{
	"ASC":  ListWorkspacesSortOrderAsc,
	"DESC": ListWorkspacesSortOrderDesc,
}

var mappingListWorkspacesSortOrderEnumLowerCase = map[string]ListWorkspacesSortOrderEnum{
	"asc":  ListWorkspacesSortOrderAsc,
	"desc": ListWorkspacesSortOrderDesc,
}

// GetListWorkspacesSortOrderEnumValues Enumerates the set of values for ListWorkspacesSortOrderEnum
func GetListWorkspacesSortOrderEnumValues() []ListWorkspacesSortOrderEnum {
	values := make([]ListWorkspacesSortOrderEnum, 0)
	for _, v := range mappingListWorkspacesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkspacesSortOrderEnumStringValues Enumerates the set of values in String for ListWorkspacesSortOrderEnum
func GetListWorkspacesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWorkspacesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkspacesSortOrderEnum(val string) (ListWorkspacesSortOrderEnum, bool) {
	enum, ok := mappingListWorkspacesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkspacesSortByEnum Enum with underlying type: string
type ListWorkspacesSortByEnum string

// Set of constants representing the allowable values for ListWorkspacesSortByEnum
const (
	ListWorkspacesSortByTimeCreated ListWorkspacesSortByEnum = "TIME_CREATED"
	ListWorkspacesSortByDisplayName ListWorkspacesSortByEnum = "DISPLAY_NAME"
	ListWorkspacesSortByTimeUpdated ListWorkspacesSortByEnum = "TIME_UPDATED"
)

var mappingListWorkspacesSortByEnum = map[string]ListWorkspacesSortByEnum{
	"TIME_CREATED": ListWorkspacesSortByTimeCreated,
	"DISPLAY_NAME": ListWorkspacesSortByDisplayName,
	"TIME_UPDATED": ListWorkspacesSortByTimeUpdated,
}

var mappingListWorkspacesSortByEnumLowerCase = map[string]ListWorkspacesSortByEnum{
	"time_created": ListWorkspacesSortByTimeCreated,
	"display_name": ListWorkspacesSortByDisplayName,
	"time_updated": ListWorkspacesSortByTimeUpdated,
}

// GetListWorkspacesSortByEnumValues Enumerates the set of values for ListWorkspacesSortByEnum
func GetListWorkspacesSortByEnumValues() []ListWorkspacesSortByEnum {
	values := make([]ListWorkspacesSortByEnum, 0)
	for _, v := range mappingListWorkspacesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkspacesSortByEnumStringValues Enumerates the set of values in String for ListWorkspacesSortByEnum
func GetListWorkspacesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListWorkspacesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkspacesSortByEnum(val string) (ListWorkspacesSortByEnum, bool) {
	enum, ok := mappingListWorkspacesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
