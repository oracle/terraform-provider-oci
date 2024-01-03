// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProtectionRule Rule that represents Request/Response Protection.
// Only actions of the following types are allowed to be referenced in this rule:
//   - CHECK
//   - RETURN_HTTP_RESPONSE
type ProtectionRule struct {

	// Rule name. Must be unique within the module.
	Name *string `mandatory:"true" json:"name"`

	// References action by name from actions defined in WebAppFirewallPolicy.
	ActionName *string `mandatory:"true" json:"actionName"`

	// An ordered list that references OCI-managed protection capabilities.
	// Referenced protection capabilities are not necessarily executed in order of appearance. Their execution order
	// is decided at runtime for improved performance.
	// The array cannot contain entries with the same pair of capability key and version more than once.
	ProtectionCapabilities []ProtectionCapability `mandatory:"true" json:"protectionCapabilities"`

	// An expression that determines whether or not the rule action should be executed.
	Condition *string `mandatory:"false" json:"condition"`

	ProtectionCapabilitySettings *ProtectionCapabilitySettings `mandatory:"false" json:"protectionCapabilitySettings"`

	// Enables/disables body inspection for this protection rule.
	// Only Protection Rules in RequestProtection can have this option enabled. Response body inspection will
	// be available at a later date.
	IsBodyInspectionEnabled *bool `mandatory:"false" json:"isBodyInspectionEnabled"`

	// The language used to parse condition from field `condition`. Available languages:
	// * **JMESPATH** an extended JMESPath language syntax.
	ConditionLanguage WebAppFirewallPolicyRuleConditionLanguageEnum `mandatory:"false" json:"conditionLanguage,omitempty"`
}

// GetName returns Name
func (m ProtectionRule) GetName() *string {
	return m.Name
}

// GetConditionLanguage returns ConditionLanguage
func (m ProtectionRule) GetConditionLanguage() WebAppFirewallPolicyRuleConditionLanguageEnum {
	return m.ConditionLanguage
}

// GetCondition returns Condition
func (m ProtectionRule) GetCondition() *string {
	return m.Condition
}

// GetActionName returns ActionName
func (m ProtectionRule) GetActionName() *string {
	return m.ActionName
}

func (m ProtectionRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProtectionRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWebAppFirewallPolicyRuleConditionLanguageEnum(string(m.ConditionLanguage)); !ok && m.ConditionLanguage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionLanguage: %s. Supported values are: %s.", m.ConditionLanguage, strings.Join(GetWebAppFirewallPolicyRuleConditionLanguageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ProtectionRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeProtectionRule ProtectionRule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeProtectionRule
	}{
		"PROTECTION",
		(MarshalTypeProtectionRule)(m),
	}

	return json.Marshal(&s)
}
