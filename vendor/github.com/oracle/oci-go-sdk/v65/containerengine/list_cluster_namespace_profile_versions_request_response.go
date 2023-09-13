// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListClusterNamespaceProfileVersionsRequest wrapper for the ListClusterNamespaceProfileVersions operation
type ListClusterNamespaceProfileVersionsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ClusterNamespaceProfileVersionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The name to filter on.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// unique ClusterNamespaceProfileVersion identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The optional order in which to sort the results.
	SortOrder ListClusterNamespaceProfileVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The optional field to sort the results by.
	SortBy ListClusterNamespaceProfileVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListClusterNamespaceProfileVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListClusterNamespaceProfileVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListClusterNamespaceProfileVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListClusterNamespaceProfileVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListClusterNamespaceProfileVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClusterNamespaceProfileVersionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetClusterNamespaceProfileVersionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListClusterNamespaceProfileVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListClusterNamespaceProfileVersionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListClusterNamespaceProfileVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListClusterNamespaceProfileVersionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListClusterNamespaceProfileVersionsResponse wrapper for the ListClusterNamespaceProfileVersions operation
type ListClusterNamespaceProfileVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ClusterNamespaceProfileVersionCollection instances
	ClusterNamespaceProfileVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListClusterNamespaceProfileVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListClusterNamespaceProfileVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListClusterNamespaceProfileVersionsSortOrderEnum Enum with underlying type: string
type ListClusterNamespaceProfileVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListClusterNamespaceProfileVersionsSortOrderEnum
const (
	ListClusterNamespaceProfileVersionsSortOrderAsc  ListClusterNamespaceProfileVersionsSortOrderEnum = "ASC"
	ListClusterNamespaceProfileVersionsSortOrderDesc ListClusterNamespaceProfileVersionsSortOrderEnum = "DESC"
)

var mappingListClusterNamespaceProfileVersionsSortOrderEnum = map[string]ListClusterNamespaceProfileVersionsSortOrderEnum{
	"ASC":  ListClusterNamespaceProfileVersionsSortOrderAsc,
	"DESC": ListClusterNamespaceProfileVersionsSortOrderDesc,
}

var mappingListClusterNamespaceProfileVersionsSortOrderEnumLowerCase = map[string]ListClusterNamespaceProfileVersionsSortOrderEnum{
	"asc":  ListClusterNamespaceProfileVersionsSortOrderAsc,
	"desc": ListClusterNamespaceProfileVersionsSortOrderDesc,
}

// GetListClusterNamespaceProfileVersionsSortOrderEnumValues Enumerates the set of values for ListClusterNamespaceProfileVersionsSortOrderEnum
func GetListClusterNamespaceProfileVersionsSortOrderEnumValues() []ListClusterNamespaceProfileVersionsSortOrderEnum {
	values := make([]ListClusterNamespaceProfileVersionsSortOrderEnum, 0)
	for _, v := range mappingListClusterNamespaceProfileVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListClusterNamespaceProfileVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListClusterNamespaceProfileVersionsSortOrderEnum
func GetListClusterNamespaceProfileVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListClusterNamespaceProfileVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClusterNamespaceProfileVersionsSortOrderEnum(val string) (ListClusterNamespaceProfileVersionsSortOrderEnum, bool) {
	enum, ok := mappingListClusterNamespaceProfileVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListClusterNamespaceProfileVersionsSortByEnum Enum with underlying type: string
type ListClusterNamespaceProfileVersionsSortByEnum string

// Set of constants representing the allowable values for ListClusterNamespaceProfileVersionsSortByEnum
const (
	ListClusterNamespaceProfileVersionsSortById          ListClusterNamespaceProfileVersionsSortByEnum = "ID"
	ListClusterNamespaceProfileVersionsSortByName        ListClusterNamespaceProfileVersionsSortByEnum = "NAME"
	ListClusterNamespaceProfileVersionsSortByTimeCreated ListClusterNamespaceProfileVersionsSortByEnum = "TIME_CREATED"
)

var mappingListClusterNamespaceProfileVersionsSortByEnum = map[string]ListClusterNamespaceProfileVersionsSortByEnum{
	"ID":           ListClusterNamespaceProfileVersionsSortById,
	"NAME":         ListClusterNamespaceProfileVersionsSortByName,
	"TIME_CREATED": ListClusterNamespaceProfileVersionsSortByTimeCreated,
}

var mappingListClusterNamespaceProfileVersionsSortByEnumLowerCase = map[string]ListClusterNamespaceProfileVersionsSortByEnum{
	"id":           ListClusterNamespaceProfileVersionsSortById,
	"name":         ListClusterNamespaceProfileVersionsSortByName,
	"time_created": ListClusterNamespaceProfileVersionsSortByTimeCreated,
}

// GetListClusterNamespaceProfileVersionsSortByEnumValues Enumerates the set of values for ListClusterNamespaceProfileVersionsSortByEnum
func GetListClusterNamespaceProfileVersionsSortByEnumValues() []ListClusterNamespaceProfileVersionsSortByEnum {
	values := make([]ListClusterNamespaceProfileVersionsSortByEnum, 0)
	for _, v := range mappingListClusterNamespaceProfileVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListClusterNamespaceProfileVersionsSortByEnumStringValues Enumerates the set of values in String for ListClusterNamespaceProfileVersionsSortByEnum
func GetListClusterNamespaceProfileVersionsSortByEnumStringValues() []string {
	return []string{
		"ID",
		"NAME",
		"TIME_CREATED",
	}
}

// GetMappingListClusterNamespaceProfileVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClusterNamespaceProfileVersionsSortByEnum(val string) (ListClusterNamespaceProfileVersionsSortByEnum, bool) {
	enum, ok := mappingListClusterNamespaceProfileVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
