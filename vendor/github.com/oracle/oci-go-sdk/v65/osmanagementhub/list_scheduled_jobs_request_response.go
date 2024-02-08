// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListScheduledJobsRequest wrapper for the ListScheduledJobs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListScheduledJobs.go.html to see an example of how to use ListScheduledJobsRequest.
type ListScheduledJobsRequest struct {

	// The OCID of the compartment that contains the resources to list.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ScheduledJobLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OCID of the managed instance for which to list resources.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The OCID of the managed instance group for which to list resources.
	ManagedInstanceGroupId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceGroupId"`

	// The OCID of the managed compartment for which to list resources.
	ManagedCompartmentId *string `mandatory:"false" contributesTo:"query" name:"managedCompartmentId"`

	// The OCID of the lifecycle stage for which to list resources.
	LifecycleStageId *string `mandatory:"false" contributesTo:"query" name:"lifecycleStageId"`

	// The operation type for which to list resources.
	OperationType ListScheduledJobsOperationTypeEnum `mandatory:"false" contributesTo:"query" name:"operationType" omitEmpty:"true"`

	// The schedule type for which to list resources.
	ScheduleType ListScheduledJobsScheduleTypeEnum `mandatory:"false" contributesTo:"query" name:"scheduleType" omitEmpty:"true"`

	// The start time after which to list all schedules, in ISO 8601 format.
	// Example: 2017-07-14T02:40:00.000Z
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The cut-off time before which to list all upcoming schedules, in ISO 8601 format.
	// Example: 2017-07-14T02:40:00.000Z
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListScheduledJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListScheduledJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If true, will only filter out restricted scheduled job.
	IsRestricted *bool `mandatory:"false" contributesTo:"query" name:"isRestricted"`

	// The OCID of the scheduled job.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Default is false. When set to true ,returns results from {compartmentId} or any of its subcompartment.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListScheduledJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListScheduledJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListScheduledJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListScheduledJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListScheduledJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduledJobLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetScheduledJobLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledJobsOperationTypeEnum(string(request.OperationType)); !ok && request.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", request.OperationType, strings.Join(GetListScheduledJobsOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledJobsScheduleTypeEnum(string(request.ScheduleType)); !ok && request.ScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduleType: %s. Supported values are: %s.", request.ScheduleType, strings.Join(GetListScheduledJobsScheduleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListScheduledJobsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListScheduledJobsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListScheduledJobsResponse wrapper for the ListScheduledJobs operation
type ListScheduledJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ScheduledJobCollection instances
	ScheduledJobCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListScheduledJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListScheduledJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListScheduledJobsOperationTypeEnum Enum with underlying type: string
type ListScheduledJobsOperationTypeEnum string

// Set of constants representing the allowable values for ListScheduledJobsOperationTypeEnum
const (
	ListScheduledJobsOperationTypeInstallPackages             ListScheduledJobsOperationTypeEnum = "INSTALL_PACKAGES"
	ListScheduledJobsOperationTypeUpdatePackages              ListScheduledJobsOperationTypeEnum = "UPDATE_PACKAGES"
	ListScheduledJobsOperationTypeRemovePackages              ListScheduledJobsOperationTypeEnum = "REMOVE_PACKAGES"
	ListScheduledJobsOperationTypeUpdateAll                   ListScheduledJobsOperationTypeEnum = "UPDATE_ALL"
	ListScheduledJobsOperationTypeUpdateSecurity              ListScheduledJobsOperationTypeEnum = "UPDATE_SECURITY"
	ListScheduledJobsOperationTypeUpdateBugfix                ListScheduledJobsOperationTypeEnum = "UPDATE_BUGFIX"
	ListScheduledJobsOperationTypeUpdateEnhancement           ListScheduledJobsOperationTypeEnum = "UPDATE_ENHANCEMENT"
	ListScheduledJobsOperationTypeUpdateOther                 ListScheduledJobsOperationTypeEnum = "UPDATE_OTHER"
	ListScheduledJobsOperationTypeUpdateKspliceUserspace      ListScheduledJobsOperationTypeEnum = "UPDATE_KSPLICE_USERSPACE"
	ListScheduledJobsOperationTypeUpdateKspliceKernel         ListScheduledJobsOperationTypeEnum = "UPDATE_KSPLICE_KERNEL"
	ListScheduledJobsOperationTypeManageModuleStreams         ListScheduledJobsOperationTypeEnum = "MANAGE_MODULE_STREAMS"
	ListScheduledJobsOperationTypeSwitchModuleStream          ListScheduledJobsOperationTypeEnum = "SWITCH_MODULE_STREAM"
	ListScheduledJobsOperationTypeAttachSoftwareSources       ListScheduledJobsOperationTypeEnum = "ATTACH_SOFTWARE_SOURCES"
	ListScheduledJobsOperationTypeDetachSoftwareSources       ListScheduledJobsOperationTypeEnum = "DETACH_SOFTWARE_SOURCES"
	ListScheduledJobsOperationTypeSyncManagementStationMirror ListScheduledJobsOperationTypeEnum = "SYNC_MANAGEMENT_STATION_MIRROR"
	ListScheduledJobsOperationTypePromoteLifecycle            ListScheduledJobsOperationTypeEnum = "PROMOTE_LIFECYCLE"
)

var mappingListScheduledJobsOperationTypeEnum = map[string]ListScheduledJobsOperationTypeEnum{
	"INSTALL_PACKAGES":               ListScheduledJobsOperationTypeInstallPackages,
	"UPDATE_PACKAGES":                ListScheduledJobsOperationTypeUpdatePackages,
	"REMOVE_PACKAGES":                ListScheduledJobsOperationTypeRemovePackages,
	"UPDATE_ALL":                     ListScheduledJobsOperationTypeUpdateAll,
	"UPDATE_SECURITY":                ListScheduledJobsOperationTypeUpdateSecurity,
	"UPDATE_BUGFIX":                  ListScheduledJobsOperationTypeUpdateBugfix,
	"UPDATE_ENHANCEMENT":             ListScheduledJobsOperationTypeUpdateEnhancement,
	"UPDATE_OTHER":                   ListScheduledJobsOperationTypeUpdateOther,
	"UPDATE_KSPLICE_USERSPACE":       ListScheduledJobsOperationTypeUpdateKspliceUserspace,
	"UPDATE_KSPLICE_KERNEL":          ListScheduledJobsOperationTypeUpdateKspliceKernel,
	"MANAGE_MODULE_STREAMS":          ListScheduledJobsOperationTypeManageModuleStreams,
	"SWITCH_MODULE_STREAM":           ListScheduledJobsOperationTypeSwitchModuleStream,
	"ATTACH_SOFTWARE_SOURCES":        ListScheduledJobsOperationTypeAttachSoftwareSources,
	"DETACH_SOFTWARE_SOURCES":        ListScheduledJobsOperationTypeDetachSoftwareSources,
	"SYNC_MANAGEMENT_STATION_MIRROR": ListScheduledJobsOperationTypeSyncManagementStationMirror,
	"PROMOTE_LIFECYCLE":              ListScheduledJobsOperationTypePromoteLifecycle,
}

var mappingListScheduledJobsOperationTypeEnumLowerCase = map[string]ListScheduledJobsOperationTypeEnum{
	"install_packages":               ListScheduledJobsOperationTypeInstallPackages,
	"update_packages":                ListScheduledJobsOperationTypeUpdatePackages,
	"remove_packages":                ListScheduledJobsOperationTypeRemovePackages,
	"update_all":                     ListScheduledJobsOperationTypeUpdateAll,
	"update_security":                ListScheduledJobsOperationTypeUpdateSecurity,
	"update_bugfix":                  ListScheduledJobsOperationTypeUpdateBugfix,
	"update_enhancement":             ListScheduledJobsOperationTypeUpdateEnhancement,
	"update_other":                   ListScheduledJobsOperationTypeUpdateOther,
	"update_ksplice_userspace":       ListScheduledJobsOperationTypeUpdateKspliceUserspace,
	"update_ksplice_kernel":          ListScheduledJobsOperationTypeUpdateKspliceKernel,
	"manage_module_streams":          ListScheduledJobsOperationTypeManageModuleStreams,
	"switch_module_stream":           ListScheduledJobsOperationTypeSwitchModuleStream,
	"attach_software_sources":        ListScheduledJobsOperationTypeAttachSoftwareSources,
	"detach_software_sources":        ListScheduledJobsOperationTypeDetachSoftwareSources,
	"sync_management_station_mirror": ListScheduledJobsOperationTypeSyncManagementStationMirror,
	"promote_lifecycle":              ListScheduledJobsOperationTypePromoteLifecycle,
}

// GetListScheduledJobsOperationTypeEnumValues Enumerates the set of values for ListScheduledJobsOperationTypeEnum
func GetListScheduledJobsOperationTypeEnumValues() []ListScheduledJobsOperationTypeEnum {
	values := make([]ListScheduledJobsOperationTypeEnum, 0)
	for _, v := range mappingListScheduledJobsOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledJobsOperationTypeEnumStringValues Enumerates the set of values in String for ListScheduledJobsOperationTypeEnum
func GetListScheduledJobsOperationTypeEnumStringValues() []string {
	return []string{
		"INSTALL_PACKAGES",
		"UPDATE_PACKAGES",
		"REMOVE_PACKAGES",
		"UPDATE_ALL",
		"UPDATE_SECURITY",
		"UPDATE_BUGFIX",
		"UPDATE_ENHANCEMENT",
		"UPDATE_OTHER",
		"UPDATE_KSPLICE_USERSPACE",
		"UPDATE_KSPLICE_KERNEL",
		"MANAGE_MODULE_STREAMS",
		"SWITCH_MODULE_STREAM",
		"ATTACH_SOFTWARE_SOURCES",
		"DETACH_SOFTWARE_SOURCES",
		"SYNC_MANAGEMENT_STATION_MIRROR",
		"PROMOTE_LIFECYCLE",
	}
}

// GetMappingListScheduledJobsOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledJobsOperationTypeEnum(val string) (ListScheduledJobsOperationTypeEnum, bool) {
	enum, ok := mappingListScheduledJobsOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListScheduledJobsScheduleTypeEnum Enum with underlying type: string
type ListScheduledJobsScheduleTypeEnum string

// Set of constants representing the allowable values for ListScheduledJobsScheduleTypeEnum
const (
	ListScheduledJobsScheduleTypeOnetime   ListScheduledJobsScheduleTypeEnum = "ONETIME"
	ListScheduledJobsScheduleTypeRecurring ListScheduledJobsScheduleTypeEnum = "RECURRING"
)

var mappingListScheduledJobsScheduleTypeEnum = map[string]ListScheduledJobsScheduleTypeEnum{
	"ONETIME":   ListScheduledJobsScheduleTypeOnetime,
	"RECURRING": ListScheduledJobsScheduleTypeRecurring,
}

var mappingListScheduledJobsScheduleTypeEnumLowerCase = map[string]ListScheduledJobsScheduleTypeEnum{
	"onetime":   ListScheduledJobsScheduleTypeOnetime,
	"recurring": ListScheduledJobsScheduleTypeRecurring,
}

// GetListScheduledJobsScheduleTypeEnumValues Enumerates the set of values for ListScheduledJobsScheduleTypeEnum
func GetListScheduledJobsScheduleTypeEnumValues() []ListScheduledJobsScheduleTypeEnum {
	values := make([]ListScheduledJobsScheduleTypeEnum, 0)
	for _, v := range mappingListScheduledJobsScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledJobsScheduleTypeEnumStringValues Enumerates the set of values in String for ListScheduledJobsScheduleTypeEnum
func GetListScheduledJobsScheduleTypeEnumStringValues() []string {
	return []string{
		"ONETIME",
		"RECURRING",
	}
}

// GetMappingListScheduledJobsScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledJobsScheduleTypeEnum(val string) (ListScheduledJobsScheduleTypeEnum, bool) {
	enum, ok := mappingListScheduledJobsScheduleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListScheduledJobsSortOrderEnum Enum with underlying type: string
type ListScheduledJobsSortOrderEnum string

// Set of constants representing the allowable values for ListScheduledJobsSortOrderEnum
const (
	ListScheduledJobsSortOrderAsc  ListScheduledJobsSortOrderEnum = "ASC"
	ListScheduledJobsSortOrderDesc ListScheduledJobsSortOrderEnum = "DESC"
)

var mappingListScheduledJobsSortOrderEnum = map[string]ListScheduledJobsSortOrderEnum{
	"ASC":  ListScheduledJobsSortOrderAsc,
	"DESC": ListScheduledJobsSortOrderDesc,
}

var mappingListScheduledJobsSortOrderEnumLowerCase = map[string]ListScheduledJobsSortOrderEnum{
	"asc":  ListScheduledJobsSortOrderAsc,
	"desc": ListScheduledJobsSortOrderDesc,
}

// GetListScheduledJobsSortOrderEnumValues Enumerates the set of values for ListScheduledJobsSortOrderEnum
func GetListScheduledJobsSortOrderEnumValues() []ListScheduledJobsSortOrderEnum {
	values := make([]ListScheduledJobsSortOrderEnum, 0)
	for _, v := range mappingListScheduledJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledJobsSortOrderEnumStringValues Enumerates the set of values in String for ListScheduledJobsSortOrderEnum
func GetListScheduledJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListScheduledJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledJobsSortOrderEnum(val string) (ListScheduledJobsSortOrderEnum, bool) {
	enum, ok := mappingListScheduledJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListScheduledJobsSortByEnum Enum with underlying type: string
type ListScheduledJobsSortByEnum string

// Set of constants representing the allowable values for ListScheduledJobsSortByEnum
const (
	ListScheduledJobsSortByTimecreated ListScheduledJobsSortByEnum = "timeCreated"
	ListScheduledJobsSortByDisplayname ListScheduledJobsSortByEnum = "displayName"
)

var mappingListScheduledJobsSortByEnum = map[string]ListScheduledJobsSortByEnum{
	"timeCreated": ListScheduledJobsSortByTimecreated,
	"displayName": ListScheduledJobsSortByDisplayname,
}

var mappingListScheduledJobsSortByEnumLowerCase = map[string]ListScheduledJobsSortByEnum{
	"timecreated": ListScheduledJobsSortByTimecreated,
	"displayname": ListScheduledJobsSortByDisplayname,
}

// GetListScheduledJobsSortByEnumValues Enumerates the set of values for ListScheduledJobsSortByEnum
func GetListScheduledJobsSortByEnumValues() []ListScheduledJobsSortByEnum {
	values := make([]ListScheduledJobsSortByEnum, 0)
	for _, v := range mappingListScheduledJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledJobsSortByEnumStringValues Enumerates the set of values in String for ListScheduledJobsSortByEnum
func GetListScheduledJobsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListScheduledJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledJobsSortByEnum(val string) (ListScheduledJobsSortByEnum, bool) {
	enum, ok := mappingListScheduledJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
