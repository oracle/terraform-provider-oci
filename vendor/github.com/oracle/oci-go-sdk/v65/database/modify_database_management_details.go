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

// ModifyDatabaseManagementDetails Data to update one or more attributes of the Database Management configuration for the database.
type ModifyDatabaseManagementDetails struct {
	CredentialDetails *DatabaseCredentialDetails `mandatory:"false" json:"credentialDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private endpoint.
	PrivateEndPointId *string `mandatory:"false" json:"privateEndPointId"`

	// The Database Management type.
	ManagementType ModifyDatabaseManagementDetailsManagementTypeEnum `mandatory:"false" json:"managementType,omitempty"`

	// The name of the Oracle Database service that will be used to connect to the database.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// Protocol used by the database connection.
	Protocol ModifyDatabaseManagementDetailsProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	// The port used to connect to the database.
	Port *int `mandatory:"false" json:"port"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure secret (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	SslSecretId *string `mandatory:"false" json:"sslSecretId"`

	// The role of the user that will be connecting to the database.
	Role ModifyDatabaseManagementDetailsRoleEnum `mandatory:"false" json:"role,omitempty"`
}

func (m ModifyDatabaseManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModifyDatabaseManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingModifyDatabaseManagementDetailsManagementTypeEnum(string(m.ManagementType)); !ok && m.ManagementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementType: %s. Supported values are: %s.", m.ManagementType, strings.Join(GetModifyDatabaseManagementDetailsManagementTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingModifyDatabaseManagementDetailsProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetModifyDatabaseManagementDetailsProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingModifyDatabaseManagementDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetModifyDatabaseManagementDetailsRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModifyDatabaseManagementDetailsManagementTypeEnum Enum with underlying type: string
type ModifyDatabaseManagementDetailsManagementTypeEnum string

// Set of constants representing the allowable values for ModifyDatabaseManagementDetailsManagementTypeEnum
const (
	ModifyDatabaseManagementDetailsManagementTypeBasic    ModifyDatabaseManagementDetailsManagementTypeEnum = "BASIC"
	ModifyDatabaseManagementDetailsManagementTypeAdvanced ModifyDatabaseManagementDetailsManagementTypeEnum = "ADVANCED"
)

var mappingModifyDatabaseManagementDetailsManagementTypeEnum = map[string]ModifyDatabaseManagementDetailsManagementTypeEnum{
	"BASIC":    ModifyDatabaseManagementDetailsManagementTypeBasic,
	"ADVANCED": ModifyDatabaseManagementDetailsManagementTypeAdvanced,
}

var mappingModifyDatabaseManagementDetailsManagementTypeEnumLowerCase = map[string]ModifyDatabaseManagementDetailsManagementTypeEnum{
	"basic":    ModifyDatabaseManagementDetailsManagementTypeBasic,
	"advanced": ModifyDatabaseManagementDetailsManagementTypeAdvanced,
}

// GetModifyDatabaseManagementDetailsManagementTypeEnumValues Enumerates the set of values for ModifyDatabaseManagementDetailsManagementTypeEnum
func GetModifyDatabaseManagementDetailsManagementTypeEnumValues() []ModifyDatabaseManagementDetailsManagementTypeEnum {
	values := make([]ModifyDatabaseManagementDetailsManagementTypeEnum, 0)
	for _, v := range mappingModifyDatabaseManagementDetailsManagementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModifyDatabaseManagementDetailsManagementTypeEnumStringValues Enumerates the set of values in String for ModifyDatabaseManagementDetailsManagementTypeEnum
func GetModifyDatabaseManagementDetailsManagementTypeEnumStringValues() []string {
	return []string{
		"BASIC",
		"ADVANCED",
	}
}

// GetMappingModifyDatabaseManagementDetailsManagementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModifyDatabaseManagementDetailsManagementTypeEnum(val string) (ModifyDatabaseManagementDetailsManagementTypeEnum, bool) {
	enum, ok := mappingModifyDatabaseManagementDetailsManagementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ModifyDatabaseManagementDetailsProtocolEnum Enum with underlying type: string
type ModifyDatabaseManagementDetailsProtocolEnum string

// Set of constants representing the allowable values for ModifyDatabaseManagementDetailsProtocolEnum
const (
	ModifyDatabaseManagementDetailsProtocolTcp  ModifyDatabaseManagementDetailsProtocolEnum = "TCP"
	ModifyDatabaseManagementDetailsProtocolTcps ModifyDatabaseManagementDetailsProtocolEnum = "TCPS"
)

var mappingModifyDatabaseManagementDetailsProtocolEnum = map[string]ModifyDatabaseManagementDetailsProtocolEnum{
	"TCP":  ModifyDatabaseManagementDetailsProtocolTcp,
	"TCPS": ModifyDatabaseManagementDetailsProtocolTcps,
}

var mappingModifyDatabaseManagementDetailsProtocolEnumLowerCase = map[string]ModifyDatabaseManagementDetailsProtocolEnum{
	"tcp":  ModifyDatabaseManagementDetailsProtocolTcp,
	"tcps": ModifyDatabaseManagementDetailsProtocolTcps,
}

// GetModifyDatabaseManagementDetailsProtocolEnumValues Enumerates the set of values for ModifyDatabaseManagementDetailsProtocolEnum
func GetModifyDatabaseManagementDetailsProtocolEnumValues() []ModifyDatabaseManagementDetailsProtocolEnum {
	values := make([]ModifyDatabaseManagementDetailsProtocolEnum, 0)
	for _, v := range mappingModifyDatabaseManagementDetailsProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetModifyDatabaseManagementDetailsProtocolEnumStringValues Enumerates the set of values in String for ModifyDatabaseManagementDetailsProtocolEnum
func GetModifyDatabaseManagementDetailsProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
	}
}

// GetMappingModifyDatabaseManagementDetailsProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModifyDatabaseManagementDetailsProtocolEnum(val string) (ModifyDatabaseManagementDetailsProtocolEnum, bool) {
	enum, ok := mappingModifyDatabaseManagementDetailsProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ModifyDatabaseManagementDetailsRoleEnum Enum with underlying type: string
type ModifyDatabaseManagementDetailsRoleEnum string

// Set of constants representing the allowable values for ModifyDatabaseManagementDetailsRoleEnum
const (
	ModifyDatabaseManagementDetailsRoleSysdba ModifyDatabaseManagementDetailsRoleEnum = "SYSDBA"
	ModifyDatabaseManagementDetailsRoleNormal ModifyDatabaseManagementDetailsRoleEnum = "NORMAL"
)

var mappingModifyDatabaseManagementDetailsRoleEnum = map[string]ModifyDatabaseManagementDetailsRoleEnum{
	"SYSDBA": ModifyDatabaseManagementDetailsRoleSysdba,
	"NORMAL": ModifyDatabaseManagementDetailsRoleNormal,
}

var mappingModifyDatabaseManagementDetailsRoleEnumLowerCase = map[string]ModifyDatabaseManagementDetailsRoleEnum{
	"sysdba": ModifyDatabaseManagementDetailsRoleSysdba,
	"normal": ModifyDatabaseManagementDetailsRoleNormal,
}

// GetModifyDatabaseManagementDetailsRoleEnumValues Enumerates the set of values for ModifyDatabaseManagementDetailsRoleEnum
func GetModifyDatabaseManagementDetailsRoleEnumValues() []ModifyDatabaseManagementDetailsRoleEnum {
	values := make([]ModifyDatabaseManagementDetailsRoleEnum, 0)
	for _, v := range mappingModifyDatabaseManagementDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetModifyDatabaseManagementDetailsRoleEnumStringValues Enumerates the set of values in String for ModifyDatabaseManagementDetailsRoleEnum
func GetModifyDatabaseManagementDetailsRoleEnumStringValues() []string {
	return []string{
		"SYSDBA",
		"NORMAL",
	}
}

// GetMappingModifyDatabaseManagementDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModifyDatabaseManagementDetailsRoleEnum(val string) (ModifyDatabaseManagementDetailsRoleEnum, bool) {
	enum, ok := mappingModifyDatabaseManagementDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
