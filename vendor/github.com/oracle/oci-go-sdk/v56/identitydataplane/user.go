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

// User The representation of User
type User struct {

	// The user's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The name of the user.
	Name *string `mandatory:"true" json:"name"`

	// If the provided password is a one-time password.
	IsOTP *bool `mandatory:"true" json:"isOTP"`

	// If mfa is activated.
	IsMfaActivated *bool `mandatory:"true" json:"isMfaActivated"`

	// If the user has been mfa verified.
	IsMfaVerified *bool `mandatory:"true" json:"isMfaVerified"`
}

func (m User) String() string {
	return common.PointerString(m)
}
