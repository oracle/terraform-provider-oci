// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListClusterNamespacesRequest wrapper for the ListClusterNamespaces operation
type ListClusterNamespacesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources with lifecycleState matching the given lifecycleState.
	LifecycleState ClusterNamespaceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The name to filter on.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// unique ClusterNamespace identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The optional order in which to sort the results.
	SortOrder ListClusterNamespacesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The optional field to sort the results by.
	SortBy ListClusterNamespacesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListClusterNamespacesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListClusterNamespacesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListClusterNamespacesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListClusterNamespacesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListClusterNamespacesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClusterNamespaceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetClusterNamespaceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListClusterNamespacesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListClusterNamespacesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListClusterNamespacesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListClusterNamespacesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListClusterNamespacesResponse wrapper for the ListClusterNamespaces operation
type ListClusterNamespacesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ClusterNamespaceCollection instances
	ClusterNamespaceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListClusterNamespacesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListClusterNamespacesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListClusterNamespacesSortOrderEnum Enum with underlying type: string
type ListClusterNamespacesSortOrderEnum string

// Set of constants representing the allowable values for ListClusterNamespacesSortOrderEnum
const (
	ListClusterNamespacesSortOrderAsc  ListClusterNamespacesSortOrderEnum = "ASC"
	ListClusterNamespacesSortOrderDesc ListClusterNamespacesSortOrderEnum = "DESC"
)

var mappingListClusterNamespacesSortOrderEnum = map[string]ListClusterNamespacesSortOrderEnum{
	"ASC":  ListClusterNamespacesSortOrderAsc,
	"DESC": ListClusterNamespacesSortOrderDesc,
}

var mappingListClusterNamespacesSortOrderEnumLowerCase = map[string]ListClusterNamespacesSortOrderEnum{
	"asc":  ListClusterNamespacesSortOrderAsc,
	"desc": ListClusterNamespacesSortOrderDesc,
}

// GetListClusterNamespacesSortOrderEnumValues Enumerates the set of values for ListClusterNamespacesSortOrderEnum
func GetListClusterNamespacesSortOrderEnumValues() []ListClusterNamespacesSortOrderEnum {
	values := make([]ListClusterNamespacesSortOrderEnum, 0)
	for _, v := range mappingListClusterNamespacesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListClusterNamespacesSortOrderEnumStringValues Enumerates the set of values in String for ListClusterNamespacesSortOrderEnum
func GetListClusterNamespacesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListClusterNamespacesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClusterNamespacesSortOrderEnum(val string) (ListClusterNamespacesSortOrderEnum, bool) {
	enum, ok := mappingListClusterNamespacesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListClusterNamespacesSortByEnum Enum with underlying type: string
type ListClusterNamespacesSortByEnum string

// Set of constants representing the allowable values for ListClusterNamespacesSortByEnum
const (
	ListClusterNamespacesSortById          ListClusterNamespacesSortByEnum = "ID"
	ListClusterNamespacesSortByName        ListClusterNamespacesSortByEnum = "NAME"
	ListClusterNamespacesSortByTimeCreated ListClusterNamespacesSortByEnum = "TIME_CREATED"
)

var mappingListClusterNamespacesSortByEnum = map[string]ListClusterNamespacesSortByEnum{
	"ID":           ListClusterNamespacesSortById,
	"NAME":         ListClusterNamespacesSortByName,
	"TIME_CREATED": ListClusterNamespacesSortByTimeCreated,
}

var mappingListClusterNamespacesSortByEnumLowerCase = map[string]ListClusterNamespacesSortByEnum{
	"id":           ListClusterNamespacesSortById,
	"name":         ListClusterNamespacesSortByName,
	"time_created": ListClusterNamespacesSortByTimeCreated,
}

// GetListClusterNamespacesSortByEnumValues Enumerates the set of values for ListClusterNamespacesSortByEnum
func GetListClusterNamespacesSortByEnumValues() []ListClusterNamespacesSortByEnum {
	values := make([]ListClusterNamespacesSortByEnum, 0)
	for _, v := range mappingListClusterNamespacesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListClusterNamespacesSortByEnumStringValues Enumerates the set of values in String for ListClusterNamespacesSortByEnum
func GetListClusterNamespacesSortByEnumStringValues() []string {
	return []string{
		"ID",
		"NAME",
		"TIME_CREATED",
	}
}

// GetMappingListClusterNamespacesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClusterNamespacesSortByEnum(val string) (ListClusterNamespacesSortByEnum, bool) {
	enum, ok := mappingListClusterNamespacesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
