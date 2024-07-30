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

// AuthenticationFactorSettingsCompliancePolicy Compliance Policy that defines actions to be taken when a condition is violated
type AuthenticationFactorSettingsCompliancePolicy struct {

	// The name of the attribute being evaluated
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Name *string `mandatory:"true" json:"name"`

	// The action to be taken if the value of the attribute is not as expected
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Action AuthenticationFactorSettingsCompliancePolicyActionEnum `mandatory:"true" json:"action"`

	// The value of the attribute to be evaluated
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`
}

func (m AuthenticationFactorSettingsCompliancePolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuthenticationFactorSettingsCompliancePolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuthenticationFactorSettingsCompliancePolicyActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetAuthenticationFactorSettingsCompliancePolicyActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuthenticationFactorSettingsCompliancePolicyActionEnum Enum with underlying type: string
type AuthenticationFactorSettingsCompliancePolicyActionEnum string

// Set of constants representing the allowable values for AuthenticationFactorSettingsCompliancePolicyActionEnum
const (
	AuthenticationFactorSettingsCompliancePolicyActionAllow  AuthenticationFactorSettingsCompliancePolicyActionEnum = "Allow"
	AuthenticationFactorSettingsCompliancePolicyActionBlock  AuthenticationFactorSettingsCompliancePolicyActionEnum = "Block"
	AuthenticationFactorSettingsCompliancePolicyActionNotify AuthenticationFactorSettingsCompliancePolicyActionEnum = "Notify"
	AuthenticationFactorSettingsCompliancePolicyActionNone   AuthenticationFactorSettingsCompliancePolicyActionEnum = "None"
)

var mappingAuthenticationFactorSettingsCompliancePolicyActionEnum = map[string]AuthenticationFactorSettingsCompliancePolicyActionEnum{
	"Allow":  AuthenticationFactorSettingsCompliancePolicyActionAllow,
	"Block":  AuthenticationFactorSettingsCompliancePolicyActionBlock,
	"Notify": AuthenticationFactorSettingsCompliancePolicyActionNotify,
	"None":   AuthenticationFactorSettingsCompliancePolicyActionNone,
}

var mappingAuthenticationFactorSettingsCompliancePolicyActionEnumLowerCase = map[string]AuthenticationFactorSettingsCompliancePolicyActionEnum{
	"allow":  AuthenticationFactorSettingsCompliancePolicyActionAllow,
	"block":  AuthenticationFactorSettingsCompliancePolicyActionBlock,
	"notify": AuthenticationFactorSettingsCompliancePolicyActionNotify,
	"none":   AuthenticationFactorSettingsCompliancePolicyActionNone,
}

// GetAuthenticationFactorSettingsCompliancePolicyActionEnumValues Enumerates the set of values for AuthenticationFactorSettingsCompliancePolicyActionEnum
func GetAuthenticationFactorSettingsCompliancePolicyActionEnumValues() []AuthenticationFactorSettingsCompliancePolicyActionEnum {
	values := make([]AuthenticationFactorSettingsCompliancePolicyActionEnum, 0)
	for _, v := range mappingAuthenticationFactorSettingsCompliancePolicyActionEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationFactorSettingsCompliancePolicyActionEnumStringValues Enumerates the set of values in String for AuthenticationFactorSettingsCompliancePolicyActionEnum
func GetAuthenticationFactorSettingsCompliancePolicyActionEnumStringValues() []string {
	return []string{
		"Allow",
		"Block",
		"Notify",
		"None",
	}
}

// GetMappingAuthenticationFactorSettingsCompliancePolicyActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationFactorSettingsCompliancePolicyActionEnum(val string) (AuthenticationFactorSettingsCompliancePolicyActionEnum, bool) {
	enum, ok := mappingAuthenticationFactorSettingsCompliancePolicyActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
