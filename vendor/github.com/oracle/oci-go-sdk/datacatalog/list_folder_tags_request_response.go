// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListFolderTagsRequest wrapper for the ListFolderTags operation
type ListFolderTagsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique folder key.
	FolderKey *string `mandatory:"true" contributesTo:"path" name:"folderKey"`

	// Immutable resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListFolderTagsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique key of the related term.
	TermKey *string `mandatory:"false" contributesTo:"query" name:"termKey"`

	// Path of the related term.
	TermPath *string `mandatory:"false" contributesTo:"query" name:"termPath"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// Specifies the fields to return in a folder tag summary response.
	Fields []ListFolderTagsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListFolderTagsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListFolderTagsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListFolderTagsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFolderTagsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFolderTagsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFolderTagsResponse wrapper for the ListFolderTags operation
type ListFolderTagsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FolderTagCollection instances
	FolderTagCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFolderTagsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFolderTagsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFolderTagsLifecycleStateEnum Enum with underlying type: string
type ListFolderTagsLifecycleStateEnum string

// Set of constants representing the allowable values for ListFolderTagsLifecycleStateEnum
const (
	ListFolderTagsLifecycleStateCreating ListFolderTagsLifecycleStateEnum = "CREATING"
	ListFolderTagsLifecycleStateActive   ListFolderTagsLifecycleStateEnum = "ACTIVE"
	ListFolderTagsLifecycleStateInactive ListFolderTagsLifecycleStateEnum = "INACTIVE"
	ListFolderTagsLifecycleStateUpdating ListFolderTagsLifecycleStateEnum = "UPDATING"
	ListFolderTagsLifecycleStateDeleting ListFolderTagsLifecycleStateEnum = "DELETING"
	ListFolderTagsLifecycleStateDeleted  ListFolderTagsLifecycleStateEnum = "DELETED"
	ListFolderTagsLifecycleStateFailed   ListFolderTagsLifecycleStateEnum = "FAILED"
	ListFolderTagsLifecycleStateMoving   ListFolderTagsLifecycleStateEnum = "MOVING"
)

var mappingListFolderTagsLifecycleState = map[string]ListFolderTagsLifecycleStateEnum{
	"CREATING": ListFolderTagsLifecycleStateCreating,
	"ACTIVE":   ListFolderTagsLifecycleStateActive,
	"INACTIVE": ListFolderTagsLifecycleStateInactive,
	"UPDATING": ListFolderTagsLifecycleStateUpdating,
	"DELETING": ListFolderTagsLifecycleStateDeleting,
	"DELETED":  ListFolderTagsLifecycleStateDeleted,
	"FAILED":   ListFolderTagsLifecycleStateFailed,
	"MOVING":   ListFolderTagsLifecycleStateMoving,
}

// GetListFolderTagsLifecycleStateEnumValues Enumerates the set of values for ListFolderTagsLifecycleStateEnum
func GetListFolderTagsLifecycleStateEnumValues() []ListFolderTagsLifecycleStateEnum {
	values := make([]ListFolderTagsLifecycleStateEnum, 0)
	for _, v := range mappingListFolderTagsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListFolderTagsFieldsEnum Enum with underlying type: string
type ListFolderTagsFieldsEnum string

// Set of constants representing the allowable values for ListFolderTagsFieldsEnum
const (
	ListFolderTagsFieldsKey             ListFolderTagsFieldsEnum = "key"
	ListFolderTagsFieldsName            ListFolderTagsFieldsEnum = "name"
	ListFolderTagsFieldsTermkey         ListFolderTagsFieldsEnum = "termKey"
	ListFolderTagsFieldsTermpath        ListFolderTagsFieldsEnum = "termPath"
	ListFolderTagsFieldsTermdescription ListFolderTagsFieldsEnum = "termDescription"
	ListFolderTagsFieldsLifecyclestate  ListFolderTagsFieldsEnum = "lifecycleState"
	ListFolderTagsFieldsTimecreated     ListFolderTagsFieldsEnum = "timeCreated"
	ListFolderTagsFieldsUri             ListFolderTagsFieldsEnum = "uri"
	ListFolderTagsFieldsGlossarykey     ListFolderTagsFieldsEnum = "glossaryKey"
	ListFolderTagsFieldsFolderkey       ListFolderTagsFieldsEnum = "folderKey"
)

var mappingListFolderTagsFields = map[string]ListFolderTagsFieldsEnum{
	"key":             ListFolderTagsFieldsKey,
	"name":            ListFolderTagsFieldsName,
	"termKey":         ListFolderTagsFieldsTermkey,
	"termPath":        ListFolderTagsFieldsTermpath,
	"termDescription": ListFolderTagsFieldsTermdescription,
	"lifecycleState":  ListFolderTagsFieldsLifecyclestate,
	"timeCreated":     ListFolderTagsFieldsTimecreated,
	"uri":             ListFolderTagsFieldsUri,
	"glossaryKey":     ListFolderTagsFieldsGlossarykey,
	"folderKey":       ListFolderTagsFieldsFolderkey,
}

// GetListFolderTagsFieldsEnumValues Enumerates the set of values for ListFolderTagsFieldsEnum
func GetListFolderTagsFieldsEnumValues() []ListFolderTagsFieldsEnum {
	values := make([]ListFolderTagsFieldsEnum, 0)
	for _, v := range mappingListFolderTagsFields {
		values = append(values, v)
	}
	return values
}

// ListFolderTagsSortByEnum Enum with underlying type: string
type ListFolderTagsSortByEnum string

// Set of constants representing the allowable values for ListFolderTagsSortByEnum
const (
	ListFolderTagsSortByTimecreated ListFolderTagsSortByEnum = "TIMECREATED"
	ListFolderTagsSortByDisplayname ListFolderTagsSortByEnum = "DISPLAYNAME"
)

var mappingListFolderTagsSortBy = map[string]ListFolderTagsSortByEnum{
	"TIMECREATED": ListFolderTagsSortByTimecreated,
	"DISPLAYNAME": ListFolderTagsSortByDisplayname,
}

// GetListFolderTagsSortByEnumValues Enumerates the set of values for ListFolderTagsSortByEnum
func GetListFolderTagsSortByEnumValues() []ListFolderTagsSortByEnum {
	values := make([]ListFolderTagsSortByEnum, 0)
	for _, v := range mappingListFolderTagsSortBy {
		values = append(values, v)
	}
	return values
}

// ListFolderTagsSortOrderEnum Enum with underlying type: string
type ListFolderTagsSortOrderEnum string

// Set of constants representing the allowable values for ListFolderTagsSortOrderEnum
const (
	ListFolderTagsSortOrderAsc  ListFolderTagsSortOrderEnum = "ASC"
	ListFolderTagsSortOrderDesc ListFolderTagsSortOrderEnum = "DESC"
)

var mappingListFolderTagsSortOrder = map[string]ListFolderTagsSortOrderEnum{
	"ASC":  ListFolderTagsSortOrderAsc,
	"DESC": ListFolderTagsSortOrderDesc,
}

// GetListFolderTagsSortOrderEnumValues Enumerates the set of values for ListFolderTagsSortOrderEnum
func GetListFolderTagsSortOrderEnumValues() []ListFolderTagsSortOrderEnum {
	values := make([]ListFolderTagsSortOrderEnum, 0)
	for _, v := range mappingListFolderTagsSortOrder {
		values = append(values, v)
	}
	return values
}
