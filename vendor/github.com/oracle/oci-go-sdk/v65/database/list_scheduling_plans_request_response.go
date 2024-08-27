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

// ListSchedulingPlansRequest wrapper for the ListSchedulingPlans operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListSchedulingPlans.go.html to see an example of how to use ListSchedulingPlansRequest.
type ListSchedulingPlansRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.
	SortBy ListSchedulingPlansSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSchedulingPlansSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState SchedulingPlanSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given scheduling policy id exactly.
	SchedulingPolicyId *string `mandatory:"false" contributesTo:"query" name:"schedulingPolicyId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given resource id exactly.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// A filter to return only resources that match the given Schedule Plan id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSchedulingPlansRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSchedulingPlansRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSchedulingPlansRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSchedulingPlansRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSchedulingPlansRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSchedulingPlansSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSchedulingPlansSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSchedulingPlansSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSchedulingPlansSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchedulingPlanSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetSchedulingPlanSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSchedulingPlansResponse wrapper for the ListSchedulingPlans operation
type ListSchedulingPlansResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SchedulingPlanCollection instances
	SchedulingPlanCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSchedulingPlansResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSchedulingPlansResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSchedulingPlansSortByEnum Enum with underlying type: string
type ListSchedulingPlansSortByEnum string

// Set of constants representing the allowable values for ListSchedulingPlansSortByEnum
const (
	ListSchedulingPlansSortByTimecreated ListSchedulingPlansSortByEnum = "TIMECREATED"
	ListSchedulingPlansSortByDisplayname ListSchedulingPlansSortByEnum = "DISPLAYNAME"
)

var mappingListSchedulingPlansSortByEnum = map[string]ListSchedulingPlansSortByEnum{
	"TIMECREATED": ListSchedulingPlansSortByTimecreated,
	"DISPLAYNAME": ListSchedulingPlansSortByDisplayname,
}

var mappingListSchedulingPlansSortByEnumLowerCase = map[string]ListSchedulingPlansSortByEnum{
	"timecreated": ListSchedulingPlansSortByTimecreated,
	"displayname": ListSchedulingPlansSortByDisplayname,
}

// GetListSchedulingPlansSortByEnumValues Enumerates the set of values for ListSchedulingPlansSortByEnum
func GetListSchedulingPlansSortByEnumValues() []ListSchedulingPlansSortByEnum {
	values := make([]ListSchedulingPlansSortByEnum, 0)
	for _, v := range mappingListSchedulingPlansSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchedulingPlansSortByEnumStringValues Enumerates the set of values in String for ListSchedulingPlansSortByEnum
func GetListSchedulingPlansSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListSchedulingPlansSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchedulingPlansSortByEnum(val string) (ListSchedulingPlansSortByEnum, bool) {
	enum, ok := mappingListSchedulingPlansSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSchedulingPlansSortOrderEnum Enum with underlying type: string
type ListSchedulingPlansSortOrderEnum string

// Set of constants representing the allowable values for ListSchedulingPlansSortOrderEnum
const (
	ListSchedulingPlansSortOrderAsc  ListSchedulingPlansSortOrderEnum = "ASC"
	ListSchedulingPlansSortOrderDesc ListSchedulingPlansSortOrderEnum = "DESC"
)

var mappingListSchedulingPlansSortOrderEnum = map[string]ListSchedulingPlansSortOrderEnum{
	"ASC":  ListSchedulingPlansSortOrderAsc,
	"DESC": ListSchedulingPlansSortOrderDesc,
}

var mappingListSchedulingPlansSortOrderEnumLowerCase = map[string]ListSchedulingPlansSortOrderEnum{
	"asc":  ListSchedulingPlansSortOrderAsc,
	"desc": ListSchedulingPlansSortOrderDesc,
}

// GetListSchedulingPlansSortOrderEnumValues Enumerates the set of values for ListSchedulingPlansSortOrderEnum
func GetListSchedulingPlansSortOrderEnumValues() []ListSchedulingPlansSortOrderEnum {
	values := make([]ListSchedulingPlansSortOrderEnum, 0)
	for _, v := range mappingListSchedulingPlansSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchedulingPlansSortOrderEnumStringValues Enumerates the set of values in String for ListSchedulingPlansSortOrderEnum
func GetListSchedulingPlansSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSchedulingPlansSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchedulingPlansSortOrderEnum(val string) (ListSchedulingPlansSortOrderEnum, bool) {
	enum, ok := mappingListSchedulingPlansSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
