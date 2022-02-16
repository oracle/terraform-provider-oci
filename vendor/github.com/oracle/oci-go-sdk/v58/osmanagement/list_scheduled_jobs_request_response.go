// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListScheduledJobsRequest wrapper for the ListScheduledJobs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListScheduledJobs.go.html to see an example of how to use ListScheduledJobsRequest.
type ListScheduledJobsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the managed instance for which to list resources.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The ID of the managed instace group for which to list resources.
	ManagedInstanceGroupId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceGroupId"`

	// The operation type for which to list resources
	OperationType ListScheduledJobsOperationTypeEnum `mandatory:"false" contributesTo:"query" name:"operationType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListScheduledJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListScheduledJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The current lifecycle state for the object.
	LifecycleState ListScheduledJobsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OS family for which to list resources.
	OsFamily ListScheduledJobsOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// If true, will only filter out restricted Autonomous Linux Scheduled Job
	IsRestricted *bool `mandatory:"false" contributesTo:"query" name:"isRestricted"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListScheduledJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListScheduledJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListScheduledJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListScheduledJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListScheduledJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListScheduledJobsOperationTypeEnum(string(request.OperationType)); !ok && request.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", request.OperationType, strings.Join(GetListScheduledJobsOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListScheduledJobsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListScheduledJobsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledJobsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListScheduledJobsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledJobsOsFamilyEnum(string(request.OsFamily)); !ok && request.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", request.OsFamily, strings.Join(GetListScheduledJobsOsFamilyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListScheduledJobsResponse wrapper for the ListScheduledJobs operation
type ListScheduledJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ScheduledJobSummary instances
	Items []ScheduledJobSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListScheduledJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListScheduledJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListScheduledJobsOperationTypeEnum Enum with underlying type: string
type ListScheduledJobsOperationTypeEnum string

// Set of constants representing the allowable values for ListScheduledJobsOperationTypeEnum
const (
	ListScheduledJobsOperationTypeInstall   ListScheduledJobsOperationTypeEnum = "INSTALL"
	ListScheduledJobsOperationTypeUpdate    ListScheduledJobsOperationTypeEnum = "UPDATE"
	ListScheduledJobsOperationTypeRemove    ListScheduledJobsOperationTypeEnum = "REMOVE"
	ListScheduledJobsOperationTypeUpdateall ListScheduledJobsOperationTypeEnum = "UPDATEALL"
)

var mappingListScheduledJobsOperationTypeEnum = map[string]ListScheduledJobsOperationTypeEnum{
	"INSTALL":   ListScheduledJobsOperationTypeInstall,
	"UPDATE":    ListScheduledJobsOperationTypeUpdate,
	"REMOVE":    ListScheduledJobsOperationTypeRemove,
	"UPDATEALL": ListScheduledJobsOperationTypeUpdateall,
}

// GetListScheduledJobsOperationTypeEnumValues Enumerates the set of values for ListScheduledJobsOperationTypeEnum
func GetListScheduledJobsOperationTypeEnumValues() []ListScheduledJobsOperationTypeEnum {
	values := make([]ListScheduledJobsOperationTypeEnum, 0)
	for _, v := range mappingListScheduledJobsOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledJobsOperationTypeEnumStringValues Enumerates the set of values in String for ListScheduledJobsOperationTypeEnum
func GetListScheduledJobsOperationTypeEnumStringValues() []string {
	return []string{
		"INSTALL",
		"UPDATE",
		"REMOVE",
		"UPDATEALL",
	}
}

// GetMappingListScheduledJobsOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledJobsOperationTypeEnum(val string) (ListScheduledJobsOperationTypeEnum, bool) {
	mappingListScheduledJobsOperationTypeEnumIgnoreCase := make(map[string]ListScheduledJobsOperationTypeEnum)
	for k, v := range mappingListScheduledJobsOperationTypeEnum {
		mappingListScheduledJobsOperationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListScheduledJobsOperationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListScheduledJobsSortOrderEnum Enum with underlying type: string
type ListScheduledJobsSortOrderEnum string

// Set of constants representing the allowable values for ListScheduledJobsSortOrderEnum
const (
	ListScheduledJobsSortOrderAsc  ListScheduledJobsSortOrderEnum = "ASC"
	ListScheduledJobsSortOrderDesc ListScheduledJobsSortOrderEnum = "DESC"
)

var mappingListScheduledJobsSortOrderEnum = map[string]ListScheduledJobsSortOrderEnum{
	"ASC":  ListScheduledJobsSortOrderAsc,
	"DESC": ListScheduledJobsSortOrderDesc,
}

// GetListScheduledJobsSortOrderEnumValues Enumerates the set of values for ListScheduledJobsSortOrderEnum
func GetListScheduledJobsSortOrderEnumValues() []ListScheduledJobsSortOrderEnum {
	values := make([]ListScheduledJobsSortOrderEnum, 0)
	for _, v := range mappingListScheduledJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledJobsSortOrderEnumStringValues Enumerates the set of values in String for ListScheduledJobsSortOrderEnum
func GetListScheduledJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListScheduledJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledJobsSortOrderEnum(val string) (ListScheduledJobsSortOrderEnum, bool) {
	mappingListScheduledJobsSortOrderEnumIgnoreCase := make(map[string]ListScheduledJobsSortOrderEnum)
	for k, v := range mappingListScheduledJobsSortOrderEnum {
		mappingListScheduledJobsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListScheduledJobsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListScheduledJobsSortByEnum Enum with underlying type: string
type ListScheduledJobsSortByEnum string

// Set of constants representing the allowable values for ListScheduledJobsSortByEnum
const (
	ListScheduledJobsSortByTimecreated ListScheduledJobsSortByEnum = "TIMECREATED"
	ListScheduledJobsSortByDisplayname ListScheduledJobsSortByEnum = "DISPLAYNAME"
)

var mappingListScheduledJobsSortByEnum = map[string]ListScheduledJobsSortByEnum{
	"TIMECREATED": ListScheduledJobsSortByTimecreated,
	"DISPLAYNAME": ListScheduledJobsSortByDisplayname,
}

// GetListScheduledJobsSortByEnumValues Enumerates the set of values for ListScheduledJobsSortByEnum
func GetListScheduledJobsSortByEnumValues() []ListScheduledJobsSortByEnum {
	values := make([]ListScheduledJobsSortByEnum, 0)
	for _, v := range mappingListScheduledJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledJobsSortByEnumStringValues Enumerates the set of values in String for ListScheduledJobsSortByEnum
func GetListScheduledJobsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListScheduledJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledJobsSortByEnum(val string) (ListScheduledJobsSortByEnum, bool) {
	mappingListScheduledJobsSortByEnumIgnoreCase := make(map[string]ListScheduledJobsSortByEnum)
	for k, v := range mappingListScheduledJobsSortByEnum {
		mappingListScheduledJobsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListScheduledJobsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListScheduledJobsLifecycleStateEnum Enum with underlying type: string
type ListScheduledJobsLifecycleStateEnum string

// Set of constants representing the allowable values for ListScheduledJobsLifecycleStateEnum
const (
	ListScheduledJobsLifecycleStateCreating ListScheduledJobsLifecycleStateEnum = "CREATING"
	ListScheduledJobsLifecycleStateUpdating ListScheduledJobsLifecycleStateEnum = "UPDATING"
	ListScheduledJobsLifecycleStateActive   ListScheduledJobsLifecycleStateEnum = "ACTIVE"
	ListScheduledJobsLifecycleStateDeleting ListScheduledJobsLifecycleStateEnum = "DELETING"
	ListScheduledJobsLifecycleStateDeleted  ListScheduledJobsLifecycleStateEnum = "DELETED"
	ListScheduledJobsLifecycleStateFailed   ListScheduledJobsLifecycleStateEnum = "FAILED"
)

var mappingListScheduledJobsLifecycleStateEnum = map[string]ListScheduledJobsLifecycleStateEnum{
	"CREATING": ListScheduledJobsLifecycleStateCreating,
	"UPDATING": ListScheduledJobsLifecycleStateUpdating,
	"ACTIVE":   ListScheduledJobsLifecycleStateActive,
	"DELETING": ListScheduledJobsLifecycleStateDeleting,
	"DELETED":  ListScheduledJobsLifecycleStateDeleted,
	"FAILED":   ListScheduledJobsLifecycleStateFailed,
}

// GetListScheduledJobsLifecycleStateEnumValues Enumerates the set of values for ListScheduledJobsLifecycleStateEnum
func GetListScheduledJobsLifecycleStateEnumValues() []ListScheduledJobsLifecycleStateEnum {
	values := make([]ListScheduledJobsLifecycleStateEnum, 0)
	for _, v := range mappingListScheduledJobsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledJobsLifecycleStateEnumStringValues Enumerates the set of values in String for ListScheduledJobsLifecycleStateEnum
func GetListScheduledJobsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListScheduledJobsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledJobsLifecycleStateEnum(val string) (ListScheduledJobsLifecycleStateEnum, bool) {
	mappingListScheduledJobsLifecycleStateEnumIgnoreCase := make(map[string]ListScheduledJobsLifecycleStateEnum)
	for k, v := range mappingListScheduledJobsLifecycleStateEnum {
		mappingListScheduledJobsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListScheduledJobsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListScheduledJobsOsFamilyEnum Enum with underlying type: string
type ListScheduledJobsOsFamilyEnum string

// Set of constants representing the allowable values for ListScheduledJobsOsFamilyEnum
const (
	ListScheduledJobsOsFamilyLinux   ListScheduledJobsOsFamilyEnum = "LINUX"
	ListScheduledJobsOsFamilyWindows ListScheduledJobsOsFamilyEnum = "WINDOWS"
	ListScheduledJobsOsFamilyAll     ListScheduledJobsOsFamilyEnum = "ALL"
)

var mappingListScheduledJobsOsFamilyEnum = map[string]ListScheduledJobsOsFamilyEnum{
	"LINUX":   ListScheduledJobsOsFamilyLinux,
	"WINDOWS": ListScheduledJobsOsFamilyWindows,
	"ALL":     ListScheduledJobsOsFamilyAll,
}

// GetListScheduledJobsOsFamilyEnumValues Enumerates the set of values for ListScheduledJobsOsFamilyEnum
func GetListScheduledJobsOsFamilyEnumValues() []ListScheduledJobsOsFamilyEnum {
	values := make([]ListScheduledJobsOsFamilyEnum, 0)
	for _, v := range mappingListScheduledJobsOsFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledJobsOsFamilyEnumStringValues Enumerates the set of values in String for ListScheduledJobsOsFamilyEnum
func GetListScheduledJobsOsFamilyEnumStringValues() []string {
	return []string{
		"LINUX",
		"WINDOWS",
		"ALL",
	}
}

// GetMappingListScheduledJobsOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledJobsOsFamilyEnum(val string) (ListScheduledJobsOsFamilyEnum, bool) {
	mappingListScheduledJobsOsFamilyEnumIgnoreCase := make(map[string]ListScheduledJobsOsFamilyEnum)
	for k, v := range mappingListScheduledJobsOsFamilyEnum {
		mappingListScheduledJobsOsFamilyEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListScheduledJobsOsFamilyEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
