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

// GroupExtOwners Group owners
// **SCIM++ Properties:**
//   - caseExact: false
//   - idcsCompositeKey: [value, type]
//   - idcsSearchable: true
//   - multiValued: true
//   - mutability: readWrite
//   - required: false
//   - returned: request
//   - type: complex
//   - uniqueness: none
type GroupExtOwners struct {

	// ID of the owner of this Group
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// Indicates the type of resource--for example, User or Group
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsDefaultValue: User
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type GroupExtOwnersTypeEnum `mandatory:"true" json:"type"`

	// The URI that corresponds to the owning Resource of this Group
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// Owner display name
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`
}

func (m GroupExtOwners) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GroupExtOwners) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGroupExtOwnersTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetGroupExtOwnersTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GroupExtOwnersTypeEnum Enum with underlying type: string
type GroupExtOwnersTypeEnum string

// Set of constants representing the allowable values for GroupExtOwnersTypeEnum
const (
	GroupExtOwnersTypeUser GroupExtOwnersTypeEnum = "User"
	GroupExtOwnersTypeApp  GroupExtOwnersTypeEnum = "App"
)

var mappingGroupExtOwnersTypeEnum = map[string]GroupExtOwnersTypeEnum{
	"User": GroupExtOwnersTypeUser,
	"App":  GroupExtOwnersTypeApp,
}

var mappingGroupExtOwnersTypeEnumLowerCase = map[string]GroupExtOwnersTypeEnum{
	"user": GroupExtOwnersTypeUser,
	"app":  GroupExtOwnersTypeApp,
}

// GetGroupExtOwnersTypeEnumValues Enumerates the set of values for GroupExtOwnersTypeEnum
func GetGroupExtOwnersTypeEnumValues() []GroupExtOwnersTypeEnum {
	values := make([]GroupExtOwnersTypeEnum, 0)
	for _, v := range mappingGroupExtOwnersTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGroupExtOwnersTypeEnumStringValues Enumerates the set of values in String for GroupExtOwnersTypeEnum
func GetGroupExtOwnersTypeEnumStringValues() []string {
	return []string{
		"User",
		"App",
	}
}

// GetMappingGroupExtOwnersTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGroupExtOwnersTypeEnum(val string) (GroupExtOwnersTypeEnum, bool) {
	enum, ok := mappingGroupExtOwnersTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
