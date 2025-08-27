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

// CreateOnboardConfigDetails A configuration of the ONBOARD type, contains fields describing Onboarding customization: policies,
// dynamic groups, user groups.
type CreateOnboardConfigDetails struct {

	// Compartment in which the configuration is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// True if customer decides marks configuration as manually configured.
	IsManuallyOnboarded *bool `mandatory:"true" json:"isManuallyOnboarded"`

	// The display name of the configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Assigned version to given onboard configuration.
	Version *string `mandatory:"false" json:"version"`

	// List of policy names assigned for onboarding
	PolicyNames []string `mandatory:"false" json:"policyNames"`

	// List of dynamic groups dedicated for Stack Monitoring.
	DynamicGroups []DynamicGroupDetails `mandatory:"false" json:"dynamicGroups"`

	// List of user groups dedicated for Stack Monitoring.
	UserGroups []GroupDetails `mandatory:"false" json:"userGroups"`

	AdditionalConfigurations *AdditionalConfigurationDetails `mandatory:"false" json:"additionalConfigurations"`
}

// GetDisplayName returns DisplayName
func (m CreateOnboardConfigDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateOnboardConfigDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateOnboardConfigDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateOnboardConfigDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateOnboardConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOnboardConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOnboardConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOnboardConfigDetails CreateOnboardConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateOnboardConfigDetails
	}{
		"ONBOARD",
		(MarshalTypeCreateOnboardConfigDetails)(m),
	}

	return json.Marshal(&s)
}
