// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListHttpRedirectsRequest wrapper for the ListHttpRedirects operation
type ListHttpRedirectsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This number is generated when the compartment is created.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated call. If unspecified, defaults to `10`.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous paginated call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The value of the sorting direction of resources in a paginated 'List' call. If unspecified, defaults to `DESC`.
	SortOrder ListHttpRedirectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the results of the List query.
	SortBy ListHttpRedirectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter redirects using a list of redirect OCIDs.
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Filter redirects using a display name.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// Filter redirects using a list of lifecycle states.
	LifecycleState []ListHttpRedirectsLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter that matches redirects created on or after the specified date and time.
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// A filter that matches redirects created before the specified date-time. Default to 1 day before now.
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHttpRedirectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHttpRedirectsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHttpRedirectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListHttpRedirectsResponse wrapper for the ListHttpRedirects operation
type ListHttpRedirectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []HttpRedirectSummary instances
	Items []HttpRedirectSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results may remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListHttpRedirectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHttpRedirectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHttpRedirectsSortOrderEnum Enum with underlying type: string
type ListHttpRedirectsSortOrderEnum string

// Set of constants representing the allowable values for ListHttpRedirectsSortOrderEnum
const (
	ListHttpRedirectsSortOrderAsc  ListHttpRedirectsSortOrderEnum = "ASC"
	ListHttpRedirectsSortOrderDesc ListHttpRedirectsSortOrderEnum = "DESC"
)

var mappingListHttpRedirectsSortOrder = map[string]ListHttpRedirectsSortOrderEnum{
	"ASC":  ListHttpRedirectsSortOrderAsc,
	"DESC": ListHttpRedirectsSortOrderDesc,
}

// GetListHttpRedirectsSortOrderEnumValues Enumerates the set of values for ListHttpRedirectsSortOrderEnum
func GetListHttpRedirectsSortOrderEnumValues() []ListHttpRedirectsSortOrderEnum {
	values := make([]ListHttpRedirectsSortOrderEnum, 0)
	for _, v := range mappingListHttpRedirectsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListHttpRedirectsSortByEnum Enum with underlying type: string
type ListHttpRedirectsSortByEnum string

// Set of constants representing the allowable values for ListHttpRedirectsSortByEnum
const (
	ListHttpRedirectsSortById          ListHttpRedirectsSortByEnum = "id"
	ListHttpRedirectsSortByDomain      ListHttpRedirectsSortByEnum = "domain"
	ListHttpRedirectsSortByTarget      ListHttpRedirectsSortByEnum = "target"
	ListHttpRedirectsSortByDisplayname ListHttpRedirectsSortByEnum = "displayName"
)

var mappingListHttpRedirectsSortBy = map[string]ListHttpRedirectsSortByEnum{
	"id":          ListHttpRedirectsSortById,
	"domain":      ListHttpRedirectsSortByDomain,
	"target":      ListHttpRedirectsSortByTarget,
	"displayName": ListHttpRedirectsSortByDisplayname,
}

// GetListHttpRedirectsSortByEnumValues Enumerates the set of values for ListHttpRedirectsSortByEnum
func GetListHttpRedirectsSortByEnumValues() []ListHttpRedirectsSortByEnum {
	values := make([]ListHttpRedirectsSortByEnum, 0)
	for _, v := range mappingListHttpRedirectsSortBy {
		values = append(values, v)
	}
	return values
}

// ListHttpRedirectsLifecycleStateEnum Enum with underlying type: string
type ListHttpRedirectsLifecycleStateEnum string

// Set of constants representing the allowable values for ListHttpRedirectsLifecycleStateEnum
const (
	ListHttpRedirectsLifecycleStateCreating ListHttpRedirectsLifecycleStateEnum = "CREATING"
	ListHttpRedirectsLifecycleStateActive   ListHttpRedirectsLifecycleStateEnum = "ACTIVE"
	ListHttpRedirectsLifecycleStateFailed   ListHttpRedirectsLifecycleStateEnum = "FAILED"
	ListHttpRedirectsLifecycleStateUpdating ListHttpRedirectsLifecycleStateEnum = "UPDATING"
	ListHttpRedirectsLifecycleStateDeleting ListHttpRedirectsLifecycleStateEnum = "DELETING"
	ListHttpRedirectsLifecycleStateDeleted  ListHttpRedirectsLifecycleStateEnum = "DELETED"
)

var mappingListHttpRedirectsLifecycleState = map[string]ListHttpRedirectsLifecycleStateEnum{
	"CREATING": ListHttpRedirectsLifecycleStateCreating,
	"ACTIVE":   ListHttpRedirectsLifecycleStateActive,
	"FAILED":   ListHttpRedirectsLifecycleStateFailed,
	"UPDATING": ListHttpRedirectsLifecycleStateUpdating,
	"DELETING": ListHttpRedirectsLifecycleStateDeleting,
	"DELETED":  ListHttpRedirectsLifecycleStateDeleted,
}

// GetListHttpRedirectsLifecycleStateEnumValues Enumerates the set of values for ListHttpRedirectsLifecycleStateEnum
func GetListHttpRedirectsLifecycleStateEnumValues() []ListHttpRedirectsLifecycleStateEnum {
	values := make([]ListHttpRedirectsLifecycleStateEnum, 0)
	for _, v := range mappingListHttpRedirectsLifecycleState {
		values = append(values, v)
	}
	return values
}
