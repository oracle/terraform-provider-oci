// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListLogsRequest wrapper for the ListLogs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/logging/ListLogs.go.html to see an example of how to use ListLogsRequest.
type ListLogsRequest struct {

	// OCID of a log group to work with.
	LogGroupId *string `mandatory:"true" contributesTo:"path" name:"logGroupId"`

	// The logType that the log object is for, whether custom or service.
	LogType ListLogsLogTypeEnum `mandatory:"false" contributesTo:"query" name:"logType" omitEmpty:"true"`

	// Service that created the log object.
	SourceService *string `mandatory:"false" contributesTo:"query" name:"sourceService"`

	// Log object resource.
	SourceResource *string `mandatory:"false" contributesTo:"query" name:"sourceResource"`

	// Resource name
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Lifecycle state of the log object
	LifecycleState ListLogsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by (one column only). Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLogsLogTypeEnum(string(request.LogType)); !ok && request.LogType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogType: %s. Supported values are: %s.", request.LogType, strings.Join(GetListLogsLogTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListLogsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLogsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLogsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLogsResponse wrapper for the ListLogs operation
type ListLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []LogSummary instances
	Items []LogSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages
	// of results exist. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogsLogTypeEnum Enum with underlying type: string
type ListLogsLogTypeEnum string

// Set of constants representing the allowable values for ListLogsLogTypeEnum
const (
	ListLogsLogTypeCustom  ListLogsLogTypeEnum = "CUSTOM"
	ListLogsLogTypeService ListLogsLogTypeEnum = "SERVICE"
)

var mappingListLogsLogTypeEnum = map[string]ListLogsLogTypeEnum{
	"CUSTOM":  ListLogsLogTypeCustom,
	"SERVICE": ListLogsLogTypeService,
}

// GetListLogsLogTypeEnumValues Enumerates the set of values for ListLogsLogTypeEnum
func GetListLogsLogTypeEnumValues() []ListLogsLogTypeEnum {
	values := make([]ListLogsLogTypeEnum, 0)
	for _, v := range mappingListLogsLogTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogsLogTypeEnumStringValues Enumerates the set of values in String for ListLogsLogTypeEnum
func GetListLogsLogTypeEnumStringValues() []string {
	return []string{
		"CUSTOM",
		"SERVICE",
	}
}

// GetMappingListLogsLogTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogsLogTypeEnum(val string) (ListLogsLogTypeEnum, bool) {
	mappingListLogsLogTypeEnumIgnoreCase := make(map[string]ListLogsLogTypeEnum)
	for k, v := range mappingListLogsLogTypeEnum {
		mappingListLogsLogTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLogsLogTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogsLifecycleStateEnum Enum with underlying type: string
type ListLogsLifecycleStateEnum string

// Set of constants representing the allowable values for ListLogsLifecycleStateEnum
const (
	ListLogsLifecycleStateCreating ListLogsLifecycleStateEnum = "CREATING"
	ListLogsLifecycleStateActive   ListLogsLifecycleStateEnum = "ACTIVE"
	ListLogsLifecycleStateUpdating ListLogsLifecycleStateEnum = "UPDATING"
	ListLogsLifecycleStateInactive ListLogsLifecycleStateEnum = "INACTIVE"
	ListLogsLifecycleStateDeleting ListLogsLifecycleStateEnum = "DELETING"
	ListLogsLifecycleStateFailed   ListLogsLifecycleStateEnum = "FAILED"
)

var mappingListLogsLifecycleStateEnum = map[string]ListLogsLifecycleStateEnum{
	"CREATING": ListLogsLifecycleStateCreating,
	"ACTIVE":   ListLogsLifecycleStateActive,
	"UPDATING": ListLogsLifecycleStateUpdating,
	"INACTIVE": ListLogsLifecycleStateInactive,
	"DELETING": ListLogsLifecycleStateDeleting,
	"FAILED":   ListLogsLifecycleStateFailed,
}

// GetListLogsLifecycleStateEnumValues Enumerates the set of values for ListLogsLifecycleStateEnum
func GetListLogsLifecycleStateEnumValues() []ListLogsLifecycleStateEnum {
	values := make([]ListLogsLifecycleStateEnum, 0)
	for _, v := range mappingListLogsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogsLifecycleStateEnumStringValues Enumerates the set of values in String for ListLogsLifecycleStateEnum
func GetListLogsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
		"DELETING",
		"FAILED",
	}
}

// GetMappingListLogsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogsLifecycleStateEnum(val string) (ListLogsLifecycleStateEnum, bool) {
	mappingListLogsLifecycleStateEnumIgnoreCase := make(map[string]ListLogsLifecycleStateEnum)
	for k, v := range mappingListLogsLifecycleStateEnum {
		mappingListLogsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLogsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogsSortByEnum Enum with underlying type: string
type ListLogsSortByEnum string

// Set of constants representing the allowable values for ListLogsSortByEnum
const (
	ListLogsSortByTimecreated ListLogsSortByEnum = "timeCreated"
	ListLogsSortByDisplayname ListLogsSortByEnum = "displayName"
)

var mappingListLogsSortByEnum = map[string]ListLogsSortByEnum{
	"timeCreated": ListLogsSortByTimecreated,
	"displayName": ListLogsSortByDisplayname,
}

// GetListLogsSortByEnumValues Enumerates the set of values for ListLogsSortByEnum
func GetListLogsSortByEnumValues() []ListLogsSortByEnum {
	values := make([]ListLogsSortByEnum, 0)
	for _, v := range mappingListLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogsSortByEnumStringValues Enumerates the set of values in String for ListLogsSortByEnum
func GetListLogsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogsSortByEnum(val string) (ListLogsSortByEnum, bool) {
	mappingListLogsSortByEnumIgnoreCase := make(map[string]ListLogsSortByEnum)
	for k, v := range mappingListLogsSortByEnum {
		mappingListLogsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLogsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogsSortOrderEnum Enum with underlying type: string
type ListLogsSortOrderEnum string

// Set of constants representing the allowable values for ListLogsSortOrderEnum
const (
	ListLogsSortOrderAsc  ListLogsSortOrderEnum = "ASC"
	ListLogsSortOrderDesc ListLogsSortOrderEnum = "DESC"
)

var mappingListLogsSortOrderEnum = map[string]ListLogsSortOrderEnum{
	"ASC":  ListLogsSortOrderAsc,
	"DESC": ListLogsSortOrderDesc,
}

// GetListLogsSortOrderEnumValues Enumerates the set of values for ListLogsSortOrderEnum
func GetListLogsSortOrderEnumValues() []ListLogsSortOrderEnum {
	values := make([]ListLogsSortOrderEnum, 0)
	for _, v := range mappingListLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogsSortOrderEnumStringValues Enumerates the set of values in String for ListLogsSortOrderEnum
func GetListLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogsSortOrderEnum(val string) (ListLogsSortOrderEnum, bool) {
	mappingListLogsSortOrderEnumIgnoreCase := make(map[string]ListLogsSortOrderEnum)
	for k, v := range mappingListLogsSortOrderEnum {
		mappingListLogsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLogsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
