// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCollectionRulesRequest wrapper for the ListCollectionRules operation
type ListCollectionRulesRequest struct {

	// The Log Analytics namespace used for the request. The namespace can be obtained by running 'oci os ns get'
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The collection rule name used for filtering. Only exact matches are returned.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The text used for collection rule filtering. Only collection rules with either a name
	// or description containing the specified text will be returned.
	DisplayText *string `mandatory:"false" contributesTo:"query" name:"displayText"`

	// The collection rule type used for filtering. Only collection rules with
	// the specified type will be returned.
	Type ListCollectionRulesTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The collection rule context type used for filtering. Only collection rules with
	// the specified context type will be returned.
	ContextType ListCollectionRulesContextTypeEnum `mandatory:"false" contributesTo:"query" name:"contextType" omitEmpty:"true"`

	// The context value used for filtering.
	ContextValue *string `mandatory:"false" contributesTo:"query" name:"contextValue"`

	// The template OCID used for filtering.
	TemplateId *string `mandatory:"false" contributesTo:"query" name:"templateId"`

	// The collection rule lifecycle state used for filtering. Currently supported
	// values are ACTIVE and DELETED.
	LifecycleState ListCollectionRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCollectionRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned templates.
	SortBy ListCollectionRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCollectionRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCollectionRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCollectionRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCollectionRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCollectionRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCollectionRulesTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListCollectionRulesTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCollectionRulesContextTypeEnum(string(request.ContextType)); !ok && request.ContextType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContextType: %s. Supported values are: %s.", request.ContextType, strings.Join(GetListCollectionRulesContextTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCollectionRulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListCollectionRulesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCollectionRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCollectionRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCollectionRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCollectionRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCollectionRulesResponse wrapper for the ListCollectionRules operation
type ListCollectionRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CollectionRuleCollection instances
	CollectionRuleCollection `presentIn:"body"`

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

func (response ListCollectionRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCollectionRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCollectionRulesTypeEnum Enum with underlying type: string
type ListCollectionRulesTypeEnum string

// Set of constants representing the allowable values for ListCollectionRulesTypeEnum
const (
	ListCollectionRulesTypeAgent ListCollectionRulesTypeEnum = "AGENT"
)

var mappingListCollectionRulesTypeEnum = map[string]ListCollectionRulesTypeEnum{
	"AGENT": ListCollectionRulesTypeAgent,
}

var mappingListCollectionRulesTypeEnumLowerCase = map[string]ListCollectionRulesTypeEnum{
	"agent": ListCollectionRulesTypeAgent,
}

// GetListCollectionRulesTypeEnumValues Enumerates the set of values for ListCollectionRulesTypeEnum
func GetListCollectionRulesTypeEnumValues() []ListCollectionRulesTypeEnum {
	values := make([]ListCollectionRulesTypeEnum, 0)
	for _, v := range mappingListCollectionRulesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListCollectionRulesTypeEnumStringValues Enumerates the set of values in String for ListCollectionRulesTypeEnum
func GetListCollectionRulesTypeEnumStringValues() []string {
	return []string{
		"AGENT",
	}
}

// GetMappingListCollectionRulesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCollectionRulesTypeEnum(val string) (ListCollectionRulesTypeEnum, bool) {
	enum, ok := mappingListCollectionRulesTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCollectionRulesContextTypeEnum Enum with underlying type: string
type ListCollectionRulesContextTypeEnum string

// Set of constants representing the allowable values for ListCollectionRulesContextTypeEnum
const (
	ListCollectionRulesContextTypeEntityType ListCollectionRulesContextTypeEnum = "ENTITY_TYPE"
	ListCollectionRulesContextTypeSource     ListCollectionRulesContextTypeEnum = "SOURCE"
)

var mappingListCollectionRulesContextTypeEnum = map[string]ListCollectionRulesContextTypeEnum{
	"ENTITY_TYPE": ListCollectionRulesContextTypeEntityType,
	"SOURCE":      ListCollectionRulesContextTypeSource,
}

var mappingListCollectionRulesContextTypeEnumLowerCase = map[string]ListCollectionRulesContextTypeEnum{
	"entity_type": ListCollectionRulesContextTypeEntityType,
	"source":      ListCollectionRulesContextTypeSource,
}

// GetListCollectionRulesContextTypeEnumValues Enumerates the set of values for ListCollectionRulesContextTypeEnum
func GetListCollectionRulesContextTypeEnumValues() []ListCollectionRulesContextTypeEnum {
	values := make([]ListCollectionRulesContextTypeEnum, 0)
	for _, v := range mappingListCollectionRulesContextTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListCollectionRulesContextTypeEnumStringValues Enumerates the set of values in String for ListCollectionRulesContextTypeEnum
func GetListCollectionRulesContextTypeEnumStringValues() []string {
	return []string{
		"ENTITY_TYPE",
		"SOURCE",
	}
}

// GetMappingListCollectionRulesContextTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCollectionRulesContextTypeEnum(val string) (ListCollectionRulesContextTypeEnum, bool) {
	enum, ok := mappingListCollectionRulesContextTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCollectionRulesLifecycleStateEnum Enum with underlying type: string
type ListCollectionRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListCollectionRulesLifecycleStateEnum
const (
	ListCollectionRulesLifecycleStateActive  ListCollectionRulesLifecycleStateEnum = "ACTIVE"
	ListCollectionRulesLifecycleStateDeleted ListCollectionRulesLifecycleStateEnum = "DELETED"
	ListCollectionRulesLifecycleStateAll     ListCollectionRulesLifecycleStateEnum = "ALL"
)

var mappingListCollectionRulesLifecycleStateEnum = map[string]ListCollectionRulesLifecycleStateEnum{
	"ACTIVE":  ListCollectionRulesLifecycleStateActive,
	"DELETED": ListCollectionRulesLifecycleStateDeleted,
	"ALL":     ListCollectionRulesLifecycleStateAll,
}

var mappingListCollectionRulesLifecycleStateEnumLowerCase = map[string]ListCollectionRulesLifecycleStateEnum{
	"active":  ListCollectionRulesLifecycleStateActive,
	"deleted": ListCollectionRulesLifecycleStateDeleted,
	"all":     ListCollectionRulesLifecycleStateAll,
}

// GetListCollectionRulesLifecycleStateEnumValues Enumerates the set of values for ListCollectionRulesLifecycleStateEnum
func GetListCollectionRulesLifecycleStateEnumValues() []ListCollectionRulesLifecycleStateEnum {
	values := make([]ListCollectionRulesLifecycleStateEnum, 0)
	for _, v := range mappingListCollectionRulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListCollectionRulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListCollectionRulesLifecycleStateEnum
func GetListCollectionRulesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"ALL",
	}
}

// GetMappingListCollectionRulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCollectionRulesLifecycleStateEnum(val string) (ListCollectionRulesLifecycleStateEnum, bool) {
	enum, ok := mappingListCollectionRulesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCollectionRulesSortOrderEnum Enum with underlying type: string
type ListCollectionRulesSortOrderEnum string

// Set of constants representing the allowable values for ListCollectionRulesSortOrderEnum
const (
	ListCollectionRulesSortOrderAsc  ListCollectionRulesSortOrderEnum = "ASC"
	ListCollectionRulesSortOrderDesc ListCollectionRulesSortOrderEnum = "DESC"
)

var mappingListCollectionRulesSortOrderEnum = map[string]ListCollectionRulesSortOrderEnum{
	"ASC":  ListCollectionRulesSortOrderAsc,
	"DESC": ListCollectionRulesSortOrderDesc,
}

var mappingListCollectionRulesSortOrderEnumLowerCase = map[string]ListCollectionRulesSortOrderEnum{
	"asc":  ListCollectionRulesSortOrderAsc,
	"desc": ListCollectionRulesSortOrderDesc,
}

// GetListCollectionRulesSortOrderEnumValues Enumerates the set of values for ListCollectionRulesSortOrderEnum
func GetListCollectionRulesSortOrderEnumValues() []ListCollectionRulesSortOrderEnum {
	values := make([]ListCollectionRulesSortOrderEnum, 0)
	for _, v := range mappingListCollectionRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCollectionRulesSortOrderEnumStringValues Enumerates the set of values in String for ListCollectionRulesSortOrderEnum
func GetListCollectionRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCollectionRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCollectionRulesSortOrderEnum(val string) (ListCollectionRulesSortOrderEnum, bool) {
	enum, ok := mappingListCollectionRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCollectionRulesSortByEnum Enum with underlying type: string
type ListCollectionRulesSortByEnum string

// Set of constants representing the allowable values for ListCollectionRulesSortByEnum
const (
	ListCollectionRulesSortByName        ListCollectionRulesSortByEnum = "name"
	ListCollectionRulesSortByTimecreated ListCollectionRulesSortByEnum = "timeCreated"
	ListCollectionRulesSortByTimeupdated ListCollectionRulesSortByEnum = "timeUpdated"
)

var mappingListCollectionRulesSortByEnum = map[string]ListCollectionRulesSortByEnum{
	"name":        ListCollectionRulesSortByName,
	"timeCreated": ListCollectionRulesSortByTimecreated,
	"timeUpdated": ListCollectionRulesSortByTimeupdated,
}

var mappingListCollectionRulesSortByEnumLowerCase = map[string]ListCollectionRulesSortByEnum{
	"name":        ListCollectionRulesSortByName,
	"timecreated": ListCollectionRulesSortByTimecreated,
	"timeupdated": ListCollectionRulesSortByTimeupdated,
}

// GetListCollectionRulesSortByEnumValues Enumerates the set of values for ListCollectionRulesSortByEnum
func GetListCollectionRulesSortByEnumValues() []ListCollectionRulesSortByEnum {
	values := make([]ListCollectionRulesSortByEnum, 0)
	for _, v := range mappingListCollectionRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCollectionRulesSortByEnumStringValues Enumerates the set of values in String for ListCollectionRulesSortByEnum
func GetListCollectionRulesSortByEnumStringValues() []string {
	return []string{
		"name",
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListCollectionRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCollectionRulesSortByEnum(val string) (ListCollectionRulesSortByEnum, bool) {
	enum, ok := mappingListCollectionRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
