// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabasePasswordCredentialDetails User provides a password to be used to connect to the database.
type DatabasePasswordCredentialDetails struct {

	// The database user's password encoded using BASE64 scheme.
	Password *string `mandatory:"true" json:"password"`

	// The user to connect to the database.
	Username *string `mandatory:"false" json:"username"`

	// The role of the database user.
	Role DatabasePasswordCredentialDetailsRoleEnum `mandatory:"false" json:"role,omitempty"`
}

func (m DatabasePasswordCredentialDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabasePasswordCredentialDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabasePasswordCredentialDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDatabasePasswordCredentialDetailsRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabasePasswordCredentialDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabasePasswordCredentialDetails DatabasePasswordCredentialDetails
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeDatabasePasswordCredentialDetails
	}{
		"PASSWORD",
		(MarshalTypeDatabasePasswordCredentialDetails)(m),
	}

	return json.Marshal(&s)
}

// DatabasePasswordCredentialDetailsRoleEnum Enum with underlying type: string
type DatabasePasswordCredentialDetailsRoleEnum string

// Set of constants representing the allowable values for DatabasePasswordCredentialDetailsRoleEnum
const (
	DatabasePasswordCredentialDetailsRoleNormal DatabasePasswordCredentialDetailsRoleEnum = "NORMAL"
	DatabasePasswordCredentialDetailsRoleSysdba DatabasePasswordCredentialDetailsRoleEnum = "SYSDBA"
	DatabasePasswordCredentialDetailsRoleSysdg  DatabasePasswordCredentialDetailsRoleEnum = "SYSDG"
)

var mappingDatabasePasswordCredentialDetailsRoleEnum = map[string]DatabasePasswordCredentialDetailsRoleEnum{
	"NORMAL": DatabasePasswordCredentialDetailsRoleNormal,
	"SYSDBA": DatabasePasswordCredentialDetailsRoleSysdba,
	"SYSDG":  DatabasePasswordCredentialDetailsRoleSysdg,
}

var mappingDatabasePasswordCredentialDetailsRoleEnumLowerCase = map[string]DatabasePasswordCredentialDetailsRoleEnum{
	"normal": DatabasePasswordCredentialDetailsRoleNormal,
	"sysdba": DatabasePasswordCredentialDetailsRoleSysdba,
	"sysdg":  DatabasePasswordCredentialDetailsRoleSysdg,
}

// GetDatabasePasswordCredentialDetailsRoleEnumValues Enumerates the set of values for DatabasePasswordCredentialDetailsRoleEnum
func GetDatabasePasswordCredentialDetailsRoleEnumValues() []DatabasePasswordCredentialDetailsRoleEnum {
	values := make([]DatabasePasswordCredentialDetailsRoleEnum, 0)
	for _, v := range mappingDatabasePasswordCredentialDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabasePasswordCredentialDetailsRoleEnumStringValues Enumerates the set of values in String for DatabasePasswordCredentialDetailsRoleEnum
func GetDatabasePasswordCredentialDetailsRoleEnumStringValues() []string {
	return []string{
		"NORMAL",
		"SYSDBA",
		"SYSDG",
	}
}

// GetMappingDatabasePasswordCredentialDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabasePasswordCredentialDetailsRoleEnum(val string) (DatabasePasswordCredentialDetailsRoleEnum, bool) {
	enum, ok := mappingDatabasePasswordCredentialDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
