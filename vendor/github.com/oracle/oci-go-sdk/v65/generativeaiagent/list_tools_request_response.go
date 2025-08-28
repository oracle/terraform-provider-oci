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

// ListToolsRequest wrapper for the ListTools operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiagent/ListTools.go.html to see an example of how to use ListToolsRequest.
type ListToolsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState ToolLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the agent.
	AgentId *string `mandatory:"false" contributesTo:"query" name:"agentId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListToolsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListToolsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListToolsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListToolsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListToolsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListToolsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListToolsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingToolLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetToolLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListToolsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListToolsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListToolsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListToolsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListToolsResponse wrapper for the ListTools operation
type ListToolsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ToolCollection instances
	ToolCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListToolsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListToolsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListToolsSortOrderEnum Enum with underlying type: string
type ListToolsSortOrderEnum string

// Set of constants representing the allowable values for ListToolsSortOrderEnum
const (
	ListToolsSortOrderAsc  ListToolsSortOrderEnum = "ASC"
	ListToolsSortOrderDesc ListToolsSortOrderEnum = "DESC"
)

var mappingListToolsSortOrderEnum = map[string]ListToolsSortOrderEnum{
	"ASC":  ListToolsSortOrderAsc,
	"DESC": ListToolsSortOrderDesc,
}

var mappingListToolsSortOrderEnumLowerCase = map[string]ListToolsSortOrderEnum{
	"asc":  ListToolsSortOrderAsc,
	"desc": ListToolsSortOrderDesc,
}

// GetListToolsSortOrderEnumValues Enumerates the set of values for ListToolsSortOrderEnum
func GetListToolsSortOrderEnumValues() []ListToolsSortOrderEnum {
	values := make([]ListToolsSortOrderEnum, 0)
	for _, v := range mappingListToolsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListToolsSortOrderEnumStringValues Enumerates the set of values in String for ListToolsSortOrderEnum
func GetListToolsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListToolsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListToolsSortOrderEnum(val string) (ListToolsSortOrderEnum, bool) {
	enum, ok := mappingListToolsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListToolsSortByEnum Enum with underlying type: string
type ListToolsSortByEnum string

// Set of constants representing the allowable values for ListToolsSortByEnum
const (
	ListToolsSortByTimecreated ListToolsSortByEnum = "timeCreated"
	ListToolsSortByDisplayname ListToolsSortByEnum = "displayName"
)

var mappingListToolsSortByEnum = map[string]ListToolsSortByEnum{
	"timeCreated": ListToolsSortByTimecreated,
	"displayName": ListToolsSortByDisplayname,
}

var mappingListToolsSortByEnumLowerCase = map[string]ListToolsSortByEnum{
	"timecreated": ListToolsSortByTimecreated,
	"displayname": ListToolsSortByDisplayname,
}

// GetListToolsSortByEnumValues Enumerates the set of values for ListToolsSortByEnum
func GetListToolsSortByEnumValues() []ListToolsSortByEnum {
	values := make([]ListToolsSortByEnum, 0)
	for _, v := range mappingListToolsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListToolsSortByEnumStringValues Enumerates the set of values in String for ListToolsSortByEnum
func GetListToolsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListToolsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListToolsSortByEnum(val string) (ListToolsSortByEnum, bool) {
	enum, ok := mappingListToolsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
