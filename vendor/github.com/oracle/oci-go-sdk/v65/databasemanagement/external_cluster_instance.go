// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalClusterInstance The details of an external cluster instance.
type ExternalClusterInstance struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external cluster instance.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the cluster instance. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the external cluster instance.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external cluster that the cluster instance belongs to.
	ExternalClusterId *string `mandatory:"true" json:"externalClusterId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system that the cluster instance is a part of.
	ExternalDbSystemId *string `mandatory:"true" json:"externalDbSystemId"`

	// The current lifecycle state of the external cluster instance.
	LifecycleState ExternalClusterInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB node.
	ExternalDbNodeId *string `mandatory:"false" json:"externalDbNodeId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external connector.
	ExternalConnectorId *string `mandatory:"false" json:"externalConnectorId"`

	// The name of the host on which the cluster instance is running.
	HostName *string `mandatory:"false" json:"hostName"`

	// The role of the cluster node.
	NodeRole ExternalClusterInstanceNodeRoleEnum `mandatory:"false" json:"nodeRole,omitempty"`

	// The Oracle base location of Cluster Ready Services (CRS).
	CrsBaseDirectory *string `mandatory:"false" json:"crsBaseDirectory"`

	// The Automatic Diagnostic Repository (ADR) home directory for the cluster instance.
	AdrHomeDirectory *string `mandatory:"false" json:"adrHomeDirectory"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the external cluster instance was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the external cluster instance was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ExternalClusterInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalClusterInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalClusterInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalClusterInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExternalClusterInstanceNodeRoleEnum(string(m.NodeRole)); !ok && m.NodeRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeRole: %s. Supported values are: %s.", m.NodeRole, strings.Join(GetExternalClusterInstanceNodeRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalClusterInstanceNodeRoleEnum Enum with underlying type: string
type ExternalClusterInstanceNodeRoleEnum string

// Set of constants representing the allowable values for ExternalClusterInstanceNodeRoleEnum
const (
	ExternalClusterInstanceNodeRoleHub  ExternalClusterInstanceNodeRoleEnum = "HUB"
	ExternalClusterInstanceNodeRoleLeaf ExternalClusterInstanceNodeRoleEnum = "LEAF"
)

var mappingExternalClusterInstanceNodeRoleEnum = map[string]ExternalClusterInstanceNodeRoleEnum{
	"HUB":  ExternalClusterInstanceNodeRoleHub,
	"LEAF": ExternalClusterInstanceNodeRoleLeaf,
}

var mappingExternalClusterInstanceNodeRoleEnumLowerCase = map[string]ExternalClusterInstanceNodeRoleEnum{
	"hub":  ExternalClusterInstanceNodeRoleHub,
	"leaf": ExternalClusterInstanceNodeRoleLeaf,
}

// GetExternalClusterInstanceNodeRoleEnumValues Enumerates the set of values for ExternalClusterInstanceNodeRoleEnum
func GetExternalClusterInstanceNodeRoleEnumValues() []ExternalClusterInstanceNodeRoleEnum {
	values := make([]ExternalClusterInstanceNodeRoleEnum, 0)
	for _, v := range mappingExternalClusterInstanceNodeRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalClusterInstanceNodeRoleEnumStringValues Enumerates the set of values in String for ExternalClusterInstanceNodeRoleEnum
func GetExternalClusterInstanceNodeRoleEnumStringValues() []string {
	return []string{
		"HUB",
		"LEAF",
	}
}

// GetMappingExternalClusterInstanceNodeRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalClusterInstanceNodeRoleEnum(val string) (ExternalClusterInstanceNodeRoleEnum, bool) {
	enum, ok := mappingExternalClusterInstanceNodeRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalClusterInstanceLifecycleStateEnum Enum with underlying type: string
type ExternalClusterInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalClusterInstanceLifecycleStateEnum
const (
	ExternalClusterInstanceLifecycleStateCreating     ExternalClusterInstanceLifecycleStateEnum = "CREATING"
	ExternalClusterInstanceLifecycleStateNotConnected ExternalClusterInstanceLifecycleStateEnum = "NOT_CONNECTED"
	ExternalClusterInstanceLifecycleStateActive       ExternalClusterInstanceLifecycleStateEnum = "ACTIVE"
	ExternalClusterInstanceLifecycleStateInactive     ExternalClusterInstanceLifecycleStateEnum = "INACTIVE"
	ExternalClusterInstanceLifecycleStateUpdating     ExternalClusterInstanceLifecycleStateEnum = "UPDATING"
	ExternalClusterInstanceLifecycleStateDeleting     ExternalClusterInstanceLifecycleStateEnum = "DELETING"
	ExternalClusterInstanceLifecycleStateDeleted      ExternalClusterInstanceLifecycleStateEnum = "DELETED"
	ExternalClusterInstanceLifecycleStateFailed       ExternalClusterInstanceLifecycleStateEnum = "FAILED"
)

var mappingExternalClusterInstanceLifecycleStateEnum = map[string]ExternalClusterInstanceLifecycleStateEnum{
	"CREATING":      ExternalClusterInstanceLifecycleStateCreating,
	"NOT_CONNECTED": ExternalClusterInstanceLifecycleStateNotConnected,
	"ACTIVE":        ExternalClusterInstanceLifecycleStateActive,
	"INACTIVE":      ExternalClusterInstanceLifecycleStateInactive,
	"UPDATING":      ExternalClusterInstanceLifecycleStateUpdating,
	"DELETING":      ExternalClusterInstanceLifecycleStateDeleting,
	"DELETED":       ExternalClusterInstanceLifecycleStateDeleted,
	"FAILED":        ExternalClusterInstanceLifecycleStateFailed,
}

var mappingExternalClusterInstanceLifecycleStateEnumLowerCase = map[string]ExternalClusterInstanceLifecycleStateEnum{
	"creating":      ExternalClusterInstanceLifecycleStateCreating,
	"not_connected": ExternalClusterInstanceLifecycleStateNotConnected,
	"active":        ExternalClusterInstanceLifecycleStateActive,
	"inactive":      ExternalClusterInstanceLifecycleStateInactive,
	"updating":      ExternalClusterInstanceLifecycleStateUpdating,
	"deleting":      ExternalClusterInstanceLifecycleStateDeleting,
	"deleted":       ExternalClusterInstanceLifecycleStateDeleted,
	"failed":        ExternalClusterInstanceLifecycleStateFailed,
}

// GetExternalClusterInstanceLifecycleStateEnumValues Enumerates the set of values for ExternalClusterInstanceLifecycleStateEnum
func GetExternalClusterInstanceLifecycleStateEnumValues() []ExternalClusterInstanceLifecycleStateEnum {
	values := make([]ExternalClusterInstanceLifecycleStateEnum, 0)
	for _, v := range mappingExternalClusterInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalClusterInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalClusterInstanceLifecycleStateEnum
func GetExternalClusterInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"NOT_CONNECTED",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingExternalClusterInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalClusterInstanceLifecycleStateEnum(val string) (ExternalClusterInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingExternalClusterInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
