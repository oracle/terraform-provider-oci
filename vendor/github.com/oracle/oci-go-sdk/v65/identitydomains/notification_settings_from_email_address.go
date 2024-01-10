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

// NotificationSettingsFromEmailAddress From email address to be used in the notification emails
// **SCIM++ Properties:**
//   - caseExact: false
//   - multiValued: false
//   - mutability: readWrite
//   - required: true
//   - returned: always
//   - type: complex
//   - uniqueness: none
type NotificationSettingsFromEmailAddress struct {

	// Value of the From email address
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	Value *string `mandatory:"true" json:"value"`

	// From address verification mode. If postmaster account is available then 'domain' mode is used or entire valid email can be verified using 'email' mode
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Validate NotificationSettingsFromEmailAddressValidateEnum `mandatory:"true" json:"validate"`

	// Validation status for the From email address
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	ValidationStatus NotificationSettingsFromEmailAddressValidationStatusEnum `mandatory:"false" json:"validationStatus,omitempty"`

	// Display name for the From email address
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m NotificationSettingsFromEmailAddress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NotificationSettingsFromEmailAddress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNotificationSettingsFromEmailAddressValidateEnum(string(m.Validate)); !ok && m.Validate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Validate: %s. Supported values are: %s.", m.Validate, strings.Join(GetNotificationSettingsFromEmailAddressValidateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingNotificationSettingsFromEmailAddressValidationStatusEnum(string(m.ValidationStatus)); !ok && m.ValidationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValidationStatus: %s. Supported values are: %s.", m.ValidationStatus, strings.Join(GetNotificationSettingsFromEmailAddressValidationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NotificationSettingsFromEmailAddressValidationStatusEnum Enum with underlying type: string
type NotificationSettingsFromEmailAddressValidationStatusEnum string

// Set of constants representing the allowable values for NotificationSettingsFromEmailAddressValidationStatusEnum
const (
	NotificationSettingsFromEmailAddressValidationStatusVerified NotificationSettingsFromEmailAddressValidationStatusEnum = "VERIFIED"
	NotificationSettingsFromEmailAddressValidationStatusPending  NotificationSettingsFromEmailAddressValidationStatusEnum = "PENDING"
)

var mappingNotificationSettingsFromEmailAddressValidationStatusEnum = map[string]NotificationSettingsFromEmailAddressValidationStatusEnum{
	"VERIFIED": NotificationSettingsFromEmailAddressValidationStatusVerified,
	"PENDING":  NotificationSettingsFromEmailAddressValidationStatusPending,
}

var mappingNotificationSettingsFromEmailAddressValidationStatusEnumLowerCase = map[string]NotificationSettingsFromEmailAddressValidationStatusEnum{
	"verified": NotificationSettingsFromEmailAddressValidationStatusVerified,
	"pending":  NotificationSettingsFromEmailAddressValidationStatusPending,
}

// GetNotificationSettingsFromEmailAddressValidationStatusEnumValues Enumerates the set of values for NotificationSettingsFromEmailAddressValidationStatusEnum
func GetNotificationSettingsFromEmailAddressValidationStatusEnumValues() []NotificationSettingsFromEmailAddressValidationStatusEnum {
	values := make([]NotificationSettingsFromEmailAddressValidationStatusEnum, 0)
	for _, v := range mappingNotificationSettingsFromEmailAddressValidationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetNotificationSettingsFromEmailAddressValidationStatusEnumStringValues Enumerates the set of values in String for NotificationSettingsFromEmailAddressValidationStatusEnum
func GetNotificationSettingsFromEmailAddressValidationStatusEnumStringValues() []string {
	return []string{
		"VERIFIED",
		"PENDING",
	}
}

// GetMappingNotificationSettingsFromEmailAddressValidationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNotificationSettingsFromEmailAddressValidationStatusEnum(val string) (NotificationSettingsFromEmailAddressValidationStatusEnum, bool) {
	enum, ok := mappingNotificationSettingsFromEmailAddressValidationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// NotificationSettingsFromEmailAddressValidateEnum Enum with underlying type: string
type NotificationSettingsFromEmailAddressValidateEnum string

// Set of constants representing the allowable values for NotificationSettingsFromEmailAddressValidateEnum
const (
	NotificationSettingsFromEmailAddressValidateEmail  NotificationSettingsFromEmailAddressValidateEnum = "email"
	NotificationSettingsFromEmailAddressValidateDomain NotificationSettingsFromEmailAddressValidateEnum = "domain"
)

var mappingNotificationSettingsFromEmailAddressValidateEnum = map[string]NotificationSettingsFromEmailAddressValidateEnum{
	"email":  NotificationSettingsFromEmailAddressValidateEmail,
	"domain": NotificationSettingsFromEmailAddressValidateDomain,
}

var mappingNotificationSettingsFromEmailAddressValidateEnumLowerCase = map[string]NotificationSettingsFromEmailAddressValidateEnum{
	"email":  NotificationSettingsFromEmailAddressValidateEmail,
	"domain": NotificationSettingsFromEmailAddressValidateDomain,
}

// GetNotificationSettingsFromEmailAddressValidateEnumValues Enumerates the set of values for NotificationSettingsFromEmailAddressValidateEnum
func GetNotificationSettingsFromEmailAddressValidateEnumValues() []NotificationSettingsFromEmailAddressValidateEnum {
	values := make([]NotificationSettingsFromEmailAddressValidateEnum, 0)
	for _, v := range mappingNotificationSettingsFromEmailAddressValidateEnum {
		values = append(values, v)
	}
	return values
}

// GetNotificationSettingsFromEmailAddressValidateEnumStringValues Enumerates the set of values in String for NotificationSettingsFromEmailAddressValidateEnum
func GetNotificationSettingsFromEmailAddressValidateEnumStringValues() []string {
	return []string{
		"email",
		"domain",
	}
}

// GetMappingNotificationSettingsFromEmailAddressValidateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNotificationSettingsFromEmailAddressValidateEnum(val string) (NotificationSettingsFromEmailAddressValidateEnum, bool) {
	enum, ok := mappingNotificationSettingsFromEmailAddressValidateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
