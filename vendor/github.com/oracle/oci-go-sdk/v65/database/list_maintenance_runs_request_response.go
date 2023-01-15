// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMaintenanceRunsRequest wrapper for the ListMaintenanceRuns operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListMaintenanceRuns.go.html to see an example of how to use ListMaintenanceRunsRequest.
type ListMaintenanceRunsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The target resource ID.
	TargetResourceId *string `mandatory:"false" contributesTo:"query" name:"targetResourceId"`

	// The type of the target resource.
	TargetResourceType MaintenanceRunSummaryTargetResourceTypeEnum `mandatory:"false" contributesTo:"query" name:"targetResourceType" omitEmpty:"true"`

	// The maintenance type.
	MaintenanceType MaintenanceRunSummaryMaintenanceTypeEnum `mandatory:"false" contributesTo:"query" name:"maintenanceType" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIME_SCHEDULED and TIME_ENDED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	// **Note:** If you do not include the availability domain filter, the resources are grouped by availability domain, then sorted.
	SortBy ListMaintenanceRunsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMaintenanceRunsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState MaintenanceRunSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given availability domain exactly.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// The sub-type of the maintenance run.
	MaintenanceSubtype MaintenanceRunSummaryMaintenanceSubtypeEnum `mandatory:"false" contributesTo:"query" name:"maintenanceSubtype" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaintenanceRunsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaintenanceRunsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaintenanceRunsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaintenanceRunsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaintenanceRunsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaintenanceRunSummaryTargetResourceTypeEnum(string(request.TargetResourceType)); !ok && request.TargetResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetResourceType: %s. Supported values are: %s.", request.TargetResourceType, strings.Join(GetMaintenanceRunSummaryTargetResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceRunSummaryMaintenanceTypeEnum(string(request.MaintenanceType)); !ok && request.MaintenanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceType: %s. Supported values are: %s.", request.MaintenanceType, strings.Join(GetMaintenanceRunSummaryMaintenanceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceRunsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaintenanceRunsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceRunsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaintenanceRunsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceRunSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMaintenanceRunSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceRunSummaryMaintenanceSubtypeEnum(string(request.MaintenanceSubtype)); !ok && request.MaintenanceSubtype != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceSubtype: %s. Supported values are: %s.", request.MaintenanceSubtype, strings.Join(GetMaintenanceRunSummaryMaintenanceSubtypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaintenanceRunsResponse wrapper for the ListMaintenanceRuns operation
type ListMaintenanceRunsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []MaintenanceRunSummary instances
	Items []MaintenanceRunSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMaintenanceRunsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaintenanceRunsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaintenanceRunsSortByEnum Enum with underlying type: string
type ListMaintenanceRunsSortByEnum string

// Set of constants representing the allowable values for ListMaintenanceRunsSortByEnum
const (
	ListMaintenanceRunsSortByTimeScheduled ListMaintenanceRunsSortByEnum = "TIME_SCHEDULED"
	ListMaintenanceRunsSortByTimeEnded     ListMaintenanceRunsSortByEnum = "TIME_ENDED"
	ListMaintenanceRunsSortByDisplayname   ListMaintenanceRunsSortByEnum = "DISPLAYNAME"
)

var mappingListMaintenanceRunsSortByEnum = map[string]ListMaintenanceRunsSortByEnum{
	"TIME_SCHEDULED": ListMaintenanceRunsSortByTimeScheduled,
	"TIME_ENDED":     ListMaintenanceRunsSortByTimeEnded,
	"DISPLAYNAME":    ListMaintenanceRunsSortByDisplayname,
}

var mappingListMaintenanceRunsSortByEnumLowerCase = map[string]ListMaintenanceRunsSortByEnum{
	"time_scheduled": ListMaintenanceRunsSortByTimeScheduled,
	"time_ended":     ListMaintenanceRunsSortByTimeEnded,
	"displayname":    ListMaintenanceRunsSortByDisplayname,
}

// GetListMaintenanceRunsSortByEnumValues Enumerates the set of values for ListMaintenanceRunsSortByEnum
func GetListMaintenanceRunsSortByEnumValues() []ListMaintenanceRunsSortByEnum {
	values := make([]ListMaintenanceRunsSortByEnum, 0)
	for _, v := range mappingListMaintenanceRunsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceRunsSortByEnumStringValues Enumerates the set of values in String for ListMaintenanceRunsSortByEnum
func GetListMaintenanceRunsSortByEnumStringValues() []string {
	return []string{
		"TIME_SCHEDULED",
		"TIME_ENDED",
		"DISPLAYNAME",
	}
}

// GetMappingListMaintenanceRunsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceRunsSortByEnum(val string) (ListMaintenanceRunsSortByEnum, bool) {
	enum, ok := mappingListMaintenanceRunsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceRunsSortOrderEnum Enum with underlying type: string
type ListMaintenanceRunsSortOrderEnum string

// Set of constants representing the allowable values for ListMaintenanceRunsSortOrderEnum
const (
	ListMaintenanceRunsSortOrderAsc  ListMaintenanceRunsSortOrderEnum = "ASC"
	ListMaintenanceRunsSortOrderDesc ListMaintenanceRunsSortOrderEnum = "DESC"
)

var mappingListMaintenanceRunsSortOrderEnum = map[string]ListMaintenanceRunsSortOrderEnum{
	"ASC":  ListMaintenanceRunsSortOrderAsc,
	"DESC": ListMaintenanceRunsSortOrderDesc,
}

var mappingListMaintenanceRunsSortOrderEnumLowerCase = map[string]ListMaintenanceRunsSortOrderEnum{
	"asc":  ListMaintenanceRunsSortOrderAsc,
	"desc": ListMaintenanceRunsSortOrderDesc,
}

// GetListMaintenanceRunsSortOrderEnumValues Enumerates the set of values for ListMaintenanceRunsSortOrderEnum
func GetListMaintenanceRunsSortOrderEnumValues() []ListMaintenanceRunsSortOrderEnum {
	values := make([]ListMaintenanceRunsSortOrderEnum, 0)
	for _, v := range mappingListMaintenanceRunsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceRunsSortOrderEnumStringValues Enumerates the set of values in String for ListMaintenanceRunsSortOrderEnum
func GetListMaintenanceRunsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaintenanceRunsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceRunsSortOrderEnum(val string) (ListMaintenanceRunsSortOrderEnum, bool) {
	enum, ok := mappingListMaintenanceRunsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
