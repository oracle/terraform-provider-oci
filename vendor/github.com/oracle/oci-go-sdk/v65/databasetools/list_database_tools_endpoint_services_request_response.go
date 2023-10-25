// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetools

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDatabaseToolsEndpointServicesRequest wrapper for the ListDatabaseToolsEndpointServices operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsEndpointServices.go.html to see an example of how to use ListDatabaseToolsEndpointServicesRequest.
type ListDatabaseToolsEndpointServicesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseToolsEndpointServicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseToolsEndpointServicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources their `lifecycleState` matches the specified `lifecycleState`.
	LifecycleState ListDatabaseToolsEndpointServicesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the entire specified name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseToolsEndpointServicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseToolsEndpointServicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseToolsEndpointServicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseToolsEndpointServicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseToolsEndpointServicesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseToolsEndpointServicesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseToolsEndpointServicesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsEndpointServicesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseToolsEndpointServicesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsEndpointServicesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDatabaseToolsEndpointServicesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseToolsEndpointServicesResponse wrapper for the ListDatabaseToolsEndpointServices operation
type ListDatabaseToolsEndpointServicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseToolsEndpointServiceCollection instances
	DatabaseToolsEndpointServiceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseToolsEndpointServicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseToolsEndpointServicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseToolsEndpointServicesSortOrderEnum Enum with underlying type: string
type ListDatabaseToolsEndpointServicesSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseToolsEndpointServicesSortOrderEnum
const (
	ListDatabaseToolsEndpointServicesSortOrderAsc  ListDatabaseToolsEndpointServicesSortOrderEnum = "ASC"
	ListDatabaseToolsEndpointServicesSortOrderDesc ListDatabaseToolsEndpointServicesSortOrderEnum = "DESC"
)

var mappingListDatabaseToolsEndpointServicesSortOrderEnum = map[string]ListDatabaseToolsEndpointServicesSortOrderEnum{
	"ASC":  ListDatabaseToolsEndpointServicesSortOrderAsc,
	"DESC": ListDatabaseToolsEndpointServicesSortOrderDesc,
}

var mappingListDatabaseToolsEndpointServicesSortOrderEnumLowerCase = map[string]ListDatabaseToolsEndpointServicesSortOrderEnum{
	"asc":  ListDatabaseToolsEndpointServicesSortOrderAsc,
	"desc": ListDatabaseToolsEndpointServicesSortOrderDesc,
}

