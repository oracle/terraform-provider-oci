// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package nosql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListIndexesRequest wrapper for the ListIndexes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/nosql/ListIndexes.go.html to see an example of how to use ListIndexesRequest.
type ListIndexesRequest struct {

	// A table name within the compartment, or a table OCID.
	TableNameOrId *string `mandatory:"true" contributesTo:"path" name:"tableNameOrId"`

	// The ID of a table's compartment. When a table is identified
	// by name, the compartmentId is often needed to provide
	// context for interpreting the name.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A shell-globbing-style (*?[]) filter for names.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Filter list by the lifecycle state of the item.
	LifecycleState ListIndexesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start
	// retrieving results. This is usually retrieved from a previous
	// list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListIndexesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be
	// provided. Default order for timeCreated is descending. Default
	// order for name is ascending. If no value is specified
	// timeCreated is default.
	SortBy ListIndexesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIndexesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIndexesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIndexesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIndexesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIndexesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListIndexesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListIndexesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIndexesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIndexesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIndexesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIndexesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIndexesResponse wrapper for the ListIndexes operation
type ListIndexesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of IndexCollection instances
	IndexCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list,
	// if this header appears in the response, then a partial list
	// might have been returned. Include this value as the `page`
	// parameter for the subsequent GET request to get the next batch
	// of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListIndexesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIndexesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIndexesLifecycleStateEnum Enum with underlying type: string
type ListIndexesLifecycleStateEnum string

// Set of constants representing the allowable values for ListIndexesLifecycleStateEnum
const (
	ListIndexesLifecycleStateAll      ListIndexesLifecycleStateEnum = "ALL"
	ListIndexesLifecycleStateCreating ListIndexesLifecycleStateEnum = "CREATING"
	ListIndexesLifecycleStateUpdating ListIndexesLifecycleStateEnum = "UPDATING"
	ListIndexesLifecycleStateActive   ListIndexesLifecycleStateEnum = "ACTIVE"
	ListIndexesLifecycleStateDeleting ListIndexesLifecycleStateEnum = "DELETING"
	ListIndexesLifecycleStateDeleted  ListIndexesLifecycleStateEnum = "DELETED"
	ListIndexesLifecycleStateFailed   ListIndexesLifecycleStateEnum = "FAILED"
	ListIndexesLifecycleStateInactive ListIndexesLifecycleStateEnum = "INACTIVE"
)

var mappingListIndexesLifecycleStateEnum = map[string]ListIndexesLifecycleStateEnum{
	"ALL":      ListIndexesLifecycleStateAll,
	"CREATING": ListIndexesLifecycleStateCreating,
	"UPDATING": ListIndexesLifecycleStateUpdating,
	"ACTIVE":   ListIndexesLifecycleStateActive,
	"DELETING": ListIndexesLifecycleStateDeleting,
	"DELETED":  ListIndexesLifecycleStateDeleted,
	"FAILED":   ListIndexesLifecycleStateFailed,
	"INACTIVE": ListIndexesLifecycleStateInactive,
}

var mappingListIndexesLifecycleStateEnumLowerCase = map[string]ListIndexesLifecycleStateEnum{
	"all":      ListIndexesLifecycleStateAll,
	"creating": ListIndexesLifecycleStateCreating,
	"updating": ListIndexesLifecycleStateUpdating,
	"active":   ListIndexesLifecycleStateActive,
	"deleting": ListIndexesLifecycleStateDeleting,
	"deleted":  ListIndexesLifecycleStateDeleted,
	"failed":   ListIndexesLifecycleStateFailed,
	"inactive": ListIndexesLifecycleStateInactive,
}

// GetListIndexesLifecycleStateEnumValues Enumerates the set of values for ListIndexesLifecycleStateEnum
func GetListIndexesLifecycleStateEnumValues() []ListIndexesLifecycleStateEnum {
	values := make([]ListIndexesLifecycleStateEnum, 0)
	for _, v := range mappingListIndexesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListIndexesLifecycleStateEnumStringValues Enumerates the set of values in String for ListIndexesLifecycleStateEnum
func GetListIndexesLifecycleStateEnumStringValues() []string {
	return []string{
		"ALL",
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingListIndexesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIndexesLifecycleStateEnum(val string) (ListIndexesLifecycleStateEnum, bool) {
	enum, ok := mappingListIndexesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIndexesSortOrderEnum Enum with underlying type: string
type ListIndexesSortOrderEnum string

// Set of constants representing the allowable values for ListIndexesSortOrderEnum
const (
	ListIndexesSortOrderAsc  ListIndexesSortOrderEnum = "ASC"
	ListIndexesSortOrderDesc ListIndexesSortOrderEnum = "DESC"
)

var mappingListIndexesSortOrderEnum = map[string]ListIndexesSortOrderEnum{
	"ASC":  ListIndexesSortOrderAsc,
	"DESC": ListIndexesSortOrderDesc,
}

var mappingListIndexesSortOrderEnumLowerCase = map[string]ListIndexesSortOrderEnum{
	"asc":  ListIndexesSortOrderAsc,
	"desc": ListIndexesSortOrderDesc,
}

// GetListIndexesSortOrderEnumValues Enumerates the set of values for ListIndexesSortOrderEnum
func GetListIndexesSortOrderEnumValues() []ListIndexesSortOrderEnum {
	values := make([]ListIndexesSortOrderEnum, 0)
	for _, v := range mappingListIndexesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIndexesSortOrderEnumStringValues Enumerates the set of values in String for ListIndexesSortOrderEnum
func GetListIndexesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIndexesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIndexesSortOrderEnum(val string) (ListIndexesSortOrderEnum, bool) {
	enum, ok := mappingListIndexesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIndexesSortByEnum Enum with underlying type: string
type ListIndexesSortByEnum string

// Set of constants representing the allowable values for ListIndexesSortByEnum
const (
	ListIndexesSortByTimecreated ListIndexesSortByEnum = "timeCreated"
	ListIndexesSortByName        ListIndexesSortByEnum = "name"
)

var mappingListIndexesSortByEnum = map[string]ListIndexesSortByEnum{
	"timeCreated": ListIndexesSortByTimecreated,
	"name":        ListIndexesSortByName,
}

var mappingListIndexesSortByEnumLowerCase = map[string]ListIndexesSortByEnum{
	"timecreated": ListIndexesSortByTimecreated,
	"name":        ListIndexesSortByName,
}

// GetListIndexesSortByEnumValues Enumerates the set of values for ListIndexesSortByEnum
func GetListIndexesSortByEnumValues() []ListIndexesSortByEnum {
	values := make([]ListIndexesSortByEnum, 0)
	for _, v := range mappingListIndexesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIndexesSortByEnumStringValues Enumerates the set of values in String for ListIndexesSortByEnum
func GetListIndexesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListIndexesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIndexesSortByEnum(val string) (ListIndexesSortByEnum, bool) {
	enum, ok := mappingListIndexesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
