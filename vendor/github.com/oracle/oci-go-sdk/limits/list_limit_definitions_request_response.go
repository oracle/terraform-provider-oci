// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package limits

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListLimitDefinitionsRequest wrapper for the ListLimitDefinitions operation
type ListLimitDefinitionsRequest struct {

	// The OCID of the parent compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The target service name.
	ServiceName *string `mandatory:"false" contributesTo:"query" name:"serviceName"`

	// Optional field, filter for a specific resource limit.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by.
	SortBy ListLimitDefinitionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'. By default it will be ascending.
	SortOrder ListLimitDefinitionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLimitDefinitionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLimitDefinitionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLimitDefinitionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLimitDefinitionsResponse wrapper for the ListLimitDefinitions operation
type ListLimitDefinitionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []LimitDefinitionSummary instances
	Items []LimitDefinitionSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLimitDefinitionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLimitDefinitionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLimitDefinitionsSortByEnum Enum with underlying type: string
type ListLimitDefinitionsSortByEnum string

// Set of constants representing the allowable values for ListLimitDefinitionsSortByEnum
const (
	ListLimitDefinitionsSortByName        ListLimitDefinitionsSortByEnum = "name"
	ListLimitDefinitionsSortByDescription ListLimitDefinitionsSortByEnum = "description"
)

var mappingListLimitDefinitionsSortBy = map[string]ListLimitDefinitionsSortByEnum{
	"name":        ListLimitDefinitionsSortByName,
	"description": ListLimitDefinitionsSortByDescription,
}

// GetListLimitDefinitionsSortByEnumValues Enumerates the set of values for ListLimitDefinitionsSortByEnum
func GetListLimitDefinitionsSortByEnumValues() []ListLimitDefinitionsSortByEnum {
	values := make([]ListLimitDefinitionsSortByEnum, 0)
	for _, v := range mappingListLimitDefinitionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListLimitDefinitionsSortOrderEnum Enum with underlying type: string
type ListLimitDefinitionsSortOrderEnum string

// Set of constants representing the allowable values for ListLimitDefinitionsSortOrderEnum
const (
	ListLimitDefinitionsSortOrderAsc  ListLimitDefinitionsSortOrderEnum = "ASC"
	ListLimitDefinitionsSortOrderDesc ListLimitDefinitionsSortOrderEnum = "DESC"
)

var mappingListLimitDefinitionsSortOrder = map[string]ListLimitDefinitionsSortOrderEnum{
	"ASC":  ListLimitDefinitionsSortOrderAsc,
	"DESC": ListLimitDefinitionsSortOrderDesc,
}

// GetListLimitDefinitionsSortOrderEnumValues Enumerates the set of values for ListLimitDefinitionsSortOrderEnum
func GetListLimitDefinitionsSortOrderEnumValues() []ListLimitDefinitionsSortOrderEnum {
	values := make([]ListLimitDefinitionsSortOrderEnum, 0)
	for _, v := range mappingListLimitDefinitionsSortOrder {
		values = append(values, v)
	}
	return values
}
