// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPatternsRequest wrapper for the ListPatterns operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListPatterns.go.html to see an example of how to use ListPatternsRequest.
type ListPatternsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListPatternsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Specifies the fields to return in a pattern summary response.
	Fields []ListPatternsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListPatternsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListPatternsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPatternsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPatternsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPatternsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPatternsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPatternsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPatternsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListPatternsLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListPatternsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListPatternsFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListPatternsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPatternsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatternsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPatternsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPatternsResponse wrapper for the ListPatterns operation
type ListPatternsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PatternCollection instances
	PatternCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPatternsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPatternsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPatternsLifecycleStateEnum Enum with underlying type: string
type ListPatternsLifecycleStateEnum string

// Set of constants representing the allowable values for ListPatternsLifecycleStateEnum
const (
	ListPatternsLifecycleStateCreating ListPatternsLifecycleStateEnum = "CREATING"
	ListPatternsLifecycleStateActive   ListPatternsLifecycleStateEnum = "ACTIVE"
	ListPatternsLifecycleStateInactive ListPatternsLifecycleStateEnum = "INACTIVE"
	ListPatternsLifecycleStateUpdating ListPatternsLifecycleStateEnum = "UPDATING"
	ListPatternsLifecycleStateDeleting ListPatternsLifecycleStateEnum = "DELETING"
	ListPatternsLifecycleStateDeleted  ListPatternsLifecycleStateEnum = "DELETED"
	ListPatternsLifecycleStateFailed   ListPatternsLifecycleStateEnum = "FAILED"
	ListPatternsLifecycleStateMoving   ListPatternsLifecycleStateEnum = "MOVING"
)

var mappingListPatternsLifecycleStateEnum = map[string]ListPatternsLifecycleStateEnum{
	"CREATING": ListPatternsLifecycleStateCreating,
	"ACTIVE":   ListPatternsLifecycleStateActive,
	"INACTIVE": ListPatternsLifecycleStateInactive,
	"UPDATING": ListPatternsLifecycleStateUpdating,
	"DELETING": ListPatternsLifecycleStateDeleting,
	"DELETED":  ListPatternsLifecycleStateDeleted,
	"FAILED":   ListPatternsLifecycleStateFailed,
	"MOVING":   ListPatternsLifecycleStateMoving,
}

var mappingListPatternsLifecycleStateEnumLowerCase = map[string]ListPatternsLifecycleStateEnum{
	"creating": ListPatternsLifecycleStateCreating,
	"active":   ListPatternsLifecycleStateActive,
	"inactive": ListPatternsLifecycleStateInactive,
	"updating": ListPatternsLifecycleStateUpdating,
	"deleting": ListPatternsLifecycleStateDeleting,
	"deleted":  ListPatternsLifecycleStateDeleted,
	"failed":   ListPatternsLifecycleStateFailed,
	"moving":   ListPatternsLifecycleStateMoving,
}

// GetListPatternsLifecycleStateEnumValues Enumerates the set of values for ListPatternsLifecycleStateEnum
func GetListPatternsLifecycleStateEnumValues() []ListPatternsLifecycleStateEnum {
	values := make([]ListPatternsLifecycleStateEnum, 0)
	for _, v := range mappingListPatternsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatternsLifecycleStateEnumStringValues Enumerates the set of values in String for ListPatternsLifecycleStateEnum
func GetListPatternsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"MOVING",
	}
}

// GetMappingListPatternsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatternsLifecycleStateEnum(val string) (ListPatternsLifecycleStateEnum, bool) {
	enum, ok := mappingListPatternsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPatternsFieldsEnum Enum with underlying type: string
type ListPatternsFieldsEnum string

// Set of constants representing the allowable values for ListPatternsFieldsEnum
const (
	ListPatternsFieldsKey            ListPatternsFieldsEnum = "key"
	ListPatternsFieldsDisplayname    ListPatternsFieldsEnum = "displayName"
	ListPatternsFieldsDescription    ListPatternsFieldsEnum = "description"
	ListPatternsFieldsCatalogid      ListPatternsFieldsEnum = "catalogId"
	ListPatternsFieldsExpression     ListPatternsFieldsEnum = "expression"
	ListPatternsFieldsLifecyclestate ListPatternsFieldsEnum = "lifecycleState"
	ListPatternsFieldsTimecreated    ListPatternsFieldsEnum = "timeCreated"
)

