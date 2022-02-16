// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListResourceTypesRequest wrapper for the ListResourceTypes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResourceTypes.go.html to see an example of how to use ListResourceTypesRequest.
type ListResourceTypesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListResourceTypesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListResourceTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName is ascending. If no value is specified displayName is default.
	SortBy ListResourceTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourceTypesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResourceTypesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListResourceTypesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceTypesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourceTypesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceTypesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourceTypesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourceTypesResponse wrapper for the ListResourceTypes operation
type ListResourceTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceTypeCollection instances
	ResourceTypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResourceTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceTypesLifecycleStateEnum Enum with underlying type: string
type ListResourceTypesLifecycleStateEnum string

// Set of constants representing the allowable values for ListResourceTypesLifecycleStateEnum
const (
	ListResourceTypesLifecycleStateCreating ListResourceTypesLifecycleStateEnum = "CREATING"
	ListResourceTypesLifecycleStateUpdating ListResourceTypesLifecycleStateEnum = "UPDATING"
	ListResourceTypesLifecycleStateActive   ListResourceTypesLifecycleStateEnum = "ACTIVE"
	ListResourceTypesLifecycleStateInactive ListResourceTypesLifecycleStateEnum = "INACTIVE"
	ListResourceTypesLifecycleStateDeleting ListResourceTypesLifecycleStateEnum = "DELETING"
	ListResourceTypesLifecycleStateDeleted  ListResourceTypesLifecycleStateEnum = "DELETED"
	ListResourceTypesLifecycleStateFailed   ListResourceTypesLifecycleStateEnum = "FAILED"
)

var mappingListResourceTypesLifecycleStateEnum = map[string]ListResourceTypesLifecycleStateEnum{
	"CREATING": ListResourceTypesLifecycleStateCreating,
	"UPDATING": ListResourceTypesLifecycleStateUpdating,
	"ACTIVE":   ListResourceTypesLifecycleStateActive,
	"INACTIVE": ListResourceTypesLifecycleStateInactive,
	"DELETING": ListResourceTypesLifecycleStateDeleting,
	"DELETED":  ListResourceTypesLifecycleStateDeleted,
	"FAILED":   ListResourceTypesLifecycleStateFailed,
}

// GetListResourceTypesLifecycleStateEnumValues Enumerates the set of values for ListResourceTypesLifecycleStateEnum
func GetListResourceTypesLifecycleStateEnumValues() []ListResourceTypesLifecycleStateEnum {
	values := make([]ListResourceTypesLifecycleStateEnum, 0)
	for _, v := range mappingListResourceTypesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceTypesLifecycleStateEnumStringValues Enumerates the set of values in String for ListResourceTypesLifecycleStateEnum
func GetListResourceTypesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListResourceTypesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceTypesLifecycleStateEnum(val string) (ListResourceTypesLifecycleStateEnum, bool) {
	mappingListResourceTypesLifecycleStateEnumIgnoreCase := make(map[string]ListResourceTypesLifecycleStateEnum)
	for k, v := range mappingListResourceTypesLifecycleStateEnum {
		mappingListResourceTypesLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListResourceTypesLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceTypesSortOrderEnum Enum with underlying type: string
type ListResourceTypesSortOrderEnum string

// Set of constants representing the allowable values for ListResourceTypesSortOrderEnum
const (
	ListResourceTypesSortOrderAsc  ListResourceTypesSortOrderEnum = "ASC"
	ListResourceTypesSortOrderDesc ListResourceTypesSortOrderEnum = "DESC"
)

var mappingListResourceTypesSortOrderEnum = map[string]ListResourceTypesSortOrderEnum{
	"ASC":  ListResourceTypesSortOrderAsc,
	"DESC": ListResourceTypesSortOrderDesc,
}

// GetListResourceTypesSortOrderEnumValues Enumerates the set of values for ListResourceTypesSortOrderEnum
func GetListResourceTypesSortOrderEnumValues() []ListResourceTypesSortOrderEnum {
	values := make([]ListResourceTypesSortOrderEnum, 0)
	for _, v := range mappingListResourceTypesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceTypesSortOrderEnumStringValues Enumerates the set of values in String for ListResourceTypesSortOrderEnum
func GetListResourceTypesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourceTypesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceTypesSortOrderEnum(val string) (ListResourceTypesSortOrderEnum, bool) {
	mappingListResourceTypesSortOrderEnumIgnoreCase := make(map[string]ListResourceTypesSortOrderEnum)
	for k, v := range mappingListResourceTypesSortOrderEnum {
		mappingListResourceTypesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListResourceTypesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceTypesSortByEnum Enum with underlying type: string
type ListResourceTypesSortByEnum string

// Set of constants representing the allowable values for ListResourceTypesSortByEnum
const (
	ListResourceTypesSortByDisplayname ListResourceTypesSortByEnum = "displayName"
	ListResourceTypesSortByRisklevel   ListResourceTypesSortByEnum = "riskLevel"
)

var mappingListResourceTypesSortByEnum = map[string]ListResourceTypesSortByEnum{
	"displayName": ListResourceTypesSortByDisplayname,
	"riskLevel":   ListResourceTypesSortByRisklevel,
}

// GetListResourceTypesSortByEnumValues Enumerates the set of values for ListResourceTypesSortByEnum
func GetListResourceTypesSortByEnumValues() []ListResourceTypesSortByEnum {
	values := make([]ListResourceTypesSortByEnum, 0)
	for _, v := range mappingListResourceTypesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceTypesSortByEnumStringValues Enumerates the set of values in String for ListResourceTypesSortByEnum
func GetListResourceTypesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"riskLevel",
	}
}

// GetMappingListResourceTypesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceTypesSortByEnum(val string) (ListResourceTypesSortByEnum, bool) {
	mappingListResourceTypesSortByEnumIgnoreCase := make(map[string]ListResourceTypesSortByEnum)
	for k, v := range mappingListResourceTypesSortByEnum {
		mappingListResourceTypesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListResourceTypesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
