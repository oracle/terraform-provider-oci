// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSqlEndpointsRequest wrapper for the ListSqlEndpoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataflow/ListSqlEndpoints.go.html to see an example of how to use ListSqlEndpointsRequest.
type ListSqlEndpointsRequest struct {

	// The OCID of the compartment in which to query resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The unique id of the SQL Endpoint.
	SqlEndpointId *string `mandatory:"false" contributesTo:"query" name:"sqlEndpointId"`

	// A filter to return only those resources whose sqlEndpointLifecycleState matches the given sqlEndpointLifecycleState.
	LifecycleState ListSqlEndpointsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The query parameter for the Spark application name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items that can be returned.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The ordering of results in ascending or descending order.
	SortOrder ListSqlEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for timeCreated is descending. The default order for displayName is ascending. If no value is specified timeCreated is used by default.
	SortBy ListSqlEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request. If provided, the returned request ID will include this value.
	// Otherwise, a random request ID will be generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlEndpointsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSqlEndpointsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSqlEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSqlEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlEndpointsResponse wrapper for the ListSqlEndpoints operation
type ListSqlEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlEndpointCollection instances
	SqlEndpointCollection `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response,
	// additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSqlEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlEndpointsLifecycleStateEnum Enum with underlying type: string
type ListSqlEndpointsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSqlEndpointsLifecycleStateEnum
const (
	ListSqlEndpointsLifecycleStateCreating       ListSqlEndpointsLifecycleStateEnum = "CREATING"
	ListSqlEndpointsLifecycleStateActive         ListSqlEndpointsLifecycleStateEnum = "ACTIVE"
	ListSqlEndpointsLifecycleStateDeleting       ListSqlEndpointsLifecycleStateEnum = "DELETING"
	ListSqlEndpointsLifecycleStateDeleted        ListSqlEndpointsLifecycleStateEnum = "DELETED"
	ListSqlEndpointsLifecycleStateFailed         ListSqlEndpointsLifecycleStateEnum = "FAILED"
	ListSqlEndpointsLifecycleStateUpdating       ListSqlEndpointsLifecycleStateEnum = "UPDATING"
	ListSqlEndpointsLifecycleStateNeedsAttention ListSqlEndpointsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListSqlEndpointsLifecycleStateInactive       ListSqlEndpointsLifecycleStateEnum = "INACTIVE"
)

var mappingListSqlEndpointsLifecycleStateEnum = map[string]ListSqlEndpointsLifecycleStateEnum{
	"CREATING":        ListSqlEndpointsLifecycleStateCreating,
	"ACTIVE":          ListSqlEndpointsLifecycleStateActive,
	"DELETING":        ListSqlEndpointsLifecycleStateDeleting,
	"DELETED":         ListSqlEndpointsLifecycleStateDeleted,
	"FAILED":          ListSqlEndpointsLifecycleStateFailed,
	"UPDATING":        ListSqlEndpointsLifecycleStateUpdating,
	"NEEDS_ATTENTION": ListSqlEndpointsLifecycleStateNeedsAttention,
	"INACTIVE":        ListSqlEndpointsLifecycleStateInactive,
}

var mappingListSqlEndpointsLifecycleStateEnumLowerCase = map[string]ListSqlEndpointsLifecycleStateEnum{
	"creating":        ListSqlEndpointsLifecycleStateCreating,
	"active":          ListSqlEndpointsLifecycleStateActive,
	"deleting":        ListSqlEndpointsLifecycleStateDeleting,
	"deleted":         ListSqlEndpointsLifecycleStateDeleted,
	"failed":          ListSqlEndpointsLifecycleStateFailed,
	"updating":        ListSqlEndpointsLifecycleStateUpdating,
	"needs_attention": ListSqlEndpointsLifecycleStateNeedsAttention,
	"inactive":        ListSqlEndpointsLifecycleStateInactive,
}

// GetListSqlEndpointsLifecycleStateEnumValues Enumerates the set of values for ListSqlEndpointsLifecycleStateEnum
func GetListSqlEndpointsLifecycleStateEnumValues() []ListSqlEndpointsLifecycleStateEnum {
	values := make([]ListSqlEndpointsLifecycleStateEnum, 0)
	for _, v := range mappingListSqlEndpointsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlEndpointsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSqlEndpointsLifecycleStateEnum
func GetListSqlEndpointsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
		"NEEDS_ATTENTION",
		"INACTIVE",
	}
}

// GetMappingListSqlEndpointsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlEndpointsLifecycleStateEnum(val string) (ListSqlEndpointsLifecycleStateEnum, bool) {
	enum, ok := mappingListSqlEndpointsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlEndpointsSortOrderEnum Enum with underlying type: string
type ListSqlEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListSqlEndpointsSortOrderEnum
const (
	ListSqlEndpointsSortOrderAsc  ListSqlEndpointsSortOrderEnum = "ASC"
	ListSqlEndpointsSortOrderDesc ListSqlEndpointsSortOrderEnum = "DESC"
)

var mappingListSqlEndpointsSortOrderEnum = map[string]ListSqlEndpointsSortOrderEnum{
	"ASC":  ListSqlEndpointsSortOrderAsc,
	"DESC": ListSqlEndpointsSortOrderDesc,
}

var mappingListSqlEndpointsSortOrderEnumLowerCase = map[string]ListSqlEndpointsSortOrderEnum{
	"asc":  ListSqlEndpointsSortOrderAsc,
	"desc": ListSqlEndpointsSortOrderDesc,
}

// GetListSqlEndpointsSortOrderEnumValues Enumerates the set of values for ListSqlEndpointsSortOrderEnum
func GetListSqlEndpointsSortOrderEnumValues() []ListSqlEndpointsSortOrderEnum {
	values := make([]ListSqlEndpointsSortOrderEnum, 0)
	for _, v := range mappingListSqlEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListSqlEndpointsSortOrderEnum
func GetListSqlEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSqlEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlEndpointsSortOrderEnum(val string) (ListSqlEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListSqlEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlEndpointsSortByEnum Enum with underlying type: string
type ListSqlEndpointsSortByEnum string

// Set of constants representing the allowable values for ListSqlEndpointsSortByEnum
const (
	ListSqlEndpointsSortById          ListSqlEndpointsSortByEnum = "id"
	ListSqlEndpointsSortByTimecreated ListSqlEndpointsSortByEnum = "timeCreated"
	ListSqlEndpointsSortByDisplayname ListSqlEndpointsSortByEnum = "displayName"
)

var mappingListSqlEndpointsSortByEnum = map[string]ListSqlEndpointsSortByEnum{
	"id":          ListSqlEndpointsSortById,
	"timeCreated": ListSqlEndpointsSortByTimecreated,
	"displayName": ListSqlEndpointsSortByDisplayname,
}

var mappingListSqlEndpointsSortByEnumLowerCase = map[string]ListSqlEndpointsSortByEnum{
	"id":          ListSqlEndpointsSortById,
	"timecreated": ListSqlEndpointsSortByTimecreated,
	"displayname": ListSqlEndpointsSortByDisplayname,
}

// GetListSqlEndpointsSortByEnumValues Enumerates the set of values for ListSqlEndpointsSortByEnum
func GetListSqlEndpointsSortByEnumValues() []ListSqlEndpointsSortByEnum {
	values := make([]ListSqlEndpointsSortByEnum, 0)
	for _, v := range mappingListSqlEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlEndpointsSortByEnumStringValues Enumerates the set of values in String for ListSqlEndpointsSortByEnum
func GetListSqlEndpointsSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSqlEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlEndpointsSortByEnum(val string) (ListSqlEndpointsSortByEnum, bool) {
	enum, ok := mappingListSqlEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
