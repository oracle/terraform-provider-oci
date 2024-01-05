// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListConnectionsRequest wrapper for the ListConnections operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListConnections.go.html to see an example of how to use ListConnectionsRequest.
type ListConnectionsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

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
	SortBy ListConnectionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListConnectionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The current state of the Database Migration Deployment.
	LifecycleState ListConnectionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConnectionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConnectionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConnectionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConnectionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConnectionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListConnectionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConnectionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConnectionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListConnectionsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConnectionsResponse wrapper for the ListConnections operation
type ListConnectionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConnectionCollection instances
	ConnectionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConnectionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConnectionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConnectionsSortByEnum Enum with underlying type: string
type ListConnectionsSortByEnum string

// Set of constants representing the allowable values for ListConnectionsSortByEnum
const (
	ListConnectionsSortByTimecreated ListConnectionsSortByEnum = "timeCreated"
	ListConnectionsSortByDisplayname ListConnectionsSortByEnum = "displayName"
)

var mappingListConnectionsSortByEnum = map[string]ListConnectionsSortByEnum{
	"timeCreated": ListConnectionsSortByTimecreated,
	"displayName": ListConnectionsSortByDisplayname,
}

var mappingListConnectionsSortByEnumLowerCase = map[string]ListConnectionsSortByEnum{
	"timecreated": ListConnectionsSortByTimecreated,
	"displayname": ListConnectionsSortByDisplayname,
}

// GetListConnectionsSortByEnumValues Enumerates the set of values for ListConnectionsSortByEnum
func GetListConnectionsSortByEnumValues() []ListConnectionsSortByEnum {
	values := make([]ListConnectionsSortByEnum, 0)
	for _, v := range mappingListConnectionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsSortByEnumStringValues Enumerates the set of values in String for ListConnectionsSortByEnum
func GetListConnectionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListConnectionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsSortByEnum(val string) (ListConnectionsSortByEnum, bool) {
	enum, ok := mappingListConnectionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectionsSortOrderEnum Enum with underlying type: string
type ListConnectionsSortOrderEnum string

// Set of constants representing the allowable values for ListConnectionsSortOrderEnum
const (
	ListConnectionsSortOrderAsc  ListConnectionsSortOrderEnum = "ASC"
	ListConnectionsSortOrderDesc ListConnectionsSortOrderEnum = "DESC"
)

var mappingListConnectionsSortOrderEnum = map[string]ListConnectionsSortOrderEnum{
	"ASC":  ListConnectionsSortOrderAsc,
	"DESC": ListConnectionsSortOrderDesc,
}

var mappingListConnectionsSortOrderEnumLowerCase = map[string]ListConnectionsSortOrderEnum{
	"asc":  ListConnectionsSortOrderAsc,
	"desc": ListConnectionsSortOrderDesc,
}

// GetListConnectionsSortOrderEnumValues Enumerates the set of values for ListConnectionsSortOrderEnum
func GetListConnectionsSortOrderEnumValues() []ListConnectionsSortOrderEnum {
	values := make([]ListConnectionsSortOrderEnum, 0)
	for _, v := range mappingListConnectionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsSortOrderEnumStringValues Enumerates the set of values in String for ListConnectionsSortOrderEnum
func GetListConnectionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConnectionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsSortOrderEnum(val string) (ListConnectionsSortOrderEnum, bool) {
	enum, ok := mappingListConnectionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectionsLifecycleStateEnum Enum with underlying type: string
type ListConnectionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListConnectionsLifecycleStateEnum
const (
	ListConnectionsLifecycleStateCreating ListConnectionsLifecycleStateEnum = "CREATING"
	ListConnectionsLifecycleStateUpdating ListConnectionsLifecycleStateEnum = "UPDATING"
	ListConnectionsLifecycleStateActive   ListConnectionsLifecycleStateEnum = "ACTIVE"
	ListConnectionsLifecycleStateInactive ListConnectionsLifecycleStateEnum = "INACTIVE"
	ListConnectionsLifecycleStateDeleting ListConnectionsLifecycleStateEnum = "DELETING"
	ListConnectionsLifecycleStateDeleted  ListConnectionsLifecycleStateEnum = "DELETED"
	ListConnectionsLifecycleStateFailed   ListConnectionsLifecycleStateEnum = "FAILED"
)

var mappingListConnectionsLifecycleStateEnum = map[string]ListConnectionsLifecycleStateEnum{
	"CREATING": ListConnectionsLifecycleStateCreating,
	"UPDATING": ListConnectionsLifecycleStateUpdating,
	"ACTIVE":   ListConnectionsLifecycleStateActive,
	"INACTIVE": ListConnectionsLifecycleStateInactive,
	"DELETING": ListConnectionsLifecycleStateDeleting,
	"DELETED":  ListConnectionsLifecycleStateDeleted,
	"FAILED":   ListConnectionsLifecycleStateFailed,
}

var mappingListConnectionsLifecycleStateEnumLowerCase = map[string]ListConnectionsLifecycleStateEnum{
	"creating": ListConnectionsLifecycleStateCreating,
	"updating": ListConnectionsLifecycleStateUpdating,
	"active":   ListConnectionsLifecycleStateActive,
	"inactive": ListConnectionsLifecycleStateInactive,
	"deleting": ListConnectionsLifecycleStateDeleting,
	"deleted":  ListConnectionsLifecycleStateDeleted,
	"failed":   ListConnectionsLifecycleStateFailed,
}

// GetListConnectionsLifecycleStateEnumValues Enumerates the set of values for ListConnectionsLifecycleStateEnum
func GetListConnectionsLifecycleStateEnumValues() []ListConnectionsLifecycleStateEnum {
	values := make([]ListConnectionsLifecycleStateEnum, 0)
	for _, v := range mappingListConnectionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListConnectionsLifecycleStateEnum
func GetListConnectionsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListConnectionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsLifecycleStateEnum(val string) (ListConnectionsLifecycleStateEnum, bool) {
	enum, ok := mappingListConnectionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
