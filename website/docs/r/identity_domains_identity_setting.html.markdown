---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_identity_setting"
sidebar_current: "docs-oci-resource-identity_domains-identity_setting"
description: |-
  Provides the Identity Setting resource in Oracle Cloud Infrastructure Identity Domains service
---

# oci_identity_domains_identity_setting
This resource provides the Identity Setting resource in Oracle Cloud Infrastructure Identity Domains service.

Replace an Identity setting.

## Example Usage

```hcl
resource "oci_identity_domains_identity_setting" "test_identity_setting" {
	#Required
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	identity_setting_id = oci_identity_domains_identity_setting.test_identity_setting.id
	schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:IdentitySettings"]

	#Optional
	posix_gid {

		#Optional
		manual_assignment_ends_at = var.identity_setting_posix_gid_manual_assignment_ends_at
		manual_assignment_starts_from = var.identity_setting_posix_gid_manual_assignment_starts_from
	}
	posix_uid {

		#Optional
		manual_assignment_ends_at = var.identity_setting_posix_uid_manual_assignment_ends_at
		manual_assignment_starts_from = var.identity_setting_posix_uid_manual_assignment_starts_from
	}
	attribute_sets = ["all"]
	attributes = ""
	authorization = var.identity_setting_authorization
	emit_locked_message_when_user_is_locked = var.identity_setting_emit_locked_message_when_user_is_locked
	external_id = "externalId"
	id = var.identity_setting_id
	my_profile {

		#Optional
		allow_end_users_to_change_their_password = true
		allow_end_users_to_link_their_support_account = true
		allow_end_users_to_manage_their_capabilities = true
		allow_end_users_to_update_their_security_settings = true
	}
	ocid = var.identity_setting_ocid
	primary_email_required = var.identity_setting_primary_email_required
	resource_type_schema_version = var.identity_setting_resource_type_schema_version
	return_inactive_over_locked_message = var.identity_setting_return_inactive_over_locked_message
	tags {
		#Required
		key = var.identity_setting_tags_key
		value = var.identity_setting_tags_value
	}
	tokens {
		#Required
		type = var.identity_setting_tokens_type

		#Optional
		expires_after = var.identity_setting_tokens_expires_after
	}
	user_allowed_to_set_recovery_email = var.identity_setting_user_allowed_to_set_recovery_email
}
```

## Argument Reference

The following arguments are supported:

* `posix_gid` - (Optional) (Updatable) A list of Posix Gid settings.

	**Added In:** 17.4.6

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* required: false
	* type: complex
	* uniqueness: none
	* `manual_assignment_ends_at` - (Optional) (Updatable) The number at which the Posix Gid Manual assignment ends.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `manual_assignment_starts_from` - (Optional) (Updatable) The number from which the Posix Gid Manual assignment starts.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
