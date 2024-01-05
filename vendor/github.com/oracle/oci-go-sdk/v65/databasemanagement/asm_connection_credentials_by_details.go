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

// AsmConnectionCredentialsByDetails The credentials used to connect to the ASM instance.
type AsmConnectionCredentialsByDetails struct {

	// The user name used to connect to the ASM instance.
	UserName *string `mandatory:"true" json:"userName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
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
	Role AsmConnectionCredentialsByDetailsRoleEnum `mandatory:"true" json:"role"`
}

func (m AsmConnectionCredentialsByDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AsmConnectionCredentialsByDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAsmConnectionCredentialsByDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetAsmConnectionCredentialsByDetailsRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AsmConnectionCredentialsByDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAsmConnectionCredentialsByDetails AsmConnectionCredentialsByDetails
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeAsmConnectionCredentialsByDetails
	}{
		"DETAILS",
		(MarshalTypeAsmConnectionCredentialsByDetails)(m),
	}

	return json.Marshal(&s)
}

// AsmConnectionCredentialsByDetailsRoleEnum Enum with underlying type: string
type AsmConnectionCredentialsByDetailsRoleEnum string

// Set of constants representing the allowable values for AsmConnectionCredentialsByDetailsRoleEnum
const (
	AsmConnectionCredentialsByDetailsRoleSysasm  AsmConnectionCredentialsByDetailsRoleEnum = "SYSASM"
	AsmConnectionCredentialsByDetailsRoleSysdba  AsmConnectionCredentialsByDetailsRoleEnum = "SYSDBA"
	AsmConnectionCredentialsByDetailsRoleSysoper AsmConnectionCredentialsByDetailsRoleEnum = "SYSOPER"
)

var mappingAsmConnectionCredentialsByDetailsRoleEnum = map[string]AsmConnectionCredentialsByDetailsRoleEnum{
	"SYSASM":  AsmConnectionCredentialsByDetailsRoleSysasm,
	"SYSDBA":  AsmConnectionCredentialsByDetailsRoleSysdba,
	"SYSOPER": AsmConnectionCredentialsByDetailsRoleSysoper,
}

var mappingAsmConnectionCredentialsByDetailsRoleEnumLowerCase = map[string]AsmConnectionCredentialsByDetailsRoleEnum{
	"sysasm":  AsmConnectionCredentialsByDetailsRoleSysasm,
	"sysdba":  AsmConnectionCredentialsByDetailsRoleSysdba,
	"sysoper": AsmConnectionCredentialsByDetailsRoleSysoper,
}

// GetAsmConnectionCredentialsByDetailsRoleEnumValues Enumerates the set of values for AsmConnectionCredentialsByDetailsRoleEnum
func GetAsmConnectionCredentialsByDetailsRoleEnumValues() []AsmConnectionCredentialsByDetailsRoleEnum {
	values := make([]AsmConnectionCredentialsByDetailsRoleEnum, 0)
	for _, v := range mappingAsmConnectionCredentialsByDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetAsmConnectionCredentialsByDetailsRoleEnumStringValues Enumerates the set of values in String for AsmConnectionCredentialsByDetailsRoleEnum
func GetAsmConnectionCredentialsByDetailsRoleEnumStringValues() []string {
	return []string{
		"SYSASM",
		"SYSDBA",
		"SYSOPER",
	}
}

// GetMappingAsmConnectionCredentialsByDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAsmConnectionCredentialsByDetailsRoleEnum(val string) (AsmConnectionCredentialsByDetailsRoleEnum, bool) {
	enum, ok := mappingAsmConnectionCredentialsByDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
