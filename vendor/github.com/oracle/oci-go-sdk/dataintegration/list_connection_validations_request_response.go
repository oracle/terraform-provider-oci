// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListConnectionValidationsRequest wrapper for the ListConnectionValidations operation
type ListConnectionValidationsRequest struct {

	// DIS workspace id
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// This filter parameter can be used to filter by the key of the object.
	Key *string `mandatory:"false" contributesTo:"query" name:"key"`

	// This filter parameter can be used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// This filter parameter can be used to filter by the identifier of the object.
	Identifier *string `mandatory:"false" contributesTo:"query" name:"identifier"`

	// This parameter allows users to specify which fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// This parameter will control pagination.  Values for the parameter should come from the `opc-next-page` or `opc-prev-page` header in previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// This parameter allows users to set the maximum number of items to return per page.  The value must be between 1 and 100 (inclusive).  Default value is 100.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// This parameter allows users to specify a sort field.  Supported sort fields are `name`, `identifier`, `timeCreated`, and `timeUpdated`.  Default sort order is the descending order of `timeCreated` (most recently created objects at the top).  Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListConnectionValidationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// This parameter is used to control the sort order.  Supported values are `ASC` (ascending) and `DESC` (descending).
	SortOrder ListConnectionValidationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConnectionValidationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConnectionValidationsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConnectionValidationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListConnectionValidationsResponse wrapper for the ListConnectionValidations operation
type ListConnectionValidationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConnectionValidationSummaryCollection instances
	ConnectionValidationSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListConnectionValidationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConnectionValidationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConnectionValidationsSortByEnum Enum with underlying type: string
type ListConnectionValidationsSortByEnum string

// Set of constants representing the allowable values for ListConnectionValidationsSortByEnum
const (
	ListConnectionValidationsSortByTimeCreated ListConnectionValidationsSortByEnum = "TIME_CREATED"
	ListConnectionValidationsSortByDisplayName ListConnectionValidationsSortByEnum = "DISPLAY_NAME"
)

var mappingListConnectionValidationsSortBy = map[string]ListConnectionValidationsSortByEnum{
	"TIME_CREATED": ListConnectionValidationsSortByTimeCreated,
	"DISPLAY_NAME": ListConnectionValidationsSortByDisplayName,
}

// GetListConnectionValidationsSortByEnumValues Enumerates the set of values for ListConnectionValidationsSortByEnum
func GetListConnectionValidationsSortByEnumValues() []ListConnectionValidationsSortByEnum {
	values := make([]ListConnectionValidationsSortByEnum, 0)
	for _, v := range mappingListConnectionValidationsSortBy {
		values = append(values, v)
	}
	return values
}

// ListConnectionValidationsSortOrderEnum Enum with underlying type: string
type ListConnectionValidationsSortOrderEnum string

// Set of constants representing the allowable values for ListConnectionValidationsSortOrderEnum
const (
	ListConnectionValidationsSortOrderAsc  ListConnectionValidationsSortOrderEnum = "ASC"
	ListConnectionValidationsSortOrderDesc ListConnectionValidationsSortOrderEnum = "DESC"
)

var mappingListConnectionValidationsSortOrder = map[string]ListConnectionValidationsSortOrderEnum{
	"ASC":  ListConnectionValidationsSortOrderAsc,
	"DESC": ListConnectionValidationsSortOrderDesc,
}

// GetListConnectionValidationsSortOrderEnumValues Enumerates the set of values for ListConnectionValidationsSortOrderEnum
func GetListConnectionValidationsSortOrderEnumValues() []ListConnectionValidationsSortOrderEnum {
	values := make([]ListConnectionValidationsSortOrderEnum, 0)
	for _, v := range mappingListConnectionValidationsSortOrder {
		values = append(values, v)
	}
	return values
}
