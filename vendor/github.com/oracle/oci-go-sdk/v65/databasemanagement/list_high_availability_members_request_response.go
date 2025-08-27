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

// ListHighAvailabilityMembersRequest wrapper for the ListHighAvailabilityMembers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListHighAvailabilityMembers.go.html to see an example of how to use ListHighAvailabilityMembersRequest.
type ListHighAvailabilityMembersRequest struct {

	// The OCID of the Managed MySQL Database.
	ManagedMySqlDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedMySqlDatabaseId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListHighAvailabilityMembersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort information by. Only one sortOrder can be used.
	// The default sort by field is ‘memberHost’.
	// The default sort order for ‘memberHost’ is ascending.
	// The ‘memberHost’ sort order is case-sensitive.
	SortBy ListHighAvailabilityMembersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHighAvailabilityMembersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHighAvailabilityMembersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHighAvailabilityMembersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHighAvailabilityMembersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListHighAvailabilityMembersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListHighAvailabilityMembersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListHighAvailabilityMembersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHighAvailabilityMembersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListHighAvailabilityMembersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListHighAvailabilityMembersResponse wrapper for the ListHighAvailabilityMembers operation
type ListHighAvailabilityMembersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedMySqlDatabaseHighAvailabilityMemberCollection instances
	ManagedMySqlDatabaseHighAvailabilityMemberCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListHighAvailabilityMembersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHighAvailabilityMembersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHighAvailabilityMembersSortOrderEnum Enum with underlying type: string
type ListHighAvailabilityMembersSortOrderEnum string

// Set of constants representing the allowable values for ListHighAvailabilityMembersSortOrderEnum
const (
	ListHighAvailabilityMembersSortOrderAsc  ListHighAvailabilityMembersSortOrderEnum = "ASC"
	ListHighAvailabilityMembersSortOrderDesc ListHighAvailabilityMembersSortOrderEnum = "DESC"
)

var mappingListHighAvailabilityMembersSortOrderEnum = map[string]ListHighAvailabilityMembersSortOrderEnum{
	"ASC":  ListHighAvailabilityMembersSortOrderAsc,
	"DESC": ListHighAvailabilityMembersSortOrderDesc,
}

var mappingListHighAvailabilityMembersSortOrderEnumLowerCase = map[string]ListHighAvailabilityMembersSortOrderEnum{
	"asc":  ListHighAvailabilityMembersSortOrderAsc,
	"desc": ListHighAvailabilityMembersSortOrderDesc,
}

// GetListHighAvailabilityMembersSortOrderEnumValues Enumerates the set of values for ListHighAvailabilityMembersSortOrderEnum
func GetListHighAvailabilityMembersSortOrderEnumValues() []ListHighAvailabilityMembersSortOrderEnum {
	values := make([]ListHighAvailabilityMembersSortOrderEnum, 0)
	for _, v := range mappingListHighAvailabilityMembersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListHighAvailabilityMembersSortOrderEnumStringValues Enumerates the set of values in String for ListHighAvailabilityMembersSortOrderEnum
func GetListHighAvailabilityMembersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListHighAvailabilityMembersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHighAvailabilityMembersSortOrderEnum(val string) (ListHighAvailabilityMembersSortOrderEnum, bool) {
	enum, ok := mappingListHighAvailabilityMembersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListHighAvailabilityMembersSortByEnum Enum with underlying type: string
type ListHighAvailabilityMembersSortByEnum string

// Set of constants representing the allowable values for ListHighAvailabilityMembersSortByEnum
const (
	ListHighAvailabilityMembersSortByMemberhost ListHighAvailabilityMembersSortByEnum = "memberHost"
)

var mappingListHighAvailabilityMembersSortByEnum = map[string]ListHighAvailabilityMembersSortByEnum{
	"memberHost": ListHighAvailabilityMembersSortByMemberhost,
}

var mappingListHighAvailabilityMembersSortByEnumLowerCase = map[string]ListHighAvailabilityMembersSortByEnum{
	"memberhost": ListHighAvailabilityMembersSortByMemberhost,
}

// GetListHighAvailabilityMembersSortByEnumValues Enumerates the set of values for ListHighAvailabilityMembersSortByEnum
func GetListHighAvailabilityMembersSortByEnumValues() []ListHighAvailabilityMembersSortByEnum {
	values := make([]ListHighAvailabilityMembersSortByEnum, 0)
	for _, v := range mappingListHighAvailabilityMembersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListHighAvailabilityMembersSortByEnumStringValues Enumerates the set of values in String for ListHighAvailabilityMembersSortByEnum
func GetListHighAvailabilityMembersSortByEnumStringValues() []string {
	return []string{
		"memberHost",
	}
}

// GetMappingListHighAvailabilityMembersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHighAvailabilityMembersSortByEnum(val string) (ListHighAvailabilityMembersSortByEnum, bool) {
	enum, ok := mappingListHighAvailabilityMembersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
