// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMaintenanceWindowsRequest wrapper for the ListMaintenanceWindows operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListMaintenanceWindows.go.html to see an example of how to use ListMaintenanceWindowsRequest.
type ListMaintenanceWindowsRequest struct {

	// The ID of the compartment in which data is listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return maintenance windows that match exact resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return maintenance windows with matching lifecycleDetails.
	LifecycleDetails ListMaintenanceWindowsLifecycleDetailsEnum `mandatory:"false" contributesTo:"query" name:"lifecycleDetails" omitEmpty:"true"`

	// A filter to return only maintenance windows with matching lifecycleState.
	Status ListMaintenanceWindowsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for mainteance window name is ascending.
	SortBy ListMaintenanceWindowsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMaintenanceWindowsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaintenanceWindowsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaintenanceWindowsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaintenanceWindowsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaintenanceWindowsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaintenanceWindowsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaintenanceWindowsLifecycleDetailsEnum(string(request.LifecycleDetails)); !ok && request.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", request.LifecycleDetails, strings.Join(GetListMaintenanceWindowsLifecycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceWindowsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListMaintenanceWindowsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceWindowsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaintenanceWindowsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceWindowsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaintenanceWindowsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaintenanceWindowsResponse wrapper for the ListMaintenanceWindows operation
type ListMaintenanceWindowsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaintenanceWindowCollection instances
	MaintenanceWindowCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. The total number of items in the result.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListMaintenanceWindowsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaintenanceWindowsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaintenanceWindowsLifecycleDetailsEnum Enum with underlying type: string
type ListMaintenanceWindowsLifecycleDetailsEnum string

// Set of constants representing the allowable values for ListMaintenanceWindowsLifecycleDetailsEnum
const (
	ListMaintenanceWindowsLifecycleDetailsInProgress ListMaintenanceWindowsLifecycleDetailsEnum = "IN_PROGRESS"
	ListMaintenanceWindowsLifecycleDetailsScheduled  ListMaintenanceWindowsLifecycleDetailsEnum = "SCHEDULED"
	ListMaintenanceWindowsLifecycleDetailsCompleted  ListMaintenanceWindowsLifecycleDetailsEnum = "COMPLETED"
)

var mappingListMaintenanceWindowsLifecycleDetailsEnum = map[string]ListMaintenanceWindowsLifecycleDetailsEnum{
	"IN_PROGRESS": ListMaintenanceWindowsLifecycleDetailsInProgress,
	"SCHEDULED":   ListMaintenanceWindowsLifecycleDetailsScheduled,
	"COMPLETED":   ListMaintenanceWindowsLifecycleDetailsCompleted,
}

var mappingListMaintenanceWindowsLifecycleDetailsEnumLowerCase = map[string]ListMaintenanceWindowsLifecycleDetailsEnum{
	"in_progress": ListMaintenanceWindowsLifecycleDetailsInProgress,
	"scheduled":   ListMaintenanceWindowsLifecycleDetailsScheduled,
	"completed":   ListMaintenanceWindowsLifecycleDetailsCompleted,
}

// GetListMaintenanceWindowsLifecycleDetailsEnumValues Enumerates the set of values for ListMaintenanceWindowsLifecycleDetailsEnum
func GetListMaintenanceWindowsLifecycleDetailsEnumValues() []ListMaintenanceWindowsLifecycleDetailsEnum {
	values := make([]ListMaintenanceWindowsLifecycleDetailsEnum, 0)
	for _, v := range mappingListMaintenanceWindowsLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceWindowsLifecycleDetailsEnumStringValues Enumerates the set of values in String for ListMaintenanceWindowsLifecycleDetailsEnum
func GetListMaintenanceWindowsLifecycleDetailsEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SCHEDULED",
		"COMPLETED",
	}
}

// GetMappingListMaintenanceWindowsLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceWindowsLifecycleDetailsEnum(val string) (ListMaintenanceWindowsLifecycleDetailsEnum, bool) {
	enum, ok := mappingListMaintenanceWindowsLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceWindowsStatusEnum Enum with underlying type: string
