// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListScheduledActivitiesRequest wrapper for the ListScheduledActivities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListScheduledActivities.go.html to see an example of how to use ListScheduledActivitiesRequest.
type ListScheduledActivitiesRequest struct {

	// unique FusionEnvironment identifier
	FusionEnvironmentId *string `mandatory:"true" contributesTo:"path" name:"fusionEnvironmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter that returns all resources that are scheduled after this date
	TimeScheduledStartGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduledStartGreaterThanOrEqualTo"`

	// A filter that returns all resources that end before this date
	TimeExpectedFinishLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeExpectedFinishLessThanOrEqualTo"`

	// A filter that returns all resources that match the specified run cycle.
	RunCycle ScheduledActivityRunCycleEnum `mandatory:"false" contributesTo:"query" name:"runCycle" omitEmpty:"true"`

	// A filter that returns all resources that match the specified status
	LifecycleState ScheduledActivityLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter that returns all resources that match the specified scheduledActivityAssociationId.
	ScheduledActivityAssociationId *string `mandatory:"false" contributesTo:"query" name:"scheduledActivityAssociationId"`

	// A filter that returns all resources that match the specified scheduledActivityPhase.
	ScheduledActivityPhase ScheduledActivityScheduledActivityPhaseEnum `mandatory:"false" contributesTo:"query" name:"scheduledActivityPhase" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListScheduledActivitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListScheduledActivitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListScheduledActivitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListScheduledActivitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListScheduledActivitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListScheduledActivitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListScheduledActivitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduledActivityRunCycleEnum(string(request.RunCycle)); !ok && request.RunCycle != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RunCycle: %s. Supported values are: %s.", request.RunCycle, strings.Join(GetScheduledActivityRunCycleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledActivityLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetScheduledActivityLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledActivityScheduledActivityPhaseEnum(string(request.ScheduledActivityPhase)); !ok && request.ScheduledActivityPhase != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduledActivityPhase: %s. Supported values are: %s.", request.ScheduledActivityPhase, strings.Join(GetScheduledActivityScheduledActivityPhaseEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledActivitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListScheduledActivitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledActivitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListScheduledActivitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListScheduledActivitiesResponse wrapper for the ListScheduledActivities operation
type ListScheduledActivitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ScheduledActivityCollection instances
	ScheduledActivityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListScheduledActivitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListScheduledActivitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListScheduledActivitiesSortOrderEnum Enum with underlying type: string
type ListScheduledActivitiesSortOrderEnum string

// Set of constants representing the allowable values for ListScheduledActivitiesSortOrderEnum
const (
	ListScheduledActivitiesSortOrderAsc  ListScheduledActivitiesSortOrderEnum = "ASC"
	ListScheduledActivitiesSortOrderDesc ListScheduledActivitiesSortOrderEnum = "DESC"
)

var mappingListScheduledActivitiesSortOrderEnum = map[string]ListScheduledActivitiesSortOrderEnum{
	"ASC":  ListScheduledActivitiesSortOrderAsc,
	"DESC": ListScheduledActivitiesSortOrderDesc,
}

var mappingListScheduledActivitiesSortOrderEnumLowerCase = map[string]ListScheduledActivitiesSortOrderEnum{
	"asc":  ListScheduledActivitiesSortOrderAsc,
	"desc": ListScheduledActivitiesSortOrderDesc,
}

// GetListScheduledActivitiesSortOrderEnumValues Enumerates the set of values for ListScheduledActivitiesSortOrderEnum
func GetListScheduledActivitiesSortOrderEnumValues() []ListScheduledActivitiesSortOrderEnum {
	values := make([]ListScheduledActivitiesSortOrderEnum, 0)
	for _, v := range mappingListScheduledActivitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledActivitiesSortOrderEnumStringValues Enumerates the set of values in String for ListScheduledActivitiesSortOrderEnum
func GetListScheduledActivitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListScheduledActivitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledActivitiesSortOrderEnum(val string) (ListScheduledActivitiesSortOrderEnum, bool) {
	enum, ok := mappingListScheduledActivitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListScheduledActivitiesSortByEnum Enum with underlying type: string
type ListScheduledActivitiesSortByEnum string

// Set of constants representing the allowable values for ListScheduledActivitiesSortByEnum
const (
	ListScheduledActivitiesSortByTimeCreated ListScheduledActivitiesSortByEnum = "TIME_CREATED"
	ListScheduledActivitiesSortByDisplayName ListScheduledActivitiesSortByEnum = "DISPLAY_NAME"
)

var mappingListScheduledActivitiesSortByEnum = map[string]ListScheduledActivitiesSortByEnum{
	"TIME_CREATED": ListScheduledActivitiesSortByTimeCreated,
	"DISPLAY_NAME": ListScheduledActivitiesSortByDisplayName,
}

var mappingListScheduledActivitiesSortByEnumLowerCase = map[string]ListScheduledActivitiesSortByEnum{
	"time_created": ListScheduledActivitiesSortByTimeCreated,
	"display_name": ListScheduledActivitiesSortByDisplayName,
}

// GetListScheduledActivitiesSortByEnumValues Enumerates the set of values for ListScheduledActivitiesSortByEnum
func GetListScheduledActivitiesSortByEnumValues() []ListScheduledActivitiesSortByEnum {
	values := make([]ListScheduledActivitiesSortByEnum, 0)
	for _, v := range mappingListScheduledActivitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledActivitiesSortByEnumStringValues Enumerates the set of values in String for ListScheduledActivitiesSortByEnum
func GetListScheduledActivitiesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListScheduledActivitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledActivitiesSortByEnum(val string) (ListScheduledActivitiesSortByEnum, bool) {
	enum, ok := mappingListScheduledActivitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
