// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// Whether the load balancer listener should resume an encrypted session by reusing the cryptographic parameters of a previous TLS session, without having to perform a full handshake again.
	// If "true", the service resumes the previous TLS encrypted session.
	// If "false", the service starts a new TLS encrypted session.
	// Enabling session resumption improves performance but provides a lower level of security. Disabling session resumption improves security but reduces performance.
	// Example: `true`
	HasSessionResumption *bool `mandatory:"false" json:"hasSessionResumption"`

	// Ids for OCI certificates service CA or CA bundles for the load balancer to trust.
	// Example: `[ocid1.cabundle.oc1.us-ashburn-1.amaaaaaaav3bgsaagl4zzyqdop5i2vuwoqewdvauuw34llqa74otq2jdsfyq]`
	TrustedCertificateAuthorityIds []string `mandatory:"false" json:"trustedCertificateAuthorityIds"`

	// Ids for OCI certificates service certificates. Currently only a single Id may be passed.
	// Example: `[ocid1.certificate.oc1.us-ashburn-1.amaaaaaaav3bgsaa5o2q7rh5nfmkkukfkogasqhk6af2opufhjlqg7m6jqzq]`
	CertificateIds []string `mandatory:"false" json:"certificateIds"`

	// A friendly name for the certificate bundle. It must be unique and it cannot be changed.
	// Valid certificate bundle names include only alphanumeric characters, dashes, and underscores.
	// Certificate bundle names cannot contain spaces. Avoid entering confidential information.
	// Example: `example_certificate_bundle`
	CertificateName *string `mandatory:"false" json:"certificateName"`

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
	// *  TLSv1.3
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SslConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSslConfigurationServerOrderPreferenceEnum(string(m.ServerOrderPreference)); !ok && m.ServerOrderPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServerOrderPreference: %s. Supported values are: %s.", m.ServerOrderPreference, strings.Join(GetSslConfigurationServerOrderPreferenceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SslConfigurationServerOrderPreferenceEnum Enum with underlying type: string
type SslConfigurationServerOrderPreferenceEnum string

// Set of constants representing the allowable values for SslConfigurationServerOrderPreferenceEnum
const (
	SslConfigurationServerOrderPreferenceEnabled  SslConfigurationServerOrderPreferenceEnum = "ENABLED"
	SslConfigurationServerOrderPreferenceDisabled SslConfigurationServerOrderPreferenceEnum = "DISABLED"
)

var mappingSslConfigurationServerOrderPreferenceEnum = map[string]SslConfigurationServerOrderPreferenceEnum{
	"ENABLED":  SslConfigurationServerOrderPreferenceEnabled,
	"DISABLED": SslConfigurationServerOrderPreferenceDisabled,
}

var mappingSslConfigurationServerOrderPreferenceEnumLowerCase = map[string]SslConfigurationServerOrderPreferenceEnum{
	"enabled":  SslConfigurationServerOrderPreferenceEnabled,
	"disabled": SslConfigurationServerOrderPreferenceDisabled,
}

// GetSslConfigurationServerOrderPreferenceEnumValues Enumerates the set of values for SslConfigurationServerOrderPreferenceEnum
func GetSslConfigurationServerOrderPreferenceEnumValues() []SslConfigurationServerOrderPreferenceEnum {
	values := make([]SslConfigurationServerOrderPreferenceEnum, 0)
	for _, v := range mappingSslConfigurationServerOrderPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetSslConfigurationServerOrderPreferenceEnumStringValues Enumerates the set of values in String for SslConfigurationServerOrderPreferenceEnum
func GetSslConfigurationServerOrderPreferenceEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingSslConfigurationServerOrderPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSslConfigurationServerOrderPreferenceEnum(val string) (SslConfigurationServerOrderPreferenceEnum, bool) {
	enum, ok := mappingSslConfigurationServerOrderPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
