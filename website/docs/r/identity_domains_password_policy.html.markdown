---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_password_policy"
sidebar_current: "docs-oci-resource-identity_domains-password_policy"
description: |-
  Provides the Password Policy resource in Oracle Cloud Infrastructure Identity Domains service
---

# oci_identity_domains_password_policy
This resource provides the Password Policy resource in Oracle Cloud Infrastructure Identity Domains service.

Create a password policy.

## Example Usage

```hcl
resource "oci_identity_domains_password_policy" "test_password_policy" {
	#Required
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	name = var.password_policy_name
	schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:PasswordPolicy"]

	#Optional
	allowed_chars = var.password_policy_allowed_chars
	attribute_sets = []
	attributes = ""
	authorization = var.password_policy_authorization
	description = var.password_policy_description
	dictionary_delimiter = var.password_policy_dictionary_delimiter
	dictionary_location = var.password_policy_dictionary_location
	dictionary_word_disallowed = var.password_policy_dictionary_word_disallowed
	disallowed_chars = var.password_policy_disallowed_chars
	disallowed_substrings = var.password_policy_disallowed_substrings
	disallowed_user_attribute_values = var.password_policy_disallowed_user_attribute_values
	distinct_characters = var.password_policy_distinct_characters
	external_id = "externalId"
	first_name_disallowed = var.password_policy_first_name_disallowed
	force_password_reset = var.password_policy_force_password_reset
	groups {
		#Required
		value = oci_identity_domains_group.test_group.id
	}
	id = var.password_policy_id
	last_name_disallowed = var.password_policy_last_name_disallowed
	lockout_duration = var.password_policy_lockout_duration
	max_incorrect_attempts = var.password_policy_max_incorrect_attempts
	max_length = var.password_policy_max_length
	max_repeated_chars = var.password_policy_max_repeated_chars
	max_special_chars = var.password_policy_max_special_chars
	min_alpha_numerals = var.password_policy_min_alpha_numerals
	min_alphas = var.password_policy_min_alphas
	min_length = var.password_policy_min_length
	min_lower_case = var.password_policy_min_lower_case
	min_numerals = var.password_policy_min_numerals
	min_password_age = var.password_policy_min_password_age
	min_special_chars = var.password_policy_min_special_chars
	min_unique_chars = var.password_policy_min_unique_chars
	min_upper_case = var.password_policy_min_upper_case
	num_passwords_in_history = var.password_policy_num_passwords_in_history
	ocid = var.password_policy_ocid
	password_expire_warning = var.password_policy_password_expire_warning
	password_expires_after = var.password_policy_password_expires_after
	password_strength = var.password_policy_password_strength
	priority = var.password_policy_priority
	required_chars = var.password_policy_required_chars
	resource_type_schema_version = var.password_policy_resource_type_schema_version
	starts_with_alphabet = var.password_policy_starts_with_alphabet
	tags {
		#Required
		key = var.password_policy_tags_key
		value = var.password_policy_tags_value
	}
	user_name_disallowed = var.password_policy_user_name_disallowed
}
```

## Argument Reference

The following arguments are supported:

* `allowed_chars` - (Optional) (Updatable) A String value whose contents indicate a set of characters that can appear, in any sequence, in a password value

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
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
* `configured_password_policy_rules` - (Optional) (Updatable) List of password policy rules that have values set. This map of stringKey:stringValue pairs can be used to aid users while setting/resetting password

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [key]
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `key` - (Required) (Updatable) The specific password policy rule

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: always
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) User-friendly text that describes a specific password policy rule

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: always
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
* `description` - (Optional) (Updatable) A String that describes the password policy

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `dictionary_delimiter` - (Optional) (Updatable) A delimiter used to separate characters in the dictionary file

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `dictionary_location` - (Optional) (Updatable) A Reference value that contains the URI of a dictionary of words not allowed to appear within a password value

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `dictionary_word_disallowed` - (Optional) (Updatable) Indicates whether the password can match a dictionary word

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `disallowed_chars` - (Optional) (Updatable) A String value whose contents indicate a set of characters that cannot appear, in any sequence, in a password value

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `disallowed_substrings` - (Optional) (Updatable) A String value whose contents indicate a set of substrings that cannot appear, in any sequence, in a password value

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `disallowed_user_attribute_values` - (Optional) (Updatable) List of User attributes whose values are not allowed in the password.

	**Added In:** 2303212224

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `distinct_characters` - (Optional) (Updatable) The number of distinct characters between old password and new password

	**Added In:** 2303212224

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
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
* `external_id` - (Optional) (Updatable) An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `first_name_disallowed` - (Optional) (Updatable) Indicates a sequence of characters that match the user's first name of given name cannot be the password. Password validation against policy will be ignored if length of first name is less than or equal to 3 characters.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `force_password_reset` - (Optional) (Updatable) Indicates whether all of the users should be forced to reset their password on the next login (to comply with new password policy changes)

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: writeOnly
	* required: false
	* returned: never
	* type: boolean
	* uniqueness: none
