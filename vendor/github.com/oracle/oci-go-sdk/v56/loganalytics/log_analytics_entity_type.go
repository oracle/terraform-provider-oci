// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum Enum with underlying type: string
type LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum string

// Set of constants representing the allowable values for LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum
const (
	LogAnalyticsEntityTypeManagementAgentEligibilityStatusEligible   LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum = "ELIGIBLE"
	LogAnalyticsEntityTypeManagementAgentEligibilityStatusIneligible LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum = "INELIGIBLE"
	LogAnalyticsEntityTypeManagementAgentEligibilityStatusUnknown    LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum = "UNKNOWN"
)

var mappingLogAnalyticsEntityTypeManagementAgentEligibilityStatus = map[string]LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum{
	"ELIGIBLE":   LogAnalyticsEntityTypeManagementAgentEligibilityStatusEligible,
	"INELIGIBLE": LogAnalyticsEntityTypeManagementAgentEligibilityStatusIneligible,
	"UNKNOWN":    LogAnalyticsEntityTypeManagementAgentEligibilityStatusUnknown,
}

// GetLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnumValues Enumerates the set of values for LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum
func GetLogAnalyticsEntityTypeManagementAgentEligibilityStatusEnumValues() []LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum {
	values := make([]LogAnalyticsEntityTypeManagementAgentEligibilityStatusEnum, 0)
	for _, v := range mappingLogAnalyticsEntityTypeManagementAgentEligibilityStatus {
		values = append(values, v)
	}
	return values
}
