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

// UserNotFoundAuthenticateUserResult The representation of UserNotFoundAuthenticateUserResult
type UserNotFoundAuthenticateUserResult struct {

	// The tenant name.
	TenantInput *string `mandatory:"true" json:"tenantInput"`

	// The user name.
	UserInput *string `mandatory:"true" json:"userInput"`

	// The resolved tenant id.
	ResolvedTenantId *string `mandatory:"true" json:"resolvedTenantId"`
}

func (m UserNotFoundAuthenticateUserResult) String() string {
	return common.PointerString(m)
}
