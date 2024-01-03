// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PluggableDatabase A pluggable database (PDB) is portable collection of schemas, schema objects, and non-schema objects that appears to an Oracle client as a non-container database. To use a PDB, it needs to be plugged into a CDB.
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized, talk to a tenancy administrator. If you are an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type PluggableDatabase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the pluggable database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the CDB.
	ContainerDatabaseId *string `mandatory:"true" json:"containerDatabaseId"`

	// The name for the pluggable database (PDB). The name is unique in the context of a Database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	PdbName *string `mandatory:"true" json:"pdbName"`

	// The current state of the pluggable database.
	LifecycleState PluggableDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the pluggable database was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// **Deprecated.** Use PluggableDatabaseNodeLevelDetails for OpenMode details.
	// The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software).
	OpenMode PluggableDatabaseOpenModeEnum `mandatory:"true" json:"openMode"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Detailed message for the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	ConnectionStrings *PluggableDatabaseConnectionStrings `mandatory:"false" json:"connectionStrings"`

	// The restricted mode of the pluggable database. If a pluggable database is opened in restricted mode,
	// the user needs both create a session and have restricted session privileges to connect to it.
	IsRestricted *bool `mandatory:"false" json:"isRestricted"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	PluggableDatabaseManagementConfig *PluggableDatabaseManagementConfig `mandatory:"false" json:"pluggableDatabaseManagementConfig"`

	RefreshableCloneConfig *PluggableDatabaseRefreshableCloneConfig `mandatory:"false" json:"refreshableCloneConfig"`

	// Pluggable Database Node Level Details.
	// Example: [{"nodeName" : "node1", "openMode" : "READ_WRITE"}, {"nodeName" : "node2", "openMode" : "READ_ONLY"}]
	PdbNodeLevelDetails []PluggableDatabaseNodeLevelDetails `mandatory:"false" json:"pdbNodeLevelDetails"`
}

