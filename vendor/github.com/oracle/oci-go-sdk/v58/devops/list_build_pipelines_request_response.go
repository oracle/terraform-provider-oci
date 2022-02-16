// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListBuildPipelinesRequest wrapper for the ListBuildPipelines operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListBuildPipelines.go.html to see an example of how to use ListBuildPipelinesRequest.
type ListBuildPipelinesRequest struct {

	// Unique identifier or OCID for listing a single resource by ID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// unique project identifier
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only build pipelines that matches the given lifecycle state.
	LifecycleState BuildPipelineLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListBuildPipelinesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for display name is ascending. If no value is specified, then the default time created value is considered.
	SortBy ListBuildPipelinesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBuildPipelinesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBuildPipelinesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBuildPipelinesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBuildPipelinesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBuildPipelinesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBuildPipelineLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBuildPipelineLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBuildPipelinesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBuildPipelinesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBuildPipelinesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBuildPipelinesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBuildPipelinesResponse wrapper for the ListBuildPipelines operation
type ListBuildPipelinesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BuildPipelineCollection instances
	BuildPipelineCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBuildPipelinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBuildPipelinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBuildPipelinesSortOrderEnum Enum with underlying type: string
type ListBuildPipelinesSortOrderEnum string

// Set of constants representing the allowable values for ListBuildPipelinesSortOrderEnum
const (
	ListBuildPipelinesSortOrderAsc  ListBuildPipelinesSortOrderEnum = "ASC"
	ListBuildPipelinesSortOrderDesc ListBuildPipelinesSortOrderEnum = "DESC"
)

var mappingListBuildPipelinesSortOrderEnum = map[string]ListBuildPipelinesSortOrderEnum{
	"ASC":  ListBuildPipelinesSortOrderAsc,
	"DESC": ListBuildPipelinesSortOrderDesc,
}

// GetListBuildPipelinesSortOrderEnumValues Enumerates the set of values for ListBuildPipelinesSortOrderEnum
func GetListBuildPipelinesSortOrderEnumValues() []ListBuildPipelinesSortOrderEnum {
	values := make([]ListBuildPipelinesSortOrderEnum, 0)
	for _, v := range mappingListBuildPipelinesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBuildPipelinesSortOrderEnumStringValues Enumerates the set of values in String for ListBuildPipelinesSortOrderEnum
func GetListBuildPipelinesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBuildPipelinesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBuildPipelinesSortOrderEnum(val string) (ListBuildPipelinesSortOrderEnum, bool) {
	mappingListBuildPipelinesSortOrderEnumIgnoreCase := make(map[string]ListBuildPipelinesSortOrderEnum)
	for k, v := range mappingListBuildPipelinesSortOrderEnum {
		mappingListBuildPipelinesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListBuildPipelinesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListBuildPipelinesSortByEnum Enum with underlying type: string
type ListBuildPipelinesSortByEnum string

// Set of constants representing the allowable values for ListBuildPipelinesSortByEnum
const (
	ListBuildPipelinesSortByTimecreated ListBuildPipelinesSortByEnum = "timeCreated"
	ListBuildPipelinesSortByDisplayname ListBuildPipelinesSortByEnum = "displayName"
)

var mappingListBuildPipelinesSortByEnum = map[string]ListBuildPipelinesSortByEnum{
	"timeCreated": ListBuildPipelinesSortByTimecreated,
	"displayName": ListBuildPipelinesSortByDisplayname,
}

// GetListBuildPipelinesSortByEnumValues Enumerates the set of values for ListBuildPipelinesSortByEnum
func GetListBuildPipelinesSortByEnumValues() []ListBuildPipelinesSortByEnum {
	values := make([]ListBuildPipelinesSortByEnum, 0)
	for _, v := range mappingListBuildPipelinesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBuildPipelinesSortByEnumStringValues Enumerates the set of values in String for ListBuildPipelinesSortByEnum
func GetListBuildPipelinesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBuildPipelinesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBuildPipelinesSortByEnum(val string) (ListBuildPipelinesSortByEnum, bool) {
	mappingListBuildPipelinesSortByEnumIgnoreCase := make(map[string]ListBuildPipelinesSortByEnum)
	for k, v := range mappingListBuildPipelinesSortByEnum {
		mappingListBuildPipelinesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListBuildPipelinesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
