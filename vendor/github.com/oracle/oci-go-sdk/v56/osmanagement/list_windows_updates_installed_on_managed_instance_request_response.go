// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListWindowsUpdatesInstalledOnManagedInstanceRequest wrapper for the ListWindowsUpdatesInstalledOnManagedInstance operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListWindowsUpdatesInstalledOnManagedInstance.go.html to see an example of how to use ListWindowsUpdatesInstalledOnManagedInstanceRequest.
type ListWindowsUpdatesInstalledOnManagedInstanceRequest struct {

	// OCID for the managed instance
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListWindowsUpdatesInstalledOnManagedInstanceSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListWindowsUpdatesInstalledOnManagedInstanceSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWindowsUpdatesInstalledOnManagedInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWindowsUpdatesInstalledOnManagedInstanceRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWindowsUpdatesInstalledOnManagedInstanceRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWindowsUpdatesInstalledOnManagedInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWindowsUpdatesInstalledOnManagedInstanceResponse wrapper for the ListWindowsUpdatesInstalledOnManagedInstance operation
type ListWindowsUpdatesInstalledOnManagedInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []InstalledWindowsUpdateSummary instances
	Items []InstalledWindowsUpdateSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWindowsUpdatesInstalledOnManagedInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWindowsUpdatesInstalledOnManagedInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWindowsUpdatesInstalledOnManagedInstanceSortOrderEnum Enum with underlying type: string
type ListWindowsUpdatesInstalledOnManagedInstanceSortOrderEnum string

// Set of constants representing the allowable values for ListWindowsUpdatesInstalledOnManagedInstanceSortOrderEnum
const (
	ListWindowsUpdatesInstalledOnManagedInstanceSortOrderAsc  ListWindowsUpdatesInstalledOnManagedInstanceSortOrderEnum = "ASC"
	ListWindowsUpdatesInstalledOnManagedInstanceSortOrderDesc ListWindowsUpdatesInstalledOnManagedInstanceSortOrderEnum = "DESC"
)

var mappingListWindowsUpdatesInstalledOnManagedInstanceSortOrder = map[string]ListWindowsUpdatesInstalledOnManagedInstanceSortOrderEnum{
	"ASC":  ListWindowsUpdatesInstalledOnManagedInstanceSortOrderAsc,
	"DESC": ListWindowsUpdatesInstalledOnManagedInstanceSortOrderDesc,
}

// GetListWindowsUpdatesInstalledOnManagedInstanceSortOrderEnumValues Enumerates the set of values for ListWindowsUpdatesInstalledOnManagedInstanceSortOrderEnum
func GetListWindowsUpdatesInstalledOnManagedInstanceSortOrderEnumValues() []ListWindowsUpdatesInstalledOnManagedInstanceSortOrderEnum {
	values := make([]ListWindowsUpdatesInstalledOnManagedInstanceSortOrderEnum, 0)
	for _, v := range mappingListWindowsUpdatesInstalledOnManagedInstanceSortOrder {
		values = append(values, v)
	}
	return values
}

// ListWindowsUpdatesInstalledOnManagedInstanceSortByEnum Enum with underlying type: string
type ListWindowsUpdatesInstalledOnManagedInstanceSortByEnum string

// Set of constants representing the allowable values for ListWindowsUpdatesInstalledOnManagedInstanceSortByEnum
const (
	ListWindowsUpdatesInstalledOnManagedInstanceSortByTimecreated ListWindowsUpdatesInstalledOnManagedInstanceSortByEnum = "TIMECREATED"
	ListWindowsUpdatesInstalledOnManagedInstanceSortByDisplayname ListWindowsUpdatesInstalledOnManagedInstanceSortByEnum = "DISPLAYNAME"
)

var mappingListWindowsUpdatesInstalledOnManagedInstanceSortBy = map[string]ListWindowsUpdatesInstalledOnManagedInstanceSortByEnum{
	"TIMECREATED": ListWindowsUpdatesInstalledOnManagedInstanceSortByTimecreated,
	"DISPLAYNAME": ListWindowsUpdatesInstalledOnManagedInstanceSortByDisplayname,
}

// GetListWindowsUpdatesInstalledOnManagedInstanceSortByEnumValues Enumerates the set of values for ListWindowsUpdatesInstalledOnManagedInstanceSortByEnum
func GetListWindowsUpdatesInstalledOnManagedInstanceSortByEnumValues() []ListWindowsUpdatesInstalledOnManagedInstanceSortByEnum {
	values := make([]ListWindowsUpdatesInstalledOnManagedInstanceSortByEnum, 0)
	for _, v := range mappingListWindowsUpdatesInstalledOnManagedInstanceSortBy {
		values = append(values, v)
	}
	return values
}
