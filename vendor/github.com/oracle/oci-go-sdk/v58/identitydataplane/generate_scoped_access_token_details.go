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

// GenerateScopedAccessTokenDetails The representation of GenerateScopedAccessTokenDetails
type GenerateScopedAccessTokenDetails struct {

	// Scope definition for the scoped access token
	Scope *string `mandatory:"true" json:"scope"`

	// A temporary public key, owned by the service. The service also owns the corresponding private key. This public
	// key will by put inside the security token by the auth service after successful validation of the certificate.
	PublicKey *string `mandatory:"true" json:"publicKey"`
}

func (m GenerateScopedAccessTokenDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateScopedAccessTokenDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
