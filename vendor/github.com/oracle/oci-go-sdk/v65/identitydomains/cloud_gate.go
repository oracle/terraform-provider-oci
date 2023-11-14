// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// CloudGate Resource representing a Cloud Gate
type CloudGate struct {

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

	// Display name/Host identifier for this Cloud Gate
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: server
	DisplayName *string `mandatory:"true" json:"displayName"`

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

	// Indicates whether this resource was created by OPC
	// **Added In:** 19.3.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: boolean
	//  - uniqueness: none
	IsOPCService *bool `mandatory:"false" json:"isOPCService"`

	// Last updated timestamp for this CloudGate's servers and mappings.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	LastModifiedTime *string `mandatory:"false" json:"lastModifiedTime"`

	// Brief description for this Cloud Gate
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Description *string `mandatory:"false" json:"description"`

	// Activation status for this Cloud Gate
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Active *bool `mandatory:"false" json:"active"`

	// Type of Cloud Gate
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type CloudGateTypeEnum `mandatory:"false" json:"type,omitempty"`

	// A list of upstream server groups
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	UpstreamServerGroups []CloudGateExtUpstreamServerGroups `mandatory:"false" json:"upstreamServerGroups"`

	// A list of upstream servers
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	UpstreamServers []CloudGateExtUpstreamServers `mandatory:"false" json:"upstreamServers"`

	OauthClient *CloudGateOauthClient `mandatory:"false" json:"oauthClient"`

	// A list of Server Blocks on this Cloud Gate
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCompositeKey: [hostName]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Servers []CloudGateExtServers `mandatory:"false" json:"servers"`

	// A list of Cloud Gate Mappings that map Apps to this Cloud Gate
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCompositeKey: [serverId, resourcePrefix]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Mappings []CloudGateExtMappings `mandatory:"false" json:"mappings"`
}

func (m CloudGate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudGate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingCloudGateTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCloudGateTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudGateTypeEnum Enum with underlying type: string
type CloudGateTypeEnum string

// Set of constants representing the allowable values for CloudGateTypeEnum
const (
	CloudGateTypeLbaas   CloudGateTypeEnum = "lbaas"
	CloudGateTypeGateway CloudGateTypeEnum = "gateway"
)

var mappingCloudGateTypeEnum = map[string]CloudGateTypeEnum{
	"lbaas":   CloudGateTypeLbaas,
	"gateway": CloudGateTypeGateway,
}

var mappingCloudGateTypeEnumLowerCase = map[string]CloudGateTypeEnum{
	"lbaas":   CloudGateTypeLbaas,
	"gateway": CloudGateTypeGateway,
}

// GetCloudGateTypeEnumValues Enumerates the set of values for CloudGateTypeEnum
func GetCloudGateTypeEnumValues() []CloudGateTypeEnum {
	values := make([]CloudGateTypeEnum, 0)
	for _, v := range mappingCloudGateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudGateTypeEnumStringValues Enumerates the set of values in String for CloudGateTypeEnum
func GetCloudGateTypeEnumStringValues() []string {
	return []string{
		"lbaas",
		"gateway",
	}
}

// GetMappingCloudGateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudGateTypeEnum(val string) (CloudGateTypeEnum, bool) {
	enum, ok := mappingCloudGateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
