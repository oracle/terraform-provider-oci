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

// Workspace A workspace is an organizational construct to keep multiple data integration solutions and their resources (data assets, data flows, tasks, and so on) separate from each other, helping you to stay organized. For example, you could have separate workspaces for development, testing, and production.
type Workspace struct {

	// A user-friendly display name for the workspace. Does not have to be unique, and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A system-generated and immutable identifier assigned to the workspace upon creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the VCN the subnet is in.
	VcnId *string `mandatory:"false" json:"vcnId"`

	// The OCID of the subnet for customer connected databases.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The IP of the custom DNS.
	DnsServerIp *string `mandatory:"false" json:"dnsServerIp"`

	// The DNS zone of the custom DNS to use to resolve names.
	DnsServerZone *string `mandatory:"false" json:"dnsServerZone"`

	// Specifies whether the private network connection is enabled or disabled.
	IsPrivateNetworkEnabled *bool `mandatory:"false" json:"isPrivateNetworkEnabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A detailed description for the workspace.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the compartment containing the workspace.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The date and time the workspace was created, in the timestamp format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the workspace was updated, in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Lifecycle states for workspaces in Data Integration Service
	// CREATING - The resource is being created and may not be usable until the entire metadata is defined
	// UPDATING - The resource is being updated and may not be usable until all changes are commited
	// DELETING - The resource is being deleted and might require deep cleanup of children.
	// ACTIVE   - The resource is valid and available for access
	// INACTIVE - The resource might be incomplete in its definition or might have been made unavailable for
	//          administrative reasons
	// DELETED  - The resource has been deleted and isn't available
	// FAILED   - The resource is in a failed state due to validation or other errors
	// STARTING - The resource is being started and may not be usable until becomes ACTIVE again
	// STOPPING - The resource is in the process of Stopping and may not be usable until it Stops or fails
	// STOPPED  - The resource is in Stopped state due to stop operation.
	LifecycleState WorkspaceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// OCID of the private endpoint associated with the container/workspace.
	EndpointId *string `mandatory:"false" json:"endpointId"`

	// Name of the private endpoint associated with the container/workspace.
	EndpointName *string `mandatory:"false" json:"endpointName"`

	// DCMS Registry ID associated with the container/workspace.
	RegistryId *string `mandatory:"false" json:"registryId"`
}

func (m Workspace) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Workspace) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWorkspaceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetWorkspaceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkspaceLifecycleStateEnum Enum with underlying type: string
type WorkspaceLifecycleStateEnum string

// Set of constants representing the allowable values for WorkspaceLifecycleStateEnum
const (
	WorkspaceLifecycleStateCreating WorkspaceLifecycleStateEnum = "CREATING"
	WorkspaceLifecycleStateActive   WorkspaceLifecycleStateEnum = "ACTIVE"
	WorkspaceLifecycleStateInactive WorkspaceLifecycleStateEnum = "INACTIVE"
	WorkspaceLifecycleStateUpdating WorkspaceLifecycleStateEnum = "UPDATING"
	WorkspaceLifecycleStateDeleting WorkspaceLifecycleStateEnum = "DELETING"
	WorkspaceLifecycleStateDeleted  WorkspaceLifecycleStateEnum = "DELETED"
	WorkspaceLifecycleStateFailed   WorkspaceLifecycleStateEnum = "FAILED"
	WorkspaceLifecycleStateStarting WorkspaceLifecycleStateEnum = "STARTING"
	WorkspaceLifecycleStateStopping WorkspaceLifecycleStateEnum = "STOPPING"
	WorkspaceLifecycleStateStopped  WorkspaceLifecycleStateEnum = "STOPPED"
)

var mappingWorkspaceLifecycleStateEnum = map[string]WorkspaceLifecycleStateEnum{
	"CREATING": WorkspaceLifecycleStateCreating,
	"ACTIVE":   WorkspaceLifecycleStateActive,
	"INACTIVE": WorkspaceLifecycleStateInactive,
	"UPDATING": WorkspaceLifecycleStateUpdating,
	"DELETING": WorkspaceLifecycleStateDeleting,
	"DELETED":  WorkspaceLifecycleStateDeleted,
	"FAILED":   WorkspaceLifecycleStateFailed,
	"STARTING": WorkspaceLifecycleStateStarting,
	"STOPPING": WorkspaceLifecycleStateStopping,
	"STOPPED":  WorkspaceLifecycleStateStopped,
}

var mappingWorkspaceLifecycleStateEnumLowerCase = map[string]WorkspaceLifecycleStateEnum{
	"creating": WorkspaceLifecycleStateCreating,
	"active":   WorkspaceLifecycleStateActive,
	"inactive": WorkspaceLifecycleStateInactive,
	"updating": WorkspaceLifecycleStateUpdating,
	"deleting": WorkspaceLifecycleStateDeleting,
	"deleted":  WorkspaceLifecycleStateDeleted,
	"failed":   WorkspaceLifecycleStateFailed,
	"starting": WorkspaceLifecycleStateStarting,
	"stopping": WorkspaceLifecycleStateStopping,
	"stopped":  WorkspaceLifecycleStateStopped,
}

// GetWorkspaceLifecycleStateEnumValues Enumerates the set of values for WorkspaceLifecycleStateEnum
func GetWorkspaceLifecycleStateEnumValues() []WorkspaceLifecycleStateEnum {
	values := make([]WorkspaceLifecycleStateEnum, 0)
	for _, v := range mappingWorkspaceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkspaceLifecycleStateEnumStringValues Enumerates the set of values in String for WorkspaceLifecycleStateEnum
func GetWorkspaceLifecycleStateEnumStringValues() []string {
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

// GetMappingWorkspaceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkspaceLifecycleStateEnum(val string) (WorkspaceLifecycleStateEnum, bool) {
	enum, ok := mappingWorkspaceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
