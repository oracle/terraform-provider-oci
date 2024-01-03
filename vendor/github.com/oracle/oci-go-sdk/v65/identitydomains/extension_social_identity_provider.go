// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExtensionSocialIdentityProvider Social Identity Provider Extension Schema
type ExtensionSocialIdentityProvider struct {

	// Whether account linking is enabled
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AccountLinkingEnabled *bool `mandatory:"true" json:"accountLinkingEnabled"`

	// Whether registration is enabled
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	RegistrationEnabled *bool `mandatory:"true" json:"registrationEnabled"`

	// Social IDP Client Application Client ID
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ConsumerKey *string `mandatory:"true" json:"consumerKey"`

	// Social IDP Client Application Client Secret
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - idcsSensitive: encrypt
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ConsumerSecret *string `mandatory:"true" json:"consumerSecret"`

	// Service Provider Name
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ServiceProviderName *string `mandatory:"true" json:"serviceProviderName"`

	// Status
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Status ExtensionSocialIdentityProviderStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Social IDP Authorization URL
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AuthzUrl *string `mandatory:"false" json:"authzUrl"`

	// Social IDP Access token URL
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AccessTokenUrl *string `mandatory:"false" json:"accessTokenUrl"`

	// Relay Param variable for Social IDP
	// **Added In:** 2305190132
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [relayParamKey]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	RelayIdpParamMappings []IdentityProviderRelayIdpParamMappings `mandatory:"false" json:"relayIdpParamMappings"`

	// Social IDP User profile URL
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ProfileUrl *string `mandatory:"false" json:"profileUrl"`

	// Scope to request
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Scope []string `mandatory:"false" json:"scope"`

	// Admin scope to request
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AdminScope []string `mandatory:"false" json:"adminScope"`

	// Social IDP allowed clock skew time
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	ClockSkewInSeconds *int `mandatory:"false" json:"clockSkewInSeconds"`

	// redirect URL for social idp
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	RedirectUrl *string `mandatory:"false" json:"redirectUrl"`

	// Discovery URL
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DiscoveryUrl *string `mandatory:"false" json:"discoveryUrl"`

	// Whether the client credential is contained in payload
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	ClientCredentialInPayload *bool `mandatory:"false" json:"clientCredentialInPayload"`

	// Id attribute used for account linking
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdAttribute *string `mandatory:"false" json:"idAttribute"`
}

func (m ExtensionSocialIdentityProvider) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionSocialIdentityProvider) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExtensionSocialIdentityProviderStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetExtensionSocialIdentityProviderStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExtensionSocialIdentityProviderStatusEnum Enum with underlying type: string
type ExtensionSocialIdentityProviderStatusEnum string

// Set of constants representing the allowable values for ExtensionSocialIdentityProviderStatusEnum
const (
	ExtensionSocialIdentityProviderStatusCreated ExtensionSocialIdentityProviderStatusEnum = "created"
	ExtensionSocialIdentityProviderStatusDeleted ExtensionSocialIdentityProviderStatusEnum = "deleted"
)

var mappingExtensionSocialIdentityProviderStatusEnum = map[string]ExtensionSocialIdentityProviderStatusEnum{
	"created": ExtensionSocialIdentityProviderStatusCreated,
	"deleted": ExtensionSocialIdentityProviderStatusDeleted,
}

var mappingExtensionSocialIdentityProviderStatusEnumLowerCase = map[string]ExtensionSocialIdentityProviderStatusEnum{
	"created": ExtensionSocialIdentityProviderStatusCreated,
	"deleted": ExtensionSocialIdentityProviderStatusDeleted,
}

// GetExtensionSocialIdentityProviderStatusEnumValues Enumerates the set of values for ExtensionSocialIdentityProviderStatusEnum
func GetExtensionSocialIdentityProviderStatusEnumValues() []ExtensionSocialIdentityProviderStatusEnum {
	values := make([]ExtensionSocialIdentityProviderStatusEnum, 0)
	for _, v := range mappingExtensionSocialIdentityProviderStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionSocialIdentityProviderStatusEnumStringValues Enumerates the set of values in String for ExtensionSocialIdentityProviderStatusEnum
func GetExtensionSocialIdentityProviderStatusEnumStringValues() []string {
	return []string{
		"created",
		"deleted",
	}
}

// GetMappingExtensionSocialIdentityProviderStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionSocialIdentityProviderStatusEnum(val string) (ExtensionSocialIdentityProviderStatusEnum, bool) {
	enum, ok := mappingExtensionSocialIdentityProviderStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
