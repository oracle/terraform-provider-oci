// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListJobsRequest wrapper for the ListJobs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListJobs.go.html to see an example of how to use ListJobsRequest.
type ListJobsRequest struct {

	// The ID of the migration in which to list resources.
	MigrationId *string `mandatory:"true" contributesTo:"query" name:"migrationId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The lifecycle state of the Migration Job.
	LifecycleState ListJobsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJobsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJobsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListJobsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJobsResponse wrapper for the ListJobs operation
type ListJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JobCollection instances
	JobCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobsSortByEnum Enum with underlying type: string
type ListJobsSortByEnum string

// Set of constants representing the allowable values for ListJobsSortByEnum
const (
	ListJobsSortByTimecreated ListJobsSortByEnum = "timeCreated"
	ListJobsSortByDisplayname ListJobsSortByEnum = "displayName"
)

var mappingListJobsSortByEnum = map[string]ListJobsSortByEnum{
	"timeCreated": ListJobsSortByTimecreated,
	"displayName": ListJobsSortByDisplayname,
}

// GetListJobsSortByEnumValues Enumerates the set of values for ListJobsSortByEnum
func GetListJobsSortByEnumValues() []ListJobsSortByEnum {
	values := make([]ListJobsSortByEnum, 0)
	for _, v := range mappingListJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobsSortByEnumStringValues Enumerates the set of values in String for ListJobsSortByEnum
func GetListJobsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobsSortByEnum(val string) (ListJobsSortByEnum, bool) {
	mappingListJobsSortByEnumIgnoreCase := make(map[string]ListJobsSortByEnum)
	for k, v := range mappingListJobsSortByEnum {
		mappingListJobsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobsSortOrderEnum Enum with underlying type: string
type ListJobsSortOrderEnum string

// Set of constants representing the allowable values for ListJobsSortOrderEnum
const (
	ListJobsSortOrderAsc  ListJobsSortOrderEnum = "ASC"
	ListJobsSortOrderDesc ListJobsSortOrderEnum = "DESC"
)

var mappingListJobsSortOrderEnum = map[string]ListJobsSortOrderEnum{
	"ASC":  ListJobsSortOrderAsc,
	"DESC": ListJobsSortOrderDesc,
}

// GetListJobsSortOrderEnumValues Enumerates the set of values for ListJobsSortOrderEnum
func GetListJobsSortOrderEnumValues() []ListJobsSortOrderEnum {
	values := make([]ListJobsSortOrderEnum, 0)
	for _, v := range mappingListJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobsSortOrderEnumStringValues Enumerates the set of values in String for ListJobsSortOrderEnum
func GetListJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobsSortOrderEnum(val string) (ListJobsSortOrderEnum, bool) {
	mappingListJobsSortOrderEnumIgnoreCase := make(map[string]ListJobsSortOrderEnum)
	for k, v := range mappingListJobsSortOrderEnum {
		mappingListJobsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobsLifecycleStateEnum Enum with underlying type: string
type ListJobsLifecycleStateEnum string

// Set of constants representing the allowable values for ListJobsLifecycleStateEnum
const (
	ListJobsLifecycleStateAccepted   ListJobsLifecycleStateEnum = "ACCEPTED"
	ListJobsLifecycleStateInProgress ListJobsLifecycleStateEnum = "IN_PROGRESS"
	ListJobsLifecycleStateUnknown    ListJobsLifecycleStateEnum = "UNKNOWN"
	ListJobsLifecycleStateTerminated ListJobsLifecycleStateEnum = "TERMINATED"
	ListJobsLifecycleStateFailed     ListJobsLifecycleStateEnum = "FAILED"
	ListJobsLifecycleStateSucceeded  ListJobsLifecycleStateEnum = "SUCCEEDED"
	ListJobsLifecycleStateWaiting    ListJobsLifecycleStateEnum = "WAITING"
	ListJobsLifecycleStateCanceling  ListJobsLifecycleStateEnum = "CANCELING"
	ListJobsLifecycleStateCanceled   ListJobsLifecycleStateEnum = "CANCELED"
)

var mappingListJobsLifecycleStateEnum = map[string]ListJobsLifecycleStateEnum{
	"ACCEPTED":    ListJobsLifecycleStateAccepted,
	"IN_PROGRESS": ListJobsLifecycleStateInProgress,
	"UNKNOWN":     ListJobsLifecycleStateUnknown,
	"TERMINATED":  ListJobsLifecycleStateTerminated,
	"FAILED":      ListJobsLifecycleStateFailed,
	"SUCCEEDED":   ListJobsLifecycleStateSucceeded,
	"WAITING":     ListJobsLifecycleStateWaiting,
	"CANCELING":   ListJobsLifecycleStateCanceling,
	"CANCELED":    ListJobsLifecycleStateCanceled,
}

// GetListJobsLifecycleStateEnumValues Enumerates the set of values for ListJobsLifecycleStateEnum
func GetListJobsLifecycleStateEnumValues() []ListJobsLifecycleStateEnum {
	values := make([]ListJobsLifecycleStateEnum, 0)
	for _, v := range mappingListJobsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobsLifecycleStateEnumStringValues Enumerates the set of values in String for ListJobsLifecycleStateEnum
func GetListJobsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"UNKNOWN",
		"TERMINATED",
		"FAILED",
		"SUCCEEDED",
		"WAITING",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingListJobsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobsLifecycleStateEnum(val string) (ListJobsLifecycleStateEnum, bool) {
	mappingListJobsLifecycleStateEnumIgnoreCase := make(map[string]ListJobsLifecycleStateEnum)
	for k, v := range mappingListJobsLifecycleStateEnum {
		mappingListJobsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
