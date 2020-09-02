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

// SslConfiguration A listener's SSL handling configuration.
// To use SSL, a listener must be associated with a Certificate.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type SslConfiguration struct {

	// The maximum depth for peer certificate chain verification.
	// Example: `3`
	VerifyDepth *int `mandatory:"true" json:"verifyDepth"`

	// Whether the load balancer listener should verify peer certificates.
	// Example: `true`
	VerifyPeerCertificate *bool `mandatory:"true" json:"verifyPeerCertificate"`

	// A friendly name for the certificate bundle. It must be unique and it cannot be changed.
	// Valid certificate bundle names include only alphanumeric characters, dashes, and underscores.
	// Certificate bundle names cannot contain spaces. Avoid entering confidential information.
	// Example: `example_certificate_bundle`
	CertificateName *string `mandatory:"true" json:"certificateName"`

	// When this attribute is set to ENABLED, the system gives preference to the server ciphers over the client
	// ciphers.
	// **Note:** This configuration is applicable only when the load balancer is acting as an SSL/HTTPS server. This
	//           field is ignored when the `SSLConfiguration` object is associated with a backend set.
	ServerOrderPreference SslConfigurationServerOrderPreferenceEnum `mandatory:"false" json:"serverOrderPreference,omitempty"`

	// The name of the cipher suite to use for HTTPS or SSL connections.
	// If this field is not specified, the default is `oci-default-ssl-cipher-suite-v1`.
	// **Notes:**
	// *  You must ensure compatibility between the specified SSL protocols and the ciphers configured in the cipher
	//    suite. Clients cannot perform an SSL handshake if there is an incompatible configuration.
	// *  You must ensure compatibility between the ciphers configured in the cipher suite and the configured
	//    certificates. For example, RSA-based ciphers require RSA certificates and ECDSA-based ciphers require ECDSA
	//    certificates.
	// *  If the cipher configuration is not modified after load balancer creation, the `GET` operation returns
	//    `oci-default-ssl-cipher-suite-v1` as the value of this field in the SSL configuration for existing listeners
	//    that predate this feature.
	// *  If the cipher configuration was modified using Oracle operations after load balancer creation, the `GET`
	//    operation returns `oci-customized-ssl-cipher-suite` as the value of this field in the SSL configuration for
	//    existing listeners that predate this feature.
	// *  The `GET` operation returns `oci-wider-compatible-ssl-cipher-suite-v1` as the value of this field in the SSL
	//    configuration for existing backend sets that predate this feature.
	// *  If the `GET` operation on a listener returns `oci-customized-ssl-cipher-suite` as the value of this field,
	//    you must specify an appropriate predefined or custom cipher suite name when updating the resource.
	// *  The `oci-customized-ssl-cipher-suite` Oracle reserved cipher suite name is not accepted as valid input for
	//    this field.
	// example: `example_cipher_suite`
	CipherSuiteName *string `mandatory:"false" json:"cipherSuiteName"`

	// A list of SSL protocols the load balancer must support for HTTPS or SSL connections.
	// The load balancer uses SSL protocols to establish a secure connection between a client and a server. A secure
	// connection ensures that all data passed between the client and the server is private.
	// The Load Balancing service supports the following protocols:
	// *  TLSv1
	// *  TLSv1.1
	// *  TLSv1.2
	// If this field is not specified, TLSv1.2 is the default.
	// **Warning:** All SSL listeners created on a given port must use the same set of SSL protocols.
	// **Notes:**
	// *  The handshake to establish an SSL connection fails if the client supports none of the specified protocols.
	// *  You must ensure compatibility between the specified SSL protocols and the ciphers configured in the cipher
	//    suite.
	// *  For all existing load balancer listeners and backend sets that predate this feature, the `GET` operation
	//    displays a list of SSL protocols currently used by those resources.
	// example: `["TLSv1.1", "TLSv1.2"]`
	Protocols []string `mandatory:"false" json:"protocols"`
}

func (m SslConfiguration) String() string {
	return common.PointerString(m)
}

// SslConfigurationServerOrderPreferenceEnum Enum with underlying type: string
type SslConfigurationServerOrderPreferenceEnum string

// Set of constants representing the allowable values for SslConfigurationServerOrderPreferenceEnum
const (
	SslConfigurationServerOrderPreferenceEnabled  SslConfigurationServerOrderPreferenceEnum = "ENABLED"
	SslConfigurationServerOrderPreferenceDisabled SslConfigurationServerOrderPreferenceEnum = "DISABLED"
)

var mappingSslConfigurationServerOrderPreference = map[string]SslConfigurationServerOrderPreferenceEnum{
	"ENABLED":  SslConfigurationServerOrderPreferenceEnabled,
	"DISABLED": SslConfigurationServerOrderPreferenceDisabled,
}

// GetSslConfigurationServerOrderPreferenceEnumValues Enumerates the set of values for SslConfigurationServerOrderPreferenceEnum
func GetSslConfigurationServerOrderPreferenceEnumValues() []SslConfigurationServerOrderPreferenceEnum {
	values := make([]SslConfigurationServerOrderPreferenceEnum, 0)
	for _, v := range mappingSslConfigurationServerOrderPreference {
		values = append(values, v)
	}
	return values
}
