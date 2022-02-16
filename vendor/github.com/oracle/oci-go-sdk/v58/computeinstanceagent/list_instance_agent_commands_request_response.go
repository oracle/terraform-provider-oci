// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package computeinstanceagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListInstanceAgentCommandsRequest wrapper for the ListInstanceAgentCommands operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/ListInstanceAgentCommands.go.html to see an example of how to use ListInstanceAgentCommandsRequest.
type ListInstanceAgentCommandsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

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
	SortBy ListInstanceAgentCommandsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The `DISPLAYNAME` sort order
	// is case sensitive.
	SortOrder ListInstanceAgentCommandsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInstanceAgentCommandsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInstanceAgentCommandsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInstanceAgentCommandsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInstanceAgentCommandsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInstanceAgentCommandsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInstanceAgentCommandsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInstanceAgentCommandsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInstanceAgentCommandsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInstanceAgentCommandsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInstanceAgentCommandsResponse wrapper for the ListInstanceAgentCommands operation
type ListInstanceAgentCommandsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []InstanceAgentCommandSummary instances
	Items []InstanceAgentCommandSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListInstanceAgentCommandsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInstanceAgentCommandsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInstanceAgentCommandsSortByEnum Enum with underlying type: string
type ListInstanceAgentCommandsSortByEnum string

// Set of constants representing the allowable values for ListInstanceAgentCommandsSortByEnum
const (
	ListInstanceAgentCommandsSortByTimecreated ListInstanceAgentCommandsSortByEnum = "TIMECREATED"
	ListInstanceAgentCommandsSortByDisplayname ListInstanceAgentCommandsSortByEnum = "DISPLAYNAME"
)

var mappingListInstanceAgentCommandsSortByEnum = map[string]ListInstanceAgentCommandsSortByEnum{
	"TIMECREATED": ListInstanceAgentCommandsSortByTimecreated,
	"DISPLAYNAME": ListInstanceAgentCommandsSortByDisplayname,
}

// GetListInstanceAgentCommandsSortByEnumValues Enumerates the set of values for ListInstanceAgentCommandsSortByEnum
func GetListInstanceAgentCommandsSortByEnumValues() []ListInstanceAgentCommandsSortByEnum {
	values := make([]ListInstanceAgentCommandsSortByEnum, 0)
	for _, v := range mappingListInstanceAgentCommandsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstanceAgentCommandsSortByEnumStringValues Enumerates the set of values in String for ListInstanceAgentCommandsSortByEnum
func GetListInstanceAgentCommandsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListInstanceAgentCommandsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstanceAgentCommandsSortByEnum(val string) (ListInstanceAgentCommandsSortByEnum, bool) {
	mappingListInstanceAgentCommandsSortByEnumIgnoreCase := make(map[string]ListInstanceAgentCommandsSortByEnum)
	for k, v := range mappingListInstanceAgentCommandsSortByEnum {
		mappingListInstanceAgentCommandsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListInstanceAgentCommandsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListInstanceAgentCommandsSortOrderEnum Enum with underlying type: string
type ListInstanceAgentCommandsSortOrderEnum string

// Set of constants representing the allowable values for ListInstanceAgentCommandsSortOrderEnum
const (
	ListInstanceAgentCommandsSortOrderAsc  ListInstanceAgentCommandsSortOrderEnum = "ASC"
	ListInstanceAgentCommandsSortOrderDesc ListInstanceAgentCommandsSortOrderEnum = "DESC"
)

var mappingListInstanceAgentCommandsSortOrderEnum = map[string]ListInstanceAgentCommandsSortOrderEnum{
	"ASC":  ListInstanceAgentCommandsSortOrderAsc,
	"DESC": ListInstanceAgentCommandsSortOrderDesc,
}

// GetListInstanceAgentCommandsSortOrderEnumValues Enumerates the set of values for ListInstanceAgentCommandsSortOrderEnum
func GetListInstanceAgentCommandsSortOrderEnumValues() []ListInstanceAgentCommandsSortOrderEnum {
	values := make([]ListInstanceAgentCommandsSortOrderEnum, 0)
	for _, v := range mappingListInstanceAgentCommandsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstanceAgentCommandsSortOrderEnumStringValues Enumerates the set of values in String for ListInstanceAgentCommandsSortOrderEnum
func GetListInstanceAgentCommandsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInstanceAgentCommandsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstanceAgentCommandsSortOrderEnum(val string) (ListInstanceAgentCommandsSortOrderEnum, bool) {
	mappingListInstanceAgentCommandsSortOrderEnumIgnoreCase := make(map[string]ListInstanceAgentCommandsSortOrderEnum)
	for k, v := range mappingListInstanceAgentCommandsSortOrderEnum {
		mappingListInstanceAgentCommandsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListInstanceAgentCommandsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
