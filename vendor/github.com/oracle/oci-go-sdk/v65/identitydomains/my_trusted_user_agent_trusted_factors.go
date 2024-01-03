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

// MyTrustedUserAgentTrustedFactors Trusted 2FA Factors
type MyTrustedUserAgentTrustedFactors struct {

	// Trusted Factor
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type *string `mandatory:"true" json:"type"`

	// trust factor creation time
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: true
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	CreationTime *string `mandatory:"true" json:"creationTime"`

	// Trusted Factor Type. Local, X509, SAML SOCIAL
	// **Added In:** 2111190457
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Category MyTrustedUserAgentTrustedFactorsCategoryEnum `mandatory:"false" json:"category,omitempty"`
}

func (m MyTrustedUserAgentTrustedFactors) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MyTrustedUserAgentTrustedFactors) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMyTrustedUserAgentTrustedFactorsCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetMyTrustedUserAgentTrustedFactorsCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MyTrustedUserAgentTrustedFactorsCategoryEnum Enum with underlying type: string
type MyTrustedUserAgentTrustedFactorsCategoryEnum string

// Set of constants representing the allowable values for MyTrustedUserAgentTrustedFactorsCategoryEnum
const (
	MyTrustedUserAgentTrustedFactorsCategorySaml       MyTrustedUserAgentTrustedFactorsCategoryEnum = "SAML"
	MyTrustedUserAgentTrustedFactorsCategoryLocal      MyTrustedUserAgentTrustedFactorsCategoryEnum = "LOCAL"
	MyTrustedUserAgentTrustedFactorsCategorySocial     MyTrustedUserAgentTrustedFactorsCategoryEnum = "SOCIAL"
	MyTrustedUserAgentTrustedFactorsCategoryX509       MyTrustedUserAgentTrustedFactorsCategoryEnum = "X509"
	MyTrustedUserAgentTrustedFactorsCategoryThirdparty MyTrustedUserAgentTrustedFactorsCategoryEnum = "THIRDPARTY"
)

var mappingMyTrustedUserAgentTrustedFactorsCategoryEnum = map[string]MyTrustedUserAgentTrustedFactorsCategoryEnum{
	"SAML":       MyTrustedUserAgentTrustedFactorsCategorySaml,
	"LOCAL":      MyTrustedUserAgentTrustedFactorsCategoryLocal,
	"SOCIAL":     MyTrustedUserAgentTrustedFactorsCategorySocial,
	"X509":       MyTrustedUserAgentTrustedFactorsCategoryX509,
	"THIRDPARTY": MyTrustedUserAgentTrustedFactorsCategoryThirdparty,
}

var mappingMyTrustedUserAgentTrustedFactorsCategoryEnumLowerCase = map[string]MyTrustedUserAgentTrustedFactorsCategoryEnum{
	"saml":       MyTrustedUserAgentTrustedFactorsCategorySaml,
	"local":      MyTrustedUserAgentTrustedFactorsCategoryLocal,
	"social":     MyTrustedUserAgentTrustedFactorsCategorySocial,
	"x509":       MyTrustedUserAgentTrustedFactorsCategoryX509,
	"thirdparty": MyTrustedUserAgentTrustedFactorsCategoryThirdparty,
}

// GetMyTrustedUserAgentTrustedFactorsCategoryEnumValues Enumerates the set of values for MyTrustedUserAgentTrustedFactorsCategoryEnum
func GetMyTrustedUserAgentTrustedFactorsCategoryEnumValues() []MyTrustedUserAgentTrustedFactorsCategoryEnum {
	values := make([]MyTrustedUserAgentTrustedFactorsCategoryEnum, 0)
	for _, v := range mappingMyTrustedUserAgentTrustedFactorsCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetMyTrustedUserAgentTrustedFactorsCategoryEnumStringValues Enumerates the set of values in String for MyTrustedUserAgentTrustedFactorsCategoryEnum
func GetMyTrustedUserAgentTrustedFactorsCategoryEnumStringValues() []string {
	return []string{
		"SAML",
		"LOCAL",
		"SOCIAL",
		"X509",
		"THIRDPARTY",
	}
}

// GetMappingMyTrustedUserAgentTrustedFactorsCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyTrustedUserAgentTrustedFactorsCategoryEnum(val string) (MyTrustedUserAgentTrustedFactorsCategoryEnum, bool) {
	enum, ok := mappingMyTrustedUserAgentTrustedFactorsCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
