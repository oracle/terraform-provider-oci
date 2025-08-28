// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSchedulesRequest wrapper for the ListSchedules operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListSchedules.go.html to see an example of how to use ListSchedulesRequest.
type ListSchedulesRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListSchedulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique Schedule identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListSchedulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListSchedulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSchedulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSchedulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSchedulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSchedulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSchedulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSchedulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSchedulesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSchedulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSchedulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSchedulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSchedulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSchedulesResponse wrapper for the ListSchedules operation
type ListSchedulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ScheduleSummary instances
	Items []ScheduleSummary `presentIn:"body"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSchedulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSchedulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSchedulesLifecycleStateEnum Enum with underlying type: string
type ListSchedulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListSchedulesLifecycleStateEnum
const (
	ListSchedulesLifecycleStateCreating ListSchedulesLifecycleStateEnum = "CREATING"
	ListSchedulesLifecycleStateActive   ListSchedulesLifecycleStateEnum = "ACTIVE"
	ListSchedulesLifecycleStateInactive ListSchedulesLifecycleStateEnum = "INACTIVE"
	ListSchedulesLifecycleStateUpdating ListSchedulesLifecycleStateEnum = "UPDATING"
	ListSchedulesLifecycleStateDeleting ListSchedulesLifecycleStateEnum = "DELETING"
	ListSchedulesLifecycleStateDeleted  ListSchedulesLifecycleStateEnum = "DELETED"
	ListSchedulesLifecycleStateFailed   ListSchedulesLifecycleStateEnum = "FAILED"
)

var mappingListSchedulesLifecycleStateEnum = map[string]ListSchedulesLifecycleStateEnum{
	"CREATING": ListSchedulesLifecycleStateCreating,
	"ACTIVE":   ListSchedulesLifecycleStateActive,
	"INACTIVE": ListSchedulesLifecycleStateInactive,
	"UPDATING": ListSchedulesLifecycleStateUpdating,
	"DELETING": ListSchedulesLifecycleStateDeleting,
	"DELETED":  ListSchedulesLifecycleStateDeleted,
	"FAILED":   ListSchedulesLifecycleStateFailed,
}

var mappingListSchedulesLifecycleStateEnumLowerCase = map[string]ListSchedulesLifecycleStateEnum{
	"creating": ListSchedulesLifecycleStateCreating,
	"active":   ListSchedulesLifecycleStateActive,
	"inactive": ListSchedulesLifecycleStateInactive,
	"updating": ListSchedulesLifecycleStateUpdating,
	"deleting": ListSchedulesLifecycleStateDeleting,
	"deleted":  ListSchedulesLifecycleStateDeleted,
	"failed":   ListSchedulesLifecycleStateFailed,
}

// GetListSchedulesLifecycleStateEnumValues Enumerates the set of values for ListSchedulesLifecycleStateEnum
func GetListSchedulesLifecycleStateEnumValues() []ListSchedulesLifecycleStateEnum {
	values := make([]ListSchedulesLifecycleStateEnum, 0)
	for _, v := range mappingListSchedulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchedulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListSchedulesLifecycleStateEnum
func GetListSchedulesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListSchedulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchedulesLifecycleStateEnum(val string) (ListSchedulesLifecycleStateEnum, bool) {
	enum, ok := mappingListSchedulesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSchedulesSortOrderEnum Enum with underlying type: string
type ListSchedulesSortOrderEnum string

// Set of constants representing the allowable values for ListSchedulesSortOrderEnum
const (
	ListSchedulesSortOrderAsc  ListSchedulesSortOrderEnum = "ASC"
	ListSchedulesSortOrderDesc ListSchedulesSortOrderEnum = "DESC"
)

var mappingListSchedulesSortOrderEnum = map[string]ListSchedulesSortOrderEnum{
	"ASC":  ListSchedulesSortOrderAsc,
	"DESC": ListSchedulesSortOrderDesc,
}

var mappingListSchedulesSortOrderEnumLowerCase = map[string]ListSchedulesSortOrderEnum{
	"asc":  ListSchedulesSortOrderAsc,
	"desc": ListSchedulesSortOrderDesc,
}

// GetListSchedulesSortOrderEnumValues Enumerates the set of values for ListSchedulesSortOrderEnum
func GetListSchedulesSortOrderEnumValues() []ListSchedulesSortOrderEnum {
	values := make([]ListSchedulesSortOrderEnum, 0)
	for _, v := range mappingListSchedulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchedulesSortOrderEnumStringValues Enumerates the set of values in String for ListSchedulesSortOrderEnum
func GetListSchedulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSchedulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchedulesSortOrderEnum(val string) (ListSchedulesSortOrderEnum, bool) {
	enum, ok := mappingListSchedulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSchedulesSortByEnum Enum with underlying type: string
type ListSchedulesSortByEnum string

// Set of constants representing the allowable values for ListSchedulesSortByEnum
const (
	ListSchedulesSortByTimecreated ListSchedulesSortByEnum = "timeCreated"
	ListSchedulesSortByDisplayname ListSchedulesSortByEnum = "displayName"
)

var mappingListSchedulesSortByEnum = map[string]ListSchedulesSortByEnum{
	"timeCreated": ListSchedulesSortByTimecreated,
	"displayName": ListSchedulesSortByDisplayname,
}

var mappingListSchedulesSortByEnumLowerCase = map[string]ListSchedulesSortByEnum{
	"timecreated": ListSchedulesSortByTimecreated,
	"displayname": ListSchedulesSortByDisplayname,
}

// GetListSchedulesSortByEnumValues Enumerates the set of values for ListSchedulesSortByEnum
func GetListSchedulesSortByEnumValues() []ListSchedulesSortByEnum {
	values := make([]ListSchedulesSortByEnum, 0)
	for _, v := range mappingListSchedulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchedulesSortByEnumStringValues Enumerates the set of values in String for ListSchedulesSortByEnum
func GetListSchedulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSchedulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchedulesSortByEnum(val string) (ListSchedulesSortByEnum, bool) {
	enum, ok := mappingListSchedulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
