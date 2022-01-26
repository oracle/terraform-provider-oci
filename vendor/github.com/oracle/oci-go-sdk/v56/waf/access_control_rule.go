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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AccessControlRule Rule that represents Request/Response Access Control.
// Only actions of the following types are allowed to be referenced in this rule:
//  * CHECK
//  * ALLOW
//  * RETURN_HTTP_RESPONSE
type AccessControlRule struct {

	// Rule name. Must be unique within the module.
	Name *string `mandatory:"true" json:"name"`

	// References action by name from actions defined in WebAppFirewallPolicy.
	ActionName *string `mandatory:"true" json:"actionName"`

	// An expression that determines whether or not the rule action should be executed.
	Condition *string `mandatory:"false" json:"condition"`

	// The language used to parse condition from field `condition`. Available languages:
	// * **JMESPATH** an extended JMESPath language syntax.
	ConditionLanguage WebAppFirewallPolicyRuleConditionLanguageEnum `mandatory:"false" json:"conditionLanguage,omitempty"`
}

//GetName returns Name
func (m AccessControlRule) GetName() *string {
	return m.Name
}

//GetConditionLanguage returns ConditionLanguage
func (m AccessControlRule) GetConditionLanguage() WebAppFirewallPolicyRuleConditionLanguageEnum {
	return m.ConditionLanguage
}

//GetCondition returns Condition
func (m AccessControlRule) GetCondition() *string {
	return m.Condition
}

//GetActionName returns ActionName
func (m AccessControlRule) GetActionName() *string {
	return m.ActionName
}

func (m AccessControlRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AccessControlRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAccessControlRule AccessControlRule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAccessControlRule
	}{
		"ACCESS_CONTROL",
		(MarshalTypeAccessControlRule)(m),
	}

	return json.Marshal(&s)
}
