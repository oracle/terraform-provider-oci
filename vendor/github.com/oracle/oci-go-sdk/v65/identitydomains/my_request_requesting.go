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

// MyRequestRequesting Requestable resource reference.
// **SCIM++ Properties:**
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: immutable
//   - required: true
//   - returned: default
//   - type: complex
//   - uniqueness: none
type MyRequestRequesting struct {

	// Resource identifier
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCsvAttributeName: requesting_id
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// Requestable type. Allowed values are Group and App.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCsvAttributeName: Requestable Type
	//  - idcsDefaultValue: Group
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type MyRequestRequestingTypeEnum `mandatory:"true" json:"type"`

	// Resource URI
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// Resource display name
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// Resource description
	// **Added In:** 2307071836
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Description *string `mandatory:"false" json:"description"`
}

func (m MyRequestRequesting) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MyRequestRequesting) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMyRequestRequestingTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMyRequestRequestingTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MyRequestRequestingTypeEnum Enum with underlying type: string
type MyRequestRequestingTypeEnum string

// Set of constants representing the allowable values for MyRequestRequestingTypeEnum
const (
	MyRequestRequestingTypeGroup MyRequestRequestingTypeEnum = "Group"
	MyRequestRequestingTypeApp   MyRequestRequestingTypeEnum = "App"
)

var mappingMyRequestRequestingTypeEnum = map[string]MyRequestRequestingTypeEnum{
	"Group": MyRequestRequestingTypeGroup,
	"App":   MyRequestRequestingTypeApp,
}

var mappingMyRequestRequestingTypeEnumLowerCase = map[string]MyRequestRequestingTypeEnum{
	"group": MyRequestRequestingTypeGroup,
	"app":   MyRequestRequestingTypeApp,
}

// GetMyRequestRequestingTypeEnumValues Enumerates the set of values for MyRequestRequestingTypeEnum
func GetMyRequestRequestingTypeEnumValues() []MyRequestRequestingTypeEnum {
	values := make([]MyRequestRequestingTypeEnum, 0)
	for _, v := range mappingMyRequestRequestingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMyRequestRequestingTypeEnumStringValues Enumerates the set of values in String for MyRequestRequestingTypeEnum
func GetMyRequestRequestingTypeEnumStringValues() []string {
	return []string{
		"Group",
		"App",
	}
}

// GetMappingMyRequestRequestingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyRequestRequestingTypeEnum(val string) (MyRequestRequestingTypeEnum, bool) {
	enum, ok := mappingMyRequestRequestingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
