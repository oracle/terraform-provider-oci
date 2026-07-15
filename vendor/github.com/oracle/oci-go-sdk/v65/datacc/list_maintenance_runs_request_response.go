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

// ListMaintenanceRunsRequest wrapper for the ListMaintenanceRuns operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListMaintenanceRuns.go.html to see an example of how to use ListMaintenanceRunsRequest.
type ListMaintenanceRunsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	// For list operations, you may provide the tenant [OCID] in this field. When a tenant OCID is provided,
	// it will be validated against the caller's tenant and then treated as tenant scope (compartmentId filtering is not applied).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that match the entire display name given. The match is case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The Database Infrastructure ID.
	InfrastructureId *string `mandatory:"false" contributesTo:"query" name:"infrastructureId"`

	// The type of the target resource.
	TargetResourceType ListMaintenanceRunsTargetResourceTypeEnum `mandatory:"false" contributesTo:"query" name:"targetResourceType" omitEmpty:"true"`

	// The maintenance type.
	MaintenanceType ListMaintenanceRunsMaintenanceTypeEnum `mandatory:"false" contributesTo:"query" name:"maintenanceType" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState ListMaintenanceRunsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sub-type of the maintenance run.
	MaintenanceSubtype ListMaintenanceRunsMaintenanceSubtypeEnum `mandatory:"false" contributesTo:"query" name:"maintenanceSubtype" omitEmpty:"true"`

	// Filter maintenance run for before given time.
	TimeAcceptedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeAcceptedLessThanOrEqualTo"`

	// Filter maintenance run for after given time.
	TimeAcceptedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeAcceptedGreaterThanOrEqualTo"`

	// The maximum number of items to return in a page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which you want to start retrieving results. This token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field by which you want to sort. You can provide only one type of sort order. The default order for `timeCreated` is descending. The default order for `displayName` is ascending. If no value is specified, then `timeCreated` is the default.
	// When listing software images within the same `version`, using `sortBy=buildIdentifier` is recommended. `buildIdentifier` is a monotonically increasing, time-ordered string marker (yyyy-mm-dd-hh:mm:ss) stored with the image.
	SortBy ListMaintenanceRunsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order that you want to use, which is either `ASC` or `DESC`.
	SortOrder ListMaintenanceRunsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request identifier.
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
	if _, ok := GetMappingListMaintenanceRunsTargetResourceTypeEnum(string(request.TargetResourceType)); !ok && request.TargetResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetResourceType: %s. Supported values are: %s.", request.TargetResourceType, strings.Join(GetListMaintenanceRunsTargetResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceRunsMaintenanceTypeEnum(string(request.MaintenanceType)); !ok && request.MaintenanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceType: %s. Supported values are: %s.", request.MaintenanceType, strings.Join(GetListMaintenanceRunsMaintenanceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceRunsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListMaintenanceRunsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceRunsMaintenanceSubtypeEnum(string(request.MaintenanceSubtype)); !ok && request.MaintenanceSubtype != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceSubtype: %s. Supported values are: %s.", request.MaintenanceSubtype, strings.Join(GetListMaintenanceRunsMaintenanceSubtypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceRunsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaintenanceRunsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceRunsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaintenanceRunsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaintenanceRunsResponse wrapper for the ListMaintenanceRuns operation
type ListMaintenanceRunsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaintenanceRunCollection instances
	MaintenanceRunCollection `presentIn:"body"`

	// Unique identifier assigned by Oracle for the request. If you need to contact
	// Oracle about a particular request, then please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then it can mean that a partial list was returned. To obtain the next batch of items, include this value as the `page` parameter for your next GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMaintenanceRunsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaintenanceRunsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaintenanceRunsTargetResourceTypeEnum Enum with underlying type: string
type ListMaintenanceRunsTargetResourceTypeEnum string

// Set of constants representing the allowable values for ListMaintenanceRunsTargetResourceTypeEnum
const (
	ListMaintenanceRunsTargetResourceTypeDbCcInfrastructure ListMaintenanceRunsTargetResourceTypeEnum = "DB_CC_INFRASTRUCTURE"
)

var mappingListMaintenanceRunsTargetResourceTypeEnum = map[string]ListMaintenanceRunsTargetResourceTypeEnum{
	"DB_CC_INFRASTRUCTURE": ListMaintenanceRunsTargetResourceTypeDbCcInfrastructure,
}

var mappingListMaintenanceRunsTargetResourceTypeEnumLowerCase = map[string]ListMaintenanceRunsTargetResourceTypeEnum{
	"db_cc_infrastructure": ListMaintenanceRunsTargetResourceTypeDbCcInfrastructure,
}

// GetListMaintenanceRunsTargetResourceTypeEnumValues Enumerates the set of values for ListMaintenanceRunsTargetResourceTypeEnum
func GetListMaintenanceRunsTargetResourceTypeEnumValues() []ListMaintenanceRunsTargetResourceTypeEnum {
	values := make([]ListMaintenanceRunsTargetResourceTypeEnum, 0)
	for _, v := range mappingListMaintenanceRunsTargetResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceRunsTargetResourceTypeEnumStringValues Enumerates the set of values in String for ListMaintenanceRunsTargetResourceTypeEnum
func GetListMaintenanceRunsTargetResourceTypeEnumStringValues() []string {
	return []string{
		"DB_CC_INFRASTRUCTURE",
	}
}

// GetMappingListMaintenanceRunsTargetResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceRunsTargetResourceTypeEnum(val string) (ListMaintenanceRunsTargetResourceTypeEnum, bool) {
	enum, ok := mappingListMaintenanceRunsTargetResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceRunsMaintenanceTypeEnum Enum with underlying type: string
type ListMaintenanceRunsMaintenanceTypeEnum string

// Set of constants representing the allowable values for ListMaintenanceRunsMaintenanceTypeEnum
const (
	ListMaintenanceRunsMaintenanceTypePlanned   ListMaintenanceRunsMaintenanceTypeEnum = "PLANNED"
	ListMaintenanceRunsMaintenanceTypeUnplanned ListMaintenanceRunsMaintenanceTypeEnum = "UNPLANNED"
)

var mappingListMaintenanceRunsMaintenanceTypeEnum = map[string]ListMaintenanceRunsMaintenanceTypeEnum{
	"PLANNED":   ListMaintenanceRunsMaintenanceTypePlanned,
	"UNPLANNED": ListMaintenanceRunsMaintenanceTypeUnplanned,
}

var mappingListMaintenanceRunsMaintenanceTypeEnumLowerCase = map[string]ListMaintenanceRunsMaintenanceTypeEnum{
	"planned":   ListMaintenanceRunsMaintenanceTypePlanned,
	"unplanned": ListMaintenanceRunsMaintenanceTypeUnplanned,
}

// GetListMaintenanceRunsMaintenanceTypeEnumValues Enumerates the set of values for ListMaintenanceRunsMaintenanceTypeEnum
func GetListMaintenanceRunsMaintenanceTypeEnumValues() []ListMaintenanceRunsMaintenanceTypeEnum {
	values := make([]ListMaintenanceRunsMaintenanceTypeEnum, 0)
	for _, v := range mappingListMaintenanceRunsMaintenanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceRunsMaintenanceTypeEnumStringValues Enumerates the set of values in String for ListMaintenanceRunsMaintenanceTypeEnum
func GetListMaintenanceRunsMaintenanceTypeEnumStringValues() []string {
	return []string{
		"PLANNED",
		"UNPLANNED",
	}
}

// GetMappingListMaintenanceRunsMaintenanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceRunsMaintenanceTypeEnum(val string) (ListMaintenanceRunsMaintenanceTypeEnum, bool) {
	enum, ok := mappingListMaintenanceRunsMaintenanceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceRunsLifecycleStateEnum Enum with underlying type: string
type ListMaintenanceRunsLifecycleStateEnum string

// Set of constants representing the allowable values for ListMaintenanceRunsLifecycleStateEnum
const (
	ListMaintenanceRunsLifecycleStateCreating       ListMaintenanceRunsLifecycleStateEnum = "CREATING"
	ListMaintenanceRunsLifecycleStateScheduled      ListMaintenanceRunsLifecycleStateEnum = "SCHEDULED"
	ListMaintenanceRunsLifecycleStateInProgress     ListMaintenanceRunsLifecycleStateEnum = "IN_PROGRESS"
	ListMaintenanceRunsLifecycleStateSucceeded      ListMaintenanceRunsLifecycleStateEnum = "SUCCEEDED"
	ListMaintenanceRunsLifecycleStateSkipped        ListMaintenanceRunsLifecycleStateEnum = "SKIPPED"
	ListMaintenanceRunsLifecycleStateFailed         ListMaintenanceRunsLifecycleStateEnum = "FAILED"
	ListMaintenanceRunsLifecycleStateUpdating       ListMaintenanceRunsLifecycleStateEnum = "UPDATING"
	ListMaintenanceRunsLifecycleStateDeleting       ListMaintenanceRunsLifecycleStateEnum = "DELETING"
	ListMaintenanceRunsLifecycleStateDeleted        ListMaintenanceRunsLifecycleStateEnum = "DELETED"
	ListMaintenanceRunsLifecycleStateCanceled       ListMaintenanceRunsLifecycleStateEnum = "CANCELED"
	ListMaintenanceRunsLifecycleStatePartialSuccess ListMaintenanceRunsLifecycleStateEnum = "PARTIAL_SUCCESS"
)

var mappingListMaintenanceRunsLifecycleStateEnum = map[string]ListMaintenanceRunsLifecycleStateEnum{
	"CREATING":        ListMaintenanceRunsLifecycleStateCreating,
	"SCHEDULED":       ListMaintenanceRunsLifecycleStateScheduled,
	"IN_PROGRESS":     ListMaintenanceRunsLifecycleStateInProgress,
	"SUCCEEDED":       ListMaintenanceRunsLifecycleStateSucceeded,
	"SKIPPED":         ListMaintenanceRunsLifecycleStateSkipped,
	"FAILED":          ListMaintenanceRunsLifecycleStateFailed,
	"UPDATING":        ListMaintenanceRunsLifecycleStateUpdating,
	"DELETING":        ListMaintenanceRunsLifecycleStateDeleting,
	"DELETED":         ListMaintenanceRunsLifecycleStateDeleted,
	"CANCELED":        ListMaintenanceRunsLifecycleStateCanceled,
	"PARTIAL_SUCCESS": ListMaintenanceRunsLifecycleStatePartialSuccess,
}

var mappingListMaintenanceRunsLifecycleStateEnumLowerCase = map[string]ListMaintenanceRunsLifecycleStateEnum{
	"creating":        ListMaintenanceRunsLifecycleStateCreating,
	"scheduled":       ListMaintenanceRunsLifecycleStateScheduled,
	"in_progress":     ListMaintenanceRunsLifecycleStateInProgress,
	"succeeded":       ListMaintenanceRunsLifecycleStateSucceeded,
	"skipped":         ListMaintenanceRunsLifecycleStateSkipped,
	"failed":          ListMaintenanceRunsLifecycleStateFailed,
	"updating":        ListMaintenanceRunsLifecycleStateUpdating,
	"deleting":        ListMaintenanceRunsLifecycleStateDeleting,
	"deleted":         ListMaintenanceRunsLifecycleStateDeleted,
	"canceled":        ListMaintenanceRunsLifecycleStateCanceled,
	"partial_success": ListMaintenanceRunsLifecycleStatePartialSuccess,
}

// GetListMaintenanceRunsLifecycleStateEnumValues Enumerates the set of values for ListMaintenanceRunsLifecycleStateEnum
func GetListMaintenanceRunsLifecycleStateEnumValues() []ListMaintenanceRunsLifecycleStateEnum {
	values := make([]ListMaintenanceRunsLifecycleStateEnum, 0)
	for _, v := range mappingListMaintenanceRunsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceRunsLifecycleStateEnumStringValues Enumerates the set of values in String for ListMaintenanceRunsLifecycleStateEnum
func GetListMaintenanceRunsLifecycleStateEnumStringValues() []string {
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

// GetMappingListMaintenanceRunsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceRunsLifecycleStateEnum(val string) (ListMaintenanceRunsLifecycleStateEnum, bool) {
	enum, ok := mappingListMaintenanceRunsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceRunsMaintenanceSubtypeEnum Enum with underlying type: string
type ListMaintenanceRunsMaintenanceSubtypeEnum string

// Set of constants representing the allowable values for ListMaintenanceRunsMaintenanceSubtypeEnum
const (
	ListMaintenanceRunsMaintenanceSubtypeYearly            ListMaintenanceRunsMaintenanceSubtypeEnum = "YEARLY"
	ListMaintenanceRunsMaintenanceSubtypeHalfyearly        ListMaintenanceRunsMaintenanceSubtypeEnum = "HALFYEARLY"
	ListMaintenanceRunsMaintenanceSubtypeQuarterly         ListMaintenanceRunsMaintenanceSubtypeEnum = "QUARTERLY"
	ListMaintenanceRunsMaintenanceSubtypeMonthly           ListMaintenanceRunsMaintenanceSubtypeEnum = "MONTHLY"
	ListMaintenanceRunsMaintenanceSubtypeDaily             ListMaintenanceRunsMaintenanceSubtypeEnum = "DAILY"
	ListMaintenanceRunsMaintenanceSubtypeHardware          ListMaintenanceRunsMaintenanceSubtypeEnum = "HARDWARE"
	ListMaintenanceRunsMaintenanceSubtypeCritical          ListMaintenanceRunsMaintenanceSubtypeEnum = "CRITICAL"
	ListMaintenanceRunsMaintenanceSubtypeInfraUpdate       ListMaintenanceRunsMaintenanceSubtypeEnum = "INFRA_UPDATE"
	ListMaintenanceRunsMaintenanceSubtypeCpsServicesUpdate ListMaintenanceRunsMaintenanceSubtypeEnum = "CPS_SERVICES_UPDATE"
	ListMaintenanceRunsMaintenanceSubtypeCpsVmUpdate       ListMaintenanceRunsMaintenanceSubtypeEnum = "CPS_VM_UPDATE"
	ListMaintenanceRunsMaintenanceSubtypeSecurityMonthly   ListMaintenanceRunsMaintenanceSubtypeEnum = "SECURITY_MONTHLY"
)

var mappingListMaintenanceRunsMaintenanceSubtypeEnum = map[string]ListMaintenanceRunsMaintenanceSubtypeEnum{
	"YEARLY":              ListMaintenanceRunsMaintenanceSubtypeYearly,
	"HALFYEARLY":          ListMaintenanceRunsMaintenanceSubtypeHalfyearly,
	"QUARTERLY":           ListMaintenanceRunsMaintenanceSubtypeQuarterly,
	"MONTHLY":             ListMaintenanceRunsMaintenanceSubtypeMonthly,
	"DAILY":               ListMaintenanceRunsMaintenanceSubtypeDaily,
	"HARDWARE":            ListMaintenanceRunsMaintenanceSubtypeHardware,
	"CRITICAL":            ListMaintenanceRunsMaintenanceSubtypeCritical,
	"INFRA_UPDATE":        ListMaintenanceRunsMaintenanceSubtypeInfraUpdate,
	"CPS_SERVICES_UPDATE": ListMaintenanceRunsMaintenanceSubtypeCpsServicesUpdate,
	"CPS_VM_UPDATE":       ListMaintenanceRunsMaintenanceSubtypeCpsVmUpdate,
	"SECURITY_MONTHLY":    ListMaintenanceRunsMaintenanceSubtypeSecurityMonthly,
}

var mappingListMaintenanceRunsMaintenanceSubtypeEnumLowerCase = map[string]ListMaintenanceRunsMaintenanceSubtypeEnum{
	"yearly":              ListMaintenanceRunsMaintenanceSubtypeYearly,
	"halfyearly":          ListMaintenanceRunsMaintenanceSubtypeHalfyearly,
	"quarterly":           ListMaintenanceRunsMaintenanceSubtypeQuarterly,
	"monthly":             ListMaintenanceRunsMaintenanceSubtypeMonthly,
	"daily":               ListMaintenanceRunsMaintenanceSubtypeDaily,
	"hardware":            ListMaintenanceRunsMaintenanceSubtypeHardware,
	"critical":            ListMaintenanceRunsMaintenanceSubtypeCritical,
	"infra_update":        ListMaintenanceRunsMaintenanceSubtypeInfraUpdate,
	"cps_services_update": ListMaintenanceRunsMaintenanceSubtypeCpsServicesUpdate,
	"cps_vm_update":       ListMaintenanceRunsMaintenanceSubtypeCpsVmUpdate,
	"security_monthly":    ListMaintenanceRunsMaintenanceSubtypeSecurityMonthly,
}

// GetListMaintenanceRunsMaintenanceSubtypeEnumValues Enumerates the set of values for ListMaintenanceRunsMaintenanceSubtypeEnum
func GetListMaintenanceRunsMaintenanceSubtypeEnumValues() []ListMaintenanceRunsMaintenanceSubtypeEnum {
	values := make([]ListMaintenanceRunsMaintenanceSubtypeEnum, 0)
	for _, v := range mappingListMaintenanceRunsMaintenanceSubtypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceRunsMaintenanceSubtypeEnumStringValues Enumerates the set of values in String for ListMaintenanceRunsMaintenanceSubtypeEnum
func GetListMaintenanceRunsMaintenanceSubtypeEnumStringValues() []string {
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

// GetMappingListMaintenanceRunsMaintenanceSubtypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceRunsMaintenanceSubtypeEnum(val string) (ListMaintenanceRunsMaintenanceSubtypeEnum, bool) {
	enum, ok := mappingListMaintenanceRunsMaintenanceSubtypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceRunsSortByEnum Enum with underlying type: string
type ListMaintenanceRunsSortByEnum string

// Set of constants representing the allowable values for ListMaintenanceRunsSortByEnum
const (
	ListMaintenanceRunsSortByTimecreated     ListMaintenanceRunsSortByEnum = "timeCreated"
	ListMaintenanceRunsSortByDisplayname     ListMaintenanceRunsSortByEnum = "displayName"
	ListMaintenanceRunsSortByBuildidentifier ListMaintenanceRunsSortByEnum = "buildIdentifier"
)

var mappingListMaintenanceRunsSortByEnum = map[string]ListMaintenanceRunsSortByEnum{
	"timeCreated":     ListMaintenanceRunsSortByTimecreated,
	"displayName":     ListMaintenanceRunsSortByDisplayname,
	"buildIdentifier": ListMaintenanceRunsSortByBuildidentifier,
}

var mappingListMaintenanceRunsSortByEnumLowerCase = map[string]ListMaintenanceRunsSortByEnum{
	"timecreated":     ListMaintenanceRunsSortByTimecreated,
	"displayname":     ListMaintenanceRunsSortByDisplayname,
	"buildidentifier": ListMaintenanceRunsSortByBuildidentifier,
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
		"timeCreated",
		"displayName",
		"buildIdentifier",
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
