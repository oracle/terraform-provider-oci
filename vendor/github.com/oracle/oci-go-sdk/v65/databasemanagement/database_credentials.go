// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseCredentials The database credentials used to perform management activity.
// Provide one of the following attribute set.
// (userName, password, role) OR (userName, secretId, role) OR (namedCredentialId)
type DatabaseCredentials struct {

	// The database user name used to perform management activity.
	UserName *string `mandatory:"false" json:"userName"`

	// The password for the database user name.
	Password *string `mandatory:"false" json:"password"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
	SecretId *string `mandatory:"false" json:"secretId"`

	// The role of the database user. Indicates whether the database user is a normal user or sysdba.
	Role DatabaseCredentialsRoleEnum `mandatory:"false" json:"role,omitempty"`
}

func (m DatabaseCredentials) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseCredentials) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseCredentialsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDatabaseCredentialsRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseCredentialsRoleEnum Enum with underlying type: string
type DatabaseCredentialsRoleEnum string

// Set of constants representing the allowable values for DatabaseCredentialsRoleEnum
const (
	DatabaseCredentialsRoleNormal DatabaseCredentialsRoleEnum = "NORMAL"
	DatabaseCredentialsRoleSysdba DatabaseCredentialsRoleEnum = "SYSDBA"
)

var mappingDatabaseCredentialsRoleEnum = map[string]DatabaseCredentialsRoleEnum{
	"NORMAL": DatabaseCredentialsRoleNormal,
	"SYSDBA": DatabaseCredentialsRoleSysdba,
}

var mappingDatabaseCredentialsRoleEnumLowerCase = map[string]DatabaseCredentialsRoleEnum{
	"normal": DatabaseCredentialsRoleNormal,
	"sysdba": DatabaseCredentialsRoleSysdba,
}

// GetDatabaseCredentialsRoleEnumValues Enumerates the set of values for DatabaseCredentialsRoleEnum
func GetDatabaseCredentialsRoleEnumValues() []DatabaseCredentialsRoleEnum {
	values := make([]DatabaseCredentialsRoleEnum, 0)
	for _, v := range mappingDatabaseCredentialsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseCredentialsRoleEnumStringValues Enumerates the set of values in String for DatabaseCredentialsRoleEnum
func GetDatabaseCredentialsRoleEnumStringValues() []string {
	return []string{
		"NORMAL",
		"SYSDBA",
	}
}

// GetMappingDatabaseCredentialsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseCredentialsRoleEnum(val string) (DatabaseCredentialsRoleEnum, bool) {
	enum, ok := mappingDatabaseCredentialsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
