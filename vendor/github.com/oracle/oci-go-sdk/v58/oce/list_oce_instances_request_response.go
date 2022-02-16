// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oce

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListOceInstancesRequest wrapper for the ListOceInstances operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oce/ListOceInstances.go.html to see an example of how to use ListOceInstancesRequest.
type ListOceInstancesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the tenancy in which to list resources.
	TenancyId *string `mandatory:"false" contributesTo:"query" name:"tenancyId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOceInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListOceInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter results on lifecycleState.
	LifecycleState ListOceInstancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOceInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOceInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOceInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOceInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOceInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOceInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOceInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOceInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOceInstancesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOceInstancesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListOceInstancesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOceInstancesResponse wrapper for the ListOceInstances operation
type ListOceInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []OceInstanceSummary instances
	Items []OceInstanceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `OceInstance`s. If this header appears in the response, then this
	// is a partial list of OceInstances. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of OceInstances.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOceInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOceInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOceInstancesSortOrderEnum Enum with underlying type: string
type ListOceInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListOceInstancesSortOrderEnum
const (
	ListOceInstancesSortOrderAsc  ListOceInstancesSortOrderEnum = "ASC"
	ListOceInstancesSortOrderDesc ListOceInstancesSortOrderEnum = "DESC"
)

var mappingListOceInstancesSortOrderEnum = map[string]ListOceInstancesSortOrderEnum{
	"ASC":  ListOceInstancesSortOrderAsc,
	"DESC": ListOceInstancesSortOrderDesc,
}

// GetListOceInstancesSortOrderEnumValues Enumerates the set of values for ListOceInstancesSortOrderEnum
func GetListOceInstancesSortOrderEnumValues() []ListOceInstancesSortOrderEnum {
	values := make([]ListOceInstancesSortOrderEnum, 0)
	for _, v := range mappingListOceInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOceInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListOceInstancesSortOrderEnum
func GetListOceInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOceInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOceInstancesSortOrderEnum(val string) (ListOceInstancesSortOrderEnum, bool) {
	mappingListOceInstancesSortOrderEnumIgnoreCase := make(map[string]ListOceInstancesSortOrderEnum)
	for k, v := range mappingListOceInstancesSortOrderEnum {
		mappingListOceInstancesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListOceInstancesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListOceInstancesSortByEnum Enum with underlying type: string
type ListOceInstancesSortByEnum string

// Set of constants representing the allowable values for ListOceInstancesSortByEnum
const (
	ListOceInstancesSortByTimecreated ListOceInstancesSortByEnum = "timeCreated"
	ListOceInstancesSortByDisplayname ListOceInstancesSortByEnum = "displayName"
)

var mappingListOceInstancesSortByEnum = map[string]ListOceInstancesSortByEnum{
	"timeCreated": ListOceInstancesSortByTimecreated,
	"displayName": ListOceInstancesSortByDisplayname,
}

// GetListOceInstancesSortByEnumValues Enumerates the set of values for ListOceInstancesSortByEnum
func GetListOceInstancesSortByEnumValues() []ListOceInstancesSortByEnum {
	values := make([]ListOceInstancesSortByEnum, 0)
	for _, v := range mappingListOceInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOceInstancesSortByEnumStringValues Enumerates the set of values in String for ListOceInstancesSortByEnum
func GetListOceInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOceInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOceInstancesSortByEnum(val string) (ListOceInstancesSortByEnum, bool) {
	mappingListOceInstancesSortByEnumIgnoreCase := make(map[string]ListOceInstancesSortByEnum)
	for k, v := range mappingListOceInstancesSortByEnum {
		mappingListOceInstancesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListOceInstancesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListOceInstancesLifecycleStateEnum Enum with underlying type: string
type ListOceInstancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListOceInstancesLifecycleStateEnum
const (
	ListOceInstancesLifecycleStateCreating ListOceInstancesLifecycleStateEnum = "CREATING"
	ListOceInstancesLifecycleStateUpdating ListOceInstancesLifecycleStateEnum = "UPDATING"
	ListOceInstancesLifecycleStateActive   ListOceInstancesLifecycleStateEnum = "ACTIVE"
	ListOceInstancesLifecycleStateDeleting ListOceInstancesLifecycleStateEnum = "DELETING"
	ListOceInstancesLifecycleStateDeleted  ListOceInstancesLifecycleStateEnum = "DELETED"
	ListOceInstancesLifecycleStateFailed   ListOceInstancesLifecycleStateEnum = "FAILED"
)

var mappingListOceInstancesLifecycleStateEnum = map[string]ListOceInstancesLifecycleStateEnum{
	"CREATING": ListOceInstancesLifecycleStateCreating,
	"UPDATING": ListOceInstancesLifecycleStateUpdating,
	"ACTIVE":   ListOceInstancesLifecycleStateActive,
	"DELETING": ListOceInstancesLifecycleStateDeleting,
	"DELETED":  ListOceInstancesLifecycleStateDeleted,
	"FAILED":   ListOceInstancesLifecycleStateFailed,
}

// GetListOceInstancesLifecycleStateEnumValues Enumerates the set of values for ListOceInstancesLifecycleStateEnum
func GetListOceInstancesLifecycleStateEnumValues() []ListOceInstancesLifecycleStateEnum {
	values := make([]ListOceInstancesLifecycleStateEnum, 0)
	for _, v := range mappingListOceInstancesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListOceInstancesLifecycleStateEnumStringValues Enumerates the set of values in String for ListOceInstancesLifecycleStateEnum
func GetListOceInstancesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListOceInstancesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOceInstancesLifecycleStateEnum(val string) (ListOceInstancesLifecycleStateEnum, bool) {
	mappingListOceInstancesLifecycleStateEnumIgnoreCase := make(map[string]ListOceInstancesLifecycleStateEnum)
	for k, v := range mappingListOceInstancesLifecycleStateEnum {
		mappingListOceInstancesLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListOceInstancesLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
