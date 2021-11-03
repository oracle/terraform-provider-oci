// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// ListIamWorkRequestErrorsRequest wrapper for the ListIamWorkRequestErrors operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListIamWorkRequestErrors.go.html to see an example of how to use ListIamWorkRequestErrorsRequest.
type ListIamWorkRequestErrorsRequest struct {

	// The OCID of the IAM work request.
	IamWorkRequestId *string `mandatory:"true" contributesTo:"path" name:"iamWorkRequestId"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListIamWorkRequestErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIamWorkRequestErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIamWorkRequestErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIamWorkRequestErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIamWorkRequestErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListIamWorkRequestErrorsResponse wrapper for the ListIamWorkRequestErrors operation
type ListIamWorkRequestErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []IamWorkRequestErrorSummary instances
	Items []IamWorkRequestErrorSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of
	// results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListIamWorkRequestErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIamWorkRequestErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIamWorkRequestErrorsSortOrderEnum Enum with underlying type: string
type ListIamWorkRequestErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListIamWorkRequestErrorsSortOrderEnum
const (
	ListIamWorkRequestErrorsSortOrderAsc  ListIamWorkRequestErrorsSortOrderEnum = "ASC"
	ListIamWorkRequestErrorsSortOrderDesc ListIamWorkRequestErrorsSortOrderEnum = "DESC"
)

var mappingListIamWorkRequestErrorsSortOrder = map[string]ListIamWorkRequestErrorsSortOrderEnum{
	"ASC":  ListIamWorkRequestErrorsSortOrderAsc,
	"DESC": ListIamWorkRequestErrorsSortOrderDesc,
}

// GetListIamWorkRequestErrorsSortOrderEnumValues Enumerates the set of values for ListIamWorkRequestErrorsSortOrderEnum
func GetListIamWorkRequestErrorsSortOrderEnumValues() []ListIamWorkRequestErrorsSortOrderEnum {
	values := make([]ListIamWorkRequestErrorsSortOrderEnum, 0)
	for _, v := range mappingListIamWorkRequestErrorsSortOrder {
		values = append(values, v)
	}
	return values
}
