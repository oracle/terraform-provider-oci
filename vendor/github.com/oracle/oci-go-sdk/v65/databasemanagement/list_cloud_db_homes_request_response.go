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

// ListCloudDbHomesRequest wrapper for the ListCloudDbHomes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudDbHomes.go.html to see an example of how to use ListCloudDbHomesRequest.
type ListCloudDbHomesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
	CloudDbSystemId *string `mandatory:"false" contributesTo:"query" name:"cloudDbSystemId"`

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
	SortBy ListCloudDbHomesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudDbHomesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudDbHomesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudDbHomesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudDbHomesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudDbHomesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudDbHomesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudDbHomesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudDbHomesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudDbHomesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudDbHomesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudDbHomesResponse wrapper for the ListCloudDbHomes operation
type ListCloudDbHomesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudDbHomeCollection instances
	CloudDbHomeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudDbHomesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudDbHomesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudDbHomesSortByEnum Enum with underlying type: string
type ListCloudDbHomesSortByEnum string

// Set of constants representing the allowable values for ListCloudDbHomesSortByEnum
const (
	ListCloudDbHomesSortByTimecreated ListCloudDbHomesSortByEnum = "TIMECREATED"
	ListCloudDbHomesSortByDisplayname ListCloudDbHomesSortByEnum = "DISPLAYNAME"
)

var mappingListCloudDbHomesSortByEnum = map[string]ListCloudDbHomesSortByEnum{
	"TIMECREATED": ListCloudDbHomesSortByTimecreated,
	"DISPLAYNAME": ListCloudDbHomesSortByDisplayname,
}

var mappingListCloudDbHomesSortByEnumLowerCase = map[string]ListCloudDbHomesSortByEnum{
	"timecreated": ListCloudDbHomesSortByTimecreated,
	"displayname": ListCloudDbHomesSortByDisplayname,
}

// GetListCloudDbHomesSortByEnumValues Enumerates the set of values for ListCloudDbHomesSortByEnum
func GetListCloudDbHomesSortByEnumValues() []ListCloudDbHomesSortByEnum {
	values := make([]ListCloudDbHomesSortByEnum, 0)
	for _, v := range mappingListCloudDbHomesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDbHomesSortByEnumStringValues Enumerates the set of values in String for ListCloudDbHomesSortByEnum
func GetListCloudDbHomesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudDbHomesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDbHomesSortByEnum(val string) (ListCloudDbHomesSortByEnum, bool) {
	enum, ok := mappingListCloudDbHomesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudDbHomesSortOrderEnum Enum with underlying type: string
type ListCloudDbHomesSortOrderEnum string

// Set of constants representing the allowable values for ListCloudDbHomesSortOrderEnum
const (
	ListCloudDbHomesSortOrderAsc  ListCloudDbHomesSortOrderEnum = "ASC"
	ListCloudDbHomesSortOrderDesc ListCloudDbHomesSortOrderEnum = "DESC"
)

var mappingListCloudDbHomesSortOrderEnum = map[string]ListCloudDbHomesSortOrderEnum{
	"ASC":  ListCloudDbHomesSortOrderAsc,
	"DESC": ListCloudDbHomesSortOrderDesc,
}

var mappingListCloudDbHomesSortOrderEnumLowerCase = map[string]ListCloudDbHomesSortOrderEnum{
	"asc":  ListCloudDbHomesSortOrderAsc,
	"desc": ListCloudDbHomesSortOrderDesc,
}

// GetListCloudDbHomesSortOrderEnumValues Enumerates the set of values for ListCloudDbHomesSortOrderEnum
func GetListCloudDbHomesSortOrderEnumValues() []ListCloudDbHomesSortOrderEnum {
	values := make([]ListCloudDbHomesSortOrderEnum, 0)
	for _, v := range mappingListCloudDbHomesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDbHomesSortOrderEnumStringValues Enumerates the set of values in String for ListCloudDbHomesSortOrderEnum
func GetListCloudDbHomesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudDbHomesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDbHomesSortOrderEnum(val string) (ListCloudDbHomesSortOrderEnum, bool) {
	enum, ok := mappingListCloudDbHomesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
