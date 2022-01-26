// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseCredentials The database credentials used to perform management activity.
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

// DatabaseCredentialsRoleEnum Enum with underlying type: string
type DatabaseCredentialsRoleEnum string

// Set of constants representing the allowable values for DatabaseCredentialsRoleEnum
const (
	DatabaseCredentialsRoleNormal DatabaseCredentialsRoleEnum = "NORMAL"
	DatabaseCredentialsRoleSysdba DatabaseCredentialsRoleEnum = "SYSDBA"
)

var mappingDatabaseCredentialsRole = map[string]DatabaseCredentialsRoleEnum{
	"NORMAL": DatabaseCredentialsRoleNormal,
	"SYSDBA": DatabaseCredentialsRoleSysdba,
}

// GetDatabaseCredentialsRoleEnumValues Enumerates the set of values for DatabaseCredentialsRoleEnum
func GetDatabaseCredentialsRoleEnumValues() []DatabaseCredentialsRoleEnum {
	values := make([]DatabaseCredentialsRoleEnum, 0)
	for _, v := range mappingDatabaseCredentialsRole {
		values = append(values, v)
	}
	return values
}
