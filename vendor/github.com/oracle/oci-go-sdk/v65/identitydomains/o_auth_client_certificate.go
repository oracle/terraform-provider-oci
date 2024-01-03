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

// OAuthClientCertificate OAuth Client Certificate schema
type OAuthClientCertificate struct {

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

	// Certificate alias
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: true
	//  - caseExact: false
	//  - mutability: readWrite
	//  - returned: always
	//  - uniqueness: none
	//  - idcsSearchable: true
	CertificateAlias *string `mandatory:"true" json:"certificateAlias"`

	// Base 64Key data attribute
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	X509Base64Certificate *string `mandatory:"true" json:"x509Base64Certificate"`

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

	// An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - caseExact: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	ExternalId *string `mandatory:"false" json:"externalId"`

	// Key store name
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - caseExact: false
	//  - returned: always
	//  - uniqueness: none
	KeyStoreName *string `mandatory:"false" json:"keyStoreName"`

	// Map
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - caseExact: false
	//  - returned: always
	//  - uniqueness: none
	Map *string `mandatory:"false" json:"map"`

	// Key store password
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - idcsSensitive: none
	//  - required: false
	//  - caseExact: true
	//  - returned: never
	//  - uniqueness: none
	KeyStorePassword *string `mandatory:"false" json:"keyStorePassword"`

	// Key store ID
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - caseExact: true
	//  - returned: never
	//  - uniqueness: none
	//  - idcsSearchable: true
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// SHA-1 Thumbprint
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - idcsSearchable: true
	//  - required: false
	//  - caseExact: true
	//  - returned: default
	//  - uniqueness: none
	Sha1Thumbprint *string `mandatory:"false" json:"sha1Thumbprint"`

	// SHA-256 Thumbprint
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - idcsSearchable: true
	//  - required: false
	//  - caseExact: true
	//  - returned: default
	//  - uniqueness: none
	Sha256Thumbprint *string `mandatory:"false" json:"sha256Thumbprint"`

	// Certificate start date
	// **SCIM++ Properties:**
	//  - type: dateTime
	//  - multiValued: false
	//  - required: false
	//  - caseExact: true
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	CertStartDate *string `mandatory:"false" json:"certStartDate"`

	// Certificate end date
	// **SCIM++ Properties:**
	//  - type: dateTime
	//  - multiValued: false
	//  - required: false
	//  - caseExact: true
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	CertEndDate *string `mandatory:"false" json:"certEndDate"`
}

func (m OAuthClientCertificate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OAuthClientCertificate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

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
