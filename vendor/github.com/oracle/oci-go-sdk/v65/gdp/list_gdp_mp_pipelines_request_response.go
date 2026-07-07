// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package gdp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListGdpMpPipelinesRequest wrapper for the ListGdpMpPipelines operation
type ListGdpMpPipelinesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique pipeline identifier.
	GdpMpPipelineId *string `mandatory:"false" contributesTo:"query" name:"gdpMpPipelineId"`

	// A filter to return only resources with a lifecycleState that matches the given lifecycleState.
	LifecycleState ListGdpMpPipelinesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListGdpMpPipelinesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListGdpMpPipelinesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGdpMpPipelinesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGdpMpPipelinesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGdpMpPipelinesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGdpMpPipelinesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGdpMpPipelinesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGdpMpPipelinesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListGdpMpPipelinesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGdpMpPipelinesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGdpMpPipelinesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGdpMpPipelinesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGdpMpPipelinesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGdpMpPipelinesResponse wrapper for the ListGdpMpPipelines operation
type ListGdpMpPipelinesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GdpMpPipelineCollection instances
	GdpMpPipelineCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGdpMpPipelinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGdpMpPipelinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGdpMpPipelinesLifecycleStateEnum Enum with underlying type: string
type ListGdpMpPipelinesLifecycleStateEnum string

// Set of constants representing the allowable values for ListGdpMpPipelinesLifecycleStateEnum
const (
	ListGdpMpPipelinesLifecycleStateCreating       ListGdpMpPipelinesLifecycleStateEnum = "CREATING"
	ListGdpMpPipelinesLifecycleStateUpdating       ListGdpMpPipelinesLifecycleStateEnum = "UPDATING"
	ListGdpMpPipelinesLifecycleStateActive         ListGdpMpPipelinesLifecycleStateEnum = "ACTIVE"
	ListGdpMpPipelinesLifecycleStateInactive       ListGdpMpPipelinesLifecycleStateEnum = "INACTIVE"
	ListGdpMpPipelinesLifecycleStateDeleting       ListGdpMpPipelinesLifecycleStateEnum = "DELETING"
	ListGdpMpPipelinesLifecycleStateDeleted        ListGdpMpPipelinesLifecycleStateEnum = "DELETED"
	ListGdpMpPipelinesLifecycleStateFailed         ListGdpMpPipelinesLifecycleStateEnum = "FAILED"
	ListGdpMpPipelinesLifecycleStateNeedsAttention ListGdpMpPipelinesLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListGdpMpPipelinesLifecycleStateEnum = map[string]ListGdpMpPipelinesLifecycleStateEnum{
	"CREATING":        ListGdpMpPipelinesLifecycleStateCreating,
	"UPDATING":        ListGdpMpPipelinesLifecycleStateUpdating,
	"ACTIVE":          ListGdpMpPipelinesLifecycleStateActive,
	"INACTIVE":        ListGdpMpPipelinesLifecycleStateInactive,
	"DELETING":        ListGdpMpPipelinesLifecycleStateDeleting,
	"DELETED":         ListGdpMpPipelinesLifecycleStateDeleted,
	"FAILED":          ListGdpMpPipelinesLifecycleStateFailed,
	"NEEDS_ATTENTION": ListGdpMpPipelinesLifecycleStateNeedsAttention,
}

var mappingListGdpMpPipelinesLifecycleStateEnumLowerCase = map[string]ListGdpMpPipelinesLifecycleStateEnum{
	"creating":        ListGdpMpPipelinesLifecycleStateCreating,
	"updating":        ListGdpMpPipelinesLifecycleStateUpdating,
	"active":          ListGdpMpPipelinesLifecycleStateActive,
	"inactive":        ListGdpMpPipelinesLifecycleStateInactive,
	"deleting":        ListGdpMpPipelinesLifecycleStateDeleting,
	"deleted":         ListGdpMpPipelinesLifecycleStateDeleted,
	"failed":          ListGdpMpPipelinesLifecycleStateFailed,
	"needs_attention": ListGdpMpPipelinesLifecycleStateNeedsAttention,
}

// GetListGdpMpPipelinesLifecycleStateEnumValues Enumerates the set of values for ListGdpMpPipelinesLifecycleStateEnum
func GetListGdpMpPipelinesLifecycleStateEnumValues() []ListGdpMpPipelinesLifecycleStateEnum {
	values := make([]ListGdpMpPipelinesLifecycleStateEnum, 0)
	for _, v := range mappingListGdpMpPipelinesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListGdpMpPipelinesLifecycleStateEnumStringValues Enumerates the set of values in String for ListGdpMpPipelinesLifecycleStateEnum
func GetListGdpMpPipelinesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListGdpMpPipelinesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGdpMpPipelinesLifecycleStateEnum(val string) (ListGdpMpPipelinesLifecycleStateEnum, bool) {
	enum, ok := mappingListGdpMpPipelinesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGdpMpPipelinesSortOrderEnum Enum with underlying type: string
type ListGdpMpPipelinesSortOrderEnum string

// Set of constants representing the allowable values for ListGdpMpPipelinesSortOrderEnum
const (
	ListGdpMpPipelinesSortOrderAsc  ListGdpMpPipelinesSortOrderEnum = "ASC"
	ListGdpMpPipelinesSortOrderDesc ListGdpMpPipelinesSortOrderEnum = "DESC"
)

var mappingListGdpMpPipelinesSortOrderEnum = map[string]ListGdpMpPipelinesSortOrderEnum{
	"ASC":  ListGdpMpPipelinesSortOrderAsc,
	"DESC": ListGdpMpPipelinesSortOrderDesc,
}

var mappingListGdpMpPipelinesSortOrderEnumLowerCase = map[string]ListGdpMpPipelinesSortOrderEnum{
	"asc":  ListGdpMpPipelinesSortOrderAsc,
	"desc": ListGdpMpPipelinesSortOrderDesc,
}

// GetListGdpMpPipelinesSortOrderEnumValues Enumerates the set of values for ListGdpMpPipelinesSortOrderEnum
func GetListGdpMpPipelinesSortOrderEnumValues() []ListGdpMpPipelinesSortOrderEnum {
	values := make([]ListGdpMpPipelinesSortOrderEnum, 0)
	for _, v := range mappingListGdpMpPipelinesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGdpMpPipelinesSortOrderEnumStringValues Enumerates the set of values in String for ListGdpMpPipelinesSortOrderEnum
func GetListGdpMpPipelinesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGdpMpPipelinesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGdpMpPipelinesSortOrderEnum(val string) (ListGdpMpPipelinesSortOrderEnum, bool) {
	enum, ok := mappingListGdpMpPipelinesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGdpMpPipelinesSortByEnum Enum with underlying type: string
type ListGdpMpPipelinesSortByEnum string

// Set of constants representing the allowable values for ListGdpMpPipelinesSortByEnum
const (
	ListGdpMpPipelinesSortByTimecreated ListGdpMpPipelinesSortByEnum = "timeCreated"
	ListGdpMpPipelinesSortByDisplayname ListGdpMpPipelinesSortByEnum = "displayName"
)

var mappingListGdpMpPipelinesSortByEnum = map[string]ListGdpMpPipelinesSortByEnum{
	"timeCreated": ListGdpMpPipelinesSortByTimecreated,
	"displayName": ListGdpMpPipelinesSortByDisplayname,
}

var mappingListGdpMpPipelinesSortByEnumLowerCase = map[string]ListGdpMpPipelinesSortByEnum{
	"timecreated": ListGdpMpPipelinesSortByTimecreated,
	"displayname": ListGdpMpPipelinesSortByDisplayname,
}

// GetListGdpMpPipelinesSortByEnumValues Enumerates the set of values for ListGdpMpPipelinesSortByEnum
func GetListGdpMpPipelinesSortByEnumValues() []ListGdpMpPipelinesSortByEnum {
	values := make([]ListGdpMpPipelinesSortByEnum, 0)
	for _, v := range mappingListGdpMpPipelinesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGdpMpPipelinesSortByEnumStringValues Enumerates the set of values in String for ListGdpMpPipelinesSortByEnum
func GetListGdpMpPipelinesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListGdpMpPipelinesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGdpMpPipelinesSortByEnum(val string) (ListGdpMpPipelinesSortByEnum, bool) {
	enum, ok := mappingListGdpMpPipelinesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
