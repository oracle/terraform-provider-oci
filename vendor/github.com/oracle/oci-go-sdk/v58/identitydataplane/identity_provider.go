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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentityProvider) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIdentityProviderProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetIdentityProviderProtocolEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IdentityProviderProtocolEnum Enum with underlying type: string
type IdentityProviderProtocolEnum string

// Set of constants representing the allowable values for IdentityProviderProtocolEnum
const (
	IdentityProviderProtocolSaml2 IdentityProviderProtocolEnum = "SAML2"
)

var mappingIdentityProviderProtocolEnum = map[string]IdentityProviderProtocolEnum{
	"SAML2": IdentityProviderProtocolSaml2,
}

// GetIdentityProviderProtocolEnumValues Enumerates the set of values for IdentityProviderProtocolEnum
func GetIdentityProviderProtocolEnumValues() []IdentityProviderProtocolEnum {
	values := make([]IdentityProviderProtocolEnum, 0)
	for _, v := range mappingIdentityProviderProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProviderProtocolEnumStringValues Enumerates the set of values in String for IdentityProviderProtocolEnum
func GetIdentityProviderProtocolEnumStringValues() []string {
	return []string{
		"SAML2",
	}
}

// GetMappingIdentityProviderProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProviderProtocolEnum(val string) (IdentityProviderProtocolEnum, bool) {
	mappingIdentityProviderProtocolEnumIgnoreCase := make(map[string]IdentityProviderProtocolEnum)
	for k, v := range mappingIdentityProviderProtocolEnum {
		mappingIdentityProviderProtocolEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingIdentityProviderProtocolEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
