// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package limits

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListServicesRequest wrapper for the ListServices operation
type ListServicesRequest struct {

	// The OCID of the parent compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The field to sort by.
	SortBy ListServicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'. By default it will be ascending.
	SortOrder ListServicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListServicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListServicesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListServicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListServicesResponse wrapper for the ListServices operation
type ListServicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ServiceSummary instances
	Items []ServiceSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListServicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListServicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListServicesSortByEnum Enum with underlying type: string
type ListServicesSortByEnum string

// Set of constants representing the allowable values for ListServicesSortByEnum
const (
	ListServicesSortByName        ListServicesSortByEnum = "name"
	ListServicesSortByDescription ListServicesSortByEnum = "description"
)

var mappingListServicesSortBy = map[string]ListServicesSortByEnum{
	"name":        ListServicesSortByName,
	"description": ListServicesSortByDescription,
}

// GetListServicesSortByEnumValues Enumerates the set of values for ListServicesSortByEnum
func GetListServicesSortByEnumValues() []ListServicesSortByEnum {
	values := make([]ListServicesSortByEnum, 0)
	for _, v := range mappingListServicesSortBy {
		values = append(values, v)
	}
	return values
}

// ListServicesSortOrderEnum Enum with underlying type: string
type ListServicesSortOrderEnum string

// Set of constants representing the allowable values for ListServicesSortOrderEnum
const (
	ListServicesSortOrderAsc  ListServicesSortOrderEnum = "ASC"
	ListServicesSortOrderDesc ListServicesSortOrderEnum = "DESC"
)

var mappingListServicesSortOrder = map[string]ListServicesSortOrderEnum{
	"ASC":  ListServicesSortOrderAsc,
	"DESC": ListServicesSortOrderDesc,
}

// GetListServicesSortOrderEnumValues Enumerates the set of values for ListServicesSortOrderEnum
func GetListServicesSortOrderEnumValues() []ListServicesSortOrderEnum {
	values := make([]ListServicesSortOrderEnum, 0)
	for _, v := range mappingListServicesSortOrder {
		values = append(values, v)
	}
	return values
}
