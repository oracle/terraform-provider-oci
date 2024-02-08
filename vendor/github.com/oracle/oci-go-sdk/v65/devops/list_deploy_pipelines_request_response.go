// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDeployPipelinesRequest wrapper for the ListDeployPipelines operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployPipelines.go.html to see an example of how to use ListDeployPipelinesRequest.
type ListDeployPipelinesRequest struct {

	// Unique identifier or OCID for listing a single resource by ID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// unique project identifier
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only DeployPipelines that matches the given lifecycleState.
	LifecycleState DeployPipelineLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListDeployPipelinesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for display name is ascending. If no value is specified, then the default time created value is considered.
	SortBy ListDeployPipelinesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeployPipelinesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeployPipelinesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeployPipelinesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeployPipelinesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeployPipelinesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeployPipelineLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDeployPipelineLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeployPipelinesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeployPipelinesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeployPipelinesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeployPipelinesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeployPipelinesResponse wrapper for the ListDeployPipelines operation
type ListDeployPipelinesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeployPipelineCollection instances
	DeployPipelineCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeployPipelinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeployPipelinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeployPipelinesSortOrderEnum Enum with underlying type: string
type ListDeployPipelinesSortOrderEnum string

// Set of constants representing the allowable values for ListDeployPipelinesSortOrderEnum
const (
	ListDeployPipelinesSortOrderAsc  ListDeployPipelinesSortOrderEnum = "ASC"
	ListDeployPipelinesSortOrderDesc ListDeployPipelinesSortOrderEnum = "DESC"
)

var mappingListDeployPipelinesSortOrderEnum = map[string]ListDeployPipelinesSortOrderEnum{
	"ASC":  ListDeployPipelinesSortOrderAsc,
	"DESC": ListDeployPipelinesSortOrderDesc,
}

var mappingListDeployPipelinesSortOrderEnumLowerCase = map[string]ListDeployPipelinesSortOrderEnum{
	"asc":  ListDeployPipelinesSortOrderAsc,
	"desc": ListDeployPipelinesSortOrderDesc,
}

// GetListDeployPipelinesSortOrderEnumValues Enumerates the set of values for ListDeployPipelinesSortOrderEnum
func GetListDeployPipelinesSortOrderEnumValues() []ListDeployPipelinesSortOrderEnum {
	values := make([]ListDeployPipelinesSortOrderEnum, 0)
	for _, v := range mappingListDeployPipelinesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeployPipelinesSortOrderEnumStringValues Enumerates the set of values in String for ListDeployPipelinesSortOrderEnum
func GetListDeployPipelinesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeployPipelinesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeployPipelinesSortOrderEnum(val string) (ListDeployPipelinesSortOrderEnum, bool) {
	enum, ok := mappingListDeployPipelinesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeployPipelinesSortByEnum Enum with underlying type: string
type ListDeployPipelinesSortByEnum string

// Set of constants representing the allowable values for ListDeployPipelinesSortByEnum
const (
	ListDeployPipelinesSortByTimecreated ListDeployPipelinesSortByEnum = "timeCreated"
	ListDeployPipelinesSortByDisplayname ListDeployPipelinesSortByEnum = "displayName"
)

var mappingListDeployPipelinesSortByEnum = map[string]ListDeployPipelinesSortByEnum{
	"timeCreated": ListDeployPipelinesSortByTimecreated,
	"displayName": ListDeployPipelinesSortByDisplayname,
}

var mappingListDeployPipelinesSortByEnumLowerCase = map[string]ListDeployPipelinesSortByEnum{
	"timecreated": ListDeployPipelinesSortByTimecreated,
	"displayname": ListDeployPipelinesSortByDisplayname,
}

// GetListDeployPipelinesSortByEnumValues Enumerates the set of values for ListDeployPipelinesSortByEnum
func GetListDeployPipelinesSortByEnumValues() []ListDeployPipelinesSortByEnum {
	values := make([]ListDeployPipelinesSortByEnum, 0)
	for _, v := range mappingListDeployPipelinesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeployPipelinesSortByEnumStringValues Enumerates the set of values in String for ListDeployPipelinesSortByEnum
func GetListDeployPipelinesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDeployPipelinesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeployPipelinesSortByEnum(val string) (ListDeployPipelinesSortByEnum, bool) {
	enum, ok := mappingListDeployPipelinesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
