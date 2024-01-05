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

// ListMonitoredResourceTasksRequest wrapper for the ListMonitoredResourceTasks operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListMonitoredResourceTasks.go.html to see an example of how to use ListMonitoredResourceTasksRequest.
type ListMonitoredResourceTasksRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment for which
	// stack monitoring resource tasks should be listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that matches with lifecycleState given.
	Status ListMonitoredResourceTasksStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for 'timeUpdated' is descending.
	SortBy ListMonitoredResourceTasksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMonitoredResourceTasksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListMonitoredResourceTasksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMonitoredResourceTasksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMonitoredResourceTasksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMonitoredResourceTasksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMonitoredResourceTasksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMonitoredResourceTasksStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListMonitoredResourceTasksStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitoredResourceTasksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMonitoredResourceTasksSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitoredResourceTasksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMonitoredResourceTasksSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMonitoredResourceTasksResponse wrapper for the ListMonitoredResourceTasks operation
type ListMonitoredResourceTasksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MonitoredResourceTasksCollection instances
	MonitoredResourceTasksCollection `presentIn:"body"`

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

func (response ListMonitoredResourceTasksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMonitoredResourceTasksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMonitoredResourceTasksStatusEnum Enum with underlying type: string
type ListMonitoredResourceTasksStatusEnum string

// Set of constants representing the allowable values for ListMonitoredResourceTasksStatusEnum
const (
	ListMonitoredResourceTasksStatusAccepted       ListMonitoredResourceTasksStatusEnum = "ACCEPTED"
	ListMonitoredResourceTasksStatusInProgress     ListMonitoredResourceTasksStatusEnum = "IN_PROGRESS"
	ListMonitoredResourceTasksStatusWaiting        ListMonitoredResourceTasksStatusEnum = "WAITING"
	ListMonitoredResourceTasksStatusFailed         ListMonitoredResourceTasksStatusEnum = "FAILED"
	ListMonitoredResourceTasksStatusSucceeded      ListMonitoredResourceTasksStatusEnum = "SUCCEEDED"
	ListMonitoredResourceTasksStatusCanceling      ListMonitoredResourceTasksStatusEnum = "CANCELING"
	ListMonitoredResourceTasksStatusCanceled       ListMonitoredResourceTasksStatusEnum = "CANCELED"
	ListMonitoredResourceTasksStatusNeedsAttention ListMonitoredResourceTasksStatusEnum = "NEEDS_ATTENTION"
)

var mappingListMonitoredResourceTasksStatusEnum = map[string]ListMonitoredResourceTasksStatusEnum{
	"ACCEPTED":        ListMonitoredResourceTasksStatusAccepted,
	"IN_PROGRESS":     ListMonitoredResourceTasksStatusInProgress,
	"WAITING":         ListMonitoredResourceTasksStatusWaiting,
	"FAILED":          ListMonitoredResourceTasksStatusFailed,
	"SUCCEEDED":       ListMonitoredResourceTasksStatusSucceeded,
	"CANCELING":       ListMonitoredResourceTasksStatusCanceling,
	"CANCELED":        ListMonitoredResourceTasksStatusCanceled,
	"NEEDS_ATTENTION": ListMonitoredResourceTasksStatusNeedsAttention,
}

var mappingListMonitoredResourceTasksStatusEnumLowerCase = map[string]ListMonitoredResourceTasksStatusEnum{
	"accepted":        ListMonitoredResourceTasksStatusAccepted,
	"in_progress":     ListMonitoredResourceTasksStatusInProgress,
	"waiting":         ListMonitoredResourceTasksStatusWaiting,
	"failed":          ListMonitoredResourceTasksStatusFailed,
	"succeeded":       ListMonitoredResourceTasksStatusSucceeded,
	"canceling":       ListMonitoredResourceTasksStatusCanceling,
	"canceled":        ListMonitoredResourceTasksStatusCanceled,
	"needs_attention": ListMonitoredResourceTasksStatusNeedsAttention,
}

// GetListMonitoredResourceTasksStatusEnumValues Enumerates the set of values for ListMonitoredResourceTasksStatusEnum
func GetListMonitoredResourceTasksStatusEnumValues() []ListMonitoredResourceTasksStatusEnum {
	values := make([]ListMonitoredResourceTasksStatusEnum, 0)
	for _, v := range mappingListMonitoredResourceTasksStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoredResourceTasksStatusEnumStringValues Enumerates the set of values in String for ListMonitoredResourceTasksStatusEnum
func GetListMonitoredResourceTasksStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListMonitoredResourceTasksStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoredResourceTasksStatusEnum(val string) (ListMonitoredResourceTasksStatusEnum, bool) {
	enum, ok := mappingListMonitoredResourceTasksStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitoredResourceTasksSortByEnum Enum with underlying type: string
type ListMonitoredResourceTasksSortByEnum string

// Set of constants representing the allowable values for ListMonitoredResourceTasksSortByEnum
const (
	ListMonitoredResourceTasksSortByTimeupdated ListMonitoredResourceTasksSortByEnum = "timeUpdated"
)

var mappingListMonitoredResourceTasksSortByEnum = map[string]ListMonitoredResourceTasksSortByEnum{
	"timeUpdated": ListMonitoredResourceTasksSortByTimeupdated,
}

var mappingListMonitoredResourceTasksSortByEnumLowerCase = map[string]ListMonitoredResourceTasksSortByEnum{
	"timeupdated": ListMonitoredResourceTasksSortByTimeupdated,
}

// GetListMonitoredResourceTasksSortByEnumValues Enumerates the set of values for ListMonitoredResourceTasksSortByEnum
func GetListMonitoredResourceTasksSortByEnumValues() []ListMonitoredResourceTasksSortByEnum {
	values := make([]ListMonitoredResourceTasksSortByEnum, 0)
	for _, v := range mappingListMonitoredResourceTasksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoredResourceTasksSortByEnumStringValues Enumerates the set of values in String for ListMonitoredResourceTasksSortByEnum
func GetListMonitoredResourceTasksSortByEnumStringValues() []string {
	return []string{
		"timeUpdated",
	}
}

// GetMappingListMonitoredResourceTasksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoredResourceTasksSortByEnum(val string) (ListMonitoredResourceTasksSortByEnum, bool) {
	enum, ok := mappingListMonitoredResourceTasksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitoredResourceTasksSortOrderEnum Enum with underlying type: string
type ListMonitoredResourceTasksSortOrderEnum string

// Set of constants representing the allowable values for ListMonitoredResourceTasksSortOrderEnum
const (
	ListMonitoredResourceTasksSortOrderAsc  ListMonitoredResourceTasksSortOrderEnum = "ASC"
	ListMonitoredResourceTasksSortOrderDesc ListMonitoredResourceTasksSortOrderEnum = "DESC"
)

var mappingListMonitoredResourceTasksSortOrderEnum = map[string]ListMonitoredResourceTasksSortOrderEnum{
	"ASC":  ListMonitoredResourceTasksSortOrderAsc,
	"DESC": ListMonitoredResourceTasksSortOrderDesc,
}

var mappingListMonitoredResourceTasksSortOrderEnumLowerCase = map[string]ListMonitoredResourceTasksSortOrderEnum{
	"asc":  ListMonitoredResourceTasksSortOrderAsc,
	"desc": ListMonitoredResourceTasksSortOrderDesc,
}

// GetListMonitoredResourceTasksSortOrderEnumValues Enumerates the set of values for ListMonitoredResourceTasksSortOrderEnum
func GetListMonitoredResourceTasksSortOrderEnumValues() []ListMonitoredResourceTasksSortOrderEnum {
	values := make([]ListMonitoredResourceTasksSortOrderEnum, 0)
	for _, v := range mappingListMonitoredResourceTasksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoredResourceTasksSortOrderEnumStringValues Enumerates the set of values in String for ListMonitoredResourceTasksSortOrderEnum
func GetListMonitoredResourceTasksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMonitoredResourceTasksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoredResourceTasksSortOrderEnum(val string) (ListMonitoredResourceTasksSortOrderEnum, bool) {
	enum, ok := mappingListMonitoredResourceTasksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
