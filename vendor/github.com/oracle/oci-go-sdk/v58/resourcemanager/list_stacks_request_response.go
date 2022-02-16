// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListStacksRequest wrapper for the ListStacks operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListStacks.go.html to see an example of how to use ListStacksRequest.
type ListStacksRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that exist in the compartment, identified by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on which to query for a stack.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter that returns only those resources that match the specified
	// lifecycle state. The state value is case-insensitive.
	// For more information about stack lifecycle states, see
	// Key Concepts (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#concepts__StackStates).
	// Allowable values:
	// - CREATING
	// - ACTIVE
	// - DELETING
	// - DELETED
	// - FAILED
	LifecycleState StackLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	// Use this filter to list a resource by name.
	// Requires `sortBy` set to `DISPLAYNAME`.
	// Alternatively, when you know the resource OCID, use the related Get operation.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to use when sorting returned resources.
	// By default, `TIMECREATED` is ordered descending.
	// By default, `DISPLAYNAME` is ordered ascending. Note that you can sort only on one field.
	SortBy ListStacksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use when sorting returned resources. Ascending (`ASC`) or descending (`DESC`).
	SortOrder ListStacksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The number of items returned in a paginated `List` call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the preceding `List` call.
	// For information about pagination, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStacksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStacksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStacksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStacksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStacksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStackLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetStackLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStacksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStacksSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStacksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStacksSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStacksResponse wrapper for the ListStacks operation
type ListStacksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []StackSummary instances
	Items []StackSummary `presentIn:"body"`

	// Unique identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of paginated list items. If the `opc-next-page`
	// header appears in the response, additional pages of results remain.
	// To receive the next page, include the header value in the `page` param.
	// If the `opc-next-page` header does not appear in the response, there
	// are no more list items to get. For more information about list pagination,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListStacksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStacksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStacksSortByEnum Enum with underlying type: string
type ListStacksSortByEnum string

// Set of constants representing the allowable values for ListStacksSortByEnum
const (
	ListStacksSortByTimecreated ListStacksSortByEnum = "TIMECREATED"
	ListStacksSortByDisplayname ListStacksSortByEnum = "DISPLAYNAME"
)

var mappingListStacksSortByEnum = map[string]ListStacksSortByEnum{
	"TIMECREATED": ListStacksSortByTimecreated,
	"DISPLAYNAME": ListStacksSortByDisplayname,
}

// GetListStacksSortByEnumValues Enumerates the set of values for ListStacksSortByEnum
func GetListStacksSortByEnumValues() []ListStacksSortByEnum {
	values := make([]ListStacksSortByEnum, 0)
	for _, v := range mappingListStacksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStacksSortByEnumStringValues Enumerates the set of values in String for ListStacksSortByEnum
func GetListStacksSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListStacksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStacksSortByEnum(val string) (ListStacksSortByEnum, bool) {
	mappingListStacksSortByEnumIgnoreCase := make(map[string]ListStacksSortByEnum)
	for k, v := range mappingListStacksSortByEnum {
		mappingListStacksSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListStacksSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListStacksSortOrderEnum Enum with underlying type: string
type ListStacksSortOrderEnum string

// Set of constants representing the allowable values for ListStacksSortOrderEnum
const (
	ListStacksSortOrderAsc  ListStacksSortOrderEnum = "ASC"
	ListStacksSortOrderDesc ListStacksSortOrderEnum = "DESC"
)

var mappingListStacksSortOrderEnum = map[string]ListStacksSortOrderEnum{
	"ASC":  ListStacksSortOrderAsc,
	"DESC": ListStacksSortOrderDesc,
}

// GetListStacksSortOrderEnumValues Enumerates the set of values for ListStacksSortOrderEnum
func GetListStacksSortOrderEnumValues() []ListStacksSortOrderEnum {
	values := make([]ListStacksSortOrderEnum, 0)
	for _, v := range mappingListStacksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStacksSortOrderEnumStringValues Enumerates the set of values in String for ListStacksSortOrderEnum
func GetListStacksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStacksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStacksSortOrderEnum(val string) (ListStacksSortOrderEnum, bool) {
	mappingListStacksSortOrderEnumIgnoreCase := make(map[string]ListStacksSortOrderEnum)
	for k, v := range mappingListStacksSortOrderEnum {
		mappingListStacksSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListStacksSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
