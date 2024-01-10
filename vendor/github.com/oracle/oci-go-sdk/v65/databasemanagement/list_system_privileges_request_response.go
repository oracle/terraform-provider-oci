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

// ListSystemPrivilegesRequest wrapper for the ListSystemPrivileges operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSystemPrivileges.go.html to see an example of how to use ListSystemPrivilegesRequest.
type ListSystemPrivilegesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The name of the user whose details are to be viewed.
	UserName *string `mandatory:"true" contributesTo:"path" name:"userName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘NAME’ is ascending. The ‘NAME’ sort order is case-sensitive.
	SortBy ListSystemPrivilegesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListSystemPrivilegesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSystemPrivilegesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSystemPrivilegesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSystemPrivilegesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSystemPrivilegesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSystemPrivilegesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSystemPrivilegesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSystemPrivilegesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSystemPrivilegesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSystemPrivilegesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSystemPrivilegesResponse wrapper for the ListSystemPrivileges operation
type ListSystemPrivilegesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SystemPrivilegeCollection instances
	SystemPrivilegeCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSystemPrivilegesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSystemPrivilegesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSystemPrivilegesSortByEnum Enum with underlying type: string
type ListSystemPrivilegesSortByEnum string

// Set of constants representing the allowable values for ListSystemPrivilegesSortByEnum
const (
	ListSystemPrivilegesSortByName ListSystemPrivilegesSortByEnum = "NAME"
)

var mappingListSystemPrivilegesSortByEnum = map[string]ListSystemPrivilegesSortByEnum{
	"NAME": ListSystemPrivilegesSortByName,
}

var mappingListSystemPrivilegesSortByEnumLowerCase = map[string]ListSystemPrivilegesSortByEnum{
	"name": ListSystemPrivilegesSortByName,
}

// GetListSystemPrivilegesSortByEnumValues Enumerates the set of values for ListSystemPrivilegesSortByEnum
func GetListSystemPrivilegesSortByEnumValues() []ListSystemPrivilegesSortByEnum {
	values := make([]ListSystemPrivilegesSortByEnum, 0)
	for _, v := range mappingListSystemPrivilegesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSystemPrivilegesSortByEnumStringValues Enumerates the set of values in String for ListSystemPrivilegesSortByEnum
func GetListSystemPrivilegesSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListSystemPrivilegesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSystemPrivilegesSortByEnum(val string) (ListSystemPrivilegesSortByEnum, bool) {
	enum, ok := mappingListSystemPrivilegesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSystemPrivilegesSortOrderEnum Enum with underlying type: string
type ListSystemPrivilegesSortOrderEnum string

// Set of constants representing the allowable values for ListSystemPrivilegesSortOrderEnum
const (
	ListSystemPrivilegesSortOrderAsc  ListSystemPrivilegesSortOrderEnum = "ASC"
	ListSystemPrivilegesSortOrderDesc ListSystemPrivilegesSortOrderEnum = "DESC"
)

var mappingListSystemPrivilegesSortOrderEnum = map[string]ListSystemPrivilegesSortOrderEnum{
	"ASC":  ListSystemPrivilegesSortOrderAsc,
	"DESC": ListSystemPrivilegesSortOrderDesc,
}

var mappingListSystemPrivilegesSortOrderEnumLowerCase = map[string]ListSystemPrivilegesSortOrderEnum{
	"asc":  ListSystemPrivilegesSortOrderAsc,
	"desc": ListSystemPrivilegesSortOrderDesc,
}

// GetListSystemPrivilegesSortOrderEnumValues Enumerates the set of values for ListSystemPrivilegesSortOrderEnum
func GetListSystemPrivilegesSortOrderEnumValues() []ListSystemPrivilegesSortOrderEnum {
	values := make([]ListSystemPrivilegesSortOrderEnum, 0)
	for _, v := range mappingListSystemPrivilegesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSystemPrivilegesSortOrderEnumStringValues Enumerates the set of values in String for ListSystemPrivilegesSortOrderEnum
func GetListSystemPrivilegesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSystemPrivilegesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSystemPrivilegesSortOrderEnum(val string) (ListSystemPrivilegesSortOrderEnum, bool) {
	enum, ok := mappingListSystemPrivilegesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
