// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagementAgentPluginsRequest wrapper for the ListManagementAgentPlugins operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementagent/ListManagementAgentPlugins.go.html to see an example of how to use ListManagementAgentPluginsRequest.
type ListManagementAgentPluginsRequest struct {

	// The OCID of the compartment to which a request will be scoped.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter to return only Management Agent Plugins having the particular display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagementAgentPluginsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Default order for displayName is ascending. If no value is specified displayName is default.
	SortBy ListManagementAgentPluginsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Filter to return only Management Agents in the particular lifecycle state.
	LifecycleState ListManagementAgentPluginsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter to return only results having the particular platform type.
	PlatformType []PlatformTypesEnum `contributesTo:"query" name:"platformType" omitEmpty:"true" collectionFormat:"multi"`

	// The ManagementAgentID of the agent from which the Management Agents to be filtered.
	AgentId *string `mandatory:"false" contributesTo:"query" name:"agentId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementAgentPluginsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementAgentPluginsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagementAgentPluginsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementAgentPluginsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagementAgentPluginsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagementAgentPluginsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagementAgentPluginsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAgentPluginsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagementAgentPluginsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAgentPluginsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListManagementAgentPluginsLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.PlatformType {
		if _, ok := GetMappingPlatformTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", val, strings.Join(GetPlatformTypesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagementAgentPluginsResponse wrapper for the ListManagementAgentPlugins operation
type ListManagementAgentPluginsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ManagementAgentPluginSummary instances
	Items []ManagementAgentPluginSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagementAgentPluginsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementAgentPluginsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementAgentPluginsSortOrderEnum Enum with underlying type: string
type ListManagementAgentPluginsSortOrderEnum string

// Set of constants representing the allowable values for ListManagementAgentPluginsSortOrderEnum
const (
	ListManagementAgentPluginsSortOrderAsc  ListManagementAgentPluginsSortOrderEnum = "ASC"
	ListManagementAgentPluginsSortOrderDesc ListManagementAgentPluginsSortOrderEnum = "DESC"
)

var mappingListManagementAgentPluginsSortOrderEnum = map[string]ListManagementAgentPluginsSortOrderEnum{
	"ASC":  ListManagementAgentPluginsSortOrderAsc,
	"DESC": ListManagementAgentPluginsSortOrderDesc,
}

var mappingListManagementAgentPluginsSortOrderEnumLowerCase = map[string]ListManagementAgentPluginsSortOrderEnum{
	"asc":  ListManagementAgentPluginsSortOrderAsc,
	"desc": ListManagementAgentPluginsSortOrderDesc,
}

// GetListManagementAgentPluginsSortOrderEnumValues Enumerates the set of values for ListManagementAgentPluginsSortOrderEnum
func GetListManagementAgentPluginsSortOrderEnumValues() []ListManagementAgentPluginsSortOrderEnum {
	values := make([]ListManagementAgentPluginsSortOrderEnum, 0)
	for _, v := range mappingListManagementAgentPluginsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentPluginsSortOrderEnumStringValues Enumerates the set of values in String for ListManagementAgentPluginsSortOrderEnum
func GetListManagementAgentPluginsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagementAgentPluginsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentPluginsSortOrderEnum(val string) (ListManagementAgentPluginsSortOrderEnum, bool) {
	enum, ok := mappingListManagementAgentPluginsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAgentPluginsSortByEnum Enum with underlying type: string
type ListManagementAgentPluginsSortByEnum string

// Set of constants representing the allowable values for ListManagementAgentPluginsSortByEnum
const (
	ListManagementAgentPluginsSortByDisplayname ListManagementAgentPluginsSortByEnum = "displayName"
)

var mappingListManagementAgentPluginsSortByEnum = map[string]ListManagementAgentPluginsSortByEnum{
	"displayName": ListManagementAgentPluginsSortByDisplayname,
}

var mappingListManagementAgentPluginsSortByEnumLowerCase = map[string]ListManagementAgentPluginsSortByEnum{
	"displayname": ListManagementAgentPluginsSortByDisplayname,
}

// GetListManagementAgentPluginsSortByEnumValues Enumerates the set of values for ListManagementAgentPluginsSortByEnum
func GetListManagementAgentPluginsSortByEnumValues() []ListManagementAgentPluginsSortByEnum {
	values := make([]ListManagementAgentPluginsSortByEnum, 0)
	for _, v := range mappingListManagementAgentPluginsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentPluginsSortByEnumStringValues Enumerates the set of values in String for ListManagementAgentPluginsSortByEnum
func GetListManagementAgentPluginsSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListManagementAgentPluginsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentPluginsSortByEnum(val string) (ListManagementAgentPluginsSortByEnum, bool) {
	enum, ok := mappingListManagementAgentPluginsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAgentPluginsLifecycleStateEnum Enum with underlying type: string
type ListManagementAgentPluginsLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagementAgentPluginsLifecycleStateEnum
const (
	ListManagementAgentPluginsLifecycleStateCreating   ListManagementAgentPluginsLifecycleStateEnum = "CREATING"
	ListManagementAgentPluginsLifecycleStateUpdating   ListManagementAgentPluginsLifecycleStateEnum = "UPDATING"
	ListManagementAgentPluginsLifecycleStateActive     ListManagementAgentPluginsLifecycleStateEnum = "ACTIVE"
	ListManagementAgentPluginsLifecycleStateInactive   ListManagementAgentPluginsLifecycleStateEnum = "INACTIVE"
	ListManagementAgentPluginsLifecycleStateTerminated ListManagementAgentPluginsLifecycleStateEnum = "TERMINATED"
	ListManagementAgentPluginsLifecycleStateDeleting   ListManagementAgentPluginsLifecycleStateEnum = "DELETING"
	ListManagementAgentPluginsLifecycleStateDeleted    ListManagementAgentPluginsLifecycleStateEnum = "DELETED"
	ListManagementAgentPluginsLifecycleStateFailed     ListManagementAgentPluginsLifecycleStateEnum = "FAILED"
)

var mappingListManagementAgentPluginsLifecycleStateEnum = map[string]ListManagementAgentPluginsLifecycleStateEnum{
	"CREATING":   ListManagementAgentPluginsLifecycleStateCreating,
	"UPDATING":   ListManagementAgentPluginsLifecycleStateUpdating,
	"ACTIVE":     ListManagementAgentPluginsLifecycleStateActive,
	"INACTIVE":   ListManagementAgentPluginsLifecycleStateInactive,
	"TERMINATED": ListManagementAgentPluginsLifecycleStateTerminated,
	"DELETING":   ListManagementAgentPluginsLifecycleStateDeleting,
	"DELETED":    ListManagementAgentPluginsLifecycleStateDeleted,
	"FAILED":     ListManagementAgentPluginsLifecycleStateFailed,
}

var mappingListManagementAgentPluginsLifecycleStateEnumLowerCase = map[string]ListManagementAgentPluginsLifecycleStateEnum{
	"creating":   ListManagementAgentPluginsLifecycleStateCreating,
	"updating":   ListManagementAgentPluginsLifecycleStateUpdating,
	"active":     ListManagementAgentPluginsLifecycleStateActive,
	"inactive":   ListManagementAgentPluginsLifecycleStateInactive,
	"terminated": ListManagementAgentPluginsLifecycleStateTerminated,
	"deleting":   ListManagementAgentPluginsLifecycleStateDeleting,
	"deleted":    ListManagementAgentPluginsLifecycleStateDeleted,
	"failed":     ListManagementAgentPluginsLifecycleStateFailed,
}

// GetListManagementAgentPluginsLifecycleStateEnumValues Enumerates the set of values for ListManagementAgentPluginsLifecycleStateEnum
func GetListManagementAgentPluginsLifecycleStateEnumValues() []ListManagementAgentPluginsLifecycleStateEnum {
	values := make([]ListManagementAgentPluginsLifecycleStateEnum, 0)
	for _, v := range mappingListManagementAgentPluginsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentPluginsLifecycleStateEnumStringValues Enumerates the set of values in String for ListManagementAgentPluginsLifecycleStateEnum
func GetListManagementAgentPluginsLifecycleStateEnumStringValues() []string {
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

// GetMappingListManagementAgentPluginsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentPluginsLifecycleStateEnum(val string) (ListManagementAgentPluginsLifecycleStateEnum, bool) {
	enum, ok := mappingListManagementAgentPluginsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
