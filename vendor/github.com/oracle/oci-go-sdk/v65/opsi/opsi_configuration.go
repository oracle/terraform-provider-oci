// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OpsiConfiguration OPSI configuration.
type OpsiConfiguration interface {

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of OPSI configuration resource.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// User-friendly display name for the OPSI configuration. The name does not have to be unique.
	GetDisplayName() *string

	// Description of OPSI configuration.
	GetDescription() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// The time at which the resource was first created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time at which the resource was last updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// OPSI configuration resource lifecycle state.
	GetLifecycleState() OpsiConfigurationLifecycleStateEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// Array of configuration item summary objects.
	GetConfigItems() []OpsiConfigurationConfigurationItemSummary
}

type opsiconfiguration struct {
	JsonData         []byte
	Id               *string                             `mandatory:"false" json:"id"`
	CompartmentId    *string                             `mandatory:"false" json:"compartmentId"`
	DisplayName      *string                             `mandatory:"false" json:"displayName"`
	Description      *string                             `mandatory:"false" json:"description"`
	FreeformTags     map[string]string                   `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{}   `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{}   `mandatory:"false" json:"systemTags"`
	TimeCreated      *common.SDKTime                     `mandatory:"false" json:"timeCreated"`
	TimeUpdated      *common.SDKTime                     `mandatory:"false" json:"timeUpdated"`
	LifecycleState   OpsiConfigurationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails *string                             `mandatory:"false" json:"lifecycleDetails"`
	ConfigItems      json.RawMessage                     `mandatory:"false" json:"configItems"`
	OpsiConfigType   string                              `json:"opsiConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *opsiconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleropsiconfiguration opsiconfiguration
	s := struct {
		Model Unmarshaleropsiconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.ConfigItems = s.Model.ConfigItems
	m.OpsiConfigType = s.Model.OpsiConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *opsiconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OpsiConfigType {
	case "UX_CONFIGURATION":
		mm := OpsiUxConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OpsiConfiguration: %s.", m.OpsiConfigType)
		return *m, nil
	}
}

// GetId returns Id
func (m opsiconfiguration) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m opsiconfiguration) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m opsiconfiguration) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m opsiconfiguration) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m opsiconfiguration) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m opsiconfiguration) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m opsiconfiguration) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetTimeCreated returns TimeCreated
func (m opsiconfiguration) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m opsiconfiguration) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m opsiconfiguration) GetLifecycleState() OpsiConfigurationLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m opsiconfiguration) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetConfigItems returns ConfigItems
func (m opsiconfiguration) GetConfigItems() json.RawMessage {
	return m.ConfigItems
}

func (m opsiconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m opsiconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOpsiConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOpsiConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
