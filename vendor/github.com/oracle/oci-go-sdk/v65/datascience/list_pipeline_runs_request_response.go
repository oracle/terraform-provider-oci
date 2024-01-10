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

// ListPipelineRunsRequest wrapper for the ListPipelineRuns operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListPipelineRuns.go.html to see an example of how to use ListPipelineRunsRequest.
type ListPipelineRunsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline.
	PipelineId *string `mandatory:"false" contributesTo:"query" name:"pipelineId"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The current state of the PipelineRun.
	LifecycleState ListPipelineRunsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

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
	SortOrder ListPipelineRunsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeAccepted`, the results are shown
	// in descending order. When you sort by `displayName`, the results are
	// shown in ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListPipelineRunsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPipelineRunsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPipelineRunsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPipelineRunsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPipelineRunsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPipelineRunsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPipelineRunsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListPipelineRunsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPipelineRunsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPipelineRunsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPipelineRunsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPipelineRunsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPipelineRunsResponse wrapper for the ListPipelineRuns operation
type ListPipelineRunsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []PipelineRunSummary instances
	Items []PipelineRunSummary `presentIn:"body"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListPipelineRunsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPipelineRunsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPipelineRunsLifecycleStateEnum Enum with underlying type: string
type ListPipelineRunsLifecycleStateEnum string

// Set of constants representing the allowable values for ListPipelineRunsLifecycleStateEnum
const (
	ListPipelineRunsLifecycleStateAccepted   ListPipelineRunsLifecycleStateEnum = "ACCEPTED"
	ListPipelineRunsLifecycleStateInProgress ListPipelineRunsLifecycleStateEnum = "IN_PROGRESS"
	ListPipelineRunsLifecycleStateFailed     ListPipelineRunsLifecycleStateEnum = "FAILED"
	ListPipelineRunsLifecycleStateSucceeded  ListPipelineRunsLifecycleStateEnum = "SUCCEEDED"
	ListPipelineRunsLifecycleStateCanceling  ListPipelineRunsLifecycleStateEnum = "CANCELING"
	ListPipelineRunsLifecycleStateCanceled   ListPipelineRunsLifecycleStateEnum = "CANCELED"
	ListPipelineRunsLifecycleStateDeleting   ListPipelineRunsLifecycleStateEnum = "DELETING"
	ListPipelineRunsLifecycleStateDeleted    ListPipelineRunsLifecycleStateEnum = "DELETED"
)

var mappingListPipelineRunsLifecycleStateEnum = map[string]ListPipelineRunsLifecycleStateEnum{
	"ACCEPTED":    ListPipelineRunsLifecycleStateAccepted,
	"IN_PROGRESS": ListPipelineRunsLifecycleStateInProgress,
	"FAILED":      ListPipelineRunsLifecycleStateFailed,
	"SUCCEEDED":   ListPipelineRunsLifecycleStateSucceeded,
	"CANCELING":   ListPipelineRunsLifecycleStateCanceling,
	"CANCELED":    ListPipelineRunsLifecycleStateCanceled,
	"DELETING":    ListPipelineRunsLifecycleStateDeleting,
	"DELETED":     ListPipelineRunsLifecycleStateDeleted,
}

var mappingListPipelineRunsLifecycleStateEnumLowerCase = map[string]ListPipelineRunsLifecycleStateEnum{
	"accepted":    ListPipelineRunsLifecycleStateAccepted,
	"in_progress": ListPipelineRunsLifecycleStateInProgress,
	"failed":      ListPipelineRunsLifecycleStateFailed,
	"succeeded":   ListPipelineRunsLifecycleStateSucceeded,
	"canceling":   ListPipelineRunsLifecycleStateCanceling,
	"canceled":    ListPipelineRunsLifecycleStateCanceled,
	"deleting":    ListPipelineRunsLifecycleStateDeleting,
	"deleted":     ListPipelineRunsLifecycleStateDeleted,
}

// GetListPipelineRunsLifecycleStateEnumValues Enumerates the set of values for ListPipelineRunsLifecycleStateEnum
func GetListPipelineRunsLifecycleStateEnumValues() []ListPipelineRunsLifecycleStateEnum {
	values := make([]ListPipelineRunsLifecycleStateEnum, 0)
	for _, v := range mappingListPipelineRunsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelineRunsLifecycleStateEnumStringValues Enumerates the set of values in String for ListPipelineRunsLifecycleStateEnum
func GetListPipelineRunsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingListPipelineRunsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelineRunsLifecycleStateEnum(val string) (ListPipelineRunsLifecycleStateEnum, bool) {
	enum, ok := mappingListPipelineRunsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPipelineRunsSortOrderEnum Enum with underlying type: string
type ListPipelineRunsSortOrderEnum string

// Set of constants representing the allowable values for ListPipelineRunsSortOrderEnum
const (
	ListPipelineRunsSortOrderAsc  ListPipelineRunsSortOrderEnum = "ASC"
	ListPipelineRunsSortOrderDesc ListPipelineRunsSortOrderEnum = "DESC"
)

var mappingListPipelineRunsSortOrderEnum = map[string]ListPipelineRunsSortOrderEnum{
	"ASC":  ListPipelineRunsSortOrderAsc,
	"DESC": ListPipelineRunsSortOrderDesc,
}

var mappingListPipelineRunsSortOrderEnumLowerCase = map[string]ListPipelineRunsSortOrderEnum{
	"asc":  ListPipelineRunsSortOrderAsc,
	"desc": ListPipelineRunsSortOrderDesc,
}

// GetListPipelineRunsSortOrderEnumValues Enumerates the set of values for ListPipelineRunsSortOrderEnum
func GetListPipelineRunsSortOrderEnumValues() []ListPipelineRunsSortOrderEnum {
	values := make([]ListPipelineRunsSortOrderEnum, 0)
	for _, v := range mappingListPipelineRunsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelineRunsSortOrderEnumStringValues Enumerates the set of values in String for ListPipelineRunsSortOrderEnum
func GetListPipelineRunsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPipelineRunsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelineRunsSortOrderEnum(val string) (ListPipelineRunsSortOrderEnum, bool) {
	enum, ok := mappingListPipelineRunsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPipelineRunsSortByEnum Enum with underlying type: string
type ListPipelineRunsSortByEnum string

// Set of constants representing the allowable values for ListPipelineRunsSortByEnum
const (
	ListPipelineRunsSortByTimeaccepted ListPipelineRunsSortByEnum = "timeAccepted"
	ListPipelineRunsSortByDisplayname  ListPipelineRunsSortByEnum = "displayName"
)

var mappingListPipelineRunsSortByEnum = map[string]ListPipelineRunsSortByEnum{
	"timeAccepted": ListPipelineRunsSortByTimeaccepted,
	"displayName":  ListPipelineRunsSortByDisplayname,
}

var mappingListPipelineRunsSortByEnumLowerCase = map[string]ListPipelineRunsSortByEnum{
	"timeaccepted": ListPipelineRunsSortByTimeaccepted,
	"displayname":  ListPipelineRunsSortByDisplayname,
}

// GetListPipelineRunsSortByEnumValues Enumerates the set of values for ListPipelineRunsSortByEnum
func GetListPipelineRunsSortByEnumValues() []ListPipelineRunsSortByEnum {
	values := make([]ListPipelineRunsSortByEnum, 0)
	for _, v := range mappingListPipelineRunsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelineRunsSortByEnumStringValues Enumerates the set of values in String for ListPipelineRunsSortByEnum
func GetListPipelineRunsSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
		"displayName",
	}
}

// GetMappingListPipelineRunsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelineRunsSortByEnum(val string) (ListPipelineRunsSortByEnum, bool) {
	enum, ok := mappingListPipelineRunsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
