// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicecatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPrivateApplicationPackagesRequest wrapper for the ListPrivateApplicationPackages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListPrivateApplicationPackages.go.html to see an example of how to use ListPrivateApplicationPackagesRequest.
type ListPrivateApplicationPackagesRequest struct {

	// The unique identifier for the private application.
	PrivateApplicationId *string `mandatory:"true" contributesTo:"query" name:"privateApplicationId"`

	// The unique identifier for the private application package.
	PrivateApplicationPackageId *string `mandatory:"false" contributesTo:"query" name:"privateApplicationPackageId"`

	// Name of the package type. If multiple package types are provided, then any resource with
	// one or more matching package types will be returned.
	PackageType []PackageTypeEnumEnum `contributesTo:"query" name:"packageType" omitEmpty:"true" collectionFormat:"multi"`

	// How many records to return. Specify a value greater than zero and less than or equal to 1000. The default is 30.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to use to sort listed results. You can only specify one field to sort by.
	// `TIMECREATED` displays results in descending order by default. You can change your
	// preference by specifying a different sort order.
	SortBy ListPrivateApplicationPackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to apply, either `ASC` or `DESC`. Default is `ASC`.
	SortOrder ListPrivateApplicationPackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Exact match name filter.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPrivateApplicationPackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPrivateApplicationPackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPrivateApplicationPackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPrivateApplicationPackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPrivateApplicationPackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.PackageType {
		if _, ok := GetMappingPackageTypeEnumEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", val, strings.Join(GetPackageTypeEnumEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListPrivateApplicationPackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPrivateApplicationPackagesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPrivateApplicationPackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPrivateApplicationPackagesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPrivateApplicationPackagesResponse wrapper for the ListPrivateApplicationPackages operation
type ListPrivateApplicationPackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PrivateApplicationPackageCollection instances
	PrivateApplicationPackageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPrivateApplicationPackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPrivateApplicationPackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPrivateApplicationPackagesSortByEnum Enum with underlying type: string
type ListPrivateApplicationPackagesSortByEnum string

// Set of constants representing the allowable values for ListPrivateApplicationPackagesSortByEnum
const (
	ListPrivateApplicationPackagesSortByTimecreated ListPrivateApplicationPackagesSortByEnum = "TIMECREATED"
	ListPrivateApplicationPackagesSortByVersion     ListPrivateApplicationPackagesSortByEnum = "VERSION"
)

var mappingListPrivateApplicationPackagesSortByEnum = map[string]ListPrivateApplicationPackagesSortByEnum{
	"TIMECREATED": ListPrivateApplicationPackagesSortByTimecreated,
	"VERSION":     ListPrivateApplicationPackagesSortByVersion,
}

var mappingListPrivateApplicationPackagesSortByEnumLowerCase = map[string]ListPrivateApplicationPackagesSortByEnum{
	"timecreated": ListPrivateApplicationPackagesSortByTimecreated,
	"version":     ListPrivateApplicationPackagesSortByVersion,
}

// GetListPrivateApplicationPackagesSortByEnumValues Enumerates the set of values for ListPrivateApplicationPackagesSortByEnum
func GetListPrivateApplicationPackagesSortByEnumValues() []ListPrivateApplicationPackagesSortByEnum {
	values := make([]ListPrivateApplicationPackagesSortByEnum, 0)
	for _, v := range mappingListPrivateApplicationPackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivateApplicationPackagesSortByEnumStringValues Enumerates the set of values in String for ListPrivateApplicationPackagesSortByEnum
func GetListPrivateApplicationPackagesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"VERSION",
	}
}

// GetMappingListPrivateApplicationPackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivateApplicationPackagesSortByEnum(val string) (ListPrivateApplicationPackagesSortByEnum, bool) {
	enum, ok := mappingListPrivateApplicationPackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPrivateApplicationPackagesSortOrderEnum Enum with underlying type: string
type ListPrivateApplicationPackagesSortOrderEnum string

// Set of constants representing the allowable values for ListPrivateApplicationPackagesSortOrderEnum
const (
	ListPrivateApplicationPackagesSortOrderAsc  ListPrivateApplicationPackagesSortOrderEnum = "ASC"
	ListPrivateApplicationPackagesSortOrderDesc ListPrivateApplicationPackagesSortOrderEnum = "DESC"
)

var mappingListPrivateApplicationPackagesSortOrderEnum = map[string]ListPrivateApplicationPackagesSortOrderEnum{
	"ASC":  ListPrivateApplicationPackagesSortOrderAsc,
	"DESC": ListPrivateApplicationPackagesSortOrderDesc,
}

var mappingListPrivateApplicationPackagesSortOrderEnumLowerCase = map[string]ListPrivateApplicationPackagesSortOrderEnum{
	"asc":  ListPrivateApplicationPackagesSortOrderAsc,
	"desc": ListPrivateApplicationPackagesSortOrderDesc,
}

// GetListPrivateApplicationPackagesSortOrderEnumValues Enumerates the set of values for ListPrivateApplicationPackagesSortOrderEnum
func GetListPrivateApplicationPackagesSortOrderEnumValues() []ListPrivateApplicationPackagesSortOrderEnum {
	values := make([]ListPrivateApplicationPackagesSortOrderEnum, 0)
	for _, v := range mappingListPrivateApplicationPackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivateApplicationPackagesSortOrderEnumStringValues Enumerates the set of values in String for ListPrivateApplicationPackagesSortOrderEnum
func GetListPrivateApplicationPackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPrivateApplicationPackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivateApplicationPackagesSortOrderEnum(val string) (ListPrivateApplicationPackagesSortOrderEnum, bool) {
	enum, ok := mappingListPrivateApplicationPackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
