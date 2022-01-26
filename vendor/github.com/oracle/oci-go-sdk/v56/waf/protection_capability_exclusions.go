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

// ProtectionCapabilityExclusions Identifies specific HTTP message parameters to exclude from inspection by a protection capability.
type ProtectionCapabilityExclusions struct {

	// List of HTTP request cookie values (by cookie name) to exclude from inspecting.
	// Example: If we have cookie 'cookieName=cookieValue' and requestCookies=['cookieName'], both 'cookieName' and 'cookieValue' will not be inspected.
	RequestCookies []string `mandatory:"false" json:"requestCookies"`

	// List of URL query parameter values from form-urlencoded XML, JSON, AMP, or POST payloads to exclude from inspecting.
	// Example: If we have query parameter 'argumentName=argumentValue' and args=['argumentName'], both 'argumentName' and 'argumentValue' will not be inspected.
	Args []string `mandatory:"false" json:"args"`
}

func (m ProtectionCapabilityExclusions) String() string {
	return common.PointerString(m)
}
