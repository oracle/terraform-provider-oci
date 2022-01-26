// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListSchedulesRequest wrapper for the ListSchedules operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListSchedules.go.html to see an example of how to use ListSchedulesRequest.
type ListSchedulesRequest struct {

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

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListSchedulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListSchedulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSchedulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSchedulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSchedulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSchedulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSchedulesResponse wrapper for the ListSchedules operation
type ListSchedulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ScheduleSummaryCollection instances
	ScheduleSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListSchedulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSchedulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSchedulesSortByEnum Enum with underlying type: string
type ListSchedulesSortByEnum string

// Set of constants representing the allowable values for ListSchedulesSortByEnum
const (
	ListSchedulesSortByTimeCreated ListSchedulesSortByEnum = "TIME_CREATED"
	ListSchedulesSortByDisplayName ListSchedulesSortByEnum = "DISPLAY_NAME"
)

var mappingListSchedulesSortBy = map[string]ListSchedulesSortByEnum{
	"TIME_CREATED": ListSchedulesSortByTimeCreated,
	"DISPLAY_NAME": ListSchedulesSortByDisplayName,
}

// GetListSchedulesSortByEnumValues Enumerates the set of values for ListSchedulesSortByEnum
func GetListSchedulesSortByEnumValues() []ListSchedulesSortByEnum {
	values := make([]ListSchedulesSortByEnum, 0)
	for _, v := range mappingListSchedulesSortBy {
		values = append(values, v)
	}
	return values
}

// ListSchedulesSortOrderEnum Enum with underlying type: string
type ListSchedulesSortOrderEnum string

// Set of constants representing the allowable values for ListSchedulesSortOrderEnum
const (
	ListSchedulesSortOrderAsc  ListSchedulesSortOrderEnum = "ASC"
	ListSchedulesSortOrderDesc ListSchedulesSortOrderEnum = "DESC"
)

var mappingListSchedulesSortOrder = map[string]ListSchedulesSortOrderEnum{
	"ASC":  ListSchedulesSortOrderAsc,
	"DESC": ListSchedulesSortOrderDesc,
}

// GetListSchedulesSortOrderEnumValues Enumerates the set of values for ListSchedulesSortOrderEnum
func GetListSchedulesSortOrderEnumValues() []ListSchedulesSortOrderEnum {
	values := make([]ListSchedulesSortOrderEnum, 0)
	for _, v := range mappingListSchedulesSortOrder {
		values = append(values, v)
	}
	return values
}
