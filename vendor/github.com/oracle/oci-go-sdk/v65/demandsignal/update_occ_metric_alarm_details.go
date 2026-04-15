// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Demand Signal API
//
// Use the OCI Control Center Demand Signal API to manage Demand Signals.
//

package demandsignal

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateOccMetricAlarmDetails Fields that can be updated for OccMetricAlarm.
type UpdateOccMetricAlarmDetails struct {

	// Human-readable name for the alarm.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description for the alarm.
	Description *string `mandatory:"false" json:"description"`

	// The current lifecycle state of the resource.
	LifecycleState UpdateOccMetricAlarmDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Alarm active status.
	IsActive *bool `mandatory:"false" json:"isActive"`

	// List of topic OCIDs for notifications.
	Subscribers []string `mandatory:"false" json:"subscribers"`

	// Frequency at which notifications should be sent.
	Frequency OccMetricAlarmFrequencyEnum `mandatory:"false" json:"frequency,omitempty"`

	// Threshold at which alarm must be triggered.
	Threshold *int `mandatory:"false" json:"threshold"`

	// Units in which threshold is being stored.
	ThresholdType UpdateOccMetricAlarmDetailsThresholdTypeEnum `mandatory:"false" json:"thresholdType,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateOccMetricAlarmDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOccMetricAlarmDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateOccMetricAlarmDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUpdateOccMetricAlarmDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccMetricAlarmFrequencyEnum(string(m.Frequency)); !ok && m.Frequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Frequency: %s. Supported values are: %s.", m.Frequency, strings.Join(GetOccMetricAlarmFrequencyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateOccMetricAlarmDetailsThresholdTypeEnum(string(m.ThresholdType)); !ok && m.ThresholdType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ThresholdType: %s. Supported values are: %s.", m.ThresholdType, strings.Join(GetUpdateOccMetricAlarmDetailsThresholdTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateOccMetricAlarmDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateOccMetricAlarmDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateOccMetricAlarmDetailsLifecycleStateEnum
const (
	UpdateOccMetricAlarmDetailsLifecycleStateCreating UpdateOccMetricAlarmDetailsLifecycleStateEnum = "CREATING"
	UpdateOccMetricAlarmDetailsLifecycleStateUpdating UpdateOccMetricAlarmDetailsLifecycleStateEnum = "UPDATING"
	UpdateOccMetricAlarmDetailsLifecycleStateActive   UpdateOccMetricAlarmDetailsLifecycleStateEnum = "ACTIVE"
	UpdateOccMetricAlarmDetailsLifecycleStateDeleting UpdateOccMetricAlarmDetailsLifecycleStateEnum = "DELETING"
	UpdateOccMetricAlarmDetailsLifecycleStateDeleted  UpdateOccMetricAlarmDetailsLifecycleStateEnum = "DELETED"
	UpdateOccMetricAlarmDetailsLifecycleStateFailed   UpdateOccMetricAlarmDetailsLifecycleStateEnum = "FAILED"
)

var mappingUpdateOccMetricAlarmDetailsLifecycleStateEnum = map[string]UpdateOccMetricAlarmDetailsLifecycleStateEnum{
	"CREATING": UpdateOccMetricAlarmDetailsLifecycleStateCreating,
	"UPDATING": UpdateOccMetricAlarmDetailsLifecycleStateUpdating,
	"ACTIVE":   UpdateOccMetricAlarmDetailsLifecycleStateActive,
	"DELETING": UpdateOccMetricAlarmDetailsLifecycleStateDeleting,
	"DELETED":  UpdateOccMetricAlarmDetailsLifecycleStateDeleted,
	"FAILED":   UpdateOccMetricAlarmDetailsLifecycleStateFailed,
}

var mappingUpdateOccMetricAlarmDetailsLifecycleStateEnumLowerCase = map[string]UpdateOccMetricAlarmDetailsLifecycleStateEnum{
	"creating": UpdateOccMetricAlarmDetailsLifecycleStateCreating,
	"updating": UpdateOccMetricAlarmDetailsLifecycleStateUpdating,
	"active":   UpdateOccMetricAlarmDetailsLifecycleStateActive,
	"deleting": UpdateOccMetricAlarmDetailsLifecycleStateDeleting,
	"deleted":  UpdateOccMetricAlarmDetailsLifecycleStateDeleted,
	"failed":   UpdateOccMetricAlarmDetailsLifecycleStateFailed,
}

// GetUpdateOccMetricAlarmDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateOccMetricAlarmDetailsLifecycleStateEnum
func GetUpdateOccMetricAlarmDetailsLifecycleStateEnumValues() []UpdateOccMetricAlarmDetailsLifecycleStateEnum {
	values := make([]UpdateOccMetricAlarmDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateOccMetricAlarmDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateOccMetricAlarmDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for UpdateOccMetricAlarmDetailsLifecycleStateEnum
func GetUpdateOccMetricAlarmDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingUpdateOccMetricAlarmDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateOccMetricAlarmDetailsLifecycleStateEnum(val string) (UpdateOccMetricAlarmDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingUpdateOccMetricAlarmDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateOccMetricAlarmDetailsThresholdTypeEnum Enum with underlying type: string
type UpdateOccMetricAlarmDetailsThresholdTypeEnum string

// Set of constants representing the allowable values for UpdateOccMetricAlarmDetailsThresholdTypeEnum
const (
	UpdateOccMetricAlarmDetailsThresholdTypePercentage UpdateOccMetricAlarmDetailsThresholdTypeEnum = "PERCENTAGE"
	UpdateOccMetricAlarmDetailsThresholdTypeUnits      UpdateOccMetricAlarmDetailsThresholdTypeEnum = "UNITS"
)

var mappingUpdateOccMetricAlarmDetailsThresholdTypeEnum = map[string]UpdateOccMetricAlarmDetailsThresholdTypeEnum{
	"PERCENTAGE": UpdateOccMetricAlarmDetailsThresholdTypePercentage,
	"UNITS":      UpdateOccMetricAlarmDetailsThresholdTypeUnits,
}

var mappingUpdateOccMetricAlarmDetailsThresholdTypeEnumLowerCase = map[string]UpdateOccMetricAlarmDetailsThresholdTypeEnum{
	"percentage": UpdateOccMetricAlarmDetailsThresholdTypePercentage,
	"units":      UpdateOccMetricAlarmDetailsThresholdTypeUnits,
}

// GetUpdateOccMetricAlarmDetailsThresholdTypeEnumValues Enumerates the set of values for UpdateOccMetricAlarmDetailsThresholdTypeEnum
func GetUpdateOccMetricAlarmDetailsThresholdTypeEnumValues() []UpdateOccMetricAlarmDetailsThresholdTypeEnum {
	values := make([]UpdateOccMetricAlarmDetailsThresholdTypeEnum, 0)
	for _, v := range mappingUpdateOccMetricAlarmDetailsThresholdTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateOccMetricAlarmDetailsThresholdTypeEnumStringValues Enumerates the set of values in String for UpdateOccMetricAlarmDetailsThresholdTypeEnum
func GetUpdateOccMetricAlarmDetailsThresholdTypeEnumStringValues() []string {
	return []string{
		"PERCENTAGE",
		"UNITS",
	}
}

// GetMappingUpdateOccMetricAlarmDetailsThresholdTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateOccMetricAlarmDetailsThresholdTypeEnum(val string) (UpdateOccMetricAlarmDetailsThresholdTypeEnum, bool) {
	enum, ok := mappingUpdateOccMetricAlarmDetailsThresholdTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
