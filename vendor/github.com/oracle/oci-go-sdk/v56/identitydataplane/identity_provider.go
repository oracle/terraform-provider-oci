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

// IdentityProvider The representation of IdentityProvider
type IdentityProvider struct {

	// The id of the provider.
	Id *string `mandatory:"true" json:"id"`

	// The name of the provider.
	Name *string `mandatory:"true" json:"name"`

	// The name of the tenant.
	TenantName *string `mandatory:"true" json:"tenantName"`

	// The id of the tenant.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The SAML endpoint where user will be redirected.
	RedirectUri *string `mandatory:"true" json:"redirectUri"`

	// The signing certificate of the provider.
	SigningCertificate *string `mandatory:"true" json:"signingCertificate"`

	// The type of the provider.
	Protocol IdentityProviderProtocolEnum `mandatory:"true" json:"protocol"`

	// The id of the service provider entity.
	ServiceProviderEntityId *string `mandatory:"true" json:"serviceProviderEntityId"`

	// Whether to force authentication.
	ForceAuthentication *bool `mandatory:"true" json:"forceAuthentication"`

	// Authentication context class refs.
	AuthnContextClassRefs []string `mandatory:"true" json:"authnContextClassRefs"`
}

func (m IdentityProvider) String() string {
	return common.PointerString(m)
}

// IdentityProviderProtocolEnum Enum with underlying type: string
type IdentityProviderProtocolEnum string

// Set of constants representing the allowable values for IdentityProviderProtocolEnum
const (
	IdentityProviderProtocolSaml2 IdentityProviderProtocolEnum = "SAML2"
)

var mappingIdentityProviderProtocol = map[string]IdentityProviderProtocolEnum{
	"SAML2": IdentityProviderProtocolSaml2,
}

// GetIdentityProviderProtocolEnumValues Enumerates the set of values for IdentityProviderProtocolEnum
func GetIdentityProviderProtocolEnumValues() []IdentityProviderProtocolEnum {
	values := make([]IdentityProviderProtocolEnum, 0)
	for _, v := range mappingIdentityProviderProtocol {
		values = append(values, v)
	}
	return values
}
