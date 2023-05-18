// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExportRequestsRequest wrapper for the ListExportRequests operation
type ListExportRequestsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListExportRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListExportRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies export status to use, either -  ALL, SUCCESSFUL, IN_PROGRESS, QUEUED, FAILED .
	Status ListExportRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// Specifies start time of a copy object request.
	TimeStartedInMillis *int64 `mandatory:"false" contributesTo:"query" name:"timeStartedInMillis"`

	// Specifies end time of a copy object request.
	TimeEndedInMillis *int64 `mandatory:"false" contributesTo:"query" name:"timeEndedInMillis"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExportRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExportRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExportRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExportRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExportRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExportRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExportRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExportRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExportRequestsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExportRequestsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListExportRequestsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExportRequestsResponse wrapper for the ListExportRequests operation
type ListExportRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExportRequestSummaryCollection instances
	ExportRequestSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExportRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExportRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExportRequestsSortOrderEnum Enum with underlying type: string
type ListExportRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListExportRequestsSortOrderEnum
const (
	ListExportRequestsSortOrderAsc  ListExportRequestsSortOrderEnum = "ASC"
	ListExportRequestsSortOrderDesc ListExportRequestsSortOrderEnum = "DESC"
)

var mappingListExportRequestsSortOrderEnum = map[string]ListExportRequestsSortOrderEnum{
	"ASC":  ListExportRequestsSortOrderAsc,
	"DESC": ListExportRequestsSortOrderDesc,
}

var mappingListExportRequestsSortOrderEnumLowerCase = map[string]ListExportRequestsSortOrderEnum{
	"asc":  ListExportRequestsSortOrderAsc,
	"desc": ListExportRequestsSortOrderDesc,
}

// GetListExportRequestsSortOrderEnumValues Enumerates the set of values for ListExportRequestsSortOrderEnum
func GetListExportRequestsSortOrderEnumValues() []ListExportRequestsSortOrderEnum {
	values := make([]ListExportRequestsSortOrderEnum, 0)
	for _, v := range mappingListExportRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExportRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListExportRequestsSortOrderEnum
func GetListExportRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExportRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExportRequestsSortOrderEnum(val string) (ListExportRequestsSortOrderEnum, bool) {
	enum, ok := mappingListExportRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExportRequestsSortByEnum Enum with underlying type: string
type ListExportRequestsSortByEnum string

// Set of constants representing the allowable values for ListExportRequestsSortByEnum
const (
	ListExportRequestsSortByTimeCreated ListExportRequestsSortByEnum = "TIME_CREATED"
	ListExportRequestsSortByDisplayName ListExportRequestsSortByEnum = "DISPLAY_NAME"
	ListExportRequestsSortByTimeUpdated ListExportRequestsSortByEnum = "TIME_UPDATED"
)

var mappingListExportRequestsSortByEnum = map[string]ListExportRequestsSortByEnum{
	"TIME_CREATED": ListExportRequestsSortByTimeCreated,
	"DISPLAY_NAME": ListExportRequestsSortByDisplayName,
	"TIME_UPDATED": ListExportRequestsSortByTimeUpdated,
}

var mappingListExportRequestsSortByEnumLowerCase = map[string]ListExportRequestsSortByEnum{
	"time_created": ListExportRequestsSortByTimeCreated,
	"display_name": ListExportRequestsSortByDisplayName,
	"time_updated": ListExportRequestsSortByTimeUpdated,
}

// GetListExportRequestsSortByEnumValues Enumerates the set of values for ListExportRequestsSortByEnum
func GetListExportRequestsSortByEnumValues() []ListExportRequestsSortByEnum {
	values := make([]ListExportRequestsSortByEnum, 0)
	for _, v := range mappingListExportRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExportRequestsSortByEnumStringValues Enumerates the set of values in String for ListExportRequestsSortByEnum
func GetListExportRequestsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListExportRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExportRequestsSortByEnum(val string) (ListExportRequestsSortByEnum, bool) {
	enum, ok := mappingListExportRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExportRequestsStatusEnum Enum with underlying type: string
type ListExportRequestsStatusEnum string

// Set of constants representing the allowable values for ListExportRequestsStatusEnum
const (
	ListExportRequestsStatusInProgress  ListExportRequestsStatusEnum = "IN_PROGRESS"
	ListExportRequestsStatusSuccessful  ListExportRequestsStatusEnum = "SUCCESSFUL"
	ListExportRequestsStatusQueued      ListExportRequestsStatusEnum = "QUEUED"
	ListExportRequestsStatusTerminating ListExportRequestsStatusEnum = "TERMINATING"
	ListExportRequestsStatusTerminated  ListExportRequestsStatusEnum = "TERMINATED"
	ListExportRequestsStatusFailed      ListExportRequestsStatusEnum = "FAILED"
)

var mappingListExportRequestsStatusEnum = map[string]ListExportRequestsStatusEnum{
	"IN_PROGRESS": ListExportRequestsStatusInProgress,
	"SUCCESSFUL":  ListExportRequestsStatusSuccessful,
	"QUEUED":      ListExportRequestsStatusQueued,
	"TERMINATING": ListExportRequestsStatusTerminating,
	"TERMINATED":  ListExportRequestsStatusTerminated,
	"FAILED":      ListExportRequestsStatusFailed,
}

var mappingListExportRequestsStatusEnumLowerCase = map[string]ListExportRequestsStatusEnum{
	"in_progress": ListExportRequestsStatusInProgress,
	"successful":  ListExportRequestsStatusSuccessful,
	"queued":      ListExportRequestsStatusQueued,
	"terminating": ListExportRequestsStatusTerminating,
	"terminated":  ListExportRequestsStatusTerminated,
	"failed":      ListExportRequestsStatusFailed,
}

// GetListExportRequestsStatusEnumValues Enumerates the set of values for ListExportRequestsStatusEnum
func GetListExportRequestsStatusEnumValues() []ListExportRequestsStatusEnum {
	values := make([]ListExportRequestsStatusEnum, 0)
	for _, v := range mappingListExportRequestsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListExportRequestsStatusEnumStringValues Enumerates the set of values in String for ListExportRequestsStatusEnum
func GetListExportRequestsStatusEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCESSFUL",
		"QUEUED",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingListExportRequestsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExportRequestsStatusEnum(val string) (ListExportRequestsStatusEnum, bool) {
	enum, ok := mappingListExportRequestsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
