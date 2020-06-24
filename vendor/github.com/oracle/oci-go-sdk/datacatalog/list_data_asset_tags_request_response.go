// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListDataAssetTagsRequest wrapper for the ListDataAssetTags operation
type ListDataAssetTagsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Immutable resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListDataAssetTagsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique key of the related term.
	TermKey *string `mandatory:"false" contributesTo:"query" name:"termKey"`

	// Path of the related term.
	TermPath *string `mandatory:"false" contributesTo:"query" name:"termPath"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// Specifies the fields to return in a data asset tag summary response.
	Fields []ListDataAssetTagsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListDataAssetTagsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataAssetTagsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListDataAssetTagsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataAssetTagsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataAssetTagsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDataAssetTagsResponse wrapper for the ListDataAssetTags operation
type ListDataAssetTagsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataAssetTagCollection instances
	DataAssetTagCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataAssetTagsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataAssetTagsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataAssetTagsLifecycleStateEnum Enum with underlying type: string
type ListDataAssetTagsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDataAssetTagsLifecycleStateEnum
const (
	ListDataAssetTagsLifecycleStateCreating ListDataAssetTagsLifecycleStateEnum = "CREATING"
	ListDataAssetTagsLifecycleStateActive   ListDataAssetTagsLifecycleStateEnum = "ACTIVE"
	ListDataAssetTagsLifecycleStateInactive ListDataAssetTagsLifecycleStateEnum = "INACTIVE"
	ListDataAssetTagsLifecycleStateUpdating ListDataAssetTagsLifecycleStateEnum = "UPDATING"
	ListDataAssetTagsLifecycleStateDeleting ListDataAssetTagsLifecycleStateEnum = "DELETING"
	ListDataAssetTagsLifecycleStateDeleted  ListDataAssetTagsLifecycleStateEnum = "DELETED"
	ListDataAssetTagsLifecycleStateFailed   ListDataAssetTagsLifecycleStateEnum = "FAILED"
	ListDataAssetTagsLifecycleStateMoving   ListDataAssetTagsLifecycleStateEnum = "MOVING"
)

var mappingListDataAssetTagsLifecycleState = map[string]ListDataAssetTagsLifecycleStateEnum{
	"CREATING": ListDataAssetTagsLifecycleStateCreating,
	"ACTIVE":   ListDataAssetTagsLifecycleStateActive,
	"INACTIVE": ListDataAssetTagsLifecycleStateInactive,
	"UPDATING": ListDataAssetTagsLifecycleStateUpdating,
	"DELETING": ListDataAssetTagsLifecycleStateDeleting,
	"DELETED":  ListDataAssetTagsLifecycleStateDeleted,
	"FAILED":   ListDataAssetTagsLifecycleStateFailed,
	"MOVING":   ListDataAssetTagsLifecycleStateMoving,
}

// GetListDataAssetTagsLifecycleStateEnumValues Enumerates the set of values for ListDataAssetTagsLifecycleStateEnum
func GetListDataAssetTagsLifecycleStateEnumValues() []ListDataAssetTagsLifecycleStateEnum {
	values := make([]ListDataAssetTagsLifecycleStateEnum, 0)
	for _, v := range mappingListDataAssetTagsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListDataAssetTagsFieldsEnum Enum with underlying type: string
type ListDataAssetTagsFieldsEnum string

// Set of constants representing the allowable values for ListDataAssetTagsFieldsEnum
const (
	ListDataAssetTagsFieldsKey             ListDataAssetTagsFieldsEnum = "key"
	ListDataAssetTagsFieldsName            ListDataAssetTagsFieldsEnum = "name"
	ListDataAssetTagsFieldsTermkey         ListDataAssetTagsFieldsEnum = "termKey"
	ListDataAssetTagsFieldsTermpath        ListDataAssetTagsFieldsEnum = "termPath"
	ListDataAssetTagsFieldsTermdescription ListDataAssetTagsFieldsEnum = "termDescription"
	ListDataAssetTagsFieldsLifecyclestate  ListDataAssetTagsFieldsEnum = "lifecycleState"
	ListDataAssetTagsFieldsTimecreated     ListDataAssetTagsFieldsEnum = "timeCreated"
	ListDataAssetTagsFieldsUri             ListDataAssetTagsFieldsEnum = "uri"
	ListDataAssetTagsFieldsGlossarykey     ListDataAssetTagsFieldsEnum = "glossaryKey"
	ListDataAssetTagsFieldsDataassetkey    ListDataAssetTagsFieldsEnum = "dataAssetKey"
)

var mappingListDataAssetTagsFields = map[string]ListDataAssetTagsFieldsEnum{
	"key":             ListDataAssetTagsFieldsKey,
	"name":            ListDataAssetTagsFieldsName,
	"termKey":         ListDataAssetTagsFieldsTermkey,
	"termPath":        ListDataAssetTagsFieldsTermpath,
	"termDescription": ListDataAssetTagsFieldsTermdescription,
	"lifecycleState":  ListDataAssetTagsFieldsLifecyclestate,
	"timeCreated":     ListDataAssetTagsFieldsTimecreated,
	"uri":             ListDataAssetTagsFieldsUri,
	"glossaryKey":     ListDataAssetTagsFieldsGlossarykey,
	"dataAssetKey":    ListDataAssetTagsFieldsDataassetkey,
}

// GetListDataAssetTagsFieldsEnumValues Enumerates the set of values for ListDataAssetTagsFieldsEnum
func GetListDataAssetTagsFieldsEnumValues() []ListDataAssetTagsFieldsEnum {
	values := make([]ListDataAssetTagsFieldsEnum, 0)
	for _, v := range mappingListDataAssetTagsFields {
		values = append(values, v)
	}
	return values
}

// ListDataAssetTagsSortByEnum Enum with underlying type: string
type ListDataAssetTagsSortByEnum string

// Set of constants representing the allowable values for ListDataAssetTagsSortByEnum
const (
	ListDataAssetTagsSortByTimecreated ListDataAssetTagsSortByEnum = "TIMECREATED"
	ListDataAssetTagsSortByDisplayname ListDataAssetTagsSortByEnum = "DISPLAYNAME"
)

var mappingListDataAssetTagsSortBy = map[string]ListDataAssetTagsSortByEnum{
	"TIMECREATED": ListDataAssetTagsSortByTimecreated,
	"DISPLAYNAME": ListDataAssetTagsSortByDisplayname,
}

// GetListDataAssetTagsSortByEnumValues Enumerates the set of values for ListDataAssetTagsSortByEnum
func GetListDataAssetTagsSortByEnumValues() []ListDataAssetTagsSortByEnum {
	values := make([]ListDataAssetTagsSortByEnum, 0)
	for _, v := range mappingListDataAssetTagsSortBy {
		values = append(values, v)
	}
	return values
}

// ListDataAssetTagsSortOrderEnum Enum with underlying type: string
type ListDataAssetTagsSortOrderEnum string

// Set of constants representing the allowable values for ListDataAssetTagsSortOrderEnum
const (
	ListDataAssetTagsSortOrderAsc  ListDataAssetTagsSortOrderEnum = "ASC"
	ListDataAssetTagsSortOrderDesc ListDataAssetTagsSortOrderEnum = "DESC"
)

var mappingListDataAssetTagsSortOrder = map[string]ListDataAssetTagsSortOrderEnum{
	"ASC":  ListDataAssetTagsSortOrderAsc,
	"DESC": ListDataAssetTagsSortOrderDesc,
}

// GetListDataAssetTagsSortOrderEnumValues Enumerates the set of values for ListDataAssetTagsSortOrderEnum
func GetListDataAssetTagsSortOrderEnumValues() []ListDataAssetTagsSortOrderEnum {
	values := make([]ListDataAssetTagsSortOrderEnum, 0)
	for _, v := range mappingListDataAssetTagsSortOrder {
		values = append(values, v)
	}
	return values
}
