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

// NetworkBondingModeDetails Details of bonding mode for Client and Backup and DR networks of an Exadata infrastructure.
type NetworkBondingModeDetails struct {

	// The network bonding mode for the Exadata infrastructure.
	ClientNetworkBondingMode NetworkBondingModeDetailsClientNetworkBondingModeEnum `mandatory:"false" json:"clientNetworkBondingMode,omitempty"`

	// The network bonding mode for the Exadata infrastructure.
	BackupNetworkBondingMode NetworkBondingModeDetailsBackupNetworkBondingModeEnum `mandatory:"false" json:"backupNetworkBondingMode,omitempty"`

	// The network bonding mode for the Exadata infrastructure.
	DrNetworkBondingMode NetworkBondingModeDetailsDrNetworkBondingModeEnum `mandatory:"false" json:"drNetworkBondingMode,omitempty"`
}

func (m NetworkBondingModeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkBondingModeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNetworkBondingModeDetailsClientNetworkBondingModeEnum(string(m.ClientNetworkBondingMode)); !ok && m.ClientNetworkBondingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClientNetworkBondingMode: %s. Supported values are: %s.", m.ClientNetworkBondingMode, strings.Join(GetNetworkBondingModeDetailsClientNetworkBondingModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNetworkBondingModeDetailsBackupNetworkBondingModeEnum(string(m.BackupNetworkBondingMode)); !ok && m.BackupNetworkBondingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupNetworkBondingMode: %s. Supported values are: %s.", m.BackupNetworkBondingMode, strings.Join(GetNetworkBondingModeDetailsBackupNetworkBondingModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNetworkBondingModeDetailsDrNetworkBondingModeEnum(string(m.DrNetworkBondingMode)); !ok && m.DrNetworkBondingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrNetworkBondingMode: %s. Supported values are: %s.", m.DrNetworkBondingMode, strings.Join(GetNetworkBondingModeDetailsDrNetworkBondingModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NetworkBondingModeDetailsClientNetworkBondingModeEnum Enum with underlying type: string
type NetworkBondingModeDetailsClientNetworkBondingModeEnum string

// Set of constants representing the allowable values for NetworkBondingModeDetailsClientNetworkBondingModeEnum
const (
	NetworkBondingModeDetailsClientNetworkBondingModeActiveBackup NetworkBondingModeDetailsClientNetworkBondingModeEnum = "ACTIVE_BACKUP"
	NetworkBondingModeDetailsClientNetworkBondingModeLacp         NetworkBondingModeDetailsClientNetworkBondingModeEnum = "LACP"
)

var mappingNetworkBondingModeDetailsClientNetworkBondingModeEnum = map[string]NetworkBondingModeDetailsClientNetworkBondingModeEnum{
	"ACTIVE_BACKUP": NetworkBondingModeDetailsClientNetworkBondingModeActiveBackup,
	"LACP":          NetworkBondingModeDetailsClientNetworkBondingModeLacp,
}

var mappingNetworkBondingModeDetailsClientNetworkBondingModeEnumLowerCase = map[string]NetworkBondingModeDetailsClientNetworkBondingModeEnum{
	"active_backup": NetworkBondingModeDetailsClientNetworkBondingModeActiveBackup,
	"lacp":          NetworkBondingModeDetailsClientNetworkBondingModeLacp,
}

// GetNetworkBondingModeDetailsClientNetworkBondingModeEnumValues Enumerates the set of values for NetworkBondingModeDetailsClientNetworkBondingModeEnum
func GetNetworkBondingModeDetailsClientNetworkBondingModeEnumValues() []NetworkBondingModeDetailsClientNetworkBondingModeEnum {
	values := make([]NetworkBondingModeDetailsClientNetworkBondingModeEnum, 0)
	for _, v := range mappingNetworkBondingModeDetailsClientNetworkBondingModeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkBondingModeDetailsClientNetworkBondingModeEnumStringValues Enumerates the set of values in String for NetworkBondingModeDetailsClientNetworkBondingModeEnum
func GetNetworkBondingModeDetailsClientNetworkBondingModeEnumStringValues() []string {
	return []string{
		"ACTIVE_BACKUP",
		"LACP",
	}
}

// GetMappingNetworkBondingModeDetailsClientNetworkBondingModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkBondingModeDetailsClientNetworkBondingModeEnum(val string) (NetworkBondingModeDetailsClientNetworkBondingModeEnum, bool) {
	enum, ok := mappingNetworkBondingModeDetailsClientNetworkBondingModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// NetworkBondingModeDetailsBackupNetworkBondingModeEnum Enum with underlying type: string
type NetworkBondingModeDetailsBackupNetworkBondingModeEnum string

// Set of constants representing the allowable values for NetworkBondingModeDetailsBackupNetworkBondingModeEnum
const (
	NetworkBondingModeDetailsBackupNetworkBondingModeActiveBackup NetworkBondingModeDetailsBackupNetworkBondingModeEnum = "ACTIVE_BACKUP"
	NetworkBondingModeDetailsBackupNetworkBondingModeLacp         NetworkBondingModeDetailsBackupNetworkBondingModeEnum = "LACP"
)

var mappingNetworkBondingModeDetailsBackupNetworkBondingModeEnum = map[string]NetworkBondingModeDetailsBackupNetworkBondingModeEnum{
	"ACTIVE_BACKUP": NetworkBondingModeDetailsBackupNetworkBondingModeActiveBackup,
	"LACP":          NetworkBondingModeDetailsBackupNetworkBondingModeLacp,
}

var mappingNetworkBondingModeDetailsBackupNetworkBondingModeEnumLowerCase = map[string]NetworkBondingModeDetailsBackupNetworkBondingModeEnum{
	"active_backup": NetworkBondingModeDetailsBackupNetworkBondingModeActiveBackup,
	"lacp":          NetworkBondingModeDetailsBackupNetworkBondingModeLacp,
}

// GetNetworkBondingModeDetailsBackupNetworkBondingModeEnumValues Enumerates the set of values for NetworkBondingModeDetailsBackupNetworkBondingModeEnum
func GetNetworkBondingModeDetailsBackupNetworkBondingModeEnumValues() []NetworkBondingModeDetailsBackupNetworkBondingModeEnum {
	values := make([]NetworkBondingModeDetailsBackupNetworkBondingModeEnum, 0)
	for _, v := range mappingNetworkBondingModeDetailsBackupNetworkBondingModeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkBondingModeDetailsBackupNetworkBondingModeEnumStringValues Enumerates the set of values in String for NetworkBondingModeDetailsBackupNetworkBondingModeEnum
func GetNetworkBondingModeDetailsBackupNetworkBondingModeEnumStringValues() []string {
	return []string{
		"ACTIVE_BACKUP",
		"LACP",
	}
}

// GetMappingNetworkBondingModeDetailsBackupNetworkBondingModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkBondingModeDetailsBackupNetworkBondingModeEnum(val string) (NetworkBondingModeDetailsBackupNetworkBondingModeEnum, bool) {
	enum, ok := mappingNetworkBondingModeDetailsBackupNetworkBondingModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// NetworkBondingModeDetailsDrNetworkBondingModeEnum Enum with underlying type: string
type NetworkBondingModeDetailsDrNetworkBondingModeEnum string

// Set of constants representing the allowable values for NetworkBondingModeDetailsDrNetworkBondingModeEnum
const (
	NetworkBondingModeDetailsDrNetworkBondingModeActiveBackup NetworkBondingModeDetailsDrNetworkBondingModeEnum = "ACTIVE_BACKUP"
	NetworkBondingModeDetailsDrNetworkBondingModeLacp         NetworkBondingModeDetailsDrNetworkBondingModeEnum = "LACP"
)

var mappingNetworkBondingModeDetailsDrNetworkBondingModeEnum = map[string]NetworkBondingModeDetailsDrNetworkBondingModeEnum{
	"ACTIVE_BACKUP": NetworkBondingModeDetailsDrNetworkBondingModeActiveBackup,
	"LACP":          NetworkBondingModeDetailsDrNetworkBondingModeLacp,
}

var mappingNetworkBondingModeDetailsDrNetworkBondingModeEnumLowerCase = map[string]NetworkBondingModeDetailsDrNetworkBondingModeEnum{
	"active_backup": NetworkBondingModeDetailsDrNetworkBondingModeActiveBackup,
	"lacp":          NetworkBondingModeDetailsDrNetworkBondingModeLacp,
}

// GetNetworkBondingModeDetailsDrNetworkBondingModeEnumValues Enumerates the set of values for NetworkBondingModeDetailsDrNetworkBondingModeEnum
func GetNetworkBondingModeDetailsDrNetworkBondingModeEnumValues() []NetworkBondingModeDetailsDrNetworkBondingModeEnum {
	values := make([]NetworkBondingModeDetailsDrNetworkBondingModeEnum, 0)
	for _, v := range mappingNetworkBondingModeDetailsDrNetworkBondingModeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkBondingModeDetailsDrNetworkBondingModeEnumStringValues Enumerates the set of values in String for NetworkBondingModeDetailsDrNetworkBondingModeEnum
func GetNetworkBondingModeDetailsDrNetworkBondingModeEnumStringValues() []string {
	return []string{
		"ACTIVE_BACKUP",
		"LACP",
	}
}

// GetMappingNetworkBondingModeDetailsDrNetworkBondingModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkBondingModeDetailsDrNetworkBondingModeEnum(val string) (NetworkBondingModeDetailsDrNetworkBondingModeEnum, bool) {
	enum, ok := mappingNetworkBondingModeDetailsDrNetworkBondingModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
