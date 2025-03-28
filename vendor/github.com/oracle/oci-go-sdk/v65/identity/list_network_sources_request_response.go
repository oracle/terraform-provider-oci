// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListNetworkSourcesRequest wrapper for the ListNetworkSources operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListNetworkSources.go.html to see an example of how to use ListNetworkSourcesRequest.
type ListNetworkSourcesRequest struct {

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListNetworkSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListNetworkSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState NetworkSourcesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNetworkSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNetworkSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNetworkSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNetworkSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNetworkSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNetworkSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNetworkSourcesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNetworkSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNetworkSourcesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetNetworkSourcesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNetworkSourcesResponse wrapper for the ListNetworkSources operation
type ListNetworkSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []NetworkSourcesSummary instances
	Items []NetworkSourcesSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNetworkSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNetworkSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNetworkSourcesSortByEnum Enum with underlying type: string
type ListNetworkSourcesSortByEnum string

// Set of constants representing the allowable values for ListNetworkSourcesSortByEnum
const (
	ListNetworkSourcesSortByTimecreated ListNetworkSourcesSortByEnum = "TIMECREATED"
	ListNetworkSourcesSortByName        ListNetworkSourcesSortByEnum = "NAME"
)

var mappingListNetworkSourcesSortByEnum = map[string]ListNetworkSourcesSortByEnum{
	"TIMECREATED": ListNetworkSourcesSortByTimecreated,
	"NAME":        ListNetworkSourcesSortByName,
}

var mappingListNetworkSourcesSortByEnumLowerCase = map[string]ListNetworkSourcesSortByEnum{
	"timecreated": ListNetworkSourcesSortByTimecreated,
	"name":        ListNetworkSourcesSortByName,
}

// GetListNetworkSourcesSortByEnumValues Enumerates the set of values for ListNetworkSourcesSortByEnum
func GetListNetworkSourcesSortByEnumValues() []ListNetworkSourcesSortByEnum {
	values := make([]ListNetworkSourcesSortByEnum, 0)
	for _, v := range mappingListNetworkSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkSourcesSortByEnumStringValues Enumerates the set of values in String for ListNetworkSourcesSortByEnum
func GetListNetworkSourcesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListNetworkSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkSourcesSortByEnum(val string) (ListNetworkSourcesSortByEnum, bool) {
	enum, ok := mappingListNetworkSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkSourcesSortOrderEnum Enum with underlying type: string
type ListNetworkSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListNetworkSourcesSortOrderEnum
const (
	ListNetworkSourcesSortOrderAsc  ListNetworkSourcesSortOrderEnum = "ASC"
	ListNetworkSourcesSortOrderDesc ListNetworkSourcesSortOrderEnum = "DESC"
)

var mappingListNetworkSourcesSortOrderEnum = map[string]ListNetworkSourcesSortOrderEnum{
	"ASC":  ListNetworkSourcesSortOrderAsc,
	"DESC": ListNetworkSourcesSortOrderDesc,
}

var mappingListNetworkSourcesSortOrderEnumLowerCase = map[string]ListNetworkSourcesSortOrderEnum{
	"asc":  ListNetworkSourcesSortOrderAsc,
	"desc": ListNetworkSourcesSortOrderDesc,
}

// GetListNetworkSourcesSortOrderEnumValues Enumerates the set of values for ListNetworkSourcesSortOrderEnum
func GetListNetworkSourcesSortOrderEnumValues() []ListNetworkSourcesSortOrderEnum {
	values := make([]ListNetworkSourcesSortOrderEnum, 0)
	for _, v := range mappingListNetworkSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListNetworkSourcesSortOrderEnum
func GetListNetworkSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNetworkSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkSourcesSortOrderEnum(val string) (ListNetworkSourcesSortOrderEnum, bool) {
	enum, ok := mappingListNetworkSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
