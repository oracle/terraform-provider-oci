// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DatabaseProperties Database related properties.
type DatabaseProperties struct {

	// Name of the database product.
	Product *string `mandatory:"true" json:"product"`

	// Database version.
	Version *string `mandatory:"true" json:"version"`

	// Database Platform name.
	// Example: `Linux x86 64-bit`
	Platform *string `mandatory:"true" json:"platform"`

	// Status of CDB/PDB/Standalone database.
	// Example: `ACTIVE`,`SUSPENDED`,`INSTANCE RECOVERY`
	Status DatabasePropertiesStatusEnum `mandatory:"true" json:"status"`

	// Database TimeZone.
	// Example: `+00:00`,`Africa/Abidjan`
	TimeZone *string `mandatory:"true" json:"timeZone"`

	// Database CharacterSet.
	// Example: `AL32UTF8`
	CharacterSet DatabasePropertiesCharacterSetEnum `mandatory:"true" json:"characterSet"`
}

func (m DatabaseProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabasePropertiesStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDatabasePropertiesStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabasePropertiesCharacterSetEnum(string(m.CharacterSet)); !ok && m.CharacterSet != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CharacterSet: %s. Supported values are: %s.", m.CharacterSet, strings.Join(GetDatabasePropertiesCharacterSetEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabasePropertiesStatusEnum Enum with underlying type: string
type DatabasePropertiesStatusEnum string

// Set of constants representing the allowable values for DatabasePropertiesStatusEnum
const (
	DatabasePropertiesStatusActive           DatabasePropertiesStatusEnum = "ACTIVE"
	DatabasePropertiesStatusSuspended        DatabasePropertiesStatusEnum = "SUSPENDED"
	DatabasePropertiesStatusInstanceRecovery DatabasePropertiesStatusEnum = "INSTANCE_RECOVERY"
	DatabasePropertiesStatusNew              DatabasePropertiesStatusEnum = "NEW"
	DatabasePropertiesStatusNormal           DatabasePropertiesStatusEnum = "NORMAL"
	DatabasePropertiesStatusUnplugged        DatabasePropertiesStatusEnum = "UNPLUGGED"
	DatabasePropertiesStatusRelocating       DatabasePropertiesStatusEnum = "RELOCATING"
	DatabasePropertiesStatusRelocated        DatabasePropertiesStatusEnum = "RELOCATED"
	DatabasePropertiesStatusRefreshing       DatabasePropertiesStatusEnum = "REFRESHING"
	DatabasePropertiesStatusUndefined        DatabasePropertiesStatusEnum = "UNDEFINED"
	DatabasePropertiesStatusUnusable         DatabasePropertiesStatusEnum = "UNUSABLE"
)

var mappingDatabasePropertiesStatusEnum = map[string]DatabasePropertiesStatusEnum{
	"ACTIVE":            DatabasePropertiesStatusActive,
	"SUSPENDED":         DatabasePropertiesStatusSuspended,
	"INSTANCE_RECOVERY": DatabasePropertiesStatusInstanceRecovery,
	"NEW":               DatabasePropertiesStatusNew,
	"NORMAL":            DatabasePropertiesStatusNormal,
	"UNPLUGGED":         DatabasePropertiesStatusUnplugged,
	"RELOCATING":        DatabasePropertiesStatusRelocating,
	"RELOCATED":         DatabasePropertiesStatusRelocated,
	"REFRESHING":        DatabasePropertiesStatusRefreshing,
	"UNDEFINED":         DatabasePropertiesStatusUndefined,
	"UNUSABLE":          DatabasePropertiesStatusUnusable,
}

var mappingDatabasePropertiesStatusEnumLowerCase = map[string]DatabasePropertiesStatusEnum{
	"active":            DatabasePropertiesStatusActive,
	"suspended":         DatabasePropertiesStatusSuspended,
	"instance_recovery": DatabasePropertiesStatusInstanceRecovery,
	"new":               DatabasePropertiesStatusNew,
	"normal":            DatabasePropertiesStatusNormal,
	"unplugged":         DatabasePropertiesStatusUnplugged,
	"relocating":        DatabasePropertiesStatusRelocating,
	"relocated":         DatabasePropertiesStatusRelocated,
	"refreshing":        DatabasePropertiesStatusRefreshing,
	"undefined":         DatabasePropertiesStatusUndefined,
	"unusable":          DatabasePropertiesStatusUnusable,
}

// GetDatabasePropertiesStatusEnumValues Enumerates the set of values for DatabasePropertiesStatusEnum
func GetDatabasePropertiesStatusEnumValues() []DatabasePropertiesStatusEnum {
	values := make([]DatabasePropertiesStatusEnum, 0)
	for _, v := range mappingDatabasePropertiesStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabasePropertiesStatusEnumStringValues Enumerates the set of values in String for DatabasePropertiesStatusEnum
func GetDatabasePropertiesStatusEnumStringValues() []string {
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

// GetMappingDatabasePropertiesStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabasePropertiesStatusEnum(val string) (DatabasePropertiesStatusEnum, bool) {
	enum, ok := mappingDatabasePropertiesStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabasePropertiesCharacterSetEnum Enum with underlying type: string
type DatabasePropertiesCharacterSetEnum string

// Set of constants representing the allowable values for DatabasePropertiesCharacterSetEnum
const (
	DatabasePropertiesCharacterSetAl24utffss DatabasePropertiesCharacterSetEnum = "AL24UTFFSS"
	DatabasePropertiesCharacterSetUtf8       DatabasePropertiesCharacterSetEnum = "UTF8"
	DatabasePropertiesCharacterSetUtfe       DatabasePropertiesCharacterSetEnum = "UTFE"
	DatabasePropertiesCharacterSetAl32utf8   DatabasePropertiesCharacterSetEnum = "AL32UTF8"
	DatabasePropertiesCharacterSetAl16utf16  DatabasePropertiesCharacterSetEnum = "AL16UTF16"
)

var mappingDatabasePropertiesCharacterSetEnum = map[string]DatabasePropertiesCharacterSetEnum{
	"AL24UTFFSS": DatabasePropertiesCharacterSetAl24utffss,
	"UTF8":       DatabasePropertiesCharacterSetUtf8,
	"UTFE":       DatabasePropertiesCharacterSetUtfe,
	"AL32UTF8":   DatabasePropertiesCharacterSetAl32utf8,
	"AL16UTF16":  DatabasePropertiesCharacterSetAl16utf16,
}

var mappingDatabasePropertiesCharacterSetEnumLowerCase = map[string]DatabasePropertiesCharacterSetEnum{
	"al24utffss": DatabasePropertiesCharacterSetAl24utffss,
	"utf8":       DatabasePropertiesCharacterSetUtf8,
	"utfe":       DatabasePropertiesCharacterSetUtfe,
	"al32utf8":   DatabasePropertiesCharacterSetAl32utf8,
	"al16utf16":  DatabasePropertiesCharacterSetAl16utf16,
}

// GetDatabasePropertiesCharacterSetEnumValues Enumerates the set of values for DatabasePropertiesCharacterSetEnum
func GetDatabasePropertiesCharacterSetEnumValues() []DatabasePropertiesCharacterSetEnum {
	values := make([]DatabasePropertiesCharacterSetEnum, 0)
	for _, v := range mappingDatabasePropertiesCharacterSetEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabasePropertiesCharacterSetEnumStringValues Enumerates the set of values in String for DatabasePropertiesCharacterSetEnum
func GetDatabasePropertiesCharacterSetEnumStringValues() []string {
	return []string{
		"AL24UTFFSS",
		"UTF8",
		"UTFE",
		"AL32UTF8",
		"AL16UTF16",
	}
}

// GetMappingDatabasePropertiesCharacterSetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabasePropertiesCharacterSetEnum(val string) (DatabasePropertiesCharacterSetEnum, bool) {
	enum, ok := mappingDatabasePropertiesCharacterSetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
