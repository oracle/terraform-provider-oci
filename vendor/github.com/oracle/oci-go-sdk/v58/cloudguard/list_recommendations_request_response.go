// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListRecommendationsRequest wrapper for the ListRecommendations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListRecommendations.go.html to see an example of how to use ListRecommendationsRequest.
type ListRecommendationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListRecommendationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for riskLevel and timeCreated is descending. If no value is specified riskLevel is default.
	SortBy ListRecommendationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The ID of the target in which to list resources.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel ListRecommendationsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListRecommendationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleDetail ListRecommendationsLifecycleDetailEnum `mandatory:"false" contributesTo:"query" name:"lifecycleDetail" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecommendationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecommendationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecommendationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecommendationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRecommendationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRecommendationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRecommendationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecommendationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRecommendationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecommendationsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListRecommendationsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecommendationsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListRecommendationsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecommendationsLifecycleDetailEnum(string(request.LifecycleDetail)); !ok && request.LifecycleDetail != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetail: %s. Supported values are: %s.", request.LifecycleDetail, strings.Join(GetListRecommendationsLifecycleDetailEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRecommendationsResponse wrapper for the ListRecommendations operation
type ListRecommendationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecommendationSummaryCollection instances
	RecommendationSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRecommendationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecommendationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecommendationsSortOrderEnum Enum with underlying type: string
type ListRecommendationsSortOrderEnum string

// Set of constants representing the allowable values for ListRecommendationsSortOrderEnum
const (
	ListRecommendationsSortOrderAsc  ListRecommendationsSortOrderEnum = "ASC"
	ListRecommendationsSortOrderDesc ListRecommendationsSortOrderEnum = "DESC"
)

var mappingListRecommendationsSortOrderEnum = map[string]ListRecommendationsSortOrderEnum{
	"ASC":  ListRecommendationsSortOrderAsc,
	"DESC": ListRecommendationsSortOrderDesc,
}

// GetListRecommendationsSortOrderEnumValues Enumerates the set of values for ListRecommendationsSortOrderEnum
func GetListRecommendationsSortOrderEnumValues() []ListRecommendationsSortOrderEnum {
	values := make([]ListRecommendationsSortOrderEnum, 0)
	for _, v := range mappingListRecommendationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendationsSortOrderEnumStringValues Enumerates the set of values in String for ListRecommendationsSortOrderEnum
func GetListRecommendationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRecommendationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendationsSortOrderEnum(val string) (ListRecommendationsSortOrderEnum, bool) {
	mappingListRecommendationsSortOrderEnumIgnoreCase := make(map[string]ListRecommendationsSortOrderEnum)
	for k, v := range mappingListRecommendationsSortOrderEnum {
		mappingListRecommendationsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRecommendationsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecommendationsSortByEnum Enum with underlying type: string
type ListRecommendationsSortByEnum string

// Set of constants representing the allowable values for ListRecommendationsSortByEnum
const (
	ListRecommendationsSortByRisklevel   ListRecommendationsSortByEnum = "riskLevel"
	ListRecommendationsSortByTimecreated ListRecommendationsSortByEnum = "timeCreated"
)

var mappingListRecommendationsSortByEnum = map[string]ListRecommendationsSortByEnum{
	"riskLevel":   ListRecommendationsSortByRisklevel,
	"timeCreated": ListRecommendationsSortByTimecreated,
}

// GetListRecommendationsSortByEnumValues Enumerates the set of values for ListRecommendationsSortByEnum
func GetListRecommendationsSortByEnumValues() []ListRecommendationsSortByEnum {
	values := make([]ListRecommendationsSortByEnum, 0)
	for _, v := range mappingListRecommendationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendationsSortByEnumStringValues Enumerates the set of values in String for ListRecommendationsSortByEnum
func GetListRecommendationsSortByEnumStringValues() []string {
	return []string{
		"riskLevel",
		"timeCreated",
	}
}

// GetMappingListRecommendationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendationsSortByEnum(val string) (ListRecommendationsSortByEnum, bool) {
	mappingListRecommendationsSortByEnumIgnoreCase := make(map[string]ListRecommendationsSortByEnum)
	for k, v := range mappingListRecommendationsSortByEnum {
		mappingListRecommendationsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRecommendationsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecommendationsAccessLevelEnum Enum with underlying type: string
type ListRecommendationsAccessLevelEnum string

// Set of constants representing the allowable values for ListRecommendationsAccessLevelEnum
const (
	ListRecommendationsAccessLevelRestricted ListRecommendationsAccessLevelEnum = "RESTRICTED"
	ListRecommendationsAccessLevelAccessible ListRecommendationsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListRecommendationsAccessLevelEnum = map[string]ListRecommendationsAccessLevelEnum{
	"RESTRICTED": ListRecommendationsAccessLevelRestricted,
	"ACCESSIBLE": ListRecommendationsAccessLevelAccessible,
}

// GetListRecommendationsAccessLevelEnumValues Enumerates the set of values for ListRecommendationsAccessLevelEnum
func GetListRecommendationsAccessLevelEnumValues() []ListRecommendationsAccessLevelEnum {
	values := make([]ListRecommendationsAccessLevelEnum, 0)
	for _, v := range mappingListRecommendationsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendationsAccessLevelEnumStringValues Enumerates the set of values in String for ListRecommendationsAccessLevelEnum
func GetListRecommendationsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListRecommendationsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendationsAccessLevelEnum(val string) (ListRecommendationsAccessLevelEnum, bool) {
	mappingListRecommendationsAccessLevelEnumIgnoreCase := make(map[string]ListRecommendationsAccessLevelEnum)
	for k, v := range mappingListRecommendationsAccessLevelEnum {
		mappingListRecommendationsAccessLevelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRecommendationsAccessLevelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecommendationsLifecycleStateEnum Enum with underlying type: string
type ListRecommendationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListRecommendationsLifecycleStateEnum
const (
	ListRecommendationsLifecycleStateCreating ListRecommendationsLifecycleStateEnum = "CREATING"
	ListRecommendationsLifecycleStateUpdating ListRecommendationsLifecycleStateEnum = "UPDATING"
	ListRecommendationsLifecycleStateActive   ListRecommendationsLifecycleStateEnum = "ACTIVE"
	ListRecommendationsLifecycleStateInactive ListRecommendationsLifecycleStateEnum = "INACTIVE"
	ListRecommendationsLifecycleStateDeleting ListRecommendationsLifecycleStateEnum = "DELETING"
	ListRecommendationsLifecycleStateDeleted  ListRecommendationsLifecycleStateEnum = "DELETED"
	ListRecommendationsLifecycleStateFailed   ListRecommendationsLifecycleStateEnum = "FAILED"
)

var mappingListRecommendationsLifecycleStateEnum = map[string]ListRecommendationsLifecycleStateEnum{
	"CREATING": ListRecommendationsLifecycleStateCreating,
	"UPDATING": ListRecommendationsLifecycleStateUpdating,
	"ACTIVE":   ListRecommendationsLifecycleStateActive,
	"INACTIVE": ListRecommendationsLifecycleStateInactive,
	"DELETING": ListRecommendationsLifecycleStateDeleting,
	"DELETED":  ListRecommendationsLifecycleStateDeleted,
	"FAILED":   ListRecommendationsLifecycleStateFailed,
}

// GetListRecommendationsLifecycleStateEnumValues Enumerates the set of values for ListRecommendationsLifecycleStateEnum
func GetListRecommendationsLifecycleStateEnumValues() []ListRecommendationsLifecycleStateEnum {
	values := make([]ListRecommendationsLifecycleStateEnum, 0)
	for _, v := range mappingListRecommendationsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendationsLifecycleStateEnumStringValues Enumerates the set of values in String for ListRecommendationsLifecycleStateEnum
func GetListRecommendationsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListRecommendationsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendationsLifecycleStateEnum(val string) (ListRecommendationsLifecycleStateEnum, bool) {
	mappingListRecommendationsLifecycleStateEnumIgnoreCase := make(map[string]ListRecommendationsLifecycleStateEnum)
	for k, v := range mappingListRecommendationsLifecycleStateEnum {
		mappingListRecommendationsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRecommendationsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecommendationsLifecycleDetailEnum Enum with underlying type: string
type ListRecommendationsLifecycleDetailEnum string

// Set of constants representing the allowable values for ListRecommendationsLifecycleDetailEnum
const (
	ListRecommendationsLifecycleDetailOpen      ListRecommendationsLifecycleDetailEnum = "OPEN"
	ListRecommendationsLifecycleDetailResolved  ListRecommendationsLifecycleDetailEnum = "RESOLVED"
	ListRecommendationsLifecycleDetailDismissed ListRecommendationsLifecycleDetailEnum = "DISMISSED"
)

var mappingListRecommendationsLifecycleDetailEnum = map[string]ListRecommendationsLifecycleDetailEnum{
	"OPEN":      ListRecommendationsLifecycleDetailOpen,
	"RESOLVED":  ListRecommendationsLifecycleDetailResolved,
	"DISMISSED": ListRecommendationsLifecycleDetailDismissed,
}

// GetListRecommendationsLifecycleDetailEnumValues Enumerates the set of values for ListRecommendationsLifecycleDetailEnum
func GetListRecommendationsLifecycleDetailEnumValues() []ListRecommendationsLifecycleDetailEnum {
	values := make([]ListRecommendationsLifecycleDetailEnum, 0)
	for _, v := range mappingListRecommendationsLifecycleDetailEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendationsLifecycleDetailEnumStringValues Enumerates the set of values in String for ListRecommendationsLifecycleDetailEnum
func GetListRecommendationsLifecycleDetailEnumStringValues() []string {
	return []string{
		"OPEN",
		"RESOLVED",
		"DISMISSED",
	}
}

// GetMappingListRecommendationsLifecycleDetailEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendationsLifecycleDetailEnum(val string) (ListRecommendationsLifecycleDetailEnum, bool) {
	mappingListRecommendationsLifecycleDetailEnumIgnoreCase := make(map[string]ListRecommendationsLifecycleDetailEnum)
	for k, v := range mappingListRecommendationsLifecycleDetailEnum {
		mappingListRecommendationsLifecycleDetailEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRecommendationsLifecycleDetailEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
