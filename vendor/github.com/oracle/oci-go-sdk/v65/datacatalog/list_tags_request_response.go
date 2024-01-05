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

// ListTagsRequest wrapper for the ListTags operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListTags.go.html to see an example of how to use ListTagsRequest.
type ListTagsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListTagsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies the fields to return in a term summary response.
	Fields []ListTagsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListTagsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTagsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListTagsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTagsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTagsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTagsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTagsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTagsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTagsLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListTagsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListTagsFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListTagsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTagsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTagsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTagsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTagsResponse wrapper for the ListTags operation
type ListTagsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TermCollection instances
	TermCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTagsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTagsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTagsLifecycleStateEnum Enum with underlying type: string
type ListTagsLifecycleStateEnum string

// Set of constants representing the allowable values for ListTagsLifecycleStateEnum
const (
	ListTagsLifecycleStateCreating ListTagsLifecycleStateEnum = "CREATING"
	ListTagsLifecycleStateActive   ListTagsLifecycleStateEnum = "ACTIVE"
	ListTagsLifecycleStateInactive ListTagsLifecycleStateEnum = "INACTIVE"
	ListTagsLifecycleStateUpdating ListTagsLifecycleStateEnum = "UPDATING"
	ListTagsLifecycleStateDeleting ListTagsLifecycleStateEnum = "DELETING"
	ListTagsLifecycleStateDeleted  ListTagsLifecycleStateEnum = "DELETED"
	ListTagsLifecycleStateFailed   ListTagsLifecycleStateEnum = "FAILED"
	ListTagsLifecycleStateMoving   ListTagsLifecycleStateEnum = "MOVING"
)

var mappingListTagsLifecycleStateEnum = map[string]ListTagsLifecycleStateEnum{
	"CREATING": ListTagsLifecycleStateCreating,
	"ACTIVE":   ListTagsLifecycleStateActive,
	"INACTIVE": ListTagsLifecycleStateInactive,
	"UPDATING": ListTagsLifecycleStateUpdating,
	"DELETING": ListTagsLifecycleStateDeleting,
	"DELETED":  ListTagsLifecycleStateDeleted,
	"FAILED":   ListTagsLifecycleStateFailed,
	"MOVING":   ListTagsLifecycleStateMoving,
}

var mappingListTagsLifecycleStateEnumLowerCase = map[string]ListTagsLifecycleStateEnum{
	"creating": ListTagsLifecycleStateCreating,
	"active":   ListTagsLifecycleStateActive,
	"inactive": ListTagsLifecycleStateInactive,
	"updating": ListTagsLifecycleStateUpdating,
	"deleting": ListTagsLifecycleStateDeleting,
	"deleted":  ListTagsLifecycleStateDeleted,
	"failed":   ListTagsLifecycleStateFailed,
	"moving":   ListTagsLifecycleStateMoving,
}

