// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMaskingReportsRequest wrapper for the ListMaskingReports operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingReports.go.html to see an example of how to use ListMaskingReportsRequest.
type ListMaskingReportsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the resources that match the specified masking policy OCID.
	MaskingPolicyId *string `mandatory:"false" contributesTo:"query" name:"maskingPolicyId"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListMaskingReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for timeMaskingFinished is descending.
	SortBy ListMaskingReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListMaskingReportsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaskingReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaskingReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaskingReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaskingReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaskingReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaskingReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaskingReportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaskingReportsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingReportsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListMaskingReportsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaskingReportsResponse wrapper for the ListMaskingReports operation
type ListMaskingReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaskingReportCollection instances
	MaskingReportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListMaskingReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaskingReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaskingReportsSortOrderEnum Enum with underlying type: string
type ListMaskingReportsSortOrderEnum string

// Set of constants representing the allowable values for ListMaskingReportsSortOrderEnum
const (
	ListMaskingReportsSortOrderAsc  ListMaskingReportsSortOrderEnum = "ASC"
	ListMaskingReportsSortOrderDesc ListMaskingReportsSortOrderEnum = "DESC"
)

var mappingListMaskingReportsSortOrderEnum = map[string]ListMaskingReportsSortOrderEnum{
	"ASC":  ListMaskingReportsSortOrderAsc,
	"DESC": ListMaskingReportsSortOrderDesc,
}

var mappingListMaskingReportsSortOrderEnumLowerCase = map[string]ListMaskingReportsSortOrderEnum{
	"asc":  ListMaskingReportsSortOrderAsc,
	"desc": ListMaskingReportsSortOrderDesc,
}

// GetListMaskingReportsSortOrderEnumValues Enumerates the set of values for ListMaskingReportsSortOrderEnum
func GetListMaskingReportsSortOrderEnumValues() []ListMaskingReportsSortOrderEnum {
	values := make([]ListMaskingReportsSortOrderEnum, 0)
	for _, v := range mappingListMaskingReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingReportsSortOrderEnumStringValues Enumerates the set of values in String for ListMaskingReportsSortOrderEnum
func GetListMaskingReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaskingReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingReportsSortOrderEnum(val string) (ListMaskingReportsSortOrderEnum, bool) {
	enum, ok := mappingListMaskingReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingReportsSortByEnum Enum with underlying type: string
type ListMaskingReportsSortByEnum string

// Set of constants representing the allowable values for ListMaskingReportsSortByEnum
const (
	ListMaskingReportsSortByTimemaskingfinished ListMaskingReportsSortByEnum = "timeMaskingFinished"
)

var mappingListMaskingReportsSortByEnum = map[string]ListMaskingReportsSortByEnum{
	"timeMaskingFinished": ListMaskingReportsSortByTimemaskingfinished,
}

var mappingListMaskingReportsSortByEnumLowerCase = map[string]ListMaskingReportsSortByEnum{
	"timemaskingfinished": ListMaskingReportsSortByTimemaskingfinished,
}

// GetListMaskingReportsSortByEnumValues Enumerates the set of values for ListMaskingReportsSortByEnum
func GetListMaskingReportsSortByEnumValues() []ListMaskingReportsSortByEnum {
	values := make([]ListMaskingReportsSortByEnum, 0)
	for _, v := range mappingListMaskingReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingReportsSortByEnumStringValues Enumerates the set of values in String for ListMaskingReportsSortByEnum
func GetListMaskingReportsSortByEnumStringValues() []string {
	return []string{
		"timeMaskingFinished",
	}
}

// GetMappingListMaskingReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingReportsSortByEnum(val string) (ListMaskingReportsSortByEnum, bool) {
	enum, ok := mappingListMaskingReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingReportsAccessLevelEnum Enum with underlying type: string
type ListMaskingReportsAccessLevelEnum string

// Set of constants representing the allowable values for ListMaskingReportsAccessLevelEnum
const (
	ListMaskingReportsAccessLevelRestricted ListMaskingReportsAccessLevelEnum = "RESTRICTED"
	ListMaskingReportsAccessLevelAccessible ListMaskingReportsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListMaskingReportsAccessLevelEnum = map[string]ListMaskingReportsAccessLevelEnum{
	"RESTRICTED": ListMaskingReportsAccessLevelRestricted,
	"ACCESSIBLE": ListMaskingReportsAccessLevelAccessible,
}

var mappingListMaskingReportsAccessLevelEnumLowerCase = map[string]ListMaskingReportsAccessLevelEnum{
	"restricted": ListMaskingReportsAccessLevelRestricted,
	"accessible": ListMaskingReportsAccessLevelAccessible,
}

// GetListMaskingReportsAccessLevelEnumValues Enumerates the set of values for ListMaskingReportsAccessLevelEnum
func GetListMaskingReportsAccessLevelEnumValues() []ListMaskingReportsAccessLevelEnum {
	values := make([]ListMaskingReportsAccessLevelEnum, 0)
	for _, v := range mappingListMaskingReportsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingReportsAccessLevelEnumStringValues Enumerates the set of values in String for ListMaskingReportsAccessLevelEnum
func GetListMaskingReportsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListMaskingReportsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingReportsAccessLevelEnum(val string) (ListMaskingReportsAccessLevelEnum, bool) {
	enum, ok := mappingListMaskingReportsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
