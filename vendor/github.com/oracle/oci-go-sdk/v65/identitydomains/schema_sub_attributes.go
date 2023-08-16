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

// SchemaSubAttributes A list specifying the contained attributes
type SchemaSubAttributes struct {

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

	// If true, indicates that the attribute value must be written to the home region and requires immediate read-after-write consistency for access flows initiated from a replica region.
	// **Added In:** 2209220956
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: never
	//  - type: boolean
	IdcsRequiresImmediateReadAfterWriteForAccessFlows *bool `mandatory:"false" json:"idcsRequiresImmediateReadAfterWriteForAccessFlows"`

	// If true, indicates that the attribute value must be written to the home region for access flows initiated from a replica region.
	// **Added In:** 2205120021
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsRequiresWriteForAccessFlows *bool `mandatory:"false" json:"idcsRequiresWriteForAccessFlows"`

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

	// Specifies if the attributes in this schema can be displayed externally
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

	// Specifies whether the schema attribute is for internal use only. Internal attributes are not exposed via REST. This attribute overrides mutability for create/update if the request is internal and the attribute internal flag is set to True. This attribute overrides the return attribute while building SCIM response attributes when both the request is internal and the schema attribute is internal.
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsInternal *bool `mandatory:"false" json:"idcsInternal"`

	// Attribute's data type--for example, String
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type SchemaSubAttributesTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Indicates the attribute's plurality
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	MultiValued *bool `mandatory:"false" json:"multiValued"`

	// Attribute's human-readable description
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
	Mutability SchemaSubAttributesMutabilityEnum `mandatory:"false" json:"mutability,omitempty"`

	// A single keyword that indicates when an attribute and associated values are returned in response to a GET request or in response to a PUT, POST, or PATCH request
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Returned SchemaSubAttributesReturnedEnum `mandatory:"false" json:"returned,omitempty"`

	// The attribute that defines the CSV header name for import/export
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsCsvAttributeName *string `mandatory:"false" json:"idcsCsvAttributeName"`

	// A single keyword value that specifies how the Service Provider enforces uniqueness of attribute values. A server MAY reject an invalid value based on uniqueness by returning an HTTP response code of 400 (Bad Request). A client MAY enforce uniqueness on the client side to a greater degree than the Service Provider enforces. For example, a client could make a value unique while the server has a uniqueness of \"none\".
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Uniqueness SchemaSubAttributesUniquenessEnum `mandatory:"false" json:"uniqueness,omitempty"`

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

	// Specifies the minimum length of this attribute
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	IdcsMinLength *int `mandatory:"false" json:"idcsMinLength"`

	// Specifies the maximum length of this attribute
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

	// If true, specifies that the sub attribute value can be set to true on one and only one instance of the CMVA.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsOnlyOneValueCanBeTrue *bool `mandatory:"false" json:"idcsOnlyOneValueCanBeTrue"`

	// **SCIM++ Properties:**
	// - caseExact: true
	// - multiValued: false
	// - mutability: readOnly
	// - required: false
	// - returned: default
	// - type: integer
	// - uniqueness: none
	// Specify a limit on the number of attribute-values that any caller will receive when requesting a CMVA attribute. If the no of CMVA instances exceeds the limit then Oracle Identity Cloud Service will throw exception. Users can choose to refine the filter on CMVA attribute.
	IdcsMaxValuesReturned *int `mandatory:"false" json:"idcsMaxValuesReturned"`

	// If true, ARM will ensure atleast one of the instances of CMVA has the attribute value set to true.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IdcsOneValueMustBeTrue *bool `mandatory:"false" json:"idcsOneValueMustBeTrue"`

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

	// Specifies whether the sub-attribute of the Resource attribute is persisted
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsValuePersisted *bool `mandatory:"false" json:"idcsValuePersisted"`

	// Specifiees if the attribute should be encrypted or hashed
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsSensitive SchemaSubAttributesIdcsSensitiveEnum `mandatory:"false" json:"idcsSensitive,omitempty"`

	// Trims any leading and trailing blanks from String values. Default is True.
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
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

	// Specifies whether changes to this attribute value will be audited
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

	// Target normalized attribute name that this normalized value of attribute gets mapped to for persistence. Only set for caseExact=false & searchable attributes. Do not use by default.
	// **Added In:** 19.1.4
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
	// **Added In:** 17.4.4
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
	// **Added In:** 17.4.4
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
	// **Added In:** 17.4.4
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsCanonicalValueSourceKeyAttrName *string `mandatory:"false" json:"idcsCanonicalValueSourceKeyAttrName"`

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

	// Indicates if the attribute is scim compliant, default is true
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsScimCompliant *bool `mandatory:"false" json:"idcsScimCompliant"`

	// Specifies if the attribute can be used for mapping with external identity sources such as AD or LDAP. If isSchemaMappable: false for the schema in which this attribute is defined, then this flag is ignored.
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsAttributeMappable *bool `mandatory:"false" json:"idcsAttributeMappable"`

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
	IdcsuiWidget SchemaSubAttributesIdcsuiWidgetEnum `mandatory:"false" json:"idcsuiWidget,omitempty"`

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
	// **Added In:** 18.3.6
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	IdcsPii *bool `mandatory:"false" json:"idcsPii"`

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

	// Specifies whether the attribute should be excluded from the BulkApi patch generated by gradle task for upgrading OOTB resources.
	// **Added In:** 2104150946
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: never
	//  - type: boolean
	IdcsExcludeFromUpgradePatch *bool `mandatory:"false" json:"idcsExcludeFromUpgradePatch"`
}

