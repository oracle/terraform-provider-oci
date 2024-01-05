// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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
	LifecycleState WorkspaceSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A detailed description about the current state of the workspace. Used to provide actionable information if the workspace is in a failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// Name of the private endpoint associated with the container/workspace. Returns null if there is none.
	EndpointName *string `mandatory:"false" json:"endpointName"`

	// DCMS endpoint associated with the container/workspace. Returns null if there is none.
	EndpointId *string `mandatory:"false" json:"endpointId"`

	// DCMS registry associated with the container/workspace. Returns null if there is none.
	RegistryId *string `mandatory:"false" json:"registryId"`
}

func (m WorkspaceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkspaceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWorkspaceSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetWorkspaceSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkspaceSummaryLifecycleStateEnum Enum with underlying type: string
type WorkspaceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for WorkspaceSummaryLifecycleStateEnum
const (
	WorkspaceSummaryLifecycleStateCreating WorkspaceSummaryLifecycleStateEnum = "CREATING"
	WorkspaceSummaryLifecycleStateActive   WorkspaceSummaryLifecycleStateEnum = "ACTIVE"
	WorkspaceSummaryLifecycleStateInactive WorkspaceSummaryLifecycleStateEnum = "INACTIVE"
	WorkspaceSummaryLifecycleStateUpdating WorkspaceSummaryLifecycleStateEnum = "UPDATING"
	WorkspaceSummaryLifecycleStateDeleting WorkspaceSummaryLifecycleStateEnum = "DELETING"
	WorkspaceSummaryLifecycleStateDeleted  WorkspaceSummaryLifecycleStateEnum = "DELETED"
	WorkspaceSummaryLifecycleStateFailed   WorkspaceSummaryLifecycleStateEnum = "FAILED"
	WorkspaceSummaryLifecycleStateStarting WorkspaceSummaryLifecycleStateEnum = "STARTING"
	WorkspaceSummaryLifecycleStateStopping WorkspaceSummaryLifecycleStateEnum = "STOPPING"
	WorkspaceSummaryLifecycleStateStopped  WorkspaceSummaryLifecycleStateEnum = "STOPPED"
)

var mappingWorkspaceSummaryLifecycleStateEnum = map[string]WorkspaceSummaryLifecycleStateEnum{
	"CREATING": WorkspaceSummaryLifecycleStateCreating,
	"ACTIVE":   WorkspaceSummaryLifecycleStateActive,
	"INACTIVE": WorkspaceSummaryLifecycleStateInactive,
	"UPDATING": WorkspaceSummaryLifecycleStateUpdating,
	"DELETING": WorkspaceSummaryLifecycleStateDeleting,
	"DELETED":  WorkspaceSummaryLifecycleStateDeleted,
	"FAILED":   WorkspaceSummaryLifecycleStateFailed,
	"STARTING": WorkspaceSummaryLifecycleStateStarting,
	"STOPPING": WorkspaceSummaryLifecycleStateStopping,
	"STOPPED":  WorkspaceSummaryLifecycleStateStopped,
}

var mappingWorkspaceSummaryLifecycleStateEnumLowerCase = map[string]WorkspaceSummaryLifecycleStateEnum{
	"creating": WorkspaceSummaryLifecycleStateCreating,
	"active":   WorkspaceSummaryLifecycleStateActive,
	"inactive": WorkspaceSummaryLifecycleStateInactive,
	"updating": WorkspaceSummaryLifecycleStateUpdating,
	"deleting": WorkspaceSummaryLifecycleStateDeleting,
	"deleted":  WorkspaceSummaryLifecycleStateDeleted,
	"failed":   WorkspaceSummaryLifecycleStateFailed,
	"starting": WorkspaceSummaryLifecycleStateStarting,
	"stopping": WorkspaceSummaryLifecycleStateStopping,
	"stopped":  WorkspaceSummaryLifecycleStateStopped,
}

// GetWorkspaceSummaryLifecycleStateEnumValues Enumerates the set of values for WorkspaceSummaryLifecycleStateEnum
func GetWorkspaceSummaryLifecycleStateEnumValues() []WorkspaceSummaryLifecycleStateEnum {
	values := make([]WorkspaceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingWorkspaceSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkspaceSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for WorkspaceSummaryLifecycleStateEnum
func GetWorkspaceSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"STARTING",
		"STOPPING",
		"STOPPED",
	}
}

// GetMappingWorkspaceSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkspaceSummaryLifecycleStateEnum(val string) (WorkspaceSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingWorkspaceSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
