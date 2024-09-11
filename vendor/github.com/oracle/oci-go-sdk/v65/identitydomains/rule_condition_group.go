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

// RuleConditionGroup Condition or ConditionGroup assigned to the rule
// **SCIM++ Properties:**
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: readWrite
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type RuleConditionGroup struct {

	// A label that indicates whether this is Condition or ConditionGroup.
	// **SCIM++ Properties:**
	//  - idcsDefaultValue: Condition
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type RuleConditionGroupTypeEnum `mandatory:"true" json:"type"`

	// Condition or ConditionGroup identifier.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"false" json:"value"`

	// ConditionGroup URI
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// Condition or ConditionGroup name
	// **Added In:** 17.4.2
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Name *string `mandatory:"false" json:"name"`
}

func (m RuleConditionGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RuleConditionGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRuleConditionGroupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRuleConditionGroupTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RuleConditionGroupTypeEnum Enum with underlying type: string
type RuleConditionGroupTypeEnum string

// Set of constants representing the allowable values for RuleConditionGroupTypeEnum
const (
	RuleConditionGroupTypeCondition      RuleConditionGroupTypeEnum = "Condition"
	RuleConditionGroupTypeConditiongroup RuleConditionGroupTypeEnum = "ConditionGroup"
)

var mappingRuleConditionGroupTypeEnum = map[string]RuleConditionGroupTypeEnum{
	"Condition":      RuleConditionGroupTypeCondition,
	"ConditionGroup": RuleConditionGroupTypeConditiongroup,
}

var mappingRuleConditionGroupTypeEnumLowerCase = map[string]RuleConditionGroupTypeEnum{
	"condition":      RuleConditionGroupTypeCondition,
	"conditiongroup": RuleConditionGroupTypeConditiongroup,
}

// GetRuleConditionGroupTypeEnumValues Enumerates the set of values for RuleConditionGroupTypeEnum
func GetRuleConditionGroupTypeEnumValues() []RuleConditionGroupTypeEnum {
	values := make([]RuleConditionGroupTypeEnum, 0)
	for _, v := range mappingRuleConditionGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleConditionGroupTypeEnumStringValues Enumerates the set of values in String for RuleConditionGroupTypeEnum
func GetRuleConditionGroupTypeEnumStringValues() []string {
	return []string{
		"Condition",
		"ConditionGroup",
	}
}

// GetMappingRuleConditionGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleConditionGroupTypeEnum(val string) (RuleConditionGroupTypeEnum, bool) {
	enum, ok := mappingRuleConditionGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