type ListMaintenanceWindowsStatusEnum string

// Set of constants representing the allowable values for ListMaintenanceWindowsStatusEnum
const (
	ListMaintenanceWindowsStatusCreating       ListMaintenanceWindowsStatusEnum = "CREATING"
	ListMaintenanceWindowsStatusUpdating       ListMaintenanceWindowsStatusEnum = "UPDATING"
	ListMaintenanceWindowsStatusInactive       ListMaintenanceWindowsStatusEnum = "INACTIVE"
	ListMaintenanceWindowsStatusActive         ListMaintenanceWindowsStatusEnum = "ACTIVE"
	ListMaintenanceWindowsStatusDeleting       ListMaintenanceWindowsStatusEnum = "DELETING"
	ListMaintenanceWindowsStatusDeleted        ListMaintenanceWindowsStatusEnum = "DELETED"
	ListMaintenanceWindowsStatusFailed         ListMaintenanceWindowsStatusEnum = "FAILED"
	ListMaintenanceWindowsStatusNeedsAttention ListMaintenanceWindowsStatusEnum = "NEEDS_ATTENTION"
)

var mappingListMaintenanceWindowsStatusEnum = map[string]ListMaintenanceWindowsStatusEnum{
	"CREATING":        ListMaintenanceWindowsStatusCreating,
	"UPDATING":        ListMaintenanceWindowsStatusUpdating,
	"INACTIVE":        ListMaintenanceWindowsStatusInactive,
	"ACTIVE":          ListMaintenanceWindowsStatusActive,
	"DELETING":        ListMaintenanceWindowsStatusDeleting,
	"DELETED":         ListMaintenanceWindowsStatusDeleted,
	"FAILED":          ListMaintenanceWindowsStatusFailed,
	"NEEDS_ATTENTION": ListMaintenanceWindowsStatusNeedsAttention,
}

var mappingListMaintenanceWindowsStatusEnumLowerCase = map[string]ListMaintenanceWindowsStatusEnum{
	"creating":        ListMaintenanceWindowsStatusCreating,
	"updating":        ListMaintenanceWindowsStatusUpdating,
	"inactive":        ListMaintenanceWindowsStatusInactive,
	"active":          ListMaintenanceWindowsStatusActive,
	"deleting":        ListMaintenanceWindowsStatusDeleting,
	"deleted":         ListMaintenanceWindowsStatusDeleted,
	"failed":          ListMaintenanceWindowsStatusFailed,
	"needs_attention": ListMaintenanceWindowsStatusNeedsAttention,
}

