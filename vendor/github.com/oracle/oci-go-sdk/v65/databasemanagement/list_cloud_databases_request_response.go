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

// ListCloudDatabasesRequest wrapper for the ListCloudDatabases operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudDatabases.go.html to see an example of how to use ListCloudDatabasesRequest.
type ListCloudDatabasesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
	CloudDbSystemId *string `mandatory:"true" contributesTo:"query" name:"cloudDbSystemId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

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
	SortBy ListCloudDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudDatabasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudDatabasesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudDatabasesResponse wrapper for the ListCloudDatabases operation
type ListCloudDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudDatabaseCollection instances
	CloudDatabaseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudDatabasesSortByEnum Enum with underlying type: string
type ListCloudDatabasesSortByEnum string

// Set of constants representing the allowable values for ListCloudDatabasesSortByEnum
const (
	ListCloudDatabasesSortByTimecreated ListCloudDatabasesSortByEnum = "TIMECREATED"
	ListCloudDatabasesSortByDisplayname ListCloudDatabasesSortByEnum = "DISPLAYNAME"
)

var mappingListCloudDatabasesSortByEnum = map[string]ListCloudDatabasesSortByEnum{
	"TIMECREATED": ListCloudDatabasesSortByTimecreated,
	"DISPLAYNAME": ListCloudDatabasesSortByDisplayname,
}

var mappingListCloudDatabasesSortByEnumLowerCase = map[string]ListCloudDatabasesSortByEnum{
	"timecreated": ListCloudDatabasesSortByTimecreated,
	"displayname": ListCloudDatabasesSortByDisplayname,
}

// GetListCloudDatabasesSortByEnumValues Enumerates the set of values for ListCloudDatabasesSortByEnum
func GetListCloudDatabasesSortByEnumValues() []ListCloudDatabasesSortByEnum {
	values := make([]ListCloudDatabasesSortByEnum, 0)
	for _, v := range mappingListCloudDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDatabasesSortByEnumStringValues Enumerates the set of values in String for ListCloudDatabasesSortByEnum
func GetListCloudDatabasesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDatabasesSortByEnum(val string) (ListCloudDatabasesSortByEnum, bool) {
	enum, ok := mappingListCloudDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudDatabasesSortOrderEnum Enum with underlying type: string
type ListCloudDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListCloudDatabasesSortOrderEnum
const (
	ListCloudDatabasesSortOrderAsc  ListCloudDatabasesSortOrderEnum = "ASC"
	ListCloudDatabasesSortOrderDesc ListCloudDatabasesSortOrderEnum = "DESC"
)

var mappingListCloudDatabasesSortOrderEnum = map[string]ListCloudDatabasesSortOrderEnum{
	"ASC":  ListCloudDatabasesSortOrderAsc,
	"DESC": ListCloudDatabasesSortOrderDesc,
}

var mappingListCloudDatabasesSortOrderEnumLowerCase = map[string]ListCloudDatabasesSortOrderEnum{
	"asc":  ListCloudDatabasesSortOrderAsc,
	"desc": ListCloudDatabasesSortOrderDesc,
}

// GetListCloudDatabasesSortOrderEnumValues Enumerates the set of values for ListCloudDatabasesSortOrderEnum
func GetListCloudDatabasesSortOrderEnumValues() []ListCloudDatabasesSortOrderEnum {
	values := make([]ListCloudDatabasesSortOrderEnum, 0)
	for _, v := range mappingListCloudDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListCloudDatabasesSortOrderEnum
func GetListCloudDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDatabasesSortOrderEnum(val string) (ListCloudDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListCloudDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
