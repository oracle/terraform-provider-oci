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

// OciConsoleSignOnPolicyConsentConsentSignedBy User or App that signs the consent.
// **SCIM++ Properties:**
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: immutable
//   - required: true
//   - returned: default
//   - type: complex
//   - uniqueness: none
type OciConsoleSignOnPolicyConsentConsentSignedBy struct {

	// Id of the User or App that signed consent.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	Value *string `mandatory:"true" json:"value"`

	// OCID of the User or App that signed consent.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	Ocid *string `mandatory:"true" json:"ocid"`

	// Name of the User or App that signed consent.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of principal that signed consent: User or App.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	Type OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum `mandatory:"true" json:"type"`
}

func (m OciConsoleSignOnPolicyConsentConsentSignedBy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciConsoleSignOnPolicyConsentConsentSignedBy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciConsoleSignOnPolicyConsentConsentSignedByTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetOciConsoleSignOnPolicyConsentConsentSignedByTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum Enum with underlying type: string
type OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum string

// Set of constants representing the allowable values for OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum
const (
	OciConsoleSignOnPolicyConsentConsentSignedByTypeUser OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum = "User"
	OciConsoleSignOnPolicyConsentConsentSignedByTypeApp  OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum = "App"
)

var mappingOciConsoleSignOnPolicyConsentConsentSignedByTypeEnum = map[string]OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum{
	"User": OciConsoleSignOnPolicyConsentConsentSignedByTypeUser,
	"App":  OciConsoleSignOnPolicyConsentConsentSignedByTypeApp,
}

var mappingOciConsoleSignOnPolicyConsentConsentSignedByTypeEnumLowerCase = map[string]OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum{
	"user": OciConsoleSignOnPolicyConsentConsentSignedByTypeUser,
	"app":  OciConsoleSignOnPolicyConsentConsentSignedByTypeApp,
}

// GetOciConsoleSignOnPolicyConsentConsentSignedByTypeEnumValues Enumerates the set of values for OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum
func GetOciConsoleSignOnPolicyConsentConsentSignedByTypeEnumValues() []OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum {
	values := make([]OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum, 0)
	for _, v := range mappingOciConsoleSignOnPolicyConsentConsentSignedByTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOciConsoleSignOnPolicyConsentConsentSignedByTypeEnumStringValues Enumerates the set of values in String for OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum
func GetOciConsoleSignOnPolicyConsentConsentSignedByTypeEnumStringValues() []string {
	return []string{
		"User",
		"App",
	}
}

// GetMappingOciConsoleSignOnPolicyConsentConsentSignedByTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciConsoleSignOnPolicyConsentConsentSignedByTypeEnum(val string) (OciConsoleSignOnPolicyConsentConsentSignedByTypeEnum, bool) {
	enum, ok := mappingOciConsoleSignOnPolicyConsentConsentSignedByTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
