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

// UpdateComputeAutoActivatePluginConfigDetails Change the details of a COMPUTE_AUTO_ACTIVATE_PLUGIN config.
type UpdateComputeAutoActivatePluginConfigDetails struct {

	// The display name of the configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// True if automatic activation of the Management Agent plugin is enabled, false if it is not enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

// GetDisplayName returns DisplayName
func (m UpdateComputeAutoActivatePluginConfigDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m UpdateComputeAutoActivatePluginConfigDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateComputeAutoActivatePluginConfigDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateComputeAutoActivatePluginConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateComputeAutoActivatePluginConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateComputeAutoActivatePluginConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateComputeAutoActivatePluginConfigDetails UpdateComputeAutoActivatePluginConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeUpdateComputeAutoActivatePluginConfigDetails
	}{
		"COMPUTE_AUTO_ACTIVATE_PLUGIN",
		(MarshalTypeUpdateComputeAutoActivatePluginConfigDetails)(m),
	}

	return json.Marshal(&s)
}
