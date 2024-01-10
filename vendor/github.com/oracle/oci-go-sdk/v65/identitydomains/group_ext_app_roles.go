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

// GroupExtAppRoles A list of appRoles that the user belongs to, either thorough direct membership, nested groups, or dynamically calculated
// **SCIM++ Properties:**
//   - caseExact: false
//   - idcsCompositeKey: [value]
//   - idcsSearchable: true
//   - multiValued: true
//   - mutability: readOnly
//   - required: false
//   - returned: request
//   - type: complex
//   - uniqueness: none
type GroupExtAppRoles struct {

	// The identifier of the appRole
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

	// The URI of the corresponding appRole resource to which the user belongs
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// A human readable name, primarily used for display purposes. READ-ONLY.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// A label indicating the attribute's function; e.g., 'direct' or 'indirect'.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	Type GroupExtAppRolesTypeEnum `mandatory:"false" json:"type,omitempty"`

	// ID of parent App. READ-ONLY.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	AppId *string `mandatory:"false" json:"appId"`

	// Name of parent App. READ-ONLY.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	AppName *string `mandatory:"false" json:"appName"`

	// If true, then the role provides administrative access privileges. READ-ONLY.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: boolean
	//  - uniqueness: none
	AdminRole *bool `mandatory:"false" json:"adminRole"`

	// The name of the legacy group associated with this AppRole.
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

func (m GroupExtAppRoles) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GroupExtAppRoles) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGroupExtAppRolesTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetGroupExtAppRolesTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GroupExtAppRolesTypeEnum Enum with underlying type: string
type GroupExtAppRolesTypeEnum string

// Set of constants representing the allowable values for GroupExtAppRolesTypeEnum
const (
	GroupExtAppRolesTypeDirect   GroupExtAppRolesTypeEnum = "direct"
	GroupExtAppRolesTypeIndirect GroupExtAppRolesTypeEnum = "indirect"
)

var mappingGroupExtAppRolesTypeEnum = map[string]GroupExtAppRolesTypeEnum{
	"direct":   GroupExtAppRolesTypeDirect,
	"indirect": GroupExtAppRolesTypeIndirect,
}

var mappingGroupExtAppRolesTypeEnumLowerCase = map[string]GroupExtAppRolesTypeEnum{
	"direct":   GroupExtAppRolesTypeDirect,
	"indirect": GroupExtAppRolesTypeIndirect,
}

// GetGroupExtAppRolesTypeEnumValues Enumerates the set of values for GroupExtAppRolesTypeEnum
func GetGroupExtAppRolesTypeEnumValues() []GroupExtAppRolesTypeEnum {
	values := make([]GroupExtAppRolesTypeEnum, 0)
	for _, v := range mappingGroupExtAppRolesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGroupExtAppRolesTypeEnumStringValues Enumerates the set of values in String for GroupExtAppRolesTypeEnum
func GetGroupExtAppRolesTypeEnumStringValues() []string {
	return []string{
		"direct",
		"indirect",
	}
}

// GetMappingGroupExtAppRolesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGroupExtAppRolesTypeEnum(val string) (GroupExtAppRolesTypeEnum, bool) {
	enum, ok := mappingGroupExtAppRolesTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
