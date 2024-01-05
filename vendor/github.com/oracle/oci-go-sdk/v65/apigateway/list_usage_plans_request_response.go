// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListUsagePlansRequest wrapper for the ListUsagePlans operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/ListUsagePlans.go.html to see an example of how to use ListUsagePlansRequest.
type ListUsagePlansRequest struct {

	// The ocid of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given lifecycle state.
	// Example: `ACTIVE`
	LifecycleState UsagePlanLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'. The default order depends on the sortBy value.
	SortOrder ListUsagePlansSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for `timeCreated` is descending. Default order for
	// `displayName` is ascending. The `displayName` sort order is case
	// sensitive.
	SortBy ListUsagePlansSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request id for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUsagePlansRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUsagePlansRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUsagePlansRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUsagePlansRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUsagePlansRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUsagePlanLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetUsagePlanLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUsagePlansSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUsagePlansSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUsagePlansSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUsagePlansSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUsagePlansResponse wrapper for the ListUsagePlans operation
type ListUsagePlansResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UsagePlanCollection instances
	UsagePlanCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request
	// id.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain. For important details about how
	// pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response,
	// additional pages of results were seen previously. For important details
	// about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListUsagePlansResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUsagePlansResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUsagePlansSortOrderEnum Enum with underlying type: string
type ListUsagePlansSortOrderEnum string

// Set of constants representing the allowable values for ListUsagePlansSortOrderEnum
const (
	ListUsagePlansSortOrderAsc  ListUsagePlansSortOrderEnum = "ASC"
	ListUsagePlansSortOrderDesc ListUsagePlansSortOrderEnum = "DESC"
)

var mappingListUsagePlansSortOrderEnum = map[string]ListUsagePlansSortOrderEnum{
	"ASC":  ListUsagePlansSortOrderAsc,
	"DESC": ListUsagePlansSortOrderDesc,
}

var mappingListUsagePlansSortOrderEnumLowerCase = map[string]ListUsagePlansSortOrderEnum{
	"asc":  ListUsagePlansSortOrderAsc,
	"desc": ListUsagePlansSortOrderDesc,
}

// GetListUsagePlansSortOrderEnumValues Enumerates the set of values for ListUsagePlansSortOrderEnum
func GetListUsagePlansSortOrderEnumValues() []ListUsagePlansSortOrderEnum {
	values := make([]ListUsagePlansSortOrderEnum, 0)
	for _, v := range mappingListUsagePlansSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUsagePlansSortOrderEnumStringValues Enumerates the set of values in String for ListUsagePlansSortOrderEnum
func GetListUsagePlansSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUsagePlansSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUsagePlansSortOrderEnum(val string) (ListUsagePlansSortOrderEnum, bool) {
	enum, ok := mappingListUsagePlansSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUsagePlansSortByEnum Enum with underlying type: string
type ListUsagePlansSortByEnum string

// Set of constants representing the allowable values for ListUsagePlansSortByEnum
const (
	ListUsagePlansSortByTimecreated ListUsagePlansSortByEnum = "timeCreated"
	ListUsagePlansSortByDisplayname ListUsagePlansSortByEnum = "displayName"
)

var mappingListUsagePlansSortByEnum = map[string]ListUsagePlansSortByEnum{
	"timeCreated": ListUsagePlansSortByTimecreated,
	"displayName": ListUsagePlansSortByDisplayname,
}

var mappingListUsagePlansSortByEnumLowerCase = map[string]ListUsagePlansSortByEnum{
	"timecreated": ListUsagePlansSortByTimecreated,
	"displayname": ListUsagePlansSortByDisplayname,
}

// GetListUsagePlansSortByEnumValues Enumerates the set of values for ListUsagePlansSortByEnum
func GetListUsagePlansSortByEnumValues() []ListUsagePlansSortByEnum {
	values := make([]ListUsagePlansSortByEnum, 0)
	for _, v := range mappingListUsagePlansSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUsagePlansSortByEnumStringValues Enumerates the set of values in String for ListUsagePlansSortByEnum
func GetListUsagePlansSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListUsagePlansSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUsagePlansSortByEnum(val string) (ListUsagePlansSortByEnum, bool) {
	enum, ok := mappingListUsagePlansSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
