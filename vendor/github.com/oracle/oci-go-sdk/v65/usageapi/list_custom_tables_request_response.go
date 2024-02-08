// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCustomTablesRequest wrapper for the ListCustomTables operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/ListCustomTables.go.html to see an example of how to use ListCustomTablesRequest.
type ListCustomTablesRequest struct {

	// The compartment ID in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The saved report ID in which to list resources.
	SavedReportId *string `mandatory:"true" contributesTo:"query" name:"savedReportId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximumimum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results.
	// This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. If not specified, the default is displayName.
	SortBy ListCustomTablesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListCustomTablesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCustomTablesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCustomTablesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCustomTablesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCustomTablesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCustomTablesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCustomTablesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCustomTablesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCustomTablesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCustomTablesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCustomTablesResponse wrapper for the ListCustomTables operation
type ListCustomTablesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CustomTableCollection instances
	CustomTableCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of custom tables. If this header appears in the response, then this
	// is a partial list of custom tables. Include this value as the `page` parameter in a subsequent
	// GET request, to get the next batch of custom tables.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCustomTablesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCustomTablesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCustomTablesSortByEnum Enum with underlying type: string
type ListCustomTablesSortByEnum string

// Set of constants representing the allowable values for ListCustomTablesSortByEnum
const (
	ListCustomTablesSortByDisplayname ListCustomTablesSortByEnum = "displayName"
)

var mappingListCustomTablesSortByEnum = map[string]ListCustomTablesSortByEnum{
	"displayName": ListCustomTablesSortByDisplayname,
}

var mappingListCustomTablesSortByEnumLowerCase = map[string]ListCustomTablesSortByEnum{
	"displayname": ListCustomTablesSortByDisplayname,
}

// GetListCustomTablesSortByEnumValues Enumerates the set of values for ListCustomTablesSortByEnum
func GetListCustomTablesSortByEnumValues() []ListCustomTablesSortByEnum {
	values := make([]ListCustomTablesSortByEnum, 0)
	for _, v := range mappingListCustomTablesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomTablesSortByEnumStringValues Enumerates the set of values in String for ListCustomTablesSortByEnum
func GetListCustomTablesSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListCustomTablesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomTablesSortByEnum(val string) (ListCustomTablesSortByEnum, bool) {
	enum, ok := mappingListCustomTablesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCustomTablesSortOrderEnum Enum with underlying type: string
type ListCustomTablesSortOrderEnum string

// Set of constants representing the allowable values for ListCustomTablesSortOrderEnum
const (
	ListCustomTablesSortOrderAsc  ListCustomTablesSortOrderEnum = "ASC"
	ListCustomTablesSortOrderDesc ListCustomTablesSortOrderEnum = "DESC"
)

var mappingListCustomTablesSortOrderEnum = map[string]ListCustomTablesSortOrderEnum{
	"ASC":  ListCustomTablesSortOrderAsc,
	"DESC": ListCustomTablesSortOrderDesc,
}

var mappingListCustomTablesSortOrderEnumLowerCase = map[string]ListCustomTablesSortOrderEnum{
	"asc":  ListCustomTablesSortOrderAsc,
	"desc": ListCustomTablesSortOrderDesc,
}

// GetListCustomTablesSortOrderEnumValues Enumerates the set of values for ListCustomTablesSortOrderEnum
func GetListCustomTablesSortOrderEnumValues() []ListCustomTablesSortOrderEnum {
	values := make([]ListCustomTablesSortOrderEnum, 0)
	for _, v := range mappingListCustomTablesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomTablesSortOrderEnumStringValues Enumerates the set of values in String for ListCustomTablesSortOrderEnum
func GetListCustomTablesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCustomTablesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomTablesSortOrderEnum(val string) (ListCustomTablesSortOrderEnum, bool) {
	enum, ok := mappingListCustomTablesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
