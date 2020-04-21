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

// CustomProtectionRuleSetting The OCID and action of a custom protection rule.
type CustomProtectionRuleSetting struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the custom protection rule.
	Id *string `mandatory:"false" json:"id"`

	// The action to take when the custom protection rule is triggered.
	// `DETECT` - Logs the request when the criteria of the custom protection rule are met. `BLOCK` - Blocks the request when the criteria of the custom protection rule are met.
	Action CustomProtectionRuleSettingActionEnum `mandatory:"false" json:"action,omitempty"`

	Exclusions []ProtectionRuleExclusion `mandatory:"false" json:"exclusions"`
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
