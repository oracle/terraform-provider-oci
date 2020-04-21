// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WaasPolicyCustomProtectionRuleSummary The OCID and action of a custom protection rule.
type WaasPolicyCustomProtectionRuleSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the custom protection rule.
	Id *string `mandatory:"false" json:"id"`

	// The user-friendly name of the custom protection rule.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The action to take when the custom protection rule is triggered.
	// `DETECT` - Logs the request when the criteria of the custom protection rule are met. `BLOCK` - Blocks the request when the criteria of the custom protection rule are met.
	Action WaasPolicyCustomProtectionRuleSummaryActionEnum `mandatory:"false" json:"action,omitempty"`

	// The list of the ModSecurity rule IDs that apply to this protection rule. For more information about ModSecurity's open source WAF rules, see Mod Security's documentation (https://www.modsecurity.org/CRS/Documentation/index.html).
	ModSecurityRuleIds []string `mandatory:"false" json:"modSecurityRuleIds"`

	Exclusions []ProtectionRuleExclusion `mandatory:"false" json:"exclusions"`
}

func (m WaasPolicyCustomProtectionRuleSummary) String() string {
	return common.PointerString(m)
}

// WaasPolicyCustomProtectionRuleSummaryActionEnum Enum with underlying type: string
type WaasPolicyCustomProtectionRuleSummaryActionEnum string

// Set of constants representing the allowable values for WaasPolicyCustomProtectionRuleSummaryActionEnum
const (
	WaasPolicyCustomProtectionRuleSummaryActionDetect WaasPolicyCustomProtectionRuleSummaryActionEnum = "DETECT"
	WaasPolicyCustomProtectionRuleSummaryActionBlock  WaasPolicyCustomProtectionRuleSummaryActionEnum = "BLOCK"
)

var mappingWaasPolicyCustomProtectionRuleSummaryAction = map[string]WaasPolicyCustomProtectionRuleSummaryActionEnum{
	"DETECT": WaasPolicyCustomProtectionRuleSummaryActionDetect,
	"BLOCK":  WaasPolicyCustomProtectionRuleSummaryActionBlock,
}

// GetWaasPolicyCustomProtectionRuleSummaryActionEnumValues Enumerates the set of values for WaasPolicyCustomProtectionRuleSummaryActionEnum
func GetWaasPolicyCustomProtectionRuleSummaryActionEnumValues() []WaasPolicyCustomProtectionRuleSummaryActionEnum {
	values := make([]WaasPolicyCustomProtectionRuleSummaryActionEnum, 0)
	for _, v := range mappingWaasPolicyCustomProtectionRuleSummaryAction {
		values = append(values, v)
	}
	return values
}
