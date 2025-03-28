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

// ListExternalAsmDiskGroupsRequest wrapper for the ListExternalAsmDiskGroups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalAsmDiskGroups.go.html to see an example of how to use ListExternalAsmDiskGroupsRequest.
type ListExternalAsmDiskGroupsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external ASM.
	ExternalAsmId *string `mandatory:"true" contributesTo:"path" name:"externalAsmId"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The
	// default sort order for `NAME` is ascending and it is case-sensitive.
	SortBy ListExternalAsmDiskGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalAsmDiskGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalAsmDiskGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalAsmDiskGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalAsmDiskGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalAsmDiskGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalAsmDiskGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalAsmDiskGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalAsmDiskGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalAsmDiskGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalAsmDiskGroupsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalAsmDiskGroupsResponse wrapper for the ListExternalAsmDiskGroups operation
type ListExternalAsmDiskGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalAsmDiskGroupCollection instances
	ExternalAsmDiskGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalAsmDiskGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalAsmDiskGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalAsmDiskGroupsSortByEnum Enum with underlying type: string
type ListExternalAsmDiskGroupsSortByEnum string

// Set of constants representing the allowable values for ListExternalAsmDiskGroupsSortByEnum
const (
	ListExternalAsmDiskGroupsSortByName ListExternalAsmDiskGroupsSortByEnum = "NAME"
)

var mappingListExternalAsmDiskGroupsSortByEnum = map[string]ListExternalAsmDiskGroupsSortByEnum{
	"NAME": ListExternalAsmDiskGroupsSortByName,
}

var mappingListExternalAsmDiskGroupsSortByEnumLowerCase = map[string]ListExternalAsmDiskGroupsSortByEnum{
	"name": ListExternalAsmDiskGroupsSortByName,
}

// GetListExternalAsmDiskGroupsSortByEnumValues Enumerates the set of values for ListExternalAsmDiskGroupsSortByEnum
func GetListExternalAsmDiskGroupsSortByEnumValues() []ListExternalAsmDiskGroupsSortByEnum {
	values := make([]ListExternalAsmDiskGroupsSortByEnum, 0)
	for _, v := range mappingListExternalAsmDiskGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalAsmDiskGroupsSortByEnumStringValues Enumerates the set of values in String for ListExternalAsmDiskGroupsSortByEnum
func GetListExternalAsmDiskGroupsSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListExternalAsmDiskGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalAsmDiskGroupsSortByEnum(val string) (ListExternalAsmDiskGroupsSortByEnum, bool) {
	enum, ok := mappingListExternalAsmDiskGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalAsmDiskGroupsSortOrderEnum Enum with underlying type: string
type ListExternalAsmDiskGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListExternalAsmDiskGroupsSortOrderEnum
const (
	ListExternalAsmDiskGroupsSortOrderAsc  ListExternalAsmDiskGroupsSortOrderEnum = "ASC"
	ListExternalAsmDiskGroupsSortOrderDesc ListExternalAsmDiskGroupsSortOrderEnum = "DESC"
)

var mappingListExternalAsmDiskGroupsSortOrderEnum = map[string]ListExternalAsmDiskGroupsSortOrderEnum{
	"ASC":  ListExternalAsmDiskGroupsSortOrderAsc,
	"DESC": ListExternalAsmDiskGroupsSortOrderDesc,
}

var mappingListExternalAsmDiskGroupsSortOrderEnumLowerCase = map[string]ListExternalAsmDiskGroupsSortOrderEnum{
	"asc":  ListExternalAsmDiskGroupsSortOrderAsc,
	"desc": ListExternalAsmDiskGroupsSortOrderDesc,
}

// GetListExternalAsmDiskGroupsSortOrderEnumValues Enumerates the set of values for ListExternalAsmDiskGroupsSortOrderEnum
func GetListExternalAsmDiskGroupsSortOrderEnumValues() []ListExternalAsmDiskGroupsSortOrderEnum {
	values := make([]ListExternalAsmDiskGroupsSortOrderEnum, 0)
	for _, v := range mappingListExternalAsmDiskGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalAsmDiskGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListExternalAsmDiskGroupsSortOrderEnum
func GetListExternalAsmDiskGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalAsmDiskGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalAsmDiskGroupsSortOrderEnum(val string) (ListExternalAsmDiskGroupsSortOrderEnum, bool) {
	enum, ok := mappingListExternalAsmDiskGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
