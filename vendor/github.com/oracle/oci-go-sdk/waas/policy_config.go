// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// PolicyConfig The configuration details for the WAAS policy.
type PolicyConfig struct {

	// The OCID of the SSL certificate to use if HTTPS is supported.
	CertificateId *string `mandatory:"false" json:"certificateId"`

	// Enable or disable HTTPS support. If true, a `certificateId` is required. If unspecified, defaults to `false`.
	IsHttpsEnabled *bool `mandatory:"false" json:"isHttpsEnabled"`

	// Force HTTP to HTTPS redirection. If unspecified, defaults to `false`.
	IsHttpsForced *bool `mandatory:"false" json:"isHttpsForced"`

	// A list of allowed TLS protocols. Only applicable when HTTPS support is enabled.
	// The TLS protocol is negotiated while the request is connecting and the most recent protocol supported by both the edge node and client browser will be selected. If no such version exists, the connection will be aborted.
	// - **TLS_V1:** corresponds to TLS 1.0 specification.
	// - **TLS_V1_1:** corresponds to TLS 1.1 specification.
	// - **TLS_V1_2:** corresponds to TLS 1.2 specification.
	// - **TLS_V1_3:** corresponds to TLS 1.3 specification.
	// Enabled TLS protocols must go in a row. For example if `TLS_v1_1` and `TLS_V1_3` are enabled, `TLS_V1_2` must be enabled too.
	TlsProtocols []PolicyConfigTlsProtocolsEnum `mandatory:"false" json:"tlsProtocols,omitempty"`

	// Enable or disable GZIP compression of origin responses. If enabled, the header `Accept-Encoding: gzip` is sent to origin, otherwise, the empty `Accept-Encoding:` header is used.
	IsOriginCompressionEnabled *bool `mandatory:"false" json:"isOriginCompressionEnabled"`

	// Enabling `isBehindCdn` allows for the collection of IP addresses from client requests if the WAF is connected to a CDN.
	IsBehindCdn *bool `mandatory:"false" json:"isBehindCdn"`

	// Specifies an HTTP header name which is treated as the connecting client's IP address. Applicable only if `isBehindCdn` is enabled.
	// The edge node reads this header and its value and sets the client IP address as specified. It does not create the header if the header is not present in the request. If the header is not present, the connecting IP address will be used as the client's true IP address. It uses the last IP address in the header's value as the true IP address.
	// Example: `X-Client-Ip: 11.1.1.1, 13.3.3.3`
	// In the case of multiple headers with the same name, only the first header will be used. It is assumed that CDN sets the correct client IP address to prevent spoofing.
	// - **X_FORWARDED_FOR:** Corresponds to `X-Forwarded-For` header name.
	// - **X_CLIENT_IP:** Corresponds to `X-Client-Ip` header name.
	// - **X_REAL_IP:** Corresponds to `X-Real-Ip` header name.
	// - **CLIENT_IP:** Corresponds to `Client-Ip` header name.
	// - **TRUE_CLIENT_IP:** Corresponds to `True-Client-Ip` header name.
	ClientAddressHeader PolicyConfigClientAddressHeaderEnum `mandatory:"false" json:"clientAddressHeader,omitempty"`

	// Enable or disable automatic content caching based on the response `cache-control` header. This feature enables the origin to act as a proxy cache. Caching is usually defined using `cache-control` header. For example `cache-control: max-age=120` means that the returned resource is valid for 120 seconds. Caching rules will overwrite this setting.
	IsCacheControlRespected *bool `mandatory:"false" json:"isCacheControlRespected"`

	// Enable or disable buffering of responses from the origin. Buffering improves overall stability in case of network issues, but slightly increases Time To First Byte.
	IsResponseBufferingEnabled *bool `mandatory:"false" json:"isResponseBufferingEnabled"`

	// The set cipher group for the configured TLS protocol. This sets the configuration for the TLS connections between clients and edge nodes only.
	// - **DEFAULT:** Cipher group supports TLS 1.0, TLS 1.1, TLS 1.2, TLS 1.3 protocols. It has the following ciphers enabled: `ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:AES:CAMELLIA:!DES-CBC3-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!MD5:!PSK:!aECDH:!EDH-DSS-DES-CBC3-SHA:!EDH-RSA-DES-CBC3-SHA:!KRB5-DES-CBC3-SHA`
	CipherGroup PolicyConfigCipherGroupEnum `mandatory:"false" json:"cipherGroup,omitempty"`

	// An object that represents a load balancing method and its properties.
	LoadBalancingMethod LoadBalancingMethod `mandatory:"false" json:"loadBalancingMethod"`

	// ModSecurity is not capable to inspect WebSockets. Therefore paths specified here have WAF disabled if Connection request header from the client has the value Upgrade (case insensitive matching) and Upgrade request header has the value websocket (case insensitive matching). Paths matches if the concatenation of request URL path and query starts with the contents of the one of `websocketPathPrefixes` array value. In All other cases challenges, like JSC, HIC and etc., remain active.
	WebsocketPathPrefixes []string `mandatory:"false" json:"websocketPathPrefixes"`

	// SNI stands for Server Name Indication and is an extension of the TLS protocol. It indicates which hostname is being contacted by the browser at the beginning of the 'handshake'-process. This allows a server to connect multiple SSL Certificates to one IP address and port.
	IsSniEnabled *bool `mandatory:"false" json:"isSniEnabled"`

	HealthChecks *HealthCheck `mandatory:"false" json:"healthChecks"`
}

