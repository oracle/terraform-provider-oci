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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RequestProtection Module that allows to enable OCI-managed protection capabilities for incoming HTTP requests.
type RequestProtection struct {

	// Ordered list of ProtectionRules. Rules are executed in order of appearance in this array.
	// ProtectionRules in this array can only use protection Capabilities of REQUEST_PROTECTION_CAPABILITY type.
	Rules []ProtectionRule `mandatory:"false" json:"rules"`

	// Maximum size of inspected HTTP message body in bytes. Actions to take if this limit is exceeded are defined in `bodyInspectionSizeLimitExceededActionName`.
	// Body inspection maximum size allowed is defined with per-tenancy limit: 8192 bytes.
	BodyInspectionSizeLimitInBytes *int `mandatory:"false" json:"bodyInspectionSizeLimitInBytes"`

	// References action by name from actions defined in WebAppFirewallPolicy. Executed if HTTP message
	// body size exceeds limit set in field `bodyInspectionSizeLimitInBytes`.
	// If this field is `null` HTTP message body will inspected up to `bodyInspectionSizeLimitInBytes` and the rest
	// will not be inspected by Protection Capabilities.
	// Allowed action types:
	// * **RETURN_HTTP_RESPONSE** terminates further execution of modules and rules and returns defined HTTP response.
	BodyInspectionSizeLimitExceededActionName *string `mandatory:"false" json:"bodyInspectionSizeLimitExceededActionName"`
}

func (m RequestProtection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestProtection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
