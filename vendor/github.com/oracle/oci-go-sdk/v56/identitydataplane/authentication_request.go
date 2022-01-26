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

// AuthenticationRequest The representation of AuthenticationRequest
type AuthenticationRequest struct {

	// The user name
	UserName *string `mandatory:"true" json:"userName"`

	// The password
	Password *string `mandatory:"true" json:"password"`

	// The name of the tenancy
	TenantName *string `mandatory:"true" json:"tenantName"`
}

func (m AuthenticationRequest) String() string {
	return common.PointerString(m)
}
