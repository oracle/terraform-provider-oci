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

// OccMetricAlarmSummary Summary of OccMetricAlarm.
type OccMetricAlarmSummary struct {

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

	// Compartment OCID in which the alarm is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Alarm active status.
	IsActive *bool `mandatory:"true" json:"isActive"`

	// List of topic OCIDs for notifications.
	Subscribers []string `mandatory:"true" json:"subscribers"`

	// Frequency at which notifications should be sent.
	Frequency OccMetricAlarmFrequencyEnum `mandatory:"true" json:"frequency"`

	// Threshold at which alarm must be triggered.
	Threshold *int `mandatory:"true" json:"threshold"`

	ResourceConfiguration BaseResourceConfiguration `mandatory:"true" json:"resourceConfiguration"`

	// Optional description for the alarm.
	Description *string `mandatory:"false" json:"description"`

	// Units in which threshold is being stored.
	ThresholdType OccMetricAlarmSummaryThresholdTypeEnum `mandatory:"false" json:"thresholdType,omitempty"`

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

func (m OccMetricAlarmSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccMetricAlarmSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccMetricAlarmLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccMetricAlarmLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccMetricAlarmFrequencyEnum(string(m.Frequency)); !ok && m.Frequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Frequency: %s. Supported values are: %s.", m.Frequency, strings.Join(GetOccMetricAlarmFrequencyEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOccMetricAlarmSummaryThresholdTypeEnum(string(m.ThresholdType)); !ok && m.ThresholdType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ThresholdType: %s. Supported values are: %s.", m.ThresholdType, strings.Join(GetOccMetricAlarmSummaryThresholdTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *OccMetricAlarmSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description           *string                                `json:"description"`
		ThresholdType         OccMetricAlarmSummaryThresholdTypeEnum `json:"thresholdType"`
		FreeformTags          map[string]string                      `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{}      `json:"definedTags"`
		SystemTags            map[string]map[string]interface{}      `json:"systemTags"`
		Id                    *string                                `json:"id"`
		DisplayName           *string                                `json:"displayName"`
		TimeCreated           *common.SDKTime                        `json:"timeCreated"`
		TimeUpdated           *common.SDKTime                        `json:"timeUpdated"`
		LifecycleState        OccMetricAlarmLifecycleStateEnum       `json:"lifecycleState"`
		CompartmentId         *string                                `json:"compartmentId"`
		IsActive              *bool                                  `json:"isActive"`
		Subscribers           []string                               `json:"subscribers"`
		Frequency             OccMetricAlarmFrequencyEnum            `json:"frequency"`
		Threshold             *int                                   `json:"threshold"`
		ResourceConfiguration baseresourceconfiguration              `json:"resourceConfiguration"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.ThresholdType = model.ThresholdType

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.CompartmentId = model.CompartmentId

	m.IsActive = model.IsActive

	m.Subscribers = make([]string, len(model.Subscribers))
	copy(m.Subscribers, model.Subscribers)
	m.Frequency = model.Frequency

	m.Threshold = model.Threshold

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

// OccMetricAlarmSummaryThresholdTypeEnum Enum with underlying type: string
type OccMetricAlarmSummaryThresholdTypeEnum string

// Set of constants representing the allowable values for OccMetricAlarmSummaryThresholdTypeEnum
const (
	OccMetricAlarmSummaryThresholdTypePercentage OccMetricAlarmSummaryThresholdTypeEnum = "PERCENTAGE"
	OccMetricAlarmSummaryThresholdTypeUnits      OccMetricAlarmSummaryThresholdTypeEnum = "UNITS"
)

var mappingOccMetricAlarmSummaryThresholdTypeEnum = map[string]OccMetricAlarmSummaryThresholdTypeEnum{
	"PERCENTAGE": OccMetricAlarmSummaryThresholdTypePercentage,
	"UNITS":      OccMetricAlarmSummaryThresholdTypeUnits,
}

var mappingOccMetricAlarmSummaryThresholdTypeEnumLowerCase = map[string]OccMetricAlarmSummaryThresholdTypeEnum{
	"percentage": OccMetricAlarmSummaryThresholdTypePercentage,
	"units":      OccMetricAlarmSummaryThresholdTypeUnits,
}

// GetOccMetricAlarmSummaryThresholdTypeEnumValues Enumerates the set of values for OccMetricAlarmSummaryThresholdTypeEnum
func GetOccMetricAlarmSummaryThresholdTypeEnumValues() []OccMetricAlarmSummaryThresholdTypeEnum {
	values := make([]OccMetricAlarmSummaryThresholdTypeEnum, 0)
	for _, v := range mappingOccMetricAlarmSummaryThresholdTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOccMetricAlarmSummaryThresholdTypeEnumStringValues Enumerates the set of values in String for OccMetricAlarmSummaryThresholdTypeEnum
func GetOccMetricAlarmSummaryThresholdTypeEnumStringValues() []string {
	return []string{
		"PERCENTAGE",
		"UNITS",
	}
}

// GetMappingOccMetricAlarmSummaryThresholdTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccMetricAlarmSummaryThresholdTypeEnum(val string) (OccMetricAlarmSummaryThresholdTypeEnum, bool) {
	enum, ok := mappingOccMetricAlarmSummaryThresholdTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
