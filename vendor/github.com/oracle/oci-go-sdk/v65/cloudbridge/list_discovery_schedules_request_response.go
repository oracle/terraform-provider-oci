// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDiscoverySchedulesRequest wrapper for the ListDiscoverySchedules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListDiscoverySchedules.go.html to see an example of how to use ListDiscoverySchedulesRequest.
type ListDiscoverySchedulesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the discovery schedule.
	DiscoveryScheduleId *string `mandatory:"false" contributesTo:"query" name:"discoveryScheduleId"`

	// The current state of the discovery schedule.
	LifecycleState ListDiscoverySchedulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. By default, the timeCreated is in descending order and displayName is in ascending order.
	SortBy ListDiscoverySchedulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListDiscoverySchedulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDiscoverySchedulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDiscoverySchedulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDiscoverySchedulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDiscoverySchedulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDiscoverySchedulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDiscoverySchedulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDiscoverySchedulesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiscoverySchedulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDiscoverySchedulesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiscoverySchedulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDiscoverySchedulesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDiscoverySchedulesResponse wrapper for the ListDiscoverySchedules operation
type ListDiscoverySchedulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DiscoveryScheduleCollection instances
	DiscoveryScheduleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDiscoverySchedulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDiscoverySchedulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDiscoverySchedulesLifecycleStateEnum Enum with underlying type: string
type ListDiscoverySchedulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDiscoverySchedulesLifecycleStateEnum
const (
	ListDiscoverySchedulesLifecycleStateActive  ListDiscoverySchedulesLifecycleStateEnum = "ACTIVE"
	ListDiscoverySchedulesLifecycleStateDeleted ListDiscoverySchedulesLifecycleStateEnum = "DELETED"
)

var mappingListDiscoverySchedulesLifecycleStateEnum = map[string]ListDiscoverySchedulesLifecycleStateEnum{
	"ACTIVE":  ListDiscoverySchedulesLifecycleStateActive,
	"DELETED": ListDiscoverySchedulesLifecycleStateDeleted,
}

var mappingListDiscoverySchedulesLifecycleStateEnumLowerCase = map[string]ListDiscoverySchedulesLifecycleStateEnum{
	"active":  ListDiscoverySchedulesLifecycleStateActive,
	"deleted": ListDiscoverySchedulesLifecycleStateDeleted,
}

// GetListDiscoverySchedulesLifecycleStateEnumValues Enumerates the set of values for ListDiscoverySchedulesLifecycleStateEnum
func GetListDiscoverySchedulesLifecycleStateEnumValues() []ListDiscoverySchedulesLifecycleStateEnum {
	values := make([]ListDiscoverySchedulesLifecycleStateEnum, 0)
	for _, v := range mappingListDiscoverySchedulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoverySchedulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListDiscoverySchedulesLifecycleStateEnum
func GetListDiscoverySchedulesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListDiscoverySchedulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoverySchedulesLifecycleStateEnum(val string) (ListDiscoverySchedulesLifecycleStateEnum, bool) {
	enum, ok := mappingListDiscoverySchedulesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiscoverySchedulesSortByEnum Enum with underlying type: string
type ListDiscoverySchedulesSortByEnum string

// Set of constants representing the allowable values for ListDiscoverySchedulesSortByEnum
const (
	ListDiscoverySchedulesSortByTimecreated ListDiscoverySchedulesSortByEnum = "timeCreated"
	ListDiscoverySchedulesSortByDisplayname ListDiscoverySchedulesSortByEnum = "displayName"
)

var mappingListDiscoverySchedulesSortByEnum = map[string]ListDiscoverySchedulesSortByEnum{
	"timeCreated": ListDiscoverySchedulesSortByTimecreated,
	"displayName": ListDiscoverySchedulesSortByDisplayname,
}

var mappingListDiscoverySchedulesSortByEnumLowerCase = map[string]ListDiscoverySchedulesSortByEnum{
	"timecreated": ListDiscoverySchedulesSortByTimecreated,
	"displayname": ListDiscoverySchedulesSortByDisplayname,
}

// GetListDiscoverySchedulesSortByEnumValues Enumerates the set of values for ListDiscoverySchedulesSortByEnum
func GetListDiscoverySchedulesSortByEnumValues() []ListDiscoverySchedulesSortByEnum {
	values := make([]ListDiscoverySchedulesSortByEnum, 0)
	for _, v := range mappingListDiscoverySchedulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoverySchedulesSortByEnumStringValues Enumerates the set of values in String for ListDiscoverySchedulesSortByEnum
func GetListDiscoverySchedulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDiscoverySchedulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoverySchedulesSortByEnum(val string) (ListDiscoverySchedulesSortByEnum, bool) {
	enum, ok := mappingListDiscoverySchedulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiscoverySchedulesSortOrderEnum Enum with underlying type: string
type ListDiscoverySchedulesSortOrderEnum string

// Set of constants representing the allowable values for ListDiscoverySchedulesSortOrderEnum
const (
	ListDiscoverySchedulesSortOrderAsc  ListDiscoverySchedulesSortOrderEnum = "ASC"
	ListDiscoverySchedulesSortOrderDesc ListDiscoverySchedulesSortOrderEnum = "DESC"
)

var mappingListDiscoverySchedulesSortOrderEnum = map[string]ListDiscoverySchedulesSortOrderEnum{
	"ASC":  ListDiscoverySchedulesSortOrderAsc,
	"DESC": ListDiscoverySchedulesSortOrderDesc,
}

var mappingListDiscoverySchedulesSortOrderEnumLowerCase = map[string]ListDiscoverySchedulesSortOrderEnum{
	"asc":  ListDiscoverySchedulesSortOrderAsc,
	"desc": ListDiscoverySchedulesSortOrderDesc,
}

// GetListDiscoverySchedulesSortOrderEnumValues Enumerates the set of values for ListDiscoverySchedulesSortOrderEnum
func GetListDiscoverySchedulesSortOrderEnumValues() []ListDiscoverySchedulesSortOrderEnum {
	values := make([]ListDiscoverySchedulesSortOrderEnum, 0)
	for _, v := range mappingListDiscoverySchedulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoverySchedulesSortOrderEnumStringValues Enumerates the set of values in String for ListDiscoverySchedulesSortOrderEnum
func GetListDiscoverySchedulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDiscoverySchedulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoverySchedulesSortOrderEnum(val string) (ListDiscoverySchedulesSortOrderEnum, bool) {
	enum, ok := mappingListDiscoverySchedulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
