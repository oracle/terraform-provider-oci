// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListTermsRequest wrapper for the ListTerms operation
type ListTermsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique glossary key.
	GlossaryKey *string `mandatory:"true" contributesTo:"path" name:"glossaryKey"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu".
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListTermsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique key of the parent term.
	ParentTermKey *string `mandatory:"false" contributesTo:"query" name:"parentTermKey"`

	// Indicates whether a term may contain child terms.
	IsAllowedToHaveChildTerms *bool `mandatory:"false" contributesTo:"query" name:"isAllowedToHaveChildTerms"`

	// Status of the approval workflow for this business term in the glossary.
	WorkflowStatus ListTermsWorkflowStatusEnum `mandatory:"false" contributesTo:"query" name:"workflowStatus" omitEmpty:"true"`

	// Full path of the resource for resources that support paths.
	Path *string `mandatory:"false" contributesTo:"query" name:"path"`

	// Specifies the fields to return in a term summary response.
	Fields []ListTermsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListTermsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTermsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListTermsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTermsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTermsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTermsResponse wrapper for the ListTerms operation
type ListTermsResponse struct {

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

func (response ListTermsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTermsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTermsLifecycleStateEnum Enum with underlying type: string
type ListTermsLifecycleStateEnum string

// Set of constants representing the allowable values for ListTermsLifecycleStateEnum
const (
	ListTermsLifecycleStateCreating ListTermsLifecycleStateEnum = "CREATING"
	ListTermsLifecycleStateActive   ListTermsLifecycleStateEnum = "ACTIVE"
	ListTermsLifecycleStateInactive ListTermsLifecycleStateEnum = "INACTIVE"
	ListTermsLifecycleStateUpdating ListTermsLifecycleStateEnum = "UPDATING"
	ListTermsLifecycleStateDeleting ListTermsLifecycleStateEnum = "DELETING"
	ListTermsLifecycleStateDeleted  ListTermsLifecycleStateEnum = "DELETED"
	ListTermsLifecycleStateFailed   ListTermsLifecycleStateEnum = "FAILED"
	ListTermsLifecycleStateMoving   ListTermsLifecycleStateEnum = "MOVING"
)

var mappingListTermsLifecycleState = map[string]ListTermsLifecycleStateEnum{
	"CREATING": ListTermsLifecycleStateCreating,
	"ACTIVE":   ListTermsLifecycleStateActive,
	"INACTIVE": ListTermsLifecycleStateInactive,
	"UPDATING": ListTermsLifecycleStateUpdating,
	"DELETING": ListTermsLifecycleStateDeleting,
	"DELETED":  ListTermsLifecycleStateDeleted,
	"FAILED":   ListTermsLifecycleStateFailed,
	"MOVING":   ListTermsLifecycleStateMoving,
}

// GetListTermsLifecycleStateEnumValues Enumerates the set of values for ListTermsLifecycleStateEnum
func GetListTermsLifecycleStateEnumValues() []ListTermsLifecycleStateEnum {
	values := make([]ListTermsLifecycleStateEnum, 0)
	for _, v := range mappingListTermsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListTermsWorkflowStatusEnum Enum with underlying type: string
type ListTermsWorkflowStatusEnum string

// Set of constants representing the allowable values for ListTermsWorkflowStatusEnum
const (
	ListTermsWorkflowStatusNew         ListTermsWorkflowStatusEnum = "NEW"
	ListTermsWorkflowStatusApproved    ListTermsWorkflowStatusEnum = "APPROVED"
	ListTermsWorkflowStatusUnderReview ListTermsWorkflowStatusEnum = "UNDER_REVIEW"
	ListTermsWorkflowStatusEscalated   ListTermsWorkflowStatusEnum = "ESCALATED"
)

var mappingListTermsWorkflowStatus = map[string]ListTermsWorkflowStatusEnum{
	"NEW":          ListTermsWorkflowStatusNew,
	"APPROVED":     ListTermsWorkflowStatusApproved,
	"UNDER_REVIEW": ListTermsWorkflowStatusUnderReview,
	"ESCALATED":    ListTermsWorkflowStatusEscalated,
}

// GetListTermsWorkflowStatusEnumValues Enumerates the set of values for ListTermsWorkflowStatusEnum
func GetListTermsWorkflowStatusEnumValues() []ListTermsWorkflowStatusEnum {
	values := make([]ListTermsWorkflowStatusEnum, 0)
	for _, v := range mappingListTermsWorkflowStatus {
		values = append(values, v)
	}
	return values
}

// ListTermsFieldsEnum Enum with underlying type: string
type ListTermsFieldsEnum string

// Set of constants representing the allowable values for ListTermsFieldsEnum
const (
	ListTermsFieldsKey                       ListTermsFieldsEnum = "key"
	ListTermsFieldsDisplayname               ListTermsFieldsEnum = "displayName"
	ListTermsFieldsDescription               ListTermsFieldsEnum = "description"
	ListTermsFieldsGlossarykey               ListTermsFieldsEnum = "glossaryKey"
	ListTermsFieldsParenttermkey             ListTermsFieldsEnum = "parentTermKey"
	ListTermsFieldsIsallowedtohavechildterms ListTermsFieldsEnum = "isAllowedToHaveChildTerms"
	ListTermsFieldsPath                      ListTermsFieldsEnum = "path"
	ListTermsFieldsLifecyclestate            ListTermsFieldsEnum = "lifecycleState"
	ListTermsFieldsTimecreated               ListTermsFieldsEnum = "timeCreated"
	ListTermsFieldsWorkflowstatus            ListTermsFieldsEnum = "workflowStatus"
	ListTermsFieldsAssociatedobjectcount     ListTermsFieldsEnum = "associatedObjectCount"
	ListTermsFieldsUri                       ListTermsFieldsEnum = "uri"
)

var mappingListTermsFields = map[string]ListTermsFieldsEnum{
	"key":                       ListTermsFieldsKey,
	"displayName":               ListTermsFieldsDisplayname,
	"description":               ListTermsFieldsDescription,
	"glossaryKey":               ListTermsFieldsGlossarykey,
	"parentTermKey":             ListTermsFieldsParenttermkey,
	"isAllowedToHaveChildTerms": ListTermsFieldsIsallowedtohavechildterms,
	"path":                  ListTermsFieldsPath,
	"lifecycleState":        ListTermsFieldsLifecyclestate,
	"timeCreated":           ListTermsFieldsTimecreated,
	"workflowStatus":        ListTermsFieldsWorkflowstatus,
	"associatedObjectCount": ListTermsFieldsAssociatedobjectcount,
	"uri": ListTermsFieldsUri,
}

// GetListTermsFieldsEnumValues Enumerates the set of values for ListTermsFieldsEnum
func GetListTermsFieldsEnumValues() []ListTermsFieldsEnum {
	values := make([]ListTermsFieldsEnum, 0)
	for _, v := range mappingListTermsFields {
		values = append(values, v)
	}
	return values
}

// ListTermsSortByEnum Enum with underlying type: string
type ListTermsSortByEnum string

// Set of constants representing the allowable values for ListTermsSortByEnum
const (
	ListTermsSortByTimecreated ListTermsSortByEnum = "TIMECREATED"
	ListTermsSortByDisplayname ListTermsSortByEnum = "DISPLAYNAME"
)

var mappingListTermsSortBy = map[string]ListTermsSortByEnum{
	"TIMECREATED": ListTermsSortByTimecreated,
	"DISPLAYNAME": ListTermsSortByDisplayname,
}

// GetListTermsSortByEnumValues Enumerates the set of values for ListTermsSortByEnum
func GetListTermsSortByEnumValues() []ListTermsSortByEnum {
	values := make([]ListTermsSortByEnum, 0)
	for _, v := range mappingListTermsSortBy {
		values = append(values, v)
	}
	return values
}

// ListTermsSortOrderEnum Enum with underlying type: string
type ListTermsSortOrderEnum string

// Set of constants representing the allowable values for ListTermsSortOrderEnum
const (
	ListTermsSortOrderAsc  ListTermsSortOrderEnum = "ASC"
	ListTermsSortOrderDesc ListTermsSortOrderEnum = "DESC"
)

var mappingListTermsSortOrder = map[string]ListTermsSortOrderEnum{
	"ASC":  ListTermsSortOrderAsc,
	"DESC": ListTermsSortOrderDesc,
}

// GetListTermsSortOrderEnumValues Enumerates the set of values for ListTermsSortOrderEnum
func GetListTermsSortOrderEnumValues() []ListTermsSortOrderEnum {
	values := make([]ListTermsSortOrderEnum, 0)
	for _, v := range mappingListTermsSortOrder {
		values = append(values, v)
	}
	return values
}
