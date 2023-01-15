// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseConnectionCredentialsByDetails User information to connect to the database. Required when performing the CreateExternalDatabaseConnectorDetails operation.
// *IMPORTANT*: Not supported for the UpdateExternalDatabaseConnectorDetails operation.
type DatabaseConnectionCredentialsByDetails struct {

	// The username that will be used to connect to the database.
	Username *string `mandatory:"true" json:"username"`

	// The password that will be used to connect to the database.
	Password *string `mandatory:"true" json:"password"`

	// The name of the credential information that used to connect to the database. The name should be in "x.y" format, where
	// the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters.
	// The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for
	// the "." character that separates the "x" and "y" portions of the name.
	// *IMPORTANT* - The name must be unique within the OCI region the credential is being created in. If you specify a name
	// that duplicates the name of another credential within the same OCI region, you may overwrite or corrupt the credential that is already
	// using the name.
	// For example: inventorydb.abc112233445566778899
	CredentialName *string `mandatory:"false" json:"credentialName"`

	// The role of the user that will be connecting to the database.
	Role DatabaseConnectionCredentialsByDetailsRoleEnum `mandatory:"true" json:"role"`
}

func (m DatabaseConnectionCredentialsByDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseConnectionCredentialsByDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseConnectionCredentialsByDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDatabaseConnectionCredentialsByDetailsRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingDatabaseConnectionCredentialsByDetailsRoleEnum = map[string]DatabaseConnectionCredentialsByDetailsRoleEnum{
	"SYSDBA": DatabaseConnectionCredentialsByDetailsRoleSysdba,
	"NORMAL": DatabaseConnectionCredentialsByDetailsRoleNormal,
}

var mappingDatabaseConnectionCredentialsByDetailsRoleEnumLowerCase = map[string]DatabaseConnectionCredentialsByDetailsRoleEnum{
	"sysdba": DatabaseConnectionCredentialsByDetailsRoleSysdba,
	"normal": DatabaseConnectionCredentialsByDetailsRoleNormal,
}

// GetDatabaseConnectionCredentialsByDetailsRoleEnumValues Enumerates the set of values for DatabaseConnectionCredentialsByDetailsRoleEnum
func GetDatabaseConnectionCredentialsByDetailsRoleEnumValues() []DatabaseConnectionCredentialsByDetailsRoleEnum {
	values := make([]DatabaseConnectionCredentialsByDetailsRoleEnum, 0)
	for _, v := range mappingDatabaseConnectionCredentialsByDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConnectionCredentialsByDetailsRoleEnumStringValues Enumerates the set of values in String for DatabaseConnectionCredentialsByDetailsRoleEnum
func GetDatabaseConnectionCredentialsByDetailsRoleEnumStringValues() []string {
	return []string{
		"SYSDBA",
		"NORMAL",
	}
}

// GetMappingDatabaseConnectionCredentialsByDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConnectionCredentialsByDetailsRoleEnum(val string) (DatabaseConnectionCredentialsByDetailsRoleEnum, bool) {
	enum, ok := mappingDatabaseConnectionCredentialsByDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
