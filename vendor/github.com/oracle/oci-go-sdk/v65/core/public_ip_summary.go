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

// PublicIpSummary Allocate public ip address response
type PublicIpSummary struct {

	// The IP address that needs to be released.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// Public IP pool name. Name of the oracle pool to which the Cidr belongs to. This refers to the OraclePoolId
	// class defined in VCNIP.
	OraclePoolId PublicIpSummaryOraclePoolIdEnum `mandatory:"false" json:"oraclePoolId,omitempty"`

	// Public IP Pool ID. Name of the public pool Id to which the Cidr belongs to.
	PublicPoolId *string `mandatory:"false" json:"publicPoolId"`
}

func (m PublicIpSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PublicIpSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPublicIpSummaryOraclePoolIdEnum(string(m.OraclePoolId)); !ok && m.OraclePoolId != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OraclePoolId: %s. Supported values are: %s.", m.OraclePoolId, strings.Join(GetPublicIpSummaryOraclePoolIdEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PublicIpSummaryOraclePoolIdEnum Enum with underlying type: string
type PublicIpSummaryOraclePoolIdEnum string

// Set of constants representing the allowable values for PublicIpSummaryOraclePoolIdEnum
const (
	PublicIpSummaryOraclePoolIdDefault     PublicIpSummaryOraclePoolIdEnum = "DEFAULT"
	PublicIpSummaryOraclePoolIdSociEgress  PublicIpSummaryOraclePoolIdEnum = "SOCI_EGRESS"
	PublicIpSummaryOraclePoolIdSociIngress PublicIpSummaryOraclePoolIdEnum = "SOCI_INGRESS"
	PublicIpSummaryOraclePoolIdOracleDev   PublicIpSummaryOraclePoolIdEnum = "ORACLE_DEV"
)

var mappingPublicIpSummaryOraclePoolIdEnum = map[string]PublicIpSummaryOraclePoolIdEnum{
	"DEFAULT":      PublicIpSummaryOraclePoolIdDefault,
	"SOCI_EGRESS":  PublicIpSummaryOraclePoolIdSociEgress,
	"SOCI_INGRESS": PublicIpSummaryOraclePoolIdSociIngress,
	"ORACLE_DEV":   PublicIpSummaryOraclePoolIdOracleDev,
}

var mappingPublicIpSummaryOraclePoolIdEnumLowerCase = map[string]PublicIpSummaryOraclePoolIdEnum{
	"default":      PublicIpSummaryOraclePoolIdDefault,
	"soci_egress":  PublicIpSummaryOraclePoolIdSociEgress,
	"soci_ingress": PublicIpSummaryOraclePoolIdSociIngress,
	"oracle_dev":   PublicIpSummaryOraclePoolIdOracleDev,
}

// GetPublicIpSummaryOraclePoolIdEnumValues Enumerates the set of values for PublicIpSummaryOraclePoolIdEnum
func GetPublicIpSummaryOraclePoolIdEnumValues() []PublicIpSummaryOraclePoolIdEnum {
	values := make([]PublicIpSummaryOraclePoolIdEnum, 0)
	for _, v := range mappingPublicIpSummaryOraclePoolIdEnum {
		values = append(values, v)
	}
	return values
}

// GetPublicIpSummaryOraclePoolIdEnumStringValues Enumerates the set of values in String for PublicIpSummaryOraclePoolIdEnum
func GetPublicIpSummaryOraclePoolIdEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"SOCI_EGRESS",
		"SOCI_INGRESS",
		"ORACLE_DEV",
	}
}

// GetMappingPublicIpSummaryOraclePoolIdEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublicIpSummaryOraclePoolIdEnum(val string) (PublicIpSummaryOraclePoolIdEnum, bool) {
	enum, ok := mappingPublicIpSummaryOraclePoolIdEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
