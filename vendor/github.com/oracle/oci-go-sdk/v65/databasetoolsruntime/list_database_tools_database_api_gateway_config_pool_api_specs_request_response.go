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

// ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest wrapper for the ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs.go.html to see an example of how to use ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest.
type ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
	DatabaseToolsDatabaseApiGatewayConfigId *string `mandatory:"true" contributesTo:"path" name:"databaseToolsDatabaseApiGatewayConfigId"`

	// The key of the pool config.
	PoolKey *string `mandatory:"true" contributesTo:"path" name:"poolKey"`

	// A filter to return only resources that match the entire specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse wrapper for the ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs operation
type ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCollection instances
	DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum Enum with underlying type: string
type ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum
const (
	ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderAsc  ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum = "ASC"
	ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderDesc ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum = "DESC"
)

var mappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum = map[string]ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum{
	"ASC":  ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderAsc,
	"DESC": ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderDesc,
}

var mappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnumLowerCase = map[string]ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum{
	"asc":  ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderAsc,
	"desc": ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderDesc,
}

// GetListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnumValues Enumerates the set of values for ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnumValues() []ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum {
	values := make([]ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum(val string) (ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum Enum with underlying type: string
type ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum
const (
	ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByTimecreated ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum = "timeCreated"
	ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByDisplayname ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum = "displayName"
)

var mappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum = map[string]ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum{
	"timeCreated": ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByTimecreated,
	"displayName": ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByDisplayname,
}

var mappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnumLowerCase = map[string]ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum{
	"timecreated": ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByTimecreated,
	"displayname": ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByDisplayname,
}

// GetListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnumValues Enumerates the set of values for ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnumValues() []ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum {
	values := make([]ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum, 0)
	for _, v := range mappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnumStringValues Enumerates the set of values in String for ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum(val string) (ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnum, bool) {
	enum, ok := mappingListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
