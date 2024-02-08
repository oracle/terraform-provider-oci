// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDynamicGroupsRequest wrapper for the ListDynamicGroups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListDynamicGroups.go.html to see an example of how to use ListDynamicGroupsRequest.
type ListDynamicGroupsRequest struct {

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListDynamicGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListDynamicGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState DynamicGroupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDynamicGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDynamicGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDynamicGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDynamicGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDynamicGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDynamicGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDynamicGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDynamicGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDynamicGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDynamicGroupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDynamicGroupLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDynamicGroupsResponse wrapper for the ListDynamicGroups operation
type ListDynamicGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DynamicGroup instances
	Items []DynamicGroup `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDynamicGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDynamicGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDynamicGroupsSortByEnum Enum with underlying type: string
type ListDynamicGroupsSortByEnum string

// Set of constants representing the allowable values for ListDynamicGroupsSortByEnum
const (
	ListDynamicGroupsSortByTimecreated ListDynamicGroupsSortByEnum = "TIMECREATED"
	ListDynamicGroupsSortByName        ListDynamicGroupsSortByEnum = "NAME"
)

var mappingListDynamicGroupsSortByEnum = map[string]ListDynamicGroupsSortByEnum{
	"TIMECREATED": ListDynamicGroupsSortByTimecreated,
	"NAME":        ListDynamicGroupsSortByName,
}

var mappingListDynamicGroupsSortByEnumLowerCase = map[string]ListDynamicGroupsSortByEnum{
	"timecreated": ListDynamicGroupsSortByTimecreated,
	"name":        ListDynamicGroupsSortByName,
}

// GetListDynamicGroupsSortByEnumValues Enumerates the set of values for ListDynamicGroupsSortByEnum
func GetListDynamicGroupsSortByEnumValues() []ListDynamicGroupsSortByEnum {
	values := make([]ListDynamicGroupsSortByEnum, 0)
	for _, v := range mappingListDynamicGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDynamicGroupsSortByEnumStringValues Enumerates the set of values in String for ListDynamicGroupsSortByEnum
func GetListDynamicGroupsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListDynamicGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDynamicGroupsSortByEnum(val string) (ListDynamicGroupsSortByEnum, bool) {
	enum, ok := mappingListDynamicGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDynamicGroupsSortOrderEnum Enum with underlying type: string
type ListDynamicGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListDynamicGroupsSortOrderEnum
const (
	ListDynamicGroupsSortOrderAsc  ListDynamicGroupsSortOrderEnum = "ASC"
	ListDynamicGroupsSortOrderDesc ListDynamicGroupsSortOrderEnum = "DESC"
)

var mappingListDynamicGroupsSortOrderEnum = map[string]ListDynamicGroupsSortOrderEnum{
	"ASC":  ListDynamicGroupsSortOrderAsc,
	"DESC": ListDynamicGroupsSortOrderDesc,
}

var mappingListDynamicGroupsSortOrderEnumLowerCase = map[string]ListDynamicGroupsSortOrderEnum{
	"asc":  ListDynamicGroupsSortOrderAsc,
	"desc": ListDynamicGroupsSortOrderDesc,
}

// GetListDynamicGroupsSortOrderEnumValues Enumerates the set of values for ListDynamicGroupsSortOrderEnum
func GetListDynamicGroupsSortOrderEnumValues() []ListDynamicGroupsSortOrderEnum {
	values := make([]ListDynamicGroupsSortOrderEnum, 0)
	for _, v := range mappingListDynamicGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDynamicGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListDynamicGroupsSortOrderEnum
func GetListDynamicGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDynamicGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDynamicGroupsSortOrderEnum(val string) (ListDynamicGroupsSortOrderEnum, bool) {
	enum, ok := mappingListDynamicGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
