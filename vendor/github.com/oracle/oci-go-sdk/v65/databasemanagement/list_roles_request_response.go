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

// ListRolesRequest wrapper for the ListRoles operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListRoles.go.html to see an example of how to use ListRolesRequest.
type ListRolesRequest struct {

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
	SortBy ListRolesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListRolesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRolesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRolesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRolesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRolesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRolesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRolesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRolesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRolesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRolesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRolesResponse wrapper for the ListRoles operation
type ListRolesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RoleCollection instances
	RoleCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListRolesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRolesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRolesSortByEnum Enum with underlying type: string
type ListRolesSortByEnum string

// Set of constants representing the allowable values for ListRolesSortByEnum
const (
	ListRolesSortByName ListRolesSortByEnum = "NAME"
)

var mappingListRolesSortByEnum = map[string]ListRolesSortByEnum{
	"NAME": ListRolesSortByName,
}

var mappingListRolesSortByEnumLowerCase = map[string]ListRolesSortByEnum{
	"name": ListRolesSortByName,
}

// GetListRolesSortByEnumValues Enumerates the set of values for ListRolesSortByEnum
func GetListRolesSortByEnumValues() []ListRolesSortByEnum {
	values := make([]ListRolesSortByEnum, 0)
	for _, v := range mappingListRolesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRolesSortByEnumStringValues Enumerates the set of values in String for ListRolesSortByEnum
func GetListRolesSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListRolesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRolesSortByEnum(val string) (ListRolesSortByEnum, bool) {
	enum, ok := mappingListRolesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRolesSortOrderEnum Enum with underlying type: string
type ListRolesSortOrderEnum string

// Set of constants representing the allowable values for ListRolesSortOrderEnum
const (
	ListRolesSortOrderAsc  ListRolesSortOrderEnum = "ASC"
	ListRolesSortOrderDesc ListRolesSortOrderEnum = "DESC"
)

var mappingListRolesSortOrderEnum = map[string]ListRolesSortOrderEnum{
	"ASC":  ListRolesSortOrderAsc,
	"DESC": ListRolesSortOrderDesc,
}

var mappingListRolesSortOrderEnumLowerCase = map[string]ListRolesSortOrderEnum{
	"asc":  ListRolesSortOrderAsc,
	"desc": ListRolesSortOrderDesc,
}

// GetListRolesSortOrderEnumValues Enumerates the set of values for ListRolesSortOrderEnum
func GetListRolesSortOrderEnumValues() []ListRolesSortOrderEnum {
	values := make([]ListRolesSortOrderEnum, 0)
	for _, v := range mappingListRolesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRolesSortOrderEnumStringValues Enumerates the set of values in String for ListRolesSortOrderEnum
func GetListRolesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRolesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRolesSortOrderEnum(val string) (ListRolesSortOrderEnum, bool) {
	enum, ok := mappingListRolesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
