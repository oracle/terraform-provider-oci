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

// ListProductLicensesRequest wrapper for the ListProductLicenses operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/ListProductLicenses.go.html to see an example of how to use ListProductLicensesRequest.
type ListProductLicensesRequest struct {

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
	SortOrder ListProductLicensesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the attribute with which to sort the rules.
	// Default: `totalLicenseUnitsConsumed`
	// * **totalLicenseUnitsConsumed:** Sorts by totalLicenseUnitsConsumed of ProductLicense.
	SortBy ListProductLicensesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProductLicensesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProductLicensesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProductLicensesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProductLicensesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProductLicensesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProductLicensesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProductLicensesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProductLicensesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProductLicensesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProductLicensesResponse wrapper for the ListProductLicenses operation
type ListProductLicensesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProductLicenseCollection instances
	ProductLicenseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProductLicensesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProductLicensesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProductLicensesSortOrderEnum Enum with underlying type: string
type ListProductLicensesSortOrderEnum string

// Set of constants representing the allowable values for ListProductLicensesSortOrderEnum
const (
	ListProductLicensesSortOrderAsc  ListProductLicensesSortOrderEnum = "ASC"
	ListProductLicensesSortOrderDesc ListProductLicensesSortOrderEnum = "DESC"
)

var mappingListProductLicensesSortOrderEnum = map[string]ListProductLicensesSortOrderEnum{
	"ASC":  ListProductLicensesSortOrderAsc,
	"DESC": ListProductLicensesSortOrderDesc,
}

var mappingListProductLicensesSortOrderEnumLowerCase = map[string]ListProductLicensesSortOrderEnum{
	"asc":  ListProductLicensesSortOrderAsc,
	"desc": ListProductLicensesSortOrderDesc,
}

// GetListProductLicensesSortOrderEnumValues Enumerates the set of values for ListProductLicensesSortOrderEnum
func GetListProductLicensesSortOrderEnumValues() []ListProductLicensesSortOrderEnum {
	values := make([]ListProductLicensesSortOrderEnum, 0)
	for _, v := range mappingListProductLicensesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProductLicensesSortOrderEnumStringValues Enumerates the set of values in String for ListProductLicensesSortOrderEnum
func GetListProductLicensesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProductLicensesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProductLicensesSortOrderEnum(val string) (ListProductLicensesSortOrderEnum, bool) {
	enum, ok := mappingListProductLicensesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProductLicensesSortByEnum Enum with underlying type: string
type ListProductLicensesSortByEnum string

// Set of constants representing the allowable values for ListProductLicensesSortByEnum
const (
	ListProductLicensesSortByTotallicenseunitsconsumed ListProductLicensesSortByEnum = "totalLicenseUnitsConsumed"
)

var mappingListProductLicensesSortByEnum = map[string]ListProductLicensesSortByEnum{
	"totalLicenseUnitsConsumed": ListProductLicensesSortByTotallicenseunitsconsumed,
}

var mappingListProductLicensesSortByEnumLowerCase = map[string]ListProductLicensesSortByEnum{
	"totallicenseunitsconsumed": ListProductLicensesSortByTotallicenseunitsconsumed,
}

// GetListProductLicensesSortByEnumValues Enumerates the set of values for ListProductLicensesSortByEnum
func GetListProductLicensesSortByEnumValues() []ListProductLicensesSortByEnum {
	values := make([]ListProductLicensesSortByEnum, 0)
	for _, v := range mappingListProductLicensesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProductLicensesSortByEnumStringValues Enumerates the set of values in String for ListProductLicensesSortByEnum
func GetListProductLicensesSortByEnumStringValues() []string {
	return []string{
		"totalLicenseUnitsConsumed",
	}
}

// GetMappingListProductLicensesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProductLicensesSortByEnum(val string) (ListProductLicensesSortByEnum, bool) {
	enum, ok := mappingListProductLicensesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
