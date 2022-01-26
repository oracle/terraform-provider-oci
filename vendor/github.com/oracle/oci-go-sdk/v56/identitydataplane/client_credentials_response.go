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

// ClientCredentialsResponse The representation of ClientCredentialsResponse
type ClientCredentialsResponse struct {

	// The access token.
	AccessToken *string `mandatory:"true" json:"access_token"`

	// The token type.
	TokenType *string `mandatory:"true" json:"token_type"`

	// The amount of time until the token expires.
	ExpiresIn *string `mandatory:"true" json:"expires_in"`
}

func (m ClientCredentialsResponse) String() string {
	return common.PointerString(m)
}
