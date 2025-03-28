// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package events

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRulesRequest wrapper for the ListRules operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/events/ListRules.go.html to see an example of how to use ListRulesRequest.
type ListRulesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this rule belongs.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return. 1 is the minimum, 50 is the maximum.
	// Default: 10
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
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
func (request ListRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRuleLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRuleLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRulesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRulesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRulesResponse wrapper for the ListRules operation
type ListRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []RuleSummary instances
	Items []RuleSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of
	// results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
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

var mappingListRulesSortByEnum = map[string]ListRulesSortByEnum{
	"TIME_CREATED": ListRulesSortByTimeCreated,
	"ID":           ListRulesSortById,
	"DISPLAY_NAME": ListRulesSortByDisplayName,
}

var mappingListRulesSortByEnumLowerCase = map[string]ListRulesSortByEnum{
	"time_created": ListRulesSortByTimeCreated,
	"id":           ListRulesSortById,
	"display_name": ListRulesSortByDisplayName,
}

// GetListRulesSortByEnumValues Enumerates the set of values for ListRulesSortByEnum
func GetListRulesSortByEnumValues() []ListRulesSortByEnum {
	values := make([]ListRulesSortByEnum, 0)
	for _, v := range mappingListRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRulesSortByEnumStringValues Enumerates the set of values in String for ListRulesSortByEnum
func GetListRulesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"ID",
		"DISPLAY_NAME",
	}
}

// GetMappingListRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRulesSortByEnum(val string) (ListRulesSortByEnum, bool) {
	enum, ok := mappingListRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRulesSortOrderEnum Enum with underlying type: string
type ListRulesSortOrderEnum string

// Set of constants representing the allowable values for ListRulesSortOrderEnum
const (
	ListRulesSortOrderAsc  ListRulesSortOrderEnum = "ASC"
	ListRulesSortOrderDesc ListRulesSortOrderEnum = "DESC"
)

var mappingListRulesSortOrderEnum = map[string]ListRulesSortOrderEnum{
	"ASC":  ListRulesSortOrderAsc,
	"DESC": ListRulesSortOrderDesc,
}

var mappingListRulesSortOrderEnumLowerCase = map[string]ListRulesSortOrderEnum{
	"asc":  ListRulesSortOrderAsc,
	"desc": ListRulesSortOrderDesc,
}

// GetListRulesSortOrderEnumValues Enumerates the set of values for ListRulesSortOrderEnum
func GetListRulesSortOrderEnumValues() []ListRulesSortOrderEnum {
	values := make([]ListRulesSortOrderEnum, 0)
	for _, v := range mappingListRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRulesSortOrderEnumStringValues Enumerates the set of values in String for ListRulesSortOrderEnum
func GetListRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRulesSortOrderEnum(val string) (ListRulesSortOrderEnum, bool) {
	enum, ok := mappingListRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
