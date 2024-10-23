// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAnnouncementsRequest wrapper for the ListAnnouncements operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListAnnouncements.go.html to see an example of how to use ListAnnouncementsRequest.
type ListAnnouncementsRequest struct {

	// Filter the list of announcements that contains the given summary value.
	SummaryContains *string `mandatory:"false" contributesTo:"query" name:"summaryContains"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAnnouncementsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort AnnouncementSummary by. Only one sort order may be provided.
	// If no value is specified timeReleased is default.
	SortBy ListAnnouncementsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAnnouncementsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAnnouncementsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAnnouncementsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAnnouncementsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAnnouncementsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAnnouncementsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAnnouncementsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAnnouncementsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAnnouncementsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAnnouncementsResponse wrapper for the ListAnnouncements operation
type ListAnnouncementsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AnnouncementCollection instances
	AnnouncementCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAnnouncementsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAnnouncementsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAnnouncementsSortOrderEnum Enum with underlying type: string
type ListAnnouncementsSortOrderEnum string

// Set of constants representing the allowable values for ListAnnouncementsSortOrderEnum
const (
	ListAnnouncementsSortOrderAsc  ListAnnouncementsSortOrderEnum = "ASC"
	ListAnnouncementsSortOrderDesc ListAnnouncementsSortOrderEnum = "DESC"
)

var mappingListAnnouncementsSortOrderEnum = map[string]ListAnnouncementsSortOrderEnum{
	"ASC":  ListAnnouncementsSortOrderAsc,
	"DESC": ListAnnouncementsSortOrderDesc,
}

var mappingListAnnouncementsSortOrderEnumLowerCase = map[string]ListAnnouncementsSortOrderEnum{
	"asc":  ListAnnouncementsSortOrderAsc,
	"desc": ListAnnouncementsSortOrderDesc,
}

// GetListAnnouncementsSortOrderEnumValues Enumerates the set of values for ListAnnouncementsSortOrderEnum
func GetListAnnouncementsSortOrderEnumValues() []ListAnnouncementsSortOrderEnum {
	values := make([]ListAnnouncementsSortOrderEnum, 0)
	for _, v := range mappingListAnnouncementsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnnouncementsSortOrderEnumStringValues Enumerates the set of values in String for ListAnnouncementsSortOrderEnum
func GetListAnnouncementsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAnnouncementsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnnouncementsSortOrderEnum(val string) (ListAnnouncementsSortOrderEnum, bool) {
	enum, ok := mappingListAnnouncementsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAnnouncementsSortByEnum Enum with underlying type: string
type ListAnnouncementsSortByEnum string

// Set of constants representing the allowable values for ListAnnouncementsSortByEnum
const (
	ListAnnouncementsSortByTimereleased ListAnnouncementsSortByEnum = "timeReleased"
	ListAnnouncementsSortBySummary      ListAnnouncementsSortByEnum = "summary"
)

var mappingListAnnouncementsSortByEnum = map[string]ListAnnouncementsSortByEnum{
	"timeReleased": ListAnnouncementsSortByTimereleased,
	"summary":      ListAnnouncementsSortBySummary,
}

var mappingListAnnouncementsSortByEnumLowerCase = map[string]ListAnnouncementsSortByEnum{
	"timereleased": ListAnnouncementsSortByTimereleased,
	"summary":      ListAnnouncementsSortBySummary,
}

// GetListAnnouncementsSortByEnumValues Enumerates the set of values for ListAnnouncementsSortByEnum
func GetListAnnouncementsSortByEnumValues() []ListAnnouncementsSortByEnum {
	values := make([]ListAnnouncementsSortByEnum, 0)
	for _, v := range mappingListAnnouncementsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnnouncementsSortByEnumStringValues Enumerates the set of values in String for ListAnnouncementsSortByEnum
func GetListAnnouncementsSortByEnumStringValues() []string {
	return []string{
		"timeReleased",
		"summary",
	}
}

// GetMappingListAnnouncementsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnnouncementsSortByEnum(val string) (ListAnnouncementsSortByEnum, bool) {
	enum, ok := mappingListAnnouncementsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
