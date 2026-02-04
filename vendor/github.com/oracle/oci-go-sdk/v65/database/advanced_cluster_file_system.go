// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdvancedClusterFileSystem Details of an advanced cluster file system.
type AdvancedClusterFileSystem struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the advanced cluster file system.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the Advanced cluster file system. The file system name is unique for a cluster.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	// The current state of the advanced cluster file system. Valid states are CREATING, AVAILABLE, UPDATING, FAILED, DELETED.
	LifecycleState AdvancedClusterFileSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The total storage allocated for advanced cluster file system in GBs.
	StorageInGBs *int `mandatory:"true" json:"storageInGBs"`

	// The mount point of file system.
	MountPoint *string `mandatory:"true" json:"mountPoint"`

	// True if the file system is mounted on all VMs within VM Cluster.
	IsMounted *bool `mandatory:"true" json:"isMounted"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Description of the advanced cluster file system.
	Description *string `mandatory:"false" json:"description"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the advanced cluster file system was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last date and time that the advanced cluster file system was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AdvancedClusterFileSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdvancedClusterFileSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAdvancedClusterFileSystemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAdvancedClusterFileSystemLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AdvancedClusterFileSystemLifecycleStateEnum Enum with underlying type: string
type AdvancedClusterFileSystemLifecycleStateEnum string

// Set of constants representing the allowable values for AdvancedClusterFileSystemLifecycleStateEnum
const (
	AdvancedClusterFileSystemLifecycleStateAvailable AdvancedClusterFileSystemLifecycleStateEnum = "AVAILABLE"
	AdvancedClusterFileSystemLifecycleStateCreating  AdvancedClusterFileSystemLifecycleStateEnum = "CREATING"
	AdvancedClusterFileSystemLifecycleStateFailed    AdvancedClusterFileSystemLifecycleStateEnum = "FAILED"
	AdvancedClusterFileSystemLifecycleStateUpdating  AdvancedClusterFileSystemLifecycleStateEnum = "UPDATING"
	AdvancedClusterFileSystemLifecycleStateDeleted   AdvancedClusterFileSystemLifecycleStateEnum = "DELETED"
	AdvancedClusterFileSystemLifecycleStateDeleting  AdvancedClusterFileSystemLifecycleStateEnum = "DELETING"
)

var mappingAdvancedClusterFileSystemLifecycleStateEnum = map[string]AdvancedClusterFileSystemLifecycleStateEnum{
	"AVAILABLE": AdvancedClusterFileSystemLifecycleStateAvailable,
	"CREATING":  AdvancedClusterFileSystemLifecycleStateCreating,
	"FAILED":    AdvancedClusterFileSystemLifecycleStateFailed,
	"UPDATING":  AdvancedClusterFileSystemLifecycleStateUpdating,
	"DELETED":   AdvancedClusterFileSystemLifecycleStateDeleted,
	"DELETING":  AdvancedClusterFileSystemLifecycleStateDeleting,
}

var mappingAdvancedClusterFileSystemLifecycleStateEnumLowerCase = map[string]AdvancedClusterFileSystemLifecycleStateEnum{
	"available": AdvancedClusterFileSystemLifecycleStateAvailable,
	"creating":  AdvancedClusterFileSystemLifecycleStateCreating,
	"failed":    AdvancedClusterFileSystemLifecycleStateFailed,
	"updating":  AdvancedClusterFileSystemLifecycleStateUpdating,
	"deleted":   AdvancedClusterFileSystemLifecycleStateDeleted,
	"deleting":  AdvancedClusterFileSystemLifecycleStateDeleting,
}

// GetAdvancedClusterFileSystemLifecycleStateEnumValues Enumerates the set of values for AdvancedClusterFileSystemLifecycleStateEnum
func GetAdvancedClusterFileSystemLifecycleStateEnumValues() []AdvancedClusterFileSystemLifecycleStateEnum {
	values := make([]AdvancedClusterFileSystemLifecycleStateEnum, 0)
	for _, v := range mappingAdvancedClusterFileSystemLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAdvancedClusterFileSystemLifecycleStateEnumStringValues Enumerates the set of values in String for AdvancedClusterFileSystemLifecycleStateEnum
func GetAdvancedClusterFileSystemLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"CREATING",
		"FAILED",
		"UPDATING",
		"DELETED",
		"DELETING",
	}
}

// GetMappingAdvancedClusterFileSystemLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdvancedClusterFileSystemLifecycleStateEnum(val string) (AdvancedClusterFileSystemLifecycleStateEnum, bool) {
	enum, ok := mappingAdvancedClusterFileSystemLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
