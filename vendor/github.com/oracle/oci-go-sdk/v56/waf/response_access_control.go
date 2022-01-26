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

// ResponseAccessControl Module that allows inspection of HTTP response properties and to return a defined HTTP response.
type ResponseAccessControl struct {

	// Ordered list of AccessControlRules. Rules are executed in order of appearance in this array.
	Rules []AccessControlRule `mandatory:"false" json:"rules"`
}

func (m ResponseAccessControl) String() string {
	return common.PointerString(m)
}
