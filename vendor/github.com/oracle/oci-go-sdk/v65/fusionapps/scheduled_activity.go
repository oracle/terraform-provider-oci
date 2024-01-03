// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduledActivity Details of scheduled activity.
type ScheduledActivity struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// scheduled activity display name, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// run cadence.
	RunCycle ScheduledActivityRunCycleEnum `mandatory:"true" json:"runCycle"`

	// FAaaS Environment Identifier.
	FusionEnvironmentId *string `mandatory:"true" json:"fusionEnvironmentId"`

	// The current state of the scheduledActivity.
	LifecycleState ScheduledActivityLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Service availability / impact during scheduled activity execution up down
	ServiceAvailability ScheduledActivityServiceAvailabilityEnum `mandatory:"true" json:"serviceAvailability"`

	// Current time the scheduled activity is scheduled to start. An RFC3339 formatted datetime string.
	TimeScheduledStart *common.SDKTime `mandatory:"true" json:"timeScheduledStart"`

	// Current time the scheduled activity is scheduled to end. An RFC3339 formatted datetime string.
	TimeExpectedFinish *common.SDKTime `mandatory:"true" json:"timeExpectedFinish"`

	// A property describing the phase of the scheduled activity.
	ScheduledActivityPhase ScheduledActivityScheduledActivityPhaseEnum `mandatory:"true" json:"scheduledActivityPhase"`

	// The unique identifier that associates a scheduled activity with others in one complete maintenance. For example, with ZDT, a complete upgrade maintenance includes 5 scheduled activities - PREPARE, EXECUTE, POST, PRE_MAINTENANCE, and POST_MAINTENANCE. All of them share the same unique identifier - scheduledActivityAssociationId.
	ScheduledActivityAssociationId *string `mandatory:"true" json:"scheduledActivityAssociationId"`

	// List of actions
	Actions []Action `mandatory:"false" json:"actions"`

	// The time the scheduled activity actually completed / cancelled / failed. An RFC3339 formatted datetime string.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Cumulative delay hours
	DelayInHours *int `mandatory:"false" json:"delayInHours"`

	// The time the scheduled activity record was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the scheduled activity record was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails ScheduledActivityLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

