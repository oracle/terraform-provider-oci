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

// ListSdksRequest wrapper for the ListSdks operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/ListSdks.go.html to see an example of how to use ListSdksRequest.
type ListSdksRequest struct {

	// The ocid of the SDK.
	SdkId *string `mandatory:"false" contributesTo:"query" name:"sdkId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given lifecycle state.
	// Example: `ACTIVE` or `DELETED`
	LifecycleState SdkLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'. The default order depends on the sortBy value.
	SortOrder ListSdksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for `timeCreated` is descending. Default order for
	// `displayName` is ascending. The `displayName` sort order is case
	// sensitive.
	SortBy ListSdksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The ocid of the API.
	ApiId *string `mandatory:"false" contributesTo:"query" name:"apiId"`

	// The client request id for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSdksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSdksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSdksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSdksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSdksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSdkLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetSdkLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSdksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSdksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSdksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSdksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSdksResponse wrapper for the ListSdks operation
type ListSdksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SdkCollection instances
	SdkCollection `presentIn:"body"`

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

func (response ListSdksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSdksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSdksSortOrderEnum Enum with underlying type: string
type ListSdksSortOrderEnum string

// Set of constants representing the allowable values for ListSdksSortOrderEnum
const (
	ListSdksSortOrderAsc  ListSdksSortOrderEnum = "ASC"
	ListSdksSortOrderDesc ListSdksSortOrderEnum = "DESC"
)

var mappingListSdksSortOrderEnum = map[string]ListSdksSortOrderEnum{
	"ASC":  ListSdksSortOrderAsc,
	"DESC": ListSdksSortOrderDesc,
}

var mappingListSdksSortOrderEnumLowerCase = map[string]ListSdksSortOrderEnum{
	"asc":  ListSdksSortOrderAsc,
	"desc": ListSdksSortOrderDesc,
}

// GetListSdksSortOrderEnumValues Enumerates the set of values for ListSdksSortOrderEnum
func GetListSdksSortOrderEnumValues() []ListSdksSortOrderEnum {
	values := make([]ListSdksSortOrderEnum, 0)
	for _, v := range mappingListSdksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSdksSortOrderEnumStringValues Enumerates the set of values in String for ListSdksSortOrderEnum
func GetListSdksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSdksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSdksSortOrderEnum(val string) (ListSdksSortOrderEnum, bool) {
	enum, ok := mappingListSdksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSdksSortByEnum Enum with underlying type: string
type ListSdksSortByEnum string

// Set of constants representing the allowable values for ListSdksSortByEnum
const (
	ListSdksSortByTimecreated ListSdksSortByEnum = "timeCreated"
	ListSdksSortByDisplayname ListSdksSortByEnum = "displayName"
)

var mappingListSdksSortByEnum = map[string]ListSdksSortByEnum{
	"timeCreated": ListSdksSortByTimecreated,
	"displayName": ListSdksSortByDisplayname,
}

var mappingListSdksSortByEnumLowerCase = map[string]ListSdksSortByEnum{
	"timecreated": ListSdksSortByTimecreated,
	"displayname": ListSdksSortByDisplayname,
}

// GetListSdksSortByEnumValues Enumerates the set of values for ListSdksSortByEnum
func GetListSdksSortByEnumValues() []ListSdksSortByEnum {
	values := make([]ListSdksSortByEnum, 0)
	for _, v := range mappingListSdksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSdksSortByEnumStringValues Enumerates the set of values in String for ListSdksSortByEnum
func GetListSdksSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSdksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSdksSortByEnum(val string) (ListSdksSortByEnum, bool) {
	enum, ok := mappingListSdksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
