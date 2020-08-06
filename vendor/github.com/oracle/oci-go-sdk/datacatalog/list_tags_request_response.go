// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListTagsRequest wrapper for the ListTags operation
type ListTagsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu".
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
func (request ListTagsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTagsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
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

var mappingListTagsLifecycleState = map[string]ListTagsLifecycleStateEnum{
	"CREATING": ListTagsLifecycleStateCreating,
	"ACTIVE":   ListTagsLifecycleStateActive,
	"INACTIVE": ListTagsLifecycleStateInactive,
	"UPDATING": ListTagsLifecycleStateUpdating,
	"DELETING": ListTagsLifecycleStateDeleting,
	"DELETED":  ListTagsLifecycleStateDeleted,
	"FAILED":   ListTagsLifecycleStateFailed,
	"MOVING":   ListTagsLifecycleStateMoving,
}

// GetListTagsLifecycleStateEnumValues Enumerates the set of values for ListTagsLifecycleStateEnum
func GetListTagsLifecycleStateEnumValues() []ListTagsLifecycleStateEnum {
	values := make([]ListTagsLifecycleStateEnum, 0)
	for _, v := range mappingListTagsLifecycleState {
		values = append(values, v)
	}
	return values
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

var mappingListTagsFields = map[string]ListTagsFieldsEnum{
	"key":                       ListTagsFieldsKey,
	"displayName":               ListTagsFieldsDisplayname,
	"description":               ListTagsFieldsDescription,
	"glossaryKey":               ListTagsFieldsGlossarykey,
	"parentTermKey":             ListTagsFieldsParenttermkey,
	"isAllowedToHaveChildTerms": ListTagsFieldsIsallowedtohavechildterms,
	"path":                  ListTagsFieldsPath,
	"lifecycleState":        ListTagsFieldsLifecyclestate,
	"timeCreated":           ListTagsFieldsTimecreated,
	"workflowStatus":        ListTagsFieldsWorkflowstatus,
	"associatedObjectCount": ListTagsFieldsAssociatedobjectcount,
	"uri": ListTagsFieldsUri,
}

// GetListTagsFieldsEnumValues Enumerates the set of values for ListTagsFieldsEnum
func GetListTagsFieldsEnumValues() []ListTagsFieldsEnum {
	values := make([]ListTagsFieldsEnum, 0)
	for _, v := range mappingListTagsFields {
		values = append(values, v)
	}
	return values
}

// ListTagsSortByEnum Enum with underlying type: string
type ListTagsSortByEnum string

// Set of constants representing the allowable values for ListTagsSortByEnum
const (
	ListTagsSortByTimecreated ListTagsSortByEnum = "TIMECREATED"
	ListTagsSortByDisplayname ListTagsSortByEnum = "DISPLAYNAME"
)

var mappingListTagsSortBy = map[string]ListTagsSortByEnum{
	"TIMECREATED": ListTagsSortByTimecreated,
	"DISPLAYNAME": ListTagsSortByDisplayname,
}

// GetListTagsSortByEnumValues Enumerates the set of values for ListTagsSortByEnum
func GetListTagsSortByEnumValues() []ListTagsSortByEnum {
	values := make([]ListTagsSortByEnum, 0)
	for _, v := range mappingListTagsSortBy {
		values = append(values, v)
	}
	return values
}

// ListTagsSortOrderEnum Enum with underlying type: string
type ListTagsSortOrderEnum string

// Set of constants representing the allowable values for ListTagsSortOrderEnum
const (
	ListTagsSortOrderAsc  ListTagsSortOrderEnum = "ASC"
	ListTagsSortOrderDesc ListTagsSortOrderEnum = "DESC"
)

var mappingListTagsSortOrder = map[string]ListTagsSortOrderEnum{
	"ASC":  ListTagsSortOrderAsc,
	"DESC": ListTagsSortOrderDesc,
}

// GetListTagsSortOrderEnumValues Enumerates the set of values for ListTagsSortOrderEnum
func GetListTagsSortOrderEnumValues() []ListTagsSortOrderEnum {
	values := make([]ListTagsSortOrderEnum, 0)
	for _, v := range mappingListTagsSortOrder {
		values = append(values, v)
	}
	return values
}
