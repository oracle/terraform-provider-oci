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

// OnboardConfigDetails A configuration of the ONBOARD type, contains fields describing Onboarding customization: policies,
// dynamic groups, user groups.
type OnboardConfigDetails struct {

	// The Unique Oracle ID (OCID) that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// True if customer decides marks configuration as manually configured.
	IsManuallyOnboarded *bool `mandatory:"true" json:"isManuallyOnboarded"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the configuration was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Config was updated.
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

	// Assigned version to given onboard configuration.
	Version *string `mandatory:"false" json:"version"`

	// List of policy names assigned for onboarding
	PolicyNames []string `mandatory:"false" json:"policyNames"`

	// List of dynamic groups dedicated for Stack Monitoring.
	DynamicGroups []DynamicGroupDetails `mandatory:"false" json:"dynamicGroups"`

	// List of user groups dedicated for Stack Monitoring.
	UserGroups []GroupDetails `mandatory:"false" json:"userGroups"`

	AdditionalConfigurations *AdditionalConfigurationDetails `mandatory:"false" json:"additionalConfigurations"`

	// The current state of the configuration.
	LifecycleState ConfigLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m OnboardConfigDetails) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m OnboardConfigDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m OnboardConfigDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m OnboardConfigDetails) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m OnboardConfigDetails) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m OnboardConfigDetails) GetLifecycleState() ConfigLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m OnboardConfigDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m OnboardConfigDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m OnboardConfigDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m OnboardConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OnboardConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OnboardConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOnboardConfigDetails OnboardConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeOnboardConfigDetails
	}{
		"ONBOARD",
		(MarshalTypeOnboardConfigDetails)(m),
	}

	return json.Marshal(&s)
}
