// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeaiagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAgentEndpointsRequest wrapper for the ListAgentEndpoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiagent/ListAgentEndpoints.go.html to see an example of how to use ListAgentEndpointsRequest.
type ListAgentEndpointsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the agent.
	AgentId *string `mandatory:"false" contributesTo:"query" name:"agentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState AgentEndpointLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAgentEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListAgentEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAgentEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAgentEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAgentEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAgentEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAgentEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAgentEndpointLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAgentEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAgentEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAgentEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAgentEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAgentEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAgentEndpointsResponse wrapper for the ListAgentEndpoints operation
type ListAgentEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AgentEndpointCollection instances
	AgentEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAgentEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAgentEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAgentEndpointsSortOrderEnum Enum with underlying type: string
type ListAgentEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListAgentEndpointsSortOrderEnum
const (
	ListAgentEndpointsSortOrderAsc  ListAgentEndpointsSortOrderEnum = "ASC"
	ListAgentEndpointsSortOrderDesc ListAgentEndpointsSortOrderEnum = "DESC"
)

var mappingListAgentEndpointsSortOrderEnum = map[string]ListAgentEndpointsSortOrderEnum{
	"ASC":  ListAgentEndpointsSortOrderAsc,
	"DESC": ListAgentEndpointsSortOrderDesc,
}

var mappingListAgentEndpointsSortOrderEnumLowerCase = map[string]ListAgentEndpointsSortOrderEnum{
	"asc":  ListAgentEndpointsSortOrderAsc,
	"desc": ListAgentEndpointsSortOrderDesc,
}

// GetListAgentEndpointsSortOrderEnumValues Enumerates the set of values for ListAgentEndpointsSortOrderEnum
func GetListAgentEndpointsSortOrderEnumValues() []ListAgentEndpointsSortOrderEnum {
	values := make([]ListAgentEndpointsSortOrderEnum, 0)
	for _, v := range mappingListAgentEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgentEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListAgentEndpointsSortOrderEnum
func GetListAgentEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAgentEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgentEndpointsSortOrderEnum(val string) (ListAgentEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListAgentEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAgentEndpointsSortByEnum Enum with underlying type: string
type ListAgentEndpointsSortByEnum string

// Set of constants representing the allowable values for ListAgentEndpointsSortByEnum
const (
	ListAgentEndpointsSortByTimecreated ListAgentEndpointsSortByEnum = "timeCreated"
	ListAgentEndpointsSortByDisplayname ListAgentEndpointsSortByEnum = "displayName"
)

var mappingListAgentEndpointsSortByEnum = map[string]ListAgentEndpointsSortByEnum{
	"timeCreated": ListAgentEndpointsSortByTimecreated,
	"displayName": ListAgentEndpointsSortByDisplayname,
}

var mappingListAgentEndpointsSortByEnumLowerCase = map[string]ListAgentEndpointsSortByEnum{
	"timecreated": ListAgentEndpointsSortByTimecreated,
	"displayname": ListAgentEndpointsSortByDisplayname,
}

// GetListAgentEndpointsSortByEnumValues Enumerates the set of values for ListAgentEndpointsSortByEnum
func GetListAgentEndpointsSortByEnumValues() []ListAgentEndpointsSortByEnum {
	values := make([]ListAgentEndpointsSortByEnum, 0)
	for _, v := range mappingListAgentEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgentEndpointsSortByEnumStringValues Enumerates the set of values in String for ListAgentEndpointsSortByEnum
func GetListAgentEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAgentEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgentEndpointsSortByEnum(val string) (ListAgentEndpointsSortByEnum, bool) {
	enum, ok := mappingListAgentEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
