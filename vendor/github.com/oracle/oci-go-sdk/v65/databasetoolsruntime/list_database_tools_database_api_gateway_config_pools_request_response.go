// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDatabaseToolsDatabaseApiGatewayConfigPoolsRequest wrapper for the ListDatabaseToolsDatabaseApiGatewayConfigPools operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListDatabaseToolsDatabaseApiGatewayConfigPools.go.html to see an example of how to use ListDatabaseToolsDatabaseApiGatewayConfigPoolsRequest.
type ListDatabaseToolsDatabaseApiGatewayConfigPoolsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
	DatabaseToolsDatabaseApiGatewayConfigId *string `mandatory:"true" contributesTo:"path" name:"databaseToolsDatabaseApiGatewayConfigId"`

	// A filter to return only resources that match the entire specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseToolsDatabaseApiGatewayConfigPoolsResponse wrapper for the ListDatabaseToolsDatabaseApiGatewayConfigPools operation
type ListDatabaseToolsDatabaseApiGatewayConfigPoolsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseToolsDatabaseApiGatewayConfigPoolCollection instances
	DatabaseToolsDatabaseApiGatewayConfigPoolCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseToolsDatabaseApiGatewayConfigPoolsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseToolsDatabaseApiGatewayConfigPoolsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum Enum with underlying type: string
type ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum
const (
	ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderAsc  ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum = "ASC"
	ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderDesc ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum = "DESC"
)

var mappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum = map[string]ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum{
	"ASC":  ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderAsc,
	"DESC": ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderDesc,
}

var mappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnumLowerCase = map[string]ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum{
	"asc":  ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderAsc,
	"desc": ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderDesc,
}

// GetListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnumValues Enumerates the set of values for ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnumValues() []ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum {
	values := make([]ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum(val string) (ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum Enum with underlying type: string
type ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum
const (
	ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByTimecreated ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum = "timeCreated"
	ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByDisplayname ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum = "displayName"
)

var mappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum = map[string]ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum{
	"timeCreated": ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByTimecreated,
	"displayName": ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByDisplayname,
}

var mappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnumLowerCase = map[string]ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum{
	"timecreated": ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByTimecreated,
	"displayname": ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByDisplayname,
}

// GetListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnumValues Enumerates the set of values for ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnumValues() []ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum {
	values := make([]ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum, 0)
	for _, v := range mappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnumStringValues Enumerates the set of values in String for ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum(val string) (ListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnum, bool) {
	enum, ok := mappingListDatabaseToolsDatabaseApiGatewayConfigPoolsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
