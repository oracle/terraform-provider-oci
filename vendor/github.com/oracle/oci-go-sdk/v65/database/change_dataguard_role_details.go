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

// ChangeDataguardRoleDetails The configuration details for change Autonomous Container Database Dataguard role
type ChangeDataguardRoleDetails struct {

	// The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled.
	Role ChangeDataguardRoleDetailsRoleEnum `mandatory:"true" json:"role"`

	// The Autonomous Container Database-Autonomous Data Guard association OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseDataguardAssociationId *string `mandatory:"true" json:"autonomousContainerDatabaseDataguardAssociationId"`

	// type of connection strings when converting database to snapshot mode
	ConnectionStringsType ChangeDataguardRoleDetailsConnectionStringsTypeEnum `mandatory:"false" json:"connectionStringsType,omitempty"`
}

func (m ChangeDataguardRoleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeDataguardRoleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingChangeDataguardRoleDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetChangeDataguardRoleDetailsRoleEnumStringValues(), ",")))
	}

	if _, ok := GetMappingChangeDataguardRoleDetailsConnectionStringsTypeEnum(string(m.ConnectionStringsType)); !ok && m.ConnectionStringsType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionStringsType: %s. Supported values are: %s.", m.ConnectionStringsType, strings.Join(GetChangeDataguardRoleDetailsConnectionStringsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChangeDataguardRoleDetailsRoleEnum Enum with underlying type: string
type ChangeDataguardRoleDetailsRoleEnum string

// Set of constants representing the allowable values for ChangeDataguardRoleDetailsRoleEnum
const (
	ChangeDataguardRoleDetailsRolePrimary         ChangeDataguardRoleDetailsRoleEnum = "PRIMARY"
	ChangeDataguardRoleDetailsRoleStandby         ChangeDataguardRoleDetailsRoleEnum = "STANDBY"
	ChangeDataguardRoleDetailsRoleDisabledStandby ChangeDataguardRoleDetailsRoleEnum = "DISABLED_STANDBY"
	ChangeDataguardRoleDetailsRoleBackupCopy      ChangeDataguardRoleDetailsRoleEnum = "BACKUP_COPY"
	ChangeDataguardRoleDetailsRoleSnapshotStandby ChangeDataguardRoleDetailsRoleEnum = "SNAPSHOT_STANDBY"
)

var mappingChangeDataguardRoleDetailsRoleEnum = map[string]ChangeDataguardRoleDetailsRoleEnum{
	"PRIMARY":          ChangeDataguardRoleDetailsRolePrimary,
	"STANDBY":          ChangeDataguardRoleDetailsRoleStandby,
	"DISABLED_STANDBY": ChangeDataguardRoleDetailsRoleDisabledStandby,
	"BACKUP_COPY":      ChangeDataguardRoleDetailsRoleBackupCopy,
	"SNAPSHOT_STANDBY": ChangeDataguardRoleDetailsRoleSnapshotStandby,
}

var mappingChangeDataguardRoleDetailsRoleEnumLowerCase = map[string]ChangeDataguardRoleDetailsRoleEnum{
	"primary":          ChangeDataguardRoleDetailsRolePrimary,
	"standby":          ChangeDataguardRoleDetailsRoleStandby,
	"disabled_standby": ChangeDataguardRoleDetailsRoleDisabledStandby,
	"backup_copy":      ChangeDataguardRoleDetailsRoleBackupCopy,
	"snapshot_standby": ChangeDataguardRoleDetailsRoleSnapshotStandby,
}

// GetChangeDataguardRoleDetailsRoleEnumValues Enumerates the set of values for ChangeDataguardRoleDetailsRoleEnum
func GetChangeDataguardRoleDetailsRoleEnumValues() []ChangeDataguardRoleDetailsRoleEnum {
	values := make([]ChangeDataguardRoleDetailsRoleEnum, 0)
	for _, v := range mappingChangeDataguardRoleDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetChangeDataguardRoleDetailsRoleEnumStringValues Enumerates the set of values in String for ChangeDataguardRoleDetailsRoleEnum
func GetChangeDataguardRoleDetailsRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
		"BACKUP_COPY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingChangeDataguardRoleDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChangeDataguardRoleDetailsRoleEnum(val string) (ChangeDataguardRoleDetailsRoleEnum, bool) {
	enum, ok := mappingChangeDataguardRoleDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ChangeDataguardRoleDetailsConnectionStringsTypeEnum Enum with underlying type: string
type ChangeDataguardRoleDetailsConnectionStringsTypeEnum string

// Set of constants representing the allowable values for ChangeDataguardRoleDetailsConnectionStringsTypeEnum
const (
	ChangeDataguardRoleDetailsConnectionStringsTypeSnapshotServices ChangeDataguardRoleDetailsConnectionStringsTypeEnum = "SNAPSHOT_SERVICES"
	ChangeDataguardRoleDetailsConnectionStringsTypePrimaryServices  ChangeDataguardRoleDetailsConnectionStringsTypeEnum = "PRIMARY_SERVICES"
)

var mappingChangeDataguardRoleDetailsConnectionStringsTypeEnum = map[string]ChangeDataguardRoleDetailsConnectionStringsTypeEnum{
	"SNAPSHOT_SERVICES": ChangeDataguardRoleDetailsConnectionStringsTypeSnapshotServices,
	"PRIMARY_SERVICES":  ChangeDataguardRoleDetailsConnectionStringsTypePrimaryServices,
}

var mappingChangeDataguardRoleDetailsConnectionStringsTypeEnumLowerCase = map[string]ChangeDataguardRoleDetailsConnectionStringsTypeEnum{
	"snapshot_services": ChangeDataguardRoleDetailsConnectionStringsTypeSnapshotServices,
	"primary_services":  ChangeDataguardRoleDetailsConnectionStringsTypePrimaryServices,
}

// GetChangeDataguardRoleDetailsConnectionStringsTypeEnumValues Enumerates the set of values for ChangeDataguardRoleDetailsConnectionStringsTypeEnum
func GetChangeDataguardRoleDetailsConnectionStringsTypeEnumValues() []ChangeDataguardRoleDetailsConnectionStringsTypeEnum {
	values := make([]ChangeDataguardRoleDetailsConnectionStringsTypeEnum, 0)
	for _, v := range mappingChangeDataguardRoleDetailsConnectionStringsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetChangeDataguardRoleDetailsConnectionStringsTypeEnumStringValues Enumerates the set of values in String for ChangeDataguardRoleDetailsConnectionStringsTypeEnum
func GetChangeDataguardRoleDetailsConnectionStringsTypeEnumStringValues() []string {
	return []string{
		"SNAPSHOT_SERVICES",
		"PRIMARY_SERVICES",
	}
}

// GetMappingChangeDataguardRoleDetailsConnectionStringsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChangeDataguardRoleDetailsConnectionStringsTypeEnum(val string) (ChangeDataguardRoleDetailsConnectionStringsTypeEnum, bool) {
	enum, ok := mappingChangeDataguardRoleDetailsConnectionStringsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
