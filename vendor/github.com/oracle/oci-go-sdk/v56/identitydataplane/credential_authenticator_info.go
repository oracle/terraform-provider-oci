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

// CredentialAuthenticatorInfo The representation of CredentialAuthenticatorInfo
type CredentialAuthenticatorInfo struct {

	// The raw credential.
	RawCredential *string `mandatory:"true" json:"rawCredential"`

	// The id of the user.
	UserId *string `mandatory:"true" json:"userId"`

	// The id of the tenant.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The name of the user.
	UserName *string `mandatory:"true" json:"userName"`

	// The name of the tenant.
	TenantName *string `mandatory:"true" json:"tenantName"`

	// The credential identifier.
	CredentialIdentifier *string `mandatory:"true" json:"credentialIdentifier"`

	// The credential list.
	CredentialList []string `mandatory:"true" json:"credentialList"`

	// The name of the service that is making this authorization request.
	Service *string `mandatory:"true" json:"service"`

	// The id of the client.
	ClientId *string `mandatory:"true" json:"clientId"`
}

func (m CredentialAuthenticatorInfo) String() string {
	return common.PointerString(m)
}
