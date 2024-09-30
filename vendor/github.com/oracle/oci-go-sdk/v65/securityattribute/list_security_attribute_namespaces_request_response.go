// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package securityattribute

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSecurityAttributeNamespacesRequest wrapper for the ListSecurityAttributeNamespaces operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/ListSecurityAttributeNamespaces.go.html to see an example of how to use ListSecurityAttributeNamespacesRequest.
type ListSecurityAttributeNamespacesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSecurityAttributeNamespacesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListSecurityAttributeNamespacesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// An optional boolean parameter indicating whether to retrieve all security attribute namespaces in subcompartments. If this
	// parameter is not specified, only the namespaces defined in the specified compartment are retrieved.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState SecurityAttributeNamespaceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecurityAttributeNamespacesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecurityAttributeNamespacesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecurityAttributeNamespacesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecurityAttributeNamespacesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecurityAttributeNamespacesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSecurityAttributeNamespacesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSecurityAttributeNamespacesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityAttributeNamespacesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSecurityAttributeNamespacesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityAttributeNamespaceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetSecurityAttributeNamespaceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecurityAttributeNamespacesResponse wrapper for the ListSecurityAttributeNamespaces operation
type ListSecurityAttributeNamespacesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SecurityAttributeNamespaceSummary instances
	Items []SecurityAttributeNamespaceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of namespaces. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSecurityAttributeNamespacesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecurityAttributeNamespacesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecurityAttributeNamespacesSortOrderEnum Enum with underlying type: string
type ListSecurityAttributeNamespacesSortOrderEnum string

// Set of constants representing the allowable values for ListSecurityAttributeNamespacesSortOrderEnum
const (
	ListSecurityAttributeNamespacesSortOrderAsc  ListSecurityAttributeNamespacesSortOrderEnum = "ASC"
	ListSecurityAttributeNamespacesSortOrderDesc ListSecurityAttributeNamespacesSortOrderEnum = "DESC"
)

var mappingListSecurityAttributeNamespacesSortOrderEnum = map[string]ListSecurityAttributeNamespacesSortOrderEnum{
	"ASC":  ListSecurityAttributeNamespacesSortOrderAsc,
	"DESC": ListSecurityAttributeNamespacesSortOrderDesc,
}

var mappingListSecurityAttributeNamespacesSortOrderEnumLowerCase = map[string]ListSecurityAttributeNamespacesSortOrderEnum{
	"asc":  ListSecurityAttributeNamespacesSortOrderAsc,
	"desc": ListSecurityAttributeNamespacesSortOrderDesc,
}

// GetListSecurityAttributeNamespacesSortOrderEnumValues Enumerates the set of values for ListSecurityAttributeNamespacesSortOrderEnum
func GetListSecurityAttributeNamespacesSortOrderEnumValues() []ListSecurityAttributeNamespacesSortOrderEnum {
	values := make([]ListSecurityAttributeNamespacesSortOrderEnum, 0)
	for _, v := range mappingListSecurityAttributeNamespacesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityAttributeNamespacesSortOrderEnumStringValues Enumerates the set of values in String for ListSecurityAttributeNamespacesSortOrderEnum
func GetListSecurityAttributeNamespacesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSecurityAttributeNamespacesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityAttributeNamespacesSortOrderEnum(val string) (ListSecurityAttributeNamespacesSortOrderEnum, bool) {
	enum, ok := mappingListSecurityAttributeNamespacesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityAttributeNamespacesSortByEnum Enum with underlying type: string
type ListSecurityAttributeNamespacesSortByEnum string

// Set of constants representing the allowable values for ListSecurityAttributeNamespacesSortByEnum
const (
	ListSecurityAttributeNamespacesSortByTimecreated ListSecurityAttributeNamespacesSortByEnum = "TIMECREATED"
	ListSecurityAttributeNamespacesSortByName        ListSecurityAttributeNamespacesSortByEnum = "NAME"
)

var mappingListSecurityAttributeNamespacesSortByEnum = map[string]ListSecurityAttributeNamespacesSortByEnum{
	"TIMECREATED": ListSecurityAttributeNamespacesSortByTimecreated,
	"NAME":        ListSecurityAttributeNamespacesSortByName,
}

var mappingListSecurityAttributeNamespacesSortByEnumLowerCase = map[string]ListSecurityAttributeNamespacesSortByEnum{
	"timecreated": ListSecurityAttributeNamespacesSortByTimecreated,
	"name":        ListSecurityAttributeNamespacesSortByName,
}

// GetListSecurityAttributeNamespacesSortByEnumValues Enumerates the set of values for ListSecurityAttributeNamespacesSortByEnum
func GetListSecurityAttributeNamespacesSortByEnumValues() []ListSecurityAttributeNamespacesSortByEnum {
	values := make([]ListSecurityAttributeNamespacesSortByEnum, 0)
	for _, v := range mappingListSecurityAttributeNamespacesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityAttributeNamespacesSortByEnumStringValues Enumerates the set of values in String for ListSecurityAttributeNamespacesSortByEnum
func GetListSecurityAttributeNamespacesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListSecurityAttributeNamespacesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityAttributeNamespacesSortByEnum(val string) (ListSecurityAttributeNamespacesSortByEnum, bool) {
	enum, ok := mappingListSecurityAttributeNamespacesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
