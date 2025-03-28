// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLogAnalyticsEntitiesRequest wrapper for the ListLogAnalyticsEntities operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsEntities.go.html to see an example of how to use ListLogAnalyticsEntitiesRequest.
type ListLogAnalyticsEntitiesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only log analytics entities whose name matches the entire name given. The match
	// is case-insensitive.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only log analytics entities whose name contains the name given. The match
	// is case-insensitive.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// A filter to return only log analytics entities whose entityTypeName matches the entire log analytics entity type name of
	// one of the entityTypeNames given in the list. The match is case-insensitive.
	EntityTypeName []string `contributesTo:"query" name:"entityTypeName" collectionFormat:"multi"`

	// A filter to return only log analytics entities whose cloudResourceId matches the cloudResourceId given.
	CloudResourceId *string `mandatory:"false" contributesTo:"query" name:"cloudResourceId"`

	// A filter to return only those log analytics entities with the specified lifecycle state. The state
	// value is case-insensitive.
	LifecycleState ListLogAnalyticsEntitiesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only log analytics entities whose lifecycleDetails contains the specified string.
	LifecycleDetailsContains *string `mandatory:"false" contributesTo:"query" name:"lifecycleDetailsContains"`

	// A filter to return only those log analytics entities whose managementAgentId is null or is not null.
	IsManagementAgentIdNull ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum `mandatory:"false" contributesTo:"query" name:"isManagementAgentIdNull" omitEmpty:"true"`

	// A filter to return only log analytics entities whose hostname matches the entire hostname given.
	Hostname *string `mandatory:"false" contributesTo:"query" name:"hostname"`

	// A filter to return only log analytics entities whose hostname contains the substring given.
	// The match is case-insensitive.
	HostnameContains *string `mandatory:"false" contributesTo:"query" name:"hostnameContains"`

	// A filter to return only log analytics entities whose sourceId matches the sourceId given.
	SourceId *string `mandatory:"false" contributesTo:"query" name:"sourceId"`

	// A filter to return only those log analytics entities with the specified auto-creation source.
	CreationSourceType []CreationSourceTypeEnum `contributesTo:"query" name:"creationSourceType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only log analytics entities whose auto-creation source details contains the specified string.
	CreationSourceDetails *string `mandatory:"false" contributesTo:"query" name:"creationSourceDetails"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLogAnalyticsEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort entities by. Only one sort order may be provided. Default order for timeCreated and timeUpdated
	// is descending. Default order for entity name is ascending. If no value is specified timeCreated is default.
	SortBy ListLogAnalyticsEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only log analytics entities whose metadata name, value and type matches the specified string.
	// Each item in the array has the format "{name}:{value}:{type}".  All inputs are case-insensitive.
	MetadataEquals []string `contributesTo:"query" name:"metadataEquals" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogAnalyticsEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogAnalyticsEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogAnalyticsEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogAnalyticsEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLogAnalyticsEntitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLogAnalyticsEntitiesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListLogAnalyticsEntitiesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogAnalyticsEntitiesIsManagementAgentIdNullEnum(string(request.IsManagementAgentIdNull)); !ok && request.IsManagementAgentIdNull != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsManagementAgentIdNull: %s. Supported values are: %s.", request.IsManagementAgentIdNull, strings.Join(GetListLogAnalyticsEntitiesIsManagementAgentIdNullEnumStringValues(), ",")))
	}
	for _, val := range request.CreationSourceType {
		if _, ok := GetMappingCreationSourceTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreationSourceType: %s. Supported values are: %s.", val, strings.Join(GetCreationSourceTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListLogAnalyticsEntitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLogAnalyticsEntitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogAnalyticsEntitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLogAnalyticsEntitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLogAnalyticsEntitiesResponse wrapper for the ListLogAnalyticsEntities operation
type ListLogAnalyticsEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsEntityCollection instances
	LogAnalyticsEntityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLogAnalyticsEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogAnalyticsEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogAnalyticsEntitiesLifecycleStateEnum Enum with underlying type: string
type ListLogAnalyticsEntitiesLifecycleStateEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntitiesLifecycleStateEnum
const (
	ListLogAnalyticsEntitiesLifecycleStateActive  ListLogAnalyticsEntitiesLifecycleStateEnum = "ACTIVE"
	ListLogAnalyticsEntitiesLifecycleStateDeleted ListLogAnalyticsEntitiesLifecycleStateEnum = "DELETED"
)

var mappingListLogAnalyticsEntitiesLifecycleStateEnum = map[string]ListLogAnalyticsEntitiesLifecycleStateEnum{
	"ACTIVE":  ListLogAnalyticsEntitiesLifecycleStateActive,
	"DELETED": ListLogAnalyticsEntitiesLifecycleStateDeleted,
}

var mappingListLogAnalyticsEntitiesLifecycleStateEnumLowerCase = map[string]ListLogAnalyticsEntitiesLifecycleStateEnum{
	"active":  ListLogAnalyticsEntitiesLifecycleStateActive,
	"deleted": ListLogAnalyticsEntitiesLifecycleStateDeleted,
}

// GetListLogAnalyticsEntitiesLifecycleStateEnumValues Enumerates the set of values for ListLogAnalyticsEntitiesLifecycleStateEnum
func GetListLogAnalyticsEntitiesLifecycleStateEnumValues() []ListLogAnalyticsEntitiesLifecycleStateEnum {
	values := make([]ListLogAnalyticsEntitiesLifecycleStateEnum, 0)
	for _, v := range mappingListLogAnalyticsEntitiesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogAnalyticsEntitiesLifecycleStateEnumStringValues Enumerates the set of values in String for ListLogAnalyticsEntitiesLifecycleStateEnum
func GetListLogAnalyticsEntitiesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListLogAnalyticsEntitiesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogAnalyticsEntitiesLifecycleStateEnum(val string) (ListLogAnalyticsEntitiesLifecycleStateEnum, bool) {
	enum, ok := mappingListLogAnalyticsEntitiesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum Enum with underlying type: string
type ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum
const (
	ListLogAnalyticsEntitiesIsManagementAgentIdNullTrue  ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum = "true"
	ListLogAnalyticsEntitiesIsManagementAgentIdNullFalse ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum = "false"
)

var mappingListLogAnalyticsEntitiesIsManagementAgentIdNullEnum = map[string]ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum{
	"true":  ListLogAnalyticsEntitiesIsManagementAgentIdNullTrue,
	"false": ListLogAnalyticsEntitiesIsManagementAgentIdNullFalse,
}

var mappingListLogAnalyticsEntitiesIsManagementAgentIdNullEnumLowerCase = map[string]ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum{
	"true":  ListLogAnalyticsEntitiesIsManagementAgentIdNullTrue,
	"false": ListLogAnalyticsEntitiesIsManagementAgentIdNullFalse,
}

// GetListLogAnalyticsEntitiesIsManagementAgentIdNullEnumValues Enumerates the set of values for ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum
func GetListLogAnalyticsEntitiesIsManagementAgentIdNullEnumValues() []ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum {
	values := make([]ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum, 0)
	for _, v := range mappingListLogAnalyticsEntitiesIsManagementAgentIdNullEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogAnalyticsEntitiesIsManagementAgentIdNullEnumStringValues Enumerates the set of values in String for ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum
func GetListLogAnalyticsEntitiesIsManagementAgentIdNullEnumStringValues() []string {
	return []string{
		"true",
		"false",
	}
}

// GetMappingListLogAnalyticsEntitiesIsManagementAgentIdNullEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogAnalyticsEntitiesIsManagementAgentIdNullEnum(val string) (ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum, bool) {
	enum, ok := mappingListLogAnalyticsEntitiesIsManagementAgentIdNullEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogAnalyticsEntitiesSortOrderEnum Enum with underlying type: string
type ListLogAnalyticsEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntitiesSortOrderEnum
const (
	ListLogAnalyticsEntitiesSortOrderAsc  ListLogAnalyticsEntitiesSortOrderEnum = "ASC"
	ListLogAnalyticsEntitiesSortOrderDesc ListLogAnalyticsEntitiesSortOrderEnum = "DESC"
)

var mappingListLogAnalyticsEntitiesSortOrderEnum = map[string]ListLogAnalyticsEntitiesSortOrderEnum{
	"ASC":  ListLogAnalyticsEntitiesSortOrderAsc,
	"DESC": ListLogAnalyticsEntitiesSortOrderDesc,
}

var mappingListLogAnalyticsEntitiesSortOrderEnumLowerCase = map[string]ListLogAnalyticsEntitiesSortOrderEnum{
	"asc":  ListLogAnalyticsEntitiesSortOrderAsc,
	"desc": ListLogAnalyticsEntitiesSortOrderDesc,
}

// GetListLogAnalyticsEntitiesSortOrderEnumValues Enumerates the set of values for ListLogAnalyticsEntitiesSortOrderEnum
func GetListLogAnalyticsEntitiesSortOrderEnumValues() []ListLogAnalyticsEntitiesSortOrderEnum {
	values := make([]ListLogAnalyticsEntitiesSortOrderEnum, 0)
	for _, v := range mappingListLogAnalyticsEntitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogAnalyticsEntitiesSortOrderEnumStringValues Enumerates the set of values in String for ListLogAnalyticsEntitiesSortOrderEnum
func GetListLogAnalyticsEntitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLogAnalyticsEntitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogAnalyticsEntitiesSortOrderEnum(val string) (ListLogAnalyticsEntitiesSortOrderEnum, bool) {
	enum, ok := mappingListLogAnalyticsEntitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogAnalyticsEntitiesSortByEnum Enum with underlying type: string
type ListLogAnalyticsEntitiesSortByEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntitiesSortByEnum
const (
	ListLogAnalyticsEntitiesSortByTimecreated ListLogAnalyticsEntitiesSortByEnum = "timeCreated"
	ListLogAnalyticsEntitiesSortByTimeupdated ListLogAnalyticsEntitiesSortByEnum = "timeUpdated"
	ListLogAnalyticsEntitiesSortByName        ListLogAnalyticsEntitiesSortByEnum = "name"
)

var mappingListLogAnalyticsEntitiesSortByEnum = map[string]ListLogAnalyticsEntitiesSortByEnum{
	"timeCreated": ListLogAnalyticsEntitiesSortByTimecreated,
	"timeUpdated": ListLogAnalyticsEntitiesSortByTimeupdated,
	"name":        ListLogAnalyticsEntitiesSortByName,
}

var mappingListLogAnalyticsEntitiesSortByEnumLowerCase = map[string]ListLogAnalyticsEntitiesSortByEnum{
	"timecreated": ListLogAnalyticsEntitiesSortByTimecreated,
	"timeupdated": ListLogAnalyticsEntitiesSortByTimeupdated,
	"name":        ListLogAnalyticsEntitiesSortByName,
}

// GetListLogAnalyticsEntitiesSortByEnumValues Enumerates the set of values for ListLogAnalyticsEntitiesSortByEnum
func GetListLogAnalyticsEntitiesSortByEnumValues() []ListLogAnalyticsEntitiesSortByEnum {
	values := make([]ListLogAnalyticsEntitiesSortByEnum, 0)
	for _, v := range mappingListLogAnalyticsEntitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogAnalyticsEntitiesSortByEnumStringValues Enumerates the set of values in String for ListLogAnalyticsEntitiesSortByEnum
func GetListLogAnalyticsEntitiesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"name",
	}
}

// GetMappingListLogAnalyticsEntitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogAnalyticsEntitiesSortByEnum(val string) (ListLogAnalyticsEntitiesSortByEnum, bool) {
	enum, ok := mappingListLogAnalyticsEntitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
