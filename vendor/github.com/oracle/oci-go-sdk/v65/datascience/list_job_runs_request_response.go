// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJobRunsRequest wrapper for the ListJobRuns operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListJobRuns.go.html to see an example of how to use ListJobRunsRequest.
type ListJobRunsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	JobId *string `mandatory:"false" contributesTo:"query" name:"jobId"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListJobRunsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown
	// in descending order. When you sort by `displayName`, the results are
	// shown in ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListJobRunsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListJobRunsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJobRunsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobRunsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJobRunsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobRunsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJobRunsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJobRunsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJobRunsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobRunsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJobRunsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobRunsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListJobRunsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJobRunsResponse wrapper for the ListJobRuns operation
type ListJobRunsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []JobRunSummary instances
	Items []JobRunSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListJobRunsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobRunsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobRunsSortOrderEnum Enum with underlying type: string
type ListJobRunsSortOrderEnum string

// Set of constants representing the allowable values for ListJobRunsSortOrderEnum
const (
	ListJobRunsSortOrderAsc  ListJobRunsSortOrderEnum = "ASC"
	ListJobRunsSortOrderDesc ListJobRunsSortOrderEnum = "DESC"
)

var mappingListJobRunsSortOrderEnum = map[string]ListJobRunsSortOrderEnum{
	"ASC":  ListJobRunsSortOrderAsc,
	"DESC": ListJobRunsSortOrderDesc,
}

var mappingListJobRunsSortOrderEnumLowerCase = map[string]ListJobRunsSortOrderEnum{
	"asc":  ListJobRunsSortOrderAsc,
	"desc": ListJobRunsSortOrderDesc,
}

// GetListJobRunsSortOrderEnumValues Enumerates the set of values for ListJobRunsSortOrderEnum
func GetListJobRunsSortOrderEnumValues() []ListJobRunsSortOrderEnum {
	values := make([]ListJobRunsSortOrderEnum, 0)
	for _, v := range mappingListJobRunsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobRunsSortOrderEnumStringValues Enumerates the set of values in String for ListJobRunsSortOrderEnum
func GetListJobRunsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJobRunsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobRunsSortOrderEnum(val string) (ListJobRunsSortOrderEnum, bool) {
	enum, ok := mappingListJobRunsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobRunsSortByEnum Enum with underlying type: string
type ListJobRunsSortByEnum string

// Set of constants representing the allowable values for ListJobRunsSortByEnum
const (
	ListJobRunsSortByTimecreated ListJobRunsSortByEnum = "timeCreated"
	ListJobRunsSortByDisplayname ListJobRunsSortByEnum = "displayName"
)

var mappingListJobRunsSortByEnum = map[string]ListJobRunsSortByEnum{
	"timeCreated": ListJobRunsSortByTimecreated,
	"displayName": ListJobRunsSortByDisplayname,
}

var mappingListJobRunsSortByEnumLowerCase = map[string]ListJobRunsSortByEnum{
	"timecreated": ListJobRunsSortByTimecreated,
	"displayname": ListJobRunsSortByDisplayname,
}

// GetListJobRunsSortByEnumValues Enumerates the set of values for ListJobRunsSortByEnum
func GetListJobRunsSortByEnumValues() []ListJobRunsSortByEnum {
	values := make([]ListJobRunsSortByEnum, 0)
	for _, v := range mappingListJobRunsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobRunsSortByEnumStringValues Enumerates the set of values in String for ListJobRunsSortByEnum
func GetListJobRunsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListJobRunsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobRunsSortByEnum(val string) (ListJobRunsSortByEnum, bool) {
	enum, ok := mappingListJobRunsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobRunsLifecycleStateEnum Enum with underlying type: string
type ListJobRunsLifecycleStateEnum string

// Set of constants representing the allowable values for ListJobRunsLifecycleStateEnum
const (
	ListJobRunsLifecycleStateAccepted       ListJobRunsLifecycleStateEnum = "ACCEPTED"
	ListJobRunsLifecycleStateInProgress     ListJobRunsLifecycleStateEnum = "IN_PROGRESS"
	ListJobRunsLifecycleStateFailed         ListJobRunsLifecycleStateEnum = "FAILED"
	ListJobRunsLifecycleStateSucceeded      ListJobRunsLifecycleStateEnum = "SUCCEEDED"
	ListJobRunsLifecycleStateCanceling      ListJobRunsLifecycleStateEnum = "CANCELING"
	ListJobRunsLifecycleStateCanceled       ListJobRunsLifecycleStateEnum = "CANCELED"
	ListJobRunsLifecycleStateDeleted        ListJobRunsLifecycleStateEnum = "DELETED"
	ListJobRunsLifecycleStateNeedsAttention ListJobRunsLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListJobRunsLifecycleStateEnum = map[string]ListJobRunsLifecycleStateEnum{
	"ACCEPTED":        ListJobRunsLifecycleStateAccepted,
	"IN_PROGRESS":     ListJobRunsLifecycleStateInProgress,
	"FAILED":          ListJobRunsLifecycleStateFailed,
	"SUCCEEDED":       ListJobRunsLifecycleStateSucceeded,
	"CANCELING":       ListJobRunsLifecycleStateCanceling,
	"CANCELED":        ListJobRunsLifecycleStateCanceled,
	"DELETED":         ListJobRunsLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListJobRunsLifecycleStateNeedsAttention,
}

var mappingListJobRunsLifecycleStateEnumLowerCase = map[string]ListJobRunsLifecycleStateEnum{
	"accepted":        ListJobRunsLifecycleStateAccepted,
	"in_progress":     ListJobRunsLifecycleStateInProgress,
	"failed":          ListJobRunsLifecycleStateFailed,
	"succeeded":       ListJobRunsLifecycleStateSucceeded,
	"canceling":       ListJobRunsLifecycleStateCanceling,
	"canceled":        ListJobRunsLifecycleStateCanceled,
	"deleted":         ListJobRunsLifecycleStateDeleted,
	"needs_attention": ListJobRunsLifecycleStateNeedsAttention,
}

// GetListJobRunsLifecycleStateEnumValues Enumerates the set of values for ListJobRunsLifecycleStateEnum
func GetListJobRunsLifecycleStateEnumValues() []ListJobRunsLifecycleStateEnum {
	values := make([]ListJobRunsLifecycleStateEnum, 0)
	for _, v := range mappingListJobRunsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobRunsLifecycleStateEnumStringValues Enumerates the set of values in String for ListJobRunsLifecycleStateEnum
func GetListJobRunsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"DELETED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListJobRunsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobRunsLifecycleStateEnum(val string) (ListJobRunsLifecycleStateEnum, bool) {
	enum, ok := mappingListJobRunsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
