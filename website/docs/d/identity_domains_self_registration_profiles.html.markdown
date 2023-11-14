---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_self_registration_profiles"
sidebar_current: "docs-oci-datasource-identity_domains-self_registration_profiles"
description: |-
  Provides the list of Self Registration Profiles in Oracle Cloud Infrastructure Identity Domains service
---

# Data Source: oci_identity_domains_self_registration_profiles
This data source provides the list of Self Registration Profiles in Oracle Cloud Infrastructure Identity Domains service.

Search for self-registration profiles.

## Example Usage

```hcl
data "oci_identity_domains_self_registration_profiles" "test_self_registration_profiles" {
	#Required
	idcs_endpoint = data.oci_identity_domain.test_domain.url

	#Optional
	self_registration_profile_count = var.self_registration_profile_self_registration_profile_count
	self_registration_profile_filter = var.self_registration_profile_self_registration_profile_filter
	attribute_sets = ["all"]
	attributes = ""
	authorization = var.self_registration_profile_authorization
	resource_type_schema_version = var.self_registration_profile_resource_type_schema_version
	start_index = var.self_registration_profile_start_index
}
```

## Argument Reference

The following arguments are supported:

* `self_registration_profile_count` - (Optional) OPTIONAL. An integer that indicates the desired maximum number of query results per page. 1000 is the largest value that you can use. See the Pagination section of the System for Cross-Domain Identity Management Protocol specification for more information. (Section 3.4.2.4).
* `self_registration_profile_filter` - (Optional) OPTIONAL. The filter string that is used to request a subset of resources. The filter string MUST be a valid filter expression. See the Filtering section of the SCIM specification for more information (Section 3.4.2.2). The string should contain at least one condition that each item must match in order to be returned in the search results. Each condition specifies an attribute, an operator, and a value. Conditions within a filter can be connected by logical operators (such as AND and OR). Sets of conditions can be grouped together using parentheses.
* `attribute_sets` - (Optional) A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If 'attributes' query parameter is also available, union of the two is fetched. Valid values - all, always, never, request, default. Values are case-insensitive.
* `attributes` - (Optional) A comma-delimited string that specifies the names of resource attributes that should be returned in the response. By default, a response that contains resource attributes contains only attributes that are defined in the schema for that resource type as returned=always or returned=default. An attribute that is defined as returned=request is returned in a response only if the request specifies its name in the value of this query parameter. If a request specifies this query parameter, the response contains the attributes that this query parameter specifies, as well as any attribute that is defined as returned=always.
* `authorization` - (Optional) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
* `idcs_endpoint` - (Required) The basic endpoint for the identity domain
* `resource_type_schema_version` - (Optional) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `start_index` - (Optional) OPTIONAL. An integer that indicates the 1-based index of the first query result. See the Pagination section of the SCIM specification for more information. (Section 3.4.2.4). The number of results pages to return. The first page is 1. Specify 2 to access the second page of results, and so on.


## Attributes Reference

The following attributes are exported:

* `self_registration_profiles` - The list of self_registration_profiles.

### SelfRegistrationProfile Reference

The following attributes are exported:

* `activation_email_required` - **SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none A Boolean value that indicates whether Account verification email is required to be sent before login or not 
* `active` - A Boolean value that indicates whether the profile is enabled or not

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `after_submit_text` - Text to be displayed on UI after doing self registration

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* idcsMultiLanguage: true
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `default` - If true, specifies that the localized attribute instance value is the default and will be returned if no localized value found for requesting user's preferred locale. One and only one instance should have this attribute set to true.

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `locale` - Type of user's locale e.g. en-CA

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(type)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Localized value of after submit text in corresponding locale

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `allowed_email_domains` - A Multivalue String value for Email domains which are valid for this profile

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
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
* `consent_text` - Consent text

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* idcsMultiLanguage: true
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `default` - If true, specifies that the localized attribute instance value is the default and will be returned if no localized value found for requesting user's preferred locale. One and only one instance should have this attribute set to true.

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `locale` - Type of user's locale e.g. en-CA

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(type)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Localized value of consent text in corresponding locale

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `consent_text_present` - A boolean value that indicates whether the consent text is present.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `default_groups` - Default groups assigned to the user

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: request
	* type: complex
	* `ref` - URI of the Default Group

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `display` - A human readable name, primarily used for display purposes. READ-ONLY.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Identifier of the Default Group.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
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
* `disallowed_email_domains` - A Multivalue String Value for Email domains to be handled as exceptions

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `display_name` - Registration page name

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* idcsMultiLanguage: true
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `default` - If true, specifies that the localized attribute instance value is the default and will be returned if no localized value found for requesting user's preferred locale. One and only one instance should have this attribute set to true.

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `locale` - Type of user's locale e.g. en-CA

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(type)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Localized value of displayName in corresponding locale

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
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
* `email_template` - Email template

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: request
	* type: complex
	* `ref` - URI of the Email Template

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `display` - A human readable name, primarily used for display purposes. READ-ONLY.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Identifier of the Email Template.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `external_id` - An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `footer_logo` - References to footer logo

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: reference
* `footer_text` - Footer text

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* idcsMultiLanguage: true
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `default` - If true, specifies that the localized attribute instance value is the default and will be returned if no localized value found for requesting user's preferred locale. One and only one instance should have this attribute set to true.

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `locale` - Type of user's locale e.g. en-CA

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(type)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Localized value of footer text in corresponding locale

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `header_logo` - Reference to header logo

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: reference
* `header_text` - Header text

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* idcsMultiLanguage: true
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `default` - If true, specifies that the localized attribute instance value is the default and will be returned if no localized value found for requesting user's preferred locale. One and only one instance should have this attribute set to true.

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `locale` - Type of user's locale e.g. en-CA

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(type)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Localized value of header text in corresponding locale

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
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
* `idcs_created_by` - The User or App who created the Resource

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: true
	* returned: default
	* type: complex
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
* `idcs_last_modified_by` - The User or App who modified the Resource

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: complex
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
* `idcs_prevented_operations` - Each value of this attribute specifies an operation that only an internal client may perform on this particular resource.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
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
* `name` - Name of the profile

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: always
	* type: string
	* uniqueness: global
* `number_of_days_redirect_url_is_valid` - Number of days redirect URL is valid

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: integer
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
* `redirect_url` - This URL will be replaced in email notification sent to user. When activation email required is set to true, user is created in \"pending verification\" state, upon clicking this link user will be able to activate himself. When activation email required is set to false, user is created in \"verified\" state, this link will be used to verify user's email.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
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
* `show_on_login_page` - A Boolean value that indicates whether the profile should be displayed on login page

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
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
* `user_attributes` - **SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none User Attributes 
	* `deletable` - If this attribute can be deleted

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `fully_qualified_attribute_name` - **SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none Fully Qualified Attribute Name 
	* `metadata` - Metadata of the user attribute

		**Added In:** 18.1.6

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `seq_number` - **SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none Sequence Number for the attribute 
	* `value` - name of the attribute

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none

