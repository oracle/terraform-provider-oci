// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProfileAnalyticsRequest wrapper for the ListProfileAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListProfileAnalytics.go.html to see an example of how to use ListProfileAnalyticsRequest.
type ListProfileAnalyticsRequest struct {

	// The OCID of the user assessment.
	UserAssessmentId *string `mandatory:"true" contributesTo:"path" name:"userAssessmentId"`

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListProfileAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only items that match the specified profile name.
	ProfileName *string `mandatory:"false" contributesTo:"query" name:"profileName"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListProfileAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListProfileAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProfileAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProfileAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProfileAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProfileAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProfileAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProfileAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListProfileAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProfileAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProfileAnalyticsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProfileAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProfileAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProfileAnalyticsResponse wrapper for the ListProfileAnalytics operation
type ListProfileAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ProfileAggregation instances
	Items []ProfileAggregation `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListProfileAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProfileAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProfileAnalyticsAccessLevelEnum Enum with underlying type: string
type ListProfileAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListProfileAnalyticsAccessLevelEnum
const (
	ListProfileAnalyticsAccessLevelRestricted ListProfileAnalyticsAccessLevelEnum = "RESTRICTED"
	ListProfileAnalyticsAccessLevelAccessible ListProfileAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListProfileAnalyticsAccessLevelEnum = map[string]ListProfileAnalyticsAccessLevelEnum{
	"RESTRICTED": ListProfileAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListProfileAnalyticsAccessLevelAccessible,
}

var mappingListProfileAnalyticsAccessLevelEnumLowerCase = map[string]ListProfileAnalyticsAccessLevelEnum{
	"restricted": ListProfileAnalyticsAccessLevelRestricted,
	"accessible": ListProfileAnalyticsAccessLevelAccessible,
}

// GetListProfileAnalyticsAccessLevelEnumValues Enumerates the set of values for ListProfileAnalyticsAccessLevelEnum
func GetListProfileAnalyticsAccessLevelEnumValues() []ListProfileAnalyticsAccessLevelEnum {
	values := make([]ListProfileAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListProfileAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfileAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListProfileAnalyticsAccessLevelEnum
func GetListProfileAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListProfileAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfileAnalyticsAccessLevelEnum(val string) (ListProfileAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListProfileAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProfileAnalyticsSortByEnum Enum with underlying type: string
type ListProfileAnalyticsSortByEnum string

// Set of constants representing the allowable values for ListProfileAnalyticsSortByEnum
const (
	ListProfileAnalyticsSortByTimecreated ListProfileAnalyticsSortByEnum = "TIMECREATED"
	ListProfileAnalyticsSortByDisplayname ListProfileAnalyticsSortByEnum = "DISPLAYNAME"
)

var mappingListProfileAnalyticsSortByEnum = map[string]ListProfileAnalyticsSortByEnum{
	"TIMECREATED": ListProfileAnalyticsSortByTimecreated,
	"DISPLAYNAME": ListProfileAnalyticsSortByDisplayname,
}

var mappingListProfileAnalyticsSortByEnumLowerCase = map[string]ListProfileAnalyticsSortByEnum{
	"timecreated": ListProfileAnalyticsSortByTimecreated,
	"displayname": ListProfileAnalyticsSortByDisplayname,
}

// GetListProfileAnalyticsSortByEnumValues Enumerates the set of values for ListProfileAnalyticsSortByEnum
func GetListProfileAnalyticsSortByEnumValues() []ListProfileAnalyticsSortByEnum {
	values := make([]ListProfileAnalyticsSortByEnum, 0)
	for _, v := range mappingListProfileAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfileAnalyticsSortByEnumStringValues Enumerates the set of values in String for ListProfileAnalyticsSortByEnum
func GetListProfileAnalyticsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListProfileAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfileAnalyticsSortByEnum(val string) (ListProfileAnalyticsSortByEnum, bool) {
	enum, ok := mappingListProfileAnalyticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProfileAnalyticsSortOrderEnum Enum with underlying type: string
type ListProfileAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for ListProfileAnalyticsSortOrderEnum
const (
	ListProfileAnalyticsSortOrderAsc  ListProfileAnalyticsSortOrderEnum = "ASC"
	ListProfileAnalyticsSortOrderDesc ListProfileAnalyticsSortOrderEnum = "DESC"
)

var mappingListProfileAnalyticsSortOrderEnum = map[string]ListProfileAnalyticsSortOrderEnum{
	"ASC":  ListProfileAnalyticsSortOrderAsc,
	"DESC": ListProfileAnalyticsSortOrderDesc,
}

var mappingListProfileAnalyticsSortOrderEnumLowerCase = map[string]ListProfileAnalyticsSortOrderEnum{
	"asc":  ListProfileAnalyticsSortOrderAsc,
	"desc": ListProfileAnalyticsSortOrderDesc,
}

// GetListProfileAnalyticsSortOrderEnumValues Enumerates the set of values for ListProfileAnalyticsSortOrderEnum
func GetListProfileAnalyticsSortOrderEnumValues() []ListProfileAnalyticsSortOrderEnum {
	values := make([]ListProfileAnalyticsSortOrderEnum, 0)
	for _, v := range mappingListProfileAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfileAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for ListProfileAnalyticsSortOrderEnum
func GetListProfileAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProfileAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfileAnalyticsSortOrderEnum(val string) (ListProfileAnalyticsSortOrderEnum, bool) {
	enum, ok := mappingListProfileAnalyticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
