// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProtectionRule The protection rule settings. Protection rules can allow, block, or trigger an alert if a request meets the parameters of an applied rule.
type ProtectionRule struct {

	// The unique key of the protection rule.
	Key *string `mandatory:"false" json:"key"`

	// The list of the ModSecurity rule IDs that apply to this protection rule. For more information about ModSecurity's open source WAF rules, see Mod Security's documentation (https://www.modsecurity.org/CRS/Documentation/index.html).
	ModSecurityRuleIds []string `mandatory:"false" json:"modSecurityRuleIds"`

	// The name of the protection rule.
	Name *string `mandatory:"false" json:"name"`

	// The description of the protection rule.
	Description *string `mandatory:"false" json:"description"`

	// The action to take when the traffic is detected as malicious. If unspecified, defaults to `OFF`.
	Action ProtectionRuleActionEnum `mandatory:"false" json:"action,omitempty"`

	// The list of labels for the protection rule.
	// **Note:** Protection rules with a `ResponseBody` label will have no effect unless `isResponseInspected` is true.
	Labels []string `mandatory:"false" json:"labels"`

	Exclusions []ProtectionRuleExclusion `mandatory:"false" json:"exclusions"`
}

func (m ProtectionRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProtectionRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingProtectionRuleActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetProtectionRuleActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProtectionRuleActionEnum Enum with underlying type: string
type ProtectionRuleActionEnum string

// Set of constants representing the allowable values for ProtectionRuleActionEnum
const (
	ProtectionRuleActionOff    ProtectionRuleActionEnum = "OFF"
	ProtectionRuleActionDetect ProtectionRuleActionEnum = "DETECT"
	ProtectionRuleActionBlock  ProtectionRuleActionEnum = "BLOCK"
)

var mappingProtectionRuleActionEnum = map[string]ProtectionRuleActionEnum{
	"OFF":    ProtectionRuleActionOff,
	"DETECT": ProtectionRuleActionDetect,
	"BLOCK":  ProtectionRuleActionBlock,
}

var mappingProtectionRuleActionEnumLowerCase = map[string]ProtectionRuleActionEnum{
	"off":    ProtectionRuleActionOff,
	"detect": ProtectionRuleActionDetect,
	"block":  ProtectionRuleActionBlock,
}

// GetProtectionRuleActionEnumValues Enumerates the set of values for ProtectionRuleActionEnum
func GetProtectionRuleActionEnumValues() []ProtectionRuleActionEnum {
	values := make([]ProtectionRuleActionEnum, 0)
	for _, v := range mappingProtectionRuleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetProtectionRuleActionEnumStringValues Enumerates the set of values in String for ProtectionRuleActionEnum
func GetProtectionRuleActionEnumStringValues() []string {
	return []string{
		"OFF",
		"DETECT",
		"BLOCK",
	}
}

// GetMappingProtectionRuleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProtectionRuleActionEnum(val string) (ProtectionRuleActionEnum, bool) {
	enum, ok := mappingProtectionRuleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
