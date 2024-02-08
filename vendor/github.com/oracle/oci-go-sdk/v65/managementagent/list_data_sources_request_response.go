// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDataSourcesRequest wrapper for the ListDataSources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementagent/ListDataSources.go.html to see an example of how to use ListDataSourcesRequest.
type ListDataSourcesRequest struct {

	// Unique Management Agent identifier
	ManagementAgentId *string `mandatory:"true" contributesTo:"path" name:"managementAgentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. If no value is specified dataSourceName is default.
	SortBy ListDataSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique name of the dataSource.
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataSourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataSourcesResponse wrapper for the ListDataSources operation
type ListDataSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DataSourceSummary instances
	Items []DataSourceSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDataSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataSourcesSortOrderEnum Enum with underlying type: string
type ListDataSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListDataSourcesSortOrderEnum
const (
	ListDataSourcesSortOrderAsc  ListDataSourcesSortOrderEnum = "ASC"
	ListDataSourcesSortOrderDesc ListDataSourcesSortOrderEnum = "DESC"
)

var mappingListDataSourcesSortOrderEnum = map[string]ListDataSourcesSortOrderEnum{
	"ASC":  ListDataSourcesSortOrderAsc,
	"DESC": ListDataSourcesSortOrderDesc,
}

var mappingListDataSourcesSortOrderEnumLowerCase = map[string]ListDataSourcesSortOrderEnum{
	"asc":  ListDataSourcesSortOrderAsc,
	"desc": ListDataSourcesSortOrderDesc,
}

// GetListDataSourcesSortOrderEnumValues Enumerates the set of values for ListDataSourcesSortOrderEnum
func GetListDataSourcesSortOrderEnumValues() []ListDataSourcesSortOrderEnum {
	values := make([]ListDataSourcesSortOrderEnum, 0)
	for _, v := range mappingListDataSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListDataSourcesSortOrderEnum
func GetListDataSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSourcesSortOrderEnum(val string) (ListDataSourcesSortOrderEnum, bool) {
	enum, ok := mappingListDataSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSourcesSortByEnum Enum with underlying type: string
type ListDataSourcesSortByEnum string

// Set of constants representing the allowable values for ListDataSourcesSortByEnum
const (
	ListDataSourcesSortByDatasourcename ListDataSourcesSortByEnum = "dataSourceName"
	ListDataSourcesSortByDatasourcetype ListDataSourcesSortByEnum = "dataSourceType"
)

var mappingListDataSourcesSortByEnum = map[string]ListDataSourcesSortByEnum{
	"dataSourceName": ListDataSourcesSortByDatasourcename,
	"dataSourceType": ListDataSourcesSortByDatasourcetype,
}

var mappingListDataSourcesSortByEnumLowerCase = map[string]ListDataSourcesSortByEnum{
	"datasourcename": ListDataSourcesSortByDatasourcename,
	"datasourcetype": ListDataSourcesSortByDatasourcetype,
}

// GetListDataSourcesSortByEnumValues Enumerates the set of values for ListDataSourcesSortByEnum
func GetListDataSourcesSortByEnumValues() []ListDataSourcesSortByEnum {
	values := make([]ListDataSourcesSortByEnum, 0)
	for _, v := range mappingListDataSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSourcesSortByEnumStringValues Enumerates the set of values in String for ListDataSourcesSortByEnum
func GetListDataSourcesSortByEnumStringValues() []string {
	return []string{
		"dataSourceName",
		"dataSourceType",
	}
}

// GetMappingListDataSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSourcesSortByEnum(val string) (ListDataSourcesSortByEnum, bool) {
	enum, ok := mappingListDataSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
