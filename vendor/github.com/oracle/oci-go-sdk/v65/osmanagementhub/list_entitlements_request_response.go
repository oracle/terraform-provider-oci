// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEntitlementsRequest wrapper for the ListEntitlements operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListEntitlements.go.html to see an example of how to use ListEntitlementsRequest.
type ListEntitlementsRequest struct {

	// The OCID of the compartment that contains the resources to list. This parameter is required.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return entitlements that match the given CSI.
	Csi *string `mandatory:"false" contributesTo:"query" name:"csi"`

	// A filter to return only profiles that match the given vendorName.
	VendorName ListEntitlementsVendorNameEnum `mandatory:"false" contributesTo:"query" name:"vendorName" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListEntitlementsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort entitlements by. Only one sort order may be provided.
	SortBy ListEntitlementsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEntitlementsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEntitlementsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEntitlementsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEntitlementsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEntitlementsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEntitlementsVendorNameEnum(string(request.VendorName)); !ok && request.VendorName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VendorName: %s. Supported values are: %s.", request.VendorName, strings.Join(GetListEntitlementsVendorNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEntitlementsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEntitlementsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEntitlementsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEntitlementsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEntitlementsResponse wrapper for the ListEntitlements operation
type ListEntitlementsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EntitlementCollection instances
	EntitlementCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEntitlementsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEntitlementsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEntitlementsVendorNameEnum Enum with underlying type: string
type ListEntitlementsVendorNameEnum string

// Set of constants representing the allowable values for ListEntitlementsVendorNameEnum
const (
	ListEntitlementsVendorNameOracle ListEntitlementsVendorNameEnum = "ORACLE"
)

var mappingListEntitlementsVendorNameEnum = map[string]ListEntitlementsVendorNameEnum{
	"ORACLE": ListEntitlementsVendorNameOracle,
}

var mappingListEntitlementsVendorNameEnumLowerCase = map[string]ListEntitlementsVendorNameEnum{
	"oracle": ListEntitlementsVendorNameOracle,
}

// GetListEntitlementsVendorNameEnumValues Enumerates the set of values for ListEntitlementsVendorNameEnum
func GetListEntitlementsVendorNameEnumValues() []ListEntitlementsVendorNameEnum {
	values := make([]ListEntitlementsVendorNameEnum, 0)
	for _, v := range mappingListEntitlementsVendorNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntitlementsVendorNameEnumStringValues Enumerates the set of values in String for ListEntitlementsVendorNameEnum
func GetListEntitlementsVendorNameEnumStringValues() []string {
	return []string{
		"ORACLE",
	}
}

// GetMappingListEntitlementsVendorNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntitlementsVendorNameEnum(val string) (ListEntitlementsVendorNameEnum, bool) {
	enum, ok := mappingListEntitlementsVendorNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntitlementsSortOrderEnum Enum with underlying type: string
type ListEntitlementsSortOrderEnum string

// Set of constants representing the allowable values for ListEntitlementsSortOrderEnum
const (
	ListEntitlementsSortOrderAsc  ListEntitlementsSortOrderEnum = "ASC"
	ListEntitlementsSortOrderDesc ListEntitlementsSortOrderEnum = "DESC"
)

var mappingListEntitlementsSortOrderEnum = map[string]ListEntitlementsSortOrderEnum{
	"ASC":  ListEntitlementsSortOrderAsc,
	"DESC": ListEntitlementsSortOrderDesc,
}

var mappingListEntitlementsSortOrderEnumLowerCase = map[string]ListEntitlementsSortOrderEnum{
	"asc":  ListEntitlementsSortOrderAsc,
	"desc": ListEntitlementsSortOrderDesc,
}

// GetListEntitlementsSortOrderEnumValues Enumerates the set of values for ListEntitlementsSortOrderEnum
func GetListEntitlementsSortOrderEnumValues() []ListEntitlementsSortOrderEnum {
	values := make([]ListEntitlementsSortOrderEnum, 0)
	for _, v := range mappingListEntitlementsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntitlementsSortOrderEnumStringValues Enumerates the set of values in String for ListEntitlementsSortOrderEnum
func GetListEntitlementsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEntitlementsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntitlementsSortOrderEnum(val string) (ListEntitlementsSortOrderEnum, bool) {
	enum, ok := mappingListEntitlementsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntitlementsSortByEnum Enum with underlying type: string
type ListEntitlementsSortByEnum string

// Set of constants representing the allowable values for ListEntitlementsSortByEnum
const (
	ListEntitlementsSortByCsi        ListEntitlementsSortByEnum = "csi"
	ListEntitlementsSortByVendorname ListEntitlementsSortByEnum = "vendorName"
)

var mappingListEntitlementsSortByEnum = map[string]ListEntitlementsSortByEnum{
	"csi":        ListEntitlementsSortByCsi,
	"vendorName": ListEntitlementsSortByVendorname,
}

var mappingListEntitlementsSortByEnumLowerCase = map[string]ListEntitlementsSortByEnum{
	"csi":        ListEntitlementsSortByCsi,
	"vendorname": ListEntitlementsSortByVendorname,
}

// GetListEntitlementsSortByEnumValues Enumerates the set of values for ListEntitlementsSortByEnum
func GetListEntitlementsSortByEnumValues() []ListEntitlementsSortByEnum {
	values := make([]ListEntitlementsSortByEnum, 0)
	for _, v := range mappingListEntitlementsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntitlementsSortByEnumStringValues Enumerates the set of values in String for ListEntitlementsSortByEnum
func GetListEntitlementsSortByEnumStringValues() []string {
	return []string{
		"csi",
		"vendorName",
	}
}

// GetMappingListEntitlementsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntitlementsSortByEnum(val string) (ListEntitlementsSortByEnum, bool) {
	enum, ok := mappingListEntitlementsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
