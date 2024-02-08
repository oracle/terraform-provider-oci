// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Stack The properties that define a stack.
// A stack is the collection of Oracle Cloud Infrastructure resources corresponding to a given Terraform configuration.
// For instructions on managing stacks, see
// Managing Stacks (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/stacks.htm).
// For more information about stacks, see
// Key Concepts (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#concepts__stackdefinition).
type Stack struct {

	// Unique identifier (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for the stack.
	Id *string `mandatory:"false" json:"id"`

	// Unique identifier (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for the compartment where the stack is located.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Human-readable name of the stack.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the stack.
	Description *string `mandatory:"false" json:"description"`

	// The date and time at which the stack was created.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current lifecycle state of the stack.
	// For more information about stack lifecycle states in Resource Manager, see
	// Key Concepts (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#concepts__StackStates).
	LifecycleState StackLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	ConfigSource ConfigSource `mandatory:"false" json:"configSource"`

	CustomTerraformProvider *CustomTerraformProvider `mandatory:"false" json:"customTerraformProvider"`

	// When `true`, the stack sources third-party Terraform providers from
	// Terraform Registry (https://registry.terraform.io/browse/providers) and allows
	// CustomTerraformProvider.
	// For more information about stack sourcing of third-party Terraform providers, see
	// Third-party Provider Configuration (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/terraformconfigresourcemanager.htm#third-party-providers).
	IsThirdPartyProviderExperienceEnabled *bool `mandatory:"false" json:"isThirdPartyProviderExperienceEnabled"`

	// Terraform variables associated with this resource.
	// Maximum number of variables supported is 250.
	// The maximum size of each variable, including both name and value, is 8192 bytes.
	// Example: `{"CompartmentId": "compartment-id-value"}`
	Variables map[string]string `mandatory:"false" json:"variables"`

	// The version of Terraform specified for the stack. Example: `0.12.x`
	TerraformVersion *string `mandatory:"false" json:"terraformVersion"`

	// Drift status of the stack.
	// Drift refers to differences between the actual (current) state of the stack and the expected (defined) state of the stack.
	StackDriftStatus StackStackDriftStatusEnum `mandatory:"false" json:"stackDriftStatus,omitempty"`

	// The date and time when the drift detection was last executed.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeDriftLastChecked *common.SDKTime `mandatory:"false" json:"timeDriftLastChecked"`

	// Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Stack) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Stack) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStackLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStackLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingStackStackDriftStatusEnum(string(m.StackDriftStatus)); !ok && m.StackDriftStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StackDriftStatus: %s. Supported values are: %s.", m.StackDriftStatus, strings.Join(GetStackStackDriftStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Stack) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id                                    *string                           `json:"id"`
		CompartmentId                         *string                           `json:"compartmentId"`
		DisplayName                           *string                           `json:"displayName"`
		Description                           *string                           `json:"description"`
		TimeCreated                           *common.SDKTime                   `json:"timeCreated"`
		LifecycleState                        StackLifecycleStateEnum           `json:"lifecycleState"`
		ConfigSource                          configsource                      `json:"configSource"`
		CustomTerraformProvider               *CustomTerraformProvider          `json:"customTerraformProvider"`
		IsThirdPartyProviderExperienceEnabled *bool                             `json:"isThirdPartyProviderExperienceEnabled"`
		Variables                             map[string]string                 `json:"variables"`
		TerraformVersion                      *string                           `json:"terraformVersion"`
		StackDriftStatus                      StackStackDriftStatusEnum         `json:"stackDriftStatus"`
		TimeDriftLastChecked                  *common.SDKTime                   `json:"timeDriftLastChecked"`
		FreeformTags                          map[string]string                 `json:"freeformTags"`
		DefinedTags                           map[string]map[string]interface{} `json:"definedTags"`
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

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	nn, e = model.ConfigSource.UnmarshalPolymorphicJSON(model.ConfigSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConfigSource = nn.(ConfigSource)
	} else {
		m.ConfigSource = nil
	}

	m.CustomTerraformProvider = model.CustomTerraformProvider

	m.IsThirdPartyProviderExperienceEnabled = model.IsThirdPartyProviderExperienceEnabled

	m.Variables = model.Variables

	m.TerraformVersion = model.TerraformVersion

	m.StackDriftStatus = model.StackDriftStatus

	m.TimeDriftLastChecked = model.TimeDriftLastChecked

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// StackLifecycleStateEnum Enum with underlying type: string
type StackLifecycleStateEnum string

// Set of constants representing the allowable values for StackLifecycleStateEnum
const (
	StackLifecycleStateCreating StackLifecycleStateEnum = "CREATING"
	StackLifecycleStateActive   StackLifecycleStateEnum = "ACTIVE"
	StackLifecycleStateDeleting StackLifecycleStateEnum = "DELETING"
	StackLifecycleStateDeleted  StackLifecycleStateEnum = "DELETED"
	StackLifecycleStateFailed   StackLifecycleStateEnum = "FAILED"
)

var mappingStackLifecycleStateEnum = map[string]StackLifecycleStateEnum{
	"CREATING": StackLifecycleStateCreating,
	"ACTIVE":   StackLifecycleStateActive,
	"DELETING": StackLifecycleStateDeleting,
	"DELETED":  StackLifecycleStateDeleted,
	"FAILED":   StackLifecycleStateFailed,
}

var mappingStackLifecycleStateEnumLowerCase = map[string]StackLifecycleStateEnum{
	"creating": StackLifecycleStateCreating,
	"active":   StackLifecycleStateActive,
	"deleting": StackLifecycleStateDeleting,
	"deleted":  StackLifecycleStateDeleted,
	"failed":   StackLifecycleStateFailed,
}

// GetStackLifecycleStateEnumValues Enumerates the set of values for StackLifecycleStateEnum
func GetStackLifecycleStateEnumValues() []StackLifecycleStateEnum {
	values := make([]StackLifecycleStateEnum, 0)
	for _, v := range mappingStackLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetStackLifecycleStateEnumStringValues Enumerates the set of values in String for StackLifecycleStateEnum
func GetStackLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingStackLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStackLifecycleStateEnum(val string) (StackLifecycleStateEnum, bool) {
	enum, ok := mappingStackLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// StackStackDriftStatusEnum Enum with underlying type: string
type StackStackDriftStatusEnum string

// Set of constants representing the allowable values for StackStackDriftStatusEnum
const (
	StackStackDriftStatusNotChecked StackStackDriftStatusEnum = "NOT_CHECKED"
	StackStackDriftStatusInSync     StackStackDriftStatusEnum = "IN_SYNC"
	StackStackDriftStatusDrifted    StackStackDriftStatusEnum = "DRIFTED"
)

var mappingStackStackDriftStatusEnum = map[string]StackStackDriftStatusEnum{
	"NOT_CHECKED": StackStackDriftStatusNotChecked,
	"IN_SYNC":     StackStackDriftStatusInSync,
	"DRIFTED":     StackStackDriftStatusDrifted,
}

var mappingStackStackDriftStatusEnumLowerCase = map[string]StackStackDriftStatusEnum{
	"not_checked": StackStackDriftStatusNotChecked,
	"in_sync":     StackStackDriftStatusInSync,
	"drifted":     StackStackDriftStatusDrifted,
}

// GetStackStackDriftStatusEnumValues Enumerates the set of values for StackStackDriftStatusEnum
func GetStackStackDriftStatusEnumValues() []StackStackDriftStatusEnum {
	values := make([]StackStackDriftStatusEnum, 0)
	for _, v := range mappingStackStackDriftStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetStackStackDriftStatusEnumStringValues Enumerates the set of values in String for StackStackDriftStatusEnum
func GetStackStackDriftStatusEnumStringValues() []string {
	return []string{
		"NOT_CHECKED",
		"IN_SYNC",
		"DRIFTED",
	}
}

// GetMappingStackStackDriftStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStackStackDriftStatusEnum(val string) (StackStackDriftStatusEnum, bool) {
	enum, ok := mappingStackStackDriftStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
