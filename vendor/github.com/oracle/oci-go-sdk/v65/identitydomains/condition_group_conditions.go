// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ConditionGroupConditions Conditions or ConditionGroup assigned to the conditionGroup
type ConditionGroupConditions struct {

	// Condition or ConditionGroup identifier
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// A label that indicates whether this is Condition or ConditionGroup.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - idcsDefaultValue: Condition
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type ConditionGroupConditionsTypeEnum `mandatory:"true" json:"type"`

	// Condition URI
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
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

func (m ConditionGroupConditions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConditionGroupConditions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConditionGroupConditionsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetConditionGroupConditionsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConditionGroupConditionsTypeEnum Enum with underlying type: string
type ConditionGroupConditionsTypeEnum string

// Set of constants representing the allowable values for ConditionGroupConditionsTypeEnum
const (
	ConditionGroupConditionsTypeCondition      ConditionGroupConditionsTypeEnum = "Condition"
	ConditionGroupConditionsTypeConditiongroup ConditionGroupConditionsTypeEnum = "ConditionGroup"
)

var mappingConditionGroupConditionsTypeEnum = map[string]ConditionGroupConditionsTypeEnum{
	"Condition":      ConditionGroupConditionsTypeCondition,
	"ConditionGroup": ConditionGroupConditionsTypeConditiongroup,
}

var mappingConditionGroupConditionsTypeEnumLowerCase = map[string]ConditionGroupConditionsTypeEnum{
	"condition":      ConditionGroupConditionsTypeCondition,
	"conditiongroup": ConditionGroupConditionsTypeConditiongroup,
}

// GetConditionGroupConditionsTypeEnumValues Enumerates the set of values for ConditionGroupConditionsTypeEnum
func GetConditionGroupConditionsTypeEnumValues() []ConditionGroupConditionsTypeEnum {
	values := make([]ConditionGroupConditionsTypeEnum, 0)
	for _, v := range mappingConditionGroupConditionsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConditionGroupConditionsTypeEnumStringValues Enumerates the set of values in String for ConditionGroupConditionsTypeEnum
func GetConditionGroupConditionsTypeEnumStringValues() []string {
	return []string{
		"Condition",
		"ConditionGroup",
	}
}

// GetMappingConditionGroupConditionsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConditionGroupConditionsTypeEnum(val string) (ConditionGroupConditionsTypeEnum, bool) {
	enum, ok := mappingConditionGroupConditionsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
