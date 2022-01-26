// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package computeinstanceagent

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListInstanceAgentPluginsRequest wrapper for the ListInstanceAgentPlugins operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/ListInstanceAgentPlugins.go.html to see an example of how to use ListInstanceAgentPluginsRequest.
type ListInstanceAgentPluginsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID of the instance.
	InstanceagentId *string `mandatory:"true" contributesTo:"path" name:"instanceagentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The plugin status
	Status ListInstanceAgentPluginsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

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
	SortBy ListInstanceAgentPluginsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The `DISPLAYNAME` sort order
	// is case sensitive.
	SortOrder ListInstanceAgentPluginsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The plugin name
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInstanceAgentPluginsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInstanceAgentPluginsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInstanceAgentPluginsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInstanceAgentPluginsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListInstanceAgentPluginsResponse wrapper for the ListInstanceAgentPlugins operation
type ListInstanceAgentPluginsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []InstanceAgentPluginSummary instances
	Items []InstanceAgentPluginSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInstanceAgentPluginsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInstanceAgentPluginsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInstanceAgentPluginsStatusEnum Enum with underlying type: string
type ListInstanceAgentPluginsStatusEnum string

// Set of constants representing the allowable values for ListInstanceAgentPluginsStatusEnum
const (
	ListInstanceAgentPluginsStatusRunning      ListInstanceAgentPluginsStatusEnum = "RUNNING"
	ListInstanceAgentPluginsStatusStopped      ListInstanceAgentPluginsStatusEnum = "STOPPED"
	ListInstanceAgentPluginsStatusNotSupported ListInstanceAgentPluginsStatusEnum = "NOT_SUPPORTED"
	ListInstanceAgentPluginsStatusInvalid      ListInstanceAgentPluginsStatusEnum = "INVALID"
)

var mappingListInstanceAgentPluginsStatus = map[string]ListInstanceAgentPluginsStatusEnum{
	"RUNNING":       ListInstanceAgentPluginsStatusRunning,
	"STOPPED":       ListInstanceAgentPluginsStatusStopped,
	"NOT_SUPPORTED": ListInstanceAgentPluginsStatusNotSupported,
	"INVALID":       ListInstanceAgentPluginsStatusInvalid,
}

// GetListInstanceAgentPluginsStatusEnumValues Enumerates the set of values for ListInstanceAgentPluginsStatusEnum
func GetListInstanceAgentPluginsStatusEnumValues() []ListInstanceAgentPluginsStatusEnum {
	values := make([]ListInstanceAgentPluginsStatusEnum, 0)
	for _, v := range mappingListInstanceAgentPluginsStatus {
		values = append(values, v)
	}
	return values
}

// ListInstanceAgentPluginsSortByEnum Enum with underlying type: string
type ListInstanceAgentPluginsSortByEnum string

// Set of constants representing the allowable values for ListInstanceAgentPluginsSortByEnum
const (
	ListInstanceAgentPluginsSortByTimecreated ListInstanceAgentPluginsSortByEnum = "TIMECREATED"
	ListInstanceAgentPluginsSortByDisplayname ListInstanceAgentPluginsSortByEnum = "DISPLAYNAME"
)

var mappingListInstanceAgentPluginsSortBy = map[string]ListInstanceAgentPluginsSortByEnum{
	"TIMECREATED": ListInstanceAgentPluginsSortByTimecreated,
	"DISPLAYNAME": ListInstanceAgentPluginsSortByDisplayname,
}

// GetListInstanceAgentPluginsSortByEnumValues Enumerates the set of values for ListInstanceAgentPluginsSortByEnum
func GetListInstanceAgentPluginsSortByEnumValues() []ListInstanceAgentPluginsSortByEnum {
	values := make([]ListInstanceAgentPluginsSortByEnum, 0)
	for _, v := range mappingListInstanceAgentPluginsSortBy {
		values = append(values, v)
	}
	return values
}

// ListInstanceAgentPluginsSortOrderEnum Enum with underlying type: string
type ListInstanceAgentPluginsSortOrderEnum string

// Set of constants representing the allowable values for ListInstanceAgentPluginsSortOrderEnum
const (
	ListInstanceAgentPluginsSortOrderAsc  ListInstanceAgentPluginsSortOrderEnum = "ASC"
	ListInstanceAgentPluginsSortOrderDesc ListInstanceAgentPluginsSortOrderEnum = "DESC"
)

var mappingListInstanceAgentPluginsSortOrder = map[string]ListInstanceAgentPluginsSortOrderEnum{
	"ASC":  ListInstanceAgentPluginsSortOrderAsc,
	"DESC": ListInstanceAgentPluginsSortOrderDesc,
}

// GetListInstanceAgentPluginsSortOrderEnumValues Enumerates the set of values for ListInstanceAgentPluginsSortOrderEnum
func GetListInstanceAgentPluginsSortOrderEnumValues() []ListInstanceAgentPluginsSortOrderEnum {
	values := make([]ListInstanceAgentPluginsSortOrderEnum, 0)
	for _, v := range mappingListInstanceAgentPluginsSortOrder {
		values = append(values, v)
	}
	return values
}
