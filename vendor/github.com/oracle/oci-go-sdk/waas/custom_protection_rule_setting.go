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

// CustomProtectionRuleSetting The OCID and action of a Custom Protection rule.
type CustomProtectionRuleSetting struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Custom Protecion rule.
	Id *string `mandatory:"false" json:"id"`

	// The action to take when the Custom Protection rule is triggered.
	Action CustomProtectionRuleSettingActionEnum `mandatory:"false" json:"action,omitempty"`
}

func (m CustomProtectionRuleSetting) String() string {
	return common.PointerString(m)
}

// CustomProtectionRuleSettingActionEnum Enum with underlying type: string
type CustomProtectionRuleSettingActionEnum string

// Set of constants representing the allowable values for CustomProtectionRuleSettingActionEnum
const (
	CustomProtectionRuleSettingActionDetect CustomProtectionRuleSettingActionEnum = "DETECT"
	CustomProtectionRuleSettingActionBlock  CustomProtectionRuleSettingActionEnum = "BLOCK"
)

var mappingCustomProtectionRuleSettingAction = map[string]CustomProtectionRuleSettingActionEnum{
	"DETECT": CustomProtectionRuleSettingActionDetect,
	"BLOCK":  CustomProtectionRuleSettingActionBlock,
}

// GetCustomProtectionRuleSettingActionEnumValues Enumerates the set of values for CustomProtectionRuleSettingActionEnum
func GetCustomProtectionRuleSettingActionEnumValues() []CustomProtectionRuleSettingActionEnum {
	values := make([]CustomProtectionRuleSettingActionEnum, 0)
	for _, v := range mappingCustomProtectionRuleSettingAction {
		values = append(values, v)
	}
	return values
}
