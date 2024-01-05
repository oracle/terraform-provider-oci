// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOpsiUxConfigurationDetails Information about OPSI UX configuration to be created.
type CreateOpsiUxConfigurationDetails struct {

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

	// Array of configuration items with custom values. All and only configuration items requiring custom values should be part of this array.
	ConfigItems []CreateConfigurationItemDetails `mandatory:"false" json:"configItems"`
}

// GetCompartmentId returns CompartmentId
func (m CreateOpsiUxConfigurationDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m CreateOpsiUxConfigurationDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateOpsiUxConfigurationDetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m CreateOpsiUxConfigurationDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateOpsiUxConfigurationDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m CreateOpsiUxConfigurationDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetConfigItems returns ConfigItems
func (m CreateOpsiUxConfigurationDetails) GetConfigItems() []CreateConfigurationItemDetails {
	return m.ConfigItems
}

func (m CreateOpsiUxConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOpsiUxConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOpsiUxConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOpsiUxConfigurationDetails CreateOpsiUxConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"opsiConfigType"`
		MarshalTypeCreateOpsiUxConfigurationDetails
	}{
		"UX_CONFIGURATION",
		(MarshalTypeCreateOpsiUxConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateOpsiUxConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId *string                           `json:"compartmentId"`
		DisplayName   *string                           `json:"displayName"`
		Description   *string                           `json:"description"`
		FreeformTags  map[string]string                 `json:"freeformTags"`
		DefinedTags   map[string]map[string]interface{} `json:"definedTags"`
		SystemTags    map[string]map[string]interface{} `json:"systemTags"`
		ConfigItems   []createconfigurationitemdetails  `json:"configItems"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.ConfigItems = make([]CreateConfigurationItemDetails, len(model.ConfigItems))
	for i, n := range model.ConfigItems {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ConfigItems[i] = nn.(CreateConfigurationItemDetails)
		} else {
			m.ConfigItems[i] = nil
		}
	}
	return
}
