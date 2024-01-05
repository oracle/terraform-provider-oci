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

// UserPhoneNumbers Phone numbers
type UserPhoneNumbers struct {

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
	Type UserPhoneNumbersTypeEnum `mandatory:"true" json:"type"`

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

func (m UserPhoneNumbers) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserPhoneNumbers) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserPhoneNumbersTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUserPhoneNumbersTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserPhoneNumbersTypeEnum Enum with underlying type: string
type UserPhoneNumbersTypeEnum string

// Set of constants representing the allowable values for UserPhoneNumbersTypeEnum
const (
	UserPhoneNumbersTypeWork     UserPhoneNumbersTypeEnum = "work"
	UserPhoneNumbersTypeHome     UserPhoneNumbersTypeEnum = "home"
	UserPhoneNumbersTypeMobile   UserPhoneNumbersTypeEnum = "mobile"
	UserPhoneNumbersTypeFax      UserPhoneNumbersTypeEnum = "fax"
	UserPhoneNumbersTypePager    UserPhoneNumbersTypeEnum = "pager"
	UserPhoneNumbersTypeOther    UserPhoneNumbersTypeEnum = "other"
	UserPhoneNumbersTypeRecovery UserPhoneNumbersTypeEnum = "recovery"
)

var mappingUserPhoneNumbersTypeEnum = map[string]UserPhoneNumbersTypeEnum{
	"work":     UserPhoneNumbersTypeWork,
	"home":     UserPhoneNumbersTypeHome,
	"mobile":   UserPhoneNumbersTypeMobile,
	"fax":      UserPhoneNumbersTypeFax,
	"pager":    UserPhoneNumbersTypePager,
	"other":    UserPhoneNumbersTypeOther,
	"recovery": UserPhoneNumbersTypeRecovery,
}

var mappingUserPhoneNumbersTypeEnumLowerCase = map[string]UserPhoneNumbersTypeEnum{
	"work":     UserPhoneNumbersTypeWork,
	"home":     UserPhoneNumbersTypeHome,
	"mobile":   UserPhoneNumbersTypeMobile,
	"fax":      UserPhoneNumbersTypeFax,
	"pager":    UserPhoneNumbersTypePager,
	"other":    UserPhoneNumbersTypeOther,
	"recovery": UserPhoneNumbersTypeRecovery,
}

// GetUserPhoneNumbersTypeEnumValues Enumerates the set of values for UserPhoneNumbersTypeEnum
func GetUserPhoneNumbersTypeEnumValues() []UserPhoneNumbersTypeEnum {
	values := make([]UserPhoneNumbersTypeEnum, 0)
	for _, v := range mappingUserPhoneNumbersTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserPhoneNumbersTypeEnumStringValues Enumerates the set of values in String for UserPhoneNumbersTypeEnum
func GetUserPhoneNumbersTypeEnumStringValues() []string {
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

// GetMappingUserPhoneNumbersTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserPhoneNumbersTypeEnum(val string) (UserPhoneNumbersTypeEnum, bool) {
	enum, ok := mappingUserPhoneNumbersTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
