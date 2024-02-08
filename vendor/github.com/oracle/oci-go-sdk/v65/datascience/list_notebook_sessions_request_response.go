// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListNotebookSessionsRequest wrapper for the ListNotebookSessions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListNotebookSessions.go.html to see an example of how to use ListNotebookSessionsRequest.
type ListNotebookSessionsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListNotebookSessionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListNotebookSessionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown
	// in descending order. When you sort by `displayName`, results are
	// shown in ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListNotebookSessionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNotebookSessionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNotebookSessionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNotebookSessionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNotebookSessionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNotebookSessionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNotebookSessionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListNotebookSessionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNotebookSessionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNotebookSessionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNotebookSessionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNotebookSessionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNotebookSessionsResponse wrapper for the ListNotebookSessions operation
type ListNotebookSessionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []NotebookSessionSummary instances
	Items []NotebookSessionSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListNotebookSessionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNotebookSessionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNotebookSessionsLifecycleStateEnum Enum with underlying type: string
type ListNotebookSessionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListNotebookSessionsLifecycleStateEnum
const (
	ListNotebookSessionsLifecycleStateCreating ListNotebookSessionsLifecycleStateEnum = "CREATING"
	ListNotebookSessionsLifecycleStateActive   ListNotebookSessionsLifecycleStateEnum = "ACTIVE"
	ListNotebookSessionsLifecycleStateDeleting ListNotebookSessionsLifecycleStateEnum = "DELETING"
	ListNotebookSessionsLifecycleStateDeleted  ListNotebookSessionsLifecycleStateEnum = "DELETED"
	ListNotebookSessionsLifecycleStateFailed   ListNotebookSessionsLifecycleStateEnum = "FAILED"
	ListNotebookSessionsLifecycleStateInactive ListNotebookSessionsLifecycleStateEnum = "INACTIVE"
	ListNotebookSessionsLifecycleStateUpdating ListNotebookSessionsLifecycleStateEnum = "UPDATING"
)

var mappingListNotebookSessionsLifecycleStateEnum = map[string]ListNotebookSessionsLifecycleStateEnum{
	"CREATING": ListNotebookSessionsLifecycleStateCreating,
	"ACTIVE":   ListNotebookSessionsLifecycleStateActive,
	"DELETING": ListNotebookSessionsLifecycleStateDeleting,
	"DELETED":  ListNotebookSessionsLifecycleStateDeleted,
	"FAILED":   ListNotebookSessionsLifecycleStateFailed,
	"INACTIVE": ListNotebookSessionsLifecycleStateInactive,
	"UPDATING": ListNotebookSessionsLifecycleStateUpdating,
}

var mappingListNotebookSessionsLifecycleStateEnumLowerCase = map[string]ListNotebookSessionsLifecycleStateEnum{
	"creating": ListNotebookSessionsLifecycleStateCreating,
	"active":   ListNotebookSessionsLifecycleStateActive,
	"deleting": ListNotebookSessionsLifecycleStateDeleting,
	"deleted":  ListNotebookSessionsLifecycleStateDeleted,
	"failed":   ListNotebookSessionsLifecycleStateFailed,
	"inactive": ListNotebookSessionsLifecycleStateInactive,
	"updating": ListNotebookSessionsLifecycleStateUpdating,
}

// GetListNotebookSessionsLifecycleStateEnumValues Enumerates the set of values for ListNotebookSessionsLifecycleStateEnum
func GetListNotebookSessionsLifecycleStateEnumValues() []ListNotebookSessionsLifecycleStateEnum {
	values := make([]ListNotebookSessionsLifecycleStateEnum, 0)
	for _, v := range mappingListNotebookSessionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListNotebookSessionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListNotebookSessionsLifecycleStateEnum
func GetListNotebookSessionsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACTIVE",
		"UPDATING",
	}
}

// GetMappingListNotebookSessionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNotebookSessionsLifecycleStateEnum(val string) (ListNotebookSessionsLifecycleStateEnum, bool) {
	enum, ok := mappingListNotebookSessionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNotebookSessionsSortOrderEnum Enum with underlying type: string
type ListNotebookSessionsSortOrderEnum string

// Set of constants representing the allowable values for ListNotebookSessionsSortOrderEnum
const (
	ListNotebookSessionsSortOrderAsc  ListNotebookSessionsSortOrderEnum = "ASC"
	ListNotebookSessionsSortOrderDesc ListNotebookSessionsSortOrderEnum = "DESC"
)

var mappingListNotebookSessionsSortOrderEnum = map[string]ListNotebookSessionsSortOrderEnum{
	"ASC":  ListNotebookSessionsSortOrderAsc,
	"DESC": ListNotebookSessionsSortOrderDesc,
}

var mappingListNotebookSessionsSortOrderEnumLowerCase = map[string]ListNotebookSessionsSortOrderEnum{
	"asc":  ListNotebookSessionsSortOrderAsc,
	"desc": ListNotebookSessionsSortOrderDesc,
}

// GetListNotebookSessionsSortOrderEnumValues Enumerates the set of values for ListNotebookSessionsSortOrderEnum
func GetListNotebookSessionsSortOrderEnumValues() []ListNotebookSessionsSortOrderEnum {
	values := make([]ListNotebookSessionsSortOrderEnum, 0)
	for _, v := range mappingListNotebookSessionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNotebookSessionsSortOrderEnumStringValues Enumerates the set of values in String for ListNotebookSessionsSortOrderEnum
func GetListNotebookSessionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNotebookSessionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNotebookSessionsSortOrderEnum(val string) (ListNotebookSessionsSortOrderEnum, bool) {
	enum, ok := mappingListNotebookSessionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNotebookSessionsSortByEnum Enum with underlying type: string
type ListNotebookSessionsSortByEnum string

// Set of constants representing the allowable values for ListNotebookSessionsSortByEnum
const (
	ListNotebookSessionsSortByTimecreated ListNotebookSessionsSortByEnum = "timeCreated"
	ListNotebookSessionsSortByDisplayname ListNotebookSessionsSortByEnum = "displayName"
)

var mappingListNotebookSessionsSortByEnum = map[string]ListNotebookSessionsSortByEnum{
	"timeCreated": ListNotebookSessionsSortByTimecreated,
	"displayName": ListNotebookSessionsSortByDisplayname,
}

var mappingListNotebookSessionsSortByEnumLowerCase = map[string]ListNotebookSessionsSortByEnum{
	"timecreated": ListNotebookSessionsSortByTimecreated,
	"displayname": ListNotebookSessionsSortByDisplayname,
}

// GetListNotebookSessionsSortByEnumValues Enumerates the set of values for ListNotebookSessionsSortByEnum
func GetListNotebookSessionsSortByEnumValues() []ListNotebookSessionsSortByEnum {
	values := make([]ListNotebookSessionsSortByEnum, 0)
	for _, v := range mappingListNotebookSessionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNotebookSessionsSortByEnumStringValues Enumerates the set of values in String for ListNotebookSessionsSortByEnum
func GetListNotebookSessionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListNotebookSessionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNotebookSessionsSortByEnum(val string) (ListNotebookSessionsSortByEnum, bool) {
	enum, ok := mappingListNotebookSessionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
