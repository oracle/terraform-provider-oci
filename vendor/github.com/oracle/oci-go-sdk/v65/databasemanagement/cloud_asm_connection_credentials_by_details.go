// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudAsmConnectionCredentialsByDetails The credentials used to connect to the ASM instance.
type CloudAsmConnectionCredentialsByDetails struct {

	// The user name used to connect to the ASM instance.
	UserName *string `mandatory:"true" json:"userName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
	PasswordSecretId *string `mandatory:"true" json:"passwordSecretId"`

	// The name of the credential information that used to connect to the DB system resource.
	// The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters,
	// and length of "y" has a maximum of 199 characters. The name strings can contain letters,
	// numbers and the underscore character only. Other characters are not valid, except for
	// the "." character that separates the "x" and "y" portions of the name.
	// *IMPORTANT* - The name must be unique within the OCI region the credential is being created in.
	// If you specify a name that duplicates the name of another credential within the same OCI region,
	// you may overwrite or corrupt the credential that is already using the name.
	// For example: inventorydb.abc112233445566778899
	CredentialName *string `mandatory:"false" json:"credentialName"`

	// The role of the user connecting to the ASM instance.
	Role CloudAsmConnectionCredentialsByDetailsRoleEnum `mandatory:"true" json:"role"`
}

func (m CloudAsmConnectionCredentialsByDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudAsmConnectionCredentialsByDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudAsmConnectionCredentialsByDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetCloudAsmConnectionCredentialsByDetailsRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CloudAsmConnectionCredentialsByDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCloudAsmConnectionCredentialsByDetails CloudAsmConnectionCredentialsByDetails
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeCloudAsmConnectionCredentialsByDetails
	}{
		"DETAILS",
		(MarshalTypeCloudAsmConnectionCredentialsByDetails)(m),
	}

	return json.Marshal(&s)
}

// CloudAsmConnectionCredentialsByDetailsRoleEnum Enum with underlying type: string
type CloudAsmConnectionCredentialsByDetailsRoleEnum string

// Set of constants representing the allowable values for CloudAsmConnectionCredentialsByDetailsRoleEnum
const (
	CloudAsmConnectionCredentialsByDetailsRoleSysasm  CloudAsmConnectionCredentialsByDetailsRoleEnum = "SYSASM"
	CloudAsmConnectionCredentialsByDetailsRoleSysdba  CloudAsmConnectionCredentialsByDetailsRoleEnum = "SYSDBA"
	CloudAsmConnectionCredentialsByDetailsRoleSysoper CloudAsmConnectionCredentialsByDetailsRoleEnum = "SYSOPER"
)

var mappingCloudAsmConnectionCredentialsByDetailsRoleEnum = map[string]CloudAsmConnectionCredentialsByDetailsRoleEnum{
	"SYSASM":  CloudAsmConnectionCredentialsByDetailsRoleSysasm,
	"SYSDBA":  CloudAsmConnectionCredentialsByDetailsRoleSysdba,
	"SYSOPER": CloudAsmConnectionCredentialsByDetailsRoleSysoper,
}

var mappingCloudAsmConnectionCredentialsByDetailsRoleEnumLowerCase = map[string]CloudAsmConnectionCredentialsByDetailsRoleEnum{
	"sysasm":  CloudAsmConnectionCredentialsByDetailsRoleSysasm,
	"sysdba":  CloudAsmConnectionCredentialsByDetailsRoleSysdba,
	"sysoper": CloudAsmConnectionCredentialsByDetailsRoleSysoper,
}

// GetCloudAsmConnectionCredentialsByDetailsRoleEnumValues Enumerates the set of values for CloudAsmConnectionCredentialsByDetailsRoleEnum
func GetCloudAsmConnectionCredentialsByDetailsRoleEnumValues() []CloudAsmConnectionCredentialsByDetailsRoleEnum {
	values := make([]CloudAsmConnectionCredentialsByDetailsRoleEnum, 0)
	for _, v := range mappingCloudAsmConnectionCredentialsByDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudAsmConnectionCredentialsByDetailsRoleEnumStringValues Enumerates the set of values in String for CloudAsmConnectionCredentialsByDetailsRoleEnum
func GetCloudAsmConnectionCredentialsByDetailsRoleEnumStringValues() []string {
	return []string{
		"SYSASM",
		"SYSDBA",
		"SYSOPER",
	}
}

// GetMappingCloudAsmConnectionCredentialsByDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudAsmConnectionCredentialsByDetailsRoleEnum(val string) (CloudAsmConnectionCredentialsByDetailsRoleEnum, bool) {
	enum, ok := mappingCloudAsmConnectionCredentialsByDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
