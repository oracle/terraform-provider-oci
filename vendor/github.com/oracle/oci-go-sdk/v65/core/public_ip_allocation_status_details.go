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

// PublicIpAllocationStatusDetails Status on allocation of public ip address
type PublicIpAllocationStatusDetails struct {

	// The IP address that needs allocation status details.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP address associated with the ip address.
	PublicIpId *string `mandatory:"true" json:"publicIpId"`

	// Public IP pool name. Name of the oracle pool to which the Cidr belongs to. This refers to the OraclePoolId
	// class defined in VCNIP.
	OraclePoolId PublicIpAllocationStatusDetailsOraclePoolIdEnum `mandatory:"false" json:"oraclePoolId,omitempty"`

	// Public IP Pool ID. Name of the public pool Id to which the Cidr belongs to.
	PublicPoolId *string `mandatory:"false" json:"publicPoolId"`

	// The Tenancy's Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for which the QoS template is applicable.
	Tenancy *string `mandatory:"false" json:"tenancy"`
}

func (m PublicIpAllocationStatusDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PublicIpAllocationStatusDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPublicIpAllocationStatusDetailsOraclePoolIdEnum(string(m.OraclePoolId)); !ok && m.OraclePoolId != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OraclePoolId: %s. Supported values are: %s.", m.OraclePoolId, strings.Join(GetPublicIpAllocationStatusDetailsOraclePoolIdEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PublicIpAllocationStatusDetailsOraclePoolIdEnum Enum with underlying type: string
type PublicIpAllocationStatusDetailsOraclePoolIdEnum string

// Set of constants representing the allowable values for PublicIpAllocationStatusDetailsOraclePoolIdEnum
const (
	PublicIpAllocationStatusDetailsOraclePoolIdDefault     PublicIpAllocationStatusDetailsOraclePoolIdEnum = "DEFAULT"
	PublicIpAllocationStatusDetailsOraclePoolIdSociEgress  PublicIpAllocationStatusDetailsOraclePoolIdEnum = "SOCI_EGRESS"
	PublicIpAllocationStatusDetailsOraclePoolIdSociIngress PublicIpAllocationStatusDetailsOraclePoolIdEnum = "SOCI_INGRESS"
	PublicIpAllocationStatusDetailsOraclePoolIdOracleDev   PublicIpAllocationStatusDetailsOraclePoolIdEnum = "ORACLE_DEV"
)

var mappingPublicIpAllocationStatusDetailsOraclePoolIdEnum = map[string]PublicIpAllocationStatusDetailsOraclePoolIdEnum{
	"DEFAULT":      PublicIpAllocationStatusDetailsOraclePoolIdDefault,
	"SOCI_EGRESS":  PublicIpAllocationStatusDetailsOraclePoolIdSociEgress,
	"SOCI_INGRESS": PublicIpAllocationStatusDetailsOraclePoolIdSociIngress,
	"ORACLE_DEV":   PublicIpAllocationStatusDetailsOraclePoolIdOracleDev,
}

var mappingPublicIpAllocationStatusDetailsOraclePoolIdEnumLowerCase = map[string]PublicIpAllocationStatusDetailsOraclePoolIdEnum{
	"default":      PublicIpAllocationStatusDetailsOraclePoolIdDefault,
	"soci_egress":  PublicIpAllocationStatusDetailsOraclePoolIdSociEgress,
	"soci_ingress": PublicIpAllocationStatusDetailsOraclePoolIdSociIngress,
	"oracle_dev":   PublicIpAllocationStatusDetailsOraclePoolIdOracleDev,
}

// GetPublicIpAllocationStatusDetailsOraclePoolIdEnumValues Enumerates the set of values for PublicIpAllocationStatusDetailsOraclePoolIdEnum
func GetPublicIpAllocationStatusDetailsOraclePoolIdEnumValues() []PublicIpAllocationStatusDetailsOraclePoolIdEnum {
	values := make([]PublicIpAllocationStatusDetailsOraclePoolIdEnum, 0)
	for _, v := range mappingPublicIpAllocationStatusDetailsOraclePoolIdEnum {
		values = append(values, v)
	}
	return values
}

// GetPublicIpAllocationStatusDetailsOraclePoolIdEnumStringValues Enumerates the set of values in String for PublicIpAllocationStatusDetailsOraclePoolIdEnum
func GetPublicIpAllocationStatusDetailsOraclePoolIdEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"SOCI_EGRESS",
		"SOCI_INGRESS",
		"ORACLE_DEV",
	}
}

// GetMappingPublicIpAllocationStatusDetailsOraclePoolIdEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublicIpAllocationStatusDetailsOraclePoolIdEnum(val string) (PublicIpAllocationStatusDetailsOraclePoolIdEnum, bool) {
	enum, ok := mappingPublicIpAllocationStatusDetailsOraclePoolIdEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
