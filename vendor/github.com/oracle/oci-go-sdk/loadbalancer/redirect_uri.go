// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RedirectUri An object that defines the redirect URI applied to the original request. The object property values compose the
// redirect URI.
// **NOTE:** The Load Balancing service cannot automatically detect or avoid infinite redirects. Be sure to provide
// meaningful, complete, and correct field values. If any component field of this object has no value, the system
// retains the value from the incoming HTTP request URI.
// For example, if you specify only the protocol field `https`, and the incoming request URI is
// `http://example.com:8080`, the resulting runtime redirect URI is `https://example.com:8080`. The system retains
// the host and port from the incoming URI and does not automatically change the port setting from `8080` to `443`.
// Be sure to configure valid percent-encoding (URL encoding) when needed.
// In addition to static string values, you can use the following tokens to construct the redirect URI. These tokens
// extract values from the incoming HTTP request URI.
// *  {protocol} : The protocol from the incoming HTTP request URI.
// *  {host}     : The domain name from the incoming HTTP request URI.
// *  {port}     : The port from the incoming HTTP request URI.
// *  {path}     : The path from the incoming HTTP request URI.
// *  {query}    : The query string from the incoming HTTP request URI.
// The tokens are case sensitive. For example, `{host}` is a valid token, but `{HOST}` is not.
// You can retain the literal characters of a token when you specify values for the path and query properties of the
// redirect URI. Use a backslash (\\) as the escape character for the \\, {, and } characters. For example, if the
// incoming HTTP request URI is `/video`, the path property value:
// `/example{path}123\{path\}`
// appears in the constructed redirect URI as:
// `/example/video123{path}`
type RedirectUri struct {

	// The HTTP protocol to use in the redirect URI.
	// When this value is null, not set, or set to `{protocol}`, the service preserves the original protocol from the
	// incoming HTTP request URI. Allowed values are:
	// *  HTTP
	// *  HTTPS
	// *  {protocol}
	// `{protocol}` is the only valid token for this property. It can appear only once in the value string.
	// Example: `HTTPS`
	Protocol *string `mandatory:"false" json:"protocol"`

	// The valid domain name (hostname) or IP address to use in the redirect URI.
	// When this value is null, not set, or set to `{host}`, the service preserves the original domain name from the
	// incoming HTTP request URI.
	// All RedirectUri tokens are valid for this property. You can use any token more than once.
	// Curly braces are valid in this property only to surround tokens, such as `{host}`
	// Examples:
	// *  **example.com** appears as `example.com` in the redirect URI.
	// *  **in{host}** appears as `inexample.com` in the redirect URI if `example.com` is the hostname in the
	//    incoming HTTP request URI.
	// *  **{port}{host}** appears as `8081example.com` in the redirect URI if `example.com` is the hostname and
	//    the port is `8081` in the incoming HTTP request URI.
	Host *string `mandatory:"false" json:"host"`

	// The communication port to use in the redirect URI.
	// Valid values include integers from 1 to 65535.
	// When this value is null, the service preserves the original port from the incoming HTTP request URI.
	// Example: `8081`
	Port *int `mandatory:"false" json:"port"`

	// The HTTP URI path to use in the redirect URI.
	// When this value is null, not set, or set to `{path}`, the service preserves the original path from the
	// incoming HTTP request URI. To omit the path from the redirect URI, set this value to an empty string, "".
	// All RedirectUri tokens are valid for this property. You can use any token more than once.
	// The path string must begin with `/` if it does not begin with the `{path}` token.
	// Examples:
	// *  __/example/video/123__ appears as `/example/video/123` in the redirect URI.
	// *  __/example{path}__ appears as `/example/video/123` in the redirect URI if `/video/123` is the path in the
	//    incoming HTTP request URI.
	// *  __{path}/123__ appears as `/example/video/123` in the redirect URI if `/example/video` is the path in the
	//    incoming HTTP request URI.
	// *  __{path}123__ appears as `/example/video123` in the redirect URI if `/example/video` is the path in the
	//    incoming HTTP request URI.
	// *  __/{host}/123__ appears as `/example.com/123` in the redirect URI if `example.com` is the hostname
	//    in the incoming HTTP request URI.
	// *  __/{host}/{port}__ appears as `/example.com/123` in the redirect URI if `example.com` is the hostname and
	//    `123` is the port in the incoming HTTP request URI.
	// *  __/{query}__ appears as `/lang=en` in the redirect URI if the query is `lang=en` in the incoming HTTP
	//    request URI.
	Path *string `mandatory:"false" json:"path"`

	// The query string to use in the redirect URI.
	// When this value is null, not set, or set to `{query}`, the service preserves the original query parameters
	// from the incoming HTTP request URI.
	// All `RedirectUri` tokens are valid for this property. You can use any token more than once.
	// If the query string does not begin with the `{query}` token, it must begin with the question mark (?) character.
	// You can specify multiple query parameters as a single string. Separate each query parameter with an ampersand
	// (&) character. To omit all incoming query parameters from the redirect URI, set this value to an empty
	// string, "".
	// If the specified query string results in a redirect URI ending with `?` or `&`, the last character is truncated.
	// For example, if the incoming URI is `http://host.com:8080/documents` and the query property value is
	// `?lang=en&{query}`, the redirect URI is `http://host.com:8080/documents?lang=en`. The system
	// truncates the final ampersand (&) because the incoming URI included no value to replace the {query} token.
	// Examples:
	// * **lang=en&time_zone=PST** appears as `lang=en&time_zone=PST` in the redirect URI.
	// * **{query}** appears as `lang=en&time_zone=PST` in the redirect URI if `lang=en&time_zone=PST` is the query
	//   string in the incoming HTTP request. If the incoming HTTP request has no query parameters, the `{query}`
	//   token renders as an empty string.
	// * **lang=en&{query}&time_zone=PST** appears as `lang=en&country=us&time_zone=PST` in the redirect URI if
	//   `country=us` is the query string in the incoming HTTP request. If the incoming HTTP request has no query
	//   parameters, this value renders as `lang=en&time_zone=PST`.
	// *  **protocol={protocol}&hostname={host}** appears as `protocol=http&hostname=example.com` in the redirect
	//    URI if the protocol is `HTTP` and the hostname is `example.com` in the incoming HTTP request.
	// *  **port={port}&hostname={host}** appears as `port=8080&hostname=example.com` in the redirect URI if the
	//    port is `8080` and the hostname is `example.com` in the incoming HTTP request URI.
	Query *string `mandatory:"false" json:"query"`
}

func (m RedirectUri) String() string {
	return common.PointerString(m)
}
