// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMaintenanceExecutionsRequest wrapper for the ListMaintenanceExecutions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListMaintenanceExecutions.go.html to see an example of how to use ListMaintenanceExecutionsRequest.
type ListMaintenanceExecutionsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	// For list operations, you may provide the tenant [OCID] in this field. When a tenant OCID is provided,
	// it will be validated against the caller's tenant and then treated as tenant scope (compartmentId filtering is not applied).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that match the entire display name given. The match is case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The Database Infrastructure ID.
	InfrastructureId *string `mandatory:"false" contributesTo:"query" name:"infrastructureId"`

	// The type of the target resource.
	TargetResourceType ListMaintenanceExecutionsTargetResourceTypeEnum `mandatory:"false" contributesTo:"query" name:"targetResourceType" omitEmpty:"true"`

	// The maintenance type.
	MaintenanceType ListMaintenanceExecutionsMaintenanceTypeEnum `mandatory:"false" contributesTo:"query" name:"maintenanceType" omitEmpty:"true"`

	// The maintenance run OCID.
	MaintenanceRunId *string `mandatory:"false" contributesTo:"query" name:"maintenanceRunId"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState ListMaintenanceExecutionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sub-type of the maintenance run.
	MaintenanceSubtype ListMaintenanceExecutionsMaintenanceSubtypeEnum `mandatory:"false" contributesTo:"query" name:"maintenanceSubtype" omitEmpty:"true"`

	// Filter maintenance run for before given time.
	TimeAcceptedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeAcceptedLessThanOrEqualTo"`

	// Filter maintenance run for after given time.
	TimeAcceptedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeAcceptedGreaterThanOrEqualTo"`

	// The maintenance execution type.
	Type ListMaintenanceExecutionsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The maximum number of items to return in a page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which you want to start retrieving results. This token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field by which you want to sort. You can provide only one type of sort order. The default order for `timeCreated` is descending. The default order for `displayName` is ascending. If no value is specified, then `timeCreated` is the default.
	// When listing software images within the same `version`, using `sortBy=buildIdentifier` is recommended. `buildIdentifier` is a monotonically increasing, time-ordered string marker (yyyy-mm-dd-hh:mm:ss) stored with the image.
	SortBy ListMaintenanceExecutionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order that you want to use, which is either `ASC` or `DESC`.
	SortOrder ListMaintenanceExecutionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaintenanceExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaintenanceExecutionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaintenanceExecutionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaintenanceExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaintenanceExecutionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaintenanceExecutionsTargetResourceTypeEnum(string(request.TargetResourceType)); !ok && request.TargetResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetResourceType: %s. Supported values are: %s.", request.TargetResourceType, strings.Join(GetListMaintenanceExecutionsTargetResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceExecutionsMaintenanceTypeEnum(string(request.MaintenanceType)); !ok && request.MaintenanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceType: %s. Supported values are: %s.", request.MaintenanceType, strings.Join(GetListMaintenanceExecutionsMaintenanceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceExecutionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListMaintenanceExecutionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceExecutionsMaintenanceSubtypeEnum(string(request.MaintenanceSubtype)); !ok && request.MaintenanceSubtype != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceSubtype: %s. Supported values are: %s.", request.MaintenanceSubtype, strings.Join(GetListMaintenanceExecutionsMaintenanceSubtypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceExecutionsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListMaintenanceExecutionsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceExecutionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaintenanceExecutionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceExecutionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaintenanceExecutionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaintenanceExecutionsResponse wrapper for the ListMaintenanceExecutions operation
type ListMaintenanceExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaintenanceExecutionCollection instances
	MaintenanceExecutionCollection `presentIn:"body"`

	// Unique identifier assigned by Oracle for the request. If you need to contact
	// Oracle about a particular request, then please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then it can mean that a partial list was returned. To obtain the next batch of items, include this value as the `page` parameter for your next GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMaintenanceExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaintenanceExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaintenanceExecutionsTargetResourceTypeEnum Enum with underlying type: string
type ListMaintenanceExecutionsTargetResourceTypeEnum string

// Set of constants representing the allowable values for ListMaintenanceExecutionsTargetResourceTypeEnum
const (
	ListMaintenanceExecutionsTargetResourceTypeDbCcInfrastructure ListMaintenanceExecutionsTargetResourceTypeEnum = "DB_CC_INFRASTRUCTURE"
)

var mappingListMaintenanceExecutionsTargetResourceTypeEnum = map[string]ListMaintenanceExecutionsTargetResourceTypeEnum{
	"DB_CC_INFRASTRUCTURE": ListMaintenanceExecutionsTargetResourceTypeDbCcInfrastructure,
}

var mappingListMaintenanceExecutionsTargetResourceTypeEnumLowerCase = map[string]ListMaintenanceExecutionsTargetResourceTypeEnum{
	"db_cc_infrastructure": ListMaintenanceExecutionsTargetResourceTypeDbCcInfrastructure,
}

// GetListMaintenanceExecutionsTargetResourceTypeEnumValues Enumerates the set of values for ListMaintenanceExecutionsTargetResourceTypeEnum
func GetListMaintenanceExecutionsTargetResourceTypeEnumValues() []ListMaintenanceExecutionsTargetResourceTypeEnum {
	values := make([]ListMaintenanceExecutionsTargetResourceTypeEnum, 0)
	for _, v := range mappingListMaintenanceExecutionsTargetResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceExecutionsTargetResourceTypeEnumStringValues Enumerates the set of values in String for ListMaintenanceExecutionsTargetResourceTypeEnum
func GetListMaintenanceExecutionsTargetResourceTypeEnumStringValues() []string {
	return []string{
		"DB_CC_INFRASTRUCTURE",
	}
}

// GetMappingListMaintenanceExecutionsTargetResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceExecutionsTargetResourceTypeEnum(val string) (ListMaintenanceExecutionsTargetResourceTypeEnum, bool) {
	enum, ok := mappingListMaintenanceExecutionsTargetResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceExecutionsMaintenanceTypeEnum Enum with underlying type: string
type ListMaintenanceExecutionsMaintenanceTypeEnum string

// Set of constants representing the allowable values for ListMaintenanceExecutionsMaintenanceTypeEnum
const (
	ListMaintenanceExecutionsMaintenanceTypePlanned   ListMaintenanceExecutionsMaintenanceTypeEnum = "PLANNED"
	ListMaintenanceExecutionsMaintenanceTypeUnplanned ListMaintenanceExecutionsMaintenanceTypeEnum = "UNPLANNED"
)

var mappingListMaintenanceExecutionsMaintenanceTypeEnum = map[string]ListMaintenanceExecutionsMaintenanceTypeEnum{
	"PLANNED":   ListMaintenanceExecutionsMaintenanceTypePlanned,
	"UNPLANNED": ListMaintenanceExecutionsMaintenanceTypeUnplanned,
}

var mappingListMaintenanceExecutionsMaintenanceTypeEnumLowerCase = map[string]ListMaintenanceExecutionsMaintenanceTypeEnum{
	"planned":   ListMaintenanceExecutionsMaintenanceTypePlanned,
	"unplanned": ListMaintenanceExecutionsMaintenanceTypeUnplanned,
}

// GetListMaintenanceExecutionsMaintenanceTypeEnumValues Enumerates the set of values for ListMaintenanceExecutionsMaintenanceTypeEnum
func GetListMaintenanceExecutionsMaintenanceTypeEnumValues() []ListMaintenanceExecutionsMaintenanceTypeEnum {
	values := make([]ListMaintenanceExecutionsMaintenanceTypeEnum, 0)
	for _, v := range mappingListMaintenanceExecutionsMaintenanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceExecutionsMaintenanceTypeEnumStringValues Enumerates the set of values in String for ListMaintenanceExecutionsMaintenanceTypeEnum
func GetListMaintenanceExecutionsMaintenanceTypeEnumStringValues() []string {
	return []string{
		"PLANNED",
		"UNPLANNED",
	}
}

// GetMappingListMaintenanceExecutionsMaintenanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceExecutionsMaintenanceTypeEnum(val string) (ListMaintenanceExecutionsMaintenanceTypeEnum, bool) {
	enum, ok := mappingListMaintenanceExecutionsMaintenanceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceExecutionsLifecycleStateEnum Enum with underlying type: string
type ListMaintenanceExecutionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListMaintenanceExecutionsLifecycleStateEnum
const (
	ListMaintenanceExecutionsLifecycleStateCreating       ListMaintenanceExecutionsLifecycleStateEnum = "CREATING"
	ListMaintenanceExecutionsLifecycleStateScheduled      ListMaintenanceExecutionsLifecycleStateEnum = "SCHEDULED"
	ListMaintenanceExecutionsLifecycleStateInProgress     ListMaintenanceExecutionsLifecycleStateEnum = "IN_PROGRESS"
	ListMaintenanceExecutionsLifecycleStateSucceeded      ListMaintenanceExecutionsLifecycleStateEnum = "SUCCEEDED"
	ListMaintenanceExecutionsLifecycleStateSkipped        ListMaintenanceExecutionsLifecycleStateEnum = "SKIPPED"
	ListMaintenanceExecutionsLifecycleStateFailed         ListMaintenanceExecutionsLifecycleStateEnum = "FAILED"
	ListMaintenanceExecutionsLifecycleStateUpdating       ListMaintenanceExecutionsLifecycleStateEnum = "UPDATING"
	ListMaintenanceExecutionsLifecycleStateDeleting       ListMaintenanceExecutionsLifecycleStateEnum = "DELETING"
	ListMaintenanceExecutionsLifecycleStateDeleted        ListMaintenanceExecutionsLifecycleStateEnum = "DELETED"
	ListMaintenanceExecutionsLifecycleStateCanceled       ListMaintenanceExecutionsLifecycleStateEnum = "CANCELED"
	ListMaintenanceExecutionsLifecycleStatePartialSuccess ListMaintenanceExecutionsLifecycleStateEnum = "PARTIAL_SUCCESS"
)

var mappingListMaintenanceExecutionsLifecycleStateEnum = map[string]ListMaintenanceExecutionsLifecycleStateEnum{
	"CREATING":        ListMaintenanceExecutionsLifecycleStateCreating,
	"SCHEDULED":       ListMaintenanceExecutionsLifecycleStateScheduled,
	"IN_PROGRESS":     ListMaintenanceExecutionsLifecycleStateInProgress,
	"SUCCEEDED":       ListMaintenanceExecutionsLifecycleStateSucceeded,
	"SKIPPED":         ListMaintenanceExecutionsLifecycleStateSkipped,
	"FAILED":          ListMaintenanceExecutionsLifecycleStateFailed,
	"UPDATING":        ListMaintenanceExecutionsLifecycleStateUpdating,
	"DELETING":        ListMaintenanceExecutionsLifecycleStateDeleting,
	"DELETED":         ListMaintenanceExecutionsLifecycleStateDeleted,
	"CANCELED":        ListMaintenanceExecutionsLifecycleStateCanceled,
	"PARTIAL_SUCCESS": ListMaintenanceExecutionsLifecycleStatePartialSuccess,
}

var mappingListMaintenanceExecutionsLifecycleStateEnumLowerCase = map[string]ListMaintenanceExecutionsLifecycleStateEnum{
	"creating":        ListMaintenanceExecutionsLifecycleStateCreating,
	"scheduled":       ListMaintenanceExecutionsLifecycleStateScheduled,
	"in_progress":     ListMaintenanceExecutionsLifecycleStateInProgress,
	"succeeded":       ListMaintenanceExecutionsLifecycleStateSucceeded,
	"skipped":         ListMaintenanceExecutionsLifecycleStateSkipped,
	"failed":          ListMaintenanceExecutionsLifecycleStateFailed,
	"updating":        ListMaintenanceExecutionsLifecycleStateUpdating,
	"deleting":        ListMaintenanceExecutionsLifecycleStateDeleting,
	"deleted":         ListMaintenanceExecutionsLifecycleStateDeleted,
	"canceled":        ListMaintenanceExecutionsLifecycleStateCanceled,
	"partial_success": ListMaintenanceExecutionsLifecycleStatePartialSuccess,
}

// GetListMaintenanceExecutionsLifecycleStateEnumValues Enumerates the set of values for ListMaintenanceExecutionsLifecycleStateEnum
func GetListMaintenanceExecutionsLifecycleStateEnumValues() []ListMaintenanceExecutionsLifecycleStateEnum {
	values := make([]ListMaintenanceExecutionsLifecycleStateEnum, 0)
	for _, v := range mappingListMaintenanceExecutionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceExecutionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListMaintenanceExecutionsLifecycleStateEnum
func GetListMaintenanceExecutionsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"SCHEDULED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"SKIPPED",
		"FAILED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"CANCELED",
		"PARTIAL_SUCCESS",
	}
}

// GetMappingListMaintenanceExecutionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceExecutionsLifecycleStateEnum(val string) (ListMaintenanceExecutionsLifecycleStateEnum, bool) {
	enum, ok := mappingListMaintenanceExecutionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceExecutionsMaintenanceSubtypeEnum Enum with underlying type: string
type ListMaintenanceExecutionsMaintenanceSubtypeEnum string

// Set of constants representing the allowable values for ListMaintenanceExecutionsMaintenanceSubtypeEnum
const (
	ListMaintenanceExecutionsMaintenanceSubtypeYearly            ListMaintenanceExecutionsMaintenanceSubtypeEnum = "YEARLY"
	ListMaintenanceExecutionsMaintenanceSubtypeHalfyearly        ListMaintenanceExecutionsMaintenanceSubtypeEnum = "HALFYEARLY"
	ListMaintenanceExecutionsMaintenanceSubtypeQuarterly         ListMaintenanceExecutionsMaintenanceSubtypeEnum = "QUARTERLY"
	ListMaintenanceExecutionsMaintenanceSubtypeMonthly           ListMaintenanceExecutionsMaintenanceSubtypeEnum = "MONTHLY"
	ListMaintenanceExecutionsMaintenanceSubtypeDaily             ListMaintenanceExecutionsMaintenanceSubtypeEnum = "DAILY"
	ListMaintenanceExecutionsMaintenanceSubtypeHardware          ListMaintenanceExecutionsMaintenanceSubtypeEnum = "HARDWARE"
	ListMaintenanceExecutionsMaintenanceSubtypeCritical          ListMaintenanceExecutionsMaintenanceSubtypeEnum = "CRITICAL"
	ListMaintenanceExecutionsMaintenanceSubtypeInfraUpdate       ListMaintenanceExecutionsMaintenanceSubtypeEnum = "INFRA_UPDATE"
	ListMaintenanceExecutionsMaintenanceSubtypeCpsServicesUpdate ListMaintenanceExecutionsMaintenanceSubtypeEnum = "CPS_SERVICES_UPDATE"
	ListMaintenanceExecutionsMaintenanceSubtypeCpsVmUpdate       ListMaintenanceExecutionsMaintenanceSubtypeEnum = "CPS_VM_UPDATE"
	ListMaintenanceExecutionsMaintenanceSubtypeSecurityMonthly   ListMaintenanceExecutionsMaintenanceSubtypeEnum = "SECURITY_MONTHLY"
)

var mappingListMaintenanceExecutionsMaintenanceSubtypeEnum = map[string]ListMaintenanceExecutionsMaintenanceSubtypeEnum{
	"YEARLY":              ListMaintenanceExecutionsMaintenanceSubtypeYearly,
	"HALFYEARLY":          ListMaintenanceExecutionsMaintenanceSubtypeHalfyearly,
	"QUARTERLY":           ListMaintenanceExecutionsMaintenanceSubtypeQuarterly,
	"MONTHLY":             ListMaintenanceExecutionsMaintenanceSubtypeMonthly,
	"DAILY":               ListMaintenanceExecutionsMaintenanceSubtypeDaily,
	"HARDWARE":            ListMaintenanceExecutionsMaintenanceSubtypeHardware,
	"CRITICAL":            ListMaintenanceExecutionsMaintenanceSubtypeCritical,
	"INFRA_UPDATE":        ListMaintenanceExecutionsMaintenanceSubtypeInfraUpdate,
	"CPS_SERVICES_UPDATE": ListMaintenanceExecutionsMaintenanceSubtypeCpsServicesUpdate,
	"CPS_VM_UPDATE":       ListMaintenanceExecutionsMaintenanceSubtypeCpsVmUpdate,
	"SECURITY_MONTHLY":    ListMaintenanceExecutionsMaintenanceSubtypeSecurityMonthly,
}

var mappingListMaintenanceExecutionsMaintenanceSubtypeEnumLowerCase = map[string]ListMaintenanceExecutionsMaintenanceSubtypeEnum{
	"yearly":              ListMaintenanceExecutionsMaintenanceSubtypeYearly,
	"halfyearly":          ListMaintenanceExecutionsMaintenanceSubtypeHalfyearly,
	"quarterly":           ListMaintenanceExecutionsMaintenanceSubtypeQuarterly,
	"monthly":             ListMaintenanceExecutionsMaintenanceSubtypeMonthly,
	"daily":               ListMaintenanceExecutionsMaintenanceSubtypeDaily,
	"hardware":            ListMaintenanceExecutionsMaintenanceSubtypeHardware,
	"critical":            ListMaintenanceExecutionsMaintenanceSubtypeCritical,
	"infra_update":        ListMaintenanceExecutionsMaintenanceSubtypeInfraUpdate,
	"cps_services_update": ListMaintenanceExecutionsMaintenanceSubtypeCpsServicesUpdate,
	"cps_vm_update":       ListMaintenanceExecutionsMaintenanceSubtypeCpsVmUpdate,
	"security_monthly":    ListMaintenanceExecutionsMaintenanceSubtypeSecurityMonthly,
}

// GetListMaintenanceExecutionsMaintenanceSubtypeEnumValues Enumerates the set of values for ListMaintenanceExecutionsMaintenanceSubtypeEnum
func GetListMaintenanceExecutionsMaintenanceSubtypeEnumValues() []ListMaintenanceExecutionsMaintenanceSubtypeEnum {
	values := make([]ListMaintenanceExecutionsMaintenanceSubtypeEnum, 0)
	for _, v := range mappingListMaintenanceExecutionsMaintenanceSubtypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceExecutionsMaintenanceSubtypeEnumStringValues Enumerates the set of values in String for ListMaintenanceExecutionsMaintenanceSubtypeEnum
func GetListMaintenanceExecutionsMaintenanceSubtypeEnumStringValues() []string {
	return []string{
		"YEARLY",
		"HALFYEARLY",
		"QUARTERLY",
		"MONTHLY",
		"DAILY",
		"HARDWARE",
		"CRITICAL",
		"INFRA_UPDATE",
		"CPS_SERVICES_UPDATE",
		"CPS_VM_UPDATE",
		"SECURITY_MONTHLY",
	}
}

// GetMappingListMaintenanceExecutionsMaintenanceSubtypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceExecutionsMaintenanceSubtypeEnum(val string) (ListMaintenanceExecutionsMaintenanceSubtypeEnum, bool) {
	enum, ok := mappingListMaintenanceExecutionsMaintenanceSubtypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceExecutionsTypeEnum Enum with underlying type: string
type ListMaintenanceExecutionsTypeEnum string

// Set of constants representing the allowable values for ListMaintenanceExecutionsTypeEnum
const (
	ListMaintenanceExecutionsTypeNotify  ListMaintenanceExecutionsTypeEnum = "NOTIFY"
	ListMaintenanceExecutionsTypeExecute ListMaintenanceExecutionsTypeEnum = "EXECUTE"
)

var mappingListMaintenanceExecutionsTypeEnum = map[string]ListMaintenanceExecutionsTypeEnum{
	"NOTIFY":  ListMaintenanceExecutionsTypeNotify,
	"EXECUTE": ListMaintenanceExecutionsTypeExecute,
}

var mappingListMaintenanceExecutionsTypeEnumLowerCase = map[string]ListMaintenanceExecutionsTypeEnum{
	"notify":  ListMaintenanceExecutionsTypeNotify,
	"execute": ListMaintenanceExecutionsTypeExecute,
}

// GetListMaintenanceExecutionsTypeEnumValues Enumerates the set of values for ListMaintenanceExecutionsTypeEnum
func GetListMaintenanceExecutionsTypeEnumValues() []ListMaintenanceExecutionsTypeEnum {
	values := make([]ListMaintenanceExecutionsTypeEnum, 0)
	for _, v := range mappingListMaintenanceExecutionsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceExecutionsTypeEnumStringValues Enumerates the set of values in String for ListMaintenanceExecutionsTypeEnum
func GetListMaintenanceExecutionsTypeEnumStringValues() []string {
	return []string{
		"NOTIFY",
		"EXECUTE",
	}
}

// GetMappingListMaintenanceExecutionsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceExecutionsTypeEnum(val string) (ListMaintenanceExecutionsTypeEnum, bool) {
	enum, ok := mappingListMaintenanceExecutionsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceExecutionsSortByEnum Enum with underlying type: string
type ListMaintenanceExecutionsSortByEnum string

// Set of constants representing the allowable values for ListMaintenanceExecutionsSortByEnum
const (
	ListMaintenanceExecutionsSortByTimecreated     ListMaintenanceExecutionsSortByEnum = "timeCreated"
	ListMaintenanceExecutionsSortByDisplayname     ListMaintenanceExecutionsSortByEnum = "displayName"
	ListMaintenanceExecutionsSortByBuildidentifier ListMaintenanceExecutionsSortByEnum = "buildIdentifier"
)

var mappingListMaintenanceExecutionsSortByEnum = map[string]ListMaintenanceExecutionsSortByEnum{
	"timeCreated":     ListMaintenanceExecutionsSortByTimecreated,
	"displayName":     ListMaintenanceExecutionsSortByDisplayname,
	"buildIdentifier": ListMaintenanceExecutionsSortByBuildidentifier,
}

var mappingListMaintenanceExecutionsSortByEnumLowerCase = map[string]ListMaintenanceExecutionsSortByEnum{
	"timecreated":     ListMaintenanceExecutionsSortByTimecreated,
	"displayname":     ListMaintenanceExecutionsSortByDisplayname,
	"buildidentifier": ListMaintenanceExecutionsSortByBuildidentifier,
}

// GetListMaintenanceExecutionsSortByEnumValues Enumerates the set of values for ListMaintenanceExecutionsSortByEnum
func GetListMaintenanceExecutionsSortByEnumValues() []ListMaintenanceExecutionsSortByEnum {
	values := make([]ListMaintenanceExecutionsSortByEnum, 0)
	for _, v := range mappingListMaintenanceExecutionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceExecutionsSortByEnumStringValues Enumerates the set of values in String for ListMaintenanceExecutionsSortByEnum
func GetListMaintenanceExecutionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"buildIdentifier",
	}
}

// GetMappingListMaintenanceExecutionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceExecutionsSortByEnum(val string) (ListMaintenanceExecutionsSortByEnum, bool) {
	enum, ok := mappingListMaintenanceExecutionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceExecutionsSortOrderEnum Enum with underlying type: string
type ListMaintenanceExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListMaintenanceExecutionsSortOrderEnum
const (
	ListMaintenanceExecutionsSortOrderAsc  ListMaintenanceExecutionsSortOrderEnum = "ASC"
	ListMaintenanceExecutionsSortOrderDesc ListMaintenanceExecutionsSortOrderEnum = "DESC"
)

var mappingListMaintenanceExecutionsSortOrderEnum = map[string]ListMaintenanceExecutionsSortOrderEnum{
	"ASC":  ListMaintenanceExecutionsSortOrderAsc,
	"DESC": ListMaintenanceExecutionsSortOrderDesc,
}

var mappingListMaintenanceExecutionsSortOrderEnumLowerCase = map[string]ListMaintenanceExecutionsSortOrderEnum{
	"asc":  ListMaintenanceExecutionsSortOrderAsc,
	"desc": ListMaintenanceExecutionsSortOrderDesc,
}

// GetListMaintenanceExecutionsSortOrderEnumValues Enumerates the set of values for ListMaintenanceExecutionsSortOrderEnum
func GetListMaintenanceExecutionsSortOrderEnumValues() []ListMaintenanceExecutionsSortOrderEnum {
	values := make([]ListMaintenanceExecutionsSortOrderEnum, 0)
	for _, v := range mappingListMaintenanceExecutionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceExecutionsSortOrderEnumStringValues Enumerates the set of values in String for ListMaintenanceExecutionsSortOrderEnum
func GetListMaintenanceExecutionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaintenanceExecutionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceExecutionsSortOrderEnum(val string) (ListMaintenanceExecutionsSortOrderEnum, bool) {
	enum, ok := mappingListMaintenanceExecutionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
