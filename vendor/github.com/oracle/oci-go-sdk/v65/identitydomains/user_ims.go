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

// UserIms User's instant messaging addresses
type UserIms struct {

	// User's instant messaging address
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

	// A label that indicates the attribute's function--for example, 'aim', 'gtalk', or 'mobile'
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type UserImsTypeEnum `mandatory:"true" json:"type"`

	// A human-readable name, primarily used for display purposes
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// A Boolean value that indicates the 'primary' or preferred attribute value for this attribute--for example, the preferred messenger or primary messenger. The primary attribute value 'true' MUST appear no more than once.
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
}

func (m UserIms) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserIms) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserImsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUserImsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserImsTypeEnum Enum with underlying type: string
type UserImsTypeEnum string

// Set of constants representing the allowable values for UserImsTypeEnum
const (
	UserImsTypeAim   UserImsTypeEnum = "aim"
	UserImsTypeGtalk UserImsTypeEnum = "gtalk"
	UserImsTypeIcq   UserImsTypeEnum = "icq"
	UserImsTypeXmpp  UserImsTypeEnum = "xmpp"
	UserImsTypeMsn   UserImsTypeEnum = "msn"
	UserImsTypeSkype UserImsTypeEnum = "skype"
	UserImsTypeQq    UserImsTypeEnum = "qq"
	UserImsTypeYahoo UserImsTypeEnum = "yahoo"
)

var mappingUserImsTypeEnum = map[string]UserImsTypeEnum{
	"aim":   UserImsTypeAim,
	"gtalk": UserImsTypeGtalk,
	"icq":   UserImsTypeIcq,
	"xmpp":  UserImsTypeXmpp,
	"msn":   UserImsTypeMsn,
	"skype": UserImsTypeSkype,
	"qq":    UserImsTypeQq,
	"yahoo": UserImsTypeYahoo,
}

var mappingUserImsTypeEnumLowerCase = map[string]UserImsTypeEnum{
	"aim":   UserImsTypeAim,
	"gtalk": UserImsTypeGtalk,
	"icq":   UserImsTypeIcq,
	"xmpp":  UserImsTypeXmpp,
	"msn":   UserImsTypeMsn,
	"skype": UserImsTypeSkype,
	"qq":    UserImsTypeQq,
	"yahoo": UserImsTypeYahoo,
}

// GetUserImsTypeEnumValues Enumerates the set of values for UserImsTypeEnum
func GetUserImsTypeEnumValues() []UserImsTypeEnum {
	values := make([]UserImsTypeEnum, 0)
	for _, v := range mappingUserImsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserImsTypeEnumStringValues Enumerates the set of values in String for UserImsTypeEnum
func GetUserImsTypeEnumStringValues() []string {
	return []string{
		"aim",
		"gtalk",
		"icq",
		"xmpp",
		"msn",
		"skype",
		"qq",
		"yahoo",
	}
}

// GetMappingUserImsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserImsTypeEnum(val string) (UserImsTypeEnum, bool) {
	enum, ok := mappingUserImsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
