// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAgentDependenciesRequest wrapper for the ListAgentDependencies operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListAgentDependencies.go.html to see an example of how to use ListAgentDependenciesRequest.
type ListAgentDependenciesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given Agent ID.
	AgentId *string `mandatory:"false" contributesTo:"query" name:"agentId"`

	// A filter to return only resources that match the given environment ID.
	EnvironmentId *string `mandatory:"false" contributesTo:"query" name:"environmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState AgentDependencyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAgentDependenciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListAgentDependenciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAgentDependenciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAgentDependenciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAgentDependenciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAgentDependenciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAgentDependenciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAgentDependencyLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAgentDependencyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAgentDependenciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAgentDependenciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAgentDependenciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAgentDependenciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAgentDependenciesResponse wrapper for the ListAgentDependencies operation
type ListAgentDependenciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AgentDependencyCollection instances
	AgentDependencyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAgentDependenciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAgentDependenciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAgentDependenciesSortOrderEnum Enum with underlying type: string
type ListAgentDependenciesSortOrderEnum string

// Set of constants representing the allowable values for ListAgentDependenciesSortOrderEnum
const (
	ListAgentDependenciesSortOrderAsc  ListAgentDependenciesSortOrderEnum = "ASC"
	ListAgentDependenciesSortOrderDesc ListAgentDependenciesSortOrderEnum = "DESC"
)

var mappingListAgentDependenciesSortOrderEnum = map[string]ListAgentDependenciesSortOrderEnum{
	"ASC":  ListAgentDependenciesSortOrderAsc,
	"DESC": ListAgentDependenciesSortOrderDesc,
}

var mappingListAgentDependenciesSortOrderEnumLowerCase = map[string]ListAgentDependenciesSortOrderEnum{
	"asc":  ListAgentDependenciesSortOrderAsc,
	"desc": ListAgentDependenciesSortOrderDesc,
}

// GetListAgentDependenciesSortOrderEnumValues Enumerates the set of values for ListAgentDependenciesSortOrderEnum
func GetListAgentDependenciesSortOrderEnumValues() []ListAgentDependenciesSortOrderEnum {
	values := make([]ListAgentDependenciesSortOrderEnum, 0)
	for _, v := range mappingListAgentDependenciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgentDependenciesSortOrderEnumStringValues Enumerates the set of values in String for ListAgentDependenciesSortOrderEnum
func GetListAgentDependenciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAgentDependenciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgentDependenciesSortOrderEnum(val string) (ListAgentDependenciesSortOrderEnum, bool) {
	enum, ok := mappingListAgentDependenciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAgentDependenciesSortByEnum Enum with underlying type: string
type ListAgentDependenciesSortByEnum string

// Set of constants representing the allowable values for ListAgentDependenciesSortByEnum
const (
	ListAgentDependenciesSortByTimecreated ListAgentDependenciesSortByEnum = "timeCreated"
	ListAgentDependenciesSortByTimeupdated ListAgentDependenciesSortByEnum = "timeUpdated"
	ListAgentDependenciesSortByDisplayname ListAgentDependenciesSortByEnum = "displayName"
)

var mappingListAgentDependenciesSortByEnum = map[string]ListAgentDependenciesSortByEnum{
	"timeCreated": ListAgentDependenciesSortByTimecreated,
	"timeUpdated": ListAgentDependenciesSortByTimeupdated,
	"displayName": ListAgentDependenciesSortByDisplayname,
}

var mappingListAgentDependenciesSortByEnumLowerCase = map[string]ListAgentDependenciesSortByEnum{
	"timecreated": ListAgentDependenciesSortByTimecreated,
	"timeupdated": ListAgentDependenciesSortByTimeupdated,
	"displayname": ListAgentDependenciesSortByDisplayname,
}

// GetListAgentDependenciesSortByEnumValues Enumerates the set of values for ListAgentDependenciesSortByEnum
func GetListAgentDependenciesSortByEnumValues() []ListAgentDependenciesSortByEnum {
	values := make([]ListAgentDependenciesSortByEnum, 0)
	for _, v := range mappingListAgentDependenciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgentDependenciesSortByEnumStringValues Enumerates the set of values in String for ListAgentDependenciesSortByEnum
func GetListAgentDependenciesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListAgentDependenciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgentDependenciesSortByEnum(val string) (ListAgentDependenciesSortByEnum, bool) {
	enum, ok := mappingListAgentDependenciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
