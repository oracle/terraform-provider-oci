// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListManagementAgentsRequest wrapper for the ListManagementAgents operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementagent/ListManagementAgents.go.html to see an example of how to use ListManagementAgentsRequest.
type ListManagementAgentsRequest struct {

	// The OCID of the compartment to which a request will be scoped.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter to return only Management Agents having the particular Plugin installed. A special pluginName of 'None' can be provided and this will return only Management Agents having no plugin installed.
	PluginName []string `contributesTo:"query" name:"pluginName" collectionFormat:"multi"`

	// Filter to return only Management Agents having the particular agent version.
	Version []string `contributesTo:"query" name:"version" collectionFormat:"multi"`

	// Filter to return only Management Agents having the particular display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter to return only Management Agents in the particular lifecycle state.
	LifecycleState ListManagementAgentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter to return only Management Agents in the particular availability status.
	AvailabilityStatus ListManagementAgentsAvailabilityStatusEnum `mandatory:"false" contributesTo:"query" name:"availabilityStatus" omitEmpty:"true"`

	// Filter to return only Management Agents having the particular agent host id.
	HostId *string `mandatory:"false" contributesTo:"query" name:"hostId"`

	// Filter to return only results having the particular platform type.
	PlatformType []PlatformTypesEnum `contributesTo:"query" name:"platformType" omitEmpty:"true" collectionFormat:"multi"`

	// true, if the agent image is manually downloaded and installed. false, if the agent is deployed as a plugin in Oracle Cloud Agent.
	IsCustomerDeployed *bool `mandatory:"false" contributesTo:"query" name:"isCustomerDeployed"`

	// A filter to return either agents or gateway types depending upon install type selected by user. By default both install type will be returned.
	InstallType ListManagementAgentsInstallTypeEnum `mandatory:"false" contributesTo:"query" name:"installType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagementAgentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListManagementAgentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementAgentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementAgentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagementAgentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementAgentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListManagementAgentsResponse wrapper for the ListManagementAgents operation
type ListManagementAgentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ManagementAgentSummary instances
	Items []ManagementAgentSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagementAgentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementAgentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementAgentsLifecycleStateEnum Enum with underlying type: string
type ListManagementAgentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagementAgentsLifecycleStateEnum
const (
	ListManagementAgentsLifecycleStateCreating   ListManagementAgentsLifecycleStateEnum = "CREATING"
	ListManagementAgentsLifecycleStateUpdating   ListManagementAgentsLifecycleStateEnum = "UPDATING"
	ListManagementAgentsLifecycleStateActive     ListManagementAgentsLifecycleStateEnum = "ACTIVE"
	ListManagementAgentsLifecycleStateInactive   ListManagementAgentsLifecycleStateEnum = "INACTIVE"
	ListManagementAgentsLifecycleStateTerminated ListManagementAgentsLifecycleStateEnum = "TERMINATED"
	ListManagementAgentsLifecycleStateDeleting   ListManagementAgentsLifecycleStateEnum = "DELETING"
	ListManagementAgentsLifecycleStateDeleted    ListManagementAgentsLifecycleStateEnum = "DELETED"
	ListManagementAgentsLifecycleStateFailed     ListManagementAgentsLifecycleStateEnum = "FAILED"
)

var mappingListManagementAgentsLifecycleState = map[string]ListManagementAgentsLifecycleStateEnum{
	"CREATING":   ListManagementAgentsLifecycleStateCreating,
	"UPDATING":   ListManagementAgentsLifecycleStateUpdating,
	"ACTIVE":     ListManagementAgentsLifecycleStateActive,
	"INACTIVE":   ListManagementAgentsLifecycleStateInactive,
	"TERMINATED": ListManagementAgentsLifecycleStateTerminated,
	"DELETING":   ListManagementAgentsLifecycleStateDeleting,
	"DELETED":    ListManagementAgentsLifecycleStateDeleted,
	"FAILED":     ListManagementAgentsLifecycleStateFailed,
}

// GetListManagementAgentsLifecycleStateEnumValues Enumerates the set of values for ListManagementAgentsLifecycleStateEnum
func GetListManagementAgentsLifecycleStateEnumValues() []ListManagementAgentsLifecycleStateEnum {
	values := make([]ListManagementAgentsLifecycleStateEnum, 0)
	for _, v := range mappingListManagementAgentsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentsAvailabilityStatusEnum Enum with underlying type: string
type ListManagementAgentsAvailabilityStatusEnum string

// Set of constants representing the allowable values for ListManagementAgentsAvailabilityStatusEnum
const (
	ListManagementAgentsAvailabilityStatusActive       ListManagementAgentsAvailabilityStatusEnum = "ACTIVE"
	ListManagementAgentsAvailabilityStatusSilent       ListManagementAgentsAvailabilityStatusEnum = "SILENT"
	ListManagementAgentsAvailabilityStatusNotAvailable ListManagementAgentsAvailabilityStatusEnum = "NOT_AVAILABLE"
)

var mappingListManagementAgentsAvailabilityStatus = map[string]ListManagementAgentsAvailabilityStatusEnum{
	"ACTIVE":        ListManagementAgentsAvailabilityStatusActive,
	"SILENT":        ListManagementAgentsAvailabilityStatusSilent,
	"NOT_AVAILABLE": ListManagementAgentsAvailabilityStatusNotAvailable,
}

// GetListManagementAgentsAvailabilityStatusEnumValues Enumerates the set of values for ListManagementAgentsAvailabilityStatusEnum
func GetListManagementAgentsAvailabilityStatusEnumValues() []ListManagementAgentsAvailabilityStatusEnum {
	values := make([]ListManagementAgentsAvailabilityStatusEnum, 0)
	for _, v := range mappingListManagementAgentsAvailabilityStatus {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentsInstallTypeEnum Enum with underlying type: string
type ListManagementAgentsInstallTypeEnum string

// Set of constants representing the allowable values for ListManagementAgentsInstallTypeEnum
const (
	ListManagementAgentsInstallTypeAgent   ListManagementAgentsInstallTypeEnum = "AGENT"
	ListManagementAgentsInstallTypeGateway ListManagementAgentsInstallTypeEnum = "GATEWAY"
)

var mappingListManagementAgentsInstallType = map[string]ListManagementAgentsInstallTypeEnum{
	"AGENT":   ListManagementAgentsInstallTypeAgent,
	"GATEWAY": ListManagementAgentsInstallTypeGateway,
}

// GetListManagementAgentsInstallTypeEnumValues Enumerates the set of values for ListManagementAgentsInstallTypeEnum
func GetListManagementAgentsInstallTypeEnumValues() []ListManagementAgentsInstallTypeEnum {
	values := make([]ListManagementAgentsInstallTypeEnum, 0)
	for _, v := range mappingListManagementAgentsInstallType {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentsSortOrderEnum Enum with underlying type: string
type ListManagementAgentsSortOrderEnum string

// Set of constants representing the allowable values for ListManagementAgentsSortOrderEnum
const (
	ListManagementAgentsSortOrderAsc  ListManagementAgentsSortOrderEnum = "ASC"
	ListManagementAgentsSortOrderDesc ListManagementAgentsSortOrderEnum = "DESC"
)

var mappingListManagementAgentsSortOrder = map[string]ListManagementAgentsSortOrderEnum{
	"ASC":  ListManagementAgentsSortOrderAsc,
	"DESC": ListManagementAgentsSortOrderDesc,
}

// GetListManagementAgentsSortOrderEnumValues Enumerates the set of values for ListManagementAgentsSortOrderEnum
func GetListManagementAgentsSortOrderEnumValues() []ListManagementAgentsSortOrderEnum {
	values := make([]ListManagementAgentsSortOrderEnum, 0)
	for _, v := range mappingListManagementAgentsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentsSortByEnum Enum with underlying type: string
type ListManagementAgentsSortByEnum string

// Set of constants representing the allowable values for ListManagementAgentsSortByEnum
const (
	ListManagementAgentsSortByTimecreated        ListManagementAgentsSortByEnum = "timeCreated"
	ListManagementAgentsSortByDisplayname        ListManagementAgentsSortByEnum = "displayName"
	ListManagementAgentsSortByHost               ListManagementAgentsSortByEnum = "host"
	ListManagementAgentsSortByAvailabilitystatus ListManagementAgentsSortByEnum = "availabilityStatus"
	ListManagementAgentsSortByPlatformtype       ListManagementAgentsSortByEnum = "platformType"
	ListManagementAgentsSortByPlugindisplaynames ListManagementAgentsSortByEnum = "pluginDisplayNames"
	ListManagementAgentsSortByVersion            ListManagementAgentsSortByEnum = "version"
)

var mappingListManagementAgentsSortBy = map[string]ListManagementAgentsSortByEnum{
	"timeCreated":        ListManagementAgentsSortByTimecreated,
	"displayName":        ListManagementAgentsSortByDisplayname,
	"host":               ListManagementAgentsSortByHost,
	"availabilityStatus": ListManagementAgentsSortByAvailabilitystatus,
	"platformType":       ListManagementAgentsSortByPlatformtype,
	"pluginDisplayNames": ListManagementAgentsSortByPlugindisplaynames,
	"version":            ListManagementAgentsSortByVersion,
}

// GetListManagementAgentsSortByEnumValues Enumerates the set of values for ListManagementAgentsSortByEnum
func GetListManagementAgentsSortByEnumValues() []ListManagementAgentsSortByEnum {
	values := make([]ListManagementAgentsSortByEnum, 0)
	for _, v := range mappingListManagementAgentsSortBy {
		values = append(values, v)
	}
	return values
}
