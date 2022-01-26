// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListLabelsRequest wrapper for the ListLabels operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLabels.go.html to see an example of how to use ListLabelsRequest.
type ListLabelsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The label name used for filtering.  Only items with, or associated with, the
	// specified label name will be returned.
	LabelName *string `mandatory:"false" contributesTo:"query" name:"labelName"`

	// The label display text used for filtering.  Only labels with the specified name or
	// description will be returned.
	LabelDisplayText *string `mandatory:"false" contributesTo:"query" name:"labelDisplayText"`

	// The system value used for filtering.  Only items with the specified system value
	// will be returned.  Valid values are built in, custom (for user defined items), or
	// all (for all items, regardless of system value).
	IsSystem ListLabelsIsSystemEnum `mandatory:"false" contributesTo:"query" name:"isSystem" omitEmpty:"true"`

	// The label priority used for filtering.  Only labels with the specified
	// priority will be returned.
	LabelPriority ListLabelsLabelPriorityEnum `mandatory:"false" contributesTo:"query" name:"labelPriority" omitEmpty:"true"`

	// A flag indicating whether or not to count the label usage per source and per rule.
	IsCountPop *bool `mandatory:"false" contributesTo:"query" name:"isCountPop"`

	// A flag indicating whether or not return the aliases used by each label.
	IsAliasPop *bool `mandatory:"false" contributesTo:"query" name:"isAliasPop"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLabelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned labels
	LabelSortBy ListLabelsLabelSortByEnum `mandatory:"false" contributesTo:"query" name:"labelSortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLabelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLabelsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLabelsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLabelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLabelsResponse wrapper for the ListLabels operation
type ListLabelsResponse struct {

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

func (response ListLabelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLabelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLabelsIsSystemEnum Enum with underlying type: string
type ListLabelsIsSystemEnum string

// Set of constants representing the allowable values for ListLabelsIsSystemEnum
const (
	ListLabelsIsSystemAll     ListLabelsIsSystemEnum = "ALL"
	ListLabelsIsSystemCustom  ListLabelsIsSystemEnum = "CUSTOM"
	ListLabelsIsSystemBuiltIn ListLabelsIsSystemEnum = "BUILT_IN"
)

var mappingListLabelsIsSystem = map[string]ListLabelsIsSystemEnum{
	"ALL":      ListLabelsIsSystemAll,
	"CUSTOM":   ListLabelsIsSystemCustom,
	"BUILT_IN": ListLabelsIsSystemBuiltIn,
}

// GetListLabelsIsSystemEnumValues Enumerates the set of values for ListLabelsIsSystemEnum
func GetListLabelsIsSystemEnumValues() []ListLabelsIsSystemEnum {
	values := make([]ListLabelsIsSystemEnum, 0)
	for _, v := range mappingListLabelsIsSystem {
		values = append(values, v)
	}
	return values
}

// ListLabelsLabelPriorityEnum Enum with underlying type: string
type ListLabelsLabelPriorityEnum string

// Set of constants representing the allowable values for ListLabelsLabelPriorityEnum
const (
	ListLabelsLabelPriorityNone   ListLabelsLabelPriorityEnum = "NONE"
	ListLabelsLabelPriorityLow    ListLabelsLabelPriorityEnum = "LOW"
	ListLabelsLabelPriorityMedium ListLabelsLabelPriorityEnum = "MEDIUM"
	ListLabelsLabelPriorityHigh   ListLabelsLabelPriorityEnum = "HIGH"
)

var mappingListLabelsLabelPriority = map[string]ListLabelsLabelPriorityEnum{
	"NONE":   ListLabelsLabelPriorityNone,
	"LOW":    ListLabelsLabelPriorityLow,
	"MEDIUM": ListLabelsLabelPriorityMedium,
	"HIGH":   ListLabelsLabelPriorityHigh,
}

// GetListLabelsLabelPriorityEnumValues Enumerates the set of values for ListLabelsLabelPriorityEnum
func GetListLabelsLabelPriorityEnumValues() []ListLabelsLabelPriorityEnum {
	values := make([]ListLabelsLabelPriorityEnum, 0)
	for _, v := range mappingListLabelsLabelPriority {
		values = append(values, v)
	}
	return values
}

// ListLabelsSortOrderEnum Enum with underlying type: string
type ListLabelsSortOrderEnum string

// Set of constants representing the allowable values for ListLabelsSortOrderEnum
const (
	ListLabelsSortOrderAsc  ListLabelsSortOrderEnum = "ASC"
	ListLabelsSortOrderDesc ListLabelsSortOrderEnum = "DESC"
)

var mappingListLabelsSortOrder = map[string]ListLabelsSortOrderEnum{
	"ASC":  ListLabelsSortOrderAsc,
	"DESC": ListLabelsSortOrderDesc,
}

// GetListLabelsSortOrderEnumValues Enumerates the set of values for ListLabelsSortOrderEnum
func GetListLabelsSortOrderEnumValues() []ListLabelsSortOrderEnum {
	values := make([]ListLabelsSortOrderEnum, 0)
	for _, v := range mappingListLabelsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListLabelsLabelSortByEnum Enum with underlying type: string
type ListLabelsLabelSortByEnum string

// Set of constants representing the allowable values for ListLabelsLabelSortByEnum
const (
	ListLabelsLabelSortByName        ListLabelsLabelSortByEnum = "name"
	ListLabelsLabelSortByPriority    ListLabelsLabelSortByEnum = "priority"
	ListLabelsLabelSortBySourceusing ListLabelsLabelSortByEnum = "sourceUsing"
)

var mappingListLabelsLabelSortBy = map[string]ListLabelsLabelSortByEnum{
	"name":        ListLabelsLabelSortByName,
	"priority":    ListLabelsLabelSortByPriority,
	"sourceUsing": ListLabelsLabelSortBySourceusing,
}

// GetListLabelsLabelSortByEnumValues Enumerates the set of values for ListLabelsLabelSortByEnum
func GetListLabelsLabelSortByEnumValues() []ListLabelsLabelSortByEnum {
	values := make([]ListLabelsLabelSortByEnum, 0)
	for _, v := range mappingListLabelsLabelSortBy {
		values = append(values, v)
	}
	return values
}