* `posix_uid` - (Optional) (Updatable) A list of Posix Uid settings.

	**Added In:** 17.4.6

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* required: false
	* type: complex
	* uniqueness: none
	* `manual_assignment_ends_at` - (Optional) (Updatable) The number at which the Posix Uid Manual assignment ends.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `manual_assignment_starts_from` - (Optional) (Updatable) The number from which the Posix Uid Manual assignment starts.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
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
* `emit_locked_message_when_user_is_locked` - (Optional) (Updatable) Indicates whether to show the 'user-is-locked' message during authentication if the user is already locked. The default value is false, which tells the system to show a generic 'authentication-failure' message. This is the most secure behavior. If the option is set to true, the system shows a more detailed 'error-message' that says the user is locked. This is more helpful but is less secure, for example, because the difference in error-messages could be used to determine which usernames exist and which do not.

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `external_id` - (Optional) (Updatable) An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
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
* `identity_setting_id` - (Required) ID of the resource
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
* `my_profile` - (Optional) (Updatable) Whether to allow users to update their own profile.

	**Added In:** 2207040824

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* required: false
	* type: complex
	* uniqueness: none
	* `allow_end_users_to_change_their_password` - (Optional) (Updatable) Whether to allow users to change their own password.

		**Added In:** 2207040824

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `allow_end_users_to_link_their_support_account` - (Optional) (Updatable) Whether to allow users to link or unlink their support accounts.

		**Added In:** 2207040824

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `allow_end_users_to_manage_their_capabilities` - (Optional) (Updatable) Whether to allow users to update their capabilities.

		**Added In:** 2207040824

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `allow_end_users_to_update_their_security_settings` - (Optional) (Updatable) Whether to allow users to update their security settings.

		**Added In:** 2207040824

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
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
* `primary_email_required` - (Optional) (Updatable) Indicates whether the primary email is required.

	**Added In:** 19.1.4

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `remove_invalid_emails` - (Optional) (Updatable) Indicates whether to remove non-RFC5322 compliant emails before creating a user.

	**Added In:** 2106170416

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `resource_type_schema_version` - (Optional) (Updatable) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `return_inactive_over_locked_message` - (Optional) (Updatable) **Added In:** 2302092332

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none Determines the PasswordAuthenticator API response when a user is both 'locked' and 'inactive.' If false (default), a 'locked' message is shown. If true, an 'inactive' message is shown. 
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
* `tokens` - (Optional) (Updatable) A list of tokens and their expiry length.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [type]
	* multiValued: true
	* required: false
	* type: complex
	* uniqueness: none
	* `expires_after` - (Optional) (Updatable) Indicates the number of minutes after which the token expires automatically.

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `type` - (Required) (Updatable) The token type.

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `user_allowed_to_set_recovery_email` - (Optional) (Updatable) Indicates whether a user is allowed to change their own recovery email.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `posix_gid` - A list of Posix Gid settings.

	**Added In:** 17.4.6

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* required: false
	* type: complex
	* uniqueness: none
	* `manual_assignment_ends_at` - The number at which the Posix Gid Manual assignment ends.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `manual_assignment_starts_from` - The number from which the Posix Gid Manual assignment starts.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
* `posix_uid` - A list of Posix Uid settings.

	**Added In:** 17.4.6

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* required: false
	* type: complex
	* uniqueness: none
	* `manual_assignment_ends_at` - The number at which the Posix Uid Manual assignment ends.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `manual_assignment_starts_from` - The number from which the Posix Uid Manual assignment starts.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
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
* `emit_locked_message_when_user_is_locked` - Indicates whether to show the 'user-is-locked' message during authentication if the user is already locked. The default value is false, which tells the system to show a generic 'authentication-failure' message. This is the most secure behavior. If the option is set to true, the system shows a more detailed 'error-message' that says the user is locked. This is more helpful but is less secure, for example, because the difference in error-messages could be used to determine which usernames exist and which do not.

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `external_id` - An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.

	**SCIM++ Properties:**
	* caseExact: false
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
* `my_profile` - Whether to allow users to update their own profile.

	**Added In:** 2207040824

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* required: false
	* type: complex
	* uniqueness: none
	* `allow_end_users_to_change_their_password` - Whether to allow users to change their own password.

		**Added In:** 2207040824

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `allow_end_users_to_link_their_support_account` - Whether to allow users to link or unlink their support accounts.

		**Added In:** 2207040824

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `allow_end_users_to_manage_their_capabilities` - Whether to allow users to update their capabilities.

		**Added In:** 2207040824

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `allow_end_users_to_update_their_security_settings` - Whether to allow users to update their security settings.

		**Added In:** 2207040824

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
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
* `primary_email_required` - Indicates whether the primary email is required.

	**Added In:** 19.1.4

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `remove_invalid_emails` - Indicates whether to remove non-RFC5322 compliant emails before creating a user.

	**Added In:** 2106170416

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `return_inactive_over_locked_message` - **Added In:** 2302092332

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none Determines the PasswordAuthenticator API response when a user is both 'locked' and 'inactive.' If false (default), a 'locked' message is shown. If true, an 'inactive' message is shown. 
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
* `tokens` - A list of tokens and their expiry length.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [type]
	* multiValued: true
	* required: false
	* type: complex
	* uniqueness: none
	* `expires_after` - Indicates the number of minutes after which the token expires automatically.

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `type` - The token type.

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `user_allowed_to_set_recovery_email` - Indicates whether a user is allowed to change their own recovery email.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Identity Setting
	* `update` - (Defaults to 20 minutes), when updating the Identity Setting
	* `delete` - (Defaults to 20 minutes), when destroying the Identity Setting


## Import

IdentitySettings can be imported using the `id`, e.g.

```
$ terraform import oci_identity_domains_identity_setting.test_identity_setting "idcsEndpoint/{idcsEndpoint}/identitySettings/{identitySettingId}" 
```

