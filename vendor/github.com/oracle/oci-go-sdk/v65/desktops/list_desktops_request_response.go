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

// ListDesktopsRequest wrapper for the ListDesktops operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/desktops/ListDesktops.go.html to see an example of how to use ListDesktopsRequest.
type ListDesktopsRequest struct {

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
	SortBy ListDesktopsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A field to indicate the sort order.
	SortOrder ListDesktopsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The unique identifier of the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID of the desktop pool.
	DesktopPoolId *string `mandatory:"false" contributesTo:"query" name:"desktopPoolId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDesktopsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDesktopsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDesktopsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDesktopsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDesktopsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDesktopsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDesktopsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDesktopsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDesktopsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDesktopsResponse wrapper for the ListDesktops operation
type ListDesktopsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DesktopCollection instances
	DesktopCollection `presentIn:"body"`

	// The unique identifier of the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDesktopsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDesktopsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDesktopsSortByEnum Enum with underlying type: string
type ListDesktopsSortByEnum string

// Set of constants representing the allowable values for ListDesktopsSortByEnum
const (
	ListDesktopsSortByTimecreated ListDesktopsSortByEnum = "TIMECREATED"
	ListDesktopsSortByDisplayname ListDesktopsSortByEnum = "DISPLAYNAME"
)

var mappingListDesktopsSortByEnum = map[string]ListDesktopsSortByEnum{
	"TIMECREATED": ListDesktopsSortByTimecreated,
	"DISPLAYNAME": ListDesktopsSortByDisplayname,
}

var mappingListDesktopsSortByEnumLowerCase = map[string]ListDesktopsSortByEnum{
	"timecreated": ListDesktopsSortByTimecreated,
	"displayname": ListDesktopsSortByDisplayname,
}

// GetListDesktopsSortByEnumValues Enumerates the set of values for ListDesktopsSortByEnum
func GetListDesktopsSortByEnumValues() []ListDesktopsSortByEnum {
	values := make([]ListDesktopsSortByEnum, 0)
	for _, v := range mappingListDesktopsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDesktopsSortByEnumStringValues Enumerates the set of values in String for ListDesktopsSortByEnum
func GetListDesktopsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDesktopsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDesktopsSortByEnum(val string) (ListDesktopsSortByEnum, bool) {
	enum, ok := mappingListDesktopsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDesktopsSortOrderEnum Enum with underlying type: string
type ListDesktopsSortOrderEnum string

// Set of constants representing the allowable values for ListDesktopsSortOrderEnum
const (
	ListDesktopsSortOrderAsc  ListDesktopsSortOrderEnum = "ASC"
	ListDesktopsSortOrderDesc ListDesktopsSortOrderEnum = "DESC"
)

var mappingListDesktopsSortOrderEnum = map[string]ListDesktopsSortOrderEnum{
	"ASC":  ListDesktopsSortOrderAsc,
	"DESC": ListDesktopsSortOrderDesc,
}

var mappingListDesktopsSortOrderEnumLowerCase = map[string]ListDesktopsSortOrderEnum{
	"asc":  ListDesktopsSortOrderAsc,
	"desc": ListDesktopsSortOrderDesc,
}

// GetListDesktopsSortOrderEnumValues Enumerates the set of values for ListDesktopsSortOrderEnum
func GetListDesktopsSortOrderEnumValues() []ListDesktopsSortOrderEnum {
	values := make([]ListDesktopsSortOrderEnum, 0)
	for _, v := range mappingListDesktopsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDesktopsSortOrderEnumStringValues Enumerates the set of values in String for ListDesktopsSortOrderEnum
func GetListDesktopsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDesktopsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDesktopsSortOrderEnum(val string) (ListDesktopsSortOrderEnum, bool) {
	enum, ok := mappingListDesktopsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
