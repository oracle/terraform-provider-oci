// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// WebAppFirewallPolicyRule Base schema for WebAppFirewallPolicyRules, including properties common to all of them.
type WebAppFirewallPolicyRule interface {

	// Rule name. Must be unique within the module.
	GetName() *string

	// References action by name from actions defined in WebAppFirewallPolicy.
	GetActionName() *string

	// The language used to parse condition from field `condition`. Available languages:
	// * **JMESPATH** an extended JMESPath language syntax.
	GetConditionLanguage() WebAppFirewallPolicyRuleConditionLanguageEnum

	// An expression that determines whether or not the rule action should be executed.
	GetCondition() *string
}

type webappfirewallpolicyrule struct {
	JsonData          []byte
	Name              *string                                       `mandatory:"true" json:"name"`
	ActionName        *string                                       `mandatory:"true" json:"actionName"`
	ConditionLanguage WebAppFirewallPolicyRuleConditionLanguageEnum `mandatory:"false" json:"conditionLanguage,omitempty"`
	Condition         *string                                       `mandatory:"false" json:"condition"`
	Type              string                                        `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *webappfirewallpolicyrule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerwebappfirewallpolicyrule webappfirewallpolicyrule
	s := struct {
		Model Unmarshalerwebappfirewallpolicyrule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.ActionName = s.Model.ActionName
	m.ConditionLanguage = s.Model.ConditionLanguage
	m.Condition = s.Model.Condition
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *webappfirewallpolicyrule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "PROTECTION":
		mm := ProtectionRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REQUEST_RATE_LIMITING":
		mm := RequestRateLimitingRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ACCESS_CONTROL":
		mm := AccessControlRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m webappfirewallpolicyrule) GetName() *string {
	return m.Name
}

//GetActionName returns ActionName
func (m webappfirewallpolicyrule) GetActionName() *string {
	return m.ActionName
}

//GetConditionLanguage returns ConditionLanguage
func (m webappfirewallpolicyrule) GetConditionLanguage() WebAppFirewallPolicyRuleConditionLanguageEnum {
	return m.ConditionLanguage
}

//GetCondition returns Condition
func (m webappfirewallpolicyrule) GetCondition() *string {
	return m.Condition
}

func (m webappfirewallpolicyrule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m webappfirewallpolicyrule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWebAppFirewallPolicyRuleConditionLanguageEnum(string(m.ConditionLanguage)); !ok && m.ConditionLanguage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionLanguage: %s. Supported values are: %s.", m.ConditionLanguage, strings.Join(GetWebAppFirewallPolicyRuleConditionLanguageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WebAppFirewallPolicyRuleConditionLanguageEnum Enum with underlying type: string
type WebAppFirewallPolicyRuleConditionLanguageEnum string

// Set of constants representing the allowable values for WebAppFirewallPolicyRuleConditionLanguageEnum
const (
	WebAppFirewallPolicyRuleConditionLanguageJmespath WebAppFirewallPolicyRuleConditionLanguageEnum = "JMESPATH"
)

var mappingWebAppFirewallPolicyRuleConditionLanguageEnum = map[string]WebAppFirewallPolicyRuleConditionLanguageEnum{
	"JMESPATH": WebAppFirewallPolicyRuleConditionLanguageJmespath,
}

// GetWebAppFirewallPolicyRuleConditionLanguageEnumValues Enumerates the set of values for WebAppFirewallPolicyRuleConditionLanguageEnum
func GetWebAppFirewallPolicyRuleConditionLanguageEnumValues() []WebAppFirewallPolicyRuleConditionLanguageEnum {
	values := make([]WebAppFirewallPolicyRuleConditionLanguageEnum, 0)
	for _, v := range mappingWebAppFirewallPolicyRuleConditionLanguageEnum {
		values = append(values, v)
	}
	return values
}

// GetWebAppFirewallPolicyRuleConditionLanguageEnumStringValues Enumerates the set of values in String for WebAppFirewallPolicyRuleConditionLanguageEnum
func GetWebAppFirewallPolicyRuleConditionLanguageEnumStringValues() []string {
	return []string{
		"JMESPATH",
	}
}

// GetMappingWebAppFirewallPolicyRuleConditionLanguageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWebAppFirewallPolicyRuleConditionLanguageEnum(val string) (WebAppFirewallPolicyRuleConditionLanguageEnum, bool) {
	mappingWebAppFirewallPolicyRuleConditionLanguageEnumIgnoreCase := make(map[string]WebAppFirewallPolicyRuleConditionLanguageEnum)
	for k, v := range mappingWebAppFirewallPolicyRuleConditionLanguageEnum {
		mappingWebAppFirewallPolicyRuleConditionLanguageEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWebAppFirewallPolicyRuleConditionLanguageEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// WebAppFirewallPolicyRuleTypeEnum Enum with underlying type: string
type WebAppFirewallPolicyRuleTypeEnum string

// Set of constants representing the allowable values for WebAppFirewallPolicyRuleTypeEnum
const (
	WebAppFirewallPolicyRuleTypeAccessControl       WebAppFirewallPolicyRuleTypeEnum = "ACCESS_CONTROL"
	WebAppFirewallPolicyRuleTypeProtection          WebAppFirewallPolicyRuleTypeEnum = "PROTECTION"
	WebAppFirewallPolicyRuleTypeRequestRateLimiting WebAppFirewallPolicyRuleTypeEnum = "REQUEST_RATE_LIMITING"
)

var mappingWebAppFirewallPolicyRuleTypeEnum = map[string]WebAppFirewallPolicyRuleTypeEnum{
	"ACCESS_CONTROL":        WebAppFirewallPolicyRuleTypeAccessControl,
	"PROTECTION":            WebAppFirewallPolicyRuleTypeProtection,
	"REQUEST_RATE_LIMITING": WebAppFirewallPolicyRuleTypeRequestRateLimiting,
}

// GetWebAppFirewallPolicyRuleTypeEnumValues Enumerates the set of values for WebAppFirewallPolicyRuleTypeEnum
func GetWebAppFirewallPolicyRuleTypeEnumValues() []WebAppFirewallPolicyRuleTypeEnum {
	values := make([]WebAppFirewallPolicyRuleTypeEnum, 0)
	for _, v := range mappingWebAppFirewallPolicyRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWebAppFirewallPolicyRuleTypeEnumStringValues Enumerates the set of values in String for WebAppFirewallPolicyRuleTypeEnum
func GetWebAppFirewallPolicyRuleTypeEnumStringValues() []string {
	return []string{
		"ACCESS_CONTROL",
		"PROTECTION",
		"REQUEST_RATE_LIMITING",
	}
}

// GetMappingWebAppFirewallPolicyRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWebAppFirewallPolicyRuleTypeEnum(val string) (WebAppFirewallPolicyRuleTypeEnum, bool) {
	mappingWebAppFirewallPolicyRuleTypeEnumIgnoreCase := make(map[string]WebAppFirewallPolicyRuleTypeEnum)
	for k, v := range mappingWebAppFirewallPolicyRuleTypeEnum {
		mappingWebAppFirewallPolicyRuleTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWebAppFirewallPolicyRuleTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
