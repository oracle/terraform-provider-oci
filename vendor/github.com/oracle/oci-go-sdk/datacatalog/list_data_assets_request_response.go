// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListDataAssetsRequest wrapper for the ListDataAssets operation
type ListDataAssetsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu".
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListDataAssetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Unique external identifier of this resource in the external source system.
	ExternalKey *string `mandatory:"false" contributesTo:"query" name:"externalKey"`

	// The key of the object type.
	TypeKey *string `mandatory:"false" contributesTo:"query" name:"typeKey"`

	// Specifies the fields to return in a data asset summary response.
	Fields []ListDataAssetsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListDataAssetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataAssetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListDataAssetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataAssetsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataAssetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDataAssetsResponse wrapper for the ListDataAssets operation
type ListDataAssetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataAssetCollection instances
	DataAssetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataAssetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataAssetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataAssetsLifecycleStateEnum Enum with underlying type: string
type ListDataAssetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDataAssetsLifecycleStateEnum
const (
	ListDataAssetsLifecycleStateCreating ListDataAssetsLifecycleStateEnum = "CREATING"
	ListDataAssetsLifecycleStateActive   ListDataAssetsLifecycleStateEnum = "ACTIVE"
	ListDataAssetsLifecycleStateInactive ListDataAssetsLifecycleStateEnum = "INACTIVE"
	ListDataAssetsLifecycleStateUpdating ListDataAssetsLifecycleStateEnum = "UPDATING"
	ListDataAssetsLifecycleStateDeleting ListDataAssetsLifecycleStateEnum = "DELETING"
	ListDataAssetsLifecycleStateDeleted  ListDataAssetsLifecycleStateEnum = "DELETED"
	ListDataAssetsLifecycleStateFailed   ListDataAssetsLifecycleStateEnum = "FAILED"
	ListDataAssetsLifecycleStateMoving   ListDataAssetsLifecycleStateEnum = "MOVING"
)

var mappingListDataAssetsLifecycleState = map[string]ListDataAssetsLifecycleStateEnum{
	"CREATING": ListDataAssetsLifecycleStateCreating,
	"ACTIVE":   ListDataAssetsLifecycleStateActive,
	"INACTIVE": ListDataAssetsLifecycleStateInactive,
	"UPDATING": ListDataAssetsLifecycleStateUpdating,
	"DELETING": ListDataAssetsLifecycleStateDeleting,
	"DELETED":  ListDataAssetsLifecycleStateDeleted,
	"FAILED":   ListDataAssetsLifecycleStateFailed,
	"MOVING":   ListDataAssetsLifecycleStateMoving,
}

// GetListDataAssetsLifecycleStateEnumValues Enumerates the set of values for ListDataAssetsLifecycleStateEnum
func GetListDataAssetsLifecycleStateEnumValues() []ListDataAssetsLifecycleStateEnum {
	values := make([]ListDataAssetsLifecycleStateEnum, 0)
	for _, v := range mappingListDataAssetsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListDataAssetsFieldsEnum Enum with underlying type: string
type ListDataAssetsFieldsEnum string

// Set of constants representing the allowable values for ListDataAssetsFieldsEnum
const (
	ListDataAssetsFieldsKey            ListDataAssetsFieldsEnum = "key"
	ListDataAssetsFieldsDisplayname    ListDataAssetsFieldsEnum = "displayName"
	ListDataAssetsFieldsDescription    ListDataAssetsFieldsEnum = "description"
	ListDataAssetsFieldsCatalogid      ListDataAssetsFieldsEnum = "catalogId"
	ListDataAssetsFieldsExternalkey    ListDataAssetsFieldsEnum = "externalKey"
	ListDataAssetsFieldsTypekey        ListDataAssetsFieldsEnum = "typeKey"
	ListDataAssetsFieldsLifecyclestate ListDataAssetsFieldsEnum = "lifecycleState"
	ListDataAssetsFieldsTimecreated    ListDataAssetsFieldsEnum = "timeCreated"
	ListDataAssetsFieldsUri            ListDataAssetsFieldsEnum = "uri"
)

var mappingListDataAssetsFields = map[string]ListDataAssetsFieldsEnum{
	"key":            ListDataAssetsFieldsKey,
	"displayName":    ListDataAssetsFieldsDisplayname,
	"description":    ListDataAssetsFieldsDescription,
	"catalogId":      ListDataAssetsFieldsCatalogid,
	"externalKey":    ListDataAssetsFieldsExternalkey,
	"typeKey":        ListDataAssetsFieldsTypekey,
	"lifecycleState": ListDataAssetsFieldsLifecyclestate,
	"timeCreated":    ListDataAssetsFieldsTimecreated,
	"uri":            ListDataAssetsFieldsUri,
}

// GetListDataAssetsFieldsEnumValues Enumerates the set of values for ListDataAssetsFieldsEnum
func GetListDataAssetsFieldsEnumValues() []ListDataAssetsFieldsEnum {
	values := make([]ListDataAssetsFieldsEnum, 0)
	for _, v := range mappingListDataAssetsFields {
		values = append(values, v)
	}
	return values
}

// ListDataAssetsSortByEnum Enum with underlying type: string
type ListDataAssetsSortByEnum string

// Set of constants representing the allowable values for ListDataAssetsSortByEnum
const (
	ListDataAssetsSortByTimecreated ListDataAssetsSortByEnum = "TIMECREATED"
	ListDataAssetsSortByDisplayname ListDataAssetsSortByEnum = "DISPLAYNAME"
)

var mappingListDataAssetsSortBy = map[string]ListDataAssetsSortByEnum{
	"TIMECREATED": ListDataAssetsSortByTimecreated,
	"DISPLAYNAME": ListDataAssetsSortByDisplayname,
}

// GetListDataAssetsSortByEnumValues Enumerates the set of values for ListDataAssetsSortByEnum
func GetListDataAssetsSortByEnumValues() []ListDataAssetsSortByEnum {
	values := make([]ListDataAssetsSortByEnum, 0)
	for _, v := range mappingListDataAssetsSortBy {
		values = append(values, v)
	}
	return values
}

// ListDataAssetsSortOrderEnum Enum with underlying type: string
type ListDataAssetsSortOrderEnum string

// Set of constants representing the allowable values for ListDataAssetsSortOrderEnum
const (
	ListDataAssetsSortOrderAsc  ListDataAssetsSortOrderEnum = "ASC"
	ListDataAssetsSortOrderDesc ListDataAssetsSortOrderEnum = "DESC"
)

var mappingListDataAssetsSortOrder = map[string]ListDataAssetsSortOrderEnum{
	"ASC":  ListDataAssetsSortOrderAsc,
	"DESC": ListDataAssetsSortOrderDesc,
}

// GetListDataAssetsSortOrderEnumValues Enumerates the set of values for ListDataAssetsSortOrderEnum
func GetListDataAssetsSortOrderEnumValues() []ListDataAssetsSortOrderEnum {
	values := make([]ListDataAssetsSortOrderEnum, 0)
	for _, v := range mappingListDataAssetsSortOrder {
		values = append(values, v)
	}
	return values
}
