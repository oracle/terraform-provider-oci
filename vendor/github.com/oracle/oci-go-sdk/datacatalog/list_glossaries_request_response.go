// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListGlossariesRequest wrapper for the ListGlossaries operation
type ListGlossariesRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu".
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListGlossariesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Specifies the fields to return in a glossary summary response.
	Fields []ListGlossariesFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListGlossariesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListGlossariesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListGlossariesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGlossariesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGlossariesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListGlossariesResponse wrapper for the ListGlossaries operation
type ListGlossariesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GlossaryCollection instances
	GlossaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGlossariesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGlossariesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGlossariesLifecycleStateEnum Enum with underlying type: string
type ListGlossariesLifecycleStateEnum string

// Set of constants representing the allowable values for ListGlossariesLifecycleStateEnum
const (
	ListGlossariesLifecycleStateCreating ListGlossariesLifecycleStateEnum = "CREATING"
	ListGlossariesLifecycleStateActive   ListGlossariesLifecycleStateEnum = "ACTIVE"
	ListGlossariesLifecycleStateInactive ListGlossariesLifecycleStateEnum = "INACTIVE"
	ListGlossariesLifecycleStateUpdating ListGlossariesLifecycleStateEnum = "UPDATING"
	ListGlossariesLifecycleStateDeleting ListGlossariesLifecycleStateEnum = "DELETING"
	ListGlossariesLifecycleStateDeleted  ListGlossariesLifecycleStateEnum = "DELETED"
	ListGlossariesLifecycleStateFailed   ListGlossariesLifecycleStateEnum = "FAILED"
	ListGlossariesLifecycleStateMoving   ListGlossariesLifecycleStateEnum = "MOVING"
)

var mappingListGlossariesLifecycleState = map[string]ListGlossariesLifecycleStateEnum{
	"CREATING": ListGlossariesLifecycleStateCreating,
	"ACTIVE":   ListGlossariesLifecycleStateActive,
	"INACTIVE": ListGlossariesLifecycleStateInactive,
	"UPDATING": ListGlossariesLifecycleStateUpdating,
	"DELETING": ListGlossariesLifecycleStateDeleting,
	"DELETED":  ListGlossariesLifecycleStateDeleted,
	"FAILED":   ListGlossariesLifecycleStateFailed,
	"MOVING":   ListGlossariesLifecycleStateMoving,
}

// GetListGlossariesLifecycleStateEnumValues Enumerates the set of values for ListGlossariesLifecycleStateEnum
func GetListGlossariesLifecycleStateEnumValues() []ListGlossariesLifecycleStateEnum {
	values := make([]ListGlossariesLifecycleStateEnum, 0)
	for _, v := range mappingListGlossariesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListGlossariesFieldsEnum Enum with underlying type: string
type ListGlossariesFieldsEnum string

// Set of constants representing the allowable values for ListGlossariesFieldsEnum
const (
	ListGlossariesFieldsKey            ListGlossariesFieldsEnum = "key"
	ListGlossariesFieldsDisplayname    ListGlossariesFieldsEnum = "displayName"
	ListGlossariesFieldsDescription    ListGlossariesFieldsEnum = "description"
	ListGlossariesFieldsCatalogid      ListGlossariesFieldsEnum = "catalogId"
	ListGlossariesFieldsLifecyclestate ListGlossariesFieldsEnum = "lifecycleState"
	ListGlossariesFieldsTimecreated    ListGlossariesFieldsEnum = "timeCreated"
	ListGlossariesFieldsUri            ListGlossariesFieldsEnum = "uri"
	ListGlossariesFieldsWorkflowstatus ListGlossariesFieldsEnum = "workflowStatus"
)

var mappingListGlossariesFields = map[string]ListGlossariesFieldsEnum{
	"key":            ListGlossariesFieldsKey,
	"displayName":    ListGlossariesFieldsDisplayname,
	"description":    ListGlossariesFieldsDescription,
	"catalogId":      ListGlossariesFieldsCatalogid,
	"lifecycleState": ListGlossariesFieldsLifecyclestate,
	"timeCreated":    ListGlossariesFieldsTimecreated,
	"uri":            ListGlossariesFieldsUri,
	"workflowStatus": ListGlossariesFieldsWorkflowstatus,
}

// GetListGlossariesFieldsEnumValues Enumerates the set of values for ListGlossariesFieldsEnum
func GetListGlossariesFieldsEnumValues() []ListGlossariesFieldsEnum {
	values := make([]ListGlossariesFieldsEnum, 0)
	for _, v := range mappingListGlossariesFields {
		values = append(values, v)
	}
	return values
}

// ListGlossariesSortByEnum Enum with underlying type: string
type ListGlossariesSortByEnum string

// Set of constants representing the allowable values for ListGlossariesSortByEnum
const (
	ListGlossariesSortByTimecreated ListGlossariesSortByEnum = "TIMECREATED"
	ListGlossariesSortByDisplayname ListGlossariesSortByEnum = "DISPLAYNAME"
)

var mappingListGlossariesSortBy = map[string]ListGlossariesSortByEnum{
	"TIMECREATED": ListGlossariesSortByTimecreated,
	"DISPLAYNAME": ListGlossariesSortByDisplayname,
}

// GetListGlossariesSortByEnumValues Enumerates the set of values for ListGlossariesSortByEnum
func GetListGlossariesSortByEnumValues() []ListGlossariesSortByEnum {
	values := make([]ListGlossariesSortByEnum, 0)
	for _, v := range mappingListGlossariesSortBy {
		values = append(values, v)
	}
	return values
}

// ListGlossariesSortOrderEnum Enum with underlying type: string
type ListGlossariesSortOrderEnum string

// Set of constants representing the allowable values for ListGlossariesSortOrderEnum
const (
	ListGlossariesSortOrderAsc  ListGlossariesSortOrderEnum = "ASC"
	ListGlossariesSortOrderDesc ListGlossariesSortOrderEnum = "DESC"
)

var mappingListGlossariesSortOrder = map[string]ListGlossariesSortOrderEnum{
	"ASC":  ListGlossariesSortOrderAsc,
	"DESC": ListGlossariesSortOrderDesc,
}

// GetListGlossariesSortOrderEnumValues Enumerates the set of values for ListGlossariesSortOrderEnum
func GetListGlossariesSortOrderEnumValues() []ListGlossariesSortOrderEnum {
	values := make([]ListGlossariesSortOrderEnum, 0)
	for _, v := range mappingListGlossariesSortOrder {
		values = append(values, v)
	}
	return values
}
