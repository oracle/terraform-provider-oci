// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListTermRelationshipsRequest wrapper for the ListTermRelationships operation
type ListTermRelationshipsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique glossary key.
	GlossaryKey *string `mandatory:"true" contributesTo:"path" name:"glossaryKey"`

	// Unique glossary term key.
	TermKey *string `mandatory:"true" contributesTo:"path" name:"termKey"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu".
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListTermRelationshipsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies the fields to return in a term relationship summary response.
	Fields []ListTermRelationshipsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListTermRelationshipsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTermRelationshipsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListTermRelationshipsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTermRelationshipsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTermRelationshipsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTermRelationshipsResponse wrapper for the ListTermRelationships operation
type ListTermRelationshipsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TermRelationshipCollection instances
	TermRelationshipCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTermRelationshipsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTermRelationshipsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTermRelationshipsLifecycleStateEnum Enum with underlying type: string
type ListTermRelationshipsLifecycleStateEnum string

// Set of constants representing the allowable values for ListTermRelationshipsLifecycleStateEnum
const (
	ListTermRelationshipsLifecycleStateCreating ListTermRelationshipsLifecycleStateEnum = "CREATING"
	ListTermRelationshipsLifecycleStateActive   ListTermRelationshipsLifecycleStateEnum = "ACTIVE"
	ListTermRelationshipsLifecycleStateInactive ListTermRelationshipsLifecycleStateEnum = "INACTIVE"
	ListTermRelationshipsLifecycleStateUpdating ListTermRelationshipsLifecycleStateEnum = "UPDATING"
	ListTermRelationshipsLifecycleStateDeleting ListTermRelationshipsLifecycleStateEnum = "DELETING"
	ListTermRelationshipsLifecycleStateDeleted  ListTermRelationshipsLifecycleStateEnum = "DELETED"
	ListTermRelationshipsLifecycleStateFailed   ListTermRelationshipsLifecycleStateEnum = "FAILED"
	ListTermRelationshipsLifecycleStateMoving   ListTermRelationshipsLifecycleStateEnum = "MOVING"
)

var mappingListTermRelationshipsLifecycleState = map[string]ListTermRelationshipsLifecycleStateEnum{
	"CREATING": ListTermRelationshipsLifecycleStateCreating,
	"ACTIVE":   ListTermRelationshipsLifecycleStateActive,
	"INACTIVE": ListTermRelationshipsLifecycleStateInactive,
	"UPDATING": ListTermRelationshipsLifecycleStateUpdating,
	"DELETING": ListTermRelationshipsLifecycleStateDeleting,
	"DELETED":  ListTermRelationshipsLifecycleStateDeleted,
	"FAILED":   ListTermRelationshipsLifecycleStateFailed,
	"MOVING":   ListTermRelationshipsLifecycleStateMoving,
}

// GetListTermRelationshipsLifecycleStateEnumValues Enumerates the set of values for ListTermRelationshipsLifecycleStateEnum
func GetListTermRelationshipsLifecycleStateEnumValues() []ListTermRelationshipsLifecycleStateEnum {
	values := make([]ListTermRelationshipsLifecycleStateEnum, 0)
	for _, v := range mappingListTermRelationshipsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListTermRelationshipsFieldsEnum Enum with underlying type: string
type ListTermRelationshipsFieldsEnum string

// Set of constants representing the allowable values for ListTermRelationshipsFieldsEnum
const (
	ListTermRelationshipsFieldsKey                    ListTermRelationshipsFieldsEnum = "key"
	ListTermRelationshipsFieldsDisplayname            ListTermRelationshipsFieldsEnum = "displayName"
	ListTermRelationshipsFieldsDescription            ListTermRelationshipsFieldsEnum = "description"
	ListTermRelationshipsFieldsRelatedtermkey         ListTermRelationshipsFieldsEnum = "relatedTermKey"
	ListTermRelationshipsFieldsRelatedtermdisplayname ListTermRelationshipsFieldsEnum = "relatedTermDisplayName"
	ListTermRelationshipsFieldsParenttermkey          ListTermRelationshipsFieldsEnum = "parentTermKey"
	ListTermRelationshipsFieldsParenttermdisplayname  ListTermRelationshipsFieldsEnum = "parentTermDisplayName"
	ListTermRelationshipsFieldsLifecyclestate         ListTermRelationshipsFieldsEnum = "lifecycleState"
	ListTermRelationshipsFieldsTimecreated            ListTermRelationshipsFieldsEnum = "timeCreated"
	ListTermRelationshipsFieldsUri                    ListTermRelationshipsFieldsEnum = "uri"
)

var mappingListTermRelationshipsFields = map[string]ListTermRelationshipsFieldsEnum{
	"key":                    ListTermRelationshipsFieldsKey,
	"displayName":            ListTermRelationshipsFieldsDisplayname,
	"description":            ListTermRelationshipsFieldsDescription,
	"relatedTermKey":         ListTermRelationshipsFieldsRelatedtermkey,
	"relatedTermDisplayName": ListTermRelationshipsFieldsRelatedtermdisplayname,
	"parentTermKey":          ListTermRelationshipsFieldsParenttermkey,
	"parentTermDisplayName":  ListTermRelationshipsFieldsParenttermdisplayname,
	"lifecycleState":         ListTermRelationshipsFieldsLifecyclestate,
	"timeCreated":            ListTermRelationshipsFieldsTimecreated,
	"uri":                    ListTermRelationshipsFieldsUri,
}

// GetListTermRelationshipsFieldsEnumValues Enumerates the set of values for ListTermRelationshipsFieldsEnum
func GetListTermRelationshipsFieldsEnumValues() []ListTermRelationshipsFieldsEnum {
	values := make([]ListTermRelationshipsFieldsEnum, 0)
	for _, v := range mappingListTermRelationshipsFields {
		values = append(values, v)
	}
	return values
}

// ListTermRelationshipsSortByEnum Enum with underlying type: string
type ListTermRelationshipsSortByEnum string

// Set of constants representing the allowable values for ListTermRelationshipsSortByEnum
const (
	ListTermRelationshipsSortByTimecreated ListTermRelationshipsSortByEnum = "TIMECREATED"
	ListTermRelationshipsSortByDisplayname ListTermRelationshipsSortByEnum = "DISPLAYNAME"
)

var mappingListTermRelationshipsSortBy = map[string]ListTermRelationshipsSortByEnum{
	"TIMECREATED": ListTermRelationshipsSortByTimecreated,
	"DISPLAYNAME": ListTermRelationshipsSortByDisplayname,
}

// GetListTermRelationshipsSortByEnumValues Enumerates the set of values for ListTermRelationshipsSortByEnum
func GetListTermRelationshipsSortByEnumValues() []ListTermRelationshipsSortByEnum {
	values := make([]ListTermRelationshipsSortByEnum, 0)
	for _, v := range mappingListTermRelationshipsSortBy {
		values = append(values, v)
	}
	return values
}

// ListTermRelationshipsSortOrderEnum Enum with underlying type: string
type ListTermRelationshipsSortOrderEnum string

// Set of constants representing the allowable values for ListTermRelationshipsSortOrderEnum
const (
	ListTermRelationshipsSortOrderAsc  ListTermRelationshipsSortOrderEnum = "ASC"
	ListTermRelationshipsSortOrderDesc ListTermRelationshipsSortOrderEnum = "DESC"
)

var mappingListTermRelationshipsSortOrder = map[string]ListTermRelationshipsSortOrderEnum{
	"ASC":  ListTermRelationshipsSortOrderAsc,
	"DESC": ListTermRelationshipsSortOrderDesc,
}

// GetListTermRelationshipsSortOrderEnumValues Enumerates the set of values for ListTermRelationshipsSortOrderEnum
func GetListTermRelationshipsSortOrderEnumValues() []ListTermRelationshipsSortOrderEnum {
	values := make([]ListTermRelationshipsSortOrderEnum, 0)
	for _, v := range mappingListTermRelationshipsSortOrder {
		values = append(values, v)
	}
	return values
}
