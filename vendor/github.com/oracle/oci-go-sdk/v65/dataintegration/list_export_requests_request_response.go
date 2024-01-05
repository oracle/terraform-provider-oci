// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListExportRequests.go.html to see an example of how to use ListExportRequestsRequest.
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
	ExportStatus ListExportRequestsExportStatusEnum `mandatory:"false" contributesTo:"query" name:"exportStatus" omitEmpty:"true"`

	// This parameter allows users to specify which view of the export object response to return. SUMMARY - Summary of the export object request will be returned. This is the default option when no value is specified. DETAILS - Details of export object request will be returned. This will include details of all the objects to be exported.
	Projection ListExportRequestsProjectionEnum `mandatory:"false" contributesTo:"query" name:"projection" omitEmpty:"true"`

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
	if _, ok := GetMappingListExportRequestsExportStatusEnum(string(request.ExportStatus)); !ok && request.ExportStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExportStatus: %s. Supported values are: %s.", request.ExportStatus, strings.Join(GetListExportRequestsExportStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExportRequestsProjectionEnum(string(request.Projection)); !ok && request.Projection != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Projection: %s. Supported values are: %s.", request.Projection, strings.Join(GetListExportRequestsProjectionEnumStringValues(), ",")))
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

// ListExportRequestsExportStatusEnum Enum with underlying type: string
type ListExportRequestsExportStatusEnum string

// Set of constants representing the allowable values for ListExportRequestsExportStatusEnum
const (
	ListExportRequestsExportStatusInProgress  ListExportRequestsExportStatusEnum = "IN_PROGRESS"
	ListExportRequestsExportStatusSuccessful  ListExportRequestsExportStatusEnum = "SUCCESSFUL"
	ListExportRequestsExportStatusQueued      ListExportRequestsExportStatusEnum = "QUEUED"
	ListExportRequestsExportStatusTerminating ListExportRequestsExportStatusEnum = "TERMINATING"
	ListExportRequestsExportStatusTerminated  ListExportRequestsExportStatusEnum = "TERMINATED"
	ListExportRequestsExportStatusFailed      ListExportRequestsExportStatusEnum = "FAILED"
)

var mappingListExportRequestsExportStatusEnum = map[string]ListExportRequestsExportStatusEnum{
	"IN_PROGRESS": ListExportRequestsExportStatusInProgress,
	"SUCCESSFUL":  ListExportRequestsExportStatusSuccessful,
	"QUEUED":      ListExportRequestsExportStatusQueued,
	"TERMINATING": ListExportRequestsExportStatusTerminating,
	"TERMINATED":  ListExportRequestsExportStatusTerminated,
	"FAILED":      ListExportRequestsExportStatusFailed,
}

var mappingListExportRequestsExportStatusEnumLowerCase = map[string]ListExportRequestsExportStatusEnum{
	"in_progress": ListExportRequestsExportStatusInProgress,
	"successful":  ListExportRequestsExportStatusSuccessful,
	"queued":      ListExportRequestsExportStatusQueued,
	"terminating": ListExportRequestsExportStatusTerminating,
	"terminated":  ListExportRequestsExportStatusTerminated,
	"failed":      ListExportRequestsExportStatusFailed,
}

// GetListExportRequestsExportStatusEnumValues Enumerates the set of values for ListExportRequestsExportStatusEnum
func GetListExportRequestsExportStatusEnumValues() []ListExportRequestsExportStatusEnum {
	values := make([]ListExportRequestsExportStatusEnum, 0)
	for _, v := range mappingListExportRequestsExportStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListExportRequestsExportStatusEnumStringValues Enumerates the set of values in String for ListExportRequestsExportStatusEnum
func GetListExportRequestsExportStatusEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCESSFUL",
		"QUEUED",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingListExportRequestsExportStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExportRequestsExportStatusEnum(val string) (ListExportRequestsExportStatusEnum, bool) {
	enum, ok := mappingListExportRequestsExportStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExportRequestsProjectionEnum Enum with underlying type: string
type ListExportRequestsProjectionEnum string

// Set of constants representing the allowable values for ListExportRequestsProjectionEnum
const (
	ListExportRequestsProjectionSummary ListExportRequestsProjectionEnum = "SUMMARY"
	ListExportRequestsProjectionDetails ListExportRequestsProjectionEnum = "DETAILS"
)

var mappingListExportRequestsProjectionEnum = map[string]ListExportRequestsProjectionEnum{
	"SUMMARY": ListExportRequestsProjectionSummary,
	"DETAILS": ListExportRequestsProjectionDetails,
}

var mappingListExportRequestsProjectionEnumLowerCase = map[string]ListExportRequestsProjectionEnum{
	"summary": ListExportRequestsProjectionSummary,
	"details": ListExportRequestsProjectionDetails,
}

// GetListExportRequestsProjectionEnumValues Enumerates the set of values for ListExportRequestsProjectionEnum
func GetListExportRequestsProjectionEnumValues() []ListExportRequestsProjectionEnum {
	values := make([]ListExportRequestsProjectionEnum, 0)
	for _, v := range mappingListExportRequestsProjectionEnum {
		values = append(values, v)
	}
	return values
}

// GetListExportRequestsProjectionEnumStringValues Enumerates the set of values in String for ListExportRequestsProjectionEnum
func GetListExportRequestsProjectionEnumStringValues() []string {
	return []string{
		"SUMMARY",
		"DETAILS",
	}
}

// GetMappingListExportRequestsProjectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExportRequestsProjectionEnum(val string) (ListExportRequestsProjectionEnum, bool) {
	enum, ok := mappingListExportRequestsProjectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
