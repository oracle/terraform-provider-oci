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

// GrantGrantor User conferring the grant to the beneficiary
// **SCIM++ Properties:**
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: readOnly
//   - idcsIgnoreReadOnlyAndImmutableRefAttrsDuringForceDelete: true
//   - required: false
//   - returned: default
//   - type: complex
type GrantGrantor struct {

	// Resource type of the grantor. Allowed values are User and App.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsDefaultValue: User
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type GrantGrantorTypeEnum `mandatory:"true" json:"type"`

	// Grantor user identifier
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"false" json:"value"`

	// Grantor URI
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// Grantor display name
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

func (m GrantGrantor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GrantGrantor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGrantGrantorTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetGrantGrantorTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GrantGrantorTypeEnum Enum with underlying type: string
type GrantGrantorTypeEnum string

// Set of constants representing the allowable values for GrantGrantorTypeEnum
const (
	GrantGrantorTypeUser                     GrantGrantorTypeEnum = "User"
	GrantGrantorTypeApp                      GrantGrantorTypeEnum = "App"
	GrantGrantorTypeGroup                    GrantGrantorTypeEnum = "Group"
	GrantGrantorTypeAppentitlementcollection GrantGrantorTypeEnum = "AppEntitlementCollection"
	GrantGrantorTypeDynamicresourcegroup     GrantGrantorTypeEnum = "DynamicResourceGroup"
)

var mappingGrantGrantorTypeEnum = map[string]GrantGrantorTypeEnum{
	"User":                     GrantGrantorTypeUser,
	"App":                      GrantGrantorTypeApp,
	"Group":                    GrantGrantorTypeGroup,
	"AppEntitlementCollection": GrantGrantorTypeAppentitlementcollection,
	"DynamicResourceGroup":     GrantGrantorTypeDynamicresourcegroup,
}

var mappingGrantGrantorTypeEnumLowerCase = map[string]GrantGrantorTypeEnum{
	"user":                     GrantGrantorTypeUser,
	"app":                      GrantGrantorTypeApp,
	"group":                    GrantGrantorTypeGroup,
	"appentitlementcollection": GrantGrantorTypeAppentitlementcollection,
	"dynamicresourcegroup":     GrantGrantorTypeDynamicresourcegroup,
}

// GetGrantGrantorTypeEnumValues Enumerates the set of values for GrantGrantorTypeEnum
func GetGrantGrantorTypeEnumValues() []GrantGrantorTypeEnum {
	values := make([]GrantGrantorTypeEnum, 0)
	for _, v := range mappingGrantGrantorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGrantGrantorTypeEnumStringValues Enumerates the set of values in String for GrantGrantorTypeEnum
func GetGrantGrantorTypeEnumStringValues() []string {
	return []string{
		"User",
		"App",
		"Group",
		"AppEntitlementCollection",
		"DynamicResourceGroup",
	}
}

// GetMappingGrantGrantorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGrantGrantorTypeEnum(val string) (GrantGrantorTypeEnum, bool) {
	enum, ok := mappingGrantGrantorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
