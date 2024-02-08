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

// ListExternalDatabasesRequest wrapper for the ListExternalDatabases operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalDatabases.go.html to see an example of how to use ListExternalDatabasesRequest.
type ListExternalDatabasesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system.
	ExternalDbSystemId *string `mandatory:"false" contributesTo:"query" name:"externalDbSystemId"`

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
	SortBy ListExternalDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalDatabasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalDatabasesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalDatabasesResponse wrapper for the ListExternalDatabases operation
type ListExternalDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalDatabaseCollection instances
	ExternalDatabaseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalDatabasesSortByEnum Enum with underlying type: string
type ListExternalDatabasesSortByEnum string

// Set of constants representing the allowable values for ListExternalDatabasesSortByEnum
const (
	ListExternalDatabasesSortByTimecreated ListExternalDatabasesSortByEnum = "TIMECREATED"
	ListExternalDatabasesSortByDisplayname ListExternalDatabasesSortByEnum = "DISPLAYNAME"
)

var mappingListExternalDatabasesSortByEnum = map[string]ListExternalDatabasesSortByEnum{
	"TIMECREATED": ListExternalDatabasesSortByTimecreated,
	"DISPLAYNAME": ListExternalDatabasesSortByDisplayname,
}

var mappingListExternalDatabasesSortByEnumLowerCase = map[string]ListExternalDatabasesSortByEnum{
	"timecreated": ListExternalDatabasesSortByTimecreated,
	"displayname": ListExternalDatabasesSortByDisplayname,
}

// GetListExternalDatabasesSortByEnumValues Enumerates the set of values for ListExternalDatabasesSortByEnum
func GetListExternalDatabasesSortByEnumValues() []ListExternalDatabasesSortByEnum {
	values := make([]ListExternalDatabasesSortByEnum, 0)
	for _, v := range mappingListExternalDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalDatabasesSortByEnumStringValues Enumerates the set of values in String for ListExternalDatabasesSortByEnum
func GetListExternalDatabasesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExternalDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalDatabasesSortByEnum(val string) (ListExternalDatabasesSortByEnum, bool) {
	enum, ok := mappingListExternalDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalDatabasesSortOrderEnum Enum with underlying type: string
type ListExternalDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListExternalDatabasesSortOrderEnum
const (
	ListExternalDatabasesSortOrderAsc  ListExternalDatabasesSortOrderEnum = "ASC"
	ListExternalDatabasesSortOrderDesc ListExternalDatabasesSortOrderEnum = "DESC"
)

var mappingListExternalDatabasesSortOrderEnum = map[string]ListExternalDatabasesSortOrderEnum{
	"ASC":  ListExternalDatabasesSortOrderAsc,
	"DESC": ListExternalDatabasesSortOrderDesc,
}

var mappingListExternalDatabasesSortOrderEnumLowerCase = map[string]ListExternalDatabasesSortOrderEnum{
	"asc":  ListExternalDatabasesSortOrderAsc,
	"desc": ListExternalDatabasesSortOrderDesc,
}

// GetListExternalDatabasesSortOrderEnumValues Enumerates the set of values for ListExternalDatabasesSortOrderEnum
func GetListExternalDatabasesSortOrderEnumValues() []ListExternalDatabasesSortOrderEnum {
	values := make([]ListExternalDatabasesSortOrderEnum, 0)
	for _, v := range mappingListExternalDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListExternalDatabasesSortOrderEnum
func GetListExternalDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalDatabasesSortOrderEnum(val string) (ListExternalDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListExternalDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