func (m PluggableDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PluggableDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPluggableDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPluggableDatabaseLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPluggableDatabaseOpenModeEnum(string(m.OpenMode)); !ok && m.OpenMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpenMode: %s. Supported values are: %s.", m.OpenMode, strings.Join(GetPluggableDatabaseOpenModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PluggableDatabaseLifecycleStateEnum Enum with underlying type: string
type PluggableDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for PluggableDatabaseLifecycleStateEnum
const (
	PluggableDatabaseLifecycleStateProvisioning      PluggableDatabaseLifecycleStateEnum = "PROVISIONING"
	PluggableDatabaseLifecycleStateAvailable         PluggableDatabaseLifecycleStateEnum = "AVAILABLE"
	PluggableDatabaseLifecycleStateTerminating       PluggableDatabaseLifecycleStateEnum = "TERMINATING"
	PluggableDatabaseLifecycleStateTerminated        PluggableDatabaseLifecycleStateEnum = "TERMINATED"
	PluggableDatabaseLifecycleStateUpdating          PluggableDatabaseLifecycleStateEnum = "UPDATING"
	PluggableDatabaseLifecycleStateFailed            PluggableDatabaseLifecycleStateEnum = "FAILED"
	PluggableDatabaseLifecycleStateRelocating        PluggableDatabaseLifecycleStateEnum = "RELOCATING"
	PluggableDatabaseLifecycleStateRelocated         PluggableDatabaseLifecycleStateEnum = "RELOCATED"
	PluggableDatabaseLifecycleStateRefreshing        PluggableDatabaseLifecycleStateEnum = "REFRESHING"
	PluggableDatabaseLifecycleStateRestoreInProgress PluggableDatabaseLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	PluggableDatabaseLifecycleStateRestoreFailed     PluggableDatabaseLifecycleStateEnum = "RESTORE_FAILED"
	PluggableDatabaseLifecycleStateBackupInProgress  PluggableDatabaseLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	PluggableDatabaseLifecycleStateDisabled          PluggableDatabaseLifecycleStateEnum = "DISABLED"
)

var mappingPluggableDatabaseLifecycleStateEnum = map[string]PluggableDatabaseLifecycleStateEnum{
	"PROVISIONING":        PluggableDatabaseLifecycleStateProvisioning,
	"AVAILABLE":           PluggableDatabaseLifecycleStateAvailable,
	"TERMINATING":         PluggableDatabaseLifecycleStateTerminating,
	"TERMINATED":          PluggableDatabaseLifecycleStateTerminated,
	"UPDATING":            PluggableDatabaseLifecycleStateUpdating,
	"FAILED":              PluggableDatabaseLifecycleStateFailed,
	"RELOCATING":          PluggableDatabaseLifecycleStateRelocating,
	"RELOCATED":           PluggableDatabaseLifecycleStateRelocated,
	"REFRESHING":          PluggableDatabaseLifecycleStateRefreshing,
	"RESTORE_IN_PROGRESS": PluggableDatabaseLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":      PluggableDatabaseLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":  PluggableDatabaseLifecycleStateBackupInProgress,
	"DISABLED":            PluggableDatabaseLifecycleStateDisabled,
}

var mappingPluggableDatabaseLifecycleStateEnumLowerCase = map[string]PluggableDatabaseLifecycleStateEnum{
	"provisioning":        PluggableDatabaseLifecycleStateProvisioning,
	"available":           PluggableDatabaseLifecycleStateAvailable,
	"terminating":         PluggableDatabaseLifecycleStateTerminating,
	"terminated":          PluggableDatabaseLifecycleStateTerminated,
	"updating":            PluggableDatabaseLifecycleStateUpdating,
	"failed":              PluggableDatabaseLifecycleStateFailed,
	"relocating":          PluggableDatabaseLifecycleStateRelocating,
	"relocated":           PluggableDatabaseLifecycleStateRelocated,
	"refreshing":          PluggableDatabaseLifecycleStateRefreshing,
	"restore_in_progress": PluggableDatabaseLifecycleStateRestoreInProgress,
	"restore_failed":      PluggableDatabaseLifecycleStateRestoreFailed,
	"backup_in_progress":  PluggableDatabaseLifecycleStateBackupInProgress,
	"disabled":            PluggableDatabaseLifecycleStateDisabled,
}

// GetPluggableDatabaseLifecycleStateEnumValues Enumerates the set of values for PluggableDatabaseLifecycleStateEnum
func GetPluggableDatabaseLifecycleStateEnumValues() []PluggableDatabaseLifecycleStateEnum {
	values := make([]PluggableDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingPluggableDatabaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPluggableDatabaseLifecycleStateEnumStringValues Enumerates the set of values in String for PluggableDatabaseLifecycleStateEnum
func GetPluggableDatabaseLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
		"UPDATING",
		"FAILED",
		"RELOCATING",
		"RELOCATED",
		"REFRESHING",
		"RESTORE_IN_PROGRESS",
		"RESTORE_FAILED",
		"BACKUP_IN_PROGRESS",
		"DISABLED",
	}
}

// GetMappingPluggableDatabaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluggableDatabaseLifecycleStateEnum(val string) (PluggableDatabaseLifecycleStateEnum, bool) {
	enum, ok := mappingPluggableDatabaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PluggableDatabaseOpenModeEnum Enum with underlying type: string
type PluggableDatabaseOpenModeEnum string

// Set of constants representing the allowable values for PluggableDatabaseOpenModeEnum
const (
	PluggableDatabaseOpenModeReadOnly  PluggableDatabaseOpenModeEnum = "READ_ONLY"
	PluggableDatabaseOpenModeReadWrite PluggableDatabaseOpenModeEnum = "READ_WRITE"
	PluggableDatabaseOpenModeMounted   PluggableDatabaseOpenModeEnum = "MOUNTED"
	PluggableDatabaseOpenModeMigrate   PluggableDatabaseOpenModeEnum = "MIGRATE"
)

var mappingPluggableDatabaseOpenModeEnum = map[string]PluggableDatabaseOpenModeEnum{
	"READ_ONLY":  PluggableDatabaseOpenModeReadOnly,
	"READ_WRITE": PluggableDatabaseOpenModeReadWrite,
	"MOUNTED":    PluggableDatabaseOpenModeMounted,
	"MIGRATE":    PluggableDatabaseOpenModeMigrate,
}

var mappingPluggableDatabaseOpenModeEnumLowerCase = map[string]PluggableDatabaseOpenModeEnum{
	"read_only":  PluggableDatabaseOpenModeReadOnly,
	"read_write": PluggableDatabaseOpenModeReadWrite,
	"mounted":    PluggableDatabaseOpenModeMounted,
	"migrate":    PluggableDatabaseOpenModeMigrate,
}

// GetPluggableDatabaseOpenModeEnumValues Enumerates the set of values for PluggableDatabaseOpenModeEnum
func GetPluggableDatabaseOpenModeEnumValues() []PluggableDatabaseOpenModeEnum {
	values := make([]PluggableDatabaseOpenModeEnum, 0)
	for _, v := range mappingPluggableDatabaseOpenModeEnum {
		values = append(values, v)
	}
	return values
}

// GetPluggableDatabaseOpenModeEnumStringValues Enumerates the set of values in String for PluggableDatabaseOpenModeEnum
func GetPluggableDatabaseOpenModeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
		"MOUNTED",
		"MIGRATE",
	}
}

// GetMappingPluggableDatabaseOpenModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluggableDatabaseOpenModeEnum(val string) (PluggableDatabaseOpenModeEnum, bool) {
	enum, ok := mappingPluggableDatabaseOpenModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
