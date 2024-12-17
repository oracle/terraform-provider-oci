// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPipelinesRequest wrapper for the ListPipelines operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListPipelines.go.html to see an example of how to use ListPipelinesRequest.
type ListPipelinesRequest struct {

	// The OCID of the compartment that contains the work request. Work requests should be scoped
	// to the same compartment as the resource the work request affects. If the work request concerns
	// multiple resources, and those resources are not in the same compartment, it is up to the service team
	// to pick the primary resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filtered list of pipelines to return for a given lifecycleState.
	LifecycleState PipelineLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filtered list of pipelines to return for a given lifecycleSubState.
	LifecycleSubState ListPipelinesLifecycleSubStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleSubState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListPipelinesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListPipelinesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

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
	if _, ok := GetMappingPipelineLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPipelineLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPipelinesLifecycleSubStateEnum(string(request.LifecycleSubState)); !ok && request.LifecycleSubState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubState: %s. Supported values are: %s.", request.LifecycleSubState, strings.Join(GetListPipelinesLifecycleSubStateEnumStringValues(), ",")))
	}
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

	// A list of PipelineCollection instances
	PipelineCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPipelinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPipelinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPipelinesLifecycleSubStateEnum Enum with underlying type: string
type ListPipelinesLifecycleSubStateEnum string

// Set of constants representing the allowable values for ListPipelinesLifecycleSubStateEnum
const (
	ListPipelinesLifecycleSubStateStarting ListPipelinesLifecycleSubStateEnum = "STARTING"
	ListPipelinesLifecycleSubStateStopping ListPipelinesLifecycleSubStateEnum = "STOPPING"
	ListPipelinesLifecycleSubStateStopped  ListPipelinesLifecycleSubStateEnum = "STOPPED"
	ListPipelinesLifecycleSubStateMoving   ListPipelinesLifecycleSubStateEnum = "MOVING"
	ListPipelinesLifecycleSubStateRunning  ListPipelinesLifecycleSubStateEnum = "RUNNING"
)

var mappingListPipelinesLifecycleSubStateEnum = map[string]ListPipelinesLifecycleSubStateEnum{
	"STARTING": ListPipelinesLifecycleSubStateStarting,
	"STOPPING": ListPipelinesLifecycleSubStateStopping,
	"STOPPED":  ListPipelinesLifecycleSubStateStopped,
	"MOVING":   ListPipelinesLifecycleSubStateMoving,
	"RUNNING":  ListPipelinesLifecycleSubStateRunning,
}

var mappingListPipelinesLifecycleSubStateEnumLowerCase = map[string]ListPipelinesLifecycleSubStateEnum{
	"starting": ListPipelinesLifecycleSubStateStarting,
	"stopping": ListPipelinesLifecycleSubStateStopping,
	"stopped":  ListPipelinesLifecycleSubStateStopped,
	"moving":   ListPipelinesLifecycleSubStateMoving,
	"running":  ListPipelinesLifecycleSubStateRunning,
}

// GetListPipelinesLifecycleSubStateEnumValues Enumerates the set of values for ListPipelinesLifecycleSubStateEnum
func GetListPipelinesLifecycleSubStateEnumValues() []ListPipelinesLifecycleSubStateEnum {
	values := make([]ListPipelinesLifecycleSubStateEnum, 0)
	for _, v := range mappingListPipelinesLifecycleSubStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelinesLifecycleSubStateEnumStringValues Enumerates the set of values in String for ListPipelinesLifecycleSubStateEnum
func GetListPipelinesLifecycleSubStateEnumStringValues() []string {
	return []string{
		"STARTING",
		"STOPPING",
		"STOPPED",
		"MOVING",
		"RUNNING",
	}
}

// GetMappingListPipelinesLifecycleSubStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelinesLifecycleSubStateEnum(val string) (ListPipelinesLifecycleSubStateEnum, bool) {
	enum, ok := mappingListPipelinesLifecycleSubStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingListPipelinesSortOrderEnumLowerCase = map[string]ListPipelinesSortOrderEnum{
	"asc":  ListPipelinesSortOrderAsc,
	"desc": ListPipelinesSortOrderDesc,
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
	enum, ok := mappingListPipelinesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPipelinesSortByEnum Enum with underlying type: string
type ListPipelinesSortByEnum string

// Set of constants representing the allowable values for ListPipelinesSortByEnum
const (
	ListPipelinesSortByTimecreated ListPipelinesSortByEnum = "timeCreated"
	ListPipelinesSortByDisplayname ListPipelinesSortByEnum = "displayName"
)

var mappingListPipelinesSortByEnum = map[string]ListPipelinesSortByEnum{
	"timeCreated": ListPipelinesSortByTimecreated,
	"displayName": ListPipelinesSortByDisplayname,
}

var mappingListPipelinesSortByEnumLowerCase = map[string]ListPipelinesSortByEnum{
	"timecreated": ListPipelinesSortByTimecreated,
	"displayname": ListPipelinesSortByDisplayname,
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
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPipelinesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelinesSortByEnum(val string) (ListPipelinesSortByEnum, bool) {
	enum, ok := mappingListPipelinesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
