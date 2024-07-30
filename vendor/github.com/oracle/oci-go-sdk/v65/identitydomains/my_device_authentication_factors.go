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

// MyDeviceAuthenticationFactors Authentication Factors
type MyDeviceAuthenticationFactors struct {

	// Authentication Factor Type
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	Type MyDeviceAuthenticationFactorsTypeEnum `mandatory:"true" json:"type"`

	// Authentication Factor Status
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	Status MyDeviceAuthenticationFactorsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Authentication Factor public key issued by client
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	PublicKey *string `mandatory:"false" json:"publicKey"`
}

func (m MyDeviceAuthenticationFactors) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MyDeviceAuthenticationFactors) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMyDeviceAuthenticationFactorsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMyDeviceAuthenticationFactorsTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMyDeviceAuthenticationFactorsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetMyDeviceAuthenticationFactorsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MyDeviceAuthenticationFactorsTypeEnum Enum with underlying type: string
type MyDeviceAuthenticationFactorsTypeEnum string

// Set of constants representing the allowable values for MyDeviceAuthenticationFactorsTypeEnum
const (
	MyDeviceAuthenticationFactorsTypeEmail             MyDeviceAuthenticationFactorsTypeEnum = "EMAIL"
	MyDeviceAuthenticationFactorsTypeSms               MyDeviceAuthenticationFactorsTypeEnum = "SMS"
	MyDeviceAuthenticationFactorsTypeTotp              MyDeviceAuthenticationFactorsTypeEnum = "TOTP"
	MyDeviceAuthenticationFactorsTypePush              MyDeviceAuthenticationFactorsTypeEnum = "PUSH"
	MyDeviceAuthenticationFactorsTypeOfflinetotp       MyDeviceAuthenticationFactorsTypeEnum = "OFFLINETOTP"
	MyDeviceAuthenticationFactorsTypeVoice             MyDeviceAuthenticationFactorsTypeEnum = "VOICE"
	MyDeviceAuthenticationFactorsTypePhoneCall         MyDeviceAuthenticationFactorsTypeEnum = "PHONE_CALL"
	MyDeviceAuthenticationFactorsTypeThirdparty        MyDeviceAuthenticationFactorsTypeEnum = "THIRDPARTY"
	MyDeviceAuthenticationFactorsTypeFidoAuthenticator MyDeviceAuthenticationFactorsTypeEnum = "FIDO_AUTHENTICATOR"
	MyDeviceAuthenticationFactorsTypeYubicoOtp         MyDeviceAuthenticationFactorsTypeEnum = "YUBICO_OTP"
)

var mappingMyDeviceAuthenticationFactorsTypeEnum = map[string]MyDeviceAuthenticationFactorsTypeEnum{
	"EMAIL":              MyDeviceAuthenticationFactorsTypeEmail,
	"SMS":                MyDeviceAuthenticationFactorsTypeSms,
	"TOTP":               MyDeviceAuthenticationFactorsTypeTotp,
	"PUSH":               MyDeviceAuthenticationFactorsTypePush,
	"OFFLINETOTP":        MyDeviceAuthenticationFactorsTypeOfflinetotp,
	"VOICE":              MyDeviceAuthenticationFactorsTypeVoice,
	"PHONE_CALL":         MyDeviceAuthenticationFactorsTypePhoneCall,
	"THIRDPARTY":         MyDeviceAuthenticationFactorsTypeThirdparty,
	"FIDO_AUTHENTICATOR": MyDeviceAuthenticationFactorsTypeFidoAuthenticator,
	"YUBICO_OTP":         MyDeviceAuthenticationFactorsTypeYubicoOtp,
}

var mappingMyDeviceAuthenticationFactorsTypeEnumLowerCase = map[string]MyDeviceAuthenticationFactorsTypeEnum{
	"email":              MyDeviceAuthenticationFactorsTypeEmail,
	"sms":                MyDeviceAuthenticationFactorsTypeSms,
	"totp":               MyDeviceAuthenticationFactorsTypeTotp,
	"push":               MyDeviceAuthenticationFactorsTypePush,
	"offlinetotp":        MyDeviceAuthenticationFactorsTypeOfflinetotp,
	"voice":              MyDeviceAuthenticationFactorsTypeVoice,
	"phone_call":         MyDeviceAuthenticationFactorsTypePhoneCall,
	"thirdparty":         MyDeviceAuthenticationFactorsTypeThirdparty,
	"fido_authenticator": MyDeviceAuthenticationFactorsTypeFidoAuthenticator,
	"yubico_otp":         MyDeviceAuthenticationFactorsTypeYubicoOtp,
}

