// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateVanityUrlDetails Input payload to update a vanity url.
type UpdateVanityUrlDetails struct {

	// PEM Private key for HTTPS connections.
	PrivateKey *string `mandatory:"true" json:"privateKey"`

	// PEM certificate for HTTPS connections.
	PublicCertificate *string `mandatory:"true" json:"publicCertificate"`

	// PEM CA certificate(s) for HTTPS connections. This may include multiple PEM certificates.
	CaCertificate *string `mandatory:"true" json:"caCertificate"`

	// Passphrase for the PEM Private key (if any).
	Passphrase *string `mandatory:"false" json:"passphrase"`
}

func (m UpdateVanityUrlDetails) String() string {
	return common.PointerString(m)
}
