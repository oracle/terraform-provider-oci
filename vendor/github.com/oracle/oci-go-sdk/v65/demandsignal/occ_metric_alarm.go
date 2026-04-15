// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Demand Signal API
//
// Use the OCI Control Center Demand Signal API to manage Demand Signals.
//

package demandsignal

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccMetricAlarm OccMetricAlarm represents an alarm configuration for OCI Control Center metric.
type OccMetricAlarm struct {

	// Unique OCID for this alarm configuration.
	Id *string `mandatory:"true" json:"id"`

	// Human-readable name for the alarm.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Creation timestamp (RFC 3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Last update timestamp (RFC 3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current lifecycle state of the resource.
	LifecycleState OccMetricAlarmLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Alarm active status.
	IsActive *bool `mandatory:"true" json:"isActive"`

	// List of topic OCIDs for notifications.
	Subscribers []string `mandatory:"true" json:"subscribers"`

	// Frequency at which notifications should be sent.
	Frequency OccMetricAlarmFrequencyEnum `mandatory:"true" json:"frequency"`

	// Threshold at which alarm must be triggered.
	Threshold *int `mandatory:"true" json:"threshold"`

	// Units in which threshold is being stored.
	ThresholdType OccMetricAlarmThresholdTypeEnum `mandatory:"true" json:"thresholdType"`

	ResourceConfiguration BaseResourceConfiguration `mandatory:"true" json:"resourceConfiguration"`

	// Optional description for the alarm.
	Description *string `mandatory:"false" json:"description"`

	// Compartment OCID in which the alarm is created.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OccMetricAlarm) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccMetricAlarm) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccMetricAlarmLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccMetricAlarmLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccMetricAlarmFrequencyEnum(string(m.Frequency)); !ok && m.Frequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Frequency: %s. Supported values are: %s.", m.Frequency, strings.Join(GetOccMetricAlarmFrequencyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccMetricAlarmThresholdTypeEnum(string(m.ThresholdType)); !ok && m.ThresholdType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ThresholdType: %s. Supported values are: %s.", m.ThresholdType, strings.Join(GetOccMetricAlarmThresholdTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *OccMetricAlarm) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description           *string                           `json:"description"`
		CompartmentId         *string                           `json:"compartmentId"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		SystemTags            map[string]map[string]interface{} `json:"systemTags"`
		Id                    *string                           `json:"id"`
		DisplayName           *string                           `json:"displayName"`
		TimeCreated           *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated           *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState        OccMetricAlarmLifecycleStateEnum  `json:"lifecycleState"`
		IsActive              *bool                             `json:"isActive"`
		Subscribers           []string                          `json:"subscribers"`
		Frequency             OccMetricAlarmFrequencyEnum       `json:"frequency"`
		Threshold             *int                              `json:"threshold"`
		ThresholdType         OccMetricAlarmThresholdTypeEnum   `json:"thresholdType"`
		ResourceConfiguration baseresourceconfiguration         `json:"resourceConfiguration"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.CompartmentId = model.CompartmentId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.IsActive = model.IsActive

	m.Subscribers = make([]string, len(model.Subscribers))
	copy(m.Subscribers, model.Subscribers)
	m.Frequency = model.Frequency

	m.Threshold = model.Threshold

	m.ThresholdType = model.ThresholdType

	nn, e = model.ResourceConfiguration.UnmarshalPolymorphicJSON(model.ResourceConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResourceConfiguration = nn.(BaseResourceConfiguration)
	} else {
		m.ResourceConfiguration = nil
	}

	return
}

// OccMetricAlarmLifecycleStateEnum Enum with underlying type: string
type OccMetricAlarmLifecycleStateEnum string

// Set of constants representing the allowable values for OccMetricAlarmLifecycleStateEnum
const (
	OccMetricAlarmLifecycleStateCreating OccMetricAlarmLifecycleStateEnum = "CREATING"
	OccMetricAlarmLifecycleStateUpdating OccMetricAlarmLifecycleStateEnum = "UPDATING"
	OccMetricAlarmLifecycleStateActive   OccMetricAlarmLifecycleStateEnum = "ACTIVE"
	OccMetricAlarmLifecycleStateDeleting OccMetricAlarmLifecycleStateEnum = "DELETING"
	OccMetricAlarmLifecycleStateDeleted  OccMetricAlarmLifecycleStateEnum = "DELETED"
	OccMetricAlarmLifecycleStateFailed   OccMetricAlarmLifecycleStateEnum = "FAILED"
)

var mappingOccMetricAlarmLifecycleStateEnum = map[string]OccMetricAlarmLifecycleStateEnum{
	"CREATING": OccMetricAlarmLifecycleStateCreating,
	"UPDATING": OccMetricAlarmLifecycleStateUpdating,
	"ACTIVE":   OccMetricAlarmLifecycleStateActive,
	"DELETING": OccMetricAlarmLifecycleStateDeleting,
	"DELETED":  OccMetricAlarmLifecycleStateDeleted,
	"FAILED":   OccMetricAlarmLifecycleStateFailed,
}

