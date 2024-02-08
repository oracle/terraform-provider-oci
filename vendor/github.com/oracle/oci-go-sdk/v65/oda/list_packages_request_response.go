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

// ListPackagesRequest wrapper for the ListPackages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListPackages.go.html to see an example of how to use ListPackagesRequest.
type ListPackagesRequest struct {

	// List only the information for this Digital Assistant instance.
	OdaInstanceId *string `mandatory:"false" contributesTo:"query" name:"odaInstanceId"`

	// Resource type identifier. Used to limit query results to the items which are applicable to the given type.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// List the packages that belong to this compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// List only the information for the package with this name. Package names are unique to a publisher and may not change.
	// Example: `My Package`
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// List only the information for the Digital Assistant instance with this user-friendly name. These names don't have to be unique and may change.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Should we return only the latest version of a package (instead of all versions)?
	IsLatestVersionOnly *bool `mandatory:"false" contributesTo:"query" name:"isLatestVersionOnly"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListPackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `TIMECREATED`.
	// The default sort order for `TIMECREATED` is descending, and the default sort order for `DISPLAYNAME` is ascending.
	SortBy ListPackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPackagesResponse wrapper for the ListPackages operation
type ListPackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []PackageSummary instances
	Items []PackageSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// When you are paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the
	// `page` query parameter for the subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of results that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListPackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPackagesSortOrderEnum Enum with underlying type: string
type ListPackagesSortOrderEnum string

// Set of constants representing the allowable values for ListPackagesSortOrderEnum
const (
	ListPackagesSortOrderAsc  ListPackagesSortOrderEnum = "ASC"
	ListPackagesSortOrderDesc ListPackagesSortOrderEnum = "DESC"
)

var mappingListPackagesSortOrderEnum = map[string]ListPackagesSortOrderEnum{
	"ASC":  ListPackagesSortOrderAsc,
	"DESC": ListPackagesSortOrderDesc,
}

var mappingListPackagesSortOrderEnumLowerCase = map[string]ListPackagesSortOrderEnum{
	"asc":  ListPackagesSortOrderAsc,
	"desc": ListPackagesSortOrderDesc,
}

// GetListPackagesSortOrderEnumValues Enumerates the set of values for ListPackagesSortOrderEnum
func GetListPackagesSortOrderEnumValues() []ListPackagesSortOrderEnum {
	values := make([]ListPackagesSortOrderEnum, 0)
	for _, v := range mappingListPackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPackagesSortOrderEnumStringValues Enumerates the set of values in String for ListPackagesSortOrderEnum
func GetListPackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPackagesSortOrderEnum(val string) (ListPackagesSortOrderEnum, bool) {
	enum, ok := mappingListPackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPackagesSortByEnum Enum with underlying type: string
type ListPackagesSortByEnum string

// Set of constants representing the allowable values for ListPackagesSortByEnum
const (
	ListPackagesSortByTimecreated ListPackagesSortByEnum = "TIMECREATED"
	ListPackagesSortByDisplayname ListPackagesSortByEnum = "DISPLAYNAME"
)

var mappingListPackagesSortByEnum = map[string]ListPackagesSortByEnum{
	"TIMECREATED": ListPackagesSortByTimecreated,
	"DISPLAYNAME": ListPackagesSortByDisplayname,
}

var mappingListPackagesSortByEnumLowerCase = map[string]ListPackagesSortByEnum{
	"timecreated": ListPackagesSortByTimecreated,
	"displayname": ListPackagesSortByDisplayname,
}

// GetListPackagesSortByEnumValues Enumerates the set of values for ListPackagesSortByEnum
func GetListPackagesSortByEnumValues() []ListPackagesSortByEnum {
	values := make([]ListPackagesSortByEnum, 0)
	for _, v := range mappingListPackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPackagesSortByEnumStringValues Enumerates the set of values in String for ListPackagesSortByEnum
func GetListPackagesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListPackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPackagesSortByEnum(val string) (ListPackagesSortByEnum, bool) {
	enum, ok := mappingListPackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