var mappingListPatternsFieldsEnum = map[string]ListPatternsFieldsEnum{
	"key":            ListPatternsFieldsKey,
	"displayName":    ListPatternsFieldsDisplayname,
	"description":    ListPatternsFieldsDescription,
	"catalogId":      ListPatternsFieldsCatalogid,
	"expression":     ListPatternsFieldsExpression,
	"lifecycleState": ListPatternsFieldsLifecyclestate,
	"timeCreated":    ListPatternsFieldsTimecreated,
}

var mappingListPatternsFieldsEnumLowerCase = map[string]ListPatternsFieldsEnum{
	"key":            ListPatternsFieldsKey,
	"displayname":    ListPatternsFieldsDisplayname,
	"description":    ListPatternsFieldsDescription,
	"catalogid":      ListPatternsFieldsCatalogid,
	"expression":     ListPatternsFieldsExpression,
	"lifecyclestate": ListPatternsFieldsLifecyclestate,
	"timecreated":    ListPatternsFieldsTimecreated,
}

// GetListPatternsFieldsEnumValues Enumerates the set of values for ListPatternsFieldsEnum
func GetListPatternsFieldsEnumValues() []ListPatternsFieldsEnum {
	values := make([]ListPatternsFieldsEnum, 0)
	for _, v := range mappingListPatternsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatternsFieldsEnumStringValues Enumerates the set of values in String for ListPatternsFieldsEnum
func GetListPatternsFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"catalogId",
		"expression",
		"lifecycleState",
		"timeCreated",
	}
}

// GetMappingListPatternsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatternsFieldsEnum(val string) (ListPatternsFieldsEnum, bool) {
	enum, ok := mappingListPatternsFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPatternsSortByEnum Enum with underlying type: string
type ListPatternsSortByEnum string

// Set of constants representing the allowable values for ListPatternsSortByEnum
const (
	ListPatternsSortByTimecreated ListPatternsSortByEnum = "TIMECREATED"
	ListPatternsSortByDisplayname ListPatternsSortByEnum = "DISPLAYNAME"
)

var mappingListPatternsSortByEnum = map[string]ListPatternsSortByEnum{
	"TIMECREATED": ListPatternsSortByTimecreated,
	"DISPLAYNAME": ListPatternsSortByDisplayname,
}

var mappingListPatternsSortByEnumLowerCase = map[string]ListPatternsSortByEnum{
	"timecreated": ListPatternsSortByTimecreated,
	"displayname": ListPatternsSortByDisplayname,
}

// GetListPatternsSortByEnumValues Enumerates the set of values for ListPatternsSortByEnum
func GetListPatternsSortByEnumValues() []ListPatternsSortByEnum {
	values := make([]ListPatternsSortByEnum, 0)
	for _, v := range mappingListPatternsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatternsSortByEnumStringValues Enumerates the set of values in String for ListPatternsSortByEnum
func GetListPatternsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListPatternsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatternsSortByEnum(val string) (ListPatternsSortByEnum, bool) {
	enum, ok := mappingListPatternsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPatternsSortOrderEnum Enum with underlying type: string
type ListPatternsSortOrderEnum string

// Set of constants representing the allowable values for ListPatternsSortOrderEnum
const (
	ListPatternsSortOrderAsc  ListPatternsSortOrderEnum = "ASC"
	ListPatternsSortOrderDesc ListPatternsSortOrderEnum = "DESC"
)

var mappingListPatternsSortOrderEnum = map[string]ListPatternsSortOrderEnum{
	"ASC":  ListPatternsSortOrderAsc,
	"DESC": ListPatternsSortOrderDesc,
}

var mappingListPatternsSortOrderEnumLowerCase = map[string]ListPatternsSortOrderEnum{
	"asc":  ListPatternsSortOrderAsc,
	"desc": ListPatternsSortOrderDesc,
}

// GetListPatternsSortOrderEnumValues Enumerates the set of values for ListPatternsSortOrderEnum
func GetListPatternsSortOrderEnumValues() []ListPatternsSortOrderEnum {
	values := make([]ListPatternsSortOrderEnum, 0)
	for _, v := range mappingListPatternsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatternsSortOrderEnumStringValues Enumerates the set of values in String for ListPatternsSortOrderEnum
func GetListPatternsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPatternsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatternsSortOrderEnum(val string) (ListPatternsSortOrderEnum, bool) {
	enum, ok := mappingListPatternsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
