// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package computeinstanceagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInstanceAgentCommandExecutionsRequest wrapper for the ListInstanceAgentCommandExecutions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/ListInstanceAgentCommandExecutions.go.html to see an example of how to use ListInstanceAgentCommandExecutionsRequest.
type ListInstanceAgentCommandExecutionsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the instance.
	InstanceId *string `mandatory:"true" contributesTo:"query" name:"instanceId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// `TIMECREATED` is descending.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListInstanceAgentCommandExecutionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The `DISPLAYNAME` sort order
	// is case sensitive.
	SortOrder ListInstanceAgentCommandExecutionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState InstanceAgentCommandExecutionSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInstanceAgentCommandExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInstanceAgentCommandExecutionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInstanceAgentCommandExecutionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInstanceAgentCommandExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInstanceAgentCommandExecutionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInstanceAgentCommandExecutionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInstanceAgentCommandExecutionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInstanceAgentCommandExecutionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInstanceAgentCommandExecutionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInstanceAgentCommandExecutionSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetInstanceAgentCommandExecutionSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInstanceAgentCommandExecutionsResponse wrapper for the ListInstanceAgentCommandExecutions operation
type ListInstanceAgentCommandExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []InstanceAgentCommandExecutionSummary instances
	Items []InstanceAgentCommandExecutionSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListInstanceAgentCommandExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInstanceAgentCommandExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInstanceAgentCommandExecutionsSortByEnum Enum with underlying type: string
type ListInstanceAgentCommandExecutionsSortByEnum string

// Set of constants representing the allowable values for ListInstanceAgentCommandExecutionsSortByEnum
const (
	ListInstanceAgentCommandExecutionsSortByTimecreated ListInstanceAgentCommandExecutionsSortByEnum = "TIMECREATED"
	ListInstanceAgentCommandExecutionsSortByDisplayname ListInstanceAgentCommandExecutionsSortByEnum = "DISPLAYNAME"
)

var mappingListInstanceAgentCommandExecutionsSortByEnum = map[string]ListInstanceAgentCommandExecutionsSortByEnum{
	"TIMECREATED": ListInstanceAgentCommandExecutionsSortByTimecreated,
	"DISPLAYNAME": ListInstanceAgentCommandExecutionsSortByDisplayname,
}

var mappingListInstanceAgentCommandExecutionsSortByEnumLowerCase = map[string]ListInstanceAgentCommandExecutionsSortByEnum{
	"timecreated": ListInstanceAgentCommandExecutionsSortByTimecreated,
	"displayname": ListInstanceAgentCommandExecutionsSortByDisplayname,
}

// GetListInstanceAgentCommandExecutionsSortByEnumValues Enumerates the set of values for ListInstanceAgentCommandExecutionsSortByEnum
func GetListInstanceAgentCommandExecutionsSortByEnumValues() []ListInstanceAgentCommandExecutionsSortByEnum {
	values := make([]ListInstanceAgentCommandExecutionsSortByEnum, 0)
	for _, v := range mappingListInstanceAgentCommandExecutionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstanceAgentCommandExecutionsSortByEnumStringValues Enumerates the set of values in String for ListInstanceAgentCommandExecutionsSortByEnum
func GetListInstanceAgentCommandExecutionsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListInstanceAgentCommandExecutionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstanceAgentCommandExecutionsSortByEnum(val string) (ListInstanceAgentCommandExecutionsSortByEnum, bool) {
	enum, ok := mappingListInstanceAgentCommandExecutionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInstanceAgentCommandExecutionsSortOrderEnum Enum with underlying type: string
type ListInstanceAgentCommandExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListInstanceAgentCommandExecutionsSortOrderEnum
const (
	ListInstanceAgentCommandExecutionsSortOrderAsc  ListInstanceAgentCommandExecutionsSortOrderEnum = "ASC"
	ListInstanceAgentCommandExecutionsSortOrderDesc ListInstanceAgentCommandExecutionsSortOrderEnum = "DESC"
)

var mappingListInstanceAgentCommandExecutionsSortOrderEnum = map[string]ListInstanceAgentCommandExecutionsSortOrderEnum{
	"ASC":  ListInstanceAgentCommandExecutionsSortOrderAsc,
	"DESC": ListInstanceAgentCommandExecutionsSortOrderDesc,
}

var mappingListInstanceAgentCommandExecutionsSortOrderEnumLowerCase = map[string]ListInstanceAgentCommandExecutionsSortOrderEnum{
	"asc":  ListInstanceAgentCommandExecutionsSortOrderAsc,
	"desc": ListInstanceAgentCommandExecutionsSortOrderDesc,
}

// GetListInstanceAgentCommandExecutionsSortOrderEnumValues Enumerates the set of values for ListInstanceAgentCommandExecutionsSortOrderEnum
func GetListInstanceAgentCommandExecutionsSortOrderEnumValues() []ListInstanceAgentCommandExecutionsSortOrderEnum {
	values := make([]ListInstanceAgentCommandExecutionsSortOrderEnum, 0)
	for _, v := range mappingListInstanceAgentCommandExecutionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstanceAgentCommandExecutionsSortOrderEnumStringValues Enumerates the set of values in String for ListInstanceAgentCommandExecutionsSortOrderEnum
func GetListInstanceAgentCommandExecutionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInstanceAgentCommandExecutionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstanceAgentCommandExecutionsSortOrderEnum(val string) (ListInstanceAgentCommandExecutionsSortOrderEnum, bool) {
	enum, ok := mappingListInstanceAgentCommandExecutionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
