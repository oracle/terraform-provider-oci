// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAgentsRequest wrapper for the ListAgents operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListAgents.go.html to see an example of how to use ListAgentsRequest.
type ListAgentsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given environment ID.
	EnvironmentId *string `mandatory:"false" contributesTo:"query" name:"environmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState AgentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given Agent ID.
	AgentId *string `mandatory:"false" contributesTo:"query" name:"agentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAgentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListAgentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAgentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAgentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAgentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAgentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAgentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAgentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAgentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAgentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAgentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAgentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAgentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAgentsResponse wrapper for the ListAgents operation
type ListAgentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AgentCollection instances
	AgentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAgentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAgentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAgentsSortOrderEnum Enum with underlying type: string
type ListAgentsSortOrderEnum string

// Set of constants representing the allowable values for ListAgentsSortOrderEnum
const (
	ListAgentsSortOrderAsc  ListAgentsSortOrderEnum = "ASC"
	ListAgentsSortOrderDesc ListAgentsSortOrderEnum = "DESC"
)

var mappingListAgentsSortOrderEnum = map[string]ListAgentsSortOrderEnum{
	"ASC":  ListAgentsSortOrderAsc,
	"DESC": ListAgentsSortOrderDesc,
}

var mappingListAgentsSortOrderEnumLowerCase = map[string]ListAgentsSortOrderEnum{
	"asc":  ListAgentsSortOrderAsc,
	"desc": ListAgentsSortOrderDesc,
}

// GetListAgentsSortOrderEnumValues Enumerates the set of values for ListAgentsSortOrderEnum
func GetListAgentsSortOrderEnumValues() []ListAgentsSortOrderEnum {
	values := make([]ListAgentsSortOrderEnum, 0)
	for _, v := range mappingListAgentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgentsSortOrderEnumStringValues Enumerates the set of values in String for ListAgentsSortOrderEnum
func GetListAgentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAgentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgentsSortOrderEnum(val string) (ListAgentsSortOrderEnum, bool) {
	enum, ok := mappingListAgentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAgentsSortByEnum Enum with underlying type: string
type ListAgentsSortByEnum string

// Set of constants representing the allowable values for ListAgentsSortByEnum
const (
	ListAgentsSortByTimecreated ListAgentsSortByEnum = "timeCreated"
	ListAgentsSortByTimeupdated ListAgentsSortByEnum = "timeUpdated"
	ListAgentsSortByDisplayname ListAgentsSortByEnum = "displayName"
)

var mappingListAgentsSortByEnum = map[string]ListAgentsSortByEnum{
	"timeCreated": ListAgentsSortByTimecreated,
	"timeUpdated": ListAgentsSortByTimeupdated,
	"displayName": ListAgentsSortByDisplayname,
}

var mappingListAgentsSortByEnumLowerCase = map[string]ListAgentsSortByEnum{
	"timecreated": ListAgentsSortByTimecreated,
	"timeupdated": ListAgentsSortByTimeupdated,
	"displayname": ListAgentsSortByDisplayname,
}

// GetListAgentsSortByEnumValues Enumerates the set of values for ListAgentsSortByEnum
func GetListAgentsSortByEnumValues() []ListAgentsSortByEnum {
	values := make([]ListAgentsSortByEnum, 0)
	for _, v := range mappingListAgentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgentsSortByEnumStringValues Enumerates the set of values in String for ListAgentsSortByEnum
func GetListAgentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListAgentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgentsSortByEnum(val string) (ListAgentsSortByEnum, bool) {
	enum, ok := mappingListAgentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
