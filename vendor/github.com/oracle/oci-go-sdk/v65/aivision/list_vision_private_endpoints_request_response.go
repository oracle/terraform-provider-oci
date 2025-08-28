// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package aivision

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListVisionPrivateEndpointsRequest wrapper for the ListVisionPrivateEndpoints operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aivision/ListVisionPrivateEndpoints.go.html to see an example of how to use ListVisionPrivateEndpointsRequest.
type ListVisionPrivateEndpointsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The filter to match projects with the given lifecycleState.
	LifecycleState VisionPrivateEndpointLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The filter to find the device with the given identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListVisionPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for timeCreated is descending. The default order for displayName is ascending.
	SortBy ListVisionPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVisionPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVisionPrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVisionPrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVisionPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVisionPrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVisionPrivateEndpointLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetVisionPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVisionPrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVisionPrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVisionPrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVisionPrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVisionPrivateEndpointsResponse wrapper for the ListVisionPrivateEndpoints operation
type ListVisionPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VisionPrivateEndpointCollection instances
	VisionPrivateEndpointCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVisionPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVisionPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVisionPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListVisionPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListVisionPrivateEndpointsSortOrderEnum
const (
	ListVisionPrivateEndpointsSortOrderAsc  ListVisionPrivateEndpointsSortOrderEnum = "ASC"
	ListVisionPrivateEndpointsSortOrderDesc ListVisionPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListVisionPrivateEndpointsSortOrderEnum = map[string]ListVisionPrivateEndpointsSortOrderEnum{
	"ASC":  ListVisionPrivateEndpointsSortOrderAsc,
	"DESC": ListVisionPrivateEndpointsSortOrderDesc,
}

var mappingListVisionPrivateEndpointsSortOrderEnumLowerCase = map[string]ListVisionPrivateEndpointsSortOrderEnum{
	"asc":  ListVisionPrivateEndpointsSortOrderAsc,
	"desc": ListVisionPrivateEndpointsSortOrderDesc,
}

// GetListVisionPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListVisionPrivateEndpointsSortOrderEnum
func GetListVisionPrivateEndpointsSortOrderEnumValues() []ListVisionPrivateEndpointsSortOrderEnum {
	values := make([]ListVisionPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListVisionPrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVisionPrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListVisionPrivateEndpointsSortOrderEnum
func GetListVisionPrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVisionPrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVisionPrivateEndpointsSortOrderEnum(val string) (ListVisionPrivateEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListVisionPrivateEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVisionPrivateEndpointsSortByEnum Enum with underlying type: string
type ListVisionPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListVisionPrivateEndpointsSortByEnum
const (
	ListVisionPrivateEndpointsSortByTimecreated ListVisionPrivateEndpointsSortByEnum = "timeCreated"
	ListVisionPrivateEndpointsSortByDisplayname ListVisionPrivateEndpointsSortByEnum = "displayName"
)

var mappingListVisionPrivateEndpointsSortByEnum = map[string]ListVisionPrivateEndpointsSortByEnum{
	"timeCreated": ListVisionPrivateEndpointsSortByTimecreated,
	"displayName": ListVisionPrivateEndpointsSortByDisplayname,
}

var mappingListVisionPrivateEndpointsSortByEnumLowerCase = map[string]ListVisionPrivateEndpointsSortByEnum{
	"timecreated": ListVisionPrivateEndpointsSortByTimecreated,
	"displayname": ListVisionPrivateEndpointsSortByDisplayname,
}

// GetListVisionPrivateEndpointsSortByEnumValues Enumerates the set of values for ListVisionPrivateEndpointsSortByEnum
func GetListVisionPrivateEndpointsSortByEnumValues() []ListVisionPrivateEndpointsSortByEnum {
	values := make([]ListVisionPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListVisionPrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVisionPrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListVisionPrivateEndpointsSortByEnum
func GetListVisionPrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListVisionPrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVisionPrivateEndpointsSortByEnum(val string) (ListVisionPrivateEndpointsSortByEnum, bool) {
	enum, ok := mappingListVisionPrivateEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
