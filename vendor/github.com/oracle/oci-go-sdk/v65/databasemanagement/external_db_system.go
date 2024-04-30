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

// ExternalDbSystem The details of an external DB system.
type ExternalDbSystem struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the DB system. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state of the external DB system resource.
	LifecycleState ExternalDbSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the external DB system was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the external DB system was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system discovery.
	DbSystemDiscoveryId *string `mandatory:"false" json:"dbSystemDiscoveryId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the management agent used during the discovery of the DB system.
	DiscoveryAgentId *string `mandatory:"false" json:"discoveryAgentId"`

	// Indicates whether the DB system is a cluster DB system or not.
	IsCluster *bool `mandatory:"false" json:"isCluster"`

	// The Oracle Grid home directory in case of cluster-based DB system and
	// Oracle home directory in case of single instance-based DB system.
	HomeDirectory *string `mandatory:"false" json:"homeDirectory"`

	DatabaseManagementConfig *ExternalDbSystemDatabaseManagementConfigDetails `mandatory:"false" json:"databaseManagementConfig"`

	StackMonitoringConfig *ExternalDbSystemStackMonitoringConfigDetails `mandatory:"false" json:"stackMonitoringConfig"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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

func (m ExternalDbSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDbSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDbSystemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalDbSystemLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalDbSystemLifecycleStateEnum Enum with underlying type: string
type ExternalDbSystemLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalDbSystemLifecycleStateEnum
const (
	ExternalDbSystemLifecycleStateCreating ExternalDbSystemLifecycleStateEnum = "CREATING"
	ExternalDbSystemLifecycleStateActive   ExternalDbSystemLifecycleStateEnum = "ACTIVE"
	ExternalDbSystemLifecycleStateUpdating ExternalDbSystemLifecycleStateEnum = "UPDATING"
	ExternalDbSystemLifecycleStateDeleting ExternalDbSystemLifecycleStateEnum = "DELETING"
	ExternalDbSystemLifecycleStateDeleted  ExternalDbSystemLifecycleStateEnum = "DELETED"
	ExternalDbSystemLifecycleStateInactive ExternalDbSystemLifecycleStateEnum = "INACTIVE"
)

var mappingExternalDbSystemLifecycleStateEnum = map[string]ExternalDbSystemLifecycleStateEnum{
	"CREATING": ExternalDbSystemLifecycleStateCreating,
	"ACTIVE":   ExternalDbSystemLifecycleStateActive,
	"UPDATING": ExternalDbSystemLifecycleStateUpdating,
	"DELETING": ExternalDbSystemLifecycleStateDeleting,
	"DELETED":  ExternalDbSystemLifecycleStateDeleted,
	"INACTIVE": ExternalDbSystemLifecycleStateInactive,
}

var mappingExternalDbSystemLifecycleStateEnumLowerCase = map[string]ExternalDbSystemLifecycleStateEnum{
	"creating": ExternalDbSystemLifecycleStateCreating,
	"active":   ExternalDbSystemLifecycleStateActive,
	"updating": ExternalDbSystemLifecycleStateUpdating,
	"deleting": ExternalDbSystemLifecycleStateDeleting,
	"deleted":  ExternalDbSystemLifecycleStateDeleted,
	"inactive": ExternalDbSystemLifecycleStateInactive,
}

// GetExternalDbSystemLifecycleStateEnumValues Enumerates the set of values for ExternalDbSystemLifecycleStateEnum
func GetExternalDbSystemLifecycleStateEnumValues() []ExternalDbSystemLifecycleStateEnum {
	values := make([]ExternalDbSystemLifecycleStateEnum, 0)
	for _, v := range mappingExternalDbSystemLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDbSystemLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalDbSystemLifecycleStateEnum
func GetExternalDbSystemLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"INACTIVE",
	}
}

// GetMappingExternalDbSystemLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDbSystemLifecycleStateEnum(val string) (ExternalDbSystemLifecycleStateEnum, bool) {
	enum, ok := mappingExternalDbSystemLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
