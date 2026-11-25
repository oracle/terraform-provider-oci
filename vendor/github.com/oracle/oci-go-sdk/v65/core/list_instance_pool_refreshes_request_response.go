// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInstancePoolRefreshesRequest wrapper for the ListInstancePoolRefreshes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListInstancePoolRefreshes.go.html to see an example of how to use ListInstancePoolRefreshesRequest.
type ListInstancePoolRefreshesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.
	InstancePoolId *string `mandatory:"true" contributesTo:"path" name:"instancePoolId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListInstancePoolRefreshesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListInstancePoolRefreshesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return refreshes that match the given lifecycle state. The state
	// value is case-insensitive.
	LifecycleState InstancePoolRefreshLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInstancePoolRefreshesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInstancePoolRefreshesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInstancePoolRefreshesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInstancePoolRefreshesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInstancePoolRefreshesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInstancePoolRefreshesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInstancePoolRefreshesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInstancePoolRefreshesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInstancePoolRefreshesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInstancePoolRefreshLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetInstancePoolRefreshLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInstancePoolRefreshesResponse wrapper for the ListInstancePoolRefreshes operation
type ListInstancePoolRefreshesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InstancePoolRefreshCollection instances
	InstancePoolRefreshCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListInstancePoolRefreshesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInstancePoolRefreshesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInstancePoolRefreshesSortByEnum Enum with underlying type: string
type ListInstancePoolRefreshesSortByEnum string

// Set of constants representing the allowable values for ListInstancePoolRefreshesSortByEnum
const (
	ListInstancePoolRefreshesSortByTimecreated ListInstancePoolRefreshesSortByEnum = "TIMECREATED"
	ListInstancePoolRefreshesSortByDisplayname ListInstancePoolRefreshesSortByEnum = "DISPLAYNAME"
)

var mappingListInstancePoolRefreshesSortByEnum = map[string]ListInstancePoolRefreshesSortByEnum{
	"TIMECREATED": ListInstancePoolRefreshesSortByTimecreated,
	"DISPLAYNAME": ListInstancePoolRefreshesSortByDisplayname,
}

var mappingListInstancePoolRefreshesSortByEnumLowerCase = map[string]ListInstancePoolRefreshesSortByEnum{
	"timecreated": ListInstancePoolRefreshesSortByTimecreated,
	"displayname": ListInstancePoolRefreshesSortByDisplayname,
}

// GetListInstancePoolRefreshesSortByEnumValues Enumerates the set of values for ListInstancePoolRefreshesSortByEnum
func GetListInstancePoolRefreshesSortByEnumValues() []ListInstancePoolRefreshesSortByEnum {
	values := make([]ListInstancePoolRefreshesSortByEnum, 0)
	for _, v := range mappingListInstancePoolRefreshesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstancePoolRefreshesSortByEnumStringValues Enumerates the set of values in String for ListInstancePoolRefreshesSortByEnum
func GetListInstancePoolRefreshesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListInstancePoolRefreshesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstancePoolRefreshesSortByEnum(val string) (ListInstancePoolRefreshesSortByEnum, bool) {
	enum, ok := mappingListInstancePoolRefreshesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInstancePoolRefreshesSortOrderEnum Enum with underlying type: string
type ListInstancePoolRefreshesSortOrderEnum string

// Set of constants representing the allowable values for ListInstancePoolRefreshesSortOrderEnum
const (
	ListInstancePoolRefreshesSortOrderAsc  ListInstancePoolRefreshesSortOrderEnum = "ASC"
	ListInstancePoolRefreshesSortOrderDesc ListInstancePoolRefreshesSortOrderEnum = "DESC"
)

var mappingListInstancePoolRefreshesSortOrderEnum = map[string]ListInstancePoolRefreshesSortOrderEnum{
	"ASC":  ListInstancePoolRefreshesSortOrderAsc,
	"DESC": ListInstancePoolRefreshesSortOrderDesc,
}

var mappingListInstancePoolRefreshesSortOrderEnumLowerCase = map[string]ListInstancePoolRefreshesSortOrderEnum{
	"asc":  ListInstancePoolRefreshesSortOrderAsc,
	"desc": ListInstancePoolRefreshesSortOrderDesc,
}

// GetListInstancePoolRefreshesSortOrderEnumValues Enumerates the set of values for ListInstancePoolRefreshesSortOrderEnum
func GetListInstancePoolRefreshesSortOrderEnumValues() []ListInstancePoolRefreshesSortOrderEnum {
	values := make([]ListInstancePoolRefreshesSortOrderEnum, 0)
	for _, v := range mappingListInstancePoolRefreshesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstancePoolRefreshesSortOrderEnumStringValues Enumerates the set of values in String for ListInstancePoolRefreshesSortOrderEnum
func GetListInstancePoolRefreshesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInstancePoolRefreshesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstancePoolRefreshesSortOrderEnum(val string) (ListInstancePoolRefreshesSortOrderEnum, bool) {
	enum, ok := mappingListInstancePoolRefreshesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
