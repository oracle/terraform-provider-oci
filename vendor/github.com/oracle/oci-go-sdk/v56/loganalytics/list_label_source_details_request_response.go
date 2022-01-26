// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListLabelSourceDetailsRequest wrapper for the ListLabelSourceDetails operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLabelSourceDetails.go.html to see an example of how to use ListLabelSourceDetailsRequest.
type ListLabelSourceDetailsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The label name used for filtering.  Only items with, or associated with, the
	// specified label name will be returned.
	LabelName *string `mandatory:"false" contributesTo:"query" name:"labelName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLabelSourceDetailsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned sources
	LabelSourceSortBy ListLabelSourceDetailsLabelSourceSortByEnum `mandatory:"false" contributesTo:"query" name:"labelSourceSortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLabelSourceDetailsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLabelSourceDetailsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLabelSourceDetailsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLabelSourceDetailsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLabelSourceDetailsResponse wrapper for the ListLabelSourceDetails operation
type ListLabelSourceDetailsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LabelSourceCollection instances
	LabelSourceCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLabelSourceDetailsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLabelSourceDetailsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLabelSourceDetailsSortOrderEnum Enum with underlying type: string
type ListLabelSourceDetailsSortOrderEnum string

// Set of constants representing the allowable values for ListLabelSourceDetailsSortOrderEnum
const (
	ListLabelSourceDetailsSortOrderAsc  ListLabelSourceDetailsSortOrderEnum = "ASC"
	ListLabelSourceDetailsSortOrderDesc ListLabelSourceDetailsSortOrderEnum = "DESC"
)

var mappingListLabelSourceDetailsSortOrder = map[string]ListLabelSourceDetailsSortOrderEnum{
	"ASC":  ListLabelSourceDetailsSortOrderAsc,
	"DESC": ListLabelSourceDetailsSortOrderDesc,
}

// GetListLabelSourceDetailsSortOrderEnumValues Enumerates the set of values for ListLabelSourceDetailsSortOrderEnum
func GetListLabelSourceDetailsSortOrderEnumValues() []ListLabelSourceDetailsSortOrderEnum {
	values := make([]ListLabelSourceDetailsSortOrderEnum, 0)
	for _, v := range mappingListLabelSourceDetailsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListLabelSourceDetailsLabelSourceSortByEnum Enum with underlying type: string
type ListLabelSourceDetailsLabelSourceSortByEnum string

// Set of constants representing the allowable values for ListLabelSourceDetailsLabelSourceSortByEnum
const (
	ListLabelSourceDetailsLabelSourceSortBySourcedisplayname     ListLabelSourceDetailsLabelSourceSortByEnum = "sourceDisplayName"
	ListLabelSourceDetailsLabelSourceSortByLabelfielddisplayname ListLabelSourceDetailsLabelSourceSortByEnum = "labelFieldDisplayName"
)

var mappingListLabelSourceDetailsLabelSourceSortBy = map[string]ListLabelSourceDetailsLabelSourceSortByEnum{
	"sourceDisplayName":     ListLabelSourceDetailsLabelSourceSortBySourcedisplayname,
	"labelFieldDisplayName": ListLabelSourceDetailsLabelSourceSortByLabelfielddisplayname,
}

// GetListLabelSourceDetailsLabelSourceSortByEnumValues Enumerates the set of values for ListLabelSourceDetailsLabelSourceSortByEnum
func GetListLabelSourceDetailsLabelSourceSortByEnumValues() []ListLabelSourceDetailsLabelSourceSortByEnum {
	values := make([]ListLabelSourceDetailsLabelSourceSortByEnum, 0)
	for _, v := range mappingListLabelSourceDetailsLabelSourceSortBy {
		values = append(values, v)
	}
	return values
}
