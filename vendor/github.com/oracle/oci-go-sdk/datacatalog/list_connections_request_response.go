// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListConnectionsRequest wrapper for the ListConnections operation
type ListConnectionsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu".
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListConnectionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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

	// Time that the resource's status was last updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeStatusUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStatusUpdated"`

	// Indicates whether this connection is the default connection.
	IsDefault *bool `mandatory:"false" contributesTo:"query" name:"isDefault"`

	// Specifies the fields to return in a connection summary response.
	Fields []ListConnectionsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListConnectionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListConnectionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListConnectionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConnectionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConnectionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListConnectionsResponse wrapper for the ListConnections operation
type ListConnectionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConnectionCollection instances
	ConnectionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConnectionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConnectionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConnectionsLifecycleStateEnum Enum with underlying type: string
type ListConnectionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListConnectionsLifecycleStateEnum
const (
	ListConnectionsLifecycleStateCreating ListConnectionsLifecycleStateEnum = "CREATING"
	ListConnectionsLifecycleStateActive   ListConnectionsLifecycleStateEnum = "ACTIVE"
	ListConnectionsLifecycleStateInactive ListConnectionsLifecycleStateEnum = "INACTIVE"
	ListConnectionsLifecycleStateUpdating ListConnectionsLifecycleStateEnum = "UPDATING"
	ListConnectionsLifecycleStateDeleting ListConnectionsLifecycleStateEnum = "DELETING"
	ListConnectionsLifecycleStateDeleted  ListConnectionsLifecycleStateEnum = "DELETED"
	ListConnectionsLifecycleStateFailed   ListConnectionsLifecycleStateEnum = "FAILED"
	ListConnectionsLifecycleStateMoving   ListConnectionsLifecycleStateEnum = "MOVING"
)

var mappingListConnectionsLifecycleState = map[string]ListConnectionsLifecycleStateEnum{
	"CREATING": ListConnectionsLifecycleStateCreating,
	"ACTIVE":   ListConnectionsLifecycleStateActive,
	"INACTIVE": ListConnectionsLifecycleStateInactive,
	"UPDATING": ListConnectionsLifecycleStateUpdating,
	"DELETING": ListConnectionsLifecycleStateDeleting,
	"DELETED":  ListConnectionsLifecycleStateDeleted,
	"FAILED":   ListConnectionsLifecycleStateFailed,
	"MOVING":   ListConnectionsLifecycleStateMoving,
}

// GetListConnectionsLifecycleStateEnumValues Enumerates the set of values for ListConnectionsLifecycleStateEnum
func GetListConnectionsLifecycleStateEnumValues() []ListConnectionsLifecycleStateEnum {
	values := make([]ListConnectionsLifecycleStateEnum, 0)
	for _, v := range mappingListConnectionsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListConnectionsFieldsEnum Enum with underlying type: string
type ListConnectionsFieldsEnum string

// Set of constants representing the allowable values for ListConnectionsFieldsEnum
const (
	ListConnectionsFieldsKey            ListConnectionsFieldsEnum = "key"
	ListConnectionsFieldsDisplayname    ListConnectionsFieldsEnum = "displayName"
	ListConnectionsFieldsDescription    ListConnectionsFieldsEnum = "description"
	ListConnectionsFieldsDataassetkey   ListConnectionsFieldsEnum = "dataAssetKey"
	ListConnectionsFieldsTypekey        ListConnectionsFieldsEnum = "typeKey"
	ListConnectionsFieldsTimecreated    ListConnectionsFieldsEnum = "timeCreated"
	ListConnectionsFieldsExternalkey    ListConnectionsFieldsEnum = "externalKey"
	ListConnectionsFieldsLifecyclestate ListConnectionsFieldsEnum = "lifecycleState"
	ListConnectionsFieldsIsdefault      ListConnectionsFieldsEnum = "isDefault"
	ListConnectionsFieldsUri            ListConnectionsFieldsEnum = "uri"
)

var mappingListConnectionsFields = map[string]ListConnectionsFieldsEnum{
	"key":            ListConnectionsFieldsKey,
	"displayName":    ListConnectionsFieldsDisplayname,
	"description":    ListConnectionsFieldsDescription,
	"dataAssetKey":   ListConnectionsFieldsDataassetkey,
	"typeKey":        ListConnectionsFieldsTypekey,
	"timeCreated":    ListConnectionsFieldsTimecreated,
	"externalKey":    ListConnectionsFieldsExternalkey,
	"lifecycleState": ListConnectionsFieldsLifecyclestate,
	"isDefault":      ListConnectionsFieldsIsdefault,
	"uri":            ListConnectionsFieldsUri,
}

// GetListConnectionsFieldsEnumValues Enumerates the set of values for ListConnectionsFieldsEnum
func GetListConnectionsFieldsEnumValues() []ListConnectionsFieldsEnum {
	values := make([]ListConnectionsFieldsEnum, 0)
	for _, v := range mappingListConnectionsFields {
		values = append(values, v)
	}
	return values
}

// ListConnectionsSortByEnum Enum with underlying type: string
type ListConnectionsSortByEnum string

// Set of constants representing the allowable values for ListConnectionsSortByEnum
const (
	ListConnectionsSortByTimecreated ListConnectionsSortByEnum = "TIMECREATED"
	ListConnectionsSortByDisplayname ListConnectionsSortByEnum = "DISPLAYNAME"
)

var mappingListConnectionsSortBy = map[string]ListConnectionsSortByEnum{
	"TIMECREATED": ListConnectionsSortByTimecreated,
	"DISPLAYNAME": ListConnectionsSortByDisplayname,
}

// GetListConnectionsSortByEnumValues Enumerates the set of values for ListConnectionsSortByEnum
func GetListConnectionsSortByEnumValues() []ListConnectionsSortByEnum {
	values := make([]ListConnectionsSortByEnum, 0)
	for _, v := range mappingListConnectionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListConnectionsSortOrderEnum Enum with underlying type: string
type ListConnectionsSortOrderEnum string

// Set of constants representing the allowable values for ListConnectionsSortOrderEnum
const (
	ListConnectionsSortOrderAsc  ListConnectionsSortOrderEnum = "ASC"
	ListConnectionsSortOrderDesc ListConnectionsSortOrderEnum = "DESC"
)

var mappingListConnectionsSortOrder = map[string]ListConnectionsSortOrderEnum{
	"ASC":  ListConnectionsSortOrderAsc,
	"DESC": ListConnectionsSortOrderDesc,
}

// GetListConnectionsSortOrderEnumValues Enumerates the set of values for ListConnectionsSortOrderEnum
func GetListConnectionsSortOrderEnumValues() []ListConnectionsSortOrderEnum {
	values := make([]ListConnectionsSortOrderEnum, 0)
	for _, v := range mappingListConnectionsSortOrder {
		values = append(values, v)
	}
	return values
}
