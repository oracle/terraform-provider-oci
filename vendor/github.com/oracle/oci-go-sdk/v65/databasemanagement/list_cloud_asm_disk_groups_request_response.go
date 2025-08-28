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

// ListCloudAsmDiskGroupsRequest wrapper for the ListCloudAsmDiskGroups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudAsmDiskGroups.go.html to see an example of how to use ListCloudAsmDiskGroupsRequest.
type ListCloudAsmDiskGroupsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM.
	CloudAsmId *string `mandatory:"true" contributesTo:"path" name:"cloudAsmId"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The
	// default sort order for `NAME` is ascending and it is case-sensitive.
	SortBy ListCloudAsmDiskGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudAsmDiskGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudAsmDiskGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudAsmDiskGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudAsmDiskGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudAsmDiskGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudAsmDiskGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudAsmDiskGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudAsmDiskGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudAsmDiskGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudAsmDiskGroupsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudAsmDiskGroupsResponse wrapper for the ListCloudAsmDiskGroups operation
type ListCloudAsmDiskGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudAsmDiskGroupCollection instances
	CloudAsmDiskGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudAsmDiskGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudAsmDiskGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudAsmDiskGroupsSortByEnum Enum with underlying type: string
type ListCloudAsmDiskGroupsSortByEnum string

// Set of constants representing the allowable values for ListCloudAsmDiskGroupsSortByEnum
const (
	ListCloudAsmDiskGroupsSortByName ListCloudAsmDiskGroupsSortByEnum = "NAME"
)

var mappingListCloudAsmDiskGroupsSortByEnum = map[string]ListCloudAsmDiskGroupsSortByEnum{
	"NAME": ListCloudAsmDiskGroupsSortByName,
}

var mappingListCloudAsmDiskGroupsSortByEnumLowerCase = map[string]ListCloudAsmDiskGroupsSortByEnum{
	"name": ListCloudAsmDiskGroupsSortByName,
}

// GetListCloudAsmDiskGroupsSortByEnumValues Enumerates the set of values for ListCloudAsmDiskGroupsSortByEnum
func GetListCloudAsmDiskGroupsSortByEnumValues() []ListCloudAsmDiskGroupsSortByEnum {
	values := make([]ListCloudAsmDiskGroupsSortByEnum, 0)
	for _, v := range mappingListCloudAsmDiskGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudAsmDiskGroupsSortByEnumStringValues Enumerates the set of values in String for ListCloudAsmDiskGroupsSortByEnum
func GetListCloudAsmDiskGroupsSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListCloudAsmDiskGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudAsmDiskGroupsSortByEnum(val string) (ListCloudAsmDiskGroupsSortByEnum, bool) {
	enum, ok := mappingListCloudAsmDiskGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudAsmDiskGroupsSortOrderEnum Enum with underlying type: string
type ListCloudAsmDiskGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListCloudAsmDiskGroupsSortOrderEnum
const (
	ListCloudAsmDiskGroupsSortOrderAsc  ListCloudAsmDiskGroupsSortOrderEnum = "ASC"
	ListCloudAsmDiskGroupsSortOrderDesc ListCloudAsmDiskGroupsSortOrderEnum = "DESC"
)

var mappingListCloudAsmDiskGroupsSortOrderEnum = map[string]ListCloudAsmDiskGroupsSortOrderEnum{
	"ASC":  ListCloudAsmDiskGroupsSortOrderAsc,
	"DESC": ListCloudAsmDiskGroupsSortOrderDesc,
}

var mappingListCloudAsmDiskGroupsSortOrderEnumLowerCase = map[string]ListCloudAsmDiskGroupsSortOrderEnum{
	"asc":  ListCloudAsmDiskGroupsSortOrderAsc,
	"desc": ListCloudAsmDiskGroupsSortOrderDesc,
}

// GetListCloudAsmDiskGroupsSortOrderEnumValues Enumerates the set of values for ListCloudAsmDiskGroupsSortOrderEnum
func GetListCloudAsmDiskGroupsSortOrderEnumValues() []ListCloudAsmDiskGroupsSortOrderEnum {
	values := make([]ListCloudAsmDiskGroupsSortOrderEnum, 0)
	for _, v := range mappingListCloudAsmDiskGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudAsmDiskGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListCloudAsmDiskGroupsSortOrderEnum
func GetListCloudAsmDiskGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudAsmDiskGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudAsmDiskGroupsSortOrderEnum(val string) (ListCloudAsmDiskGroupsSortOrderEnum, bool) {
	enum, ok := mappingListCloudAsmDiskGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
