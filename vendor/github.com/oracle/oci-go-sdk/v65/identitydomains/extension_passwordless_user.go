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

// ExtensionPasswordlessUser This extension defines attributes used to manage Passwordless-Factor Authentication within a service provider. The extension is typically applied to a User resource, but MAY be applied to other resources that use MFA.
type ExtensionPasswordlessUser struct {

	// Authentication Factor Type
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	FactorType ExtensionPasswordlessUserFactorTypeEnum `mandatory:"false" json:"factorType,omitempty"`

	// Authentication Factor Method
	// **Added In:** 2009232244
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	FactorMethod *string `mandatory:"false" json:"factorMethod"`

	FactorIdentifier *UserExtFactorIdentifier `mandatory:"false" json:"factorIdentifier"`
}

func (m ExtensionPasswordlessUser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionPasswordlessUser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExtensionPasswordlessUserFactorTypeEnum(string(m.FactorType)); !ok && m.FactorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FactorType: %s. Supported values are: %s.", m.FactorType, strings.Join(GetExtensionPasswordlessUserFactorTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExtensionPasswordlessUserFactorTypeEnum Enum with underlying type: string
type ExtensionPasswordlessUserFactorTypeEnum string

// Set of constants representing the allowable values for ExtensionPasswordlessUserFactorTypeEnum
const (
	ExtensionPasswordlessUserFactorTypeEmail             ExtensionPasswordlessUserFactorTypeEnum = "EMAIL"
	ExtensionPasswordlessUserFactorTypeSms               ExtensionPasswordlessUserFactorTypeEnum = "SMS"
	ExtensionPasswordlessUserFactorTypePhoneCall         ExtensionPasswordlessUserFactorTypeEnum = "PHONE_CALL"
	ExtensionPasswordlessUserFactorTypeTotp              ExtensionPasswordlessUserFactorTypeEnum = "TOTP"
	ExtensionPasswordlessUserFactorTypePush              ExtensionPasswordlessUserFactorTypeEnum = "PUSH"
	ExtensionPasswordlessUserFactorTypeOfflinetotp       ExtensionPasswordlessUserFactorTypeEnum = "OFFLINETOTP"
	ExtensionPasswordlessUserFactorTypeSecurityQuestions ExtensionPasswordlessUserFactorTypeEnum = "SECURITY_QUESTIONS"
	ExtensionPasswordlessUserFactorTypeVoice             ExtensionPasswordlessUserFactorTypeEnum = "VOICE"
	ExtensionPasswordlessUserFactorTypeUsernamePassword  ExtensionPasswordlessUserFactorTypeEnum = "USERNAME_PASSWORD"
	ExtensionPasswordlessUserFactorTypeThirdparty        ExtensionPasswordlessUserFactorTypeEnum = "THIRDPARTY"
	ExtensionPasswordlessUserFactorTypeFidoAuthenticator ExtensionPasswordlessUserFactorTypeEnum = "FIDO_AUTHENTICATOR"
)

var mappingExtensionPasswordlessUserFactorTypeEnum = map[string]ExtensionPasswordlessUserFactorTypeEnum{
	"EMAIL":              ExtensionPasswordlessUserFactorTypeEmail,
	"SMS":                ExtensionPasswordlessUserFactorTypeSms,
	"PHONE_CALL":         ExtensionPasswordlessUserFactorTypePhoneCall,
	"TOTP":               ExtensionPasswordlessUserFactorTypeTotp,
	"PUSH":               ExtensionPasswordlessUserFactorTypePush,
	"OFFLINETOTP":        ExtensionPasswordlessUserFactorTypeOfflinetotp,
	"SECURITY_QUESTIONS": ExtensionPasswordlessUserFactorTypeSecurityQuestions,
	"VOICE":              ExtensionPasswordlessUserFactorTypeVoice,
	"USERNAME_PASSWORD":  ExtensionPasswordlessUserFactorTypeUsernamePassword,
	"THIRDPARTY":         ExtensionPasswordlessUserFactorTypeThirdparty,
	"FIDO_AUTHENTICATOR": ExtensionPasswordlessUserFactorTypeFidoAuthenticator,
}

var mappingExtensionPasswordlessUserFactorTypeEnumLowerCase = map[string]ExtensionPasswordlessUserFactorTypeEnum{
	"email":              ExtensionPasswordlessUserFactorTypeEmail,
	"sms":                ExtensionPasswordlessUserFactorTypeSms,
	"phone_call":         ExtensionPasswordlessUserFactorTypePhoneCall,
	"totp":               ExtensionPasswordlessUserFactorTypeTotp,
	"push":               ExtensionPasswordlessUserFactorTypePush,
	"offlinetotp":        ExtensionPasswordlessUserFactorTypeOfflinetotp,
	"security_questions": ExtensionPasswordlessUserFactorTypeSecurityQuestions,
	"voice":              ExtensionPasswordlessUserFactorTypeVoice,
	"username_password":  ExtensionPasswordlessUserFactorTypeUsernamePassword,
	"thirdparty":         ExtensionPasswordlessUserFactorTypeThirdparty,
	"fido_authenticator": ExtensionPasswordlessUserFactorTypeFidoAuthenticator,
}

// GetExtensionPasswordlessUserFactorTypeEnumValues Enumerates the set of values for ExtensionPasswordlessUserFactorTypeEnum
func GetExtensionPasswordlessUserFactorTypeEnumValues() []ExtensionPasswordlessUserFactorTypeEnum {
	values := make([]ExtensionPasswordlessUserFactorTypeEnum, 0)
	for _, v := range mappingExtensionPasswordlessUserFactorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionPasswordlessUserFactorTypeEnumStringValues Enumerates the set of values in String for ExtensionPasswordlessUserFactorTypeEnum
func GetExtensionPasswordlessUserFactorTypeEnumStringValues() []string {
	return []string{
		"EMAIL",
		"SMS",
		"PHONE_CALL",
		"TOTP",
		"PUSH",
		"OFFLINETOTP",
		"SECURITY_QUESTIONS",
		"VOICE",
		"USERNAME_PASSWORD",
		"THIRDPARTY",
		"FIDO_AUTHENTICATOR",
	}
}

// GetMappingExtensionPasswordlessUserFactorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionPasswordlessUserFactorTypeEnum(val string) (ExtensionPasswordlessUserFactorTypeEnum, bool) {
	enum, ok := mappingExtensionPasswordlessUserFactorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
