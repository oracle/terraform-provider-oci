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

// ListProxyUsersRequest wrapper for the ListProxyUsers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListProxyUsers.go.html to see an example of how to use ListProxyUsersRequest.
type ListProxyUsersRequest struct {

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
	SortBy ListProxyUsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListProxyUsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListProxyUsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProxyUsersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProxyUsersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProxyUsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProxyUsersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProxyUsersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProxyUsersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProxyUsersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProxyUsersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProxyUsersResponse wrapper for the ListProxyUsers operation
type ListProxyUsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProxyUserCollection instances
	ProxyUserCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListProxyUsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProxyUsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProxyUsersSortByEnum Enum with underlying type: string
type ListProxyUsersSortByEnum string

// Set of constants representing the allowable values for ListProxyUsersSortByEnum
const (
	ListProxyUsersSortByName ListProxyUsersSortByEnum = "NAME"
)

var mappingListProxyUsersSortByEnum = map[string]ListProxyUsersSortByEnum{
	"NAME": ListProxyUsersSortByName,
}

var mappingListProxyUsersSortByEnumLowerCase = map[string]ListProxyUsersSortByEnum{
	"name": ListProxyUsersSortByName,
}

// GetListProxyUsersSortByEnumValues Enumerates the set of values for ListProxyUsersSortByEnum
func GetListProxyUsersSortByEnumValues() []ListProxyUsersSortByEnum {
	values := make([]ListProxyUsersSortByEnum, 0)
	for _, v := range mappingListProxyUsersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProxyUsersSortByEnumStringValues Enumerates the set of values in String for ListProxyUsersSortByEnum
func GetListProxyUsersSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListProxyUsersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProxyUsersSortByEnum(val string) (ListProxyUsersSortByEnum, bool) {
	enum, ok := mappingListProxyUsersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProxyUsersSortOrderEnum Enum with underlying type: string
type ListProxyUsersSortOrderEnum string

// Set of constants representing the allowable values for ListProxyUsersSortOrderEnum
const (
	ListProxyUsersSortOrderAsc  ListProxyUsersSortOrderEnum = "ASC"
	ListProxyUsersSortOrderDesc ListProxyUsersSortOrderEnum = "DESC"
)

var mappingListProxyUsersSortOrderEnum = map[string]ListProxyUsersSortOrderEnum{
	"ASC":  ListProxyUsersSortOrderAsc,
	"DESC": ListProxyUsersSortOrderDesc,
}

var mappingListProxyUsersSortOrderEnumLowerCase = map[string]ListProxyUsersSortOrderEnum{
	"asc":  ListProxyUsersSortOrderAsc,
	"desc": ListProxyUsersSortOrderDesc,
}

// GetListProxyUsersSortOrderEnumValues Enumerates the set of values for ListProxyUsersSortOrderEnum
func GetListProxyUsersSortOrderEnumValues() []ListProxyUsersSortOrderEnum {
	values := make([]ListProxyUsersSortOrderEnum, 0)
	for _, v := range mappingListProxyUsersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProxyUsersSortOrderEnumStringValues Enumerates the set of values in String for ListProxyUsersSortOrderEnum
func GetListProxyUsersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProxyUsersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProxyUsersSortOrderEnum(val string) (ListProxyUsersSortOrderEnum, bool) {
	enum, ok := mappingListProxyUsersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
