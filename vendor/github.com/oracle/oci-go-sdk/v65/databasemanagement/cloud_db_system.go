// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudDbSystem The details of a cloud DB system.
type CloudDbSystem struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent cloud DB Infrastructure. For VM Dbsystems ,
	// it will be the DBSystem Id. For ExaCS and ExaCC,  it will be the cloudVmClusterId and vmClusterId respectively.
	DbaasParentInfrastructureId *string `mandatory:"true" json:"dbaasParentInfrastructureId"`

	// The deployment type of cloud dbsystem.
	DeploymentType CloudDbSystemDeploymentTypeEnum `mandatory:"true" json:"deploymentType"`

	// The user-friendly name for the DB system. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state of the cloud DB system resource.
	LifecycleState CloudDbSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the cloud DB system was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the cloud DB system was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system discovery.
	DbSystemDiscoveryId *string `mandatory:"false" json:"dbSystemDiscoveryId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used during the discovery of the DB system.
	DiscoveryAgentId *string `mandatory:"false" json:"discoveryAgentId"`

	// Indicates whether the DB system is a cluster DB system or not.
	IsCluster *bool `mandatory:"false" json:"isCluster"`

	// The Oracle Grid home directory in case of cluster-based DB system and
	// Oracle home directory in case of single instance-based DB system.
	HomeDirectory *string `mandatory:"false" json:"homeDirectory"`

	DatabaseManagementConfig *CloudDbSystemDatabaseManagementConfigDetails `mandatory:"false" json:"databaseManagementConfig"`

	StackMonitoringConfig *CloudDbSystemStackMonitoringConfigDetails `mandatory:"false" json:"stackMonitoringConfig"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CloudDbSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudDbSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudDbSystemDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetCloudDbSystemDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCloudDbSystemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudDbSystemLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudDbSystemLifecycleStateEnum Enum with underlying type: string
type CloudDbSystemLifecycleStateEnum string

// Set of constants representing the allowable values for CloudDbSystemLifecycleStateEnum
const (
	CloudDbSystemLifecycleStateCreating CloudDbSystemLifecycleStateEnum = "CREATING"
	CloudDbSystemLifecycleStateActive   CloudDbSystemLifecycleStateEnum = "ACTIVE"
	CloudDbSystemLifecycleStateUpdating CloudDbSystemLifecycleStateEnum = "UPDATING"
	CloudDbSystemLifecycleStateDeleting CloudDbSystemLifecycleStateEnum = "DELETING"
	CloudDbSystemLifecycleStateDeleted  CloudDbSystemLifecycleStateEnum = "DELETED"
	CloudDbSystemLifecycleStateInactive CloudDbSystemLifecycleStateEnum = "INACTIVE"
)

var mappingCloudDbSystemLifecycleStateEnum = map[string]CloudDbSystemLifecycleStateEnum{
	"CREATING": CloudDbSystemLifecycleStateCreating,
	"ACTIVE":   CloudDbSystemLifecycleStateActive,
	"UPDATING": CloudDbSystemLifecycleStateUpdating,
	"DELETING": CloudDbSystemLifecycleStateDeleting,
	"DELETED":  CloudDbSystemLifecycleStateDeleted,
	"INACTIVE": CloudDbSystemLifecycleStateInactive,
}

var mappingCloudDbSystemLifecycleStateEnumLowerCase = map[string]CloudDbSystemLifecycleStateEnum{
	"creating": CloudDbSystemLifecycleStateCreating,
	"active":   CloudDbSystemLifecycleStateActive,
	"updating": CloudDbSystemLifecycleStateUpdating,
	"deleting": CloudDbSystemLifecycleStateDeleting,
	"deleted":  CloudDbSystemLifecycleStateDeleted,
	"inactive": CloudDbSystemLifecycleStateInactive,
}

// GetCloudDbSystemLifecycleStateEnumValues Enumerates the set of values for CloudDbSystemLifecycleStateEnum
func GetCloudDbSystemLifecycleStateEnumValues() []CloudDbSystemLifecycleStateEnum {
	values := make([]CloudDbSystemLifecycleStateEnum, 0)
	for _, v := range mappingCloudDbSystemLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudDbSystemLifecycleStateEnumStringValues Enumerates the set of values in String for CloudDbSystemLifecycleStateEnum
func GetCloudDbSystemLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"INACTIVE",
	}
}

// GetMappingCloudDbSystemLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudDbSystemLifecycleStateEnum(val string) (CloudDbSystemLifecycleStateEnum, bool) {
	enum, ok := mappingCloudDbSystemLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
