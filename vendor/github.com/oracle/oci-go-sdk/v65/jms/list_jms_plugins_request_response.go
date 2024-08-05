// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJmsPluginsRequest wrapper for the ListJmsPlugins operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListJmsPlugins.go.html to see an example of how to use ListJmsPluginsRequest.
type ListJmsPluginsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Flag to determine whether the info should be gathered only in the compartment or in the compartment and its subcompartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the JmsPlugin.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The ID of the Fleet.
	FleetId *string `mandatory:"false" contributesTo:"query" name:"fleetId"`

	// The ManagementAgent (OMA) or Instance (OCA) OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) that identifies the Agent.
	AgentId *string `mandatory:"false" contributesTo:"query" name:"agentId"`

	// Filter JmsPlugin with its lifecycle state.
	LifecycleState ListJmsPluginsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter JmsPlugin with its availability status.
	AvailabilityStatus ListJmsPluginsAvailabilityStatusEnum `mandatory:"false" contributesTo:"query" name:"availabilityStatus" omitEmpty:"true"`

	// If present, only plugins with a registration time before this parameter are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeRegisteredLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeRegisteredLessThanOrEqualTo"`

	// If present, only plugins with a last seen time before this parameter are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeLastSeenLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastSeenLessThanOrEqualTo"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListJmsPluginsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort JmsPlugin. Only one sort order may be provided.
	// Default order is **descending**.
	// If no value is specified _timeLastSeen_ is default.
	SortBy ListJmsPluginsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter the list with hostname contains the given value.
	HostnameContains *string `mandatory:"false" contributesTo:"query" name:"hostnameContains"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJmsPluginsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJmsPluginsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJmsPluginsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJmsPluginsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJmsPluginsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJmsPluginsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListJmsPluginsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJmsPluginsAvailabilityStatusEnum(string(request.AvailabilityStatus)); !ok && request.AvailabilityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailabilityStatus: %s. Supported values are: %s.", request.AvailabilityStatus, strings.Join(GetListJmsPluginsAvailabilityStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJmsPluginsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJmsPluginsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJmsPluginsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJmsPluginsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJmsPluginsResponse wrapper for the ListJmsPlugins operation
type ListJmsPluginsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JmsPluginCollection instances
	JmsPluginCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJmsPluginsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJmsPluginsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJmsPluginsLifecycleStateEnum Enum with underlying type: string
type ListJmsPluginsLifecycleStateEnum string

// Set of constants representing the allowable values for ListJmsPluginsLifecycleStateEnum
const (
	ListJmsPluginsLifecycleStateActive         ListJmsPluginsLifecycleStateEnum = "ACTIVE"
	ListJmsPluginsLifecycleStateInactive       ListJmsPluginsLifecycleStateEnum = "INACTIVE"
	ListJmsPluginsLifecycleStateNeedsAttention ListJmsPluginsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListJmsPluginsLifecycleStateDeleted        ListJmsPluginsLifecycleStateEnum = "DELETED"
)

var mappingListJmsPluginsLifecycleStateEnum = map[string]ListJmsPluginsLifecycleStateEnum{
	"ACTIVE":          ListJmsPluginsLifecycleStateActive,
	"INACTIVE":        ListJmsPluginsLifecycleStateInactive,
	"NEEDS_ATTENTION": ListJmsPluginsLifecycleStateNeedsAttention,
	"DELETED":         ListJmsPluginsLifecycleStateDeleted,
}

var mappingListJmsPluginsLifecycleStateEnumLowerCase = map[string]ListJmsPluginsLifecycleStateEnum{
	"active":          ListJmsPluginsLifecycleStateActive,
	"inactive":        ListJmsPluginsLifecycleStateInactive,
	"needs_attention": ListJmsPluginsLifecycleStateNeedsAttention,
	"deleted":         ListJmsPluginsLifecycleStateDeleted,
}

// GetListJmsPluginsLifecycleStateEnumValues Enumerates the set of values for ListJmsPluginsLifecycleStateEnum
func GetListJmsPluginsLifecycleStateEnumValues() []ListJmsPluginsLifecycleStateEnum {
	values := make([]ListJmsPluginsLifecycleStateEnum, 0)
	for _, v := range mappingListJmsPluginsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListJmsPluginsLifecycleStateEnumStringValues Enumerates the set of values in String for ListJmsPluginsLifecycleStateEnum
func GetListJmsPluginsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"NEEDS_ATTENTION",
		"DELETED",
	}
}

// GetMappingListJmsPluginsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJmsPluginsLifecycleStateEnum(val string) (ListJmsPluginsLifecycleStateEnum, bool) {
	enum, ok := mappingListJmsPluginsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJmsPluginsAvailabilityStatusEnum Enum with underlying type: string
type ListJmsPluginsAvailabilityStatusEnum string

// Set of constants representing the allowable values for ListJmsPluginsAvailabilityStatusEnum
const (
	ListJmsPluginsAvailabilityStatusActive       ListJmsPluginsAvailabilityStatusEnum = "ACTIVE"
	ListJmsPluginsAvailabilityStatusSilent       ListJmsPluginsAvailabilityStatusEnum = "SILENT"
	ListJmsPluginsAvailabilityStatusNotAvailable ListJmsPluginsAvailabilityStatusEnum = "NOT_AVAILABLE"
)

var mappingListJmsPluginsAvailabilityStatusEnum = map[string]ListJmsPluginsAvailabilityStatusEnum{
	"ACTIVE":        ListJmsPluginsAvailabilityStatusActive,
	"SILENT":        ListJmsPluginsAvailabilityStatusSilent,
	"NOT_AVAILABLE": ListJmsPluginsAvailabilityStatusNotAvailable,
}

var mappingListJmsPluginsAvailabilityStatusEnumLowerCase = map[string]ListJmsPluginsAvailabilityStatusEnum{
	"active":        ListJmsPluginsAvailabilityStatusActive,
	"silent":        ListJmsPluginsAvailabilityStatusSilent,
	"not_available": ListJmsPluginsAvailabilityStatusNotAvailable,
}

// GetListJmsPluginsAvailabilityStatusEnumValues Enumerates the set of values for ListJmsPluginsAvailabilityStatusEnum
func GetListJmsPluginsAvailabilityStatusEnumValues() []ListJmsPluginsAvailabilityStatusEnum {
	values := make([]ListJmsPluginsAvailabilityStatusEnum, 0)
	for _, v := range mappingListJmsPluginsAvailabilityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListJmsPluginsAvailabilityStatusEnumStringValues Enumerates the set of values in String for ListJmsPluginsAvailabilityStatusEnum
func GetListJmsPluginsAvailabilityStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"SILENT",
		"NOT_AVAILABLE",
	}
}

// GetMappingListJmsPluginsAvailabilityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJmsPluginsAvailabilityStatusEnum(val string) (ListJmsPluginsAvailabilityStatusEnum, bool) {
	enum, ok := mappingListJmsPluginsAvailabilityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJmsPluginsSortOrderEnum Enum with underlying type: string
type ListJmsPluginsSortOrderEnum string

// Set of constants representing the allowable values for ListJmsPluginsSortOrderEnum
const (
	ListJmsPluginsSortOrderAsc  ListJmsPluginsSortOrderEnum = "ASC"
	ListJmsPluginsSortOrderDesc ListJmsPluginsSortOrderEnum = "DESC"
)

var mappingListJmsPluginsSortOrderEnum = map[string]ListJmsPluginsSortOrderEnum{
	"ASC":  ListJmsPluginsSortOrderAsc,
	"DESC": ListJmsPluginsSortOrderDesc,
}

var mappingListJmsPluginsSortOrderEnumLowerCase = map[string]ListJmsPluginsSortOrderEnum{
	"asc":  ListJmsPluginsSortOrderAsc,
	"desc": ListJmsPluginsSortOrderDesc,
}

// GetListJmsPluginsSortOrderEnumValues Enumerates the set of values for ListJmsPluginsSortOrderEnum
func GetListJmsPluginsSortOrderEnumValues() []ListJmsPluginsSortOrderEnum {
	values := make([]ListJmsPluginsSortOrderEnum, 0)
	for _, v := range mappingListJmsPluginsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJmsPluginsSortOrderEnumStringValues Enumerates the set of values in String for ListJmsPluginsSortOrderEnum
func GetListJmsPluginsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJmsPluginsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJmsPluginsSortOrderEnum(val string) (ListJmsPluginsSortOrderEnum, bool) {
	enum, ok := mappingListJmsPluginsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJmsPluginsSortByEnum Enum with underlying type: string
type ListJmsPluginsSortByEnum string

// Set of constants representing the allowable values for ListJmsPluginsSortByEnum
const (
	ListJmsPluginsSortById                 ListJmsPluginsSortByEnum = "id"
	ListJmsPluginsSortByTimelastseen       ListJmsPluginsSortByEnum = "timeLastSeen"
	ListJmsPluginsSortByTimeregistered     ListJmsPluginsSortByEnum = "timeRegistered"
	ListJmsPluginsSortByHostname           ListJmsPluginsSortByEnum = "hostname"
	ListJmsPluginsSortByAgentid            ListJmsPluginsSortByEnum = "agentId"
	ListJmsPluginsSortByAgenttype          ListJmsPluginsSortByEnum = "agentType"
	ListJmsPluginsSortByLifecyclestate     ListJmsPluginsSortByEnum = "lifecycleState"
	ListJmsPluginsSortByAvailabilitystatus ListJmsPluginsSortByEnum = "availabilityStatus"
	ListJmsPluginsSortByFleetid            ListJmsPluginsSortByEnum = "fleetId"
	ListJmsPluginsSortByCompartmentid      ListJmsPluginsSortByEnum = "compartmentId"
	ListJmsPluginsSortByOsfamily           ListJmsPluginsSortByEnum = "osFamily"
	ListJmsPluginsSortByOsarchitecture     ListJmsPluginsSortByEnum = "osArchitecture"
	ListJmsPluginsSortByOsdistribution     ListJmsPluginsSortByEnum = "osDistribution"
	ListJmsPluginsSortByPluginversion      ListJmsPluginsSortByEnum = "pluginVersion"
)

var mappingListJmsPluginsSortByEnum = map[string]ListJmsPluginsSortByEnum{
	"id":                 ListJmsPluginsSortById,
	"timeLastSeen":       ListJmsPluginsSortByTimelastseen,
	"timeRegistered":     ListJmsPluginsSortByTimeregistered,
	"hostname":           ListJmsPluginsSortByHostname,
	"agentId":            ListJmsPluginsSortByAgentid,
	"agentType":          ListJmsPluginsSortByAgenttype,
	"lifecycleState":     ListJmsPluginsSortByLifecyclestate,
	"availabilityStatus": ListJmsPluginsSortByAvailabilitystatus,
	"fleetId":            ListJmsPluginsSortByFleetid,
	"compartmentId":      ListJmsPluginsSortByCompartmentid,
	"osFamily":           ListJmsPluginsSortByOsfamily,
	"osArchitecture":     ListJmsPluginsSortByOsarchitecture,
	"osDistribution":     ListJmsPluginsSortByOsdistribution,
	"pluginVersion":      ListJmsPluginsSortByPluginversion,
}

var mappingListJmsPluginsSortByEnumLowerCase = map[string]ListJmsPluginsSortByEnum{
	"id":                 ListJmsPluginsSortById,
	"timelastseen":       ListJmsPluginsSortByTimelastseen,
	"timeregistered":     ListJmsPluginsSortByTimeregistered,
	"hostname":           ListJmsPluginsSortByHostname,
	"agentid":            ListJmsPluginsSortByAgentid,
	"agenttype":          ListJmsPluginsSortByAgenttype,
	"lifecyclestate":     ListJmsPluginsSortByLifecyclestate,
	"availabilitystatus": ListJmsPluginsSortByAvailabilitystatus,
	"fleetid":            ListJmsPluginsSortByFleetid,
	"compartmentid":      ListJmsPluginsSortByCompartmentid,
	"osfamily":           ListJmsPluginsSortByOsfamily,
	"osarchitecture":     ListJmsPluginsSortByOsarchitecture,
	"osdistribution":     ListJmsPluginsSortByOsdistribution,
	"pluginversion":      ListJmsPluginsSortByPluginversion,
}

// GetListJmsPluginsSortByEnumValues Enumerates the set of values for ListJmsPluginsSortByEnum
func GetListJmsPluginsSortByEnumValues() []ListJmsPluginsSortByEnum {
	values := make([]ListJmsPluginsSortByEnum, 0)
	for _, v := range mappingListJmsPluginsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJmsPluginsSortByEnumStringValues Enumerates the set of values in String for ListJmsPluginsSortByEnum
func GetListJmsPluginsSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeLastSeen",
		"timeRegistered",
		"hostname",
		"agentId",
		"agentType",
		"lifecycleState",
		"availabilityStatus",
		"fleetId",
		"compartmentId",
		"osFamily",
		"osArchitecture",
		"osDistribution",
		"pluginVersion",
	}
}

// GetMappingListJmsPluginsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJmsPluginsSortByEnum(val string) (ListJmsPluginsSortByEnum, bool) {
	enum, ok := mappingListJmsPluginsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
