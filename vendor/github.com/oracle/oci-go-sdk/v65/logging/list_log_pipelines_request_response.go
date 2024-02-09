// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLogPipelinesRequest wrapper for the ListLogPipelines operation
type ListLogPipelinesRequest struct {

	// Compartment OCID to list resources in. See compartmentIdInSubtree
	//      for nested compartments traversal.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Resource name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Lifecycle state of the log pipeline.
	LifecycleState ListLogPipelinesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by (one column only). Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListLogPipelinesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListLogPipelinesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogPipelinesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogPipelinesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogPipelinesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogPipelinesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLogPipelinesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLogPipelinesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListLogPipelinesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogPipelinesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLogPipelinesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogPipelinesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLogPipelinesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLogPipelinesResponse wrapper for the ListLogPipelines operation
type ListLogPipelinesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogPipelineSummaryCollection instances
	LogPipelineSummaryCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages
	// of results exist. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLogPipelinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogPipelinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogPipelinesLifecycleStateEnum Enum with underlying type: string
type ListLogPipelinesLifecycleStateEnum string

// Set of constants representing the allowable values for ListLogPipelinesLifecycleStateEnum
const (
	ListLogPipelinesLifecycleStateCreating ListLogPipelinesLifecycleStateEnum = "CREATING"
	ListLogPipelinesLifecycleStateActive   ListLogPipelinesLifecycleStateEnum = "ACTIVE"
	ListLogPipelinesLifecycleStateUpdating ListLogPipelinesLifecycleStateEnum = "UPDATING"
	ListLogPipelinesLifecycleStateInactive ListLogPipelinesLifecycleStateEnum = "INACTIVE"
	ListLogPipelinesLifecycleStateDeleting ListLogPipelinesLifecycleStateEnum = "DELETING"
	ListLogPipelinesLifecycleStateFailed   ListLogPipelinesLifecycleStateEnum = "FAILED"
)

var mappingListLogPipelinesLifecycleStateEnum = map[string]ListLogPipelinesLifecycleStateEnum{
	"CREATING": ListLogPipelinesLifecycleStateCreating,
	"ACTIVE":   ListLogPipelinesLifecycleStateActive,
	"UPDATING": ListLogPipelinesLifecycleStateUpdating,
	"INACTIVE": ListLogPipelinesLifecycleStateInactive,
	"DELETING": ListLogPipelinesLifecycleStateDeleting,
	"FAILED":   ListLogPipelinesLifecycleStateFailed,
}

var mappingListLogPipelinesLifecycleStateEnumLowerCase = map[string]ListLogPipelinesLifecycleStateEnum{
	"creating": ListLogPipelinesLifecycleStateCreating,
	"active":   ListLogPipelinesLifecycleStateActive,
	"updating": ListLogPipelinesLifecycleStateUpdating,
	"inactive": ListLogPipelinesLifecycleStateInactive,
	"deleting": ListLogPipelinesLifecycleStateDeleting,
	"failed":   ListLogPipelinesLifecycleStateFailed,
}

// GetListLogPipelinesLifecycleStateEnumValues Enumerates the set of values for ListLogPipelinesLifecycleStateEnum
func GetListLogPipelinesLifecycleStateEnumValues() []ListLogPipelinesLifecycleStateEnum {
	values := make([]ListLogPipelinesLifecycleStateEnum, 0)
	for _, v := range mappingListLogPipelinesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogPipelinesLifecycleStateEnumStringValues Enumerates the set of values in String for ListLogPipelinesLifecycleStateEnum
func GetListLogPipelinesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
		"DELETING",
		"FAILED",
	}
}

// GetMappingListLogPipelinesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogPipelinesLifecycleStateEnum(val string) (ListLogPipelinesLifecycleStateEnum, bool) {
	enum, ok := mappingListLogPipelinesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogPipelinesSortByEnum Enum with underlying type: string
type ListLogPipelinesSortByEnum string

// Set of constants representing the allowable values for ListLogPipelinesSortByEnum
const (
	ListLogPipelinesSortByTimecreated ListLogPipelinesSortByEnum = "timeCreated"
	ListLogPipelinesSortByDisplayname ListLogPipelinesSortByEnum = "displayName"
)

var mappingListLogPipelinesSortByEnum = map[string]ListLogPipelinesSortByEnum{
	"timeCreated": ListLogPipelinesSortByTimecreated,
	"displayName": ListLogPipelinesSortByDisplayname,
}

var mappingListLogPipelinesSortByEnumLowerCase = map[string]ListLogPipelinesSortByEnum{
	"timecreated": ListLogPipelinesSortByTimecreated,
	"displayname": ListLogPipelinesSortByDisplayname,
}

// GetListLogPipelinesSortByEnumValues Enumerates the set of values for ListLogPipelinesSortByEnum
func GetListLogPipelinesSortByEnumValues() []ListLogPipelinesSortByEnum {
	values := make([]ListLogPipelinesSortByEnum, 0)
	for _, v := range mappingListLogPipelinesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogPipelinesSortByEnumStringValues Enumerates the set of values in String for ListLogPipelinesSortByEnum
func GetListLogPipelinesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListLogPipelinesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogPipelinesSortByEnum(val string) (ListLogPipelinesSortByEnum, bool) {
	enum, ok := mappingListLogPipelinesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogPipelinesSortOrderEnum Enum with underlying type: string
type ListLogPipelinesSortOrderEnum string

// Set of constants representing the allowable values for ListLogPipelinesSortOrderEnum
const (
	ListLogPipelinesSortOrderAsc  ListLogPipelinesSortOrderEnum = "ASC"
	ListLogPipelinesSortOrderDesc ListLogPipelinesSortOrderEnum = "DESC"
)

var mappingListLogPipelinesSortOrderEnum = map[string]ListLogPipelinesSortOrderEnum{
	"ASC":  ListLogPipelinesSortOrderAsc,
	"DESC": ListLogPipelinesSortOrderDesc,
}

var mappingListLogPipelinesSortOrderEnumLowerCase = map[string]ListLogPipelinesSortOrderEnum{
	"asc":  ListLogPipelinesSortOrderAsc,
	"desc": ListLogPipelinesSortOrderDesc,
}

// GetListLogPipelinesSortOrderEnumValues Enumerates the set of values for ListLogPipelinesSortOrderEnum
func GetListLogPipelinesSortOrderEnumValues() []ListLogPipelinesSortOrderEnum {
	values := make([]ListLogPipelinesSortOrderEnum, 0)
	for _, v := range mappingListLogPipelinesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogPipelinesSortOrderEnumStringValues Enumerates the set of values in String for ListLogPipelinesSortOrderEnum
func GetListLogPipelinesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLogPipelinesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogPipelinesSortOrderEnum(val string) (ListLogPipelinesSortOrderEnum, bool) {
	enum, ok := mappingListLogPipelinesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
