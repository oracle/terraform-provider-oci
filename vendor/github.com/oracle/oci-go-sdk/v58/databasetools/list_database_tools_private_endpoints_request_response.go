// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetools

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListDatabaseToolsPrivateEndpointsRequest wrapper for the ListDatabaseToolsPrivateEndpoints operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsPrivateEndpoints.go.html to see an example of how to use ListDatabaseToolsPrivateEndpointsRequest.
type ListDatabaseToolsPrivateEndpointsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their subnetId matches the given subnetId.
	SubnetId *string `mandatory:"false" contributesTo:"query" name:"subnetId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseToolsPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseToolsPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources their type matches the given type.
	EndpointServiceId *string `mandatory:"false" contributesTo:"query" name:"endpointServiceId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListDatabaseToolsPrivateEndpointsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseToolsPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseToolsPrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseToolsPrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseToolsPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseToolsPrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseToolsPrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseToolsPrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsPrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseToolsPrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsPrivateEndpointsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDatabaseToolsPrivateEndpointsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseToolsPrivateEndpointsResponse wrapper for the ListDatabaseToolsPrivateEndpoints operation
type ListDatabaseToolsPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseToolsPrivateEndpointCollection instances
	DatabaseToolsPrivateEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseToolsPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseToolsPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseToolsPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListDatabaseToolsPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseToolsPrivateEndpointsSortOrderEnum
