// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// UsersRequest wrapper for the Users operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/Users.go.html to see an example of how to use UsersRequest.
type UsersRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy UsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder UsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UsersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UsersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UsersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUsersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetUsersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUsersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetUsersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UsersResponse wrapper for the Users operation
type UsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of string instances
	Value *string `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response UsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// UsersSortByEnum Enum with underlying type: string
type UsersSortByEnum string

// Set of constants representing the allowable values for UsersSortByEnum
const (
	UsersSortByTimecreated UsersSortByEnum = "TIMECREATED"
	UsersSortByDisplayname UsersSortByEnum = "DISPLAYNAME"
)

var mappingUsersSortByEnum = map[string]UsersSortByEnum{
	"TIMECREATED": UsersSortByTimecreated,
	"DISPLAYNAME": UsersSortByDisplayname,
}

// GetUsersSortByEnumValues Enumerates the set of values for UsersSortByEnum
func GetUsersSortByEnumValues() []UsersSortByEnum {
	values := make([]UsersSortByEnum, 0)
	for _, v := range mappingUsersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetUsersSortByEnumStringValues Enumerates the set of values in String for UsersSortByEnum
func GetUsersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingUsersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUsersSortByEnum(val string) (UsersSortByEnum, bool) {
	mappingUsersSortByEnumIgnoreCase := make(map[string]UsersSortByEnum)
	for k, v := range mappingUsersSortByEnum {
		mappingUsersSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUsersSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UsersSortOrderEnum Enum with underlying type: string
type UsersSortOrderEnum string

// Set of constants representing the allowable values for UsersSortOrderEnum
const (
	UsersSortOrderAsc  UsersSortOrderEnum = "ASC"
	UsersSortOrderDesc UsersSortOrderEnum = "DESC"
)

var mappingUsersSortOrderEnum = map[string]UsersSortOrderEnum{
	"ASC":  UsersSortOrderAsc,
	"DESC": UsersSortOrderDesc,
}

// GetUsersSortOrderEnumValues Enumerates the set of values for UsersSortOrderEnum
func GetUsersSortOrderEnumValues() []UsersSortOrderEnum {
	values := make([]UsersSortOrderEnum, 0)
	for _, v := range mappingUsersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetUsersSortOrderEnumStringValues Enumerates the set of values in String for UsersSortOrderEnum
func GetUsersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingUsersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUsersSortOrderEnum(val string) (UsersSortOrderEnum, bool) {
	mappingUsersSortOrderEnumIgnoreCase := make(map[string]UsersSortOrderEnum)
	for k, v := range mappingUsersSortOrderEnum {
		mappingUsersSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUsersSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
