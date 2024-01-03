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

// SchemaAttributes A complex type that specifies the set of Resource attributes
type SchemaAttributes struct {

	// Attribute's name
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
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
	//  - type: string
	//  - uniqueness: none
	IdcsDisplayNameMessageId *string `mandatory:"false" json:"idcsDisplayNameMessageId"`

	// Specifies if the attributes in this schema should be hidden externally
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsRtsaHideAttribute *bool `mandatory:"false" json:"idcsRtsaHideAttribute"`

	LocalizedDisplayName *SchemaLocalizedDisplayName `mandatory:"false" json:"localizedDisplayName"`

	// The attribute's data type--for example, String
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type SchemaAttributesTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Indicates the attribute's plurality
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
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
	//  - type: string
	//  - uniqueness: none
	Description *string `mandatory:"false" json:"description"`

	// Specifies if the attribute is required
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	Required *bool `mandatory:"false" json:"required"`

	// Specifies if the attribute is required
	// **Added In:** 2305190132
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsReturnEmptyWhenNull *bool `mandatory:"false" json:"idcsReturnEmptyWhenNull"`

	// A collection of canonical values. Applicable Service Providers MUST specify the canonical types specified in the core schema specification--for example, \"work\", \"home\".
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CanonicalValues []string `mandatory:"false" json:"canonicalValues"`

	// Specifies the default value for an attribute. The value must be one from canonicalValues if defined.
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsDefaultValue *string `mandatory:"false" json:"idcsDefaultValue"`

	// A collection of Localized canonical values.
	// **SCIM++ Properties:**
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	LocalizedCanonicalValues []SchemaLocalizedCanonicalValues `mandatory:"false" json:"localizedCanonicalValues"`

	// Specifies if the String attribute is case-sensitive
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
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
	//  - type: string
	//  - uniqueness: none
	Mutability SchemaAttributesMutabilityEnum `mandatory:"false" json:"mutability,omitempty"`

	// A single keyword that indicates when an attribute and associated values are returned in response to a GET request or in response to a PUT, POST, or PATCH request
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Returned SchemaAttributesReturnedEnum `mandatory:"false" json:"returned,omitempty"`

	// A single keyword value that specifies how the Service Provider enforces uniqueness of attribute values. A server MAY reject an invalid value based on uniqueness by returning an HTTP response code of 400 (Bad Request). A client MAY enforce uniqueness on the client side to a greater degree than the Service Provider enforces. For example, a client could make a value unique while the server has the uniqueness of \"none\".
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Uniqueness SchemaAttributesUniquenessEnum `mandatory:"false" json:"uniqueness,omitempty"`

	// The attribute defining the CSV header name for import/export
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsCsvAttributeName *string `mandatory:"false" json:"idcsCsvAttributeName"`

	// Specifies the mapping between external identity source attributes and Oracle Identity Cloud Service complex attributes (e.g. email => emails[work].value)
	// **SCIM++ Properties:**
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: complex
	IdcsComplexAttributeNameMappings []SchemaIdcsComplexAttributeNameMappings `mandatory:"false" json:"idcsComplexAttributeNameMappings"`

	// Maps to ICF target attribute name
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
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
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsICFRequired *bool `mandatory:"false" json:"idcsICFRequired"`

	// Maps to ICF data type
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsICFAttributeType SchemaAttributesIdcsICFAttributeTypeEnum `mandatory:"false" json:"idcsICFAttributeType,omitempty"`

	// Csv meta data for those resource type attributes which can be imported / exported from / to csv.
	// **SCIM++ Properties:**
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	IdcsCsvAttributeNameMappings []SchemaIdcsCsvAttributeNameMappings `mandatory:"false" json:"idcsCsvAttributeNameMappings"`

	// The names of the Resource types that may be referenced--for example, User. This is only applicable for attributes that are of the \"reference\" data type.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ReferenceTypes []string `mandatory:"false" json:"referenceTypes"`

	// Indicates that the schema has been deprecated since version
	// **Deprecated Since: 19.3.3**
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	IdcsDeprecatedSinceVersion *int `mandatory:"false" json:"idcsDeprecatedSinceVersion"`

	// Indicates that the schema has been added since version
	// **Deprecated Since: 19.3.3**
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	IdcsAddedSinceVersion *int `mandatory:"false" json:"idcsAddedSinceVersion"`

	// Indicates that the schema has been deprecated since this release number
	// **Added In:** 17.3.4
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	IdcsDeprecatedSinceReleaseNumber *string `mandatory:"false" json:"idcsDeprecatedSinceReleaseNumber"`

	// Indicates that the schema has been added since this release number
	// **Added In:** 17.3.4
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	IdcsAddedSinceReleaseNumber *string `mandatory:"false" json:"idcsAddedSinceReleaseNumber"`

	// Specifies the minimum length of the attribute
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	IdcsMinLength *int `mandatory:"false" json:"idcsMinLength"`

	// Specifies the maximum length of the attribute
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	IdcsMaxLength *int `mandatory:"false" json:"idcsMaxLength"`

	// Specifies the minimum value of the integer attribute
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: integer
	IdcsMinValue *int `mandatory:"false" json:"idcsMinValue"`

	// Specifies the maximum value of the integer attribute
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: integer
	IdcsMaxValue *int `mandatory:"false" json:"idcsMaxValue"`

	// If true, specifies that the attribute can have multiple language values set for the attribute on which this is set.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readOnly
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
	//  - returned: default
	//  - type: string
	IdcsRefResourceAttributes []string `mandatory:"false" json:"idcsRefResourceAttributes"`

	// Specifies the indirectly referenced Resources
	// **SCIM++ Properties:**
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	IdcsIndirectRefResourceAttributes []string `mandatory:"false" json:"idcsIndirectRefResourceAttributes"`

	// Sequence tracking ID name for the attribute
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	IdcsAutoIncrementSeqName *string `mandatory:"false" json:"idcsAutoIncrementSeqName"`

	// Specifies whether the value of the Resource attribute is persisted
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsValuePersisted *bool `mandatory:"false" json:"idcsValuePersisted"`

	// Flag to specify if the attribute should be encrypted or hashed
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsSensitive SchemaAttributesIdcsSensitiveEnum `mandatory:"false" json:"idcsSensitive,omitempty"`

	// Specifies whether the schema attribute is for internal use only. Internal attributes are not exposed via REST. This attribute overrides mutability for create/update if the request is internal and the attribute internalflag is set to True. This attribute overrides the return attribute while building SCIM response attributes when both the request is internal and the schema attribute is internal.
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsInternal *bool `mandatory:"false" json:"idcsInternal"`

	// Trims any leading and trailing blanks from String values. Default is True.
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsTrimStringValue *bool `mandatory:"false" json:"idcsTrimStringValue"`

	// Specifies whether this attribute can be included in a search filter
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsSearchable *bool `mandatory:"false" json:"idcsSearchable"`

	// Specifies whether this attribute value was generated
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsGenerated *bool `mandatory:"false" json:"idcsGenerated"`

	// Specifies whether changes to this attribute value are audited
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
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
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsTargetAttributeName *string `mandatory:"false" json:"idcsTargetAttributeName"`

	// Contains the canonical name of the other attribute sharing the same idcsTargetAttributeName
	// **Added In:** 2209122038
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	IdcsMapsToSameTargetAttributeNameAs *string `mandatory:"false" json:"idcsMapsToSameTargetAttributeNameAs"`

	// Target normalized attribute name that this normalized value of attribute gets mapped to for persistence. Only set for caseExact=false & searchable attributes. Do not use by default.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsTargetNormAttributeName *string `mandatory:"false" json:"idcsTargetNormAttributeName"`

	// Old Target attribute name from child table for CSVA attribute prior to migration. This maintains this attribute used to get mapped to for persistence
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsTargetAttributeNameToMigrateFrom *string `mandatory:"false" json:"idcsTargetAttributeNameToMigrateFrom"`

	// Target index name created for this attribute for performance
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsTargetUniqueConstraintName *string `mandatory:"false" json:"idcsTargetUniqueConstraintName"`

	// Specifies the mapper to use when mapping this attribute value to DataProvider-specific semantics
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
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
	//  - type: string
	//  - uniqueness: none
	IdcsDisplayName *string `mandatory:"false" json:"idcsDisplayName"`

	// Specifies the Resource type to read from for dynamic canonical values
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
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

	// Specifies the Resource type ID to read from for dynamic canonical values
	// **Added In:** 17.4.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsCanonicalValueSourceResourceTypeID *string `mandatory:"false" json:"idcsCanonicalValueSourceResourceTypeID"`

	// Display name for the canonical value attribute name.
	// **Added In:** 17.4.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsCanonicalValueSourceDisplayAttrName *string `mandatory:"false" json:"idcsCanonicalValueSourceDisplayAttrName"`

	// Source key attribute for the canonical value.
	// **Added In:** 17.4.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsCanonicalValueSourceKeyAttrName *string `mandatory:"false" json:"idcsCanonicalValueSourceKeyAttrName"`

	// Type of the canonical value.
	// **Added In:** 17.4.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsCanonicalValueType SchemaAttributesIdcsCanonicalValueTypeEnum `mandatory:"false" json:"idcsCanonicalValueType,omitempty"`

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
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsCompositeKey []string `mandatory:"false" json:"idcsCompositeKey"`

	// **SCIM++ Properties:**
	// - caseExact: false
	// - multiValued: false
	// - mutability: readOnly
	// - required: false
	// - returned: default
	// - type: boolean
	// - uniqueness: none
	// Whether the CMVA attribute will be fetched or not for current resource in AbstractResourceManager update operation before calling data provider update. Default is true.
	IdcsFetchComplexAttributeValues *bool `mandatory:"false" json:"idcsFetchComplexAttributeValues"`

	// Indicates if the attribute is scim compliant, default is true
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsScimCompliant *bool `mandatory:"false" json:"idcsScimCompliant"`

	// Specifies if the attribute can be used for mapping with external identity sources such as AD or LDAP. If isSchemaMappable: false for the schema in which this attribute is defined, then this flag is ignored
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsAttributeMappable *bool `mandatory:"false" json:"idcsAttributeMappable"`

	// If true, ARM should not validate the value of the attribute since it will be converted/migrated to another attribute internally by the manager which will build valid post, put, or patch payload, depending on the client requested operation
	// **Added In:** 18.2.2
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsValuePersistedInOtherAttribute *bool `mandatory:"false" json:"idcsValuePersistedInOtherAttribute"`

	// Specifies whether the attribute is PII (personal information). False by default for all attributes. If attribute with idcsPii = true, it's value must be obfuscated before it's written to the Oracle Identity Cloud Service system logs.
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsPii *bool `mandatory:"false" json:"idcsPii"`

	// Specifies whether the attribute should be excluded from the BulkApi patch generated by gradle task for upgrading OOTB resources.
	// **Added In:** 2104150946
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: never
	//  - type: boolean
	IdcsExcludeFromUpgradePatch *bool `mandatory:"false" json:"idcsExcludeFromUpgradePatch"`

	// A list specifying the contained attributes
	// **SCIM++ Properties:**
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	SubAttributes []SchemaSubAttributes `mandatory:"false" json:"subAttributes"`

	// Specifies the referenced Resource attribute
	// **Deprecated Since: 17.3.4**
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	IdcsRefResourceAttribute *string `mandatory:"false" json:"idcsRefResourceAttribute"`

	// Specifies whether the attribute is cacheable. True by default for all attributes. If attribute with idcsAttributeCachable = false, is present \"attributesToGet\" while executing GET/SEARCH on cacheable resource, Cache is missed and data is fetched from Data Provider.
	// **Added In:** 17.3.4
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsAttributeCacheable *bool `mandatory:"false" json:"idcsAttributeCacheable"`

	// Metadata used by Oracle Identity Cloud Service UI to sequence the attributes displayed on the Account Form.
	// **Added In:** 17.4.2
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	IdcsuiOrder *int `mandatory:"false" json:"idcsuiOrder"`

	// Metadata used by Oracle Identity Cloud Service UI to validate the attribute values using regular expressions.
	// **Added In:** 17.4.2
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	IdcsuiRegexp *string `mandatory:"false" json:"idcsuiRegexp"`

	// Metadata used by Oracle Identity Cloud Service UI to decide whether the attribute must be displayed on the Account Form.
	// **Added In:** 17.4.2
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsuiVisible *bool `mandatory:"false" json:"idcsuiVisible"`

	// Metadata used by Oracle Identity Cloud Service UI to render a widget for this attribute on the Account Form.
	// **Added In:** 17.4.2
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	IdcsuiWidget SchemaAttributesIdcsuiWidgetEnum `mandatory:"false" json:"idcsuiWidget,omitempty"`

	// The list of features that require this attribute
	// **Deprecated Since: 19.1.6**
	// **SCIM++ Properties:**
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	IdcsFeatures []SchemaAttributesIdcsFeaturesEnum `mandatory:"false" json:"idcsFeatures,omitempty"`

	// A subset of \"canonicalValues\" that are not supported when the \"optionalPii\" feature is disabled in GlobalConfig.
	// **Deprecated Since: 19.1.6**
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsOptionalPiiCanonicalValues []string `mandatory:"false" json:"idcsOptionalPiiCanonicalValues"`

	// Specifies if the value of the attribute should be sanitized using OWASP library for HTML content
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsSanitize *bool `mandatory:"false" json:"idcsSanitize"`

	// Specifies whether the attribute from resource schema should override from common schema with the same name.
	// **Added In:** 2102181953
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsOverrideCommonAttribute *bool `mandatory:"false" json:"idcsOverrideCommonAttribute"`

	// Specifies whether the readOnly and immutable reference attributes should be ignored when forceDelete=true.
	// **Added In:** 2104220644
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsIgnoreReadOnlyAndImmutableRefAttrsDuringForceDelete *bool `mandatory:"false" json:"idcsIgnoreReadOnlyAndImmutableRefAttrsDuringForceDelete"`

	// Set this attribute to True if the resource is eligibal for update while system is in readonly mode.
	// **Added In:** 2106170416
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsAllowUpdatesInReadOnlyMode *bool `mandatory:"false" json:"idcsAllowUpdatesInReadOnlyMode"`

	// Set this attribute to True if the pagination is required on an attribute.
	// **Added In:** 2202230830
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsPaginateResponse *bool `mandatory:"false" json:"idcsPaginateResponse"`

	// If true, indicates that the attribute value must be written to the home region for access flows initiated from a replica region.
	// **Added In:** 2209220956
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsRequiresWriteForAccessFlows *bool `mandatory:"false" json:"idcsRequiresWriteForAccessFlows"`

	// If true, indicates that the attribute value must be written to the home region and requires immediate read-after-write consistency for access flows initiated from a replica region.
	// **Added In:** 2209220956
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: never
	//  - type: boolean
	IdcsRequiresImmediateReadAfterWriteForAccessFlows *bool `mandatory:"false" json:"idcsRequiresImmediateReadAfterWriteForAccessFlows"`
}

