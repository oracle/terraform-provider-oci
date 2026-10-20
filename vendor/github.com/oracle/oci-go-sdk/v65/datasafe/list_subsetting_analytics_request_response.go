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

// ListSubsettingAnalyticsRequest wrapper for the ListSubsettingAnalytics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSubsettingAnalytics.go.html to see an example of how to use ListSubsettingAnalyticsRequest.
type ListSubsettingAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Attribute by which the subsetting analytics data should be grouped.
	GroupBy ListSubsettingAnalyticsGroupByEnum `mandatory:"false" contributesTo:"query" name:"groupBy" omitEmpty:"true"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only the resources that match the specified subsetting policy OCID.
	SubsettingPolicyId *string `mandatory:"false" contributesTo:"query" name:"subsettingPolicyId"`

	// A filter to return the target database group that matches the specified OCID.
	TargetDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"targetDatabaseGroupId"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for all the fields is ascending.
	SortBy ListSubsettingAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSubsettingAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only the resources that were created after the specified date and time, as defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// Search for resources that were created before a specific date.
	// Specifying this parameter corresponding `timeCreatedLessThan`
	// parameter will retrieve all resources created before the
	// specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubsettingAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubsettingAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubsettingAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubsettingAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubsettingAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubsettingAnalyticsGroupByEnum(string(request.GroupBy)); !ok && request.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", request.GroupBy, strings.Join(GetListSubsettingAnalyticsGroupByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubsettingAnalyticsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubsettingAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubsettingAnalyticsResponse wrapper for the ListSubsettingAnalytics operation
type ListSubsettingAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubsettingAnalyticsCollection instances
	SubsettingAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSubsettingAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubsettingAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubsettingAnalyticsGroupByEnum Enum with underlying type: string
type ListSubsettingAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListSubsettingAnalyticsGroupByEnum
const (
	ListSubsettingAnalyticsGroupByTargetid            ListSubsettingAnalyticsGroupByEnum = "targetId"
	ListSubsettingAnalyticsGroupByPolicyid            ListSubsettingAnalyticsGroupByEnum = "policyId"
	ListSubsettingAnalyticsGroupByTargetidandpolicyid ListSubsettingAnalyticsGroupByEnum = "targetIdAndPolicyId"
)

var mappingListSubsettingAnalyticsGroupByEnum = map[string]ListSubsettingAnalyticsGroupByEnum{
	"targetId":            ListSubsettingAnalyticsGroupByTargetid,
	"policyId":            ListSubsettingAnalyticsGroupByPolicyid,
	"targetIdAndPolicyId": ListSubsettingAnalyticsGroupByTargetidandpolicyid,
}

var mappingListSubsettingAnalyticsGroupByEnumLowerCase = map[string]ListSubsettingAnalyticsGroupByEnum{
	"targetid":            ListSubsettingAnalyticsGroupByTargetid,
	"policyid":            ListSubsettingAnalyticsGroupByPolicyid,
	"targetidandpolicyid": ListSubsettingAnalyticsGroupByTargetidandpolicyid,
}

// GetListSubsettingAnalyticsGroupByEnumValues Enumerates the set of values for ListSubsettingAnalyticsGroupByEnum
func GetListSubsettingAnalyticsGroupByEnumValues() []ListSubsettingAnalyticsGroupByEnum {
	values := make([]ListSubsettingAnalyticsGroupByEnum, 0)
	for _, v := range mappingListSubsettingAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListSubsettingAnalyticsGroupByEnum
func GetListSubsettingAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"targetId",
		"policyId",
		"targetIdAndPolicyId",
	}
}

// GetMappingListSubsettingAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingAnalyticsGroupByEnum(val string) (ListSubsettingAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListSubsettingAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingAnalyticsSortByEnum Enum with underlying type: string
type ListSubsettingAnalyticsSortByEnum string

// Set of constants representing the allowable values for ListSubsettingAnalyticsSortByEnum
const (
	ListSubsettingAnalyticsSortByTimelastsubsetted ListSubsettingAnalyticsSortByEnum = "timeLastSubsetted"
)

var mappingListSubsettingAnalyticsSortByEnum = map[string]ListSubsettingAnalyticsSortByEnum{
	"timeLastSubsetted": ListSubsettingAnalyticsSortByTimelastsubsetted,
}

var mappingListSubsettingAnalyticsSortByEnumLowerCase = map[string]ListSubsettingAnalyticsSortByEnum{
	"timelastsubsetted": ListSubsettingAnalyticsSortByTimelastsubsetted,
}

// GetListSubsettingAnalyticsSortByEnumValues Enumerates the set of values for ListSubsettingAnalyticsSortByEnum
func GetListSubsettingAnalyticsSortByEnumValues() []ListSubsettingAnalyticsSortByEnum {
	values := make([]ListSubsettingAnalyticsSortByEnum, 0)
	for _, v := range mappingListSubsettingAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingAnalyticsSortByEnumStringValues Enumerates the set of values in String for ListSubsettingAnalyticsSortByEnum
func GetListSubsettingAnalyticsSortByEnumStringValues() []string {
	return []string{
		"timeLastSubsetted",
	}
}

// GetMappingListSubsettingAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingAnalyticsSortByEnum(val string) (ListSubsettingAnalyticsSortByEnum, bool) {
	enum, ok := mappingListSubsettingAnalyticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingAnalyticsSortOrderEnum Enum with underlying type: string
type ListSubsettingAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for ListSubsettingAnalyticsSortOrderEnum
const (
	ListSubsettingAnalyticsSortOrderAsc  ListSubsettingAnalyticsSortOrderEnum = "ASC"
	ListSubsettingAnalyticsSortOrderDesc ListSubsettingAnalyticsSortOrderEnum = "DESC"
)

var mappingListSubsettingAnalyticsSortOrderEnum = map[string]ListSubsettingAnalyticsSortOrderEnum{
	"ASC":  ListSubsettingAnalyticsSortOrderAsc,
	"DESC": ListSubsettingAnalyticsSortOrderDesc,
}

var mappingListSubsettingAnalyticsSortOrderEnumLowerCase = map[string]ListSubsettingAnalyticsSortOrderEnum{
	"asc":  ListSubsettingAnalyticsSortOrderAsc,
	"desc": ListSubsettingAnalyticsSortOrderDesc,
}

// GetListSubsettingAnalyticsSortOrderEnumValues Enumerates the set of values for ListSubsettingAnalyticsSortOrderEnum
func GetListSubsettingAnalyticsSortOrderEnumValues() []ListSubsettingAnalyticsSortOrderEnum {
	values := make([]ListSubsettingAnalyticsSortOrderEnum, 0)
	for _, v := range mappingListSubsettingAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for ListSubsettingAnalyticsSortOrderEnum
func GetListSubsettingAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubsettingAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingAnalyticsSortOrderEnum(val string) (ListSubsettingAnalyticsSortOrderEnum, bool) {
	enum, ok := mappingListSubsettingAnalyticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
