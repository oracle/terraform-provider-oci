// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ConfigSummary Summary of the configuration.
type ConfigSummary interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// Compartment Identifier.
	GetCompartmentId() *string

	// The current state of the configuration.
	GetLifecycleState() ConfigLifecycleStateEnum

	// Config Identifier, can be renamed.
	GetDisplayName() *string

	// The time the the configuration was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the configuration was updated.
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

type configsummary struct {
	JsonData       []byte
	DisplayName    *string                           `mandatory:"false" json:"displayName"`
	TimeCreated    *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated    *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags     map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id             *string                           `mandatory:"true" json:"id"`
	CompartmentId  *string                           `mandatory:"true" json:"compartmentId"`
	LifecycleState ConfigLifecycleStateEnum          `mandatory:"true" json:"lifecycleState"`
	ConfigType     string                            `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *configsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigsummary configsummary
	s := struct {
		Model Unmarshalerconfigsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.LifecycleState = s.Model.LifecycleState
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "AUTO_PROMOTE":
		mm := AutoPromoteConfigSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LICENSE_AUTO_ASSIGN":
		mm := LicenseAutoAssignConfigSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LICENSE_ENTERPRISE_EXTENSIBILITY":
		mm := LicenseEnterpriseExtensibilityConfigSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ConfigSummary: %s.", m.ConfigType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m configsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m configsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m configsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m configsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m configsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m configsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m configsummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m configsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetLifecycleState returns LifecycleState
func (m configsummary) GetLifecycleState() ConfigLifecycleStateEnum {
	return m.LifecycleState
}

func (m configsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m configsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
