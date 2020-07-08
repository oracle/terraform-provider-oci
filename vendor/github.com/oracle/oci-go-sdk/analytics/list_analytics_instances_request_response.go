// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package analytics

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAnalyticsInstancesRequest wrapper for the ListAnalyticsInstances operation
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
func (request ListAnalyticsInstancesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAnalyticsInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
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

var mappingListAnalyticsInstancesCapacityType = map[string]ListAnalyticsInstancesCapacityTypeEnum{
	"OLPU_COUNT": ListAnalyticsInstancesCapacityTypeOlpuCount,
	"USER_COUNT": ListAnalyticsInstancesCapacityTypeUserCount,
}

// GetListAnalyticsInstancesCapacityTypeEnumValues Enumerates the set of values for ListAnalyticsInstancesCapacityTypeEnum
func GetListAnalyticsInstancesCapacityTypeEnumValues() []ListAnalyticsInstancesCapacityTypeEnum {
	values := make([]ListAnalyticsInstancesCapacityTypeEnum, 0)
	for _, v := range mappingListAnalyticsInstancesCapacityType {
		values = append(values, v)
	}
	return values
}

// ListAnalyticsInstancesFeatureSetEnum Enum with underlying type: string
type ListAnalyticsInstancesFeatureSetEnum string

// Set of constants representing the allowable values for ListAnalyticsInstancesFeatureSetEnum
const (
	ListAnalyticsInstancesFeatureSetSelfServiceAnalytics ListAnalyticsInstancesFeatureSetEnum = "SELF_SERVICE_ANALYTICS"
	ListAnalyticsInstancesFeatureSetEnterpriseAnalytics  ListAnalyticsInstancesFeatureSetEnum = "ENTERPRISE_ANALYTICS"
)

var mappingListAnalyticsInstancesFeatureSet = map[string]ListAnalyticsInstancesFeatureSetEnum{
	"SELF_SERVICE_ANALYTICS": ListAnalyticsInstancesFeatureSetSelfServiceAnalytics,
	"ENTERPRISE_ANALYTICS":   ListAnalyticsInstancesFeatureSetEnterpriseAnalytics,
}

// GetListAnalyticsInstancesFeatureSetEnumValues Enumerates the set of values for ListAnalyticsInstancesFeatureSetEnum
func GetListAnalyticsInstancesFeatureSetEnumValues() []ListAnalyticsInstancesFeatureSetEnum {
	values := make([]ListAnalyticsInstancesFeatureSetEnum, 0)
	for _, v := range mappingListAnalyticsInstancesFeatureSet {
		values = append(values, v)
	}
	return values
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

var mappingListAnalyticsInstancesLifecycleState = map[string]ListAnalyticsInstancesLifecycleStateEnum{
	"ACTIVE":   ListAnalyticsInstancesLifecycleStateActive,
	"CREATING": ListAnalyticsInstancesLifecycleStateCreating,
	"DELETED":  ListAnalyticsInstancesLifecycleStateDeleted,
	"DELETING": ListAnalyticsInstancesLifecycleStateDeleting,
	"FAILED":   ListAnalyticsInstancesLifecycleStateFailed,
	"INACTIVE": ListAnalyticsInstancesLifecycleStateInactive,
	"UPDATING": ListAnalyticsInstancesLifecycleStateUpdating,
}

// GetListAnalyticsInstancesLifecycleStateEnumValues Enumerates the set of values for ListAnalyticsInstancesLifecycleStateEnum
func GetListAnalyticsInstancesLifecycleStateEnumValues() []ListAnalyticsInstancesLifecycleStateEnum {
	values := make([]ListAnalyticsInstancesLifecycleStateEnum, 0)
	for _, v := range mappingListAnalyticsInstancesLifecycleState {
		values = append(values, v)
	}
	return values
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

var mappingListAnalyticsInstancesSortBy = map[string]ListAnalyticsInstancesSortByEnum{
	"capacityType":   ListAnalyticsInstancesSortByCapacitytype,
	"capacityValue":  ListAnalyticsInstancesSortByCapacityvalue,
	"featureSet":     ListAnalyticsInstancesSortByFeatureset,
	"lifecycleState": ListAnalyticsInstancesSortByLifecyclestate,
	"name":           ListAnalyticsInstancesSortByName,
	"timeCreated":    ListAnalyticsInstancesSortByTimecreated,
}

// GetListAnalyticsInstancesSortByEnumValues Enumerates the set of values for ListAnalyticsInstancesSortByEnum
func GetListAnalyticsInstancesSortByEnumValues() []ListAnalyticsInstancesSortByEnum {
	values := make([]ListAnalyticsInstancesSortByEnum, 0)
	for _, v := range mappingListAnalyticsInstancesSortBy {
		values = append(values, v)
	}
	return values
}

// ListAnalyticsInstancesSortOrderEnum Enum with underlying type: string
type ListAnalyticsInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListAnalyticsInstancesSortOrderEnum
const (
	ListAnalyticsInstancesSortOrderAsc  ListAnalyticsInstancesSortOrderEnum = "ASC"
	ListAnalyticsInstancesSortOrderDesc ListAnalyticsInstancesSortOrderEnum = "DESC"
)

var mappingListAnalyticsInstancesSortOrder = map[string]ListAnalyticsInstancesSortOrderEnum{
	"ASC":  ListAnalyticsInstancesSortOrderAsc,
	"DESC": ListAnalyticsInstancesSortOrderDesc,
}

// GetListAnalyticsInstancesSortOrderEnumValues Enumerates the set of values for ListAnalyticsInstancesSortOrderEnum
func GetListAnalyticsInstancesSortOrderEnumValues() []ListAnalyticsInstancesSortOrderEnum {
	values := make([]ListAnalyticsInstancesSortOrderEnum, 0)
	for _, v := range mappingListAnalyticsInstancesSortOrder {
		values = append(values, v)
	}
	return values
}
