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

// UserExtAppRoles A list of all AppRoles to which this User belongs directly, indirectly or implicitly. The User could belong directly because the User is a member of the AppRole, could belong indirectly because the User is a member of a Group that is a member of the AppRole, or could belong implicitly because the AppRole is public.
// **SCIM++ Properties:**
//   - idcsCompositeKey: [value]
//   - multiValued: true
//   - mutability: readOnly
//   - required: false
//   - returned: request
//   - type: complex
//   - uniqueness: none
type UserExtAppRoles struct {

	// The Id of the AppRole assigned to the User.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// The URI of the AppRole assigned to the User.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// The display name of the AppRole assigned to the User.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// The kind of membership this User has in the AppRole. A value of 'direct' indicates that the User is a member of the AppRole.  A value of  'indirect' indicates that the User is a member of a Group that is a member of the AppRole.  A value of 'implicit' indicates that the AppRole is public.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	Type UserExtAppRolesTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The ID of the App that defines this AppRole.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AppId *string `mandatory:"false" json:"appId"`

	// The name (Client ID) of the App that defines this AppRole.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AppName *string `mandatory:"false" json:"appName"`

	// If true, then the role provides administrative access privileges. READ-ONLY.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AdminRole *bool `mandatory:"false" json:"adminRole"`

	// The name (if any) under which this AppRole should appear in this User's group-memberships for reasons of backward compatibility. Oracle Identity Cloud Service distinguishes between Groups and AppRoles, but some services still expect AppRoles appear as if they were service-instance-specific Groups.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LegacyGroupName *string `mandatory:"false" json:"legacyGroupName"`
}

func (m UserExtAppRoles) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserExtAppRoles) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUserExtAppRolesTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUserExtAppRolesTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserExtAppRolesTypeEnum Enum with underlying type: string
type UserExtAppRolesTypeEnum string

// Set of constants representing the allowable values for UserExtAppRolesTypeEnum
const (
	UserExtAppRolesTypeDirect   UserExtAppRolesTypeEnum = "direct"
	UserExtAppRolesTypeIndirect UserExtAppRolesTypeEnum = "indirect"
	UserExtAppRolesTypeImplicit UserExtAppRolesTypeEnum = "implicit"
)

var mappingUserExtAppRolesTypeEnum = map[string]UserExtAppRolesTypeEnum{
	"direct":   UserExtAppRolesTypeDirect,
	"indirect": UserExtAppRolesTypeIndirect,
	"implicit": UserExtAppRolesTypeImplicit,
}

var mappingUserExtAppRolesTypeEnumLowerCase = map[string]UserExtAppRolesTypeEnum{
	"direct":   UserExtAppRolesTypeDirect,
	"indirect": UserExtAppRolesTypeIndirect,
	"implicit": UserExtAppRolesTypeImplicit,
}

// GetUserExtAppRolesTypeEnumValues Enumerates the set of values for UserExtAppRolesTypeEnum
func GetUserExtAppRolesTypeEnumValues() []UserExtAppRolesTypeEnum {
	values := make([]UserExtAppRolesTypeEnum, 0)
	for _, v := range mappingUserExtAppRolesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserExtAppRolesTypeEnumStringValues Enumerates the set of values in String for UserExtAppRolesTypeEnum
func GetUserExtAppRolesTypeEnumStringValues() []string {
	return []string{
		"direct",
		"indirect",
		"implicit",
	}
}

// GetMappingUserExtAppRolesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserExtAppRolesTypeEnum(val string) (UserExtAppRolesTypeEnum, bool) {
	enum, ok := mappingUserExtAppRolesTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
