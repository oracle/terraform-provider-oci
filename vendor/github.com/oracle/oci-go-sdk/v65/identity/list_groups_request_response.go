// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListGroupsRequest wrapper for the ListGroups operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListGroups.go.html to see an example of how to use ListGroupsRequest.
type ListGroupsRequest struct {

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
	SortBy ListGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState GroupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGroupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetGroupLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGroupsResponse wrapper for the ListGroups operation
type ListGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []Group instances
	Items []Group `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGroupsSortByEnum Enum with underlying type: string
type ListGroupsSortByEnum string

// Set of constants representing the allowable values for ListGroupsSortByEnum
const (
	ListGroupsSortByTimecreated ListGroupsSortByEnum = "TIMECREATED"
	ListGroupsSortByName        ListGroupsSortByEnum = "NAME"
)

var mappingListGroupsSortByEnum = map[string]ListGroupsSortByEnum{
	"TIMECREATED": ListGroupsSortByTimecreated,
	"NAME":        ListGroupsSortByName,
}

var mappingListGroupsSortByEnumLowerCase = map[string]ListGroupsSortByEnum{
	"timecreated": ListGroupsSortByTimecreated,
	"name":        ListGroupsSortByName,
}

// GetListGroupsSortByEnumValues Enumerates the set of values for ListGroupsSortByEnum
func GetListGroupsSortByEnumValues() []ListGroupsSortByEnum {
	values := make([]ListGroupsSortByEnum, 0)
	for _, v := range mappingListGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGroupsSortByEnumStringValues Enumerates the set of values in String for ListGroupsSortByEnum
func GetListGroupsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGroupsSortByEnum(val string) (ListGroupsSortByEnum, bool) {
	enum, ok := mappingListGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGroupsSortOrderEnum Enum with underlying type: string
type ListGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListGroupsSortOrderEnum
const (
	ListGroupsSortOrderAsc  ListGroupsSortOrderEnum = "ASC"
	ListGroupsSortOrderDesc ListGroupsSortOrderEnum = "DESC"
)

var mappingListGroupsSortOrderEnum = map[string]ListGroupsSortOrderEnum{
	"ASC":  ListGroupsSortOrderAsc,
	"DESC": ListGroupsSortOrderDesc,
}

var mappingListGroupsSortOrderEnumLowerCase = map[string]ListGroupsSortOrderEnum{
	"asc":  ListGroupsSortOrderAsc,
	"desc": ListGroupsSortOrderDesc,
}

// GetListGroupsSortOrderEnumValues Enumerates the set of values for ListGroupsSortOrderEnum
func GetListGroupsSortOrderEnumValues() []ListGroupsSortOrderEnum {
	values := make([]ListGroupsSortOrderEnum, 0)
	for _, v := range mappingListGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListGroupsSortOrderEnum
func GetListGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGroupsSortOrderEnum(val string) (ListGroupsSortOrderEnum, bool) {
	enum, ok := mappingListGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
