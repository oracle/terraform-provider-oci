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

// ListExternalDbSystemsRequest wrapper for the ListExternalDbSystems operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalDbSystems.go.html to see an example of how to use ListExternalDbSystemsRequest.
type ListExternalDbSystemsRequest struct {

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
	SortBy ListExternalDbSystemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalDbSystemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalDbSystemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalDbSystemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalDbSystemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalDbSystemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalDbSystemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalDbSystemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalDbSystemsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalDbSystemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalDbSystemsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalDbSystemsResponse wrapper for the ListExternalDbSystems operation
type ListExternalDbSystemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalDbSystemCollection instances
	ExternalDbSystemCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalDbSystemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalDbSystemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalDbSystemsSortByEnum Enum with underlying type: string
type ListExternalDbSystemsSortByEnum string

// Set of constants representing the allowable values for ListExternalDbSystemsSortByEnum
const (
	ListExternalDbSystemsSortByTimecreated ListExternalDbSystemsSortByEnum = "TIMECREATED"
	ListExternalDbSystemsSortByDisplayname ListExternalDbSystemsSortByEnum = "DISPLAYNAME"
)

var mappingListExternalDbSystemsSortByEnum = map[string]ListExternalDbSystemsSortByEnum{
	"TIMECREATED": ListExternalDbSystemsSortByTimecreated,
	"DISPLAYNAME": ListExternalDbSystemsSortByDisplayname,
}

var mappingListExternalDbSystemsSortByEnumLowerCase = map[string]ListExternalDbSystemsSortByEnum{
	"timecreated": ListExternalDbSystemsSortByTimecreated,
	"displayname": ListExternalDbSystemsSortByDisplayname,
}

// GetListExternalDbSystemsSortByEnumValues Enumerates the set of values for ListExternalDbSystemsSortByEnum
func GetListExternalDbSystemsSortByEnumValues() []ListExternalDbSystemsSortByEnum {
	values := make([]ListExternalDbSystemsSortByEnum, 0)
	for _, v := range mappingListExternalDbSystemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalDbSystemsSortByEnumStringValues Enumerates the set of values in String for ListExternalDbSystemsSortByEnum
func GetListExternalDbSystemsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExternalDbSystemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalDbSystemsSortByEnum(val string) (ListExternalDbSystemsSortByEnum, bool) {
	enum, ok := mappingListExternalDbSystemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalDbSystemsSortOrderEnum Enum with underlying type: string
type ListExternalDbSystemsSortOrderEnum string

// Set of constants representing the allowable values for ListExternalDbSystemsSortOrderEnum
const (
	ListExternalDbSystemsSortOrderAsc  ListExternalDbSystemsSortOrderEnum = "ASC"
	ListExternalDbSystemsSortOrderDesc ListExternalDbSystemsSortOrderEnum = "DESC"
)

var mappingListExternalDbSystemsSortOrderEnum = map[string]ListExternalDbSystemsSortOrderEnum{
	"ASC":  ListExternalDbSystemsSortOrderAsc,
	"DESC": ListExternalDbSystemsSortOrderDesc,
}

var mappingListExternalDbSystemsSortOrderEnumLowerCase = map[string]ListExternalDbSystemsSortOrderEnum{
	"asc":  ListExternalDbSystemsSortOrderAsc,
	"desc": ListExternalDbSystemsSortOrderDesc,
}

// GetListExternalDbSystemsSortOrderEnumValues Enumerates the set of values for ListExternalDbSystemsSortOrderEnum
func GetListExternalDbSystemsSortOrderEnumValues() []ListExternalDbSystemsSortOrderEnum {
	values := make([]ListExternalDbSystemsSortOrderEnum, 0)
	for _, v := range mappingListExternalDbSystemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalDbSystemsSortOrderEnumStringValues Enumerates the set of values in String for ListExternalDbSystemsSortOrderEnum
func GetListExternalDbSystemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalDbSystemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalDbSystemsSortOrderEnum(val string) (ListExternalDbSystemsSortOrderEnum, bool) {
	enum, ok := mappingListExternalDbSystemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
