// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLabelSourceDetailsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLabelSourceDetailsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLabelSourceDetailsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLabelSourceDetailsLabelSourceSortByEnum(string(request.LabelSourceSortBy)); !ok && request.LabelSourceSortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LabelSourceSortBy: %s. Supported values are: %s.", request.LabelSourceSortBy, strings.Join(GetListLabelSourceDetailsLabelSourceSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListLabelSourceDetailsSortOrderEnum = map[string]ListLabelSourceDetailsSortOrderEnum{
	"ASC":  ListLabelSourceDetailsSortOrderAsc,
	"DESC": ListLabelSourceDetailsSortOrderDesc,
}

// GetListLabelSourceDetailsSortOrderEnumValues Enumerates the set of values for ListLabelSourceDetailsSortOrderEnum
func GetListLabelSourceDetailsSortOrderEnumValues() []ListLabelSourceDetailsSortOrderEnum {
	values := make([]ListLabelSourceDetailsSortOrderEnum, 0)
	for _, v := range mappingListLabelSourceDetailsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLabelSourceDetailsSortOrderEnumStringValues Enumerates the set of values in String for ListLabelSourceDetailsSortOrderEnum
func GetListLabelSourceDetailsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLabelSourceDetailsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLabelSourceDetailsSortOrderEnum(val string) (ListLabelSourceDetailsSortOrderEnum, bool) {
	mappingListLabelSourceDetailsSortOrderEnumIgnoreCase := make(map[string]ListLabelSourceDetailsSortOrderEnum)
	for k, v := range mappingListLabelSourceDetailsSortOrderEnum {
		mappingListLabelSourceDetailsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLabelSourceDetailsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListLabelSourceDetailsLabelSourceSortByEnum Enum with underlying type: string
type ListLabelSourceDetailsLabelSourceSortByEnum string

// Set of constants representing the allowable values for ListLabelSourceDetailsLabelSourceSortByEnum
const (
	ListLabelSourceDetailsLabelSourceSortBySourcedisplayname     ListLabelSourceDetailsLabelSourceSortByEnum = "sourceDisplayName"
	ListLabelSourceDetailsLabelSourceSortByLabelfielddisplayname ListLabelSourceDetailsLabelSourceSortByEnum = "labelFieldDisplayName"
)

var mappingListLabelSourceDetailsLabelSourceSortByEnum = map[string]ListLabelSourceDetailsLabelSourceSortByEnum{
	"sourceDisplayName":     ListLabelSourceDetailsLabelSourceSortBySourcedisplayname,
	"labelFieldDisplayName": ListLabelSourceDetailsLabelSourceSortByLabelfielddisplayname,
}

// GetListLabelSourceDetailsLabelSourceSortByEnumValues Enumerates the set of values for ListLabelSourceDetailsLabelSourceSortByEnum
func GetListLabelSourceDetailsLabelSourceSortByEnumValues() []ListLabelSourceDetailsLabelSourceSortByEnum {
	values := make([]ListLabelSourceDetailsLabelSourceSortByEnum, 0)
	for _, v := range mappingListLabelSourceDetailsLabelSourceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLabelSourceDetailsLabelSourceSortByEnumStringValues Enumerates the set of values in String for ListLabelSourceDetailsLabelSourceSortByEnum
func GetListLabelSourceDetailsLabelSourceSortByEnumStringValues() []string {
	return []string{
		"sourceDisplayName",
		"labelFieldDisplayName",
	}
}

// GetMappingListLabelSourceDetailsLabelSourceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLabelSourceDetailsLabelSourceSortByEnum(val string) (ListLabelSourceDetailsLabelSourceSortByEnum, bool) {
	mappingListLabelSourceDetailsLabelSourceSortByEnumIgnoreCase := make(map[string]ListLabelSourceDetailsLabelSourceSortByEnum)
	for k, v := range mappingListLabelSourceDetailsLabelSourceSortByEnum {
		mappingListLabelSourceDetailsLabelSourceSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLabelSourceDetailsLabelSourceSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
