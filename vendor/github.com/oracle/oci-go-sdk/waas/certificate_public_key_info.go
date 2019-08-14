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

// CertificatePublicKeyInfo Information about the public key and the algorithm used by the public key.
type CertificatePublicKeyInfo struct {

	// The algorithm identifier and parameters for the public key.
	Algorithm *string `mandatory:"false" json:"algorithm"`

	// The private key exponent.
	Exponent *int `mandatory:"false" json:"exponent"`

	// The number of bits in a key used by a cryptographic algorithm.
	KeySize *int `mandatory:"false" json:"keySize"`
}

func (m CertificatePublicKeyInfo) String() string {
	return common.PointerString(m)
}
