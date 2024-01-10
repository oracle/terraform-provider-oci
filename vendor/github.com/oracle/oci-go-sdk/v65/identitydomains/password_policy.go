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

// PasswordPolicy PasswordPolicy resource.
type PasswordPolicy struct {

	// REQUIRED. The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. This specification defines URIs for User, Group, and a standard \"enterprise\" extension. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Schemas []string `mandatory:"true" json:"schemas"`

	// A String that is the name of the policy to display to the user. This is the only mandatory attribute for a password policy.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: server
	Name *string `mandatory:"true" json:"name"`

	// Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: global
	Id *string `mandatory:"false" json:"id"`

	// Unique OCI identifier for the SCIM Resource.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: global
	Ocid *string `mandatory:"false" json:"ocid"`

	Meta *Meta `mandatory:"false" json:"meta"`

	IdcsCreatedBy *IdcsCreatedBy `mandatory:"false" json:"idcsCreatedBy"`

	IdcsLastModifiedBy *IdcsLastModifiedBy `mandatory:"false" json:"idcsLastModifiedBy"`

	// Each value of this attribute specifies an operation that only an internal client may perform on this particular resource.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	IdcsPreventedOperations []IdcsPreventedOperationsEnum `mandatory:"false" json:"idcsPreventedOperations,omitempty"`

	// A list of tags on this resource.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [key, value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Tags []Tags `mandatory:"false" json:"tags"`

	// A boolean flag indicating this resource in the process of being deleted. Usually set to true when synchronous deletion of the resource would take too long.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	DeleteInProgress *bool `mandatory:"false" json:"deleteInProgress"`

	// The release number when the resource was upgraded.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	IdcsLastUpgradedInRelease *string `mandatory:"false" json:"idcsLastUpgradedInRelease"`

	// OCI Domain Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DomainOcid *string `mandatory:"false" json:"domainOcid"`

	// OCI Compartment Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CompartmentOcid *string `mandatory:"false" json:"compartmentOcid"`

	// OCI Tenant Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	TenancyOcid *string `mandatory:"false" json:"tenancyOcid"`

	// An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ExternalId *string `mandatory:"false" json:"externalId"`

	// A String that describes the password policy
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Description *string `mandatory:"false" json:"description"`

	// The maximum password length (in characters). A value of 0 or no value indicates no maximum length restriction.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxLength *int `mandatory:"false" json:"maxLength"`

	// The minimum password length (in characters). A value of 0 or no value indicates no minimum length restriction.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MinLength *int `mandatory:"false" json:"minLength"`

	// The minimum number of alphabetic characters in a password.  A value of 0 or no value indicates no minimum alphas restriction.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MinAlphas *int `mandatory:"false" json:"minAlphas"`

	// The minimum number of numeric characters in a password.  A value of 0 or no value indicates no minimum numeric character restriction.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MinNumerals *int `mandatory:"false" json:"minNumerals"`

	// The minimum number of a combination of alphabetic and numeric characters in a password.  A value of 0 or no value indicates no minimum alphanumeric character restriction.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MinAlphaNumerals *int `mandatory:"false" json:"minAlphaNumerals"`

	// The minimum number of special characters in a password. A value of 0 or no value indicates no minimum special characters restriction.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MinSpecialChars *int `mandatory:"false" json:"minSpecialChars"`

	// The maximum number of special characters in a password.  A value of 0 or no value indicates no maximum special characters restriction.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxSpecialChars *int `mandatory:"false" json:"maxSpecialChars"`

	// The minimum number of lowercase alphabetic characters in a password.  A value of 0 or no value indicates no minimum lowercase restriction.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MinLowerCase *int `mandatory:"false" json:"minLowerCase"`

	// The minimum number of uppercase alphabetic characters in a password. A value of 0 or no value indicates no minimum uppercase restriction.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MinUpperCase *int `mandatory:"false" json:"minUpperCase"`

	// The minimum number of unique characters in a password.  A value of 0 or no value indicates no minimum unique characters restriction.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MinUniqueChars *int `mandatory:"false" json:"minUniqueChars"`

	// The maximum number of repeated characters allowed in a password.  A value of 0 or no value indicates no such restriction.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxRepeatedChars *int `mandatory:"false" json:"maxRepeatedChars"`

	// Indicates that the password must begin with an alphabetic character
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	StartsWithAlphabet *bool `mandatory:"false" json:"startsWithAlphabet"`

	// Indicates a sequence of characters that match the user's first name of given name cannot be the password. Password validation against policy will be ignored if length of first name is less than or equal to 3 characters.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	FirstNameDisallowed *bool `mandatory:"false" json:"firstNameDisallowed"`

	// Indicates a sequence of characters that match the user's last name of given name cannot be the password. Password validation against policy will be ignored if length of last name is less than or equal to 3 characters.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	LastNameDisallowed *bool `mandatory:"false" json:"lastNameDisallowed"`

	// Indicates a sequence of characters that match the username cannot be the password. Password validation against policy will be ignored if length of user name is less than or equal to 3 characters.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	UserNameDisallowed *bool `mandatory:"false" json:"userNameDisallowed"`

	// List of User attributes whose values are not allowed in the password.
	// **Added In:** 2303212224
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DisallowedUserAttributeValues []string `mandatory:"false" json:"disallowedUserAttributeValues"`

	// Minimum time after which the user can resubmit the reset password request
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MinPasswordAge *int `mandatory:"false" json:"minPasswordAge"`

	// The number of days after which the password expires automatically
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	PasswordExpiresAfter *int `mandatory:"false" json:"passwordExpiresAfter"`

	// An integer indicating the number of days before which the user should be warned about password expiry.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	PasswordExpireWarning *int `mandatory:"false" json:"passwordExpireWarning"`

	// A String value whose contents indicate a set of characters that must appear, in any sequence, in a password value
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	RequiredChars *string `mandatory:"false" json:"requiredChars"`

	// A String value whose contents indicate a set of characters that cannot appear, in any sequence, in a password value
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DisallowedChars *string `mandatory:"false" json:"disallowedChars"`

	// A String value whose contents indicate a set of characters that can appear, in any sequence, in a password value
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AllowedChars *string `mandatory:"false" json:"allowedChars"`

	// A String value whose contents indicate a set of substrings that cannot appear, in any sequence, in a password value
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DisallowedSubstrings []string `mandatory:"false" json:"disallowedSubstrings"`

	// Indicates whether the password can match a dictionary word
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	DictionaryWordDisallowed *bool `mandatory:"false" json:"dictionaryWordDisallowed"`

	// A Reference value that contains the URI of a dictionary of words not allowed to appear within a password value
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DictionaryLocation *string `mandatory:"false" json:"dictionaryLocation"`

	// A delimiter used to separate characters in the dictionary file
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DictionaryDelimiter *string `mandatory:"false" json:"dictionaryDelimiter"`

	// An integer that represents the maximum number of failed logins before an account is locked
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxIncorrectAttempts *int `mandatory:"false" json:"maxIncorrectAttempts"`

	// The time period in minutes to lock out a user account when the threshold of invalid login attempts is reached. The available range is from 5 through 1440 minutes (24 hours).
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	LockoutDuration *int `mandatory:"false" json:"lockoutDuration"`

	// The number of passwords that will be kept in history that may not be used as a password
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	NumPasswordsInHistory *int `mandatory:"false" json:"numPasswordsInHistory"`

	// Indicates whether the password policy is configured as Simple, Standard, or Custom.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PasswordStrength PasswordPolicyPasswordStrengthEnum `mandatory:"false" json:"passwordStrength,omitempty"`

	// Indicates whether all of the users should be forced to reset their password on the next login (to comply with new password policy changes)
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: writeOnly
	//  - required: false
	//  - returned: never
	//  - type: boolean
	//  - uniqueness: none
	ForcePasswordReset *bool `mandatory:"false" json:"forcePasswordReset"`

	// The number of distinct characters between old password and new password
	// **Added In:** 2303212224
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	DistinctCharacters *int `mandatory:"false" json:"distinctCharacters"`

	// Password policy priority
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - idcsMinValue: 1
	//  - uniqueness: server
	Priority *int `mandatory:"false" json:"priority"`

	// A list of groups that the password policy belongs to.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	Groups []PasswordPolicyGroups `mandatory:"false" json:"groups"`

	// List of password policy rules that have values set. This map of stringKey:stringValue pairs can be used to aid users while setting/resetting password
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [key]
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	ConfiguredPasswordPolicyRules []PasswordPolicyConfiguredPasswordPolicyRules `mandatory:"false" json:"configuredPasswordPolicyRules"`
}

