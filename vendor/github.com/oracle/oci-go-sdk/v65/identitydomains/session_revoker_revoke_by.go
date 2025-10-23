// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// SessionRevokerRevokeBy Identifies the resource (by value/ocid and type), for which associated sessions are to be deleted
// **SCIM++ Properties:**
//   - caseExact: false
//   - idcsSearchable: false
//   - multiValued: false
//   - mutability: writeOnly
//   - required: true
//   - returned: never
//   - type: complex
//   - uniqueness: none
type SessionRevokerRevokeBy struct {

	// **SCIM++ Properties:**
	// - caseExact: false
	// - multiValued: false
	// - required: true
	// - type: string
	// - uniqueness: none
	// - idcsRequiresWriteForAccessFlows: true
	// - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	// The type of the resource
	Type SessionRevokerRevokeByTypeEnum `mandatory:"true" json:"type"`

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

func (m SessionRevokerRevokeBy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SessionRevokerRevokeBy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSessionRevokerRevokeByTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSessionRevokerRevokeByTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SessionRevokerRevokeByTypeEnum Enum with underlying type: string
type SessionRevokerRevokeByTypeEnum string

// Set of constants representing the allowable values for SessionRevokerRevokeByTypeEnum
const (
	SessionRevokerRevokeByTypeSession SessionRevokerRevokeByTypeEnum = "Session"
	SessionRevokerRevokeByTypeApp     SessionRevokerRevokeByTypeEnum = "App"
	SessionRevokerRevokeByTypeUser    SessionRevokerRevokeByTypeEnum = "User"
)

var mappingSessionRevokerRevokeByTypeEnum = map[string]SessionRevokerRevokeByTypeEnum{
	"Session": SessionRevokerRevokeByTypeSession,
	"App":     SessionRevokerRevokeByTypeApp,
	"User":    SessionRevokerRevokeByTypeUser,
}

var mappingSessionRevokerRevokeByTypeEnumLowerCase = map[string]SessionRevokerRevokeByTypeEnum{
	"session": SessionRevokerRevokeByTypeSession,
	"app":     SessionRevokerRevokeByTypeApp,
	"user":    SessionRevokerRevokeByTypeUser,
}

// GetSessionRevokerRevokeByTypeEnumValues Enumerates the set of values for SessionRevokerRevokeByTypeEnum
func GetSessionRevokerRevokeByTypeEnumValues() []SessionRevokerRevokeByTypeEnum {
	values := make([]SessionRevokerRevokeByTypeEnum, 0)
	for _, v := range mappingSessionRevokerRevokeByTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSessionRevokerRevokeByTypeEnumStringValues Enumerates the set of values in String for SessionRevokerRevokeByTypeEnum
func GetSessionRevokerRevokeByTypeEnumStringValues() []string {
	return []string{
		"Session",
		"App",
		"User",
	}
}

// GetMappingSessionRevokerRevokeByTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSessionRevokerRevokeByTypeEnum(val string) (SessionRevokerRevokeByTypeEnum, bool) {
	enum, ok := mappingSessionRevokerRevokeByTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
