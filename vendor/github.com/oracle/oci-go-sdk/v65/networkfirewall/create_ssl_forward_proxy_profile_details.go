// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSslForwardProxyProfileDetails Request for creating SSLForwardProxy used on the firewall policy rules.
type CreateSslForwardProxyProfileDetails struct {

	// Name of the decryption profile.
	Name *string `mandatory:"true" json:"name"`

	// Whether to block sessions if server's certificate is expired.
	IsExpiredCertificateBlocked *bool `mandatory:"false" json:"isExpiredCertificateBlocked"`

	// Whether to block sessions if server's certificate is issued by an untrusted certificate authority (CA).
	IsUntrustedIssuerBlocked *bool `mandatory:"false" json:"isUntrustedIssuerBlocked"`

	// Whether to block sessions if the revocation status check for server's certificate
	// does not succeed within the maximum allowed time (defaulting to 5 seconds).
	IsRevocationStatusTimeoutBlocked *bool `mandatory:"false" json:"isRevocationStatusTimeoutBlocked"`

	// Whether to block sessions if SSL version is not supported.
	IsUnsupportedVersionBlocked *bool `mandatory:"false" json:"isUnsupportedVersionBlocked"`

	// Whether to block sessions if SSL cipher suite is not supported.
	IsUnsupportedCipherBlocked *bool `mandatory:"false" json:"isUnsupportedCipherBlocked"`

	// Whether to block sessions if the revocation status check for server's certificate results in "unknown".
	IsUnknownRevocationStatusBlocked *bool `mandatory:"false" json:"isUnknownRevocationStatusBlocked"`

	// Whether to block sessions if the server's certificate uses extensions other than key usage and/or extended key usage.
	AreCertificateExtensionsRestricted *bool `mandatory:"false" json:"areCertificateExtensionsRestricted"`

	// Whether to automatically append SAN to impersonating certificate if server certificate is missing SAN.
	IsAutoIncludeAltName *bool `mandatory:"false" json:"isAutoIncludeAltName"`

	// Whether to block sessions if the firewall is temporarily unable to decrypt their traffic.
	IsOutOfCapacityBlocked *bool `mandatory:"false" json:"isOutOfCapacityBlocked"`
}

// GetName returns Name
func (m CreateSslForwardProxyProfileDetails) GetName() *string {
	return m.Name
}

func (m CreateSslForwardProxyProfileDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSslForwardProxyProfileDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateSslForwardProxyProfileDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSslForwardProxyProfileDetails CreateSslForwardProxyProfileDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateSslForwardProxyProfileDetails
	}{
		"SSL_FORWARD_PROXY",
		(MarshalTypeCreateSslForwardProxyProfileDetails)(m),
	}

	return json.Marshal(&s)
}
