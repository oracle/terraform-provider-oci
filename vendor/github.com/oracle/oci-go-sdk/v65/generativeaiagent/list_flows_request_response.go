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

// ListFlowsRequest wrapper for the ListFlows operation
type ListFlowsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState FlowLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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
	SortOrder ListFlowsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListFlowsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFlowsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFlowsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFlowsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFlowsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFlowsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFlowLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetFlowLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFlowsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFlowsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFlowsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFlowsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFlowsResponse wrapper for the ListFlows operation
type ListFlowsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FlowCollection instances
	FlowCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFlowsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFlowsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFlowsSortOrderEnum Enum with underlying type: string
type ListFlowsSortOrderEnum string

// Set of constants representing the allowable values for ListFlowsSortOrderEnum
const (
	ListFlowsSortOrderAsc  ListFlowsSortOrderEnum = "ASC"
	ListFlowsSortOrderDesc ListFlowsSortOrderEnum = "DESC"
)

var mappingListFlowsSortOrderEnum = map[string]ListFlowsSortOrderEnum{
	"ASC":  ListFlowsSortOrderAsc,
	"DESC": ListFlowsSortOrderDesc,
}

var mappingListFlowsSortOrderEnumLowerCase = map[string]ListFlowsSortOrderEnum{
	"asc":  ListFlowsSortOrderAsc,
	"desc": ListFlowsSortOrderDesc,
}

// GetListFlowsSortOrderEnumValues Enumerates the set of values for ListFlowsSortOrderEnum
func GetListFlowsSortOrderEnumValues() []ListFlowsSortOrderEnum {
	values := make([]ListFlowsSortOrderEnum, 0)
	for _, v := range mappingListFlowsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFlowsSortOrderEnumStringValues Enumerates the set of values in String for ListFlowsSortOrderEnum
func GetListFlowsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFlowsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFlowsSortOrderEnum(val string) (ListFlowsSortOrderEnum, bool) {
	enum, ok := mappingListFlowsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFlowsSortByEnum Enum with underlying type: string
type ListFlowsSortByEnum string

// Set of constants representing the allowable values for ListFlowsSortByEnum
const (
	ListFlowsSortByTimecreated ListFlowsSortByEnum = "timeCreated"
	ListFlowsSortByDisplayname ListFlowsSortByEnum = "displayName"
)

var mappingListFlowsSortByEnum = map[string]ListFlowsSortByEnum{
	"timeCreated": ListFlowsSortByTimecreated,
	"displayName": ListFlowsSortByDisplayname,
}

var mappingListFlowsSortByEnumLowerCase = map[string]ListFlowsSortByEnum{
	"timecreated": ListFlowsSortByTimecreated,
	"displayname": ListFlowsSortByDisplayname,
}

// GetListFlowsSortByEnumValues Enumerates the set of values for ListFlowsSortByEnum
func GetListFlowsSortByEnumValues() []ListFlowsSortByEnum {
	values := make([]ListFlowsSortByEnum, 0)
	for _, v := range mappingListFlowsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFlowsSortByEnumStringValues Enumerates the set of values in String for ListFlowsSortByEnum
func GetListFlowsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListFlowsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFlowsSortByEnum(val string) (ListFlowsSortByEnum, bool) {
	enum, ok := mappingListFlowsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
