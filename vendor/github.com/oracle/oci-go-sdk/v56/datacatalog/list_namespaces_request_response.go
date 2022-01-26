// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListNamespacesRequest wrapper for the ListNamespaces operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListNamespaces.go.html to see an example of how to use ListNamespacesRequest.
type ListNamespacesRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListNamespacesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListNamespacesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListNamespacesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the fields to return in a namespace summary response.
	Fields []ListNamespacesFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

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

func (request ListNamespacesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNamespacesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNamespacesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNamespacesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListNamespacesResponse wrapper for the ListNamespaces operation
type ListNamespacesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NamespaceCollection instances
	NamespaceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNamespacesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNamespacesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNamespacesLifecycleStateEnum Enum with underlying type: string
type ListNamespacesLifecycleStateEnum string

// Set of constants representing the allowable values for ListNamespacesLifecycleStateEnum
const (
	ListNamespacesLifecycleStateCreating ListNamespacesLifecycleStateEnum = "CREATING"
	ListNamespacesLifecycleStateActive   ListNamespacesLifecycleStateEnum = "ACTIVE"
	ListNamespacesLifecycleStateInactive ListNamespacesLifecycleStateEnum = "INACTIVE"
	ListNamespacesLifecycleStateUpdating ListNamespacesLifecycleStateEnum = "UPDATING"
	ListNamespacesLifecycleStateDeleting ListNamespacesLifecycleStateEnum = "DELETING"
	ListNamespacesLifecycleStateDeleted  ListNamespacesLifecycleStateEnum = "DELETED"
	ListNamespacesLifecycleStateFailed   ListNamespacesLifecycleStateEnum = "FAILED"
	ListNamespacesLifecycleStateMoving   ListNamespacesLifecycleStateEnum = "MOVING"
)

var mappingListNamespacesLifecycleState = map[string]ListNamespacesLifecycleStateEnum{
	"CREATING": ListNamespacesLifecycleStateCreating,
	"ACTIVE":   ListNamespacesLifecycleStateActive,
	"INACTIVE": ListNamespacesLifecycleStateInactive,
	"UPDATING": ListNamespacesLifecycleStateUpdating,
	"DELETING": ListNamespacesLifecycleStateDeleting,
	"DELETED":  ListNamespacesLifecycleStateDeleted,
	"FAILED":   ListNamespacesLifecycleStateFailed,
	"MOVING":   ListNamespacesLifecycleStateMoving,
}

// GetListNamespacesLifecycleStateEnumValues Enumerates the set of values for ListNamespacesLifecycleStateEnum
func GetListNamespacesLifecycleStateEnumValues() []ListNamespacesLifecycleStateEnum {
	values := make([]ListNamespacesLifecycleStateEnum, 0)
	for _, v := range mappingListNamespacesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListNamespacesSortByEnum Enum with underlying type: string
type ListNamespacesSortByEnum string

// Set of constants representing the allowable values for ListNamespacesSortByEnum
const (
	ListNamespacesSortByTimecreated ListNamespacesSortByEnum = "TIMECREATED"
	ListNamespacesSortByDisplayname ListNamespacesSortByEnum = "DISPLAYNAME"
)

var mappingListNamespacesSortBy = map[string]ListNamespacesSortByEnum{
	"TIMECREATED": ListNamespacesSortByTimecreated,
	"DISPLAYNAME": ListNamespacesSortByDisplayname,
}

// GetListNamespacesSortByEnumValues Enumerates the set of values for ListNamespacesSortByEnum
func GetListNamespacesSortByEnumValues() []ListNamespacesSortByEnum {
	values := make([]ListNamespacesSortByEnum, 0)
	for _, v := range mappingListNamespacesSortBy {
		values = append(values, v)
	}
	return values
}

// ListNamespacesSortOrderEnum Enum with underlying type: string
type ListNamespacesSortOrderEnum string

// Set of constants representing the allowable values for ListNamespacesSortOrderEnum
const (
	ListNamespacesSortOrderAsc  ListNamespacesSortOrderEnum = "ASC"
	ListNamespacesSortOrderDesc ListNamespacesSortOrderEnum = "DESC"
)

var mappingListNamespacesSortOrder = map[string]ListNamespacesSortOrderEnum{
	"ASC":  ListNamespacesSortOrderAsc,
	"DESC": ListNamespacesSortOrderDesc,
}

// GetListNamespacesSortOrderEnumValues Enumerates the set of values for ListNamespacesSortOrderEnum
func GetListNamespacesSortOrderEnumValues() []ListNamespacesSortOrderEnum {
	values := make([]ListNamespacesSortOrderEnum, 0)
	for _, v := range mappingListNamespacesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListNamespacesFieldsEnum Enum with underlying type: string
type ListNamespacesFieldsEnum string

// Set of constants representing the allowable values for ListNamespacesFieldsEnum
const (
	ListNamespacesFieldsKey            ListNamespacesFieldsEnum = "key"
	ListNamespacesFieldsDisplayname    ListNamespacesFieldsEnum = "displayName"
	ListNamespacesFieldsDescription    ListNamespacesFieldsEnum = "description"
	ListNamespacesFieldsLifecyclestate ListNamespacesFieldsEnum = "lifecycleState"
	ListNamespacesFieldsTimecreated    ListNamespacesFieldsEnum = "timeCreated"
)

var mappingListNamespacesFields = map[string]ListNamespacesFieldsEnum{
	"key":            ListNamespacesFieldsKey,
	"displayName":    ListNamespacesFieldsDisplayname,
	"description":    ListNamespacesFieldsDescription,
	"lifecycleState": ListNamespacesFieldsLifecyclestate,
	"timeCreated":    ListNamespacesFieldsTimecreated,
}

// GetListNamespacesFieldsEnumValues Enumerates the set of values for ListNamespacesFieldsEnum
func GetListNamespacesFieldsEnumValues() []ListNamespacesFieldsEnum {
	values := make([]ListNamespacesFieldsEnum, 0)
	for _, v := range mappingListNamespacesFields {
		values = append(values, v)
	}
	return values
}
