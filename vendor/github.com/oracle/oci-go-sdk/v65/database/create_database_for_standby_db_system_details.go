// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateDatabaseForStandbyDbSystemDetails Details for creating a database for a standby db system with dataguard.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateDatabaseForStandbyDbSystemDetails struct {

	// For SYS, SYSTEM, and PDB Admin, enter the same password as the primary admin password.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The protection mode of this Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum `mandatory:"true" json:"protectionMode"`

	// The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
	// * MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
	// * MAXIMUM_PERFORMANCE - ASYNC
	// * MAXIMUM_PROTECTION - SYNC
	// For more information, see
	// Redo Transport Services (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only transport type currently supported by the Database service is ASYNC.
	TransportType CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum `mandatory:"true" json:"transportType"`

	// For TDE Wallet, enter the same password as the primary wallet password.
	TdeWalletPassword *string `mandatory:"false" json:"tdeWalletPassword"`

	// The database software image OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`

	// True if active Data Guard is enabled.
	IsActiveDataGuardEnabled *bool `mandatory:"false" json:"isActiveDataGuardEnabled"`

	// The database domain. In a distributed database system, DB_DOMAIN specifies the logical location of the database within the network structure.
	DbDomain *string `mandatory:"false" json:"dbDomain"`

	// Specifies a prefix for the `Oracle SID` of the database to be created.
	SidPrefix *string `mandatory:"false" json:"sidPrefix"`

	// The `DB_UNIQUE_NAME` of the Oracle Database.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	SourceEncryptionKeyLocationDetails EncryptionKeyLocationDetails `mandatory:"false" json:"sourceEncryptionKeyLocationDetails"`

	DbBackupConfig *DbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	DatabaseFreeformTags map[string]string `mandatory:"false" json:"databaseFreeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DatabaseDefinedTags map[string]map[string]interface{} `mandatory:"false" json:"databaseDefinedTags"`
}

func (m CreateDatabaseForStandbyDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseForStandbyDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetCreateDatabaseForStandbyDbSystemDetailsProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetCreateDatabaseForStandbyDbSystemDetailsTransportTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDatabaseForStandbyDbSystemDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TdeWalletPassword                  *string                                                   `json:"tdeWalletPassword"`
		DatabaseSoftwareImageId            *string                                                   `json:"databaseSoftwareImageId"`
		IsActiveDataGuardEnabled           *bool                                                     `json:"isActiveDataGuardEnabled"`
		DbDomain                           *string                                                   `json:"dbDomain"`
		SidPrefix                          *string                                                   `json:"sidPrefix"`
		DbUniqueName                       *string                                                   `json:"dbUniqueName"`
		SourceEncryptionKeyLocationDetails encryptionkeylocationdetails                              `json:"sourceEncryptionKeyLocationDetails"`
		DbBackupConfig                     *DbBackupConfig                                           `json:"dbBackupConfig"`
		DatabaseFreeformTags               map[string]string                                         `json:"databaseFreeformTags"`
		DatabaseDefinedTags                map[string]map[string]interface{}                         `json:"databaseDefinedTags"`
		AdminPassword                      *string                                                   `json:"adminPassword"`
		ProtectionMode                     CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum `json:"protectionMode"`
		TransportType                      CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum  `json:"transportType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TdeWalletPassword = model.TdeWalletPassword

	m.DatabaseSoftwareImageId = model.DatabaseSoftwareImageId

	m.IsActiveDataGuardEnabled = model.IsActiveDataGuardEnabled

	m.DbDomain = model.DbDomain

	m.SidPrefix = model.SidPrefix

	m.DbUniqueName = model.DbUniqueName

	nn, e = model.SourceEncryptionKeyLocationDetails.UnmarshalPolymorphicJSON(model.SourceEncryptionKeyLocationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceEncryptionKeyLocationDetails = nn.(EncryptionKeyLocationDetails)
	} else {
		m.SourceEncryptionKeyLocationDetails = nil
	}

	m.DbBackupConfig = model.DbBackupConfig

	m.DatabaseFreeformTags = model.DatabaseFreeformTags

	m.DatabaseDefinedTags = model.DatabaseDefinedTags

	m.AdminPassword = model.AdminPassword

	m.ProtectionMode = model.ProtectionMode

	m.TransportType = model.TransportType

	return
}

// CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum Enum with underlying type: string
type CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum string

// Set of constants representing the allowable values for CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum
const (
	CreateDatabaseForStandbyDbSystemDetailsProtectionModeAvailability CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	CreateDatabaseForStandbyDbSystemDetailsProtectionModePerformance  CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	CreateDatabaseForStandbyDbSystemDetailsProtectionModeProtection   CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum = "MAXIMUM_PROTECTION"
)

var mappingCreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum = map[string]CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": CreateDatabaseForStandbyDbSystemDetailsProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  CreateDatabaseForStandbyDbSystemDetailsProtectionModePerformance,
	"MAXIMUM_PROTECTION":   CreateDatabaseForStandbyDbSystemDetailsProtectionModeProtection,
}

var mappingCreateDatabaseForStandbyDbSystemDetailsProtectionModeEnumLowerCase = map[string]CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum{
	"maximum_availability": CreateDatabaseForStandbyDbSystemDetailsProtectionModeAvailability,
	"maximum_performance":  CreateDatabaseForStandbyDbSystemDetailsProtectionModePerformance,
	"maximum_protection":   CreateDatabaseForStandbyDbSystemDetailsProtectionModeProtection,
}

// GetCreateDatabaseForStandbyDbSystemDetailsProtectionModeEnumValues Enumerates the set of values for CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum
func GetCreateDatabaseForStandbyDbSystemDetailsProtectionModeEnumValues() []CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum {
	values := make([]CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum, 0)
	for _, v := range mappingCreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseForStandbyDbSystemDetailsProtectionModeEnumStringValues Enumerates the set of values in String for CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum
func GetCreateDatabaseForStandbyDbSystemDetailsProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
		"MAXIMUM_PROTECTION",
	}
}

// GetMappingCreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum(val string) (CreateDatabaseForStandbyDbSystemDetailsProtectionModeEnum, bool) {
	enum, ok := mappingCreateDatabaseForStandbyDbSystemDetailsProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum Enum with underlying type: string
type CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum string

// Set of constants representing the allowable values for CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum
const (
	CreateDatabaseForStandbyDbSystemDetailsTransportTypeSync     CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum = "SYNC"
	CreateDatabaseForStandbyDbSystemDetailsTransportTypeAsync    CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum = "ASYNC"
	CreateDatabaseForStandbyDbSystemDetailsTransportTypeFastsync CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum = "FASTSYNC"
)

var mappingCreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum = map[string]CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum{
	"SYNC":     CreateDatabaseForStandbyDbSystemDetailsTransportTypeSync,
	"ASYNC":    CreateDatabaseForStandbyDbSystemDetailsTransportTypeAsync,
	"FASTSYNC": CreateDatabaseForStandbyDbSystemDetailsTransportTypeFastsync,
}

var mappingCreateDatabaseForStandbyDbSystemDetailsTransportTypeEnumLowerCase = map[string]CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum{
	"sync":     CreateDatabaseForStandbyDbSystemDetailsTransportTypeSync,
	"async":    CreateDatabaseForStandbyDbSystemDetailsTransportTypeAsync,
	"fastsync": CreateDatabaseForStandbyDbSystemDetailsTransportTypeFastsync,
}

// GetCreateDatabaseForStandbyDbSystemDetailsTransportTypeEnumValues Enumerates the set of values for CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum
func GetCreateDatabaseForStandbyDbSystemDetailsTransportTypeEnumValues() []CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum {
	values := make([]CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum, 0)
	for _, v := range mappingCreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseForStandbyDbSystemDetailsTransportTypeEnumStringValues Enumerates the set of values in String for CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum
func GetCreateDatabaseForStandbyDbSystemDetailsTransportTypeEnumStringValues() []string {
	return []string{
		"SYNC",
		"ASYNC",
		"FASTSYNC",
	}
}

// GetMappingCreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum(val string) (CreateDatabaseForStandbyDbSystemDetailsTransportTypeEnum, bool) {
	enum, ok := mappingCreateDatabaseForStandbyDbSystemDetailsTransportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
