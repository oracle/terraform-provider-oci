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

// ListSqlTuningSetsRequest wrapper for the ListSqlTuningSets operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSqlTuningSets.go.html to see an example of how to use ListSqlTuningSetsRequest.
type ListSqlTuningSetsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The owner of the SQL tuning set.
	Owner *string `mandatory:"false" contributesTo:"query" name:"owner"`

	// Allow searching the name of the SQL tuning set by partial matching. The search is case insensitive.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// The option to sort the SQL tuning set summary data.
	SortBy ListSqlTuningSetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListSqlTuningSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlTuningSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlTuningSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlTuningSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlTuningSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlTuningSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlTuningSetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSqlTuningSetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlTuningSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSqlTuningSetsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlTuningSetsResponse wrapper for the ListSqlTuningSets operation
type ListSqlTuningSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlTuningSetCollection instances
	SqlTuningSetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSqlTuningSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlTuningSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlTuningSetsSortByEnum Enum with underlying type: string
type ListSqlTuningSetsSortByEnum string

// Set of constants representing the allowable values for ListSqlTuningSetsSortByEnum
const (
	ListSqlTuningSetsSortByName ListSqlTuningSetsSortByEnum = "NAME"
)

var mappingListSqlTuningSetsSortByEnum = map[string]ListSqlTuningSetsSortByEnum{
	"NAME": ListSqlTuningSetsSortByName,
}

var mappingListSqlTuningSetsSortByEnumLowerCase = map[string]ListSqlTuningSetsSortByEnum{
	"name": ListSqlTuningSetsSortByName,
}

// GetListSqlTuningSetsSortByEnumValues Enumerates the set of values for ListSqlTuningSetsSortByEnum
func GetListSqlTuningSetsSortByEnumValues() []ListSqlTuningSetsSortByEnum {
	values := make([]ListSqlTuningSetsSortByEnum, 0)
	for _, v := range mappingListSqlTuningSetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlTuningSetsSortByEnumStringValues Enumerates the set of values in String for ListSqlTuningSetsSortByEnum
func GetListSqlTuningSetsSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListSqlTuningSetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlTuningSetsSortByEnum(val string) (ListSqlTuningSetsSortByEnum, bool) {
	enum, ok := mappingListSqlTuningSetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlTuningSetsSortOrderEnum Enum with underlying type: string
type ListSqlTuningSetsSortOrderEnum string

// Set of constants representing the allowable values for ListSqlTuningSetsSortOrderEnum
const (
	ListSqlTuningSetsSortOrderAsc  ListSqlTuningSetsSortOrderEnum = "ASC"
	ListSqlTuningSetsSortOrderDesc ListSqlTuningSetsSortOrderEnum = "DESC"
)

var mappingListSqlTuningSetsSortOrderEnum = map[string]ListSqlTuningSetsSortOrderEnum{
	"ASC":  ListSqlTuningSetsSortOrderAsc,
	"DESC": ListSqlTuningSetsSortOrderDesc,
}

var mappingListSqlTuningSetsSortOrderEnumLowerCase = map[string]ListSqlTuningSetsSortOrderEnum{
	"asc":  ListSqlTuningSetsSortOrderAsc,
	"desc": ListSqlTuningSetsSortOrderDesc,
}

// GetListSqlTuningSetsSortOrderEnumValues Enumerates the set of values for ListSqlTuningSetsSortOrderEnum
func GetListSqlTuningSetsSortOrderEnumValues() []ListSqlTuningSetsSortOrderEnum {
	values := make([]ListSqlTuningSetsSortOrderEnum, 0)
	for _, v := range mappingListSqlTuningSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlTuningSetsSortOrderEnumStringValues Enumerates the set of values in String for ListSqlTuningSetsSortOrderEnum
func GetListSqlTuningSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSqlTuningSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlTuningSetsSortOrderEnum(val string) (ListSqlTuningSetsSortOrderEnum, bool) {
	enum, ok := mappingListSqlTuningSetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
