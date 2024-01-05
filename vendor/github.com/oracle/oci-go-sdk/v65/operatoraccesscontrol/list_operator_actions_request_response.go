// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package operatoraccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOperatorActionsRequest wrapper for the ListOperatorActions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ListOperatorActions.go.html to see an example of how to use ListOperatorActionsRequest.
type ListOperatorActionsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only lists of resources that match the entire given service type.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// A filter to return only resources whose lifecycleState matches the given OperatorAction lifecycleState.
	LifecycleState ListOperatorActionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOperatorActionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListOperatorActionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOperatorActionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOperatorActionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOperatorActionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOperatorActionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOperatorActionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOperatorActionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListOperatorActionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOperatorActionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOperatorActionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOperatorActionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOperatorActionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOperatorActionsResponse wrapper for the ListOperatorActions operation
type ListOperatorActionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OperatorActionCollection instances
	OperatorActionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOperatorActionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOperatorActionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOperatorActionsLifecycleStateEnum Enum with underlying type: string
type ListOperatorActionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListOperatorActionsLifecycleStateEnum
const (
	ListOperatorActionsLifecycleStateActive   ListOperatorActionsLifecycleStateEnum = "ACTIVE"
	ListOperatorActionsLifecycleStateInactive ListOperatorActionsLifecycleStateEnum = "INACTIVE"
)

var mappingListOperatorActionsLifecycleStateEnum = map[string]ListOperatorActionsLifecycleStateEnum{
	"ACTIVE":   ListOperatorActionsLifecycleStateActive,
	"INACTIVE": ListOperatorActionsLifecycleStateInactive,
}

var mappingListOperatorActionsLifecycleStateEnumLowerCase = map[string]ListOperatorActionsLifecycleStateEnum{
	"active":   ListOperatorActionsLifecycleStateActive,
	"inactive": ListOperatorActionsLifecycleStateInactive,
}

// GetListOperatorActionsLifecycleStateEnumValues Enumerates the set of values for ListOperatorActionsLifecycleStateEnum
func GetListOperatorActionsLifecycleStateEnumValues() []ListOperatorActionsLifecycleStateEnum {
	values := make([]ListOperatorActionsLifecycleStateEnum, 0)
	for _, v := range mappingListOperatorActionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListOperatorActionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListOperatorActionsLifecycleStateEnum
func GetListOperatorActionsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingListOperatorActionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOperatorActionsLifecycleStateEnum(val string) (ListOperatorActionsLifecycleStateEnum, bool) {
	enum, ok := mappingListOperatorActionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOperatorActionsSortOrderEnum Enum with underlying type: string
type ListOperatorActionsSortOrderEnum string

// Set of constants representing the allowable values for ListOperatorActionsSortOrderEnum
const (
	ListOperatorActionsSortOrderAsc  ListOperatorActionsSortOrderEnum = "ASC"
	ListOperatorActionsSortOrderDesc ListOperatorActionsSortOrderEnum = "DESC"
)

var mappingListOperatorActionsSortOrderEnum = map[string]ListOperatorActionsSortOrderEnum{
	"ASC":  ListOperatorActionsSortOrderAsc,
	"DESC": ListOperatorActionsSortOrderDesc,
}

var mappingListOperatorActionsSortOrderEnumLowerCase = map[string]ListOperatorActionsSortOrderEnum{
	"asc":  ListOperatorActionsSortOrderAsc,
	"desc": ListOperatorActionsSortOrderDesc,
}

// GetListOperatorActionsSortOrderEnumValues Enumerates the set of values for ListOperatorActionsSortOrderEnum
func GetListOperatorActionsSortOrderEnumValues() []ListOperatorActionsSortOrderEnum {
	values := make([]ListOperatorActionsSortOrderEnum, 0)
	for _, v := range mappingListOperatorActionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOperatorActionsSortOrderEnumStringValues Enumerates the set of values in String for ListOperatorActionsSortOrderEnum
func GetListOperatorActionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOperatorActionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOperatorActionsSortOrderEnum(val string) (ListOperatorActionsSortOrderEnum, bool) {
	enum, ok := mappingListOperatorActionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOperatorActionsSortByEnum Enum with underlying type: string
type ListOperatorActionsSortByEnum string

// Set of constants representing the allowable values for ListOperatorActionsSortByEnum
const (
	ListOperatorActionsSortByTimecreated ListOperatorActionsSortByEnum = "timeCreated"
	ListOperatorActionsSortByDisplayname ListOperatorActionsSortByEnum = "displayName"
)

var mappingListOperatorActionsSortByEnum = map[string]ListOperatorActionsSortByEnum{
	"timeCreated": ListOperatorActionsSortByTimecreated,
	"displayName": ListOperatorActionsSortByDisplayname,
}

var mappingListOperatorActionsSortByEnumLowerCase = map[string]ListOperatorActionsSortByEnum{
	"timecreated": ListOperatorActionsSortByTimecreated,
	"displayname": ListOperatorActionsSortByDisplayname,
}

// GetListOperatorActionsSortByEnumValues Enumerates the set of values for ListOperatorActionsSortByEnum
func GetListOperatorActionsSortByEnumValues() []ListOperatorActionsSortByEnum {
	values := make([]ListOperatorActionsSortByEnum, 0)
	for _, v := range mappingListOperatorActionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOperatorActionsSortByEnumStringValues Enumerates the set of values in String for ListOperatorActionsSortByEnum
func GetListOperatorActionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOperatorActionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOperatorActionsSortByEnum(val string) (ListOperatorActionsSortByEnum, bool) {
	enum, ok := mappingListOperatorActionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
