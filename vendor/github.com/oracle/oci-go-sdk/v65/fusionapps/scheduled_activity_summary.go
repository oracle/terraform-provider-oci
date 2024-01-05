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

// ScheduledActivitySummary Summary of the scheduled activity for a Fusion environment.
type ScheduledActivitySummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A friendly name for the scheduled activity. Can be changed later.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The run cadence of this scheduled activity. Valid values are Quarterly, Monthly, OneOff, and Vertex.
	RunCycle ScheduledActivityRunCycleEnum `mandatory:"true" json:"runCycle"`

	// The OCID of the Fusion environment for the scheduled activity.
	FusionEnvironmentId *string `mandatory:"true" json:"fusionEnvironmentId"`

	// The current state of the scheduled activity. Valid values are Scheduled, In progress , Failed, Completed.
	LifecycleState ScheduledActivityLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Current time the scheduled activity is scheduled to start. An RFC3339 formatted datetime string.
	TimeScheduledStart *common.SDKTime `mandatory:"true" json:"timeScheduledStart"`

	// Current time the scheduled activity is scheduled to end. An RFC3339 formatted datetime string.
	TimeExpectedFinish *common.SDKTime `mandatory:"true" json:"timeExpectedFinish"`

	// Service availability / impact during scheduled activity execution, up down
	ServiceAvailability ScheduledActivityServiceAvailabilityEnum `mandatory:"true" json:"serviceAvailability"`

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
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// The time the scheduled activity record was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ScheduledActivitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledActivitySummary) ValidateEnumValue() (bool, error) {
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

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ScheduledActivitySummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Actions                        []action                                    `json:"actions"`
		TimeFinished                   *common.SDKTime                             `json:"timeFinished"`
		DelayInHours                   *int                                        `json:"delayInHours"`
		TimeAccepted                   *common.SDKTime                             `json:"timeAccepted"`
		TimeUpdated                    *common.SDKTime                             `json:"timeUpdated"`
		LifecycleDetails               *string                                     `json:"lifecycleDetails"`
		FreeformTags                   map[string]string                           `json:"freeformTags"`
		DefinedTags                    map[string]map[string]interface{}           `json:"definedTags"`
		Id                             *string                                     `json:"id"`
		DisplayName                    *string                                     `json:"displayName"`
		RunCycle                       ScheduledActivityRunCycleEnum               `json:"runCycle"`
		FusionEnvironmentId            *string                                     `json:"fusionEnvironmentId"`
		LifecycleState                 ScheduledActivityLifecycleStateEnum         `json:"lifecycleState"`
		TimeScheduledStart             *common.SDKTime                             `json:"timeScheduledStart"`
		TimeExpectedFinish             *common.SDKTime                             `json:"timeExpectedFinish"`
		ServiceAvailability            ScheduledActivityServiceAvailabilityEnum    `json:"serviceAvailability"`
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

	m.TimeAccepted = model.TimeAccepted

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.RunCycle = model.RunCycle

	m.FusionEnvironmentId = model.FusionEnvironmentId

	m.LifecycleState = model.LifecycleState

	m.TimeScheduledStart = model.TimeScheduledStart

	m.TimeExpectedFinish = model.TimeExpectedFinish

	m.ServiceAvailability = model.ServiceAvailability

	m.ScheduledActivityPhase = model.ScheduledActivityPhase

	m.ScheduledActivityAssociationId = model.ScheduledActivityAssociationId

	return
}