func (m ScheduledActivity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledActivity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduledActivityRunCycleEnum(string(m.RunCycle)); !ok && m.RunCycle != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RunCycle: %s. Supported values are: %s.", m.RunCycle, strings.Join(GetScheduledActivityRunCycleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledActivityLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduledActivityLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledActivityServiceAvailabilityEnum(string(m.ServiceAvailability)); !ok && m.ServiceAvailability != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceAvailability: %s. Supported values are: %s.", m.ServiceAvailability, strings.Join(GetScheduledActivityServiceAvailabilityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledActivityScheduledActivityPhaseEnum(string(m.ScheduledActivityPhase)); !ok && m.ScheduledActivityPhase != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduledActivityPhase: %s. Supported values are: %s.", m.ScheduledActivityPhase, strings.Join(GetScheduledActivityScheduledActivityPhaseEnumStringValues(), ",")))
	}

	if _, ok := GetMappingScheduledActivityLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetScheduledActivityLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ScheduledActivity) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Actions                        []action                                    `json:"actions"`
		TimeFinished                   *common.SDKTime                             `json:"timeFinished"`
		DelayInHours                   *int                                        `json:"delayInHours"`
		TimeCreated                    *common.SDKTime                             `json:"timeCreated"`
		TimeUpdated                    *common.SDKTime                             `json:"timeUpdated"`
		LifecycleDetails               ScheduledActivityLifecycleDetailsEnum       `json:"lifecycleDetails"`
		Id                             *string                                     `json:"id"`
		DisplayName                    *string                                     `json:"displayName"`
		RunCycle                       ScheduledActivityRunCycleEnum               `json:"runCycle"`
		FusionEnvironmentId            *string                                     `json:"fusionEnvironmentId"`
		LifecycleState                 ScheduledActivityLifecycleStateEnum         `json:"lifecycleState"`
		ServiceAvailability            ScheduledActivityServiceAvailabilityEnum    `json:"serviceAvailability"`
		TimeScheduledStart             *common.SDKTime                             `json:"timeScheduledStart"`
		TimeExpectedFinish             *common.SDKTime                             `json:"timeExpectedFinish"`
		ScheduledActivityPhase         ScheduledActivityScheduledActivityPhaseEnum `json:"scheduledActivityPhase"`
		ScheduledActivityAssociationId *string                                     `json:"scheduledActivityAssociationId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Actions = make([]Action, len(model.Actions))
	for i, n := range model.Actions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Actions[i] = nn.(Action)
		} else {
			m.Actions[i] = nil
		}
	}
	m.TimeFinished = model.TimeFinished

	m.DelayInHours = model.DelayInHours

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.RunCycle = model.RunCycle

	m.FusionEnvironmentId = model.FusionEnvironmentId

	m.LifecycleState = model.LifecycleState

	m.ServiceAvailability = model.ServiceAvailability

	m.TimeScheduledStart = model.TimeScheduledStart

	m.TimeExpectedFinish = model.TimeExpectedFinish

	m.ScheduledActivityPhase = model.ScheduledActivityPhase

	m.ScheduledActivityAssociationId = model.ScheduledActivityAssociationId

	return
}

// ScheduledActivityRunCycleEnum Enum with underlying type: string
type ScheduledActivityRunCycleEnum string

// Set of constants representing the allowable values for ScheduledActivityRunCycleEnum
const (
	ScheduledActivityRunCycleQuarterly ScheduledActivityRunCycleEnum = "QUARTERLY"
	ScheduledActivityRunCycleMonthly   ScheduledActivityRunCycleEnum = "MONTHLY"
	ScheduledActivityRunCycleOneoff    ScheduledActivityRunCycleEnum = "ONEOFF"
	ScheduledActivityRunCycleVertex    ScheduledActivityRunCycleEnum = "VERTEX"
)

var mappingScheduledActivityRunCycleEnum = map[string]ScheduledActivityRunCycleEnum{
	"QUARTERLY": ScheduledActivityRunCycleQuarterly,
	"MONTHLY":   ScheduledActivityRunCycleMonthly,
	"ONEOFF":    ScheduledActivityRunCycleOneoff,
	"VERTEX":    ScheduledActivityRunCycleVertex,
}

var mappingScheduledActivityRunCycleEnumLowerCase = map[string]ScheduledActivityRunCycleEnum{
	"quarterly": ScheduledActivityRunCycleQuarterly,
	"monthly":   ScheduledActivityRunCycleMonthly,
	"oneoff":    ScheduledActivityRunCycleOneoff,
	"vertex":    ScheduledActivityRunCycleVertex,
}

// GetScheduledActivityRunCycleEnumValues Enumerates the set of values for ScheduledActivityRunCycleEnum
func GetScheduledActivityRunCycleEnumValues() []ScheduledActivityRunCycleEnum {
	values := make([]ScheduledActivityRunCycleEnum, 0)
	for _, v := range mappingScheduledActivityRunCycleEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledActivityRunCycleEnumStringValues Enumerates the set of values in String for ScheduledActivityRunCycleEnum
func GetScheduledActivityRunCycleEnumStringValues() []string {
	return []string{
		"QUARTERLY",
		"MONTHLY",
		"ONEOFF",
		"VERTEX",
	}
}

// GetMappingScheduledActivityRunCycleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledActivityRunCycleEnum(val string) (ScheduledActivityRunCycleEnum, bool) {
	enum, ok := mappingScheduledActivityRunCycleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduledActivityLifecycleStateEnum Enum with underlying type: string
type ScheduledActivityLifecycleStateEnum string

// Set of constants representing the allowable values for ScheduledActivityLifecycleStateEnum
const (
	ScheduledActivityLifecycleStateAccepted   ScheduledActivityLifecycleStateEnum = "ACCEPTED"
	ScheduledActivityLifecycleStateInProgress ScheduledActivityLifecycleStateEnum = "IN_PROGRESS"
	ScheduledActivityLifecycleStateFailed     ScheduledActivityLifecycleStateEnum = "FAILED"
	ScheduledActivityLifecycleStateSucceeded  ScheduledActivityLifecycleStateEnum = "SUCCEEDED"
	ScheduledActivityLifecycleStateCanceled   ScheduledActivityLifecycleStateEnum = "CANCELED"
)

var mappingScheduledActivityLifecycleStateEnum = map[string]ScheduledActivityLifecycleStateEnum{
	"ACCEPTED":    ScheduledActivityLifecycleStateAccepted,
	"IN_PROGRESS": ScheduledActivityLifecycleStateInProgress,
	"FAILED":      ScheduledActivityLifecycleStateFailed,
	"SUCCEEDED":   ScheduledActivityLifecycleStateSucceeded,
	"CANCELED":    ScheduledActivityLifecycleStateCanceled,
}

var mappingScheduledActivityLifecycleStateEnumLowerCase = map[string]ScheduledActivityLifecycleStateEnum{
	"accepted":    ScheduledActivityLifecycleStateAccepted,
	"in_progress": ScheduledActivityLifecycleStateInProgress,
	"failed":      ScheduledActivityLifecycleStateFailed,
	"succeeded":   ScheduledActivityLifecycleStateSucceeded,
	"canceled":    ScheduledActivityLifecycleStateCanceled,
}

// GetScheduledActivityLifecycleStateEnumValues Enumerates the set of values for ScheduledActivityLifecycleStateEnum
func GetScheduledActivityLifecycleStateEnumValues() []ScheduledActivityLifecycleStateEnum {
	values := make([]ScheduledActivityLifecycleStateEnum, 0)
	for _, v := range mappingScheduledActivityLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledActivityLifecycleStateEnumStringValues Enumerates the set of values in String for ScheduledActivityLifecycleStateEnum
func GetScheduledActivityLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELED",
	}
}

// GetMappingScheduledActivityLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledActivityLifecycleStateEnum(val string) (ScheduledActivityLifecycleStateEnum, bool) {
	enum, ok := mappingScheduledActivityLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduledActivityServiceAvailabilityEnum Enum with underlying type: string
type ScheduledActivityServiceAvailabilityEnum string

// Set of constants representing the allowable values for ScheduledActivityServiceAvailabilityEnum
const (
	ScheduledActivityServiceAvailabilityAvailable   ScheduledActivityServiceAvailabilityEnum = "AVAILABLE"
	ScheduledActivityServiceAvailabilityUnavailable ScheduledActivityServiceAvailabilityEnum = "UNAVAILABLE"
)

var mappingScheduledActivityServiceAvailabilityEnum = map[string]ScheduledActivityServiceAvailabilityEnum{
	"AVAILABLE":   ScheduledActivityServiceAvailabilityAvailable,
	"UNAVAILABLE": ScheduledActivityServiceAvailabilityUnavailable,
}

var mappingScheduledActivityServiceAvailabilityEnumLowerCase = map[string]ScheduledActivityServiceAvailabilityEnum{
	"available":   ScheduledActivityServiceAvailabilityAvailable,
	"unavailable": ScheduledActivityServiceAvailabilityUnavailable,
}

// GetScheduledActivityServiceAvailabilityEnumValues Enumerates the set of values for ScheduledActivityServiceAvailabilityEnum
func GetScheduledActivityServiceAvailabilityEnumValues() []ScheduledActivityServiceAvailabilityEnum {
	values := make([]ScheduledActivityServiceAvailabilityEnum, 0)
	for _, v := range mappingScheduledActivityServiceAvailabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledActivityServiceAvailabilityEnumStringValues Enumerates the set of values in String for ScheduledActivityServiceAvailabilityEnum
func GetScheduledActivityServiceAvailabilityEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"UNAVAILABLE",
	}
}

// GetMappingScheduledActivityServiceAvailabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledActivityServiceAvailabilityEnum(val string) (ScheduledActivityServiceAvailabilityEnum, bool) {
	enum, ok := mappingScheduledActivityServiceAvailabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduledActivityLifecycleDetailsEnum Enum with underlying type: string
type ScheduledActivityLifecycleDetailsEnum string

// Set of constants representing the allowable values for ScheduledActivityLifecycleDetailsEnum
const (
	ScheduledActivityLifecycleDetailsNone               ScheduledActivityLifecycleDetailsEnum = "NONE"
	ScheduledActivityLifecycleDetailsRollbackaccepted   ScheduledActivityLifecycleDetailsEnum = "ROLLBACKACCEPTED"
	ScheduledActivityLifecycleDetailsRollbackinprogress ScheduledActivityLifecycleDetailsEnum = "ROLLBACKINPROGRESS"
	ScheduledActivityLifecycleDetailsRollbacksucceeded  ScheduledActivityLifecycleDetailsEnum = "ROLLBACKSUCCEEDED"
	ScheduledActivityLifecycleDetailsRollbackfailed     ScheduledActivityLifecycleDetailsEnum = "ROLLBACKFAILED"
)

var mappingScheduledActivityLifecycleDetailsEnum = map[string]ScheduledActivityLifecycleDetailsEnum{
	"NONE":               ScheduledActivityLifecycleDetailsNone,
	"ROLLBACKACCEPTED":   ScheduledActivityLifecycleDetailsRollbackaccepted,
	"ROLLBACKINPROGRESS": ScheduledActivityLifecycleDetailsRollbackinprogress,
	"ROLLBACKSUCCEEDED":  ScheduledActivityLifecycleDetailsRollbacksucceeded,
	"ROLLBACKFAILED":     ScheduledActivityLifecycleDetailsRollbackfailed,
}

var mappingScheduledActivityLifecycleDetailsEnumLowerCase = map[string]ScheduledActivityLifecycleDetailsEnum{
	"none":               ScheduledActivityLifecycleDetailsNone,
	"rollbackaccepted":   ScheduledActivityLifecycleDetailsRollbackaccepted,
	"rollbackinprogress": ScheduledActivityLifecycleDetailsRollbackinprogress,
	"rollbacksucceeded":  ScheduledActivityLifecycleDetailsRollbacksucceeded,
	"rollbackfailed":     ScheduledActivityLifecycleDetailsRollbackfailed,
}

// GetScheduledActivityLifecycleDetailsEnumValues Enumerates the set of values for ScheduledActivityLifecycleDetailsEnum
func GetScheduledActivityLifecycleDetailsEnumValues() []ScheduledActivityLifecycleDetailsEnum {
	values := make([]ScheduledActivityLifecycleDetailsEnum, 0)
	for _, v := range mappingScheduledActivityLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledActivityLifecycleDetailsEnumStringValues Enumerates the set of values in String for ScheduledActivityLifecycleDetailsEnum
func GetScheduledActivityLifecycleDetailsEnumStringValues() []string {
	return []string{
		"NONE",
		"ROLLBACKACCEPTED",
		"ROLLBACKINPROGRESS",
		"ROLLBACKSUCCEEDED",
		"ROLLBACKFAILED",
	}
}

// GetMappingScheduledActivityLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledActivityLifecycleDetailsEnum(val string) (ScheduledActivityLifecycleDetailsEnum, bool) {
	enum, ok := mappingScheduledActivityLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduledActivityScheduledActivityPhaseEnum Enum with underlying type: string
type ScheduledActivityScheduledActivityPhaseEnum string

// Set of constants representing the allowable values for ScheduledActivityScheduledActivityPhaseEnum
const (
	ScheduledActivityScheduledActivityPhasePreMaintenance  ScheduledActivityScheduledActivityPhaseEnum = "PRE_MAINTENANCE"
	ScheduledActivityScheduledActivityPhaseMaintenance     ScheduledActivityScheduledActivityPhaseEnum = "MAINTENANCE"
	ScheduledActivityScheduledActivityPhasePostMaintenance ScheduledActivityScheduledActivityPhaseEnum = "POST_MAINTENANCE"
)

var mappingScheduledActivityScheduledActivityPhaseEnum = map[string]ScheduledActivityScheduledActivityPhaseEnum{
	"PRE_MAINTENANCE":  ScheduledActivityScheduledActivityPhasePreMaintenance,
	"MAINTENANCE":      ScheduledActivityScheduledActivityPhaseMaintenance,
	"POST_MAINTENANCE": ScheduledActivityScheduledActivityPhasePostMaintenance,
}

var mappingScheduledActivityScheduledActivityPhaseEnumLowerCase = map[string]ScheduledActivityScheduledActivityPhaseEnum{
	"pre_maintenance":  ScheduledActivityScheduledActivityPhasePreMaintenance,
	"maintenance":      ScheduledActivityScheduledActivityPhaseMaintenance,
	"post_maintenance": ScheduledActivityScheduledActivityPhasePostMaintenance,
}

// GetScheduledActivityScheduledActivityPhaseEnumValues Enumerates the set of values for ScheduledActivityScheduledActivityPhaseEnum
func GetScheduledActivityScheduledActivityPhaseEnumValues() []ScheduledActivityScheduledActivityPhaseEnum {
	values := make([]ScheduledActivityScheduledActivityPhaseEnum, 0)
	for _, v := range mappingScheduledActivityScheduledActivityPhaseEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledActivityScheduledActivityPhaseEnumStringValues Enumerates the set of values in String for ScheduledActivityScheduledActivityPhaseEnum
func GetScheduledActivityScheduledActivityPhaseEnumStringValues() []string {
	return []string{
		"PRE_MAINTENANCE",
		"MAINTENANCE",
		"POST_MAINTENANCE",
	}
}

// GetMappingScheduledActivityScheduledActivityPhaseEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledActivityScheduledActivityPhaseEnum(val string) (ScheduledActivityScheduledActivityPhaseEnum, bool) {
	enum, ok := mappingScheduledActivityScheduledActivityPhaseEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
