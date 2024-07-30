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

// Condition Condition resource. A unit that captures a condition.
type Condition struct {

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

	// Condition name
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: global
	Name *string `mandatory:"true" json:"name"`

	// AttributeName - RHS of condition
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// **SCIM++ Properties:**
	// - caseExact: true
	// - idcsSearchable: false
	// - multiValued: false
	// - mutability: readWrite
	// - required: true
	// - returned: default
	// - type: string
	// - uniqueness: none
	// Operator in the condition. It support all SCIM operators like eq, gt, lt, le, sw etc
	Operator ConditionOperatorEnum `mandatory:"true" json:"operator"`

	// attributeValue - RHS of condition
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AttributeValue *string `mandatory:"true" json:"attributeValue"`

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

	// An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value.  The value of the externalId attribute is always issued be the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ExternalId *string `mandatory:"false" json:"externalId"`

	// Condition Description
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Description *string `mandatory:"false" json:"description"`

	// Evaluate the condition if this expression returns true, else skip condition evaluation
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	EvaluateConditionIf *string `mandatory:"false" json:"evaluateConditionIf"`
}

func (m Condition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Condition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConditionOperatorEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetConditionOperatorEnumStringValues(), ",")))
	}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConditionOperatorEnum Enum with underlying type: string
type ConditionOperatorEnum string

// Set of constants representing the allowable values for ConditionOperatorEnum
const (
	ConditionOperatorEq    ConditionOperatorEnum = "eq"
	ConditionOperatorNe    ConditionOperatorEnum = "ne"
	ConditionOperatorCo    ConditionOperatorEnum = "co"
	ConditionOperatorCoany ConditionOperatorEnum = "coany"
	ConditionOperatorSw    ConditionOperatorEnum = "sw"
	ConditionOperatorEw    ConditionOperatorEnum = "ew"
	ConditionOperatorGt    ConditionOperatorEnum = "gt"
	ConditionOperatorGe    ConditionOperatorEnum = "ge"
	ConditionOperatorLt    ConditionOperatorEnum = "lt"
	ConditionOperatorLe    ConditionOperatorEnum = "le"
	ConditionOperatorIn    ConditionOperatorEnum = "in"
	ConditionOperatorNin   ConditionOperatorEnum = "nin"
)

var mappingConditionOperatorEnum = map[string]ConditionOperatorEnum{
	"eq":    ConditionOperatorEq,
	"ne":    ConditionOperatorNe,
	"co":    ConditionOperatorCo,
	"coany": ConditionOperatorCoany,
	"sw":    ConditionOperatorSw,
	"ew":    ConditionOperatorEw,
	"gt":    ConditionOperatorGt,
	"ge":    ConditionOperatorGe,
	"lt":    ConditionOperatorLt,
	"le":    ConditionOperatorLe,
	"in":    ConditionOperatorIn,
	"nin":   ConditionOperatorNin,
}

var mappingConditionOperatorEnumLowerCase = map[string]ConditionOperatorEnum{
	"eq":    ConditionOperatorEq,
	"ne":    ConditionOperatorNe,
	"co":    ConditionOperatorCo,
	"coany": ConditionOperatorCoany,
	"sw":    ConditionOperatorSw,
	"ew":    ConditionOperatorEw,
	"gt":    ConditionOperatorGt,
	"ge":    ConditionOperatorGe,
	"lt":    ConditionOperatorLt,
	"le":    ConditionOperatorLe,
	"in":    ConditionOperatorIn,
	"nin":   ConditionOperatorNin,
}

// GetConditionOperatorEnumValues Enumerates the set of values for ConditionOperatorEnum
func GetConditionOperatorEnumValues() []ConditionOperatorEnum {
	values := make([]ConditionOperatorEnum, 0)
	for _, v := range mappingConditionOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetConditionOperatorEnumStringValues Enumerates the set of values in String for ConditionOperatorEnum
func GetConditionOperatorEnumStringValues() []string {
	return []string{
		"eq",
		"ne",
		"co",
		"coany",
		"sw",
		"ew",
		"gt",
		"ge",
		"lt",
		"le",
		"in",
		"nin",
	}
}

// GetMappingConditionOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConditionOperatorEnum(val string) (ConditionOperatorEnum, bool) {
	enum, ok := mappingConditionOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
