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

// ResourceTypeSchemaAttribute ResourceTypeSchemaAttribute Schema Definition
type ResourceTypeSchemaAttribute struct {

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

	// Attribute's name
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - idcsSearchable: true
	//  - type: string
	//  - uniqueness: none
	Name *string `mandatory:"false" json:"name"`

	// Localized schema attribute display name for use by UI client  for displaying attribute labels
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - idcsSearchable: true
	//  - type: string
	//  - uniqueness: none
	IdcsDisplayNameMessageId *string `mandatory:"false" json:"idcsDisplayNameMessageId"`

	// ResourceType this attribute belongs to.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - idcsSearchable: true
	//  - uniqueness: none
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// Schema URN string that this attribute belongs to
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - idcsSearchable: true
	//  - uniqueness: none
	IdcsSchemaUrn *string `mandatory:"false" json:"idcsSchemaUrn"`

	// Fully qualified name of this attribute
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - idcsSearchable: true
	//  - uniqueness: none
	IdcsFullyQualifiedName *string `mandatory:"false" json:"idcsFullyQualifiedName"`

	// custom attribute flag.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - idcsSearchable: true
	//  - uniqueness: none
	IdcsCustomAttribute *bool `mandatory:"false" json:"idcsCustomAttribute"`

	// The attribute's data type--for example, String
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type ResourceTypeSchemaAttributeTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Indicates the attribute's plurality
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: boolean
	MultiValued *bool `mandatory:"false" json:"multiValued"`

	// The attribute's human-readable description
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - idcsSearchable: true
	//  - type: string
	//  - uniqueness: none
	Description *string `mandatory:"false" json:"description"`

	// Specifies if the attribute is required
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: boolean
	Required *bool `mandatory:"false" json:"required"`

	// A collection of canonical values. Applicable Service Providers MUST specify the canonical types specified in the core schema specification--for example, \"work\", \"home\".
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CanonicalValues []string `mandatory:"false" json:"canonicalValues"`

	// Specifies if the String attribute is case-sensitive
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: boolean
	CaseExact *bool `mandatory:"false" json:"caseExact"`

	// Specifies if the attribute is mutable
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - idcsSearchable: true
	//  - type: string
	//  - uniqueness: none
	Mutability ResourceTypeSchemaAttributeMutabilityEnum `mandatory:"false" json:"mutability,omitempty"`

	// Specifies User mutability for this attribute
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	EndUserMutability ResourceTypeSchemaAttributeEndUserMutabilityEnum `mandatory:"false" json:"endUserMutability,omitempty"`

	// Specifies the list of User mutabilities allowed
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	EndUserMutabilityAllowedValues []ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum `mandatory:"false" json:"endUserMutabilityAllowedValues,omitempty"`

	// A single keyword that indicates when an attribute and associated values are returned in response to a GET request or in response to a PUT, POST, or PATCH request
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - idcsSearchable: true
	//  - type: string
	//  - uniqueness: none
	Returned ResourceTypeSchemaAttributeReturnedEnum `mandatory:"false" json:"returned,omitempty"`

	// A single keyword value that specifies how the Service Provider enforces uniqueness of attribute values. A server MAY reject an invalid value based on uniqueness by returning an HTTP response code of 400 (Bad Request). A client MAY enforce uniqueness on the client side to a greater degree than the Service Provider enforces. For example, a client could make a value unique while the server has the uniqueness of \"none\".
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - required: false
	//  - returned: default
	//  - idcsSearchable: true
	//  - type: string
	//  - uniqueness: none
	Uniqueness ResourceTypeSchemaAttributeUniquenessEnum `mandatory:"false" json:"uniqueness,omitempty"`

	// The attribute defining the CSV column header name for import/export
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsCsvColumnHeaderName *string `mandatory:"false" json:"idcsCsvColumnHeaderName"`

	// Maps to ICF target attribute name
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsICFBundleAttributeName *string `mandatory:"false" json:"idcsICFBundleAttributeName"`

	// Metadata to identify the ICF required attribute
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsICFRequired *bool `mandatory:"false" json:"idcsICFRequired"`

	// Maps to ICF data type
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsSearchable: true
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsICFAttributeType ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum `mandatory:"false" json:"idcsICFAttributeType,omitempty"`

	// The names of the Resource types that may be referenced--for example, User. This is only applicable for attributes that are of the \"reference\" data type.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ReferenceTypes []string `mandatory:"false" json:"referenceTypes"`

	// Indicates that the schema has been deprecated since version
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: integer
	IdcsDeprecatedSinceVersion *int `mandatory:"false" json:"idcsDeprecatedSinceVersion"`

	// Indicates that the schema has been added since version
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: integer
	IdcsAddedSinceVersion *int `mandatory:"false" json:"idcsAddedSinceVersion"`

	// Indicates that the schema has been deprecated since this release number
	// **Added In:** 17.3.4
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsSearchable: true
	//  - required: false
	//  - returned: default
	//  - type: string
	IdcsDeprecatedSinceReleaseNumber *string `mandatory:"false" json:"idcsDeprecatedSinceReleaseNumber"`

	// Indicates that the schema has been added since this release number
	// **Added In:** 17.3.4
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsSearchable: true
	//  - required: false
	//  - returned: default
	//  - type: string
	IdcsAddedSinceReleaseNumber *string `mandatory:"false" json:"idcsAddedSinceReleaseNumber"`

	// Specifies the minimum length of the attribute
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsSearchable: true
	//  - required: false
	//  - returned: default
	//  - type: integer
	IdcsMinLength *int `mandatory:"false" json:"idcsMinLength"`

	// Specifies the maximum length of the attribute
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: integer
	IdcsMaxLength *int `mandatory:"false" json:"idcsMaxLength"`

	// Specifies the minimum value of the integer attribute
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: integer
	IdcsMinValue *int `mandatory:"false" json:"idcsMinValue"`

	// Specifies the maximum value of the integer attribute
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsSearchable: true
	//  - required: false
	//  - returned: default
	//  - type: integer
	IdcsMaxValue *int `mandatory:"false" json:"idcsMaxValue"`

	// If true, specifies that the attribute can have multiple language values set for the attribute on which this is set.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsSearchable: true
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsMultiLanguage *bool `mandatory:"false" json:"idcsMultiLanguage"`

	// Specifies the directly referenced Resources
	// **SCIM++ Properties:**
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	IdcsRefResourceAttributes []string `mandatory:"false" json:"idcsRefResourceAttributes"`

	// Specifies the indirectly referenced Resources
	// **SCIM++ Properties:**
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	IdcsIndirectRefResourceAttributes []string `mandatory:"false" json:"idcsIndirectRefResourceAttributes"`

	// Sequence tracking ID name for the attribute
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	IdcsAutoIncrementSeqName *string `mandatory:"false" json:"idcsAutoIncrementSeqName"`

	// Specifies whether the value of the Resource attribute is persisted
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: boolean
	IdcsValuePersisted *bool `mandatory:"false" json:"idcsValuePersisted"`

	// Flag to specify if the attribute should be encrypted or hashed
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsSensitive ResourceTypeSchemaAttributeIdcsSensitiveEnum `mandatory:"false" json:"idcsSensitive,omitempty"`

	// Specifies whether the schema attribute is for internal use only. Internal attributes are not exposed via REST. This attribute overrides mutability for create/update if the request is internal and the attribute internalflag is set to True. This attribute overrides the return attribute while building SCIM response attributes when both the request is internal and the schema attribute is internal.
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: boolean
	IdcsInternal *bool `mandatory:"false" json:"idcsInternal"`

	// Trims any leading and trailing blanks from String values. Default is True.
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: boolean
	IdcsTrimStringValue *bool `mandatory:"false" json:"idcsTrimStringValue"`

	// Specifies whether this attribute can be included in a search filter
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: boolean
	IdcsSearchable *bool `mandatory:"false" json:"idcsSearchable"`

	// Specifies whether this attribute value was generated
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: boolean
	IdcsGenerated *bool `mandatory:"false" json:"idcsGenerated"`

	// Specifies whether changes to this attribute value are audited
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsSearchable: true
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsAuditable *bool `mandatory:"false" json:"idcsAuditable"`

	// Target attribute name that this attribute gets mapped to for persistence
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsTargetAttributeName *string `mandatory:"false" json:"idcsTargetAttributeName"`

	// Target index name created for this attribute for performance
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsTargetUniqueConstraintName *string `mandatory:"false" json:"idcsTargetUniqueConstraintName"`

	// Target normalized attribute name that this normalized value of attribute gets mapped to for persistence. Only set for caseExact=false & searchable attributes. Do not use by default.
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsTargetNormAttributeName *string `mandatory:"false" json:"idcsTargetNormAttributeName"`

	// Old Target attribute name from child table for CSVA attribute prior to migration. This maintains this attribute used to get mapped to for persistence
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsTargetAttributeNameToMigrateFrom *string `mandatory:"false" json:"idcsTargetAttributeNameToMigrateFrom"`

	// Specifies the mapper to use when mapping this attribute value to DataProvider-specific semantics
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsToTargetMapper *string `mandatory:"false" json:"idcsToTargetMapper"`

	// Specifies the mapper to use when mapping this attribute value from DataProvider-specific semantics
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsFromTargetMapper *string `mandatory:"false" json:"idcsFromTargetMapper"`

	// Specifies the user-friendly displayable attribute name or catalog key used for localization
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - idcsSearchable: true
	//  - type: string
	//  - uniqueness: none
	IdcsDisplayName *string `mandatory:"false" json:"idcsDisplayName"`

	// Specifies the Resource type to read from for dynamic canonical values
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsCanonicalValueSourceResourceType *string `mandatory:"false" json:"idcsCanonicalValueSourceResourceType"`

	// Filter to use when getting canonical values for this schema attribute
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsCanonicalValueSourceFilter *string `mandatory:"false" json:"idcsCanonicalValueSourceFilter"`

	// Validate payload reference value during create, replace, and update. Default is True.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsValidateReference *bool `mandatory:"false" json:"idcsValidateReference"`

	// The set of one or more sub attributes' names of a CMVA, whose values uniquely identify an instance of a CMVA
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsCompositeKey []string `mandatory:"false" json:"idcsCompositeKey"`

	// **SCIM++ Properties:**
	// - caseExact: false
	// - multiValued: false
	// - mutability: readOnly
	// - required: false
	// - idcsSearchable: true
	// - returned: default
	// - type: boolean
	// - uniqueness: none
	// Whether the CMVA attribute will be fetched or not for current resource in AbstractResourceManager update operation before calling data provider update. Default is true.
	IdcsFetchComplexAttributeValues *bool `mandatory:"false" json:"idcsFetchComplexAttributeValues"`

	// Indicates if the attribute is scim compliant, default is true
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsSearchable: true
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsScimCompliant *bool `mandatory:"false" json:"idcsScimCompliant"`

	// Specifies if the attribute can be used for mapping with external identity sources such as AD or LDAP. If isSchemaMappable: false for the schema in which this attribute is defined, then this flag is ignored
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - idcsSearchable: true
	//  - returned: default
	//  - type: boolean
	IdcsAttributeMappable *bool `mandatory:"false" json:"idcsAttributeMappable"`

	// Specifies the referenced Resource attribute
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsSearchable: true
	//  - required: false
	//  - returned: default
	//  - type: string
	IdcsRefResourceAttribute *string `mandatory:"false" json:"idcsRefResourceAttribute"`

	// Specifies whether the attribute is cacheable. True by default for all attributes. If attribute with idcsAttributeCachable = false, is present \"attributesToGet\" while executing GET/SEARCH on cacheable resource, Cache is missed and data is fetched from Data Provider.
	// **Added In:** 17.3.4
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsSearchable: true
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsAttributeCacheable *bool `mandatory:"false" json:"idcsAttributeCacheable"`
}

func (m ResourceTypeSchemaAttribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceTypeSchemaAttribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingResourceTypeSchemaAttributeTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetResourceTypeSchemaAttributeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResourceTypeSchemaAttributeMutabilityEnum(string(m.Mutability)); !ok && m.Mutability != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mutability: %s. Supported values are: %s.", m.Mutability, strings.Join(GetResourceTypeSchemaAttributeMutabilityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResourceTypeSchemaAttributeEndUserMutabilityEnum(string(m.EndUserMutability)); !ok && m.EndUserMutability != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EndUserMutability: %s. Supported values are: %s.", m.EndUserMutability, strings.Join(GetResourceTypeSchemaAttributeEndUserMutabilityEnumStringValues(), ",")))
	}
	for _, val := range m.EndUserMutabilityAllowedValues {
		if _, ok := GetMappingResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EndUserMutabilityAllowedValues: %s. Supported values are: %s.", val, strings.Join(GetResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingResourceTypeSchemaAttributeReturnedEnum(string(m.Returned)); !ok && m.Returned != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Returned: %s. Supported values are: %s.", m.Returned, strings.Join(GetResourceTypeSchemaAttributeReturnedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResourceTypeSchemaAttributeUniquenessEnum(string(m.Uniqueness)); !ok && m.Uniqueness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Uniqueness: %s. Supported values are: %s.", m.Uniqueness, strings.Join(GetResourceTypeSchemaAttributeUniquenessEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum(string(m.IdcsICFAttributeType)); !ok && m.IdcsICFAttributeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsICFAttributeType: %s. Supported values are: %s.", m.IdcsICFAttributeType, strings.Join(GetResourceTypeSchemaAttributeIdcsICFAttributeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResourceTypeSchemaAttributeIdcsSensitiveEnum(string(m.IdcsSensitive)); !ok && m.IdcsSensitive != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsSensitive: %s. Supported values are: %s.", m.IdcsSensitive, strings.Join(GetResourceTypeSchemaAttributeIdcsSensitiveEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResourceTypeSchemaAttributeTypeEnum Enum with underlying type: string
type ResourceTypeSchemaAttributeTypeEnum string

// Set of constants representing the allowable values for ResourceTypeSchemaAttributeTypeEnum
const (
	ResourceTypeSchemaAttributeTypeString    ResourceTypeSchemaAttributeTypeEnum = "string"
	ResourceTypeSchemaAttributeTypeComplex   ResourceTypeSchemaAttributeTypeEnum = "complex"
	ResourceTypeSchemaAttributeTypeBoolean   ResourceTypeSchemaAttributeTypeEnum = "boolean"
	ResourceTypeSchemaAttributeTypeDecimal   ResourceTypeSchemaAttributeTypeEnum = "decimal"
	ResourceTypeSchemaAttributeTypeInteger   ResourceTypeSchemaAttributeTypeEnum = "integer"
	ResourceTypeSchemaAttributeTypeDatetime  ResourceTypeSchemaAttributeTypeEnum = "dateTime"
	ResourceTypeSchemaAttributeTypeReference ResourceTypeSchemaAttributeTypeEnum = "reference"
	ResourceTypeSchemaAttributeTypeBinary    ResourceTypeSchemaAttributeTypeEnum = "binary"
)

var mappingResourceTypeSchemaAttributeTypeEnum = map[string]ResourceTypeSchemaAttributeTypeEnum{
	"string":    ResourceTypeSchemaAttributeTypeString,
	"complex":   ResourceTypeSchemaAttributeTypeComplex,
	"boolean":   ResourceTypeSchemaAttributeTypeBoolean,
	"decimal":   ResourceTypeSchemaAttributeTypeDecimal,
	"integer":   ResourceTypeSchemaAttributeTypeInteger,
	"dateTime":  ResourceTypeSchemaAttributeTypeDatetime,
	"reference": ResourceTypeSchemaAttributeTypeReference,
	"binary":    ResourceTypeSchemaAttributeTypeBinary,
}

var mappingResourceTypeSchemaAttributeTypeEnumLowerCase = map[string]ResourceTypeSchemaAttributeTypeEnum{
	"string":    ResourceTypeSchemaAttributeTypeString,
	"complex":   ResourceTypeSchemaAttributeTypeComplex,
	"boolean":   ResourceTypeSchemaAttributeTypeBoolean,
	"decimal":   ResourceTypeSchemaAttributeTypeDecimal,
	"integer":   ResourceTypeSchemaAttributeTypeInteger,
	"datetime":  ResourceTypeSchemaAttributeTypeDatetime,
	"reference": ResourceTypeSchemaAttributeTypeReference,
	"binary":    ResourceTypeSchemaAttributeTypeBinary,
}

// GetResourceTypeSchemaAttributeTypeEnumValues Enumerates the set of values for ResourceTypeSchemaAttributeTypeEnum
func GetResourceTypeSchemaAttributeTypeEnumValues() []ResourceTypeSchemaAttributeTypeEnum {
	values := make([]ResourceTypeSchemaAttributeTypeEnum, 0)
	for _, v := range mappingResourceTypeSchemaAttributeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceTypeSchemaAttributeTypeEnumStringValues Enumerates the set of values in String for ResourceTypeSchemaAttributeTypeEnum
func GetResourceTypeSchemaAttributeTypeEnumStringValues() []string {
	return []string{
		"string",
		"complex",
		"boolean",
		"decimal",
		"integer",
		"dateTime",
		"reference",
		"binary",
	}
}

// GetMappingResourceTypeSchemaAttributeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceTypeSchemaAttributeTypeEnum(val string) (ResourceTypeSchemaAttributeTypeEnum, bool) {
	enum, ok := mappingResourceTypeSchemaAttributeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ResourceTypeSchemaAttributeMutabilityEnum Enum with underlying type: string
type ResourceTypeSchemaAttributeMutabilityEnum string

// Set of constants representing the allowable values for ResourceTypeSchemaAttributeMutabilityEnum
const (
	ResourceTypeSchemaAttributeMutabilityReadonly  ResourceTypeSchemaAttributeMutabilityEnum = "readOnly"
	ResourceTypeSchemaAttributeMutabilityReadwrite ResourceTypeSchemaAttributeMutabilityEnum = "readWrite"
	ResourceTypeSchemaAttributeMutabilityImmutable ResourceTypeSchemaAttributeMutabilityEnum = "immutable"
	ResourceTypeSchemaAttributeMutabilityWriteonly ResourceTypeSchemaAttributeMutabilityEnum = "writeOnly"
)

var mappingResourceTypeSchemaAttributeMutabilityEnum = map[string]ResourceTypeSchemaAttributeMutabilityEnum{
	"readOnly":  ResourceTypeSchemaAttributeMutabilityReadonly,
	"readWrite": ResourceTypeSchemaAttributeMutabilityReadwrite,
	"immutable": ResourceTypeSchemaAttributeMutabilityImmutable,
	"writeOnly": ResourceTypeSchemaAttributeMutabilityWriteonly,
}

var mappingResourceTypeSchemaAttributeMutabilityEnumLowerCase = map[string]ResourceTypeSchemaAttributeMutabilityEnum{
	"readonly":  ResourceTypeSchemaAttributeMutabilityReadonly,
	"readwrite": ResourceTypeSchemaAttributeMutabilityReadwrite,
	"immutable": ResourceTypeSchemaAttributeMutabilityImmutable,
	"writeonly": ResourceTypeSchemaAttributeMutabilityWriteonly,
}

// GetResourceTypeSchemaAttributeMutabilityEnumValues Enumerates the set of values for ResourceTypeSchemaAttributeMutabilityEnum
func GetResourceTypeSchemaAttributeMutabilityEnumValues() []ResourceTypeSchemaAttributeMutabilityEnum {
	values := make([]ResourceTypeSchemaAttributeMutabilityEnum, 0)
	for _, v := range mappingResourceTypeSchemaAttributeMutabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceTypeSchemaAttributeMutabilityEnumStringValues Enumerates the set of values in String for ResourceTypeSchemaAttributeMutabilityEnum
func GetResourceTypeSchemaAttributeMutabilityEnumStringValues() []string {
	return []string{
		"readOnly",
		"readWrite",
		"immutable",
		"writeOnly",
	}
}

// GetMappingResourceTypeSchemaAttributeMutabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceTypeSchemaAttributeMutabilityEnum(val string) (ResourceTypeSchemaAttributeMutabilityEnum, bool) {
	enum, ok := mappingResourceTypeSchemaAttributeMutabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ResourceTypeSchemaAttributeEndUserMutabilityEnum Enum with underlying type: string
type ResourceTypeSchemaAttributeEndUserMutabilityEnum string

// Set of constants representing the allowable values for ResourceTypeSchemaAttributeEndUserMutabilityEnum
const (
	ResourceTypeSchemaAttributeEndUserMutabilityReadonly  ResourceTypeSchemaAttributeEndUserMutabilityEnum = "readOnly"
	ResourceTypeSchemaAttributeEndUserMutabilityReadwrite ResourceTypeSchemaAttributeEndUserMutabilityEnum = "readWrite"
	ResourceTypeSchemaAttributeEndUserMutabilityImmutable ResourceTypeSchemaAttributeEndUserMutabilityEnum = "immutable"
	ResourceTypeSchemaAttributeEndUserMutabilityHidden    ResourceTypeSchemaAttributeEndUserMutabilityEnum = "hidden"
)

var mappingResourceTypeSchemaAttributeEndUserMutabilityEnum = map[string]ResourceTypeSchemaAttributeEndUserMutabilityEnum{
	"readOnly":  ResourceTypeSchemaAttributeEndUserMutabilityReadonly,
	"readWrite": ResourceTypeSchemaAttributeEndUserMutabilityReadwrite,
	"immutable": ResourceTypeSchemaAttributeEndUserMutabilityImmutable,
	"hidden":    ResourceTypeSchemaAttributeEndUserMutabilityHidden,
}

var mappingResourceTypeSchemaAttributeEndUserMutabilityEnumLowerCase = map[string]ResourceTypeSchemaAttributeEndUserMutabilityEnum{
	"readonly":  ResourceTypeSchemaAttributeEndUserMutabilityReadonly,
	"readwrite": ResourceTypeSchemaAttributeEndUserMutabilityReadwrite,
	"immutable": ResourceTypeSchemaAttributeEndUserMutabilityImmutable,
	"hidden":    ResourceTypeSchemaAttributeEndUserMutabilityHidden,
}

// GetResourceTypeSchemaAttributeEndUserMutabilityEnumValues Enumerates the set of values for ResourceTypeSchemaAttributeEndUserMutabilityEnum
func GetResourceTypeSchemaAttributeEndUserMutabilityEnumValues() []ResourceTypeSchemaAttributeEndUserMutabilityEnum {
	values := make([]ResourceTypeSchemaAttributeEndUserMutabilityEnum, 0)
	for _, v := range mappingResourceTypeSchemaAttributeEndUserMutabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceTypeSchemaAttributeEndUserMutabilityEnumStringValues Enumerates the set of values in String for ResourceTypeSchemaAttributeEndUserMutabilityEnum
func GetResourceTypeSchemaAttributeEndUserMutabilityEnumStringValues() []string {
	return []string{
		"readOnly",
		"readWrite",
		"immutable",
		"hidden",
	}
}

// GetMappingResourceTypeSchemaAttributeEndUserMutabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceTypeSchemaAttributeEndUserMutabilityEnum(val string) (ResourceTypeSchemaAttributeEndUserMutabilityEnum, bool) {
	enum, ok := mappingResourceTypeSchemaAttributeEndUserMutabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum Enum with underlying type: string
type ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum string

// Set of constants representing the allowable values for ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum
const (
	ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesReadonly  ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum = "readOnly"
	ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesReadwrite ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum = "readWrite"
	ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesImmutable ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum = "immutable"
	ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesHidden    ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum = "hidden"
)

var mappingResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum = map[string]ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum{
	"readOnly":  ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesReadonly,
	"readWrite": ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesReadwrite,
	"immutable": ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesImmutable,
	"hidden":    ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesHidden,
}

var mappingResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnumLowerCase = map[string]ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum{
	"readonly":  ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesReadonly,
	"readwrite": ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesReadwrite,
	"immutable": ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesImmutable,
	"hidden":    ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesHidden,
}

// GetResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnumValues Enumerates the set of values for ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum
func GetResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnumValues() []ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum {
	values := make([]ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum, 0)
	for _, v := range mappingResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnumStringValues Enumerates the set of values in String for ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum
func GetResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnumStringValues() []string {
	return []string{
		"readOnly",
		"readWrite",
		"immutable",
		"hidden",
	}
}

// GetMappingResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum(val string) (ResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnum, bool) {
	enum, ok := mappingResourceTypeSchemaAttributeEndUserMutabilityAllowedValuesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ResourceTypeSchemaAttributeReturnedEnum Enum with underlying type: string
type ResourceTypeSchemaAttributeReturnedEnum string

// Set of constants representing the allowable values for ResourceTypeSchemaAttributeReturnedEnum
const (
	ResourceTypeSchemaAttributeReturnedAlways  ResourceTypeSchemaAttributeReturnedEnum = "always"
	ResourceTypeSchemaAttributeReturnedNever   ResourceTypeSchemaAttributeReturnedEnum = "never"
	ResourceTypeSchemaAttributeReturnedDefault ResourceTypeSchemaAttributeReturnedEnum = "default"
	ResourceTypeSchemaAttributeReturnedRequest ResourceTypeSchemaAttributeReturnedEnum = "request"
)

var mappingResourceTypeSchemaAttributeReturnedEnum = map[string]ResourceTypeSchemaAttributeReturnedEnum{
	"always":  ResourceTypeSchemaAttributeReturnedAlways,
	"never":   ResourceTypeSchemaAttributeReturnedNever,
	"default": ResourceTypeSchemaAttributeReturnedDefault,
	"request": ResourceTypeSchemaAttributeReturnedRequest,
}

var mappingResourceTypeSchemaAttributeReturnedEnumLowerCase = map[string]ResourceTypeSchemaAttributeReturnedEnum{
	"always":  ResourceTypeSchemaAttributeReturnedAlways,
	"never":   ResourceTypeSchemaAttributeReturnedNever,
	"default": ResourceTypeSchemaAttributeReturnedDefault,
	"request": ResourceTypeSchemaAttributeReturnedRequest,
}

// GetResourceTypeSchemaAttributeReturnedEnumValues Enumerates the set of values for ResourceTypeSchemaAttributeReturnedEnum
func GetResourceTypeSchemaAttributeReturnedEnumValues() []ResourceTypeSchemaAttributeReturnedEnum {
	values := make([]ResourceTypeSchemaAttributeReturnedEnum, 0)
	for _, v := range mappingResourceTypeSchemaAttributeReturnedEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceTypeSchemaAttributeReturnedEnumStringValues Enumerates the set of values in String for ResourceTypeSchemaAttributeReturnedEnum
func GetResourceTypeSchemaAttributeReturnedEnumStringValues() []string {
	return []string{
		"always",
		"never",
		"default",
		"request",
	}
}

// GetMappingResourceTypeSchemaAttributeReturnedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceTypeSchemaAttributeReturnedEnum(val string) (ResourceTypeSchemaAttributeReturnedEnum, bool) {
	enum, ok := mappingResourceTypeSchemaAttributeReturnedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ResourceTypeSchemaAttributeUniquenessEnum Enum with underlying type: string
type ResourceTypeSchemaAttributeUniquenessEnum string

// Set of constants representing the allowable values for ResourceTypeSchemaAttributeUniquenessEnum
const (
	ResourceTypeSchemaAttributeUniquenessNone   ResourceTypeSchemaAttributeUniquenessEnum = "none"
	ResourceTypeSchemaAttributeUniquenessServer ResourceTypeSchemaAttributeUniquenessEnum = "server"
	ResourceTypeSchemaAttributeUniquenessGlobal ResourceTypeSchemaAttributeUniquenessEnum = "global"
)

var mappingResourceTypeSchemaAttributeUniquenessEnum = map[string]ResourceTypeSchemaAttributeUniquenessEnum{
	"none":   ResourceTypeSchemaAttributeUniquenessNone,
	"server": ResourceTypeSchemaAttributeUniquenessServer,
	"global": ResourceTypeSchemaAttributeUniquenessGlobal,
}

var mappingResourceTypeSchemaAttributeUniquenessEnumLowerCase = map[string]ResourceTypeSchemaAttributeUniquenessEnum{
	"none":   ResourceTypeSchemaAttributeUniquenessNone,
	"server": ResourceTypeSchemaAttributeUniquenessServer,
	"global": ResourceTypeSchemaAttributeUniquenessGlobal,
}

// GetResourceTypeSchemaAttributeUniquenessEnumValues Enumerates the set of values for ResourceTypeSchemaAttributeUniquenessEnum
func GetResourceTypeSchemaAttributeUniquenessEnumValues() []ResourceTypeSchemaAttributeUniquenessEnum {
	values := make([]ResourceTypeSchemaAttributeUniquenessEnum, 0)
	for _, v := range mappingResourceTypeSchemaAttributeUniquenessEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceTypeSchemaAttributeUniquenessEnumStringValues Enumerates the set of values in String for ResourceTypeSchemaAttributeUniquenessEnum
func GetResourceTypeSchemaAttributeUniquenessEnumStringValues() []string {
	return []string{
		"none",
		"server",
		"global",
	}
}

// GetMappingResourceTypeSchemaAttributeUniquenessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceTypeSchemaAttributeUniquenessEnum(val string) (ResourceTypeSchemaAttributeUniquenessEnum, bool) {
	enum, ok := mappingResourceTypeSchemaAttributeUniquenessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum Enum with underlying type: string
type ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum string

// Set of constants representing the allowable values for ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum
const (
	ResourceTypeSchemaAttributeIdcsICFAttributeTypeString        ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = "string"
	ResourceTypeSchemaAttributeIdcsICFAttributeTypeLong          ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = "long"
	ResourceTypeSchemaAttributeIdcsICFAttributeTypeChar          ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = "char"
	ResourceTypeSchemaAttributeIdcsICFAttributeTypeDouble        ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = "double"
	ResourceTypeSchemaAttributeIdcsICFAttributeTypeFloat         ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = "float"
	ResourceTypeSchemaAttributeIdcsICFAttributeTypeInteger       ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = "integer"
	ResourceTypeSchemaAttributeIdcsICFAttributeTypeBoolean       ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = "boolean"
	ResourceTypeSchemaAttributeIdcsICFAttributeTypeBytes         ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = "bytes"
	ResourceTypeSchemaAttributeIdcsICFAttributeTypeBigdecimal    ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = "bigdecimal"
	ResourceTypeSchemaAttributeIdcsICFAttributeTypeBiginteger    ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = "biginteger"
	ResourceTypeSchemaAttributeIdcsICFAttributeTypeGuardedbytes  ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = "guardedbytes"
	ResourceTypeSchemaAttributeIdcsICFAttributeTypeGuardedstring ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = "guardedstring"
)

var mappingResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum = map[string]ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum{
	"string":        ResourceTypeSchemaAttributeIdcsICFAttributeTypeString,
	"long":          ResourceTypeSchemaAttributeIdcsICFAttributeTypeLong,
	"char":          ResourceTypeSchemaAttributeIdcsICFAttributeTypeChar,
	"double":        ResourceTypeSchemaAttributeIdcsICFAttributeTypeDouble,
	"float":         ResourceTypeSchemaAttributeIdcsICFAttributeTypeFloat,
	"integer":       ResourceTypeSchemaAttributeIdcsICFAttributeTypeInteger,
	"boolean":       ResourceTypeSchemaAttributeIdcsICFAttributeTypeBoolean,
	"bytes":         ResourceTypeSchemaAttributeIdcsICFAttributeTypeBytes,
	"bigdecimal":    ResourceTypeSchemaAttributeIdcsICFAttributeTypeBigdecimal,
	"biginteger":    ResourceTypeSchemaAttributeIdcsICFAttributeTypeBiginteger,
	"guardedbytes":  ResourceTypeSchemaAttributeIdcsICFAttributeTypeGuardedbytes,
	"guardedstring": ResourceTypeSchemaAttributeIdcsICFAttributeTypeGuardedstring,
}

var mappingResourceTypeSchemaAttributeIdcsICFAttributeTypeEnumLowerCase = map[string]ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum{
	"string":        ResourceTypeSchemaAttributeIdcsICFAttributeTypeString,
	"long":          ResourceTypeSchemaAttributeIdcsICFAttributeTypeLong,
	"char":          ResourceTypeSchemaAttributeIdcsICFAttributeTypeChar,
	"double":        ResourceTypeSchemaAttributeIdcsICFAttributeTypeDouble,
	"float":         ResourceTypeSchemaAttributeIdcsICFAttributeTypeFloat,
	"integer":       ResourceTypeSchemaAttributeIdcsICFAttributeTypeInteger,
	"boolean":       ResourceTypeSchemaAttributeIdcsICFAttributeTypeBoolean,
	"bytes":         ResourceTypeSchemaAttributeIdcsICFAttributeTypeBytes,
	"bigdecimal":    ResourceTypeSchemaAttributeIdcsICFAttributeTypeBigdecimal,
	"biginteger":    ResourceTypeSchemaAttributeIdcsICFAttributeTypeBiginteger,
	"guardedbytes":  ResourceTypeSchemaAttributeIdcsICFAttributeTypeGuardedbytes,
	"guardedstring": ResourceTypeSchemaAttributeIdcsICFAttributeTypeGuardedstring,
}

// GetResourceTypeSchemaAttributeIdcsICFAttributeTypeEnumValues Enumerates the set of values for ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum
func GetResourceTypeSchemaAttributeIdcsICFAttributeTypeEnumValues() []ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum {
	values := make([]ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum, 0)
	for _, v := range mappingResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceTypeSchemaAttributeIdcsICFAttributeTypeEnumStringValues Enumerates the set of values in String for ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum
func GetResourceTypeSchemaAttributeIdcsICFAttributeTypeEnumStringValues() []string {
	return []string{
		"string",
		"long",
		"char",
		"double",
		"float",
		"integer",
		"boolean",
		"bytes",
		"bigdecimal",
		"biginteger",
		"guardedbytes",
		"guardedstring",
	}
}

// GetMappingResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum(val string) (ResourceTypeSchemaAttributeIdcsICFAttributeTypeEnum, bool) {
	enum, ok := mappingResourceTypeSchemaAttributeIdcsICFAttributeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ResourceTypeSchemaAttributeIdcsSensitiveEnum Enum with underlying type: string
type ResourceTypeSchemaAttributeIdcsSensitiveEnum string

// Set of constants representing the allowable values for ResourceTypeSchemaAttributeIdcsSensitiveEnum
const (
	ResourceTypeSchemaAttributeIdcsSensitiveEncrypt  ResourceTypeSchemaAttributeIdcsSensitiveEnum = "encrypt"
	ResourceTypeSchemaAttributeIdcsSensitiveHash     ResourceTypeSchemaAttributeIdcsSensitiveEnum = "hash"
	ResourceTypeSchemaAttributeIdcsSensitiveHashSc   ResourceTypeSchemaAttributeIdcsSensitiveEnum = "hash_sc"
	ResourceTypeSchemaAttributeIdcsSensitiveChecksum ResourceTypeSchemaAttributeIdcsSensitiveEnum = "checksum"
	ResourceTypeSchemaAttributeIdcsSensitiveNone     ResourceTypeSchemaAttributeIdcsSensitiveEnum = "none"
)

var mappingResourceTypeSchemaAttributeIdcsSensitiveEnum = map[string]ResourceTypeSchemaAttributeIdcsSensitiveEnum{
	"encrypt":  ResourceTypeSchemaAttributeIdcsSensitiveEncrypt,
	"hash":     ResourceTypeSchemaAttributeIdcsSensitiveHash,
	"hash_sc":  ResourceTypeSchemaAttributeIdcsSensitiveHashSc,
	"checksum": ResourceTypeSchemaAttributeIdcsSensitiveChecksum,
	"none":     ResourceTypeSchemaAttributeIdcsSensitiveNone,
}

var mappingResourceTypeSchemaAttributeIdcsSensitiveEnumLowerCase = map[string]ResourceTypeSchemaAttributeIdcsSensitiveEnum{
	"encrypt":  ResourceTypeSchemaAttributeIdcsSensitiveEncrypt,
	"hash":     ResourceTypeSchemaAttributeIdcsSensitiveHash,
	"hash_sc":  ResourceTypeSchemaAttributeIdcsSensitiveHashSc,
	"checksum": ResourceTypeSchemaAttributeIdcsSensitiveChecksum,
	"none":     ResourceTypeSchemaAttributeIdcsSensitiveNone,
}

// GetResourceTypeSchemaAttributeIdcsSensitiveEnumValues Enumerates the set of values for ResourceTypeSchemaAttributeIdcsSensitiveEnum
func GetResourceTypeSchemaAttributeIdcsSensitiveEnumValues() []ResourceTypeSchemaAttributeIdcsSensitiveEnum {
	values := make([]ResourceTypeSchemaAttributeIdcsSensitiveEnum, 0)
	for _, v := range mappingResourceTypeSchemaAttributeIdcsSensitiveEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceTypeSchemaAttributeIdcsSensitiveEnumStringValues Enumerates the set of values in String for ResourceTypeSchemaAttributeIdcsSensitiveEnum
func GetResourceTypeSchemaAttributeIdcsSensitiveEnumStringValues() []string {
	return []string{
		"encrypt",
		"hash",
		"hash_sc",
		"checksum",
		"none",
	}
}

// GetMappingResourceTypeSchemaAttributeIdcsSensitiveEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceTypeSchemaAttributeIdcsSensitiveEnum(val string) (ResourceTypeSchemaAttributeIdcsSensitiveEnum, bool) {
	enum, ok := mappingResourceTypeSchemaAttributeIdcsSensitiveEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
