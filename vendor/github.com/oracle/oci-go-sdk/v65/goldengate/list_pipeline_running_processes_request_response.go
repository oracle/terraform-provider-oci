// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPipelineRunningProcessesRequest wrapper for the ListPipelineRunningProcesses operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListPipelineRunningProcesses.go.html to see an example of how to use ListPipelineRunningProcessesRequest.
type ListPipelineRunningProcessesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the pipeline created.
	PipelineId *string `mandatory:"true" contributesTo:"path" name:"pipelineId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListPipelineRunningProcessesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListPipelineRunningProcessesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPipelineRunningProcessesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPipelineRunningProcessesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPipelineRunningProcessesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPipelineRunningProcessesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPipelineRunningProcessesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPipelineRunningProcessesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPipelineRunningProcessesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPipelineRunningProcessesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPipelineRunningProcessesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPipelineRunningProcessesResponse wrapper for the ListPipelineRunningProcesses operation
type ListPipelineRunningProcessesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PipelineRunningProcessCollection instances
	PipelineRunningProcessCollection `presentIn:"body"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListPipelineRunningProcessesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPipelineRunningProcessesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPipelineRunningProcessesSortOrderEnum Enum with underlying type: string
type ListPipelineRunningProcessesSortOrderEnum string

// Set of constants representing the allowable values for ListPipelineRunningProcessesSortOrderEnum
const (
	ListPipelineRunningProcessesSortOrderAsc  ListPipelineRunningProcessesSortOrderEnum = "ASC"
	ListPipelineRunningProcessesSortOrderDesc ListPipelineRunningProcessesSortOrderEnum = "DESC"
)

var mappingListPipelineRunningProcessesSortOrderEnum = map[string]ListPipelineRunningProcessesSortOrderEnum{
	"ASC":  ListPipelineRunningProcessesSortOrderAsc,
	"DESC": ListPipelineRunningProcessesSortOrderDesc,
}

var mappingListPipelineRunningProcessesSortOrderEnumLowerCase = map[string]ListPipelineRunningProcessesSortOrderEnum{
	"asc":  ListPipelineRunningProcessesSortOrderAsc,
	"desc": ListPipelineRunningProcessesSortOrderDesc,
}

// GetListPipelineRunningProcessesSortOrderEnumValues Enumerates the set of values for ListPipelineRunningProcessesSortOrderEnum
func GetListPipelineRunningProcessesSortOrderEnumValues() []ListPipelineRunningProcessesSortOrderEnum {
	values := make([]ListPipelineRunningProcessesSortOrderEnum, 0)
	for _, v := range mappingListPipelineRunningProcessesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelineRunningProcessesSortOrderEnumStringValues Enumerates the set of values in String for ListPipelineRunningProcessesSortOrderEnum
func GetListPipelineRunningProcessesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPipelineRunningProcessesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelineRunningProcessesSortOrderEnum(val string) (ListPipelineRunningProcessesSortOrderEnum, bool) {
	enum, ok := mappingListPipelineRunningProcessesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPipelineRunningProcessesSortByEnum Enum with underlying type: string
type ListPipelineRunningProcessesSortByEnum string

// Set of constants representing the allowable values for ListPipelineRunningProcessesSortByEnum
const (
	ListPipelineRunningProcessesSortByTimecreated ListPipelineRunningProcessesSortByEnum = "timeCreated"
	ListPipelineRunningProcessesSortByDisplayname ListPipelineRunningProcessesSortByEnum = "displayName"
)

var mappingListPipelineRunningProcessesSortByEnum = map[string]ListPipelineRunningProcessesSortByEnum{
	"timeCreated": ListPipelineRunningProcessesSortByTimecreated,
	"displayName": ListPipelineRunningProcessesSortByDisplayname,
}

var mappingListPipelineRunningProcessesSortByEnumLowerCase = map[string]ListPipelineRunningProcessesSortByEnum{
	"timecreated": ListPipelineRunningProcessesSortByTimecreated,
	"displayname": ListPipelineRunningProcessesSortByDisplayname,
}

// GetListPipelineRunningProcessesSortByEnumValues Enumerates the set of values for ListPipelineRunningProcessesSortByEnum
func GetListPipelineRunningProcessesSortByEnumValues() []ListPipelineRunningProcessesSortByEnum {
	values := make([]ListPipelineRunningProcessesSortByEnum, 0)
	for _, v := range mappingListPipelineRunningProcessesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelineRunningProcessesSortByEnumStringValues Enumerates the set of values in String for ListPipelineRunningProcessesSortByEnum
func GetListPipelineRunningProcessesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPipelineRunningProcessesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelineRunningProcessesSortByEnum(val string) (ListPipelineRunningProcessesSortByEnum, bool) {
	enum, ok := mappingListPipelineRunningProcessesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
