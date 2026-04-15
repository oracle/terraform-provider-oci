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

// CreateOccMetricAlarmDetails Payload for creating an OccMetricAlarm.
type CreateOccMetricAlarmDetails struct {

	// Human-readable name for the alarm.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment OCID in which the alarm is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Alarm active status.
	IsActive *bool `mandatory:"true" json:"isActive"`

	// Frequency at which notifications should be sent.
	Frequency OccMetricAlarmFrequencyEnum `mandatory:"true" json:"frequency"`

	// Threshold at which alarm must be triggered.
	Threshold *int `mandatory:"true" json:"threshold"`

	ResourceConfiguration BaseResourceConfiguration `mandatory:"true" json:"resourceConfiguration"`

	// Optional description for the alarm.
	Description *string `mandatory:"false" json:"description"`

	// The current lifecycle state of the resource.
	LifecycleState CreateOccMetricAlarmDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// List of topic OCIDs for notifications.
	Subscribers []string `mandatory:"false" json:"subscribers"`

	// Units in which threshold is being stored.
	ThresholdType CreateOccMetricAlarmDetailsThresholdTypeEnum `mandatory:"false" json:"thresholdType,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOccMetricAlarmDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOccMetricAlarmDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccMetricAlarmFrequencyEnum(string(m.Frequency)); !ok && m.Frequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Frequency: %s. Supported values are: %s.", m.Frequency, strings.Join(GetOccMetricAlarmFrequencyEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCreateOccMetricAlarmDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCreateOccMetricAlarmDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateOccMetricAlarmDetailsThresholdTypeEnum(string(m.ThresholdType)); !ok && m.ThresholdType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ThresholdType: %s. Supported values are: %s.", m.ThresholdType, strings.Join(GetCreateOccMetricAlarmDetailsThresholdTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateOccMetricAlarmDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description           *string                                       `json:"description"`
		LifecycleState        CreateOccMetricAlarmDetailsLifecycleStateEnum `json:"lifecycleState"`
		Subscribers           []string                                      `json:"subscribers"`
		ThresholdType         CreateOccMetricAlarmDetailsThresholdTypeEnum  `json:"thresholdType"`
		FreeformTags          map[string]string                             `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{}             `json:"definedTags"`
		DisplayName           *string                                       `json:"displayName"`
		CompartmentId         *string                                       `json:"compartmentId"`
		IsActive              *bool                                         `json:"isActive"`
		Frequency             OccMetricAlarmFrequencyEnum                   `json:"frequency"`
		Threshold             *int                                          `json:"threshold"`
		ResourceConfiguration baseresourceconfiguration                     `json:"resourceConfiguration"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.LifecycleState = model.LifecycleState

	m.Subscribers = make([]string, len(model.Subscribers))
	copy(m.Subscribers, model.Subscribers)
	m.ThresholdType = model.ThresholdType

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.IsActive = model.IsActive

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

// CreateOccMetricAlarmDetailsLifecycleStateEnum Enum with underlying type: string
type CreateOccMetricAlarmDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for CreateOccMetricAlarmDetailsLifecycleStateEnum
const (
	CreateOccMetricAlarmDetailsLifecycleStateCreating CreateOccMetricAlarmDetailsLifecycleStateEnum = "CREATING"
	CreateOccMetricAlarmDetailsLifecycleStateUpdating CreateOccMetricAlarmDetailsLifecycleStateEnum = "UPDATING"
	CreateOccMetricAlarmDetailsLifecycleStateActive   CreateOccMetricAlarmDetailsLifecycleStateEnum = "ACTIVE"
	CreateOccMetricAlarmDetailsLifecycleStateDeleting CreateOccMetricAlarmDetailsLifecycleStateEnum = "DELETING"
	CreateOccMetricAlarmDetailsLifecycleStateDeleted  CreateOccMetricAlarmDetailsLifecycleStateEnum = "DELETED"
	CreateOccMetricAlarmDetailsLifecycleStateFailed   CreateOccMetricAlarmDetailsLifecycleStateEnum = "FAILED"
)

var mappingCreateOccMetricAlarmDetailsLifecycleStateEnum = map[string]CreateOccMetricAlarmDetailsLifecycleStateEnum{
	"CREATING": CreateOccMetricAlarmDetailsLifecycleStateCreating,
	"UPDATING": CreateOccMetricAlarmDetailsLifecycleStateUpdating,
	"ACTIVE":   CreateOccMetricAlarmDetailsLifecycleStateActive,
	"DELETING": CreateOccMetricAlarmDetailsLifecycleStateDeleting,
	"DELETED":  CreateOccMetricAlarmDetailsLifecycleStateDeleted,
	"FAILED":   CreateOccMetricAlarmDetailsLifecycleStateFailed,
}

var mappingCreateOccMetricAlarmDetailsLifecycleStateEnumLowerCase = map[string]CreateOccMetricAlarmDetailsLifecycleStateEnum{
	"creating": CreateOccMetricAlarmDetailsLifecycleStateCreating,
	"updating": CreateOccMetricAlarmDetailsLifecycleStateUpdating,
	"active":   CreateOccMetricAlarmDetailsLifecycleStateActive,
	"deleting": CreateOccMetricAlarmDetailsLifecycleStateDeleting,
	"deleted":  CreateOccMetricAlarmDetailsLifecycleStateDeleted,
	"failed":   CreateOccMetricAlarmDetailsLifecycleStateFailed,
}

// GetCreateOccMetricAlarmDetailsLifecycleStateEnumValues Enumerates the set of values for CreateOccMetricAlarmDetailsLifecycleStateEnum
func GetCreateOccMetricAlarmDetailsLifecycleStateEnumValues() []CreateOccMetricAlarmDetailsLifecycleStateEnum {
	values := make([]CreateOccMetricAlarmDetailsLifecycleStateEnum, 0)
	for _, v := range mappingCreateOccMetricAlarmDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOccMetricAlarmDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for CreateOccMetricAlarmDetailsLifecycleStateEnum
func GetCreateOccMetricAlarmDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCreateOccMetricAlarmDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOccMetricAlarmDetailsLifecycleStateEnum(val string) (CreateOccMetricAlarmDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingCreateOccMetricAlarmDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateOccMetricAlarmDetailsThresholdTypeEnum Enum with underlying type: string
type CreateOccMetricAlarmDetailsThresholdTypeEnum string

// Set of constants representing the allowable values for CreateOccMetricAlarmDetailsThresholdTypeEnum
const (
	CreateOccMetricAlarmDetailsThresholdTypePercentage CreateOccMetricAlarmDetailsThresholdTypeEnum = "PERCENTAGE"
	CreateOccMetricAlarmDetailsThresholdTypeUnits      CreateOccMetricAlarmDetailsThresholdTypeEnum = "UNITS"
)

var mappingCreateOccMetricAlarmDetailsThresholdTypeEnum = map[string]CreateOccMetricAlarmDetailsThresholdTypeEnum{
	"PERCENTAGE": CreateOccMetricAlarmDetailsThresholdTypePercentage,
	"UNITS":      CreateOccMetricAlarmDetailsThresholdTypeUnits,
}

var mappingCreateOccMetricAlarmDetailsThresholdTypeEnumLowerCase = map[string]CreateOccMetricAlarmDetailsThresholdTypeEnum{
	"percentage": CreateOccMetricAlarmDetailsThresholdTypePercentage,
	"units":      CreateOccMetricAlarmDetailsThresholdTypeUnits,
}

// GetCreateOccMetricAlarmDetailsThresholdTypeEnumValues Enumerates the set of values for CreateOccMetricAlarmDetailsThresholdTypeEnum
func GetCreateOccMetricAlarmDetailsThresholdTypeEnumValues() []CreateOccMetricAlarmDetailsThresholdTypeEnum {
	values := make([]CreateOccMetricAlarmDetailsThresholdTypeEnum, 0)
	for _, v := range mappingCreateOccMetricAlarmDetailsThresholdTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOccMetricAlarmDetailsThresholdTypeEnumStringValues Enumerates the set of values in String for CreateOccMetricAlarmDetailsThresholdTypeEnum
func GetCreateOccMetricAlarmDetailsThresholdTypeEnumStringValues() []string {
	return []string{
		"PERCENTAGE",
		"UNITS",
	}
}

// GetMappingCreateOccMetricAlarmDetailsThresholdTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOccMetricAlarmDetailsThresholdTypeEnum(val string) (CreateOccMetricAlarmDetailsThresholdTypeEnum, bool) {
	enum, ok := mappingCreateOccMetricAlarmDetailsThresholdTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
