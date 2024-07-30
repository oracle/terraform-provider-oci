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

// MyGroupMembers The group members. <b>Important:</b> When requesting group members, a maximum of 10,000 members can be returned in a single request. If the response contains more than 10,000 members, the request will fail. Use 'startIndex' and 'count' to return members in pages instead of in a single response, for example: #attributes=members[startIndex=1%26count=10]. This REST API is SCIM compliant.
type MyGroupMembers struct {

	// The ID of the member of this Group
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// Indicates the type of resource, for example, User or Group.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - idcsDefaultValue: User
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type MyGroupMembersTypeEnum `mandatory:"true" json:"type"`

	// The date and time that the member was added to the group.
	// **Added In:** 2012271618
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	DateAdded *string `mandatory:"false" json:"dateAdded"`

	// The OCID of the member of this group.
	// **Added In:** 2012271618
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Ocid *string `mandatory:"false" json:"ocid"`

	// The membership OCID.
	// **Added In:** 2102181953
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	MembershipOcid *string `mandatory:"false" json:"membershipOcid"`

	// The URI that corresponds to the member Resource of this group.
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

	// The member's display name.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// The member's name.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Name *string `mandatory:"false" json:"name"`
}

func (m MyGroupMembers) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MyGroupMembers) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMyGroupMembersTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMyGroupMembersTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MyGroupMembersTypeEnum Enum with underlying type: string
type MyGroupMembersTypeEnum string

// Set of constants representing the allowable values for MyGroupMembersTypeEnum
const (
	MyGroupMembersTypeUser MyGroupMembersTypeEnum = "User"
)

var mappingMyGroupMembersTypeEnum = map[string]MyGroupMembersTypeEnum{
	"User": MyGroupMembersTypeUser,
}

var mappingMyGroupMembersTypeEnumLowerCase = map[string]MyGroupMembersTypeEnum{
	"user": MyGroupMembersTypeUser,
}

// GetMyGroupMembersTypeEnumValues Enumerates the set of values for MyGroupMembersTypeEnum
func GetMyGroupMembersTypeEnumValues() []MyGroupMembersTypeEnum {
	values := make([]MyGroupMembersTypeEnum, 0)
	for _, v := range mappingMyGroupMembersTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMyGroupMembersTypeEnumStringValues Enumerates the set of values in String for MyGroupMembersTypeEnum
func GetMyGroupMembersTypeEnumStringValues() []string {
	return []string{
		"User",
	}
}

// GetMappingMyGroupMembersTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyGroupMembersTypeEnum(val string) (MyGroupMembersTypeEnum, bool) {
	enum, ok := mappingMyGroupMembersTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