// GetListDatabaseToolsEndpointServicesSortOrderEnumValues Enumerates the set of values for ListDatabaseToolsEndpointServicesSortOrderEnum
func GetListDatabaseToolsEndpointServicesSortOrderEnumValues() []ListDatabaseToolsEndpointServicesSortOrderEnum {
	values := make([]ListDatabaseToolsEndpointServicesSortOrderEnum, 0)
	for _, v := range mappingListDatabaseToolsEndpointServicesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsEndpointServicesSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseToolsEndpointServicesSortOrderEnum
func GetListDatabaseToolsEndpointServicesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseToolsEndpointServicesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsEndpointServicesSortOrderEnum(val string) (ListDatabaseToolsEndpointServicesSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseToolsEndpointServicesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsEndpointServicesSortByEnum Enum with underlying type: string
type ListDatabaseToolsEndpointServicesSortByEnum string

// Set of constants representing the allowable values for ListDatabaseToolsEndpointServicesSortByEnum
const (
	ListDatabaseToolsEndpointServicesSortByTimecreated ListDatabaseToolsEndpointServicesSortByEnum = "timeCreated"
	ListDatabaseToolsEndpointServicesSortByDisplayname ListDatabaseToolsEndpointServicesSortByEnum = "displayName"
)

var mappingListDatabaseToolsEndpointServicesSortByEnum = map[string]ListDatabaseToolsEndpointServicesSortByEnum{
	"timeCreated": ListDatabaseToolsEndpointServicesSortByTimecreated,
	"displayName": ListDatabaseToolsEndpointServicesSortByDisplayname,
}

var mappingListDatabaseToolsEndpointServicesSortByEnumLowerCase = map[string]ListDatabaseToolsEndpointServicesSortByEnum{
	"timecreated": ListDatabaseToolsEndpointServicesSortByTimecreated,
	"displayname": ListDatabaseToolsEndpointServicesSortByDisplayname,
}

// GetListDatabaseToolsEndpointServicesSortByEnumValues Enumerates the set of values for ListDatabaseToolsEndpointServicesSortByEnum
func GetListDatabaseToolsEndpointServicesSortByEnumValues() []ListDatabaseToolsEndpointServicesSortByEnum {
	values := make([]ListDatabaseToolsEndpointServicesSortByEnum, 0)
	for _, v := range mappingListDatabaseToolsEndpointServicesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsEndpointServicesSortByEnumStringValues Enumerates the set of values in String for ListDatabaseToolsEndpointServicesSortByEnum
func GetListDatabaseToolsEndpointServicesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatabaseToolsEndpointServicesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsEndpointServicesSortByEnum(val string) (ListDatabaseToolsEndpointServicesSortByEnum, bool) {
	enum, ok := mappingListDatabaseToolsEndpointServicesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsEndpointServicesLifecycleStateEnum Enum with underlying type: string
type ListDatabaseToolsEndpointServicesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatabaseToolsEndpointServicesLifecycleStateEnum
const (
	ListDatabaseToolsEndpointServicesLifecycleStateCreating ListDatabaseToolsEndpointServicesLifecycleStateEnum = "CREATING"
	ListDatabaseToolsEndpointServicesLifecycleStateUpdating ListDatabaseToolsEndpointServicesLifecycleStateEnum = "UPDATING"
	ListDatabaseToolsEndpointServicesLifecycleStateActive   ListDatabaseToolsEndpointServicesLifecycleStateEnum = "ACTIVE"
	ListDatabaseToolsEndpointServicesLifecycleStateDeleting ListDatabaseToolsEndpointServicesLifecycleStateEnum = "DELETING"
	ListDatabaseToolsEndpointServicesLifecycleStateDeleted  ListDatabaseToolsEndpointServicesLifecycleStateEnum = "DELETED"
	ListDatabaseToolsEndpointServicesLifecycleStateFailed   ListDatabaseToolsEndpointServicesLifecycleStateEnum = "FAILED"
	ListDatabaseToolsEndpointServicesLifecycleStateInactive ListDatabaseToolsEndpointServicesLifecycleStateEnum = "INACTIVE"
)

var mappingListDatabaseToolsEndpointServicesLifecycleStateEnum = map[string]ListDatabaseToolsEndpointServicesLifecycleStateEnum{
	"CREATING": ListDatabaseToolsEndpointServicesLifecycleStateCreating,
	"UPDATING": ListDatabaseToolsEndpointServicesLifecycleStateUpdating,
	"ACTIVE":   ListDatabaseToolsEndpointServicesLifecycleStateActive,
	"DELETING": ListDatabaseToolsEndpointServicesLifecycleStateDeleting,
	"DELETED":  ListDatabaseToolsEndpointServicesLifecycleStateDeleted,
	"FAILED":   ListDatabaseToolsEndpointServicesLifecycleStateFailed,
	"INACTIVE": ListDatabaseToolsEndpointServicesLifecycleStateInactive,
}

var mappingListDatabaseToolsEndpointServicesLifecycleStateEnumLowerCase = map[string]ListDatabaseToolsEndpointServicesLifecycleStateEnum{
	"creating": ListDatabaseToolsEndpointServicesLifecycleStateCreating,
	"updating": ListDatabaseToolsEndpointServicesLifecycleStateUpdating,
	"active":   ListDatabaseToolsEndpointServicesLifecycleStateActive,
	"deleting": ListDatabaseToolsEndpointServicesLifecycleStateDeleting,
	"deleted":  ListDatabaseToolsEndpointServicesLifecycleStateDeleted,
	"failed":   ListDatabaseToolsEndpointServicesLifecycleStateFailed,
	"inactive": ListDatabaseToolsEndpointServicesLifecycleStateInactive,
}

// GetListDatabaseToolsEndpointServicesLifecycleStateEnumValues Enumerates the set of values for ListDatabaseToolsEndpointServicesLifecycleStateEnum
func GetListDatabaseToolsEndpointServicesLifecycleStateEnumValues() []ListDatabaseToolsEndpointServicesLifecycleStateEnum {
	values := make([]ListDatabaseToolsEndpointServicesLifecycleStateEnum, 0)
	for _, v := range mappingListDatabaseToolsEndpointServicesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsEndpointServicesLifecycleStateEnumStringValues Enumerates the set of values in String for ListDatabaseToolsEndpointServicesLifecycleStateEnum
func GetListDatabaseToolsEndpointServicesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingListDatabaseToolsEndpointServicesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsEndpointServicesLifecycleStateEnum(val string) (ListDatabaseToolsEndpointServicesLifecycleStateEnum, bool) {
	enum, ok := mappingListDatabaseToolsEndpointServicesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
