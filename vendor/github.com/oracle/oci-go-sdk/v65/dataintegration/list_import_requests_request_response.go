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

// ListImportRequestsRequest wrapper for the ListImportRequests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListImportRequests.go.html to see an example of how to use ListImportRequestsRequest.
type ListImportRequestsRequest struct {

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
	SortOrder ListImportRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListImportRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies import status to use, either -  ALL, SUCCESSFUL, IN_PROGRESS, QUEUED, FAILED .
	ImportStatus ListImportRequestsImportStatusEnum `mandatory:"false" contributesTo:"query" name:"importStatus" omitEmpty:"true"`

	// This parameter allows users to specify which view of the import object response to return. SUMMARY - Summary of the import object request will be returned. This is the default option when no value is specified. DETAILS - Details of import object request will be returned. This will include details of all the objects to be exported.
	Projection ListImportRequestsProjectionEnum `mandatory:"false" contributesTo:"query" name:"projection" omitEmpty:"true"`

	// Specifies start time of a copy object request.
	TimeStartedInMillis *int64 `mandatory:"false" contributesTo:"query" name:"timeStartedInMillis"`

	// Specifies end time of a copy object request.
	TimeEndedInMillis *int64 `mandatory:"false" contributesTo:"query" name:"timeEndedInMillis"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListImportRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListImportRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListImportRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListImportRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListImportRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListImportRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListImportRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListImportRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListImportRequestsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListImportRequestsImportStatusEnum(string(request.ImportStatus)); !ok && request.ImportStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImportStatus: %s. Supported values are: %s.", request.ImportStatus, strings.Join(GetListImportRequestsImportStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListImportRequestsProjectionEnum(string(request.Projection)); !ok && request.Projection != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Projection: %s. Supported values are: %s.", request.Projection, strings.Join(GetListImportRequestsProjectionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListImportRequestsResponse wrapper for the ListImportRequests operation
type ListImportRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ImportRequestSummaryCollection instances
	ImportRequestSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListImportRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListImportRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListImportRequestsSortOrderEnum Enum with underlying type: string
type ListImportRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListImportRequestsSortOrderEnum
const (
	ListImportRequestsSortOrderAsc  ListImportRequestsSortOrderEnum = "ASC"
	ListImportRequestsSortOrderDesc ListImportRequestsSortOrderEnum = "DESC"
)

var mappingListImportRequestsSortOrderEnum = map[string]ListImportRequestsSortOrderEnum{
	"ASC":  ListImportRequestsSortOrderAsc,
	"DESC": ListImportRequestsSortOrderDesc,
}

var mappingListImportRequestsSortOrderEnumLowerCase = map[string]ListImportRequestsSortOrderEnum{
	"asc":  ListImportRequestsSortOrderAsc,
	"desc": ListImportRequestsSortOrderDesc,
}

// GetListImportRequestsSortOrderEnumValues Enumerates the set of values for ListImportRequestsSortOrderEnum
func GetListImportRequestsSortOrderEnumValues() []ListImportRequestsSortOrderEnum {
	values := make([]ListImportRequestsSortOrderEnum, 0)
	for _, v := range mappingListImportRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListImportRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListImportRequestsSortOrderEnum
func GetListImportRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListImportRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListImportRequestsSortOrderEnum(val string) (ListImportRequestsSortOrderEnum, bool) {
	enum, ok := mappingListImportRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListImportRequestsSortByEnum Enum with underlying type: string
type ListImportRequestsSortByEnum string

// Set of constants representing the allowable values for ListImportRequestsSortByEnum
const (
	ListImportRequestsSortByTimeCreated ListImportRequestsSortByEnum = "TIME_CREATED"
	ListImportRequestsSortByDisplayName ListImportRequestsSortByEnum = "DISPLAY_NAME"
	ListImportRequestsSortByTimeUpdated ListImportRequestsSortByEnum = "TIME_UPDATED"
)

var mappingListImportRequestsSortByEnum = map[string]ListImportRequestsSortByEnum{
	"TIME_CREATED": ListImportRequestsSortByTimeCreated,
	"DISPLAY_NAME": ListImportRequestsSortByDisplayName,
	"TIME_UPDATED": ListImportRequestsSortByTimeUpdated,
}

var mappingListImportRequestsSortByEnumLowerCase = map[string]ListImportRequestsSortByEnum{
	"time_created": ListImportRequestsSortByTimeCreated,
	"display_name": ListImportRequestsSortByDisplayName,
	"time_updated": ListImportRequestsSortByTimeUpdated,
}

// GetListImportRequestsSortByEnumValues Enumerates the set of values for ListImportRequestsSortByEnum
func GetListImportRequestsSortByEnumValues() []ListImportRequestsSortByEnum {
	values := make([]ListImportRequestsSortByEnum, 0)
	for _, v := range mappingListImportRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListImportRequestsSortByEnumStringValues Enumerates the set of values in String for ListImportRequestsSortByEnum
func GetListImportRequestsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListImportRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListImportRequestsSortByEnum(val string) (ListImportRequestsSortByEnum, bool) {
	enum, ok := mappingListImportRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListImportRequestsImportStatusEnum Enum with underlying type: string
type ListImportRequestsImportStatusEnum string

// Set of constants representing the allowable values for ListImportRequestsImportStatusEnum
const (
	ListImportRequestsImportStatusInProgress  ListImportRequestsImportStatusEnum = "IN_PROGRESS"
	ListImportRequestsImportStatusSuccessful  ListImportRequestsImportStatusEnum = "SUCCESSFUL"
	ListImportRequestsImportStatusQueued      ListImportRequestsImportStatusEnum = "QUEUED"
	ListImportRequestsImportStatusTerminating ListImportRequestsImportStatusEnum = "TERMINATING"
	ListImportRequestsImportStatusTerminated  ListImportRequestsImportStatusEnum = "TERMINATED"
	ListImportRequestsImportStatusFailed      ListImportRequestsImportStatusEnum = "FAILED"
)

var mappingListImportRequestsImportStatusEnum = map[string]ListImportRequestsImportStatusEnum{
	"IN_PROGRESS": ListImportRequestsImportStatusInProgress,
	"SUCCESSFUL":  ListImportRequestsImportStatusSuccessful,
	"QUEUED":      ListImportRequestsImportStatusQueued,
	"TERMINATING": ListImportRequestsImportStatusTerminating,
	"TERMINATED":  ListImportRequestsImportStatusTerminated,
	"FAILED":      ListImportRequestsImportStatusFailed,
}

var mappingListImportRequestsImportStatusEnumLowerCase = map[string]ListImportRequestsImportStatusEnum{
	"in_progress": ListImportRequestsImportStatusInProgress,
	"successful":  ListImportRequestsImportStatusSuccessful,
	"queued":      ListImportRequestsImportStatusQueued,
	"terminating": ListImportRequestsImportStatusTerminating,
	"terminated":  ListImportRequestsImportStatusTerminated,
	"failed":      ListImportRequestsImportStatusFailed,
}

// GetListImportRequestsImportStatusEnumValues Enumerates the set of values for ListImportRequestsImportStatusEnum
func GetListImportRequestsImportStatusEnumValues() []ListImportRequestsImportStatusEnum {
	values := make([]ListImportRequestsImportStatusEnum, 0)
	for _, v := range mappingListImportRequestsImportStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListImportRequestsImportStatusEnumStringValues Enumerates the set of values in String for ListImportRequestsImportStatusEnum
func GetListImportRequestsImportStatusEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCESSFUL",
		"QUEUED",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingListImportRequestsImportStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListImportRequestsImportStatusEnum(val string) (ListImportRequestsImportStatusEnum, bool) {
	enum, ok := mappingListImportRequestsImportStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListImportRequestsProjectionEnum Enum with underlying type: string
type ListImportRequestsProjectionEnum string

// Set of constants representing the allowable values for ListImportRequestsProjectionEnum
const (
	ListImportRequestsProjectionSummary ListImportRequestsProjectionEnum = "SUMMARY"
	ListImportRequestsProjectionDetails ListImportRequestsProjectionEnum = "DETAILS"
)

var mappingListImportRequestsProjectionEnum = map[string]ListImportRequestsProjectionEnum{
	"SUMMARY": ListImportRequestsProjectionSummary,
	"DETAILS": ListImportRequestsProjectionDetails,
}

var mappingListImportRequestsProjectionEnumLowerCase = map[string]ListImportRequestsProjectionEnum{
	"summary": ListImportRequestsProjectionSummary,
	"details": ListImportRequestsProjectionDetails,
}

// GetListImportRequestsProjectionEnumValues Enumerates the set of values for ListImportRequestsProjectionEnum
func GetListImportRequestsProjectionEnumValues() []ListImportRequestsProjectionEnum {
	values := make([]ListImportRequestsProjectionEnum, 0)
	for _, v := range mappingListImportRequestsProjectionEnum {
		values = append(values, v)
	}
	return values
}

// GetListImportRequestsProjectionEnumStringValues Enumerates the set of values in String for ListImportRequestsProjectionEnum
func GetListImportRequestsProjectionEnumStringValues() []string {
	return []string{
		"SUMMARY",
		"DETAILS",
	}
}

// GetMappingListImportRequestsProjectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListImportRequestsProjectionEnum(val string) (ListImportRequestsProjectionEnum, bool) {
	enum, ok := mappingListImportRequestsProjectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
