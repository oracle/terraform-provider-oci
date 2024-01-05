// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package secrets

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSecretBundleVersionsRequest wrapper for the ListSecretBundleVersions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/secrets/ListSecretBundleVersions.go.html to see an example of how to use ListSecretBundleVersionsRequest.
type ListSecretBundleVersionsRequest struct {

	// The OCID of the secret.
	SecretId *string `mandatory:"true" contributesTo:"path" name:"secretId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated "List" call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call. For information about
	// pagination, see List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can specify only one sort order. The default
	// order for `VERSION_NUMBER` is descending.
	SortBy ListSecretBundleVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSecretBundleVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecretBundleVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecretBundleVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecretBundleVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecretBundleVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecretBundleVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSecretBundleVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSecretBundleVersionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecretBundleVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSecretBundleVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecretBundleVersionsResponse wrapper for the ListSecretBundleVersions operation
type ListSecretBundleVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SecretBundleVersionSummary instances
	Items []SecretBundleVersionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSecretBundleVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecretBundleVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecretBundleVersionsSortByEnum Enum with underlying type: string
type ListSecretBundleVersionsSortByEnum string

// Set of constants representing the allowable values for ListSecretBundleVersionsSortByEnum
const (
	ListSecretBundleVersionsSortByVersionNumber ListSecretBundleVersionsSortByEnum = "VERSION_NUMBER"
)

var mappingListSecretBundleVersionsSortByEnum = map[string]ListSecretBundleVersionsSortByEnum{
	"VERSION_NUMBER": ListSecretBundleVersionsSortByVersionNumber,
}

var mappingListSecretBundleVersionsSortByEnumLowerCase = map[string]ListSecretBundleVersionsSortByEnum{
	"version_number": ListSecretBundleVersionsSortByVersionNumber,
}

// GetListSecretBundleVersionsSortByEnumValues Enumerates the set of values for ListSecretBundleVersionsSortByEnum
func GetListSecretBundleVersionsSortByEnumValues() []ListSecretBundleVersionsSortByEnum {
	values := make([]ListSecretBundleVersionsSortByEnum, 0)
	for _, v := range mappingListSecretBundleVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecretBundleVersionsSortByEnumStringValues Enumerates the set of values in String for ListSecretBundleVersionsSortByEnum
func GetListSecretBundleVersionsSortByEnumStringValues() []string {
	return []string{
		"VERSION_NUMBER",
	}
}

// GetMappingListSecretBundleVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecretBundleVersionsSortByEnum(val string) (ListSecretBundleVersionsSortByEnum, bool) {
	enum, ok := mappingListSecretBundleVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecretBundleVersionsSortOrderEnum Enum with underlying type: string
type ListSecretBundleVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListSecretBundleVersionsSortOrderEnum
const (
	ListSecretBundleVersionsSortOrderAsc  ListSecretBundleVersionsSortOrderEnum = "ASC"
	ListSecretBundleVersionsSortOrderDesc ListSecretBundleVersionsSortOrderEnum = "DESC"
)

var mappingListSecretBundleVersionsSortOrderEnum = map[string]ListSecretBundleVersionsSortOrderEnum{
	"ASC":  ListSecretBundleVersionsSortOrderAsc,
	"DESC": ListSecretBundleVersionsSortOrderDesc,
}

var mappingListSecretBundleVersionsSortOrderEnumLowerCase = map[string]ListSecretBundleVersionsSortOrderEnum{
	"asc":  ListSecretBundleVersionsSortOrderAsc,
	"desc": ListSecretBundleVersionsSortOrderDesc,
}

// GetListSecretBundleVersionsSortOrderEnumValues Enumerates the set of values for ListSecretBundleVersionsSortOrderEnum
func GetListSecretBundleVersionsSortOrderEnumValues() []ListSecretBundleVersionsSortOrderEnum {
	values := make([]ListSecretBundleVersionsSortOrderEnum, 0)
	for _, v := range mappingListSecretBundleVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecretBundleVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListSecretBundleVersionsSortOrderEnum
func GetListSecretBundleVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSecretBundleVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecretBundleVersionsSortOrderEnum(val string) (ListSecretBundleVersionsSortOrderEnum, bool) {
	enum, ok := mappingListSecretBundleVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
