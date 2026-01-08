// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// PatchTaskSummary Summary of a patch task.
type PatchTaskSummary struct {

	// Unique identifier that is system generated
	Key *int64 `mandatory:"true" json:"key"`

	// PatchOperation Identifier which contains this patch task
	OperationId *string `mandatory:"true" json:"operationId"`

	// The type of the PatchTask
	Type PatchTaskTypeEnum `mandatory:"true" json:"type"`

	// The unique identifier of the resource being patched
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The work state of the PatchTask.
	WorkState PatchTaskWorkStateEnum `mandatory:"true" json:"workState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Name of the resource being patched
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// If true then patch task is skipped
	IsSkipped *bool `mandatory:"false" json:"isSkipped"`

	// Reason for the patch task operation skip
	ReasonSkipped *string `mandatory:"false" json:"reasonSkipped"`

	// The time the PatchTask was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the PatchTask actually started. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the PatchTask actually completed. An RFC3339 formatted datetime string
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The elapsed time for the PatchTask in seconds.
	TimeElapsedInSeconds *int64 `mandatory:"false" json:"timeElapsedInSeconds"`

	// Patch task status.
	Status PatchTaskStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The total number of retries done for the PatchTask
	RetryCount *int `mandatory:"false" json:"retryCount"`

	DeployDetails *DeployDetails `mandatory:"false" json:"deployDetails"`

	UpdateDetails *UpdateDetails `mandatory:"false" json:"updateDetails"`

	MigrateListenerDetails *MigrateListenerDetails `mandatory:"false" json:"migrateListenerDetails"`

	CleanupDetails *CleanupDetails `mandatory:"false" json:"cleanupDetails"`

	RollbackListenerDetails *RollbackListenerDetails `mandatory:"false" json:"rollbackListenerDetails"`

	RollbackDetails *RollbackDetails `mandatory:"false" json:"rollbackDetails"`

	Messages *PatchTaskMessages `mandatory:"false" json:"messages"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PatchTaskSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchTaskSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchTaskTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPatchTaskTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchTaskWorkStateEnum(string(m.WorkState)); !ok && m.WorkState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkState: %s. Supported values are: %s.", m.WorkState, strings.Join(GetPatchTaskWorkStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPatchTaskStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPatchTaskStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
