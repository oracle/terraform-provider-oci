// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDataMaskRulesRequest wrapper for the ListDataMaskRules operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDataMaskRules.go.html to see an example of how to use ListDataMaskRulesRequest.
type ListDataMaskRulesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListDataMaskRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel ListDataMaskRulesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataMaskRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDataMaskRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The status of the dataMaskRule.
	DataMaskRuleStatus ListDataMaskRulesDataMaskRuleStatusEnum `mandatory:"false" contributesTo:"query" name:"dataMaskRuleStatus" omitEmpty:"true"`

	// OCID of target
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// OCID of iamGroup
	IamGroupId *string `mandatory:"false" contributesTo:"query" name:"iamGroupId"`

	// Type of target
	TargetType *string `mandatory:"false" contributesTo:"query" name:"targetType"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataMaskRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataMaskRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataMaskRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataMaskRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDataMaskRulesResponse wrapper for the ListDataMaskRules operation
type ListDataMaskRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataMaskRuleCollection instances
	DataMaskRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataMaskRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataMaskRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataMaskRulesLifecycleStateEnum Enum with underlying type: string
type ListDataMaskRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDataMaskRulesLifecycleStateEnum
const (
	ListDataMaskRulesLifecycleStateCreating ListDataMaskRulesLifecycleStateEnum = "CREATING"
	ListDataMaskRulesLifecycleStateUpdating ListDataMaskRulesLifecycleStateEnum = "UPDATING"
	ListDataMaskRulesLifecycleStateActive   ListDataMaskRulesLifecycleStateEnum = "ACTIVE"
	ListDataMaskRulesLifecycleStateInactive ListDataMaskRulesLifecycleStateEnum = "INACTIVE"
	ListDataMaskRulesLifecycleStateDeleting ListDataMaskRulesLifecycleStateEnum = "DELETING"
	ListDataMaskRulesLifecycleStateDeleted  ListDataMaskRulesLifecycleStateEnum = "DELETED"
	ListDataMaskRulesLifecycleStateFailed   ListDataMaskRulesLifecycleStateEnum = "FAILED"
)

var mappingListDataMaskRulesLifecycleState = map[string]ListDataMaskRulesLifecycleStateEnum{
	"CREATING": ListDataMaskRulesLifecycleStateCreating,
	"UPDATING": ListDataMaskRulesLifecycleStateUpdating,
	"ACTIVE":   ListDataMaskRulesLifecycleStateActive,
	"INACTIVE": ListDataMaskRulesLifecycleStateInactive,
	"DELETING": ListDataMaskRulesLifecycleStateDeleting,
	"DELETED":  ListDataMaskRulesLifecycleStateDeleted,
	"FAILED":   ListDataMaskRulesLifecycleStateFailed,
}

// GetListDataMaskRulesLifecycleStateEnumValues Enumerates the set of values for ListDataMaskRulesLifecycleStateEnum
func GetListDataMaskRulesLifecycleStateEnumValues() []ListDataMaskRulesLifecycleStateEnum {
	values := make([]ListDataMaskRulesLifecycleStateEnum, 0)
	for _, v := range mappingListDataMaskRulesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListDataMaskRulesAccessLevelEnum Enum with underlying type: string
type ListDataMaskRulesAccessLevelEnum string

// Set of constants representing the allowable values for ListDataMaskRulesAccessLevelEnum
const (
	ListDataMaskRulesAccessLevelRestricted ListDataMaskRulesAccessLevelEnum = "RESTRICTED"
	ListDataMaskRulesAccessLevelAccessible ListDataMaskRulesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListDataMaskRulesAccessLevel = map[string]ListDataMaskRulesAccessLevelEnum{
	"RESTRICTED": ListDataMaskRulesAccessLevelRestricted,
	"ACCESSIBLE": ListDataMaskRulesAccessLevelAccessible,
}

// GetListDataMaskRulesAccessLevelEnumValues Enumerates the set of values for ListDataMaskRulesAccessLevelEnum
func GetListDataMaskRulesAccessLevelEnumValues() []ListDataMaskRulesAccessLevelEnum {
	values := make([]ListDataMaskRulesAccessLevelEnum, 0)
	for _, v := range mappingListDataMaskRulesAccessLevel {
		values = append(values, v)
	}
	return values
}

// ListDataMaskRulesSortOrderEnum Enum with underlying type: string
type ListDataMaskRulesSortOrderEnum string

// Set of constants representing the allowable values for ListDataMaskRulesSortOrderEnum
const (
	ListDataMaskRulesSortOrderAsc  ListDataMaskRulesSortOrderEnum = "ASC"
	ListDataMaskRulesSortOrderDesc ListDataMaskRulesSortOrderEnum = "DESC"
)

var mappingListDataMaskRulesSortOrder = map[string]ListDataMaskRulesSortOrderEnum{
	"ASC":  ListDataMaskRulesSortOrderAsc,
	"DESC": ListDataMaskRulesSortOrderDesc,
}

// GetListDataMaskRulesSortOrderEnumValues Enumerates the set of values for ListDataMaskRulesSortOrderEnum
func GetListDataMaskRulesSortOrderEnumValues() []ListDataMaskRulesSortOrderEnum {
	values := make([]ListDataMaskRulesSortOrderEnum, 0)
	for _, v := range mappingListDataMaskRulesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDataMaskRulesSortByEnum Enum with underlying type: string
type ListDataMaskRulesSortByEnum string

// Set of constants representing the allowable values for ListDataMaskRulesSortByEnum
const (
	ListDataMaskRulesSortByTimecreated ListDataMaskRulesSortByEnum = "timeCreated"
	ListDataMaskRulesSortByDisplayname ListDataMaskRulesSortByEnum = "displayName"
)

var mappingListDataMaskRulesSortBy = map[string]ListDataMaskRulesSortByEnum{
	"timeCreated": ListDataMaskRulesSortByTimecreated,
	"displayName": ListDataMaskRulesSortByDisplayname,
}

// GetListDataMaskRulesSortByEnumValues Enumerates the set of values for ListDataMaskRulesSortByEnum
func GetListDataMaskRulesSortByEnumValues() []ListDataMaskRulesSortByEnum {
	values := make([]ListDataMaskRulesSortByEnum, 0)
	for _, v := range mappingListDataMaskRulesSortBy {
		values = append(values, v)
	}
	return values
}

// ListDataMaskRulesDataMaskRuleStatusEnum Enum with underlying type: string
type ListDataMaskRulesDataMaskRuleStatusEnum string

// Set of constants representing the allowable values for ListDataMaskRulesDataMaskRuleStatusEnum
const (
	ListDataMaskRulesDataMaskRuleStatusEnabled  ListDataMaskRulesDataMaskRuleStatusEnum = "ENABLED"
	ListDataMaskRulesDataMaskRuleStatusDisabled ListDataMaskRulesDataMaskRuleStatusEnum = "DISABLED"
)

var mappingListDataMaskRulesDataMaskRuleStatus = map[string]ListDataMaskRulesDataMaskRuleStatusEnum{
	"ENABLED":  ListDataMaskRulesDataMaskRuleStatusEnabled,
	"DISABLED": ListDataMaskRulesDataMaskRuleStatusDisabled,
}

// GetListDataMaskRulesDataMaskRuleStatusEnumValues Enumerates the set of values for ListDataMaskRulesDataMaskRuleStatusEnum
func GetListDataMaskRulesDataMaskRuleStatusEnumValues() []ListDataMaskRulesDataMaskRuleStatusEnum {
	values := make([]ListDataMaskRulesDataMaskRuleStatusEnum, 0)
	for _, v := range mappingListDataMaskRulesDataMaskRuleStatus {
		values = append(values, v)
	}
	return values
}
