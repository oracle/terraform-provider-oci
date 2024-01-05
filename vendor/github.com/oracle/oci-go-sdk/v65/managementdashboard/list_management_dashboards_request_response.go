// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementdashboard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagementDashboardsRequest wrapper for the ListManagementDashboards operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/ListManagementDashboards.go.html to see an example of how to use ListManagementDashboardsRequest.
type ListManagementDashboardsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page on which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagementDashboardsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is the default.
	SortBy ListManagementDashboardsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementDashboardsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementDashboardsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagementDashboardsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementDashboardsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagementDashboardsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagementDashboardsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagementDashboardsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementDashboardsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagementDashboardsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagementDashboardsResponse wrapper for the ListManagementDashboards operation
type ListManagementDashboardsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagementDashboardCollection instances
	ManagementDashboardCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagementDashboardsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementDashboardsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementDashboardsSortOrderEnum Enum with underlying type: string
type ListManagementDashboardsSortOrderEnum string

// Set of constants representing the allowable values for ListManagementDashboardsSortOrderEnum
const (
	ListManagementDashboardsSortOrderAsc  ListManagementDashboardsSortOrderEnum = "ASC"
	ListManagementDashboardsSortOrderDesc ListManagementDashboardsSortOrderEnum = "DESC"
)

var mappingListManagementDashboardsSortOrderEnum = map[string]ListManagementDashboardsSortOrderEnum{
	"ASC":  ListManagementDashboardsSortOrderAsc,
	"DESC": ListManagementDashboardsSortOrderDesc,
}

var mappingListManagementDashboardsSortOrderEnumLowerCase = map[string]ListManagementDashboardsSortOrderEnum{
	"asc":  ListManagementDashboardsSortOrderAsc,
	"desc": ListManagementDashboardsSortOrderDesc,
}

// GetListManagementDashboardsSortOrderEnumValues Enumerates the set of values for ListManagementDashboardsSortOrderEnum
func GetListManagementDashboardsSortOrderEnumValues() []ListManagementDashboardsSortOrderEnum {
	values := make([]ListManagementDashboardsSortOrderEnum, 0)
	for _, v := range mappingListManagementDashboardsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementDashboardsSortOrderEnumStringValues Enumerates the set of values in String for ListManagementDashboardsSortOrderEnum
func GetListManagementDashboardsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagementDashboardsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementDashboardsSortOrderEnum(val string) (ListManagementDashboardsSortOrderEnum, bool) {
	enum, ok := mappingListManagementDashboardsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementDashboardsSortByEnum Enum with underlying type: string
type ListManagementDashboardsSortByEnum string

// Set of constants representing the allowable values for ListManagementDashboardsSortByEnum
const (
	ListManagementDashboardsSortByTimecreated ListManagementDashboardsSortByEnum = "timeCreated"
	ListManagementDashboardsSortByDisplayname ListManagementDashboardsSortByEnum = "displayName"
)

var mappingListManagementDashboardsSortByEnum = map[string]ListManagementDashboardsSortByEnum{
	"timeCreated": ListManagementDashboardsSortByTimecreated,
	"displayName": ListManagementDashboardsSortByDisplayname,
}

var mappingListManagementDashboardsSortByEnumLowerCase = map[string]ListManagementDashboardsSortByEnum{
	"timecreated": ListManagementDashboardsSortByTimecreated,
	"displayname": ListManagementDashboardsSortByDisplayname,
}

// GetListManagementDashboardsSortByEnumValues Enumerates the set of values for ListManagementDashboardsSortByEnum
func GetListManagementDashboardsSortByEnumValues() []ListManagementDashboardsSortByEnum {
	values := make([]ListManagementDashboardsSortByEnum, 0)
	for _, v := range mappingListManagementDashboardsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementDashboardsSortByEnumStringValues Enumerates the set of values in String for ListManagementDashboardsSortByEnum
func GetListManagementDashboardsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagementDashboardsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementDashboardsSortByEnum(val string) (ListManagementDashboardsSortByEnum, bool) {
	enum, ok := mappingListManagementDashboardsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