* `groups` - (Optional) (Updatable) A list of groups that the password policy belongs to.

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - (Optional) (Updatable) Group Display Name

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) The URI of the corresponding Group resource to which the password policy belongs

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) The identifier of the group.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: always
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
* `last_name_disallowed` - (Optional) (Updatable) Indicates a sequence of characters that match the user's last name of given name cannot be the password. Password validation against policy will be ignored if length of last name is less than or equal to 3 characters.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `lockout_duration` - (Optional) (Updatable) The time period in minutes to lock out a user account when the threshold of invalid login attempts is reached. The available range is from 5 through 1440 minutes (24 hours).

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `max_incorrect_attempts` - (Optional) (Updatable) An integer that represents the maximum number of failed logins before an account is locked

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `max_length` - (Optional) (Updatable) The maximum password length (in characters). A value of 0 or no value indicates no maximum length restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `max_repeated_chars` - (Optional) (Updatable) The maximum number of repeated characters allowed in a password.  A value of 0 or no value indicates no such restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `max_special_chars` - (Optional) (Updatable) The maximum number of special characters in a password.  A value of 0 or no value indicates no maximum special characters restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
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
* `min_alpha_numerals` - (Optional) (Updatable) The minimum number of a combination of alphabetic and numeric characters in a password.  A value of 0 or no value indicates no minimum alphanumeric character restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_alphas` - (Optional) (Updatable) The minimum number of alphabetic characters in a password.  A value of 0 or no value indicates no minimum alphas restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_length` - (Optional) (Updatable) The minimum password length (in characters). A value of 0 or no value indicates no minimum length restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_lower_case` - (Optional) (Updatable) The minimum number of lowercase alphabetic characters in a password.  A value of 0 or no value indicates no minimum lowercase restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_numerals` - (Optional) (Updatable) The minimum number of numeric characters in a password.  A value of 0 or no value indicates no minimum numeric character restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_password_age` - (Optional) (Updatable) Minimum time after which the user can resubmit the reset password request

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_special_chars` - (Optional) (Updatable) The minimum number of special characters in a password. A value of 0 or no value indicates no minimum special characters restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_unique_chars` - (Optional) (Updatable) The minimum number of unique characters in a password.  A value of 0 or no value indicates no minimum unique characters restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_upper_case` - (Optional) (Updatable) The minimum number of uppercase alphabetic characters in a password. A value of 0 or no value indicates no minimum uppercase restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `name` - (Required) (Updatable) A String that is the name of the policy to display to the user. This is the only mandatory attribute for a password policy.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: true
	* returned: always
	* type: string
	* uniqueness: server
* `num_passwords_in_history` - (Optional) (Updatable) The number of passwords that will be kept in history that may not be used as a password

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
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
* `password_expire_warning` - (Optional) (Updatable) An integer indicating the number of days before which the user should be warned about password expiry.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `password_expires_after` - (Optional) (Updatable) The number of days after which the password expires automatically

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `password_strength` - (Optional) (Updatable) Indicates whether the password policy is configured as Simple, Standard, or Custom.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `priority` - (Optional) (Updatable) Password policy priority

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* idcsMinValue: 1
	* uniqueness: server
* `required_chars` - (Optional) (Updatable) A String value whose contents indicate a set of characters that must appear, in any sequence, in a password value

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
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
* `starts_with_alphabet` - (Optional) (Updatable) Indicates that the password must begin with an alphabetic character

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
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
* `user_name_disallowed` - (Optional) (Updatable) Indicates a sequence of characters that match the username cannot be the password. Password validation against policy will be ignored if length of user name is less than or equal to 3 characters.

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

