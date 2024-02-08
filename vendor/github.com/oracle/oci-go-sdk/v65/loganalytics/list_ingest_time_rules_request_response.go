// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListIngestTimeRulesRequest wrapper for the ListIngestTimeRules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListIngestTimeRules.go.html to see an example of how to use ListIngestTimeRulesRequest.
type ListIngestTimeRulesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return rules whose displayName matches in whole or in part the
	// specified value. The match is case-insensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The rule lifecycle state used for filtering. Currently supported
	// values are ACTIVE and DELETED.
	LifecycleState ListIngestTimeRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The ingest time rule condition kind used for filtering. Only rules with conditions
	// of the specified kind will be returned.
	ConditionKind ListIngestTimeRulesConditionKindEnum `mandatory:"false" contributesTo:"query" name:"conditionKind" omitEmpty:"true"`

	// The field name used for filtering. Only rules using the
	// specified field name will be returned.
	FieldName *string `mandatory:"false" contributesTo:"query" name:"fieldName"`

	// The field value used for filtering. Only rules using the
	// specified field value will be returned.
	FieldValue *string `mandatory:"false" contributesTo:"query" name:"fieldValue"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListIngestTimeRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListIngestTimeRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIngestTimeRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIngestTimeRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIngestTimeRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIngestTimeRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIngestTimeRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListIngestTimeRulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListIngestTimeRulesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIngestTimeRulesConditionKindEnum(string(request.ConditionKind)); !ok && request.ConditionKind != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionKind: %s. Supported values are: %s.", request.ConditionKind, strings.Join(GetListIngestTimeRulesConditionKindEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIngestTimeRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIngestTimeRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIngestTimeRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIngestTimeRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIngestTimeRulesResponse wrapper for the ListIngestTimeRules operation
type ListIngestTimeRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of IngestTimeRuleSummaryCollection instances
	IngestTimeRuleSummaryCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListIngestTimeRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIngestTimeRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIngestTimeRulesLifecycleStateEnum Enum with underlying type: string
type ListIngestTimeRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListIngestTimeRulesLifecycleStateEnum
const (
	ListIngestTimeRulesLifecycleStateActive  ListIngestTimeRulesLifecycleStateEnum = "ACTIVE"
	ListIngestTimeRulesLifecycleStateDeleted ListIngestTimeRulesLifecycleStateEnum = "DELETED"
)

var mappingListIngestTimeRulesLifecycleStateEnum = map[string]ListIngestTimeRulesLifecycleStateEnum{
	"ACTIVE":  ListIngestTimeRulesLifecycleStateActive,
	"DELETED": ListIngestTimeRulesLifecycleStateDeleted,
}

var mappingListIngestTimeRulesLifecycleStateEnumLowerCase = map[string]ListIngestTimeRulesLifecycleStateEnum{
	"active":  ListIngestTimeRulesLifecycleStateActive,
	"deleted": ListIngestTimeRulesLifecycleStateDeleted,
}

// GetListIngestTimeRulesLifecycleStateEnumValues Enumerates the set of values for ListIngestTimeRulesLifecycleStateEnum
func GetListIngestTimeRulesLifecycleStateEnumValues() []ListIngestTimeRulesLifecycleStateEnum {
	values := make([]ListIngestTimeRulesLifecycleStateEnum, 0)
	for _, v := range mappingListIngestTimeRulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListIngestTimeRulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListIngestTimeRulesLifecycleStateEnum
func GetListIngestTimeRulesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListIngestTimeRulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIngestTimeRulesLifecycleStateEnum(val string) (ListIngestTimeRulesLifecycleStateEnum, bool) {
	enum, ok := mappingListIngestTimeRulesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIngestTimeRulesConditionKindEnum Enum with underlying type: string
type ListIngestTimeRulesConditionKindEnum string

// Set of constants representing the allowable values for ListIngestTimeRulesConditionKindEnum
const (
	ListIngestTimeRulesConditionKindField ListIngestTimeRulesConditionKindEnum = "FIELD"
)

var mappingListIngestTimeRulesConditionKindEnum = map[string]ListIngestTimeRulesConditionKindEnum{
	"FIELD": ListIngestTimeRulesConditionKindField,
}

var mappingListIngestTimeRulesConditionKindEnumLowerCase = map[string]ListIngestTimeRulesConditionKindEnum{
	"field": ListIngestTimeRulesConditionKindField,
}

// GetListIngestTimeRulesConditionKindEnumValues Enumerates the set of values for ListIngestTimeRulesConditionKindEnum
func GetListIngestTimeRulesConditionKindEnumValues() []ListIngestTimeRulesConditionKindEnum {
	values := make([]ListIngestTimeRulesConditionKindEnum, 0)
	for _, v := range mappingListIngestTimeRulesConditionKindEnum {
		values = append(values, v)
	}
	return values
}

// GetListIngestTimeRulesConditionKindEnumStringValues Enumerates the set of values in String for ListIngestTimeRulesConditionKindEnum
func GetListIngestTimeRulesConditionKindEnumStringValues() []string {
	return []string{
		"FIELD",
	}
}

// GetMappingListIngestTimeRulesConditionKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIngestTimeRulesConditionKindEnum(val string) (ListIngestTimeRulesConditionKindEnum, bool) {
	enum, ok := mappingListIngestTimeRulesConditionKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIngestTimeRulesSortOrderEnum Enum with underlying type: string
type ListIngestTimeRulesSortOrderEnum string

// Set of constants representing the allowable values for ListIngestTimeRulesSortOrderEnum
const (
	ListIngestTimeRulesSortOrderAsc  ListIngestTimeRulesSortOrderEnum = "ASC"
	ListIngestTimeRulesSortOrderDesc ListIngestTimeRulesSortOrderEnum = "DESC"
)

var mappingListIngestTimeRulesSortOrderEnum = map[string]ListIngestTimeRulesSortOrderEnum{
	"ASC":  ListIngestTimeRulesSortOrderAsc,
	"DESC": ListIngestTimeRulesSortOrderDesc,
}

var mappingListIngestTimeRulesSortOrderEnumLowerCase = map[string]ListIngestTimeRulesSortOrderEnum{
	"asc":  ListIngestTimeRulesSortOrderAsc,
	"desc": ListIngestTimeRulesSortOrderDesc,
}

// GetListIngestTimeRulesSortOrderEnumValues Enumerates the set of values for ListIngestTimeRulesSortOrderEnum
func GetListIngestTimeRulesSortOrderEnumValues() []ListIngestTimeRulesSortOrderEnum {
	values := make([]ListIngestTimeRulesSortOrderEnum, 0)
	for _, v := range mappingListIngestTimeRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIngestTimeRulesSortOrderEnumStringValues Enumerates the set of values in String for ListIngestTimeRulesSortOrderEnum
func GetListIngestTimeRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIngestTimeRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIngestTimeRulesSortOrderEnum(val string) (ListIngestTimeRulesSortOrderEnum, bool) {
	enum, ok := mappingListIngestTimeRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIngestTimeRulesSortByEnum Enum with underlying type: string
type ListIngestTimeRulesSortByEnum string

// Set of constants representing the allowable values for ListIngestTimeRulesSortByEnum
const (
	ListIngestTimeRulesSortByTimecreated ListIngestTimeRulesSortByEnum = "timeCreated"
	ListIngestTimeRulesSortByTimeupdated ListIngestTimeRulesSortByEnum = "timeUpdated"
	ListIngestTimeRulesSortByDisplayname ListIngestTimeRulesSortByEnum = "displayName"
)

var mappingListIngestTimeRulesSortByEnum = map[string]ListIngestTimeRulesSortByEnum{
	"timeCreated": ListIngestTimeRulesSortByTimecreated,
	"timeUpdated": ListIngestTimeRulesSortByTimeupdated,
	"displayName": ListIngestTimeRulesSortByDisplayname,
}

var mappingListIngestTimeRulesSortByEnumLowerCase = map[string]ListIngestTimeRulesSortByEnum{
	"timecreated": ListIngestTimeRulesSortByTimecreated,
	"timeupdated": ListIngestTimeRulesSortByTimeupdated,
	"displayname": ListIngestTimeRulesSortByDisplayname,
}

// GetListIngestTimeRulesSortByEnumValues Enumerates the set of values for ListIngestTimeRulesSortByEnum
func GetListIngestTimeRulesSortByEnumValues() []ListIngestTimeRulesSortByEnum {
	values := make([]ListIngestTimeRulesSortByEnum, 0)
	for _, v := range mappingListIngestTimeRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIngestTimeRulesSortByEnumStringValues Enumerates the set of values in String for ListIngestTimeRulesSortByEnum
func GetListIngestTimeRulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListIngestTimeRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIngestTimeRulesSortByEnum(val string) (ListIngestTimeRulesSortByEnum, bool) {
	enum, ok := mappingListIngestTimeRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
