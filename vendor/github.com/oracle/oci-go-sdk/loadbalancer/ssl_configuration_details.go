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

// SslConfigurationDetails The load balancer's SSL handling configuration details.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type SslConfigurationDetails struct {

	// A friendly name for the certificate bundle. It must be unique and it cannot be changed.
	// Valid certificate bundle names include only alphanumeric characters, dashes, and underscores.
	// Certificate bundle names cannot contain spaces. Avoid entering confidential information.
	// Example: `example_certificate_bundle`
	CertificateName *string `mandatory:"true" json:"certificateName"`

	// Whether the load balancer listener should verify peer certificates.
	// Example: `true`
	VerifyPeerCertificate *bool `mandatory:"false" json:"verifyPeerCertificate"`

	// The maximum depth for peer certificate chain verification.
	// Example: `3`
	VerifyDepth *int `mandatory:"false" json:"verifyDepth"`
}

func (m SslConfigurationDetails) String() string {
	return common.PointerString(m)
}
