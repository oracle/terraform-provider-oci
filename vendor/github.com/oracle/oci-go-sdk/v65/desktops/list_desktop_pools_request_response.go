// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package desktops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDesktopPoolsRequest wrapper for the ListDesktopPools operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/desktops/ListDesktopPools.go.html to see an example of how to use ListDesktopPoolsRequest.
type ListDesktopPoolsRequest struct {

	// The OCID of the compartment of the desktop pool.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// A filter to return only results with the given displayName.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only results with the given OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only results with the given lifecycleState.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`

	// The maximum number of results to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A field to sort by.
	SortBy ListDesktopPoolsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A field to indicate the sort order.
	SortOrder ListDesktopPoolsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The unique identifier of the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDesktopPoolsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDesktopPoolsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDesktopPoolsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDesktopPoolsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDesktopPoolsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDesktopPoolsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDesktopPoolsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDesktopPoolsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDesktopPoolsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDesktopPoolsResponse wrapper for the ListDesktopPools operation
type ListDesktopPoolsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DesktopPoolCollection instances
	DesktopPoolCollection `presentIn:"body"`

	// The unique identifier of the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDesktopPoolsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDesktopPoolsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDesktopPoolsSortByEnum Enum with underlying type: string
type ListDesktopPoolsSortByEnum string

// Set of constants representing the allowable values for ListDesktopPoolsSortByEnum
const (
	ListDesktopPoolsSortByTimecreated ListDesktopPoolsSortByEnum = "TIMECREATED"
	ListDesktopPoolsSortByDisplayname ListDesktopPoolsSortByEnum = "DISPLAYNAME"
)

var mappingListDesktopPoolsSortByEnum = map[string]ListDesktopPoolsSortByEnum{
	"TIMECREATED": ListDesktopPoolsSortByTimecreated,
	"DISPLAYNAME": ListDesktopPoolsSortByDisplayname,
}

var mappingListDesktopPoolsSortByEnumLowerCase = map[string]ListDesktopPoolsSortByEnum{
	"timecreated": ListDesktopPoolsSortByTimecreated,
	"displayname": ListDesktopPoolsSortByDisplayname,
}

// GetListDesktopPoolsSortByEnumValues Enumerates the set of values for ListDesktopPoolsSortByEnum
func GetListDesktopPoolsSortByEnumValues() []ListDesktopPoolsSortByEnum {
	values := make([]ListDesktopPoolsSortByEnum, 0)
	for _, v := range mappingListDesktopPoolsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDesktopPoolsSortByEnumStringValues Enumerates the set of values in String for ListDesktopPoolsSortByEnum
func GetListDesktopPoolsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDesktopPoolsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDesktopPoolsSortByEnum(val string) (ListDesktopPoolsSortByEnum, bool) {
	enum, ok := mappingListDesktopPoolsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDesktopPoolsSortOrderEnum Enum with underlying type: string
type ListDesktopPoolsSortOrderEnum string

// Set of constants representing the allowable values for ListDesktopPoolsSortOrderEnum
const (
	ListDesktopPoolsSortOrderAsc  ListDesktopPoolsSortOrderEnum = "ASC"
	ListDesktopPoolsSortOrderDesc ListDesktopPoolsSortOrderEnum = "DESC"
)

var mappingListDesktopPoolsSortOrderEnum = map[string]ListDesktopPoolsSortOrderEnum{
	"ASC":  ListDesktopPoolsSortOrderAsc,
	"DESC": ListDesktopPoolsSortOrderDesc,
}

var mappingListDesktopPoolsSortOrderEnumLowerCase = map[string]ListDesktopPoolsSortOrderEnum{
	"asc":  ListDesktopPoolsSortOrderAsc,
	"desc": ListDesktopPoolsSortOrderDesc,
}

// GetListDesktopPoolsSortOrderEnumValues Enumerates the set of values for ListDesktopPoolsSortOrderEnum
func GetListDesktopPoolsSortOrderEnumValues() []ListDesktopPoolsSortOrderEnum {
	values := make([]ListDesktopPoolsSortOrderEnum, 0)
	for _, v := range mappingListDesktopPoolsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDesktopPoolsSortOrderEnumStringValues Enumerates the set of values in String for ListDesktopPoolsSortOrderEnum
func GetListDesktopPoolsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDesktopPoolsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDesktopPoolsSortOrderEnum(val string) (ListDesktopPoolsSortOrderEnum, bool) {
	enum, ok := mappingListDesktopPoolsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
