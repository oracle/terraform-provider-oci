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

// IdcsLastModifiedBy The User or App who modified the Resource
// **SCIM++ Properties:**
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: readOnly
//   - required: false
//   - returned: default
//   - type: complex
type IdcsLastModifiedBy struct {

	// The ID of the SCIM resource that represents the User or App who modified this Resource
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

	// The URI of the SCIM resource that represents the User or App who modified this Resource
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

	// The type of resource, User or App, that modified this Resource
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type IdcsLastModifiedByTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The displayName of the User or App who modified this Resource
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

	// The OCID of the SCIM resource that represents the User or App who modified this Resource
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

func (m IdcsLastModifiedBy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdcsLastModifiedBy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIdcsLastModifiedByTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetIdcsLastModifiedByTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IdcsLastModifiedByTypeEnum Enum with underlying type: string
type IdcsLastModifiedByTypeEnum string

// Set of constants representing the allowable values for IdcsLastModifiedByTypeEnum
const (
	IdcsLastModifiedByTypeUser IdcsLastModifiedByTypeEnum = "User"
	IdcsLastModifiedByTypeApp  IdcsLastModifiedByTypeEnum = "App"
)

var mappingIdcsLastModifiedByTypeEnum = map[string]IdcsLastModifiedByTypeEnum{
	"User": IdcsLastModifiedByTypeUser,
	"App":  IdcsLastModifiedByTypeApp,
}

var mappingIdcsLastModifiedByTypeEnumLowerCase = map[string]IdcsLastModifiedByTypeEnum{
	"user": IdcsLastModifiedByTypeUser,
	"app":  IdcsLastModifiedByTypeApp,
}

// GetIdcsLastModifiedByTypeEnumValues Enumerates the set of values for IdcsLastModifiedByTypeEnum
func GetIdcsLastModifiedByTypeEnumValues() []IdcsLastModifiedByTypeEnum {
	values := make([]IdcsLastModifiedByTypeEnum, 0)
	for _, v := range mappingIdcsLastModifiedByTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIdcsLastModifiedByTypeEnumStringValues Enumerates the set of values in String for IdcsLastModifiedByTypeEnum
func GetIdcsLastModifiedByTypeEnumStringValues() []string {
	return []string{
		"User",
		"App",
	}
}

// GetMappingIdcsLastModifiedByTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdcsLastModifiedByTypeEnum(val string) (IdcsLastModifiedByTypeEnum, bool) {
	enum, ok := mappingIdcsLastModifiedByTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
