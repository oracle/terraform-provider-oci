// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDistributedDatabasePrivateEndpointsRequest wrapper for the ListDistributedDatabasePrivateEndpoints operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ListDistributedDatabasePrivateEndpoints.go.html to see an example of how to use ListDistributedDatabasePrivateEndpointsRequest.
type ListDistributedDatabasePrivateEndpointsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState DistributedDatabasePrivateEndpointLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListDistributedDatabasePrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListDistributedDatabasePrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only private endpoint that match the entire name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDistributedDatabasePrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDistributedDatabasePrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDistributedDatabasePrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDistributedDatabasePrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDistributedDatabasePrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabasePrivateEndpointLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDistributedDatabasePrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDistributedDatabasePrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDistributedDatabasePrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDistributedDatabasePrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDistributedDatabasePrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDistributedDatabasePrivateEndpointsResponse wrapper for the ListDistributedDatabasePrivateEndpoints operation
type ListDistributedDatabasePrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DistributedDatabasePrivateEndpointCollection instances
	DistributedDatabasePrivateEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDistributedDatabasePrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDistributedDatabasePrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDistributedDatabasePrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListDistributedDatabasePrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListDistributedDatabasePrivateEndpointsSortOrderEnum
const (
	ListDistributedDatabasePrivateEndpointsSortOrderAsc  ListDistributedDatabasePrivateEndpointsSortOrderEnum = "ASC"
	ListDistributedDatabasePrivateEndpointsSortOrderDesc ListDistributedDatabasePrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListDistributedDatabasePrivateEndpointsSortOrderEnum = map[string]ListDistributedDatabasePrivateEndpointsSortOrderEnum{
	"ASC":  ListDistributedDatabasePrivateEndpointsSortOrderAsc,
	"DESC": ListDistributedDatabasePrivateEndpointsSortOrderDesc,
}

var mappingListDistributedDatabasePrivateEndpointsSortOrderEnumLowerCase = map[string]ListDistributedDatabasePrivateEndpointsSortOrderEnum{
	"asc":  ListDistributedDatabasePrivateEndpointsSortOrderAsc,
	"desc": ListDistributedDatabasePrivateEndpointsSortOrderDesc,
}

// GetListDistributedDatabasePrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListDistributedDatabasePrivateEndpointsSortOrderEnum
func GetListDistributedDatabasePrivateEndpointsSortOrderEnumValues() []ListDistributedDatabasePrivateEndpointsSortOrderEnum {
	values := make([]ListDistributedDatabasePrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListDistributedDatabasePrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDistributedDatabasePrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListDistributedDatabasePrivateEndpointsSortOrderEnum
func GetListDistributedDatabasePrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDistributedDatabasePrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDistributedDatabasePrivateEndpointsSortOrderEnum(val string) (ListDistributedDatabasePrivateEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListDistributedDatabasePrivateEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDistributedDatabasePrivateEndpointsSortByEnum Enum with underlying type: string
type ListDistributedDatabasePrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListDistributedDatabasePrivateEndpointsSortByEnum
const (
	ListDistributedDatabasePrivateEndpointsSortByTimecreated ListDistributedDatabasePrivateEndpointsSortByEnum = "timeCreated"
	ListDistributedDatabasePrivateEndpointsSortByTimeupdated ListDistributedDatabasePrivateEndpointsSortByEnum = "timeUpdated"
)

var mappingListDistributedDatabasePrivateEndpointsSortByEnum = map[string]ListDistributedDatabasePrivateEndpointsSortByEnum{
	"timeCreated": ListDistributedDatabasePrivateEndpointsSortByTimecreated,
	"timeUpdated": ListDistributedDatabasePrivateEndpointsSortByTimeupdated,
}

var mappingListDistributedDatabasePrivateEndpointsSortByEnumLowerCase = map[string]ListDistributedDatabasePrivateEndpointsSortByEnum{
	"timecreated": ListDistributedDatabasePrivateEndpointsSortByTimecreated,
	"timeupdated": ListDistributedDatabasePrivateEndpointsSortByTimeupdated,
}

// GetListDistributedDatabasePrivateEndpointsSortByEnumValues Enumerates the set of values for ListDistributedDatabasePrivateEndpointsSortByEnum
func GetListDistributedDatabasePrivateEndpointsSortByEnumValues() []ListDistributedDatabasePrivateEndpointsSortByEnum {
	values := make([]ListDistributedDatabasePrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListDistributedDatabasePrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDistributedDatabasePrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListDistributedDatabasePrivateEndpointsSortByEnum
func GetListDistributedDatabasePrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListDistributedDatabasePrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDistributedDatabasePrivateEndpointsSortByEnum(val string) (ListDistributedDatabasePrivateEndpointsSortByEnum, bool) {
	enum, ok := mappingListDistributedDatabasePrivateEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
