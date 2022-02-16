// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Jwk) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
