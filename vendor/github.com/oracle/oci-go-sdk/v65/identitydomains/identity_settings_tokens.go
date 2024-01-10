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

// IdentitySettingsTokens A list of tokens and their expiry length.
type IdentitySettingsTokens struct {

	// The token type.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type IdentitySettingsTokensTypeEnum `mandatory:"true" json:"type"`

	// Indicates the number of minutes after which the token expires automatically.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	ExpiresAfter *int `mandatory:"false" json:"expiresAfter"`
}

func (m IdentitySettingsTokens) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentitySettingsTokens) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIdentitySettingsTokensTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetIdentitySettingsTokensTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IdentitySettingsTokensTypeEnum Enum with underlying type: string
type IdentitySettingsTokensTypeEnum string

// Set of constants representing the allowable values for IdentitySettingsTokensTypeEnum
const (
	IdentitySettingsTokensTypeEmailverification IdentitySettingsTokensTypeEnum = "emailVerification"
	IdentitySettingsTokensTypePasswordreset     IdentitySettingsTokensTypeEnum = "passwordReset"
	IdentitySettingsTokensTypeCreateuser        IdentitySettingsTokensTypeEnum = "createUser"
)

var mappingIdentitySettingsTokensTypeEnum = map[string]IdentitySettingsTokensTypeEnum{
	"emailVerification": IdentitySettingsTokensTypeEmailverification,
	"passwordReset":     IdentitySettingsTokensTypePasswordreset,
	"createUser":        IdentitySettingsTokensTypeCreateuser,
}

var mappingIdentitySettingsTokensTypeEnumLowerCase = map[string]IdentitySettingsTokensTypeEnum{
	"emailverification": IdentitySettingsTokensTypeEmailverification,
	"passwordreset":     IdentitySettingsTokensTypePasswordreset,
	"createuser":        IdentitySettingsTokensTypeCreateuser,
}

// GetIdentitySettingsTokensTypeEnumValues Enumerates the set of values for IdentitySettingsTokensTypeEnum
func GetIdentitySettingsTokensTypeEnumValues() []IdentitySettingsTokensTypeEnum {
	values := make([]IdentitySettingsTokensTypeEnum, 0)
	for _, v := range mappingIdentitySettingsTokensTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentitySettingsTokensTypeEnumStringValues Enumerates the set of values in String for IdentitySettingsTokensTypeEnum
func GetIdentitySettingsTokensTypeEnumStringValues() []string {
	return []string{
		"emailVerification",
		"passwordReset",
		"createUser",
	}
}

// GetMappingIdentitySettingsTokensTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentitySettingsTokensTypeEnum(val string) (IdentitySettingsTokensTypeEnum, bool) {
	enum, ok := mappingIdentitySettingsTokensTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
