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

// IdcsCreatedBy The User or App who created the Resource
// **SCIM++ Properties:**
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: readOnly
//   - required: true
//   - returned: default
//   - type: complex
type IdcsCreatedBy struct {

	// The ID of the SCIM resource that represents the User or App who created this Resource
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

	// The URI of the SCIM resource that represents the User or App who created this Resource
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

	// The type of resource, User or App, that created this Resource
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type IdcsCreatedByTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The displayName of the User or App who created this Resource
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// The OCID of the SCIM resource that represents the User or App who created this Resource
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Ocid *string `mandatory:"false" json:"ocid"`
}

func (m IdcsCreatedBy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdcsCreatedBy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIdcsCreatedByTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetIdcsCreatedByTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IdcsCreatedByTypeEnum Enum with underlying type: string
type IdcsCreatedByTypeEnum string

// Set of constants representing the allowable values for IdcsCreatedByTypeEnum
const (
	IdcsCreatedByTypeUser IdcsCreatedByTypeEnum = "User"
	IdcsCreatedByTypeApp  IdcsCreatedByTypeEnum = "App"
)

var mappingIdcsCreatedByTypeEnum = map[string]IdcsCreatedByTypeEnum{
	"User": IdcsCreatedByTypeUser,
	"App":  IdcsCreatedByTypeApp,
}

var mappingIdcsCreatedByTypeEnumLowerCase = map[string]IdcsCreatedByTypeEnum{
	"user": IdcsCreatedByTypeUser,
	"app":  IdcsCreatedByTypeApp,
}

// GetIdcsCreatedByTypeEnumValues Enumerates the set of values for IdcsCreatedByTypeEnum
func GetIdcsCreatedByTypeEnumValues() []IdcsCreatedByTypeEnum {
	values := make([]IdcsCreatedByTypeEnum, 0)
	for _, v := range mappingIdcsCreatedByTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIdcsCreatedByTypeEnumStringValues Enumerates the set of values in String for IdcsCreatedByTypeEnum
func GetIdcsCreatedByTypeEnumStringValues() []string {
	return []string{
		"User",
		"App",
	}
}

// GetMappingIdcsCreatedByTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdcsCreatedByTypeEnum(val string) (IdcsCreatedByTypeEnum, bool) {
	enum, ok := mappingIdcsCreatedByTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
