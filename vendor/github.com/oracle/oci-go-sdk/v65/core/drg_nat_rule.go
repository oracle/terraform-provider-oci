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

// DrgNatRule A DRG NAT rule specifies the mapping between an original source/destination IP and a translated source/destination IP. The source needs to be described in the context of packet entering the DRG.
type DrgNatRule struct {

	// The Oracle-assigned ID of the DRG NAT rule.
	Id *string `mandatory:"true" json:"id"`

	// The DRG NAT policy's Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	DrgNatPolicyId *string `mandatory:"true" json:"drgNatPolicyId"`

	// The priority associated with each DRG NAT rule.
	DrgNatRulePriority *int64 `mandatory:"true" json:"drgNatRulePriority"`

	// Indicates whether the DRG NAT rule is stateful or stateless.
	DrgNatRuleType DrgNatRuleDrgNatRuleTypeEnum `mandatory:"true" json:"drgNatRuleType"`

	// Represents the range of IP addresses to match against when routing traffic.
	// Original CIDR range for Source NAT.
	// Potential values:
	//   * An IPv4 address range in CIDR notation. For example: `192.168.1.0/24`.
	OriginalSource *string `mandatory:"false" json:"originalSource"`

	// Represents the range of IP addresses to match against when routing traffic.
	// Translated CIDR range for Source NAT.
	// Potential values:
	//   * An IPv4 address range in CIDR notation. For example: `192.168.1.0/24`.
	TranslatedSource *string `mandatory:"false" json:"translatedSource"`

	// Represents the range of IP addresses to match against when routing traffic.
	// Original CIDR range for Destination NAT.
	// Potential values:
	//   * An IPv4 address range in CIDR notation. For example: `192.168.1.0/24`.
	OriginalDestination *string `mandatory:"false" json:"originalDestination"`

	// Represents the range of IP addresses to match against when routing traffic.
	// Translated CIDR range for Destination NAT.
	// Potential values:
	//   * An IPv4 address range in CIDR notation. For example: `192.168.1.0/24`.
	TranslatedDestination *string `mandatory:"false" json:"translatedDestination"`
}

func (m DrgNatRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgNatRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrgNatRuleDrgNatRuleTypeEnum(string(m.DrgNatRuleType)); !ok && m.DrgNatRuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgNatRuleType: %s. Supported values are: %s.", m.DrgNatRuleType, strings.Join(GetDrgNatRuleDrgNatRuleTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrgNatRuleDrgNatRuleTypeEnum Enum with underlying type: string
type DrgNatRuleDrgNatRuleTypeEnum string

// Set of constants representing the allowable values for DrgNatRuleDrgNatRuleTypeEnum
const (
	DrgNatRuleDrgNatRuleTypeStateful  DrgNatRuleDrgNatRuleTypeEnum = "STATEFUL"
	DrgNatRuleDrgNatRuleTypeStateless DrgNatRuleDrgNatRuleTypeEnum = "STATELESS"
)

var mappingDrgNatRuleDrgNatRuleTypeEnum = map[string]DrgNatRuleDrgNatRuleTypeEnum{
	"STATEFUL":  DrgNatRuleDrgNatRuleTypeStateful,
	"STATELESS": DrgNatRuleDrgNatRuleTypeStateless,
}

var mappingDrgNatRuleDrgNatRuleTypeEnumLowerCase = map[string]DrgNatRuleDrgNatRuleTypeEnum{
	"stateful":  DrgNatRuleDrgNatRuleTypeStateful,
	"stateless": DrgNatRuleDrgNatRuleTypeStateless,
}

// GetDrgNatRuleDrgNatRuleTypeEnumValues Enumerates the set of values for DrgNatRuleDrgNatRuleTypeEnum
func GetDrgNatRuleDrgNatRuleTypeEnumValues() []DrgNatRuleDrgNatRuleTypeEnum {
	values := make([]DrgNatRuleDrgNatRuleTypeEnum, 0)
	for _, v := range mappingDrgNatRuleDrgNatRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgNatRuleDrgNatRuleTypeEnumStringValues Enumerates the set of values in String for DrgNatRuleDrgNatRuleTypeEnum
func GetDrgNatRuleDrgNatRuleTypeEnumStringValues() []string {
	return []string{
		"STATEFUL",
		"STATELESS",
	}
}

// GetMappingDrgNatRuleDrgNatRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgNatRuleDrgNatRuleTypeEnum(val string) (DrgNatRuleDrgNatRuleTypeEnum, bool) {
	enum, ok := mappingDrgNatRuleDrgNatRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
