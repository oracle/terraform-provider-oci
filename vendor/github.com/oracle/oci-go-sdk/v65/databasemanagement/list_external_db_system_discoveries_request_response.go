// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExternalDbSystemDiscoveriesRequest wrapper for the ListExternalDbSystemDiscoveries operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalDbSystemDiscoveries.go.html to see an example of how to use ListExternalDbSystemDiscoveriesRequest.
type ListExternalDbSystemDiscoveriesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
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
	SortBy ListExternalDbSystemDiscoveriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalDbSystemDiscoveriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalDbSystemDiscoveriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalDbSystemDiscoveriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalDbSystemDiscoveriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalDbSystemDiscoveriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalDbSystemDiscoveriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalDbSystemDiscoveriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalDbSystemDiscoveriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalDbSystemDiscoveriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalDbSystemDiscoveriesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalDbSystemDiscoveriesResponse wrapper for the ListExternalDbSystemDiscoveries operation
type ListExternalDbSystemDiscoveriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalDbSystemDiscoveryCollection instances
	ExternalDbSystemDiscoveryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalDbSystemDiscoveriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalDbSystemDiscoveriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalDbSystemDiscoveriesSortByEnum Enum with underlying type: string
type ListExternalDbSystemDiscoveriesSortByEnum string

// Set of constants representing the allowable values for ListExternalDbSystemDiscoveriesSortByEnum
const (
	ListExternalDbSystemDiscoveriesSortByTimecreated ListExternalDbSystemDiscoveriesSortByEnum = "TIMECREATED"
	ListExternalDbSystemDiscoveriesSortByDisplayname ListExternalDbSystemDiscoveriesSortByEnum = "DISPLAYNAME"
)

var mappingListExternalDbSystemDiscoveriesSortByEnum = map[string]ListExternalDbSystemDiscoveriesSortByEnum{
	"TIMECREATED": ListExternalDbSystemDiscoveriesSortByTimecreated,
	"DISPLAYNAME": ListExternalDbSystemDiscoveriesSortByDisplayname,
}

var mappingListExternalDbSystemDiscoveriesSortByEnumLowerCase = map[string]ListExternalDbSystemDiscoveriesSortByEnum{
	"timecreated": ListExternalDbSystemDiscoveriesSortByTimecreated,
	"displayname": ListExternalDbSystemDiscoveriesSortByDisplayname,
}

// GetListExternalDbSystemDiscoveriesSortByEnumValues Enumerates the set of values for ListExternalDbSystemDiscoveriesSortByEnum
func GetListExternalDbSystemDiscoveriesSortByEnumValues() []ListExternalDbSystemDiscoveriesSortByEnum {
	values := make([]ListExternalDbSystemDiscoveriesSortByEnum, 0)
	for _, v := range mappingListExternalDbSystemDiscoveriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalDbSystemDiscoveriesSortByEnumStringValues Enumerates the set of values in String for ListExternalDbSystemDiscoveriesSortByEnum
func GetListExternalDbSystemDiscoveriesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExternalDbSystemDiscoveriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalDbSystemDiscoveriesSortByEnum(val string) (ListExternalDbSystemDiscoveriesSortByEnum, bool) {
	enum, ok := mappingListExternalDbSystemDiscoveriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalDbSystemDiscoveriesSortOrderEnum Enum with underlying type: string
type ListExternalDbSystemDiscoveriesSortOrderEnum string

// Set of constants representing the allowable values for ListExternalDbSystemDiscoveriesSortOrderEnum
const (
	ListExternalDbSystemDiscoveriesSortOrderAsc  ListExternalDbSystemDiscoveriesSortOrderEnum = "ASC"
	ListExternalDbSystemDiscoveriesSortOrderDesc ListExternalDbSystemDiscoveriesSortOrderEnum = "DESC"
)

var mappingListExternalDbSystemDiscoveriesSortOrderEnum = map[string]ListExternalDbSystemDiscoveriesSortOrderEnum{
	"ASC":  ListExternalDbSystemDiscoveriesSortOrderAsc,
	"DESC": ListExternalDbSystemDiscoveriesSortOrderDesc,
}

var mappingListExternalDbSystemDiscoveriesSortOrderEnumLowerCase = map[string]ListExternalDbSystemDiscoveriesSortOrderEnum{
	"asc":  ListExternalDbSystemDiscoveriesSortOrderAsc,
	"desc": ListExternalDbSystemDiscoveriesSortOrderDesc,
}

// GetListExternalDbSystemDiscoveriesSortOrderEnumValues Enumerates the set of values for ListExternalDbSystemDiscoveriesSortOrderEnum
func GetListExternalDbSystemDiscoveriesSortOrderEnumValues() []ListExternalDbSystemDiscoveriesSortOrderEnum {
	values := make([]ListExternalDbSystemDiscoveriesSortOrderEnum, 0)
	for _, v := range mappingListExternalDbSystemDiscoveriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalDbSystemDiscoveriesSortOrderEnumStringValues Enumerates the set of values in String for ListExternalDbSystemDiscoveriesSortOrderEnum
func GetListExternalDbSystemDiscoveriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalDbSystemDiscoveriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalDbSystemDiscoveriesSortOrderEnum(val string) (ListExternalDbSystemDiscoveriesSortOrderEnum, bool) {
	enum, ok := mappingListExternalDbSystemDiscoveriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
