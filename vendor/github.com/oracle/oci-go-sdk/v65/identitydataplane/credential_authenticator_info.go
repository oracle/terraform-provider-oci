// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Data Plane API
//
// APIs for managing identity data plane services. For example, use this API to create a scoped-access security token. To manage identity domains (for example, creating or deleting an identity domain) or to manage resources (for example, users and groups) within the default identity domain, see IAM API (https://docs.oracle.com/iaas/api/#/en/identity/).
//

package identitydataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CredentialAuthenticatorInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