func (m PolicyConfig) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *PolicyConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CertificateId              *string                             `json:"certificateId"`
		IsHttpsEnabled             *bool                               `json:"isHttpsEnabled"`
		IsHttpsForced              *bool                               `json:"isHttpsForced"`
		TlsProtocols               []PolicyConfigTlsProtocolsEnum      `json:"tlsProtocols"`
		IsOriginCompressionEnabled *bool                               `json:"isOriginCompressionEnabled"`
		IsBehindCdn                *bool                               `json:"isBehindCdn"`
		ClientAddressHeader        PolicyConfigClientAddressHeaderEnum `json:"clientAddressHeader"`
		IsCacheControlRespected    *bool                               `json:"isCacheControlRespected"`
		IsResponseBufferingEnabled *bool                               `json:"isResponseBufferingEnabled"`
		CipherGroup                PolicyConfigCipherGroupEnum         `json:"cipherGroup"`
		LoadBalancingMethod        loadbalancingmethod                 `json:"loadBalancingMethod"`
		WebsocketPathPrefixes      []string                            `json:"websocketPathPrefixes"`
		IsSniEnabled               *bool                               `json:"isSniEnabled"`
		HealthChecks               *HealthCheck                        `json:"healthChecks"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CertificateId = model.CertificateId

	m.IsHttpsEnabled = model.IsHttpsEnabled

	m.IsHttpsForced = model.IsHttpsForced

	m.TlsProtocols = make([]PolicyConfigTlsProtocolsEnum, len(model.TlsProtocols))
	for i, n := range model.TlsProtocols {
		m.TlsProtocols[i] = n
	}

	m.IsOriginCompressionEnabled = model.IsOriginCompressionEnabled

	m.IsBehindCdn = model.IsBehindCdn

	m.ClientAddressHeader = model.ClientAddressHeader

	m.IsCacheControlRespected = model.IsCacheControlRespected

	m.IsResponseBufferingEnabled = model.IsResponseBufferingEnabled

	m.CipherGroup = model.CipherGroup

	nn, e = model.LoadBalancingMethod.UnmarshalPolymorphicJSON(model.LoadBalancingMethod.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LoadBalancingMethod = nn.(LoadBalancingMethod)
	} else {
		m.LoadBalancingMethod = nil
	}

	m.WebsocketPathPrefixes = make([]string, len(model.WebsocketPathPrefixes))
	for i, n := range model.WebsocketPathPrefixes {
		m.WebsocketPathPrefixes[i] = n
	}

	m.IsSniEnabled = model.IsSniEnabled

	m.HealthChecks = model.HealthChecks

	return
}

// PolicyConfigTlsProtocolsEnum Enum with underlying type: string
type PolicyConfigTlsProtocolsEnum string

// Set of constants representing the allowable values for PolicyConfigTlsProtocolsEnum
const (
	PolicyConfigTlsProtocolsV1  PolicyConfigTlsProtocolsEnum = "TLS_V1"
	PolicyConfigTlsProtocolsV11 PolicyConfigTlsProtocolsEnum = "TLS_V1_1"
	PolicyConfigTlsProtocolsV12 PolicyConfigTlsProtocolsEnum = "TLS_V1_2"
	PolicyConfigTlsProtocolsV13 PolicyConfigTlsProtocolsEnum = "TLS_V1_3"
)

var mappingPolicyConfigTlsProtocols = map[string]PolicyConfigTlsProtocolsEnum{
	"TLS_V1":   PolicyConfigTlsProtocolsV1,
	"TLS_V1_1": PolicyConfigTlsProtocolsV11,
	"TLS_V1_2": PolicyConfigTlsProtocolsV12,
	"TLS_V1_3": PolicyConfigTlsProtocolsV13,
}

// GetPolicyConfigTlsProtocolsEnumValues Enumerates the set of values for PolicyConfigTlsProtocolsEnum
func GetPolicyConfigTlsProtocolsEnumValues() []PolicyConfigTlsProtocolsEnum {
	values := make([]PolicyConfigTlsProtocolsEnum, 0)
	for _, v := range mappingPolicyConfigTlsProtocols {
		values = append(values, v)
	}
	return values
}

// PolicyConfigClientAddressHeaderEnum Enum with underlying type: string
type PolicyConfigClientAddressHeaderEnum string

// Set of constants representing the allowable values for PolicyConfigClientAddressHeaderEnum
const (
	PolicyConfigClientAddressHeaderXForwardedFor PolicyConfigClientAddressHeaderEnum = "X_FORWARDED_FOR"
	PolicyConfigClientAddressHeaderXClientIp     PolicyConfigClientAddressHeaderEnum = "X_CLIENT_IP"
	PolicyConfigClientAddressHeaderXRealIp       PolicyConfigClientAddressHeaderEnum = "X_REAL_IP"
	PolicyConfigClientAddressHeaderClientIp      PolicyConfigClientAddressHeaderEnum = "CLIENT_IP"
	PolicyConfigClientAddressHeaderTrueClientIp  PolicyConfigClientAddressHeaderEnum = "TRUE_CLIENT_IP"
)

var mappingPolicyConfigClientAddressHeader = map[string]PolicyConfigClientAddressHeaderEnum{
	"X_FORWARDED_FOR": PolicyConfigClientAddressHeaderXForwardedFor,
	"X_CLIENT_IP":     PolicyConfigClientAddressHeaderXClientIp,
	"X_REAL_IP":       PolicyConfigClientAddressHeaderXRealIp,
	"CLIENT_IP":       PolicyConfigClientAddressHeaderClientIp,
	"TRUE_CLIENT_IP":  PolicyConfigClientAddressHeaderTrueClientIp,
}

// GetPolicyConfigClientAddressHeaderEnumValues Enumerates the set of values for PolicyConfigClientAddressHeaderEnum
func GetPolicyConfigClientAddressHeaderEnumValues() []PolicyConfigClientAddressHeaderEnum {
	values := make([]PolicyConfigClientAddressHeaderEnum, 0)
	for _, v := range mappingPolicyConfigClientAddressHeader {
		values = append(values, v)
	}
	return values
}

// PolicyConfigCipherGroupEnum Enum with underlying type: string
type PolicyConfigCipherGroupEnum string

// Set of constants representing the allowable values for PolicyConfigCipherGroupEnum
const (
	PolicyConfigCipherGroupDefault PolicyConfigCipherGroupEnum = "DEFAULT"
)

var mappingPolicyConfigCipherGroup = map[string]PolicyConfigCipherGroupEnum{
	"DEFAULT": PolicyConfigCipherGroupDefault,
}

// GetPolicyConfigCipherGroupEnumValues Enumerates the set of values for PolicyConfigCipherGroupEnum
func GetPolicyConfigCipherGroupEnumValues() []PolicyConfigCipherGroupEnum {
	values := make([]PolicyConfigCipherGroupEnum, 0)
	for _, v := range mappingPolicyConfigCipherGroup {
		values = append(values, v)
	}
	return values
}
