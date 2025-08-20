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

// CreateStandbyDetails Standby Creation Details.
type CreateStandbyDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source database.
	SourceDatabaseId *string `mandatory:"true" json:"sourceDatabaseId"`

	// The administrator password of the primary database in this Data Guard association.
	// **The password MUST be the same as the primary admin password.**
	DatabaseAdminPassword *string `mandatory:"true" json:"databaseAdminPassword"`

	// The TDE wallet password of the source database specified by 'sourceDatabaseId'.
	SourceTdeWalletPassword *string `mandatory:"true" json:"sourceTdeWalletPassword"`

	// The protection mode of this Data Guard. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode CreateStandbyDetailsProtectionModeEnum `mandatory:"true" json:"protectionMode"`

	// The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
	// * MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
	// * MAXIMUM_PERFORMANCE - ASYNC
	// * MAXIMUM_PROTECTION - SYNC
	// For more information, see
	// Redo Transport Services (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only transport type currently supported by the Database service is ASYNC.
	TransportType CreateStandbyDetailsTransportTypeEnum `mandatory:"true" json:"transportType"`

	SourceEncryptionKeyLocationDetails EncryptionKeyLocationDetails `mandatory:"false" json:"sourceEncryptionKeyLocationDetails"`

	// True if active Data Guard is enabled.
	IsActiveDataGuardEnabled *bool `mandatory:"false" json:"isActiveDataGuardEnabled"`

	// Specifies the `DB_UNIQUE_NAME` of the peer database to be created.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// Specifies a prefix for the `Oracle SID` of the database to be created.
	SidPrefix *string `mandatory:"false" json:"sidPrefix"`

	StorageSizeDetails *DatabaseStorageSizeDetails `mandatory:"false" json:"storageSizeDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateStandbyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateStandbyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateStandbyDetailsProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetCreateStandbyDetailsProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateStandbyDetailsTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetCreateStandbyDetailsTransportTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateStandbyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SourceEncryptionKeyLocationDetails encryptionkeylocationdetails           `json:"sourceEncryptionKeyLocationDetails"`
		IsActiveDataGuardEnabled           *bool                                  `json:"isActiveDataGuardEnabled"`
		DbUniqueName                       *string                                `json:"dbUniqueName"`
		SidPrefix                          *string                                `json:"sidPrefix"`
		StorageSizeDetails                 *DatabaseStorageSizeDetails            `json:"storageSizeDetails"`
		FreeformTags                       map[string]string                      `json:"freeformTags"`
		DefinedTags                        map[string]map[string]interface{}      `json:"definedTags"`
		SourceDatabaseId                   *string                                `json:"sourceDatabaseId"`
		DatabaseAdminPassword              *string                                `json:"databaseAdminPassword"`
		SourceTdeWalletPassword            *string                                `json:"sourceTdeWalletPassword"`
		ProtectionMode                     CreateStandbyDetailsProtectionModeEnum `json:"protectionMode"`
		TransportType                      CreateStandbyDetailsTransportTypeEnum  `json:"transportType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.SourceEncryptionKeyLocationDetails.UnmarshalPolymorphicJSON(model.SourceEncryptionKeyLocationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceEncryptionKeyLocationDetails = nn.(EncryptionKeyLocationDetails)
	} else {
		m.SourceEncryptionKeyLocationDetails = nil
	}

	m.IsActiveDataGuardEnabled = model.IsActiveDataGuardEnabled

	m.DbUniqueName = model.DbUniqueName

	m.SidPrefix = model.SidPrefix

	m.StorageSizeDetails = model.StorageSizeDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SourceDatabaseId = model.SourceDatabaseId

	m.DatabaseAdminPassword = model.DatabaseAdminPassword

	m.SourceTdeWalletPassword = model.SourceTdeWalletPassword

	m.ProtectionMode = model.ProtectionMode

	m.TransportType = model.TransportType

	return
}

// CreateStandbyDetailsProtectionModeEnum Enum with underlying type: string
type CreateStandbyDetailsProtectionModeEnum string

