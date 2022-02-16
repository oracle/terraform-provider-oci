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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUploadsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUploadsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUploadsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUploadsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUploadsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUploadsWarningsFilterEnum(string(request.WarningsFilter)); !ok && request.WarningsFilter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WarningsFilter: %s. Supported values are: %s.", request.WarningsFilter, strings.Join(GetListUploadsWarningsFilterEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListUploadsSortOrderEnum = map[string]ListUploadsSortOrderEnum{
	"ASC":  ListUploadsSortOrderAsc,
	"DESC": ListUploadsSortOrderDesc,
}

// GetListUploadsSortOrderEnumValues Enumerates the set of values for ListUploadsSortOrderEnum
func GetListUploadsSortOrderEnumValues() []ListUploadsSortOrderEnum {
	values := make([]ListUploadsSortOrderEnum, 0)
	for _, v := range mappingListUploadsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUploadsSortOrderEnumStringValues Enumerates the set of values in String for ListUploadsSortOrderEnum
func GetListUploadsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUploadsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUploadsSortOrderEnum(val string) (ListUploadsSortOrderEnum, bool) {
	mappingListUploadsSortOrderEnumIgnoreCase := make(map[string]ListUploadsSortOrderEnum)
	for k, v := range mappingListUploadsSortOrderEnum {
		mappingListUploadsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUploadsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListUploadsSortByEnum Enum with underlying type: string
type ListUploadsSortByEnum string

// Set of constants representing the allowable values for ListUploadsSortByEnum
const (
	ListUploadsSortByTimeupdated ListUploadsSortByEnum = "timeUpdated"
	ListUploadsSortByTimecreated ListUploadsSortByEnum = "timeCreated"
	ListUploadsSortByName        ListUploadsSortByEnum = "name"
)

var mappingListUploadsSortByEnum = map[string]ListUploadsSortByEnum{
	"timeUpdated": ListUploadsSortByTimeupdated,
	"timeCreated": ListUploadsSortByTimecreated,
	"name":        ListUploadsSortByName,
}

// GetListUploadsSortByEnumValues Enumerates the set of values for ListUploadsSortByEnum
func GetListUploadsSortByEnumValues() []ListUploadsSortByEnum {
	values := make([]ListUploadsSortByEnum, 0)
	for _, v := range mappingListUploadsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUploadsSortByEnumStringValues Enumerates the set of values in String for ListUploadsSortByEnum
func GetListUploadsSortByEnumStringValues() []string {
	return []string{
		"timeUpdated",
		"timeCreated",
		"name",
	}
}

// GetMappingListUploadsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUploadsSortByEnum(val string) (ListUploadsSortByEnum, bool) {
	mappingListUploadsSortByEnumIgnoreCase := make(map[string]ListUploadsSortByEnum)
	for k, v := range mappingListUploadsSortByEnum {
		mappingListUploadsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUploadsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListUploadsWarningsFilterEnum Enum with underlying type: string
type ListUploadsWarningsFilterEnum string

// Set of constants representing the allowable values for ListUploadsWarningsFilterEnum
const (
	ListUploadsWarningsFilterWithWarnings    ListUploadsWarningsFilterEnum = "WITH_WARNINGS"
	ListUploadsWarningsFilterWithoutWarnings ListUploadsWarningsFilterEnum = "WITHOUT_WARNINGS"
	ListUploadsWarningsFilterAll             ListUploadsWarningsFilterEnum = "ALL"
)

var mappingListUploadsWarningsFilterEnum = map[string]ListUploadsWarningsFilterEnum{
	"WITH_WARNINGS":    ListUploadsWarningsFilterWithWarnings,
	"WITHOUT_WARNINGS": ListUploadsWarningsFilterWithoutWarnings,
	"ALL":              ListUploadsWarningsFilterAll,
}

// GetListUploadsWarningsFilterEnumValues Enumerates the set of values for ListUploadsWarningsFilterEnum
func GetListUploadsWarningsFilterEnumValues() []ListUploadsWarningsFilterEnum {
	values := make([]ListUploadsWarningsFilterEnum, 0)
	for _, v := range mappingListUploadsWarningsFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetListUploadsWarningsFilterEnumStringValues Enumerates the set of values in String for ListUploadsWarningsFilterEnum
func GetListUploadsWarningsFilterEnumStringValues() []string {
	return []string{
		"WITH_WARNINGS",
		"WITHOUT_WARNINGS",
		"ALL",
	}
}

// GetMappingListUploadsWarningsFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUploadsWarningsFilterEnum(val string) (ListUploadsWarningsFilterEnum, bool) {
	mappingListUploadsWarningsFilterEnumIgnoreCase := make(map[string]ListUploadsWarningsFilterEnum)
	for k, v := range mappingListUploadsWarningsFilterEnum {
		mappingListUploadsWarningsFilterEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUploadsWarningsFilterEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
