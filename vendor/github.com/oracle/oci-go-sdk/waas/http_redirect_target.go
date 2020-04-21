// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// HttpRedirectTarget The representation of HttpRedirectTarget
type HttpRedirectTarget struct {

	// The protocol used for the target, http or https.
	Protocol HttpRedirectTargetProtocolEnum `mandatory:"true" json:"protocol"`

	// The host portion of the redirect.
	Host *string `mandatory:"true" json:"host"`

	// The path component of the target URL (e.g., "/path/to/resource" in "https://target.example.com/path/to/resource?redirected"), which can be empty, static, or request-copying, or request-prefixing. Use of \ is not permitted except to escape a following \, {, or }. An empty value is treated the same as static "/". A static value must begin with a leading "/", optionally followed by other path characters. A request-copying value must exactly match "{path}", and will be replaced with the path component of the request URL (including its initial "/"). A request-prefixing value must start with "/" and end with a non-escaped "{path}", which will be replaced with the path component of the request URL (including its initial "/"). Only one such replacement token is allowed.
	Path *string `mandatory:"true" json:"path"`

	// The query component of the target URL (e.g., "?redirected" in "https://target.example.com/path/to/resource?redirected"), which can be empty, static, or request-copying. Use of \ is not permitted except to escape a following \, {, or }. An empty value results in a redirection target URL with no query component. A static value must begin with a leading "?", optionally followed by other query characters. A request-copying value must exactly match "{query}", and will be replaced with the query component of the request URL (including a leading "?" if and only if the request URL includes a query component).
	Query *string `mandatory:"true" json:"query"`

	// Port number of the target destination of the redirect, default to match protocol
	Port *int `mandatory:"false" json:"port"`
}

func (m HttpRedirectTarget) String() string {
	return common.PointerString(m)
}

// HttpRedirectTargetProtocolEnum Enum with underlying type: string
type HttpRedirectTargetProtocolEnum string

// Set of constants representing the allowable values for HttpRedirectTargetProtocolEnum
const (
	HttpRedirectTargetProtocolHttp  HttpRedirectTargetProtocolEnum = "HTTP"
	HttpRedirectTargetProtocolHttps HttpRedirectTargetProtocolEnum = "HTTPS"
)

var mappingHttpRedirectTargetProtocol = map[string]HttpRedirectTargetProtocolEnum{
	"HTTP":  HttpRedirectTargetProtocolHttp,
	"HTTPS": HttpRedirectTargetProtocolHttps,
}

// GetHttpRedirectTargetProtocolEnumValues Enumerates the set of values for HttpRedirectTargetProtocolEnum
func GetHttpRedirectTargetProtocolEnumValues() []HttpRedirectTargetProtocolEnum {
	values := make([]HttpRedirectTargetProtocolEnum, 0)
	for _, v := range mappingHttpRedirectTargetProtocol {
		values = append(values, v)
	}
	return values
}
