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

// ListTablesRequest wrapper for the ListTables operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/nosql/ListTables.go.html to see an example of how to use ListTablesRequest.
type ListTablesRequest struct {

	// The ID of a table's compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A shell-globbing-style (*?[]) filter for names.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start
	// retrieving results. This is usually retrieved from a previous
	// list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListTablesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be
	// provided. Default order for timeCreated is descending. Default
	// order for name is ascending. If no value is specified
	// timeCreated is default.
	SortBy ListTablesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Filter list by the lifecycle state of the item.
	LifecycleState ListTablesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTablesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTablesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTablesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTablesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTablesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTablesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTablesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTablesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTablesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTablesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTablesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTablesResponse wrapper for the ListTables operation
type ListTablesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TableCollection instances
	TableCollection `presentIn:"body"`

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

func (response ListTablesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTablesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTablesSortOrderEnum Enum with underlying type: string
type ListTablesSortOrderEnum string

// Set of constants representing the allowable values for ListTablesSortOrderEnum
const (
	ListTablesSortOrderAsc  ListTablesSortOrderEnum = "ASC"
	ListTablesSortOrderDesc ListTablesSortOrderEnum = "DESC"
)

var mappingListTablesSortOrderEnum = map[string]ListTablesSortOrderEnum{
	"ASC":  ListTablesSortOrderAsc,
	"DESC": ListTablesSortOrderDesc,
}

var mappingListTablesSortOrderEnumLowerCase = map[string]ListTablesSortOrderEnum{
	"asc":  ListTablesSortOrderAsc,
	"desc": ListTablesSortOrderDesc,
}

// GetListTablesSortOrderEnumValues Enumerates the set of values for ListTablesSortOrderEnum
func GetListTablesSortOrderEnumValues() []ListTablesSortOrderEnum {
	values := make([]ListTablesSortOrderEnum, 0)
	for _, v := range mappingListTablesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTablesSortOrderEnumStringValues Enumerates the set of values in String for ListTablesSortOrderEnum
func GetListTablesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTablesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTablesSortOrderEnum(val string) (ListTablesSortOrderEnum, bool) {
	enum, ok := mappingListTablesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTablesSortByEnum Enum with underlying type: string
type ListTablesSortByEnum string

// Set of constants representing the allowable values for ListTablesSortByEnum
const (
	ListTablesSortByTimecreated ListTablesSortByEnum = "timeCreated"
	ListTablesSortByName        ListTablesSortByEnum = "name"
)

var mappingListTablesSortByEnum = map[string]ListTablesSortByEnum{
	"timeCreated": ListTablesSortByTimecreated,
	"name":        ListTablesSortByName,
}

var mappingListTablesSortByEnumLowerCase = map[string]ListTablesSortByEnum{
	"timecreated": ListTablesSortByTimecreated,
	"name":        ListTablesSortByName,
}

// GetListTablesSortByEnumValues Enumerates the set of values for ListTablesSortByEnum
func GetListTablesSortByEnumValues() []ListTablesSortByEnum {
	values := make([]ListTablesSortByEnum, 0)
	for _, v := range mappingListTablesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTablesSortByEnumStringValues Enumerates the set of values in String for ListTablesSortByEnum
func GetListTablesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListTablesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTablesSortByEnum(val string) (ListTablesSortByEnum, bool) {
	enum, ok := mappingListTablesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTablesLifecycleStateEnum Enum with underlying type: string
type ListTablesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTablesLifecycleStateEnum
const (
	ListTablesLifecycleStateAll      ListTablesLifecycleStateEnum = "ALL"
	ListTablesLifecycleStateCreating ListTablesLifecycleStateEnum = "CREATING"
	ListTablesLifecycleStateUpdating ListTablesLifecycleStateEnum = "UPDATING"
	ListTablesLifecycleStateActive   ListTablesLifecycleStateEnum = "ACTIVE"
	ListTablesLifecycleStateDeleting ListTablesLifecycleStateEnum = "DELETING"
	ListTablesLifecycleStateDeleted  ListTablesLifecycleStateEnum = "DELETED"
	ListTablesLifecycleStateFailed   ListTablesLifecycleStateEnum = "FAILED"
	ListTablesLifecycleStateInactive ListTablesLifecycleStateEnum = "INACTIVE"
)

var mappingListTablesLifecycleStateEnum = map[string]ListTablesLifecycleStateEnum{
	"ALL":      ListTablesLifecycleStateAll,
	"CREATING": ListTablesLifecycleStateCreating,
	"UPDATING": ListTablesLifecycleStateUpdating,
	"ACTIVE":   ListTablesLifecycleStateActive,
	"DELETING": ListTablesLifecycleStateDeleting,
	"DELETED":  ListTablesLifecycleStateDeleted,
	"FAILED":   ListTablesLifecycleStateFailed,
	"INACTIVE": ListTablesLifecycleStateInactive,
}

var mappingListTablesLifecycleStateEnumLowerCase = map[string]ListTablesLifecycleStateEnum{
	"all":      ListTablesLifecycleStateAll,
	"creating": ListTablesLifecycleStateCreating,
	"updating": ListTablesLifecycleStateUpdating,
	"active":   ListTablesLifecycleStateActive,
	"deleting": ListTablesLifecycleStateDeleting,
	"deleted":  ListTablesLifecycleStateDeleted,
	"failed":   ListTablesLifecycleStateFailed,
	"inactive": ListTablesLifecycleStateInactive,
}

// GetListTablesLifecycleStateEnumValues Enumerates the set of values for ListTablesLifecycleStateEnum
func GetListTablesLifecycleStateEnumValues() []ListTablesLifecycleStateEnum {
	values := make([]ListTablesLifecycleStateEnum, 0)
	for _, v := range mappingListTablesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTablesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTablesLifecycleStateEnum
func GetListTablesLifecycleStateEnumStringValues() []string {
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

// GetMappingListTablesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTablesLifecycleStateEnum(val string) (ListTablesLifecycleStateEnum, bool) {
	enum, ok := mappingListTablesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
