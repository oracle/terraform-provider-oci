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

// UserExtDelegatedAuthenticationTargetApp If set, indicates the user's preferred authentication target app. If not set and the user's \"syncedFromApp\" is set and is enabled for delegated authentication, it is used. Otherwise, the user authenticates locally to Oracle Identity Cloud Service.
// **Added In:** 17.4.6
// **SCIM++ Properties:**
//   - idcsCompositeKey: [value]
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: readWrite
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type UserExtDelegatedAuthenticationTargetApp struct {

	// App identifier
	// **Added In:** 17.4.6
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
	// **Added In:** 17.4.6
	// **SCIM++ Properties:**
	//  - idcsDefaultValue: IdentitySource
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type UserExtDelegatedAuthenticationTargetAppTypeEnum `mandatory:"true" json:"type"`

	// App URI
	// **Added In:** 17.4.6
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
	// **Added In:** 17.4.6
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

func (m UserExtDelegatedAuthenticationTargetApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserExtDelegatedAuthenticationTargetApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserExtDelegatedAuthenticationTargetAppTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUserExtDelegatedAuthenticationTargetAppTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserExtDelegatedAuthenticationTargetAppTypeEnum Enum with underlying type: string
type UserExtDelegatedAuthenticationTargetAppTypeEnum string

// Set of constants representing the allowable values for UserExtDelegatedAuthenticationTargetAppTypeEnum
const (
	UserExtDelegatedAuthenticationTargetAppTypeApp            UserExtDelegatedAuthenticationTargetAppTypeEnum = "App"
	UserExtDelegatedAuthenticationTargetAppTypeIdentitysource UserExtDelegatedAuthenticationTargetAppTypeEnum = "IdentitySource"
)

var mappingUserExtDelegatedAuthenticationTargetAppTypeEnum = map[string]UserExtDelegatedAuthenticationTargetAppTypeEnum{
	"App":            UserExtDelegatedAuthenticationTargetAppTypeApp,
	"IdentitySource": UserExtDelegatedAuthenticationTargetAppTypeIdentitysource,
}

var mappingUserExtDelegatedAuthenticationTargetAppTypeEnumLowerCase = map[string]UserExtDelegatedAuthenticationTargetAppTypeEnum{
	"app":            UserExtDelegatedAuthenticationTargetAppTypeApp,
	"identitysource": UserExtDelegatedAuthenticationTargetAppTypeIdentitysource,
}

// GetUserExtDelegatedAuthenticationTargetAppTypeEnumValues Enumerates the set of values for UserExtDelegatedAuthenticationTargetAppTypeEnum
func GetUserExtDelegatedAuthenticationTargetAppTypeEnumValues() []UserExtDelegatedAuthenticationTargetAppTypeEnum {
	values := make([]UserExtDelegatedAuthenticationTargetAppTypeEnum, 0)
	for _, v := range mappingUserExtDelegatedAuthenticationTargetAppTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserExtDelegatedAuthenticationTargetAppTypeEnumStringValues Enumerates the set of values in String for UserExtDelegatedAuthenticationTargetAppTypeEnum
func GetUserExtDelegatedAuthenticationTargetAppTypeEnumStringValues() []string {
	return []string{
		"App",
		"IdentitySource",
	}
}

// GetMappingUserExtDelegatedAuthenticationTargetAppTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserExtDelegatedAuthenticationTargetAppTypeEnum(val string) (UserExtDelegatedAuthenticationTargetAppTypeEnum, bool) {
	enum, ok := mappingUserExtDelegatedAuthenticationTargetAppTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
