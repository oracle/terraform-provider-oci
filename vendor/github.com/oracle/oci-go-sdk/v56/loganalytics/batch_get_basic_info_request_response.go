// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// BatchGetBasicInfoRequest wrapper for the BatchGetBasicInfo operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/BatchGetBasicInfo.go.html to see an example of how to use BatchGetBasicInfoRequest.
type BatchGetBasicInfoRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// List of label names to get information on
	BasicDetails LabelNames `contributesTo:"body"`

	// A flag specifying whether or not to include information on deleted labels.
	IsIncludeDeleted *bool `mandatory:"true" contributesTo:"query" name:"isIncludeDeleted"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder BatchGetBasicInfoSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned labels
	BasicLabelSortBy BatchGetBasicInfoBasicLabelSortByEnum `mandatory:"false" contributesTo:"query" name:"basicLabelSortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request BatchGetBasicInfoRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request BatchGetBasicInfoRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request BatchGetBasicInfoRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request BatchGetBasicInfoRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// BatchGetBasicInfoResponse wrapper for the BatchGetBasicInfo operation
type BatchGetBasicInfoResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsLabelCollection instances
	LogAnalyticsLabelCollection `presentIn:"body"`

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

func (response BatchGetBasicInfoResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response BatchGetBasicInfoResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// BatchGetBasicInfoSortOrderEnum Enum with underlying type: string
type BatchGetBasicInfoSortOrderEnum string

// Set of constants representing the allowable values for BatchGetBasicInfoSortOrderEnum
const (
	BatchGetBasicInfoSortOrderAsc  BatchGetBasicInfoSortOrderEnum = "ASC"
	BatchGetBasicInfoSortOrderDesc BatchGetBasicInfoSortOrderEnum = "DESC"
)

var mappingBatchGetBasicInfoSortOrder = map[string]BatchGetBasicInfoSortOrderEnum{
	"ASC":  BatchGetBasicInfoSortOrderAsc,
	"DESC": BatchGetBasicInfoSortOrderDesc,
}

// GetBatchGetBasicInfoSortOrderEnumValues Enumerates the set of values for BatchGetBasicInfoSortOrderEnum
func GetBatchGetBasicInfoSortOrderEnumValues() []BatchGetBasicInfoSortOrderEnum {
	values := make([]BatchGetBasicInfoSortOrderEnum, 0)
	for _, v := range mappingBatchGetBasicInfoSortOrder {
		values = append(values, v)
	}
	return values
}

// BatchGetBasicInfoBasicLabelSortByEnum Enum with underlying type: string
type BatchGetBasicInfoBasicLabelSortByEnum string

// Set of constants representing the allowable values for BatchGetBasicInfoBasicLabelSortByEnum
const (
	BatchGetBasicInfoBasicLabelSortByName     BatchGetBasicInfoBasicLabelSortByEnum = "name"
	BatchGetBasicInfoBasicLabelSortByPriority BatchGetBasicInfoBasicLabelSortByEnum = "priority"
)

var mappingBatchGetBasicInfoBasicLabelSortBy = map[string]BatchGetBasicInfoBasicLabelSortByEnum{
	"name":     BatchGetBasicInfoBasicLabelSortByName,
	"priority": BatchGetBasicInfoBasicLabelSortByPriority,
}

// GetBatchGetBasicInfoBasicLabelSortByEnumValues Enumerates the set of values for BatchGetBasicInfoBasicLabelSortByEnum
func GetBatchGetBasicInfoBasicLabelSortByEnumValues() []BatchGetBasicInfoBasicLabelSortByEnum {
	values := make([]BatchGetBasicInfoBasicLabelSortByEnum, 0)
	for _, v := range mappingBatchGetBasicInfoBasicLabelSortBy {
		values = append(values, v)
	}
	return values
}
