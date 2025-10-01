// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourceanalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListResourceAnalyticsInstancesRequest wrapper for the ListResourceAnalyticsInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/ListResourceAnalyticsInstances.go.html to see an example of how to use ListResourceAnalyticsInstancesRequest.
type ListResourceAnalyticsInstancesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState ResourceAnalyticsInstanceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListResourceAnalyticsInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `TIME_CREATED`
	// is descending. Default order for `DISPLAY_NAME` is ascending.
	SortBy ListResourceAnalyticsInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceAnalyticsInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceAnalyticsInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceAnalyticsInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceAnalyticsInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourceAnalyticsInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourceAnalyticsInstanceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetResourceAnalyticsInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceAnalyticsInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourceAnalyticsInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceAnalyticsInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourceAnalyticsInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourceAnalyticsInstancesResponse wrapper for the ListResourceAnalyticsInstances operation
type ListResourceAnalyticsInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceAnalyticsInstanceCollection instances
	ResourceAnalyticsInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResourceAnalyticsInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceAnalyticsInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceAnalyticsInstancesSortOrderEnum Enum with underlying type: string
type ListResourceAnalyticsInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListResourceAnalyticsInstancesSortOrderEnum
const (
	ListResourceAnalyticsInstancesSortOrderAsc  ListResourceAnalyticsInstancesSortOrderEnum = "ASC"
	ListResourceAnalyticsInstancesSortOrderDesc ListResourceAnalyticsInstancesSortOrderEnum = "DESC"
)

var mappingListResourceAnalyticsInstancesSortOrderEnum = map[string]ListResourceAnalyticsInstancesSortOrderEnum{
	"ASC":  ListResourceAnalyticsInstancesSortOrderAsc,
	"DESC": ListResourceAnalyticsInstancesSortOrderDesc,
}

var mappingListResourceAnalyticsInstancesSortOrderEnumLowerCase = map[string]ListResourceAnalyticsInstancesSortOrderEnum{
	"asc":  ListResourceAnalyticsInstancesSortOrderAsc,
	"desc": ListResourceAnalyticsInstancesSortOrderDesc,
}

// GetListResourceAnalyticsInstancesSortOrderEnumValues Enumerates the set of values for ListResourceAnalyticsInstancesSortOrderEnum
func GetListResourceAnalyticsInstancesSortOrderEnumValues() []ListResourceAnalyticsInstancesSortOrderEnum {
	values := make([]ListResourceAnalyticsInstancesSortOrderEnum, 0)
	for _, v := range mappingListResourceAnalyticsInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceAnalyticsInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListResourceAnalyticsInstancesSortOrderEnum
func GetListResourceAnalyticsInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourceAnalyticsInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceAnalyticsInstancesSortOrderEnum(val string) (ListResourceAnalyticsInstancesSortOrderEnum, bool) {
	enum, ok := mappingListResourceAnalyticsInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceAnalyticsInstancesSortByEnum Enum with underlying type: string
type ListResourceAnalyticsInstancesSortByEnum string

// Set of constants representing the allowable values for ListResourceAnalyticsInstancesSortByEnum
const (
	ListResourceAnalyticsInstancesSortByTimeCreated ListResourceAnalyticsInstancesSortByEnum = "TIME_CREATED"
	ListResourceAnalyticsInstancesSortByDisplayName ListResourceAnalyticsInstancesSortByEnum = "DISPLAY_NAME"
)

var mappingListResourceAnalyticsInstancesSortByEnum = map[string]ListResourceAnalyticsInstancesSortByEnum{
	"TIME_CREATED": ListResourceAnalyticsInstancesSortByTimeCreated,
	"DISPLAY_NAME": ListResourceAnalyticsInstancesSortByDisplayName,
}

var mappingListResourceAnalyticsInstancesSortByEnumLowerCase = map[string]ListResourceAnalyticsInstancesSortByEnum{
	"time_created": ListResourceAnalyticsInstancesSortByTimeCreated,
	"display_name": ListResourceAnalyticsInstancesSortByDisplayName,
}

// GetListResourceAnalyticsInstancesSortByEnumValues Enumerates the set of values for ListResourceAnalyticsInstancesSortByEnum
func GetListResourceAnalyticsInstancesSortByEnumValues() []ListResourceAnalyticsInstancesSortByEnum {
	values := make([]ListResourceAnalyticsInstancesSortByEnum, 0)
	for _, v := range mappingListResourceAnalyticsInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceAnalyticsInstancesSortByEnumStringValues Enumerates the set of values in String for ListResourceAnalyticsInstancesSortByEnum
func GetListResourceAnalyticsInstancesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListResourceAnalyticsInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceAnalyticsInstancesSortByEnum(val string) (ListResourceAnalyticsInstancesSortByEnum, bool) {
	enum, ok := mappingListResourceAnalyticsInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
