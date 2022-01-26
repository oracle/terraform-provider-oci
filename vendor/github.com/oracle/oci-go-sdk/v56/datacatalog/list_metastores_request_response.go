// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListMetastoresRequest wrapper for the ListMetastores operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListMetastores.go.html to see an example of how to use ListMetastoresRequest.
type ListMetastoresRequest struct {

	// The OCID of the compartment where you want to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListMetastoresLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListMetastoresSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListMetastoresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMetastoresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMetastoresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMetastoresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMetastoresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListMetastoresResponse wrapper for the ListMetastores operation
type ListMetastoresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []MetastoreSummary instances
	Items []MetastoreSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMetastoresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMetastoresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMetastoresLifecycleStateEnum Enum with underlying type: string
type ListMetastoresLifecycleStateEnum string

// Set of constants representing the allowable values for ListMetastoresLifecycleStateEnum
const (
	ListMetastoresLifecycleStateCreating ListMetastoresLifecycleStateEnum = "CREATING"
	ListMetastoresLifecycleStateActive   ListMetastoresLifecycleStateEnum = "ACTIVE"
	ListMetastoresLifecycleStateInactive ListMetastoresLifecycleStateEnum = "INACTIVE"
	ListMetastoresLifecycleStateUpdating ListMetastoresLifecycleStateEnum = "UPDATING"
	ListMetastoresLifecycleStateDeleting ListMetastoresLifecycleStateEnum = "DELETING"
	ListMetastoresLifecycleStateDeleted  ListMetastoresLifecycleStateEnum = "DELETED"
	ListMetastoresLifecycleStateFailed   ListMetastoresLifecycleStateEnum = "FAILED"
	ListMetastoresLifecycleStateMoving   ListMetastoresLifecycleStateEnum = "MOVING"
)

var mappingListMetastoresLifecycleState = map[string]ListMetastoresLifecycleStateEnum{
	"CREATING": ListMetastoresLifecycleStateCreating,
	"ACTIVE":   ListMetastoresLifecycleStateActive,
	"INACTIVE": ListMetastoresLifecycleStateInactive,
	"UPDATING": ListMetastoresLifecycleStateUpdating,
	"DELETING": ListMetastoresLifecycleStateDeleting,
	"DELETED":  ListMetastoresLifecycleStateDeleted,
	"FAILED":   ListMetastoresLifecycleStateFailed,
	"MOVING":   ListMetastoresLifecycleStateMoving,
}

// GetListMetastoresLifecycleStateEnumValues Enumerates the set of values for ListMetastoresLifecycleStateEnum
func GetListMetastoresLifecycleStateEnumValues() []ListMetastoresLifecycleStateEnum {
	values := make([]ListMetastoresLifecycleStateEnum, 0)
	for _, v := range mappingListMetastoresLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListMetastoresSortOrderEnum Enum with underlying type: string
type ListMetastoresSortOrderEnum string

// Set of constants representing the allowable values for ListMetastoresSortOrderEnum
const (
	ListMetastoresSortOrderAsc  ListMetastoresSortOrderEnum = "ASC"
	ListMetastoresSortOrderDesc ListMetastoresSortOrderEnum = "DESC"
)

var mappingListMetastoresSortOrder = map[string]ListMetastoresSortOrderEnum{
	"ASC":  ListMetastoresSortOrderAsc,
	"DESC": ListMetastoresSortOrderDesc,
}

// GetListMetastoresSortOrderEnumValues Enumerates the set of values for ListMetastoresSortOrderEnum
func GetListMetastoresSortOrderEnumValues() []ListMetastoresSortOrderEnum {
	values := make([]ListMetastoresSortOrderEnum, 0)
	for _, v := range mappingListMetastoresSortOrder {
		values = append(values, v)
	}
	return values
}

// ListMetastoresSortByEnum Enum with underlying type: string
type ListMetastoresSortByEnum string

// Set of constants representing the allowable values for ListMetastoresSortByEnum
const (
	ListMetastoresSortByTimecreated ListMetastoresSortByEnum = "TIMECREATED"
	ListMetastoresSortByDisplayname ListMetastoresSortByEnum = "DISPLAYNAME"
)

var mappingListMetastoresSortBy = map[string]ListMetastoresSortByEnum{
	"TIMECREATED": ListMetastoresSortByTimecreated,
	"DISPLAYNAME": ListMetastoresSortByDisplayname,
}

// GetListMetastoresSortByEnumValues Enumerates the set of values for ListMetastoresSortByEnum
func GetListMetastoresSortByEnumValues() []ListMetastoresSortByEnum {
	values := make([]ListMetastoresSortByEnum, 0)
	for _, v := range mappingListMetastoresSortBy {
		values = append(values, v)
	}
	return values
}
