// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTaskSchedulesRequest wrapper for the ListTaskSchedules operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListTaskSchedules.go.html to see an example of how to use ListTaskSchedulesRequest.
type ListTaskSchedulesRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The application key.
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// Used to filter by the key of the object.
	Key []string `contributesTo:"query" name:"key" collectionFormat:"multi"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Used to filter by the identifier of the object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// Used to filter by the object type of the object. It can be suffixed with an optional filter operator InSubtree. If this operator is not specified, then exact match is considered. <br><br><B>Examples:</B><br> <ul> <li><B>?type=DATA_LOADER_TASK&typeInSubtree=false</B> returns all objects of type data loader task</li> <li><B>?type=DATA_LOADER_TASK</B> returns all objects of type data loader task</li> <li><B>?type=DATA_LOADER_TASK&typeInSubtree=true</B> returns all objects of type data loader task</li> </ul>
	Type []string `contributesTo:"query" name:"type" collectionFormat:"multi"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListTaskSchedulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListTaskSchedulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This filter parameter can be used to filter task schedule by its state.
	IsEnabled *bool `mandatory:"false" contributesTo:"query" name:"isEnabled"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTaskSchedulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTaskSchedulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTaskSchedulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTaskSchedulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTaskSchedulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTaskSchedulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTaskSchedulesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTaskSchedulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTaskSchedulesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTaskSchedulesResponse wrapper for the ListTaskSchedules operation
type ListTaskSchedulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TaskScheduleSummaryCollection instances
	TaskScheduleSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListTaskSchedulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTaskSchedulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTaskSchedulesSortByEnum Enum with underlying type: string
type ListTaskSchedulesSortByEnum string

// Set of constants representing the allowable values for ListTaskSchedulesSortByEnum
const (
	ListTaskSchedulesSortByTimeCreated ListTaskSchedulesSortByEnum = "TIME_CREATED"
	ListTaskSchedulesSortByDisplayName ListTaskSchedulesSortByEnum = "DISPLAY_NAME"
	ListTaskSchedulesSortByTimeUpdated ListTaskSchedulesSortByEnum = "TIME_UPDATED"
)

var mappingListTaskSchedulesSortByEnum = map[string]ListTaskSchedulesSortByEnum{
	"TIME_CREATED": ListTaskSchedulesSortByTimeCreated,
	"DISPLAY_NAME": ListTaskSchedulesSortByDisplayName,
	"TIME_UPDATED": ListTaskSchedulesSortByTimeUpdated,
}

var mappingListTaskSchedulesSortByEnumLowerCase = map[string]ListTaskSchedulesSortByEnum{
	"time_created": ListTaskSchedulesSortByTimeCreated,
	"display_name": ListTaskSchedulesSortByDisplayName,
	"time_updated": ListTaskSchedulesSortByTimeUpdated,
}

// GetListTaskSchedulesSortByEnumValues Enumerates the set of values for ListTaskSchedulesSortByEnum
func GetListTaskSchedulesSortByEnumValues() []ListTaskSchedulesSortByEnum {
	values := make([]ListTaskSchedulesSortByEnum, 0)
	for _, v := range mappingListTaskSchedulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTaskSchedulesSortByEnumStringValues Enumerates the set of values in String for ListTaskSchedulesSortByEnum
func GetListTaskSchedulesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListTaskSchedulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTaskSchedulesSortByEnum(val string) (ListTaskSchedulesSortByEnum, bool) {
	enum, ok := mappingListTaskSchedulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTaskSchedulesSortOrderEnum Enum with underlying type: string
type ListTaskSchedulesSortOrderEnum string

// Set of constants representing the allowable values for ListTaskSchedulesSortOrderEnum
const (
	ListTaskSchedulesSortOrderAsc  ListTaskSchedulesSortOrderEnum = "ASC"
	ListTaskSchedulesSortOrderDesc ListTaskSchedulesSortOrderEnum = "DESC"
)

var mappingListTaskSchedulesSortOrderEnum = map[string]ListTaskSchedulesSortOrderEnum{
	"ASC":  ListTaskSchedulesSortOrderAsc,
	"DESC": ListTaskSchedulesSortOrderDesc,
}

var mappingListTaskSchedulesSortOrderEnumLowerCase = map[string]ListTaskSchedulesSortOrderEnum{
	"asc":  ListTaskSchedulesSortOrderAsc,
	"desc": ListTaskSchedulesSortOrderDesc,
}

// GetListTaskSchedulesSortOrderEnumValues Enumerates the set of values for ListTaskSchedulesSortOrderEnum
func GetListTaskSchedulesSortOrderEnumValues() []ListTaskSchedulesSortOrderEnum {
	values := make([]ListTaskSchedulesSortOrderEnum, 0)
	for _, v := range mappingListTaskSchedulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTaskSchedulesSortOrderEnumStringValues Enumerates the set of values in String for ListTaskSchedulesSortOrderEnum
func GetListTaskSchedulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTaskSchedulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTaskSchedulesSortOrderEnum(val string) (ListTaskSchedulesSortOrderEnum, bool) {
	enum, ok := mappingListTaskSchedulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
