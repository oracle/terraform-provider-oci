// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListZprTagNamespacesRequest wrapper for the ListZprTagNamespaces operation
type ListZprTagNamespacesRequest struct {

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListZprTagNamespacesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListZprTagNamespacesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// An optional boolean parameter indicating whether to retrieve all zpr tag namespaces in subcompartments. If this
	// parameter is not specified, only the zpr tag namespaces defined in the specified compartment are retrieved.
	ShouldIncludeSubcompartments *bool `mandatory:"false" contributesTo:"query" name:"shouldIncludeSubcompartments"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState ZprTagNamespaceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListZprTagNamespacesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListZprTagNamespacesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListZprTagNamespacesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListZprTagNamespacesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListZprTagNamespacesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListZprTagNamespacesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListZprTagNamespacesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListZprTagNamespacesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListZprTagNamespacesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingZprTagNamespaceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetZprTagNamespaceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListZprTagNamespacesResponse wrapper for the ListZprTagNamespaces operation
type ListZprTagNamespacesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ZprTagNamespaceCollection instances
	ZprTagNamespaceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of tagNamespaces. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListZprTagNamespacesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListZprTagNamespacesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListZprTagNamespacesSortOrderEnum Enum with underlying type: string
type ListZprTagNamespacesSortOrderEnum string

// Set of constants representing the allowable values for ListZprTagNamespacesSortOrderEnum
const (
	ListZprTagNamespacesSortOrderAsc  ListZprTagNamespacesSortOrderEnum = "ASC"
	ListZprTagNamespacesSortOrderDesc ListZprTagNamespacesSortOrderEnum = "DESC"
)

var mappingListZprTagNamespacesSortOrderEnum = map[string]ListZprTagNamespacesSortOrderEnum{
	"ASC":  ListZprTagNamespacesSortOrderAsc,
	"DESC": ListZprTagNamespacesSortOrderDesc,
}

var mappingListZprTagNamespacesSortOrderEnumLowerCase = map[string]ListZprTagNamespacesSortOrderEnum{
	"asc":  ListZprTagNamespacesSortOrderAsc,
	"desc": ListZprTagNamespacesSortOrderDesc,
}

// GetListZprTagNamespacesSortOrderEnumValues Enumerates the set of values for ListZprTagNamespacesSortOrderEnum
func GetListZprTagNamespacesSortOrderEnumValues() []ListZprTagNamespacesSortOrderEnum {
	values := make([]ListZprTagNamespacesSortOrderEnum, 0)
	for _, v := range mappingListZprTagNamespacesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprTagNamespacesSortOrderEnumStringValues Enumerates the set of values in String for ListZprTagNamespacesSortOrderEnum
func GetListZprTagNamespacesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListZprTagNamespacesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprTagNamespacesSortOrderEnum(val string) (ListZprTagNamespacesSortOrderEnum, bool) {
	enum, ok := mappingListZprTagNamespacesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListZprTagNamespacesSortByEnum Enum with underlying type: string
type ListZprTagNamespacesSortByEnum string

// Set of constants representing the allowable values for ListZprTagNamespacesSortByEnum
const (
	ListZprTagNamespacesSortByTimecreated ListZprTagNamespacesSortByEnum = "TIMECREATED"
	ListZprTagNamespacesSortByName        ListZprTagNamespacesSortByEnum = "NAME"
)

var mappingListZprTagNamespacesSortByEnum = map[string]ListZprTagNamespacesSortByEnum{
	"TIMECREATED": ListZprTagNamespacesSortByTimecreated,
	"NAME":        ListZprTagNamespacesSortByName,
}

var mappingListZprTagNamespacesSortByEnumLowerCase = map[string]ListZprTagNamespacesSortByEnum{
	"timecreated": ListZprTagNamespacesSortByTimecreated,
	"name":        ListZprTagNamespacesSortByName,
}

// GetListZprTagNamespacesSortByEnumValues Enumerates the set of values for ListZprTagNamespacesSortByEnum
func GetListZprTagNamespacesSortByEnumValues() []ListZprTagNamespacesSortByEnum {
	values := make([]ListZprTagNamespacesSortByEnum, 0)
	for _, v := range mappingListZprTagNamespacesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprTagNamespacesSortByEnumStringValues Enumerates the set of values in String for ListZprTagNamespacesSortByEnum
func GetListZprTagNamespacesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListZprTagNamespacesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprTagNamespacesSortByEnum(val string) (ListZprTagNamespacesSortByEnum, bool) {
	enum, ok := mappingListZprTagNamespacesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