* `allowed_chars` - A String value whose contents indicate a set of characters that can appear, in any sequence, in a password value

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
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
* `configured_password_policy_rules` - List of password policy rules that have values set. This map of stringKey:stringValue pairs can be used to aid users while setting/resetting password

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [key]
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `key` - The specific password policy rule

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: always
		* type: string
		* uniqueness: none
	* `value` - User-friendly text that describes a specific password policy rule

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: always
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
* `description` - A String that describes the password policy

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `dictionary_delimiter` - A delimiter used to separate characters in the dictionary file

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `dictionary_location` - A Reference value that contains the URI of a dictionary of words not allowed to appear within a password value

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `dictionary_word_disallowed` - Indicates whether the password can match a dictionary word

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `disallowed_chars` - A String value whose contents indicate a set of characters that cannot appear, in any sequence, in a password value

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `disallowed_substrings` - A String value whose contents indicate a set of substrings that cannot appear, in any sequence, in a password value

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `disallowed_user_attribute_values` - List of User attributes whose values are not allowed in the password.

	**Added In:** 2303212224

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `distinct_characters` - The number of distinct characters between old password and new password

	**Added In:** 2303212224

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
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
* `external_id` - An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `first_name_disallowed` - Indicates a sequence of characters that match the user's first name of given name cannot be the password. Password validation against policy will be ignored if length of first name is less than or equal to 3 characters.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `force_password_reset` - Indicates whether all of the users should be forced to reset their password on the next login (to comply with new password policy changes)

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: writeOnly
	* required: false
	* returned: never
	* type: boolean
	* uniqueness: none
* `groups` - A list of groups that the password policy belongs to.

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - Group Display Name

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - The URI of the corresponding Group resource to which the password policy belongs

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - The identifier of the group.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: always
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
* `last_name_disallowed` - Indicates a sequence of characters that match the user's last name of given name cannot be the password. Password validation against policy will be ignored if length of last name is less than or equal to 3 characters.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `lockout_duration` - The time period in minutes to lock out a user account when the threshold of invalid login attempts is reached. The available range is from 5 through 1440 minutes (24 hours).

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `max_incorrect_attempts` - An integer that represents the maximum number of failed logins before an account is locked

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `max_length` - The maximum password length (in characters). A value of 0 or no value indicates no maximum length restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `max_repeated_chars` - The maximum number of repeated characters allowed in a password.  A value of 0 or no value indicates no such restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `max_special_chars` - The maximum number of special characters in a password.  A value of 0 or no value indicates no maximum special characters restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
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
* `min_alpha_numerals` - The minimum number of a combination of alphabetic and numeric characters in a password.  A value of 0 or no value indicates no minimum alphanumeric character restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_alphas` - The minimum number of alphabetic characters in a password.  A value of 0 or no value indicates no minimum alphas restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_length` - The minimum password length (in characters). A value of 0 or no value indicates no minimum length restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_lower_case` - The minimum number of lowercase alphabetic characters in a password.  A value of 0 or no value indicates no minimum lowercase restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_numerals` - The minimum number of numeric characters in a password.  A value of 0 or no value indicates no minimum numeric character restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_password_age` - Minimum time after which the user can resubmit the reset password request

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_special_chars` - The minimum number of special characters in a password. A value of 0 or no value indicates no minimum special characters restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_unique_chars` - The minimum number of unique characters in a password.  A value of 0 or no value indicates no minimum unique characters restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `min_upper_case` - The minimum number of uppercase alphabetic characters in a password. A value of 0 or no value indicates no minimum uppercase restriction.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `name` - A String that is the name of the policy to display to the user. This is the only mandatory attribute for a password policy.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: true
	* returned: always
	* type: string
	* uniqueness: server
* `num_passwords_in_history` - The number of passwords that will be kept in history that may not be used as a password

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
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
* `password_expire_warning` - An integer indicating the number of days before which the user should be warned about password expiry.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `password_expires_after` - The number of days after which the password expires automatically

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `password_strength` - Indicates whether the password policy is configured as Simple, Standard, or Custom.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `priority` - Password policy priority

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* idcsMinValue: 1
	* uniqueness: server
* `required_chars` - A String value whose contents indicate a set of characters that must appear, in any sequence, in a password value

	**SCIM++ Properties:**
	* caseExact: false
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
* `starts_with_alphabet` - Indicates that the password must begin with an alphabetic character

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
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
* `user_name_disallowed` - Indicates a sequence of characters that match the username cannot be the password. Password validation against policy will be ignored if length of user name is less than or equal to 3 characters.

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
	* `create` - (Defaults to 20 minutes), when creating the Password Policy
	* `update` - (Defaults to 20 minutes), when updating the Password Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Password Policy


## Import

PasswordPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_identity_domains_password_policy.test_password_policy "idcsEndpoint/{idcsEndpoint}/passwordPolicies/{passwordPolicyId}" 
```

