// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomProtectionRuleSetting) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCustomProtectionRuleSettingActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetCustomProtectionRuleSettingActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CustomProtectionRuleSettingActionEnum Enum with underlying type: string
type CustomProtectionRuleSettingActionEnum string

// Set of constants representing the allowable values for CustomProtectionRuleSettingActionEnum
const (
	CustomProtectionRuleSettingActionDetect CustomProtectionRuleSettingActionEnum = "DETECT"
	CustomProtectionRuleSettingActionBlock  CustomProtectionRuleSettingActionEnum = "BLOCK"
)

var mappingCustomProtectionRuleSettingActionEnum = map[string]CustomProtectionRuleSettingActionEnum{
	"DETECT": CustomProtectionRuleSettingActionDetect,
	"BLOCK":  CustomProtectionRuleSettingActionBlock,
}

// GetCustomProtectionRuleSettingActionEnumValues Enumerates the set of values for CustomProtectionRuleSettingActionEnum
func GetCustomProtectionRuleSettingActionEnumValues() []CustomProtectionRuleSettingActionEnum {
	values := make([]CustomProtectionRuleSettingActionEnum, 0)
	for _, v := range mappingCustomProtectionRuleSettingActionEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomProtectionRuleSettingActionEnumStringValues Enumerates the set of values in String for CustomProtectionRuleSettingActionEnum
func GetCustomProtectionRuleSettingActionEnumStringValues() []string {
	return []string{
		"DETECT",
		"BLOCK",
	}
}

// GetMappingCustomProtectionRuleSettingActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomProtectionRuleSettingActionEnum(val string) (CustomProtectionRuleSettingActionEnum, bool) {
	mappingCustomProtectionRuleSettingActionEnumIgnoreCase := make(map[string]CustomProtectionRuleSettingActionEnum)
	for k, v := range mappingCustomProtectionRuleSettingActionEnum {
		mappingCustomProtectionRuleSettingActionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCustomProtectionRuleSettingActionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
