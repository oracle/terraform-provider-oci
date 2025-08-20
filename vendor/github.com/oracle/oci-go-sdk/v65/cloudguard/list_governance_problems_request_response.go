// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListGovernanceProblemsRequest wrapper for the ListGovernanceProblems operation
type ListGovernanceProblemsRequest struct {

	// The OCID of the compartment that contains the governance target
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Start time for a filter. If start time is not specified, start time will be set to current time - 30 days.
	TimeLastDetectedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastDetectedGreaterThanOrEqualTo"`

	// End time for a filter. If end time is not specified, end time will be set to current time.
	TimeLastDetectedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastDetectedLessThanOrEqualTo"`

	// Start time for a filter. If start time is not specified, start time will be set to current time - 30 days.
	TimeFirstDetectedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeFirstDetectedGreaterThanOrEqualTo"`

	// End time for a filter. If end time is not specified, end time will be set to current time.
	TimeFirstDetectedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeFirstDetectedLessThanOrEqualTo"`

	// The field life cycle state. Only one state can be provided. Default value for state is active.
	LifecycleDetail ListGovernanceProblemsLifecycleDetailEnum `mandatory:"false" contributesTo:"query" name:"lifecycleDetail" omitEmpty:"true"`

	// The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListGovernanceProblemsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// OCI monitoring region.
	Region *string `mandatory:"false" contributesTo:"query" name:"region"`

	// Risk level of the problem.
	RiskLevel *string `mandatory:"false" contributesTo:"query" name:"riskLevel"`

	// Resource type associated with the resource.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// City of the problem.
	City *string `mandatory:"false" contributesTo:"query" name:"city"`

	// State or province of the problem.
	State *string `mandatory:"false" contributesTo:"query" name:"state"`

	// Country of the problem.
	Country *string `mandatory:"false" contributesTo:"query" name:"country"`

	// User-defined label associated with the problem.
	Label *string `mandatory:"false" contributesTo:"query" name:"label"`

	// Comma seperated list of detector rule IDs to be passed in to match against Problems.
	DetectorRuleIdList []string `contributesTo:"query" name:"detectorRuleIdList" collectionFormat:"multi"`

	// The field to list the problems by detector type.
	DetectorType ListGovernanceProblemsDetectorTypeEnum `mandatory:"false" contributesTo:"query" name:"detectorType" omitEmpty:"true"`

	// Setting this to `SECURITY_ZONE` returns only security zone-related violations.
	ProblemCategory ListGovernanceProblemsProblemCategoryEnum `mandatory:"false" contributesTo:"query" name:"problemCategory" omitEmpty:"true"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel ListGovernanceProblemsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The ID of the resource associated with the problem.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListGovernanceProblemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for riskLevel, timeLastDetected and resourceName is descending. Default order for riskLevel and resourceName is ascending. If no value is specified timeLastDetected is default.
	SortBy ListGovernanceProblemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The OCID of the governance target
	GovernanceTargetId *string `mandatory:"false" contributesTo:"query" name:"governanceTargetId"`

	// The name of the subject tenant
	SubjectTenantName *string `mandatory:"false" contributesTo:"query" name:"subjectTenantName"`

	// The subject tenant name prefix
	SubjectTenantNamePrefix *string `mandatory:"false" contributesTo:"query" name:"subjectTenantNamePrefix"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGovernanceProblemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGovernanceProblemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGovernanceProblemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGovernanceProblemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGovernanceProblemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGovernanceProblemsLifecycleDetailEnum(string(request.LifecycleDetail)); !ok && request.LifecycleDetail != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetail: %s. Supported values are: %s.", request.LifecycleDetail, strings.Join(GetListGovernanceProblemsLifecycleDetailEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceProblemsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListGovernanceProblemsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceProblemsDetectorTypeEnum(string(request.DetectorType)); !ok && request.DetectorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DetectorType: %s. Supported values are: %s.", request.DetectorType, strings.Join(GetListGovernanceProblemsDetectorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceProblemsProblemCategoryEnum(string(request.ProblemCategory)); !ok && request.ProblemCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProblemCategory: %s. Supported values are: %s.", request.ProblemCategory, strings.Join(GetListGovernanceProblemsProblemCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceProblemsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListGovernanceProblemsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceProblemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGovernanceProblemsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceProblemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGovernanceProblemsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGovernanceProblemsResponse wrapper for the ListGovernanceProblems operation
type ListGovernanceProblemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GovernanceProblemCollection instances
	GovernanceProblemCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGovernanceProblemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGovernanceProblemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGovernanceProblemsLifecycleDetailEnum Enum with underlying type: string
type ListGovernanceProblemsLifecycleDetailEnum string

// Set of constants representing the allowable values for ListGovernanceProblemsLifecycleDetailEnum
const (
	ListGovernanceProblemsLifecycleDetailOpen      ListGovernanceProblemsLifecycleDetailEnum = "OPEN"
	ListGovernanceProblemsLifecycleDetailResolved  ListGovernanceProblemsLifecycleDetailEnum = "RESOLVED"
	ListGovernanceProblemsLifecycleDetailDismissed ListGovernanceProblemsLifecycleDetailEnum = "DISMISSED"
	ListGovernanceProblemsLifecycleDetailDeleted   ListGovernanceProblemsLifecycleDetailEnum = "DELETED"
)

var mappingListGovernanceProblemsLifecycleDetailEnum = map[string]ListGovernanceProblemsLifecycleDetailEnum{
	"OPEN":      ListGovernanceProblemsLifecycleDetailOpen,
	"RESOLVED":  ListGovernanceProblemsLifecycleDetailResolved,
	"DISMISSED": ListGovernanceProblemsLifecycleDetailDismissed,
	"DELETED":   ListGovernanceProblemsLifecycleDetailDeleted,
}

var mappingListGovernanceProblemsLifecycleDetailEnumLowerCase = map[string]ListGovernanceProblemsLifecycleDetailEnum{
	"open":      ListGovernanceProblemsLifecycleDetailOpen,
	"resolved":  ListGovernanceProblemsLifecycleDetailResolved,
	"dismissed": ListGovernanceProblemsLifecycleDetailDismissed,
	"deleted":   ListGovernanceProblemsLifecycleDetailDeleted,
}

// GetListGovernanceProblemsLifecycleDetailEnumValues Enumerates the set of values for ListGovernanceProblemsLifecycleDetailEnum
func GetListGovernanceProblemsLifecycleDetailEnumValues() []ListGovernanceProblemsLifecycleDetailEnum {
	values := make([]ListGovernanceProblemsLifecycleDetailEnum, 0)
	for _, v := range mappingListGovernanceProblemsLifecycleDetailEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceProblemsLifecycleDetailEnumStringValues Enumerates the set of values in String for ListGovernanceProblemsLifecycleDetailEnum
func GetListGovernanceProblemsLifecycleDetailEnumStringValues() []string {
	return []string{
		"OPEN",
		"RESOLVED",
		"DISMISSED",
		"DELETED",
	}
}

// GetMappingListGovernanceProblemsLifecycleDetailEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceProblemsLifecycleDetailEnum(val string) (ListGovernanceProblemsLifecycleDetailEnum, bool) {
	enum, ok := mappingListGovernanceProblemsLifecycleDetailEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceProblemsLifecycleStateEnum Enum with underlying type: string
type ListGovernanceProblemsLifecycleStateEnum string

// Set of constants representing the allowable values for ListGovernanceProblemsLifecycleStateEnum
const (
	ListGovernanceProblemsLifecycleStateActive   ListGovernanceProblemsLifecycleStateEnum = "ACTIVE"
	ListGovernanceProblemsLifecycleStateInactive ListGovernanceProblemsLifecycleStateEnum = "INACTIVE"
)

var mappingListGovernanceProblemsLifecycleStateEnum = map[string]ListGovernanceProblemsLifecycleStateEnum{
	"ACTIVE":   ListGovernanceProblemsLifecycleStateActive,
	"INACTIVE": ListGovernanceProblemsLifecycleStateInactive,
}

var mappingListGovernanceProblemsLifecycleStateEnumLowerCase = map[string]ListGovernanceProblemsLifecycleStateEnum{
	"active":   ListGovernanceProblemsLifecycleStateActive,
	"inactive": ListGovernanceProblemsLifecycleStateInactive,
}

// GetListGovernanceProblemsLifecycleStateEnumValues Enumerates the set of values for ListGovernanceProblemsLifecycleStateEnum
func GetListGovernanceProblemsLifecycleStateEnumValues() []ListGovernanceProblemsLifecycleStateEnum {
	values := make([]ListGovernanceProblemsLifecycleStateEnum, 0)
	for _, v := range mappingListGovernanceProblemsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceProblemsLifecycleStateEnumStringValues Enumerates the set of values in String for ListGovernanceProblemsLifecycleStateEnum
func GetListGovernanceProblemsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingListGovernanceProblemsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceProblemsLifecycleStateEnum(val string) (ListGovernanceProblemsLifecycleStateEnum, bool) {
	enum, ok := mappingListGovernanceProblemsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceProblemsDetectorTypeEnum Enum with underlying type: string
type ListGovernanceProblemsDetectorTypeEnum string

// Set of constants representing the allowable values for ListGovernanceProblemsDetectorTypeEnum
const (
	ListGovernanceProblemsDetectorTypeIaasActivityDetector          ListGovernanceProblemsDetectorTypeEnum = "IAAS_ACTIVITY_DETECTOR"
	ListGovernanceProblemsDetectorTypeIaasConfigurationDetector     ListGovernanceProblemsDetectorTypeEnum = "IAAS_CONFIGURATION_DETECTOR"
	ListGovernanceProblemsDetectorTypeIaasThreatDetector            ListGovernanceProblemsDetectorTypeEnum = "IAAS_THREAT_DETECTOR"
	ListGovernanceProblemsDetectorTypeIaasLoggingDetector           ListGovernanceProblemsDetectorTypeEnum = "IAAS_LOGGING_DETECTOR"
	ListGovernanceProblemsDetectorTypeIaasInsightDetector           ListGovernanceProblemsDetectorTypeEnum = "IAAS_INSIGHT_DETECTOR"
	ListGovernanceProblemsDetectorTypeIaasLogInsightDetector        ListGovernanceProblemsDetectorTypeEnum = "IAAS_LOG_INSIGHT_DETECTOR"
	ListGovernanceProblemsDetectorTypeSaasFaActivityDetector        ListGovernanceProblemsDetectorTypeEnum = "SAAS_FA_ACTIVITY_DETECTOR"
	ListGovernanceProblemsDetectorTypeIaasInstanceSecurityDetector  ListGovernanceProblemsDetectorTypeEnum = "IAAS_INSTANCE_SECURITY_DETECTOR"
	ListGovernanceProblemsDetectorTypeIaasContainerSecurityDetector ListGovernanceProblemsDetectorTypeEnum = "IAAS_CONTAINER_SECURITY_DETECTOR"
)

var mappingListGovernanceProblemsDetectorTypeEnum = map[string]ListGovernanceProblemsDetectorTypeEnum{
	"IAAS_ACTIVITY_DETECTOR":           ListGovernanceProblemsDetectorTypeIaasActivityDetector,
	"IAAS_CONFIGURATION_DETECTOR":      ListGovernanceProblemsDetectorTypeIaasConfigurationDetector,
	"IAAS_THREAT_DETECTOR":             ListGovernanceProblemsDetectorTypeIaasThreatDetector,
	"IAAS_LOGGING_DETECTOR":            ListGovernanceProblemsDetectorTypeIaasLoggingDetector,
	"IAAS_INSIGHT_DETECTOR":            ListGovernanceProblemsDetectorTypeIaasInsightDetector,
	"IAAS_LOG_INSIGHT_DETECTOR":        ListGovernanceProblemsDetectorTypeIaasLogInsightDetector,
	"SAAS_FA_ACTIVITY_DETECTOR":        ListGovernanceProblemsDetectorTypeSaasFaActivityDetector,
	"IAAS_INSTANCE_SECURITY_DETECTOR":  ListGovernanceProblemsDetectorTypeIaasInstanceSecurityDetector,
	"IAAS_CONTAINER_SECURITY_DETECTOR": ListGovernanceProblemsDetectorTypeIaasContainerSecurityDetector,
}

var mappingListGovernanceProblemsDetectorTypeEnumLowerCase = map[string]ListGovernanceProblemsDetectorTypeEnum{
	"iaas_activity_detector":           ListGovernanceProblemsDetectorTypeIaasActivityDetector,
	"iaas_configuration_detector":      ListGovernanceProblemsDetectorTypeIaasConfigurationDetector,
	"iaas_threat_detector":             ListGovernanceProblemsDetectorTypeIaasThreatDetector,
	"iaas_logging_detector":            ListGovernanceProblemsDetectorTypeIaasLoggingDetector,
	"iaas_insight_detector":            ListGovernanceProblemsDetectorTypeIaasInsightDetector,
	"iaas_log_insight_detector":        ListGovernanceProblemsDetectorTypeIaasLogInsightDetector,
	"saas_fa_activity_detector":        ListGovernanceProblemsDetectorTypeSaasFaActivityDetector,
	"iaas_instance_security_detector":  ListGovernanceProblemsDetectorTypeIaasInstanceSecurityDetector,
	"iaas_container_security_detector": ListGovernanceProblemsDetectorTypeIaasContainerSecurityDetector,
}

// GetListGovernanceProblemsDetectorTypeEnumValues Enumerates the set of values for ListGovernanceProblemsDetectorTypeEnum
func GetListGovernanceProblemsDetectorTypeEnumValues() []ListGovernanceProblemsDetectorTypeEnum {
	values := make([]ListGovernanceProblemsDetectorTypeEnum, 0)
	for _, v := range mappingListGovernanceProblemsDetectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceProblemsDetectorTypeEnumStringValues Enumerates the set of values in String for ListGovernanceProblemsDetectorTypeEnum
func GetListGovernanceProblemsDetectorTypeEnumStringValues() []string {
	return []string{
		"IAAS_ACTIVITY_DETECTOR",
		"IAAS_CONFIGURATION_DETECTOR",
		"IAAS_THREAT_DETECTOR",
		"IAAS_LOGGING_DETECTOR",
		"IAAS_INSIGHT_DETECTOR",
		"IAAS_LOG_INSIGHT_DETECTOR",
		"SAAS_FA_ACTIVITY_DETECTOR",
		"IAAS_INSTANCE_SECURITY_DETECTOR",
		"IAAS_CONTAINER_SECURITY_DETECTOR",
	}
}

// GetMappingListGovernanceProblemsDetectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceProblemsDetectorTypeEnum(val string) (ListGovernanceProblemsDetectorTypeEnum, bool) {
	enum, ok := mappingListGovernanceProblemsDetectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceProblemsProblemCategoryEnum Enum with underlying type: string
type ListGovernanceProblemsProblemCategoryEnum string

// Set of constants representing the allowable values for ListGovernanceProblemsProblemCategoryEnum
const (
	ListGovernanceProblemsProblemCategorySecurityZone ListGovernanceProblemsProblemCategoryEnum = "SECURITY_ZONE"
)

var mappingListGovernanceProblemsProblemCategoryEnum = map[string]ListGovernanceProblemsProblemCategoryEnum{
	"SECURITY_ZONE": ListGovernanceProblemsProblemCategorySecurityZone,
}

var mappingListGovernanceProblemsProblemCategoryEnumLowerCase = map[string]ListGovernanceProblemsProblemCategoryEnum{
	"security_zone": ListGovernanceProblemsProblemCategorySecurityZone,
}

// GetListGovernanceProblemsProblemCategoryEnumValues Enumerates the set of values for ListGovernanceProblemsProblemCategoryEnum
func GetListGovernanceProblemsProblemCategoryEnumValues() []ListGovernanceProblemsProblemCategoryEnum {
	values := make([]ListGovernanceProblemsProblemCategoryEnum, 0)
	for _, v := range mappingListGovernanceProblemsProblemCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceProblemsProblemCategoryEnumStringValues Enumerates the set of values in String for ListGovernanceProblemsProblemCategoryEnum
func GetListGovernanceProblemsProblemCategoryEnumStringValues() []string {
	return []string{
		"SECURITY_ZONE",
	}
}

// GetMappingListGovernanceProblemsProblemCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceProblemsProblemCategoryEnum(val string) (ListGovernanceProblemsProblemCategoryEnum, bool) {
	enum, ok := mappingListGovernanceProblemsProblemCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceProblemsAccessLevelEnum Enum with underlying type: string
type ListGovernanceProblemsAccessLevelEnum string

// Set of constants representing the allowable values for ListGovernanceProblemsAccessLevelEnum
const (
	ListGovernanceProblemsAccessLevelRestricted ListGovernanceProblemsAccessLevelEnum = "RESTRICTED"
	ListGovernanceProblemsAccessLevelAccessible ListGovernanceProblemsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListGovernanceProblemsAccessLevelEnum = map[string]ListGovernanceProblemsAccessLevelEnum{
	"RESTRICTED": ListGovernanceProblemsAccessLevelRestricted,
	"ACCESSIBLE": ListGovernanceProblemsAccessLevelAccessible,
}

var mappingListGovernanceProblemsAccessLevelEnumLowerCase = map[string]ListGovernanceProblemsAccessLevelEnum{
	"restricted": ListGovernanceProblemsAccessLevelRestricted,
	"accessible": ListGovernanceProblemsAccessLevelAccessible,
}

// GetListGovernanceProblemsAccessLevelEnumValues Enumerates the set of values for ListGovernanceProblemsAccessLevelEnum
func GetListGovernanceProblemsAccessLevelEnumValues() []ListGovernanceProblemsAccessLevelEnum {
	values := make([]ListGovernanceProblemsAccessLevelEnum, 0)
	for _, v := range mappingListGovernanceProblemsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceProblemsAccessLevelEnumStringValues Enumerates the set of values in String for ListGovernanceProblemsAccessLevelEnum
func GetListGovernanceProblemsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListGovernanceProblemsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceProblemsAccessLevelEnum(val string) (ListGovernanceProblemsAccessLevelEnum, bool) {
	enum, ok := mappingListGovernanceProblemsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceProblemsSortOrderEnum Enum with underlying type: string
type ListGovernanceProblemsSortOrderEnum string

// Set of constants representing the allowable values for ListGovernanceProblemsSortOrderEnum
const (
	ListGovernanceProblemsSortOrderAsc  ListGovernanceProblemsSortOrderEnum = "ASC"
	ListGovernanceProblemsSortOrderDesc ListGovernanceProblemsSortOrderEnum = "DESC"
)

var mappingListGovernanceProblemsSortOrderEnum = map[string]ListGovernanceProblemsSortOrderEnum{
	"ASC":  ListGovernanceProblemsSortOrderAsc,
	"DESC": ListGovernanceProblemsSortOrderDesc,
}

var mappingListGovernanceProblemsSortOrderEnumLowerCase = map[string]ListGovernanceProblemsSortOrderEnum{
	"asc":  ListGovernanceProblemsSortOrderAsc,
	"desc": ListGovernanceProblemsSortOrderDesc,
}

// GetListGovernanceProblemsSortOrderEnumValues Enumerates the set of values for ListGovernanceProblemsSortOrderEnum
func GetListGovernanceProblemsSortOrderEnumValues() []ListGovernanceProblemsSortOrderEnum {
	values := make([]ListGovernanceProblemsSortOrderEnum, 0)
	for _, v := range mappingListGovernanceProblemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceProblemsSortOrderEnumStringValues Enumerates the set of values in String for ListGovernanceProblemsSortOrderEnum
func GetListGovernanceProblemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGovernanceProblemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceProblemsSortOrderEnum(val string) (ListGovernanceProblemsSortOrderEnum, bool) {
	enum, ok := mappingListGovernanceProblemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceProblemsSortByEnum Enum with underlying type: string
type ListGovernanceProblemsSortByEnum string

// Set of constants representing the allowable values for ListGovernanceProblemsSortByEnum
const (
	ListGovernanceProblemsSortByRisklevel        ListGovernanceProblemsSortByEnum = "riskLevel"
	ListGovernanceProblemsSortByTimelastdetected ListGovernanceProblemsSortByEnum = "timeLastDetected"
	ListGovernanceProblemsSortByResourcename     ListGovernanceProblemsSortByEnum = "resourceName"
)

var mappingListGovernanceProblemsSortByEnum = map[string]ListGovernanceProblemsSortByEnum{
	"riskLevel":        ListGovernanceProblemsSortByRisklevel,
	"timeLastDetected": ListGovernanceProblemsSortByTimelastdetected,
	"resourceName":     ListGovernanceProblemsSortByResourcename,
}

var mappingListGovernanceProblemsSortByEnumLowerCase = map[string]ListGovernanceProblemsSortByEnum{
	"risklevel":        ListGovernanceProblemsSortByRisklevel,
	"timelastdetected": ListGovernanceProblemsSortByTimelastdetected,
	"resourcename":     ListGovernanceProblemsSortByResourcename,
}

// GetListGovernanceProblemsSortByEnumValues Enumerates the set of values for ListGovernanceProblemsSortByEnum
func GetListGovernanceProblemsSortByEnumValues() []ListGovernanceProblemsSortByEnum {
	values := make([]ListGovernanceProblemsSortByEnum, 0)
	for _, v := range mappingListGovernanceProblemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceProblemsSortByEnumStringValues Enumerates the set of values in String for ListGovernanceProblemsSortByEnum
func GetListGovernanceProblemsSortByEnumStringValues() []string {
	return []string{
		"riskLevel",
		"timeLastDetected",
		"resourceName",
	}
}

// GetMappingListGovernanceProblemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceProblemsSortByEnum(val string) (ListGovernanceProblemsSortByEnum, bool) {
	enum, ok := mappingListGovernanceProblemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
