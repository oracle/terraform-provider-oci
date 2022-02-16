// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListApisRequest wrapper for the ListApis operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/ListApis.go.html to see an example of how to use ListApisRequest.
type ListApisRequest struct {

	// The ocid of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given lifecycle state.
	// Example: `ACTIVE`
	LifecycleState ApiSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'. The default order depends on the sortBy value.
	SortOrder ListApisSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for `timeCreated` is descending. Default order for
	// `displayName` is ascending. The `displayName` sort order is case
	// sensitive.
	SortBy ListApisSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request id for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApisRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApisRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApisRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApisRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListApisRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApiSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetApiSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApisSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListApisSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApisSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListApisSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListApisResponse wrapper for the ListApis operation
type ListApisResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ApiCollection instances
	ApiCollection `presentIn:"body"`

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

func (response ListApisResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApisResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApisSortOrderEnum Enum with underlying type: string
type ListApisSortOrderEnum string

// Set of constants representing the allowable values for ListApisSortOrderEnum
const (
	ListApisSortOrderAsc  ListApisSortOrderEnum = "ASC"
	ListApisSortOrderDesc ListApisSortOrderEnum = "DESC"
)

var mappingListApisSortOrderEnum = map[string]ListApisSortOrderEnum{
	"ASC":  ListApisSortOrderAsc,
	"DESC": ListApisSortOrderDesc,
}

// GetListApisSortOrderEnumValues Enumerates the set of values for ListApisSortOrderEnum
func GetListApisSortOrderEnumValues() []ListApisSortOrderEnum {
	values := make([]ListApisSortOrderEnum, 0)
	for _, v := range mappingListApisSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListApisSortOrderEnumStringValues Enumerates the set of values in String for ListApisSortOrderEnum
func GetListApisSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListApisSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApisSortOrderEnum(val string) (ListApisSortOrderEnum, bool) {
	mappingListApisSortOrderEnumIgnoreCase := make(map[string]ListApisSortOrderEnum)
	for k, v := range mappingListApisSortOrderEnum {
		mappingListApisSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListApisSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListApisSortByEnum Enum with underlying type: string
type ListApisSortByEnum string

// Set of constants representing the allowable values for ListApisSortByEnum
const (
	ListApisSortByTimecreated ListApisSortByEnum = "timeCreated"
	ListApisSortByDisplayname ListApisSortByEnum = "displayName"
)

var mappingListApisSortByEnum = map[string]ListApisSortByEnum{
	"timeCreated": ListApisSortByTimecreated,
	"displayName": ListApisSortByDisplayname,
}

// GetListApisSortByEnumValues Enumerates the set of values for ListApisSortByEnum
func GetListApisSortByEnumValues() []ListApisSortByEnum {
	values := make([]ListApisSortByEnum, 0)
	for _, v := range mappingListApisSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListApisSortByEnumStringValues Enumerates the set of values in String for ListApisSortByEnum
func GetListApisSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListApisSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApisSortByEnum(val string) (ListApisSortByEnum, bool) {
	mappingListApisSortByEnumIgnoreCase := make(map[string]ListApisSortByEnum)
	for k, v := range mappingListApisSortByEnum {
		mappingListApisSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListApisSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
