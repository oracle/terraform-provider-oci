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

// AuthenticationFactorSettingsDuoSecuritySettings Settings related to Duo Security
// **Added In:** 19.2.1
// **SCIM++ Properties:**
//   - idcsSearchable: false
//   - multiValued: false
//   - mutability: readWrite
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type AuthenticationFactorSettingsDuoSecuritySettings struct {

	// Integration key from Duo Security authenticator
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IntegrationKey *string `mandatory:"true" json:"integrationKey"`

	// Secret key from Duo Security authenticator
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SecretKey *string `mandatory:"true" json:"secretKey"`

	// Hostname to access the Duo security account
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ApiHostname *string `mandatory:"true" json:"apiHostname"`

	// User attribute mapping value
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	UserMappingAttribute AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum `mandatory:"true" json:"userMappingAttribute"`

	// Attestation key to attest the request and response between Duo Security
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: never
	//  - type: string
	//  - uniqueness: none
	AttestationKey *string `mandatory:"false" json:"attestationKey"`
}

func (m AuthenticationFactorSettingsDuoSecuritySettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuthenticationFactorSettingsDuoSecuritySettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum(string(m.UserMappingAttribute)); !ok && m.UserMappingAttribute != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UserMappingAttribute: %s. Supported values are: %s.", m.UserMappingAttribute, strings.Join(GetAuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum Enum with underlying type: string
type AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum string

// Set of constants representing the allowable values for AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum
const (
	AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributePrimaryemail AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum = "primaryEmail"
	AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeUsername     AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum = "userName"
	AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeGivenname    AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum = "givenName"
)

var mappingAuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum = map[string]AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum{
	"primaryEmail": AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributePrimaryemail,
	"userName":     AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeUsername,
	"givenName":    AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeGivenname,
}

var mappingAuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnumLowerCase = map[string]AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum{
	"primaryemail": AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributePrimaryemail,
	"username":     AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeUsername,
	"givenname":    AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeGivenname,
}

// GetAuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnumValues Enumerates the set of values for AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum
func GetAuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnumValues() []AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum {
	values := make([]AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum, 0)
	for _, v := range mappingAuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnumStringValues Enumerates the set of values in String for AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum
func GetAuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnumStringValues() []string {
	return []string{
		"primaryEmail",
		"userName",
		"givenName",
	}
}

// GetMappingAuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum(val string) (AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum, bool) {
	enum, ok := mappingAuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
