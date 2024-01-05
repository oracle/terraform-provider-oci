// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ailanguage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEndpointsRequest wrapper for the ListEndpoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/ListEndpoints.go.html to see an example of how to use ListEndpointsRequest.
type ListEndpointsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID of the endpoint.
	EndpointId *string `mandatory:"false" contributesTo:"query" name:"endpointId"`

	// The ID of the project for which to list the objects.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the trained model for which to list the endpoints.
	ModelId *string `mandatory:"false" contributesTo:"query" name:"modelId"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState EndpointLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown
	// in descending order. When you sort by `displayName`, the results are
	// shown in ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEndpointLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEndpointsResponse wrapper for the ListEndpoints operation
type ListEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EndpointCollection instances
	EndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEndpointsSortOrderEnum Enum with underlying type: string
type ListEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListEndpointsSortOrderEnum
const (
	ListEndpointsSortOrderAsc  ListEndpointsSortOrderEnum = "ASC"
	ListEndpointsSortOrderDesc ListEndpointsSortOrderEnum = "DESC"
)

var mappingListEndpointsSortOrderEnum = map[string]ListEndpointsSortOrderEnum{
	"ASC":  ListEndpointsSortOrderAsc,
	"DESC": ListEndpointsSortOrderDesc,
}

var mappingListEndpointsSortOrderEnumLowerCase = map[string]ListEndpointsSortOrderEnum{
	"asc":  ListEndpointsSortOrderAsc,
	"desc": ListEndpointsSortOrderDesc,
}

// GetListEndpointsSortOrderEnumValues Enumerates the set of values for ListEndpointsSortOrderEnum
func GetListEndpointsSortOrderEnumValues() []ListEndpointsSortOrderEnum {
	values := make([]ListEndpointsSortOrderEnum, 0)
	for _, v := range mappingListEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListEndpointsSortOrderEnum
func GetListEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEndpointsSortOrderEnum(val string) (ListEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEndpointsSortByEnum Enum with underlying type: string
type ListEndpointsSortByEnum string

// Set of constants representing the allowable values for ListEndpointsSortByEnum
const (
	ListEndpointsSortByTimecreated ListEndpointsSortByEnum = "timeCreated"
	ListEndpointsSortByDisplayname ListEndpointsSortByEnum = "displayName"
)

var mappingListEndpointsSortByEnum = map[string]ListEndpointsSortByEnum{
	"timeCreated": ListEndpointsSortByTimecreated,
	"displayName": ListEndpointsSortByDisplayname,
}

var mappingListEndpointsSortByEnumLowerCase = map[string]ListEndpointsSortByEnum{
	"timecreated": ListEndpointsSortByTimecreated,
	"displayname": ListEndpointsSortByDisplayname,
}

// GetListEndpointsSortByEnumValues Enumerates the set of values for ListEndpointsSortByEnum
func GetListEndpointsSortByEnumValues() []ListEndpointsSortByEnum {
	values := make([]ListEndpointsSortByEnum, 0)
	for _, v := range mappingListEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEndpointsSortByEnumStringValues Enumerates the set of values in String for ListEndpointsSortByEnum
func GetListEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEndpointsSortByEnum(val string) (ListEndpointsSortByEnum, bool) {
	enum, ok := mappingListEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