// GetListTagsLifecycleStateEnumValues Enumerates the set of values for ListTagsLifecycleStateEnum
func GetListTagsLifecycleStateEnumValues() []ListTagsLifecycleStateEnum {
	values := make([]ListTagsLifecycleStateEnum, 0)
	for _, v := range mappingListTagsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTagsLifecycleStateEnumStringValues Enumerates the set of values in String for ListTagsLifecycleStateEnum
func GetListTagsLifecycleStateEnumStringValues() []string {
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

// GetMappingListTagsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTagsLifecycleStateEnum(val string) (ListTagsLifecycleStateEnum, bool) {
	enum, ok := mappingListTagsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTagsFieldsEnum Enum with underlying type: string
type ListTagsFieldsEnum string

// Set of constants representing the allowable values for ListTagsFieldsEnum
const (
	ListTagsFieldsKey                       ListTagsFieldsEnum = "key"
	ListTagsFieldsDisplayname               ListTagsFieldsEnum = "displayName"
	ListTagsFieldsDescription               ListTagsFieldsEnum = "description"
	ListTagsFieldsGlossarykey               ListTagsFieldsEnum = "glossaryKey"
	ListTagsFieldsParenttermkey             ListTagsFieldsEnum = "parentTermKey"
	ListTagsFieldsIsallowedtohavechildterms ListTagsFieldsEnum = "isAllowedToHaveChildTerms"
	ListTagsFieldsPath                      ListTagsFieldsEnum = "path"
	ListTagsFieldsLifecyclestate            ListTagsFieldsEnum = "lifecycleState"
	ListTagsFieldsTimecreated               ListTagsFieldsEnum = "timeCreated"
	ListTagsFieldsWorkflowstatus            ListTagsFieldsEnum = "workflowStatus"
	ListTagsFieldsAssociatedobjectcount     ListTagsFieldsEnum = "associatedObjectCount"
	ListTagsFieldsUri                       ListTagsFieldsEnum = "uri"
)

var mappingListTagsFieldsEnum = map[string]ListTagsFieldsEnum{
	"key":                       ListTagsFieldsKey,
	"displayName":               ListTagsFieldsDisplayname,
	"description":               ListTagsFieldsDescription,
	"glossaryKey":               ListTagsFieldsGlossarykey,
	"parentTermKey":             ListTagsFieldsParenttermkey,
	"isAllowedToHaveChildTerms": ListTagsFieldsIsallowedtohavechildterms,
	"path":                      ListTagsFieldsPath,
	"lifecycleState":            ListTagsFieldsLifecyclestate,
	"timeCreated":               ListTagsFieldsTimecreated,
	"workflowStatus":            ListTagsFieldsWorkflowstatus,
	"associatedObjectCount":     ListTagsFieldsAssociatedobjectcount,
	"uri":                       ListTagsFieldsUri,
}

var mappingListTagsFieldsEnumLowerCase = map[string]ListTagsFieldsEnum{
	"key":                       ListTagsFieldsKey,
	"displayname":               ListTagsFieldsDisplayname,
	"description":               ListTagsFieldsDescription,
	"glossarykey":               ListTagsFieldsGlossarykey,
	"parenttermkey":             ListTagsFieldsParenttermkey,
	"isallowedtohavechildterms": ListTagsFieldsIsallowedtohavechildterms,
	"path":                      ListTagsFieldsPath,
	"lifecyclestate":            ListTagsFieldsLifecyclestate,
	"timecreated":               ListTagsFieldsTimecreated,
	"workflowstatus":            ListTagsFieldsWorkflowstatus,
	"associatedobjectcount":     ListTagsFieldsAssociatedobjectcount,
	"uri":                       ListTagsFieldsUri,
}

// GetListTagsFieldsEnumValues Enumerates the set of values for ListTagsFieldsEnum
func GetListTagsFieldsEnumValues() []ListTagsFieldsEnum {
	values := make([]ListTagsFieldsEnum, 0)
	for _, v := range mappingListTagsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListTagsFieldsEnumStringValues Enumerates the set of values in String for ListTagsFieldsEnum
func GetListTagsFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"glossaryKey",
		"parentTermKey",
		"isAllowedToHaveChildTerms",
		"path",
		"lifecycleState",
		"timeCreated",
		"workflowStatus",
		"associatedObjectCount",
		"uri",
	}
}

// GetMappingListTagsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTagsFieldsEnum(val string) (ListTagsFieldsEnum, bool) {
	enum, ok := mappingListTagsFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTagsSortByEnum Enum with underlying type: string
type ListTagsSortByEnum string

// Set of constants representing the allowable values for ListTagsSortByEnum
const (
	ListTagsSortByTimecreated ListTagsSortByEnum = "TIMECREATED"
	ListTagsSortByDisplayname ListTagsSortByEnum = "DISPLAYNAME"
)

var mappingListTagsSortByEnum = map[string]ListTagsSortByEnum{
	"TIMECREATED": ListTagsSortByTimecreated,
	"DISPLAYNAME": ListTagsSortByDisplayname,
}

var mappingListTagsSortByEnumLowerCase = map[string]ListTagsSortByEnum{
	"timecreated": ListTagsSortByTimecreated,
	"displayname": ListTagsSortByDisplayname,
}

// GetListTagsSortByEnumValues Enumerates the set of values for ListTagsSortByEnum
func GetListTagsSortByEnumValues() []ListTagsSortByEnum {
	values := make([]ListTagsSortByEnum, 0)
	for _, v := range mappingListTagsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTagsSortByEnumStringValues Enumerates the set of values in String for ListTagsSortByEnum
func GetListTagsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListTagsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTagsSortByEnum(val string) (ListTagsSortByEnum, bool) {
	enum, ok := mappingListTagsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTagsSortOrderEnum Enum with underlying type: string
type ListTagsSortOrderEnum string

// Set of constants representing the allowable values for ListTagsSortOrderEnum
const (
	ListTagsSortOrderAsc  ListTagsSortOrderEnum = "ASC"
	ListTagsSortOrderDesc ListTagsSortOrderEnum = "DESC"
)

var mappingListTagsSortOrderEnum = map[string]ListTagsSortOrderEnum{
	"ASC":  ListTagsSortOrderAsc,
	"DESC": ListTagsSortOrderDesc,
}

var mappingListTagsSortOrderEnumLowerCase = map[string]ListTagsSortOrderEnum{
	"asc":  ListTagsSortOrderAsc,
	"desc": ListTagsSortOrderDesc,
}

// GetListTagsSortOrderEnumValues Enumerates the set of values for ListTagsSortOrderEnum
func GetListTagsSortOrderEnumValues() []ListTagsSortOrderEnum {
	values := make([]ListTagsSortOrderEnum, 0)
	for _, v := range mappingListTagsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTagsSortOrderEnumStringValues Enumerates the set of values in String for ListTagsSortOrderEnum
func GetListTagsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTagsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTagsSortOrderEnum(val string) (ListTagsSortOrderEnum, bool) {
	enum, ok := mappingListTagsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
