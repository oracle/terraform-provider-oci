// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WaasPolicyCustomProtectionRuleSummary Summary information about a Custom Protection rule.
type WaasPolicyCustomProtectionRuleSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Custom Protection rule.
	Id *string `mandatory:"false" json:"id"`

	// The user-friendly name of the Custom Protection rule.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The action to take when the Custom Protection rule is triggered.
	Action WaasPolicyCustomProtectionRuleSummaryActionEnum `mandatory:"false" json:"action,omitempty"`

	// The list of the ModSecurity rule IDs that apply to this protection rule. For more information about ModSecurity's open source WAF rules, see Mod Security's documentation (https://www.modsecurity.org/CRS/Documentation/index.html).
	ModSecurityRuleIds []string `mandatory:"false" json:"modSecurityRuleIds"`
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
