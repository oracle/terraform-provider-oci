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

// RestoreOciConsolePolicy Schema to restoring OCI Console Policy to Factory Defaults.
type RestoreOciConsolePolicy struct {

	// **SCIM++ Properties:**
	// - caseExact: false
	// - idcsSearchable: false
	// - multiValued: true
	// - mutability: readWrite
	// - required: true
	// - returned: default
	// - type: string
	// - uniqueness: none
	// REQUIRED. The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. This specification defines URIs for User, Group, and a standard \"enterprise\" extension. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior.
	Schemas []string `mandatory:"true" json:"schemas"`

	// **SCIM++ Properties:**
	// - idcsSearchable: false
	// - multiValued: false
	// - required: true
	// - mutability: writeOnly
	// - returned: default
	// - type: boolean
	// Consent to be provided for restoring the Oci Console SignOn Policy to Factory Defaults. Defaults to false
	Consent *bool `mandatory:"true" json:"consent"`

	// **SCIM++ Properties:**
	// - idcsSearchable: false
	// - multiValued: false
	// - mutability: writeOnly
	// - required: true
	// - returned: default
	// - type: string
	// Detailed reason when domain admin opts to restore the Oci Console SignOn Policy to Factory Defaults
	Reason *string `mandatory:"true" json:"reason"`

	// **SCIM++ Properties:**
	// - caseExact: false
	// - idcsSearchable: true
	// - multiValued: false
	// - mutability: readOnly
	// - required: false
	// - returned: always
	// - type: string
	// - uniqueness: global
	// Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.
	Id *string `mandatory:"false" json:"id"`

	// **SCIM++ Properties:**
	// - caseExact: true
	// - idcsSearchable: true
	// - multiValued: false
	// - mutability: immutable
	// - required: false
	// - returned: default
	// - type: string
	// - uniqueness: global
	// Unique OCI identifier for the SCIM Resource.
	Ocid *string `mandatory:"false" json:"ocid"`

	Meta *Meta `mandatory:"false" json:"meta"`

	IdcsCreatedBy *IdcsCreatedBy `mandatory:"false" json:"idcsCreatedBy"`

	IdcsLastModifiedBy *IdcsLastModifiedBy `mandatory:"false" json:"idcsLastModifiedBy"`

	// **SCIM++ Properties:**
	// - idcsSearchable: false
	// - multiValued: true
	// - mutability: readOnly
	// - required: false
	// - returned: request
	// - type: string
	// - uniqueness: none
	// Each value of this attribute specifies an operation that only an internal client may perform on this particular resource.
	IdcsPreventedOperations []IdcsPreventedOperationsEnum `mandatory:"false" json:"idcsPreventedOperations,omitempty"`

	// **SCIM++ Properties:**
	// - idcsCompositeKey: [key, value]
	// - idcsCsvAttributeNameMappings: [[columnHeaderName:Tag Key, mapsTo:tags.key], [columnHeaderName:Tag Value, mapsTo:tags.value]]
	// - idcsSearchable: true
	// - multiValued: true
	// - mutability: readWrite
	// - required: false
	// - returned: request
	// - type: complex
	// - uniqueness: none
	// A list of tags on this resource.
	Tags []Tags `mandatory:"false" json:"tags"`

	// **SCIM++ Properties:**
	// - caseExact: false
	// - idcsSearchable: true
	// - multiValued: false
	// - mutability: readOnly
	// - required: false
	// - returned: default
	// - type: boolean
	// - uniqueness: none
	// A boolean flag indicating this resource in the process of being deleted. Usually set to true when synchronous deletion of the resource would take too long.
	DeleteInProgress *bool `mandatory:"false" json:"deleteInProgress"`

	// **SCIM++ Properties:**
	// - caseExact: false
	// - idcsSearchable: false
	// - multiValued: false
	// - mutability: readOnly
	// - required: false
	// - returned: request
	// - type: string
	// - uniqueness: none
	// The release number when the resource was upgraded.
	IdcsLastUpgradedInRelease *string `mandatory:"false" json:"idcsLastUpgradedInRelease"`

	// **SCIM++ Properties:**
	// - caseExact: false
	// - idcsSearchable: false
	// - multiValued: false
	// - mutability: readOnly
	// - required: false
	// - returned: default
	// - type: string
	// - uniqueness: none
	// OCI Domain Id (ocid) in which the resource lives.
	DomainOcid *string `mandatory:"false" json:"domainOcid"`

	// **SCIM++ Properties:**
	// - caseExact: false
	// - idcsSearchable: false
	// - multiValued: false
	// - mutability: readOnly
	// - required: false
	// - returned: default
	// - type: string
	// - uniqueness: none
	// OCI Compartment Id (ocid) in which the resource lives.
	CompartmentOcid *string `mandatory:"false" json:"compartmentOcid"`

	// **SCIM++ Properties:**
	// - caseExact: false
	// - idcsSearchable: false
	// - multiValued: false
	// - mutability: readOnly
	// - required: false
	// - returned: default
	// - type: string
	// - uniqueness: none
	// OCI Tenant Id (ocid) in which the resource lives.
	TenancyOcid *string `mandatory:"false" json:"tenancyOcid"`
}

func (m RestoreOciConsolePolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestoreOciConsolePolicy) ValidateEnumValue() (bool, error) {
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
