// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// SearchCriteriaRequest wrapper for the SearchCriteria operation
type SearchCriteriaRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// The information used to create an extended search results.
	SearchCriteriaDetails SearchCriteria `contributesTo:"body"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Immutable resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState SearchCriteriaLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A search timeout string (for example, timeout=4000ms), bounding the search request to be executed within the
	// specified time value and bail with the hits accumulated up to that point when expired.
	// Defaults to no timeout.
	Timeout *string `mandatory:"false" contributesTo:"query" name:"timeout"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy SearchCriteriaSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SearchCriteriaSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request SearchCriteriaRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SearchCriteriaRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SearchCriteriaRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SearchCriteriaResponse wrapper for the SearchCriteria operation
type SearchCriteriaResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SearchResultCollection instances
	SearchResultCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SearchCriteriaResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SearchCriteriaResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SearchCriteriaLifecycleStateEnum Enum with underlying type: string
type SearchCriteriaLifecycleStateEnum string

// Set of constants representing the allowable values for SearchCriteriaLifecycleStateEnum
const (
	SearchCriteriaLifecycleStateCreating SearchCriteriaLifecycleStateEnum = "CREATING"
	SearchCriteriaLifecycleStateActive   SearchCriteriaLifecycleStateEnum = "ACTIVE"
	SearchCriteriaLifecycleStateInactive SearchCriteriaLifecycleStateEnum = "INACTIVE"
	SearchCriteriaLifecycleStateUpdating SearchCriteriaLifecycleStateEnum = "UPDATING"
	SearchCriteriaLifecycleStateDeleting SearchCriteriaLifecycleStateEnum = "DELETING"
	SearchCriteriaLifecycleStateDeleted  SearchCriteriaLifecycleStateEnum = "DELETED"
	SearchCriteriaLifecycleStateFailed   SearchCriteriaLifecycleStateEnum = "FAILED"
	SearchCriteriaLifecycleStateMoving   SearchCriteriaLifecycleStateEnum = "MOVING"
)

var mappingSearchCriteriaLifecycleState = map[string]SearchCriteriaLifecycleStateEnum{
	"CREATING": SearchCriteriaLifecycleStateCreating,
	"ACTIVE":   SearchCriteriaLifecycleStateActive,
	"INACTIVE": SearchCriteriaLifecycleStateInactive,
	"UPDATING": SearchCriteriaLifecycleStateUpdating,
	"DELETING": SearchCriteriaLifecycleStateDeleting,
	"DELETED":  SearchCriteriaLifecycleStateDeleted,
	"FAILED":   SearchCriteriaLifecycleStateFailed,
	"MOVING":   SearchCriteriaLifecycleStateMoving,
}

// GetSearchCriteriaLifecycleStateEnumValues Enumerates the set of values for SearchCriteriaLifecycleStateEnum
func GetSearchCriteriaLifecycleStateEnumValues() []SearchCriteriaLifecycleStateEnum {
	values := make([]SearchCriteriaLifecycleStateEnum, 0)
	for _, v := range mappingSearchCriteriaLifecycleState {
		values = append(values, v)
	}
	return values
}

// SearchCriteriaSortByEnum Enum with underlying type: string
type SearchCriteriaSortByEnum string

// Set of constants representing the allowable values for SearchCriteriaSortByEnum
const (
	SearchCriteriaSortByTimecreated SearchCriteriaSortByEnum = "TIMECREATED"
	SearchCriteriaSortByDisplayname SearchCriteriaSortByEnum = "DISPLAYNAME"
)

var mappingSearchCriteriaSortBy = map[string]SearchCriteriaSortByEnum{
	"TIMECREATED": SearchCriteriaSortByTimecreated,
	"DISPLAYNAME": SearchCriteriaSortByDisplayname,
}

// GetSearchCriteriaSortByEnumValues Enumerates the set of values for SearchCriteriaSortByEnum
func GetSearchCriteriaSortByEnumValues() []SearchCriteriaSortByEnum {
	values := make([]SearchCriteriaSortByEnum, 0)
	for _, v := range mappingSearchCriteriaSortBy {
		values = append(values, v)
	}
	return values
}

// SearchCriteriaSortOrderEnum Enum with underlying type: string
type SearchCriteriaSortOrderEnum string

// Set of constants representing the allowable values for SearchCriteriaSortOrderEnum
const (
	SearchCriteriaSortOrderAsc  SearchCriteriaSortOrderEnum = "ASC"
	SearchCriteriaSortOrderDesc SearchCriteriaSortOrderEnum = "DESC"
)

var mappingSearchCriteriaSortOrder = map[string]SearchCriteriaSortOrderEnum{
	"ASC":  SearchCriteriaSortOrderAsc,
	"DESC": SearchCriteriaSortOrderDesc,
}

// GetSearchCriteriaSortOrderEnumValues Enumerates the set of values for SearchCriteriaSortOrderEnum
func GetSearchCriteriaSortOrderEnumValues() []SearchCriteriaSortOrderEnum {
	values := make([]SearchCriteriaSortOrderEnum, 0)
	for _, v := range mappingSearchCriteriaSortOrder {
		values = append(values, v)
	}
	return values
}
