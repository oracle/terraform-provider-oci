---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_social_identity_provider"
sidebar_current: "docs-oci-datasource-identity_domains-social_identity_provider"
description: |-
  Provides details about a specific Social Identity Provider in Oracle Cloud Infrastructure Identity Domains service
---

# Data Source: oci_identity_domains_social_identity_provider
This data source provides details about a specific Social Identity Provider resource in Oracle Cloud Infrastructure Identity Domains service.

Get a Social Identity Provider

## Example Usage

```hcl
data "oci_identity_domains_social_identity_provider" "test_social_identity_provider" {
	#Required
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	social_identity_provider_id = oci_identity_identity_provider.test_identity_provider.id

	#Optional
	authorization = var.social_identity_provider_authorization
	resource_type_schema_version = var.social_identity_provider_resource_type_schema_version
}
```

## Argument Reference

The following arguments are supported:

* `authorization` - (Optional) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
* `idcs_endpoint` - (Required) The basic endpoint for the identity domain
* `resource_type_schema_version` - (Optional) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `social_identity_provider_id` - (Required) ID of the resource


## Attributes Reference

The following attributes are exported:

* `access_token_url` - Social IDP Access token URL

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `account_linking_enabled` - Whether account linking is enabled

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `admin_scope` - Admin scope to request

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `authz_url` - Social IDP Authorization URL

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `auto_redirect_enabled` - Whether social auto redirect is enabled. The IDP policy should be configured with only one Social IDP, and without username/password selected.

	**Added In:** 2310202314

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `client_credential_in_payload` - Whether the client credential is contained in payload

	**Added In:** 18.4.2

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `clock_skew_in_seconds` - Social IDP allowed clock skew time

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
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
* `consumer_key` - Social IDP Client Application Client ID

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: string
	* uniqueness: none
* `consumer_secret` - Social IDP Client Application Client Secret

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* idcsSensitive: encrypt
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
* `description` - Social IDP description

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `discovery_url` - Discovery URL

	**Added In:** 17.4.6

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
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
* `enabled` - Whether the IDP is enabled or not

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `icon_url` - ICON URL for social idp

	**Added In:** 17.4.6

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
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
* `id_attribute` - Id attribute used for account linking

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: immutable
	* required: false
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
	* `_ref` - The URI of the SCIM resource that represents the User or App who created this Resource

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
	* `_ref` - The URI of the SCIM resource that represents the User or App who modified this Resource

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
* `jit_prov_assigned_groups` - Lists the groups each social JIT-provisioned user is a member. Just-in-Time user-provisioning applies this static list when jitProvGroupStaticListEnabled:true.

	**Added In:** 2309290043

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `_ref` - Group URI

		**Added In:** 2309290043

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `display` - A human readable name, primarily used for display purposes. READ-ONLY.

		**Added In:** 2309290043

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `value` - Group identifier

		**Added In:** 2309290043

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `jit_prov_group_static_list_enabled` - Set to true to indicate Social JIT User Provisioning Groups should be assigned from a static list

	**Added In:** 2309290043

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
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
* `name` - Social provider name

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: always
	* type: string
	* uniqueness: global
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
* `profile_url` - Social IDP User profile URL

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `redirect_url` - redirect URL for social idp

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `refresh_token_url` - Social IDP Refresh token URL

	**Added In:** 19.1.4

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `registration_enabled` - Whether registration is enabled

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `relay_idp_param_mappings` - Relay Param variable for Social IDP

	**Added In:** 2305190132

	**SCIM++ Properties:**
	* idcsCompositeKey: [relayParamKey]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `relay_param_key` - Key or name of the relayParam.

		**Added In:** 2305190132

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `relay_param_value` - Value of the relayParam (if defined)

		**Added In:** 2305190132

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
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
* `scope` - Scope to request

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `service_provider_name` - Service Provider Name

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: true
	* returned: default
	* type: string
	* uniqueness: none
* `show_on_login` - Whether show on login

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `social_jit_provisioning_enabled` - Whether Social JIT Provisioning is enabled

	**Added In:** 2307282043

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `status` - Status

	**Added In:** 17.4.6

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
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

