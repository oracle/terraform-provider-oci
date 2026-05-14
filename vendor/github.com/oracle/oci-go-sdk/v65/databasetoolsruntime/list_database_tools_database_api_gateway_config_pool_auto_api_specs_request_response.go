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

// ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest wrapper for the ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs.go.html to see an example of how to use ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest.
type ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest struct {

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
	SortOrder ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse wrapper for the ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs operation
type ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecCollection instances
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum Enum with underlying type: string
type ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum
const (
	ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderAsc  ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum = "ASC"
	ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderDesc ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum = "DESC"
)

var mappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum = map[string]ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum{
	"ASC":  ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderAsc,
	"DESC": ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderDesc,
}

var mappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnumLowerCase = map[string]ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum{
	"asc":  ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderAsc,
	"desc": ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderDesc,
}

// GetListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnumValues Enumerates the set of values for ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnumValues() []ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum {
	values := make([]ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum(val string) (ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum Enum with underlying type: string
type ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum
const (
	ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByTimecreated ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum = "timeCreated"
	ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByDisplayname ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum = "displayName"
)

var mappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum = map[string]ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum{
	"timeCreated": ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByTimecreated,
	"displayName": ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByDisplayname,
}

var mappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnumLowerCase = map[string]ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum{
	"timecreated": ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByTimecreated,
	"displayname": ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByDisplayname,
}

// GetListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnumValues Enumerates the set of values for ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnumValues() []ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum {
	values := make([]ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum, 0)
	for _, v := range mappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnumStringValues Enumerates the set of values in String for ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum(val string) (ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnum, bool) {
	enum, ok := mappingListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
