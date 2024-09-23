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

// ListSchedulerJobsRequest wrapper for the ListSchedulerJobs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListSchedulerJobs.go.html to see an example of how to use ListSchedulerJobsRequest.
type ListSchedulerJobsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState SchedulerJobLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// unique Fleet identifier
	FleetId *string `mandatory:"false" contributesTo:"query" name:"fleetId"`

	// Scheduled Time
	TimeScheduledGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduledGreaterThanOrEqualTo"`

	// Scheduled Time
	TimeScheduledLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduledLessThan"`

	// Fetch next remediation Job
	IsRemediationJobNeeded *bool `mandatory:"false" contributesTo:"query" name:"isRemediationJobNeeded"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique SchedulerJob identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// SchedulerJob Definition identifier
	DefintionId *string `mandatory:"false" contributesTo:"query" name:"defintionId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSchedulerJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated and timeScheduled is descending. Default order for displayName is ascending.
	SortBy ListSchedulerJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSchedulerJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSchedulerJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSchedulerJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSchedulerJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSchedulerJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSchedulerJobLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetSchedulerJobLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSchedulerJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSchedulerJobsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSchedulerJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSchedulerJobsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSchedulerJobsResponse wrapper for the ListSchedulerJobs operation
type ListSchedulerJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SchedulerJobCollection instances
	SchedulerJobCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSchedulerJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSchedulerJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSchedulerJobsSortOrderEnum Enum with underlying type: string
type ListSchedulerJobsSortOrderEnum string

// Set of constants representing the allowable values for ListSchedulerJobsSortOrderEnum
const (
	ListSchedulerJobsSortOrderAsc  ListSchedulerJobsSortOrderEnum = "ASC"
	ListSchedulerJobsSortOrderDesc ListSchedulerJobsSortOrderEnum = "DESC"
)

var mappingListSchedulerJobsSortOrderEnum = map[string]ListSchedulerJobsSortOrderEnum{
	"ASC":  ListSchedulerJobsSortOrderAsc,
	"DESC": ListSchedulerJobsSortOrderDesc,
}

var mappingListSchedulerJobsSortOrderEnumLowerCase = map[string]ListSchedulerJobsSortOrderEnum{
	"asc":  ListSchedulerJobsSortOrderAsc,
	"desc": ListSchedulerJobsSortOrderDesc,
}

// GetListSchedulerJobsSortOrderEnumValues Enumerates the set of values for ListSchedulerJobsSortOrderEnum
func GetListSchedulerJobsSortOrderEnumValues() []ListSchedulerJobsSortOrderEnum {
	values := make([]ListSchedulerJobsSortOrderEnum, 0)
	for _, v := range mappingListSchedulerJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchedulerJobsSortOrderEnumStringValues Enumerates the set of values in String for ListSchedulerJobsSortOrderEnum
func GetListSchedulerJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSchedulerJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchedulerJobsSortOrderEnum(val string) (ListSchedulerJobsSortOrderEnum, bool) {
	enum, ok := mappingListSchedulerJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSchedulerJobsSortByEnum Enum with underlying type: string
type ListSchedulerJobsSortByEnum string

// Set of constants representing the allowable values for ListSchedulerJobsSortByEnum
const (
	ListSchedulerJobsSortByTimecreated   ListSchedulerJobsSortByEnum = "timeCreated"
	ListSchedulerJobsSortByTimescheduled ListSchedulerJobsSortByEnum = "timeScheduled"
	ListSchedulerJobsSortByDisplayname   ListSchedulerJobsSortByEnum = "displayName"
)

var mappingListSchedulerJobsSortByEnum = map[string]ListSchedulerJobsSortByEnum{
	"timeCreated":   ListSchedulerJobsSortByTimecreated,
	"timeScheduled": ListSchedulerJobsSortByTimescheduled,
	"displayName":   ListSchedulerJobsSortByDisplayname,
}

var mappingListSchedulerJobsSortByEnumLowerCase = map[string]ListSchedulerJobsSortByEnum{
	"timecreated":   ListSchedulerJobsSortByTimecreated,
	"timescheduled": ListSchedulerJobsSortByTimescheduled,
	"displayname":   ListSchedulerJobsSortByDisplayname,
}

// GetListSchedulerJobsSortByEnumValues Enumerates the set of values for ListSchedulerJobsSortByEnum
func GetListSchedulerJobsSortByEnumValues() []ListSchedulerJobsSortByEnum {
	values := make([]ListSchedulerJobsSortByEnum, 0)
	for _, v := range mappingListSchedulerJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchedulerJobsSortByEnumStringValues Enumerates the set of values in String for ListSchedulerJobsSortByEnum
func GetListSchedulerJobsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeScheduled",
		"displayName",
	}
}

// GetMappingListSchedulerJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchedulerJobsSortByEnum(val string) (ListSchedulerJobsSortByEnum, bool) {
	enum, ok := mappingListSchedulerJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
