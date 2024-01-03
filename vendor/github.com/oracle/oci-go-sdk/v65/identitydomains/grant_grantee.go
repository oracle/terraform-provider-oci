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

// GrantGrantee Grantee beneficiary. The grantee may be a User, Group, App or DynamicResourceGroup.
// **SCIM++ Properties:**
//   - idcsCsvAttributeNameMappings: [[columnHeaderName:Grantee Name, csvColumnForResolvingResourceType:Grantee Type, mapsTo:grantee.value], [columnHeaderName:Grantee Type, mapsTo:grantee.type]]
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: immutable
//   - required: true
//   - returned: default
//   - type: complex
//   - uniqueness: none
type GrantGrantee struct {

	// Grantee identifier
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCsvAttributeName: Member
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// Grantee resource type. Allowed values are User, Group, App and DynamicResourceGroup.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCsvAttributeName: Member Type
	//  - idcsDefaultValue: User
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type GrantGranteeTypeEnum `mandatory:"true" json:"type"`

	// Grantee URI
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// Grantee display name
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`
}

func (m GrantGrantee) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GrantGrantee) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGrantGranteeTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetGrantGranteeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GrantGranteeTypeEnum Enum with underlying type: string
type GrantGranteeTypeEnum string

// Set of constants representing the allowable values for GrantGranteeTypeEnum
const (
	GrantGranteeTypeUser                 GrantGranteeTypeEnum = "User"
	GrantGranteeTypeGroup                GrantGranteeTypeEnum = "Group"
	GrantGranteeTypeApp                  GrantGranteeTypeEnum = "App"
	GrantGranteeTypeDynamicresourcegroup GrantGranteeTypeEnum = "DynamicResourceGroup"
)

var mappingGrantGranteeTypeEnum = map[string]GrantGranteeTypeEnum{
	"User":                 GrantGranteeTypeUser,
	"Group":                GrantGranteeTypeGroup,
	"App":                  GrantGranteeTypeApp,
	"DynamicResourceGroup": GrantGranteeTypeDynamicresourcegroup,
}

var mappingGrantGranteeTypeEnumLowerCase = map[string]GrantGranteeTypeEnum{
	"user":                 GrantGranteeTypeUser,
	"group":                GrantGranteeTypeGroup,
	"app":                  GrantGranteeTypeApp,
	"dynamicresourcegroup": GrantGranteeTypeDynamicresourcegroup,
}

// GetGrantGranteeTypeEnumValues Enumerates the set of values for GrantGranteeTypeEnum
func GetGrantGranteeTypeEnumValues() []GrantGranteeTypeEnum {
	values := make([]GrantGranteeTypeEnum, 0)
	for _, v := range mappingGrantGranteeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGrantGranteeTypeEnumStringValues Enumerates the set of values in String for GrantGranteeTypeEnum
func GetGrantGranteeTypeEnumStringValues() []string {
	return []string{
		"User",
		"Group",
		"App",
		"DynamicResourceGroup",
	}
}

// GetMappingGrantGranteeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGrantGranteeTypeEnum(val string) (GrantGranteeTypeEnum, bool) {
	enum, ok := mappingGrantGranteeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
