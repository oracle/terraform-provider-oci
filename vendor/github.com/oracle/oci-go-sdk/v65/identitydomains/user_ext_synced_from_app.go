// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
// Use this pattern to construct endpoints for identity domains: `https://<domainURL>/admin/v1/`. See Finding an Identity Domain URL (https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/locate-identity-domain-url.htm) to locate the domain URL you need.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UserExtSyncedFromApp Managed App or an Identity Source from where the user is synced. If enabled, this Managed App or Identity Source can be used for performing delegated authentication.
// **Added In:** 18.2.6
// **SCIM++ Properties:**
//   - idcsCompositeKey: [value]
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: readOnly
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type UserExtSyncedFromApp struct {

	// App identifier
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// A label that indicates whether this is an App or IdentitySource.
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - idcsDefaultValue: IdentitySource
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type UserExtSyncedFromAppTypeEnum `mandatory:"true" json:"type"`

	// App URI
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// App Display Name
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`
}

func (m UserExtSyncedFromApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserExtSyncedFromApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserExtSyncedFromAppTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUserExtSyncedFromAppTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserExtSyncedFromAppTypeEnum Enum with underlying type: string
type UserExtSyncedFromAppTypeEnum string

// Set of constants representing the allowable values for UserExtSyncedFromAppTypeEnum
const (
	UserExtSyncedFromAppTypeIdentitysource   UserExtSyncedFromAppTypeEnum = "IdentitySource"
	UserExtSyncedFromAppTypeApp              UserExtSyncedFromAppTypeEnum = "App"
	UserExtSyncedFromAppTypeIdentityprovider UserExtSyncedFromAppTypeEnum = "IdentityProvider"
)

var mappingUserExtSyncedFromAppTypeEnum = map[string]UserExtSyncedFromAppTypeEnum{
	"IdentitySource":   UserExtSyncedFromAppTypeIdentitysource,
	"App":              UserExtSyncedFromAppTypeApp,
	"IdentityProvider": UserExtSyncedFromAppTypeIdentityprovider,
}

var mappingUserExtSyncedFromAppTypeEnumLowerCase = map[string]UserExtSyncedFromAppTypeEnum{
	"identitysource":   UserExtSyncedFromAppTypeIdentitysource,
	"app":              UserExtSyncedFromAppTypeApp,
	"identityprovider": UserExtSyncedFromAppTypeIdentityprovider,
}

// GetUserExtSyncedFromAppTypeEnumValues Enumerates the set of values for UserExtSyncedFromAppTypeEnum
func GetUserExtSyncedFromAppTypeEnumValues() []UserExtSyncedFromAppTypeEnum {
	values := make([]UserExtSyncedFromAppTypeEnum, 0)
	for _, v := range mappingUserExtSyncedFromAppTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserExtSyncedFromAppTypeEnumStringValues Enumerates the set of values in String for UserExtSyncedFromAppTypeEnum
func GetUserExtSyncedFromAppTypeEnumStringValues() []string {
	return []string{
		"IdentitySource",
		"App",
		"IdentityProvider",
	}
}

// GetMappingUserExtSyncedFromAppTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserExtSyncedFromAppTypeEnum(val string) (UserExtSyncedFromAppTypeEnum, bool) {
	enum, ok := mappingUserExtSyncedFromAppTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
