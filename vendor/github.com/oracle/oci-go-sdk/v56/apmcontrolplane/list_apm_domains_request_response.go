// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmcontrolplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListApmDomainsRequest wrapper for the ListApmDomains operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmcontrolplane/ListApmDomains.go.html to see an example of how to use ListApmDomainsRequest.
type ListApmDomainsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given life-cycle state.
	LifecycleState ListApmDomainsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListApmDomainsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListApmDomainsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApmDomainsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApmDomainsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApmDomainsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApmDomainsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListApmDomainsResponse wrapper for the ListApmDomains operation
type ListApmDomainsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ApmDomainSummary instances
	Items []ApmDomainSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListApmDomainsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApmDomainsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApmDomainsLifecycleStateEnum Enum with underlying type: string
type ListApmDomainsLifecycleStateEnum string

// Set of constants representing the allowable values for ListApmDomainsLifecycleStateEnum
const (
	ListApmDomainsLifecycleStateCreating ListApmDomainsLifecycleStateEnum = "CREATING"
	ListApmDomainsLifecycleStateUpdating ListApmDomainsLifecycleStateEnum = "UPDATING"
	ListApmDomainsLifecycleStateActive   ListApmDomainsLifecycleStateEnum = "ACTIVE"
	ListApmDomainsLifecycleStateDeleting ListApmDomainsLifecycleStateEnum = "DELETING"
	ListApmDomainsLifecycleStateDeleted  ListApmDomainsLifecycleStateEnum = "DELETED"
	ListApmDomainsLifecycleStateFailed   ListApmDomainsLifecycleStateEnum = "FAILED"
)

var mappingListApmDomainsLifecycleState = map[string]ListApmDomainsLifecycleStateEnum{
	"CREATING": ListApmDomainsLifecycleStateCreating,
	"UPDATING": ListApmDomainsLifecycleStateUpdating,
	"ACTIVE":   ListApmDomainsLifecycleStateActive,
	"DELETING": ListApmDomainsLifecycleStateDeleting,
	"DELETED":  ListApmDomainsLifecycleStateDeleted,
	"FAILED":   ListApmDomainsLifecycleStateFailed,
}

// GetListApmDomainsLifecycleStateEnumValues Enumerates the set of values for ListApmDomainsLifecycleStateEnum
func GetListApmDomainsLifecycleStateEnumValues() []ListApmDomainsLifecycleStateEnum {
	values := make([]ListApmDomainsLifecycleStateEnum, 0)
	for _, v := range mappingListApmDomainsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListApmDomainsSortOrderEnum Enum with underlying type: string
type ListApmDomainsSortOrderEnum string

// Set of constants representing the allowable values for ListApmDomainsSortOrderEnum
const (
	ListApmDomainsSortOrderAsc  ListApmDomainsSortOrderEnum = "ASC"
	ListApmDomainsSortOrderDesc ListApmDomainsSortOrderEnum = "DESC"
)

var mappingListApmDomainsSortOrder = map[string]ListApmDomainsSortOrderEnum{
	"ASC":  ListApmDomainsSortOrderAsc,
	"DESC": ListApmDomainsSortOrderDesc,
}

// GetListApmDomainsSortOrderEnumValues Enumerates the set of values for ListApmDomainsSortOrderEnum
func GetListApmDomainsSortOrderEnumValues() []ListApmDomainsSortOrderEnum {
	values := make([]ListApmDomainsSortOrderEnum, 0)
	for _, v := range mappingListApmDomainsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListApmDomainsSortByEnum Enum with underlying type: string
type ListApmDomainsSortByEnum string

// Set of constants representing the allowable values for ListApmDomainsSortByEnum
const (
	ListApmDomainsSortByTimecreated ListApmDomainsSortByEnum = "timeCreated"
	ListApmDomainsSortByDisplayname ListApmDomainsSortByEnum = "displayName"
)

var mappingListApmDomainsSortBy = map[string]ListApmDomainsSortByEnum{
	"timeCreated": ListApmDomainsSortByTimecreated,
	"displayName": ListApmDomainsSortByDisplayname,
}

// GetListApmDomainsSortByEnumValues Enumerates the set of values for ListApmDomainsSortByEnum
func GetListApmDomainsSortByEnumValues() []ListApmDomainsSortByEnum {
	values := make([]ListApmDomainsSortByEnum, 0)
	for _, v := range mappingListApmDomainsSortBy {
		values = append(values, v)
	}
	return values
}
