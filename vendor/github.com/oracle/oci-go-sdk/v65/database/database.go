// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// Database The representation of Database
type Database struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database name.
	DbName *string `mandatory:"true" json:"dbName"`

	// A system-generated name for the database to ensure uniqueness within an Oracle Data Guard group (a primary database and its standby databases). The unique name cannot be changed.
	DbUniqueName *string `mandatory:"true" json:"dbUniqueName"`

	// The current state of the database.
	LifecycleState DatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The character set for the database.
	CharacterSet *string `mandatory:"false" json:"characterSet"`

	// The national character set for the database.
	NcharacterSet *string `mandatory:"false" json:"ncharacterSet"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId *string `mandatory:"false" json:"dbHomeId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster.
	VmClusterId *string `mandatory:"false" json:"vmClusterId"`

	// The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
	PdbName *string `mandatory:"false" json:"pdbName"`

	// The database workload type.
	DbWorkload *string `mandatory:"false" json:"dbWorkload"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time when the latest database backup was created.
	LastBackupTimestamp *common.SDKTime `mandatory:"false" json:"lastBackupTimestamp"`

	DbBackupConfig *DbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Connection strings used to connect to the Oracle Database.
	ConnectionStrings *DatabaseConnectionStrings `mandatory:"false" json:"connectionStrings"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Point in time recovery timeStamp of the source database at which cloned database system is cloned from the source database system, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339)
	SourceDatabasePointInTimeRecoveryTimestamp *common.SDKTime `mandatory:"false" json:"sourceDatabasePointInTimeRecoveryTimestamp"`

	// The database software image OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`

	// True if the database is a container database.
	IsCdb *bool `mandatory:"false" json:"isCdb"`

	DatabaseManagementConfig *CloudDatabaseManagementConfig `mandatory:"false" json:"databaseManagementConfig"`

	// Specifies a prefix for the `Oracle SID` of the database to be created.
	SidPrefix *string `mandatory:"false" json:"sidPrefix"`
}

func (m Database) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Database) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseLifecycleStateEnum Enum with underlying type: string
type DatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseLifecycleStateEnum
const (
	DatabaseLifecycleStateProvisioning     DatabaseLifecycleStateEnum = "PROVISIONING"
	DatabaseLifecycleStateAvailable        DatabaseLifecycleStateEnum = "AVAILABLE"
	DatabaseLifecycleStateUpdating         DatabaseLifecycleStateEnum = "UPDATING"
	DatabaseLifecycleStateBackupInProgress DatabaseLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	DatabaseLifecycleStateUpgrading        DatabaseLifecycleStateEnum = "UPGRADING"
	DatabaseLifecycleStateConverting       DatabaseLifecycleStateEnum = "CONVERTING"
	DatabaseLifecycleStateTerminating      DatabaseLifecycleStateEnum = "TERMINATING"
	DatabaseLifecycleStateTerminated       DatabaseLifecycleStateEnum = "TERMINATED"
	DatabaseLifecycleStateRestoreFailed    DatabaseLifecycleStateEnum = "RESTORE_FAILED"
	DatabaseLifecycleStateFailed           DatabaseLifecycleStateEnum = "FAILED"
)

var mappingDatabaseLifecycleStateEnum = map[string]DatabaseLifecycleStateEnum{
	"PROVISIONING":       DatabaseLifecycleStateProvisioning,
	"AVAILABLE":          DatabaseLifecycleStateAvailable,
	"UPDATING":           DatabaseLifecycleStateUpdating,
	"BACKUP_IN_PROGRESS": DatabaseLifecycleStateBackupInProgress,
	"UPGRADING":          DatabaseLifecycleStateUpgrading,
	"CONVERTING":         DatabaseLifecycleStateConverting,
	"TERMINATING":        DatabaseLifecycleStateTerminating,
	"TERMINATED":         DatabaseLifecycleStateTerminated,
	"RESTORE_FAILED":     DatabaseLifecycleStateRestoreFailed,
	"FAILED":             DatabaseLifecycleStateFailed,
}

var mappingDatabaseLifecycleStateEnumLowerCase = map[string]DatabaseLifecycleStateEnum{
	"provisioning":       DatabaseLifecycleStateProvisioning,
	"available":          DatabaseLifecycleStateAvailable,
	"updating":           DatabaseLifecycleStateUpdating,
	"backup_in_progress": DatabaseLifecycleStateBackupInProgress,
	"upgrading":          DatabaseLifecycleStateUpgrading,
	"converting":         DatabaseLifecycleStateConverting,
	"terminating":        DatabaseLifecycleStateTerminating,
	"terminated":         DatabaseLifecycleStateTerminated,
	"restore_failed":     DatabaseLifecycleStateRestoreFailed,
	"failed":             DatabaseLifecycleStateFailed,
}

// GetDatabaseLifecycleStateEnumValues Enumerates the set of values for DatabaseLifecycleStateEnum
func GetDatabaseLifecycleStateEnumValues() []DatabaseLifecycleStateEnum {
	values := make([]DatabaseLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseLifecycleStateEnumStringValues Enumerates the set of values in String for DatabaseLifecycleStateEnum
func GetDatabaseLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"BACKUP_IN_PROGRESS",
		"UPGRADING",
		"CONVERTING",
		"TERMINATING",
		"TERMINATED",
		"RESTORE_FAILED",
		"FAILED",
	}
}

// GetMappingDatabaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseLifecycleStateEnum(val string) (DatabaseLifecycleStateEnum, bool) {
	enum, ok := mappingDatabaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
