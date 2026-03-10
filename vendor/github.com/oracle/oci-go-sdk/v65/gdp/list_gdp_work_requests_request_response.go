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

// ListGdpWorkRequestsRequest wrapper for the ListGdpWorkRequests operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/gdp/ListGdpWorkRequests.go.html to see an example of how to use ListGdpWorkRequestsRequest.
type ListGdpWorkRequestsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The ID of the asynchronous work request.
	WorkRequestId *string `mandatory:"false" contributesTo:"query" name:"workRequestId"`

	// A filter to return only resources with a lifecycleState that matches the given OperationStatus.
	Status ListGdpWorkRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The ID of the resource affected by the work request.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListGdpWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeAccepted is descending.
	SortBy ListGdpWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGdpWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGdpWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGdpWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGdpWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGdpWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGdpWorkRequestsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListGdpWorkRequestsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGdpWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGdpWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGdpWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGdpWorkRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGdpWorkRequestsResponse wrapper for the ListGdpWorkRequests operation
type ListGdpWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GdpWorkRequestSummaryCollection instances
	GdpWorkRequestSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGdpWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGdpWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGdpWorkRequestsStatusEnum Enum with underlying type: string
type ListGdpWorkRequestsStatusEnum string

// Set of constants representing the allowable values for ListGdpWorkRequestsStatusEnum
const (
	ListGdpWorkRequestsStatusAccepted       ListGdpWorkRequestsStatusEnum = "ACCEPTED"
	ListGdpWorkRequestsStatusInProgress     ListGdpWorkRequestsStatusEnum = "IN_PROGRESS"
	ListGdpWorkRequestsStatusWaiting        ListGdpWorkRequestsStatusEnum = "WAITING"
	ListGdpWorkRequestsStatusNeedsAttention ListGdpWorkRequestsStatusEnum = "NEEDS_ATTENTION"
	ListGdpWorkRequestsStatusFailed         ListGdpWorkRequestsStatusEnum = "FAILED"
	ListGdpWorkRequestsStatusSucceeded      ListGdpWorkRequestsStatusEnum = "SUCCEEDED"
	ListGdpWorkRequestsStatusCanceling      ListGdpWorkRequestsStatusEnum = "CANCELING"
	ListGdpWorkRequestsStatusCanceled       ListGdpWorkRequestsStatusEnum = "CANCELED"
)

var mappingListGdpWorkRequestsStatusEnum = map[string]ListGdpWorkRequestsStatusEnum{
	"ACCEPTED":        ListGdpWorkRequestsStatusAccepted,
	"IN_PROGRESS":     ListGdpWorkRequestsStatusInProgress,
	"WAITING":         ListGdpWorkRequestsStatusWaiting,
	"NEEDS_ATTENTION": ListGdpWorkRequestsStatusNeedsAttention,
	"FAILED":          ListGdpWorkRequestsStatusFailed,
	"SUCCEEDED":       ListGdpWorkRequestsStatusSucceeded,
	"CANCELING":       ListGdpWorkRequestsStatusCanceling,
	"CANCELED":        ListGdpWorkRequestsStatusCanceled,
}

var mappingListGdpWorkRequestsStatusEnumLowerCase = map[string]ListGdpWorkRequestsStatusEnum{
	"accepted":        ListGdpWorkRequestsStatusAccepted,
	"in_progress":     ListGdpWorkRequestsStatusInProgress,
	"waiting":         ListGdpWorkRequestsStatusWaiting,
	"needs_attention": ListGdpWorkRequestsStatusNeedsAttention,
	"failed":          ListGdpWorkRequestsStatusFailed,
	"succeeded":       ListGdpWorkRequestsStatusSucceeded,
	"canceling":       ListGdpWorkRequestsStatusCanceling,
	"canceled":        ListGdpWorkRequestsStatusCanceled,
}

// GetListGdpWorkRequestsStatusEnumValues Enumerates the set of values for ListGdpWorkRequestsStatusEnum
func GetListGdpWorkRequestsStatusEnumValues() []ListGdpWorkRequestsStatusEnum {
	values := make([]ListGdpWorkRequestsStatusEnum, 0)
	for _, v := range mappingListGdpWorkRequestsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListGdpWorkRequestsStatusEnumStringValues Enumerates the set of values in String for ListGdpWorkRequestsStatusEnum
func GetListGdpWorkRequestsStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"NEEDS_ATTENTION",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingListGdpWorkRequestsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGdpWorkRequestsStatusEnum(val string) (ListGdpWorkRequestsStatusEnum, bool) {
	enum, ok := mappingListGdpWorkRequestsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGdpWorkRequestsSortOrderEnum Enum with underlying type: string
type ListGdpWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListGdpWorkRequestsSortOrderEnum
const (
	ListGdpWorkRequestsSortOrderAsc  ListGdpWorkRequestsSortOrderEnum = "ASC"
	ListGdpWorkRequestsSortOrderDesc ListGdpWorkRequestsSortOrderEnum = "DESC"
)

var mappingListGdpWorkRequestsSortOrderEnum = map[string]ListGdpWorkRequestsSortOrderEnum{
	"ASC":  ListGdpWorkRequestsSortOrderAsc,
	"DESC": ListGdpWorkRequestsSortOrderDesc,
}

var mappingListGdpWorkRequestsSortOrderEnumLowerCase = map[string]ListGdpWorkRequestsSortOrderEnum{
	"asc":  ListGdpWorkRequestsSortOrderAsc,
	"desc": ListGdpWorkRequestsSortOrderDesc,
}

// GetListGdpWorkRequestsSortOrderEnumValues Enumerates the set of values for ListGdpWorkRequestsSortOrderEnum
func GetListGdpWorkRequestsSortOrderEnumValues() []ListGdpWorkRequestsSortOrderEnum {
	values := make([]ListGdpWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListGdpWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGdpWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListGdpWorkRequestsSortOrderEnum
func GetListGdpWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGdpWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGdpWorkRequestsSortOrderEnum(val string) (ListGdpWorkRequestsSortOrderEnum, bool) {
	enum, ok := mappingListGdpWorkRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGdpWorkRequestsSortByEnum Enum with underlying type: string
type ListGdpWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListGdpWorkRequestsSortByEnum
const (
	ListGdpWorkRequestsSortByTimeaccepted ListGdpWorkRequestsSortByEnum = "timeAccepted"
)

var mappingListGdpWorkRequestsSortByEnum = map[string]ListGdpWorkRequestsSortByEnum{
	"timeAccepted": ListGdpWorkRequestsSortByTimeaccepted,
}

var mappingListGdpWorkRequestsSortByEnumLowerCase = map[string]ListGdpWorkRequestsSortByEnum{
	"timeaccepted": ListGdpWorkRequestsSortByTimeaccepted,
}

// GetListGdpWorkRequestsSortByEnumValues Enumerates the set of values for ListGdpWorkRequestsSortByEnum
func GetListGdpWorkRequestsSortByEnumValues() []ListGdpWorkRequestsSortByEnum {
	values := make([]ListGdpWorkRequestsSortByEnum, 0)
	for _, v := range mappingListGdpWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGdpWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListGdpWorkRequestsSortByEnum
func GetListGdpWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
	}
}

// GetMappingListGdpWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGdpWorkRequestsSortByEnum(val string) (ListGdpWorkRequestsSortByEnum, bool) {
	enum, ok := mappingListGdpWorkRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
