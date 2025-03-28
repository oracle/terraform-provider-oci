// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagementAgentsRequest wrapper for the ListManagementAgents operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementagent/ListManagementAgents.go.html to see an example of how to use ListManagementAgentsRequest.
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

	// Filter to return only results having the particular gatewayId.
	GatewayId []string `contributesTo:"query" name:"gatewayId" collectionFormat:"multi"`

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

	// if set to true then it fetches resources for all compartments where user has access to else only on the compartment specified.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// When the value is "ACCESSIBLE", insufficient permissions for a compartment will filter out resources in that compartment without rejecting the request.
	AccessLevel *string `mandatory:"false" contributesTo:"query" name:"accessLevel"`

	// The type of the dataSource.
	DataSourceType ListManagementAgentsDataSourceTypeEnum `mandatory:"false" contributesTo:"query" name:"dataSourceType" omitEmpty:"true"`

	// Unique name of the dataSource.
	DataSourceName []string `contributesTo:"query" name:"dataSourceName" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementAgentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementAgentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagementAgentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagementAgentsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListManagementAgentsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAgentsAvailabilityStatusEnum(string(request.AvailabilityStatus)); !ok && request.AvailabilityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailabilityStatus: %s. Supported values are: %s.", request.AvailabilityStatus, strings.Join(GetListManagementAgentsAvailabilityStatusEnumStringValues(), ",")))
	}
	for _, val := range request.PlatformType {
		if _, ok := GetMappingPlatformTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", val, strings.Join(GetPlatformTypesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListManagementAgentsInstallTypeEnum(string(request.InstallType)); !ok && request.InstallType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstallType: %s. Supported values are: %s.", request.InstallType, strings.Join(GetListManagementAgentsInstallTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAgentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagementAgentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAgentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagementAgentsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAgentsDataSourceTypeEnum(string(request.DataSourceType)); !ok && request.DataSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataSourceType: %s. Supported values are: %s.", request.DataSourceType, strings.Join(GetListManagementAgentsDataSourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListManagementAgentsLifecycleStateEnum = map[string]ListManagementAgentsLifecycleStateEnum{
	"CREATING":   ListManagementAgentsLifecycleStateCreating,
	"UPDATING":   ListManagementAgentsLifecycleStateUpdating,
	"ACTIVE":     ListManagementAgentsLifecycleStateActive,
	"INACTIVE":   ListManagementAgentsLifecycleStateInactive,
	"TERMINATED": ListManagementAgentsLifecycleStateTerminated,
	"DELETING":   ListManagementAgentsLifecycleStateDeleting,
	"DELETED":    ListManagementAgentsLifecycleStateDeleted,
	"FAILED":     ListManagementAgentsLifecycleStateFailed,
}

var mappingListManagementAgentsLifecycleStateEnumLowerCase = map[string]ListManagementAgentsLifecycleStateEnum{
	"creating":   ListManagementAgentsLifecycleStateCreating,
	"updating":   ListManagementAgentsLifecycleStateUpdating,
	"active":     ListManagementAgentsLifecycleStateActive,
	"inactive":   ListManagementAgentsLifecycleStateInactive,
	"terminated": ListManagementAgentsLifecycleStateTerminated,
	"deleting":   ListManagementAgentsLifecycleStateDeleting,
	"deleted":    ListManagementAgentsLifecycleStateDeleted,
	"failed":     ListManagementAgentsLifecycleStateFailed,
}

// GetListManagementAgentsLifecycleStateEnumValues Enumerates the set of values for ListManagementAgentsLifecycleStateEnum
func GetListManagementAgentsLifecycleStateEnumValues() []ListManagementAgentsLifecycleStateEnum {
	values := make([]ListManagementAgentsLifecycleStateEnum, 0)
	for _, v := range mappingListManagementAgentsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentsLifecycleStateEnumStringValues Enumerates the set of values in String for ListManagementAgentsLifecycleStateEnum
func GetListManagementAgentsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"TERMINATED",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListManagementAgentsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentsLifecycleStateEnum(val string) (ListManagementAgentsLifecycleStateEnum, bool) {
	enum, ok := mappingListManagementAgentsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAgentsAvailabilityStatusEnum Enum with underlying type: string
type ListManagementAgentsAvailabilityStatusEnum string

// Set of constants representing the allowable values for ListManagementAgentsAvailabilityStatusEnum
const (
	ListManagementAgentsAvailabilityStatusActive       ListManagementAgentsAvailabilityStatusEnum = "ACTIVE"
	ListManagementAgentsAvailabilityStatusSilent       ListManagementAgentsAvailabilityStatusEnum = "SILENT"
	ListManagementAgentsAvailabilityStatusNotAvailable ListManagementAgentsAvailabilityStatusEnum = "NOT_AVAILABLE"
)

var mappingListManagementAgentsAvailabilityStatusEnum = map[string]ListManagementAgentsAvailabilityStatusEnum{
	"ACTIVE":        ListManagementAgentsAvailabilityStatusActive,
	"SILENT":        ListManagementAgentsAvailabilityStatusSilent,
	"NOT_AVAILABLE": ListManagementAgentsAvailabilityStatusNotAvailable,
}

var mappingListManagementAgentsAvailabilityStatusEnumLowerCase = map[string]ListManagementAgentsAvailabilityStatusEnum{
	"active":        ListManagementAgentsAvailabilityStatusActive,
	"silent":        ListManagementAgentsAvailabilityStatusSilent,
	"not_available": ListManagementAgentsAvailabilityStatusNotAvailable,
}

// GetListManagementAgentsAvailabilityStatusEnumValues Enumerates the set of values for ListManagementAgentsAvailabilityStatusEnum
func GetListManagementAgentsAvailabilityStatusEnumValues() []ListManagementAgentsAvailabilityStatusEnum {
	values := make([]ListManagementAgentsAvailabilityStatusEnum, 0)
	for _, v := range mappingListManagementAgentsAvailabilityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentsAvailabilityStatusEnumStringValues Enumerates the set of values in String for ListManagementAgentsAvailabilityStatusEnum
func GetListManagementAgentsAvailabilityStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"SILENT",
		"NOT_AVAILABLE",
	}
}

// GetMappingListManagementAgentsAvailabilityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentsAvailabilityStatusEnum(val string) (ListManagementAgentsAvailabilityStatusEnum, bool) {
	enum, ok := mappingListManagementAgentsAvailabilityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAgentsInstallTypeEnum Enum with underlying type: string
type ListManagementAgentsInstallTypeEnum string

// Set of constants representing the allowable values for ListManagementAgentsInstallTypeEnum
const (
	ListManagementAgentsInstallTypeAgent   ListManagementAgentsInstallTypeEnum = "AGENT"
	ListManagementAgentsInstallTypeGateway ListManagementAgentsInstallTypeEnum = "GATEWAY"
)

var mappingListManagementAgentsInstallTypeEnum = map[string]ListManagementAgentsInstallTypeEnum{
	"AGENT":   ListManagementAgentsInstallTypeAgent,
	"GATEWAY": ListManagementAgentsInstallTypeGateway,
}

var mappingListManagementAgentsInstallTypeEnumLowerCase = map[string]ListManagementAgentsInstallTypeEnum{
	"agent":   ListManagementAgentsInstallTypeAgent,
	"gateway": ListManagementAgentsInstallTypeGateway,
}

// GetListManagementAgentsInstallTypeEnumValues Enumerates the set of values for ListManagementAgentsInstallTypeEnum
func GetListManagementAgentsInstallTypeEnumValues() []ListManagementAgentsInstallTypeEnum {
	values := make([]ListManagementAgentsInstallTypeEnum, 0)
	for _, v := range mappingListManagementAgentsInstallTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentsInstallTypeEnumStringValues Enumerates the set of values in String for ListManagementAgentsInstallTypeEnum
func GetListManagementAgentsInstallTypeEnumStringValues() []string {
	return []string{
		"AGENT",
		"GATEWAY",
	}
}

// GetMappingListManagementAgentsInstallTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentsInstallTypeEnum(val string) (ListManagementAgentsInstallTypeEnum, bool) {
	enum, ok := mappingListManagementAgentsInstallTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAgentsSortOrderEnum Enum with underlying type: string
type ListManagementAgentsSortOrderEnum string

// Set of constants representing the allowable values for ListManagementAgentsSortOrderEnum
const (
	ListManagementAgentsSortOrderAsc  ListManagementAgentsSortOrderEnum = "ASC"
	ListManagementAgentsSortOrderDesc ListManagementAgentsSortOrderEnum = "DESC"
)

var mappingListManagementAgentsSortOrderEnum = map[string]ListManagementAgentsSortOrderEnum{
	"ASC":  ListManagementAgentsSortOrderAsc,
	"DESC": ListManagementAgentsSortOrderDesc,
}

var mappingListManagementAgentsSortOrderEnumLowerCase = map[string]ListManagementAgentsSortOrderEnum{
	"asc":  ListManagementAgentsSortOrderAsc,
	"desc": ListManagementAgentsSortOrderDesc,
}

// GetListManagementAgentsSortOrderEnumValues Enumerates the set of values for ListManagementAgentsSortOrderEnum
func GetListManagementAgentsSortOrderEnumValues() []ListManagementAgentsSortOrderEnum {
	values := make([]ListManagementAgentsSortOrderEnum, 0)
	for _, v := range mappingListManagementAgentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentsSortOrderEnumStringValues Enumerates the set of values in String for ListManagementAgentsSortOrderEnum
func GetListManagementAgentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagementAgentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentsSortOrderEnum(val string) (ListManagementAgentsSortOrderEnum, bool) {
	enum, ok := mappingListManagementAgentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingListManagementAgentsSortByEnum = map[string]ListManagementAgentsSortByEnum{
	"timeCreated":        ListManagementAgentsSortByTimecreated,
	"displayName":        ListManagementAgentsSortByDisplayname,
	"host":               ListManagementAgentsSortByHost,
	"availabilityStatus": ListManagementAgentsSortByAvailabilitystatus,
	"platformType":       ListManagementAgentsSortByPlatformtype,
	"pluginDisplayNames": ListManagementAgentsSortByPlugindisplaynames,
	"version":            ListManagementAgentsSortByVersion,
}

var mappingListManagementAgentsSortByEnumLowerCase = map[string]ListManagementAgentsSortByEnum{
	"timecreated":        ListManagementAgentsSortByTimecreated,
	"displayname":        ListManagementAgentsSortByDisplayname,
	"host":               ListManagementAgentsSortByHost,
	"availabilitystatus": ListManagementAgentsSortByAvailabilitystatus,
	"platformtype":       ListManagementAgentsSortByPlatformtype,
	"plugindisplaynames": ListManagementAgentsSortByPlugindisplaynames,
	"version":            ListManagementAgentsSortByVersion,
}

// GetListManagementAgentsSortByEnumValues Enumerates the set of values for ListManagementAgentsSortByEnum
func GetListManagementAgentsSortByEnumValues() []ListManagementAgentsSortByEnum {
	values := make([]ListManagementAgentsSortByEnum, 0)
	for _, v := range mappingListManagementAgentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentsSortByEnumStringValues Enumerates the set of values in String for ListManagementAgentsSortByEnum
func GetListManagementAgentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"host",
		"availabilityStatus",
		"platformType",
		"pluginDisplayNames",
		"version",
	}
}

// GetMappingListManagementAgentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentsSortByEnum(val string) (ListManagementAgentsSortByEnum, bool) {
	enum, ok := mappingListManagementAgentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAgentsDataSourceTypeEnum Enum with underlying type: string
type ListManagementAgentsDataSourceTypeEnum string

// Set of constants representing the allowable values for ListManagementAgentsDataSourceTypeEnum
const (
	ListManagementAgentsDataSourceTypeKubernetesCluster ListManagementAgentsDataSourceTypeEnum = "KUBERNETES_CLUSTER"
	ListManagementAgentsDataSourceTypePrometheusEmitter ListManagementAgentsDataSourceTypeEnum = "PROMETHEUS_EMITTER"
)

var mappingListManagementAgentsDataSourceTypeEnum = map[string]ListManagementAgentsDataSourceTypeEnum{
	"KUBERNETES_CLUSTER": ListManagementAgentsDataSourceTypeKubernetesCluster,
	"PROMETHEUS_EMITTER": ListManagementAgentsDataSourceTypePrometheusEmitter,
}

var mappingListManagementAgentsDataSourceTypeEnumLowerCase = map[string]ListManagementAgentsDataSourceTypeEnum{
	"kubernetes_cluster": ListManagementAgentsDataSourceTypeKubernetesCluster,
	"prometheus_emitter": ListManagementAgentsDataSourceTypePrometheusEmitter,
}

// GetListManagementAgentsDataSourceTypeEnumValues Enumerates the set of values for ListManagementAgentsDataSourceTypeEnum
func GetListManagementAgentsDataSourceTypeEnumValues() []ListManagementAgentsDataSourceTypeEnum {
	values := make([]ListManagementAgentsDataSourceTypeEnum, 0)
	for _, v := range mappingListManagementAgentsDataSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentsDataSourceTypeEnumStringValues Enumerates the set of values in String for ListManagementAgentsDataSourceTypeEnum
func GetListManagementAgentsDataSourceTypeEnumStringValues() []string {
	return []string{
		"KUBERNETES_CLUSTER",
		"PROMETHEUS_EMITTER",
	}
}

// GetMappingListManagementAgentsDataSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentsDataSourceTypeEnum(val string) (ListManagementAgentsDataSourceTypeEnum, bool) {
	enum, ok := mappingListManagementAgentsDataSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
