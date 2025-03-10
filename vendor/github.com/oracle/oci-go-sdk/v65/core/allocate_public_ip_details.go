// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AllocatePublicIpDetails Allocates the details of public ip address
type AllocatePublicIpDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP address associated with the NAT gateway.
	PublicIpId *string `mandatory:"true" json:"publicIpId"`

	// The Tenancy's Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for which the QoS template is applicable.
	Tenancy *string `mandatory:"true" json:"tenancy"`

	// Public IP pool name. Name of the oracle pool to which the Cidr belongs to. This refers to the OraclePoolId
	// class defined in VCNIP.
	OraclePoolId AllocatePublicIpDetailsOraclePoolIdEnum `mandatory:"true" json:"oraclePoolId"`

	// Public IP Pool ID. Name of the public pool Id to which the Cidr belongs to.
	PublicPoolId *string `mandatory:"true" json:"publicPoolId"`
}

func (m AllocatePublicIpDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AllocatePublicIpDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAllocatePublicIpDetailsOraclePoolIdEnum(string(m.OraclePoolId)); !ok && m.OraclePoolId != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OraclePoolId: %s. Supported values are: %s.", m.OraclePoolId, strings.Join(GetAllocatePublicIpDetailsOraclePoolIdEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AllocatePublicIpDetailsOraclePoolIdEnum Enum with underlying type: string
type AllocatePublicIpDetailsOraclePoolIdEnum string

// Set of constants representing the allowable values for AllocatePublicIpDetailsOraclePoolIdEnum
const (
	AllocatePublicIpDetailsOraclePoolIdDefault     AllocatePublicIpDetailsOraclePoolIdEnum = "DEFAULT"
	AllocatePublicIpDetailsOraclePoolIdSociEgress  AllocatePublicIpDetailsOraclePoolIdEnum = "SOCI_EGRESS"
	AllocatePublicIpDetailsOraclePoolIdSociIngress AllocatePublicIpDetailsOraclePoolIdEnum = "SOCI_INGRESS"
	AllocatePublicIpDetailsOraclePoolIdOracleDev   AllocatePublicIpDetailsOraclePoolIdEnum = "ORACLE_DEV"
)

var mappingAllocatePublicIpDetailsOraclePoolIdEnum = map[string]AllocatePublicIpDetailsOraclePoolIdEnum{
	"DEFAULT":      AllocatePublicIpDetailsOraclePoolIdDefault,
	"SOCI_EGRESS":  AllocatePublicIpDetailsOraclePoolIdSociEgress,
	"SOCI_INGRESS": AllocatePublicIpDetailsOraclePoolIdSociIngress,
	"ORACLE_DEV":   AllocatePublicIpDetailsOraclePoolIdOracleDev,
}

var mappingAllocatePublicIpDetailsOraclePoolIdEnumLowerCase = map[string]AllocatePublicIpDetailsOraclePoolIdEnum{
	"default":      AllocatePublicIpDetailsOraclePoolIdDefault,
	"soci_egress":  AllocatePublicIpDetailsOraclePoolIdSociEgress,
	"soci_ingress": AllocatePublicIpDetailsOraclePoolIdSociIngress,
	"oracle_dev":   AllocatePublicIpDetailsOraclePoolIdOracleDev,
}

// GetAllocatePublicIpDetailsOraclePoolIdEnumValues Enumerates the set of values for AllocatePublicIpDetailsOraclePoolIdEnum
func GetAllocatePublicIpDetailsOraclePoolIdEnumValues() []AllocatePublicIpDetailsOraclePoolIdEnum {
	values := make([]AllocatePublicIpDetailsOraclePoolIdEnum, 0)
	for _, v := range mappingAllocatePublicIpDetailsOraclePoolIdEnum {
		values = append(values, v)
	}
	return values
}

// GetAllocatePublicIpDetailsOraclePoolIdEnumStringValues Enumerates the set of values in String for AllocatePublicIpDetailsOraclePoolIdEnum
func GetAllocatePublicIpDetailsOraclePoolIdEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"SOCI_EGRESS",
		"SOCI_INGRESS",
		"ORACLE_DEV",
	}
}

// GetMappingAllocatePublicIpDetailsOraclePoolIdEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAllocatePublicIpDetailsOraclePoolIdEnum(val string) (AllocatePublicIpDetailsOraclePoolIdEnum, bool) {
	enum, ok := mappingAllocatePublicIpDetailsOraclePoolIdEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
