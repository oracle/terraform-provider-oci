// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bastion

import (
	"github.com/oracle/oci-go-sdk/v46/common"
	"net/http"
)

// ListSessionsRequest wrapper for the ListSessions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bastion/ListSessions.go.html to see an example of how to use ListSessionsRequest.
type ListSessionsRequest struct {

	// The unique identifier (OCID) of the bastion in which to list sessions.
	BastionId *string `mandatory:"true" contributesTo:"query" name:"bastionId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	SessionLifecycleState ListSessionsSessionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"sessionLifecycleState" omitEmpty:"true"`

	// The unique identifier (OCID) of the session in which to list resources.
	SessionId *string `mandatory:"false" contributesTo:"query" name:"sessionId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListSessionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListSessionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSessionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSessionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSessionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSessionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSessionsResponse wrapper for the ListSessions operation
type ListSessionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SessionSummary instances
	Items []SessionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSessionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSessionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSessionsSessionLifecycleStateEnum Enum with underlying type: string
type ListSessionsSessionLifecycleStateEnum string

// Set of constants representing the allowable values for ListSessionsSessionLifecycleStateEnum
const (
	ListSessionsSessionLifecycleStateCreating ListSessionsSessionLifecycleStateEnum = "CREATING"
	ListSessionsSessionLifecycleStateActive   ListSessionsSessionLifecycleStateEnum = "ACTIVE"
	ListSessionsSessionLifecycleStateDeleting ListSessionsSessionLifecycleStateEnum = "DELETING"
	ListSessionsSessionLifecycleStateDeleted  ListSessionsSessionLifecycleStateEnum = "DELETED"
	ListSessionsSessionLifecycleStateFailed   ListSessionsSessionLifecycleStateEnum = "FAILED"
)

var mappingListSessionsSessionLifecycleState = map[string]ListSessionsSessionLifecycleStateEnum{
	"CREATING": ListSessionsSessionLifecycleStateCreating,
	"ACTIVE":   ListSessionsSessionLifecycleStateActive,
	"DELETING": ListSessionsSessionLifecycleStateDeleting,
	"DELETED":  ListSessionsSessionLifecycleStateDeleted,
	"FAILED":   ListSessionsSessionLifecycleStateFailed,
}

// GetListSessionsSessionLifecycleStateEnumValues Enumerates the set of values for ListSessionsSessionLifecycleStateEnum
func GetListSessionsSessionLifecycleStateEnumValues() []ListSessionsSessionLifecycleStateEnum {
	values := make([]ListSessionsSessionLifecycleStateEnum, 0)
	for _, v := range mappingListSessionsSessionLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListSessionsSortOrderEnum Enum with underlying type: string
type ListSessionsSortOrderEnum string

// Set of constants representing the allowable values for ListSessionsSortOrderEnum
const (
	ListSessionsSortOrderAsc  ListSessionsSortOrderEnum = "ASC"
	ListSessionsSortOrderDesc ListSessionsSortOrderEnum = "DESC"
)

var mappingListSessionsSortOrder = map[string]ListSessionsSortOrderEnum{
	"ASC":  ListSessionsSortOrderAsc,
	"DESC": ListSessionsSortOrderDesc,
}

// GetListSessionsSortOrderEnumValues Enumerates the set of values for ListSessionsSortOrderEnum
func GetListSessionsSortOrderEnumValues() []ListSessionsSortOrderEnum {
	values := make([]ListSessionsSortOrderEnum, 0)
	for _, v := range mappingListSessionsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListSessionsSortByEnum Enum with underlying type: string
type ListSessionsSortByEnum string

// Set of constants representing the allowable values for ListSessionsSortByEnum
const (
	ListSessionsSortByTimecreated ListSessionsSortByEnum = "timeCreated"
	ListSessionsSortByDisplayname ListSessionsSortByEnum = "displayName"
)

var mappingListSessionsSortBy = map[string]ListSessionsSortByEnum{
	"timeCreated": ListSessionsSortByTimecreated,
	"displayName": ListSessionsSortByDisplayname,
}

// GetListSessionsSortByEnumValues Enumerates the set of values for ListSessionsSortByEnum
func GetListSessionsSortByEnumValues() []ListSessionsSortByEnum {
	values := make([]ListSessionsSortByEnum, 0)
	for _, v := range mappingListSessionsSortBy {
		values = append(values, v)
	}
	return values
}
