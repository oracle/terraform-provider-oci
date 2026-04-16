// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityConfig Security configuration for the Managed Kafka destination.
type SecurityConfig struct {

	// Security protocol of the Managed Kafka.
	SecurityProtocol SecurityConfigSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`

	// OCID of the vault containing the SASL credentials or MTLS client certificate.
	VaultId *string `mandatory:"true" json:"vaultId"`

	// OCID of the vault containing the trust store certificate to verify brokers identity.
	TrustStoreVaultId *string `mandatory:"false" json:"trustStoreVaultId"`
}

func (m SecurityConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecurityConfigSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetSecurityConfigSecurityProtocolEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecurityConfigSecurityProtocolEnum Enum with underlying type: string
type SecurityConfigSecurityProtocolEnum string

// Set of constants representing the allowable values for SecurityConfigSecurityProtocolEnum
const (
	SecurityConfigSecurityProtocolSasl SecurityConfigSecurityProtocolEnum = "SASL"
	SecurityConfigSecurityProtocolMtls SecurityConfigSecurityProtocolEnum = "MTLS"
)

var mappingSecurityConfigSecurityProtocolEnum = map[string]SecurityConfigSecurityProtocolEnum{
	"SASL": SecurityConfigSecurityProtocolSasl,
	"MTLS": SecurityConfigSecurityProtocolMtls,
}

var mappingSecurityConfigSecurityProtocolEnumLowerCase = map[string]SecurityConfigSecurityProtocolEnum{
	"sasl": SecurityConfigSecurityProtocolSasl,
	"mtls": SecurityConfigSecurityProtocolMtls,
}

// GetSecurityConfigSecurityProtocolEnumValues Enumerates the set of values for SecurityConfigSecurityProtocolEnum
func GetSecurityConfigSecurityProtocolEnumValues() []SecurityConfigSecurityProtocolEnum {
	values := make([]SecurityConfigSecurityProtocolEnum, 0)
	for _, v := range mappingSecurityConfigSecurityProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityConfigSecurityProtocolEnumStringValues Enumerates the set of values in String for SecurityConfigSecurityProtocolEnum
func GetSecurityConfigSecurityProtocolEnumStringValues() []string {
	return []string{
		"SASL",
		"MTLS",
	}
}

// GetMappingSecurityConfigSecurityProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityConfigSecurityProtocolEnum(val string) (SecurityConfigSecurityProtocolEnum, bool) {
	enum, ok := mappingSecurityConfigSecurityProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
