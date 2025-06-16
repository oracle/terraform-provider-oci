// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchOperationSummary Summary of a patch operation.
type PatchOperationSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// PatchOperation Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the PatchOperation was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the PatchOperation.
	LifecycleState PatchOperationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// If true, only validates that prerequisites are met
	IsPrerequisitesOnly *bool `mandatory:"false" json:"isPrerequisitesOnly"`

	// If true, only quick validations are run. This is applicable only when isPrerequisitesOnly=true.
	IsQuickPrerequisitesCheck *bool `mandatory:"false" json:"isQuickPrerequisitesCheck"`

	// The deployment type of the patch operation.
	DeploymentType DeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`

	// Working directory for staging binaries and temporary files
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// The unique identifier of the user who invoked this operation
	InvokedByUserId *string `mandatory:"false" json:"invokedByUserId"`

	// The time the PatchOperation was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time the PatchOperation was actually started. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the PatchOperation was actually completed. An RFC3339 formatted datetime string
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The elapsed time for the PatchOperation in seconds.
	TimeElapsedInSeconds *int64 `mandatory:"false" json:"timeElapsedInSeconds"`

	// The work state of the PatchOperation.
	WorkState PatchOperationWorkStateEnum `mandatory:"false" json:"workState,omitempty"`

	// Patch operation status.
	Status PatchOperationStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The total number of warnings found during the PatchOperation execution
	WarningsCount *int `mandatory:"false" json:"warningsCount"`

	// The number of resources patched by the PatchOperation
	ResourceCount *int `mandatory:"false" json:"resourceCount"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PatchOperationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchOperationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchOperationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPatchOperationLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchOperationWorkStateEnum(string(m.WorkState)); !ok && m.WorkState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkState: %s. Supported values are: %s.", m.WorkState, strings.Join(GetPatchOperationWorkStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchOperationStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPatchOperationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
