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

// MeIms User's instant messaging addresses
type MeIms struct {

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
	Type MeImsTypeEnum `mandatory:"true" json:"type"`

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

func (m MeIms) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MeIms) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMeImsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMeImsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MeImsTypeEnum Enum with underlying type: string
type MeImsTypeEnum string

// Set of constants representing the allowable values for MeImsTypeEnum
const (
	MeImsTypeAim   MeImsTypeEnum = "aim"
	MeImsTypeGtalk MeImsTypeEnum = "gtalk"
	MeImsTypeIcq   MeImsTypeEnum = "icq"
	MeImsTypeXmpp  MeImsTypeEnum = "xmpp"
	MeImsTypeMsn   MeImsTypeEnum = "msn"
	MeImsTypeSkype MeImsTypeEnum = "skype"
	MeImsTypeQq    MeImsTypeEnum = "qq"
	MeImsTypeYahoo MeImsTypeEnum = "yahoo"
)

var mappingMeImsTypeEnum = map[string]MeImsTypeEnum{
	"aim":   MeImsTypeAim,
	"gtalk": MeImsTypeGtalk,
	"icq":   MeImsTypeIcq,
	"xmpp":  MeImsTypeXmpp,
	"msn":   MeImsTypeMsn,
	"skype": MeImsTypeSkype,
	"qq":    MeImsTypeQq,
	"yahoo": MeImsTypeYahoo,
}

var mappingMeImsTypeEnumLowerCase = map[string]MeImsTypeEnum{
	"aim":   MeImsTypeAim,
	"gtalk": MeImsTypeGtalk,
	"icq":   MeImsTypeIcq,
	"xmpp":  MeImsTypeXmpp,
	"msn":   MeImsTypeMsn,
	"skype": MeImsTypeSkype,
	"qq":    MeImsTypeQq,
	"yahoo": MeImsTypeYahoo,
}

// GetMeImsTypeEnumValues Enumerates the set of values for MeImsTypeEnum
func GetMeImsTypeEnumValues() []MeImsTypeEnum {
	values := make([]MeImsTypeEnum, 0)
	for _, v := range mappingMeImsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMeImsTypeEnumStringValues Enumerates the set of values in String for MeImsTypeEnum
func GetMeImsTypeEnumStringValues() []string {
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

// GetMappingMeImsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMeImsTypeEnum(val string) (MeImsTypeEnum, bool) {
	enum, ok := mappingMeImsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