// GetListMaintenanceWindowsStatusEnumValues Enumerates the set of values for ListMaintenanceWindowsStatusEnum
func GetListMaintenanceWindowsStatusEnumValues() []ListMaintenanceWindowsStatusEnum {
	values := make([]ListMaintenanceWindowsStatusEnum, 0)
	for _, v := range mappingListMaintenanceWindowsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceWindowsStatusEnumStringValues Enumerates the set of values in String for ListMaintenanceWindowsStatusEnum
func GetListMaintenanceWindowsStatusEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"INACTIVE",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListMaintenanceWindowsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceWindowsStatusEnum(val string) (ListMaintenanceWindowsStatusEnum, bool) {
	enum, ok := mappingListMaintenanceWindowsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceWindowsSortByEnum Enum with underlying type: string
type ListMaintenanceWindowsSortByEnum string

// Set of constants representing the allowable values for ListMaintenanceWindowsSortByEnum
const (
	ListMaintenanceWindowsSortByName        ListMaintenanceWindowsSortByEnum = "NAME"
	ListMaintenanceWindowsSortByStartTime   ListMaintenanceWindowsSortByEnum = "START_TIME"
	ListMaintenanceWindowsSortByEndTime     ListMaintenanceWindowsSortByEnum = "END_TIME"
	ListMaintenanceWindowsSortByTimeCreated ListMaintenanceWindowsSortByEnum = "TIME_CREATED"
	ListMaintenanceWindowsSortByTimeUpdated ListMaintenanceWindowsSortByEnum = "TIME_UPDATED"
)

var mappingListMaintenanceWindowsSortByEnum = map[string]ListMaintenanceWindowsSortByEnum{
	"NAME":         ListMaintenanceWindowsSortByName,
	"START_TIME":   ListMaintenanceWindowsSortByStartTime,
	"END_TIME":     ListMaintenanceWindowsSortByEndTime,
	"TIME_CREATED": ListMaintenanceWindowsSortByTimeCreated,
	"TIME_UPDATED": ListMaintenanceWindowsSortByTimeUpdated,
}

var mappingListMaintenanceWindowsSortByEnumLowerCase = map[string]ListMaintenanceWindowsSortByEnum{
	"name":         ListMaintenanceWindowsSortByName,
	"start_time":   ListMaintenanceWindowsSortByStartTime,
	"end_time":     ListMaintenanceWindowsSortByEndTime,
	"time_created": ListMaintenanceWindowsSortByTimeCreated,
	"time_updated": ListMaintenanceWindowsSortByTimeUpdated,
}

// GetListMaintenanceWindowsSortByEnumValues Enumerates the set of values for ListMaintenanceWindowsSortByEnum
func GetListMaintenanceWindowsSortByEnumValues() []ListMaintenanceWindowsSortByEnum {
	values := make([]ListMaintenanceWindowsSortByEnum, 0)
	for _, v := range mappingListMaintenanceWindowsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceWindowsSortByEnumStringValues Enumerates the set of values in String for ListMaintenanceWindowsSortByEnum
func GetListMaintenanceWindowsSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"START_TIME",
		"END_TIME",
		"TIME_CREATED",
		"TIME_UPDATED",
	}
}

// GetMappingListMaintenanceWindowsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceWindowsSortByEnum(val string) (ListMaintenanceWindowsSortByEnum, bool) {
	enum, ok := mappingListMaintenanceWindowsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceWindowsSortOrderEnum Enum with underlying type: string
type ListMaintenanceWindowsSortOrderEnum string

// Set of constants representing the allowable values for ListMaintenanceWindowsSortOrderEnum
const (
	ListMaintenanceWindowsSortOrderAsc  ListMaintenanceWindowsSortOrderEnum = "ASC"
	ListMaintenanceWindowsSortOrderDesc ListMaintenanceWindowsSortOrderEnum = "DESC"
)

var mappingListMaintenanceWindowsSortOrderEnum = map[string]ListMaintenanceWindowsSortOrderEnum{
	"ASC":  ListMaintenanceWindowsSortOrderAsc,
	"DESC": ListMaintenanceWindowsSortOrderDesc,
}

var mappingListMaintenanceWindowsSortOrderEnumLowerCase = map[string]ListMaintenanceWindowsSortOrderEnum{
	"asc":  ListMaintenanceWindowsSortOrderAsc,
	"desc": ListMaintenanceWindowsSortOrderDesc,
}

// GetListMaintenanceWindowsSortOrderEnumValues Enumerates the set of values for ListMaintenanceWindowsSortOrderEnum
func GetListMaintenanceWindowsSortOrderEnumValues() []ListMaintenanceWindowsSortOrderEnum {
	values := make([]ListMaintenanceWindowsSortOrderEnum, 0)
	for _, v := range mappingListMaintenanceWindowsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceWindowsSortOrderEnumStringValues Enumerates the set of values in String for ListMaintenanceWindowsSortOrderEnum
func GetListMaintenanceWindowsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaintenanceWindowsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceWindowsSortOrderEnum(val string) (ListMaintenanceWindowsSortOrderEnum, bool) {
	enum, ok := mappingListMaintenanceWindowsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
