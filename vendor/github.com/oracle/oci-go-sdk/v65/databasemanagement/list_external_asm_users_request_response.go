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

// ListExternalAsmUsersRequest wrapper for the ListExternalAsmUsers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalAsmUsers.go.html to see an example of how to use ListExternalAsmUsersRequest.
type ListExternalAsmUsersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external ASM.
	ExternalAsmId *string `mandatory:"true" contributesTo:"path" name:"externalAsmId"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The
	// default sort order for `NAME` is ascending and it is case-sensitive.
	SortBy ListExternalAsmUsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalAsmUsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalAsmUsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalAsmUsersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalAsmUsersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalAsmUsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalAsmUsersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalAsmUsersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalAsmUsersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalAsmUsersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalAsmUsersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalAsmUsersResponse wrapper for the ListExternalAsmUsers operation
type ListExternalAsmUsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalAsmUserCollection instances
	ExternalAsmUserCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalAsmUsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalAsmUsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalAsmUsersSortByEnum Enum with underlying type: string
type ListExternalAsmUsersSortByEnum string

// Set of constants representing the allowable values for ListExternalAsmUsersSortByEnum
const (
	ListExternalAsmUsersSortByName ListExternalAsmUsersSortByEnum = "NAME"
)

var mappingListExternalAsmUsersSortByEnum = map[string]ListExternalAsmUsersSortByEnum{
	"NAME": ListExternalAsmUsersSortByName,
}

var mappingListExternalAsmUsersSortByEnumLowerCase = map[string]ListExternalAsmUsersSortByEnum{
	"name": ListExternalAsmUsersSortByName,
}

// GetListExternalAsmUsersSortByEnumValues Enumerates the set of values for ListExternalAsmUsersSortByEnum
func GetListExternalAsmUsersSortByEnumValues() []ListExternalAsmUsersSortByEnum {
	values := make([]ListExternalAsmUsersSortByEnum, 0)
	for _, v := range mappingListExternalAsmUsersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalAsmUsersSortByEnumStringValues Enumerates the set of values in String for ListExternalAsmUsersSortByEnum
func GetListExternalAsmUsersSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListExternalAsmUsersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalAsmUsersSortByEnum(val string) (ListExternalAsmUsersSortByEnum, bool) {
	enum, ok := mappingListExternalAsmUsersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalAsmUsersSortOrderEnum Enum with underlying type: string
type ListExternalAsmUsersSortOrderEnum string

// Set of constants representing the allowable values for ListExternalAsmUsersSortOrderEnum
const (
	ListExternalAsmUsersSortOrderAsc  ListExternalAsmUsersSortOrderEnum = "ASC"
	ListExternalAsmUsersSortOrderDesc ListExternalAsmUsersSortOrderEnum = "DESC"
)

var mappingListExternalAsmUsersSortOrderEnum = map[string]ListExternalAsmUsersSortOrderEnum{
	"ASC":  ListExternalAsmUsersSortOrderAsc,
	"DESC": ListExternalAsmUsersSortOrderDesc,
}

var mappingListExternalAsmUsersSortOrderEnumLowerCase = map[string]ListExternalAsmUsersSortOrderEnum{
	"asc":  ListExternalAsmUsersSortOrderAsc,
	"desc": ListExternalAsmUsersSortOrderDesc,
}

// GetListExternalAsmUsersSortOrderEnumValues Enumerates the set of values for ListExternalAsmUsersSortOrderEnum
func GetListExternalAsmUsersSortOrderEnumValues() []ListExternalAsmUsersSortOrderEnum {
	values := make([]ListExternalAsmUsersSortOrderEnum, 0)
	for _, v := range mappingListExternalAsmUsersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalAsmUsersSortOrderEnumStringValues Enumerates the set of values in String for ListExternalAsmUsersSortOrderEnum
func GetListExternalAsmUsersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalAsmUsersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalAsmUsersSortOrderEnum(val string) (ListExternalAsmUsersSortOrderEnum, bool) {
	enum, ok := mappingListExternalAsmUsersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
