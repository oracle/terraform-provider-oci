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

// LogAnalyticsEntityType Description of log analytics entity type.
type LogAnalyticsEntityType struct {

	// Log analytics entity type name.
	Name *string `mandatory:"true" json:"name"`

	// Internal name for the log analytics entity type.
	InternalName *string `mandatory:"true" json:"internalName"`

	// Log analytics entity type category. Category will be used for grouping and filtering.
	Category *string `mandatory:"true" json:"category"`

	// Log analytics entity type group. That can be CLOUD (OCI) or NON_CLOUD otherwise.
	CloudType EntityCloudTypeEnum `mandatory:"true" json:"cloudType"`

	// The current lifecycle state of the log analytics entity.
	LifecycleState EntityLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Time the log analytics entity type was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time the log analytics entity type was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The parameters used in file patterns specified in log sources for this log analytics entity type.
	Properties []EntityTypeProperty `mandatory:"false" json:"properties"`

	// This field indicates whether logs for entities of this type can be collected using a management agent.
	ManagementAgentEligibilityStatus LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum `mandatory:"false" json:"managementAgentEligibilityStatus,omitempty"`
}

func (m LogAnalyticsEntityType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsEntityType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEntityCloudTypeEnum(string(m.CloudType)); !ok && m.CloudType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CloudType: %s. Supported values are: %s.", m.CloudType, strings.Join(GetEntityCloudTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEntityLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEntityLifecycleStatesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum(string(m.ManagementAgentEligibilityStatus)); !ok && m.ManagementAgentEligibilityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementAgentEligibilityStatus: %s. Supported values are: %s.", m.ManagementAgentEligibilityStatus, strings.Join(GetLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum Enum with underlying type: string
type LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum string

// Set of constants representing the allowable values for LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum
const (
	LogAnalyticsEntityTypeManagementAgentEligibilityStatusEligible   LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum = "ELIGIBLE"
	LogAnalyticsEntityTypeManagementAgentEligibilityStatusIneligible LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum = "INELIGIBLE"
	LogAnalyticsEntityTypeManagementAgentEligibilityStatusUnknown    LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum = "UNKNOWN"
)

var mappingLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum = map[string]LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum{
	"ELIGIBLE":   LogAnalyticsEntityTypeManagementAgentEligibilityStatusEligible,
	"INELIGIBLE": LogAnalyticsEntityTypeManagementAgentEligibilityStatusIneligible,
	"UNKNOWN":    LogAnalyticsEntityTypeManagementAgentEligibilityStatusUnknown,
}

var mappingLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnumLowerCase = map[string]LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum{
	"eligible":   LogAnalyticsEntityTypeManagementAgentEligibilityStatusEligible,
	"ineligible": LogAnalyticsEntityTypeManagementAgentEligibilityStatusIneligible,
	"unknown":    LogAnalyticsEntityTypeManagementAgentEligibilityStatusUnknown,
}

// GetLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnumValues Enumerates the set of values for LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum
func GetLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnumValues() []LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum {
	values := make([]LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum, 0)
	for _, v := range mappingLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnumStringValues Enumerates the set of values in String for LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum
func GetLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnumStringValues() []string {
	return []string{
		"ELIGIBLE",
		"INELIGIBLE",
		"UNKNOWN",
	}
}

// GetMappingLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum(val string) (LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum, bool) {
	enum, ok := mappingLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
