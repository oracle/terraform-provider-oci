// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListManagementAgentImagesRequest wrapper for the ListManagementAgentImages operation
type ListManagementAgentImagesRequest struct {

	// The ID of the compartment from which the Management Agents to be listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagementAgentImagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for platformType is descending. Default order for version is descending. If no value is specified platformType is default.
	SortBy ListManagementAgentImagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire platform name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Filter to return only Management Agents in the particular lifecycle state.
	LifecycleState ListManagementAgentImagesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementAgentImagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementAgentImagesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementAgentImagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListManagementAgentImagesResponse wrapper for the ListManagementAgentImages operation
type ListManagementAgentImagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ManagementAgentImageSummary instances
	Items []ManagementAgentImageSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagementAgentImagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementAgentImagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementAgentImagesSortOrderEnum Enum with underlying type: string
type ListManagementAgentImagesSortOrderEnum string

// Set of constants representing the allowable values for ListManagementAgentImagesSortOrderEnum
const (
	ListManagementAgentImagesSortOrderAsc  ListManagementAgentImagesSortOrderEnum = "ASC"
	ListManagementAgentImagesSortOrderDesc ListManagementAgentImagesSortOrderEnum = "DESC"
)

var mappingListManagementAgentImagesSortOrder = map[string]ListManagementAgentImagesSortOrderEnum{
	"ASC":  ListManagementAgentImagesSortOrderAsc,
	"DESC": ListManagementAgentImagesSortOrderDesc,
}

// GetListManagementAgentImagesSortOrderEnumValues Enumerates the set of values for ListManagementAgentImagesSortOrderEnum
func GetListManagementAgentImagesSortOrderEnumValues() []ListManagementAgentImagesSortOrderEnum {
	values := make([]ListManagementAgentImagesSortOrderEnum, 0)
	for _, v := range mappingListManagementAgentImagesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentImagesSortByEnum Enum with underlying type: string
type ListManagementAgentImagesSortByEnum string

// Set of constants representing the allowable values for ListManagementAgentImagesSortByEnum
const (
	ListManagementAgentImagesSortByPlatformtype ListManagementAgentImagesSortByEnum = "platformType"
	ListManagementAgentImagesSortByVersion      ListManagementAgentImagesSortByEnum = "version"
)

var mappingListManagementAgentImagesSortBy = map[string]ListManagementAgentImagesSortByEnum{
	"platformType": ListManagementAgentImagesSortByPlatformtype,
	"version":      ListManagementAgentImagesSortByVersion,
}

// GetListManagementAgentImagesSortByEnumValues Enumerates the set of values for ListManagementAgentImagesSortByEnum
func GetListManagementAgentImagesSortByEnumValues() []ListManagementAgentImagesSortByEnum {
	values := make([]ListManagementAgentImagesSortByEnum, 0)
	for _, v := range mappingListManagementAgentImagesSortBy {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentImagesLifecycleStateEnum Enum with underlying type: string
type ListManagementAgentImagesLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagementAgentImagesLifecycleStateEnum
const (
	ListManagementAgentImagesLifecycleStateCreating   ListManagementAgentImagesLifecycleStateEnum = "CREATING"
	ListManagementAgentImagesLifecycleStateUpdating   ListManagementAgentImagesLifecycleStateEnum = "UPDATING"
	ListManagementAgentImagesLifecycleStateActive     ListManagementAgentImagesLifecycleStateEnum = "ACTIVE"
	ListManagementAgentImagesLifecycleStateInactive   ListManagementAgentImagesLifecycleStateEnum = "INACTIVE"
	ListManagementAgentImagesLifecycleStateTerminated ListManagementAgentImagesLifecycleStateEnum = "TERMINATED"
	ListManagementAgentImagesLifecycleStateDeleting   ListManagementAgentImagesLifecycleStateEnum = "DELETING"
	ListManagementAgentImagesLifecycleStateDeleted    ListManagementAgentImagesLifecycleStateEnum = "DELETED"
	ListManagementAgentImagesLifecycleStateFailed     ListManagementAgentImagesLifecycleStateEnum = "FAILED"
)

var mappingListManagementAgentImagesLifecycleState = map[string]ListManagementAgentImagesLifecycleStateEnum{
	"CREATING":   ListManagementAgentImagesLifecycleStateCreating,
	"UPDATING":   ListManagementAgentImagesLifecycleStateUpdating,
	"ACTIVE":     ListManagementAgentImagesLifecycleStateActive,
	"INACTIVE":   ListManagementAgentImagesLifecycleStateInactive,
	"TERMINATED": ListManagementAgentImagesLifecycleStateTerminated,
	"DELETING":   ListManagementAgentImagesLifecycleStateDeleting,
	"DELETED":    ListManagementAgentImagesLifecycleStateDeleted,
	"FAILED":     ListManagementAgentImagesLifecycleStateFailed,
}

// GetListManagementAgentImagesLifecycleStateEnumValues Enumerates the set of values for ListManagementAgentImagesLifecycleStateEnum
func GetListManagementAgentImagesLifecycleStateEnumValues() []ListManagementAgentImagesLifecycleStateEnum {
	values := make([]ListManagementAgentImagesLifecycleStateEnum, 0)
	for _, v := range mappingListManagementAgentImagesLifecycleState {
		values = append(values, v)
	}
	return values
}
