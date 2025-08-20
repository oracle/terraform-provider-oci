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

// ConvertStandbyAutonomousContainerDatabaseDetails The configuration details for change Autonomous Container Database Dataguard role
type ConvertStandbyAutonomousContainerDatabaseDetails struct {

	// The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled.
	Role ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum `mandatory:"true" json:"role"`

	// type of connection strings when converting database to snapshot mode
	ConnectionStringsType ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum `mandatory:"false" json:"connectionStringsType,omitempty"`
}

func (m ConvertStandbyAutonomousContainerDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConvertStandbyAutonomousContainerDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetConvertStandbyAutonomousContainerDatabaseDetailsRoleEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum(string(m.ConnectionStringsType)); !ok && m.ConnectionStringsType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionStringsType: %s. Supported values are: %s.", m.ConnectionStringsType, strings.Join(GetConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum Enum with underlying type: string
type ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum string

// Set of constants representing the allowable values for ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum
const (
	ConvertStandbyAutonomousContainerDatabaseDetailsRolePrimary         ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum = "PRIMARY"
	ConvertStandbyAutonomousContainerDatabaseDetailsRoleStandby         ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum = "STANDBY"
	ConvertStandbyAutonomousContainerDatabaseDetailsRoleDisabledStandby ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum = "DISABLED_STANDBY"
	ConvertStandbyAutonomousContainerDatabaseDetailsRoleBackupCopy      ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum = "BACKUP_COPY"
	ConvertStandbyAutonomousContainerDatabaseDetailsRoleSnapshotStandby ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum = "SNAPSHOT_STANDBY"
)

var mappingConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum = map[string]ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum{
	"PRIMARY":          ConvertStandbyAutonomousContainerDatabaseDetailsRolePrimary,
	"STANDBY":          ConvertStandbyAutonomousContainerDatabaseDetailsRoleStandby,
	"DISABLED_STANDBY": ConvertStandbyAutonomousContainerDatabaseDetailsRoleDisabledStandby,
	"BACKUP_COPY":      ConvertStandbyAutonomousContainerDatabaseDetailsRoleBackupCopy,
	"SNAPSHOT_STANDBY": ConvertStandbyAutonomousContainerDatabaseDetailsRoleSnapshotStandby,
}

var mappingConvertStandbyAutonomousContainerDatabaseDetailsRoleEnumLowerCase = map[string]ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum{
	"primary":          ConvertStandbyAutonomousContainerDatabaseDetailsRolePrimary,
	"standby":          ConvertStandbyAutonomousContainerDatabaseDetailsRoleStandby,
	"disabled_standby": ConvertStandbyAutonomousContainerDatabaseDetailsRoleDisabledStandby,
	"backup_copy":      ConvertStandbyAutonomousContainerDatabaseDetailsRoleBackupCopy,
	"snapshot_standby": ConvertStandbyAutonomousContainerDatabaseDetailsRoleSnapshotStandby,
}

// GetConvertStandbyAutonomousContainerDatabaseDetailsRoleEnumValues Enumerates the set of values for ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum
func GetConvertStandbyAutonomousContainerDatabaseDetailsRoleEnumValues() []ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum {
	values := make([]ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum, 0)
	for _, v := range mappingConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetConvertStandbyAutonomousContainerDatabaseDetailsRoleEnumStringValues Enumerates the set of values in String for ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum
func GetConvertStandbyAutonomousContainerDatabaseDetailsRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
		"BACKUP_COPY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum(val string) (ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum, bool) {
	enum, ok := mappingConvertStandbyAutonomousContainerDatabaseDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum Enum with underlying type: string
type ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum string

// Set of constants representing the allowable values for ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum
const (
	ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeSnapshotServices ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum = "SNAPSHOT_SERVICES"
	ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypePrimaryServices  ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum = "PRIMARY_SERVICES"
)

var mappingConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum = map[string]ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum{
	"SNAPSHOT_SERVICES": ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeSnapshotServices,
	"PRIMARY_SERVICES":  ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypePrimaryServices,
}

var mappingConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnumLowerCase = map[string]ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum{
	"snapshot_services": ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeSnapshotServices,
	"primary_services":  ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypePrimaryServices,
}

// GetConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnumValues Enumerates the set of values for ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum
func GetConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnumValues() []ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum {
	values := make([]ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum, 0)
	for _, v := range mappingConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnumStringValues Enumerates the set of values in String for ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum
func GetConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnumStringValues() []string {
	return []string{
		"SNAPSHOT_SERVICES",
		"PRIMARY_SERVICES",
	}
}

// GetMappingConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum(val string) (ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum, bool) {
	enum, ok := mappingConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
