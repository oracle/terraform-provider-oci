---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_self_registration_profile"
sidebar_current: "docs-oci-resource-identity_domains-self_registration_profile"
description: |-
  Provides the Self Registration Profile resource in Oracle Cloud Infrastructure Identity Domains service
---

# oci_identity_domains_self_registration_profile
This resource provides the Self Registration Profile resource in Oracle Cloud Infrastructure Identity Domains service.

Create a self-registration profile.

## Example Usage

```hcl
resource "oci_identity_domains_self_registration_profile" "test_self_registration_profile" {
	#Required
	activation_email_required = var.self_registration_profile_activation_email_required
	consent_text_present = var.self_registration_profile_consent_text_present
	display_name {
		#Required
		locale = var.self_registration_profile_display_name_locale
		value = var.self_registration_profile_display_name_value

		#Optional
		default = var.self_registration_profile_display_name_default
	}
	email_template {
		#Required
		value = var.self_registration_profile_email_template_value
	}
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	name = var.self_registration_profile_name
	number_of_days_redirect_url_is_valid = var.self_registration_profile_number_of_days_redirect_url_is_valid
	redirect_url = var.self_registration_profile_redirect_url
	schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:SelfRegistrationProfile"]
	show_on_login_page = var.self_registration_profile_show_on_login_page

	#Optional
	active = var.self_registration_profile_active
	after_submit_text {
		#Required
		locale = var.self_registration_profile_after_submit_text_locale
		value = var.self_registration_profile_after_submit_text_value

		#Optional
		default = var.self_registration_profile_after_submit_text_default
	}
	allowed_email_domains = var.self_registration_profile_allowed_email_domains
	attribute_sets = ["all"]
	attributes = ""
	authorization = var.self_registration_profile_authorization
	consent_text {
		#Required
		locale = var.self_registration_profile_consent_text_locale
		value = var.self_registration_profile_consent_text_value

		#Optional
		default = var.self_registration_profile_consent_text_default
	}
	default_groups {
		#Required
		value = var.self_registration_profile_default_groups_value
	}
	disallowed_email_domains = var.self_registration_profile_disallowed_email_domains
	external_id = "externalId"
	footer_logo = var.self_registration_profile_footer_logo
	footer_text {
		#Required
		locale = var.self_registration_profile_footer_text_locale
		value = var.self_registration_profile_footer_text_value

		#Optional
		default = var.self_registration_profile_footer_text_default
	}
	header_logo = var.self_registration_profile_header_logo
	header_text {
		#Required
		locale = var.self_registration_profile_header_text_locale
		value = var.self_registration_profile_header_text_value

		#Optional
		default = var.self_registration_profile_header_text_default
	}
	id = var.self_registration_profile_id
	ocid = var.self_registration_profile_ocid
	resource_type_schema_version = var.self_registration_profile_resource_type_schema_version
	tags {
		#Required
		key = var.self_registration_profile_tags_key
		value = var.self_registration_profile_tags_value
	}
	user_attributes {
		#Required
		seq_number = var.self_registration_profile_user_attributes_seq_number
		value = var.self_registration_profile_user_attributes_value

		#Optional
		fully_qualified_attribute_name = var.self_registration_profile_user_attributes_fully_qualified_attribute_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `activation_email_required` - (Required) (Updatable) **SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none A Boolean value that indicates whether Account verification email is required to be sent before login or not 
* `active` - (Optional) (Updatable) A Boolean value that indicates whether the profile is enabled or not

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `after_submit_text` - (Optional) (Updatable) Text to be displayed on UI after doing self registration

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
	* `default` - (Optional) (Updatable) If true, specifies that the localized attribute instance value is the default and will be returned if no localized value found for requesting user's preferred locale. One and only one instance should have this attribute set to true.

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `locale` - (Required) (Updatable) Type of user's locale e.g. en-CA

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
	* `value` - (Required) (Updatable) Localized value of after submit text in corresponding locale

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `allowed_email_domains` - (Optional) (Updatable) A Multivalue String value for Email domains which are valid for this profile

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `attribute_sets` - (Optional) (Updatable) A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If 'attributes' query parameter is also available, union of the two is fetched. Valid values - all, always, never, request, default. Values are case-insensitive.
* `attributes` - (Optional) (Updatable) A comma-delimited string that specifies the names of resource attributes that should be returned in the response. By default, a response that contains resource attributes contains only attributes that are defined in the schema for that resource type as returned=always or returned=default. An attribute that is defined as returned=request is returned in a response only if the request specifies its name in the value of this query parameter. If a request specifies this query parameter, the response contains the attributes that this query parameter specifies, as well as any attribute that is defined as returned=always.
* `authorization` - (Optional) (Updatable) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
* `compartment_ocid` - (Optional) (Updatable) Oracle Cloud Infrastructure Compartment Id (ocid) in which the resource lives.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `consent_text` - (Optional) (Updatable) Consent text

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
	* `default` - (Optional) (Updatable) If true, specifies that the localized attribute instance value is the default and will be returned if no localized value found for requesting user's preferred locale. One and only one instance should have this attribute set to true.

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `locale` - (Required) (Updatable) Type of user's locale e.g. en-CA

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
	* `value` - (Required) (Updatable) Localized value of consent text in corresponding locale

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `consent_text_present` - (Required) (Updatable) A boolean value that indicates whether the consent text is present.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `default_groups` - (Optional) (Updatable) Default groups assigned to the user

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: request
	* type: complex
	* `ref` - (Optional) (Updatable) URI of the Default Group

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `display` - (Optional) (Updatable) A human readable name, primarily used for display purposes. READ-ONLY.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) Identifier of the Default Group.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `delete_in_progress` - (Optional) (Updatable) A boolean flag indicating this resource in the process of being deleted. Usually set to true when synchronous deletion of the resource would take too long.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `disallowed_email_domains` - (Optional) (Updatable) A Multivalue String Value for Email domains to be handled as exceptions

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `display_name` - (Required) (Updatable) Registration page name

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
	* `default` - (Optional) (Updatable) If true, specifies that the localized attribute instance value is the default and will be returned if no localized value found for requesting user's preferred locale. One and only one instance should have this attribute set to true.

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `locale` - (Required) (Updatable) Type of user's locale e.g. en-CA

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
	* `value` - (Required) (Updatable) Localized value of displayName in corresponding locale

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `domain_ocid` - (Optional) (Updatable) Oracle Cloud Infrastructure Domain Id (ocid) in which the resource lives.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `email_template` - (Required) (Updatable) Email template

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: request
	* type: complex
	* `ref` - (Optional) (Updatable) URI of the Email Template

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `display` - (Optional) (Updatable) A human readable name, primarily used for display purposes. READ-ONLY.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) Identifier of the Email Template.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `external_id` - (Optional) (Updatable) An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `footer_logo` - (Optional) (Updatable) References to footer logo

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: reference
* `footer_text` - (Optional) (Updatable) Footer text

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
	* `default` - (Optional) (Updatable) If true, specifies that the localized attribute instance value is the default and will be returned if no localized value found for requesting user's preferred locale. One and only one instance should have this attribute set to true.

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `locale` - (Required) (Updatable) Type of user's locale e.g. en-CA

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
	* `value` - (Required) (Updatable) Localized value of footer text in corresponding locale

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `header_logo` - (Optional) (Updatable) Reference to header logo

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: reference
* `header_text` - (Optional) (Updatable) Header text

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
	* `default` - (Optional) (Updatable) If true, specifies that the localized attribute instance value is the default and will be returned if no localized value found for requesting user's preferred locale. One and only one instance should have this attribute set to true.

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `locale` - (Required) (Updatable) Type of user's locale e.g. en-CA

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
	* `value` - (Required) (Updatable) Localized value of header text in corresponding locale

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `id` - (Optional) (Updatable) Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: always
	* type: string
	* uniqueness: global
* `idcs_created_by` - (Optional) (Updatable) The User or App who created the Resource

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: true
	* returned: default
	* type: complex
	* `ref` - (Optional) (Updatable) The URI of the SCIM resource that represents the User or App who created this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `display` - (Optional) (Updatable) The displayName of the User or App who created this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocid` - (Optional) (Updatable) The OCID of the SCIM resource that represents the User or App who created this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* returned: default
		* type: string
		* uniqueness: none
	* `type` - (Optional) (Updatable) The type of resource, User or App, that created this Resource

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) The ID of the SCIM resource that represents the User or App who created this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `idcs_endpoint` - (Required) The basic endpoint for the identity domain
* `idcs_last_modified_by` - (Optional) (Updatable) The User or App who modified the Resource

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: complex
	* `ref` - (Optional) (Updatable) The URI of the SCIM resource that represents the User or App who modified this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `display` - (Optional) (Updatable) The displayName of the User or App who modified this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocid` - (Optional) (Updatable) The OCID of the SCIM resource that represents the User or App who modified this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* returned: default
		* type: string
		* uniqueness: none
	* `type` - (Optional) (Updatable) The type of resource, User or App, that modified this Resource

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) The ID of the SCIM resource that represents the User or App who modified this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `idcs_last_upgraded_in_release` - (Optional) (Updatable) The release number when the resource was upgraded.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `idcs_prevented_operations` - (Optional) (Updatable) Each value of this attribute specifies an operation that only an internal client may perform on this particular resource.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `meta` - (Optional) (Updatable) A complex attribute that contains resource metadata. All sub-attributes are OPTIONAL.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Created Date, mapsTo:meta.created]]
	* type: complex
	* `created` - (Optional) (Updatable) The DateTime the Resource was added to the Service Provider

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
		* uniqueness: none
	* `last_modified` - (Optional) (Updatable) The most recent DateTime that the details of this Resource were updated at the Service Provider. If this Resource has never been modified since its initial creation, the value MUST be the same as the value of created. The attribute MUST be a DateTime.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
		* uniqueness: none
	* `location` - (Optional) (Updatable) The URI of the Resource being returned. This value MUST be the same as the Location HTTP response header.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `resource_type` - (Optional) (Updatable) Name of the resource type of the resource--for example, Users or Groups

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `version` - (Optional) (Updatable) The version of the Resource being returned. This value must be the same as the ETag HTTP response header.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `name` - (Required) (Updatable) Name of the profile

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: always
	* type: string
	* uniqueness: global
* `number_of_days_redirect_url_is_valid` - (Required) (Updatable) Number of days redirect URL is valid

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: integer
	* uniqueness: none
* `ocid` - (Optional) (Updatable) Unique Oracle Cloud Infrastructure identifier for the SCIM Resource.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: string
	* uniqueness: global
* `redirect_url` - (Required) (Updatable) This URL will be replaced in email notification sent to user. When activation email required is set to true, user is created in \"pending verification\" state, upon clicking this link user will be able to activate himself. When activation email required is set to false, user is created in \"verified\" state, this link will be used to verify user's email.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: string
	* uniqueness: none
* `resource_type_schema_version` - (Optional) (Updatable) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `schemas` - (Required) (Updatable) REQUIRED. The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. This specification defines URIs for User, Group, and a standard \"enterprise\" extension. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: true
	* returned: default
	* type: string
	* uniqueness: none
* `show_on_login_page` - (Required) (Updatable) A Boolean value that indicates whether the profile should be displayed on login page

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `tags` - (Optional) (Updatable) A list of tags on this resource.

	**SCIM++ Properties:**
	* idcsCompositeKey: [key, value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `key` - (Required) (Updatable) Key or name of the tag.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) Value of the tag.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `tenancy_ocid` - (Optional) (Updatable) Oracle Cloud Infrastructure Tenant Id (ocid) in which the resource lives.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `user_attributes` - (Optional) (Updatable) **SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none User Attributes 
	* `deletable` - (Optional) (Updatable) If this attribute can be deleted

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `fully_qualified_attribute_name` - (Optional) (Updatable) **SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none Fully Qualified Attribute Name 
	* `metadata` - (Optional) (Updatable) Metadata of the user attribute

		**Added In:** 18.1.6

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `seq_number` - (Required) (Updatable) **SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none Sequence Number for the attribute 
	* `value` - (Required) (Updatable) name of the attribute

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Self Registration Profile
	* `update` - (Defaults to 20 minutes), when updating the Self Registration Profile
	* `delete` - (Defaults to 20 minutes), when destroying the Self Registration Profile


## Import

SelfRegistrationProfiles can be imported using the `id`, e.g.

```
$ terraform import oci_identity_domains_self_registration_profile.test_self_registration_profile "idcsEndpoint/{idcsEndpoint}/selfRegistrationProfiles/{selfRegistrationProfileId}" 
```

