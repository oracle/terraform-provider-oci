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

// MePhoneNumbers Phone numbers
type MePhoneNumbers struct {

	// User's phone number
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// A label that indicates the attribute's function- for example, 'work', 'home', or 'mobile'
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type MePhoneNumbersTypeEnum `mandatory:"true" json:"type"`

	// A human-readable name, primarily used for display purposes. READ ONLY
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

	// A Boolean value that indicates the 'primary' or preferred attribute value for this attribute--for example, the preferred phone number or primary phone number. The primary attribute value 'true' MUST appear no more than once.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Primary *bool `mandatory:"false" json:"primary"`

	// A Boolean value that indicates if the phone number is verified.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Verified *bool `mandatory:"false" json:"verified"`
}

func (m MePhoneNumbers) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MePhoneNumbers) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMePhoneNumbersTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMePhoneNumbersTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MePhoneNumbersTypeEnum Enum with underlying type: string
type MePhoneNumbersTypeEnum string

// Set of constants representing the allowable values for MePhoneNumbersTypeEnum
const (
	MePhoneNumbersTypeWork     MePhoneNumbersTypeEnum = "work"
	MePhoneNumbersTypeHome     MePhoneNumbersTypeEnum = "home"
	MePhoneNumbersTypeMobile   MePhoneNumbersTypeEnum = "mobile"
	MePhoneNumbersTypeFax      MePhoneNumbersTypeEnum = "fax"
	MePhoneNumbersTypePager    MePhoneNumbersTypeEnum = "pager"
	MePhoneNumbersTypeOther    MePhoneNumbersTypeEnum = "other"
	MePhoneNumbersTypeRecovery MePhoneNumbersTypeEnum = "recovery"
)

var mappingMePhoneNumbersTypeEnum = map[string]MePhoneNumbersTypeEnum{
	"work":     MePhoneNumbersTypeWork,
	"home":     MePhoneNumbersTypeHome,
	"mobile":   MePhoneNumbersTypeMobile,
	"fax":      MePhoneNumbersTypeFax,
	"pager":    MePhoneNumbersTypePager,
	"other":    MePhoneNumbersTypeOther,
	"recovery": MePhoneNumbersTypeRecovery,
}

var mappingMePhoneNumbersTypeEnumLowerCase = map[string]MePhoneNumbersTypeEnum{
	"work":     MePhoneNumbersTypeWork,
	"home":     MePhoneNumbersTypeHome,
	"mobile":   MePhoneNumbersTypeMobile,
	"fax":      MePhoneNumbersTypeFax,
	"pager":    MePhoneNumbersTypePager,
	"other":    MePhoneNumbersTypeOther,
	"recovery": MePhoneNumbersTypeRecovery,
}

// GetMePhoneNumbersTypeEnumValues Enumerates the set of values for MePhoneNumbersTypeEnum
func GetMePhoneNumbersTypeEnumValues() []MePhoneNumbersTypeEnum {
	values := make([]MePhoneNumbersTypeEnum, 0)
	for _, v := range mappingMePhoneNumbersTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMePhoneNumbersTypeEnumStringValues Enumerates the set of values in String for MePhoneNumbersTypeEnum
func GetMePhoneNumbersTypeEnumStringValues() []string {
	return []string{
		"work",
		"home",
		"mobile",
		"fax",
		"pager",
		"other",
		"recovery",
	}
}

// GetMappingMePhoneNumbersTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMePhoneNumbersTypeEnum(val string) (MePhoneNumbersTypeEnum, bool) {
	enum, ok := mappingMePhoneNumbersTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
