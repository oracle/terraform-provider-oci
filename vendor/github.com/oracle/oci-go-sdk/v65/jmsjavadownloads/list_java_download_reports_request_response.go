// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jmsjavadownloads

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJavaDownloadReportsRequest wrapper for the ListJavaDownloadReports operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListJavaDownloadReports.go.html to see an example of how to use ListJavaDownloadReportsRequest.
type ListJavaDownloadReportsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenancy.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListJavaDownloadReportsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Java download report identifier.
	JavaDownloadReportId *string `mandatory:"false" contributesTo:"query" name:"javaDownloadReportId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListJavaDownloadReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. If no value is specified, _timeCreated_ is the default.
	SortBy ListJavaDownloadReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJavaDownloadReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJavaDownloadReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJavaDownloadReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJavaDownloadReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJavaDownloadReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJavaDownloadReportsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListJavaDownloadReportsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaDownloadReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJavaDownloadReportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaDownloadReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJavaDownloadReportsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJavaDownloadReportsResponse wrapper for the ListJavaDownloadReports operation
type ListJavaDownloadReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JavaDownloadReportCollection instances
	JavaDownloadReportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJavaDownloadReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJavaDownloadReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJavaDownloadReportsLifecycleStateEnum Enum with underlying type: string
type ListJavaDownloadReportsLifecycleStateEnum string

// Set of constants representing the allowable values for ListJavaDownloadReportsLifecycleStateEnum
const (
	ListJavaDownloadReportsLifecycleStateActive         ListJavaDownloadReportsLifecycleStateEnum = "ACTIVE"
	ListJavaDownloadReportsLifecycleStateCreating       ListJavaDownloadReportsLifecycleStateEnum = "CREATING"
	ListJavaDownloadReportsLifecycleStateDeleted        ListJavaDownloadReportsLifecycleStateEnum = "DELETED"
	ListJavaDownloadReportsLifecycleStateDeleting       ListJavaDownloadReportsLifecycleStateEnum = "DELETING"
	ListJavaDownloadReportsLifecycleStateFailed         ListJavaDownloadReportsLifecycleStateEnum = "FAILED"
	ListJavaDownloadReportsLifecycleStateNeedsAttention ListJavaDownloadReportsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListJavaDownloadReportsLifecycleStateUpdating       ListJavaDownloadReportsLifecycleStateEnum = "UPDATING"
)

var mappingListJavaDownloadReportsLifecycleStateEnum = map[string]ListJavaDownloadReportsLifecycleStateEnum{
	"ACTIVE":          ListJavaDownloadReportsLifecycleStateActive,
	"CREATING":        ListJavaDownloadReportsLifecycleStateCreating,
	"DELETED":         ListJavaDownloadReportsLifecycleStateDeleted,
	"DELETING":        ListJavaDownloadReportsLifecycleStateDeleting,
	"FAILED":          ListJavaDownloadReportsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListJavaDownloadReportsLifecycleStateNeedsAttention,
	"UPDATING":        ListJavaDownloadReportsLifecycleStateUpdating,
}

var mappingListJavaDownloadReportsLifecycleStateEnumLowerCase = map[string]ListJavaDownloadReportsLifecycleStateEnum{
	"active":          ListJavaDownloadReportsLifecycleStateActive,
	"creating":        ListJavaDownloadReportsLifecycleStateCreating,
	"deleted":         ListJavaDownloadReportsLifecycleStateDeleted,
	"deleting":        ListJavaDownloadReportsLifecycleStateDeleting,
	"failed":          ListJavaDownloadReportsLifecycleStateFailed,
	"needs_attention": ListJavaDownloadReportsLifecycleStateNeedsAttention,
	"updating":        ListJavaDownloadReportsLifecycleStateUpdating,
}

// GetListJavaDownloadReportsLifecycleStateEnumValues Enumerates the set of values for ListJavaDownloadReportsLifecycleStateEnum
func GetListJavaDownloadReportsLifecycleStateEnumValues() []ListJavaDownloadReportsLifecycleStateEnum {
	values := make([]ListJavaDownloadReportsLifecycleStateEnum, 0)
	for _, v := range mappingListJavaDownloadReportsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaDownloadReportsLifecycleStateEnumStringValues Enumerates the set of values in String for ListJavaDownloadReportsLifecycleStateEnum
func GetListJavaDownloadReportsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"NEEDS_ATTENTION",
		"UPDATING",
	}
}

// GetMappingListJavaDownloadReportsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaDownloadReportsLifecycleStateEnum(val string) (ListJavaDownloadReportsLifecycleStateEnum, bool) {
	enum, ok := mappingListJavaDownloadReportsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaDownloadReportsSortOrderEnum Enum with underlying type: string
type ListJavaDownloadReportsSortOrderEnum string

// Set of constants representing the allowable values for ListJavaDownloadReportsSortOrderEnum
const (
	ListJavaDownloadReportsSortOrderAsc  ListJavaDownloadReportsSortOrderEnum = "ASC"
	ListJavaDownloadReportsSortOrderDesc ListJavaDownloadReportsSortOrderEnum = "DESC"
)

var mappingListJavaDownloadReportsSortOrderEnum = map[string]ListJavaDownloadReportsSortOrderEnum{
	"ASC":  ListJavaDownloadReportsSortOrderAsc,
	"DESC": ListJavaDownloadReportsSortOrderDesc,
}

var mappingListJavaDownloadReportsSortOrderEnumLowerCase = map[string]ListJavaDownloadReportsSortOrderEnum{
	"asc":  ListJavaDownloadReportsSortOrderAsc,
	"desc": ListJavaDownloadReportsSortOrderDesc,
}

// GetListJavaDownloadReportsSortOrderEnumValues Enumerates the set of values for ListJavaDownloadReportsSortOrderEnum
func GetListJavaDownloadReportsSortOrderEnumValues() []ListJavaDownloadReportsSortOrderEnum {
	values := make([]ListJavaDownloadReportsSortOrderEnum, 0)
	for _, v := range mappingListJavaDownloadReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaDownloadReportsSortOrderEnumStringValues Enumerates the set of values in String for ListJavaDownloadReportsSortOrderEnum
func GetListJavaDownloadReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJavaDownloadReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaDownloadReportsSortOrderEnum(val string) (ListJavaDownloadReportsSortOrderEnum, bool) {
	enum, ok := mappingListJavaDownloadReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaDownloadReportsSortByEnum Enum with underlying type: string
type ListJavaDownloadReportsSortByEnum string

// Set of constants representing the allowable values for ListJavaDownloadReportsSortByEnum
const (
	ListJavaDownloadReportsSortByTimecreated ListJavaDownloadReportsSortByEnum = "timeCreated"
	ListJavaDownloadReportsSortByDisplayname ListJavaDownloadReportsSortByEnum = "displayName"
)

var mappingListJavaDownloadReportsSortByEnum = map[string]ListJavaDownloadReportsSortByEnum{
	"timeCreated": ListJavaDownloadReportsSortByTimecreated,
	"displayName": ListJavaDownloadReportsSortByDisplayname,
}

var mappingListJavaDownloadReportsSortByEnumLowerCase = map[string]ListJavaDownloadReportsSortByEnum{
	"timecreated": ListJavaDownloadReportsSortByTimecreated,
	"displayname": ListJavaDownloadReportsSortByDisplayname,
}

// GetListJavaDownloadReportsSortByEnumValues Enumerates the set of values for ListJavaDownloadReportsSortByEnum
func GetListJavaDownloadReportsSortByEnumValues() []ListJavaDownloadReportsSortByEnum {
	values := make([]ListJavaDownloadReportsSortByEnum, 0)
	for _, v := range mappingListJavaDownloadReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaDownloadReportsSortByEnumStringValues Enumerates the set of values in String for ListJavaDownloadReportsSortByEnum
func GetListJavaDownloadReportsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListJavaDownloadReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaDownloadReportsSortByEnum(val string) (ListJavaDownloadReportsSortByEnum, bool) {
	enum, ok := mappingListJavaDownloadReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
