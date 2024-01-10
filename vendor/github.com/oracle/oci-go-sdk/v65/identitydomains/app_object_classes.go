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

// AppObjectClasses Object classes
// **SCIM++ Properties:**
//   - idcsCompositeKey: [value]
//   - idcsSearchable: true
//   - multiValued: true
//   - mutability: readOnly
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type AppObjectClasses struct {

	// Object class template identifier
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// Object Class type. Allowed values are AccountObjectClass, ManagedObjectClass.
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsDefaultValue: AccountObjectClass
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type AppObjectClassesTypeEnum `mandatory:"true" json:"type"`

	// Object class URI
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// Object class display name
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// Object class resource type
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// If true, the object class represents an account. The isAccountObjectClass attribute value 'true' MUST appear no more than once.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsAccountObjectClass *bool `mandatory:"false" json:"isAccountObjectClass"`
}

func (m AppObjectClasses) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppObjectClasses) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAppObjectClassesTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAppObjectClassesTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppObjectClassesTypeEnum Enum with underlying type: string
type AppObjectClassesTypeEnum string

// Set of constants representing the allowable values for AppObjectClassesTypeEnum
const (
	AppObjectClassesTypeAccountobjectclass AppObjectClassesTypeEnum = "AccountObjectClass"
	AppObjectClassesTypeManagedobjectclass AppObjectClassesTypeEnum = "ManagedObjectClass"
)

var mappingAppObjectClassesTypeEnum = map[string]AppObjectClassesTypeEnum{
	"AccountObjectClass": AppObjectClassesTypeAccountobjectclass,
	"ManagedObjectClass": AppObjectClassesTypeManagedobjectclass,
}

var mappingAppObjectClassesTypeEnumLowerCase = map[string]AppObjectClassesTypeEnum{
	"accountobjectclass": AppObjectClassesTypeAccountobjectclass,
	"managedobjectclass": AppObjectClassesTypeManagedobjectclass,
}

// GetAppObjectClassesTypeEnumValues Enumerates the set of values for AppObjectClassesTypeEnum
func GetAppObjectClassesTypeEnumValues() []AppObjectClassesTypeEnum {
	values := make([]AppObjectClassesTypeEnum, 0)
	for _, v := range mappingAppObjectClassesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppObjectClassesTypeEnumStringValues Enumerates the set of values in String for AppObjectClassesTypeEnum
func GetAppObjectClassesTypeEnumStringValues() []string {
	return []string{
		"AccountObjectClass",
		"ManagedObjectClass",
	}
}

// GetMappingAppObjectClassesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppObjectClassesTypeEnum(val string) (AppObjectClassesTypeEnum, bool) {
	enum, ok := mappingAppObjectClassesTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
