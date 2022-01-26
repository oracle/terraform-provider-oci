// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package limits

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListQuotasRequest wrapper for the ListQuotas operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limits/ListQuotas.go.html to see an example of how to use ListQuotasRequest.
type ListQuotasRequest struct {

	// The OCID of the parent compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// name
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Filters returned quotas based on the given state.
	LifecycleState ListQuotasLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'. By default, it is ascending.
	SortOrder ListQuotasSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Time created is default ordered as descending. Display name is default ordered as ascending.
	SortBy ListQuotasSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListQuotasRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListQuotasRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListQuotasRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListQuotasRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListQuotasResponse wrapper for the ListQuotas operation
type ListQuotasResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []QuotaSummary instances
	Items []QuotaSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListQuotasResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListQuotasResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListQuotasLifecycleStateEnum Enum with underlying type: string
type ListQuotasLifecycleStateEnum string

// Set of constants representing the allowable values for ListQuotasLifecycleStateEnum
const (
	ListQuotasLifecycleStateActive ListQuotasLifecycleStateEnum = "ACTIVE"
)

var mappingListQuotasLifecycleState = map[string]ListQuotasLifecycleStateEnum{
	"ACTIVE": ListQuotasLifecycleStateActive,
}

// GetListQuotasLifecycleStateEnumValues Enumerates the set of values for ListQuotasLifecycleStateEnum
func GetListQuotasLifecycleStateEnumValues() []ListQuotasLifecycleStateEnum {
	values := make([]ListQuotasLifecycleStateEnum, 0)
	for _, v := range mappingListQuotasLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListQuotasSortOrderEnum Enum with underlying type: string
type ListQuotasSortOrderEnum string

// Set of constants representing the allowable values for ListQuotasSortOrderEnum
const (
	ListQuotasSortOrderAsc  ListQuotasSortOrderEnum = "ASC"
	ListQuotasSortOrderDesc ListQuotasSortOrderEnum = "DESC"
)

var mappingListQuotasSortOrder = map[string]ListQuotasSortOrderEnum{
	"ASC":  ListQuotasSortOrderAsc,
	"DESC": ListQuotasSortOrderDesc,
}

// GetListQuotasSortOrderEnumValues Enumerates the set of values for ListQuotasSortOrderEnum
func GetListQuotasSortOrderEnumValues() []ListQuotasSortOrderEnum {
	values := make([]ListQuotasSortOrderEnum, 0)
	for _, v := range mappingListQuotasSortOrder {
		values = append(values, v)
	}
	return values
}

// ListQuotasSortByEnum Enum with underlying type: string
type ListQuotasSortByEnum string

// Set of constants representing the allowable values for ListQuotasSortByEnum
const (
	ListQuotasSortByName        ListQuotasSortByEnum = "NAME"
	ListQuotasSortByTimecreated ListQuotasSortByEnum = "TIMECREATED"
)

var mappingListQuotasSortBy = map[string]ListQuotasSortByEnum{
	"NAME":        ListQuotasSortByName,
	"TIMECREATED": ListQuotasSortByTimecreated,
}

// GetListQuotasSortByEnumValues Enumerates the set of values for ListQuotasSortByEnum
func GetListQuotasSortByEnumValues() []ListQuotasSortByEnum {
	values := make([]ListQuotasSortByEnum, 0)
	for _, v := range mappingListQuotasSortBy {
		values = append(values, v)
	}
	return values
}
