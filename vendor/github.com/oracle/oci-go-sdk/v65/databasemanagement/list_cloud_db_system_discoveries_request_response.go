// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCloudDbSystemDiscoveriesRequest wrapper for the ListCloudDbSystemDiscoveries operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudDbSystemDiscoveries.go.html to see an example of how to use ListCloudDbSystemDiscoveriesRequest.
type ListCloudDbSystemDiscoveriesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to only return the resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for `TIMECREATED` is descending and the default sort order for `DISPLAYNAME` is ascending.
	// The `DISPLAYNAME` sort order is case-sensitive.
	SortBy ListCloudDbSystemDiscoveriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudDbSystemDiscoveriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudDbSystemDiscoveriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudDbSystemDiscoveriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudDbSystemDiscoveriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudDbSystemDiscoveriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudDbSystemDiscoveriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudDbSystemDiscoveriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudDbSystemDiscoveriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudDbSystemDiscoveriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudDbSystemDiscoveriesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudDbSystemDiscoveriesResponse wrapper for the ListCloudDbSystemDiscoveries operation
type ListCloudDbSystemDiscoveriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudDbSystemDiscoveryCollection instances
	CloudDbSystemDiscoveryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudDbSystemDiscoveriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudDbSystemDiscoveriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudDbSystemDiscoveriesSortByEnum Enum with underlying type: string
type ListCloudDbSystemDiscoveriesSortByEnum string

// Set of constants representing the allowable values for ListCloudDbSystemDiscoveriesSortByEnum
const (
	ListCloudDbSystemDiscoveriesSortByTimecreated ListCloudDbSystemDiscoveriesSortByEnum = "TIMECREATED"
	ListCloudDbSystemDiscoveriesSortByDisplayname ListCloudDbSystemDiscoveriesSortByEnum = "DISPLAYNAME"
)

var mappingListCloudDbSystemDiscoveriesSortByEnum = map[string]ListCloudDbSystemDiscoveriesSortByEnum{
	"TIMECREATED": ListCloudDbSystemDiscoveriesSortByTimecreated,
	"DISPLAYNAME": ListCloudDbSystemDiscoveriesSortByDisplayname,
}

var mappingListCloudDbSystemDiscoveriesSortByEnumLowerCase = map[string]ListCloudDbSystemDiscoveriesSortByEnum{
	"timecreated": ListCloudDbSystemDiscoveriesSortByTimecreated,
	"displayname": ListCloudDbSystemDiscoveriesSortByDisplayname,
}

// GetListCloudDbSystemDiscoveriesSortByEnumValues Enumerates the set of values for ListCloudDbSystemDiscoveriesSortByEnum
func GetListCloudDbSystemDiscoveriesSortByEnumValues() []ListCloudDbSystemDiscoveriesSortByEnum {
	values := make([]ListCloudDbSystemDiscoveriesSortByEnum, 0)
	for _, v := range mappingListCloudDbSystemDiscoveriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDbSystemDiscoveriesSortByEnumStringValues Enumerates the set of values in String for ListCloudDbSystemDiscoveriesSortByEnum
func GetListCloudDbSystemDiscoveriesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudDbSystemDiscoveriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDbSystemDiscoveriesSortByEnum(val string) (ListCloudDbSystemDiscoveriesSortByEnum, bool) {
	enum, ok := mappingListCloudDbSystemDiscoveriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudDbSystemDiscoveriesSortOrderEnum Enum with underlying type: string
type ListCloudDbSystemDiscoveriesSortOrderEnum string

// Set of constants representing the allowable values for ListCloudDbSystemDiscoveriesSortOrderEnum
const (
	ListCloudDbSystemDiscoveriesSortOrderAsc  ListCloudDbSystemDiscoveriesSortOrderEnum = "ASC"
	ListCloudDbSystemDiscoveriesSortOrderDesc ListCloudDbSystemDiscoveriesSortOrderEnum = "DESC"
)

var mappingListCloudDbSystemDiscoveriesSortOrderEnum = map[string]ListCloudDbSystemDiscoveriesSortOrderEnum{
	"ASC":  ListCloudDbSystemDiscoveriesSortOrderAsc,
	"DESC": ListCloudDbSystemDiscoveriesSortOrderDesc,
}

var mappingListCloudDbSystemDiscoveriesSortOrderEnumLowerCase = map[string]ListCloudDbSystemDiscoveriesSortOrderEnum{
	"asc":  ListCloudDbSystemDiscoveriesSortOrderAsc,
	"desc": ListCloudDbSystemDiscoveriesSortOrderDesc,
}

// GetListCloudDbSystemDiscoveriesSortOrderEnumValues Enumerates the set of values for ListCloudDbSystemDiscoveriesSortOrderEnum
func GetListCloudDbSystemDiscoveriesSortOrderEnumValues() []ListCloudDbSystemDiscoveriesSortOrderEnum {
	values := make([]ListCloudDbSystemDiscoveriesSortOrderEnum, 0)
	for _, v := range mappingListCloudDbSystemDiscoveriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDbSystemDiscoveriesSortOrderEnumStringValues Enumerates the set of values in String for ListCloudDbSystemDiscoveriesSortOrderEnum
func GetListCloudDbSystemDiscoveriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudDbSystemDiscoveriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDbSystemDiscoveriesSortOrderEnum(val string) (ListCloudDbSystemDiscoveriesSortOrderEnum, bool) {
	enum, ok := mappingListCloudDbSystemDiscoveriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
