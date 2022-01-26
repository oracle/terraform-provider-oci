// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListResponderRecipeResponderRulesRequest wrapper for the ListResponderRecipeResponderRules operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResponderRecipeResponderRules.go.html to see an example of how to use ListResponderRecipeResponderRulesRequest.
type ListResponderRecipeResponderRulesRequest struct {

	// OCID of ResponderRecipe
	ResponderRecipeId *string `mandatory:"true" contributesTo:"path" name:"responderRecipeId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListResponderRecipeResponderRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListResponderRecipeResponderRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName is ascending. If no value is specified displayName is default.
	SortBy ListResponderRecipeResponderRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResponderRecipeResponderRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResponderRecipeResponderRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResponderRecipeResponderRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResponderRecipeResponderRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListResponderRecipeResponderRulesResponse wrapper for the ListResponderRecipeResponderRules operation
type ListResponderRecipeResponderRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResponderRecipeResponderRuleCollection instances
	ResponderRecipeResponderRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResponderRecipeResponderRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResponderRecipeResponderRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResponderRecipeResponderRulesLifecycleStateEnum Enum with underlying type: string
type ListResponderRecipeResponderRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListResponderRecipeResponderRulesLifecycleStateEnum
const (
	ListResponderRecipeResponderRulesLifecycleStateCreating ListResponderRecipeResponderRulesLifecycleStateEnum = "CREATING"
	ListResponderRecipeResponderRulesLifecycleStateUpdating ListResponderRecipeResponderRulesLifecycleStateEnum = "UPDATING"
	ListResponderRecipeResponderRulesLifecycleStateActive   ListResponderRecipeResponderRulesLifecycleStateEnum = "ACTIVE"
	ListResponderRecipeResponderRulesLifecycleStateInactive ListResponderRecipeResponderRulesLifecycleStateEnum = "INACTIVE"
	ListResponderRecipeResponderRulesLifecycleStateDeleting ListResponderRecipeResponderRulesLifecycleStateEnum = "DELETING"
	ListResponderRecipeResponderRulesLifecycleStateDeleted  ListResponderRecipeResponderRulesLifecycleStateEnum = "DELETED"
	ListResponderRecipeResponderRulesLifecycleStateFailed   ListResponderRecipeResponderRulesLifecycleStateEnum = "FAILED"
)

var mappingListResponderRecipeResponderRulesLifecycleState = map[string]ListResponderRecipeResponderRulesLifecycleStateEnum{
	"CREATING": ListResponderRecipeResponderRulesLifecycleStateCreating,
	"UPDATING": ListResponderRecipeResponderRulesLifecycleStateUpdating,
	"ACTIVE":   ListResponderRecipeResponderRulesLifecycleStateActive,
	"INACTIVE": ListResponderRecipeResponderRulesLifecycleStateInactive,
	"DELETING": ListResponderRecipeResponderRulesLifecycleStateDeleting,
	"DELETED":  ListResponderRecipeResponderRulesLifecycleStateDeleted,
	"FAILED":   ListResponderRecipeResponderRulesLifecycleStateFailed,
}

// GetListResponderRecipeResponderRulesLifecycleStateEnumValues Enumerates the set of values for ListResponderRecipeResponderRulesLifecycleStateEnum
func GetListResponderRecipeResponderRulesLifecycleStateEnumValues() []ListResponderRecipeResponderRulesLifecycleStateEnum {
	values := make([]ListResponderRecipeResponderRulesLifecycleStateEnum, 0)
	for _, v := range mappingListResponderRecipeResponderRulesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListResponderRecipeResponderRulesSortOrderEnum Enum with underlying type: string
type ListResponderRecipeResponderRulesSortOrderEnum string

// Set of constants representing the allowable values for ListResponderRecipeResponderRulesSortOrderEnum
const (
	ListResponderRecipeResponderRulesSortOrderAsc  ListResponderRecipeResponderRulesSortOrderEnum = "ASC"
	ListResponderRecipeResponderRulesSortOrderDesc ListResponderRecipeResponderRulesSortOrderEnum = "DESC"
)

var mappingListResponderRecipeResponderRulesSortOrder = map[string]ListResponderRecipeResponderRulesSortOrderEnum{
	"ASC":  ListResponderRecipeResponderRulesSortOrderAsc,
	"DESC": ListResponderRecipeResponderRulesSortOrderDesc,
}

// GetListResponderRecipeResponderRulesSortOrderEnumValues Enumerates the set of values for ListResponderRecipeResponderRulesSortOrderEnum
func GetListResponderRecipeResponderRulesSortOrderEnumValues() []ListResponderRecipeResponderRulesSortOrderEnum {
	values := make([]ListResponderRecipeResponderRulesSortOrderEnum, 0)
	for _, v := range mappingListResponderRecipeResponderRulesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListResponderRecipeResponderRulesSortByEnum Enum with underlying type: string
type ListResponderRecipeResponderRulesSortByEnum string

// Set of constants representing the allowable values for ListResponderRecipeResponderRulesSortByEnum
const (
	ListResponderRecipeResponderRulesSortByDisplayname ListResponderRecipeResponderRulesSortByEnum = "displayName"
	ListResponderRecipeResponderRulesSortByRisklevel   ListResponderRecipeResponderRulesSortByEnum = "riskLevel"
)

var mappingListResponderRecipeResponderRulesSortBy = map[string]ListResponderRecipeResponderRulesSortByEnum{
	"displayName": ListResponderRecipeResponderRulesSortByDisplayname,
	"riskLevel":   ListResponderRecipeResponderRulesSortByRisklevel,
}

// GetListResponderRecipeResponderRulesSortByEnumValues Enumerates the set of values for ListResponderRecipeResponderRulesSortByEnum
func GetListResponderRecipeResponderRulesSortByEnumValues() []ListResponderRecipeResponderRulesSortByEnum {
	values := make([]ListResponderRecipeResponderRulesSortByEnum, 0)
	for _, v := range mappingListResponderRecipeResponderRulesSortBy {
		values = append(values, v)
	}
	return values
}
