// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOpensearchClusterPipelinesRequest wrapper for the ListOpensearchClusterPipelines operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opensearch/ListOpensearchClusterPipelines.go.html to see an example of how to use ListOpensearchClusterPipelinesRequest.
type ListOpensearchClusterPipelinesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState OpensearchClusterPipelineLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return pipelines whose any component has the given pipelineComponentId.
	PipelineComponentId *string `mandatory:"false" contributesTo:"query" name:"pipelineComponentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique OpensearchClusterPipeline identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOpensearchClusterPipelinesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListOpensearchClusterPipelinesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOpensearchClusterPipelinesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOpensearchClusterPipelinesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOpensearchClusterPipelinesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOpensearchClusterPipelinesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOpensearchClusterPipelinesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOpensearchClusterPipelineLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOpensearchClusterPipelineLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOpensearchClusterPipelinesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOpensearchClusterPipelinesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOpensearchClusterPipelinesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOpensearchClusterPipelinesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOpensearchClusterPipelinesResponse wrapper for the ListOpensearchClusterPipelines operation
type ListOpensearchClusterPipelinesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OpensearchClusterPipelineCollection instances
	OpensearchClusterPipelineCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOpensearchClusterPipelinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOpensearchClusterPipelinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOpensearchClusterPipelinesSortOrderEnum Enum with underlying type: string
type ListOpensearchClusterPipelinesSortOrderEnum string

// Set of constants representing the allowable values for ListOpensearchClusterPipelinesSortOrderEnum
const (
	ListOpensearchClusterPipelinesSortOrderAsc  ListOpensearchClusterPipelinesSortOrderEnum = "ASC"
	ListOpensearchClusterPipelinesSortOrderDesc ListOpensearchClusterPipelinesSortOrderEnum = "DESC"
)

var mappingListOpensearchClusterPipelinesSortOrderEnum = map[string]ListOpensearchClusterPipelinesSortOrderEnum{
	"ASC":  ListOpensearchClusterPipelinesSortOrderAsc,
	"DESC": ListOpensearchClusterPipelinesSortOrderDesc,
}

var mappingListOpensearchClusterPipelinesSortOrderEnumLowerCase = map[string]ListOpensearchClusterPipelinesSortOrderEnum{
	"asc":  ListOpensearchClusterPipelinesSortOrderAsc,
	"desc": ListOpensearchClusterPipelinesSortOrderDesc,
}

// GetListOpensearchClusterPipelinesSortOrderEnumValues Enumerates the set of values for ListOpensearchClusterPipelinesSortOrderEnum
func GetListOpensearchClusterPipelinesSortOrderEnumValues() []ListOpensearchClusterPipelinesSortOrderEnum {
	values := make([]ListOpensearchClusterPipelinesSortOrderEnum, 0)
	for _, v := range mappingListOpensearchClusterPipelinesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpensearchClusterPipelinesSortOrderEnumStringValues Enumerates the set of values in String for ListOpensearchClusterPipelinesSortOrderEnum
func GetListOpensearchClusterPipelinesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOpensearchClusterPipelinesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpensearchClusterPipelinesSortOrderEnum(val string) (ListOpensearchClusterPipelinesSortOrderEnum, bool) {
	enum, ok := mappingListOpensearchClusterPipelinesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOpensearchClusterPipelinesSortByEnum Enum with underlying type: string
type ListOpensearchClusterPipelinesSortByEnum string

// Set of constants representing the allowable values for ListOpensearchClusterPipelinesSortByEnum
const (
	ListOpensearchClusterPipelinesSortByTimecreated ListOpensearchClusterPipelinesSortByEnum = "timeCreated"
	ListOpensearchClusterPipelinesSortByDisplayname ListOpensearchClusterPipelinesSortByEnum = "displayName"
)

var mappingListOpensearchClusterPipelinesSortByEnum = map[string]ListOpensearchClusterPipelinesSortByEnum{
	"timeCreated": ListOpensearchClusterPipelinesSortByTimecreated,
	"displayName": ListOpensearchClusterPipelinesSortByDisplayname,
}

var mappingListOpensearchClusterPipelinesSortByEnumLowerCase = map[string]ListOpensearchClusterPipelinesSortByEnum{
	"timecreated": ListOpensearchClusterPipelinesSortByTimecreated,
	"displayname": ListOpensearchClusterPipelinesSortByDisplayname,
}

// GetListOpensearchClusterPipelinesSortByEnumValues Enumerates the set of values for ListOpensearchClusterPipelinesSortByEnum
func GetListOpensearchClusterPipelinesSortByEnumValues() []ListOpensearchClusterPipelinesSortByEnum {
	values := make([]ListOpensearchClusterPipelinesSortByEnum, 0)
	for _, v := range mappingListOpensearchClusterPipelinesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpensearchClusterPipelinesSortByEnumStringValues Enumerates the set of values in String for ListOpensearchClusterPipelinesSortByEnum
func GetListOpensearchClusterPipelinesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOpensearchClusterPipelinesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpensearchClusterPipelinesSortByEnum(val string) (ListOpensearchClusterPipelinesSortByEnum, bool) {
	enum, ok := mappingListOpensearchClusterPipelinesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
