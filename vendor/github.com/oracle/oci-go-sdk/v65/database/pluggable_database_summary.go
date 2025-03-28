// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// PluggableDatabaseSummary A pluggable database (PDB) is portable collection of schemas, schema objects, and non-schema objects that appears to an Oracle client as a non-container database. To use a PDB, it needs to be plugged into a CDB.
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized, talk to a tenancy administrator. If you are an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type PluggableDatabaseSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pluggable database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB.
	ContainerDatabaseId *string `mandatory:"true" json:"containerDatabaseId"`

	// The name for the pluggable database (PDB). The name is unique in the context of a Database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	PdbName *string `mandatory:"true" json:"pdbName"`

	// The current state of the pluggable database.
	LifecycleState PluggableDatabaseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the pluggable database was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// **Deprecated.** Use PluggableDatabaseNodeLevelDetails for OpenMode details.
	// The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software).
	OpenMode PluggableDatabaseSummaryOpenModeEnum `mandatory:"true" json:"openMode"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Detailed message for the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	ConnectionStrings *PluggableDatabaseConnectionStrings `mandatory:"false" json:"connectionStrings"`

	// The restricted mode of the pluggable database. If a pluggable database is opened in restricted mode,
	// the user needs both create a session and have restricted session privileges to connect to it.
	IsRestricted *bool `mandatory:"false" json:"isRestricted"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	PluggableDatabaseManagementConfig *PluggableDatabaseManagementConfig `mandatory:"false" json:"pluggableDatabaseManagementConfig"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	RefreshableCloneConfig *PluggableDatabaseRefreshableCloneConfig `mandatory:"false" json:"refreshableCloneConfig"`

	// Pluggable Database Node Level Details.
	// Example: [{"nodeName" : "node1", "openMode" : "READ_WRITE"}, {"nodeName" : "node2", "openMode" : "READ_ONLY"}]
	PdbNodeLevelDetails []PluggableDatabaseNodeLevelDetails `mandatory:"false" json:"pdbNodeLevelDetails"`
}

func (m PluggableDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PluggableDatabaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPluggableDatabaseSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPluggableDatabaseSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPluggableDatabaseSummaryOpenModeEnum(string(m.OpenMode)); !ok && m.OpenMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpenMode: %s. Supported values are: %s.", m.OpenMode, strings.Join(GetPluggableDatabaseSummaryOpenModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PluggableDatabaseSummaryLifecycleStateEnum Enum with underlying type: string
type PluggableDatabaseSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for PluggableDatabaseSummaryLifecycleStateEnum
const (
	PluggableDatabaseSummaryLifecycleStateProvisioning      PluggableDatabaseSummaryLifecycleStateEnum = "PROVISIONING"
	PluggableDatabaseSummaryLifecycleStateAvailable         PluggableDatabaseSummaryLifecycleStateEnum = "AVAILABLE"
	PluggableDatabaseSummaryLifecycleStateTerminating       PluggableDatabaseSummaryLifecycleStateEnum = "TERMINATING"
	PluggableDatabaseSummaryLifecycleStateTerminated        PluggableDatabaseSummaryLifecycleStateEnum = "TERMINATED"
	PluggableDatabaseSummaryLifecycleStateUpdating          PluggableDatabaseSummaryLifecycleStateEnum = "UPDATING"
	PluggableDatabaseSummaryLifecycleStateFailed            PluggableDatabaseSummaryLifecycleStateEnum = "FAILED"
	PluggableDatabaseSummaryLifecycleStateRelocating        PluggableDatabaseSummaryLifecycleStateEnum = "RELOCATING"
	PluggableDatabaseSummaryLifecycleStateRelocated         PluggableDatabaseSummaryLifecycleStateEnum = "RELOCATED"
	PluggableDatabaseSummaryLifecycleStateRefreshing        PluggableDatabaseSummaryLifecycleStateEnum = "REFRESHING"
	PluggableDatabaseSummaryLifecycleStateRestoreInProgress PluggableDatabaseSummaryLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	PluggableDatabaseSummaryLifecycleStateRestoreFailed     PluggableDatabaseSummaryLifecycleStateEnum = "RESTORE_FAILED"
	PluggableDatabaseSummaryLifecycleStateBackupInProgress  PluggableDatabaseSummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	PluggableDatabaseSummaryLifecycleStateDisabled          PluggableDatabaseSummaryLifecycleStateEnum = "DISABLED"
)

var mappingPluggableDatabaseSummaryLifecycleStateEnum = map[string]PluggableDatabaseSummaryLifecycleStateEnum{
	"PROVISIONING":        PluggableDatabaseSummaryLifecycleStateProvisioning,
	"AVAILABLE":           PluggableDatabaseSummaryLifecycleStateAvailable,
	"TERMINATING":         PluggableDatabaseSummaryLifecycleStateTerminating,
	"TERMINATED":          PluggableDatabaseSummaryLifecycleStateTerminated,
	"UPDATING":            PluggableDatabaseSummaryLifecycleStateUpdating,
	"FAILED":              PluggableDatabaseSummaryLifecycleStateFailed,
	"RELOCATING":          PluggableDatabaseSummaryLifecycleStateRelocating,
	"RELOCATED":           PluggableDatabaseSummaryLifecycleStateRelocated,
	"REFRESHING":          PluggableDatabaseSummaryLifecycleStateRefreshing,
	"RESTORE_IN_PROGRESS": PluggableDatabaseSummaryLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":      PluggableDatabaseSummaryLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":  PluggableDatabaseSummaryLifecycleStateBackupInProgress,
	"DISABLED":            PluggableDatabaseSummaryLifecycleStateDisabled,
}

var mappingPluggableDatabaseSummaryLifecycleStateEnumLowerCase = map[string]PluggableDatabaseSummaryLifecycleStateEnum{
	"provisioning":        PluggableDatabaseSummaryLifecycleStateProvisioning,
	"available":           PluggableDatabaseSummaryLifecycleStateAvailable,
	"terminating":         PluggableDatabaseSummaryLifecycleStateTerminating,
	"terminated":          PluggableDatabaseSummaryLifecycleStateTerminated,
	"updating":            PluggableDatabaseSummaryLifecycleStateUpdating,
	"failed":              PluggableDatabaseSummaryLifecycleStateFailed,
	"relocating":          PluggableDatabaseSummaryLifecycleStateRelocating,
	"relocated":           PluggableDatabaseSummaryLifecycleStateRelocated,
	"refreshing":          PluggableDatabaseSummaryLifecycleStateRefreshing,
	"restore_in_progress": PluggableDatabaseSummaryLifecycleStateRestoreInProgress,
	"restore_failed":      PluggableDatabaseSummaryLifecycleStateRestoreFailed,
	"backup_in_progress":  PluggableDatabaseSummaryLifecycleStateBackupInProgress,
	"disabled":            PluggableDatabaseSummaryLifecycleStateDisabled,
}

// GetPluggableDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for PluggableDatabaseSummaryLifecycleStateEnum
func GetPluggableDatabaseSummaryLifecycleStateEnumValues() []PluggableDatabaseSummaryLifecycleStateEnum {
	values := make([]PluggableDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingPluggableDatabaseSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPluggableDatabaseSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for PluggableDatabaseSummaryLifecycleStateEnum
func GetPluggableDatabaseSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingPluggableDatabaseSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluggableDatabaseSummaryLifecycleStateEnum(val string) (PluggableDatabaseSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingPluggableDatabaseSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PluggableDatabaseSummaryOpenModeEnum Enum with underlying type: string
type PluggableDatabaseSummaryOpenModeEnum string

// Set of constants representing the allowable values for PluggableDatabaseSummaryOpenModeEnum
const (
	PluggableDatabaseSummaryOpenModeReadOnly  PluggableDatabaseSummaryOpenModeEnum = "READ_ONLY"
	PluggableDatabaseSummaryOpenModeReadWrite PluggableDatabaseSummaryOpenModeEnum = "READ_WRITE"
	PluggableDatabaseSummaryOpenModeMounted   PluggableDatabaseSummaryOpenModeEnum = "MOUNTED"
	PluggableDatabaseSummaryOpenModeMigrate   PluggableDatabaseSummaryOpenModeEnum = "MIGRATE"
)

var mappingPluggableDatabaseSummaryOpenModeEnum = map[string]PluggableDatabaseSummaryOpenModeEnum{
	"READ_ONLY":  PluggableDatabaseSummaryOpenModeReadOnly,
	"READ_WRITE": PluggableDatabaseSummaryOpenModeReadWrite,
	"MOUNTED":    PluggableDatabaseSummaryOpenModeMounted,
	"MIGRATE":    PluggableDatabaseSummaryOpenModeMigrate,
}

var mappingPluggableDatabaseSummaryOpenModeEnumLowerCase = map[string]PluggableDatabaseSummaryOpenModeEnum{
	"read_only":  PluggableDatabaseSummaryOpenModeReadOnly,
	"read_write": PluggableDatabaseSummaryOpenModeReadWrite,
	"mounted":    PluggableDatabaseSummaryOpenModeMounted,
	"migrate":    PluggableDatabaseSummaryOpenModeMigrate,
}

// GetPluggableDatabaseSummaryOpenModeEnumValues Enumerates the set of values for PluggableDatabaseSummaryOpenModeEnum
func GetPluggableDatabaseSummaryOpenModeEnumValues() []PluggableDatabaseSummaryOpenModeEnum {
	values := make([]PluggableDatabaseSummaryOpenModeEnum, 0)
	for _, v := range mappingPluggableDatabaseSummaryOpenModeEnum {
		values = append(values, v)
	}
	return values
}

// GetPluggableDatabaseSummaryOpenModeEnumStringValues Enumerates the set of values in String for PluggableDatabaseSummaryOpenModeEnum
func GetPluggableDatabaseSummaryOpenModeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
		"MOUNTED",
		"MIGRATE",
	}
}

// GetMappingPluggableDatabaseSummaryOpenModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluggableDatabaseSummaryOpenModeEnum(val string) (PluggableDatabaseSummaryOpenModeEnum, bool) {
	enum, ok := mappingPluggableDatabaseSummaryOpenModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
