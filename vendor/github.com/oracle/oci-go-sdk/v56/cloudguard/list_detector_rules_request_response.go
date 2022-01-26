// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDetectorRulesRequest wrapper for the ListDetectorRules operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDetectorRules.go.html to see an example of how to use ListDetectorRulesRequest.
type ListDetectorRulesRequest struct {

	// The Name of Detector.
	DetectorId *string `mandatory:"true" contributesTo:"path" name:"detectorId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListDetectorRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDetectorRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDetectorRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDetectorRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDetectorRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDetectorRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDetectorRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDetectorRulesResponse wrapper for the ListDetectorRules operation
type ListDetectorRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DetectorRuleCollection instances
	DetectorRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDetectorRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDetectorRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDetectorRulesLifecycleStateEnum Enum with underlying type: string
type ListDetectorRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDetectorRulesLifecycleStateEnum
const (
	ListDetectorRulesLifecycleStateCreating ListDetectorRulesLifecycleStateEnum = "CREATING"
	ListDetectorRulesLifecycleStateUpdating ListDetectorRulesLifecycleStateEnum = "UPDATING"
	ListDetectorRulesLifecycleStateActive   ListDetectorRulesLifecycleStateEnum = "ACTIVE"
	ListDetectorRulesLifecycleStateInactive ListDetectorRulesLifecycleStateEnum = "INACTIVE"
	ListDetectorRulesLifecycleStateDeleting ListDetectorRulesLifecycleStateEnum = "DELETING"
	ListDetectorRulesLifecycleStateDeleted  ListDetectorRulesLifecycleStateEnum = "DELETED"
	ListDetectorRulesLifecycleStateFailed   ListDetectorRulesLifecycleStateEnum = "FAILED"
)

var mappingListDetectorRulesLifecycleState = map[string]ListDetectorRulesLifecycleStateEnum{
	"CREATING": ListDetectorRulesLifecycleStateCreating,
	"UPDATING": ListDetectorRulesLifecycleStateUpdating,
	"ACTIVE":   ListDetectorRulesLifecycleStateActive,
	"INACTIVE": ListDetectorRulesLifecycleStateInactive,
	"DELETING": ListDetectorRulesLifecycleStateDeleting,
	"DELETED":  ListDetectorRulesLifecycleStateDeleted,
	"FAILED":   ListDetectorRulesLifecycleStateFailed,
}

// GetListDetectorRulesLifecycleStateEnumValues Enumerates the set of values for ListDetectorRulesLifecycleStateEnum
func GetListDetectorRulesLifecycleStateEnumValues() []ListDetectorRulesLifecycleStateEnum {
	values := make([]ListDetectorRulesLifecycleStateEnum, 0)
	for _, v := range mappingListDetectorRulesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListDetectorRulesSortOrderEnum Enum with underlying type: string
type ListDetectorRulesSortOrderEnum string

// Set of constants representing the allowable values for ListDetectorRulesSortOrderEnum
const (
	ListDetectorRulesSortOrderAsc  ListDetectorRulesSortOrderEnum = "ASC"
	ListDetectorRulesSortOrderDesc ListDetectorRulesSortOrderEnum = "DESC"
)

var mappingListDetectorRulesSortOrder = map[string]ListDetectorRulesSortOrderEnum{
	"ASC":  ListDetectorRulesSortOrderAsc,
	"DESC": ListDetectorRulesSortOrderDesc,
}

// GetListDetectorRulesSortOrderEnumValues Enumerates the set of values for ListDetectorRulesSortOrderEnum
func GetListDetectorRulesSortOrderEnumValues() []ListDetectorRulesSortOrderEnum {
	values := make([]ListDetectorRulesSortOrderEnum, 0)
	for _, v := range mappingListDetectorRulesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDetectorRulesSortByEnum Enum with underlying type: string
type ListDetectorRulesSortByEnum string

// Set of constants representing the allowable values for ListDetectorRulesSortByEnum
const (
	ListDetectorRulesSortByTimecreated ListDetectorRulesSortByEnum = "timeCreated"
	ListDetectorRulesSortByDisplayname ListDetectorRulesSortByEnum = "displayName"
)

var mappingListDetectorRulesSortBy = map[string]ListDetectorRulesSortByEnum{
	"timeCreated": ListDetectorRulesSortByTimecreated,
	"displayName": ListDetectorRulesSortByDisplayname,
}

// GetListDetectorRulesSortByEnumValues Enumerates the set of values for ListDetectorRulesSortByEnum
func GetListDetectorRulesSortByEnumValues() []ListDetectorRulesSortByEnum {
	values := make([]ListDetectorRulesSortByEnum, 0)
	for _, v := range mappingListDetectorRulesSortBy {
		values = append(values, v)
	}
	return values
}
