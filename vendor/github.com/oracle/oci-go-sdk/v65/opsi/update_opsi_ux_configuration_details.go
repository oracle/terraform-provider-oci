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

// UpdateOpsiUxConfigurationDetails Information to be updated in OPSI UX configuration.
type UpdateOpsiUxConfigurationDetails struct {

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
	// This array overwrites the existing custom configuration items array for this resource.
	ConfigItems []UpdateConfigurationItemDetails `mandatory:"false" json:"configItems"`
}

// GetDisplayName returns DisplayName
func (m UpdateOpsiUxConfigurationDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m UpdateOpsiUxConfigurationDetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m UpdateOpsiUxConfigurationDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateOpsiUxConfigurationDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m UpdateOpsiUxConfigurationDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetConfigItems returns ConfigItems
func (m UpdateOpsiUxConfigurationDetails) GetConfigItems() []UpdateConfigurationItemDetails {
	return m.ConfigItems
}

func (m UpdateOpsiUxConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOpsiUxConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOpsiUxConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOpsiUxConfigurationDetails UpdateOpsiUxConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"opsiConfigType"`
		MarshalTypeUpdateOpsiUxConfigurationDetails
	}{
		"UX_CONFIGURATION",
		(MarshalTypeUpdateOpsiUxConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateOpsiUxConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName  *string                           `json:"displayName"`
		Description  *string                           `json:"description"`
		FreeformTags map[string]string                 `json:"freeformTags"`
		DefinedTags  map[string]map[string]interface{} `json:"definedTags"`
		SystemTags   map[string]map[string]interface{} `json:"systemTags"`
		ConfigItems  []updateconfigurationitemdetails  `json:"configItems"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.ConfigItems = make([]UpdateConfigurationItemDetails, len(model.ConfigItems))
	for i, n := range model.ConfigItems {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ConfigItems[i] = nn.(UpdateConfigurationItemDetails)
		} else {
			m.ConfigItems[i] = nil
		}
	}
	return
}
