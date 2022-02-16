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

// ListBuildPipelineStagesRequest wrapper for the ListBuildPipelineStages operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListBuildPipelineStages.go.html to see an example of how to use ListBuildPipelineStagesRequest.
type ListBuildPipelineStagesRequest struct {

	// Unique identifier or OCID for listing a single resource by ID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID of the parent build pipeline.
	BuildPipelineId *string `mandatory:"false" contributesTo:"query" name:"buildPipelineId"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return the stages that matches the given lifecycle state.
	LifecycleState BuildPipelineStageLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListBuildPipelineStagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for display name is ascending. If no value is specified, then the default time created value is considered.
	SortBy ListBuildPipelineStagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBuildPipelineStagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBuildPipelineStagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBuildPipelineStagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBuildPipelineStagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBuildPipelineStagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBuildPipelineStageLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBuildPipelineStageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBuildPipelineStagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBuildPipelineStagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBuildPipelineStagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBuildPipelineStagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBuildPipelineStagesResponse wrapper for the ListBuildPipelineStages operation
type ListBuildPipelineStagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BuildPipelineStageCollection instances
	BuildPipelineStageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBuildPipelineStagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBuildPipelineStagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBuildPipelineStagesSortOrderEnum Enum with underlying type: string
type ListBuildPipelineStagesSortOrderEnum string

// Set of constants representing the allowable values for ListBuildPipelineStagesSortOrderEnum
const (
	ListBuildPipelineStagesSortOrderAsc  ListBuildPipelineStagesSortOrderEnum = "ASC"
	ListBuildPipelineStagesSortOrderDesc ListBuildPipelineStagesSortOrderEnum = "DESC"
)

var mappingListBuildPipelineStagesSortOrderEnum = map[string]ListBuildPipelineStagesSortOrderEnum{
	"ASC":  ListBuildPipelineStagesSortOrderAsc,
	"DESC": ListBuildPipelineStagesSortOrderDesc,
}

// GetListBuildPipelineStagesSortOrderEnumValues Enumerates the set of values for ListBuildPipelineStagesSortOrderEnum
func GetListBuildPipelineStagesSortOrderEnumValues() []ListBuildPipelineStagesSortOrderEnum {
	values := make([]ListBuildPipelineStagesSortOrderEnum, 0)
	for _, v := range mappingListBuildPipelineStagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBuildPipelineStagesSortOrderEnumStringValues Enumerates the set of values in String for ListBuildPipelineStagesSortOrderEnum
func GetListBuildPipelineStagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBuildPipelineStagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBuildPipelineStagesSortOrderEnum(val string) (ListBuildPipelineStagesSortOrderEnum, bool) {
	mappingListBuildPipelineStagesSortOrderEnumIgnoreCase := make(map[string]ListBuildPipelineStagesSortOrderEnum)
	for k, v := range mappingListBuildPipelineStagesSortOrderEnum {
		mappingListBuildPipelineStagesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListBuildPipelineStagesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListBuildPipelineStagesSortByEnum Enum with underlying type: string
type ListBuildPipelineStagesSortByEnum string

// Set of constants representing the allowable values for ListBuildPipelineStagesSortByEnum
const (
	ListBuildPipelineStagesSortByTimecreated ListBuildPipelineStagesSortByEnum = "timeCreated"
	ListBuildPipelineStagesSortByDisplayname ListBuildPipelineStagesSortByEnum = "displayName"
)

var mappingListBuildPipelineStagesSortByEnum = map[string]ListBuildPipelineStagesSortByEnum{
	"timeCreated": ListBuildPipelineStagesSortByTimecreated,
	"displayName": ListBuildPipelineStagesSortByDisplayname,
}

// GetListBuildPipelineStagesSortByEnumValues Enumerates the set of values for ListBuildPipelineStagesSortByEnum
func GetListBuildPipelineStagesSortByEnumValues() []ListBuildPipelineStagesSortByEnum {
	values := make([]ListBuildPipelineStagesSortByEnum, 0)
	for _, v := range mappingListBuildPipelineStagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBuildPipelineStagesSortByEnumStringValues Enumerates the set of values in String for ListBuildPipelineStagesSortByEnum
func GetListBuildPipelineStagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBuildPipelineStagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBuildPipelineStagesSortByEnum(val string) (ListBuildPipelineStagesSortByEnum, bool) {
	mappingListBuildPipelineStagesSortByEnumIgnoreCase := make(map[string]ListBuildPipelineStagesSortByEnum)
	for k, v := range mappingListBuildPipelineStagesSortByEnum {
		mappingListBuildPipelineStagesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListBuildPipelineStagesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
