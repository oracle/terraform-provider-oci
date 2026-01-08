// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SessionRevokerIdcsLockedBy The User or App who locked the Resource
// **SCIM++ Properties:**
//   - idcsSearchable: false
//   - multiValued: false
//   - mutability: readOnly
//   - required: false
//   - returned: default
//   - type: complex
type SessionRevokerIdcsLockedBy struct {

	// The ID of the SCIM resource that represents the User or App who LOCKED this Resource
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"false" json:"value"`

	// The OCID of the SCIM resource that represents the User or App who LOCKED this Resource
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Ocid *string `mandatory:"false" json:"ocid"`

	// The URI of the SCIM resource that represents the User or App who Locked this Resource
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// The type of resource, User or App, that locked this Resource
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type SessionRevokerIdcsLockedByTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The displayName of the User or App who locked this Resource
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
}

func (m SessionRevokerIdcsLockedBy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SessionRevokerIdcsLockedBy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSessionRevokerIdcsLockedByTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSessionRevokerIdcsLockedByTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SessionRevokerIdcsLockedByTypeEnum Enum with underlying type: string
type SessionRevokerIdcsLockedByTypeEnum string

// Set of constants representing the allowable values for SessionRevokerIdcsLockedByTypeEnum
const (
	SessionRevokerIdcsLockedByTypeUser SessionRevokerIdcsLockedByTypeEnum = "User"
	SessionRevokerIdcsLockedByTypeApp  SessionRevokerIdcsLockedByTypeEnum = "App"
)

var mappingSessionRevokerIdcsLockedByTypeEnum = map[string]SessionRevokerIdcsLockedByTypeEnum{
	"User": SessionRevokerIdcsLockedByTypeUser,
	"App":  SessionRevokerIdcsLockedByTypeApp,
}

var mappingSessionRevokerIdcsLockedByTypeEnumLowerCase = map[string]SessionRevokerIdcsLockedByTypeEnum{
	"user": SessionRevokerIdcsLockedByTypeUser,
	"app":  SessionRevokerIdcsLockedByTypeApp,
}

// GetSessionRevokerIdcsLockedByTypeEnumValues Enumerates the set of values for SessionRevokerIdcsLockedByTypeEnum
func GetSessionRevokerIdcsLockedByTypeEnumValues() []SessionRevokerIdcsLockedByTypeEnum {
	values := make([]SessionRevokerIdcsLockedByTypeEnum, 0)
	for _, v := range mappingSessionRevokerIdcsLockedByTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSessionRevokerIdcsLockedByTypeEnumStringValues Enumerates the set of values in String for SessionRevokerIdcsLockedByTypeEnum
func GetSessionRevokerIdcsLockedByTypeEnumStringValues() []string {
	return []string{
		"User",
		"App",
	}
}

// GetMappingSessionRevokerIdcsLockedByTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSessionRevokerIdcsLockedByTypeEnum(val string) (SessionRevokerIdcsLockedByTypeEnum, bool) {
	enum, ok := mappingSessionRevokerIdcsLockedByTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
