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

// ListCloudAsmUsersRequest wrapper for the ListCloudAsmUsers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudAsmUsers.go.html to see an example of how to use ListCloudAsmUsersRequest.
type ListCloudAsmUsersRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM.
	CloudAsmId *string `mandatory:"true" contributesTo:"path" name:"cloudAsmId"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The
	// default sort order for `NAME` is ascending and it is case-sensitive.
	SortBy ListCloudAsmUsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudAsmUsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudAsmUsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudAsmUsersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudAsmUsersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudAsmUsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudAsmUsersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudAsmUsersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudAsmUsersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudAsmUsersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudAsmUsersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudAsmUsersResponse wrapper for the ListCloudAsmUsers operation
type ListCloudAsmUsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudAsmUserCollection instances
	CloudAsmUserCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudAsmUsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudAsmUsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudAsmUsersSortByEnum Enum with underlying type: string
type ListCloudAsmUsersSortByEnum string

// Set of constants representing the allowable values for ListCloudAsmUsersSortByEnum
const (
	ListCloudAsmUsersSortByName ListCloudAsmUsersSortByEnum = "NAME"
)

var mappingListCloudAsmUsersSortByEnum = map[string]ListCloudAsmUsersSortByEnum{
	"NAME": ListCloudAsmUsersSortByName,
}

var mappingListCloudAsmUsersSortByEnumLowerCase = map[string]ListCloudAsmUsersSortByEnum{
	"name": ListCloudAsmUsersSortByName,
}

// GetListCloudAsmUsersSortByEnumValues Enumerates the set of values for ListCloudAsmUsersSortByEnum
func GetListCloudAsmUsersSortByEnumValues() []ListCloudAsmUsersSortByEnum {
	values := make([]ListCloudAsmUsersSortByEnum, 0)
	for _, v := range mappingListCloudAsmUsersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudAsmUsersSortByEnumStringValues Enumerates the set of values in String for ListCloudAsmUsersSortByEnum
func GetListCloudAsmUsersSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListCloudAsmUsersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudAsmUsersSortByEnum(val string) (ListCloudAsmUsersSortByEnum, bool) {
	enum, ok := mappingListCloudAsmUsersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudAsmUsersSortOrderEnum Enum with underlying type: string
type ListCloudAsmUsersSortOrderEnum string

// Set of constants representing the allowable values for ListCloudAsmUsersSortOrderEnum
const (
	ListCloudAsmUsersSortOrderAsc  ListCloudAsmUsersSortOrderEnum = "ASC"
	ListCloudAsmUsersSortOrderDesc ListCloudAsmUsersSortOrderEnum = "DESC"
)

var mappingListCloudAsmUsersSortOrderEnum = map[string]ListCloudAsmUsersSortOrderEnum{
	"ASC":  ListCloudAsmUsersSortOrderAsc,
	"DESC": ListCloudAsmUsersSortOrderDesc,
}

var mappingListCloudAsmUsersSortOrderEnumLowerCase = map[string]ListCloudAsmUsersSortOrderEnum{
	"asc":  ListCloudAsmUsersSortOrderAsc,
	"desc": ListCloudAsmUsersSortOrderDesc,
}

// GetListCloudAsmUsersSortOrderEnumValues Enumerates the set of values for ListCloudAsmUsersSortOrderEnum
func GetListCloudAsmUsersSortOrderEnumValues() []ListCloudAsmUsersSortOrderEnum {
	values := make([]ListCloudAsmUsersSortOrderEnum, 0)
	for _, v := range mappingListCloudAsmUsersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudAsmUsersSortOrderEnumStringValues Enumerates the set of values in String for ListCloudAsmUsersSortOrderEnum
func GetListCloudAsmUsersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudAsmUsersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudAsmUsersSortOrderEnum(val string) (ListCloudAsmUsersSortOrderEnum, bool) {
	enum, ok := mappingListCloudAsmUsersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
