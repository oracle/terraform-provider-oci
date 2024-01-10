// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSoftwareSourcesRequest wrapper for the ListSoftwareSources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListSoftwareSources.go.html to see an example of how to use ListSoftwareSourcesRequest.
type ListSoftwareSourcesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListSoftwareSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListSoftwareSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The current lifecycle state for the object.
	LifecycleState ListSoftwareSourcesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSoftwareSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSoftwareSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSoftwareSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSoftwareSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSoftwareSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSoftwareSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSoftwareSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSoftwareSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSoftwareSourcesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSoftwareSourcesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSoftwareSourcesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSoftwareSourcesResponse wrapper for the ListSoftwareSources operation
type ListSoftwareSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SoftwareSourceSummary instances
	Items []SoftwareSourceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSoftwareSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSoftwareSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSoftwareSourcesSortOrderEnum Enum with underlying type: string
type ListSoftwareSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListSoftwareSourcesSortOrderEnum
const (
	ListSoftwareSourcesSortOrderAsc  ListSoftwareSourcesSortOrderEnum = "ASC"
	ListSoftwareSourcesSortOrderDesc ListSoftwareSourcesSortOrderEnum = "DESC"
)

var mappingListSoftwareSourcesSortOrderEnum = map[string]ListSoftwareSourcesSortOrderEnum{
	"ASC":  ListSoftwareSourcesSortOrderAsc,
	"DESC": ListSoftwareSourcesSortOrderDesc,
}

var mappingListSoftwareSourcesSortOrderEnumLowerCase = map[string]ListSoftwareSourcesSortOrderEnum{
	"asc":  ListSoftwareSourcesSortOrderAsc,
	"desc": ListSoftwareSourcesSortOrderDesc,
}

// GetListSoftwareSourcesSortOrderEnumValues Enumerates the set of values for ListSoftwareSourcesSortOrderEnum
func GetListSoftwareSourcesSortOrderEnumValues() []ListSoftwareSourcesSortOrderEnum {
	values := make([]ListSoftwareSourcesSortOrderEnum, 0)
	for _, v := range mappingListSoftwareSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListSoftwareSourcesSortOrderEnum
func GetListSoftwareSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSoftwareSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareSourcesSortOrderEnum(val string) (ListSoftwareSourcesSortOrderEnum, bool) {
	enum, ok := mappingListSoftwareSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwareSourcesSortByEnum Enum with underlying type: string
type ListSoftwareSourcesSortByEnum string

// Set of constants representing the allowable values for ListSoftwareSourcesSortByEnum
const (
	ListSoftwareSourcesSortByTimecreated ListSoftwareSourcesSortByEnum = "TIMECREATED"
	ListSoftwareSourcesSortByDisplayname ListSoftwareSourcesSortByEnum = "DISPLAYNAME"
)

var mappingListSoftwareSourcesSortByEnum = map[string]ListSoftwareSourcesSortByEnum{
	"TIMECREATED": ListSoftwareSourcesSortByTimecreated,
	"DISPLAYNAME": ListSoftwareSourcesSortByDisplayname,
}

var mappingListSoftwareSourcesSortByEnumLowerCase = map[string]ListSoftwareSourcesSortByEnum{
	"timecreated": ListSoftwareSourcesSortByTimecreated,
	"displayname": ListSoftwareSourcesSortByDisplayname,
}

// GetListSoftwareSourcesSortByEnumValues Enumerates the set of values for ListSoftwareSourcesSortByEnum
func GetListSoftwareSourcesSortByEnumValues() []ListSoftwareSourcesSortByEnum {
	values := make([]ListSoftwareSourcesSortByEnum, 0)
	for _, v := range mappingListSoftwareSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareSourcesSortByEnumStringValues Enumerates the set of values in String for ListSoftwareSourcesSortByEnum
func GetListSoftwareSourcesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListSoftwareSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareSourcesSortByEnum(val string) (ListSoftwareSourcesSortByEnum, bool) {
	enum, ok := mappingListSoftwareSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwareSourcesLifecycleStateEnum Enum with underlying type: string
type ListSoftwareSourcesLifecycleStateEnum string

// Set of constants representing the allowable values for ListSoftwareSourcesLifecycleStateEnum
const (
	ListSoftwareSourcesLifecycleStateCreating ListSoftwareSourcesLifecycleStateEnum = "CREATING"
	ListSoftwareSourcesLifecycleStateUpdating ListSoftwareSourcesLifecycleStateEnum = "UPDATING"
	ListSoftwareSourcesLifecycleStateActive   ListSoftwareSourcesLifecycleStateEnum = "ACTIVE"
	ListSoftwareSourcesLifecycleStateDeleting ListSoftwareSourcesLifecycleStateEnum = "DELETING"
	ListSoftwareSourcesLifecycleStateDeleted  ListSoftwareSourcesLifecycleStateEnum = "DELETED"
	ListSoftwareSourcesLifecycleStateFailed   ListSoftwareSourcesLifecycleStateEnum = "FAILED"
)

var mappingListSoftwareSourcesLifecycleStateEnum = map[string]ListSoftwareSourcesLifecycleStateEnum{
	"CREATING": ListSoftwareSourcesLifecycleStateCreating,
	"UPDATING": ListSoftwareSourcesLifecycleStateUpdating,
	"ACTIVE":   ListSoftwareSourcesLifecycleStateActive,
	"DELETING": ListSoftwareSourcesLifecycleStateDeleting,
	"DELETED":  ListSoftwareSourcesLifecycleStateDeleted,
	"FAILED":   ListSoftwareSourcesLifecycleStateFailed,
}

var mappingListSoftwareSourcesLifecycleStateEnumLowerCase = map[string]ListSoftwareSourcesLifecycleStateEnum{
	"creating": ListSoftwareSourcesLifecycleStateCreating,
	"updating": ListSoftwareSourcesLifecycleStateUpdating,
	"active":   ListSoftwareSourcesLifecycleStateActive,
	"deleting": ListSoftwareSourcesLifecycleStateDeleting,
	"deleted":  ListSoftwareSourcesLifecycleStateDeleted,
	"failed":   ListSoftwareSourcesLifecycleStateFailed,
}

// GetListSoftwareSourcesLifecycleStateEnumValues Enumerates the set of values for ListSoftwareSourcesLifecycleStateEnum
func GetListSoftwareSourcesLifecycleStateEnumValues() []ListSoftwareSourcesLifecycleStateEnum {
	values := make([]ListSoftwareSourcesLifecycleStateEnum, 0)
	for _, v := range mappingListSoftwareSourcesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareSourcesLifecycleStateEnumStringValues Enumerates the set of values in String for ListSoftwareSourcesLifecycleStateEnum
func GetListSoftwareSourcesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListSoftwareSourcesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareSourcesLifecycleStateEnum(val string) (ListSoftwareSourcesLifecycleStateEnum, bool) {
	enum, ok := mappingListSoftwareSourcesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
