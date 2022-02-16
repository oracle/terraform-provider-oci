// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListPipelinesRequest wrapper for the ListPipelines operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListPipelines.go.html to see an example of how to use ListPipelinesRequest.
type ListPipelinesRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Used to filter by the project or the folder object.
	AggregatorKey *string `mandatory:"false" contributesTo:"query" name:"aggregatorKey"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Used to filter by the identifier of the object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListPipelinesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListPipelinesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPipelinesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPipelinesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPipelinesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPipelinesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPipelinesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPipelinesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPipelinesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPipelinesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPipelinesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPipelinesResponse wrapper for the ListPipelines operation
type ListPipelinesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PipelineSummaryCollection instances
	PipelineSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListPipelinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPipelinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPipelinesSortOrderEnum Enum with underlying type: string
type ListPipelinesSortOrderEnum string

// Set of constants representing the allowable values for ListPipelinesSortOrderEnum
const (
	ListPipelinesSortOrderAsc  ListPipelinesSortOrderEnum = "ASC"
	ListPipelinesSortOrderDesc ListPipelinesSortOrderEnum = "DESC"
)

var mappingListPipelinesSortOrderEnum = map[string]ListPipelinesSortOrderEnum{
	"ASC":  ListPipelinesSortOrderAsc,
	"DESC": ListPipelinesSortOrderDesc,
}

// GetListPipelinesSortOrderEnumValues Enumerates the set of values for ListPipelinesSortOrderEnum
func GetListPipelinesSortOrderEnumValues() []ListPipelinesSortOrderEnum {
	values := make([]ListPipelinesSortOrderEnum, 0)
	for _, v := range mappingListPipelinesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelinesSortOrderEnumStringValues Enumerates the set of values in String for ListPipelinesSortOrderEnum
func GetListPipelinesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPipelinesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelinesSortOrderEnum(val string) (ListPipelinesSortOrderEnum, bool) {
	mappingListPipelinesSortOrderEnumIgnoreCase := make(map[string]ListPipelinesSortOrderEnum)
	for k, v := range mappingListPipelinesSortOrderEnum {
		mappingListPipelinesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListPipelinesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListPipelinesSortByEnum Enum with underlying type: string
type ListPipelinesSortByEnum string

// Set of constants representing the allowable values for ListPipelinesSortByEnum
const (
	ListPipelinesSortByTimeCreated ListPipelinesSortByEnum = "TIME_CREATED"
	ListPipelinesSortByDisplayName ListPipelinesSortByEnum = "DISPLAY_NAME"
)

var mappingListPipelinesSortByEnum = map[string]ListPipelinesSortByEnum{
	"TIME_CREATED": ListPipelinesSortByTimeCreated,
	"DISPLAY_NAME": ListPipelinesSortByDisplayName,
}

// GetListPipelinesSortByEnumValues Enumerates the set of values for ListPipelinesSortByEnum
func GetListPipelinesSortByEnumValues() []ListPipelinesSortByEnum {
	values := make([]ListPipelinesSortByEnum, 0)
	for _, v := range mappingListPipelinesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelinesSortByEnumStringValues Enumerates the set of values in String for ListPipelinesSortByEnum
func GetListPipelinesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListPipelinesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelinesSortByEnum(val string) (ListPipelinesSortByEnum, bool) {
	mappingListPipelinesSortByEnumIgnoreCase := make(map[string]ListPipelinesSortByEnum)
	for k, v := range mappingListPipelinesSortByEnum {
		mappingListPipelinesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListPipelinesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
