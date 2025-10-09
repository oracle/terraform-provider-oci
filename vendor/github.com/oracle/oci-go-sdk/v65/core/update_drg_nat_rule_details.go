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

// UpdateDrgNatRuleDetails DRG NAT rules to update in the DRG NAT policy.
type UpdateDrgNatRuleDetails struct {

	// The Oracle-assigned ID of the DRG NAT rule.
	Id *string `mandatory:"true" json:"id"`

	// The priority associated with each DRG NAT rule.
	DrgNatRulePriority *int64 `mandatory:"false" json:"drgNatRulePriority"`

	// Indicates whether the DRG NAT rule is either stateful or stateless.
	DrgNatRuleType UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum `mandatory:"false" json:"drgNatRuleType,omitempty"`

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

func (m UpdateDrgNatRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDrgNatRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateDrgNatRuleDetailsDrgNatRuleTypeEnum(string(m.DrgNatRuleType)); !ok && m.DrgNatRuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgNatRuleType: %s. Supported values are: %s.", m.DrgNatRuleType, strings.Join(GetUpdateDrgNatRuleDetailsDrgNatRuleTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum Enum with underlying type: string
type UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum string

// Set of constants representing the allowable values for UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum
const (
	UpdateDrgNatRuleDetailsDrgNatRuleTypeStateful  UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum = "STATEFUL"
	UpdateDrgNatRuleDetailsDrgNatRuleTypeStateless UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum = "STATELESS"
)

var mappingUpdateDrgNatRuleDetailsDrgNatRuleTypeEnum = map[string]UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum{
	"STATEFUL":  UpdateDrgNatRuleDetailsDrgNatRuleTypeStateful,
	"STATELESS": UpdateDrgNatRuleDetailsDrgNatRuleTypeStateless,
}

var mappingUpdateDrgNatRuleDetailsDrgNatRuleTypeEnumLowerCase = map[string]UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum{
	"stateful":  UpdateDrgNatRuleDetailsDrgNatRuleTypeStateful,
	"stateless": UpdateDrgNatRuleDetailsDrgNatRuleTypeStateless,
}

// GetUpdateDrgNatRuleDetailsDrgNatRuleTypeEnumValues Enumerates the set of values for UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum
func GetUpdateDrgNatRuleDetailsDrgNatRuleTypeEnumValues() []UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum {
	values := make([]UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum, 0)
	for _, v := range mappingUpdateDrgNatRuleDetailsDrgNatRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDrgNatRuleDetailsDrgNatRuleTypeEnumStringValues Enumerates the set of values in String for UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum
func GetUpdateDrgNatRuleDetailsDrgNatRuleTypeEnumStringValues() []string {
	return []string{
		"STATEFUL",
		"STATELESS",
	}
}

// GetMappingUpdateDrgNatRuleDetailsDrgNatRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDrgNatRuleDetailsDrgNatRuleTypeEnum(val string) (UpdateDrgNatRuleDetailsDrgNatRuleTypeEnum, bool) {
	enum, ok := mappingUpdateDrgNatRuleDetailsDrgNatRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
