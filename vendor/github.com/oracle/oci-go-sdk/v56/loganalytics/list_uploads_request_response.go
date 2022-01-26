// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListUploadsRequest wrapper for the ListUploads operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListUploads.go.html to see an example of how to use ListUploadsRequest.
type ListUploadsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Name of the upload container.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only uploads whose name contains the given name.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListUploadsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeUpdated is descending.
	// Default order for name is ascending. If no value is specified timeUpdated is default.
	SortBy ListUploadsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Use this for filtering uploads w.r.t warnings. Only one value is allowed. If no value is specified then ALL is taken as default,
	// which means that all the uploads with and without warnings will be returned.
	WarningsFilter ListUploadsWarningsFilterEnum `mandatory:"false" contributesTo:"query" name:"warningsFilter" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUploadsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUploadsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUploadsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUploadsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListUploadsResponse wrapper for the ListUploads operation
type ListUploadsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UploadCollection instances
	UploadCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Total count.
	OpcTotalItems *int64 `presentIn:"header" name:"opc-total-items"`
}

func (response ListUploadsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUploadsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUploadsSortOrderEnum Enum with underlying type: string
type ListUploadsSortOrderEnum string

// Set of constants representing the allowable values for ListUploadsSortOrderEnum
const (
	ListUploadsSortOrderAsc  ListUploadsSortOrderEnum = "ASC"
	ListUploadsSortOrderDesc ListUploadsSortOrderEnum = "DESC"
)

var mappingListUploadsSortOrder = map[string]ListUploadsSortOrderEnum{
	"ASC":  ListUploadsSortOrderAsc,
	"DESC": ListUploadsSortOrderDesc,
}

// GetListUploadsSortOrderEnumValues Enumerates the set of values for ListUploadsSortOrderEnum
func GetListUploadsSortOrderEnumValues() []ListUploadsSortOrderEnum {
	values := make([]ListUploadsSortOrderEnum, 0)
	for _, v := range mappingListUploadsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListUploadsSortByEnum Enum with underlying type: string
type ListUploadsSortByEnum string

// Set of constants representing the allowable values for ListUploadsSortByEnum
const (
	ListUploadsSortByTimeupdated ListUploadsSortByEnum = "timeUpdated"
	ListUploadsSortByTimecreated ListUploadsSortByEnum = "timeCreated"
	ListUploadsSortByName        ListUploadsSortByEnum = "name"
)

var mappingListUploadsSortBy = map[string]ListUploadsSortByEnum{
	"timeUpdated": ListUploadsSortByTimeupdated,
	"timeCreated": ListUploadsSortByTimecreated,
	"name":        ListUploadsSortByName,
}

// GetListUploadsSortByEnumValues Enumerates the set of values for ListUploadsSortByEnum
func GetListUploadsSortByEnumValues() []ListUploadsSortByEnum {
	values := make([]ListUploadsSortByEnum, 0)
	for _, v := range mappingListUploadsSortBy {
		values = append(values, v)
	}
	return values
}

// ListUploadsWarningsFilterEnum Enum with underlying type: string
type ListUploadsWarningsFilterEnum string

// Set of constants representing the allowable values for ListUploadsWarningsFilterEnum
const (
	ListUploadsWarningsFilterWithWarnings    ListUploadsWarningsFilterEnum = "WITH_WARNINGS"
	ListUploadsWarningsFilterWithoutWarnings ListUploadsWarningsFilterEnum = "WITHOUT_WARNINGS"
	ListUploadsWarningsFilterAll             ListUploadsWarningsFilterEnum = "ALL"
)

var mappingListUploadsWarningsFilter = map[string]ListUploadsWarningsFilterEnum{
	"WITH_WARNINGS":    ListUploadsWarningsFilterWithWarnings,
	"WITHOUT_WARNINGS": ListUploadsWarningsFilterWithoutWarnings,
	"ALL":              ListUploadsWarningsFilterAll,
}

// GetListUploadsWarningsFilterEnumValues Enumerates the set of values for ListUploadsWarningsFilterEnum
func GetListUploadsWarningsFilterEnumValues() []ListUploadsWarningsFilterEnum {
	values := make([]ListUploadsWarningsFilterEnum, 0)
	for _, v := range mappingListUploadsWarningsFilter {
		values = append(values, v)
	}
	return values
}
