// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEventsAnalyticsRequest wrapper for the ListEventsAnalytics operation
type ListEventsAnalyticsRequest struct {

	// A filter that returns events analytics starting from date provided.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStart *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeStart"`

	// A filter that returns events analytics starting up to date provided.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeEnd *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeEnd"`

	// A filter that returns events analytics grouped by specified resource.
	// Example: `2016-08-25T21:10:29.600Z`
	GroupBy ListEventsAnalyticsGroupByEnum `mandatory:"true" contributesTo:"query" name:"groupBy" omitEmpty:"true"`

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only events that match the type provided.
	Type ListEventsAnalyticsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. This filter returns resources associated with the specified resource.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListEventsAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName is ascending.
	SortBy ListEventsAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Indicates whether to list only resources managed by the Autonomous Linux service.
	IsManagedByAutonomousLinux *bool `mandatory:"false" contributesTo:"query" name:"isManagedByAutonomousLinux"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEventsAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEventsAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEventsAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEventsAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEventsAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEventsAnalyticsGroupByEnum(string(request.GroupBy)); !ok && request.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", request.GroupBy, strings.Join(GetListEventsAnalyticsGroupByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEventsAnalyticsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListEventsAnalyticsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEventsAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEventsAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEventsAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEventsAnalyticsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEventsAnalyticsResponse wrapper for the ListEventsAnalytics operation
type ListEventsAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EventsAnalyticsCollection instances
	EventsAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEventsAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEventsAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEventsAnalyticsGroupByEnum Enum with underlying type: string
type ListEventsAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListEventsAnalyticsGroupByEnum
const (
	ListEventsAnalyticsGroupByResourceid     ListEventsAnalyticsGroupByEnum = "resourceId"
	ListEventsAnalyticsGroupByEventtype      ListEventsAnalyticsGroupByEnum = "eventType"
	ListEventsAnalyticsGroupByLifecyclestate ListEventsAnalyticsGroupByEnum = "lifecycleState"
)

var mappingListEventsAnalyticsGroupByEnum = map[string]ListEventsAnalyticsGroupByEnum{
	"resourceId":     ListEventsAnalyticsGroupByResourceid,
	"eventType":      ListEventsAnalyticsGroupByEventtype,
	"lifecycleState": ListEventsAnalyticsGroupByLifecyclestate,
}

var mappingListEventsAnalyticsGroupByEnumLowerCase = map[string]ListEventsAnalyticsGroupByEnum{
	"resourceid":     ListEventsAnalyticsGroupByResourceid,
	"eventtype":      ListEventsAnalyticsGroupByEventtype,
	"lifecyclestate": ListEventsAnalyticsGroupByLifecyclestate,
}

// GetListEventsAnalyticsGroupByEnumValues Enumerates the set of values for ListEventsAnalyticsGroupByEnum
func GetListEventsAnalyticsGroupByEnumValues() []ListEventsAnalyticsGroupByEnum {
	values := make([]ListEventsAnalyticsGroupByEnum, 0)
	for _, v := range mappingListEventsAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEventsAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListEventsAnalyticsGroupByEnum
func GetListEventsAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"resourceId",
		"eventType",
		"lifecycleState",
	}
}

// GetMappingListEventsAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEventsAnalyticsGroupByEnum(val string) (ListEventsAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListEventsAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEventsAnalyticsTypeEnum Enum with underlying type: string
type ListEventsAnalyticsTypeEnum string

// Set of constants representing the allowable values for ListEventsAnalyticsTypeEnum
const (
	ListEventsAnalyticsTypeKernelOops        ListEventsAnalyticsTypeEnum = "KERNEL_OOPS"
	ListEventsAnalyticsTypeKernelCrash       ListEventsAnalyticsTypeEnum = "KERNEL_CRASH"
	ListEventsAnalyticsTypeExploitAttempt    ListEventsAnalyticsTypeEnum = "EXPLOIT_ATTEMPT"
	ListEventsAnalyticsTypeSoftwareUpdate    ListEventsAnalyticsTypeEnum = "SOFTWARE_UPDATE"
	ListEventsAnalyticsTypeKspliceUpdate     ListEventsAnalyticsTypeEnum = "KSPLICE_UPDATE"
	ListEventsAnalyticsTypeSoftwareSource    ListEventsAnalyticsTypeEnum = "SOFTWARE_SOURCE"
	ListEventsAnalyticsTypeAgent             ListEventsAnalyticsTypeEnum = "AGENT"
	ListEventsAnalyticsTypeManagementStation ListEventsAnalyticsTypeEnum = "MANAGEMENT_STATION"
	ListEventsAnalyticsTypeSysadmin          ListEventsAnalyticsTypeEnum = "SYSADMIN"
	ListEventsAnalyticsTypeReboot            ListEventsAnalyticsTypeEnum = "REBOOT"
	ListEventsAnalyticsTypeWindowsUpdate     ListEventsAnalyticsTypeEnum = "WINDOWS_UPDATE"
	ListEventsAnalyticsTypeScheduledJob      ListEventsAnalyticsTypeEnum = "SCHEDULED_JOB"
	ListEventsAnalyticsTypeRegistration      ListEventsAnalyticsTypeEnum = "REGISTRATION"
	ListEventsAnalyticsTypeSnapUpdate        ListEventsAnalyticsTypeEnum = "SNAP_UPDATE"
	ListEventsAnalyticsTypeReport            ListEventsAnalyticsTypeEnum = "REPORT"
)

var mappingListEventsAnalyticsTypeEnum = map[string]ListEventsAnalyticsTypeEnum{
	"KERNEL_OOPS":        ListEventsAnalyticsTypeKernelOops,
	"KERNEL_CRASH":       ListEventsAnalyticsTypeKernelCrash,
	"EXPLOIT_ATTEMPT":    ListEventsAnalyticsTypeExploitAttempt,
	"SOFTWARE_UPDATE":    ListEventsAnalyticsTypeSoftwareUpdate,
	"KSPLICE_UPDATE":     ListEventsAnalyticsTypeKspliceUpdate,
	"SOFTWARE_SOURCE":    ListEventsAnalyticsTypeSoftwareSource,
	"AGENT":              ListEventsAnalyticsTypeAgent,
	"MANAGEMENT_STATION": ListEventsAnalyticsTypeManagementStation,
	"SYSADMIN":           ListEventsAnalyticsTypeSysadmin,
	"REBOOT":             ListEventsAnalyticsTypeReboot,
	"WINDOWS_UPDATE":     ListEventsAnalyticsTypeWindowsUpdate,
	"SCHEDULED_JOB":      ListEventsAnalyticsTypeScheduledJob,
	"REGISTRATION":       ListEventsAnalyticsTypeRegistration,
	"SNAP_UPDATE":        ListEventsAnalyticsTypeSnapUpdate,
	"REPORT":             ListEventsAnalyticsTypeReport,
}

var mappingListEventsAnalyticsTypeEnumLowerCase = map[string]ListEventsAnalyticsTypeEnum{
	"kernel_oops":        ListEventsAnalyticsTypeKernelOops,
	"kernel_crash":       ListEventsAnalyticsTypeKernelCrash,
	"exploit_attempt":    ListEventsAnalyticsTypeExploitAttempt,
	"software_update":    ListEventsAnalyticsTypeSoftwareUpdate,
	"ksplice_update":     ListEventsAnalyticsTypeKspliceUpdate,
	"software_source":    ListEventsAnalyticsTypeSoftwareSource,
	"agent":              ListEventsAnalyticsTypeAgent,
	"management_station": ListEventsAnalyticsTypeManagementStation,
	"sysadmin":           ListEventsAnalyticsTypeSysadmin,
	"reboot":             ListEventsAnalyticsTypeReboot,
	"windows_update":     ListEventsAnalyticsTypeWindowsUpdate,
	"scheduled_job":      ListEventsAnalyticsTypeScheduledJob,
	"registration":       ListEventsAnalyticsTypeRegistration,
	"snap_update":        ListEventsAnalyticsTypeSnapUpdate,
	"report":             ListEventsAnalyticsTypeReport,
}

// GetListEventsAnalyticsTypeEnumValues Enumerates the set of values for ListEventsAnalyticsTypeEnum
func GetListEventsAnalyticsTypeEnumValues() []ListEventsAnalyticsTypeEnum {
	values := make([]ListEventsAnalyticsTypeEnum, 0)
	for _, v := range mappingListEventsAnalyticsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListEventsAnalyticsTypeEnumStringValues Enumerates the set of values in String for ListEventsAnalyticsTypeEnum
func GetListEventsAnalyticsTypeEnumStringValues() []string {
	return []string{
		"KERNEL_OOPS",
		"KERNEL_CRASH",
		"EXPLOIT_ATTEMPT",
		"SOFTWARE_UPDATE",
		"KSPLICE_UPDATE",
		"SOFTWARE_SOURCE",
		"AGENT",
		"MANAGEMENT_STATION",
		"SYSADMIN",
		"REBOOT",
		"WINDOWS_UPDATE",
		"SCHEDULED_JOB",
		"REGISTRATION",
		"SNAP_UPDATE",
		"REPORT",
	}
}

// GetMappingListEventsAnalyticsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEventsAnalyticsTypeEnum(val string) (ListEventsAnalyticsTypeEnum, bool) {
	enum, ok := mappingListEventsAnalyticsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEventsAnalyticsSortOrderEnum Enum with underlying type: string
type ListEventsAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for ListEventsAnalyticsSortOrderEnum
const (
	ListEventsAnalyticsSortOrderAsc  ListEventsAnalyticsSortOrderEnum = "ASC"
	ListEventsAnalyticsSortOrderDesc ListEventsAnalyticsSortOrderEnum = "DESC"
)

var mappingListEventsAnalyticsSortOrderEnum = map[string]ListEventsAnalyticsSortOrderEnum{
	"ASC":  ListEventsAnalyticsSortOrderAsc,
	"DESC": ListEventsAnalyticsSortOrderDesc,
}

var mappingListEventsAnalyticsSortOrderEnumLowerCase = map[string]ListEventsAnalyticsSortOrderEnum{
	"asc":  ListEventsAnalyticsSortOrderAsc,
	"desc": ListEventsAnalyticsSortOrderDesc,
}

// GetListEventsAnalyticsSortOrderEnumValues Enumerates the set of values for ListEventsAnalyticsSortOrderEnum
func GetListEventsAnalyticsSortOrderEnumValues() []ListEventsAnalyticsSortOrderEnum {
	values := make([]ListEventsAnalyticsSortOrderEnum, 0)
	for _, v := range mappingListEventsAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEventsAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for ListEventsAnalyticsSortOrderEnum
func GetListEventsAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEventsAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEventsAnalyticsSortOrderEnum(val string) (ListEventsAnalyticsSortOrderEnum, bool) {
	enum, ok := mappingListEventsAnalyticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEventsAnalyticsSortByEnum Enum with underlying type: string
type ListEventsAnalyticsSortByEnum string

// Set of constants representing the allowable values for ListEventsAnalyticsSortByEnum
const (
	ListEventsAnalyticsSortByTimecreated    ListEventsAnalyticsSortByEnum = "timeCreated"
	ListEventsAnalyticsSortByTimeoccurredat ListEventsAnalyticsSortByEnum = "timeOccurredAt"
	ListEventsAnalyticsSortByDisplayname    ListEventsAnalyticsSortByEnum = "displayName"
)

var mappingListEventsAnalyticsSortByEnum = map[string]ListEventsAnalyticsSortByEnum{
	"timeCreated":    ListEventsAnalyticsSortByTimecreated,
	"timeOccurredAt": ListEventsAnalyticsSortByTimeoccurredat,
	"displayName":    ListEventsAnalyticsSortByDisplayname,
}

var mappingListEventsAnalyticsSortByEnumLowerCase = map[string]ListEventsAnalyticsSortByEnum{
	"timecreated":    ListEventsAnalyticsSortByTimecreated,
	"timeoccurredat": ListEventsAnalyticsSortByTimeoccurredat,
	"displayname":    ListEventsAnalyticsSortByDisplayname,
}

// GetListEventsAnalyticsSortByEnumValues Enumerates the set of values for ListEventsAnalyticsSortByEnum
func GetListEventsAnalyticsSortByEnumValues() []ListEventsAnalyticsSortByEnum {
	values := make([]ListEventsAnalyticsSortByEnum, 0)
	for _, v := range mappingListEventsAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEventsAnalyticsSortByEnumStringValues Enumerates the set of values in String for ListEventsAnalyticsSortByEnum
func GetListEventsAnalyticsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeOccurredAt",
		"displayName",
	}
}

// GetMappingListEventsAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEventsAnalyticsSortByEnum(val string) (ListEventsAnalyticsSortByEnum, bool) {
	enum, ok := mappingListEventsAnalyticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
