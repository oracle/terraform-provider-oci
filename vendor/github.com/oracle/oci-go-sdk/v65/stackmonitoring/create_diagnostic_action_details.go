// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateDiagnosticActionDetails Diagnostic request model.
type CreateDiagnosticActionDetails interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The primary OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that will be used for the diagnostic action.
	GetMonitoredResourceId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type creatediagnosticactiondetails struct {
	JsonData             []byte
	DisplayName          *string                           `mandatory:"false" json:"displayName"`
	FreeformTags         map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags          map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	CompartmentId        *string                           `mandatory:"true" json:"compartmentId"`
	MonitoredResourceId  *string                           `mandatory:"true" json:"monitoredResourceId"`
	DiagnosticActionType string                            `json:"diagnosticActionType"`
}

// UnmarshalJSON unmarshals json
func (m *creatediagnosticactiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatediagnosticactiondetails creatediagnosticactiondetails
	s := struct {
		Model Unmarshalercreatediagnosticactiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.MonitoredResourceId = s.Model.MonitoredResourceId
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.DiagnosticActionType = s.Model.DiagnosticActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *creatediagnosticactiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DiagnosticActionType {
	case "TOP_PROCESS":
		mm := CreateDiagnosticActionTopProcessDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDiagnosticActionDetails: %s.", m.DiagnosticActionType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m creatediagnosticactiondetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m creatediagnosticactiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m creatediagnosticactiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetCompartmentId returns CompartmentId
func (m creatediagnosticactiondetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetMonitoredResourceId returns MonitoredResourceId
func (m creatediagnosticactiondetails) GetMonitoredResourceId() *string {
	return m.MonitoredResourceId
}

func (m creatediagnosticactiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m creatediagnosticactiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
