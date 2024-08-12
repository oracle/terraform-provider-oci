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

// OciConsoleSignOnPolicyConsentModifiedResource The modified Policy, Rule, ConditionGroup or Condition during consent signing.
// **SCIM++ Properties:**
//   - idcsSearchable: false
//   - multiValued: false
//   - mutability: immutable
//   - required: true
//   - returned: default
//   - type: complex
//   - uniqueness: none
type OciConsoleSignOnPolicyConsentModifiedResource struct {

	// Modified Policy, Rule, ConditionGroup or Condition Id.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	Value *string `mandatory:"true" json:"value"`

	// The modified Policy, Rule, ConditionGroup, or Condition OCID.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	Ocid *string `mandatory:"true" json:"ocid"`

	// The Modified Resource type - Policy, Rule, ConditionGroup, or Condition. A label that indicates the resource type.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - idcsDefaultValue: Policy
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum `mandatory:"true" json:"type"`
}

func (m OciConsoleSignOnPolicyConsentModifiedResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciConsoleSignOnPolicyConsentModifiedResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciConsoleSignOnPolicyConsentModifiedResourceTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetOciConsoleSignOnPolicyConsentModifiedResourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum Enum with underlying type: string
type OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum string

// Set of constants representing the allowable values for OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum
const (
	OciConsoleSignOnPolicyConsentModifiedResourceTypePolicy         OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum = "Policy"
	OciConsoleSignOnPolicyConsentModifiedResourceTypeRule           OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum = "Rule"
	OciConsoleSignOnPolicyConsentModifiedResourceTypeConditiongroup OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum = "ConditionGroup"
	OciConsoleSignOnPolicyConsentModifiedResourceTypeCondition      OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum = "Condition"
)

var mappingOciConsoleSignOnPolicyConsentModifiedResourceTypeEnum = map[string]OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum{
	"Policy":         OciConsoleSignOnPolicyConsentModifiedResourceTypePolicy,
	"Rule":           OciConsoleSignOnPolicyConsentModifiedResourceTypeRule,
	"ConditionGroup": OciConsoleSignOnPolicyConsentModifiedResourceTypeConditiongroup,
	"Condition":      OciConsoleSignOnPolicyConsentModifiedResourceTypeCondition,
}

var mappingOciConsoleSignOnPolicyConsentModifiedResourceTypeEnumLowerCase = map[string]OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum{
	"policy":         OciConsoleSignOnPolicyConsentModifiedResourceTypePolicy,
	"rule":           OciConsoleSignOnPolicyConsentModifiedResourceTypeRule,
	"conditiongroup": OciConsoleSignOnPolicyConsentModifiedResourceTypeConditiongroup,
	"condition":      OciConsoleSignOnPolicyConsentModifiedResourceTypeCondition,
}

// GetOciConsoleSignOnPolicyConsentModifiedResourceTypeEnumValues Enumerates the set of values for OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum
func GetOciConsoleSignOnPolicyConsentModifiedResourceTypeEnumValues() []OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum {
	values := make([]OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum, 0)
	for _, v := range mappingOciConsoleSignOnPolicyConsentModifiedResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOciConsoleSignOnPolicyConsentModifiedResourceTypeEnumStringValues Enumerates the set of values in String for OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum
func GetOciConsoleSignOnPolicyConsentModifiedResourceTypeEnumStringValues() []string {
	return []string{
		"Policy",
		"Rule",
		"ConditionGroup",
		"Condition",
	}
}

// GetMappingOciConsoleSignOnPolicyConsentModifiedResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciConsoleSignOnPolicyConsentModifiedResourceTypeEnum(val string) (OciConsoleSignOnPolicyConsentModifiedResourceTypeEnum, bool) {
	enum, ok := mappingOciConsoleSignOnPolicyConsentModifiedResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
