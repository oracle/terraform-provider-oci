// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// EnablePluggableDatabaseManagementDetails Data to enable the Database Management service for the pluggable database.
type EnablePluggableDatabaseManagementDetails struct {
	CredentialDetails *DatabaseCredentialDetails `mandatory:"true" json:"credentialDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private endpoint.
	PrivateEndPointId *string `mandatory:"true" json:"privateEndPointId"`

	// The name of the Oracle Database service that will be used to connect to the database.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// Protocol used by the database connection.
	Protocol EnablePluggableDatabaseManagementDetailsProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	// The port used to connect to the pluggable database.
	Port *int `mandatory:"false" json:"port"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure secret (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	SslSecretId *string `mandatory:"false" json:"sslSecretId"`

	// The role of the user that will be connecting to the pluggable database.
	Role EnablePluggableDatabaseManagementDetailsRoleEnum `mandatory:"false" json:"role,omitempty"`
}

func (m EnablePluggableDatabaseManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnablePluggableDatabaseManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEnablePluggableDatabaseManagementDetailsProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetEnablePluggableDatabaseManagementDetailsProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEnablePluggableDatabaseManagementDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetEnablePluggableDatabaseManagementDetailsRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EnablePluggableDatabaseManagementDetailsProtocolEnum Enum with underlying type: string
type EnablePluggableDatabaseManagementDetailsProtocolEnum string

// Set of constants representing the allowable values for EnablePluggableDatabaseManagementDetailsProtocolEnum
const (
	EnablePluggableDatabaseManagementDetailsProtocolTcp  EnablePluggableDatabaseManagementDetailsProtocolEnum = "TCP"
	EnablePluggableDatabaseManagementDetailsProtocolTcps EnablePluggableDatabaseManagementDetailsProtocolEnum = "TCPS"
)

var mappingEnablePluggableDatabaseManagementDetailsProtocolEnum = map[string]EnablePluggableDatabaseManagementDetailsProtocolEnum{
	"TCP":  EnablePluggableDatabaseManagementDetailsProtocolTcp,
	"TCPS": EnablePluggableDatabaseManagementDetailsProtocolTcps,
}

var mappingEnablePluggableDatabaseManagementDetailsProtocolEnumLowerCase = map[string]EnablePluggableDatabaseManagementDetailsProtocolEnum{
	"tcp":  EnablePluggableDatabaseManagementDetailsProtocolTcp,
	"tcps": EnablePluggableDatabaseManagementDetailsProtocolTcps,
}

// GetEnablePluggableDatabaseManagementDetailsProtocolEnumValues Enumerates the set of values for EnablePluggableDatabaseManagementDetailsProtocolEnum
func GetEnablePluggableDatabaseManagementDetailsProtocolEnumValues() []EnablePluggableDatabaseManagementDetailsProtocolEnum {
	values := make([]EnablePluggableDatabaseManagementDetailsProtocolEnum, 0)
	for _, v := range mappingEnablePluggableDatabaseManagementDetailsProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetEnablePluggableDatabaseManagementDetailsProtocolEnumStringValues Enumerates the set of values in String for EnablePluggableDatabaseManagementDetailsProtocolEnum
func GetEnablePluggableDatabaseManagementDetailsProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
	}
}

// GetMappingEnablePluggableDatabaseManagementDetailsProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnablePluggableDatabaseManagementDetailsProtocolEnum(val string) (EnablePluggableDatabaseManagementDetailsProtocolEnum, bool) {
	enum, ok := mappingEnablePluggableDatabaseManagementDetailsProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EnablePluggableDatabaseManagementDetailsRoleEnum Enum with underlying type: string
type EnablePluggableDatabaseManagementDetailsRoleEnum string

// Set of constants representing the allowable values for EnablePluggableDatabaseManagementDetailsRoleEnum
const (
	EnablePluggableDatabaseManagementDetailsRoleSysdba EnablePluggableDatabaseManagementDetailsRoleEnum = "SYSDBA"
	EnablePluggableDatabaseManagementDetailsRoleNormal EnablePluggableDatabaseManagementDetailsRoleEnum = "NORMAL"
)

var mappingEnablePluggableDatabaseManagementDetailsRoleEnum = map[string]EnablePluggableDatabaseManagementDetailsRoleEnum{
	"SYSDBA": EnablePluggableDatabaseManagementDetailsRoleSysdba,
	"NORMAL": EnablePluggableDatabaseManagementDetailsRoleNormal,
}

var mappingEnablePluggableDatabaseManagementDetailsRoleEnumLowerCase = map[string]EnablePluggableDatabaseManagementDetailsRoleEnum{
	"sysdba": EnablePluggableDatabaseManagementDetailsRoleSysdba,
	"normal": EnablePluggableDatabaseManagementDetailsRoleNormal,
}

// GetEnablePluggableDatabaseManagementDetailsRoleEnumValues Enumerates the set of values for EnablePluggableDatabaseManagementDetailsRoleEnum
func GetEnablePluggableDatabaseManagementDetailsRoleEnumValues() []EnablePluggableDatabaseManagementDetailsRoleEnum {
	values := make([]EnablePluggableDatabaseManagementDetailsRoleEnum, 0)
	for _, v := range mappingEnablePluggableDatabaseManagementDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetEnablePluggableDatabaseManagementDetailsRoleEnumStringValues Enumerates the set of values in String for EnablePluggableDatabaseManagementDetailsRoleEnum
func GetEnablePluggableDatabaseManagementDetailsRoleEnumStringValues() []string {
	return []string{
		"SYSDBA",
		"NORMAL",
	}
}

// GetMappingEnablePluggableDatabaseManagementDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnablePluggableDatabaseManagementDetailsRoleEnum(val string) (EnablePluggableDatabaseManagementDetailsRoleEnum, bool) {
	enum, ok := mappingEnablePluggableDatabaseManagementDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
