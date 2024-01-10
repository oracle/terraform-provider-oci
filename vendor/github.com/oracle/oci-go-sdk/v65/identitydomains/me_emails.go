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

// MeEmails A complex attribute representing emails
type MeEmails struct {

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
	Type MeEmailsTypeEnum `mandatory:"true" json:"type"`

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

func (m MeEmails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MeEmails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMeEmailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMeEmailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MeEmailsTypeEnum Enum with underlying type: string
type MeEmailsTypeEnum string

// Set of constants representing the allowable values for MeEmailsTypeEnum
const (
	MeEmailsTypeWork     MeEmailsTypeEnum = "work"
	MeEmailsTypeHome     MeEmailsTypeEnum = "home"
	MeEmailsTypeOther    MeEmailsTypeEnum = "other"
	MeEmailsTypeRecovery MeEmailsTypeEnum = "recovery"
)

var mappingMeEmailsTypeEnum = map[string]MeEmailsTypeEnum{
	"work":     MeEmailsTypeWork,
	"home":     MeEmailsTypeHome,
	"other":    MeEmailsTypeOther,
	"recovery": MeEmailsTypeRecovery,
}

var mappingMeEmailsTypeEnumLowerCase = map[string]MeEmailsTypeEnum{
	"work":     MeEmailsTypeWork,
	"home":     MeEmailsTypeHome,
	"other":    MeEmailsTypeOther,
	"recovery": MeEmailsTypeRecovery,
}

// GetMeEmailsTypeEnumValues Enumerates the set of values for MeEmailsTypeEnum
func GetMeEmailsTypeEnumValues() []MeEmailsTypeEnum {
	values := make([]MeEmailsTypeEnum, 0)
	for _, v := range mappingMeEmailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMeEmailsTypeEnumStringValues Enumerates the set of values in String for MeEmailsTypeEnum
func GetMeEmailsTypeEnumStringValues() []string {
	return []string{
		"work",
		"home",
		"other",
		"recovery",
	}
}

// GetMappingMeEmailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMeEmailsTypeEnum(val string) (MeEmailsTypeEnum, bool) {
	enum, ok := mappingMeEmailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
