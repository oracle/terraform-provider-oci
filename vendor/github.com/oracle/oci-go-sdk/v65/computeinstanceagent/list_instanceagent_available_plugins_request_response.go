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

// ListInstanceagentAvailablePluginsRequest wrapper for the ListInstanceagentAvailablePlugins operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/ListInstanceagentAvailablePlugins.go.html to see an example of how to use ListInstanceagentAvailablePluginsRequest.
type ListInstanceagentAvailablePluginsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OS for which the plugin is supported.
	// Examples of OperatingSystemQueryParam:OperatingSystemVersionQueryParam are as follows:
	// 'CentOS' '6.10' , 'CentOS Linux' '7', 'CentOS Linux' '8',
	// 'Oracle Linux Server' '6.10', 'Oracle Linux Server' '8.0',
	// 'Red Hat Enterprise Linux Server' '7.8',
	// 'Windows' '10', 'Windows' '2008ServerR2', 'Windows' '2012ServerR2', 'Windows' '7', 'Windows' '8.1'
	OsName *string `mandatory:"true" contributesTo:"query" name:"osName"`

	// The OS version for which the plugin is supported.
	OsVersion *string `mandatory:"true" contributesTo:"query" name:"osVersion"`

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
	SortBy ListInstanceagentAvailablePluginsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The `DISPLAYNAME` sort order
	// is case sensitive.
	SortOrder ListInstanceagentAvailablePluginsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The plugin name
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInstanceagentAvailablePluginsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInstanceagentAvailablePluginsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInstanceagentAvailablePluginsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInstanceagentAvailablePluginsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInstanceagentAvailablePluginsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInstanceagentAvailablePluginsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInstanceagentAvailablePluginsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInstanceagentAvailablePluginsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInstanceagentAvailablePluginsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInstanceagentAvailablePluginsResponse wrapper for the ListInstanceagentAvailablePlugins operation
type ListInstanceagentAvailablePluginsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AvailablePluginSummary instances
	Items []AvailablePluginSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInstanceagentAvailablePluginsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInstanceagentAvailablePluginsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInstanceagentAvailablePluginsSortByEnum Enum with underlying type: string
type ListInstanceagentAvailablePluginsSortByEnum string

// Set of constants representing the allowable values for ListInstanceagentAvailablePluginsSortByEnum
const (
	ListInstanceagentAvailablePluginsSortByTimecreated ListInstanceagentAvailablePluginsSortByEnum = "TIMECREATED"
	ListInstanceagentAvailablePluginsSortByDisplayname ListInstanceagentAvailablePluginsSortByEnum = "DISPLAYNAME"
)

var mappingListInstanceagentAvailablePluginsSortByEnum = map[string]ListInstanceagentAvailablePluginsSortByEnum{
	"TIMECREATED": ListInstanceagentAvailablePluginsSortByTimecreated,
	"DISPLAYNAME": ListInstanceagentAvailablePluginsSortByDisplayname,
}

var mappingListInstanceagentAvailablePluginsSortByEnumLowerCase = map[string]ListInstanceagentAvailablePluginsSortByEnum{
	"timecreated": ListInstanceagentAvailablePluginsSortByTimecreated,
	"displayname": ListInstanceagentAvailablePluginsSortByDisplayname,
}

// GetListInstanceagentAvailablePluginsSortByEnumValues Enumerates the set of values for ListInstanceagentAvailablePluginsSortByEnum
func GetListInstanceagentAvailablePluginsSortByEnumValues() []ListInstanceagentAvailablePluginsSortByEnum {
	values := make([]ListInstanceagentAvailablePluginsSortByEnum, 0)
	for _, v := range mappingListInstanceagentAvailablePluginsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstanceagentAvailablePluginsSortByEnumStringValues Enumerates the set of values in String for ListInstanceagentAvailablePluginsSortByEnum
func GetListInstanceagentAvailablePluginsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListInstanceagentAvailablePluginsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstanceagentAvailablePluginsSortByEnum(val string) (ListInstanceagentAvailablePluginsSortByEnum, bool) {
	enum, ok := mappingListInstanceagentAvailablePluginsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInstanceagentAvailablePluginsSortOrderEnum Enum with underlying type: string
type ListInstanceagentAvailablePluginsSortOrderEnum string

// Set of constants representing the allowable values for ListInstanceagentAvailablePluginsSortOrderEnum
const (
	ListInstanceagentAvailablePluginsSortOrderAsc  ListInstanceagentAvailablePluginsSortOrderEnum = "ASC"
	ListInstanceagentAvailablePluginsSortOrderDesc ListInstanceagentAvailablePluginsSortOrderEnum = "DESC"
)

var mappingListInstanceagentAvailablePluginsSortOrderEnum = map[string]ListInstanceagentAvailablePluginsSortOrderEnum{
	"ASC":  ListInstanceagentAvailablePluginsSortOrderAsc,
	"DESC": ListInstanceagentAvailablePluginsSortOrderDesc,
}

var mappingListInstanceagentAvailablePluginsSortOrderEnumLowerCase = map[string]ListInstanceagentAvailablePluginsSortOrderEnum{
	"asc":  ListInstanceagentAvailablePluginsSortOrderAsc,
	"desc": ListInstanceagentAvailablePluginsSortOrderDesc,
}

// GetListInstanceagentAvailablePluginsSortOrderEnumValues Enumerates the set of values for ListInstanceagentAvailablePluginsSortOrderEnum
func GetListInstanceagentAvailablePluginsSortOrderEnumValues() []ListInstanceagentAvailablePluginsSortOrderEnum {
	values := make([]ListInstanceagentAvailablePluginsSortOrderEnum, 0)
	for _, v := range mappingListInstanceagentAvailablePluginsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstanceagentAvailablePluginsSortOrderEnumStringValues Enumerates the set of values in String for ListInstanceagentAvailablePluginsSortOrderEnum
func GetListInstanceagentAvailablePluginsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInstanceagentAvailablePluginsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstanceagentAvailablePluginsSortOrderEnum(val string) (ListInstanceagentAvailablePluginsSortOrderEnum, bool) {
	enum, ok := mappingListInstanceagentAvailablePluginsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
