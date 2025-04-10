// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package sch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListConnectorPluginsRequest wrapper for the ListConnectorPlugins operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/sch/ListConnectorPlugins.go.html to see an example of how to use ListConnectorPluginsRequest.
type ListConnectorPluginsRequest struct {

	// A filter to return only resources that match the given lifecycle state.
	// Example: `ACTIVE`
	LifecycleState ListConnectorPluginsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	// Example: `example_service_connector`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given connector plugin name ignoring case.
	// Example: `QueueSource`
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The maximum number of results per page, or items to return
	// in a paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListConnectorPluginsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for `timeCreated` is descending.
	// Default order for `displayName` is ascending. If no value is specified `timeCreated` is default.
	SortBy ListConnectorPluginsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConnectorPluginsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConnectorPluginsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConnectorPluginsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConnectorPluginsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConnectorPluginsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListConnectorPluginsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListConnectorPluginsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectorPluginsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConnectorPluginsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectorPluginsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConnectorPluginsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConnectorPluginsResponse wrapper for the ListConnectorPlugins operation
type ListConnectorPluginsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConnectorPluginCollection instances
	ConnectorPluginCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain. For important details about
	// how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination.  When this header appears in the response,
	// previous pages of results exist. For important details about
	// how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListConnectorPluginsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConnectorPluginsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConnectorPluginsLifecycleStateEnum Enum with underlying type: string
type ListConnectorPluginsLifecycleStateEnum string

// Set of constants representing the allowable values for ListConnectorPluginsLifecycleStateEnum
const (
	ListConnectorPluginsLifecycleStateCreating       ListConnectorPluginsLifecycleStateEnum = "CREATING"
	ListConnectorPluginsLifecycleStateUpdating       ListConnectorPluginsLifecycleStateEnum = "UPDATING"
	ListConnectorPluginsLifecycleStateActive         ListConnectorPluginsLifecycleStateEnum = "ACTIVE"
	ListConnectorPluginsLifecycleStateInactive       ListConnectorPluginsLifecycleStateEnum = "INACTIVE"
	ListConnectorPluginsLifecycleStateNeedsAttention ListConnectorPluginsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListConnectorPluginsLifecycleStateDeleting       ListConnectorPluginsLifecycleStateEnum = "DELETING"
	ListConnectorPluginsLifecycleStateDeleted        ListConnectorPluginsLifecycleStateEnum = "DELETED"
	ListConnectorPluginsLifecycleStateFailed         ListConnectorPluginsLifecycleStateEnum = "FAILED"
)

var mappingListConnectorPluginsLifecycleStateEnum = map[string]ListConnectorPluginsLifecycleStateEnum{
	"CREATING":        ListConnectorPluginsLifecycleStateCreating,
	"UPDATING":        ListConnectorPluginsLifecycleStateUpdating,
	"ACTIVE":          ListConnectorPluginsLifecycleStateActive,
	"INACTIVE":        ListConnectorPluginsLifecycleStateInactive,
	"NEEDS_ATTENTION": ListConnectorPluginsLifecycleStateNeedsAttention,
	"DELETING":        ListConnectorPluginsLifecycleStateDeleting,
	"DELETED":         ListConnectorPluginsLifecycleStateDeleted,
	"FAILED":          ListConnectorPluginsLifecycleStateFailed,
}

var mappingListConnectorPluginsLifecycleStateEnumLowerCase = map[string]ListConnectorPluginsLifecycleStateEnum{
	"creating":        ListConnectorPluginsLifecycleStateCreating,
	"updating":        ListConnectorPluginsLifecycleStateUpdating,
	"active":          ListConnectorPluginsLifecycleStateActive,
	"inactive":        ListConnectorPluginsLifecycleStateInactive,
	"needs_attention": ListConnectorPluginsLifecycleStateNeedsAttention,
	"deleting":        ListConnectorPluginsLifecycleStateDeleting,
	"deleted":         ListConnectorPluginsLifecycleStateDeleted,
	"failed":          ListConnectorPluginsLifecycleStateFailed,
}

// GetListConnectorPluginsLifecycleStateEnumValues Enumerates the set of values for ListConnectorPluginsLifecycleStateEnum
func GetListConnectorPluginsLifecycleStateEnumValues() []ListConnectorPluginsLifecycleStateEnum {
	values := make([]ListConnectorPluginsLifecycleStateEnum, 0)
	for _, v := range mappingListConnectorPluginsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectorPluginsLifecycleStateEnumStringValues Enumerates the set of values in String for ListConnectorPluginsLifecycleStateEnum
func GetListConnectorPluginsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListConnectorPluginsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectorPluginsLifecycleStateEnum(val string) (ListConnectorPluginsLifecycleStateEnum, bool) {
	enum, ok := mappingListConnectorPluginsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectorPluginsSortOrderEnum Enum with underlying type: string
type ListConnectorPluginsSortOrderEnum string

// Set of constants representing the allowable values for ListConnectorPluginsSortOrderEnum
const (
	ListConnectorPluginsSortOrderAsc  ListConnectorPluginsSortOrderEnum = "ASC"
	ListConnectorPluginsSortOrderDesc ListConnectorPluginsSortOrderEnum = "DESC"
)

var mappingListConnectorPluginsSortOrderEnum = map[string]ListConnectorPluginsSortOrderEnum{
	"ASC":  ListConnectorPluginsSortOrderAsc,
	"DESC": ListConnectorPluginsSortOrderDesc,
}

var mappingListConnectorPluginsSortOrderEnumLowerCase = map[string]ListConnectorPluginsSortOrderEnum{
	"asc":  ListConnectorPluginsSortOrderAsc,
	"desc": ListConnectorPluginsSortOrderDesc,
}

// GetListConnectorPluginsSortOrderEnumValues Enumerates the set of values for ListConnectorPluginsSortOrderEnum
func GetListConnectorPluginsSortOrderEnumValues() []ListConnectorPluginsSortOrderEnum {
	values := make([]ListConnectorPluginsSortOrderEnum, 0)
	for _, v := range mappingListConnectorPluginsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectorPluginsSortOrderEnumStringValues Enumerates the set of values in String for ListConnectorPluginsSortOrderEnum
func GetListConnectorPluginsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConnectorPluginsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectorPluginsSortOrderEnum(val string) (ListConnectorPluginsSortOrderEnum, bool) {
	enum, ok := mappingListConnectorPluginsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectorPluginsSortByEnum Enum with underlying type: string
type ListConnectorPluginsSortByEnum string

// Set of constants representing the allowable values for ListConnectorPluginsSortByEnum
const (
	ListConnectorPluginsSortByTimecreated ListConnectorPluginsSortByEnum = "timeCreated"
	ListConnectorPluginsSortByDisplayname ListConnectorPluginsSortByEnum = "displayName"
)

var mappingListConnectorPluginsSortByEnum = map[string]ListConnectorPluginsSortByEnum{
	"timeCreated": ListConnectorPluginsSortByTimecreated,
	"displayName": ListConnectorPluginsSortByDisplayname,
}

var mappingListConnectorPluginsSortByEnumLowerCase = map[string]ListConnectorPluginsSortByEnum{
	"timecreated": ListConnectorPluginsSortByTimecreated,
	"displayname": ListConnectorPluginsSortByDisplayname,
}

// GetListConnectorPluginsSortByEnumValues Enumerates the set of values for ListConnectorPluginsSortByEnum
func GetListConnectorPluginsSortByEnumValues() []ListConnectorPluginsSortByEnum {
	values := make([]ListConnectorPluginsSortByEnum, 0)
	for _, v := range mappingListConnectorPluginsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectorPluginsSortByEnumStringValues Enumerates the set of values in String for ListConnectorPluginsSortByEnum
func GetListConnectorPluginsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListConnectorPluginsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectorPluginsSortByEnum(val string) (ListConnectorPluginsSortByEnum, bool) {
	enum, ok := mappingListConnectorPluginsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
