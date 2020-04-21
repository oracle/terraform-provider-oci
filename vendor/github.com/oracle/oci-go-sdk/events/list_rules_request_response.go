// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package events

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListRulesRequest wrapper for the ListRules operation
type ListRulesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this rule belongs.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return. 1 is the minimum, 50 is the maximum.
	// Default: 10
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only rules that match the lifecycle state in this parameter.
	// Example: `Creating`
	LifecycleState RuleLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only rules with descriptions that match the displayName string
	// in this parameter.
	// Example: `"This rule sends a notification upon completion of DbaaS backup."`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Specifies the attribute with which to sort the rules.
	// Default: `timeCreated`
	// * **TIME_CREATED:** Sorts by timeCreated.
	// * **DISPLAY_NAME:** Sorts by displayName.
	// * **ID:** Sorts by id.
	SortBy ListRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order.
	// * **ASC:** Ascending sort order.
	// * **DESC:** Descending sort order.
	SortOrder ListRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRulesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListRulesResponse wrapper for the ListRules operation
type ListRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []RuleSummary instances
	Items []RuleSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of
	// results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRulesSortByEnum Enum with underlying type: string
type ListRulesSortByEnum string

// Set of constants representing the allowable values for ListRulesSortByEnum
const (
	ListRulesSortByTimeCreated ListRulesSortByEnum = "TIME_CREATED"
	ListRulesSortById          ListRulesSortByEnum = "ID"
	ListRulesSortByDisplayName ListRulesSortByEnum = "DISPLAY_NAME"
)

var mappingListRulesSortBy = map[string]ListRulesSortByEnum{
	"TIME_CREATED": ListRulesSortByTimeCreated,
	"ID":           ListRulesSortById,
	"DISPLAY_NAME": ListRulesSortByDisplayName,
}

// GetListRulesSortByEnumValues Enumerates the set of values for ListRulesSortByEnum
func GetListRulesSortByEnumValues() []ListRulesSortByEnum {
	values := make([]ListRulesSortByEnum, 0)
	for _, v := range mappingListRulesSortBy {
		values = append(values, v)
	}
	return values
}

// ListRulesSortOrderEnum Enum with underlying type: string
type ListRulesSortOrderEnum string

// Set of constants representing the allowable values for ListRulesSortOrderEnum
const (
	ListRulesSortOrderAsc  ListRulesSortOrderEnum = "ASC"
	ListRulesSortOrderDesc ListRulesSortOrderEnum = "DESC"
)

var mappingListRulesSortOrder = map[string]ListRulesSortOrderEnum{
	"ASC":  ListRulesSortOrderAsc,
	"DESC": ListRulesSortOrderDesc,
}

// GetListRulesSortOrderEnumValues Enumerates the set of values for ListRulesSortOrderEnum
func GetListRulesSortOrderEnumValues() []ListRulesSortOrderEnum {
	values := make([]ListRulesSortOrderEnum, 0)
	for _, v := range mappingListRulesSortOrder {
		values = append(values, v)
	}
	return values
}
