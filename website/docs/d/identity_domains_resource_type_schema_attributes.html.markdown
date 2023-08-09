---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_resource_type_schema_attributes"
sidebar_current: "docs-oci-datasource-identity_domains-resource_type_schema_attributes"
description: |-
  Provides the list of Resource Type Schema Attributes in Oracle Cloud Infrastructure Identity Domains service
---

# Data Source: oci_identity_domains_resource_type_schema_attributes
This data source provides the list of Resource Type Schema Attributes in Oracle Cloud Infrastructure Identity Domains service.

Search Resource Type Schema Attributes

## Example Usage

```hcl
data "oci_identity_domains_resource_type_schema_attributes" "test_resource_type_schema_attributes" {
	#Required
	idcs_endpoint = data.oci_identity_domain.test_domain.url

	#Optional
	resource_type_schema_attribute_count = var.resource_type_schema_attribute_resource_type_schema_attribute_count
	resource_type_schema_attribute_filter = var.resource_type_schema_attribute_resource_type_schema_attribute_filter
	attribute_sets = ["all"]
	attributes = ""
	authorization = var.resource_type_schema_attribute_authorization
	resource_type_schema_version = var.resource_type_schema_attribute_resource_type_schema_version
	start_index = var.resource_type_schema_attribute_start_index
}
```

## Argument Reference

The following arguments are supported:

* `resource_type_schema_attribute_count` - (Optional) OPTIONAL. An integer that indicates the desired maximum number of query results per page. 1000 is the largest value that you can use. See the Pagination section of the System for Cross-Domain Identity Management Protocol specification for more information. (Section 3.4.2.4).
* `resource_type_schema_attribute_filter` - (Optional) OPTIONAL. The filter string that is used to request a subset of resources. The filter string MUST be a valid filter expression. See the Filtering section of the SCIM specification for more information (Section 3.4.2.2). The string should contain at least one condition that each item must match in order to be returned in the search results. Each condition specifies an attribute, an operator, and a value. Conditions within a filter can be connected by logical operators (such as AND and OR). Sets of conditions can be grouped together using parentheses.
* `attribute_sets` - (Optional) A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If 'attributes' query parameter is also available, union of the two is fetched. Valid values - all, always, never, request, default. Values are case-insensitive.
* `attributes` - (Optional) A comma-delimited string that specifies the names of resource attributes that should be returned in the response. By default, a response that contains resource attributes contains only attributes that are defined in the schema for that resource type as returned=always or returned=default. An attribute that is defined as returned=request is returned in a response only if the request specifies its name in the value of this query parameter. If a request specifies this query parameter, the response contains the attributes that this query parameter specifies, as well as any attribute that is defined as returned=always.
* `authorization` - (Optional) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
* `idcs_endpoint` - (Required) The basic endpoint for the identity domain
* `resource_type_schema_attribute_count` - (Optional) OPTIONAL. An integer that indicates the desired maximum number of query results per page. 1000 is the largest value that you can use. See the Pagination section of the System for Cross-Domain Identity Management Protocol specification for more information. (Section 3.4.2.4).
* `resource_type_schema_attribute_filter` - (Optional) OPTIONAL. The filter string that is used to request a subset of resources. The filter string MUST be a valid filter expression. See the Filtering section of the SCIM specification for more information (Section 3.4.2.2). The string should contain at least one condition that each item must match in order to be returned in the search results. Each condition specifies an attribute, an operator, and a value. Conditions within a filter can be connected by logical operators (such as AND and OR). Sets of conditions can be grouped together using parentheses.
* `resource_type_schema_version` - (Optional) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `start_index` - (Optional) OPTIONAL. An integer that indicates the 1-based index of the first query result. See the Pagination section of the SCIM specification for more information. (Section 3.4.2.4). The number of results pages to return. The first page is 1. Specify 2 to access the second page of results, and so on.


## Attributes Reference

The following attributes are exported:

* `resource_type_schema_attributes` - The list of resource_type_schema_attributes.

### ResourceTypeSchemaAttribute Reference

The following attributes are exported:

