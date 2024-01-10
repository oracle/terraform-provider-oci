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

// GroupExtSyncedFromApp The entity that created this Group.
// **Added In:** 18.4.2
// **SCIM++ Properties:**
//   - idcsCompositeKey: [value]
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: readOnly
//   - required: false
//   - returned: request
//   - type: complex
//   - uniqueness: none
type GroupExtSyncedFromApp struct {

	// The ID of the App.
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// The type of the entity that created this Group.
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - idcsDefaultValue: App
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type GroupExtSyncedFromAppTypeEnum `mandatory:"true" json:"type"`

	// App URI
	// **Added In:** 18.4.2
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
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`
}

func (m GroupExtSyncedFromApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GroupExtSyncedFromApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGroupExtSyncedFromAppTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetGroupExtSyncedFromAppTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GroupExtSyncedFromAppTypeEnum Enum with underlying type: string
type GroupExtSyncedFromAppTypeEnum string

// Set of constants representing the allowable values for GroupExtSyncedFromAppTypeEnum
const (
	GroupExtSyncedFromAppTypeApp GroupExtSyncedFromAppTypeEnum = "App"
)

var mappingGroupExtSyncedFromAppTypeEnum = map[string]GroupExtSyncedFromAppTypeEnum{
	"App": GroupExtSyncedFromAppTypeApp,
}

var mappingGroupExtSyncedFromAppTypeEnumLowerCase = map[string]GroupExtSyncedFromAppTypeEnum{
	"app": GroupExtSyncedFromAppTypeApp,
}

// GetGroupExtSyncedFromAppTypeEnumValues Enumerates the set of values for GroupExtSyncedFromAppTypeEnum
func GetGroupExtSyncedFromAppTypeEnumValues() []GroupExtSyncedFromAppTypeEnum {
	values := make([]GroupExtSyncedFromAppTypeEnum, 0)
	for _, v := range mappingGroupExtSyncedFromAppTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGroupExtSyncedFromAppTypeEnumStringValues Enumerates the set of values in String for GroupExtSyncedFromAppTypeEnum
func GetGroupExtSyncedFromAppTypeEnumStringValues() []string {
	return []string{
		"App",
	}
}

// GetMappingGroupExtSyncedFromAppTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGroupExtSyncedFromAppTypeEnum(val string) (GroupExtSyncedFromAppTypeEnum, bool) {
	enum, ok := mappingGroupExtSyncedFromAppTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
