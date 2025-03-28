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

// MeGroups A list of groups that the user belongs to, either thorough direct membership, nested groups, or dynamically calculated
type MeGroups struct {

	// The identifier of the User's group.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// The OCID of the User's group.
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
	Ocid *string `mandatory:"false" json:"ocid"`

	// The URI of the corresponding Group resource to which the user belongs
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// A human readable name, primarily used for display purposes. READ-ONLY.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// A human readable name for Group as defined by the Service Consumer. READ-ONLY.
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	NonUniqueDisplay *string `mandatory:"false" json:"nonUniqueDisplay"`

	// An identifier for the Resource as defined by the Service Consumer. READ-ONLY.
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ExternalId *string `mandatory:"false" json:"externalId"`

	// A label indicating the attribute's function; e.g., 'direct' or 'indirect'.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	Type MeGroupsTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The membership OCID.
	// **Added In:** 2103141444
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	MembershipOcid *string `mandatory:"false" json:"membershipOcid"`

	// Date when the member is Added to the group
	// **Added In:** 2105200541
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	DateAdded *string `mandatory:"false" json:"dateAdded"`
}

func (m MeGroups) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MeGroups) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMeGroupsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMeGroupsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MeGroupsTypeEnum Enum with underlying type: string
type MeGroupsTypeEnum string

// Set of constants representing the allowable values for MeGroupsTypeEnum
const (
	MeGroupsTypeDirect   MeGroupsTypeEnum = "direct"
	MeGroupsTypeIndirect MeGroupsTypeEnum = "indirect"
)

var mappingMeGroupsTypeEnum = map[string]MeGroupsTypeEnum{
	"direct":   MeGroupsTypeDirect,
	"indirect": MeGroupsTypeIndirect,
}

var mappingMeGroupsTypeEnumLowerCase = map[string]MeGroupsTypeEnum{
	"direct":   MeGroupsTypeDirect,
	"indirect": MeGroupsTypeIndirect,
}

// GetMeGroupsTypeEnumValues Enumerates the set of values for MeGroupsTypeEnum
func GetMeGroupsTypeEnumValues() []MeGroupsTypeEnum {
	values := make([]MeGroupsTypeEnum, 0)
	for _, v := range mappingMeGroupsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMeGroupsTypeEnumStringValues Enumerates the set of values in String for MeGroupsTypeEnum
func GetMeGroupsTypeEnumStringValues() []string {
	return []string{
		"direct",
		"indirect",
	}
}

// GetMappingMeGroupsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMeGroupsTypeEnum(val string) (MeGroupsTypeEnum, bool) {
	enum, ok := mappingMeGroupsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
