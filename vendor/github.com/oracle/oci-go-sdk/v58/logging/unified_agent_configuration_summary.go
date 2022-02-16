// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UnifiedAgentConfigurationSummary Unified Agent configuration summary object returned by the list API.
type UnifiedAgentConfigurationSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly display name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The pipeline state.
	LifecycleState LogLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Type of Unified Agent service configuration.
	ConfigurationType UnifiedAgentServiceConfigurationTypesEnum `mandatory:"true" json:"configurationType"`

	// State of unified agent service configuration.
	ConfigurationState UnifiedAgentServiceConfigurationStatesEnum `mandatory:"true" json:"configurationState"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the resource was last modified.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`
}

func (m UnifiedAgentConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLogLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLogLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUnifiedAgentServiceConfigurationTypesEnum(string(m.ConfigurationType)); !ok && m.ConfigurationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigurationType: %s. Supported values are: %s.", m.ConfigurationType, strings.Join(GetUnifiedAgentServiceConfigurationTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUnifiedAgentServiceConfigurationStatesEnum(string(m.ConfigurationState)); !ok && m.ConfigurationState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigurationState: %s. Supported values are: %s.", m.ConfigurationState, strings.Join(GetUnifiedAgentServiceConfigurationStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
