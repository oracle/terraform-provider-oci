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

// ListAlertPoliciesRequest wrapper for the ListAlertPolicies operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAlertPolicies.go.html to see an example of how to use ListAlertPoliciesRequest.
type ListAlertPoliciesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return policy by it's OCID.
	AlertPolicyId *string `mandatory:"false" contributesTo:"query" name:"alertPolicyId"`

	// An optional filter to return only alert policies of a certain type.
	Type ListAlertPoliciesTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// An optional filter to return only alert policies that are user-defined or not.
	IsUserDefined *bool `mandatory:"false" contributesTo:"query" name:"isUserDefined"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// An optional filter to return only alert policies that have the given life-cycle state.
	LifecycleState ListAlertPoliciesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListAlertPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListAlertPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

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

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListAlertPoliciesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAlertPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAlertPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAlertPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAlertPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAlertPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAlertPoliciesTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListAlertPoliciesTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertPoliciesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAlertPoliciesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAlertPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAlertPoliciesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertPoliciesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAlertPoliciesAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAlertPoliciesResponse wrapper for the ListAlertPolicies operation
type ListAlertPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AlertPolicyCollection instances
	AlertPolicyCollection `presentIn:"body"`

	// For optimistic concurrency control. For more information, see ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven)
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAlertPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAlertPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAlertPoliciesTypeEnum Enum with underlying type: string
type ListAlertPoliciesTypeEnum string

// Set of constants representing the allowable values for ListAlertPoliciesTypeEnum
const (
	ListAlertPoliciesTypeAuditing           ListAlertPoliciesTypeEnum = "AUDITING"
	ListAlertPoliciesTypeSecurityAssessment ListAlertPoliciesTypeEnum = "SECURITY_ASSESSMENT"
	ListAlertPoliciesTypeUserAssessment     ListAlertPoliciesTypeEnum = "USER_ASSESSMENT"
)

var mappingListAlertPoliciesTypeEnum = map[string]ListAlertPoliciesTypeEnum{
	"AUDITING":            ListAlertPoliciesTypeAuditing,
	"SECURITY_ASSESSMENT": ListAlertPoliciesTypeSecurityAssessment,
	"USER_ASSESSMENT":     ListAlertPoliciesTypeUserAssessment,
}

// GetListAlertPoliciesTypeEnumValues Enumerates the set of values for ListAlertPoliciesTypeEnum
func GetListAlertPoliciesTypeEnumValues() []ListAlertPoliciesTypeEnum {
	values := make([]ListAlertPoliciesTypeEnum, 0)
	for _, v := range mappingListAlertPoliciesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertPoliciesTypeEnumStringValues Enumerates the set of values in String for ListAlertPoliciesTypeEnum
func GetListAlertPoliciesTypeEnumStringValues() []string {
	return []string{
		"AUDITING",
		"SECURITY_ASSESSMENT",
		"USER_ASSESSMENT",
	}
}

// GetMappingListAlertPoliciesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertPoliciesTypeEnum(val string) (ListAlertPoliciesTypeEnum, bool) {
	mappingListAlertPoliciesTypeEnumIgnoreCase := make(map[string]ListAlertPoliciesTypeEnum)
	for k, v := range mappingListAlertPoliciesTypeEnum {
		mappingListAlertPoliciesTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAlertPoliciesTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertPoliciesLifecycleStateEnum Enum with underlying type: string
type ListAlertPoliciesLifecycleStateEnum string

// Set of constants representing the allowable values for ListAlertPoliciesLifecycleStateEnum
const (
	ListAlertPoliciesLifecycleStateCreating ListAlertPoliciesLifecycleStateEnum = "CREATING"
	ListAlertPoliciesLifecycleStateUpdating ListAlertPoliciesLifecycleStateEnum = "UPDATING"
	ListAlertPoliciesLifecycleStateActive   ListAlertPoliciesLifecycleStateEnum = "ACTIVE"
	ListAlertPoliciesLifecycleStateDeleting ListAlertPoliciesLifecycleStateEnum = "DELETING"
	ListAlertPoliciesLifecycleStateDeleted  ListAlertPoliciesLifecycleStateEnum = "DELETED"
	ListAlertPoliciesLifecycleStateFailed   ListAlertPoliciesLifecycleStateEnum = "FAILED"
)

var mappingListAlertPoliciesLifecycleStateEnum = map[string]ListAlertPoliciesLifecycleStateEnum{
	"CREATING": ListAlertPoliciesLifecycleStateCreating,
	"UPDATING": ListAlertPoliciesLifecycleStateUpdating,
	"ACTIVE":   ListAlertPoliciesLifecycleStateActive,
	"DELETING": ListAlertPoliciesLifecycleStateDeleting,
	"DELETED":  ListAlertPoliciesLifecycleStateDeleted,
	"FAILED":   ListAlertPoliciesLifecycleStateFailed,
}

// GetListAlertPoliciesLifecycleStateEnumValues Enumerates the set of values for ListAlertPoliciesLifecycleStateEnum
func GetListAlertPoliciesLifecycleStateEnumValues() []ListAlertPoliciesLifecycleStateEnum {
	values := make([]ListAlertPoliciesLifecycleStateEnum, 0)
	for _, v := range mappingListAlertPoliciesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertPoliciesLifecycleStateEnumStringValues Enumerates the set of values in String for ListAlertPoliciesLifecycleStateEnum
func GetListAlertPoliciesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListAlertPoliciesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertPoliciesLifecycleStateEnum(val string) (ListAlertPoliciesLifecycleStateEnum, bool) {
	mappingListAlertPoliciesLifecycleStateEnumIgnoreCase := make(map[string]ListAlertPoliciesLifecycleStateEnum)
	for k, v := range mappingListAlertPoliciesLifecycleStateEnum {
		mappingListAlertPoliciesLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAlertPoliciesLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertPoliciesSortOrderEnum Enum with underlying type: string
type ListAlertPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListAlertPoliciesSortOrderEnum
const (
	ListAlertPoliciesSortOrderAsc  ListAlertPoliciesSortOrderEnum = "ASC"
	ListAlertPoliciesSortOrderDesc ListAlertPoliciesSortOrderEnum = "DESC"
)

var mappingListAlertPoliciesSortOrderEnum = map[string]ListAlertPoliciesSortOrderEnum{
	"ASC":  ListAlertPoliciesSortOrderAsc,
	"DESC": ListAlertPoliciesSortOrderDesc,
}

// GetListAlertPoliciesSortOrderEnumValues Enumerates the set of values for ListAlertPoliciesSortOrderEnum
func GetListAlertPoliciesSortOrderEnumValues() []ListAlertPoliciesSortOrderEnum {
	values := make([]ListAlertPoliciesSortOrderEnum, 0)
	for _, v := range mappingListAlertPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListAlertPoliciesSortOrderEnum
func GetListAlertPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAlertPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertPoliciesSortOrderEnum(val string) (ListAlertPoliciesSortOrderEnum, bool) {
	mappingListAlertPoliciesSortOrderEnumIgnoreCase := make(map[string]ListAlertPoliciesSortOrderEnum)
	for k, v := range mappingListAlertPoliciesSortOrderEnum {
		mappingListAlertPoliciesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAlertPoliciesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertPoliciesSortByEnum Enum with underlying type: string
type ListAlertPoliciesSortByEnum string

// Set of constants representing the allowable values for ListAlertPoliciesSortByEnum
const (
	ListAlertPoliciesSortByDisplayname ListAlertPoliciesSortByEnum = "displayName"
	ListAlertPoliciesSortByTimecreated ListAlertPoliciesSortByEnum = "timeCreated"
)

var mappingListAlertPoliciesSortByEnum = map[string]ListAlertPoliciesSortByEnum{
	"displayName": ListAlertPoliciesSortByDisplayname,
	"timeCreated": ListAlertPoliciesSortByTimecreated,
}

// GetListAlertPoliciesSortByEnumValues Enumerates the set of values for ListAlertPoliciesSortByEnum
func GetListAlertPoliciesSortByEnumValues() []ListAlertPoliciesSortByEnum {
	values := make([]ListAlertPoliciesSortByEnum, 0)
	for _, v := range mappingListAlertPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertPoliciesSortByEnumStringValues Enumerates the set of values in String for ListAlertPoliciesSortByEnum
func GetListAlertPoliciesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListAlertPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertPoliciesSortByEnum(val string) (ListAlertPoliciesSortByEnum, bool) {
	mappingListAlertPoliciesSortByEnumIgnoreCase := make(map[string]ListAlertPoliciesSortByEnum)
	for k, v := range mappingListAlertPoliciesSortByEnum {
		mappingListAlertPoliciesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAlertPoliciesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertPoliciesAccessLevelEnum Enum with underlying type: string
type ListAlertPoliciesAccessLevelEnum string

// Set of constants representing the allowable values for ListAlertPoliciesAccessLevelEnum
const (
	ListAlertPoliciesAccessLevelRestricted ListAlertPoliciesAccessLevelEnum = "RESTRICTED"
	ListAlertPoliciesAccessLevelAccessible ListAlertPoliciesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAlertPoliciesAccessLevelEnum = map[string]ListAlertPoliciesAccessLevelEnum{
	"RESTRICTED": ListAlertPoliciesAccessLevelRestricted,
	"ACCESSIBLE": ListAlertPoliciesAccessLevelAccessible,
}

// GetListAlertPoliciesAccessLevelEnumValues Enumerates the set of values for ListAlertPoliciesAccessLevelEnum
func GetListAlertPoliciesAccessLevelEnumValues() []ListAlertPoliciesAccessLevelEnum {
	values := make([]ListAlertPoliciesAccessLevelEnum, 0)
	for _, v := range mappingListAlertPoliciesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertPoliciesAccessLevelEnumStringValues Enumerates the set of values in String for ListAlertPoliciesAccessLevelEnum
func GetListAlertPoliciesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAlertPoliciesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertPoliciesAccessLevelEnum(val string) (ListAlertPoliciesAccessLevelEnum, bool) {
	mappingListAlertPoliciesAccessLevelEnumIgnoreCase := make(map[string]ListAlertPoliciesAccessLevelEnum)
	for k, v := range mappingListAlertPoliciesAccessLevelEnum {
		mappingListAlertPoliciesAccessLevelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAlertPoliciesAccessLevelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