func (m SchemaSubAttributes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchemaSubAttributes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSchemaSubAttributesTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSchemaSubAttributesTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchemaSubAttributesMutabilityEnum(string(m.Mutability)); !ok && m.Mutability != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mutability: %s. Supported values are: %s.", m.Mutability, strings.Join(GetSchemaSubAttributesMutabilityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchemaSubAttributesReturnedEnum(string(m.Returned)); !ok && m.Returned != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Returned: %s. Supported values are: %s.", m.Returned, strings.Join(GetSchemaSubAttributesReturnedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchemaSubAttributesUniquenessEnum(string(m.Uniqueness)); !ok && m.Uniqueness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Uniqueness: %s. Supported values are: %s.", m.Uniqueness, strings.Join(GetSchemaSubAttributesUniquenessEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchemaSubAttributesIdcsSensitiveEnum(string(m.IdcsSensitive)); !ok && m.IdcsSensitive != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsSensitive: %s. Supported values are: %s.", m.IdcsSensitive, strings.Join(GetSchemaSubAttributesIdcsSensitiveEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSchemaSubAttributesIdcsuiWidgetEnum(string(m.IdcsuiWidget)); !ok && m.IdcsuiWidget != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsuiWidget: %s. Supported values are: %s.", m.IdcsuiWidget, strings.Join(GetSchemaSubAttributesIdcsuiWidgetEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SchemaSubAttributesTypeEnum Enum with underlying type: string
type SchemaSubAttributesTypeEnum string

// Set of constants representing the allowable values for SchemaSubAttributesTypeEnum
const (
	SchemaSubAttributesTypeString    SchemaSubAttributesTypeEnum = "string"
	SchemaSubAttributesTypeComplex   SchemaSubAttributesTypeEnum = "complex"
	SchemaSubAttributesTypeBoolean   SchemaSubAttributesTypeEnum = "boolean"
	SchemaSubAttributesTypeDecimal   SchemaSubAttributesTypeEnum = "decimal"
	SchemaSubAttributesTypeInteger   SchemaSubAttributesTypeEnum = "integer"
	SchemaSubAttributesTypeDatetime  SchemaSubAttributesTypeEnum = "dateTime"
	SchemaSubAttributesTypeReference SchemaSubAttributesTypeEnum = "reference"
	SchemaSubAttributesTypeBinary    SchemaSubAttributesTypeEnum = "binary"
)

var mappingSchemaSubAttributesTypeEnum = map[string]SchemaSubAttributesTypeEnum{
	"string":    SchemaSubAttributesTypeString,
	"complex":   SchemaSubAttributesTypeComplex,
	"boolean":   SchemaSubAttributesTypeBoolean,
	"decimal":   SchemaSubAttributesTypeDecimal,
	"integer":   SchemaSubAttributesTypeInteger,
	"dateTime":  SchemaSubAttributesTypeDatetime,
	"reference": SchemaSubAttributesTypeReference,
	"binary":    SchemaSubAttributesTypeBinary,
}

var mappingSchemaSubAttributesTypeEnumLowerCase = map[string]SchemaSubAttributesTypeEnum{
	"string":    SchemaSubAttributesTypeString,
	"complex":   SchemaSubAttributesTypeComplex,
	"boolean":   SchemaSubAttributesTypeBoolean,
	"decimal":   SchemaSubAttributesTypeDecimal,
	"integer":   SchemaSubAttributesTypeInteger,
	"datetime":  SchemaSubAttributesTypeDatetime,
	"reference": SchemaSubAttributesTypeReference,
	"binary":    SchemaSubAttributesTypeBinary,
}

// GetSchemaSubAttributesTypeEnumValues Enumerates the set of values for SchemaSubAttributesTypeEnum
func GetSchemaSubAttributesTypeEnumValues() []SchemaSubAttributesTypeEnum {
	values := make([]SchemaSubAttributesTypeEnum, 0)
	for _, v := range mappingSchemaSubAttributesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaSubAttributesTypeEnumStringValues Enumerates the set of values in String for SchemaSubAttributesTypeEnum
func GetSchemaSubAttributesTypeEnumStringValues() []string {
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

// GetMappingSchemaSubAttributesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaSubAttributesTypeEnum(val string) (SchemaSubAttributesTypeEnum, bool) {
	enum, ok := mappingSchemaSubAttributesTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaSubAttributesMutabilityEnum Enum with underlying type: string
type SchemaSubAttributesMutabilityEnum string

// Set of constants representing the allowable values for SchemaSubAttributesMutabilityEnum
const (
	SchemaSubAttributesMutabilityReadonly  SchemaSubAttributesMutabilityEnum = "readOnly"
	SchemaSubAttributesMutabilityReadwrite SchemaSubAttributesMutabilityEnum = "readWrite"
	SchemaSubAttributesMutabilityImmutable SchemaSubAttributesMutabilityEnum = "immutable"
	SchemaSubAttributesMutabilityWriteonly SchemaSubAttributesMutabilityEnum = "writeOnly"
)

var mappingSchemaSubAttributesMutabilityEnum = map[string]SchemaSubAttributesMutabilityEnum{
	"readOnly":  SchemaSubAttributesMutabilityReadonly,
	"readWrite": SchemaSubAttributesMutabilityReadwrite,
	"immutable": SchemaSubAttributesMutabilityImmutable,
	"writeOnly": SchemaSubAttributesMutabilityWriteonly,
}

var mappingSchemaSubAttributesMutabilityEnumLowerCase = map[string]SchemaSubAttributesMutabilityEnum{
	"readonly":  SchemaSubAttributesMutabilityReadonly,
	"readwrite": SchemaSubAttributesMutabilityReadwrite,
	"immutable": SchemaSubAttributesMutabilityImmutable,
	"writeonly": SchemaSubAttributesMutabilityWriteonly,
}

// GetSchemaSubAttributesMutabilityEnumValues Enumerates the set of values for SchemaSubAttributesMutabilityEnum
func GetSchemaSubAttributesMutabilityEnumValues() []SchemaSubAttributesMutabilityEnum {
	values := make([]SchemaSubAttributesMutabilityEnum, 0)
	for _, v := range mappingSchemaSubAttributesMutabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaSubAttributesMutabilityEnumStringValues Enumerates the set of values in String for SchemaSubAttributesMutabilityEnum
func GetSchemaSubAttributesMutabilityEnumStringValues() []string {
	return []string{
		"readOnly",
		"readWrite",
		"immutable",
		"writeOnly",
	}
}

// GetMappingSchemaSubAttributesMutabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaSubAttributesMutabilityEnum(val string) (SchemaSubAttributesMutabilityEnum, bool) {
	enum, ok := mappingSchemaSubAttributesMutabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaSubAttributesReturnedEnum Enum with underlying type: string
type SchemaSubAttributesReturnedEnum string

// Set of constants representing the allowable values for SchemaSubAttributesReturnedEnum
const (
	SchemaSubAttributesReturnedAlways  SchemaSubAttributesReturnedEnum = "always"
	SchemaSubAttributesReturnedNever   SchemaSubAttributesReturnedEnum = "never"
	SchemaSubAttributesReturnedDefault SchemaSubAttributesReturnedEnum = "default"
	SchemaSubAttributesReturnedRequest SchemaSubAttributesReturnedEnum = "request"
)

var mappingSchemaSubAttributesReturnedEnum = map[string]SchemaSubAttributesReturnedEnum{
	"always":  SchemaSubAttributesReturnedAlways,
	"never":   SchemaSubAttributesReturnedNever,
	"default": SchemaSubAttributesReturnedDefault,
	"request": SchemaSubAttributesReturnedRequest,
}

var mappingSchemaSubAttributesReturnedEnumLowerCase = map[string]SchemaSubAttributesReturnedEnum{
	"always":  SchemaSubAttributesReturnedAlways,
	"never":   SchemaSubAttributesReturnedNever,
	"default": SchemaSubAttributesReturnedDefault,
	"request": SchemaSubAttributesReturnedRequest,
}

// GetSchemaSubAttributesReturnedEnumValues Enumerates the set of values for SchemaSubAttributesReturnedEnum
func GetSchemaSubAttributesReturnedEnumValues() []SchemaSubAttributesReturnedEnum {
	values := make([]SchemaSubAttributesReturnedEnum, 0)
	for _, v := range mappingSchemaSubAttributesReturnedEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaSubAttributesReturnedEnumStringValues Enumerates the set of values in String for SchemaSubAttributesReturnedEnum
func GetSchemaSubAttributesReturnedEnumStringValues() []string {
	return []string{
		"always",
		"never",
		"default",
		"request",
	}
}

// GetMappingSchemaSubAttributesReturnedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaSubAttributesReturnedEnum(val string) (SchemaSubAttributesReturnedEnum, bool) {
	enum, ok := mappingSchemaSubAttributesReturnedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaSubAttributesUniquenessEnum Enum with underlying type: string
type SchemaSubAttributesUniquenessEnum string

// Set of constants representing the allowable values for SchemaSubAttributesUniquenessEnum
const (
	SchemaSubAttributesUniquenessNone   SchemaSubAttributesUniquenessEnum = "none"
	SchemaSubAttributesUniquenessServer SchemaSubAttributesUniquenessEnum = "server"
	SchemaSubAttributesUniquenessGlobal SchemaSubAttributesUniquenessEnum = "global"
)

var mappingSchemaSubAttributesUniquenessEnum = map[string]SchemaSubAttributesUniquenessEnum{
	"none":   SchemaSubAttributesUniquenessNone,
	"server": SchemaSubAttributesUniquenessServer,
	"global": SchemaSubAttributesUniquenessGlobal,
}

var mappingSchemaSubAttributesUniquenessEnumLowerCase = map[string]SchemaSubAttributesUniquenessEnum{
	"none":   SchemaSubAttributesUniquenessNone,
	"server": SchemaSubAttributesUniquenessServer,
	"global": SchemaSubAttributesUniquenessGlobal,
}

// GetSchemaSubAttributesUniquenessEnumValues Enumerates the set of values for SchemaSubAttributesUniquenessEnum
func GetSchemaSubAttributesUniquenessEnumValues() []SchemaSubAttributesUniquenessEnum {
	values := make([]SchemaSubAttributesUniquenessEnum, 0)
	for _, v := range mappingSchemaSubAttributesUniquenessEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaSubAttributesUniquenessEnumStringValues Enumerates the set of values in String for SchemaSubAttributesUniquenessEnum
func GetSchemaSubAttributesUniquenessEnumStringValues() []string {
	return []string{
		"none",
		"server",
		"global",
	}
}

// GetMappingSchemaSubAttributesUniquenessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaSubAttributesUniquenessEnum(val string) (SchemaSubAttributesUniquenessEnum, bool) {
	enum, ok := mappingSchemaSubAttributesUniquenessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaSubAttributesIdcsSensitiveEnum Enum with underlying type: string
type SchemaSubAttributesIdcsSensitiveEnum string

// Set of constants representing the allowable values for SchemaSubAttributesIdcsSensitiveEnum
const (
	SchemaSubAttributesIdcsSensitiveEncrypt  SchemaSubAttributesIdcsSensitiveEnum = "encrypt"
	SchemaSubAttributesIdcsSensitiveHash     SchemaSubAttributesIdcsSensitiveEnum = "hash"
	SchemaSubAttributesIdcsSensitiveHashSc   SchemaSubAttributesIdcsSensitiveEnum = "hash_sc"
	SchemaSubAttributesIdcsSensitiveChecksum SchemaSubAttributesIdcsSensitiveEnum = "checksum"
	SchemaSubAttributesIdcsSensitiveNone     SchemaSubAttributesIdcsSensitiveEnum = "none"
)

var mappingSchemaSubAttributesIdcsSensitiveEnum = map[string]SchemaSubAttributesIdcsSensitiveEnum{
	"encrypt":  SchemaSubAttributesIdcsSensitiveEncrypt,
	"hash":     SchemaSubAttributesIdcsSensitiveHash,
	"hash_sc":  SchemaSubAttributesIdcsSensitiveHashSc,
	"checksum": SchemaSubAttributesIdcsSensitiveChecksum,
	"none":     SchemaSubAttributesIdcsSensitiveNone,
}

var mappingSchemaSubAttributesIdcsSensitiveEnumLowerCase = map[string]SchemaSubAttributesIdcsSensitiveEnum{
	"encrypt":  SchemaSubAttributesIdcsSensitiveEncrypt,
	"hash":     SchemaSubAttributesIdcsSensitiveHash,
	"hash_sc":  SchemaSubAttributesIdcsSensitiveHashSc,
	"checksum": SchemaSubAttributesIdcsSensitiveChecksum,
	"none":     SchemaSubAttributesIdcsSensitiveNone,
}

// GetSchemaSubAttributesIdcsSensitiveEnumValues Enumerates the set of values for SchemaSubAttributesIdcsSensitiveEnum
func GetSchemaSubAttributesIdcsSensitiveEnumValues() []SchemaSubAttributesIdcsSensitiveEnum {
	values := make([]SchemaSubAttributesIdcsSensitiveEnum, 0)
	for _, v := range mappingSchemaSubAttributesIdcsSensitiveEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaSubAttributesIdcsSensitiveEnumStringValues Enumerates the set of values in String for SchemaSubAttributesIdcsSensitiveEnum
func GetSchemaSubAttributesIdcsSensitiveEnumStringValues() []string {
	return []string{
		"encrypt",
		"hash",
		"hash_sc",
		"checksum",
		"none",
	}
}

// GetMappingSchemaSubAttributesIdcsSensitiveEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaSubAttributesIdcsSensitiveEnum(val string) (SchemaSubAttributesIdcsSensitiveEnum, bool) {
	enum, ok := mappingSchemaSubAttributesIdcsSensitiveEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaSubAttributesIdcsuiWidgetEnum Enum with underlying type: string
type SchemaSubAttributesIdcsuiWidgetEnum string

// Set of constants representing the allowable values for SchemaSubAttributesIdcsuiWidgetEnum
const (
	SchemaSubAttributesIdcsuiWidgetInputtext SchemaSubAttributesIdcsuiWidgetEnum = "inputtext"
	SchemaSubAttributesIdcsuiWidgetCheckbox  SchemaSubAttributesIdcsuiWidgetEnum = "checkbox"
	SchemaSubAttributesIdcsuiWidgetTextarea  SchemaSubAttributesIdcsuiWidgetEnum = "textarea"
	SchemaSubAttributesIdcsuiWidgetCombobox  SchemaSubAttributesIdcsuiWidgetEnum = "combobox"
)

var mappingSchemaSubAttributesIdcsuiWidgetEnum = map[string]SchemaSubAttributesIdcsuiWidgetEnum{
	"inputtext": SchemaSubAttributesIdcsuiWidgetInputtext,
	"checkbox":  SchemaSubAttributesIdcsuiWidgetCheckbox,
	"textarea":  SchemaSubAttributesIdcsuiWidgetTextarea,
	"combobox":  SchemaSubAttributesIdcsuiWidgetCombobox,
}

var mappingSchemaSubAttributesIdcsuiWidgetEnumLowerCase = map[string]SchemaSubAttributesIdcsuiWidgetEnum{
	"inputtext": SchemaSubAttributesIdcsuiWidgetInputtext,
	"checkbox":  SchemaSubAttributesIdcsuiWidgetCheckbox,
	"textarea":  SchemaSubAttributesIdcsuiWidgetTextarea,
	"combobox":  SchemaSubAttributesIdcsuiWidgetCombobox,
}

// GetSchemaSubAttributesIdcsuiWidgetEnumValues Enumerates the set of values for SchemaSubAttributesIdcsuiWidgetEnum
func GetSchemaSubAttributesIdcsuiWidgetEnumValues() []SchemaSubAttributesIdcsuiWidgetEnum {
	values := make([]SchemaSubAttributesIdcsuiWidgetEnum, 0)
	for _, v := range mappingSchemaSubAttributesIdcsuiWidgetEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaSubAttributesIdcsuiWidgetEnumStringValues Enumerates the set of values in String for SchemaSubAttributesIdcsuiWidgetEnum
func GetSchemaSubAttributesIdcsuiWidgetEnumStringValues() []string {
	return []string{
		"inputtext",
		"checkbox",
		"textarea",
		"combobox",
	}
}

// GetMappingSchemaSubAttributesIdcsuiWidgetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaSubAttributesIdcsuiWidgetEnum(val string) (SchemaSubAttributesIdcsuiWidgetEnum, bool) {
	enum, ok := mappingSchemaSubAttributesIdcsuiWidgetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
