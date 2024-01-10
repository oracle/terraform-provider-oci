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

// AppRoleMembers AppRole members - when requesting members attribute, it is recommended to use startIndex and count to return members in pages instead of in a single response, eg : #attributes=members[startIndex=1%26count=10]
type AppRoleMembers struct {

	// ID of the member of this AppRole
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCsvAttributeName: Member
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// Indicates the type of Resource--for example, User, Group or DynamicResourceGroup
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCsvAttributeName: Member Type
	//  - idcsDefaultValue: User
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type AppRoleMembersTypeEnum `mandatory:"true" json:"type"`

	// The URI corresponding to the member Resource of this Group
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// Member display name
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`
}

func (m AppRoleMembers) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppRoleMembers) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAppRoleMembersTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAppRoleMembersTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppRoleMembersTypeEnum Enum with underlying type: string
type AppRoleMembersTypeEnum string

// Set of constants representing the allowable values for AppRoleMembersTypeEnum
const (
	AppRoleMembersTypeUser                 AppRoleMembersTypeEnum = "User"
	AppRoleMembersTypeGroup                AppRoleMembersTypeEnum = "Group"
	AppRoleMembersTypeDynamicresourcegroup AppRoleMembersTypeEnum = "DynamicResourceGroup"
)

var mappingAppRoleMembersTypeEnum = map[string]AppRoleMembersTypeEnum{
	"User":                 AppRoleMembersTypeUser,
	"Group":                AppRoleMembersTypeGroup,
	"DynamicResourceGroup": AppRoleMembersTypeDynamicresourcegroup,
}

var mappingAppRoleMembersTypeEnumLowerCase = map[string]AppRoleMembersTypeEnum{
	"user":                 AppRoleMembersTypeUser,
	"group":                AppRoleMembersTypeGroup,
	"dynamicresourcegroup": AppRoleMembersTypeDynamicresourcegroup,
}

// GetAppRoleMembersTypeEnumValues Enumerates the set of values for AppRoleMembersTypeEnum
func GetAppRoleMembersTypeEnumValues() []AppRoleMembersTypeEnum {
	values := make([]AppRoleMembersTypeEnum, 0)
	for _, v := range mappingAppRoleMembersTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppRoleMembersTypeEnumStringValues Enumerates the set of values in String for AppRoleMembersTypeEnum
func GetAppRoleMembersTypeEnumStringValues() []string {
	return []string{
		"User",
		"Group",
		"DynamicResourceGroup",
	}
}

// GetMappingAppRoleMembersTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppRoleMembersTypeEnum(val string) (AppRoleMembersTypeEnum, bool) {
	enum, ok := mappingAppRoleMembersTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
