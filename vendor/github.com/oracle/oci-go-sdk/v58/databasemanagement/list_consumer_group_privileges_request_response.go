// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListConsumerGroupPrivilegesRequest wrapper for the ListConsumerGroupPrivileges operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListConsumerGroupPrivileges.go.html to see an example of how to use ListConsumerGroupPrivilegesRequest.
type ListConsumerGroupPrivilegesRequest struct {

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
	SortBy ListConsumerGroupPrivilegesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListConsumerGroupPrivilegesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConsumerGroupPrivilegesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConsumerGroupPrivilegesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConsumerGroupPrivilegesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConsumerGroupPrivilegesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConsumerGroupPrivilegesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListConsumerGroupPrivilegesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConsumerGroupPrivilegesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConsumerGroupPrivilegesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConsumerGroupPrivilegesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConsumerGroupPrivilegesResponse wrapper for the ListConsumerGroupPrivileges operation
type ListConsumerGroupPrivilegesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConsumerGroupPrivilegeCollection instances
	ConsumerGroupPrivilegeCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListConsumerGroupPrivilegesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConsumerGroupPrivilegesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConsumerGroupPrivilegesSortByEnum Enum with underlying type: string
type ListConsumerGroupPrivilegesSortByEnum string

// Set of constants representing the allowable values for ListConsumerGroupPrivilegesSortByEnum
const (
	ListConsumerGroupPrivilegesSortByName ListConsumerGroupPrivilegesSortByEnum = "NAME"
)

var mappingListConsumerGroupPrivilegesSortByEnum = map[string]ListConsumerGroupPrivilegesSortByEnum{
	"NAME": ListConsumerGroupPrivilegesSortByName,
}

// GetListConsumerGroupPrivilegesSortByEnumValues Enumerates the set of values for ListConsumerGroupPrivilegesSortByEnum
func GetListConsumerGroupPrivilegesSortByEnumValues() []ListConsumerGroupPrivilegesSortByEnum {
	values := make([]ListConsumerGroupPrivilegesSortByEnum, 0)
	for _, v := range mappingListConsumerGroupPrivilegesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConsumerGroupPrivilegesSortByEnumStringValues Enumerates the set of values in String for ListConsumerGroupPrivilegesSortByEnum
func GetListConsumerGroupPrivilegesSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListConsumerGroupPrivilegesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConsumerGroupPrivilegesSortByEnum(val string) (ListConsumerGroupPrivilegesSortByEnum, bool) {
	mappingListConsumerGroupPrivilegesSortByEnumIgnoreCase := make(map[string]ListConsumerGroupPrivilegesSortByEnum)
	for k, v := range mappingListConsumerGroupPrivilegesSortByEnum {
		mappingListConsumerGroupPrivilegesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListConsumerGroupPrivilegesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListConsumerGroupPrivilegesSortOrderEnum Enum with underlying type: string
type ListConsumerGroupPrivilegesSortOrderEnum string

// Set of constants representing the allowable values for ListConsumerGroupPrivilegesSortOrderEnum
const (
	ListConsumerGroupPrivilegesSortOrderAsc  ListConsumerGroupPrivilegesSortOrderEnum = "ASC"
	ListConsumerGroupPrivilegesSortOrderDesc ListConsumerGroupPrivilegesSortOrderEnum = "DESC"
)

var mappingListConsumerGroupPrivilegesSortOrderEnum = map[string]ListConsumerGroupPrivilegesSortOrderEnum{
	"ASC":  ListConsumerGroupPrivilegesSortOrderAsc,
	"DESC": ListConsumerGroupPrivilegesSortOrderDesc,
}

// GetListConsumerGroupPrivilegesSortOrderEnumValues Enumerates the set of values for ListConsumerGroupPrivilegesSortOrderEnum
func GetListConsumerGroupPrivilegesSortOrderEnumValues() []ListConsumerGroupPrivilegesSortOrderEnum {
	values := make([]ListConsumerGroupPrivilegesSortOrderEnum, 0)
	for _, v := range mappingListConsumerGroupPrivilegesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConsumerGroupPrivilegesSortOrderEnumStringValues Enumerates the set of values in String for ListConsumerGroupPrivilegesSortOrderEnum
func GetListConsumerGroupPrivilegesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConsumerGroupPrivilegesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConsumerGroupPrivilegesSortOrderEnum(val string) (ListConsumerGroupPrivilegesSortOrderEnum, bool) {
	mappingListConsumerGroupPrivilegesSortOrderEnumIgnoreCase := make(map[string]ListConsumerGroupPrivilegesSortOrderEnum)
	for k, v := range mappingListConsumerGroupPrivilegesSortOrderEnum {
		mappingListConsumerGroupPrivilegesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListConsumerGroupPrivilegesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