const (
	ListDatabaseToolsPrivateEndpointsSortOrderAsc  ListDatabaseToolsPrivateEndpointsSortOrderEnum = "ASC"
	ListDatabaseToolsPrivateEndpointsSortOrderDesc ListDatabaseToolsPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListDatabaseToolsPrivateEndpointsSortOrderEnum = map[string]ListDatabaseToolsPrivateEndpointsSortOrderEnum{
	"ASC":  ListDatabaseToolsPrivateEndpointsSortOrderAsc,
	"DESC": ListDatabaseToolsPrivateEndpointsSortOrderDesc,
}

// GetListDatabaseToolsPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListDatabaseToolsPrivateEndpointsSortOrderEnum
func GetListDatabaseToolsPrivateEndpointsSortOrderEnumValues() []ListDatabaseToolsPrivateEndpointsSortOrderEnum {
	values := make([]ListDatabaseToolsPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseToolsPrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsPrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseToolsPrivateEndpointsSortOrderEnum
func GetListDatabaseToolsPrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseToolsPrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsPrivateEndpointsSortOrderEnum(val string) (ListDatabaseToolsPrivateEndpointsSortOrderEnum, bool) {
	mappingListDatabaseToolsPrivateEndpointsSortOrderEnumIgnoreCase := make(map[string]ListDatabaseToolsPrivateEndpointsSortOrderEnum)
	for k, v := range mappingListDatabaseToolsPrivateEndpointsSortOrderEnum {
		mappingListDatabaseToolsPrivateEndpointsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDatabaseToolsPrivateEndpointsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsPrivateEndpointsSortByEnum Enum with underlying type: string
type ListDatabaseToolsPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseToolsPrivateEndpointsSortByEnum
const (
	ListDatabaseToolsPrivateEndpointsSortByTimecreated ListDatabaseToolsPrivateEndpointsSortByEnum = "timeCreated"
	ListDatabaseToolsPrivateEndpointsSortByDisplayname ListDatabaseToolsPrivateEndpointsSortByEnum = "displayName"
)

var mappingListDatabaseToolsPrivateEndpointsSortByEnum = map[string]ListDatabaseToolsPrivateEndpointsSortByEnum{
	"timeCreated": ListDatabaseToolsPrivateEndpointsSortByTimecreated,
	"displayName": ListDatabaseToolsPrivateEndpointsSortByDisplayname,
}

// GetListDatabaseToolsPrivateEndpointsSortByEnumValues Enumerates the set of values for ListDatabaseToolsPrivateEndpointsSortByEnum
func GetListDatabaseToolsPrivateEndpointsSortByEnumValues() []ListDatabaseToolsPrivateEndpointsSortByEnum {
	values := make([]ListDatabaseToolsPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListDatabaseToolsPrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsPrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListDatabaseToolsPrivateEndpointsSortByEnum
func GetListDatabaseToolsPrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatabaseToolsPrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsPrivateEndpointsSortByEnum(val string) (ListDatabaseToolsPrivateEndpointsSortByEnum, bool) {
	mappingListDatabaseToolsPrivateEndpointsSortByEnumIgnoreCase := make(map[string]ListDatabaseToolsPrivateEndpointsSortByEnum)
	for k, v := range mappingListDatabaseToolsPrivateEndpointsSortByEnum {
		mappingListDatabaseToolsPrivateEndpointsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDatabaseToolsPrivateEndpointsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsPrivateEndpointsLifecycleStateEnum Enum with underlying type: string
type ListDatabaseToolsPrivateEndpointsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatabaseToolsPrivateEndpointsLifecycleStateEnum
const (
	ListDatabaseToolsPrivateEndpointsLifecycleStateCreating ListDatabaseToolsPrivateEndpointsLifecycleStateEnum = "CREATING"
	ListDatabaseToolsPrivateEndpointsLifecycleStateUpdating ListDatabaseToolsPrivateEndpointsLifecycleStateEnum = "UPDATING"
	ListDatabaseToolsPrivateEndpointsLifecycleStateActive   ListDatabaseToolsPrivateEndpointsLifecycleStateEnum = "ACTIVE"
	ListDatabaseToolsPrivateEndpointsLifecycleStateDeleting ListDatabaseToolsPrivateEndpointsLifecycleStateEnum = "DELETING"
	ListDatabaseToolsPrivateEndpointsLifecycleStateDeleted  ListDatabaseToolsPrivateEndpointsLifecycleStateEnum = "DELETED"
	ListDatabaseToolsPrivateEndpointsLifecycleStateFailed   ListDatabaseToolsPrivateEndpointsLifecycleStateEnum = "FAILED"
)

var mappingListDatabaseToolsPrivateEndpointsLifecycleStateEnum = map[string]ListDatabaseToolsPrivateEndpointsLifecycleStateEnum{
	"CREATING": ListDatabaseToolsPrivateEndpointsLifecycleStateCreating,
	"UPDATING": ListDatabaseToolsPrivateEndpointsLifecycleStateUpdating,
	"ACTIVE":   ListDatabaseToolsPrivateEndpointsLifecycleStateActive,
	"DELETING": ListDatabaseToolsPrivateEndpointsLifecycleStateDeleting,
	"DELETED":  ListDatabaseToolsPrivateEndpointsLifecycleStateDeleted,
	"FAILED":   ListDatabaseToolsPrivateEndpointsLifecycleStateFailed,
}

// GetListDatabaseToolsPrivateEndpointsLifecycleStateEnumValues Enumerates the set of values for ListDatabaseToolsPrivateEndpointsLifecycleStateEnum
func GetListDatabaseToolsPrivateEndpointsLifecycleStateEnumValues() []ListDatabaseToolsPrivateEndpointsLifecycleStateEnum {
	values := make([]ListDatabaseToolsPrivateEndpointsLifecycleStateEnum, 0)
	for _, v := range mappingListDatabaseToolsPrivateEndpointsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsPrivateEndpointsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDatabaseToolsPrivateEndpointsLifecycleStateEnum
func GetListDatabaseToolsPrivateEndpointsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListDatabaseToolsPrivateEndpointsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsPrivateEndpointsLifecycleStateEnum(val string) (ListDatabaseToolsPrivateEndpointsLifecycleStateEnum, bool) {
	mappingListDatabaseToolsPrivateEndpointsLifecycleStateEnumIgnoreCase := make(map[string]ListDatabaseToolsPrivateEndpointsLifecycleStateEnum)
	for k, v := range mappingListDatabaseToolsPrivateEndpointsLifecycleStateEnum {
		mappingListDatabaseToolsPrivateEndpointsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDatabaseToolsPrivateEndpointsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
