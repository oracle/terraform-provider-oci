// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementdashboard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOobManagementDashboardsRequest wrapper for the ListOobManagementDashboards operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/ListOobManagementDashboards.go.html to see an example of how to use ListOobManagementDashboardsRequest.
type ListOobManagementDashboardsRequest struct {

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
	SortOrder ListOobManagementDashboardsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is the default.
	SortBy ListOobManagementDashboardsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOobManagementDashboardsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOobManagementDashboardsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOobManagementDashboardsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOobManagementDashboardsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOobManagementDashboardsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOobManagementDashboardsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOobManagementDashboardsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOobManagementDashboardsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOobManagementDashboardsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOobManagementDashboardsResponse wrapper for the ListOobManagementDashboards operation
type ListOobManagementDashboardsResponse struct {

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

func (response ListOobManagementDashboardsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOobManagementDashboardsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOobManagementDashboardsSortOrderEnum Enum with underlying type: string
type ListOobManagementDashboardsSortOrderEnum string

// Set of constants representing the allowable values for ListOobManagementDashboardsSortOrderEnum
const (
	ListOobManagementDashboardsSortOrderAsc  ListOobManagementDashboardsSortOrderEnum = "ASC"
	ListOobManagementDashboardsSortOrderDesc ListOobManagementDashboardsSortOrderEnum = "DESC"
)

var mappingListOobManagementDashboardsSortOrderEnum = map[string]ListOobManagementDashboardsSortOrderEnum{
	"ASC":  ListOobManagementDashboardsSortOrderAsc,
	"DESC": ListOobManagementDashboardsSortOrderDesc,
}

var mappingListOobManagementDashboardsSortOrderEnumLowerCase = map[string]ListOobManagementDashboardsSortOrderEnum{
	"asc":  ListOobManagementDashboardsSortOrderAsc,
	"desc": ListOobManagementDashboardsSortOrderDesc,
}

// GetListOobManagementDashboardsSortOrderEnumValues Enumerates the set of values for ListOobManagementDashboardsSortOrderEnum
func GetListOobManagementDashboardsSortOrderEnumValues() []ListOobManagementDashboardsSortOrderEnum {
	values := make([]ListOobManagementDashboardsSortOrderEnum, 0)
	for _, v := range mappingListOobManagementDashboardsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOobManagementDashboardsSortOrderEnumStringValues Enumerates the set of values in String for ListOobManagementDashboardsSortOrderEnum
func GetListOobManagementDashboardsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOobManagementDashboardsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOobManagementDashboardsSortOrderEnum(val string) (ListOobManagementDashboardsSortOrderEnum, bool) {
	enum, ok := mappingListOobManagementDashboardsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOobManagementDashboardsSortByEnum Enum with underlying type: string
type ListOobManagementDashboardsSortByEnum string

// Set of constants representing the allowable values for ListOobManagementDashboardsSortByEnum
const (
	ListOobManagementDashboardsSortByTimecreated ListOobManagementDashboardsSortByEnum = "timeCreated"
	ListOobManagementDashboardsSortByDisplayname ListOobManagementDashboardsSortByEnum = "displayName"
)

var mappingListOobManagementDashboardsSortByEnum = map[string]ListOobManagementDashboardsSortByEnum{
	"timeCreated": ListOobManagementDashboardsSortByTimecreated,
	"displayName": ListOobManagementDashboardsSortByDisplayname,
}

var mappingListOobManagementDashboardsSortByEnumLowerCase = map[string]ListOobManagementDashboardsSortByEnum{
	"timecreated": ListOobManagementDashboardsSortByTimecreated,
	"displayname": ListOobManagementDashboardsSortByDisplayname,
}

// GetListOobManagementDashboardsSortByEnumValues Enumerates the set of values for ListOobManagementDashboardsSortByEnum
func GetListOobManagementDashboardsSortByEnumValues() []ListOobManagementDashboardsSortByEnum {
	values := make([]ListOobManagementDashboardsSortByEnum, 0)
	for _, v := range mappingListOobManagementDashboardsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOobManagementDashboardsSortByEnumStringValues Enumerates the set of values in String for ListOobManagementDashboardsSortByEnum
func GetListOobManagementDashboardsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOobManagementDashboardsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOobManagementDashboardsSortByEnum(val string) (ListOobManagementDashboardsSortByEnum, bool) {
	enum, ok := mappingListOobManagementDashboardsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
