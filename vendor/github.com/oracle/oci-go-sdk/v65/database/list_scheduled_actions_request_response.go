// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListScheduledActionsRequest wrapper for the ListScheduledActions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListScheduledActions.go.html to see an example of how to use ListScheduledActionsRequest.
type ListScheduledActionsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListScheduledActionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given service type exactly.
	ServiceType *string `mandatory:"false" contributesTo:"query" name:"serviceType"`

	// A filter to return only resources that match the given scheduling policy id exactly.
	SchedulingPlanId *string `mandatory:"false" contributesTo:"query" name:"schedulingPlanId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.
	SortBy ListScheduledActionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the given Scheduled Action id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState ScheduledActionSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListScheduledActionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListScheduledActionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListScheduledActionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListScheduledActionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListScheduledActionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListScheduledActionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListScheduledActionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledActionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListScheduledActionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledActionSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetScheduledActionSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListScheduledActionsResponse wrapper for the ListScheduledActions operation
type ListScheduledActionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ScheduledActionCollection instances
	ScheduledActionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListScheduledActionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListScheduledActionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListScheduledActionsSortOrderEnum Enum with underlying type: string
type ListScheduledActionsSortOrderEnum string

// Set of constants representing the allowable values for ListScheduledActionsSortOrderEnum
const (
	ListScheduledActionsSortOrderAsc  ListScheduledActionsSortOrderEnum = "ASC"
	ListScheduledActionsSortOrderDesc ListScheduledActionsSortOrderEnum = "DESC"
)

var mappingListScheduledActionsSortOrderEnum = map[string]ListScheduledActionsSortOrderEnum{
	"ASC":  ListScheduledActionsSortOrderAsc,
	"DESC": ListScheduledActionsSortOrderDesc,
}

var mappingListScheduledActionsSortOrderEnumLowerCase = map[string]ListScheduledActionsSortOrderEnum{
	"asc":  ListScheduledActionsSortOrderAsc,
	"desc": ListScheduledActionsSortOrderDesc,
}

// GetListScheduledActionsSortOrderEnumValues Enumerates the set of values for ListScheduledActionsSortOrderEnum
func GetListScheduledActionsSortOrderEnumValues() []ListScheduledActionsSortOrderEnum {
	values := make([]ListScheduledActionsSortOrderEnum, 0)
	for _, v := range mappingListScheduledActionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledActionsSortOrderEnumStringValues Enumerates the set of values in String for ListScheduledActionsSortOrderEnum
func GetListScheduledActionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListScheduledActionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledActionsSortOrderEnum(val string) (ListScheduledActionsSortOrderEnum, bool) {
	enum, ok := mappingListScheduledActionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListScheduledActionsSortByEnum Enum with underlying type: string
type ListScheduledActionsSortByEnum string

// Set of constants representing the allowable values for ListScheduledActionsSortByEnum
const (
	ListScheduledActionsSortByTimecreated ListScheduledActionsSortByEnum = "TIMECREATED"
	ListScheduledActionsSortByDisplayname ListScheduledActionsSortByEnum = "DISPLAYNAME"
)

var mappingListScheduledActionsSortByEnum = map[string]ListScheduledActionsSortByEnum{
	"TIMECREATED": ListScheduledActionsSortByTimecreated,
	"DISPLAYNAME": ListScheduledActionsSortByDisplayname,
}

var mappingListScheduledActionsSortByEnumLowerCase = map[string]ListScheduledActionsSortByEnum{
	"timecreated": ListScheduledActionsSortByTimecreated,
	"displayname": ListScheduledActionsSortByDisplayname,
}

// GetListScheduledActionsSortByEnumValues Enumerates the set of values for ListScheduledActionsSortByEnum
func GetListScheduledActionsSortByEnumValues() []ListScheduledActionsSortByEnum {
	values := make([]ListScheduledActionsSortByEnum, 0)
	for _, v := range mappingListScheduledActionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledActionsSortByEnumStringValues Enumerates the set of values in String for ListScheduledActionsSortByEnum
func GetListScheduledActionsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListScheduledActionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledActionsSortByEnum(val string) (ListScheduledActionsSortByEnum, bool) {
	enum, ok := mappingListScheduledActionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
