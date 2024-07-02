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

// OpsiUxConfiguration OPSI UX configuration.
type OpsiUxConfiguration struct {

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of OPSI configuration resource.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// User-friendly display name for the OPSI configuration. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of OPSI configuration.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The time at which the resource was first created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time at which the resource was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Array of configuration item summary objects.
	ConfigItems []OpsiConfigurationConfigurationItemSummary `mandatory:"false" json:"configItems"`

	// OPSI configuration resource lifecycle state.
	LifecycleState OpsiConfigurationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m OpsiUxConfiguration) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m OpsiUxConfiguration) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m OpsiUxConfiguration) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m OpsiUxConfiguration) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m OpsiUxConfiguration) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m OpsiUxConfiguration) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m OpsiUxConfiguration) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetTimeCreated returns TimeCreated
func (m OpsiUxConfiguration) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m OpsiUxConfiguration) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m OpsiUxConfiguration) GetLifecycleState() OpsiConfigurationLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m OpsiUxConfiguration) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetConfigItems returns ConfigItems
func (m OpsiUxConfiguration) GetConfigItems() []OpsiConfigurationConfigurationItemSummary {
	return m.ConfigItems
}

func (m OpsiUxConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpsiUxConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOpsiConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOpsiConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OpsiUxConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOpsiUxConfiguration OpsiUxConfiguration
	s := struct {
		DiscriminatorParam string `json:"opsiConfigType"`
		MarshalTypeOpsiUxConfiguration
	}{
		"UX_CONFIGURATION",
		(MarshalTypeOpsiUxConfiguration)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OpsiUxConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id               *string                                     `json:"id"`
		CompartmentId    *string                                     `json:"compartmentId"`
		DisplayName      *string                                     `json:"displayName"`
		Description      *string                                     `json:"description"`
		FreeformTags     map[string]string                           `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{}           `json:"definedTags"`
		SystemTags       map[string]map[string]interface{}           `json:"systemTags"`
		TimeCreated      *common.SDKTime                             `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                             `json:"timeUpdated"`
		LifecycleState   OpsiConfigurationLifecycleStateEnum         `json:"lifecycleState"`
		LifecycleDetails *string                                     `json:"lifecycleDetails"`
		ConfigItems      []opsiconfigurationconfigurationitemsummary `json:"configItems"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.ConfigItems = make([]OpsiConfigurationConfigurationItemSummary, len(model.ConfigItems))
	for i, n := range model.ConfigItems {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ConfigItems[i] = nn.(OpsiConfigurationConfigurationItemSummary)
		} else {
			m.ConfigItems[i] = nil
		}
	}
	return
}