func (m SchemaAttributes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchemaAttributes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSchemaAttributesTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSchemaAttributesTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchemaAttributesMutabilityEnum(string(m.Mutability)); !ok && m.Mutability != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mutability: %s. Supported values are: %s.", m.Mutability, strings.Join(GetSchemaAttributesMutabilityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchemaAttributesReturnedEnum(string(m.Returned)); !ok && m.Returned != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Returned: %s. Supported values are: %s.", m.Returned, strings.Join(GetSchemaAttributesReturnedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchemaAttributesUniquenessEnum(string(m.Uniqueness)); !ok && m.Uniqueness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Uniqueness: %s. Supported values are: %s.", m.Uniqueness, strings.Join(GetSchemaAttributesUniquenessEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchemaAttributesIdcsICFAttributeTypeEnum(string(m.IdcsICFAttributeType)); !ok && m.IdcsICFAttributeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsICFAttributeType: %s. Supported values are: %s.", m.IdcsICFAttributeType, strings.Join(GetSchemaAttributesIdcsICFAttributeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchemaAttributesIdcsSensitiveEnum(string(m.IdcsSensitive)); !ok && m.IdcsSensitive != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsSensitive: %s. Supported values are: %s.", m.IdcsSensitive, strings.Join(GetSchemaAttributesIdcsSensitiveEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchemaAttributesIdcsCanonicalValueTypeEnum(string(m.IdcsCanonicalValueType)); !ok && m.IdcsCanonicalValueType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsCanonicalValueType: %s. Supported values are: %s.", m.IdcsCanonicalValueType, strings.Join(GetSchemaAttributesIdcsCanonicalValueTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchemaAttributesIdcsuiWidgetEnum(string(m.IdcsuiWidget)); !ok && m.IdcsuiWidget != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsuiWidget: %s. Supported values are: %s.", m.IdcsuiWidget, strings.Join(GetSchemaAttributesIdcsuiWidgetEnumStringValues(), ",")))
	}
	for _, val := range m.IdcsFeatures {
		if _, ok := GetMappingSchemaAttributesIdcsFeaturesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsFeatures: %s. Supported values are: %s.", val, strings.Join(GetSchemaAttributesIdcsFeaturesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SchemaAttributesTypeEnum Enum with underlying type: string
type SchemaAttributesTypeEnum string

// Set of constants representing the allowable values for SchemaAttributesTypeEnum
const (
	SchemaAttributesTypeString    SchemaAttributesTypeEnum = "string"
	SchemaAttributesTypeComplex   SchemaAttributesTypeEnum = "complex"
	SchemaAttributesTypeBoolean   SchemaAttributesTypeEnum = "boolean"
	SchemaAttributesTypeDecimal   SchemaAttributesTypeEnum = "decimal"
	SchemaAttributesTypeInteger   SchemaAttributesTypeEnum = "integer"
	SchemaAttributesTypeDatetime  SchemaAttributesTypeEnum = "dateTime"
	SchemaAttributesTypeReference SchemaAttributesTypeEnum = "reference"
	SchemaAttributesTypeBinary    SchemaAttributesTypeEnum = "binary"
)

var mappingSchemaAttributesTypeEnum = map[string]SchemaAttributesTypeEnum{
	"string":    SchemaAttributesTypeString,
	"complex":   SchemaAttributesTypeComplex,
	"boolean":   SchemaAttributesTypeBoolean,
	"decimal":   SchemaAttributesTypeDecimal,
	"integer":   SchemaAttributesTypeInteger,
	"dateTime":  SchemaAttributesTypeDatetime,
	"reference": SchemaAttributesTypeReference,
	"binary":    SchemaAttributesTypeBinary,
}

var mappingSchemaAttributesTypeEnumLowerCase = map[string]SchemaAttributesTypeEnum{
	"string":    SchemaAttributesTypeString,
	"complex":   SchemaAttributesTypeComplex,
	"boolean":   SchemaAttributesTypeBoolean,
	"decimal":   SchemaAttributesTypeDecimal,
	"integer":   SchemaAttributesTypeInteger,
	"datetime":  SchemaAttributesTypeDatetime,
	"reference": SchemaAttributesTypeReference,
	"binary":    SchemaAttributesTypeBinary,
}

// GetSchemaAttributesTypeEnumValues Enumerates the set of values for SchemaAttributesTypeEnum
func GetSchemaAttributesTypeEnumValues() []SchemaAttributesTypeEnum {
	values := make([]SchemaAttributesTypeEnum, 0)
	for _, v := range mappingSchemaAttributesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaAttributesTypeEnumStringValues Enumerates the set of values in String for SchemaAttributesTypeEnum
func GetSchemaAttributesTypeEnumStringValues() []string {
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

// GetMappingSchemaAttributesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaAttributesTypeEnum(val string) (SchemaAttributesTypeEnum, bool) {
	enum, ok := mappingSchemaAttributesTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaAttributesMutabilityEnum Enum with underlying type: string
type SchemaAttributesMutabilityEnum string

// Set of constants representing the allowable values for SchemaAttributesMutabilityEnum
const (
	SchemaAttributesMutabilityReadonly  SchemaAttributesMutabilityEnum = "readOnly"
	SchemaAttributesMutabilityReadwrite SchemaAttributesMutabilityEnum = "readWrite"
	SchemaAttributesMutabilityImmutable SchemaAttributesMutabilityEnum = "immutable"
	SchemaAttributesMutabilityWriteonly SchemaAttributesMutabilityEnum = "writeOnly"
)

var mappingSchemaAttributesMutabilityEnum = map[string]SchemaAttributesMutabilityEnum{
	"readOnly":  SchemaAttributesMutabilityReadonly,
	"readWrite": SchemaAttributesMutabilityReadwrite,
	"immutable": SchemaAttributesMutabilityImmutable,
	"writeOnly": SchemaAttributesMutabilityWriteonly,
}

var mappingSchemaAttributesMutabilityEnumLowerCase = map[string]SchemaAttributesMutabilityEnum{
	"readonly":  SchemaAttributesMutabilityReadonly,
	"readwrite": SchemaAttributesMutabilityReadwrite,
	"immutable": SchemaAttributesMutabilityImmutable,
	"writeonly": SchemaAttributesMutabilityWriteonly,
}

// GetSchemaAttributesMutabilityEnumValues Enumerates the set of values for SchemaAttributesMutabilityEnum
func GetSchemaAttributesMutabilityEnumValues() []SchemaAttributesMutabilityEnum {
	values := make([]SchemaAttributesMutabilityEnum, 0)
	for _, v := range mappingSchemaAttributesMutabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaAttributesMutabilityEnumStringValues Enumerates the set of values in String for SchemaAttributesMutabilityEnum
func GetSchemaAttributesMutabilityEnumStringValues() []string {
	return []string{
		"readOnly",
		"readWrite",
		"immutable",
		"writeOnly",
	}
}

// GetMappingSchemaAttributesMutabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaAttributesMutabilityEnum(val string) (SchemaAttributesMutabilityEnum, bool) {
	enum, ok := mappingSchemaAttributesMutabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaAttributesReturnedEnum Enum with underlying type: string
type SchemaAttributesReturnedEnum string

// Set of constants representing the allowable values for SchemaAttributesReturnedEnum
const (
	SchemaAttributesReturnedAlways  SchemaAttributesReturnedEnum = "always"
	SchemaAttributesReturnedNever   SchemaAttributesReturnedEnum = "never"
	SchemaAttributesReturnedDefault SchemaAttributesReturnedEnum = "default"
	SchemaAttributesReturnedRequest SchemaAttributesReturnedEnum = "request"
)

var mappingSchemaAttributesReturnedEnum = map[string]SchemaAttributesReturnedEnum{
	"always":  SchemaAttributesReturnedAlways,
	"never":   SchemaAttributesReturnedNever,
	"default": SchemaAttributesReturnedDefault,
	"request": SchemaAttributesReturnedRequest,
}

var mappingSchemaAttributesReturnedEnumLowerCase = map[string]SchemaAttributesReturnedEnum{
	"always":  SchemaAttributesReturnedAlways,
	"never":   SchemaAttributesReturnedNever,
	"default": SchemaAttributesReturnedDefault,
	"request": SchemaAttributesReturnedRequest,
}

// GetSchemaAttributesReturnedEnumValues Enumerates the set of values for SchemaAttributesReturnedEnum
func GetSchemaAttributesReturnedEnumValues() []SchemaAttributesReturnedEnum {
	values := make([]SchemaAttributesReturnedEnum, 0)
	for _, v := range mappingSchemaAttributesReturnedEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaAttributesReturnedEnumStringValues Enumerates the set of values in String for SchemaAttributesReturnedEnum
func GetSchemaAttributesReturnedEnumStringValues() []string {
	return []string{
		"always",
		"never",
		"default",
		"request",
	}
}

// GetMappingSchemaAttributesReturnedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaAttributesReturnedEnum(val string) (SchemaAttributesReturnedEnum, bool) {
	enum, ok := mappingSchemaAttributesReturnedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaAttributesUniquenessEnum Enum with underlying type: string
type SchemaAttributesUniquenessEnum string

// Set of constants representing the allowable values for SchemaAttributesUniquenessEnum
const (
	SchemaAttributesUniquenessNone   SchemaAttributesUniquenessEnum = "none"
	SchemaAttributesUniquenessServer SchemaAttributesUniquenessEnum = "server"
	SchemaAttributesUniquenessGlobal SchemaAttributesUniquenessEnum = "global"
)

var mappingSchemaAttributesUniquenessEnum = map[string]SchemaAttributesUniquenessEnum{
	"none":   SchemaAttributesUniquenessNone,
	"server": SchemaAttributesUniquenessServer,
	"global": SchemaAttributesUniquenessGlobal,
}

var mappingSchemaAttributesUniquenessEnumLowerCase = map[string]SchemaAttributesUniquenessEnum{
	"none":   SchemaAttributesUniquenessNone,
	"server": SchemaAttributesUniquenessServer,
	"global": SchemaAttributesUniquenessGlobal,
}

// GetSchemaAttributesUniquenessEnumValues Enumerates the set of values for SchemaAttributesUniquenessEnum
func GetSchemaAttributesUniquenessEnumValues() []SchemaAttributesUniquenessEnum {
	values := make([]SchemaAttributesUniquenessEnum, 0)
	for _, v := range mappingSchemaAttributesUniquenessEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaAttributesUniquenessEnumStringValues Enumerates the set of values in String for SchemaAttributesUniquenessEnum
func GetSchemaAttributesUniquenessEnumStringValues() []string {
	return []string{
		"none",
		"server",
		"global",
	}
}

// GetMappingSchemaAttributesUniquenessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaAttributesUniquenessEnum(val string) (SchemaAttributesUniquenessEnum, bool) {
	enum, ok := mappingSchemaAttributesUniquenessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaAttributesIdcsICFAttributeTypeEnum Enum with underlying type: string
type SchemaAttributesIdcsICFAttributeTypeEnum string

// Set of constants representing the allowable values for SchemaAttributesIdcsICFAttributeTypeEnum
const (
	SchemaAttributesIdcsICFAttributeTypeString        SchemaAttributesIdcsICFAttributeTypeEnum = "string"
	SchemaAttributesIdcsICFAttributeTypeLong          SchemaAttributesIdcsICFAttributeTypeEnum = "long"
	SchemaAttributesIdcsICFAttributeTypeChar          SchemaAttributesIdcsICFAttributeTypeEnum = "char"
	SchemaAttributesIdcsICFAttributeTypeDouble        SchemaAttributesIdcsICFAttributeTypeEnum = "double"
	SchemaAttributesIdcsICFAttributeTypeFloat         SchemaAttributesIdcsICFAttributeTypeEnum = "float"
	SchemaAttributesIdcsICFAttributeTypeInteger       SchemaAttributesIdcsICFAttributeTypeEnum = "integer"
	SchemaAttributesIdcsICFAttributeTypeBoolean       SchemaAttributesIdcsICFAttributeTypeEnum = "boolean"
	SchemaAttributesIdcsICFAttributeTypeBytes         SchemaAttributesIdcsICFAttributeTypeEnum = "bytes"
	SchemaAttributesIdcsICFAttributeTypeBigdecimal    SchemaAttributesIdcsICFAttributeTypeEnum = "bigdecimal"
	SchemaAttributesIdcsICFAttributeTypeBiginteger    SchemaAttributesIdcsICFAttributeTypeEnum = "biginteger"
	SchemaAttributesIdcsICFAttributeTypeGuardedbytes  SchemaAttributesIdcsICFAttributeTypeEnum = "guardedbytes"
	SchemaAttributesIdcsICFAttributeTypeGuardedstring SchemaAttributesIdcsICFAttributeTypeEnum = "guardedstring"
)

var mappingSchemaAttributesIdcsICFAttributeTypeEnum = map[string]SchemaAttributesIdcsICFAttributeTypeEnum{
	"string":        SchemaAttributesIdcsICFAttributeTypeString,
	"long":          SchemaAttributesIdcsICFAttributeTypeLong,
	"char":          SchemaAttributesIdcsICFAttributeTypeChar,
	"double":        SchemaAttributesIdcsICFAttributeTypeDouble,
	"float":         SchemaAttributesIdcsICFAttributeTypeFloat,
	"integer":       SchemaAttributesIdcsICFAttributeTypeInteger,
	"boolean":       SchemaAttributesIdcsICFAttributeTypeBoolean,
	"bytes":         SchemaAttributesIdcsICFAttributeTypeBytes,
	"bigdecimal":    SchemaAttributesIdcsICFAttributeTypeBigdecimal,
	"biginteger":    SchemaAttributesIdcsICFAttributeTypeBiginteger,
	"guardedbytes":  SchemaAttributesIdcsICFAttributeTypeGuardedbytes,
	"guardedstring": SchemaAttributesIdcsICFAttributeTypeGuardedstring,
}

var mappingSchemaAttributesIdcsICFAttributeTypeEnumLowerCase = map[string]SchemaAttributesIdcsICFAttributeTypeEnum{
	"string":        SchemaAttributesIdcsICFAttributeTypeString,
	"long":          SchemaAttributesIdcsICFAttributeTypeLong,
	"char":          SchemaAttributesIdcsICFAttributeTypeChar,
	"double":        SchemaAttributesIdcsICFAttributeTypeDouble,
	"float":         SchemaAttributesIdcsICFAttributeTypeFloat,
	"integer":       SchemaAttributesIdcsICFAttributeTypeInteger,
	"boolean":       SchemaAttributesIdcsICFAttributeTypeBoolean,
	"bytes":         SchemaAttributesIdcsICFAttributeTypeBytes,
	"bigdecimal":    SchemaAttributesIdcsICFAttributeTypeBigdecimal,
	"biginteger":    SchemaAttributesIdcsICFAttributeTypeBiginteger,
	"guardedbytes":  SchemaAttributesIdcsICFAttributeTypeGuardedbytes,
	"guardedstring": SchemaAttributesIdcsICFAttributeTypeGuardedstring,
}

// GetSchemaAttributesIdcsICFAttributeTypeEnumValues Enumerates the set of values for SchemaAttributesIdcsICFAttributeTypeEnum
func GetSchemaAttributesIdcsICFAttributeTypeEnumValues() []SchemaAttributesIdcsICFAttributeTypeEnum {
	values := make([]SchemaAttributesIdcsICFAttributeTypeEnum, 0)
	for _, v := range mappingSchemaAttributesIdcsICFAttributeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaAttributesIdcsICFAttributeTypeEnumStringValues Enumerates the set of values in String for SchemaAttributesIdcsICFAttributeTypeEnum
func GetSchemaAttributesIdcsICFAttributeTypeEnumStringValues() []string {
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

// GetMappingSchemaAttributesIdcsICFAttributeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaAttributesIdcsICFAttributeTypeEnum(val string) (SchemaAttributesIdcsICFAttributeTypeEnum, bool) {
	enum, ok := mappingSchemaAttributesIdcsICFAttributeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaAttributesIdcsSensitiveEnum Enum with underlying type: string
type SchemaAttributesIdcsSensitiveEnum string

// Set of constants representing the allowable values for SchemaAttributesIdcsSensitiveEnum
const (
	SchemaAttributesIdcsSensitiveEncrypt  SchemaAttributesIdcsSensitiveEnum = "encrypt"
	SchemaAttributesIdcsSensitiveHash     SchemaAttributesIdcsSensitiveEnum = "hash"
	SchemaAttributesIdcsSensitiveHashSc   SchemaAttributesIdcsSensitiveEnum = "hash_sc"
	SchemaAttributesIdcsSensitiveChecksum SchemaAttributesIdcsSensitiveEnum = "checksum"
	SchemaAttributesIdcsSensitiveNone     SchemaAttributesIdcsSensitiveEnum = "none"
)

var mappingSchemaAttributesIdcsSensitiveEnum = map[string]SchemaAttributesIdcsSensitiveEnum{
	"encrypt":  SchemaAttributesIdcsSensitiveEncrypt,
	"hash":     SchemaAttributesIdcsSensitiveHash,
	"hash_sc":  SchemaAttributesIdcsSensitiveHashSc,
	"checksum": SchemaAttributesIdcsSensitiveChecksum,
	"none":     SchemaAttributesIdcsSensitiveNone,
}

var mappingSchemaAttributesIdcsSensitiveEnumLowerCase = map[string]SchemaAttributesIdcsSensitiveEnum{
	"encrypt":  SchemaAttributesIdcsSensitiveEncrypt,
	"hash":     SchemaAttributesIdcsSensitiveHash,
	"hash_sc":  SchemaAttributesIdcsSensitiveHashSc,
	"checksum": SchemaAttributesIdcsSensitiveChecksum,
	"none":     SchemaAttributesIdcsSensitiveNone,
}

// GetSchemaAttributesIdcsSensitiveEnumValues Enumerates the set of values for SchemaAttributesIdcsSensitiveEnum
func GetSchemaAttributesIdcsSensitiveEnumValues() []SchemaAttributesIdcsSensitiveEnum {
	values := make([]SchemaAttributesIdcsSensitiveEnum, 0)
	for _, v := range mappingSchemaAttributesIdcsSensitiveEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaAttributesIdcsSensitiveEnumStringValues Enumerates the set of values in String for SchemaAttributesIdcsSensitiveEnum
func GetSchemaAttributesIdcsSensitiveEnumStringValues() []string {
	return []string{
		"encrypt",
		"hash",
		"hash_sc",
		"checksum",
		"none",
	}
}

// GetMappingSchemaAttributesIdcsSensitiveEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaAttributesIdcsSensitiveEnum(val string) (SchemaAttributesIdcsSensitiveEnum, bool) {
	enum, ok := mappingSchemaAttributesIdcsSensitiveEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaAttributesIdcsCanonicalValueTypeEnum Enum with underlying type: string
type SchemaAttributesIdcsCanonicalValueTypeEnum string

// Set of constants representing the allowable values for SchemaAttributesIdcsCanonicalValueTypeEnum
const (
	SchemaAttributesIdcsCanonicalValueTypeDynamic SchemaAttributesIdcsCanonicalValueTypeEnum = "dynamic"
	SchemaAttributesIdcsCanonicalValueTypeStatic  SchemaAttributesIdcsCanonicalValueTypeEnum = "static"
)

var mappingSchemaAttributesIdcsCanonicalValueTypeEnum = map[string]SchemaAttributesIdcsCanonicalValueTypeEnum{
	"dynamic": SchemaAttributesIdcsCanonicalValueTypeDynamic,
	"static":  SchemaAttributesIdcsCanonicalValueTypeStatic,
}

var mappingSchemaAttributesIdcsCanonicalValueTypeEnumLowerCase = map[string]SchemaAttributesIdcsCanonicalValueTypeEnum{
	"dynamic": SchemaAttributesIdcsCanonicalValueTypeDynamic,
	"static":  SchemaAttributesIdcsCanonicalValueTypeStatic,
}

// GetSchemaAttributesIdcsCanonicalValueTypeEnumValues Enumerates the set of values for SchemaAttributesIdcsCanonicalValueTypeEnum
func GetSchemaAttributesIdcsCanonicalValueTypeEnumValues() []SchemaAttributesIdcsCanonicalValueTypeEnum {
	values := make([]SchemaAttributesIdcsCanonicalValueTypeEnum, 0)
	for _, v := range mappingSchemaAttributesIdcsCanonicalValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaAttributesIdcsCanonicalValueTypeEnumStringValues Enumerates the set of values in String for SchemaAttributesIdcsCanonicalValueTypeEnum
func GetSchemaAttributesIdcsCanonicalValueTypeEnumStringValues() []string {
	return []string{
		"dynamic",
		"static",
	}
}

// GetMappingSchemaAttributesIdcsCanonicalValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaAttributesIdcsCanonicalValueTypeEnum(val string) (SchemaAttributesIdcsCanonicalValueTypeEnum, bool) {
	enum, ok := mappingSchemaAttributesIdcsCanonicalValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaAttributesIdcsuiWidgetEnum Enum with underlying type: string
type SchemaAttributesIdcsuiWidgetEnum string

// Set of constants representing the allowable values for SchemaAttributesIdcsuiWidgetEnum
const (
	SchemaAttributesIdcsuiWidgetInputtext SchemaAttributesIdcsuiWidgetEnum = "inputtext"
	SchemaAttributesIdcsuiWidgetCheckbox  SchemaAttributesIdcsuiWidgetEnum = "checkbox"
	SchemaAttributesIdcsuiWidgetTextarea  SchemaAttributesIdcsuiWidgetEnum = "textarea"
	SchemaAttributesIdcsuiWidgetCombobox  SchemaAttributesIdcsuiWidgetEnum = "combobox"
)

var mappingSchemaAttributesIdcsuiWidgetEnum = map[string]SchemaAttributesIdcsuiWidgetEnum{
	"inputtext": SchemaAttributesIdcsuiWidgetInputtext,
	"checkbox":  SchemaAttributesIdcsuiWidgetCheckbox,
	"textarea":  SchemaAttributesIdcsuiWidgetTextarea,
	"combobox":  SchemaAttributesIdcsuiWidgetCombobox,
}

var mappingSchemaAttributesIdcsuiWidgetEnumLowerCase = map[string]SchemaAttributesIdcsuiWidgetEnum{
	"inputtext": SchemaAttributesIdcsuiWidgetInputtext,
	"checkbox":  SchemaAttributesIdcsuiWidgetCheckbox,
	"textarea":  SchemaAttributesIdcsuiWidgetTextarea,
	"combobox":  SchemaAttributesIdcsuiWidgetCombobox,
}

// GetSchemaAttributesIdcsuiWidgetEnumValues Enumerates the set of values for SchemaAttributesIdcsuiWidgetEnum
func GetSchemaAttributesIdcsuiWidgetEnumValues() []SchemaAttributesIdcsuiWidgetEnum {
	values := make([]SchemaAttributesIdcsuiWidgetEnum, 0)
	for _, v := range mappingSchemaAttributesIdcsuiWidgetEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaAttributesIdcsuiWidgetEnumStringValues Enumerates the set of values in String for SchemaAttributesIdcsuiWidgetEnum
func GetSchemaAttributesIdcsuiWidgetEnumStringValues() []string {
	return []string{
		"inputtext",
		"checkbox",
		"textarea",
		"combobox",
	}
}

// GetMappingSchemaAttributesIdcsuiWidgetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaAttributesIdcsuiWidgetEnum(val string) (SchemaAttributesIdcsuiWidgetEnum, bool) {
	enum, ok := mappingSchemaAttributesIdcsuiWidgetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaAttributesIdcsFeaturesEnum Enum with underlying type: string
type SchemaAttributesIdcsFeaturesEnum string

// Set of constants representing the allowable values for SchemaAttributesIdcsFeaturesEnum
const (
	SchemaAttributesIdcsFeaturesOptionalpii         SchemaAttributesIdcsFeaturesEnum = "optionalPii"
	SchemaAttributesIdcsFeaturesMfa                 SchemaAttributesIdcsFeaturesEnum = "mfa"
	SchemaAttributesIdcsFeaturesSocial              SchemaAttributesIdcsFeaturesEnum = "social"
	SchemaAttributesIdcsFeaturesSchemacustomization SchemaAttributesIdcsFeaturesEnum = "schemaCustomization"
)

var mappingSchemaAttributesIdcsFeaturesEnum = map[string]SchemaAttributesIdcsFeaturesEnum{
	"optionalPii":         SchemaAttributesIdcsFeaturesOptionalpii,
	"mfa":                 SchemaAttributesIdcsFeaturesMfa,
	"social":              SchemaAttributesIdcsFeaturesSocial,
	"schemaCustomization": SchemaAttributesIdcsFeaturesSchemacustomization,
}

var mappingSchemaAttributesIdcsFeaturesEnumLowerCase = map[string]SchemaAttributesIdcsFeaturesEnum{
	"optionalpii":         SchemaAttributesIdcsFeaturesOptionalpii,
	"mfa":                 SchemaAttributesIdcsFeaturesMfa,
	"social":              SchemaAttributesIdcsFeaturesSocial,
	"schemacustomization": SchemaAttributesIdcsFeaturesSchemacustomization,
}

// GetSchemaAttributesIdcsFeaturesEnumValues Enumerates the set of values for SchemaAttributesIdcsFeaturesEnum
func GetSchemaAttributesIdcsFeaturesEnumValues() []SchemaAttributesIdcsFeaturesEnum {
	values := make([]SchemaAttributesIdcsFeaturesEnum, 0)
	for _, v := range mappingSchemaAttributesIdcsFeaturesEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaAttributesIdcsFeaturesEnumStringValues Enumerates the set of values in String for SchemaAttributesIdcsFeaturesEnum
func GetSchemaAttributesIdcsFeaturesEnumStringValues() []string {
	return []string{
		"optionalPii",
		"mfa",
		"social",
		"schemaCustomization",
	}
}

// GetMappingSchemaAttributesIdcsFeaturesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaAttributesIdcsFeaturesEnum(val string) (SchemaAttributesIdcsFeaturesEnum, bool) {
	enum, ok := mappingSchemaAttributesIdcsFeaturesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
