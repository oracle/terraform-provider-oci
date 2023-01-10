// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMeshesRequest wrapper for the ListMeshes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListMeshes.go.html to see an example of how to use ListMeshesRequest.
type ListMeshesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire displayName given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMeshesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListMeshesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the life cycle state given.
	LifecycleState MeshLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Mesh identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMeshesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMeshesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMeshesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMeshesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMeshesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMeshesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMeshesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMeshesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMeshesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMeshLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMeshLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMeshesResponse wrapper for the ListMeshes operation
type ListMeshesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MeshCollection instances
	MeshCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMeshesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMeshesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMeshesSortOrderEnum Enum with underlying type: string
type ListMeshesSortOrderEnum string

// Set of constants representing the allowable values for ListMeshesSortOrderEnum
const (
	ListMeshesSortOrderAsc  ListMeshesSortOrderEnum = "ASC"
	ListMeshesSortOrderDesc ListMeshesSortOrderEnum = "DESC"
)

var mappingListMeshesSortOrderEnum = map[string]ListMeshesSortOrderEnum{
	"ASC":  ListMeshesSortOrderAsc,
	"DESC": ListMeshesSortOrderDesc,
}

var mappingListMeshesSortOrderEnumLowerCase = map[string]ListMeshesSortOrderEnum{
	"asc":  ListMeshesSortOrderAsc,
	"desc": ListMeshesSortOrderDesc,
}

// GetListMeshesSortOrderEnumValues Enumerates the set of values for ListMeshesSortOrderEnum
func GetListMeshesSortOrderEnumValues() []ListMeshesSortOrderEnum {
	values := make([]ListMeshesSortOrderEnum, 0)
	for _, v := range mappingListMeshesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMeshesSortOrderEnumStringValues Enumerates the set of values in String for ListMeshesSortOrderEnum
func GetListMeshesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMeshesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMeshesSortOrderEnum(val string) (ListMeshesSortOrderEnum, bool) {
	enum, ok := mappingListMeshesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMeshesSortByEnum Enum with underlying type: string
type ListMeshesSortByEnum string

// Set of constants representing the allowable values for ListMeshesSortByEnum
const (
	ListMeshesSortById          ListMeshesSortByEnum = "id"
	ListMeshesSortByTimecreated ListMeshesSortByEnum = "timeCreated"
	ListMeshesSortByDisplayname ListMeshesSortByEnum = "displayName"
)

var mappingListMeshesSortByEnum = map[string]ListMeshesSortByEnum{
	"id":          ListMeshesSortById,
	"timeCreated": ListMeshesSortByTimecreated,
	"displayName": ListMeshesSortByDisplayname,
}

var mappingListMeshesSortByEnumLowerCase = map[string]ListMeshesSortByEnum{
	"id":          ListMeshesSortById,
	"timecreated": ListMeshesSortByTimecreated,
	"displayname": ListMeshesSortByDisplayname,
}

// GetListMeshesSortByEnumValues Enumerates the set of values for ListMeshesSortByEnum
func GetListMeshesSortByEnumValues() []ListMeshesSortByEnum {
	values := make([]ListMeshesSortByEnum, 0)
	for _, v := range mappingListMeshesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMeshesSortByEnumStringValues Enumerates the set of values in String for ListMeshesSortByEnum
func GetListMeshesSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMeshesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMeshesSortByEnum(val string) (ListMeshesSortByEnum, bool) {
	enum, ok := mappingListMeshesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
