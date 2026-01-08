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

// MySessionRevokerRevokeBy Identifies the resource (by value/ocid and type), for which associated sessions are to be deleted
// **SCIM++ Properties:**
//   - caseExact: false
//   - idcsSearchable: false
//   - multiValued: false
//   - mutability: writeOnly
//   - required: true
//   - returned: never
//   - type: complex
//   - uniqueness: none
type MySessionRevokerRevokeBy struct {

	// **SCIM++ Properties:**
	// - caseExact: false
	// - multiValued: false
	// - required: true
	// - type: string
	// - uniqueness: none
	// - idcsRequiresWriteForAccessFlows: true
	// - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	// The type of the resource
	Type MySessionRevokerRevokeByTypeEnum `mandatory:"true" json:"type"`

	// **SCIM++ Properties:**
	// - caseExact: true
	// - multiValued: false
	// - required: false
	// - type: string
	// - uniqueness: none
	// - idcsRequiresWriteForAccessFlows: true
	// - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	// The identifier of the resource
	Value *string `mandatory:"false" json:"value"`

	// **SCIM++ Properties:**
	// - caseExact: true
	// - multiValued: false
	// - required: false
	// - type: string
	// - uniqueness: none
	// - idcsRequiresWriteForAccessFlows: true
	// - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	// The OCID of the resource
	Ocid *string `mandatory:"false" json:"ocid"`
}

func (m MySessionRevokerRevokeBy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySessionRevokerRevokeBy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMySessionRevokerRevokeByTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMySessionRevokerRevokeByTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MySessionRevokerRevokeByTypeEnum Enum with underlying type: string
type MySessionRevokerRevokeByTypeEnum string

// Set of constants representing the allowable values for MySessionRevokerRevokeByTypeEnum
const (
	MySessionRevokerRevokeByTypeSession MySessionRevokerRevokeByTypeEnum = "Session"
	MySessionRevokerRevokeByTypeApp     MySessionRevokerRevokeByTypeEnum = "App"
	MySessionRevokerRevokeByTypeUser    MySessionRevokerRevokeByTypeEnum = "User"
)

var mappingMySessionRevokerRevokeByTypeEnum = map[string]MySessionRevokerRevokeByTypeEnum{
	"Session": MySessionRevokerRevokeByTypeSession,
	"App":     MySessionRevokerRevokeByTypeApp,
	"User":    MySessionRevokerRevokeByTypeUser,
}

var mappingMySessionRevokerRevokeByTypeEnumLowerCase = map[string]MySessionRevokerRevokeByTypeEnum{
	"session": MySessionRevokerRevokeByTypeSession,
	"app":     MySessionRevokerRevokeByTypeApp,
	"user":    MySessionRevokerRevokeByTypeUser,
}

// GetMySessionRevokerRevokeByTypeEnumValues Enumerates the set of values for MySessionRevokerRevokeByTypeEnum
func GetMySessionRevokerRevokeByTypeEnumValues() []MySessionRevokerRevokeByTypeEnum {
	values := make([]MySessionRevokerRevokeByTypeEnum, 0)
	for _, v := range mappingMySessionRevokerRevokeByTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMySessionRevokerRevokeByTypeEnumStringValues Enumerates the set of values in String for MySessionRevokerRevokeByTypeEnum
func GetMySessionRevokerRevokeByTypeEnumStringValues() []string {
	return []string{
		"Session",
		"App",
		"User",
	}
}

// GetMappingMySessionRevokerRevokeByTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySessionRevokerRevokeByTypeEnum(val string) (MySessionRevokerRevokeByTypeEnum, bool) {
	enum, ok := mappingMySessionRevokerRevokeByTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
