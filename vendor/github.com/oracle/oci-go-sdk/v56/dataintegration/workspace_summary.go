// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// WorkspaceSummary Summary details of a workspace.
type WorkspaceSummary struct {

	// A system-generated and immutable identifier assigned to the workspace upon creation.
	Id *string `mandatory:"false" json:"id"`

	// A user defined description for the workspace.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly display name that is changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the compartment that contains the workspace.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The date and time the workspace was created, in the timestamp format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the workspace was updated, in the timestamp format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current state of the workspace.
	LifecycleState WorkspaceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A detailed description about the current state of the workspace. Used to provide actionable information if the workspace is in a failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`
}

func (m WorkspaceSummary) String() string {
	return common.PointerString(m)
}
