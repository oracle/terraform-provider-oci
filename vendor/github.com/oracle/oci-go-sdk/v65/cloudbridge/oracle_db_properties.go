// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleDbProperties Oracle database related properties.
type OracleDbProperties struct {

	// Name of the database product.
	Product *string `mandatory:"true" json:"product"`

	// Database version.
	Version *string `mandatory:"true" json:"version"`

	// Database Platform name.
	// Example: `Linux x86 64-bit`
	Platform *string `mandatory:"true" json:"platform"`

	// Status of CDB/PDB/Standalone database.
	// Example: `ACTIVE`,`SUSPENDED`,`INSTANCE RECOVERY`
	Status OracleDbPropertiesStatusEnum `mandatory:"true" json:"status"`

	// Database TimeZone.
	// Example: `+00:00`,`Africa/Abidjan`
	TimeZone *string `mandatory:"true" json:"timeZone"`

	// Database CharacterSet.
	// Example: `AL32UTF8`
	CharacterSet OracleDbPropertiesCharacterSetEnum `mandatory:"true" json:"characterSet"`

	// Database identifier.
	// Example: 1 for root, 2 for PDB template seed and 3 or above for PDB's
	ConnectionId *string `mandatory:"true" json:"connectionId"`

	// Type of Database.
	// Example: `CDB`,`PDB`,`NON_CDB`
	Type OracleDbPropertiesTypeEnum `mandatory:"true" json:"type"`

	// Database service name.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// List of connection details.
	ConnectionDetails []OracleDbConnectionDetails `mandatory:"true" json:"connectionDetails"`

	// List of listeners.
	Listeners []string `mandatory:"true" json:"listeners"`
}

func (m OracleDbProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleDbProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbPropertiesStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOracleDbPropertiesStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOracleDbPropertiesCharacterSetEnum(string(m.CharacterSet)); !ok && m.CharacterSet != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CharacterSet: %s. Supported values are: %s.", m.CharacterSet, strings.Join(GetOracleDbPropertiesCharacterSetEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOracleDbPropertiesTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetOracleDbPropertiesTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OracleDbPropertiesStatusEnum Enum with underlying type: string
type OracleDbPropertiesStatusEnum string

// Set of constants representing the allowable values for OracleDbPropertiesStatusEnum
const (
	OracleDbPropertiesStatusActive           OracleDbPropertiesStatusEnum = "ACTIVE"
	OracleDbPropertiesStatusSuspended        OracleDbPropertiesStatusEnum = "SUSPENDED"
	OracleDbPropertiesStatusInstanceRecovery OracleDbPropertiesStatusEnum = "INSTANCE_RECOVERY"
	OracleDbPropertiesStatusNew              OracleDbPropertiesStatusEnum = "NEW"
	OracleDbPropertiesStatusNormal           OracleDbPropertiesStatusEnum = "NORMAL"
	OracleDbPropertiesStatusUnplugged        OracleDbPropertiesStatusEnum = "UNPLUGGED"
	OracleDbPropertiesStatusRelocating       OracleDbPropertiesStatusEnum = "RELOCATING"
	OracleDbPropertiesStatusRelocated        OracleDbPropertiesStatusEnum = "RELOCATED"
	OracleDbPropertiesStatusRefreshing       OracleDbPropertiesStatusEnum = "REFRESHING"
	OracleDbPropertiesStatusUndefined        OracleDbPropertiesStatusEnum = "UNDEFINED"
	OracleDbPropertiesStatusUnusable         OracleDbPropertiesStatusEnum = "UNUSABLE"
)

var mappingOracleDbPropertiesStatusEnum = map[string]OracleDbPropertiesStatusEnum{
	"ACTIVE":            OracleDbPropertiesStatusActive,
	"SUSPENDED":         OracleDbPropertiesStatusSuspended,
	"INSTANCE_RECOVERY": OracleDbPropertiesStatusInstanceRecovery,
	"NEW":               OracleDbPropertiesStatusNew,
	"NORMAL":            OracleDbPropertiesStatusNormal,
	"UNPLUGGED":         OracleDbPropertiesStatusUnplugged,
	"RELOCATING":        OracleDbPropertiesStatusRelocating,
	"RELOCATED":         OracleDbPropertiesStatusRelocated,
	"REFRESHING":        OracleDbPropertiesStatusRefreshing,
	"UNDEFINED":         OracleDbPropertiesStatusUndefined,
	"UNUSABLE":          OracleDbPropertiesStatusUnusable,
}

var mappingOracleDbPropertiesStatusEnumLowerCase = map[string]OracleDbPropertiesStatusEnum{
	"active":            OracleDbPropertiesStatusActive,
	"suspended":         OracleDbPropertiesStatusSuspended,
	"instance_recovery": OracleDbPropertiesStatusInstanceRecovery,
	"new":               OracleDbPropertiesStatusNew,
	"normal":            OracleDbPropertiesStatusNormal,
	"unplugged":         OracleDbPropertiesStatusUnplugged,
	"relocating":        OracleDbPropertiesStatusRelocating,
	"relocated":         OracleDbPropertiesStatusRelocated,
	"refreshing":        OracleDbPropertiesStatusRefreshing,
	"undefined":         OracleDbPropertiesStatusUndefined,
	"unusable":          OracleDbPropertiesStatusUnusable,
}

// GetOracleDbPropertiesStatusEnumValues Enumerates the set of values for OracleDbPropertiesStatusEnum
func GetOracleDbPropertiesStatusEnumValues() []OracleDbPropertiesStatusEnum {
	values := make([]OracleDbPropertiesStatusEnum, 0)
	for _, v := range mappingOracleDbPropertiesStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDbPropertiesStatusEnumStringValues Enumerates the set of values in String for OracleDbPropertiesStatusEnum
func GetOracleDbPropertiesStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"SUSPENDED",
		"INSTANCE_RECOVERY",
		"NEW",
		"NORMAL",
		"UNPLUGGED",
		"RELOCATING",
		"RELOCATED",
		"REFRESHING",
		"UNDEFINED",
		"UNUSABLE",
	}
}

// GetMappingOracleDbPropertiesStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDbPropertiesStatusEnum(val string) (OracleDbPropertiesStatusEnum, bool) {
	enum, ok := mappingOracleDbPropertiesStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OracleDbPropertiesCharacterSetEnum Enum with underlying type: string
type OracleDbPropertiesCharacterSetEnum string

// Set of constants representing the allowable values for OracleDbPropertiesCharacterSetEnum
const (
	OracleDbPropertiesCharacterSetAl24utffss OracleDbPropertiesCharacterSetEnum = "AL24UTFFSS"
	OracleDbPropertiesCharacterSetUtf8       OracleDbPropertiesCharacterSetEnum = "UTF8"
	OracleDbPropertiesCharacterSetUtfe       OracleDbPropertiesCharacterSetEnum = "UTFE"
	OracleDbPropertiesCharacterSetAl32utf8   OracleDbPropertiesCharacterSetEnum = "AL32UTF8"
	OracleDbPropertiesCharacterSetAl16utf16  OracleDbPropertiesCharacterSetEnum = "AL16UTF16"
)

var mappingOracleDbPropertiesCharacterSetEnum = map[string]OracleDbPropertiesCharacterSetEnum{
	"AL24UTFFSS": OracleDbPropertiesCharacterSetAl24utffss,
	"UTF8":       OracleDbPropertiesCharacterSetUtf8,
	"UTFE":       OracleDbPropertiesCharacterSetUtfe,
	"AL32UTF8":   OracleDbPropertiesCharacterSetAl32utf8,
	"AL16UTF16":  OracleDbPropertiesCharacterSetAl16utf16,
}

var mappingOracleDbPropertiesCharacterSetEnumLowerCase = map[string]OracleDbPropertiesCharacterSetEnum{
	"al24utffss": OracleDbPropertiesCharacterSetAl24utffss,
	"utf8":       OracleDbPropertiesCharacterSetUtf8,
	"utfe":       OracleDbPropertiesCharacterSetUtfe,
	"al32utf8":   OracleDbPropertiesCharacterSetAl32utf8,
	"al16utf16":  OracleDbPropertiesCharacterSetAl16utf16,
}

// GetOracleDbPropertiesCharacterSetEnumValues Enumerates the set of values for OracleDbPropertiesCharacterSetEnum
func GetOracleDbPropertiesCharacterSetEnumValues() []OracleDbPropertiesCharacterSetEnum {
	values := make([]OracleDbPropertiesCharacterSetEnum, 0)
	for _, v := range mappingOracleDbPropertiesCharacterSetEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDbPropertiesCharacterSetEnumStringValues Enumerates the set of values in String for OracleDbPropertiesCharacterSetEnum
func GetOracleDbPropertiesCharacterSetEnumStringValues() []string {
	return []string{
		"AL24UTFFSS",
		"UTF8",
		"UTFE",
		"AL32UTF8",
		"AL16UTF16",
	}
}

// GetMappingOracleDbPropertiesCharacterSetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDbPropertiesCharacterSetEnum(val string) (OracleDbPropertiesCharacterSetEnum, bool) {
	enum, ok := mappingOracleDbPropertiesCharacterSetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OracleDbPropertiesTypeEnum Enum with underlying type: string
type OracleDbPropertiesTypeEnum string

// Set of constants representing the allowable values for OracleDbPropertiesTypeEnum
const (
	OracleDbPropertiesTypeCdb    OracleDbPropertiesTypeEnum = "CDB"
	OracleDbPropertiesTypeNonCdb OracleDbPropertiesTypeEnum = "NON_CDB"
	OracleDbPropertiesTypePdb    OracleDbPropertiesTypeEnum = "PDB"
)

var mappingOracleDbPropertiesTypeEnum = map[string]OracleDbPropertiesTypeEnum{
	"CDB":     OracleDbPropertiesTypeCdb,
	"NON_CDB": OracleDbPropertiesTypeNonCdb,
	"PDB":     OracleDbPropertiesTypePdb,
}

var mappingOracleDbPropertiesTypeEnumLowerCase = map[string]OracleDbPropertiesTypeEnum{
	"cdb":     OracleDbPropertiesTypeCdb,
	"non_cdb": OracleDbPropertiesTypeNonCdb,
	"pdb":     OracleDbPropertiesTypePdb,
}

// GetOracleDbPropertiesTypeEnumValues Enumerates the set of values for OracleDbPropertiesTypeEnum
func GetOracleDbPropertiesTypeEnumValues() []OracleDbPropertiesTypeEnum {
	values := make([]OracleDbPropertiesTypeEnum, 0)
	for _, v := range mappingOracleDbPropertiesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDbPropertiesTypeEnumStringValues Enumerates the set of values in String for OracleDbPropertiesTypeEnum
func GetOracleDbPropertiesTypeEnumStringValues() []string {
	return []string{
		"CDB",
		"NON_CDB",
		"PDB",
	}
}

// GetMappingOracleDbPropertiesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDbPropertiesTypeEnum(val string) (OracleDbPropertiesTypeEnum, bool) {
	enum, ok := mappingOracleDbPropertiesTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
