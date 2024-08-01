// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetricData, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.cloud.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AlarmSuppressionSummary A summary of properties for the specified alarm suppression.
type AlarmSuppressionSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm suppression.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the alarm suppression.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	AlarmSuppressionTarget AlarmSuppressionTarget `mandatory:"true" json:"alarmSuppressionTarget"`

	// The level of this alarm suppression.
	// `ALARM` indicates a suppression of the entire alarm, regardless of dimension.
	// `DIMENSION` indicates a suppression configured for specified dimensions.
	Level AlarmSuppressionLevelEnum `mandatory:"true" json:"level"`

	// A user-friendly name for the alarm suppression. It does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The start date and time for the suppression to take place, inclusive. Format defined by RFC3339.
	// Example: `2023-02-01T01:02:29.600Z`
	TimeSuppressFrom *common.SDKTime `mandatory:"true" json:"timeSuppressFrom"`

	// The end date and time for the suppression to take place, inclusive. Format defined by RFC3339.
	// Example: `2023-02-01T02:02:29.600Z`
	TimeSuppressUntil *common.SDKTime `mandatory:"true" json:"timeSuppressUntil"`

	// The current lifecycle state of the alarm suppression.
	// Example: `DELETED`
	LifecycleState AlarmSuppressionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the alarm suppression was created. Format defined by RFC3339.
	// Example: `2023-02-01T01:02:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the alarm suppression was last updated (deleted). Format defined by RFC3339.
	// Example: `2023-02-03T01:02:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Array of all preconditions for alarm suppression.
	// Example: `[{
	//   conditionType: "RECURRENCE",
	//   suppressionRecurrence: "FRQ=DAILY;BYHOUR=10",
	//   suppressionDuration: "PT1H"
	// }]`
	SuppressionConditions []SuppressionCondition `mandatory:"false" json:"suppressionConditions"`

	// Human-readable reason for this alarm suppression.
	// It does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Oracle recommends including tracking information for the event or associated work,
	// such as a ticket number.
	// Example: `Planned outage due to change IT-1234.`
	Description *string `mandatory:"false" json:"description"`

	// Configured dimension filter for suppressing alarm state entries that include the set of specified dimension key-value pairs.
	// Example: `{"resourceId": "ocid1.instance.region1.phx.exampleuniqueID"}`
	Dimensions map[string]string `mandatory:"false" json:"dimensions"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AlarmSuppressionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlarmSuppressionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlarmSuppressionLevelEnum(string(m.Level)); !ok && m.Level != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Level: %s. Supported values are: %s.", m.Level, strings.Join(GetAlarmSuppressionLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAlarmSuppressionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAlarmSuppressionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AlarmSuppressionSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SuppressionConditions  []suppressioncondition             `json:"suppressionConditions"`
		Description            *string                            `json:"description"`
		Dimensions             map[string]string                  `json:"dimensions"`
		FreeformTags           map[string]string                  `json:"freeformTags"`
		DefinedTags            map[string]map[string]interface{}  `json:"definedTags"`
		Id                     *string                            `json:"id"`
		CompartmentId          *string                            `json:"compartmentId"`
		AlarmSuppressionTarget alarmsuppressiontarget             `json:"alarmSuppressionTarget"`
		Level                  AlarmSuppressionLevelEnum          `json:"level"`
		DisplayName            *string                            `json:"displayName"`
		TimeSuppressFrom       *common.SDKTime                    `json:"timeSuppressFrom"`
		TimeSuppressUntil      *common.SDKTime                    `json:"timeSuppressUntil"`
		LifecycleState         AlarmSuppressionLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated            *common.SDKTime                    `json:"timeCreated"`
		TimeUpdated            *common.SDKTime                    `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SuppressionConditions = make([]SuppressionCondition, len(model.SuppressionConditions))
	for i, n := range model.SuppressionConditions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.SuppressionConditions[i] = nn.(SuppressionCondition)
		} else {
			m.SuppressionConditions[i] = nil
		}
	}
	m.Description = model.Description

	m.Dimensions = model.Dimensions

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	nn, e = model.AlarmSuppressionTarget.UnmarshalPolymorphicJSON(model.AlarmSuppressionTarget.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AlarmSuppressionTarget = nn.(AlarmSuppressionTarget)
	} else {
		m.AlarmSuppressionTarget = nil
	}

	m.Level = model.Level

	m.DisplayName = model.DisplayName

	m.TimeSuppressFrom = model.TimeSuppressFrom

	m.TimeSuppressUntil = model.TimeSuppressUntil

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}
