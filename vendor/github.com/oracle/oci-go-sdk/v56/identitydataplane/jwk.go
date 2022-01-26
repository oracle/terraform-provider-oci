// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Jwk The representation of Jwk
type Jwk struct {

	// The modulus.
	N *string `mandatory:"true" json:"n"`

	// The exponent.
	E *string `mandatory:"true" json:"e"`

	// The key id.
	Kid *string `mandatory:"true" json:"kid"`

	// The key use.
	Use *string `mandatory:"true" json:"use"`

	// The algorithm.
	Alg *string `mandatory:"true" json:"alg"`

	// The key type.
	Kty *string `mandatory:"true" json:"kty"`
}

func (m Jwk) String() string {
	return common.PointerString(m)
}
