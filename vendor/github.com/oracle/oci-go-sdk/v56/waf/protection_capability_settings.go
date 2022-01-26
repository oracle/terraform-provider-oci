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

// ProtectionCapabilitySettings Settings for protection capabilities
type ProtectionCapabilitySettings struct {

	// Maximum number of arguments allowed. Used in protection capability 920380: Number of Arguments Limits.
	MaxNumberOfArguments *int `mandatory:"false" json:"maxNumberOfArguments"`

	// Maximum allowed length of a single argument. Used in protection capability 920370: Limit argument value length.
	MaxSingleArgumentLength *int `mandatory:"false" json:"maxSingleArgumentLength"`

	// Maximum allowed total length of all arguments. Used in protection capability 920390: Limit arguments total length.
	MaxTotalArgumentLength *int `mandatory:"false" json:"maxTotalArgumentLength"`

	// Maximum number of headers allowed in an HTTP request. Used in protection capability 9200014: Limit Number of Request Headers.
	MaxHttpRequestHeaders *int `mandatory:"false" json:"maxHttpRequestHeaders"`

	// Maximum allowed length of headers in an HTTP request. Used in protection capability: 9200024: Limit length of request header size.
	MaxHttpRequestHeaderLength *int `mandatory:"false" json:"maxHttpRequestHeaderLength"`

	// List of allowed HTTP methods. Each value as a RFC7230 formated token string.
	// Used in protection capability 911100: Restrict HTTP Request Methods.
	AllowedHttpMethods []string `mandatory:"false" json:"allowedHttpMethods"`
}

func (m ProtectionCapabilitySettings) String() string {
	return common.PointerString(m)
}
