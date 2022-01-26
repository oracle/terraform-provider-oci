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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// RequestAccessControl Module that allows inspection of HTTP request properties and to return a defined HTTP response.
// In this module, rules with the name 'Default Action' are not allowed, since this name is reserved for default action logs.
type RequestAccessControl struct {

	// References an default Action to take if no AccessControlRule was matched. Allowed action types:
	// * **ALLOW** continues execution of other modules and their rules.
	// * **RETURN_HTTP_RESPONSE** terminates further execution of modules and rules and returns defined HTTP response.
	DefaultActionName *string `mandatory:"true" json:"defaultActionName"`

	// Ordered list of AccessControlRules. Rules are executed in order of appearance in this array.
	Rules []AccessControlRule `mandatory:"false" json:"rules"`
}

func (m RequestAccessControl) String() string {
	return common.PointerString(m)
}
