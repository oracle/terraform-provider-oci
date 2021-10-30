// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// ListIamWorkRequestLogsRequest wrapper for the ListIamWorkRequestLogs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListIamWorkRequestLogs.go.html to see an example of how to use ListIamWorkRequestLogsRequest.
type ListIamWorkRequestLogsRequest struct {

	// The OCID of the IAM work request.
	IamWorkRequestId *string `mandatory:"true" contributesTo:"path" name:"iamWorkRequestId"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListIamWorkRequestLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIamWorkRequestLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIamWorkRequestLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIamWorkRequestLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIamWorkRequestLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListIamWorkRequestLogsResponse wrapper for the ListIamWorkRequestLogs operation
type ListIamWorkRequestLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []IamWorkRequestLogSummary instances
	Items []IamWorkRequestLogSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The number of seconds that the client should wait before polling again.
	RetryAfter *float32 `presentIn:"header" name:"retry-after"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListIamWorkRequestLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIamWorkRequestLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIamWorkRequestLogsSortOrderEnum Enum with underlying type: string
type ListIamWorkRequestLogsSortOrderEnum string

// Set of constants representing the allowable values for ListIamWorkRequestLogsSortOrderEnum
const (
	ListIamWorkRequestLogsSortOrderAsc  ListIamWorkRequestLogsSortOrderEnum = "ASC"
	ListIamWorkRequestLogsSortOrderDesc ListIamWorkRequestLogsSortOrderEnum = "DESC"
)

var mappingListIamWorkRequestLogsSortOrder = map[string]ListIamWorkRequestLogsSortOrderEnum{
	"ASC":  ListIamWorkRequestLogsSortOrderAsc,
	"DESC": ListIamWorkRequestLogsSortOrderDesc,
}

// GetListIamWorkRequestLogsSortOrderEnumValues Enumerates the set of values for ListIamWorkRequestLogsSortOrderEnum
func GetListIamWorkRequestLogsSortOrderEnumValues() []ListIamWorkRequestLogsSortOrderEnum {
	values := make([]ListIamWorkRequestLogsSortOrderEnum, 0)
	for _, v := range mappingListIamWorkRequestLogsSortOrder {
		values = append(values, v)
	}
	return values
}
