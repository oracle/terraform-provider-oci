// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSubsettingReportsRequest wrapper for the ListSubsettingReports operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSubsettingReports.go.html to see an example of how to use ListSubsettingReportsRequest.
type ListSubsettingReportsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the resources that match the specified subsetting policy OCID.
	SubsettingPolicyId *string `mandatory:"false" contributesTo:"query" name:"subsettingPolicyId"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return the target database group that matches the specified OCID.
	TargetDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"targetDatabaseGroupId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSubsettingReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for timeSubsettingFinished is descending.
	SortBy ListSubsettingReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSubsettingReportsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubsettingReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubsettingReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubsettingReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubsettingReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubsettingReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubsettingReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubsettingReportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubsettingReportsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingReportsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSubsettingReportsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubsettingReportsResponse wrapper for the ListSubsettingReports operation
type ListSubsettingReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubsettingReportCollection instances
	SubsettingReportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSubsettingReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubsettingReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubsettingReportsSortOrderEnum Enum with underlying type: string
type ListSubsettingReportsSortOrderEnum string

// Set of constants representing the allowable values for ListSubsettingReportsSortOrderEnum
const (
	ListSubsettingReportsSortOrderAsc  ListSubsettingReportsSortOrderEnum = "ASC"
	ListSubsettingReportsSortOrderDesc ListSubsettingReportsSortOrderEnum = "DESC"
)

var mappingListSubsettingReportsSortOrderEnum = map[string]ListSubsettingReportsSortOrderEnum{
	"ASC":  ListSubsettingReportsSortOrderAsc,
	"DESC": ListSubsettingReportsSortOrderDesc,
}

var mappingListSubsettingReportsSortOrderEnumLowerCase = map[string]ListSubsettingReportsSortOrderEnum{
	"asc":  ListSubsettingReportsSortOrderAsc,
	"desc": ListSubsettingReportsSortOrderDesc,
}

// GetListSubsettingReportsSortOrderEnumValues Enumerates the set of values for ListSubsettingReportsSortOrderEnum
func GetListSubsettingReportsSortOrderEnumValues() []ListSubsettingReportsSortOrderEnum {
	values := make([]ListSubsettingReportsSortOrderEnum, 0)
	for _, v := range mappingListSubsettingReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingReportsSortOrderEnumStringValues Enumerates the set of values in String for ListSubsettingReportsSortOrderEnum
func GetListSubsettingReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubsettingReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingReportsSortOrderEnum(val string) (ListSubsettingReportsSortOrderEnum, bool) {
	enum, ok := mappingListSubsettingReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingReportsSortByEnum Enum with underlying type: string
type ListSubsettingReportsSortByEnum string

// Set of constants representing the allowable values for ListSubsettingReportsSortByEnum
const (
	ListSubsettingReportsSortByTimesubsettingfinished ListSubsettingReportsSortByEnum = "timeSubsettingFinished"
)

var mappingListSubsettingReportsSortByEnum = map[string]ListSubsettingReportsSortByEnum{
	"timeSubsettingFinished": ListSubsettingReportsSortByTimesubsettingfinished,
}

var mappingListSubsettingReportsSortByEnumLowerCase = map[string]ListSubsettingReportsSortByEnum{
	"timesubsettingfinished": ListSubsettingReportsSortByTimesubsettingfinished,
}

// GetListSubsettingReportsSortByEnumValues Enumerates the set of values for ListSubsettingReportsSortByEnum
func GetListSubsettingReportsSortByEnumValues() []ListSubsettingReportsSortByEnum {
	values := make([]ListSubsettingReportsSortByEnum, 0)
	for _, v := range mappingListSubsettingReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingReportsSortByEnumStringValues Enumerates the set of values in String for ListSubsettingReportsSortByEnum
func GetListSubsettingReportsSortByEnumStringValues() []string {
	return []string{
		"timeSubsettingFinished",
	}
}

// GetMappingListSubsettingReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingReportsSortByEnum(val string) (ListSubsettingReportsSortByEnum, bool) {
	enum, ok := mappingListSubsettingReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingReportsAccessLevelEnum Enum with underlying type: string
type ListSubsettingReportsAccessLevelEnum string

// Set of constants representing the allowable values for ListSubsettingReportsAccessLevelEnum
const (
	ListSubsettingReportsAccessLevelRestricted ListSubsettingReportsAccessLevelEnum = "RESTRICTED"
	ListSubsettingReportsAccessLevelAccessible ListSubsettingReportsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSubsettingReportsAccessLevelEnum = map[string]ListSubsettingReportsAccessLevelEnum{
	"RESTRICTED": ListSubsettingReportsAccessLevelRestricted,
	"ACCESSIBLE": ListSubsettingReportsAccessLevelAccessible,
}

var mappingListSubsettingReportsAccessLevelEnumLowerCase = map[string]ListSubsettingReportsAccessLevelEnum{
	"restricted": ListSubsettingReportsAccessLevelRestricted,
	"accessible": ListSubsettingReportsAccessLevelAccessible,
}

// GetListSubsettingReportsAccessLevelEnumValues Enumerates the set of values for ListSubsettingReportsAccessLevelEnum
func GetListSubsettingReportsAccessLevelEnumValues() []ListSubsettingReportsAccessLevelEnum {
	values := make([]ListSubsettingReportsAccessLevelEnum, 0)
	for _, v := range mappingListSubsettingReportsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingReportsAccessLevelEnumStringValues Enumerates the set of values in String for ListSubsettingReportsAccessLevelEnum
func GetListSubsettingReportsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSubsettingReportsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingReportsAccessLevelEnum(val string) (ListSubsettingReportsAccessLevelEnum, bool) {
	enum, ok := mappingListSubsettingReportsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
