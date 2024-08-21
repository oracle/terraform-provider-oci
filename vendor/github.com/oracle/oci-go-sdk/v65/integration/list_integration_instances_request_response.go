// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package integration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListIntegrationInstancesRequest wrapper for the ListIntegrationInstances operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/integration/ListIntegrationInstances.go.html to see an example of how to use ListIntegrationInstancesRequest.
type ListIntegrationInstancesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Life cycle state to query on.
	LifecycleState ListIntegrationInstancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListIntegrationInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order
	// for TIMECREATED is descending. Default order for DISPLAYNAME is
	// ascending. If no value is specified TIMECREATED is default.
	SortBy ListIntegrationInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIntegrationInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIntegrationInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIntegrationInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIntegrationInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIntegrationInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListIntegrationInstancesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListIntegrationInstancesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIntegrationInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIntegrationInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIntegrationInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIntegrationInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIntegrationInstancesResponse wrapper for the ListIntegrationInstances operation
type ListIntegrationInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []IntegrationInstanceSummary instances
	Items []IntegrationInstanceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, additional pages of results have been previously returned
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`
}

func (response ListIntegrationInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIntegrationInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIntegrationInstancesLifecycleStateEnum Enum with underlying type: string
type ListIntegrationInstancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListIntegrationInstancesLifecycleStateEnum
const (
	ListIntegrationInstancesLifecycleStateCreating ListIntegrationInstancesLifecycleStateEnum = "CREATING"
	ListIntegrationInstancesLifecycleStateUpdating ListIntegrationInstancesLifecycleStateEnum = "UPDATING"
	ListIntegrationInstancesLifecycleStateActive   ListIntegrationInstancesLifecycleStateEnum = "ACTIVE"
	ListIntegrationInstancesLifecycleStateInactive ListIntegrationInstancesLifecycleStateEnum = "INACTIVE"
	ListIntegrationInstancesLifecycleStateDeleting ListIntegrationInstancesLifecycleStateEnum = "DELETING"
	ListIntegrationInstancesLifecycleStateDeleted  ListIntegrationInstancesLifecycleStateEnum = "DELETED"
	ListIntegrationInstancesLifecycleStateFailed   ListIntegrationInstancesLifecycleStateEnum = "FAILED"
	ListIntegrationInstancesLifecycleStateStandby  ListIntegrationInstancesLifecycleStateEnum = "STANDBY"
)

var mappingListIntegrationInstancesLifecycleStateEnum = map[string]ListIntegrationInstancesLifecycleStateEnum{
	"CREATING": ListIntegrationInstancesLifecycleStateCreating,
	"UPDATING": ListIntegrationInstancesLifecycleStateUpdating,
	"ACTIVE":   ListIntegrationInstancesLifecycleStateActive,
	"INACTIVE": ListIntegrationInstancesLifecycleStateInactive,
	"DELETING": ListIntegrationInstancesLifecycleStateDeleting,
	"DELETED":  ListIntegrationInstancesLifecycleStateDeleted,
	"FAILED":   ListIntegrationInstancesLifecycleStateFailed,
	"STANDBY":  ListIntegrationInstancesLifecycleStateStandby,
}

var mappingListIntegrationInstancesLifecycleStateEnumLowerCase = map[string]ListIntegrationInstancesLifecycleStateEnum{
	"creating": ListIntegrationInstancesLifecycleStateCreating,
	"updating": ListIntegrationInstancesLifecycleStateUpdating,
	"active":   ListIntegrationInstancesLifecycleStateActive,
	"inactive": ListIntegrationInstancesLifecycleStateInactive,
	"deleting": ListIntegrationInstancesLifecycleStateDeleting,
	"deleted":  ListIntegrationInstancesLifecycleStateDeleted,
	"failed":   ListIntegrationInstancesLifecycleStateFailed,
	"standby":  ListIntegrationInstancesLifecycleStateStandby,
}

// GetListIntegrationInstancesLifecycleStateEnumValues Enumerates the set of values for ListIntegrationInstancesLifecycleStateEnum
func GetListIntegrationInstancesLifecycleStateEnumValues() []ListIntegrationInstancesLifecycleStateEnum {
	values := make([]ListIntegrationInstancesLifecycleStateEnum, 0)
	for _, v := range mappingListIntegrationInstancesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListIntegrationInstancesLifecycleStateEnumStringValues Enumerates the set of values in String for ListIntegrationInstancesLifecycleStateEnum
func GetListIntegrationInstancesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"STANDBY",
	}
}

// GetMappingListIntegrationInstancesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIntegrationInstancesLifecycleStateEnum(val string) (ListIntegrationInstancesLifecycleStateEnum, bool) {
	enum, ok := mappingListIntegrationInstancesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIntegrationInstancesSortOrderEnum Enum with underlying type: string
type ListIntegrationInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListIntegrationInstancesSortOrderEnum
const (
	ListIntegrationInstancesSortOrderAsc  ListIntegrationInstancesSortOrderEnum = "ASC"
	ListIntegrationInstancesSortOrderDesc ListIntegrationInstancesSortOrderEnum = "DESC"
)

var mappingListIntegrationInstancesSortOrderEnum = map[string]ListIntegrationInstancesSortOrderEnum{
	"ASC":  ListIntegrationInstancesSortOrderAsc,
	"DESC": ListIntegrationInstancesSortOrderDesc,
}

var mappingListIntegrationInstancesSortOrderEnumLowerCase = map[string]ListIntegrationInstancesSortOrderEnum{
	"asc":  ListIntegrationInstancesSortOrderAsc,
	"desc": ListIntegrationInstancesSortOrderDesc,
}

// GetListIntegrationInstancesSortOrderEnumValues Enumerates the set of values for ListIntegrationInstancesSortOrderEnum
func GetListIntegrationInstancesSortOrderEnumValues() []ListIntegrationInstancesSortOrderEnum {
	values := make([]ListIntegrationInstancesSortOrderEnum, 0)
	for _, v := range mappingListIntegrationInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIntegrationInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListIntegrationInstancesSortOrderEnum
func GetListIntegrationInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIntegrationInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIntegrationInstancesSortOrderEnum(val string) (ListIntegrationInstancesSortOrderEnum, bool) {
	enum, ok := mappingListIntegrationInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIntegrationInstancesSortByEnum Enum with underlying type: string
type ListIntegrationInstancesSortByEnum string

// Set of constants representing the allowable values for ListIntegrationInstancesSortByEnum
const (
	ListIntegrationInstancesSortByTimecreated ListIntegrationInstancesSortByEnum = "TIMECREATED"
	ListIntegrationInstancesSortByDisplayname ListIntegrationInstancesSortByEnum = "DISPLAYNAME"
)

var mappingListIntegrationInstancesSortByEnum = map[string]ListIntegrationInstancesSortByEnum{
	"TIMECREATED": ListIntegrationInstancesSortByTimecreated,
	"DISPLAYNAME": ListIntegrationInstancesSortByDisplayname,
}

var mappingListIntegrationInstancesSortByEnumLowerCase = map[string]ListIntegrationInstancesSortByEnum{
	"timecreated": ListIntegrationInstancesSortByTimecreated,
	"displayname": ListIntegrationInstancesSortByDisplayname,
}

// GetListIntegrationInstancesSortByEnumValues Enumerates the set of values for ListIntegrationInstancesSortByEnum
func GetListIntegrationInstancesSortByEnumValues() []ListIntegrationInstancesSortByEnum {
	values := make([]ListIntegrationInstancesSortByEnum, 0)
	for _, v := range mappingListIntegrationInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIntegrationInstancesSortByEnumStringValues Enumerates the set of values in String for ListIntegrationInstancesSortByEnum
func GetListIntegrationInstancesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListIntegrationInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIntegrationInstancesSortByEnum(val string) (ListIntegrationInstancesSortByEnum, bool) {
	enum, ok := mappingListIntegrationInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
