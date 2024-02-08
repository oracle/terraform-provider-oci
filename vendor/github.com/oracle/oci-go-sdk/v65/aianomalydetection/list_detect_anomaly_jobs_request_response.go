// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package aianomalydetection

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDetectAnomalyJobsRequest wrapper for the ListDetectAnomalyJobs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aianomalydetection/ListDetectAnomalyJobs.go.html to see an example of how to use ListDetectAnomalyJobsRequest.
type ListDetectAnomalyJobsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the trained model for which to list the resources.
	ModelId *string `mandatory:"false" contributesTo:"query" name:"modelId"`

	// The ID of the project for which to list the objects.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// Unique Async Job identifier
	DetectAnomalyJobId *string `mandatory:"false" contributesTo:"query" name:"detectAnomalyJobId"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState DetectAnomalyJobLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDetectAnomalyJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDetectAnomalyJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDetectAnomalyJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDetectAnomalyJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDetectAnomalyJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDetectAnomalyJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDetectAnomalyJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDetectAnomalyJobLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDetectAnomalyJobLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDetectAnomalyJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDetectAnomalyJobsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDetectAnomalyJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDetectAnomalyJobsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDetectAnomalyJobsResponse wrapper for the ListDetectAnomalyJobs operation
type ListDetectAnomalyJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DetectAnomalyJobCollection instances
	DetectAnomalyJobCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// precedent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDetectAnomalyJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDetectAnomalyJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDetectAnomalyJobsSortOrderEnum Enum with underlying type: string
type ListDetectAnomalyJobsSortOrderEnum string

// Set of constants representing the allowable values for ListDetectAnomalyJobsSortOrderEnum
const (
	ListDetectAnomalyJobsSortOrderAsc  ListDetectAnomalyJobsSortOrderEnum = "ASC"
	ListDetectAnomalyJobsSortOrderDesc ListDetectAnomalyJobsSortOrderEnum = "DESC"
)

var mappingListDetectAnomalyJobsSortOrderEnum = map[string]ListDetectAnomalyJobsSortOrderEnum{
	"ASC":  ListDetectAnomalyJobsSortOrderAsc,
	"DESC": ListDetectAnomalyJobsSortOrderDesc,
}

var mappingListDetectAnomalyJobsSortOrderEnumLowerCase = map[string]ListDetectAnomalyJobsSortOrderEnum{
	"asc":  ListDetectAnomalyJobsSortOrderAsc,
	"desc": ListDetectAnomalyJobsSortOrderDesc,
}

// GetListDetectAnomalyJobsSortOrderEnumValues Enumerates the set of values for ListDetectAnomalyJobsSortOrderEnum
func GetListDetectAnomalyJobsSortOrderEnumValues() []ListDetectAnomalyJobsSortOrderEnum {
	values := make([]ListDetectAnomalyJobsSortOrderEnum, 0)
	for _, v := range mappingListDetectAnomalyJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectAnomalyJobsSortOrderEnumStringValues Enumerates the set of values in String for ListDetectAnomalyJobsSortOrderEnum
func GetListDetectAnomalyJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDetectAnomalyJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectAnomalyJobsSortOrderEnum(val string) (ListDetectAnomalyJobsSortOrderEnum, bool) {
	enum, ok := mappingListDetectAnomalyJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDetectAnomalyJobsSortByEnum Enum with underlying type: string
type ListDetectAnomalyJobsSortByEnum string

// Set of constants representing the allowable values for ListDetectAnomalyJobsSortByEnum
const (
	ListDetectAnomalyJobsSortByTimecreated ListDetectAnomalyJobsSortByEnum = "timeCreated"
	ListDetectAnomalyJobsSortByDisplayname ListDetectAnomalyJobsSortByEnum = "displayName"
)

var mappingListDetectAnomalyJobsSortByEnum = map[string]ListDetectAnomalyJobsSortByEnum{
	"timeCreated": ListDetectAnomalyJobsSortByTimecreated,
	"displayName": ListDetectAnomalyJobsSortByDisplayname,
}

var mappingListDetectAnomalyJobsSortByEnumLowerCase = map[string]ListDetectAnomalyJobsSortByEnum{
	"timecreated": ListDetectAnomalyJobsSortByTimecreated,
	"displayname": ListDetectAnomalyJobsSortByDisplayname,
}

// GetListDetectAnomalyJobsSortByEnumValues Enumerates the set of values for ListDetectAnomalyJobsSortByEnum
func GetListDetectAnomalyJobsSortByEnumValues() []ListDetectAnomalyJobsSortByEnum {
	values := make([]ListDetectAnomalyJobsSortByEnum, 0)
	for _, v := range mappingListDetectAnomalyJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectAnomalyJobsSortByEnumStringValues Enumerates the set of values in String for ListDetectAnomalyJobsSortByEnum
func GetListDetectAnomalyJobsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDetectAnomalyJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectAnomalyJobsSortByEnum(val string) (ListDetectAnomalyJobsSortByEnum, bool) {
	enum, ok := mappingListDetectAnomalyJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
