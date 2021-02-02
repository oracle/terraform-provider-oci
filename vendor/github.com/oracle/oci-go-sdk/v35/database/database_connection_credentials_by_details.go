// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v35/common"
)

// DatabaseConnectionCredentialsByDetails User information to connect to the database.
type DatabaseConnectionCredentialsByDetails struct {

	// The username that will be used to connect to the database.
	Username *string `mandatory:"true" json:"username"`

	// The password that will be used to connect to the database.
	Password *string `mandatory:"true" json:"password"`

	// The name of the credential information that used to connect to the database.
	CredentialName *string `mandatory:"false" json:"credentialName"`

	// The role of the user that will be connecting to the database.
	Role DatabaseConnectionCredentialsByDetailsRoleEnum `mandatory:"true" json:"role"`
}

func (m DatabaseConnectionCredentialsByDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DatabaseConnectionCredentialsByDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseConnectionCredentialsByDetails DatabaseConnectionCredentialsByDetails
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeDatabaseConnectionCredentialsByDetails
	}{
		"DETAILS",
		(MarshalTypeDatabaseConnectionCredentialsByDetails)(m),
	}

	return json.Marshal(&s)
}

// DatabaseConnectionCredentialsByDetailsRoleEnum Enum with underlying type: string
type DatabaseConnectionCredentialsByDetailsRoleEnum string

// Set of constants representing the allowable values for DatabaseConnectionCredentialsByDetailsRoleEnum
const (
	DatabaseConnectionCredentialsByDetailsRoleSysdba DatabaseConnectionCredentialsByDetailsRoleEnum = "SYSDBA"
	DatabaseConnectionCredentialsByDetailsRoleNormal DatabaseConnectionCredentialsByDetailsRoleEnum = "NORMAL"
)

var mappingDatabaseConnectionCredentialsByDetailsRole = map[string]DatabaseConnectionCredentialsByDetailsRoleEnum{
	"SYSDBA": DatabaseConnectionCredentialsByDetailsRoleSysdba,
	"NORMAL": DatabaseConnectionCredentialsByDetailsRoleNormal,
}

// GetDatabaseConnectionCredentialsByDetailsRoleEnumValues Enumerates the set of values for DatabaseConnectionCredentialsByDetailsRoleEnum
func GetDatabaseConnectionCredentialsByDetailsRoleEnumValues() []DatabaseConnectionCredentialsByDetailsRoleEnum {
	values := make([]DatabaseConnectionCredentialsByDetailsRoleEnum, 0)
	for _, v := range mappingDatabaseConnectionCredentialsByDetailsRole {
		values = append(values, v)
	}
	return values
}
