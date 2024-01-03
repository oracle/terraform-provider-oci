// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogAnalyticsEntityTypeSummary Summary of a log analytics entity type.
type LogAnalyticsEntityTypeSummary struct {

	// Log analytics entity type name.
	Name *string `mandatory:"true" json:"name"`

	// Internal name for the log analytics entity type.
	InternalName *string `mandatory:"true" json:"internalName"`

	// Log analytics entity type category. Category will be used for grouping and filtering.
	Category *string `mandatory:"true" json:"category"`

	// Log analytics entity type group. This can be CLOUD (OCI) or NON_CLOUD otherwise.
	CloudType EntityCloudTypeEnum `mandatory:"true" json:"cloudType"`

	// The current lifecycle state of the log analytics entity type.
	LifecycleState EntityLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Time the log analytics entity type was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time the log analytics entity type was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// This field indicates whether logs for entities of this type can be collected using a management agent.
	ManagementAgentEligibilityStatus LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum `mandatory:"false" json:"managementAgentEligibilityStatus,omitempty"`
}

func (m LogAnalyticsEntityTypeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsEntityTypeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEntityCloudTypeEnum(string(m.CloudType)); !ok && m.CloudType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CloudType: %s. Supported values are: %s.", m.CloudType, strings.Join(GetEntityCloudTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEntityLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEntityLifecycleStatesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum(string(m.ManagementAgentEligibilityStatus)); !ok && m.ManagementAgentEligibilityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementAgentEligibilityStatus: %s. Supported values are: %s.", m.ManagementAgentEligibilityStatus, strings.Join(GetLogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum Enum with underlying type: string
type LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum string

// Set of constants representing the allowable values for LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum
const (
	LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEligible   LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum = "ELIGIBLE"
	LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusIneligible LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum = "INELIGIBLE"
	LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusUnknown    LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum = "UNKNOWN"
)

var mappingLogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum = map[string]LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum{
	"ELIGIBLE":   LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEligible,
	"INELIGIBLE": LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusIneligible,
	"UNKNOWN":    LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusUnknown,
}

var mappingLogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnumLowerCase = map[string]LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum{
	"eligible":   LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEligible,
	"ineligible": LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusIneligible,
	"unknown":    LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusUnknown,
}

// GetLogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnumValues Enumerates the set of values for LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum
func GetLogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnumValues() []LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum {
	values := make([]LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum, 0)
	for _, v := range mappingLogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnumStringValues Enumerates the set of values in String for LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum
func GetLogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnumStringValues() []string {
	return []string{
		"ELIGIBLE",
		"INELIGIBLE",
		"UNKNOWN",
	}
}

// GetMappingLogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum(val string) (LogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnum, bool) {
	enum, ok := mappingLogAnalyticsEntityTypeSummaryManagementAgentEligibilityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
