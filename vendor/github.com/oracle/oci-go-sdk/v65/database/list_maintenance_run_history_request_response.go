// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMaintenanceRunHistoryRequest wrapper for the ListMaintenanceRunHistory operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListMaintenanceRunHistory.go.html to see an example of how to use ListMaintenanceRunHistoryRequest.
type ListMaintenanceRunHistoryRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
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
	SortBy ListMaintenanceRunHistorySortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMaintenanceRunHistorySortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The state of the maintenance run history.
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

func (request ListMaintenanceRunHistoryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaintenanceRunHistoryRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaintenanceRunHistoryRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaintenanceRunHistoryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaintenanceRunHistoryRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaintenanceRunSummaryTargetResourceTypeEnum(string(request.TargetResourceType)); !ok && request.TargetResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetResourceType: %s. Supported values are: %s.", request.TargetResourceType, strings.Join(GetMaintenanceRunSummaryTargetResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceRunSummaryMaintenanceTypeEnum(string(request.MaintenanceType)); !ok && request.MaintenanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceType: %s. Supported values are: %s.", request.MaintenanceType, strings.Join(GetMaintenanceRunSummaryMaintenanceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceRunHistorySortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaintenanceRunHistorySortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceRunHistorySortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaintenanceRunHistorySortOrderEnumStringValues(), ",")))
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

// ListMaintenanceRunHistoryResponse wrapper for the ListMaintenanceRunHistory operation
type ListMaintenanceRunHistoryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []MaintenanceRunHistorySummary instances
	Items []MaintenanceRunHistorySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMaintenanceRunHistoryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaintenanceRunHistoryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaintenanceRunHistorySortByEnum Enum with underlying type: string
type ListMaintenanceRunHistorySortByEnum string

// Set of constants representing the allowable values for ListMaintenanceRunHistorySortByEnum
const (
	ListMaintenanceRunHistorySortByTimeScheduled ListMaintenanceRunHistorySortByEnum = "TIME_SCHEDULED"
	ListMaintenanceRunHistorySortByTimeEnded     ListMaintenanceRunHistorySortByEnum = "TIME_ENDED"
	ListMaintenanceRunHistorySortByDisplayname   ListMaintenanceRunHistorySortByEnum = "DISPLAYNAME"
)

var mappingListMaintenanceRunHistorySortByEnum = map[string]ListMaintenanceRunHistorySortByEnum{
	"TIME_SCHEDULED": ListMaintenanceRunHistorySortByTimeScheduled,
	"TIME_ENDED":     ListMaintenanceRunHistorySortByTimeEnded,
	"DISPLAYNAME":    ListMaintenanceRunHistorySortByDisplayname,
}

var mappingListMaintenanceRunHistorySortByEnumLowerCase = map[string]ListMaintenanceRunHistorySortByEnum{
	"time_scheduled": ListMaintenanceRunHistorySortByTimeScheduled,
	"time_ended":     ListMaintenanceRunHistorySortByTimeEnded,
	"displayname":    ListMaintenanceRunHistorySortByDisplayname,
}

// GetListMaintenanceRunHistorySortByEnumValues Enumerates the set of values for ListMaintenanceRunHistorySortByEnum
func GetListMaintenanceRunHistorySortByEnumValues() []ListMaintenanceRunHistorySortByEnum {
	values := make([]ListMaintenanceRunHistorySortByEnum, 0)
	for _, v := range mappingListMaintenanceRunHistorySortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceRunHistorySortByEnumStringValues Enumerates the set of values in String for ListMaintenanceRunHistorySortByEnum
func GetListMaintenanceRunHistorySortByEnumStringValues() []string {
	return []string{
		"TIME_SCHEDULED",
		"TIME_ENDED",
		"DISPLAYNAME",
	}
}

// GetMappingListMaintenanceRunHistorySortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceRunHistorySortByEnum(val string) (ListMaintenanceRunHistorySortByEnum, bool) {
	enum, ok := mappingListMaintenanceRunHistorySortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceRunHistorySortOrderEnum Enum with underlying type: string
type ListMaintenanceRunHistorySortOrderEnum string

// Set of constants representing the allowable values for ListMaintenanceRunHistorySortOrderEnum
const (
	ListMaintenanceRunHistorySortOrderAsc  ListMaintenanceRunHistorySortOrderEnum = "ASC"
	ListMaintenanceRunHistorySortOrderDesc ListMaintenanceRunHistorySortOrderEnum = "DESC"
)

var mappingListMaintenanceRunHistorySortOrderEnum = map[string]ListMaintenanceRunHistorySortOrderEnum{
	"ASC":  ListMaintenanceRunHistorySortOrderAsc,
	"DESC": ListMaintenanceRunHistorySortOrderDesc,
}

var mappingListMaintenanceRunHistorySortOrderEnumLowerCase = map[string]ListMaintenanceRunHistorySortOrderEnum{
	"asc":  ListMaintenanceRunHistorySortOrderAsc,
	"desc": ListMaintenanceRunHistorySortOrderDesc,
}

// GetListMaintenanceRunHistorySortOrderEnumValues Enumerates the set of values for ListMaintenanceRunHistorySortOrderEnum
func GetListMaintenanceRunHistorySortOrderEnumValues() []ListMaintenanceRunHistorySortOrderEnum {
	values := make([]ListMaintenanceRunHistorySortOrderEnum, 0)
	for _, v := range mappingListMaintenanceRunHistorySortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceRunHistorySortOrderEnumStringValues Enumerates the set of values in String for ListMaintenanceRunHistorySortOrderEnum
func GetListMaintenanceRunHistorySortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaintenanceRunHistorySortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceRunHistorySortOrderEnum(val string) (ListMaintenanceRunHistorySortOrderEnum, bool) {
	enum, ok := mappingListMaintenanceRunHistorySortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
