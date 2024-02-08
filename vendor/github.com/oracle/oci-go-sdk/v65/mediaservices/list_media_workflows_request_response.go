// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMediaWorkflowsRequest wrapper for the ListMediaWorkflows operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaWorkflows.go.html to see an example of how to use ListMediaWorkflowsRequest.
type ListMediaWorkflowsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique MediaWorkflow identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only the resources with lifecycleState matching the given lifecycleState.
	LifecycleState MediaWorkflowLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMediaWorkflowsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default
	// order for displayName is ascending.
	SortBy ListMediaWorkflowsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A token representing the position at which to start retrieving results. This must come from the
	// `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMediaWorkflowsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMediaWorkflowsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMediaWorkflowsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMediaWorkflowsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMediaWorkflowsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMediaWorkflowLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMediaWorkflowLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaWorkflowsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMediaWorkflowsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaWorkflowsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMediaWorkflowsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMediaWorkflowsResponse wrapper for the ListMediaWorkflows operation
type ListMediaWorkflowsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MediaWorkflowCollection instances
	MediaWorkflowCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMediaWorkflowsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMediaWorkflowsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMediaWorkflowsSortOrderEnum Enum with underlying type: string
type ListMediaWorkflowsSortOrderEnum string

// Set of constants representing the allowable values for ListMediaWorkflowsSortOrderEnum
const (
	ListMediaWorkflowsSortOrderAsc  ListMediaWorkflowsSortOrderEnum = "ASC"
	ListMediaWorkflowsSortOrderDesc ListMediaWorkflowsSortOrderEnum = "DESC"
)

var mappingListMediaWorkflowsSortOrderEnum = map[string]ListMediaWorkflowsSortOrderEnum{
	"ASC":  ListMediaWorkflowsSortOrderAsc,
	"DESC": ListMediaWorkflowsSortOrderDesc,
}

var mappingListMediaWorkflowsSortOrderEnumLowerCase = map[string]ListMediaWorkflowsSortOrderEnum{
	"asc":  ListMediaWorkflowsSortOrderAsc,
	"desc": ListMediaWorkflowsSortOrderDesc,
}

// GetListMediaWorkflowsSortOrderEnumValues Enumerates the set of values for ListMediaWorkflowsSortOrderEnum
func GetListMediaWorkflowsSortOrderEnumValues() []ListMediaWorkflowsSortOrderEnum {
	values := make([]ListMediaWorkflowsSortOrderEnum, 0)
	for _, v := range mappingListMediaWorkflowsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaWorkflowsSortOrderEnumStringValues Enumerates the set of values in String for ListMediaWorkflowsSortOrderEnum
func GetListMediaWorkflowsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMediaWorkflowsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaWorkflowsSortOrderEnum(val string) (ListMediaWorkflowsSortOrderEnum, bool) {
	enum, ok := mappingListMediaWorkflowsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMediaWorkflowsSortByEnum Enum with underlying type: string
type ListMediaWorkflowsSortByEnum string

// Set of constants representing the allowable values for ListMediaWorkflowsSortByEnum
const (
	ListMediaWorkflowsSortByTimecreated ListMediaWorkflowsSortByEnum = "timeCreated"
	ListMediaWorkflowsSortByDisplayname ListMediaWorkflowsSortByEnum = "displayName"
)

var mappingListMediaWorkflowsSortByEnum = map[string]ListMediaWorkflowsSortByEnum{
	"timeCreated": ListMediaWorkflowsSortByTimecreated,
	"displayName": ListMediaWorkflowsSortByDisplayname,
}

var mappingListMediaWorkflowsSortByEnumLowerCase = map[string]ListMediaWorkflowsSortByEnum{
	"timecreated": ListMediaWorkflowsSortByTimecreated,
	"displayname": ListMediaWorkflowsSortByDisplayname,
}

// GetListMediaWorkflowsSortByEnumValues Enumerates the set of values for ListMediaWorkflowsSortByEnum
func GetListMediaWorkflowsSortByEnumValues() []ListMediaWorkflowsSortByEnum {
	values := make([]ListMediaWorkflowsSortByEnum, 0)
	for _, v := range mappingListMediaWorkflowsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaWorkflowsSortByEnumStringValues Enumerates the set of values in String for ListMediaWorkflowsSortByEnum
func GetListMediaWorkflowsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMediaWorkflowsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaWorkflowsSortByEnum(val string) (ListMediaWorkflowsSortByEnum, bool) {
	enum, ok := mappingListMediaWorkflowsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
