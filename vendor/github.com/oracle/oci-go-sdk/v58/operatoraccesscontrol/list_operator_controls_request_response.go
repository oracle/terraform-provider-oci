// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package operatoraccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListOperatorControlsRequest wrapper for the ListOperatorControls operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ListOperatorControls.go.html to see an example of how to use ListOperatorControlsRequest.
type ListOperatorControlsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycleState matches the given OperatorControl lifecycleState.
	LifecycleState ListOperatorControlsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return OperatorControl that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only lists of resources that match the entire given service type.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOperatorControlsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListOperatorControlsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOperatorControlsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOperatorControlsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOperatorControlsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOperatorControlsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOperatorControlsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOperatorControlsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListOperatorControlsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOperatorControlsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOperatorControlsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOperatorControlsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOperatorControlsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOperatorControlsResponse wrapper for the ListOperatorControls operation
type ListOperatorControlsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OperatorControlCollection instances
	OperatorControlCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOperatorControlsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOperatorControlsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOperatorControlsLifecycleStateEnum Enum with underlying type: string
type ListOperatorControlsLifecycleStateEnum string

// Set of constants representing the allowable values for ListOperatorControlsLifecycleStateEnum
const (
	ListOperatorControlsLifecycleStateCreated    ListOperatorControlsLifecycleStateEnum = "CREATED"
	ListOperatorControlsLifecycleStateAssigned   ListOperatorControlsLifecycleStateEnum = "ASSIGNED"
	ListOperatorControlsLifecycleStateUnassigned ListOperatorControlsLifecycleStateEnum = "UNASSIGNED"
	ListOperatorControlsLifecycleStateDeleted    ListOperatorControlsLifecycleStateEnum = "DELETED"
)

var mappingListOperatorControlsLifecycleStateEnum = map[string]ListOperatorControlsLifecycleStateEnum{
	"CREATED":    ListOperatorControlsLifecycleStateCreated,
	"ASSIGNED":   ListOperatorControlsLifecycleStateAssigned,
	"UNASSIGNED": ListOperatorControlsLifecycleStateUnassigned,
	"DELETED":    ListOperatorControlsLifecycleStateDeleted,
}

// GetListOperatorControlsLifecycleStateEnumValues Enumerates the set of values for ListOperatorControlsLifecycleStateEnum
func GetListOperatorControlsLifecycleStateEnumValues() []ListOperatorControlsLifecycleStateEnum {
	values := make([]ListOperatorControlsLifecycleStateEnum, 0)
	for _, v := range mappingListOperatorControlsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListOperatorControlsLifecycleStateEnumStringValues Enumerates the set of values in String for ListOperatorControlsLifecycleStateEnum
func GetListOperatorControlsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATED",
		"ASSIGNED",
		"UNASSIGNED",
		"DELETED",
	}
}

// GetMappingListOperatorControlsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOperatorControlsLifecycleStateEnum(val string) (ListOperatorControlsLifecycleStateEnum, bool) {
	mappingListOperatorControlsLifecycleStateEnumIgnoreCase := make(map[string]ListOperatorControlsLifecycleStateEnum)
	for k, v := range mappingListOperatorControlsLifecycleStateEnum {
		mappingListOperatorControlsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListOperatorControlsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListOperatorControlsSortOrderEnum Enum with underlying type: string
type ListOperatorControlsSortOrderEnum string

// Set of constants representing the allowable values for ListOperatorControlsSortOrderEnum
const (
	ListOperatorControlsSortOrderAsc  ListOperatorControlsSortOrderEnum = "ASC"
	ListOperatorControlsSortOrderDesc ListOperatorControlsSortOrderEnum = "DESC"
)

var mappingListOperatorControlsSortOrderEnum = map[string]ListOperatorControlsSortOrderEnum{
	"ASC":  ListOperatorControlsSortOrderAsc,
	"DESC": ListOperatorControlsSortOrderDesc,
}

// GetListOperatorControlsSortOrderEnumValues Enumerates the set of values for ListOperatorControlsSortOrderEnum
func GetListOperatorControlsSortOrderEnumValues() []ListOperatorControlsSortOrderEnum {
	values := make([]ListOperatorControlsSortOrderEnum, 0)
	for _, v := range mappingListOperatorControlsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOperatorControlsSortOrderEnumStringValues Enumerates the set of values in String for ListOperatorControlsSortOrderEnum
func GetListOperatorControlsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOperatorControlsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOperatorControlsSortOrderEnum(val string) (ListOperatorControlsSortOrderEnum, bool) {
	mappingListOperatorControlsSortOrderEnumIgnoreCase := make(map[string]ListOperatorControlsSortOrderEnum)
	for k, v := range mappingListOperatorControlsSortOrderEnum {
		mappingListOperatorControlsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListOperatorControlsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListOperatorControlsSortByEnum Enum with underlying type: string
type ListOperatorControlsSortByEnum string

// Set of constants representing the allowable values for ListOperatorControlsSortByEnum
const (
	ListOperatorControlsSortByTimecreated ListOperatorControlsSortByEnum = "timeCreated"
	ListOperatorControlsSortByDisplayname ListOperatorControlsSortByEnum = "displayName"
)

var mappingListOperatorControlsSortByEnum = map[string]ListOperatorControlsSortByEnum{
	"timeCreated": ListOperatorControlsSortByTimecreated,
	"displayName": ListOperatorControlsSortByDisplayname,
}

// GetListOperatorControlsSortByEnumValues Enumerates the set of values for ListOperatorControlsSortByEnum
func GetListOperatorControlsSortByEnumValues() []ListOperatorControlsSortByEnum {
	values := make([]ListOperatorControlsSortByEnum, 0)
	for _, v := range mappingListOperatorControlsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOperatorControlsSortByEnumStringValues Enumerates the set of values in String for ListOperatorControlsSortByEnum
func GetListOperatorControlsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOperatorControlsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOperatorControlsSortByEnum(val string) (ListOperatorControlsSortByEnum, bool) {
	mappingListOperatorControlsSortByEnumIgnoreCase := make(map[string]ListOperatorControlsSortByEnum)
	for k, v := range mappingListOperatorControlsSortByEnum {
		mappingListOperatorControlsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListOperatorControlsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
