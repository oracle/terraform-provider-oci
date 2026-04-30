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

// UpdateDatabaseDetails Details to update a database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateDatabaseDetails struct {
	DbBackupConfig *DbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId *string `mandatory:"false" json:"dbHomeId"`

	// A new strong password for SYS, SYSTEM, and the plugbable database ADMIN user. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numeric, and two special characters. The special characters must be _, \#, or -.
	NewAdminPassword *string `mandatory:"false" json:"newAdminPassword"`

	// The existing TDE wallet password. You must provide the existing password in order to set a new TDE wallet password.
	OldTdeWalletPassword *string `mandatory:"false" json:"oldTdeWalletPassword"`

	// The new password to open the TDE wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numeric, and two special characters. The special characters must be _, \#, or -.
	NewTdeWalletPassword *string `mandatory:"false" json:"newTdeWalletPassword"`

	StorageSizeDetails *DatabaseStorageSizeDetails `mandatory:"false" json:"storageSizeDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	ManagedSoftwareUpdateDetails *ManagedSoftwareUpdateInputDetails `mandatory:"false" json:"managedSoftwareUpdateDetails"`

	PatchOptions *PatchOptions `mandatory:"false" json:"patchOptions"`

	// The administrator password of the primary database in this Data Guard association.
	// **The password MUST be the same as the primary admin password.**
	DatabaseAdminPassword *string `mandatory:"false" json:"databaseAdminPassword"`

	// The protection mode of this Data Guard. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode UpdateDatabaseDetailsProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
	// * MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
	// * MAXIMUM_PERFORMANCE - ASYNC
	// * MAXIMUM_PROTECTION - SYNC
	// For more information, see
	// Redo Transport Services (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only transport type currently supported by the Database service is ASYNC.
	TransportType UpdateDatabaseDetailsTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`

	// True if active Data Guard is enabled.
	IsActiveDataGuardEnabled *bool `mandatory:"false" json:"isActiveDataGuardEnabled"`

	AutoFailoverConfiguration *AutoFailoverConfiguration `mandatory:"false" json:"autoFailoverConfiguration"`

	DataSafeRegistrationDetails *DataSafeRegistrationRequestDetails `mandatory:"false" json:"dataSafeRegistrationDetails"`
}

func (m UpdateDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateDatabaseDetailsProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetUpdateDatabaseDetailsProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateDatabaseDetailsTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetUpdateDatabaseDetailsTransportTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDatabaseDetailsProtectionModeEnum Enum with underlying type: string
type UpdateDatabaseDetailsProtectionModeEnum string

// Set of constants representing the allowable values for UpdateDatabaseDetailsProtectionModeEnum
const (
	UpdateDatabaseDetailsProtectionModeAvailability UpdateDatabaseDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	UpdateDatabaseDetailsProtectionModePerformance  UpdateDatabaseDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	UpdateDatabaseDetailsProtectionModeProtection   UpdateDatabaseDetailsProtectionModeEnum = "MAXIMUM_PROTECTION"
)

var mappingUpdateDatabaseDetailsProtectionModeEnum = map[string]UpdateDatabaseDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": UpdateDatabaseDetailsProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  UpdateDatabaseDetailsProtectionModePerformance,
	"MAXIMUM_PROTECTION":   UpdateDatabaseDetailsProtectionModeProtection,
}

var mappingUpdateDatabaseDetailsProtectionModeEnumLowerCase = map[string]UpdateDatabaseDetailsProtectionModeEnum{
	"maximum_availability": UpdateDatabaseDetailsProtectionModeAvailability,
	"maximum_performance":  UpdateDatabaseDetailsProtectionModePerformance,
	"maximum_protection":   UpdateDatabaseDetailsProtectionModeProtection,
}

// GetUpdateDatabaseDetailsProtectionModeEnumValues Enumerates the set of values for UpdateDatabaseDetailsProtectionModeEnum
func GetUpdateDatabaseDetailsProtectionModeEnumValues() []UpdateDatabaseDetailsProtectionModeEnum {
	values := make([]UpdateDatabaseDetailsProtectionModeEnum, 0)
	for _, v := range mappingUpdateDatabaseDetailsProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDatabaseDetailsProtectionModeEnumStringValues Enumerates the set of values in String for UpdateDatabaseDetailsProtectionModeEnum
func GetUpdateDatabaseDetailsProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
		"MAXIMUM_PROTECTION",
	}
}

// GetMappingUpdateDatabaseDetailsProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDatabaseDetailsProtectionModeEnum(val string) (UpdateDatabaseDetailsProtectionModeEnum, bool) {
	enum, ok := mappingUpdateDatabaseDetailsProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateDatabaseDetailsTransportTypeEnum Enum with underlying type: string
type UpdateDatabaseDetailsTransportTypeEnum string

// Set of constants representing the allowable values for UpdateDatabaseDetailsTransportTypeEnum
const (
	UpdateDatabaseDetailsTransportTypeSync     UpdateDatabaseDetailsTransportTypeEnum = "SYNC"
	UpdateDatabaseDetailsTransportTypeAsync    UpdateDatabaseDetailsTransportTypeEnum = "ASYNC"
	UpdateDatabaseDetailsTransportTypeFastsync UpdateDatabaseDetailsTransportTypeEnum = "FASTSYNC"
)

var mappingUpdateDatabaseDetailsTransportTypeEnum = map[string]UpdateDatabaseDetailsTransportTypeEnum{
	"SYNC":     UpdateDatabaseDetailsTransportTypeSync,
	"ASYNC":    UpdateDatabaseDetailsTransportTypeAsync,
	"FASTSYNC": UpdateDatabaseDetailsTransportTypeFastsync,
}

var mappingUpdateDatabaseDetailsTransportTypeEnumLowerCase = map[string]UpdateDatabaseDetailsTransportTypeEnum{
	"sync":     UpdateDatabaseDetailsTransportTypeSync,
	"async":    UpdateDatabaseDetailsTransportTypeAsync,
	"fastsync": UpdateDatabaseDetailsTransportTypeFastsync,
}

// GetUpdateDatabaseDetailsTransportTypeEnumValues Enumerates the set of values for UpdateDatabaseDetailsTransportTypeEnum
func GetUpdateDatabaseDetailsTransportTypeEnumValues() []UpdateDatabaseDetailsTransportTypeEnum {
	values := make([]UpdateDatabaseDetailsTransportTypeEnum, 0)
	for _, v := range mappingUpdateDatabaseDetailsTransportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDatabaseDetailsTransportTypeEnumStringValues Enumerates the set of values in String for UpdateDatabaseDetailsTransportTypeEnum
func GetUpdateDatabaseDetailsTransportTypeEnumStringValues() []string {
	return []string{
		"SYNC",
		"ASYNC",
		"FASTSYNC",
	}
}

// GetMappingUpdateDatabaseDetailsTransportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDatabaseDetailsTransportTypeEnum(val string) (UpdateDatabaseDetailsTransportTypeEnum, bool) {
	enum, ok := mappingUpdateDatabaseDetailsTransportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
