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

// ListDataMaskingActivitiesRequest wrapper for the ListDataMaskingActivities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListDataMaskingActivities.go.html to see an example of how to use ListDataMaskingActivitiesRequest.
type ListDataMaskingActivitiesRequest struct {

	// unique FusionEnvironment identifier
	FusionEnvironmentId *string `mandatory:"true" contributesTo:"path" name:"fusionEnvironmentId"`

	// A filter that returns all resources that match the specified status
	LifecycleState DataMaskingActivityLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataMaskingActivitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDataMaskingActivitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataMaskingActivitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataMaskingActivitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataMaskingActivitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataMaskingActivitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataMaskingActivitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataMaskingActivityLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDataMaskingActivityLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataMaskingActivitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataMaskingActivitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataMaskingActivitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataMaskingActivitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataMaskingActivitiesResponse wrapper for the ListDataMaskingActivities operation
type ListDataMaskingActivitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataMaskingActivityCollection instances
	DataMaskingActivityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataMaskingActivitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataMaskingActivitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataMaskingActivitiesSortOrderEnum Enum with underlying type: string
type ListDataMaskingActivitiesSortOrderEnum string

// Set of constants representing the allowable values for ListDataMaskingActivitiesSortOrderEnum
const (
	ListDataMaskingActivitiesSortOrderAsc  ListDataMaskingActivitiesSortOrderEnum = "ASC"
	ListDataMaskingActivitiesSortOrderDesc ListDataMaskingActivitiesSortOrderEnum = "DESC"
)

var mappingListDataMaskingActivitiesSortOrderEnum = map[string]ListDataMaskingActivitiesSortOrderEnum{
	"ASC":  ListDataMaskingActivitiesSortOrderAsc,
	"DESC": ListDataMaskingActivitiesSortOrderDesc,
}

var mappingListDataMaskingActivitiesSortOrderEnumLowerCase = map[string]ListDataMaskingActivitiesSortOrderEnum{
	"asc":  ListDataMaskingActivitiesSortOrderAsc,
	"desc": ListDataMaskingActivitiesSortOrderDesc,
}

// GetListDataMaskingActivitiesSortOrderEnumValues Enumerates the set of values for ListDataMaskingActivitiesSortOrderEnum
func GetListDataMaskingActivitiesSortOrderEnumValues() []ListDataMaskingActivitiesSortOrderEnum {
	values := make([]ListDataMaskingActivitiesSortOrderEnum, 0)
	for _, v := range mappingListDataMaskingActivitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataMaskingActivitiesSortOrderEnumStringValues Enumerates the set of values in String for ListDataMaskingActivitiesSortOrderEnum
func GetListDataMaskingActivitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataMaskingActivitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataMaskingActivitiesSortOrderEnum(val string) (ListDataMaskingActivitiesSortOrderEnum, bool) {
	enum, ok := mappingListDataMaskingActivitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataMaskingActivitiesSortByEnum Enum with underlying type: string
type ListDataMaskingActivitiesSortByEnum string

// Set of constants representing the allowable values for ListDataMaskingActivitiesSortByEnum
const (
	ListDataMaskingActivitiesSortByTimeCreated ListDataMaskingActivitiesSortByEnum = "TIME_CREATED"
	ListDataMaskingActivitiesSortByDisplayName ListDataMaskingActivitiesSortByEnum = "DISPLAY_NAME"
)

var mappingListDataMaskingActivitiesSortByEnum = map[string]ListDataMaskingActivitiesSortByEnum{
	"TIME_CREATED": ListDataMaskingActivitiesSortByTimeCreated,
	"DISPLAY_NAME": ListDataMaskingActivitiesSortByDisplayName,
}

var mappingListDataMaskingActivitiesSortByEnumLowerCase = map[string]ListDataMaskingActivitiesSortByEnum{
	"time_created": ListDataMaskingActivitiesSortByTimeCreated,
	"display_name": ListDataMaskingActivitiesSortByDisplayName,
}

// GetListDataMaskingActivitiesSortByEnumValues Enumerates the set of values for ListDataMaskingActivitiesSortByEnum
func GetListDataMaskingActivitiesSortByEnumValues() []ListDataMaskingActivitiesSortByEnum {
	values := make([]ListDataMaskingActivitiesSortByEnum, 0)
	for _, v := range mappingListDataMaskingActivitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataMaskingActivitiesSortByEnumStringValues Enumerates the set of values in String for ListDataMaskingActivitiesSortByEnum
func GetListDataMaskingActivitiesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListDataMaskingActivitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataMaskingActivitiesSortByEnum(val string) (ListDataMaskingActivitiesSortByEnum, bool) {
	enum, ok := mappingListDataMaskingActivitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
