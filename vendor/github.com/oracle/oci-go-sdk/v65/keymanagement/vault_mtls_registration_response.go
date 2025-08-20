// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VaultMtlsRegistrationResponse The representation of VaultMtlsRegistrationResponse
type VaultMtlsRegistrationResponse struct {

	// The mTLS DNS endpoint for performing crypto operations..
	MtlsCryptoDns *string `mandatory:"true" json:"mtlsCryptoDns"`

	// The mtls vault registering state.
	Status VaultMtlsRegistrationResponseStatusEnum `mandatory:"true" json:"status"`
}

func (m VaultMtlsRegistrationResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VaultMtlsRegistrationResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVaultMtlsRegistrationResponseStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetVaultMtlsRegistrationResponseStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VaultMtlsRegistrationResponseStatusEnum Enum with underlying type: string
type VaultMtlsRegistrationResponseStatusEnum string

// Set of constants representing the allowable values for VaultMtlsRegistrationResponseStatusEnum
const (
	VaultMtlsRegistrationResponseStatusCreating VaultMtlsRegistrationResponseStatusEnum = "CREATING"
	VaultMtlsRegistrationResponseStatusCreated  VaultMtlsRegistrationResponseStatusEnum = "CREATED"
)

var mappingVaultMtlsRegistrationResponseStatusEnum = map[string]VaultMtlsRegistrationResponseStatusEnum{
	"CREATING": VaultMtlsRegistrationResponseStatusCreating,
	"CREATED":  VaultMtlsRegistrationResponseStatusCreated,
}

var mappingVaultMtlsRegistrationResponseStatusEnumLowerCase = map[string]VaultMtlsRegistrationResponseStatusEnum{
	"creating": VaultMtlsRegistrationResponseStatusCreating,
	"created":  VaultMtlsRegistrationResponseStatusCreated,
}

// GetVaultMtlsRegistrationResponseStatusEnumValues Enumerates the set of values for VaultMtlsRegistrationResponseStatusEnum
func GetVaultMtlsRegistrationResponseStatusEnumValues() []VaultMtlsRegistrationResponseStatusEnum {
	values := make([]VaultMtlsRegistrationResponseStatusEnum, 0)
	for _, v := range mappingVaultMtlsRegistrationResponseStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetVaultMtlsRegistrationResponseStatusEnumStringValues Enumerates the set of values in String for VaultMtlsRegistrationResponseStatusEnum
func GetVaultMtlsRegistrationResponseStatusEnumStringValues() []string {
	return []string{
		"CREATING",
		"CREATED",
	}
}

// GetMappingVaultMtlsRegistrationResponseStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVaultMtlsRegistrationResponseStatusEnum(val string) (VaultMtlsRegistrationResponseStatusEnum, bool) {
	enum, ok := mappingVaultMtlsRegistrationResponseStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
