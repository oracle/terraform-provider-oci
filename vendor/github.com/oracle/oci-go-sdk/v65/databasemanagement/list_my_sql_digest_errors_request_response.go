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

// ListMySqlDigestErrorsRequest wrapper for the ListMySqlDigestErrors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListMySqlDigestErrors.go.html to see an example of how to use ListMySqlDigestErrorsRequest.
type ListMySqlDigestErrorsRequest struct {

	// The OCID of the Managed MySQL Database.
	ManagedMySqlDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedMySqlDatabaseId"`

	// The digest of a MySQL normalized query.
	Digest *string `mandatory:"true" contributesTo:"query" name:"digest"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used.
	// The default sort by field is ‘occurrenceCount’.
	// The default sort order for ‘occurrenceCount’ is descending.
	SortBy ListMySqlDigestErrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListMySqlDigestErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMySqlDigestErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMySqlDigestErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMySqlDigestErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMySqlDigestErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMySqlDigestErrorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMySqlDigestErrorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMySqlDigestErrorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMySqlDigestErrorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMySqlDigestErrorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMySqlDigestErrorsResponse wrapper for the ListMySqlDigestErrors operation
type ListMySqlDigestErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MySqlDigestErrorsCollection instances
	MySqlDigestErrorsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMySqlDigestErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMySqlDigestErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMySqlDigestErrorsSortByEnum Enum with underlying type: string
type ListMySqlDigestErrorsSortByEnum string

// Set of constants representing the allowable values for ListMySqlDigestErrorsSortByEnum
const (
	ListMySqlDigestErrorsSortByOccurrencecount ListMySqlDigestErrorsSortByEnum = "occurrenceCount"
)

var mappingListMySqlDigestErrorsSortByEnum = map[string]ListMySqlDigestErrorsSortByEnum{
	"occurrenceCount": ListMySqlDigestErrorsSortByOccurrencecount,
}

var mappingListMySqlDigestErrorsSortByEnumLowerCase = map[string]ListMySqlDigestErrorsSortByEnum{
	"occurrencecount": ListMySqlDigestErrorsSortByOccurrencecount,
}

// GetListMySqlDigestErrorsSortByEnumValues Enumerates the set of values for ListMySqlDigestErrorsSortByEnum
func GetListMySqlDigestErrorsSortByEnumValues() []ListMySqlDigestErrorsSortByEnum {
	values := make([]ListMySqlDigestErrorsSortByEnum, 0)
	for _, v := range mappingListMySqlDigestErrorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMySqlDigestErrorsSortByEnumStringValues Enumerates the set of values in String for ListMySqlDigestErrorsSortByEnum
func GetListMySqlDigestErrorsSortByEnumStringValues() []string {
	return []string{
		"occurrenceCount",
	}
}

// GetMappingListMySqlDigestErrorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMySqlDigestErrorsSortByEnum(val string) (ListMySqlDigestErrorsSortByEnum, bool) {
	enum, ok := mappingListMySqlDigestErrorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMySqlDigestErrorsSortOrderEnum Enum with underlying type: string
type ListMySqlDigestErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListMySqlDigestErrorsSortOrderEnum
const (
	ListMySqlDigestErrorsSortOrderAsc  ListMySqlDigestErrorsSortOrderEnum = "ASC"
	ListMySqlDigestErrorsSortOrderDesc ListMySqlDigestErrorsSortOrderEnum = "DESC"
)

var mappingListMySqlDigestErrorsSortOrderEnum = map[string]ListMySqlDigestErrorsSortOrderEnum{
	"ASC":  ListMySqlDigestErrorsSortOrderAsc,
	"DESC": ListMySqlDigestErrorsSortOrderDesc,
}

var mappingListMySqlDigestErrorsSortOrderEnumLowerCase = map[string]ListMySqlDigestErrorsSortOrderEnum{
	"asc":  ListMySqlDigestErrorsSortOrderAsc,
	"desc": ListMySqlDigestErrorsSortOrderDesc,
}

// GetListMySqlDigestErrorsSortOrderEnumValues Enumerates the set of values for ListMySqlDigestErrorsSortOrderEnum
func GetListMySqlDigestErrorsSortOrderEnumValues() []ListMySqlDigestErrorsSortOrderEnum {
	values := make([]ListMySqlDigestErrorsSortOrderEnum, 0)
	for _, v := range mappingListMySqlDigestErrorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMySqlDigestErrorsSortOrderEnumStringValues Enumerates the set of values in String for ListMySqlDigestErrorsSortOrderEnum
func GetListMySqlDigestErrorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMySqlDigestErrorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMySqlDigestErrorsSortOrderEnum(val string) (ListMySqlDigestErrorsSortOrderEnum, bool) {
	enum, ok := mappingListMySqlDigestErrorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
