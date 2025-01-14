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

// DiagnosticActionTopProcessDetails Diagnostic request model for Top Process
type DiagnosticActionTopProcessDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Diagnostic Action
	Id *string `mandatory:"true" json:"id"`

	// The time the diagnostic action was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The resource to be monitored.
	MonitoredResourceId *string `mandatory:"true" json:"monitoredResourceId"`

	// number of process to return
	Limit *int `mandatory:"true" json:"limit"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the diagnostic action was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the Resource.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetCompartmentId returns CompartmentId
func (m DiagnosticActionTopProcessDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetId returns Id
func (m DiagnosticActionTopProcessDetails) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m DiagnosticActionTopProcessDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetLifecycleState returns LifecycleState
func (m DiagnosticActionTopProcessDetails) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m DiagnosticActionTopProcessDetails) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DiagnosticActionTopProcessDetails) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetMonitoredResourceId returns MonitoredResourceId
func (m DiagnosticActionTopProcessDetails) GetMonitoredResourceId() *string {
	return m.MonitoredResourceId
}

// GetFreeformTags returns FreeformTags
func (m DiagnosticActionTopProcessDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m DiagnosticActionTopProcessDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m DiagnosticActionTopProcessDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m DiagnosticActionTopProcessDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiagnosticActionTopProcessDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DiagnosticActionTopProcessDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDiagnosticActionTopProcessDetails DiagnosticActionTopProcessDetails
	s := struct {
		DiscriminatorParam string `json:"diagnosticActionType"`
		MarshalTypeDiagnosticActionTopProcessDetails
	}{
		"TOP_PROCESS",
		(MarshalTypeDiagnosticActionTopProcessDetails)(m),
	}

	return json.Marshal(&s)
}
