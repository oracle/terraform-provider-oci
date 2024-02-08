// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListReplicationSchedulesRequest wrapper for the ListReplicationSchedules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudmigrations/ListReplicationSchedules.go.html to see an example of how to use ListReplicationSchedulesRequest.
type ListReplicationSchedulesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The current state of the replication schedule.
	LifecycleState ReplicationScheduleLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire given display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique replication schedule identifier in query
	ReplicationScheduleId *string `mandatory:"false" contributesTo:"query" name:"replicationScheduleId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of the previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListReplicationSchedulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. The default order for 'timeCreated' is descending. The default order for 'displayName' is ascending.
	SortBy ListReplicationSchedulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReplicationSchedulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReplicationSchedulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReplicationSchedulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReplicationSchedulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListReplicationSchedulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReplicationScheduleLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetReplicationScheduleLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReplicationSchedulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListReplicationSchedulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReplicationSchedulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListReplicationSchedulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListReplicationSchedulesResponse wrapper for the ListReplicationSchedules operation
type ListReplicationSchedulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ReplicationScheduleCollection instances
	ReplicationScheduleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListReplicationSchedulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReplicationSchedulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReplicationSchedulesSortOrderEnum Enum with underlying type: string
type ListReplicationSchedulesSortOrderEnum string

// Set of constants representing the allowable values for ListReplicationSchedulesSortOrderEnum
const (
	ListReplicationSchedulesSortOrderAsc  ListReplicationSchedulesSortOrderEnum = "ASC"
	ListReplicationSchedulesSortOrderDesc ListReplicationSchedulesSortOrderEnum = "DESC"
)

var mappingListReplicationSchedulesSortOrderEnum = map[string]ListReplicationSchedulesSortOrderEnum{
	"ASC":  ListReplicationSchedulesSortOrderAsc,
	"DESC": ListReplicationSchedulesSortOrderDesc,
}

var mappingListReplicationSchedulesSortOrderEnumLowerCase = map[string]ListReplicationSchedulesSortOrderEnum{
	"asc":  ListReplicationSchedulesSortOrderAsc,
	"desc": ListReplicationSchedulesSortOrderDesc,
}

// GetListReplicationSchedulesSortOrderEnumValues Enumerates the set of values for ListReplicationSchedulesSortOrderEnum
func GetListReplicationSchedulesSortOrderEnumValues() []ListReplicationSchedulesSortOrderEnum {
	values := make([]ListReplicationSchedulesSortOrderEnum, 0)
	for _, v := range mappingListReplicationSchedulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListReplicationSchedulesSortOrderEnumStringValues Enumerates the set of values in String for ListReplicationSchedulesSortOrderEnum
func GetListReplicationSchedulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListReplicationSchedulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReplicationSchedulesSortOrderEnum(val string) (ListReplicationSchedulesSortOrderEnum, bool) {
	enum, ok := mappingListReplicationSchedulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReplicationSchedulesSortByEnum Enum with underlying type: string
type ListReplicationSchedulesSortByEnum string

// Set of constants representing the allowable values for ListReplicationSchedulesSortByEnum
const (
	ListReplicationSchedulesSortByTimecreated ListReplicationSchedulesSortByEnum = "timeCreated"
	ListReplicationSchedulesSortByDisplayname ListReplicationSchedulesSortByEnum = "displayName"
)

var mappingListReplicationSchedulesSortByEnum = map[string]ListReplicationSchedulesSortByEnum{
	"timeCreated": ListReplicationSchedulesSortByTimecreated,
	"displayName": ListReplicationSchedulesSortByDisplayname,
}

var mappingListReplicationSchedulesSortByEnumLowerCase = map[string]ListReplicationSchedulesSortByEnum{
	"timecreated": ListReplicationSchedulesSortByTimecreated,
	"displayname": ListReplicationSchedulesSortByDisplayname,
}

// GetListReplicationSchedulesSortByEnumValues Enumerates the set of values for ListReplicationSchedulesSortByEnum
func GetListReplicationSchedulesSortByEnumValues() []ListReplicationSchedulesSortByEnum {
	values := make([]ListReplicationSchedulesSortByEnum, 0)
	for _, v := range mappingListReplicationSchedulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListReplicationSchedulesSortByEnumStringValues Enumerates the set of values in String for ListReplicationSchedulesSortByEnum
func GetListReplicationSchedulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListReplicationSchedulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReplicationSchedulesSortByEnum(val string) (ListReplicationSchedulesSortByEnum, bool) {
	enum, ok := mappingListReplicationSchedulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
