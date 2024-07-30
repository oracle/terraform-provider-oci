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

// ExtensionAdaptiveUser This extension defines attributes to manage user's risk score.
type ExtensionAdaptiveUser struct {

	// Risk Level
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	RiskLevel ExtensionAdaptiveUserRiskLevelEnum `mandatory:"false" json:"riskLevel,omitempty"`

	// The risk score pertaining to the user.
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [value]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	RiskScores []UserExtRiskScores `mandatory:"false" json:"riskScores"`
}

func (m ExtensionAdaptiveUser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionAdaptiveUser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExtensionAdaptiveUserRiskLevelEnum(string(m.RiskLevel)); !ok && m.RiskLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RiskLevel: %s. Supported values are: %s.", m.RiskLevel, strings.Join(GetExtensionAdaptiveUserRiskLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExtensionAdaptiveUserRiskLevelEnum Enum with underlying type: string
type ExtensionAdaptiveUserRiskLevelEnum string

// Set of constants representing the allowable values for ExtensionAdaptiveUserRiskLevelEnum
const (
	ExtensionAdaptiveUserRiskLevelLow    ExtensionAdaptiveUserRiskLevelEnum = "LOW"
	ExtensionAdaptiveUserRiskLevelMedium ExtensionAdaptiveUserRiskLevelEnum = "MEDIUM"
	ExtensionAdaptiveUserRiskLevelHigh   ExtensionAdaptiveUserRiskLevelEnum = "HIGH"
)

var mappingExtensionAdaptiveUserRiskLevelEnum = map[string]ExtensionAdaptiveUserRiskLevelEnum{
	"LOW":    ExtensionAdaptiveUserRiskLevelLow,
	"MEDIUM": ExtensionAdaptiveUserRiskLevelMedium,
	"HIGH":   ExtensionAdaptiveUserRiskLevelHigh,
}

var mappingExtensionAdaptiveUserRiskLevelEnumLowerCase = map[string]ExtensionAdaptiveUserRiskLevelEnum{
	"low":    ExtensionAdaptiveUserRiskLevelLow,
	"medium": ExtensionAdaptiveUserRiskLevelMedium,
	"high":   ExtensionAdaptiveUserRiskLevelHigh,
}

// GetExtensionAdaptiveUserRiskLevelEnumValues Enumerates the set of values for ExtensionAdaptiveUserRiskLevelEnum
func GetExtensionAdaptiveUserRiskLevelEnumValues() []ExtensionAdaptiveUserRiskLevelEnum {
	values := make([]ExtensionAdaptiveUserRiskLevelEnum, 0)
	for _, v := range mappingExtensionAdaptiveUserRiskLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionAdaptiveUserRiskLevelEnumStringValues Enumerates the set of values in String for ExtensionAdaptiveUserRiskLevelEnum
func GetExtensionAdaptiveUserRiskLevelEnumStringValues() []string {
	return []string{
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

// GetMappingExtensionAdaptiveUserRiskLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionAdaptiveUserRiskLevelEnum(val string) (ExtensionAdaptiveUserRiskLevelEnum, bool) {
	enum, ok := mappingExtensionAdaptiveUserRiskLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
