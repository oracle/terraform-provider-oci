// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IpCapacity A public IP capacity CIDR owned by VCN CP with its allocation status. This structure will be used to transfer
// staging table entries from VCNIP to VCNCP.
type IpCapacity struct {

	// CIDR of the public IP capacity
	Cidr *string `mandatory:"false" json:"cidr"`

	// Public IP pool. Name of the oracle pool Id to which the Cidr belongs to. This refers to the OraclePoolId
	// class defined in VCNIP.
	OraclePoolId IpCapacityOraclePoolIdEnum `mandatory:"false" json:"oraclePoolId,omitempty"`

	// Status of the public IP CIDR whether it is currently available to allocate to users.
	StagingCidrState IpCapacityStagingCidrStateEnum `mandatory:"false" json:"stagingCidrState,omitempty"`

	// Size of the address/CIDR configured to be allocated from this capacity.
	PrefixLength *int `mandatory:"false" json:"prefixLength"`
}

func (m IpCapacity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IpCapacity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIpCapacityOraclePoolIdEnum(string(m.OraclePoolId)); !ok && m.OraclePoolId != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OraclePoolId: %s. Supported values are: %s.", m.OraclePoolId, strings.Join(GetIpCapacityOraclePoolIdEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIpCapacityStagingCidrStateEnum(string(m.StagingCidrState)); !ok && m.StagingCidrState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StagingCidrState: %s. Supported values are: %s.", m.StagingCidrState, strings.Join(GetIpCapacityStagingCidrStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IpCapacityOraclePoolIdEnum Enum with underlying type: string
type IpCapacityOraclePoolIdEnum string

// Set of constants representing the allowable values for IpCapacityOraclePoolIdEnum
const (
	IpCapacityOraclePoolIdDefault     IpCapacityOraclePoolIdEnum = "DEFAULT"
	IpCapacityOraclePoolIdSociEgress  IpCapacityOraclePoolIdEnum = "SOCI_EGRESS"
	IpCapacityOraclePoolIdSociIngress IpCapacityOraclePoolIdEnum = "SOCI_INGRESS"
	IpCapacityOraclePoolIdOracleDev   IpCapacityOraclePoolIdEnum = "ORACLE_DEV"
)

var mappingIpCapacityOraclePoolIdEnum = map[string]IpCapacityOraclePoolIdEnum{
	"DEFAULT":      IpCapacityOraclePoolIdDefault,
	"SOCI_EGRESS":  IpCapacityOraclePoolIdSociEgress,
	"SOCI_INGRESS": IpCapacityOraclePoolIdSociIngress,
	"ORACLE_DEV":   IpCapacityOraclePoolIdOracleDev,
}

var mappingIpCapacityOraclePoolIdEnumLowerCase = map[string]IpCapacityOraclePoolIdEnum{
	"default":      IpCapacityOraclePoolIdDefault,
	"soci_egress":  IpCapacityOraclePoolIdSociEgress,
	"soci_ingress": IpCapacityOraclePoolIdSociIngress,
	"oracle_dev":   IpCapacityOraclePoolIdOracleDev,
}

// GetIpCapacityOraclePoolIdEnumValues Enumerates the set of values for IpCapacityOraclePoolIdEnum
func GetIpCapacityOraclePoolIdEnumValues() []IpCapacityOraclePoolIdEnum {
	values := make([]IpCapacityOraclePoolIdEnum, 0)
	for _, v := range mappingIpCapacityOraclePoolIdEnum {
		values = append(values, v)
	}
	return values
}

// GetIpCapacityOraclePoolIdEnumStringValues Enumerates the set of values in String for IpCapacityOraclePoolIdEnum
func GetIpCapacityOraclePoolIdEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"SOCI_EGRESS",
		"SOCI_INGRESS",
		"ORACLE_DEV",
	}
}

// GetMappingIpCapacityOraclePoolIdEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIpCapacityOraclePoolIdEnum(val string) (IpCapacityOraclePoolIdEnum, bool) {
	enum, ok := mappingIpCapacityOraclePoolIdEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IpCapacityStagingCidrStateEnum Enum with underlying type: string
type IpCapacityStagingCidrStateEnum string

// Set of constants representing the allowable values for IpCapacityStagingCidrStateEnum
const (
	IpCapacityStagingCidrStateProvisioned IpCapacityStagingCidrStateEnum = "PROVISIONED"
	IpCapacityStagingCidrStateInCapacity  IpCapacityStagingCidrStateEnum = "IN_CAPACITY"
)

var mappingIpCapacityStagingCidrStateEnum = map[string]IpCapacityStagingCidrStateEnum{
	"PROVISIONED": IpCapacityStagingCidrStateProvisioned,
	"IN_CAPACITY": IpCapacityStagingCidrStateInCapacity,
}

var mappingIpCapacityStagingCidrStateEnumLowerCase = map[string]IpCapacityStagingCidrStateEnum{
	"provisioned": IpCapacityStagingCidrStateProvisioned,
	"in_capacity": IpCapacityStagingCidrStateInCapacity,
}

// GetIpCapacityStagingCidrStateEnumValues Enumerates the set of values for IpCapacityStagingCidrStateEnum
func GetIpCapacityStagingCidrStateEnumValues() []IpCapacityStagingCidrStateEnum {
	values := make([]IpCapacityStagingCidrStateEnum, 0)
	for _, v := range mappingIpCapacityStagingCidrStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIpCapacityStagingCidrStateEnumStringValues Enumerates the set of values in String for IpCapacityStagingCidrStateEnum
func GetIpCapacityStagingCidrStateEnumStringValues() []string {
	return []string{
		"PROVISIONED",
		"IN_CAPACITY",
	}
}

// GetMappingIpCapacityStagingCidrStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIpCapacityStagingCidrStateEnum(val string) (IpCapacityStagingCidrStateEnum, bool) {
	enum, ok := mappingIpCapacityStagingCidrStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
