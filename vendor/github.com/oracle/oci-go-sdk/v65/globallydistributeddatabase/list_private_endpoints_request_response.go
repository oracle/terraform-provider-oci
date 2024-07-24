// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package globallydistributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPrivateEndpointsRequest wrapper for the ListPrivateEndpoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ListPrivateEndpoints.go.html to see an example of how to use ListPrivateEndpointsRequest.
type ListPrivateEndpointsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState PrivateEndpointLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only private endpoint that match the entire name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivateEndpointLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPrivateEndpointsResponse wrapper for the ListPrivateEndpoints operation
type ListPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PrivateEndpointCollection instances
	PrivateEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListPrivateEndpointsSortOrderEnum
const (
	ListPrivateEndpointsSortOrderAsc  ListPrivateEndpointsSortOrderEnum = "ASC"
	ListPrivateEndpointsSortOrderDesc ListPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListPrivateEndpointsSortOrderEnum = map[string]ListPrivateEndpointsSortOrderEnum{
	"ASC":  ListPrivateEndpointsSortOrderAsc,
	"DESC": ListPrivateEndpointsSortOrderDesc,
}

var mappingListPrivateEndpointsSortOrderEnumLowerCase = map[string]ListPrivateEndpointsSortOrderEnum{
	"asc":  ListPrivateEndpointsSortOrderAsc,
	"desc": ListPrivateEndpointsSortOrderDesc,
}

// GetListPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListPrivateEndpointsSortOrderEnum
func GetListPrivateEndpointsSortOrderEnumValues() []ListPrivateEndpointsSortOrderEnum {
	values := make([]ListPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListPrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListPrivateEndpointsSortOrderEnum
func GetListPrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivateEndpointsSortOrderEnum(val string) (ListPrivateEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListPrivateEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPrivateEndpointsSortByEnum Enum with underlying type: string
type ListPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListPrivateEndpointsSortByEnum
const (
	ListPrivateEndpointsSortByTimecreated ListPrivateEndpointsSortByEnum = "timeCreated"
	ListPrivateEndpointsSortByTimeupdated ListPrivateEndpointsSortByEnum = "timeUpdated"
)

var mappingListPrivateEndpointsSortByEnum = map[string]ListPrivateEndpointsSortByEnum{
	"timeCreated": ListPrivateEndpointsSortByTimecreated,
	"timeUpdated": ListPrivateEndpointsSortByTimeupdated,
}

var mappingListPrivateEndpointsSortByEnumLowerCase = map[string]ListPrivateEndpointsSortByEnum{
	"timecreated": ListPrivateEndpointsSortByTimecreated,
	"timeupdated": ListPrivateEndpointsSortByTimeupdated,
}

// GetListPrivateEndpointsSortByEnumValues Enumerates the set of values for ListPrivateEndpointsSortByEnum
func GetListPrivateEndpointsSortByEnumValues() []ListPrivateEndpointsSortByEnum {
	values := make([]ListPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListPrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListPrivateEndpointsSortByEnum
func GetListPrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListPrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivateEndpointsSortByEnum(val string) (ListPrivateEndpointsSortByEnum, bool) {
	enum, ok := mappingListPrivateEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
