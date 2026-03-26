// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// VaultMtlsUpdateResponse The response of the update mTLS vault api call.
type VaultMtlsUpdateResponse struct {

	// The mTLS DNS endpoint for performing crypto operations.
	MtlsCryptoDns *string `mandatory:"true" json:"mtlsCryptoDns"`

	// Identifier of the fleet associated with the mTLS connection.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// The mtls vault registering state.
	Status VaultMtlsUpdateResponseStatusEnum `mandatory:"true" json:"status"`
}

func (m VaultMtlsUpdateResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VaultMtlsUpdateResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVaultMtlsUpdateResponseStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetVaultMtlsUpdateResponseStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VaultMtlsUpdateResponseStatusEnum Enum with underlying type: string
type VaultMtlsUpdateResponseStatusEnum string

// Set of constants representing the allowable values for VaultMtlsUpdateResponseStatusEnum
const (
	VaultMtlsUpdateResponseStatusCreating VaultMtlsUpdateResponseStatusEnum = "CREATING"
	VaultMtlsUpdateResponseStatusCreated  VaultMtlsUpdateResponseStatusEnum = "CREATED"
)

var mappingVaultMtlsUpdateResponseStatusEnum = map[string]VaultMtlsUpdateResponseStatusEnum{
	"CREATING": VaultMtlsUpdateResponseStatusCreating,
	"CREATED":  VaultMtlsUpdateResponseStatusCreated,
}

var mappingVaultMtlsUpdateResponseStatusEnumLowerCase = map[string]VaultMtlsUpdateResponseStatusEnum{
	"creating": VaultMtlsUpdateResponseStatusCreating,
	"created":  VaultMtlsUpdateResponseStatusCreated,
}

// GetVaultMtlsUpdateResponseStatusEnumValues Enumerates the set of values for VaultMtlsUpdateResponseStatusEnum
func GetVaultMtlsUpdateResponseStatusEnumValues() []VaultMtlsUpdateResponseStatusEnum {
	values := make([]VaultMtlsUpdateResponseStatusEnum, 0)
	for _, v := range mappingVaultMtlsUpdateResponseStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetVaultMtlsUpdateResponseStatusEnumStringValues Enumerates the set of values in String for VaultMtlsUpdateResponseStatusEnum
func GetVaultMtlsUpdateResponseStatusEnumStringValues() []string {
	return []string{
		"CREATING",
		"CREATED",
	}
}

// GetMappingVaultMtlsUpdateResponseStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVaultMtlsUpdateResponseStatusEnum(val string) (VaultMtlsUpdateResponseStatusEnum, bool) {
	enum, ok := mappingVaultMtlsUpdateResponseStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
