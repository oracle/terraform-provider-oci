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

// MyTrustedUserAgent This schema defines the attributes of Trusted User Agents owned by users. Multi-Factor Authentication uses Trusted User Agents to authenticate users.  A User Agent is software application that a user uses to issue requests.
//
//	For example, a User Agent could be a particular browser (possibly one of several executing on a desktop or laptop) or a particular mobile application (again, one of several executing on a particular mobile device).
//	A User Agent is trusted once the Multi-Factor Authentication has verified it in some way.
type MyTrustedUserAgent struct {

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

	// The name of the User Agent that the user wants the system to trust and to use in Multi-Factor Authentication.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Name *string `mandatory:"true" json:"name"`

	// Trust token for the user agent. This is a random string value that will be updated whenever a token that has been issued is verified successfully.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - idcsSensitive: none
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	TrustToken *string `mandatory:"true" json:"trustToken"`

	User *MyTrustedUserAgentUser `mandatory:"true" json:"user"`

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

	// Trust token issued geo-location.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Location *string `mandatory:"false" json:"location"`

	// User agent platform for which the trust token has been issued.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Platform *string `mandatory:"false" json:"platform"`

	// Validation period of the trust token.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	ExpiryTime *string `mandatory:"false" json:"expiryTime"`

	// Indicates when this token was used lastime.
	// **Added In:** 2111190457
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	LastUsedOn *string `mandatory:"false" json:"lastUsedOn"`

	// The token type being created. This token is used as trusted and kmsi token.
	// **Added In:** 2111190457
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	TokenType MyTrustedUserAgentTokenTypeEnum `mandatory:"false" json:"tokenType,omitempty"`

	// Trusted 2FA Factors
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCompositeKey: [type]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	TrustedFactors []MyTrustedUserAgentTrustedFactors `mandatory:"false" json:"trustedFactors"`
}

func (m MyTrustedUserAgent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MyTrustedUserAgent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingMyTrustedUserAgentTokenTypeEnum(string(m.TokenType)); !ok && m.TokenType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TokenType: %s. Supported values are: %s.", m.TokenType, strings.Join(GetMyTrustedUserAgentTokenTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MyTrustedUserAgentTokenTypeEnum Enum with underlying type: string
type MyTrustedUserAgentTokenTypeEnum string

// Set of constants representing the allowable values for MyTrustedUserAgentTokenTypeEnum
const (
	MyTrustedUserAgentTokenTypeKmsi    MyTrustedUserAgentTokenTypeEnum = "KMSI"
	MyTrustedUserAgentTokenTypeTrusted MyTrustedUserAgentTokenTypeEnum = "TRUSTED"
)

var mappingMyTrustedUserAgentTokenTypeEnum = map[string]MyTrustedUserAgentTokenTypeEnum{
	"KMSI":    MyTrustedUserAgentTokenTypeKmsi,
	"TRUSTED": MyTrustedUserAgentTokenTypeTrusted,
}

var mappingMyTrustedUserAgentTokenTypeEnumLowerCase = map[string]MyTrustedUserAgentTokenTypeEnum{
	"kmsi":    MyTrustedUserAgentTokenTypeKmsi,
	"trusted": MyTrustedUserAgentTokenTypeTrusted,
}

// GetMyTrustedUserAgentTokenTypeEnumValues Enumerates the set of values for MyTrustedUserAgentTokenTypeEnum
func GetMyTrustedUserAgentTokenTypeEnumValues() []MyTrustedUserAgentTokenTypeEnum {
	values := make([]MyTrustedUserAgentTokenTypeEnum, 0)
	for _, v := range mappingMyTrustedUserAgentTokenTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMyTrustedUserAgentTokenTypeEnumStringValues Enumerates the set of values in String for MyTrustedUserAgentTokenTypeEnum
func GetMyTrustedUserAgentTokenTypeEnumStringValues() []string {
	return []string{
		"KMSI",
		"TRUSTED",
	}
}

// GetMappingMyTrustedUserAgentTokenTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyTrustedUserAgentTokenTypeEnum(val string) (MyTrustedUserAgentTokenTypeEnum, bool) {
	enum, ok := mappingMyTrustedUserAgentTokenTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
