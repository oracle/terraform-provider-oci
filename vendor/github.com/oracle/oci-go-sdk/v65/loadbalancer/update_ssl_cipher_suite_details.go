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

// UpdateSslCipherSuiteDetails The configuration details for updating an SSL cipher suite.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateSslCipherSuiteDetails struct {

	// A list of SSL ciphers the load balancer must support for HTTPS or SSL connections.
	// The following ciphers are valid values for this property:
	// *  __TLSv1.3 ciphers__
	//         "TLS_AES_128_GCM_SHA256"
	//         "TLS_AES_256_GCM_SHA384"
	//         "TLS_CHACHA20_POLY1305_SHA256"
	//         "TLS_AES_128_CCM_SHA256"
	//         "TLS_AES_128_CCM_8_SHA256"
	// *  __TLSv1.2 ciphers__
	//         "AES128-GCM-SHA256"
	//         "AES128-SHA256"
	//         "AES256-GCM-SHA384"
	//         "AES256-SHA256"
	//         "DH-DSS-AES128-GCM-SHA256"
	//         "DH-DSS-AES128-SHA256"
	//         "DH-DSS-AES256-GCM-SHA384"
	//         "DH-DSS-AES256-SHA256"
	//         "DH-RSA-AES128-GCM-SHA256"
	//         "DH-RSA-AES128-SHA256"
	//         "DH-RSA-AES256-GCM-SHA384"
	//         "DH-RSA-AES256-SHA256"
	//         "DHE-DSS-AES128-GCM-SHA256"
	//         "DHE-DSS-AES128-SHA256"
	//         "DHE-DSS-AES256-GCM-SHA384"
	//         "DHE-DSS-AES256-SHA256"
	//         "DHE-RSA-AES128-GCM-SHA256"
	//         "DHE-RSA-AES128-SHA256"
	//         "DHE-RSA-AES256-GCM-SHA384"
	//         "DHE-RSA-AES256-SHA256"
	//         "ECDH-ECDSA-AES128-GCM-SHA256"
	//         "ECDH-ECDSA-AES128-SHA256"
	//         "ECDH-ECDSA-AES256-GCM-SHA384"
	//         "ECDH-ECDSA-AES256-SHA384"
	//         "ECDH-RSA-AES128-GCM-SHA256"
	//         "ECDH-RSA-AES128-SHA256"
	//         "ECDH-RSA-AES256-GCM-SHA384"
	//         "ECDH-RSA-AES256-SHA384"
	//         "ECDHE-ECDSA-AES128-GCM-SHA256"
	//         "ECDHE-ECDSA-AES128-SHA256"
	//         "ECDHE-ECDSA-AES256-GCM-SHA384"
	//         "ECDHE-ECDSA-AES256-SHA384"
	//         "ECDHE-RSA-AES128-GCM-SHA256"
	//         "ECDHE-RSA-AES128-SHA256"
	//         "ECDHE-RSA-AES256-GCM-SHA384"
	//         "ECDHE-RSA-AES256-SHA384"
	// *  __TLSv1 ciphers also supported by TLSv1.2__
	//         "AES128-SHA"
	//         "AES256-SHA"
	//         "CAMELLIA128-SHA"
	//         "CAMELLIA256-SHA"
	//         "DES-CBC3-SHA"
	//         "DH-DSS-AES128-SHA"
	//         "DH-DSS-AES256-SHA"
	//         "DH-DSS-CAMELLIA128-SHA"
	//         "DH-DSS-CAMELLIA256-SHA"
	//         "DH-DSS-DES-CBC3-SHAv"
	//         "DH-DSS-SEED-SHA"
	//         "DH-RSA-AES128-SHA"
	//         "DH-RSA-AES256-SHA"
	//         "DH-RSA-CAMELLIA128-SHA"
	//         "DH-RSA-CAMELLIA256-SHA"
	//         "DH-RSA-DES-CBC3-SHA"
	//         "DH-RSA-SEED-SHA"
	//         "DHE-DSS-AES128-SHA"
	//         "DHE-DSS-AES256-SHA"
	//         "DHE-DSS-CAMELLIA128-SHA"
	//         "DHE-DSS-CAMELLIA256-SHA"
	//         "DHE-DSS-DES-CBC3-SHA"
	//         "DHE-DSS-SEED-SHA"
	//         "DHE-RSA-AES128-SHA"
	//         "DHE-RSA-AES256-SHA"
	//         "DHE-RSA-CAMELLIA128-SHA"
	//         "DHE-RSA-CAMELLIA256-SHA"
	//         "DHE-RSA-DES-CBC3-SHA"
	//         "DHE-RSA-SEED-SHA"
	//         "ECDH-ECDSA-AES128-SHA"
	//         "ECDH-ECDSA-AES256-SHA"
	//         "ECDH-ECDSA-DES-CBC3-SHA"
	//         "ECDH-ECDSA-RC4-SHA"
	//         "ECDH-RSA-AES128-SHA"
	//         "ECDH-RSA-AES256-SHA"
	//         "ECDH-RSA-DES-CBC3-SHA"
	//         "ECDH-RSA-RC4-SHA"
	//         "ECDHE-ECDSA-AES128-SHA"
	//         "ECDHE-ECDSA-AES256-SHA"
	//         "ECDHE-ECDSA-DES-CBC3-SHA"
	//         "ECDHE-ECDSA-RC4-SHA"
	//         "ECDHE-RSA-AES128-SHA"
	//         "ECDHE-RSA-AES256-SHA"
	//         "ECDHE-RSA-DES-CBC3-SHA"
	//         "ECDHE-RSA-RC4-SHA"
	//         "IDEA-CBC-SHA"
	//         "KRB5-DES-CBC3-MD5"
	//         "KRB5-DES-CBC3-SHA"
	//         "KRB5-IDEA-CBC-MD5"
	//         "KRB5-IDEA-CBC-SHA"
	//         "KRB5-RC4-MD5"
	//         "KRB5-RC4-SHA"
	//         "PSK-3DES-EDE-CBC-SHA"
	//         "PSK-AES128-CBC-SHA"
	//         "PSK-AES256-CBC-SHA"
	//         "PSK-RC4-SHA"
	//         "RC4-MD5"
	//         "RC4-SHA"
	//         "SEED-SHA"
	// example: `["ECDHE-RSA-AES256-GCM-SHA384","ECDHE-ECDSA-AES256-GCM-SHA384","ECDHE-RSA-AES128-GCM-SHA256"]`
	Ciphers []string `mandatory:"true" json:"ciphers"`
}

func (m UpdateSslCipherSuiteDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSslCipherSuiteDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
