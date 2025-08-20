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

// CreateComputeAutoActivatePluginConfigDetails The details of a COMPUTE_AUTO_ACTIVATE_PLUGIN configuration.
type CreateComputeAutoActivatePluginConfigDetails struct {

	// Compartment in which the configuration is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// True if automatic activation of the Management Agent plugin is enabled, false if it is not enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The display name of the configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetDisplayName returns DisplayName
func (m CreateComputeAutoActivatePluginConfigDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateComputeAutoActivatePluginConfigDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateComputeAutoActivatePluginConfigDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateComputeAutoActivatePluginConfigDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateComputeAutoActivatePluginConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateComputeAutoActivatePluginConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateComputeAutoActivatePluginConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateComputeAutoActivatePluginConfigDetails CreateComputeAutoActivatePluginConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateComputeAutoActivatePluginConfigDetails
	}{
		"COMPUTE_AUTO_ACTIVATE_PLUGIN",
		(MarshalTypeCreateComputeAutoActivatePluginConfigDetails)(m),
	}

	return json.Marshal(&s)
}
