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

// ListClusterNamespaceProfilesRequest wrapper for the ListClusterNamespaceProfiles operation
type ListClusterNamespaceProfilesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ClusterNamespaceProfileLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique ClusterNamespaceProfile identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The optional order in which to sort the results.
	SortOrder ListClusterNamespaceProfilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The optional field to sort the results by.
	SortBy ListClusterNamespaceProfilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListClusterNamespaceProfilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListClusterNamespaceProfilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListClusterNamespaceProfilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListClusterNamespaceProfilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListClusterNamespaceProfilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClusterNamespaceProfileLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetClusterNamespaceProfileLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListClusterNamespaceProfilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListClusterNamespaceProfilesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListClusterNamespaceProfilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListClusterNamespaceProfilesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListClusterNamespaceProfilesResponse wrapper for the ListClusterNamespaceProfiles operation
type ListClusterNamespaceProfilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ClusterNamespaceProfileCollection instances
	ClusterNamespaceProfileCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListClusterNamespaceProfilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListClusterNamespaceProfilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListClusterNamespaceProfilesSortOrderEnum Enum with underlying type: string
type ListClusterNamespaceProfilesSortOrderEnum string

// Set of constants representing the allowable values for ListClusterNamespaceProfilesSortOrderEnum
const (
	ListClusterNamespaceProfilesSortOrderAsc  ListClusterNamespaceProfilesSortOrderEnum = "ASC"
	ListClusterNamespaceProfilesSortOrderDesc ListClusterNamespaceProfilesSortOrderEnum = "DESC"
)

var mappingListClusterNamespaceProfilesSortOrderEnum = map[string]ListClusterNamespaceProfilesSortOrderEnum{
	"ASC":  ListClusterNamespaceProfilesSortOrderAsc,
	"DESC": ListClusterNamespaceProfilesSortOrderDesc,
}

var mappingListClusterNamespaceProfilesSortOrderEnumLowerCase = map[string]ListClusterNamespaceProfilesSortOrderEnum{
	"asc":  ListClusterNamespaceProfilesSortOrderAsc,
	"desc": ListClusterNamespaceProfilesSortOrderDesc,
}

// GetListClusterNamespaceProfilesSortOrderEnumValues Enumerates the set of values for ListClusterNamespaceProfilesSortOrderEnum
func GetListClusterNamespaceProfilesSortOrderEnumValues() []ListClusterNamespaceProfilesSortOrderEnum {
	values := make([]ListClusterNamespaceProfilesSortOrderEnum, 0)
	for _, v := range mappingListClusterNamespaceProfilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListClusterNamespaceProfilesSortOrderEnumStringValues Enumerates the set of values in String for ListClusterNamespaceProfilesSortOrderEnum
func GetListClusterNamespaceProfilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListClusterNamespaceProfilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClusterNamespaceProfilesSortOrderEnum(val string) (ListClusterNamespaceProfilesSortOrderEnum, bool) {
	enum, ok := mappingListClusterNamespaceProfilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListClusterNamespaceProfilesSortByEnum Enum with underlying type: string
type ListClusterNamespaceProfilesSortByEnum string

// Set of constants representing the allowable values for ListClusterNamespaceProfilesSortByEnum
const (
	ListClusterNamespaceProfilesSortById          ListClusterNamespaceProfilesSortByEnum = "ID"
	ListClusterNamespaceProfilesSortByName        ListClusterNamespaceProfilesSortByEnum = "NAME"
	ListClusterNamespaceProfilesSortByTimeCreated ListClusterNamespaceProfilesSortByEnum = "TIME_CREATED"
)

var mappingListClusterNamespaceProfilesSortByEnum = map[string]ListClusterNamespaceProfilesSortByEnum{
	"ID":           ListClusterNamespaceProfilesSortById,
	"NAME":         ListClusterNamespaceProfilesSortByName,
	"TIME_CREATED": ListClusterNamespaceProfilesSortByTimeCreated,
}

var mappingListClusterNamespaceProfilesSortByEnumLowerCase = map[string]ListClusterNamespaceProfilesSortByEnum{
	"id":           ListClusterNamespaceProfilesSortById,
	"name":         ListClusterNamespaceProfilesSortByName,
	"time_created": ListClusterNamespaceProfilesSortByTimeCreated,
}

// GetListClusterNamespaceProfilesSortByEnumValues Enumerates the set of values for ListClusterNamespaceProfilesSortByEnum
func GetListClusterNamespaceProfilesSortByEnumValues() []ListClusterNamespaceProfilesSortByEnum {
	values := make([]ListClusterNamespaceProfilesSortByEnum, 0)
	for _, v := range mappingListClusterNamespaceProfilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListClusterNamespaceProfilesSortByEnumStringValues Enumerates the set of values in String for ListClusterNamespaceProfilesSortByEnum
func GetListClusterNamespaceProfilesSortByEnumStringValues() []string {
	return []string{
		"ID",
		"NAME",
		"TIME_CREATED",
	}
}

// GetMappingListClusterNamespaceProfilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClusterNamespaceProfilesSortByEnum(val string) (ListClusterNamespaceProfilesSortByEnum, bool) {
	enum, ok := mappingListClusterNamespaceProfilesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