// GetMyDeviceAuthenticationFactorsTypeEnumValues Enumerates the set of values for MyDeviceAuthenticationFactorsTypeEnum
func GetMyDeviceAuthenticationFactorsTypeEnumValues() []MyDeviceAuthenticationFactorsTypeEnum {
	values := make([]MyDeviceAuthenticationFactorsTypeEnum, 0)
	for _, v := range mappingMyDeviceAuthenticationFactorsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMyDeviceAuthenticationFactorsTypeEnumStringValues Enumerates the set of values in String for MyDeviceAuthenticationFactorsTypeEnum
func GetMyDeviceAuthenticationFactorsTypeEnumStringValues() []string {
	return []string{
		"EMAIL",
		"SMS",
		"TOTP",
		"PUSH",
		"OFFLINETOTP",
		"VOICE",
		"PHONE_CALL",
		"THIRDPARTY",
		"FIDO_AUTHENTICATOR",
		"YUBICO_OTP",
	}
}

// GetMappingMyDeviceAuthenticationFactorsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyDeviceAuthenticationFactorsTypeEnum(val string) (MyDeviceAuthenticationFactorsTypeEnum, bool) {
	enum, ok := mappingMyDeviceAuthenticationFactorsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MyDeviceAuthenticationFactorsStatusEnum Enum with underlying type: string
type MyDeviceAuthenticationFactorsStatusEnum string

// Set of constants representing the allowable values for MyDeviceAuthenticationFactorsStatusEnum
const (
	MyDeviceAuthenticationFactorsStatusInitiated  MyDeviceAuthenticationFactorsStatusEnum = "INITIATED"
	MyDeviceAuthenticationFactorsStatusInprogress MyDeviceAuthenticationFactorsStatusEnum = "INPROGRESS"
	MyDeviceAuthenticationFactorsStatusEnrolled   MyDeviceAuthenticationFactorsStatusEnum = "ENROLLED"
	MyDeviceAuthenticationFactorsStatusLocked     MyDeviceAuthenticationFactorsStatusEnum = "LOCKED"
	MyDeviceAuthenticationFactorsStatusInactive   MyDeviceAuthenticationFactorsStatusEnum = "INACTIVE"
	MyDeviceAuthenticationFactorsStatusBlocked    MyDeviceAuthenticationFactorsStatusEnum = "BLOCKED"
)

var mappingMyDeviceAuthenticationFactorsStatusEnum = map[string]MyDeviceAuthenticationFactorsStatusEnum{
	"INITIATED":  MyDeviceAuthenticationFactorsStatusInitiated,
	"INPROGRESS": MyDeviceAuthenticationFactorsStatusInprogress,
	"ENROLLED":   MyDeviceAuthenticationFactorsStatusEnrolled,
	"LOCKED":     MyDeviceAuthenticationFactorsStatusLocked,
	"INACTIVE":   MyDeviceAuthenticationFactorsStatusInactive,
	"BLOCKED":    MyDeviceAuthenticationFactorsStatusBlocked,
}

var mappingMyDeviceAuthenticationFactorsStatusEnumLowerCase = map[string]MyDeviceAuthenticationFactorsStatusEnum{
	"initiated":  MyDeviceAuthenticationFactorsStatusInitiated,
	"inprogress": MyDeviceAuthenticationFactorsStatusInprogress,
	"enrolled":   MyDeviceAuthenticationFactorsStatusEnrolled,
	"locked":     MyDeviceAuthenticationFactorsStatusLocked,
	"inactive":   MyDeviceAuthenticationFactorsStatusInactive,
	"blocked":    MyDeviceAuthenticationFactorsStatusBlocked,
}

// GetMyDeviceAuthenticationFactorsStatusEnumValues Enumerates the set of values for MyDeviceAuthenticationFactorsStatusEnum
func GetMyDeviceAuthenticationFactorsStatusEnumValues() []MyDeviceAuthenticationFactorsStatusEnum {
	values := make([]MyDeviceAuthenticationFactorsStatusEnum, 0)
	for _, v := range mappingMyDeviceAuthenticationFactorsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetMyDeviceAuthenticationFactorsStatusEnumStringValues Enumerates the set of values in String for MyDeviceAuthenticationFactorsStatusEnum
func GetMyDeviceAuthenticationFactorsStatusEnumStringValues() []string {
	return []string{
		"INITIATED",
		"INPROGRESS",
		"ENROLLED",
		"LOCKED",
		"INACTIVE",
		"BLOCKED",
	}
}

// GetMappingMyDeviceAuthenticationFactorsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyDeviceAuthenticationFactorsStatusEnum(val string) (MyDeviceAuthenticationFactorsStatusEnum, bool) {
	enum, ok := mappingMyDeviceAuthenticationFactorsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
