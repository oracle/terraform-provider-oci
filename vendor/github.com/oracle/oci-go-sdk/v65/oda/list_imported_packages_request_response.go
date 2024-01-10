// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListImportedPackagesRequest wrapper for the ListImportedPackages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListImportedPackages.go.html to see an example of how to use ListImportedPackagesRequest.
type ListImportedPackagesRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// List only the information for the package with this name. Package names are unique to a publisher and may not change.
	// Example: `My Package`
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListImportedPackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `TIMECREATED`.
	// The default sort order for `TIMECREATED` is descending, and the default sort order for `DISPLAYNAME` is ascending.
	SortBy ListImportedPackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListImportedPackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListImportedPackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListImportedPackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListImportedPackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListImportedPackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListImportedPackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListImportedPackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListImportedPackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListImportedPackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListImportedPackagesResponse wrapper for the ListImportedPackages operation
type ListImportedPackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ImportedPackageSummary instances
	Items []ImportedPackageSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// When you are paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the
	// `page` query parameter for the subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListImportedPackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListImportedPackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListImportedPackagesSortOrderEnum Enum with underlying type: string
type ListImportedPackagesSortOrderEnum string

// Set of constants representing the allowable values for ListImportedPackagesSortOrderEnum
const (
	ListImportedPackagesSortOrderAsc  ListImportedPackagesSortOrderEnum = "ASC"
	ListImportedPackagesSortOrderDesc ListImportedPackagesSortOrderEnum = "DESC"
)

var mappingListImportedPackagesSortOrderEnum = map[string]ListImportedPackagesSortOrderEnum{
	"ASC":  ListImportedPackagesSortOrderAsc,
	"DESC": ListImportedPackagesSortOrderDesc,
}

var mappingListImportedPackagesSortOrderEnumLowerCase = map[string]ListImportedPackagesSortOrderEnum{
	"asc":  ListImportedPackagesSortOrderAsc,
	"desc": ListImportedPackagesSortOrderDesc,
}

// GetListImportedPackagesSortOrderEnumValues Enumerates the set of values for ListImportedPackagesSortOrderEnum
func GetListImportedPackagesSortOrderEnumValues() []ListImportedPackagesSortOrderEnum {
	values := make([]ListImportedPackagesSortOrderEnum, 0)
	for _, v := range mappingListImportedPackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListImportedPackagesSortOrderEnumStringValues Enumerates the set of values in String for ListImportedPackagesSortOrderEnum
func GetListImportedPackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListImportedPackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListImportedPackagesSortOrderEnum(val string) (ListImportedPackagesSortOrderEnum, bool) {
	enum, ok := mappingListImportedPackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListImportedPackagesSortByEnum Enum with underlying type: string
type ListImportedPackagesSortByEnum string

// Set of constants representing the allowable values for ListImportedPackagesSortByEnum
const (
	ListImportedPackagesSortByTimecreated ListImportedPackagesSortByEnum = "TIMECREATED"
	ListImportedPackagesSortByDisplayname ListImportedPackagesSortByEnum = "DISPLAYNAME"
)

var mappingListImportedPackagesSortByEnum = map[string]ListImportedPackagesSortByEnum{
	"TIMECREATED": ListImportedPackagesSortByTimecreated,
	"DISPLAYNAME": ListImportedPackagesSortByDisplayname,
}

var mappingListImportedPackagesSortByEnumLowerCase = map[string]ListImportedPackagesSortByEnum{
	"timecreated": ListImportedPackagesSortByTimecreated,
	"displayname": ListImportedPackagesSortByDisplayname,
}

// GetListImportedPackagesSortByEnumValues Enumerates the set of values for ListImportedPackagesSortByEnum
func GetListImportedPackagesSortByEnumValues() []ListImportedPackagesSortByEnum {
	values := make([]ListImportedPackagesSortByEnum, 0)
	for _, v := range mappingListImportedPackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListImportedPackagesSortByEnumStringValues Enumerates the set of values in String for ListImportedPackagesSortByEnum
func GetListImportedPackagesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListImportedPackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListImportedPackagesSortByEnum(val string) (ListImportedPackagesSortByEnum, bool) {
	enum, ok := mappingListImportedPackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
