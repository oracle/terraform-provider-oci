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

// ModifyPluggableDatabaseManagementDetails Data to update one or more attributes of the Database Management configuration for the pluggable database.
type ModifyPluggableDatabaseManagementDetails struct {
	CredentialDetails *DatabaseCredentialDetails `mandatory:"false" json:"credentialDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private endpoint.
	PrivateEndPointId *string `mandatory:"false" json:"privateEndPointId"`

	// The name of the Oracle Database service that will be used to connect to the database.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// Protocol used by the database connection.
	Protocol ModifyPluggableDatabaseManagementDetailsProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	// The port used to connect to the database.
	Port *int `mandatory:"false" json:"port"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure secret (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	SslSecretId *string `mandatory:"false" json:"sslSecretId"`

	// The role of the user that will be connecting to the database.
	Role ModifyPluggableDatabaseManagementDetailsRoleEnum `mandatory:"false" json:"role,omitempty"`
}

func (m ModifyPluggableDatabaseManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModifyPluggableDatabaseManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingModifyPluggableDatabaseManagementDetailsProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetModifyPluggableDatabaseManagementDetailsProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingModifyPluggableDatabaseManagementDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetModifyPluggableDatabaseManagementDetailsRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModifyPluggableDatabaseManagementDetailsProtocolEnum Enum with underlying type: string
type ModifyPluggableDatabaseManagementDetailsProtocolEnum string

// Set of constants representing the allowable values for ModifyPluggableDatabaseManagementDetailsProtocolEnum
const (
	ModifyPluggableDatabaseManagementDetailsProtocolTcp  ModifyPluggableDatabaseManagementDetailsProtocolEnum = "TCP"
	ModifyPluggableDatabaseManagementDetailsProtocolTcps ModifyPluggableDatabaseManagementDetailsProtocolEnum = "TCPS"
)

var mappingModifyPluggableDatabaseManagementDetailsProtocolEnum = map[string]ModifyPluggableDatabaseManagementDetailsProtocolEnum{
	"TCP":  ModifyPluggableDatabaseManagementDetailsProtocolTcp,
	"TCPS": ModifyPluggableDatabaseManagementDetailsProtocolTcps,
}

var mappingModifyPluggableDatabaseManagementDetailsProtocolEnumLowerCase = map[string]ModifyPluggableDatabaseManagementDetailsProtocolEnum{
	"tcp":  ModifyPluggableDatabaseManagementDetailsProtocolTcp,
	"tcps": ModifyPluggableDatabaseManagementDetailsProtocolTcps,
}

// GetModifyPluggableDatabaseManagementDetailsProtocolEnumValues Enumerates the set of values for ModifyPluggableDatabaseManagementDetailsProtocolEnum
func GetModifyPluggableDatabaseManagementDetailsProtocolEnumValues() []ModifyPluggableDatabaseManagementDetailsProtocolEnum {
	values := make([]ModifyPluggableDatabaseManagementDetailsProtocolEnum, 0)
	for _, v := range mappingModifyPluggableDatabaseManagementDetailsProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetModifyPluggableDatabaseManagementDetailsProtocolEnumStringValues Enumerates the set of values in String for ModifyPluggableDatabaseManagementDetailsProtocolEnum
func GetModifyPluggableDatabaseManagementDetailsProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
	}
}

// GetMappingModifyPluggableDatabaseManagementDetailsProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModifyPluggableDatabaseManagementDetailsProtocolEnum(val string) (ModifyPluggableDatabaseManagementDetailsProtocolEnum, bool) {
	enum, ok := mappingModifyPluggableDatabaseManagementDetailsProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ModifyPluggableDatabaseManagementDetailsRoleEnum Enum with underlying type: string
type ModifyPluggableDatabaseManagementDetailsRoleEnum string

// Set of constants representing the allowable values for ModifyPluggableDatabaseManagementDetailsRoleEnum
const (
	ModifyPluggableDatabaseManagementDetailsRoleSysdba ModifyPluggableDatabaseManagementDetailsRoleEnum = "SYSDBA"
	ModifyPluggableDatabaseManagementDetailsRoleNormal ModifyPluggableDatabaseManagementDetailsRoleEnum = "NORMAL"
)

var mappingModifyPluggableDatabaseManagementDetailsRoleEnum = map[string]ModifyPluggableDatabaseManagementDetailsRoleEnum{
	"SYSDBA": ModifyPluggableDatabaseManagementDetailsRoleSysdba,
	"NORMAL": ModifyPluggableDatabaseManagementDetailsRoleNormal,
}

var mappingModifyPluggableDatabaseManagementDetailsRoleEnumLowerCase = map[string]ModifyPluggableDatabaseManagementDetailsRoleEnum{
	"sysdba": ModifyPluggableDatabaseManagementDetailsRoleSysdba,
	"normal": ModifyPluggableDatabaseManagementDetailsRoleNormal,
}

// GetModifyPluggableDatabaseManagementDetailsRoleEnumValues Enumerates the set of values for ModifyPluggableDatabaseManagementDetailsRoleEnum
func GetModifyPluggableDatabaseManagementDetailsRoleEnumValues() []ModifyPluggableDatabaseManagementDetailsRoleEnum {
	values := make([]ModifyPluggableDatabaseManagementDetailsRoleEnum, 0)
	for _, v := range mappingModifyPluggableDatabaseManagementDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetModifyPluggableDatabaseManagementDetailsRoleEnumStringValues Enumerates the set of values in String for ModifyPluggableDatabaseManagementDetailsRoleEnum
func GetModifyPluggableDatabaseManagementDetailsRoleEnumStringValues() []string {
	return []string{
		"SYSDBA",
		"NORMAL",
	}
}

// GetMappingModifyPluggableDatabaseManagementDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModifyPluggableDatabaseManagementDetailsRoleEnum(val string) (ModifyPluggableDatabaseManagementDetailsRoleEnum, bool) {
	enum, ok := mappingModifyPluggableDatabaseManagementDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
