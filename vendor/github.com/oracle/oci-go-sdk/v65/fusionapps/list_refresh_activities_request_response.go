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

// ListRefreshActivitiesRequest wrapper for the ListRefreshActivities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListRefreshActivities.go.html to see an example of how to use ListRefreshActivitiesRequest.
type ListRefreshActivitiesRequest struct {

	// unique FusionEnvironment identifier
	FusionEnvironmentId *string `mandatory:"true" contributesTo:"path" name:"fusionEnvironmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter that returns all resources that are scheduled after this date
	TimeScheduledStartGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduledStartGreaterThanOrEqualTo"`

	// A filter that returns all resources that end before this date
	TimeExpectedFinishLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeExpectedFinishLessThanOrEqualTo"`

	// A filter that returns all resources that match the specified status
	LifecycleState RefreshActivityLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListRefreshActivitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListRefreshActivitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRefreshActivitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRefreshActivitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRefreshActivitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRefreshActivitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRefreshActivitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRefreshActivityLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRefreshActivityLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRefreshActivitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRefreshActivitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRefreshActivitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRefreshActivitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRefreshActivitiesResponse wrapper for the ListRefreshActivities operation
type ListRefreshActivitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RefreshActivityCollection instances
	RefreshActivityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRefreshActivitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRefreshActivitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRefreshActivitiesSortOrderEnum Enum with underlying type: string
type ListRefreshActivitiesSortOrderEnum string

// Set of constants representing the allowable values for ListRefreshActivitiesSortOrderEnum
const (
	ListRefreshActivitiesSortOrderAsc  ListRefreshActivitiesSortOrderEnum = "ASC"
	ListRefreshActivitiesSortOrderDesc ListRefreshActivitiesSortOrderEnum = "DESC"
)

var mappingListRefreshActivitiesSortOrderEnum = map[string]ListRefreshActivitiesSortOrderEnum{
	"ASC":  ListRefreshActivitiesSortOrderAsc,
	"DESC": ListRefreshActivitiesSortOrderDesc,
}

var mappingListRefreshActivitiesSortOrderEnumLowerCase = map[string]ListRefreshActivitiesSortOrderEnum{
	"asc":  ListRefreshActivitiesSortOrderAsc,
	"desc": ListRefreshActivitiesSortOrderDesc,
}

// GetListRefreshActivitiesSortOrderEnumValues Enumerates the set of values for ListRefreshActivitiesSortOrderEnum
func GetListRefreshActivitiesSortOrderEnumValues() []ListRefreshActivitiesSortOrderEnum {
	values := make([]ListRefreshActivitiesSortOrderEnum, 0)
	for _, v := range mappingListRefreshActivitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRefreshActivitiesSortOrderEnumStringValues Enumerates the set of values in String for ListRefreshActivitiesSortOrderEnum
func GetListRefreshActivitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRefreshActivitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRefreshActivitiesSortOrderEnum(val string) (ListRefreshActivitiesSortOrderEnum, bool) {
	enum, ok := mappingListRefreshActivitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRefreshActivitiesSortByEnum Enum with underlying type: string
type ListRefreshActivitiesSortByEnum string

// Set of constants representing the allowable values for ListRefreshActivitiesSortByEnum
const (
	ListRefreshActivitiesSortByTimeCreated ListRefreshActivitiesSortByEnum = "TIME_CREATED"
	ListRefreshActivitiesSortByDisplayName ListRefreshActivitiesSortByEnum = "DISPLAY_NAME"
)

var mappingListRefreshActivitiesSortByEnum = map[string]ListRefreshActivitiesSortByEnum{
	"TIME_CREATED": ListRefreshActivitiesSortByTimeCreated,
	"DISPLAY_NAME": ListRefreshActivitiesSortByDisplayName,
}

var mappingListRefreshActivitiesSortByEnumLowerCase = map[string]ListRefreshActivitiesSortByEnum{
	"time_created": ListRefreshActivitiesSortByTimeCreated,
	"display_name": ListRefreshActivitiesSortByDisplayName,
}

// GetListRefreshActivitiesSortByEnumValues Enumerates the set of values for ListRefreshActivitiesSortByEnum
func GetListRefreshActivitiesSortByEnumValues() []ListRefreshActivitiesSortByEnum {
	values := make([]ListRefreshActivitiesSortByEnum, 0)
	for _, v := range mappingListRefreshActivitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRefreshActivitiesSortByEnumStringValues Enumerates the set of values in String for ListRefreshActivitiesSortByEnum
func GetListRefreshActivitiesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListRefreshActivitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRefreshActivitiesSortByEnum(val string) (ListRefreshActivitiesSortByEnum, bool) {
	enum, ok := mappingListRefreshActivitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
