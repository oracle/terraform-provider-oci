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

// UserEmails A complex attribute representing emails
type UserEmails struct {

	// Email address
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

	// Type of email address
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type UserEmailsTypeEnum `mandatory:"true" json:"type"`

	// A Boolean value that indicates whether the email address is the primary email address. The primary attribute value 'true' MUST appear no more than once.
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

	// A Boolean value that indicates whether the email address is the secondary email address. The secondary attribute value 'true' MUST appear no more than once.
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Secondary *bool `mandatory:"false" json:"secondary"`

	// A Boolean value that indicates whether or not the e-mail address is verified
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Verified *bool `mandatory:"false" json:"verified"`

	// Pending e-mail address verification
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PendingVerificationData *string `mandatory:"false" json:"pendingVerificationData"`
}

func (m UserEmails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserEmails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserEmailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUserEmailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserEmailsTypeEnum Enum with underlying type: string
type UserEmailsTypeEnum string

// Set of constants representing the allowable values for UserEmailsTypeEnum
const (
	UserEmailsTypeWork     UserEmailsTypeEnum = "work"
	UserEmailsTypeHome     UserEmailsTypeEnum = "home"
	UserEmailsTypeOther    UserEmailsTypeEnum = "other"
	UserEmailsTypeRecovery UserEmailsTypeEnum = "recovery"
)

var mappingUserEmailsTypeEnum = map[string]UserEmailsTypeEnum{
	"work":     UserEmailsTypeWork,
	"home":     UserEmailsTypeHome,
	"other":    UserEmailsTypeOther,
	"recovery": UserEmailsTypeRecovery,
}

var mappingUserEmailsTypeEnumLowerCase = map[string]UserEmailsTypeEnum{
	"work":     UserEmailsTypeWork,
	"home":     UserEmailsTypeHome,
	"other":    UserEmailsTypeOther,
	"recovery": UserEmailsTypeRecovery,
}

// GetUserEmailsTypeEnumValues Enumerates the set of values for UserEmailsTypeEnum
func GetUserEmailsTypeEnumValues() []UserEmailsTypeEnum {
	values := make([]UserEmailsTypeEnum, 0)
	for _, v := range mappingUserEmailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserEmailsTypeEnumStringValues Enumerates the set of values in String for UserEmailsTypeEnum
func GetUserEmailsTypeEnumStringValues() []string {
	return []string{
		"work",
		"home",
		"other",
		"recovery",
	}
}

// GetMappingUserEmailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserEmailsTypeEnum(val string) (UserEmailsTypeEnum, bool) {
	enum, ok := mappingUserEmailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
