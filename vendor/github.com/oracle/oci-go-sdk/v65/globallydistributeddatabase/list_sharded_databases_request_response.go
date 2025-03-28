// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package globallydistributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListShardedDatabasesRequest wrapper for the ListShardedDatabases operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ListShardedDatabases.go.html to see an example of how to use ListShardedDatabasesRequest.
type ListShardedDatabasesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ShardedDatabaseLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListShardedDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListShardedDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only sharded databases that match the entire name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListShardedDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListShardedDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListShardedDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListShardedDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListShardedDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShardedDatabaseLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetShardedDatabaseLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListShardedDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListShardedDatabasesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListShardedDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListShardedDatabasesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListShardedDatabasesResponse wrapper for the ListShardedDatabases operation
type ListShardedDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ShardedDatabaseCollection instances
	ShardedDatabaseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListShardedDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListShardedDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListShardedDatabasesSortOrderEnum Enum with underlying type: string
type ListShardedDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListShardedDatabasesSortOrderEnum
const (
	ListShardedDatabasesSortOrderAsc  ListShardedDatabasesSortOrderEnum = "ASC"
	ListShardedDatabasesSortOrderDesc ListShardedDatabasesSortOrderEnum = "DESC"
)

var mappingListShardedDatabasesSortOrderEnum = map[string]ListShardedDatabasesSortOrderEnum{
	"ASC":  ListShardedDatabasesSortOrderAsc,
	"DESC": ListShardedDatabasesSortOrderDesc,
}

var mappingListShardedDatabasesSortOrderEnumLowerCase = map[string]ListShardedDatabasesSortOrderEnum{
	"asc":  ListShardedDatabasesSortOrderAsc,
	"desc": ListShardedDatabasesSortOrderDesc,
}

// GetListShardedDatabasesSortOrderEnumValues Enumerates the set of values for ListShardedDatabasesSortOrderEnum
func GetListShardedDatabasesSortOrderEnumValues() []ListShardedDatabasesSortOrderEnum {
	values := make([]ListShardedDatabasesSortOrderEnum, 0)
	for _, v := range mappingListShardedDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListShardedDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListShardedDatabasesSortOrderEnum
func GetListShardedDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListShardedDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListShardedDatabasesSortOrderEnum(val string) (ListShardedDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListShardedDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListShardedDatabasesSortByEnum Enum with underlying type: string
type ListShardedDatabasesSortByEnum string

// Set of constants representing the allowable values for ListShardedDatabasesSortByEnum
const (
	ListShardedDatabasesSortByTimecreated ListShardedDatabasesSortByEnum = "timeCreated"
	ListShardedDatabasesSortByTimeupdated ListShardedDatabasesSortByEnum = "timeUpdated"
)

var mappingListShardedDatabasesSortByEnum = map[string]ListShardedDatabasesSortByEnum{
	"timeCreated": ListShardedDatabasesSortByTimecreated,
	"timeUpdated": ListShardedDatabasesSortByTimeupdated,
}

var mappingListShardedDatabasesSortByEnumLowerCase = map[string]ListShardedDatabasesSortByEnum{
	"timecreated": ListShardedDatabasesSortByTimecreated,
	"timeupdated": ListShardedDatabasesSortByTimeupdated,
}

// GetListShardedDatabasesSortByEnumValues Enumerates the set of values for ListShardedDatabasesSortByEnum
func GetListShardedDatabasesSortByEnumValues() []ListShardedDatabasesSortByEnum {
	values := make([]ListShardedDatabasesSortByEnum, 0)
	for _, v := range mappingListShardedDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListShardedDatabasesSortByEnumStringValues Enumerates the set of values in String for ListShardedDatabasesSortByEnum
func GetListShardedDatabasesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListShardedDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListShardedDatabasesSortByEnum(val string) (ListShardedDatabasesSortByEnum, bool) {
	enum, ok := mappingListShardedDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
