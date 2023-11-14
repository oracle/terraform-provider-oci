// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// SettingsTenantCustomClaims Custom claims associated with the specific tenant
type SettingsTenantCustomClaims struct {

	// Custom claim name
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: server
	Name *string `mandatory:"true" json:"name"`

	// Custom claim value
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// Indicates under what scenario the custom claim will be return
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Mode SettingsTenantCustomClaimsModeEnum `mandatory:"true" json:"mode"`

	// Indicates if the custom claim is an expression
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Expression *bool `mandatory:"true" json:"expression"`

	// Indicates if the custom claim is associated with all scopes
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AllScopes *bool `mandatory:"true" json:"allScopes"`

	// Indicates what type of token the custom claim will be embedded
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	TokenType SettingsTenantCustomClaimsTokenTypeEnum `mandatory:"true" json:"tokenType"`

	// Scopes associated with a specific custom claim
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Scopes []string `mandatory:"false" json:"scopes"`
}

func (m SettingsTenantCustomClaims) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SettingsTenantCustomClaims) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSettingsTenantCustomClaimsModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetSettingsTenantCustomClaimsModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSettingsTenantCustomClaimsTokenTypeEnum(string(m.TokenType)); !ok && m.TokenType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TokenType: %s. Supported values are: %s.", m.TokenType, strings.Join(GetSettingsTenantCustomClaimsTokenTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SettingsTenantCustomClaimsModeEnum Enum with underlying type: string
type SettingsTenantCustomClaimsModeEnum string

// Set of constants representing the allowable values for SettingsTenantCustomClaimsModeEnum
const (
	SettingsTenantCustomClaimsModeAlways  SettingsTenantCustomClaimsModeEnum = "always"
	SettingsTenantCustomClaimsModeRequest SettingsTenantCustomClaimsModeEnum = "request"
	SettingsTenantCustomClaimsModeNever   SettingsTenantCustomClaimsModeEnum = "never"
)

var mappingSettingsTenantCustomClaimsModeEnum = map[string]SettingsTenantCustomClaimsModeEnum{
	"always":  SettingsTenantCustomClaimsModeAlways,
	"request": SettingsTenantCustomClaimsModeRequest,
	"never":   SettingsTenantCustomClaimsModeNever,
}

var mappingSettingsTenantCustomClaimsModeEnumLowerCase = map[string]SettingsTenantCustomClaimsModeEnum{
	"always":  SettingsTenantCustomClaimsModeAlways,
	"request": SettingsTenantCustomClaimsModeRequest,
	"never":   SettingsTenantCustomClaimsModeNever,
}

// GetSettingsTenantCustomClaimsModeEnumValues Enumerates the set of values for SettingsTenantCustomClaimsModeEnum
func GetSettingsTenantCustomClaimsModeEnumValues() []SettingsTenantCustomClaimsModeEnum {
	values := make([]SettingsTenantCustomClaimsModeEnum, 0)
	for _, v := range mappingSettingsTenantCustomClaimsModeEnum {
		values = append(values, v)
	}
	return values
}

// GetSettingsTenantCustomClaimsModeEnumStringValues Enumerates the set of values in String for SettingsTenantCustomClaimsModeEnum
func GetSettingsTenantCustomClaimsModeEnumStringValues() []string {
	return []string{
		"always",
		"request",
		"never",
	}
}

// GetMappingSettingsTenantCustomClaimsModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSettingsTenantCustomClaimsModeEnum(val string) (SettingsTenantCustomClaimsModeEnum, bool) {
	enum, ok := mappingSettingsTenantCustomClaimsModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SettingsTenantCustomClaimsTokenTypeEnum Enum with underlying type: string
type SettingsTenantCustomClaimsTokenTypeEnum string

// Set of constants representing the allowable values for SettingsTenantCustomClaimsTokenTypeEnum
const (
	SettingsTenantCustomClaimsTokenTypeAt   SettingsTenantCustomClaimsTokenTypeEnum = "AT"
	SettingsTenantCustomClaimsTokenTypeIt   SettingsTenantCustomClaimsTokenTypeEnum = "IT"
	SettingsTenantCustomClaimsTokenTypeBoth SettingsTenantCustomClaimsTokenTypeEnum = "BOTH"
)

var mappingSettingsTenantCustomClaimsTokenTypeEnum = map[string]SettingsTenantCustomClaimsTokenTypeEnum{
	"AT":   SettingsTenantCustomClaimsTokenTypeAt,
	"IT":   SettingsTenantCustomClaimsTokenTypeIt,
	"BOTH": SettingsTenantCustomClaimsTokenTypeBoth,
}

var mappingSettingsTenantCustomClaimsTokenTypeEnumLowerCase = map[string]SettingsTenantCustomClaimsTokenTypeEnum{
	"at":   SettingsTenantCustomClaimsTokenTypeAt,
	"it":   SettingsTenantCustomClaimsTokenTypeIt,
	"both": SettingsTenantCustomClaimsTokenTypeBoth,
}

// GetSettingsTenantCustomClaimsTokenTypeEnumValues Enumerates the set of values for SettingsTenantCustomClaimsTokenTypeEnum
func GetSettingsTenantCustomClaimsTokenTypeEnumValues() []SettingsTenantCustomClaimsTokenTypeEnum {
	values := make([]SettingsTenantCustomClaimsTokenTypeEnum, 0)
	for _, v := range mappingSettingsTenantCustomClaimsTokenTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSettingsTenantCustomClaimsTokenTypeEnumStringValues Enumerates the set of values in String for SettingsTenantCustomClaimsTokenTypeEnum
func GetSettingsTenantCustomClaimsTokenTypeEnumStringValues() []string {
	return []string{
		"AT",
		"IT",
		"BOTH",
	}
}

// GetMappingSettingsTenantCustomClaimsTokenTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSettingsTenantCustomClaimsTokenTypeEnum(val string) (SettingsTenantCustomClaimsTokenTypeEnum, bool) {
	enum, ok := mappingSettingsTenantCustomClaimsTokenTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
