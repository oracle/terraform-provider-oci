// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListApplicationVipsRequest wrapper for the ListApplicationVips operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListApplicationVips.go.html to see an example of how to use ListApplicationVipsRequest.
type ListApplicationVipsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud VM cluster associated with the application virtual IP (VIP) address.
	CloudVmClusterId *string `mandatory:"true" contributesTo:"query" name:"cloudVmClusterId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListApplicationVipsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for TIMECREATED is descending.
	// Default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListApplicationVipsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState ApplicationVipSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApplicationVipsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApplicationVipsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApplicationVipsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApplicationVipsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListApplicationVipsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListApplicationVipsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListApplicationVipsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApplicationVipsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListApplicationVipsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingApplicationVipSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetApplicationVipSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListApplicationVipsResponse wrapper for the ListApplicationVips operation
type ListApplicationVipsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ApplicationVipSummary instances
	Items []ApplicationVipSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListApplicationVipsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApplicationVipsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApplicationVipsSortOrderEnum Enum with underlying type: string
type ListApplicationVipsSortOrderEnum string

// Set of constants representing the allowable values for ListApplicationVipsSortOrderEnum
const (
	ListApplicationVipsSortOrderAsc  ListApplicationVipsSortOrderEnum = "ASC"
	ListApplicationVipsSortOrderDesc ListApplicationVipsSortOrderEnum = "DESC"
)

var mappingListApplicationVipsSortOrderEnum = map[string]ListApplicationVipsSortOrderEnum{
	"ASC":  ListApplicationVipsSortOrderAsc,
	"DESC": ListApplicationVipsSortOrderDesc,
}

var mappingListApplicationVipsSortOrderEnumLowerCase = map[string]ListApplicationVipsSortOrderEnum{
	"asc":  ListApplicationVipsSortOrderAsc,
	"desc": ListApplicationVipsSortOrderDesc,
}

// GetListApplicationVipsSortOrderEnumValues Enumerates the set of values for ListApplicationVipsSortOrderEnum
func GetListApplicationVipsSortOrderEnumValues() []ListApplicationVipsSortOrderEnum {
	values := make([]ListApplicationVipsSortOrderEnum, 0)
	for _, v := range mappingListApplicationVipsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicationVipsSortOrderEnumStringValues Enumerates the set of values in String for ListApplicationVipsSortOrderEnum
func GetListApplicationVipsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListApplicationVipsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicationVipsSortOrderEnum(val string) (ListApplicationVipsSortOrderEnum, bool) {
	enum, ok := mappingListApplicationVipsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListApplicationVipsSortByEnum Enum with underlying type: string
type ListApplicationVipsSortByEnum string

// Set of constants representing the allowable values for ListApplicationVipsSortByEnum
const (
	ListApplicationVipsSortByDisplayname ListApplicationVipsSortByEnum = "DISPLAYNAME"
	ListApplicationVipsSortByTimecreated ListApplicationVipsSortByEnum = "TIMECREATED"
)

var mappingListApplicationVipsSortByEnum = map[string]ListApplicationVipsSortByEnum{
	"DISPLAYNAME": ListApplicationVipsSortByDisplayname,
	"TIMECREATED": ListApplicationVipsSortByTimecreated,
}

var mappingListApplicationVipsSortByEnumLowerCase = map[string]ListApplicationVipsSortByEnum{
	"displayname": ListApplicationVipsSortByDisplayname,
	"timecreated": ListApplicationVipsSortByTimecreated,
}

// GetListApplicationVipsSortByEnumValues Enumerates the set of values for ListApplicationVipsSortByEnum
func GetListApplicationVipsSortByEnumValues() []ListApplicationVipsSortByEnum {
	values := make([]ListApplicationVipsSortByEnum, 0)
	for _, v := range mappingListApplicationVipsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicationVipsSortByEnumStringValues Enumerates the set of values in String for ListApplicationVipsSortByEnum
func GetListApplicationVipsSortByEnumStringValues() []string {
	return []string{
		"DISPLAYNAME",
		"TIMECREATED",
	}
}

// GetMappingListApplicationVipsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicationVipsSortByEnum(val string) (ListApplicationVipsSortByEnum, bool) {
	enum, ok := mappingListApplicationVipsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
