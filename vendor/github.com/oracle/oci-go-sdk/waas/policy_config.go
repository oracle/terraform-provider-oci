// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
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

	// A list of allowed TLS protocols. Only applicable when HTTPS support is enabled. It affects client's connection to the edge nodes. The most secure TLS version will be chosen.
	// - **TLS_V1:** corresponds to TLS 1.0 specification.
	// - **TLS_V1_1:** corresponds to TLS 1.1 specification.
	// - **TLS_V1_2:** corresponds to TLS 1.2 specification.
	// - **TLS_V1_3:** corresponds to TLS 1.3 specification.
	// Enabled TLS protocols must go in a row. For example if TLS_v1_1 and TLS_V1_3 are enabled, TLS_V1_2 must be enabled too.
	TlsProtocols []PolicyConfigTlsProtocolsEnum `mandatory:"false" json:"tlsProtocols,omitempty"`

	// Enable or disable GZIP compression of origin responses. If enabled, the header `Accept-Encoding: gzip` is sent to origin, otherwise - empty `Accept-Encoding:` header is used.
	IsOriginCompressionEnabled *bool `mandatory:"false" json:"isOriginCompressionEnabled"`

	// Enable or disable the use of CDN. It allows to specify true client IP address if clients do not connect directly to us.
	IsBehindCdn *bool `mandatory:"false" json:"isBehindCdn"`

	// The HTTP header used to pass the client IP address from the CDN if `isBehindCdn` is enabled. This feature consumes the header and its value as the true client IP address. It does not create the header. Using trusted chains (for example `X-Client-Ip: 11.1.1.1, 13.3.3.3`), the last IP address in the list will be used as true client IP address. In case of multiple headers with the same name, the first one will be used. If the header is not present it will use the connecting IP address as the true client IP address. It's assumed that CDN sets the correct client IP address and prevents spoofing.
	// - **X_FORWARDED_FOR:** Corresponds to `X-Forwarded-For` header name.
	// - **X_CLIENT_IP:** Corresponds to `X-Client-Ip` header name.
	// - **X_REAL_IP:** Corresponds to `X-Real-Ip` header name.
	// - **CLIENT_IP:** Corresponds to `Client-Ip` header name.
	// - **TRUE_CLIENT_IP:** Corresponds to `True-Client-Ip` header name.
	ClientAddressHeader PolicyConfigClientAddressHeaderEnum `mandatory:"false" json:"clientAddressHeader,omitempty"`

	// Enable or disable automatic content caching based on the response `cache-control` header. This feature enables the origin to act as a proxy cache. Caching policies are usually defined using `cache-control` header. For example `cache-control: max-age=120` means that the returned resource is valid for 120 seconds. Caching rules will overwrite this setting.
	IsCacheControlRespected *bool `mandatory:"false" json:"isCacheControlRespected"`

	// Enable or disable buffering of responses from the origin. Buffering improves overall stability in case of network issues, but slightly increases Time To First Byte.
	IsResponseBufferingEnabled *bool `mandatory:"false" json:"isResponseBufferingEnabled"`

	// The cipher group
	// - **DEFAULT:** Cipher group supports TLS 1.0, TLS 1.1, TLS 1.2, TLS 1.3 protocols. It has the following ciphers enabled: `ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:AES:CAMELLIA:!DES-CBC3-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!MD5:!PSK:!aECDH:!EDH-DSS-DES-CBC3-SHA:!EDH-RSA-DES-CBC3-SHA:!KRB5-DES-CBC3-SHA`
	CipherGroup PolicyConfigCipherGroupEnum `mandatory:"false" json:"cipherGroup,omitempty"`
}

func (m PolicyConfig) String() string {
	return common.PointerString(m)
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
