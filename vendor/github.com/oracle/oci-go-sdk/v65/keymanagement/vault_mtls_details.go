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

// VaultMtlsDetails Register mTLS configuration request details for a vault.
type VaultMtlsDetails struct {

	// Identifier of the fleet associated with the mTLS connection.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// The mtls vault registering state.
	Status VaultMtlsDetailsStatusEnum `mandatory:"true" json:"status"`
}

func (m VaultMtlsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VaultMtlsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVaultMtlsDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetVaultMtlsDetailsStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VaultMtlsDetailsStatusEnum Enum with underlying type: string
type VaultMtlsDetailsStatusEnum string

// Set of constants representing the allowable values for VaultMtlsDetailsStatusEnum
const (
	VaultMtlsDetailsStatusCreating VaultMtlsDetailsStatusEnum = "CREATING"
	VaultMtlsDetailsStatusCreated  VaultMtlsDetailsStatusEnum = "CREATED"
)

var mappingVaultMtlsDetailsStatusEnum = map[string]VaultMtlsDetailsStatusEnum{
	"CREATING": VaultMtlsDetailsStatusCreating,
	"CREATED":  VaultMtlsDetailsStatusCreated,
}

var mappingVaultMtlsDetailsStatusEnumLowerCase = map[string]VaultMtlsDetailsStatusEnum{
	"creating": VaultMtlsDetailsStatusCreating,
	"created":  VaultMtlsDetailsStatusCreated,
}

// GetVaultMtlsDetailsStatusEnumValues Enumerates the set of values for VaultMtlsDetailsStatusEnum
func GetVaultMtlsDetailsStatusEnumValues() []VaultMtlsDetailsStatusEnum {
	values := make([]VaultMtlsDetailsStatusEnum, 0)
	for _, v := range mappingVaultMtlsDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetVaultMtlsDetailsStatusEnumStringValues Enumerates the set of values in String for VaultMtlsDetailsStatusEnum
func GetVaultMtlsDetailsStatusEnumStringValues() []string {
	return []string{
		"CREATING",
		"CREATED",
	}
}

// GetMappingVaultMtlsDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVaultMtlsDetailsStatusEnum(val string) (VaultMtlsDetailsStatusEnum, bool) {
	enum, ok := mappingVaultMtlsDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
