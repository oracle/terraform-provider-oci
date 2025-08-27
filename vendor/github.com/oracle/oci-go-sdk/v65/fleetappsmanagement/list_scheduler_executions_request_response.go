// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSchedulerExecutionsRequest wrapper for the ListSchedulerExecutions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListSchedulerExecutions.go.html to see an example of how to use ListSchedulerExecutionsRequest.
type ListSchedulerExecutionsRequest struct {

	// The ID of the compartment in which to list resources.
	// Empty only if the resource OCID query param is not specified.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Scheduled Time
	TimeScheduledGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduledGreaterThanOrEqualTo"`

	// Scheduled Time
	TimeScheduledLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduledLessThan"`

	// SchedulerDefinition identifier
	SchedulerDefintionId *string `mandatory:"false" contributesTo:"query" name:"schedulerDefintionId"`

	// SchedulerJob identifier filter
	SchedulerJobId *string `mandatory:"false" contributesTo:"query" name:"schedulerJobId"`

	// ResourceId filter (Example FleetId)
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// A filter to return only schedule definitions whose associated runbookId matches the given runbookId.
	RunbookId *string `mandatory:"false" contributesTo:"query" name:"runbookId"`

	// RunbookVersion Name filter
	RunbookVersionName *string `mandatory:"false" contributesTo:"query" name:"runbookVersionName"`

	// A filter to return only resources their subState matches the given subState.
	Substate *string `mandatory:"false" contributesTo:"query" name:"substate"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSchedulerExecutionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.  Default order for timeCreated and timeScheduled is descending.
	SortBy ListSchedulerExecutionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSchedulerExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSchedulerExecutionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSchedulerExecutionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSchedulerExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSchedulerExecutionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSchedulerExecutionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSchedulerExecutionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSchedulerExecutionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSchedulerExecutionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSchedulerExecutionsResponse wrapper for the ListSchedulerExecutions operation
type ListSchedulerExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SchedulerExecutionCollection instances
	SchedulerExecutionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSchedulerExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSchedulerExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSchedulerExecutionsSortOrderEnum Enum with underlying type: string
type ListSchedulerExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListSchedulerExecutionsSortOrderEnum
const (
	ListSchedulerExecutionsSortOrderAsc  ListSchedulerExecutionsSortOrderEnum = "ASC"
	ListSchedulerExecutionsSortOrderDesc ListSchedulerExecutionsSortOrderEnum = "DESC"
)

var mappingListSchedulerExecutionsSortOrderEnum = map[string]ListSchedulerExecutionsSortOrderEnum{
	"ASC":  ListSchedulerExecutionsSortOrderAsc,
	"DESC": ListSchedulerExecutionsSortOrderDesc,
}

var mappingListSchedulerExecutionsSortOrderEnumLowerCase = map[string]ListSchedulerExecutionsSortOrderEnum{
	"asc":  ListSchedulerExecutionsSortOrderAsc,
	"desc": ListSchedulerExecutionsSortOrderDesc,
}

// GetListSchedulerExecutionsSortOrderEnumValues Enumerates the set of values for ListSchedulerExecutionsSortOrderEnum
func GetListSchedulerExecutionsSortOrderEnumValues() []ListSchedulerExecutionsSortOrderEnum {
	values := make([]ListSchedulerExecutionsSortOrderEnum, 0)
	for _, v := range mappingListSchedulerExecutionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchedulerExecutionsSortOrderEnumStringValues Enumerates the set of values in String for ListSchedulerExecutionsSortOrderEnum
func GetListSchedulerExecutionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSchedulerExecutionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchedulerExecutionsSortOrderEnum(val string) (ListSchedulerExecutionsSortOrderEnum, bool) {
	enum, ok := mappingListSchedulerExecutionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSchedulerExecutionsSortByEnum Enum with underlying type: string
type ListSchedulerExecutionsSortByEnum string

// Set of constants representing the allowable values for ListSchedulerExecutionsSortByEnum
const (
	ListSchedulerExecutionsSortByTimecreated   ListSchedulerExecutionsSortByEnum = "timeCreated"
	ListSchedulerExecutionsSortByTimescheduled ListSchedulerExecutionsSortByEnum = "timeScheduled"
)

var mappingListSchedulerExecutionsSortByEnum = map[string]ListSchedulerExecutionsSortByEnum{
	"timeCreated":   ListSchedulerExecutionsSortByTimecreated,
	"timeScheduled": ListSchedulerExecutionsSortByTimescheduled,
}

var mappingListSchedulerExecutionsSortByEnumLowerCase = map[string]ListSchedulerExecutionsSortByEnum{
	"timecreated":   ListSchedulerExecutionsSortByTimecreated,
	"timescheduled": ListSchedulerExecutionsSortByTimescheduled,
}

// GetListSchedulerExecutionsSortByEnumValues Enumerates the set of values for ListSchedulerExecutionsSortByEnum
func GetListSchedulerExecutionsSortByEnumValues() []ListSchedulerExecutionsSortByEnum {
	values := make([]ListSchedulerExecutionsSortByEnum, 0)
	for _, v := range mappingListSchedulerExecutionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchedulerExecutionsSortByEnumStringValues Enumerates the set of values in String for ListSchedulerExecutionsSortByEnum
func GetListSchedulerExecutionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeScheduled",
	}
}

// GetMappingListSchedulerExecutionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchedulerExecutionsSortByEnum(val string) (ListSchedulerExecutionsSortByEnum, bool) {
	enum, ok := mappingListSchedulerExecutionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
