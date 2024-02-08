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

// ListRulesRequest wrapper for the ListRules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListRules.go.html to see an example of how to use ListRulesRequest.
type ListRulesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return rules whose displayName matches in whole or in part the
	// specified value. The match is case-insensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The rule kind used for filtering. Only rules of the specified
	// kind will be returned.
	Kind ListRulesKindEnum `mandatory:"false" contributesTo:"query" name:"kind" omitEmpty:"true"`

	// The rule lifecycle state used for filtering. Currently supported
	// values are ACTIVE and DELETED.
	LifecycleState ListRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
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
	if _, ok := GetMappingListRulesKindEnum(string(request.Kind)); !ok && request.Kind != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Kind: %s. Supported values are: %s.", request.Kind, strings.Join(GetListRulesKindEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListRulesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRulesSortByEnumStringValues(), ",")))
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

	// A list of RuleSummaryCollection instances
	RuleSummaryCollection `presentIn:"body"`

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

func (response ListRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRulesKindEnum Enum with underlying type: string
type ListRulesKindEnum string

// Set of constants representing the allowable values for ListRulesKindEnum
const (
	ListRulesKindIngestTime  ListRulesKindEnum = "INGEST_TIME"
	ListRulesKindSavedSearch ListRulesKindEnum = "SAVED_SEARCH"
	ListRulesKindAll         ListRulesKindEnum = "ALL"
)

var mappingListRulesKindEnum = map[string]ListRulesKindEnum{
	"INGEST_TIME":  ListRulesKindIngestTime,
	"SAVED_SEARCH": ListRulesKindSavedSearch,
	"ALL":          ListRulesKindAll,
}

var mappingListRulesKindEnumLowerCase = map[string]ListRulesKindEnum{
	"ingest_time":  ListRulesKindIngestTime,
	"saved_search": ListRulesKindSavedSearch,
	"all":          ListRulesKindAll,
}

// GetListRulesKindEnumValues Enumerates the set of values for ListRulesKindEnum
func GetListRulesKindEnumValues() []ListRulesKindEnum {
	values := make([]ListRulesKindEnum, 0)
	for _, v := range mappingListRulesKindEnum {
		values = append(values, v)
	}
	return values
}

// GetListRulesKindEnumStringValues Enumerates the set of values in String for ListRulesKindEnum
func GetListRulesKindEnumStringValues() []string {
	return []string{
		"INGEST_TIME",
		"SAVED_SEARCH",
		"ALL",
	}
}

// GetMappingListRulesKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRulesKindEnum(val string) (ListRulesKindEnum, bool) {
	enum, ok := mappingListRulesKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRulesLifecycleStateEnum Enum with underlying type: string
type ListRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListRulesLifecycleStateEnum
const (
	ListRulesLifecycleStateActive  ListRulesLifecycleStateEnum = "ACTIVE"
	ListRulesLifecycleStateDeleted ListRulesLifecycleStateEnum = "DELETED"
)

var mappingListRulesLifecycleStateEnum = map[string]ListRulesLifecycleStateEnum{
	"ACTIVE":  ListRulesLifecycleStateActive,
	"DELETED": ListRulesLifecycleStateDeleted,
}

var mappingListRulesLifecycleStateEnumLowerCase = map[string]ListRulesLifecycleStateEnum{
	"active":  ListRulesLifecycleStateActive,
	"deleted": ListRulesLifecycleStateDeleted,
}

// GetListRulesLifecycleStateEnumValues Enumerates the set of values for ListRulesLifecycleStateEnum
func GetListRulesLifecycleStateEnumValues() []ListRulesLifecycleStateEnum {
	values := make([]ListRulesLifecycleStateEnum, 0)
	for _, v := range mappingListRulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListRulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListRulesLifecycleStateEnum
func GetListRulesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListRulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRulesLifecycleStateEnum(val string) (ListRulesLifecycleStateEnum, bool) {
	enum, ok := mappingListRulesLifecycleStateEnumLowerCase[strings.ToLower(val)]
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

// ListRulesSortByEnum Enum with underlying type: string
type ListRulesSortByEnum string

// Set of constants representing the allowable values for ListRulesSortByEnum
const (
	ListRulesSortByTimecreated ListRulesSortByEnum = "timeCreated"
	ListRulesSortByTimeupdated ListRulesSortByEnum = "timeUpdated"
	ListRulesSortByDisplayname ListRulesSortByEnum = "displayName"
)

var mappingListRulesSortByEnum = map[string]ListRulesSortByEnum{
	"timeCreated": ListRulesSortByTimecreated,
	"timeUpdated": ListRulesSortByTimeupdated,
	"displayName": ListRulesSortByDisplayname,
}

var mappingListRulesSortByEnumLowerCase = map[string]ListRulesSortByEnum{
	"timecreated": ListRulesSortByTimecreated,
	"timeupdated": ListRulesSortByTimeupdated,
	"displayname": ListRulesSortByDisplayname,
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
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRulesSortByEnum(val string) (ListRulesSortByEnum, bool) {
	enum, ok := mappingListRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
