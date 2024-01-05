// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRuntimePipelinesRequest wrapper for the ListRuntimePipelines operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListRuntimePipelines.go.html to see an example of how to use ListRuntimePipelinesRequest.
type ListRuntimePipelinesRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The application key.
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Used to filter by the key of the object.
	Key []string `contributesTo:"query" name:"key" collectionFormat:"multi"`

	// Unique key of the aggregator
	AggregatorKey *string `mandatory:"false" contributesTo:"query" name:"aggregatorKey"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Used to filter by the identifier of the object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListRuntimePipelinesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListRuntimePipelinesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique type of the aggregator
	AggregatorType []string `contributesTo:"query" name:"aggregatorType" collectionFormat:"multi"`

	// This filter parameter can be used to filter by model specific queryable fields of the object <br><br><B>Examples:-</B><br> <ul> <li><B>?filter=status eq Failed</B> returns all objects that have a status field with value Failed</li> </ul>
	Filter []string `contributesTo:"query" name:"filter" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRuntimePipelinesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRuntimePipelinesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRuntimePipelinesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRuntimePipelinesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRuntimePipelinesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRuntimePipelinesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRuntimePipelinesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRuntimePipelinesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRuntimePipelinesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRuntimePipelinesResponse wrapper for the ListRuntimePipelines operation
type ListRuntimePipelinesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RuntimePipelineSummaryCollection instances
	RuntimePipelineSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `RuntimePipeline`s. If this header appears in the response, then this
	// is a partial list of RuntimePipeline. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of RuntimePipelines.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRuntimePipelinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRuntimePipelinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRuntimePipelinesSortOrderEnum Enum with underlying type: string
type ListRuntimePipelinesSortOrderEnum string

// Set of constants representing the allowable values for ListRuntimePipelinesSortOrderEnum
const (
	ListRuntimePipelinesSortOrderAsc  ListRuntimePipelinesSortOrderEnum = "ASC"
	ListRuntimePipelinesSortOrderDesc ListRuntimePipelinesSortOrderEnum = "DESC"
)

var mappingListRuntimePipelinesSortOrderEnum = map[string]ListRuntimePipelinesSortOrderEnum{
	"ASC":  ListRuntimePipelinesSortOrderAsc,
	"DESC": ListRuntimePipelinesSortOrderDesc,
}

var mappingListRuntimePipelinesSortOrderEnumLowerCase = map[string]ListRuntimePipelinesSortOrderEnum{
	"asc":  ListRuntimePipelinesSortOrderAsc,
	"desc": ListRuntimePipelinesSortOrderDesc,
}

// GetListRuntimePipelinesSortOrderEnumValues Enumerates the set of values for ListRuntimePipelinesSortOrderEnum
func GetListRuntimePipelinesSortOrderEnumValues() []ListRuntimePipelinesSortOrderEnum {
	values := make([]ListRuntimePipelinesSortOrderEnum, 0)
	for _, v := range mappingListRuntimePipelinesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRuntimePipelinesSortOrderEnumStringValues Enumerates the set of values in String for ListRuntimePipelinesSortOrderEnum
func GetListRuntimePipelinesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRuntimePipelinesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRuntimePipelinesSortOrderEnum(val string) (ListRuntimePipelinesSortOrderEnum, bool) {
	enum, ok := mappingListRuntimePipelinesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRuntimePipelinesSortByEnum Enum with underlying type: string
type ListRuntimePipelinesSortByEnum string

// Set of constants representing the allowable values for ListRuntimePipelinesSortByEnum
const (
	ListRuntimePipelinesSortByTimeCreated ListRuntimePipelinesSortByEnum = "TIME_CREATED"
	ListRuntimePipelinesSortByDisplayName ListRuntimePipelinesSortByEnum = "DISPLAY_NAME"
	ListRuntimePipelinesSortByTimeUpdated ListRuntimePipelinesSortByEnum = "TIME_UPDATED"
)

var mappingListRuntimePipelinesSortByEnum = map[string]ListRuntimePipelinesSortByEnum{
	"TIME_CREATED": ListRuntimePipelinesSortByTimeCreated,
	"DISPLAY_NAME": ListRuntimePipelinesSortByDisplayName,
	"TIME_UPDATED": ListRuntimePipelinesSortByTimeUpdated,
}

var mappingListRuntimePipelinesSortByEnumLowerCase = map[string]ListRuntimePipelinesSortByEnum{
	"time_created": ListRuntimePipelinesSortByTimeCreated,
	"display_name": ListRuntimePipelinesSortByDisplayName,
	"time_updated": ListRuntimePipelinesSortByTimeUpdated,
}

// GetListRuntimePipelinesSortByEnumValues Enumerates the set of values for ListRuntimePipelinesSortByEnum
func GetListRuntimePipelinesSortByEnumValues() []ListRuntimePipelinesSortByEnum {
	values := make([]ListRuntimePipelinesSortByEnum, 0)
	for _, v := range mappingListRuntimePipelinesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRuntimePipelinesSortByEnumStringValues Enumerates the set of values in String for ListRuntimePipelinesSortByEnum
func GetListRuntimePipelinesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListRuntimePipelinesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRuntimePipelinesSortByEnum(val string) (ListRuntimePipelinesSortByEnum, bool) {
	enum, ok := mappingListRuntimePipelinesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
