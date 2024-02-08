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

// ListCopyObjectRequestsRequest wrapper for the ListCopyObjectRequests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListCopyObjectRequests.go.html to see an example of how to use ListCopyObjectRequestsRequest.
type ListCopyObjectRequestsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListCopyObjectRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListCopyObjectRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies copy status to use, either -  ALL, SUCCESSFUL, IN_PROGRESS, QUEUED, FAILED .
	CopyStatus ListCopyObjectRequestsCopyStatusEnum `mandatory:"false" contributesTo:"query" name:"copyStatus" omitEmpty:"true"`

	// This parameter allows users to specify which view of the copy object response to return. SUMMARY - Summary of the copy object response will be returned. This is the default option when no value is specified. DETAILS - Details of copy object response will be returned. This will include details of all the objects to be copied.
	Projection ListCopyObjectRequestsProjectionEnum `mandatory:"false" contributesTo:"query" name:"projection" omitEmpty:"true"`

	// Specifies start time of a copy object request.
	TimeStartedInMillis *int64 `mandatory:"false" contributesTo:"query" name:"timeStartedInMillis"`

	// Specifies end time of a copy object request.
	TimeEndedInMillis *int64 `mandatory:"false" contributesTo:"query" name:"timeEndedInMillis"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCopyObjectRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCopyObjectRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCopyObjectRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCopyObjectRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCopyObjectRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCopyObjectRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCopyObjectRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCopyObjectRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCopyObjectRequestsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCopyObjectRequestsCopyStatusEnum(string(request.CopyStatus)); !ok && request.CopyStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CopyStatus: %s. Supported values are: %s.", request.CopyStatus, strings.Join(GetListCopyObjectRequestsCopyStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCopyObjectRequestsProjectionEnum(string(request.Projection)); !ok && request.Projection != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Projection: %s. Supported values are: %s.", request.Projection, strings.Join(GetListCopyObjectRequestsProjectionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCopyObjectRequestsResponse wrapper for the ListCopyObjectRequests operation
type ListCopyObjectRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CopyObjectRequestSummaryCollection instances
	CopyObjectRequestSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCopyObjectRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCopyObjectRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCopyObjectRequestsSortOrderEnum Enum with underlying type: string
type ListCopyObjectRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListCopyObjectRequestsSortOrderEnum
const (
	ListCopyObjectRequestsSortOrderAsc  ListCopyObjectRequestsSortOrderEnum = "ASC"
	ListCopyObjectRequestsSortOrderDesc ListCopyObjectRequestsSortOrderEnum = "DESC"
)

var mappingListCopyObjectRequestsSortOrderEnum = map[string]ListCopyObjectRequestsSortOrderEnum{
	"ASC":  ListCopyObjectRequestsSortOrderAsc,
	"DESC": ListCopyObjectRequestsSortOrderDesc,
}

var mappingListCopyObjectRequestsSortOrderEnumLowerCase = map[string]ListCopyObjectRequestsSortOrderEnum{
	"asc":  ListCopyObjectRequestsSortOrderAsc,
	"desc": ListCopyObjectRequestsSortOrderDesc,
}

// GetListCopyObjectRequestsSortOrderEnumValues Enumerates the set of values for ListCopyObjectRequestsSortOrderEnum
func GetListCopyObjectRequestsSortOrderEnumValues() []ListCopyObjectRequestsSortOrderEnum {
	values := make([]ListCopyObjectRequestsSortOrderEnum, 0)
	for _, v := range mappingListCopyObjectRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCopyObjectRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListCopyObjectRequestsSortOrderEnum
func GetListCopyObjectRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCopyObjectRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCopyObjectRequestsSortOrderEnum(val string) (ListCopyObjectRequestsSortOrderEnum, bool) {
	enum, ok := mappingListCopyObjectRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCopyObjectRequestsSortByEnum Enum with underlying type: string
type ListCopyObjectRequestsSortByEnum string

// Set of constants representing the allowable values for ListCopyObjectRequestsSortByEnum
const (
	ListCopyObjectRequestsSortByTimeCreated ListCopyObjectRequestsSortByEnum = "TIME_CREATED"
	ListCopyObjectRequestsSortByDisplayName ListCopyObjectRequestsSortByEnum = "DISPLAY_NAME"
	ListCopyObjectRequestsSortByTimeUpdated ListCopyObjectRequestsSortByEnum = "TIME_UPDATED"
)

var mappingListCopyObjectRequestsSortByEnum = map[string]ListCopyObjectRequestsSortByEnum{
	"TIME_CREATED": ListCopyObjectRequestsSortByTimeCreated,
	"DISPLAY_NAME": ListCopyObjectRequestsSortByDisplayName,
	"TIME_UPDATED": ListCopyObjectRequestsSortByTimeUpdated,
}

var mappingListCopyObjectRequestsSortByEnumLowerCase = map[string]ListCopyObjectRequestsSortByEnum{
	"time_created": ListCopyObjectRequestsSortByTimeCreated,
	"display_name": ListCopyObjectRequestsSortByDisplayName,
	"time_updated": ListCopyObjectRequestsSortByTimeUpdated,
}

// GetListCopyObjectRequestsSortByEnumValues Enumerates the set of values for ListCopyObjectRequestsSortByEnum
func GetListCopyObjectRequestsSortByEnumValues() []ListCopyObjectRequestsSortByEnum {
	values := make([]ListCopyObjectRequestsSortByEnum, 0)
	for _, v := range mappingListCopyObjectRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCopyObjectRequestsSortByEnumStringValues Enumerates the set of values in String for ListCopyObjectRequestsSortByEnum
func GetListCopyObjectRequestsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListCopyObjectRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCopyObjectRequestsSortByEnum(val string) (ListCopyObjectRequestsSortByEnum, bool) {
	enum, ok := mappingListCopyObjectRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCopyObjectRequestsCopyStatusEnum Enum with underlying type: string
type ListCopyObjectRequestsCopyStatusEnum string

// Set of constants representing the allowable values for ListCopyObjectRequestsCopyStatusEnum
const (
	ListCopyObjectRequestsCopyStatusInProgress  ListCopyObjectRequestsCopyStatusEnum = "IN_PROGRESS"
	ListCopyObjectRequestsCopyStatusSuccessful  ListCopyObjectRequestsCopyStatusEnum = "SUCCESSFUL"
	ListCopyObjectRequestsCopyStatusQueued      ListCopyObjectRequestsCopyStatusEnum = "QUEUED"
	ListCopyObjectRequestsCopyStatusTerminating ListCopyObjectRequestsCopyStatusEnum = "TERMINATING"
	ListCopyObjectRequestsCopyStatusTerminated  ListCopyObjectRequestsCopyStatusEnum = "TERMINATED"
	ListCopyObjectRequestsCopyStatusFailed      ListCopyObjectRequestsCopyStatusEnum = "FAILED"
	ListCopyObjectRequestsCopyStatusAll         ListCopyObjectRequestsCopyStatusEnum = "ALL"
)

var mappingListCopyObjectRequestsCopyStatusEnum = map[string]ListCopyObjectRequestsCopyStatusEnum{
	"IN_PROGRESS": ListCopyObjectRequestsCopyStatusInProgress,
	"SUCCESSFUL":  ListCopyObjectRequestsCopyStatusSuccessful,
	"QUEUED":      ListCopyObjectRequestsCopyStatusQueued,
	"TERMINATING": ListCopyObjectRequestsCopyStatusTerminating,
	"TERMINATED":  ListCopyObjectRequestsCopyStatusTerminated,
	"FAILED":      ListCopyObjectRequestsCopyStatusFailed,
	"ALL":         ListCopyObjectRequestsCopyStatusAll,
}

var mappingListCopyObjectRequestsCopyStatusEnumLowerCase = map[string]ListCopyObjectRequestsCopyStatusEnum{
	"in_progress": ListCopyObjectRequestsCopyStatusInProgress,
	"successful":  ListCopyObjectRequestsCopyStatusSuccessful,
	"queued":      ListCopyObjectRequestsCopyStatusQueued,
	"terminating": ListCopyObjectRequestsCopyStatusTerminating,
	"terminated":  ListCopyObjectRequestsCopyStatusTerminated,
	"failed":      ListCopyObjectRequestsCopyStatusFailed,
	"all":         ListCopyObjectRequestsCopyStatusAll,
}

// GetListCopyObjectRequestsCopyStatusEnumValues Enumerates the set of values for ListCopyObjectRequestsCopyStatusEnum
func GetListCopyObjectRequestsCopyStatusEnumValues() []ListCopyObjectRequestsCopyStatusEnum {
	values := make([]ListCopyObjectRequestsCopyStatusEnum, 0)
	for _, v := range mappingListCopyObjectRequestsCopyStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListCopyObjectRequestsCopyStatusEnumStringValues Enumerates the set of values in String for ListCopyObjectRequestsCopyStatusEnum
func GetListCopyObjectRequestsCopyStatusEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCESSFUL",
		"QUEUED",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"ALL",
	}
}

// GetMappingListCopyObjectRequestsCopyStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCopyObjectRequestsCopyStatusEnum(val string) (ListCopyObjectRequestsCopyStatusEnum, bool) {
	enum, ok := mappingListCopyObjectRequestsCopyStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCopyObjectRequestsProjectionEnum Enum with underlying type: string
type ListCopyObjectRequestsProjectionEnum string

// Set of constants representing the allowable values for ListCopyObjectRequestsProjectionEnum
const (
	ListCopyObjectRequestsProjectionSummary ListCopyObjectRequestsProjectionEnum = "SUMMARY"
	ListCopyObjectRequestsProjectionDetails ListCopyObjectRequestsProjectionEnum = "DETAILS"
)

var mappingListCopyObjectRequestsProjectionEnum = map[string]ListCopyObjectRequestsProjectionEnum{
	"SUMMARY": ListCopyObjectRequestsProjectionSummary,
	"DETAILS": ListCopyObjectRequestsProjectionDetails,
}

var mappingListCopyObjectRequestsProjectionEnumLowerCase = map[string]ListCopyObjectRequestsProjectionEnum{
	"summary": ListCopyObjectRequestsProjectionSummary,
	"details": ListCopyObjectRequestsProjectionDetails,
}

// GetListCopyObjectRequestsProjectionEnumValues Enumerates the set of values for ListCopyObjectRequestsProjectionEnum
func GetListCopyObjectRequestsProjectionEnumValues() []ListCopyObjectRequestsProjectionEnum {
	values := make([]ListCopyObjectRequestsProjectionEnum, 0)
	for _, v := range mappingListCopyObjectRequestsProjectionEnum {
		values = append(values, v)
	}
	return values
}

// GetListCopyObjectRequestsProjectionEnumStringValues Enumerates the set of values in String for ListCopyObjectRequestsProjectionEnum
func GetListCopyObjectRequestsProjectionEnumStringValues() []string {
	return []string{
		"SUMMARY",
		"DETAILS",
	}
}

// GetMappingListCopyObjectRequestsProjectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCopyObjectRequestsProjectionEnum(val string) (ListCopyObjectRequestsProjectionEnum, bool) {
	enum, ok := mappingListCopyObjectRequestsProjectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
