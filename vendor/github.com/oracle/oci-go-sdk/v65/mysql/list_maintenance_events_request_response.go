// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMaintenanceEventsRequest wrapper for the ListMaintenanceEvents operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ListMaintenanceEvents.go.html to see an example of how to use ListMaintenanceEventsRequest.
type ListMaintenanceEventsRequest struct {

	// The DB System OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"true" contributesTo:"path" name:"dbSystemId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The MySQL version before the maintenance event.
	MysqlVersionBeforeMaintenance *string `mandatory:"false" contributesTo:"query" name:"mysqlVersionBeforeMaintenance"`

	// The MySQL version after the maintenance event.
	MysqlVersionAfterMaintenance *string `mandatory:"false" contributesTo:"query" name:"mysqlVersionAfterMaintenance"`

	// How the maintenance event was triggered.
	MaintenanceType ListMaintenanceEventsMaintenanceTypeEnum `mandatory:"false" contributesTo:"query" name:"maintenanceType" omitEmpty:"true"`

	// The nature of the maintenance event.
	MaintenanceAction ListMaintenanceEventsMaintenanceActionEnum `mandatory:"false" contributesTo:"query" name:"maintenanceAction" omitEmpty:"true"`

	// The last status of the maintenance event.
	MaintenanceStatus MaintenanceEventMaintenanceStatusEnum `mandatory:"false" contributesTo:"query" name:"maintenanceStatus" omitEmpty:"true"`

	// The maximum number of items to return in a paginated list call. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaasAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous list call. For information about pagination, see List
	// Pagination (https://docs.oracle.com/iaasAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided.
	// Time fields are default ordered as descending.
	SortBy ListMaintenanceEventsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ASC or DESC).
	SortOrder ListMaintenanceEventsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaintenanceEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaintenanceEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaintenanceEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaintenanceEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaintenanceEventsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaintenanceEventsMaintenanceTypeEnum(string(request.MaintenanceType)); !ok && request.MaintenanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceType: %s. Supported values are: %s.", request.MaintenanceType, strings.Join(GetListMaintenanceEventsMaintenanceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceEventsMaintenanceActionEnum(string(request.MaintenanceAction)); !ok && request.MaintenanceAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceAction: %s. Supported values are: %s.", request.MaintenanceAction, strings.Join(GetListMaintenanceEventsMaintenanceActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceEventMaintenanceStatusEnum(string(request.MaintenanceStatus)); !ok && request.MaintenanceStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceStatus: %s. Supported values are: %s.", request.MaintenanceStatus, strings.Join(GetMaintenanceEventMaintenanceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceEventsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaintenanceEventsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceEventsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaintenanceEventsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaintenanceEventsResponse wrapper for the ListMaintenanceEvents operation
type ListMaintenanceEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []MaintenanceEvent instances
	Items []MaintenanceEvent `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMaintenanceEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaintenanceEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaintenanceEventsMaintenanceTypeEnum Enum with underlying type: string
type ListMaintenanceEventsMaintenanceTypeEnum string

// Set of constants representing the allowable values for ListMaintenanceEventsMaintenanceTypeEnum
const (
	ListMaintenanceEventsMaintenanceTypeAutomatic ListMaintenanceEventsMaintenanceTypeEnum = "AUTOMATIC"
	ListMaintenanceEventsMaintenanceTypeManual    ListMaintenanceEventsMaintenanceTypeEnum = "MANUAL"
	ListMaintenanceEventsMaintenanceTypeShape     ListMaintenanceEventsMaintenanceTypeEnum = "SHAPE"
)

var mappingListMaintenanceEventsMaintenanceTypeEnum = map[string]ListMaintenanceEventsMaintenanceTypeEnum{
	"AUTOMATIC": ListMaintenanceEventsMaintenanceTypeAutomatic,
	"MANUAL":    ListMaintenanceEventsMaintenanceTypeManual,
	"SHAPE":     ListMaintenanceEventsMaintenanceTypeShape,
}

var mappingListMaintenanceEventsMaintenanceTypeEnumLowerCase = map[string]ListMaintenanceEventsMaintenanceTypeEnum{
	"automatic": ListMaintenanceEventsMaintenanceTypeAutomatic,
	"manual":    ListMaintenanceEventsMaintenanceTypeManual,
	"shape":     ListMaintenanceEventsMaintenanceTypeShape,
}

// GetListMaintenanceEventsMaintenanceTypeEnumValues Enumerates the set of values for ListMaintenanceEventsMaintenanceTypeEnum
func GetListMaintenanceEventsMaintenanceTypeEnumValues() []ListMaintenanceEventsMaintenanceTypeEnum {
	values := make([]ListMaintenanceEventsMaintenanceTypeEnum, 0)
	for _, v := range mappingListMaintenanceEventsMaintenanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceEventsMaintenanceTypeEnumStringValues Enumerates the set of values in String for ListMaintenanceEventsMaintenanceTypeEnum
func GetListMaintenanceEventsMaintenanceTypeEnumStringValues() []string {
	return []string{
		"AUTOMATIC",
		"MANUAL",
		"SHAPE",
	}
}

// GetMappingListMaintenanceEventsMaintenanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceEventsMaintenanceTypeEnum(val string) (ListMaintenanceEventsMaintenanceTypeEnum, bool) {
	enum, ok := mappingListMaintenanceEventsMaintenanceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceEventsMaintenanceActionEnum Enum with underlying type: string
type ListMaintenanceEventsMaintenanceActionEnum string

// Set of constants representing the allowable values for ListMaintenanceEventsMaintenanceActionEnum
const (
	ListMaintenanceEventsMaintenanceActionDatabase     ListMaintenanceEventsMaintenanceActionEnum = "DATABASE"
	ListMaintenanceEventsMaintenanceActionOsUpdate     ListMaintenanceEventsMaintenanceActionEnum = "OS_UPDATE"
	ListMaintenanceEventsMaintenanceActionOnlineUpdate ListMaintenanceEventsMaintenanceActionEnum = "ONLINE_UPDATE"
	ListMaintenanceEventsMaintenanceActionHardware     ListMaintenanceEventsMaintenanceActionEnum = "HARDWARE"
)

var mappingListMaintenanceEventsMaintenanceActionEnum = map[string]ListMaintenanceEventsMaintenanceActionEnum{
	"DATABASE":      ListMaintenanceEventsMaintenanceActionDatabase,
	"OS_UPDATE":     ListMaintenanceEventsMaintenanceActionOsUpdate,
	"ONLINE_UPDATE": ListMaintenanceEventsMaintenanceActionOnlineUpdate,
	"HARDWARE":      ListMaintenanceEventsMaintenanceActionHardware,
}

var mappingListMaintenanceEventsMaintenanceActionEnumLowerCase = map[string]ListMaintenanceEventsMaintenanceActionEnum{
	"database":      ListMaintenanceEventsMaintenanceActionDatabase,
	"os_update":     ListMaintenanceEventsMaintenanceActionOsUpdate,
	"online_update": ListMaintenanceEventsMaintenanceActionOnlineUpdate,
	"hardware":      ListMaintenanceEventsMaintenanceActionHardware,
}

// GetListMaintenanceEventsMaintenanceActionEnumValues Enumerates the set of values for ListMaintenanceEventsMaintenanceActionEnum
func GetListMaintenanceEventsMaintenanceActionEnumValues() []ListMaintenanceEventsMaintenanceActionEnum {
	values := make([]ListMaintenanceEventsMaintenanceActionEnum, 0)
	for _, v := range mappingListMaintenanceEventsMaintenanceActionEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceEventsMaintenanceActionEnumStringValues Enumerates the set of values in String for ListMaintenanceEventsMaintenanceActionEnum
func GetListMaintenanceEventsMaintenanceActionEnumStringValues() []string {
	return []string{
		"DATABASE",
		"OS_UPDATE",
		"ONLINE_UPDATE",
		"HARDWARE",
	}
}

// GetMappingListMaintenanceEventsMaintenanceActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceEventsMaintenanceActionEnum(val string) (ListMaintenanceEventsMaintenanceActionEnum, bool) {
	enum, ok := mappingListMaintenanceEventsMaintenanceActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceEventsSortByEnum Enum with underlying type: string
type ListMaintenanceEventsSortByEnum string

// Set of constants representing the allowable values for ListMaintenanceEventsSortByEnum
const (
	ListMaintenanceEventsSortByTimecreated   ListMaintenanceEventsSortByEnum = "timeCreated"
	ListMaintenanceEventsSortByTimescheduled ListMaintenanceEventsSortByEnum = "timeScheduled"
	ListMaintenanceEventsSortByTimestarted   ListMaintenanceEventsSortByEnum = "timeStarted"
	ListMaintenanceEventsSortByTimeended     ListMaintenanceEventsSortByEnum = "timeEnded"
)

var mappingListMaintenanceEventsSortByEnum = map[string]ListMaintenanceEventsSortByEnum{
	"timeCreated":   ListMaintenanceEventsSortByTimecreated,
	"timeScheduled": ListMaintenanceEventsSortByTimescheduled,
	"timeStarted":   ListMaintenanceEventsSortByTimestarted,
	"timeEnded":     ListMaintenanceEventsSortByTimeended,
}

var mappingListMaintenanceEventsSortByEnumLowerCase = map[string]ListMaintenanceEventsSortByEnum{
	"timecreated":   ListMaintenanceEventsSortByTimecreated,
	"timescheduled": ListMaintenanceEventsSortByTimescheduled,
	"timestarted":   ListMaintenanceEventsSortByTimestarted,
	"timeended":     ListMaintenanceEventsSortByTimeended,
}

// GetListMaintenanceEventsSortByEnumValues Enumerates the set of values for ListMaintenanceEventsSortByEnum
func GetListMaintenanceEventsSortByEnumValues() []ListMaintenanceEventsSortByEnum {
	values := make([]ListMaintenanceEventsSortByEnum, 0)
	for _, v := range mappingListMaintenanceEventsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceEventsSortByEnumStringValues Enumerates the set of values in String for ListMaintenanceEventsSortByEnum
func GetListMaintenanceEventsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeScheduled",
		"timeStarted",
		"timeEnded",
	}
}

// GetMappingListMaintenanceEventsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceEventsSortByEnum(val string) (ListMaintenanceEventsSortByEnum, bool) {
	enum, ok := mappingListMaintenanceEventsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceEventsSortOrderEnum Enum with underlying type: string
type ListMaintenanceEventsSortOrderEnum string

// Set of constants representing the allowable values for ListMaintenanceEventsSortOrderEnum
const (
	ListMaintenanceEventsSortOrderAsc  ListMaintenanceEventsSortOrderEnum = "ASC"
	ListMaintenanceEventsSortOrderDesc ListMaintenanceEventsSortOrderEnum = "DESC"
)

var mappingListMaintenanceEventsSortOrderEnum = map[string]ListMaintenanceEventsSortOrderEnum{
	"ASC":  ListMaintenanceEventsSortOrderAsc,
	"DESC": ListMaintenanceEventsSortOrderDesc,
}

var mappingListMaintenanceEventsSortOrderEnumLowerCase = map[string]ListMaintenanceEventsSortOrderEnum{
	"asc":  ListMaintenanceEventsSortOrderAsc,
	"desc": ListMaintenanceEventsSortOrderDesc,
}

// GetListMaintenanceEventsSortOrderEnumValues Enumerates the set of values for ListMaintenanceEventsSortOrderEnum
func GetListMaintenanceEventsSortOrderEnumValues() []ListMaintenanceEventsSortOrderEnum {
	values := make([]ListMaintenanceEventsSortOrderEnum, 0)
	for _, v := range mappingListMaintenanceEventsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceEventsSortOrderEnumStringValues Enumerates the set of values in String for ListMaintenanceEventsSortOrderEnum
func GetListMaintenanceEventsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaintenanceEventsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceEventsSortOrderEnum(val string) (ListMaintenanceEventsSortOrderEnum, bool) {
	enum, ok := mappingListMaintenanceEventsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
