// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListUserAnalyticsRequest wrapper for the ListUserAnalytics operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUserAnalytics.go.html to see an example of how to use ListUserAnalyticsRequest.
type ListUserAnalyticsRequest struct {

	// The OCID of the user assessment.
	UserAssessmentId *string `mandatory:"true" contributesTo:"path" name:"userAssessmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListUserAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to return only items that match the specified user category.
	UserCategory *string `mandatory:"false" contributesTo:"query" name:"userCategory"`

	// A filter to return only items that match the specified user key.
	UserKey *string `mandatory:"false" contributesTo:"query" name:"userKey"`

	// A filter to return only items that match the specified account status.
	AccountStatus *string `mandatory:"false" contributesTo:"query" name:"accountStatus"`

	// A filter to return only items that match the specified authentication type.
	AuthenticationType *string `mandatory:"false" contributesTo:"query" name:"authenticationType"`

	// A filter to return only items that match the specified user name.
	UserName *string `mandatory:"false" contributesTo:"query" name:"userName"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return users whose last login time in the database is greater than or equal to the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeLastLoginGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastLoginGreaterThanOrEqualTo"`

	// A filter to return users whose last login time in the database is less than the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeLastLoginLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastLoginLessThan"`

	// A filter to return users whose creation time in the database is greater than or equal to the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeUserCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUserCreatedGreaterThanOrEqualTo"`

	// A filter to return users whose creation time in the database is less than the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeUserCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUserCreatedLessThan"`

	// A filter to return users whose last password change in the database is greater than or equal to the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimePasswordLastChangedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timePasswordLastChangedGreaterThanOrEqualTo"`

	// A filter to return users whose last password change in the database is less than the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimePasswordLastChangedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timePasswordLastChangedLessThan"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListUserAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order (sortOrder). The default order for userName is ascending.
	SortBy ListUserAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUserAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUserAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUserAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUserAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUserAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUserAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListUserAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUserAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUserAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUserAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUserAnalyticsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUserAnalyticsResponse wrapper for the ListUserAnalytics operation
type ListUserAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []UserAggregation instances
	Items []UserAggregation `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListUserAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUserAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUserAnalyticsAccessLevelEnum Enum with underlying type: string
type ListUserAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListUserAnalyticsAccessLevelEnum
const (
	ListUserAnalyticsAccessLevelRestricted ListUserAnalyticsAccessLevelEnum = "RESTRICTED"
	ListUserAnalyticsAccessLevelAccessible ListUserAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListUserAnalyticsAccessLevelEnum = map[string]ListUserAnalyticsAccessLevelEnum{
	"RESTRICTED": ListUserAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListUserAnalyticsAccessLevelAccessible,
}

// GetListUserAnalyticsAccessLevelEnumValues Enumerates the set of values for ListUserAnalyticsAccessLevelEnum
func GetListUserAnalyticsAccessLevelEnumValues() []ListUserAnalyticsAccessLevelEnum {
	values := make([]ListUserAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListUserAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListUserAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListUserAnalyticsAccessLevelEnum
func GetListUserAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListUserAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUserAnalyticsAccessLevelEnum(val string) (ListUserAnalyticsAccessLevelEnum, bool) {
	mappingListUserAnalyticsAccessLevelEnumIgnoreCase := make(map[string]ListUserAnalyticsAccessLevelEnum)
	for k, v := range mappingListUserAnalyticsAccessLevelEnum {
		mappingListUserAnalyticsAccessLevelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUserAnalyticsAccessLevelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListUserAnalyticsSortOrderEnum Enum with underlying type: string
type ListUserAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for ListUserAnalyticsSortOrderEnum
const (
	ListUserAnalyticsSortOrderAsc  ListUserAnalyticsSortOrderEnum = "ASC"
	ListUserAnalyticsSortOrderDesc ListUserAnalyticsSortOrderEnum = "DESC"
)

var mappingListUserAnalyticsSortOrderEnum = map[string]ListUserAnalyticsSortOrderEnum{
	"ASC":  ListUserAnalyticsSortOrderAsc,
	"DESC": ListUserAnalyticsSortOrderDesc,
}

// GetListUserAnalyticsSortOrderEnumValues Enumerates the set of values for ListUserAnalyticsSortOrderEnum
func GetListUserAnalyticsSortOrderEnumValues() []ListUserAnalyticsSortOrderEnum {
	values := make([]ListUserAnalyticsSortOrderEnum, 0)
	for _, v := range mappingListUserAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUserAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for ListUserAnalyticsSortOrderEnum
func GetListUserAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUserAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUserAnalyticsSortOrderEnum(val string) (ListUserAnalyticsSortOrderEnum, bool) {
	mappingListUserAnalyticsSortOrderEnumIgnoreCase := make(map[string]ListUserAnalyticsSortOrderEnum)
	for k, v := range mappingListUserAnalyticsSortOrderEnum {
		mappingListUserAnalyticsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUserAnalyticsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListUserAnalyticsSortByEnum Enum with underlying type: string
type ListUserAnalyticsSortByEnum string

// Set of constants representing the allowable values for ListUserAnalyticsSortByEnum
const (
	ListUserAnalyticsSortByUsername            ListUserAnalyticsSortByEnum = "userName"
	ListUserAnalyticsSortByUsercategory        ListUserAnalyticsSortByEnum = "userCategory"
	ListUserAnalyticsSortByAccountstatus       ListUserAnalyticsSortByEnum = "accountStatus"
	ListUserAnalyticsSortByTimelastlogin       ListUserAnalyticsSortByEnum = "timeLastLogin"
	ListUserAnalyticsSortByTargetid            ListUserAnalyticsSortByEnum = "targetId"
	ListUserAnalyticsSortByTimeusercreated     ListUserAnalyticsSortByEnum = "timeUserCreated"
	ListUserAnalyticsSortByAuthenticationtype  ListUserAnalyticsSortByEnum = "authenticationType"
	ListUserAnalyticsSortByTimepasswordchanged ListUserAnalyticsSortByEnum = "timePasswordChanged"
)

var mappingListUserAnalyticsSortByEnum = map[string]ListUserAnalyticsSortByEnum{
	"userName":            ListUserAnalyticsSortByUsername,
	"userCategory":        ListUserAnalyticsSortByUsercategory,
	"accountStatus":       ListUserAnalyticsSortByAccountstatus,
	"timeLastLogin":       ListUserAnalyticsSortByTimelastlogin,
	"targetId":            ListUserAnalyticsSortByTargetid,
	"timeUserCreated":     ListUserAnalyticsSortByTimeusercreated,
	"authenticationType":  ListUserAnalyticsSortByAuthenticationtype,
	"timePasswordChanged": ListUserAnalyticsSortByTimepasswordchanged,
}

// GetListUserAnalyticsSortByEnumValues Enumerates the set of values for ListUserAnalyticsSortByEnum
func GetListUserAnalyticsSortByEnumValues() []ListUserAnalyticsSortByEnum {
	values := make([]ListUserAnalyticsSortByEnum, 0)
	for _, v := range mappingListUserAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUserAnalyticsSortByEnumStringValues Enumerates the set of values in String for ListUserAnalyticsSortByEnum
func GetListUserAnalyticsSortByEnumStringValues() []string {
	return []string{
		"userName",
		"userCategory",
		"accountStatus",
		"timeLastLogin",
		"targetId",
		"timeUserCreated",
		"authenticationType",
		"timePasswordChanged",
	}
}

// GetMappingListUserAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUserAnalyticsSortByEnum(val string) (ListUserAnalyticsSortByEnum, bool) {
	mappingListUserAnalyticsSortByEnumIgnoreCase := make(map[string]ListUserAnalyticsSortByEnum)
	for k, v := range mappingListUserAnalyticsSortByEnum {
		mappingListUserAnalyticsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUserAnalyticsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