func (m PasswordPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PasswordPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingPasswordPolicyPasswordStrengthEnum(string(m.PasswordStrength)); !ok && m.PasswordStrength != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PasswordStrength: %s. Supported values are: %s.", m.PasswordStrength, strings.Join(GetPasswordPolicyPasswordStrengthEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PasswordPolicyPasswordStrengthEnum Enum with underlying type: string
type PasswordPolicyPasswordStrengthEnum string

// Set of constants representing the allowable values for PasswordPolicyPasswordStrengthEnum
const (
	PasswordPolicyPasswordStrengthSimple   PasswordPolicyPasswordStrengthEnum = "Simple"
	PasswordPolicyPasswordStrengthStandard PasswordPolicyPasswordStrengthEnum = "Standard"
	PasswordPolicyPasswordStrengthCustom   PasswordPolicyPasswordStrengthEnum = "Custom"
)

var mappingPasswordPolicyPasswordStrengthEnum = map[string]PasswordPolicyPasswordStrengthEnum{
	"Simple":   PasswordPolicyPasswordStrengthSimple,
	"Standard": PasswordPolicyPasswordStrengthStandard,
	"Custom":   PasswordPolicyPasswordStrengthCustom,
}

var mappingPasswordPolicyPasswordStrengthEnumLowerCase = map[string]PasswordPolicyPasswordStrengthEnum{
	"simple":   PasswordPolicyPasswordStrengthSimple,
	"standard": PasswordPolicyPasswordStrengthStandard,
	"custom":   PasswordPolicyPasswordStrengthCustom,
}

// GetPasswordPolicyPasswordStrengthEnumValues Enumerates the set of values for PasswordPolicyPasswordStrengthEnum
func GetPasswordPolicyPasswordStrengthEnumValues() []PasswordPolicyPasswordStrengthEnum {
	values := make([]PasswordPolicyPasswordStrengthEnum, 0)
	for _, v := range mappingPasswordPolicyPasswordStrengthEnum {
		values = append(values, v)
	}
	return values
}

// GetPasswordPolicyPasswordStrengthEnumStringValues Enumerates the set of values in String for PasswordPolicyPasswordStrengthEnum
func GetPasswordPolicyPasswordStrengthEnumStringValues() []string {
	return []string{
		"Simple",
		"Standard",
		"Custom",
	}
}

// GetMappingPasswordPolicyPasswordStrengthEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPasswordPolicyPasswordStrengthEnum(val string) (PasswordPolicyPasswordStrengthEnum, bool) {
	enum, ok := mappingPasswordPolicyPasswordStrengthEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