// Set of constants representing the allowable values for CreateStandbyDetailsProtectionModeEnum
const (
	CreateStandbyDetailsProtectionModeAvailability CreateStandbyDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	CreateStandbyDetailsProtectionModePerformance  CreateStandbyDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	CreateStandbyDetailsProtectionModeProtection   CreateStandbyDetailsProtectionModeEnum = "MAXIMUM_PROTECTION"
)

var mappingCreateStandbyDetailsProtectionModeEnum = map[string]CreateStandbyDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": CreateStandbyDetailsProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  CreateStandbyDetailsProtectionModePerformance,
	"MAXIMUM_PROTECTION":   CreateStandbyDetailsProtectionModeProtection,
}

var mappingCreateStandbyDetailsProtectionModeEnumLowerCase = map[string]CreateStandbyDetailsProtectionModeEnum{
	"maximum_availability": CreateStandbyDetailsProtectionModeAvailability,
	"maximum_performance":  CreateStandbyDetailsProtectionModePerformance,
	"maximum_protection":   CreateStandbyDetailsProtectionModeProtection,
}

// GetCreateStandbyDetailsProtectionModeEnumValues Enumerates the set of values for CreateStandbyDetailsProtectionModeEnum
func GetCreateStandbyDetailsProtectionModeEnumValues() []CreateStandbyDetailsProtectionModeEnum {
	values := make([]CreateStandbyDetailsProtectionModeEnum, 0)
	for _, v := range mappingCreateStandbyDetailsProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateStandbyDetailsProtectionModeEnumStringValues Enumerates the set of values in String for CreateStandbyDetailsProtectionModeEnum
func GetCreateStandbyDetailsProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
		"MAXIMUM_PROTECTION",
	}
}

// GetMappingCreateStandbyDetailsProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateStandbyDetailsProtectionModeEnum(val string) (CreateStandbyDetailsProtectionModeEnum, bool) {
	enum, ok := mappingCreateStandbyDetailsProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateStandbyDetailsTransportTypeEnum Enum with underlying type: string
type CreateStandbyDetailsTransportTypeEnum string

// Set of constants representing the allowable values for CreateStandbyDetailsTransportTypeEnum
const (
	CreateStandbyDetailsTransportTypeSync     CreateStandbyDetailsTransportTypeEnum = "SYNC"
	CreateStandbyDetailsTransportTypeAsync    CreateStandbyDetailsTransportTypeEnum = "ASYNC"
	CreateStandbyDetailsTransportTypeFastsync CreateStandbyDetailsTransportTypeEnum = "FASTSYNC"
)

var mappingCreateStandbyDetailsTransportTypeEnum = map[string]CreateStandbyDetailsTransportTypeEnum{
	"SYNC":     CreateStandbyDetailsTransportTypeSync,
	"ASYNC":    CreateStandbyDetailsTransportTypeAsync,
	"FASTSYNC": CreateStandbyDetailsTransportTypeFastsync,
}

var mappingCreateStandbyDetailsTransportTypeEnumLowerCase = map[string]CreateStandbyDetailsTransportTypeEnum{
	"sync":     CreateStandbyDetailsTransportTypeSync,
	"async":    CreateStandbyDetailsTransportTypeAsync,
	"fastsync": CreateStandbyDetailsTransportTypeFastsync,
}

// GetCreateStandbyDetailsTransportTypeEnumValues Enumerates the set of values for CreateStandbyDetailsTransportTypeEnum
func GetCreateStandbyDetailsTransportTypeEnumValues() []CreateStandbyDetailsTransportTypeEnum {
	values := make([]CreateStandbyDetailsTransportTypeEnum, 0)
	for _, v := range mappingCreateStandbyDetailsTransportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateStandbyDetailsTransportTypeEnumStringValues Enumerates the set of values in String for CreateStandbyDetailsTransportTypeEnum
func GetCreateStandbyDetailsTransportTypeEnumStringValues() []string {
	return []string{
		"SYNC",
		"ASYNC",
		"FASTSYNC",
	}
}

// GetMappingCreateStandbyDetailsTransportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateStandbyDetailsTransportTypeEnum(val string) (CreateStandbyDetailsTransportTypeEnum, bool) {
	enum, ok := mappingCreateStandbyDetailsTransportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
