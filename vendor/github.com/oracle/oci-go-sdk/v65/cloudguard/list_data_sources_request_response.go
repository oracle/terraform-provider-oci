// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDataSourcesRequest wrapper for the ListDataSources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDataSources.go.html to see an example of how to use ListDataSourcesRequest.
type ListDataSourcesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources their feedProvider matches the given DataSourceFeedProvider.
	DataSourceFeedProvider ListDataSourcesDataSourceFeedProviderEnum `mandatory:"false" contributesTo:"query" name:"dataSourceFeedProvider" omitEmpty:"true"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListDataSourcesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources their query type matches the given LoggingQueryType.
	LoggingQueryType ListDataSourcesLoggingQueryTypeEnum `mandatory:"false" contributesTo:"query" name:"loggingQueryType" omitEmpty:"true"`

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
	AccessLevel ListDataSourcesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDataSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataSourcesDataSourceFeedProviderEnum(string(request.DataSourceFeedProvider)); !ok && request.DataSourceFeedProvider != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataSourceFeedProvider: %s. Supported values are: %s.", request.DataSourceFeedProvider, strings.Join(GetListDataSourcesDataSourceFeedProviderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSourcesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDataSourcesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSourcesLoggingQueryTypeEnum(string(request.LoggingQueryType)); !ok && request.LoggingQueryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LoggingQueryType: %s. Supported values are: %s.", request.LoggingQueryType, strings.Join(GetListDataSourcesLoggingQueryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSourcesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListDataSourcesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataSourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataSourcesResponse wrapper for the ListDataSources operation
type ListDataSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataSourceCollection instances
	DataSourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataSourcesDataSourceFeedProviderEnum Enum with underlying type: string
type ListDataSourcesDataSourceFeedProviderEnum string

// Set of constants representing the allowable values for ListDataSourcesDataSourceFeedProviderEnum
const (
	ListDataSourcesDataSourceFeedProviderLoggingquery ListDataSourcesDataSourceFeedProviderEnum = "LOGGINGQUERY"
)

var mappingListDataSourcesDataSourceFeedProviderEnum = map[string]ListDataSourcesDataSourceFeedProviderEnum{
	"LOGGINGQUERY": ListDataSourcesDataSourceFeedProviderLoggingquery,
}

var mappingListDataSourcesDataSourceFeedProviderEnumLowerCase = map[string]ListDataSourcesDataSourceFeedProviderEnum{
	"loggingquery": ListDataSourcesDataSourceFeedProviderLoggingquery,
}

// GetListDataSourcesDataSourceFeedProviderEnumValues Enumerates the set of values for ListDataSourcesDataSourceFeedProviderEnum
func GetListDataSourcesDataSourceFeedProviderEnumValues() []ListDataSourcesDataSourceFeedProviderEnum {
	values := make([]ListDataSourcesDataSourceFeedProviderEnum, 0)
	for _, v := range mappingListDataSourcesDataSourceFeedProviderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSourcesDataSourceFeedProviderEnumStringValues Enumerates the set of values in String for ListDataSourcesDataSourceFeedProviderEnum
func GetListDataSourcesDataSourceFeedProviderEnumStringValues() []string {
	return []string{
		"LOGGINGQUERY",
	}
}

// GetMappingListDataSourcesDataSourceFeedProviderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSourcesDataSourceFeedProviderEnum(val string) (ListDataSourcesDataSourceFeedProviderEnum, bool) {
	enum, ok := mappingListDataSourcesDataSourceFeedProviderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSourcesLifecycleStateEnum Enum with underlying type: string
type ListDataSourcesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDataSourcesLifecycleStateEnum
const (
	ListDataSourcesLifecycleStateCreating ListDataSourcesLifecycleStateEnum = "CREATING"
	ListDataSourcesLifecycleStateUpdating ListDataSourcesLifecycleStateEnum = "UPDATING"
	ListDataSourcesLifecycleStateActive   ListDataSourcesLifecycleStateEnum = "ACTIVE"
	ListDataSourcesLifecycleStateInactive ListDataSourcesLifecycleStateEnum = "INACTIVE"
	ListDataSourcesLifecycleStateDeleting ListDataSourcesLifecycleStateEnum = "DELETING"
	ListDataSourcesLifecycleStateDeleted  ListDataSourcesLifecycleStateEnum = "DELETED"
	ListDataSourcesLifecycleStateFailed   ListDataSourcesLifecycleStateEnum = "FAILED"
)

var mappingListDataSourcesLifecycleStateEnum = map[string]ListDataSourcesLifecycleStateEnum{
	"CREATING": ListDataSourcesLifecycleStateCreating,
	"UPDATING": ListDataSourcesLifecycleStateUpdating,
	"ACTIVE":   ListDataSourcesLifecycleStateActive,
	"INACTIVE": ListDataSourcesLifecycleStateInactive,
	"DELETING": ListDataSourcesLifecycleStateDeleting,
	"DELETED":  ListDataSourcesLifecycleStateDeleted,
	"FAILED":   ListDataSourcesLifecycleStateFailed,
}

var mappingListDataSourcesLifecycleStateEnumLowerCase = map[string]ListDataSourcesLifecycleStateEnum{
	"creating": ListDataSourcesLifecycleStateCreating,
	"updating": ListDataSourcesLifecycleStateUpdating,
	"active":   ListDataSourcesLifecycleStateActive,
	"inactive": ListDataSourcesLifecycleStateInactive,
	"deleting": ListDataSourcesLifecycleStateDeleting,
	"deleted":  ListDataSourcesLifecycleStateDeleted,
	"failed":   ListDataSourcesLifecycleStateFailed,
}

// GetListDataSourcesLifecycleStateEnumValues Enumerates the set of values for ListDataSourcesLifecycleStateEnum
func GetListDataSourcesLifecycleStateEnumValues() []ListDataSourcesLifecycleStateEnum {
	values := make([]ListDataSourcesLifecycleStateEnum, 0)
	for _, v := range mappingListDataSourcesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSourcesLifecycleStateEnumStringValues Enumerates the set of values in String for ListDataSourcesLifecycleStateEnum
func GetListDataSourcesLifecycleStateEnumStringValues() []string {
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

// GetMappingListDataSourcesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSourcesLifecycleStateEnum(val string) (ListDataSourcesLifecycleStateEnum, bool) {
	enum, ok := mappingListDataSourcesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSourcesLoggingQueryTypeEnum Enum with underlying type: string
type ListDataSourcesLoggingQueryTypeEnum string

// Set of constants representing the allowable values for ListDataSourcesLoggingQueryTypeEnum
const (
	ListDataSourcesLoggingQueryTypeInsight ListDataSourcesLoggingQueryTypeEnum = "INSIGHT"
)

var mappingListDataSourcesLoggingQueryTypeEnum = map[string]ListDataSourcesLoggingQueryTypeEnum{
	"INSIGHT": ListDataSourcesLoggingQueryTypeInsight,
}

var mappingListDataSourcesLoggingQueryTypeEnumLowerCase = map[string]ListDataSourcesLoggingQueryTypeEnum{
	"insight": ListDataSourcesLoggingQueryTypeInsight,
}

// GetListDataSourcesLoggingQueryTypeEnumValues Enumerates the set of values for ListDataSourcesLoggingQueryTypeEnum
func GetListDataSourcesLoggingQueryTypeEnumValues() []ListDataSourcesLoggingQueryTypeEnum {
	values := make([]ListDataSourcesLoggingQueryTypeEnum, 0)
	for _, v := range mappingListDataSourcesLoggingQueryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSourcesLoggingQueryTypeEnumStringValues Enumerates the set of values in String for ListDataSourcesLoggingQueryTypeEnum
func GetListDataSourcesLoggingQueryTypeEnumStringValues() []string {
	return []string{
		"INSIGHT",
	}
}

// GetMappingListDataSourcesLoggingQueryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSourcesLoggingQueryTypeEnum(val string) (ListDataSourcesLoggingQueryTypeEnum, bool) {
	enum, ok := mappingListDataSourcesLoggingQueryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSourcesAccessLevelEnum Enum with underlying type: string
type ListDataSourcesAccessLevelEnum string

// Set of constants representing the allowable values for ListDataSourcesAccessLevelEnum
const (
	ListDataSourcesAccessLevelRestricted ListDataSourcesAccessLevelEnum = "RESTRICTED"
	ListDataSourcesAccessLevelAccessible ListDataSourcesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListDataSourcesAccessLevelEnum = map[string]ListDataSourcesAccessLevelEnum{
	"RESTRICTED": ListDataSourcesAccessLevelRestricted,
	"ACCESSIBLE": ListDataSourcesAccessLevelAccessible,
}

var mappingListDataSourcesAccessLevelEnumLowerCase = map[string]ListDataSourcesAccessLevelEnum{
	"restricted": ListDataSourcesAccessLevelRestricted,
	"accessible": ListDataSourcesAccessLevelAccessible,
}

// GetListDataSourcesAccessLevelEnumValues Enumerates the set of values for ListDataSourcesAccessLevelEnum
func GetListDataSourcesAccessLevelEnumValues() []ListDataSourcesAccessLevelEnum {
	values := make([]ListDataSourcesAccessLevelEnum, 0)
	for _, v := range mappingListDataSourcesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSourcesAccessLevelEnumStringValues Enumerates the set of values in String for ListDataSourcesAccessLevelEnum
func GetListDataSourcesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListDataSourcesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSourcesAccessLevelEnum(val string) (ListDataSourcesAccessLevelEnum, bool) {
	enum, ok := mappingListDataSourcesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSourcesSortOrderEnum Enum with underlying type: string
type ListDataSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListDataSourcesSortOrderEnum
const (
	ListDataSourcesSortOrderAsc  ListDataSourcesSortOrderEnum = "ASC"
	ListDataSourcesSortOrderDesc ListDataSourcesSortOrderEnum = "DESC"
)

var mappingListDataSourcesSortOrderEnum = map[string]ListDataSourcesSortOrderEnum{
	"ASC":  ListDataSourcesSortOrderAsc,
	"DESC": ListDataSourcesSortOrderDesc,
}

var mappingListDataSourcesSortOrderEnumLowerCase = map[string]ListDataSourcesSortOrderEnum{
	"asc":  ListDataSourcesSortOrderAsc,
	"desc": ListDataSourcesSortOrderDesc,
}

// GetListDataSourcesSortOrderEnumValues Enumerates the set of values for ListDataSourcesSortOrderEnum
func GetListDataSourcesSortOrderEnumValues() []ListDataSourcesSortOrderEnum {
	values := make([]ListDataSourcesSortOrderEnum, 0)
	for _, v := range mappingListDataSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListDataSourcesSortOrderEnum
func GetListDataSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSourcesSortOrderEnum(val string) (ListDataSourcesSortOrderEnum, bool) {
	enum, ok := mappingListDataSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSourcesSortByEnum Enum with underlying type: string
type ListDataSourcesSortByEnum string

// Set of constants representing the allowable values for ListDataSourcesSortByEnum
const (
	ListDataSourcesSortByTimecreated ListDataSourcesSortByEnum = "timeCreated"
	ListDataSourcesSortByDisplayname ListDataSourcesSortByEnum = "displayName"
)

var mappingListDataSourcesSortByEnum = map[string]ListDataSourcesSortByEnum{
	"timeCreated": ListDataSourcesSortByTimecreated,
	"displayName": ListDataSourcesSortByDisplayname,
}

var mappingListDataSourcesSortByEnumLowerCase = map[string]ListDataSourcesSortByEnum{
	"timecreated": ListDataSourcesSortByTimecreated,
	"displayname": ListDataSourcesSortByDisplayname,
}

// GetListDataSourcesSortByEnumValues Enumerates the set of values for ListDataSourcesSortByEnum
func GetListDataSourcesSortByEnumValues() []ListDataSourcesSortByEnum {
	values := make([]ListDataSourcesSortByEnum, 0)
	for _, v := range mappingListDataSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSourcesSortByEnumStringValues Enumerates the set of values in String for ListDataSourcesSortByEnum
func GetListDataSourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDataSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSourcesSortByEnum(val string) (ListDataSourcesSortByEnum, bool) {
	enum, ok := mappingListDataSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
