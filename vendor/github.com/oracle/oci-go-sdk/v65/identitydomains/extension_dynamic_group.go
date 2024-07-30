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

// ExtensionDynamicGroup Dynamic Group
type ExtensionDynamicGroup struct {

	// Membership type
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	MembershipType ExtensionDynamicGroupMembershipTypeEnum `mandatory:"false" json:"membershipType,omitempty"`

	// Membership rule
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	MembershipRule *string `mandatory:"false" json:"membershipRule"`
}

func (m ExtensionDynamicGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionDynamicGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExtensionDynamicGroupMembershipTypeEnum(string(m.MembershipType)); !ok && m.MembershipType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MembershipType: %s. Supported values are: %s.", m.MembershipType, strings.Join(GetExtensionDynamicGroupMembershipTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExtensionDynamicGroupMembershipTypeEnum Enum with underlying type: string
type ExtensionDynamicGroupMembershipTypeEnum string

// Set of constants representing the allowable values for ExtensionDynamicGroupMembershipTypeEnum
const (
	ExtensionDynamicGroupMembershipTypeStatic  ExtensionDynamicGroupMembershipTypeEnum = "static"
	ExtensionDynamicGroupMembershipTypeDynamic ExtensionDynamicGroupMembershipTypeEnum = "dynamic"
)

var mappingExtensionDynamicGroupMembershipTypeEnum = map[string]ExtensionDynamicGroupMembershipTypeEnum{
	"static":  ExtensionDynamicGroupMembershipTypeStatic,
	"dynamic": ExtensionDynamicGroupMembershipTypeDynamic,
}

var mappingExtensionDynamicGroupMembershipTypeEnumLowerCase = map[string]ExtensionDynamicGroupMembershipTypeEnum{
	"static":  ExtensionDynamicGroupMembershipTypeStatic,
	"dynamic": ExtensionDynamicGroupMembershipTypeDynamic,
}

// GetExtensionDynamicGroupMembershipTypeEnumValues Enumerates the set of values for ExtensionDynamicGroupMembershipTypeEnum
func GetExtensionDynamicGroupMembershipTypeEnumValues() []ExtensionDynamicGroupMembershipTypeEnum {
	values := make([]ExtensionDynamicGroupMembershipTypeEnum, 0)
	for _, v := range mappingExtensionDynamicGroupMembershipTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionDynamicGroupMembershipTypeEnumStringValues Enumerates the set of values in String for ExtensionDynamicGroupMembershipTypeEnum
func GetExtensionDynamicGroupMembershipTypeEnumStringValues() []string {
	return []string{
		"static",
		"dynamic",
	}
}

// GetMappingExtensionDynamicGroupMembershipTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionDynamicGroupMembershipTypeEnum(val string) (ExtensionDynamicGroupMembershipTypeEnum, bool) {
	enum, ok := mappingExtensionDynamicGroupMembershipTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
