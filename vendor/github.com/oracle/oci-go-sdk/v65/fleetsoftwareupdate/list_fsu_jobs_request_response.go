// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFsuJobsRequest wrapper for the ListFsuJobs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuJobs.go.html to see an example of how to use ListFsuJobsRequest.
type ListFsuJobsRequest struct {

	// The ID of the compartment in which to list resources.
	FsuActionId *string `mandatory:"true" contributesTo:"query" name:"fsuActionId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState ListFsuJobsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListFsuJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFsuJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFsuJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFsuJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFsuJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFsuJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFsuJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFsuJobsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFsuJobsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFsuJobsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFsuJobsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFsuJobsResponse wrapper for the ListFsuJobs operation
type ListFsuJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FsuJobCollection instances
	FsuJobCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFsuJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFsuJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFsuJobsLifecycleStateEnum Enum with underlying type: string
type ListFsuJobsLifecycleStateEnum string

// Set of constants representing the allowable values for ListFsuJobsLifecycleStateEnum
const (
	ListFsuJobsLifecycleStateAccepted       ListFsuJobsLifecycleStateEnum = "ACCEPTED"
	ListFsuJobsLifecycleStateInProgress     ListFsuJobsLifecycleStateEnum = "IN_PROGRESS"
	ListFsuJobsLifecycleStateUnknown        ListFsuJobsLifecycleStateEnum = "UNKNOWN"
	ListFsuJobsLifecycleStateTerminated     ListFsuJobsLifecycleStateEnum = "TERMINATED"
	ListFsuJobsLifecycleStateFailed         ListFsuJobsLifecycleStateEnum = "FAILED"
	ListFsuJobsLifecycleStateNeedsAttention ListFsuJobsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListFsuJobsLifecycleStateSucceeded      ListFsuJobsLifecycleStateEnum = "SUCCEEDED"
	ListFsuJobsLifecycleStateWaiting        ListFsuJobsLifecycleStateEnum = "WAITING"
	ListFsuJobsLifecycleStateCanceling      ListFsuJobsLifecycleStateEnum = "CANCELING"
	ListFsuJobsLifecycleStateCanceled       ListFsuJobsLifecycleStateEnum = "CANCELED"
)

var mappingListFsuJobsLifecycleStateEnum = map[string]ListFsuJobsLifecycleStateEnum{
	"ACCEPTED":        ListFsuJobsLifecycleStateAccepted,
	"IN_PROGRESS":     ListFsuJobsLifecycleStateInProgress,
	"UNKNOWN":         ListFsuJobsLifecycleStateUnknown,
	"TERMINATED":      ListFsuJobsLifecycleStateTerminated,
	"FAILED":          ListFsuJobsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListFsuJobsLifecycleStateNeedsAttention,
	"SUCCEEDED":       ListFsuJobsLifecycleStateSucceeded,
	"WAITING":         ListFsuJobsLifecycleStateWaiting,
	"CANCELING":       ListFsuJobsLifecycleStateCanceling,
	"CANCELED":        ListFsuJobsLifecycleStateCanceled,
}

var mappingListFsuJobsLifecycleStateEnumLowerCase = map[string]ListFsuJobsLifecycleStateEnum{
	"accepted":        ListFsuJobsLifecycleStateAccepted,
	"in_progress":     ListFsuJobsLifecycleStateInProgress,
	"unknown":         ListFsuJobsLifecycleStateUnknown,
	"terminated":      ListFsuJobsLifecycleStateTerminated,
	"failed":          ListFsuJobsLifecycleStateFailed,
	"needs_attention": ListFsuJobsLifecycleStateNeedsAttention,
	"succeeded":       ListFsuJobsLifecycleStateSucceeded,
	"waiting":         ListFsuJobsLifecycleStateWaiting,
	"canceling":       ListFsuJobsLifecycleStateCanceling,
	"canceled":        ListFsuJobsLifecycleStateCanceled,
}

// GetListFsuJobsLifecycleStateEnumValues Enumerates the set of values for ListFsuJobsLifecycleStateEnum
func GetListFsuJobsLifecycleStateEnumValues() []ListFsuJobsLifecycleStateEnum {
	values := make([]ListFsuJobsLifecycleStateEnum, 0)
	for _, v := range mappingListFsuJobsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuJobsLifecycleStateEnumStringValues Enumerates the set of values in String for ListFsuJobsLifecycleStateEnum
func GetListFsuJobsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"UNKNOWN",
		"TERMINATED",
		"FAILED",
		"NEEDS_ATTENTION",
		"SUCCEEDED",
		"WAITING",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingListFsuJobsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuJobsLifecycleStateEnum(val string) (ListFsuJobsLifecycleStateEnum, bool) {
	enum, ok := mappingListFsuJobsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuJobsSortByEnum Enum with underlying type: string
type ListFsuJobsSortByEnum string

// Set of constants representing the allowable values for ListFsuJobsSortByEnum
const (
	ListFsuJobsSortByTimecreated ListFsuJobsSortByEnum = "timeCreated"
	ListFsuJobsSortByDisplayname ListFsuJobsSortByEnum = "displayName"
)

var mappingListFsuJobsSortByEnum = map[string]ListFsuJobsSortByEnum{
	"timeCreated": ListFsuJobsSortByTimecreated,
	"displayName": ListFsuJobsSortByDisplayname,
}

var mappingListFsuJobsSortByEnumLowerCase = map[string]ListFsuJobsSortByEnum{
	"timecreated": ListFsuJobsSortByTimecreated,
	"displayname": ListFsuJobsSortByDisplayname,
}

// GetListFsuJobsSortByEnumValues Enumerates the set of values for ListFsuJobsSortByEnum
func GetListFsuJobsSortByEnumValues() []ListFsuJobsSortByEnum {
	values := make([]ListFsuJobsSortByEnum, 0)
	for _, v := range mappingListFsuJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuJobsSortByEnumStringValues Enumerates the set of values in String for ListFsuJobsSortByEnum
func GetListFsuJobsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListFsuJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuJobsSortByEnum(val string) (ListFsuJobsSortByEnum, bool) {
	enum, ok := mappingListFsuJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuJobsSortOrderEnum Enum with underlying type: string
type ListFsuJobsSortOrderEnum string

// Set of constants representing the allowable values for ListFsuJobsSortOrderEnum
const (
	ListFsuJobsSortOrderAsc  ListFsuJobsSortOrderEnum = "ASC"
	ListFsuJobsSortOrderDesc ListFsuJobsSortOrderEnum = "DESC"
)

var mappingListFsuJobsSortOrderEnum = map[string]ListFsuJobsSortOrderEnum{
	"ASC":  ListFsuJobsSortOrderAsc,
	"DESC": ListFsuJobsSortOrderDesc,
}

var mappingListFsuJobsSortOrderEnumLowerCase = map[string]ListFsuJobsSortOrderEnum{
	"asc":  ListFsuJobsSortOrderAsc,
	"desc": ListFsuJobsSortOrderDesc,
}

// GetListFsuJobsSortOrderEnumValues Enumerates the set of values for ListFsuJobsSortOrderEnum
func GetListFsuJobsSortOrderEnumValues() []ListFsuJobsSortOrderEnum {
	values := make([]ListFsuJobsSortOrderEnum, 0)
	for _, v := range mappingListFsuJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuJobsSortOrderEnumStringValues Enumerates the set of values in String for ListFsuJobsSortOrderEnum
func GetListFsuJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFsuJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuJobsSortOrderEnum(val string) (ListFsuJobsSortOrderEnum, bool) {
	enum, ok := mappingListFsuJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
