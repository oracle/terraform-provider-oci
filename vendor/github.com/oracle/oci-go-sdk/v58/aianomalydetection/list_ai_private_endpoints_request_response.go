// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package aianomalydetection

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListAiPrivateEndpointsRequest wrapper for the ListAiPrivateEndpoints operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aianomalydetection/ListAiPrivateEndpoints.go.html to see an example of how to use ListAiPrivateEndpointsRequest.
type ListAiPrivateEndpointsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState AiPrivateEndpointLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique AiPrivateEndpoint identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAiPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListAiPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAiPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAiPrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAiPrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAiPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAiPrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAiPrivateEndpointLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAiPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAiPrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAiPrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAiPrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAiPrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAiPrivateEndpointsResponse wrapper for the ListAiPrivateEndpoints operation
type ListAiPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AiPrivateEndpointCollection instances
	AiPrivateEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAiPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAiPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAiPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListAiPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListAiPrivateEndpointsSortOrderEnum
const (
	ListAiPrivateEndpointsSortOrderAsc  ListAiPrivateEndpointsSortOrderEnum = "ASC"
	ListAiPrivateEndpointsSortOrderDesc ListAiPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListAiPrivateEndpointsSortOrderEnum = map[string]ListAiPrivateEndpointsSortOrderEnum{
	"ASC":  ListAiPrivateEndpointsSortOrderAsc,
	"DESC": ListAiPrivateEndpointsSortOrderDesc,
}

// GetListAiPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListAiPrivateEndpointsSortOrderEnum
func GetListAiPrivateEndpointsSortOrderEnumValues() []ListAiPrivateEndpointsSortOrderEnum {
	values := make([]ListAiPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListAiPrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAiPrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListAiPrivateEndpointsSortOrderEnum
func GetListAiPrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAiPrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAiPrivateEndpointsSortOrderEnum(val string) (ListAiPrivateEndpointsSortOrderEnum, bool) {
	mappingListAiPrivateEndpointsSortOrderEnumIgnoreCase := make(map[string]ListAiPrivateEndpointsSortOrderEnum)
	for k, v := range mappingListAiPrivateEndpointsSortOrderEnum {
		mappingListAiPrivateEndpointsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAiPrivateEndpointsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAiPrivateEndpointsSortByEnum Enum with underlying type: string
type ListAiPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListAiPrivateEndpointsSortByEnum
const (
	ListAiPrivateEndpointsSortByTimecreated ListAiPrivateEndpointsSortByEnum = "timeCreated"
	ListAiPrivateEndpointsSortByDisplayname ListAiPrivateEndpointsSortByEnum = "displayName"
)

var mappingListAiPrivateEndpointsSortByEnum = map[string]ListAiPrivateEndpointsSortByEnum{
	"timeCreated": ListAiPrivateEndpointsSortByTimecreated,
	"displayName": ListAiPrivateEndpointsSortByDisplayname,
}

// GetListAiPrivateEndpointsSortByEnumValues Enumerates the set of values for ListAiPrivateEndpointsSortByEnum
func GetListAiPrivateEndpointsSortByEnumValues() []ListAiPrivateEndpointsSortByEnum {
	values := make([]ListAiPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListAiPrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAiPrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListAiPrivateEndpointsSortByEnum
func GetListAiPrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAiPrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAiPrivateEndpointsSortByEnum(val string) (ListAiPrivateEndpointsSortByEnum, bool) {
	mappingListAiPrivateEndpointsSortByEnumIgnoreCase := make(map[string]ListAiPrivateEndpointsSortByEnum)
	for k, v := range mappingListAiPrivateEndpointsSortByEnum {
		mappingListAiPrivateEndpointsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAiPrivateEndpointsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