* `resources` - A multi-valued list of complex objects containing the requested resources. This MAY be a subset of the full set of resources if pagination is requested. REQUIRED if "totalResults" is non-zero.
	* `canonical_values` - A collection of canonical values. Applicable Service Providers MUST specify the canonical types specified in the core schema specification--for example, \"work\", \"home\".

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `case_exact` - Specifies if the String attribute is case-sensitive

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: boolean
	* `compartment_ocid` - Oracle Cloud Infrastructure Compartment Id (ocid) in which the resource lives.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `delete_in_progress` - A boolean flag indicating this resource in the process of being deleted. Usually set to true when synchronous deletion of the resource would take too long.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `description` - The attribute's human-readable description

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* idcsSearchable: true
		* type: string
		* uniqueness: none
	* `domain_ocid` - Oracle Cloud Infrastructure Domain Id (ocid) in which the resource lives.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `end_user_mutability` - Specifies User mutability for this attribute

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `end_user_mutability_allowed_values` - Specifies the list of User mutabilities allowed

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `id` - Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: always
		* type: string
		* uniqueness: global
	* `idcs_added_since_release_number` - Indicates that the schema has been added since this release number

		**Added In:** 17.3.4

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* idcsSearchable: true
		* required: false
		* returned: default
		* type: string
	* `idcs_added_since_version` - Indicates that the schema has been added since version

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: integer
	* `idcs_attribute_cacheable` - Specifies whether the attribute is cacheable. True by default for all attributes. If attribute with idcsAttributeCachable = false, is present \"attributesToGet\" while executing GET/SEARCH on cacheable resource, Cache is missed and data is fetched from Data Provider.

		**Added In:** 17.3.4

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* idcsSearchable: true
		* required: false
		* returned: default
		* type: boolean
	* `idcs_attribute_mappable` - Specifies if the attribute can be used for mapping with external identity sources such as AD or LDAP. If isSchemaMappable: false for the schema in which this attribute is defined, then this flag is ignored

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: boolean
	* `idcs_auditable` - Specifies whether changes to this attribute value are audited

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* idcsSearchable: true
		* required: false
		* returned: default
		* type: boolean
	* `idcs_auto_increment_seq_name` - Sequence tracking ID name for the attribute

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
	* `idcs_canonical_value_source_filter` - Filter to use when getting canonical values for this schema attribute

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_canonical_value_source_resource_type` - Specifies the Resource type to read from for dynamic canonical values

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_composite_key` - The set of one or more sub attributes' names of a CMVA, whose values uniquely identify an instance of a CMVA

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_created_by` - The User or App who created the Resource

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: complex
		* `display` - The displayName of the User or App who created this Resource

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ocid` - The OCID of the SCIM resource that represents the User or App who created this Resource

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI of the SCIM resource that represents the User or App who created this Resource

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - The type of resource, User or App, that created this Resource

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - The ID of the SCIM resource that represents the User or App who created this Resource

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `idcs_csv_column_header_name` - The attribute defining the CSV column header name for import/export

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_custom_attribute` - custom attribute flag.

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* idcsSearchable: true
		* uniqueness: none
	* `idcs_deprecated_since_release_number` - Indicates that the schema has been deprecated since this release number

		**Added In:** 17.3.4

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* idcsSearchable: true
		* required: false
		* returned: default
		* type: string
	* `idcs_deprecated_since_version` - Indicates that the schema has been deprecated since version

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: integer
	* `idcs_display_name` - Specifies the user-friendly displayable attribute name or catalog key used for localization

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* idcsSearchable: true
		* type: string
		* uniqueness: none
	* `idcs_display_name_message_id` - Localized schema attribute display name for use by UI client  for displaying attribute labels

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* idcsSearchable: true
		* type: string
		* uniqueness: none
	* `idcs_fetch_complex_attribute_values` - **SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: boolean
		* uniqueness: none Whether the CMVA attribute will be fetched or not for current resource in AbstractResourceManager update operation before calling data provider update. Default is true. 
	* `idcs_from_target_mapper` - Specifies the mapper to use when mapping this attribute value from DataProvider-specific semantics

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_fully_qualified_name` - Fully qualified name of this attribute

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* idcsSearchable: true
		* uniqueness: none
	* `idcs_generated` - Specifies whether this attribute value was generated

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: boolean
	* `idcs_icf_attribute_type` - Maps to ICF data type

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* idcsSearchable: true
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_icf_bundle_attribute_name` - Maps to ICF target attribute name

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_icf_required` - Metadata to identify the ICF required attribute

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `idcs_indirect_ref_resource_attributes` - Specifies the indirectly referenced Resources

		**SCIM++ Properties:**
		* multiValued: true
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
	* `idcs_internal` - Specifies whether the schema attribute is for internal use only. Internal attributes are not exposed via REST. This attribute overrides mutability for create/update if the request is internal and the attribute internalflag is set to True. This attribute overrides the return attribute while building SCIM response attributes when both the request is internal and the schema attribute is internal.

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: boolean
	* `idcs_last_modified_by` - The User or App who modified the Resource

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* `display` - The displayName of the User or App who modified this Resource

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ocid` - The OCID of the SCIM resource that represents the User or App who modified this Resource

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI of the SCIM resource that represents the User or App who modified this Resource

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - The type of resource, User or App, that modified this Resource

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - The ID of the SCIM resource that represents the User or App who modified this Resource

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `idcs_last_upgraded_in_release` - The release number when the resource was upgraded.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `idcs_max_length` - Specifies the maximum length of the attribute

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: integer
	* `idcs_max_value` - Specifies the maximum value of the integer attribute

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* idcsSearchable: true
		* required: false
		* returned: default
		* type: integer
	* `idcs_min_length` - Specifies the minimum length of the attribute

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* idcsSearchable: true
		* required: false
		* returned: default
		* type: integer
	* `idcs_min_value` - Specifies the minimum value of the integer attribute

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: integer
	* `idcs_multi_language` - If true, specifies that the attribute can have multiple language values set for the attribute on which this is set.

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readOnly
		* idcsSearchable: true
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `idcs_prevented_operations` - Each value of this attribute specifies an operation that only an internal client may perform on this particular resource.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `idcs_ref_resource_attribute` - Specifies the referenced Resource attribute

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* idcsSearchable: true
		* required: false
		* returned: default
		* type: string
	* `idcs_ref_resource_attributes` - Specifies the directly referenced Resources

		**SCIM++ Properties:**
		* multiValued: true
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
	* `idcs_schema_urn` - Schema URN string that this attribute belongs to

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* idcsSearchable: true
		* uniqueness: none
	* `idcs_scim_compliant` - Indicates if the attribute is scim compliant, default is true

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* idcsSearchable: true
		* required: false
		* returned: default
		* type: boolean
	* `idcs_searchable` - Specifies whether this attribute can be included in a search filter

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: boolean
	* `idcs_sensitive` - Flag to specify if the attribute should be encrypted or hashed

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_target_attribute_name` - Target attribute name that this attribute gets mapped to for persistence

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_target_attribute_name_to_migrate_from` - Old Target attribute name from child table for CSVA attribute prior to migration. This maintains this attribute used to get mapped to for persistence

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_target_norm_attribute_name` - Target normalized attribute name that this normalized value of attribute gets mapped to for persistence. Only set for caseExact=false & searchable attributes. Do not use by default.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_target_unique_constraint_name` - Target index name created for this attribute for performance

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_to_target_mapper` - Specifies the mapper to use when mapping this attribute value to DataProvider-specific semantics

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `idcs_trim_string_value` - Trims any leading and trailing blanks from String values. Default is True.

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: boolean
	* `idcs_validate_reference` - Validate payload reference value during create, replace, and update. Default is True.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `idcs_value_persisted` - Specifies whether the value of the Resource attribute is persisted

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: boolean
	* `meta` - A complex attribute that contains resource metadata. All sub-attributes are OPTIONAL.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Created Date, mapsTo:meta.created]]
		* type: complex
		* `created` - The DateTime the Resource was added to the Service Provider

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: dateTime
			* uniqueness: none
		* `last_modified` - The most recent DateTime that the details of this Resource were updated at the Service Provider. If this Resource has never been modified since its initial creation, the value MUST be the same as the value of created. The attribute MUST be a DateTime.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: dateTime
			* uniqueness: none
		* `location` - The URI of the Resource being returned. This value MUST be the same as the Location HTTP response header.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `resource_type` - Name of the resource type of the resource--for example, Users or Groups

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `version` - The version of the Resource being returned. This value must be the same as the ETag HTTP response header.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `multi_valued` - Indicates the attribute's plurality

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: boolean
	* `mutability` - Specifies if the attribute is mutable

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* idcsSearchable: true
		* type: string
		* uniqueness: none
	* `name` - Attribute's name

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* idcsSearchable: true
		* type: string
		* uniqueness: none
	* `ocid` - Unique Oracle Cloud Infrastructure identifier for the SCIM Resource.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: string
		* uniqueness: global
	* `reference_types` - The names of the Resource types that may be referenced--for example, User. This is only applicable for attributes that are of the \"reference\" data type.

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `required` - Specifies if the attribute is required

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: boolean
	* `resource_type` - ResourceType this attribute belongs to.

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* idcsSearchable: true
		* uniqueness: none
	* `returned` - A single keyword that indicates when an attribute and associated values are returned in response to a GET request or in response to a PUT, POST, or PATCH request

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* idcsSearchable: true
		* type: string
		* uniqueness: none
	* `schemas` - REQUIRED. The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. This specification defines URIs for User, Group, and a standard \"enterprise\" extension. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `tags` - A list of tags on this resource.

		**SCIM++ Properties:**
		* idcsCompositeKey: [key, value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `key` - Key or name of the tag.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - Value of the tag.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `tenancy_ocid` - Oracle Cloud Infrastructure Tenant Id (ocid) in which the resource lives.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `type` - The attribute's data type--for example, String

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* idcsSearchable: true
		* returned: default
		* type: string
		* uniqueness: none
	* `uniqueness` - A single keyword value that specifies how the Service Provider enforces uniqueness of attribute values. A server MAY reject an invalid value based on uniqueness by returning an HTTP response code of 400 (Bad Request). A client MAY enforce uniqueness on the client side to a greater degree than the Service Provider enforces. For example, a client could make a value unique while the server has the uniqueness of \"none\".

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* required: false
		* returned: default
		* idcsSearchable: true
		* type: string
		* uniqueness: none
* `items_per_page` - The number of resources returned in a list response page. REQUIRED when partial results returned due to pagination.
* `schemas` - The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior. REQUIRED.
* `start_index` - The 1-based index of the first result in the current set of list results.  REQUIRED when partial results returned due to pagination.
* `total_results` - The total number of results returned by the list or query operation.  The value may be larger than the number of resources returned such as when returning a single page of results where multiple pages are available. REQUIRED.

