// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Database The representation of Database
type Database struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
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

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId *string `mandatory:"false" json:"dbHomeId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
	VmClusterId *string `mandatory:"false" json:"vmClusterId"`

	// The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
	PdbName *string `mandatory:"false" json:"pdbName"`

	// **Deprecated.** The dbWorkload field has been deprecated for Exadata Database Service on Dedicated Infrastructure, Exadata Database Service on Cloud@Customer, and Base Database Service.
	// Support for this attribute will end in November 2023. You may choose to update your custom scripts to exclude the dbWorkload attribute. After November 2023 if you pass a value to the dbWorkload attribute, it will be ignored.
	// The database workload type.
	DbWorkload *string `mandatory:"false" json:"dbWorkload"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time when the latest database backup was created.
	LastBackupTimestamp *common.SDKTime `mandatory:"false" json:"lastBackupTimestamp"`

	// The duration when the latest database backup created.
	LastBackupDurationInSeconds *int `mandatory:"false" json:"lastBackupDurationInSeconds"`

	// The date and time when the latest database backup failed.
	LastFailedBackupTimestamp *common.SDKTime `mandatory:"false" json:"lastFailedBackupTimestamp"`

	DbBackupConfig *DbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

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

	// The Connection strings used to connect to the Oracle Database.
	ConnectionStrings *DatabaseConnectionStrings `mandatory:"false" json:"connectionStrings"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous AI Database Serverless does not use key versions, hence is not applicable for Autonomous AI Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Point in time recovery timeStamp of the source database at which cloned database system is cloned from the source database system, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339)
	SourceDatabasePointInTimeRecoveryTimestamp *common.SDKTime `mandatory:"false" json:"sourceDatabasePointInTimeRecoveryTimestamp"`

	// The database software image OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`

	// True if the database is a container database.
	IsCdb *bool `mandatory:"false" json:"isCdb"`

	DatabaseManagementConfig *CloudDatabaseManagementConfig `mandatory:"false" json:"databaseManagementConfig"`

	// Specifies a prefix for the `Oracle SID` of the database to be created.
	SidPrefix *string `mandatory:"false" json:"sidPrefix"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// The wallet name for Oracle Key Vault.
	KeyStoreWalletName *string `mandatory:"false" json:"keyStoreWalletName"`

	DataGuardGroup *DataGuardGroup `mandatory:"false" json:"dataGuardGroup"`

	EncryptionKeyLocationDetails EncryptionKeyLocationDetails `mandatory:"false" json:"encryptionKeyLocationDetails"`

	StorageSizeDetails *DatabaseStorageSizeResponseDetails `mandatory:"false" json:"storageSizeDetails"`
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
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Database) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CharacterSet                               *string                             `json:"characterSet"`
		NcharacterSet                              *string                             `json:"ncharacterSet"`
		DbHomeId                                   *string                             `json:"dbHomeId"`
		DbSystemId                                 *string                             `json:"dbSystemId"`
		VmClusterId                                *string                             `json:"vmClusterId"`
		PdbName                                    *string                             `json:"pdbName"`
		DbWorkload                                 *string                             `json:"dbWorkload"`
		LifecycleDetails                           *string                             `json:"lifecycleDetails"`
		TimeCreated                                *common.SDKTime                     `json:"timeCreated"`
		LastBackupTimestamp                        *common.SDKTime                     `json:"lastBackupTimestamp"`
		LastBackupDurationInSeconds                *int                                `json:"lastBackupDurationInSeconds"`
		LastFailedBackupTimestamp                  *common.SDKTime                     `json:"lastFailedBackupTimestamp"`
		DbBackupConfig                             *DbBackupConfig                     `json:"dbBackupConfig"`
		FreeformTags                               map[string]string                   `json:"freeformTags"`
		DefinedTags                                map[string]map[string]interface{}   `json:"definedTags"`
		SystemTags                                 map[string]map[string]interface{}   `json:"systemTags"`
		ConnectionStrings                          *DatabaseConnectionStrings          `json:"connectionStrings"`
		KmsKeyId                                   *string                             `json:"kmsKeyId"`
		KmsKeyVersionId                            *string                             `json:"kmsKeyVersionId"`
		VaultId                                    *string                             `json:"vaultId"`
		SourceDatabasePointInTimeRecoveryTimestamp *common.SDKTime                     `json:"sourceDatabasePointInTimeRecoveryTimestamp"`
		DatabaseSoftwareImageId                    *string                             `json:"databaseSoftwareImageId"`
		IsCdb                                      *bool                               `json:"isCdb"`
		DatabaseManagementConfig                   *CloudDatabaseManagementConfig      `json:"databaseManagementConfig"`
		SidPrefix                                  *string                             `json:"sidPrefix"`
		KeyStoreId                                 *string                             `json:"keyStoreId"`
		KeyStoreWalletName                         *string                             `json:"keyStoreWalletName"`
		DataGuardGroup                             *DataGuardGroup                     `json:"dataGuardGroup"`
		EncryptionKeyLocationDetails               encryptionkeylocationdetails        `json:"encryptionKeyLocationDetails"`
		StorageSizeDetails                         *DatabaseStorageSizeResponseDetails `json:"storageSizeDetails"`
		Id                                         *string                             `json:"id"`
		CompartmentId                              *string                             `json:"compartmentId"`
		DbName                                     *string                             `json:"dbName"`
		DbUniqueName                               *string                             `json:"dbUniqueName"`
		LifecycleState                             DatabaseLifecycleStateEnum          `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CharacterSet = model.CharacterSet

	m.NcharacterSet = model.NcharacterSet

	m.DbHomeId = model.DbHomeId

	m.DbSystemId = model.DbSystemId

	m.VmClusterId = model.VmClusterId

	m.PdbName = model.PdbName

	m.DbWorkload = model.DbWorkload

	m.LifecycleDetails = model.LifecycleDetails

	m.TimeCreated = model.TimeCreated

	m.LastBackupTimestamp = model.LastBackupTimestamp

	m.LastBackupDurationInSeconds = model.LastBackupDurationInSeconds

	m.LastFailedBackupTimestamp = model.LastFailedBackupTimestamp

	m.DbBackupConfig = model.DbBackupConfig

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.ConnectionStrings = model.ConnectionStrings

	m.KmsKeyId = model.KmsKeyId

	m.KmsKeyVersionId = model.KmsKeyVersionId

	m.VaultId = model.VaultId

	m.SourceDatabasePointInTimeRecoveryTimestamp = model.SourceDatabasePointInTimeRecoveryTimestamp

	m.DatabaseSoftwareImageId = model.DatabaseSoftwareImageId

	m.IsCdb = model.IsCdb

	m.DatabaseManagementConfig = model.DatabaseManagementConfig

	m.SidPrefix = model.SidPrefix

	m.KeyStoreId = model.KeyStoreId

	m.KeyStoreWalletName = model.KeyStoreWalletName

	m.DataGuardGroup = model.DataGuardGroup

	nn, e = model.EncryptionKeyLocationDetails.UnmarshalPolymorphicJSON(model.EncryptionKeyLocationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EncryptionKeyLocationDetails = nn.(EncryptionKeyLocationDetails)
	} else {
		m.EncryptionKeyLocationDetails = nil
	}

	m.StorageSizeDetails = model.StorageSizeDetails

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DbName = model.DbName

	m.DbUniqueName = model.DbUniqueName

	m.LifecycleState = model.LifecycleState

	return
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
