// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateVanityUrlDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
