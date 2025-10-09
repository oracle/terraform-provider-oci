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

// AddDrgNatRuleDetails Details used to add a DRG NAT rule.
type AddDrgNatRuleDetails struct {

	// The priority associated with each DRG NAT rule.
	DrgNatRulePriority *int64 `mandatory:"true" json:"drgNatRulePriority"`

	// Indicates whether the DRG NAT rule is stateful or stateless.
	DrgNatRuleType AddDrgNatRuleDetailsDrgNatRuleTypeEnum `mandatory:"true" json:"drgNatRuleType"`

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

func (m AddDrgNatRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddDrgNatRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddDrgNatRuleDetailsDrgNatRuleTypeEnum(string(m.DrgNatRuleType)); !ok && m.DrgNatRuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgNatRuleType: %s. Supported values are: %s.", m.DrgNatRuleType, strings.Join(GetAddDrgNatRuleDetailsDrgNatRuleTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddDrgNatRuleDetailsDrgNatRuleTypeEnum Enum with underlying type: string
type AddDrgNatRuleDetailsDrgNatRuleTypeEnum string

// Set of constants representing the allowable values for AddDrgNatRuleDetailsDrgNatRuleTypeEnum
const (
	AddDrgNatRuleDetailsDrgNatRuleTypeStateful  AddDrgNatRuleDetailsDrgNatRuleTypeEnum = "STATEFUL"
	AddDrgNatRuleDetailsDrgNatRuleTypeStateless AddDrgNatRuleDetailsDrgNatRuleTypeEnum = "STATELESS"
)

var mappingAddDrgNatRuleDetailsDrgNatRuleTypeEnum = map[string]AddDrgNatRuleDetailsDrgNatRuleTypeEnum{
	"STATEFUL":  AddDrgNatRuleDetailsDrgNatRuleTypeStateful,
	"STATELESS": AddDrgNatRuleDetailsDrgNatRuleTypeStateless,
}

var mappingAddDrgNatRuleDetailsDrgNatRuleTypeEnumLowerCase = map[string]AddDrgNatRuleDetailsDrgNatRuleTypeEnum{
	"stateful":  AddDrgNatRuleDetailsDrgNatRuleTypeStateful,
	"stateless": AddDrgNatRuleDetailsDrgNatRuleTypeStateless,
}

// GetAddDrgNatRuleDetailsDrgNatRuleTypeEnumValues Enumerates the set of values for AddDrgNatRuleDetailsDrgNatRuleTypeEnum
func GetAddDrgNatRuleDetailsDrgNatRuleTypeEnumValues() []AddDrgNatRuleDetailsDrgNatRuleTypeEnum {
	values := make([]AddDrgNatRuleDetailsDrgNatRuleTypeEnum, 0)
	for _, v := range mappingAddDrgNatRuleDetailsDrgNatRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddDrgNatRuleDetailsDrgNatRuleTypeEnumStringValues Enumerates the set of values in String for AddDrgNatRuleDetailsDrgNatRuleTypeEnum
func GetAddDrgNatRuleDetailsDrgNatRuleTypeEnumStringValues() []string {
	return []string{
		"STATEFUL",
		"STATELESS",
	}
}

// GetMappingAddDrgNatRuleDetailsDrgNatRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddDrgNatRuleDetailsDrgNatRuleTypeEnum(val string) (AddDrgNatRuleDetailsDrgNatRuleTypeEnum, bool) {
	enum, ok := mappingAddDrgNatRuleDetailsDrgNatRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
