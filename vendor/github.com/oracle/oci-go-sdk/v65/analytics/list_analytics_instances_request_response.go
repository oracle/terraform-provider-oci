// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package analytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAnalyticsInstancesRequest wrapper for the ListAnalyticsInstances operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/ListAnalyticsInstances.go.html to see an example of how to use ListAnalyticsInstancesRequest.
type ListAnalyticsInstancesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to only return resources matching the capacity type enum. Values are
	// case-insensitive.
	CapacityType ListAnalyticsInstancesCapacityTypeEnum `mandatory:"false" contributesTo:"query" name:"capacityType" omitEmpty:"true"`

	// A filter to only return resources matching the feature set. Values are
	// case-insensitive.
	FeatureSet ListAnalyticsInstancesFeatureSetEnum `mandatory:"false" contributesTo:"query" name:"featureSet" omitEmpty:"true"`

	// A filter to only return resources matching the lifecycle state. The state
	// value is case-insensitive.
	LifecycleState ListAnalyticsInstancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by (one column only). Default sort order is
	// ascending exception of `timeCreated` column (descending).
	SortBy ListAnalyticsInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAnalyticsInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAnalyticsInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAnalyticsInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAnalyticsInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAnalyticsInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAnalyticsInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAnalyticsInstancesCapacityTypeEnum(string(request.CapacityType)); !ok && request.CapacityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CapacityType: %s. Supported values are: %s.", request.CapacityType, strings.Join(GetListAnalyticsInstancesCapacityTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAnalyticsInstancesFeatureSetEnum(string(request.FeatureSet)); !ok && request.FeatureSet != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FeatureSet: %s. Supported values are: %s.", request.FeatureSet, strings.Join(GetListAnalyticsInstancesFeatureSetEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAnalyticsInstancesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAnalyticsInstancesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAnalyticsInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAnalyticsInstancesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAnalyticsInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAnalyticsInstancesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAnalyticsInstancesResponse wrapper for the ListAnalyticsInstances operation
type ListAnalyticsInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AnalyticsInstanceSummary instances
	Items []AnalyticsInstanceSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAnalyticsInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAnalyticsInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAnalyticsInstancesCapacityTypeEnum Enum with underlying type: string
type ListAnalyticsInstancesCapacityTypeEnum string

// Set of constants representing the allowable values for ListAnalyticsInstancesCapacityTypeEnum
const (
	ListAnalyticsInstancesCapacityTypeOlpuCount ListAnalyticsInstancesCapacityTypeEnum = "OLPU_COUNT"
	ListAnalyticsInstancesCapacityTypeUserCount ListAnalyticsInstancesCapacityTypeEnum = "USER_COUNT"
)

var mappingListAnalyticsInstancesCapacityTypeEnum = map[string]ListAnalyticsInstancesCapacityTypeEnum{
	"OLPU_COUNT": ListAnalyticsInstancesCapacityTypeOlpuCount,
	"USER_COUNT": ListAnalyticsInstancesCapacityTypeUserCount,
}

var mappingListAnalyticsInstancesCapacityTypeEnumLowerCase = map[string]ListAnalyticsInstancesCapacityTypeEnum{
	"olpu_count": ListAnalyticsInstancesCapacityTypeOlpuCount,
	"user_count": ListAnalyticsInstancesCapacityTypeUserCount,
}

// GetListAnalyticsInstancesCapacityTypeEnumValues Enumerates the set of values for ListAnalyticsInstancesCapacityTypeEnum
func GetListAnalyticsInstancesCapacityTypeEnumValues() []ListAnalyticsInstancesCapacityTypeEnum {
	values := make([]ListAnalyticsInstancesCapacityTypeEnum, 0)
	for _, v := range mappingListAnalyticsInstancesCapacityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnalyticsInstancesCapacityTypeEnumStringValues Enumerates the set of values in String for ListAnalyticsInstancesCapacityTypeEnum
func GetListAnalyticsInstancesCapacityTypeEnumStringValues() []string {
	return []string{
		"OLPU_COUNT",
		"USER_COUNT",
	}
}

// GetMappingListAnalyticsInstancesCapacityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnalyticsInstancesCapacityTypeEnum(val string) (ListAnalyticsInstancesCapacityTypeEnum, bool) {
	enum, ok := mappingListAnalyticsInstancesCapacityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAnalyticsInstancesFeatureSetEnum Enum with underlying type: string
type ListAnalyticsInstancesFeatureSetEnum string

// Set of constants representing the allowable values for ListAnalyticsInstancesFeatureSetEnum
const (
	ListAnalyticsInstancesFeatureSetSelfServiceAnalytics ListAnalyticsInstancesFeatureSetEnum = "SELF_SERVICE_ANALYTICS"
	ListAnalyticsInstancesFeatureSetEnterpriseAnalytics  ListAnalyticsInstancesFeatureSetEnum = "ENTERPRISE_ANALYTICS"
)

var mappingListAnalyticsInstancesFeatureSetEnum = map[string]ListAnalyticsInstancesFeatureSetEnum{
	"SELF_SERVICE_ANALYTICS": ListAnalyticsInstancesFeatureSetSelfServiceAnalytics,
	"ENTERPRISE_ANALYTICS":   ListAnalyticsInstancesFeatureSetEnterpriseAnalytics,
}

var mappingListAnalyticsInstancesFeatureSetEnumLowerCase = map[string]ListAnalyticsInstancesFeatureSetEnum{
	"self_service_analytics": ListAnalyticsInstancesFeatureSetSelfServiceAnalytics,
	"enterprise_analytics":   ListAnalyticsInstancesFeatureSetEnterpriseAnalytics,
}

// GetListAnalyticsInstancesFeatureSetEnumValues Enumerates the set of values for ListAnalyticsInstancesFeatureSetEnum
func GetListAnalyticsInstancesFeatureSetEnumValues() []ListAnalyticsInstancesFeatureSetEnum {
	values := make([]ListAnalyticsInstancesFeatureSetEnum, 0)
	for _, v := range mappingListAnalyticsInstancesFeatureSetEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnalyticsInstancesFeatureSetEnumStringValues Enumerates the set of values in String for ListAnalyticsInstancesFeatureSetEnum
func GetListAnalyticsInstancesFeatureSetEnumStringValues() []string {
	return []string{
		"SELF_SERVICE_ANALYTICS",
		"ENTERPRISE_ANALYTICS",
	}
}

// GetMappingListAnalyticsInstancesFeatureSetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnalyticsInstancesFeatureSetEnum(val string) (ListAnalyticsInstancesFeatureSetEnum, bool) {
	enum, ok := mappingListAnalyticsInstancesFeatureSetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAnalyticsInstancesLifecycleStateEnum Enum with underlying type: string
type ListAnalyticsInstancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListAnalyticsInstancesLifecycleStateEnum
const (
	ListAnalyticsInstancesLifecycleStateActive   ListAnalyticsInstancesLifecycleStateEnum = "ACTIVE"
	ListAnalyticsInstancesLifecycleStateCreating ListAnalyticsInstancesLifecycleStateEnum = "CREATING"
	ListAnalyticsInstancesLifecycleStateDeleted  ListAnalyticsInstancesLifecycleStateEnum = "DELETED"
	ListAnalyticsInstancesLifecycleStateDeleting ListAnalyticsInstancesLifecycleStateEnum = "DELETING"
	ListAnalyticsInstancesLifecycleStateFailed   ListAnalyticsInstancesLifecycleStateEnum = "FAILED"
	ListAnalyticsInstancesLifecycleStateInactive ListAnalyticsInstancesLifecycleStateEnum = "INACTIVE"
	ListAnalyticsInstancesLifecycleStateUpdating ListAnalyticsInstancesLifecycleStateEnum = "UPDATING"
)

var mappingListAnalyticsInstancesLifecycleStateEnum = map[string]ListAnalyticsInstancesLifecycleStateEnum{
	"ACTIVE":   ListAnalyticsInstancesLifecycleStateActive,
	"CREATING": ListAnalyticsInstancesLifecycleStateCreating,
	"DELETED":  ListAnalyticsInstancesLifecycleStateDeleted,
	"DELETING": ListAnalyticsInstancesLifecycleStateDeleting,
	"FAILED":   ListAnalyticsInstancesLifecycleStateFailed,
	"INACTIVE": ListAnalyticsInstancesLifecycleStateInactive,
	"UPDATING": ListAnalyticsInstancesLifecycleStateUpdating,
}

var mappingListAnalyticsInstancesLifecycleStateEnumLowerCase = map[string]ListAnalyticsInstancesLifecycleStateEnum{
	"active":   ListAnalyticsInstancesLifecycleStateActive,
	"creating": ListAnalyticsInstancesLifecycleStateCreating,
	"deleted":  ListAnalyticsInstancesLifecycleStateDeleted,
	"deleting": ListAnalyticsInstancesLifecycleStateDeleting,
	"failed":   ListAnalyticsInstancesLifecycleStateFailed,
	"inactive": ListAnalyticsInstancesLifecycleStateInactive,
	"updating": ListAnalyticsInstancesLifecycleStateUpdating,
}

// GetListAnalyticsInstancesLifecycleStateEnumValues Enumerates the set of values for ListAnalyticsInstancesLifecycleStateEnum
func GetListAnalyticsInstancesLifecycleStateEnumValues() []ListAnalyticsInstancesLifecycleStateEnum {
	values := make([]ListAnalyticsInstancesLifecycleStateEnum, 0)
	for _, v := range mappingListAnalyticsInstancesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnalyticsInstancesLifecycleStateEnumStringValues Enumerates the set of values in String for ListAnalyticsInstancesLifecycleStateEnum
func GetListAnalyticsInstancesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"INACTIVE",
		"UPDATING",
	}
}

// GetMappingListAnalyticsInstancesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnalyticsInstancesLifecycleStateEnum(val string) (ListAnalyticsInstancesLifecycleStateEnum, bool) {
	enum, ok := mappingListAnalyticsInstancesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAnalyticsInstancesSortByEnum Enum with underlying type: string
type ListAnalyticsInstancesSortByEnum string

// Set of constants representing the allowable values for ListAnalyticsInstancesSortByEnum
const (
	ListAnalyticsInstancesSortByCapacitytype   ListAnalyticsInstancesSortByEnum = "capacityType"
	ListAnalyticsInstancesSortByCapacityvalue  ListAnalyticsInstancesSortByEnum = "capacityValue"
	ListAnalyticsInstancesSortByFeatureset     ListAnalyticsInstancesSortByEnum = "featureSet"
	ListAnalyticsInstancesSortByLifecyclestate ListAnalyticsInstancesSortByEnum = "lifecycleState"
	ListAnalyticsInstancesSortByName           ListAnalyticsInstancesSortByEnum = "name"
	ListAnalyticsInstancesSortByTimecreated    ListAnalyticsInstancesSortByEnum = "timeCreated"
)

var mappingListAnalyticsInstancesSortByEnum = map[string]ListAnalyticsInstancesSortByEnum{
	"capacityType":   ListAnalyticsInstancesSortByCapacitytype,
	"capacityValue":  ListAnalyticsInstancesSortByCapacityvalue,
	"featureSet":     ListAnalyticsInstancesSortByFeatureset,
	"lifecycleState": ListAnalyticsInstancesSortByLifecyclestate,
	"name":           ListAnalyticsInstancesSortByName,
	"timeCreated":    ListAnalyticsInstancesSortByTimecreated,
}

var mappingListAnalyticsInstancesSortByEnumLowerCase = map[string]ListAnalyticsInstancesSortByEnum{
	"capacitytype":   ListAnalyticsInstancesSortByCapacitytype,
	"capacityvalue":  ListAnalyticsInstancesSortByCapacityvalue,
	"featureset":     ListAnalyticsInstancesSortByFeatureset,
	"lifecyclestate": ListAnalyticsInstancesSortByLifecyclestate,
	"name":           ListAnalyticsInstancesSortByName,
	"timecreated":    ListAnalyticsInstancesSortByTimecreated,
}

// GetListAnalyticsInstancesSortByEnumValues Enumerates the set of values for ListAnalyticsInstancesSortByEnum
func GetListAnalyticsInstancesSortByEnumValues() []ListAnalyticsInstancesSortByEnum {
	values := make([]ListAnalyticsInstancesSortByEnum, 0)
	for _, v := range mappingListAnalyticsInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnalyticsInstancesSortByEnumStringValues Enumerates the set of values in String for ListAnalyticsInstancesSortByEnum
func GetListAnalyticsInstancesSortByEnumStringValues() []string {
	return []string{
		"capacityType",
		"capacityValue",
		"featureSet",
		"lifecycleState",
		"name",
		"timeCreated",
	}
}

// GetMappingListAnalyticsInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnalyticsInstancesSortByEnum(val string) (ListAnalyticsInstancesSortByEnum, bool) {
	enum, ok := mappingListAnalyticsInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAnalyticsInstancesSortOrderEnum Enum with underlying type: string
type ListAnalyticsInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListAnalyticsInstancesSortOrderEnum
const (
	ListAnalyticsInstancesSortOrderAsc  ListAnalyticsInstancesSortOrderEnum = "ASC"
	ListAnalyticsInstancesSortOrderDesc ListAnalyticsInstancesSortOrderEnum = "DESC"
)

var mappingListAnalyticsInstancesSortOrderEnum = map[string]ListAnalyticsInstancesSortOrderEnum{
	"ASC":  ListAnalyticsInstancesSortOrderAsc,
	"DESC": ListAnalyticsInstancesSortOrderDesc,
}

var mappingListAnalyticsInstancesSortOrderEnumLowerCase = map[string]ListAnalyticsInstancesSortOrderEnum{
	"asc":  ListAnalyticsInstancesSortOrderAsc,
	"desc": ListAnalyticsInstancesSortOrderDesc,
}

// GetListAnalyticsInstancesSortOrderEnumValues Enumerates the set of values for ListAnalyticsInstancesSortOrderEnum
func GetListAnalyticsInstancesSortOrderEnumValues() []ListAnalyticsInstancesSortOrderEnum {
	values := make([]ListAnalyticsInstancesSortOrderEnum, 0)
	for _, v := range mappingListAnalyticsInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnalyticsInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListAnalyticsInstancesSortOrderEnum
func GetListAnalyticsInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAnalyticsInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnalyticsInstancesSortOrderEnum(val string) (ListAnalyticsInstancesSortOrderEnum, bool) {
	enum, ok := mappingListAnalyticsInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
