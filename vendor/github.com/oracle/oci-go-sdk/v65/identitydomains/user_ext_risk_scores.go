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

// UserExtRiskScores The risk score pertaining to the user.
// **Added In:** 18.1.6
// **SCIM++ Properties:**
//   - caseExact: false
//   - idcsCompositeKey: [value]
//   - multiValued: true
//   - mutability: readWrite
//   - required: false
//   - returned: request
//   - type: complex
//   - uniqueness: none
type UserExtRiskScores struct {

	// Risk Provider Profile: Identifier for the provider service from which the risk score was received.
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// Risk Score value
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: integer
	//  - uniqueness: none
	//  - idcsMaxValue: 100
	//  - idcsMinValue: 0
	Score *int `mandatory:"true" json:"score"`

	// Risk Level
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	RiskLevel UserExtRiskScoresRiskLevelEnum `mandatory:"true" json:"riskLevel"`

	// Last update timestamp for the risk score
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: dateTime
	//  - uniqueness: none
	LastUpdateTimestamp *string `mandatory:"true" json:"lastUpdateTimestamp"`

	// Risk Provider Profile URI: URI that corresponds to risk source identifier.
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// Risk Provider Profile Source
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Source *string `mandatory:"false" json:"source"`

	// Risk Provider Profile status
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Status *string `mandatory:"false" json:"status"`
}

func (m UserExtRiskScores) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserExtRiskScores) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserExtRiskScoresRiskLevelEnum(string(m.RiskLevel)); !ok && m.RiskLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RiskLevel: %s. Supported values are: %s.", m.RiskLevel, strings.Join(GetUserExtRiskScoresRiskLevelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserExtRiskScoresRiskLevelEnum Enum with underlying type: string
type UserExtRiskScoresRiskLevelEnum string

// Set of constants representing the allowable values for UserExtRiskScoresRiskLevelEnum
const (
	UserExtRiskScoresRiskLevelLow    UserExtRiskScoresRiskLevelEnum = "LOW"
	UserExtRiskScoresRiskLevelMedium UserExtRiskScoresRiskLevelEnum = "MEDIUM"
	UserExtRiskScoresRiskLevelHigh   UserExtRiskScoresRiskLevelEnum = "HIGH"
)

var mappingUserExtRiskScoresRiskLevelEnum = map[string]UserExtRiskScoresRiskLevelEnum{
	"LOW":    UserExtRiskScoresRiskLevelLow,
	"MEDIUM": UserExtRiskScoresRiskLevelMedium,
	"HIGH":   UserExtRiskScoresRiskLevelHigh,
}

var mappingUserExtRiskScoresRiskLevelEnumLowerCase = map[string]UserExtRiskScoresRiskLevelEnum{
	"low":    UserExtRiskScoresRiskLevelLow,
	"medium": UserExtRiskScoresRiskLevelMedium,
	"high":   UserExtRiskScoresRiskLevelHigh,
}

// GetUserExtRiskScoresRiskLevelEnumValues Enumerates the set of values for UserExtRiskScoresRiskLevelEnum
func GetUserExtRiskScoresRiskLevelEnumValues() []UserExtRiskScoresRiskLevelEnum {
	values := make([]UserExtRiskScoresRiskLevelEnum, 0)
	for _, v := range mappingUserExtRiskScoresRiskLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetUserExtRiskScoresRiskLevelEnumStringValues Enumerates the set of values in String for UserExtRiskScoresRiskLevelEnum
func GetUserExtRiskScoresRiskLevelEnumStringValues() []string {
	return []string{
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

// GetMappingUserExtRiskScoresRiskLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserExtRiskScoresRiskLevelEnum(val string) (UserExtRiskScoresRiskLevelEnum, bool) {
	enum, ok := mappingUserExtRiskScoresRiskLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
