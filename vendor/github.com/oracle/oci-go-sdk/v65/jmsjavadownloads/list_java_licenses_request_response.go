// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jmsjavadownloads

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJavaLicensesRequest wrapper for the ListJavaLicenses operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListJavaLicenses.go.html to see an example of how to use ListJavaLicensesRequest.
type ListJavaLicensesRequest struct {

	// Unique Java license type.
	LicenseType ListJavaLicensesLicenseTypeEnum `mandatory:"false" contributesTo:"query" name:"licenseType" omitEmpty:"true"`

	// A filter to return only resources that match the display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListJavaLicensesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. If no value is specified, _licenseType_ is the default.
	SortBy ListJavaLicensesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJavaLicensesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJavaLicensesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJavaLicensesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJavaLicensesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJavaLicensesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJavaLicensesLicenseTypeEnum(string(request.LicenseType)); !ok && request.LicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseType: %s. Supported values are: %s.", request.LicenseType, strings.Join(GetListJavaLicensesLicenseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaLicensesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJavaLicensesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaLicensesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJavaLicensesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJavaLicensesResponse wrapper for the ListJavaLicenses operation
type ListJavaLicensesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JavaLicenseCollection instances
	JavaLicenseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJavaLicensesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJavaLicensesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJavaLicensesLicenseTypeEnum Enum with underlying type: string
type ListJavaLicensesLicenseTypeEnum string

// Set of constants representing the allowable values for ListJavaLicensesLicenseTypeEnum
const (
	ListJavaLicensesLicenseTypeOtn        ListJavaLicensesLicenseTypeEnum = "OTN"
	ListJavaLicensesLicenseTypeNftc       ListJavaLicensesLicenseTypeEnum = "NFTC"
	ListJavaLicensesLicenseTypeRestricted ListJavaLicensesLicenseTypeEnum = "RESTRICTED"
)

var mappingListJavaLicensesLicenseTypeEnum = map[string]ListJavaLicensesLicenseTypeEnum{
	"OTN":        ListJavaLicensesLicenseTypeOtn,
	"NFTC":       ListJavaLicensesLicenseTypeNftc,
	"RESTRICTED": ListJavaLicensesLicenseTypeRestricted,
}

var mappingListJavaLicensesLicenseTypeEnumLowerCase = map[string]ListJavaLicensesLicenseTypeEnum{
	"otn":        ListJavaLicensesLicenseTypeOtn,
	"nftc":       ListJavaLicensesLicenseTypeNftc,
	"restricted": ListJavaLicensesLicenseTypeRestricted,
}

// GetListJavaLicensesLicenseTypeEnumValues Enumerates the set of values for ListJavaLicensesLicenseTypeEnum
func GetListJavaLicensesLicenseTypeEnumValues() []ListJavaLicensesLicenseTypeEnum {
	values := make([]ListJavaLicensesLicenseTypeEnum, 0)
	for _, v := range mappingListJavaLicensesLicenseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaLicensesLicenseTypeEnumStringValues Enumerates the set of values in String for ListJavaLicensesLicenseTypeEnum
func GetListJavaLicensesLicenseTypeEnumStringValues() []string {
	return []string{
		"OTN",
		"NFTC",
		"RESTRICTED",
	}
}

// GetMappingListJavaLicensesLicenseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaLicensesLicenseTypeEnum(val string) (ListJavaLicensesLicenseTypeEnum, bool) {
	enum, ok := mappingListJavaLicensesLicenseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaLicensesSortOrderEnum Enum with underlying type: string
type ListJavaLicensesSortOrderEnum string

// Set of constants representing the allowable values for ListJavaLicensesSortOrderEnum
const (
	ListJavaLicensesSortOrderAsc  ListJavaLicensesSortOrderEnum = "ASC"
	ListJavaLicensesSortOrderDesc ListJavaLicensesSortOrderEnum = "DESC"
)

var mappingListJavaLicensesSortOrderEnum = map[string]ListJavaLicensesSortOrderEnum{
	"ASC":  ListJavaLicensesSortOrderAsc,
	"DESC": ListJavaLicensesSortOrderDesc,
}

var mappingListJavaLicensesSortOrderEnumLowerCase = map[string]ListJavaLicensesSortOrderEnum{
	"asc":  ListJavaLicensesSortOrderAsc,
	"desc": ListJavaLicensesSortOrderDesc,
}

// GetListJavaLicensesSortOrderEnumValues Enumerates the set of values for ListJavaLicensesSortOrderEnum
func GetListJavaLicensesSortOrderEnumValues() []ListJavaLicensesSortOrderEnum {
	values := make([]ListJavaLicensesSortOrderEnum, 0)
	for _, v := range mappingListJavaLicensesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaLicensesSortOrderEnumStringValues Enumerates the set of values in String for ListJavaLicensesSortOrderEnum
func GetListJavaLicensesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJavaLicensesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaLicensesSortOrderEnum(val string) (ListJavaLicensesSortOrderEnum, bool) {
	enum, ok := mappingListJavaLicensesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaLicensesSortByEnum Enum with underlying type: string
type ListJavaLicensesSortByEnum string

// Set of constants representing the allowable values for ListJavaLicensesSortByEnum
const (
	ListJavaLicensesSortByLicensetype ListJavaLicensesSortByEnum = "licenseType"
	ListJavaLicensesSortByDisplayname ListJavaLicensesSortByEnum = "displayName"
)

var mappingListJavaLicensesSortByEnum = map[string]ListJavaLicensesSortByEnum{
	"licenseType": ListJavaLicensesSortByLicensetype,
	"displayName": ListJavaLicensesSortByDisplayname,
}

var mappingListJavaLicensesSortByEnumLowerCase = map[string]ListJavaLicensesSortByEnum{
	"licensetype": ListJavaLicensesSortByLicensetype,
	"displayname": ListJavaLicensesSortByDisplayname,
}

// GetListJavaLicensesSortByEnumValues Enumerates the set of values for ListJavaLicensesSortByEnum
func GetListJavaLicensesSortByEnumValues() []ListJavaLicensesSortByEnum {
	values := make([]ListJavaLicensesSortByEnum, 0)
	for _, v := range mappingListJavaLicensesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaLicensesSortByEnumStringValues Enumerates the set of values in String for ListJavaLicensesSortByEnum
func GetListJavaLicensesSortByEnumStringValues() []string {
	return []string{
		"licenseType",
		"displayName",
	}
}

// GetMappingListJavaLicensesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaLicensesSortByEnum(val string) (ListJavaLicensesSortByEnum, bool) {
	enum, ok := mappingListJavaLicensesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
