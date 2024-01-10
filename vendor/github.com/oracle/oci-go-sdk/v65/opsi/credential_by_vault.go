// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CredentialByVault Vault Credential Details to connect to the database.
type CredentialByVault struct {

	// Credential source name that had been added in Management Agent wallet. This is supplied in the External Database Service.
	CredentialSourceName *string `mandatory:"true" json:"credentialSourceName"`

	// database user name.
	UserName *string `mandatory:"false" json:"userName"`

	// The secret OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) mapping to the database credentials.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Secret where the database keystore contents are stored. This is used for TCPS support in BM/VM/ExaCS cases.
	WalletSecretId *string `mandatory:"false" json:"walletSecretId"`

	// database user role.
	Role CredentialByVaultRoleEnum `mandatory:"false" json:"role,omitempty"`
}

// GetCredentialSourceName returns CredentialSourceName
func (m CredentialByVault) GetCredentialSourceName() *string {
	return m.CredentialSourceName
}

func (m CredentialByVault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CredentialByVault) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCredentialByVaultRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetCredentialByVaultRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CredentialByVault) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCredentialByVault CredentialByVault
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeCredentialByVault
	}{
		"CREDENTIALS_BY_VAULT",
		(MarshalTypeCredentialByVault)(m),
	}

	return json.Marshal(&s)
}

// CredentialByVaultRoleEnum Enum with underlying type: string
type CredentialByVaultRoleEnum string

// Set of constants representing the allowable values for CredentialByVaultRoleEnum
const (
	CredentialByVaultRoleNormal CredentialByVaultRoleEnum = "NORMAL"
)

var mappingCredentialByVaultRoleEnum = map[string]CredentialByVaultRoleEnum{
	"NORMAL": CredentialByVaultRoleNormal,
}

var mappingCredentialByVaultRoleEnumLowerCase = map[string]CredentialByVaultRoleEnum{
	"normal": CredentialByVaultRoleNormal,
}

// GetCredentialByVaultRoleEnumValues Enumerates the set of values for CredentialByVaultRoleEnum
func GetCredentialByVaultRoleEnumValues() []CredentialByVaultRoleEnum {
	values := make([]CredentialByVaultRoleEnum, 0)
	for _, v := range mappingCredentialByVaultRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetCredentialByVaultRoleEnumStringValues Enumerates the set of values in String for CredentialByVaultRoleEnum
func GetCredentialByVaultRoleEnumStringValues() []string {
	return []string{
		"NORMAL",
	}
}

// GetMappingCredentialByVaultRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCredentialByVaultRoleEnum(val string) (CredentialByVaultRoleEnum, bool) {
	enum, ok := mappingCredentialByVaultRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
