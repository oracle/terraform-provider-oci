// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package vbsinst

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListVbsInstancesRequest wrapper for the ListVbsInstances operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vbsinst/ListVbsInstances.go.html to see an example of how to use ListVbsInstancesRequest.
type ListVbsInstancesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListVbsInstancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// unique VbsInstance identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListVbsInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListVbsInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVbsInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVbsInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVbsInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVbsInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVbsInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListVbsInstancesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListVbsInstancesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVbsInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVbsInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVbsInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVbsInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVbsInstancesResponse wrapper for the ListVbsInstances operation
type ListVbsInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VbsInstanceSummaryCollection instances
	VbsInstanceSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVbsInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVbsInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVbsInstancesLifecycleStateEnum Enum with underlying type: string
type ListVbsInstancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListVbsInstancesLifecycleStateEnum
const (
	ListVbsInstancesLifecycleStateCreating ListVbsInstancesLifecycleStateEnum = "CREATING"
	ListVbsInstancesLifecycleStateUpdating ListVbsInstancesLifecycleStateEnum = "UPDATING"
	ListVbsInstancesLifecycleStateActive   ListVbsInstancesLifecycleStateEnum = "ACTIVE"
	ListVbsInstancesLifecycleStateDeleting ListVbsInstancesLifecycleStateEnum = "DELETING"
	ListVbsInstancesLifecycleStateDeleted  ListVbsInstancesLifecycleStateEnum = "DELETED"
	ListVbsInstancesLifecycleStateFailed   ListVbsInstancesLifecycleStateEnum = "FAILED"
)

var mappingListVbsInstancesLifecycleStateEnum = map[string]ListVbsInstancesLifecycleStateEnum{
	"CREATING": ListVbsInstancesLifecycleStateCreating,
	"UPDATING": ListVbsInstancesLifecycleStateUpdating,
	"ACTIVE":   ListVbsInstancesLifecycleStateActive,
	"DELETING": ListVbsInstancesLifecycleStateDeleting,
	"DELETED":  ListVbsInstancesLifecycleStateDeleted,
	"FAILED":   ListVbsInstancesLifecycleStateFailed,
}

var mappingListVbsInstancesLifecycleStateEnumLowerCase = map[string]ListVbsInstancesLifecycleStateEnum{
	"creating": ListVbsInstancesLifecycleStateCreating,
	"updating": ListVbsInstancesLifecycleStateUpdating,
	"active":   ListVbsInstancesLifecycleStateActive,
	"deleting": ListVbsInstancesLifecycleStateDeleting,
	"deleted":  ListVbsInstancesLifecycleStateDeleted,
	"failed":   ListVbsInstancesLifecycleStateFailed,
}

// GetListVbsInstancesLifecycleStateEnumValues Enumerates the set of values for ListVbsInstancesLifecycleStateEnum
func GetListVbsInstancesLifecycleStateEnumValues() []ListVbsInstancesLifecycleStateEnum {
	values := make([]ListVbsInstancesLifecycleStateEnum, 0)
	for _, v := range mappingListVbsInstancesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListVbsInstancesLifecycleStateEnumStringValues Enumerates the set of values in String for ListVbsInstancesLifecycleStateEnum
func GetListVbsInstancesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListVbsInstancesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVbsInstancesLifecycleStateEnum(val string) (ListVbsInstancesLifecycleStateEnum, bool) {
	enum, ok := mappingListVbsInstancesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVbsInstancesSortOrderEnum Enum with underlying type: string
type ListVbsInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListVbsInstancesSortOrderEnum
const (
	ListVbsInstancesSortOrderAsc  ListVbsInstancesSortOrderEnum = "ASC"
	ListVbsInstancesSortOrderDesc ListVbsInstancesSortOrderEnum = "DESC"
)

var mappingListVbsInstancesSortOrderEnum = map[string]ListVbsInstancesSortOrderEnum{
	"ASC":  ListVbsInstancesSortOrderAsc,
	"DESC": ListVbsInstancesSortOrderDesc,
}

var mappingListVbsInstancesSortOrderEnumLowerCase = map[string]ListVbsInstancesSortOrderEnum{
	"asc":  ListVbsInstancesSortOrderAsc,
	"desc": ListVbsInstancesSortOrderDesc,
}

// GetListVbsInstancesSortOrderEnumValues Enumerates the set of values for ListVbsInstancesSortOrderEnum
func GetListVbsInstancesSortOrderEnumValues() []ListVbsInstancesSortOrderEnum {
	values := make([]ListVbsInstancesSortOrderEnum, 0)
	for _, v := range mappingListVbsInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVbsInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListVbsInstancesSortOrderEnum
func GetListVbsInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVbsInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVbsInstancesSortOrderEnum(val string) (ListVbsInstancesSortOrderEnum, bool) {
	enum, ok := mappingListVbsInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVbsInstancesSortByEnum Enum with underlying type: string
type ListVbsInstancesSortByEnum string

// Set of constants representing the allowable values for ListVbsInstancesSortByEnum
const (
	ListVbsInstancesSortByTimecreated ListVbsInstancesSortByEnum = "timeCreated"
	ListVbsInstancesSortByDisplayname ListVbsInstancesSortByEnum = "displayName"
)

var mappingListVbsInstancesSortByEnum = map[string]ListVbsInstancesSortByEnum{
	"timeCreated": ListVbsInstancesSortByTimecreated,
	"displayName": ListVbsInstancesSortByDisplayname,
}

var mappingListVbsInstancesSortByEnumLowerCase = map[string]ListVbsInstancesSortByEnum{
	"timecreated": ListVbsInstancesSortByTimecreated,
	"displayname": ListVbsInstancesSortByDisplayname,
}

// GetListVbsInstancesSortByEnumValues Enumerates the set of values for ListVbsInstancesSortByEnum
func GetListVbsInstancesSortByEnumValues() []ListVbsInstancesSortByEnum {
	values := make([]ListVbsInstancesSortByEnum, 0)
	for _, v := range mappingListVbsInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVbsInstancesSortByEnumStringValues Enumerates the set of values in String for ListVbsInstancesSortByEnum
func GetListVbsInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListVbsInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVbsInstancesSortByEnum(val string) (ListVbsInstancesSortByEnum, bool) {
	enum, ok := mappingListVbsInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
