// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProblemsRequest wrapper for the ListProblems operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListProblems.go.html to see an example of how to use ListProblemsRequest.
type ListProblemsRequest struct {

	// The OCID of the compartment in which to list resources.
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
	LifecycleDetail ListProblemsLifecycleDetailEnum `mandatory:"false" contributesTo:"query" name:"lifecycleDetail" omitEmpty:"true"`

	// The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListProblemsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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
	DetectorType ListProblemsDetectorTypeEnum `mandatory:"false" contributesTo:"query" name:"detectorType" omitEmpty:"true"`

	// The ID of the target in which to list resources.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// Setting this to `SECURITY_ZONE` returns only security zone-related violations.
	ProblemCategory ListProblemsProblemCategoryEnum `mandatory:"false" contributesTo:"query" name:"problemCategory" omitEmpty:"true"`

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
	AccessLevel ListProblemsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The ID of the resource associated with the problem.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListProblemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for riskLevel, timeLastDetected and resourceName is descending. Default order for riskLevel and resourceName is ascending. If no value is specified timeLastDetected is default.
	SortBy ListProblemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProblemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProblemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProblemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProblemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProblemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProblemsLifecycleDetailEnum(string(request.LifecycleDetail)); !ok && request.LifecycleDetail != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetail: %s. Supported values are: %s.", request.LifecycleDetail, strings.Join(GetListProblemsLifecycleDetailEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProblemsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListProblemsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProblemsDetectorTypeEnum(string(request.DetectorType)); !ok && request.DetectorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DetectorType: %s. Supported values are: %s.", request.DetectorType, strings.Join(GetListProblemsDetectorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProblemsProblemCategoryEnum(string(request.ProblemCategory)); !ok && request.ProblemCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProblemCategory: %s. Supported values are: %s.", request.ProblemCategory, strings.Join(GetListProblemsProblemCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProblemsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListProblemsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProblemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProblemsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProblemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProblemsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProblemsResponse wrapper for the ListProblems operation
type ListProblemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProblemCollection instances
	ProblemCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProblemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProblemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProblemsLifecycleDetailEnum Enum with underlying type: string
type ListProblemsLifecycleDetailEnum string

// Set of constants representing the allowable values for ListProblemsLifecycleDetailEnum
const (
	ListProblemsLifecycleDetailOpen      ListProblemsLifecycleDetailEnum = "OPEN"
	ListProblemsLifecycleDetailResolved  ListProblemsLifecycleDetailEnum = "RESOLVED"
	ListProblemsLifecycleDetailDismissed ListProblemsLifecycleDetailEnum = "DISMISSED"
	ListProblemsLifecycleDetailDeleted   ListProblemsLifecycleDetailEnum = "DELETED"
)

var mappingListProblemsLifecycleDetailEnum = map[string]ListProblemsLifecycleDetailEnum{
	"OPEN":      ListProblemsLifecycleDetailOpen,
	"RESOLVED":  ListProblemsLifecycleDetailResolved,
	"DISMISSED": ListProblemsLifecycleDetailDismissed,
	"DELETED":   ListProblemsLifecycleDetailDeleted,
}

var mappingListProblemsLifecycleDetailEnumLowerCase = map[string]ListProblemsLifecycleDetailEnum{
	"open":      ListProblemsLifecycleDetailOpen,
	"resolved":  ListProblemsLifecycleDetailResolved,
	"dismissed": ListProblemsLifecycleDetailDismissed,
	"deleted":   ListProblemsLifecycleDetailDeleted,
}

// GetListProblemsLifecycleDetailEnumValues Enumerates the set of values for ListProblemsLifecycleDetailEnum
func GetListProblemsLifecycleDetailEnumValues() []ListProblemsLifecycleDetailEnum {
	values := make([]ListProblemsLifecycleDetailEnum, 0)
	for _, v := range mappingListProblemsLifecycleDetailEnum {
		values = append(values, v)
	}
	return values
}

// GetListProblemsLifecycleDetailEnumStringValues Enumerates the set of values in String for ListProblemsLifecycleDetailEnum
func GetListProblemsLifecycleDetailEnumStringValues() []string {
	return []string{
		"OPEN",
		"RESOLVED",
		"DISMISSED",
		"DELETED",
	}
}

// GetMappingListProblemsLifecycleDetailEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProblemsLifecycleDetailEnum(val string) (ListProblemsLifecycleDetailEnum, bool) {
	enum, ok := mappingListProblemsLifecycleDetailEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProblemsLifecycleStateEnum Enum with underlying type: string
type ListProblemsLifecycleStateEnum string

// Set of constants representing the allowable values for ListProblemsLifecycleStateEnum
const (
	ListProblemsLifecycleStateActive   ListProblemsLifecycleStateEnum = "ACTIVE"
	ListProblemsLifecycleStateInactive ListProblemsLifecycleStateEnum = "INACTIVE"
)

var mappingListProblemsLifecycleStateEnum = map[string]ListProblemsLifecycleStateEnum{
	"ACTIVE":   ListProblemsLifecycleStateActive,
	"INACTIVE": ListProblemsLifecycleStateInactive,
}

var mappingListProblemsLifecycleStateEnumLowerCase = map[string]ListProblemsLifecycleStateEnum{
	"active":   ListProblemsLifecycleStateActive,
	"inactive": ListProblemsLifecycleStateInactive,
}

// GetListProblemsLifecycleStateEnumValues Enumerates the set of values for ListProblemsLifecycleStateEnum
func GetListProblemsLifecycleStateEnumValues() []ListProblemsLifecycleStateEnum {
	values := make([]ListProblemsLifecycleStateEnum, 0)
	for _, v := range mappingListProblemsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListProblemsLifecycleStateEnumStringValues Enumerates the set of values in String for ListProblemsLifecycleStateEnum
func GetListProblemsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingListProblemsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProblemsLifecycleStateEnum(val string) (ListProblemsLifecycleStateEnum, bool) {
	enum, ok := mappingListProblemsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProblemsDetectorTypeEnum Enum with underlying type: string
type ListProblemsDetectorTypeEnum string

// Set of constants representing the allowable values for ListProblemsDetectorTypeEnum
const (
	ListProblemsDetectorTypeActivityDetector         ListProblemsDetectorTypeEnum = "IAAS_ACTIVITY_DETECTOR"
	ListProblemsDetectorTypeConfigurationDetector    ListProblemsDetectorTypeEnum = "IAAS_CONFIGURATION_DETECTOR"
	ListProblemsDetectorTypeThreatDetector           ListProblemsDetectorTypeEnum = "IAAS_THREAT_DETECTOR"
	ListProblemsDetectorTypeLogInsightDetector       ListProblemsDetectorTypeEnum = "IAAS_LOG_INSIGHT_DETECTOR"
	ListProblemsDetectorTypeInstanceSecurityDetector ListProblemsDetectorTypeEnum = "IAAS_INSTANCE_SECURITY_DETECTOR"
)

var mappingListProblemsDetectorTypeEnum = map[string]ListProblemsDetectorTypeEnum{
	"IAAS_ACTIVITY_DETECTOR":          ListProblemsDetectorTypeActivityDetector,
	"IAAS_CONFIGURATION_DETECTOR":     ListProblemsDetectorTypeConfigurationDetector,
	"IAAS_THREAT_DETECTOR":            ListProblemsDetectorTypeThreatDetector,
	"IAAS_LOG_INSIGHT_DETECTOR":       ListProblemsDetectorTypeLogInsightDetector,
	"IAAS_INSTANCE_SECURITY_DETECTOR": ListProblemsDetectorTypeInstanceSecurityDetector,
}

var mappingListProblemsDetectorTypeEnumLowerCase = map[string]ListProblemsDetectorTypeEnum{
	"iaas_activity_detector":          ListProblemsDetectorTypeActivityDetector,
	"iaas_configuration_detector":     ListProblemsDetectorTypeConfigurationDetector,
	"iaas_threat_detector":            ListProblemsDetectorTypeThreatDetector,
	"iaas_log_insight_detector":       ListProblemsDetectorTypeLogInsightDetector,
	"iaas_instance_security_detector": ListProblemsDetectorTypeInstanceSecurityDetector,
}

// GetListProblemsDetectorTypeEnumValues Enumerates the set of values for ListProblemsDetectorTypeEnum
func GetListProblemsDetectorTypeEnumValues() []ListProblemsDetectorTypeEnum {
	values := make([]ListProblemsDetectorTypeEnum, 0)
	for _, v := range mappingListProblemsDetectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListProblemsDetectorTypeEnumStringValues Enumerates the set of values in String for ListProblemsDetectorTypeEnum
func GetListProblemsDetectorTypeEnumStringValues() []string {
	return []string{
		"IAAS_ACTIVITY_DETECTOR",
		"IAAS_CONFIGURATION_DETECTOR",
		"IAAS_THREAT_DETECTOR",
		"IAAS_LOG_INSIGHT_DETECTOR",
		"IAAS_INSTANCE_SECURITY_DETECTOR",
	}
}

// GetMappingListProblemsDetectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProblemsDetectorTypeEnum(val string) (ListProblemsDetectorTypeEnum, bool) {
	enum, ok := mappingListProblemsDetectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProblemsProblemCategoryEnum Enum with underlying type: string
type ListProblemsProblemCategoryEnum string

// Set of constants representing the allowable values for ListProblemsProblemCategoryEnum
const (
	ListProblemsProblemCategorySecurityZone ListProblemsProblemCategoryEnum = "SECURITY_ZONE"
)

var mappingListProblemsProblemCategoryEnum = map[string]ListProblemsProblemCategoryEnum{
	"SECURITY_ZONE": ListProblemsProblemCategorySecurityZone,
}

var mappingListProblemsProblemCategoryEnumLowerCase = map[string]ListProblemsProblemCategoryEnum{
	"security_zone": ListProblemsProblemCategorySecurityZone,
}

// GetListProblemsProblemCategoryEnumValues Enumerates the set of values for ListProblemsProblemCategoryEnum
func GetListProblemsProblemCategoryEnumValues() []ListProblemsProblemCategoryEnum {
	values := make([]ListProblemsProblemCategoryEnum, 0)
	for _, v := range mappingListProblemsProblemCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetListProblemsProblemCategoryEnumStringValues Enumerates the set of values in String for ListProblemsProblemCategoryEnum
func GetListProblemsProblemCategoryEnumStringValues() []string {
	return []string{
		"SECURITY_ZONE",
	}
}

// GetMappingListProblemsProblemCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProblemsProblemCategoryEnum(val string) (ListProblemsProblemCategoryEnum, bool) {
	enum, ok := mappingListProblemsProblemCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProblemsAccessLevelEnum Enum with underlying type: string
type ListProblemsAccessLevelEnum string

// Set of constants representing the allowable values for ListProblemsAccessLevelEnum
const (
	ListProblemsAccessLevelRestricted ListProblemsAccessLevelEnum = "RESTRICTED"
	ListProblemsAccessLevelAccessible ListProblemsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListProblemsAccessLevelEnum = map[string]ListProblemsAccessLevelEnum{
	"RESTRICTED": ListProblemsAccessLevelRestricted,
	"ACCESSIBLE": ListProblemsAccessLevelAccessible,
}

var mappingListProblemsAccessLevelEnumLowerCase = map[string]ListProblemsAccessLevelEnum{
	"restricted": ListProblemsAccessLevelRestricted,
	"accessible": ListProblemsAccessLevelAccessible,
}

// GetListProblemsAccessLevelEnumValues Enumerates the set of values for ListProblemsAccessLevelEnum
func GetListProblemsAccessLevelEnumValues() []ListProblemsAccessLevelEnum {
	values := make([]ListProblemsAccessLevelEnum, 0)
	for _, v := range mappingListProblemsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListProblemsAccessLevelEnumStringValues Enumerates the set of values in String for ListProblemsAccessLevelEnum
func GetListProblemsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListProblemsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProblemsAccessLevelEnum(val string) (ListProblemsAccessLevelEnum, bool) {
	enum, ok := mappingListProblemsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProblemsSortOrderEnum Enum with underlying type: string
type ListProblemsSortOrderEnum string

// Set of constants representing the allowable values for ListProblemsSortOrderEnum
const (
	ListProblemsSortOrderAsc  ListProblemsSortOrderEnum = "ASC"
	ListProblemsSortOrderDesc ListProblemsSortOrderEnum = "DESC"
)

var mappingListProblemsSortOrderEnum = map[string]ListProblemsSortOrderEnum{
	"ASC":  ListProblemsSortOrderAsc,
	"DESC": ListProblemsSortOrderDesc,
}

var mappingListProblemsSortOrderEnumLowerCase = map[string]ListProblemsSortOrderEnum{
	"asc":  ListProblemsSortOrderAsc,
	"desc": ListProblemsSortOrderDesc,
}

// GetListProblemsSortOrderEnumValues Enumerates the set of values for ListProblemsSortOrderEnum
func GetListProblemsSortOrderEnumValues() []ListProblemsSortOrderEnum {
	values := make([]ListProblemsSortOrderEnum, 0)
	for _, v := range mappingListProblemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProblemsSortOrderEnumStringValues Enumerates the set of values in String for ListProblemsSortOrderEnum
func GetListProblemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProblemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProblemsSortOrderEnum(val string) (ListProblemsSortOrderEnum, bool) {
	enum, ok := mappingListProblemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProblemsSortByEnum Enum with underlying type: string
type ListProblemsSortByEnum string

// Set of constants representing the allowable values for ListProblemsSortByEnum
const (
	ListProblemsSortByRisklevel        ListProblemsSortByEnum = "riskLevel"
	ListProblemsSortByTimelastdetected ListProblemsSortByEnum = "timeLastDetected"
	ListProblemsSortByResourcename     ListProblemsSortByEnum = "resourceName"
)

var mappingListProblemsSortByEnum = map[string]ListProblemsSortByEnum{
	"riskLevel":        ListProblemsSortByRisklevel,
	"timeLastDetected": ListProblemsSortByTimelastdetected,
	"resourceName":     ListProblemsSortByResourcename,
}

var mappingListProblemsSortByEnumLowerCase = map[string]ListProblemsSortByEnum{
	"risklevel":        ListProblemsSortByRisklevel,
	"timelastdetected": ListProblemsSortByTimelastdetected,
	"resourcename":     ListProblemsSortByResourcename,
}

// GetListProblemsSortByEnumValues Enumerates the set of values for ListProblemsSortByEnum
func GetListProblemsSortByEnumValues() []ListProblemsSortByEnum {
	values := make([]ListProblemsSortByEnum, 0)
	for _, v := range mappingListProblemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProblemsSortByEnumStringValues Enumerates the set of values in String for ListProblemsSortByEnum
func GetListProblemsSortByEnumStringValues() []string {
	return []string{
		"riskLevel",
		"timeLastDetected",
		"resourceName",
	}
}

// GetMappingListProblemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProblemsSortByEnum(val string) (ListProblemsSortByEnum, bool) {
	enum, ok := mappingListProblemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
