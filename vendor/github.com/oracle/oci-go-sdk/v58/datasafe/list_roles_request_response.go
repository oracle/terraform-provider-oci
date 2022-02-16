// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListRolesRequest wrapper for the ListRoles operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListRoles.go.html to see an example of how to use ListRolesRequest.
type ListRolesRequest struct {

	// The OCID of the Data Safe target database.
	TargetDatabaseId *string `mandatory:"true" contributesTo:"path" name:"targetDatabaseId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only a specific role based on role name.
	RoleName []string `contributesTo:"query" name:"roleName" collectionFormat:"multi"`

	// A filter to return roles based on whether they are maintained by oracle or not.
	IsOracleMaintained *bool `mandatory:"false" contributesTo:"query" name:"isOracleMaintained"`

	// A filter to return roles based on authentication type.
	AuthenticationType *string `mandatory:"false" contributesTo:"query" name:"authenticationType"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListRolesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	SortBy ListRolesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only items if role name contains a specific string.
	RoleNameContains *string `mandatory:"false" contributesTo:"query" name:"roleNameContains"`

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
	if _, ok := GetMappingListRolesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRolesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRolesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRolesSortByEnumStringValues(), ",")))
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

	// A list of []RoleSummary instances
	Items []RoleSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListRolesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRolesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
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
	mappingListRolesSortOrderEnumIgnoreCase := make(map[string]ListRolesSortOrderEnum)
	for k, v := range mappingListRolesSortOrderEnum {
		mappingListRolesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRolesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListRolesSortByEnum Enum with underlying type: string
type ListRolesSortByEnum string

// Set of constants representing the allowable values for ListRolesSortByEnum
const (
	ListRolesSortByRolename ListRolesSortByEnum = "ROLENAME"
)

var mappingListRolesSortByEnum = map[string]ListRolesSortByEnum{
	"ROLENAME": ListRolesSortByRolename,
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
		"ROLENAME",
	}
}

// GetMappingListRolesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRolesSortByEnum(val string) (ListRolesSortByEnum, bool) {
	mappingListRolesSortByEnumIgnoreCase := make(map[string]ListRolesSortByEnum)
	for k, v := range mappingListRolesSortByEnum {
		mappingListRolesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRolesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
