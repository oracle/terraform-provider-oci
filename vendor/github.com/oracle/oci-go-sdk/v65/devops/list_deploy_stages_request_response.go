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

// ListDeployStagesRequest wrapper for the ListDeployStages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployStages.go.html to see an example of how to use ListDeployStagesRequest.
type ListDeployStagesRequest struct {

	// Unique identifier or OCID for listing a single resource by ID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The ID of the parent pipeline.
	DeployPipelineId *string `mandatory:"false" contributesTo:"query" name:"deployPipelineId"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only deployment stages that matches the given lifecycle state.
	LifecycleState DeployStageLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListDeployStagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for display name is ascending. If no value is specified, then the default time created value is considered.
	SortBy ListDeployStagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeployStagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeployStagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeployStagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeployStagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeployStagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeployStageLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDeployStageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeployStagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeployStagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeployStagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeployStagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeployStagesResponse wrapper for the ListDeployStages operation
type ListDeployStagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeployStageCollection instances
	DeployStageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeployStagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeployStagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeployStagesSortOrderEnum Enum with underlying type: string
type ListDeployStagesSortOrderEnum string

// Set of constants representing the allowable values for ListDeployStagesSortOrderEnum
const (
	ListDeployStagesSortOrderAsc  ListDeployStagesSortOrderEnum = "ASC"
	ListDeployStagesSortOrderDesc ListDeployStagesSortOrderEnum = "DESC"
)

var mappingListDeployStagesSortOrderEnum = map[string]ListDeployStagesSortOrderEnum{
	"ASC":  ListDeployStagesSortOrderAsc,
	"DESC": ListDeployStagesSortOrderDesc,
}

var mappingListDeployStagesSortOrderEnumLowerCase = map[string]ListDeployStagesSortOrderEnum{
	"asc":  ListDeployStagesSortOrderAsc,
	"desc": ListDeployStagesSortOrderDesc,
}

// GetListDeployStagesSortOrderEnumValues Enumerates the set of values for ListDeployStagesSortOrderEnum
func GetListDeployStagesSortOrderEnumValues() []ListDeployStagesSortOrderEnum {
	values := make([]ListDeployStagesSortOrderEnum, 0)
	for _, v := range mappingListDeployStagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeployStagesSortOrderEnumStringValues Enumerates the set of values in String for ListDeployStagesSortOrderEnum
func GetListDeployStagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeployStagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeployStagesSortOrderEnum(val string) (ListDeployStagesSortOrderEnum, bool) {
	enum, ok := mappingListDeployStagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeployStagesSortByEnum Enum with underlying type: string
type ListDeployStagesSortByEnum string

// Set of constants representing the allowable values for ListDeployStagesSortByEnum
const (
	ListDeployStagesSortByTimecreated ListDeployStagesSortByEnum = "timeCreated"
	ListDeployStagesSortByDisplayname ListDeployStagesSortByEnum = "displayName"
)

var mappingListDeployStagesSortByEnum = map[string]ListDeployStagesSortByEnum{
	"timeCreated": ListDeployStagesSortByTimecreated,
	"displayName": ListDeployStagesSortByDisplayname,
}

var mappingListDeployStagesSortByEnumLowerCase = map[string]ListDeployStagesSortByEnum{
	"timecreated": ListDeployStagesSortByTimecreated,
	"displayname": ListDeployStagesSortByDisplayname,
}

// GetListDeployStagesSortByEnumValues Enumerates the set of values for ListDeployStagesSortByEnum
func GetListDeployStagesSortByEnumValues() []ListDeployStagesSortByEnum {
	values := make([]ListDeployStagesSortByEnum, 0)
	for _, v := range mappingListDeployStagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeployStagesSortByEnumStringValues Enumerates the set of values in String for ListDeployStagesSortByEnum
func GetListDeployStagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDeployStagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeployStagesSortByEnum(val string) (ListDeployStagesSortByEnum, bool) {
	enum, ok := mappingListDeployStagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