var mappingOccMetricAlarmLifecycleStateEnumLowerCase = map[string]OccMetricAlarmLifecycleStateEnum{
	"creating": OccMetricAlarmLifecycleStateCreating,
	"updating": OccMetricAlarmLifecycleStateUpdating,
	"active":   OccMetricAlarmLifecycleStateActive,
	"deleting": OccMetricAlarmLifecycleStateDeleting,
	"deleted":  OccMetricAlarmLifecycleStateDeleted,
	"failed":   OccMetricAlarmLifecycleStateFailed,
}

// GetOccMetricAlarmLifecycleStateEnumValues Enumerates the set of values for OccMetricAlarmLifecycleStateEnum
func GetOccMetricAlarmLifecycleStateEnumValues() []OccMetricAlarmLifecycleStateEnum {
	values := make([]OccMetricAlarmLifecycleStateEnum, 0)
	for _, v := range mappingOccMetricAlarmLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOccMetricAlarmLifecycleStateEnumStringValues Enumerates the set of values in String for OccMetricAlarmLifecycleStateEnum
func GetOccMetricAlarmLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOccMetricAlarmLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccMetricAlarmLifecycleStateEnum(val string) (OccMetricAlarmLifecycleStateEnum, bool) {
	enum, ok := mappingOccMetricAlarmLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OccMetricAlarmFrequencyEnum Enum with underlying type: string
type OccMetricAlarmFrequencyEnum string

// Set of constants representing the allowable values for OccMetricAlarmFrequencyEnum
const (
	OccMetricAlarmFrequencyDaily   OccMetricAlarmFrequencyEnum = "DAILY"
	OccMetricAlarmFrequencyWeekly  OccMetricAlarmFrequencyEnum = "WEEKLY"
	OccMetricAlarmFrequencyMonthly OccMetricAlarmFrequencyEnum = "MONTHLY"
)

var mappingOccMetricAlarmFrequencyEnum = map[string]OccMetricAlarmFrequencyEnum{
	"DAILY":   OccMetricAlarmFrequencyDaily,
	"WEEKLY":  OccMetricAlarmFrequencyWeekly,
	"MONTHLY": OccMetricAlarmFrequencyMonthly,
}

var mappingOccMetricAlarmFrequencyEnumLowerCase = map[string]OccMetricAlarmFrequencyEnum{
	"daily":   OccMetricAlarmFrequencyDaily,
	"weekly":  OccMetricAlarmFrequencyWeekly,
	"monthly": OccMetricAlarmFrequencyMonthly,
}

// GetOccMetricAlarmFrequencyEnumValues Enumerates the set of values for OccMetricAlarmFrequencyEnum
func GetOccMetricAlarmFrequencyEnumValues() []OccMetricAlarmFrequencyEnum {
	values := make([]OccMetricAlarmFrequencyEnum, 0)
	for _, v := range mappingOccMetricAlarmFrequencyEnum {
		values = append(values, v)
	}
	return values
}

// GetOccMetricAlarmFrequencyEnumStringValues Enumerates the set of values in String for OccMetricAlarmFrequencyEnum
func GetOccMetricAlarmFrequencyEnumStringValues() []string {
	return []string{
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
}

// GetMappingOccMetricAlarmFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccMetricAlarmFrequencyEnum(val string) (OccMetricAlarmFrequencyEnum, bool) {
	enum, ok := mappingOccMetricAlarmFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OccMetricAlarmThresholdTypeEnum Enum with underlying type: string
type OccMetricAlarmThresholdTypeEnum string

// Set of constants representing the allowable values for OccMetricAlarmThresholdTypeEnum
const (
	OccMetricAlarmThresholdTypePercentage OccMetricAlarmThresholdTypeEnum = "PERCENTAGE"
	OccMetricAlarmThresholdTypeUnits      OccMetricAlarmThresholdTypeEnum = "UNITS"
)

var mappingOccMetricAlarmThresholdTypeEnum = map[string]OccMetricAlarmThresholdTypeEnum{
	"PERCENTAGE": OccMetricAlarmThresholdTypePercentage,
	"UNITS":      OccMetricAlarmThresholdTypeUnits,
}

var mappingOccMetricAlarmThresholdTypeEnumLowerCase = map[string]OccMetricAlarmThresholdTypeEnum{
	"percentage": OccMetricAlarmThresholdTypePercentage,
	"units":      OccMetricAlarmThresholdTypeUnits,
}

// GetOccMetricAlarmThresholdTypeEnumValues Enumerates the set of values for OccMetricAlarmThresholdTypeEnum
func GetOccMetricAlarmThresholdTypeEnumValues() []OccMetricAlarmThresholdTypeEnum {
	values := make([]OccMetricAlarmThresholdTypeEnum, 0)
	for _, v := range mappingOccMetricAlarmThresholdTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOccMetricAlarmThresholdTypeEnumStringValues Enumerates the set of values in String for OccMetricAlarmThresholdTypeEnum
func GetOccMetricAlarmThresholdTypeEnumStringValues() []string {
	return []string{
		"PERCENTAGE",
		"UNITS",
	}
}

// GetMappingOccMetricAlarmThresholdTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccMetricAlarmThresholdTypeEnum(val string) (OccMetricAlarmThresholdTypeEnum, bool) {
	enum, ok := mappingOccMetricAlarmThresholdTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
