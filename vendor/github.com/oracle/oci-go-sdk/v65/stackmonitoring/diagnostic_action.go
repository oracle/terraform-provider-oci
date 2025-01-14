// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiagnosticAction Diagnostic request model.
type DiagnosticAction interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Diagnostic Action
	GetId() *string

	// The current state of the Resource.
	GetLifecycleState() LifecycleStateEnum

	// The time the diagnostic action was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The resource to be monitored.
	GetMonitoredResourceId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The time the diagnostic action was last updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type diagnosticaction struct {
	JsonData             []byte
	DisplayName          *string                           `mandatory:"false" json:"displayName"`
	TimeUpdated          *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	FreeformTags         map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags          map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags           map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	CompartmentId        *string                           `mandatory:"true" json:"compartmentId"`
	Id                   *string                           `mandatory:"true" json:"id"`
	LifecycleState       LifecycleStateEnum                `mandatory:"true" json:"lifecycleState"`
	TimeCreated          *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	MonitoredResourceId  *string                           `mandatory:"true" json:"monitoredResourceId"`
	DiagnosticActionType string                            `json:"diagnosticActionType"`
}

// UnmarshalJSON unmarshals json
func (m *diagnosticaction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdiagnosticaction diagnosticaction
	s := struct {
		Model Unmarshalerdiagnosticaction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.Id = s.Model.Id
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.MonitoredResourceId = s.Model.MonitoredResourceId
	m.DisplayName = s.Model.DisplayName
	m.TimeUpdated = s.Model.TimeUpdated
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.DiagnosticActionType = s.Model.DiagnosticActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *diagnosticaction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DiagnosticActionType {
	case "TOP_PROCESS":
		mm := DiagnosticActionTopProcessDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DiagnosticAction: %s.", m.DiagnosticActionType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m diagnosticaction) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeUpdated returns TimeUpdated
func (m diagnosticaction) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m diagnosticaction) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m diagnosticaction) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m diagnosticaction) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetCompartmentId returns CompartmentId
func (m diagnosticaction) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetId returns Id
func (m diagnosticaction) GetId() *string {
	return m.Id
}

// GetLifecycleState returns LifecycleState
func (m diagnosticaction) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m diagnosticaction) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetMonitoredResourceId returns MonitoredResourceId
func (m diagnosticaction) GetMonitoredResourceId() *string {
	return m.MonitoredResourceId
}

func (m diagnosticaction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m diagnosticaction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiagnosticActionDiagnosticActionTypeEnum Enum with underlying type: string
type DiagnosticActionDiagnosticActionTypeEnum string

// Set of constants representing the allowable values for DiagnosticActionDiagnosticActionTypeEnum
const (
	DiagnosticActionDiagnosticActionTypeTopProcess DiagnosticActionDiagnosticActionTypeEnum = "TOP_PROCESS"
)

var mappingDiagnosticActionDiagnosticActionTypeEnum = map[string]DiagnosticActionDiagnosticActionTypeEnum{
	"TOP_PROCESS": DiagnosticActionDiagnosticActionTypeTopProcess,
}

var mappingDiagnosticActionDiagnosticActionTypeEnumLowerCase = map[string]DiagnosticActionDiagnosticActionTypeEnum{
	"top_process": DiagnosticActionDiagnosticActionTypeTopProcess,
}

// GetDiagnosticActionDiagnosticActionTypeEnumValues Enumerates the set of values for DiagnosticActionDiagnosticActionTypeEnum
func GetDiagnosticActionDiagnosticActionTypeEnumValues() []DiagnosticActionDiagnosticActionTypeEnum {
	values := make([]DiagnosticActionDiagnosticActionTypeEnum, 0)
	for _, v := range mappingDiagnosticActionDiagnosticActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiagnosticActionDiagnosticActionTypeEnumStringValues Enumerates the set of values in String for DiagnosticActionDiagnosticActionTypeEnum
func GetDiagnosticActionDiagnosticActionTypeEnumStringValues() []string {
	return []string{
		"TOP_PROCESS",
	}
}

// GetMappingDiagnosticActionDiagnosticActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiagnosticActionDiagnosticActionTypeEnum(val string) (DiagnosticActionDiagnosticActionTypeEnum, bool) {
	enum, ok := mappingDiagnosticActionDiagnosticActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
