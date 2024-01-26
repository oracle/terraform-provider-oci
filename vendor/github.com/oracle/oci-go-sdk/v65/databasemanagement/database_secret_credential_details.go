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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseSecretCredentialDetails User provides a secret OCID, which will be used to retrieve the password to connect to the database.
type DatabaseSecretCredentialDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Secret
	// where the database password is stored.
	PasswordSecretId *string `mandatory:"true" json:"passwordSecretId"`

	// The user to connect to the database.
	Username *string `mandatory:"false" json:"username"`

	// The role of the database user.
	Role DatabaseSecretCredentialDetailsRoleEnum `mandatory:"false" json:"role,omitempty"`
}

func (m DatabaseSecretCredentialDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseSecretCredentialDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseSecretCredentialDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDatabaseSecretCredentialDetailsRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseSecretCredentialDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseSecretCredentialDetails DatabaseSecretCredentialDetails
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeDatabaseSecretCredentialDetails
	}{
		"SECRET",
		(MarshalTypeDatabaseSecretCredentialDetails)(m),
	}

	return json.Marshal(&s)
}

// DatabaseSecretCredentialDetailsRoleEnum Enum with underlying type: string
type DatabaseSecretCredentialDetailsRoleEnum string

// Set of constants representing the allowable values for DatabaseSecretCredentialDetailsRoleEnum
const (
	DatabaseSecretCredentialDetailsRoleNormal DatabaseSecretCredentialDetailsRoleEnum = "NORMAL"
	DatabaseSecretCredentialDetailsRoleSysdba DatabaseSecretCredentialDetailsRoleEnum = "SYSDBA"
)

var mappingDatabaseSecretCredentialDetailsRoleEnum = map[string]DatabaseSecretCredentialDetailsRoleEnum{
	"NORMAL": DatabaseSecretCredentialDetailsRoleNormal,
	"SYSDBA": DatabaseSecretCredentialDetailsRoleSysdba,
}

var mappingDatabaseSecretCredentialDetailsRoleEnumLowerCase = map[string]DatabaseSecretCredentialDetailsRoleEnum{
	"normal": DatabaseSecretCredentialDetailsRoleNormal,
	"sysdba": DatabaseSecretCredentialDetailsRoleSysdba,
}

// GetDatabaseSecretCredentialDetailsRoleEnumValues Enumerates the set of values for DatabaseSecretCredentialDetailsRoleEnum
func GetDatabaseSecretCredentialDetailsRoleEnumValues() []DatabaseSecretCredentialDetailsRoleEnum {
	values := make([]DatabaseSecretCredentialDetailsRoleEnum, 0)
	for _, v := range mappingDatabaseSecretCredentialDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseSecretCredentialDetailsRoleEnumStringValues Enumerates the set of values in String for DatabaseSecretCredentialDetailsRoleEnum
func GetDatabaseSecretCredentialDetailsRoleEnumStringValues() []string {
	return []string{
		"NORMAL",
		"SYSDBA",
	}
}

// GetMappingDatabaseSecretCredentialDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseSecretCredentialDetailsRoleEnum(val string) (DatabaseSecretCredentialDetailsRoleEnum, bool) {
	enum, ok := mappingDatabaseSecretCredentialDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
