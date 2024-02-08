// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package licensemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTopUtilizedProductLicensesRequest wrapper for the ListTopUtilizedProductLicenses operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/ListTopUtilizedProductLicenses.go.html to see an example of how to use ListTopUtilizedProductLicensesRequest.
type ListTopUtilizedProductLicensesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) used for the license record, product license, and configuration.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Indicates if the given compartment is the root compartment.
	IsCompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"isCompartmentIdInSubtree"`

	// The sort order to use, whether `ASC` or `DESC`.
	SortOrder ListTopUtilizedProductLicensesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the attribute with which to sort the rules.
	// Default: `totalLicenseUnitsConsumed`
	// * **totalLicenseUnitsConsumed:** Sorts by totalLicenseUnitsConsumed of ProductLicense.
	SortBy ListTopUtilizedProductLicensesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTopUtilizedProductLicensesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTopUtilizedProductLicensesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTopUtilizedProductLicensesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTopUtilizedProductLicensesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTopUtilizedProductLicensesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTopUtilizedProductLicensesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTopUtilizedProductLicensesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTopUtilizedProductLicensesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTopUtilizedProductLicensesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTopUtilizedProductLicensesResponse wrapper for the ListTopUtilizedProductLicenses operation
type ListTopUtilizedProductLicensesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TopUtilizedProductLicenseCollection instances
	TopUtilizedProductLicenseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTopUtilizedProductLicensesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTopUtilizedProductLicensesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTopUtilizedProductLicensesSortOrderEnum Enum with underlying type: string
type ListTopUtilizedProductLicensesSortOrderEnum string

// Set of constants representing the allowable values for ListTopUtilizedProductLicensesSortOrderEnum
const (
	ListTopUtilizedProductLicensesSortOrderAsc  ListTopUtilizedProductLicensesSortOrderEnum = "ASC"
	ListTopUtilizedProductLicensesSortOrderDesc ListTopUtilizedProductLicensesSortOrderEnum = "DESC"
)

var mappingListTopUtilizedProductLicensesSortOrderEnum = map[string]ListTopUtilizedProductLicensesSortOrderEnum{
	"ASC":  ListTopUtilizedProductLicensesSortOrderAsc,
	"DESC": ListTopUtilizedProductLicensesSortOrderDesc,
}

var mappingListTopUtilizedProductLicensesSortOrderEnumLowerCase = map[string]ListTopUtilizedProductLicensesSortOrderEnum{
	"asc":  ListTopUtilizedProductLicensesSortOrderAsc,
	"desc": ListTopUtilizedProductLicensesSortOrderDesc,
}

// GetListTopUtilizedProductLicensesSortOrderEnumValues Enumerates the set of values for ListTopUtilizedProductLicensesSortOrderEnum
func GetListTopUtilizedProductLicensesSortOrderEnumValues() []ListTopUtilizedProductLicensesSortOrderEnum {
	values := make([]ListTopUtilizedProductLicensesSortOrderEnum, 0)
	for _, v := range mappingListTopUtilizedProductLicensesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTopUtilizedProductLicensesSortOrderEnumStringValues Enumerates the set of values in String for ListTopUtilizedProductLicensesSortOrderEnum
func GetListTopUtilizedProductLicensesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTopUtilizedProductLicensesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTopUtilizedProductLicensesSortOrderEnum(val string) (ListTopUtilizedProductLicensesSortOrderEnum, bool) {
	enum, ok := mappingListTopUtilizedProductLicensesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTopUtilizedProductLicensesSortByEnum Enum with underlying type: string
type ListTopUtilizedProductLicensesSortByEnum string

// Set of constants representing the allowable values for ListTopUtilizedProductLicensesSortByEnum
const (
	ListTopUtilizedProductLicensesSortByTotallicenseunitsconsumed ListTopUtilizedProductLicensesSortByEnum = "totalLicenseUnitsConsumed"
)

var mappingListTopUtilizedProductLicensesSortByEnum = map[string]ListTopUtilizedProductLicensesSortByEnum{
	"totalLicenseUnitsConsumed": ListTopUtilizedProductLicensesSortByTotallicenseunitsconsumed,
}

var mappingListTopUtilizedProductLicensesSortByEnumLowerCase = map[string]ListTopUtilizedProductLicensesSortByEnum{
	"totallicenseunitsconsumed": ListTopUtilizedProductLicensesSortByTotallicenseunitsconsumed,
}

// GetListTopUtilizedProductLicensesSortByEnumValues Enumerates the set of values for ListTopUtilizedProductLicensesSortByEnum
func GetListTopUtilizedProductLicensesSortByEnumValues() []ListTopUtilizedProductLicensesSortByEnum {
	values := make([]ListTopUtilizedProductLicensesSortByEnum, 0)
	for _, v := range mappingListTopUtilizedProductLicensesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTopUtilizedProductLicensesSortByEnumStringValues Enumerates the set of values in String for ListTopUtilizedProductLicensesSortByEnum
func GetListTopUtilizedProductLicensesSortByEnumStringValues() []string {
	return []string{
		"totalLicenseUnitsConsumed",
	}
}

// GetMappingListTopUtilizedProductLicensesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTopUtilizedProductLicensesSortByEnum(val string) (ListTopUtilizedProductLicensesSortByEnum, bool) {
	enum, ok := mappingListTopUtilizedProductLicensesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
