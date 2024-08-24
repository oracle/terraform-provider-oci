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

// ListDesktopPoolDesktopsRequest wrapper for the ListDesktopPoolDesktops operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/desktops/ListDesktopPoolDesktops.go.html to see an example of how to use ListDesktopPoolDesktopsRequest.
type ListDesktopPoolDesktopsRequest struct {

	// The OCID of the compartment of the desktop pool.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID of the desktop pool.
	DesktopPoolId *string `mandatory:"true" contributesTo:"path" name:"desktopPoolId"`

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
	SortBy ListDesktopPoolDesktopsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A field to indicate the sort order.
	SortOrder ListDesktopPoolDesktopsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The unique identifier of the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDesktopPoolDesktopsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDesktopPoolDesktopsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDesktopPoolDesktopsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDesktopPoolDesktopsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDesktopPoolDesktopsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDesktopPoolDesktopsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDesktopPoolDesktopsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDesktopPoolDesktopsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDesktopPoolDesktopsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDesktopPoolDesktopsResponse wrapper for the ListDesktopPoolDesktops operation
type ListDesktopPoolDesktopsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DesktopPoolDesktopCollection instances
	DesktopPoolDesktopCollection `presentIn:"body"`

	// The unique identifier of the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDesktopPoolDesktopsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDesktopPoolDesktopsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDesktopPoolDesktopsSortByEnum Enum with underlying type: string
type ListDesktopPoolDesktopsSortByEnum string

// Set of constants representing the allowable values for ListDesktopPoolDesktopsSortByEnum
const (
	ListDesktopPoolDesktopsSortByTimecreated ListDesktopPoolDesktopsSortByEnum = "TIMECREATED"
	ListDesktopPoolDesktopsSortByDisplayname ListDesktopPoolDesktopsSortByEnum = "DISPLAYNAME"
)

var mappingListDesktopPoolDesktopsSortByEnum = map[string]ListDesktopPoolDesktopsSortByEnum{
	"TIMECREATED": ListDesktopPoolDesktopsSortByTimecreated,
	"DISPLAYNAME": ListDesktopPoolDesktopsSortByDisplayname,
}

var mappingListDesktopPoolDesktopsSortByEnumLowerCase = map[string]ListDesktopPoolDesktopsSortByEnum{
	"timecreated": ListDesktopPoolDesktopsSortByTimecreated,
	"displayname": ListDesktopPoolDesktopsSortByDisplayname,
}

// GetListDesktopPoolDesktopsSortByEnumValues Enumerates the set of values for ListDesktopPoolDesktopsSortByEnum
func GetListDesktopPoolDesktopsSortByEnumValues() []ListDesktopPoolDesktopsSortByEnum {
	values := make([]ListDesktopPoolDesktopsSortByEnum, 0)
	for _, v := range mappingListDesktopPoolDesktopsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDesktopPoolDesktopsSortByEnumStringValues Enumerates the set of values in String for ListDesktopPoolDesktopsSortByEnum
func GetListDesktopPoolDesktopsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDesktopPoolDesktopsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDesktopPoolDesktopsSortByEnum(val string) (ListDesktopPoolDesktopsSortByEnum, bool) {
	enum, ok := mappingListDesktopPoolDesktopsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDesktopPoolDesktopsSortOrderEnum Enum with underlying type: string
type ListDesktopPoolDesktopsSortOrderEnum string

// Set of constants representing the allowable values for ListDesktopPoolDesktopsSortOrderEnum
const (
	ListDesktopPoolDesktopsSortOrderAsc  ListDesktopPoolDesktopsSortOrderEnum = "ASC"
	ListDesktopPoolDesktopsSortOrderDesc ListDesktopPoolDesktopsSortOrderEnum = "DESC"
)

var mappingListDesktopPoolDesktopsSortOrderEnum = map[string]ListDesktopPoolDesktopsSortOrderEnum{
	"ASC":  ListDesktopPoolDesktopsSortOrderAsc,
	"DESC": ListDesktopPoolDesktopsSortOrderDesc,
}

var mappingListDesktopPoolDesktopsSortOrderEnumLowerCase = map[string]ListDesktopPoolDesktopsSortOrderEnum{
	"asc":  ListDesktopPoolDesktopsSortOrderAsc,
	"desc": ListDesktopPoolDesktopsSortOrderDesc,
}

// GetListDesktopPoolDesktopsSortOrderEnumValues Enumerates the set of values for ListDesktopPoolDesktopsSortOrderEnum
func GetListDesktopPoolDesktopsSortOrderEnumValues() []ListDesktopPoolDesktopsSortOrderEnum {
	values := make([]ListDesktopPoolDesktopsSortOrderEnum, 0)
	for _, v := range mappingListDesktopPoolDesktopsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDesktopPoolDesktopsSortOrderEnumStringValues Enumerates the set of values in String for ListDesktopPoolDesktopsSortOrderEnum
func GetListDesktopPoolDesktopsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDesktopPoolDesktopsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDesktopPoolDesktopsSortOrderEnum(val string) (ListDesktopPoolDesktopsSortOrderEnum, bool) {
	enum, ok := mappingListDesktopPoolDesktopsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
