---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_user"
sidebar_current: "docs-oci-resource-identity_domains-user"
description: |-
  Provides the User resource in Oracle Cloud Infrastructure Identity Domains service
---

# oci_identity_domains_user
This resource provides the User resource in Oracle Cloud Infrastructure Identity Domains service.

Create a user.

## Example Usage

```hcl
resource "oci_identity_domains_user" "test_user" {
	#Required
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	name {
		#Required
		family_name = var.user_name_family_name

		#Optional
		formatted = var.user_name_formatted
		given_name = var.user_name_given_name
		honorific_prefix = var.user_name_honorific_prefix
		honorific_suffix = var.user_name_honorific_suffix
		middle_name = var.user_name_middle_name
	}
	schemas = ["urn:ietf:params:scim:schemas:core:2.0:User"]
	user_name = "userName"
	/* Note: In most cases, a primary email is REQUIRED to create a user. Otherwise you might get a 400 error. Please see "emails" block below. */

	#Optional
	active = var.user_active
	addresses {
		#Required
		type = var.user_addresses_type

		#Optional
		country = var.user_addresses_country
		formatted = var.user_addresses_formatted
		locality = var.user_addresses_locality
		postal_code = var.user_addresses_postal_code
		primary = var.user_addresses_primary
		region = var.user_addresses_region
		street_address = var.user_addresses_street_address
	}
	attribute_sets = []
	attributes = ""
	authorization = var.user_authorization
	description = var.user_description
	display_name = var.user_display_name

	/* One and only one "emails" block needs to have "primary" set to true */
	emails {
		#Required
		type = var.user_emails_type
		value = var.user_emails_value

		#Optional
		primary = true
		secondary = var.user_emails_secondary
		verified = var.user_emails_verified
	}
	/* Note:
      If a new user is created without a recovery email being set, we automatically add one using the primary email value,
      to ensure the account can be recovered (when account recovery feature is enabled in the current domain).
      So it is recommended to set an email of type "recovery" like below. If not, it is expected to see an update about 
      recovery email when plan/apply after creation.
    */
	emails {
		#Required
		type = "recovery"
		value = var.user_emails_value
	}
	entitlements {
		#Required
		type = var.user_entitlements_type
		value = var.user_entitlements_value

		#Optional
		display = var.user_entitlements_display
		primary = var.user_entitlements_primary
	}
	external_id = "externalId"
	id = var.user_id
	ims {
		#Required
		type = var.user_ims_type
		value = var.user_ims_value

		#Optional
		display = var.user_ims_display
		primary = var.user_ims_primary
	}
	locale = var.user_locale
	nick_name = var.user_nick_name
	ocid = var.user_ocid
	password = var.user_password
	phone_numbers {
		#Required
		type = var.user_phone_numbers_type
		value = var.user_phone_numbers_value

		#Optional
		primary = var.user_phone_numbers_primary
	}
	photos {
		#Required
		type = var.user_photos_type
		value = var.user_photos_value

		#Optional
		display = var.user_photos_display
		primary = var.user_photos_primary
	}
	preferred_language = var.user_preferred_language
	profile_url = var.user_profile_url
	resource_type_schema_version = var.user_resource_type_schema_version
	roles {
		#Required
		type = var.user_roles_type
		value = var.user_roles_value

		#Optional
		display = var.user_roles_display
		primary = var.user_roles_primary
	}
	tags {
		#Required
		key = var.user_tags_key
		value = var.user_tags_value
	}
	timezone = var.user_timezone
	title = var.user_title
	urnietfparamsscimschemasextensionenterprise20user {

		#Optional
		cost_center = var.user_urnietfparamsscimschemasextensionenterprise20user_cost_center
		department = var.user_urnietfparamsscimschemasextensionenterprise20user_department
		division = var.user_urnietfparamsscimschemasextensionenterprise20user_division
		employee_number = var.user_urnietfparamsscimschemasextensionenterprise20user_employee_number
		manager {

			#Optional
			value = var.user_urnietfparamsscimschemasextensionenterprise20user_manager_value
		}
		organization = var.user_urnietfparamsscimschemasextensionenterprise20user_organization
	}
	urnietfparamsscimschemasoracleidcsextension_oci_tags {

		#Optional
		defined_tags {
			#Required
			key = var.user_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_key
			namespace = var.user_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_namespace
			value = var.user_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_value
		}
		freeform_tags {
			#Required
			key = var.user_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_key
			value = var.user_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_value
		}
	}
	urnietfparamsscimschemasoracleidcsextensionadaptive_user {

		#Optional
		risk_level = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_level
		risk_scores {
			#Required
			last_update_timestamp = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_last_update_timestamp
			risk_level = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_risk_level
			score = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_score
			value = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_value

			#Optional
			source = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_source
			status = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_status
		}
	}
	urnietfparamsscimschemasoracleidcsextensioncapabilities_user {

		#Optional
		can_use_api_keys = true
		can_use_auth_tokens = true
		can_use_console_password = true
		can_use_customer_secret_keys = true
		can_use_db_credentials = true
		can_use_oauth2client_credentials = true
		can_use_smtp_credentials = true
	}
	urnietfparamsscimschemasoracleidcsextensiondb_credentials_user {

		#Optional
		db_user_name = "dbUserName"
	}
	urnietfparamsscimschemasoracleidcsextensionkerberos_user_user {

		#Optional
		realm_users {
			#Required
			value = var.user_urnietfparamsscimschemasoracleidcsextensionkerberos_user_user_realm_users_value

			#Optional
			principal_name = var.user_urnietfparamsscimschemasoracleidcsextensionkerberos_user_user_realm_users_principal_name
			realm_name = var.user_urnietfparamsscimschemasoracleidcsextensionkerberos_user_user_realm_users_realm_name
		}
	}
	urnietfparamsscimschemasoracleidcsextensionmfa_user {

		#Optional
		bypass_codes {
			#Required
			value = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_bypass_codes_value
		}
		devices {
			#Required
			value = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_value

			#Optional
			authentication_method = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_authentication_method
			display = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_display
			factor_status = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_factor_status
			factor_type = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_factor_type
			last_sync_time = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_last_sync_time
			status = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_status
			third_party_vendor_name = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_third_party_vendor_name
		}
		login_attempts = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_login_attempts
		mfa_enabled_on = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_mfa_enabled_on
		mfa_ignored_apps = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_mfa_ignored_apps
		mfa_status = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_mfa_status
		preferred_authentication_factor = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_preferred_authentication_factor
		preferred_authentication_method = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_preferred_authentication_method
		preferred_device {
			#Required
			value = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_preferred_device_value

			#Optional
			display = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_preferred_device_display
		}
		preferred_third_party_vendor = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_preferred_third_party_vendor
		trusted_user_agents {
			#Required
			value = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_trusted_user_agents_value

			#Optional
			display = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_trusted_user_agents_display
		}
	}
	urnietfparamsscimschemasoracleidcsextensionpasswordless_user {

		#Optional
		factor_identifier {
			#Required
			value = var.user_urnietfparamsscimschemasoracleidcsextensionpasswordless_user_factor_identifier_value

			#Optional
			display = var.user_urnietfparamsscimschemasoracleidcsextensionpasswordless_user_factor_identifier_display
		}
		factor_method = var.user_urnietfparamsscimschemasoracleidcsextensionpasswordless_user_factor_method
		factor_type = var.user_urnietfparamsscimschemasoracleidcsextensionpasswordless_user_factor_type
	}
	urnietfparamsscimschemasoracleidcsextensionposix_user {

		#Optional
		gecos = var.user_urnietfparamsscimschemasoracleidcsextensionposix_user_gecos
		gid_number = var.user_urnietfparamsscimschemasoracleidcsextensionposix_user_gid_number
		home_directory = var.user_urnietfparamsscimschemasoracleidcsextensionposix_user_home_directory
		login_shell = var.user_urnietfparamsscimschemasoracleidcsextensionposix_user_login_shell
		uid_number = var.user_urnietfparamsscimschemasoracleidcsextensionposix_user_uid_number
	}
	urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user {

		#Optional
		sec_questions {
			#Required
			answer = var.user_urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user_sec_questions_answer
			value = var.user_urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user_sec_questions_value

			#Optional
			hint_text = var.user_urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user_sec_questions_hint_text
		}
	}
	urnietfparamsscimschemasoracleidcsextensionself_change_user {

		#Optional
		allow_self_change = var.user_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change
	}
	urnietfparamsscimschemasoracleidcsextensionself_registration_user {
		#Required
		self_registration_profile {
			#Required
			value = var.user_urnietfparamsscimschemasoracleidcsextensionself_registration_user_self_registration_profile_value

			#Optional
			display = var.user_urnietfparamsscimschemasoracleidcsextensionself_registration_user_self_registration_profile_display
		}

		#Optional
		consent_granted = var.user_urnietfparamsscimschemasoracleidcsextensionself_registration_user_consent_granted
		user_token = var.user_urnietfparamsscimschemasoracleidcsextensionself_registration_user_user_token
	}
	urnietfparamsscimschemasoracleidcsextensionsff_user {

		#Optional
		sff_auth_keys = var.user_urnietfparamsscimschemasoracleidcsextensionsff_user_sff_auth_keys
	}
	urnietfparamsscimschemasoracleidcsextensionsocial_account_user {

		#Optional
		social_accounts {
			#Required
			value = var.user_urnietfparamsscimschemasoracleidcsextensionsocial_account_user_social_accounts_value

			#Optional
			display = var.user_urnietfparamsscimschemasoracleidcsextensionsocial_account_user_social_accounts_display
		}
	}
	urnietfparamsscimschemasoracleidcsextensionterms_of_use_user {

		#Optional
		terms_of_use_consents {
			#Required
			value = var.user_urnietfparamsscimschemasoracleidcsextensionterms_of_use_user_terms_of_use_consents_value
		}
	}
	urnietfparamsscimschemasoracleidcsextensionuser_state_user {

		#Optional
		locked {

			#Optional
			expired = var.user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_locked_expired
			lock_date = var.user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_locked_lock_date
			on = var.user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_locked_on
			reason = var.user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_locked_reason
		}
		max_concurrent_sessions = var.user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_max_concurrent_sessions
		recovery_locked {

			#Optional
			lock_date = var.user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_recovery_locked_lock_date
			on = var.user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_recovery_locked_on
		}
	}
	urnietfparamsscimschemasoracleidcsextensionuser_user {

		#Optional
		user_provider = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_user_provider
		account_recovery_required = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_account_recovery_required
		bypass_notification = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_bypass_notification
		creation_mechanism = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_creation_mechanism
		delegated_authentication_target_app {
			#Required
			type = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_delegated_authentication_target_app_type
			value = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_delegated_authentication_target_app_value

			#Optional
			display = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_delegated_authentication_target_app_display
		}
		do_not_show_getting_started = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_do_not_show_getting_started
		is_authentication_delegated = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_is_authentication_delegated
		is_federated_user = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_is_federated_user
		is_group_membership_normalized = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_is_group_membership_normalized
		is_group_membership_synced_to_users_groups = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_is_group_membership_synced_to_users_groups
		notification_email_template_id = oci_identity_domains_notification_email_template.test_notification_email_template.id
		preferred_ui_landing_page = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_preferred_ui_landing_page
		synced_from_app {
			#Required
			type = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_synced_from_app_type
			value = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_synced_from_app_value

			#Optional
			display = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_synced_from_app_display
		}
		user_flow_controlled_by_external_client = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_user_flow_controlled_by_external_client
	}
	user_type = var.user_user_type
	x509certificates {
		#Required
		value = var.user_x509certificates_value

		#Optional
		display = var.user_x509certificates_display
		primary = var.user_x509certificates_primary
		type = var.user_x509certificates_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `active` - (Optional) (Updatable) User status

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Active
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Active]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `addresses` - (Optional) (Updatable) A physical mailing address for this User, as described in (address Element). Canonical Type Values of work, home, and other. The value attribute is a complex type with the following sub-attributes.

	**SCIM++ Properties:**
	* idcsCompositeKey: [type]
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Work Address Street, deprecatedColumnHeaderName:Work Street Address, mapsTo:addresses[work].streetAddress], [columnHeaderName:Work Address Locality, deprecatedColumnHeaderName:Work City, mapsTo:addresses[work].locality], [columnHeaderName:Work Address Region, deprecatedColumnHeaderName:Work State, mapsTo:addresses[work].region], [columnHeaderName:Work Address Postal Code, deprecatedColumnHeaderName:Work Postal Code, mapsTo:addresses[work].postalCode], [columnHeaderName:Work Address Country, deprecatedColumnHeaderName:Work Country, mapsTo:addresses[work].country], [columnHeaderName:Work Address Formatted, mapsTo:addresses[work].formatted], [columnHeaderName:Home Address Formatted, mapsTo:addresses[home].formatted], [columnHeaderName:Other Address Formatted, mapsTo:addresses[other].formatted], [columnHeaderName:Home Address Street, mapsTo:addresses[home].streetAddress], [columnHeaderName:Other Address Street, mapsTo:addresses[other].streetAddress], [columnHeaderName:Home Address Locality, mapsTo:addresses[home].locality], [columnHeaderName:Other Address Locality, mapsTo:addresses[other].locality], [columnHeaderName:Home Address Region, mapsTo:addresses[home].region], [columnHeaderName:Other Address Region, mapsTo:addresses[other].region], [columnHeaderName:Home Address Country, mapsTo:addresses[home].country], [columnHeaderName:Other Address Country, mapsTo:addresses[other].country], [columnHeaderName:Home Address Postal Code, mapsTo:addresses[home].postalCode], [columnHeaderName:Other Address Postal Code, mapsTo:addresses[other].postalCode], [columnHeaderName:Primary Address Type, mapsTo:addresses[$(type)].primary]]
	* idcsPii: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `country` - (Optional) (Updatable) The country name component.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCanonicalValueSourceFilter: attrName eq "countries" and attrValues.value eq "upper($(country))"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `formatted` - (Optional) (Updatable) The full mailing address, formatted for display or use with a mailing label. This attribute MAY contain newlines.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `locality` - (Optional) (Updatable) The city or locality component.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `postal_code` - (Optional) (Updatable) The zipcode or postal code component.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - (Optional) (Updatable) A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `region` - (Optional) (Updatable) The state or region component.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `street_address` - (Optional) (Updatable) The full street address component, which may include house number, street name, PO BOX, and multi-line extended street address information. This attribute MAY contain newlines.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `type` - (Required) (Updatable) A label indicating the attribute's function; e.g., 'work' or 'home'.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
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
* `description` - (Optional) (Updatable) Description of the user

	**Added In:** 2012271618

	**SCIM++ Properties:**
	* caseExact: false
	* idcsPii: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `display_name` - (Optional) (Updatable) Display name

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Display Name
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Display Name]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
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
* `emails` - (Optional) (Updatable) A complex attribute representing emails

	**SCIM++ Properties:**
	* idcsCompositeKey: [value, type]
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Work Email, mapsTo:emails[work].value], [columnHeaderName:Home Email, mapsTo:emails[home].value], [columnHeaderName:Primary Email Type, mapsTo:emails[$(type)].primary], [columnHeaderName:Other Email, mapsTo:emails[other].value], [columnHeaderName:Recovery Email, mapsTo:emails[recovery].value], [columnHeaderName:Work Email Verified, mapsTo:emails[work].verified], [columnHeaderName:Home Email Verified, mapsTo:emails[home].verified], [columnHeaderName:Other Email Verified, mapsTo:emails[other].verified], [columnHeaderName:Recovery Email Verified, mapsTo:emails[recovery].verified]]
	* idcsPii: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `pending_verification_data` - (Optional) (Updatable) Pending e-mail address verification

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - (Optional) (Updatable) A Boolean value that indicates whether the email address is the primary email address. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `secondary` - (Optional) (Updatable) A Boolean value that indicates whether the email address is the secondary email address. The secondary attribute value 'true' MUST appear no more than once.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - (Required) (Updatable) Type of email address

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) Email address

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `verified` - (Optional) (Updatable) A Boolean value that indicates whether or not the e-mail address is verified

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
* `entitlements` - (Optional) (Updatable) A list of entitlements for the User that represent a thing the User has.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value, type]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - (Optional) (Updatable) A human readable name, primarily used for display purposes.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - (Optional) (Updatable) A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - (Required) (Updatable) A label indicating the attribute's function.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) The value of an entitlement.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `external_id` - (Optional) (Updatable) An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeNameMappings: [[columnHeaderName:External Id]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `groups` - (Optional) (Updatable) A list of groups that the user belongs to, either thorough direct membership, nested groups, or dynamically calculated

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `date_added` - (Optional) (Updatable) Date when the member is Added to the group

		**Added In:** 2105200541

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
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
	* `external_id` - (Optional) (Updatable) An identifier for the Resource as defined by the Service Consumer. READ-ONLY.

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `membership_ocid` - (Optional) (Updatable) The membership OCID.

		**Added In:** 2103141444

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `non_unique_display` - (Optional) (Updatable) A human readable name for Group as defined by the Service Consumer. READ-ONLY.

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocid` - (Optional) (Updatable) The OCID of the User's group.

		**Added In:** 2102181953

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) The URI of the corresponding Group resource to which the user belongs

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - (Optional) (Updatable) A label indicating the attribute's function; e.g., 'direct' or 'indirect'.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) The identifier of the User's group.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
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
* `ims` - (Optional) (Updatable) User's instant messaging addresses

	**SCIM++ Properties:**
	* idcsCompositeKey: [value, type]
	* idcsPii: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - (Optional) (Updatable) A human-readable name, primarily used for display purposes

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - (Optional) (Updatable) A Boolean value that indicates the 'primary' or preferred attribute value for this attribute--for example, the preferred messenger or primary messenger. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - (Required) (Updatable) A label that indicates the attribute's function--for example, 'aim', 'gtalk', or 'mobile'

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) User's instant messaging address

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `locale` - (Optional) (Updatable) Used to indicate the User's default location for purposes of localizing items such as currency, date and time format, numerical representations, and so on.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Locale
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Locale]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
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
* `name` - (Required) (Updatable) A complex attribute that contains attributes representing the name

	**SCIM++ Properties:**
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Formatted Name, mapsTo:name.formatted], [columnHeaderName:Honorific Prefix, mapsTo:name.honorificPrefix], [columnHeaderName:First Name, mapsTo:name.givenName], [columnHeaderName:Middle Name, mapsTo:name.middleName], [columnHeaderName:Last Name, mapsTo:name.familyName], [columnHeaderName:Honorific Suffix, mapsTo:name.honorificSuffix]]
	* idcsPii: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `family_name` - (Required) (Updatable) Last name

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Last Name
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `formatted` - (Optional) (Updatable) Full name

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `given_name` - (Optional) (Updatable) First name

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: First Name
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `honorific_prefix` - (Optional) (Updatable) Prefix

		**SCIM++ Properties:**
		* idcsCsvAttributeName: Honorific Prefix
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `honorific_suffix` - (Optional) (Updatable) Suffix

		**SCIM++ Properties:**
		* idcsCsvAttributeName: Honorific Suffix
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `middle_name` - (Optional) (Updatable) Middle name

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Middle Name
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `nick_name` - (Optional) (Updatable) Nick name

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Nick Name
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Nick Name]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
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
* `password` - (Optional) (Updatable) Password attribute. Max length for password is controlled via Password Policy.

	**SCIM++ Properties:**
	* idcsCsvAttributeName: Password
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Password]]
	* idcsPii: true
	* idcsSearchable: false
	* idcsSensitive: hash
	* multiValued: false
	* mutability: writeOnly
	* required: false
	* returned: never
	* type: string
	* uniqueness: none
* `phone_numbers` - (Optional) (Updatable) Phone numbers

	**SCIM++ Properties:**
	* idcsCompositeKey: [value, type]
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Work Phone, mapsTo:phoneNumbers[work].value], [columnHeaderName:Mobile No, mapsTo:phoneNumbers[mobile].value], [columnHeaderName:Home Phone, mapsTo:phoneNumbers[home].value], [columnHeaderName:Fax, mapsTo:phoneNumbers[fax].value], [columnHeaderName:Pager, mapsTo:phoneNumbers[pager].value], [columnHeaderName:Other Phone, mapsTo:phoneNumbers[other].value], [columnHeaderName:Recovery Phone, mapsTo:phoneNumbers[recovery].value], [columnHeaderName:Primary Phone Type, mapsTo:phoneNumbers[$(type)].primary]]
	* idcsPii: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - (Optional) (Updatable) A human-readable name, primarily used for display purposes. READ ONLY

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - (Optional) (Updatable) A Boolean value that indicates the 'primary' or preferred attribute value for this attribute--for example, the preferred phone number or primary phone number. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - (Required) (Updatable) A label that indicates the attribute's function- for example, 'work', 'home', or 'mobile'

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) User's phone number

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `verified` - (Optional) (Updatable) A Boolean value that indicates if the phone number is verified.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
* `photos` - (Optional) (Updatable) URLs of photos for the User

	**SCIM++ Properties:**
	* idcsCompositeKey: [value, type]
	* idcsPii: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - (Optional) (Updatable) A human readable name, primarily used for display purposes.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - (Optional) (Updatable) A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g., the preferred photo or thumbnail. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - (Required) (Updatable) A label indicating the attribute's function; e.g., 'photo' or 'thumbnail'.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) URL of a photo for the User

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: reference
		* uniqueness: none
* `preferred_language` - (Optional) (Updatable) User's preferred written or spoken language used for localized user interfaces

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Preferred Language
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Preferred Language]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `profile_url` - (Optional) (Updatable) A fully-qualified URL to a page representing the User's online profile

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Profile URL
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Profile Url]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: reference
	* uniqueness: none
* `resource_type_schema_version` - (Optional) (Updatable) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `roles` - (Optional) (Updatable) A list of roles for the User that collectively represent who the User is; e.g., 'Student', 'Faculty'.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value, type]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - (Optional) (Updatable) A human readable name, primarily used for display purposes.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - (Optional) (Updatable) A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - (Required) (Updatable) A label indicating the attribute's function.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) The value of a role.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
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
* `timezone` - (Optional) (Updatable) User's timezone

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCanonicalValueSourceFilter: attrName eq "timezones" and attrValues.value eq "$(timezone)"
	* idcsCanonicalValueSourceResourceType: AllowedValue
	* idcsCsvAttributeName: TimeZone
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Time Zone, deprecatedColumnHeaderName:TimeZone]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `title` - (Optional) (Updatable) Title

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Title
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Title]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `urnietfparamsscimschemasextensionenterprise20user` - (Optional) (Updatable) Enterprise User
	* `cost_center` - (Optional) (Updatable) Identifies the name of a cost center.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Cost Center
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Cost Center]]
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `department` - (Optional) (Updatable) Identifies the name of a department.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Department
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Department]]
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `division` - (Optional) (Updatable) Identifies the name of a division.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Division
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Division]]
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `employee_number` - (Optional) (Updatable) Numeric or alphanumeric identifier assigned to  a person, typically based on order of hire or association with an organization.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Employee Number
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Employee Number]]
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `manager` - (Optional) (Updatable) The User's manager. A complex type that optionally allows Service Providers to represent organizational hierarchy by referencing the 'id' attribute of another User.

		**SCIM++ Properties:**
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Manager, deprecatedColumnHeaderName:Manager Name, mapsTo:manager.value]]
		* idcsPii: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display_name` - (Optional) (Updatable) The displayName of the User's manager. OPTIONAL and READ-ONLY.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the SCIM resource representing the User's manager.  RECOMMENDED.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Optional) (Updatable) The id of the SCIM resource representing  the User's  manager.  RECOMMENDED.

			**SCIM++ Properties:**
			* idcsCsvAttributeName: Manager Name
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `organization` - (Optional) (Updatable) Identifies the name of an organization.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Organization
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Organization Name, deprecatedColumnHeaderName:Organization]]
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextension_oci_tags` - (Optional) (Updatable) Oracle Cloud Infrastructure Tags.
	* `defined_tags` - (Optional) (Updatable) Oracle Cloud Infrastructure Defined Tags

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsCompositeKey: [namespace, key, value]
		* type: complex
		* idcsSearchable: true
		* required: false
		* mutability: readWrite
		* multiValued: true
		* returned: default
		* `key` - (Required) (Updatable) Oracle Cloud Infrastructure Tag key

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* type: string
			* required: true
			* mutability: readWrite
			* returned: default
			* idcsSearchable: true
			* uniqueness: none
		* `namespace` - (Required) (Updatable) Oracle Cloud Infrastructure Tag namespace

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* type: string
			* required: true
			* mutability: readWrite
			* returned: default
			* idcsSearchable: true
			* uniqueness: none
		* `value` - (Required) (Updatable) Oracle Cloud Infrastructure Tag value

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* required: true
			* idcsReturnEmptyWhenNull: true
			* mutability: readWrite
			* returned: default
			* type: string
			* idcsSearchable: true
			* uniqueness: none
	* `freeform_tags` - (Optional) (Updatable) Oracle Cloud Infrastructure Freeform Tags

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsCompositeKey: [key, value]
		* idcsSearchable: true
		* type: complex
		* required: false
		* mutability: readWrite
		* returned: default
		* multiValued: true
		* `key` - (Required) (Updatable) Oracle Cloud Infrastructure Tag key

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* type: string
			* required: true
			* mutability: readWrite
			* returned: default
			* idcsSearchable: true
			* uniqueness: none
		* `value` - (Required) (Updatable) Oracle Cloud Infrastructure Tag value

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* required: true
			* idcsReturnEmptyWhenNull: true
			* mutability: readWrite
			* returned: default
			* type: string
			* idcsSearchable: true
			* uniqueness: none
	* `tag_slug` - (Optional) (Updatable) Oracle Cloud Infrastructure Tag slug

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* type: binary
		* mutability: readOnly
		* returned: request
* `urnietfparamsscimschemasoracleidcsextensionadaptive_user` - (Optional) (Updatable) This extension defines attributes to manage user's risk score.
	* `risk_level` - (Optional) (Updatable) Risk Level

		**Added In:** 18.1.6

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `risk_scores` - (Optional) (Updatable) The risk score pertaining to the user.

		**Added In:** 18.1.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `last_update_timestamp` - (Required) (Updatable) Last update timestamp for the risk score

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: dateTime
			* uniqueness: none
		* `ref` - (Optional) (Updatable) Risk Provider Profile URI: URI that corresponds to risk source identifier.

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: reference
			* uniqueness: none
		* `risk_level` - (Required) (Updatable) Risk Level

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
		* `score` - (Required) (Updatable) Risk Score value

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: integer
			* uniqueness: none
			* idcsMaxValue: 100
			* idcsMinValue: 0
		* `source` - (Optional) (Updatable) Risk Provider Profile Source

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `status` - (Optional) (Updatable) Risk Provider Profile status

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `value` - (Required) (Updatable) Risk Provider Profile: Identifier for the provider service from which the risk score was received.

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensioncapabilities_user` - (Optional) (Updatable) User's Capabilities
	* `can_use_api_keys` - (Optional) (Updatable) Indicates whether a user can use API keys.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_auth_tokens` - (Optional) (Updatable) Indicates whether a user can use Auth tokens.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_console` - (Optional) (Updatable) Specifies whether user can access the Console.

		**Added In:** 2206280902

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_console_password` - (Optional) (Updatable) Indicates whether a user can use Console passwords.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_customer_secret_keys` - (Optional) (Updatable) Indicates whether a user can use customer secret keys.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_db_credentials` - (Optional) (Updatable) Indicates whether a user can use database credentials.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_oauth2client_credentials` - (Optional) (Updatable) Indicates whether a user can use OAuth2 client credentials.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_smtp_credentials` - (Optional) (Updatable) Indicates whether a user can use SMTP credentials.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensiondb_credentials_user` - (Optional) (Updatable) The database credentials user extension.
	* `db_login_attempts` - (Optional) (Updatable) The number of failed login attempts. The value is reset to 0 after a successful login.

		**Added In:** 2102181953

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `db_user_name` - (Optional) (Updatable) The database username.

		**Added In:** 2102181953

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* type: string
		* returned: request
		* caseExact: false
		* uniqueness: none
		* idcsSearchable: true
* `urnietfparamsscimschemasoracleidcsextensiondb_user_user` - (Optional) (Updatable) DB User extension
	* `db_global_roles` - (Optional) (Updatable) DB global roles to which the user is granted access.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsSensitive: none
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `domain_level_schema` - (Optional) (Updatable) DB domain level schema to which the user is granted access.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsSensitive: none
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `instance_level_schema` - (Optional) (Updatable) DB instance level schema to which the user is granted access.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsSensitive: none
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `is_db_user` - (Optional) (Updatable) If true, indicates this is a database user.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `password_verifiers` - (Optional) (Updatable) Password Verifiers for DB User.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsCompositeKey: [type]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `type` - (Required) (Updatable) Type of database password verifier (for example, MR-SHA512 or SSHA).

			**Added In:** 18.2.2

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - (Required) (Updatable) Hash value of database password verifier.

			**Added In:** 18.2.2

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* idcsSensitive: none
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionkerberos_user_user` - (Optional) (Updatable) Kerberos User extension
	* `realm_users` - (Optional) (Updatable) A list of kerberos realm users for an Oracle Identity Cloud Service User

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `principal_name` - (Optional) (Updatable) Principal Name of the KerberosRealmUser associated with the Oracle Identity Cloud Service User.

			**SCIM++ Properties:**
			* idcsPii: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `realm_name` - (Optional) (Updatable) Realm Name for the KerberosRealmUser associated with the Oracle Identity Cloud Service User.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding KerberosRealmUser resource associated with the Oracle Identity Cloud Service User.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) id of the KerberosRealmUser associated with the Oracle Identity Cloud Service User.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionmfa_user` - (Optional) (Updatable) This extension defines attributes used to manage Multi-Factor Authentication within a service provider. The extension is typically applied to a User resource, but MAY be applied to other resources that use MFA.
	* `bypass_codes` - (Optional) (Updatable) A list of bypass codes that belongs to the user.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding BypassCode resource which belongs to user

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) The user's bypass code identifier.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
	* `devices` - (Optional) (Updatable) A list of devices enrolled by the user.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `authentication_method` - (Optional) (Updatable) The authentication method.

			**Added In:** 2009232244

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `display` - (Optional) (Updatable) A human readable name, primarily used for display purposes. READ-ONLY.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `factor_status` - (Optional) (Updatable) The device authentication factor status.

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `factor_type` - (Optional) (Updatable) The device authentication factor type.

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `last_sync_time` - (Optional) (Updatable) The last sync time for device.

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: dateTime
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding Device resource which belongs to user.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `status` - (Optional) (Updatable) The device's status.

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `third_party_vendor_name` - (Optional) (Updatable) The third-party factor vendor name.

			**Added In:** 2009232244

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - (Required) (Updatable) The user's device identifier.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
	* `login_attempts` - (Optional) (Updatable) The number of incorrect multi factor authentication sign in attempts made by this user. The user is  locked if this reaches the threshold specified in the maxIncorrectAttempts attribute in AuthenticationFactorSettings.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* idcsRequiresImmediateReadAfterWriteForAccessFlows: true
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `mfa_enabled_on` - (Optional) (Updatable) The date when the user enrolled in multi factor authentication. This will be set to null, when the user resets their factors.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `mfa_ignored_apps` - (Optional) (Updatable) User MFA Ignored Apps Identifiers

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `mfa_status` - (Optional) (Updatable) The user opted for MFA.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `preferred_authentication_factor` - (Optional) (Updatable) The preferred authentication factor type.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `preferred_authentication_method` - (Optional) (Updatable) The preferred authentication method.

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `preferred_device` - (Optional) (Updatable) The user's preferred device.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - (Optional) (Updatable) The device display name.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI that corresponds to the device resource.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) The user's preferred device identifier.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `preferred_third_party_vendor` - (Optional) (Updatable) The preferred third-party vendor name.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `trusted_user_agents` - (Optional) (Updatable) A list of trusted User Agents owned by this user. Multi-Factored Authentication uses Trusted User Agents to authenticate users.  A User Agent is software application that a user uses to issue requests. For example, a User Agent could be a particular browser (possibly one of several executing on a desktop or laptop) or a particular mobile application (again, oneof several executing on a particular mobile device). A User Agent is trusted once the Multi-Factor Authentication has verified it in some way.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `display` - (Optional) (Updatable) A human-readable identifier for this trusted user agent, used primarily for display purposes. READ-ONLY.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding trusted user agent resource.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) The user's trusted user agent identifier.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionpassword_state_user` - (Optional) (Updatable) This extension defines attributes used to manage account passwords within a Service Provider. The extension is typically applied to a User resource, but MAY be applied to other resources that use passwords.
	* `applicable_password_policy` - (Optional) (Updatable) Applicable Password Policy

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `display` - (Optional) (Updatable) Password Policy Display Name

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `priority` - (Optional) (Updatable) PasswordPolicy priority

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding PasswordPolicy resource.

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) The identifier of the password policy.

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
	* `cant_change` - (Optional) (Updatable) Indicates that the current password MAY NOT be changed and all other password expiry settings SHALL be ignored

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `cant_expire` - (Optional) (Updatable) Indicates that the password expiry policy will not be applied for the current Resource

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `expired` - (Optional) (Updatable) Indicates that the password has expired

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `last_failed_validation_date` - (Optional) (Updatable) A DateTime that specifies the date and time when last failed password validation was set

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `last_successful_set_date` - (Optional) (Updatable) A DateTime that specifies the date and time when the current password was set

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `last_successful_validation_date` - (Optional) (Updatable) A DateTime that specifies the date and time when last successful password validation was set

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `must_change` - (Optional) (Updatable) Indicates that the subject password value MUST change on next login. If not changed, typically the account is locked. The value may be set indirectly when the subject's current password expires or directly set by an administrator.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionpasswordless_user` - (Optional) (Updatable) This extension defines attributes used to manage Passwordless-Factor Authentication within a service provider. The extension is typically applied to a User resource, but MAY be applied to other resources that use MFA.
	* `factor_identifier` - (Optional) (Updatable) Factor Identifier ID

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - (Optional) (Updatable) Factor Identifier display name

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI that corresponds to the device resource

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) The identifier of the User's preferred device

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `factor_method` - (Optional) (Updatable) Authentication Factor Method

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `factor_type` - (Optional) (Updatable) Authentication Factor Type

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionposix_user` - (Optional) (Updatable) POSIX User extension
	* `gecos` - (Optional) (Updatable) General information about the POSIX account such as their real name and phone number

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `gid_number` - (Optional) (Updatable) Primary Group identifier of the POSIX user

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `home_directory` - (Optional) (Updatable) The absolute path to the home directory of the POSIX account

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `login_shell` - (Optional) (Updatable) The path to the login shell of the POSIX account

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `uid_number` - (Optional) (Updatable) Integer uniquely identifying a user in a POSIX administrative domain

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: integer
		* uniqueness: server
* `urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user` - (Optional) (Updatable) This extension defines the attributes used to store the security questions of a user.
	* `sec_questions` - (Optional) (Updatable) The schema used to mnage security question and answers provided by a user for account recovery and/or MFA. While setting up security questions, a user can also provide a hint for the answer.

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `answer` - (Required) (Updatable) The answer provided by a user for a security question.

			**SCIM++ Properties:**
			* idcsCsvAttributeName: Answer
			* idcsSearchable: false
			* idcsSensitive: hash
			* multiValued: false
			* mutability: writeOnly
			* required: true
			* returned: never
			* type: string
			* uniqueness: none
			* idcsPii: true
		* `hint_text` - (Optional) (Updatable) The hint for an answer that's given by user when setting up a security question.

			**SCIM++ Properties:**
			* caseExact: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding Security Question resource.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) The identifier of the question selected by the user when setting up a security question.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionself_change_user` - (Optional) (Updatable) Controls whether a user can update themselves or not via User related APIs
	* `allow_self_change` - (Optional) (Updatable) If true, allows requesting user to update themselves. If false, requesting user can't update themself (default).

		**Added In:** 2205182039

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: writeOnly
		* required: false
		* returned: never
		* type: boolean
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionself_registration_user` - (Optional) (Updatable) This extension defines attributes used to manage self registration profile linked to the user.
	* `consent_granted` - (Optional) (Updatable) A boolean value that indicates whether the consent is granted.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `self_registration_profile` - (Required) (Updatable) Self registration profile used when user is self registered.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: request
		* type: complex
		* uniqueness: none
		* `display` - (Optional) (Updatable) A human readable name, primarily used for display purposes. READ-ONLY.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: request
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) URI of the profile.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) Self Registration Profile Id

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
	* `user_token` - (Optional) (Updatable) User token used for auto-login.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionsff_user` - (Optional) (Updatable) SFF Auth Keys User extension
	* `sff_auth_keys` - (Optional) (Updatable) SFF auth keys clob

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionsocial_account_user` - (Optional) (Updatable) Social User extension
	* `social_accounts` - (Optional) (Updatable) Description:

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* idcsPii: true
		* type: complex
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
		* `ref` - (Optional) (Updatable) The URI of the corresponding SocialAccount resource linked with the user

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) 
* `urnietfparamsscimschemasoracleidcsextensionterms_of_use_user` - (Optional) (Updatable) Terms Of Use extension
	* `terms_of_use_consents` - (Optional) (Updatable) Description:

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none 
		* `ref` - (Optional) (Updatable) The URI of the corresponding TermsOfUseConsent resource linked with the user

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) 
* `urnietfparamsscimschemasoracleidcsextensionuser_credentials_user` - (Optional) (Updatable) User's credentials
	* `api_keys` - (Optional) (Updatable) A list of API keys corresponding to user.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `key` - (Optional) (Updatable) The user's API key value.

			**Added In:** 2106240046

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ocid` - (Optional) (Updatable) The user's API key OCID.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding ApiKey resource to which the user belongs.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Optional) (Updatable) The user's API key identifier.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `auth_tokens` - (Optional) (Updatable) A list of Auth tokens corresponding to user.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `ocid` - (Optional) (Updatable) The user's Auth token OCID.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding AuthToken resource to which the user belongs.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Optional) (Updatable) The user's Auth token identifier.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `customer_secret_keys` - (Optional) (Updatable) A list of customer secret keys corresponding to user.

		**Added In:** 2102181953

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `ocid` - (Optional) (Updatable) The user's customer secret key OCID.

			**Added In:** 2102181953

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding CustomerSecretKey resource to which the user belongs.

			**Added In:** 2102181953

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Optional) (Updatable) The user's customer secret key identifier.

			**Added In:** 2102181953

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `db_credentials` - (Optional) (Updatable) A list of database credentials corresponding to user.

		**Added In:** 2102181953

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `ocid` - (Optional) (Updatable) The user's database credential OCID.

			**Added In:** 2102181953

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding UserDbCredential resource to which the user belongs.

			**Added In:** 2102181953

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Optional) (Updatable) The user's database credential identifier.

			**Added In:** 2102181953

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `o_auth2client_credentials` - (Optional) (Updatable) A list of OAuth2 client credentials corresponding to a user.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `ocid` - (Optional) (Updatable) The user's OAuth2 client credential OCID.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding OAuth2ClientCredential resource to which the user belongs.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Optional) (Updatable) The user's OAuth2 client credential identifier.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `smtp_credentials` - (Optional) (Updatable) A list of SMTP credentials corresponding to user.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `ocid` - (Optional) (Updatable) The user's Auth token OCID.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding SmtpCredential resource to which the user belongs.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Optional) (Updatable) The user's SMTP credential identifier.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionuser_state_user` - (Optional) (Updatable) This extension defines the attributes used to manage account passwords within a service provider. The extension is typically applied to a User resource, but MAY be applied to other resources that use passwords.
	* `last_failed_login_date` - (Optional) (Updatable) The last failed login date.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsAllowUpdatesInReadOnlyMode: true
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `last_successful_login_date` - (Optional) (Updatable) The last successful login date.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* idcsAllowUpdatesInReadOnlyMode: true
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `locked` - (Optional) (Updatable) A complex attribute that indicates an account is locked (blocking any new sessions).

		**SCIM++ Properties:**
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Locked, mapsTo:locked.on], [columnHeaderName:Locked Reason, mapsTo:locked.reason], [columnHeaderName:Locked Date, mapsTo:locked.lockDate]]
		* idcsSearchable: false
		* idcsAllowUpdatesInReadOnlyMode: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `expired` - (Optional) (Updatable) Indicates whether the user password is expired. If this value is false, password expiry is still evaluated during user login.

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: request
			* type: boolean
			* uniqueness: none
		* `lock_date` - (Optional) (Updatable) The date and time that the current resource was locked.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* idcsRequiresWriteForAccessFlows: true
			* required: false
			* returned: default
			* type: dateTime
			* uniqueness: none
		* `on` - (Optional) (Updatable) Indicates that the account is locked.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* idcsRequiresWriteForAccessFlows: true
			* idcsRequiresImmediateReadAfterWriteForAccessFlows: true
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `reason` - (Optional) (Updatable) Indicates the reason for locking the account. Valid values are: 0 - failed password login attempts, 1 - admin lock, 2 - failed reset password attempts, 3 - failed MFA login attempts, 4 - failed MFA login attempts for federated user, 5 - failed Database login attempts

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* idcsRequiresWriteForAccessFlows: true
			* idcsRequiresImmediateReadAfterWriteForAccessFlows: true
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
	* `login_attempts` - (Optional) (Updatable) The number of failed login attempts. The value is reset to 0 after a successful login.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsAllowUpdatesInReadOnlyMode: true
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* idcsRequiresImmediateReadAfterWriteForAccessFlows: true
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `max_concurrent_sessions` - (Optional) (Updatable) The maximum number of concurrent sessions for a user.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsMaxValue: 999
		* idcsMinValue: 1
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `previous_successful_login_date` - (Optional) (Updatable) The previous successful login date.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `recovery_attempts` - (Optional) (Updatable) The number of failed recovery attempts. The value is reset to 0 after a successful login.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `recovery_enroll_attempts` - (Optional) (Updatable) The number of failed account recovery enrollment attempts.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `recovery_locked` - (Optional) (Updatable) A complex attribute that indicates a password recovery is locked (blocking any new sessions).

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `lock_date` - (Optional) (Updatable) The date and time that the current resource was locked.

			**Added In:** 19.1.4

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* idcsRequiresWriteForAccessFlows: true
			* required: false
			* returned: default
			* type: dateTime
			* uniqueness: none
		* `on` - (Optional) (Updatable) Indicates that the recovery is locked.

			**Added In:** 19.1.4

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* idcsRequiresWriteForAccessFlows: true
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionuser_user` - (Optional) (Updatable) Oracle Identity Cloud Service User
	* `user_provider` - (Optional) (Updatable) Registration provider

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `account_recovery_required` - (Optional) (Updatable) Boolean value to prompt user to setup account recovery during login.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `accounts` - (Optional) (Updatable) Accounts assigned to this User. Each value of this attribute refers to an app-specific identity that is owned by this User. Therefore, this attribute is a convenience that allows one to see on each User the Apps to which that User has access.

		**SCIM++ Properties:**
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `active` - (Optional) (Updatable) Status of the account

			**Added In:** 17.4.6

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `app_id` - (Optional) (Updatable) The ID of the App to which this Account gives access.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `name` - (Optional) (Updatable) Name of the account assigned to the User.

			**Added In:** 17.4.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the Account assigned to the User.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Optional) (Updatable) The Id of the Account assigned to the User.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `app_roles` - (Optional) (Updatable) A list of all AppRoles to which this User belongs directly, indirectly or implicitly. The User could belong directly because the User is a member of the AppRole, could belong indirectly because the User is a member of a Group that is a member of the AppRole, or could belong implicitly because the AppRole is public.

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `admin_role` - (Optional) (Updatable) If true, then the role provides administrative access privileges. READ-ONLY.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `app_id` - (Optional) (Updatable) The ID of the App that defines this AppRole.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `app_name` - (Optional) (Updatable) The name (Client ID) of the App that defines this AppRole.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `display` - (Optional) (Updatable) The display name of the AppRole assigned to the User.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `legacy_group_name` - (Optional) (Updatable) The name (if any) under which this AppRole should appear in this User's group-memberships for reasons of backward compatibility. Oracle Identity Cloud Service distinguishes between Groups and AppRoles, but some services still expect AppRoles appear as if they were service-instance-specific Groups.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the AppRole assigned to the User.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - (Optional) (Updatable) The kind of membership this User has in the AppRole. A value of 'direct' indicates that the User is a member of the AppRole.  A value of  'indirect' indicates that the User is a member of a Group that is a member of the AppRole.  A value of 'implicit' indicates that the AppRole is public.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: request
			* type: string
			* uniqueness: none
		* `value` - (Required) (Updatable) The Id of the AppRole assigned to the User.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
	* `applicable_authentication_target_app` - (Optional) (Updatable) The app against which the user will authenticate. The value is not persisted but rather calculated. If the user's delegatedAuthenticationTargetApp is set, that value is returned. Otherwise, the app returned by evaluating the user's applicable Delegated Authentication Policy is returned.

		**Added In:** 18.1.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `display` - (Optional) (Updatable) App Display Name

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) App URI

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `target_request_timeout` - (Optional) (Updatable) Timeout interval for Synchronization TargetAction in milliseconds

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `type` - (Required) (Updatable) A label that indicates whether this is an App or IdentitySource.

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - (Optional) (Updatable) App identifier

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: true
			* multiValued: false
			* mutability: readOnly
			* returned: default
			* type: string
			* uniqueness: none
	* `bypass_notification` - (Optional) (Updatable) A Boolean value indicating whether or not to send email notification after creating the user. This attribute is not used in update/replace operations.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeNameMappings: [[columnHeaderName:ByPass Notification]]
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: never
		* type: boolean
		* uniqueness: none
	* `creation_mechanism` - (Optional) (Updatable) User creation mechanism

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeNameMappings: [[defaultValue:import]]
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `delegated_authentication_target_app` - (Optional) (Updatable) If set, indicates the user's preferred authentication target app. If not set and the user's \"syncedFromApp\" is set and is enabled for delegated authentication, it is used. Otherwise, the user authenticates locally to Oracle Identity Cloud Service.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - (Optional) (Updatable) App Display Name

			**Added In:** 17.4.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) App URI

			**Added In:** 17.4.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - (Required) (Updatable) A label that indicates whether this is an App or IdentitySource.

			**Added In:** 17.4.6

			**SCIM++ Properties:**
			* idcsDefaultValue: IdentitySource
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - (Required) (Updatable) App identifier

			**Added In:** 17.4.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `do_not_show_getting_started` - (Optional) (Updatable) A Boolean value indicating whether or not to hide the getting started page

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `grants` - (Optional) (Updatable) Grants to this User. Each value of this attribute refers to a Grant to this User of some App (and optionally of some entitlement). Therefore, this attribute is a convenience that allows one to see on each User all of the Grants to that User.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `app_id` - (Optional) (Updatable) The ID of the App in this Grant.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `grant_mechanism` - (Optional) (Updatable) Each value of grantMechanism indicates how (or by what component) some App (or App-Entitlement) was granted. A customer or the UI should use only grantMechanism values that start with 'ADMINISTRATOR':
			* 'ADMINISTRATOR_TO_USER' is for a direct grant to a specific User.
			* 'ADMINISTRATOR_TO_GROUP' is for a grant to a specific Group, which results in indirect grants to Users who are members of that Group.
			* 'ADMINISTRATOR_TO_APP' is for a grant to a specific App.  The grantee (client) App gains access to the granted (server) App.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `grantor_id` - (Optional) (Updatable) Grantor identifier

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of this Grant to this User.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Optional) (Updatable) The ID of this Grant to this User.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `group_membership_last_modified` - (Optional) (Updatable) Specifies date time when a User's group membership was last modified.

		**Added In:** 2304270343

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `idcs_app_roles_limited_to_groups` - (Optional) (Updatable) Description:

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsCompositeKey: [value, idcsAppRoleId]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex 
		* `display` - (Optional) (Updatable) Group display name

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `idcs_app_role_id` - (Required) (Updatable) The id of the Oracle Identity Cloud Service AppRole grant limited to one or more Groups.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsCsvAttributeName: IDCS AppRole Name
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `ocid` - (Optional) (Updatable) The ocid of a Group the AppRole Grant is limited to

			**Added In:** 2202230830

			**SCIM++ Properties:**
			* idcsCsvAttributeName: Group Ocid
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the SCIM resource representing the Group manager.  RECOMMENDED.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) The id of a Group the AppRole Grant is limited to

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsCsvAttributeName: Group Name
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `is_account_recovery_enrolled` - (Optional) (Updatable) A Boolean value indicating whether or not a user is enrolled for account recovery

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `is_authentication_delegated` - (Optional) (Updatable) A Boolean value indicating whether or not authentication request by this user should be delegated to a remote app. This value should be true only when the User was originally synced from an app which is enabled for delegated authentication

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: never
		* type: boolean
		* uniqueness: none
	* `is_federated_user` - (Optional) (Updatable) A Boolean value indicating whether or not the user is federated.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Federated
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Federated]]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_group_membership_normalized` - (Optional) (Updatable) A Boolean value indicating whether or not group membership is normalized for this user.

		**Deprecated Since: 19.3.3**

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: never
		* type: boolean
		* uniqueness: none
	* `is_group_membership_synced_to_users_groups` - (Optional) (Updatable) A Boolean value Indicates whether this User's group membership has been sync'ed from Group.members to UsersGroups.

		**Added In:** 19.3.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: never
		* type: boolean
		* uniqueness: none
	* `notification_email_template_id` - (Optional) (Updatable) Specifies the EmailTemplate to be used when sending notification to the user this request is for. If specified, it overrides the default EmailTemplate for this event.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: writeOnly
		* required: false
		* returned: never
		* type: string
		* uniqueness: none
	* `preferred_ui_landing_page` - (Optional) (Updatable) User's preferred landing page following login, logout and reset password.

		**Added In:** 2302092332

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `status` - (Optional) (Updatable) A supplemental status indicating the reason why a user is disabled

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `support_accounts` - (Optional) (Updatable) A list of Support Accounts corresponding to user.

		**Added In:** 2103141444

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `user_provider` - (Optional) (Updatable) User Support Account Provider

			**Added In:** 2103141444

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ocid` - (Optional) (Updatable) The OCID of the user's support account.

			**Added In:** 2103141444

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding Support Account resource to which the user belongs

			**Added In:** 2103141444

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `user_id` - (Optional) (Updatable) User Support User Id

			**Added In:** 2103141444

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `value` - (Optional) (Updatable) The identifier of the User's support Account.

			**Added In:** 2103141444

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `synced_from_app` - (Optional) (Updatable) Managed App or an Identity Source from where the user is synced. If enabled, this Managed App or Identity Source can be used for performing delegated authentication.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - (Optional) (Updatable) App Display Name

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) App URI

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - (Required) (Updatable) A label that indicates whether this is an App or IdentitySource.

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* idcsDefaultValue: IdentitySource
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - (Required) (Updatable) App identifier

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `user_flow_controlled_by_external_client` - (Optional) (Updatable) A Boolean value indicating whether to bypass notification and return user token to be used by an external client to control the user flow.

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: never
		* type: boolean
		* uniqueness: none
	* `user_token` - (Optional) (Updatable) User token returned if userFlowControlledByExternalClient is true

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `ref` - (Optional) (Updatable) User Token URI

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Optional) (Updatable) User Token identifier

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
* `user_name` - (Required) (Updatable) User name

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: User ID
	* idcsCsvAttributeNameMappings: [[columnHeaderName:User Name, deprecatedColumnHeaderName:User ID]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: always
	* type: string
	* uniqueness: global
* `user_type` - (Optional) (Updatable) Used to identify the organization-to-user relationship

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: User Type
	* idcsCsvAttributeNameMappings: [[columnHeaderName:User Type]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `x509certificates` - (Optional) (Updatable) A list of certificates issued to the User.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - (Optional) (Updatable) A human readable name, primarily used for display purposes.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - (Optional) (Updatable) A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - (Optional) (Updatable) A label indicating the attribute's function.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) The value of a X509 certificate.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: binary
		* uniqueness: none


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `active` - User status

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Active
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Active]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `addresses` - A physical mailing address for this User, as described in (address Element). Canonical Type Values of work, home, and other. The value attribute is a complex type with the following sub-attributes.

	**SCIM++ Properties:**
	* idcsCompositeKey: [type]
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Work Address Street, deprecatedColumnHeaderName:Work Street Address, mapsTo:addresses[work].streetAddress], [columnHeaderName:Work Address Locality, deprecatedColumnHeaderName:Work City, mapsTo:addresses[work].locality], [columnHeaderName:Work Address Region, deprecatedColumnHeaderName:Work State, mapsTo:addresses[work].region], [columnHeaderName:Work Address Postal Code, deprecatedColumnHeaderName:Work Postal Code, mapsTo:addresses[work].postalCode], [columnHeaderName:Work Address Country, deprecatedColumnHeaderName:Work Country, mapsTo:addresses[work].country], [columnHeaderName:Work Address Formatted, mapsTo:addresses[work].formatted], [columnHeaderName:Home Address Formatted, mapsTo:addresses[home].formatted], [columnHeaderName:Other Address Formatted, mapsTo:addresses[other].formatted], [columnHeaderName:Home Address Street, mapsTo:addresses[home].streetAddress], [columnHeaderName:Other Address Street, mapsTo:addresses[other].streetAddress], [columnHeaderName:Home Address Locality, mapsTo:addresses[home].locality], [columnHeaderName:Other Address Locality, mapsTo:addresses[other].locality], [columnHeaderName:Home Address Region, mapsTo:addresses[home].region], [columnHeaderName:Other Address Region, mapsTo:addresses[other].region], [columnHeaderName:Home Address Country, mapsTo:addresses[home].country], [columnHeaderName:Other Address Country, mapsTo:addresses[other].country], [columnHeaderName:Home Address Postal Code, mapsTo:addresses[home].postalCode], [columnHeaderName:Other Address Postal Code, mapsTo:addresses[other].postalCode], [columnHeaderName:Primary Address Type, mapsTo:addresses[$(type)].primary]]
	* idcsPii: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `country` - The country name component.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCanonicalValueSourceFilter: attrName eq "countries" and attrValues.value eq "upper($(country))"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `formatted` - The full mailing address, formatted for display or use with a mailing label. This attribute MAY contain newlines.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `locality` - The city or locality component.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `postal_code` - The zipcode or postal code component.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `region` - The state or region component.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `street_address` - The full street address component, which may include house number, street name, PO BOX, and multi-line extended street address information. This attribute MAY contain newlines.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `type` - A label indicating the attribute's function; e.g., 'work' or 'home'.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
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
* `description` - Description of the user

	**Added In:** 2012271618

	**SCIM++ Properties:**
	* caseExact: false
	* idcsPii: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `display_name` - Display name

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Display Name
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Display Name]]
	* idcsPii: true
	* idcsSearchable: true
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
* `emails` - A complex attribute representing emails

	**SCIM++ Properties:**
	* idcsCompositeKey: [value, type]
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Work Email, mapsTo:emails[work].value], [columnHeaderName:Home Email, mapsTo:emails[home].value], [columnHeaderName:Primary Email Type, mapsTo:emails[$(type)].primary], [columnHeaderName:Other Email, mapsTo:emails[other].value], [columnHeaderName:Recovery Email, mapsTo:emails[recovery].value], [columnHeaderName:Work Email Verified, mapsTo:emails[work].verified], [columnHeaderName:Home Email Verified, mapsTo:emails[home].verified], [columnHeaderName:Other Email Verified, mapsTo:emails[other].verified], [columnHeaderName:Recovery Email Verified, mapsTo:emails[recovery].verified]]
	* idcsPii: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `pending_verification_data` - Pending e-mail address verification

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - A Boolean value that indicates whether the email address is the primary email address. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `secondary` - A Boolean value that indicates whether the email address is the secondary email address. The secondary attribute value 'true' MUST appear no more than once.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - Type of email address

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Email address

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `verified` - A Boolean value that indicates whether or not the e-mail address is verified

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
* `entitlements` - A list of entitlements for the User that represent a thing the User has.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value, type]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - A human readable name, primarily used for display purposes.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - A label indicating the attribute's function.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - The value of an entitlement.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `external_id` - An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeNameMappings: [[columnHeaderName:External Id]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `groups` - A list of groups that the user belongs to, either thorough direct membership, nested groups, or dynamically calculated

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `date_added` - Date when the member is Added to the group

		**Added In:** 2105200541

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
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
	* `external_id` - An identifier for the Resource as defined by the Service Consumer. READ-ONLY.

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `membership_ocid` - The membership OCID.

		**Added In:** 2103141444

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `non_unique_display` - A human readable name for Group as defined by the Service Consumer. READ-ONLY.

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocid` - The OCID of the User's group.

		**Added In:** 2102181953

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - The URI of the corresponding Group resource to which the user belongs

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - A label indicating the attribute's function; e.g., 'direct' or 'indirect'.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `value` - The identifier of the User's group.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
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
* `ims` - User's instant messaging addresses

	**SCIM++ Properties:**
	* idcsCompositeKey: [value, type]
	* idcsPii: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - A human-readable name, primarily used for display purposes

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - A Boolean value that indicates the 'primary' or preferred attribute value for this attribute--for example, the preferred messenger or primary messenger. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - A label that indicates the attribute's function--for example, 'aim', 'gtalk', or 'mobile'

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - User's instant messaging address

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `locale` - Used to indicate the User's default location for purposes of localizing items such as currency, date and time format, numerical representations, and so on.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Locale
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Locale]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
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
* `name` - A complex attribute that contains attributes representing the name

	**SCIM++ Properties:**
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Formatted Name, mapsTo:name.formatted], [columnHeaderName:Honorific Prefix, mapsTo:name.honorificPrefix], [columnHeaderName:First Name, mapsTo:name.givenName], [columnHeaderName:Middle Name, mapsTo:name.middleName], [columnHeaderName:Last Name, mapsTo:name.familyName], [columnHeaderName:Honorific Suffix, mapsTo:name.honorificSuffix]]
	* idcsPii: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `family_name` - Last name

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Last Name
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `formatted` - Full name

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `given_name` - First name

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: First Name
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `honorific_prefix` - Prefix

		**SCIM++ Properties:**
		* idcsCsvAttributeName: Honorific Prefix
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `honorific_suffix` - Suffix

		**SCIM++ Properties:**
		* idcsCsvAttributeName: Honorific Suffix
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `middle_name` - Middle name

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Middle Name
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `nick_name` - Nick name

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Nick Name
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Nick Name]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
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
* `password` - Password attribute. Max length for password is controlled via Password Policy.

	**SCIM++ Properties:**
	* idcsCsvAttributeName: Password
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Password]]
	* idcsPii: true
	* idcsSearchable: false
	* idcsSensitive: hash
	* multiValued: false
	* mutability: writeOnly
	* required: false
	* returned: never
	* type: string
	* uniqueness: none
* `phone_numbers` - Phone numbers

	**SCIM++ Properties:**
	* idcsCompositeKey: [value, type]
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Work Phone, mapsTo:phoneNumbers[work].value], [columnHeaderName:Mobile No, mapsTo:phoneNumbers[mobile].value], [columnHeaderName:Home Phone, mapsTo:phoneNumbers[home].value], [columnHeaderName:Fax, mapsTo:phoneNumbers[fax].value], [columnHeaderName:Pager, mapsTo:phoneNumbers[pager].value], [columnHeaderName:Other Phone, mapsTo:phoneNumbers[other].value], [columnHeaderName:Recovery Phone, mapsTo:phoneNumbers[recovery].value], [columnHeaderName:Primary Phone Type, mapsTo:phoneNumbers[$(type)].primary]]
	* idcsPii: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - A human-readable name, primarily used for display purposes. READ ONLY

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - A Boolean value that indicates the 'primary' or preferred attribute value for this attribute--for example, the preferred phone number or primary phone number. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - A label that indicates the attribute's function- for example, 'work', 'home', or 'mobile'

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - User's phone number

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `verified` - A Boolean value that indicates if the phone number is verified.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
* `photos` - URLs of photos for the User

	**SCIM++ Properties:**
	* idcsCompositeKey: [value, type]
	* idcsPii: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - A human readable name, primarily used for display purposes.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g., the preferred photo or thumbnail. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - A label indicating the attribute's function; e.g., 'photo' or 'thumbnail'.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - URL of a photo for the User

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: reference
		* uniqueness: none
* `preferred_language` - User's preferred written or spoken language used for localized user interfaces

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Preferred Language
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Preferred Language]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `profile_url` - A fully-qualified URL to a page representing the User's online profile

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Profile URL
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Profile Url]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: reference
	* uniqueness: none
* `roles` - A list of roles for the User that collectively represent who the User is; e.g., 'Student', 'Faculty'.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value, type]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - A human readable name, primarily used for display purposes.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - A label indicating the attribute's function.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - The value of a role.

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
* `timezone` - User's timezone

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCanonicalValueSourceFilter: attrName eq "timezones" and attrValues.value eq "$(timezone)"
	* idcsCanonicalValueSourceResourceType: AllowedValue
	* idcsCsvAttributeName: TimeZone
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Time Zone, deprecatedColumnHeaderName:TimeZone]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `title` - Title

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Title
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Title]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `urnietfparamsscimschemasextensionenterprise20user` - Enterprise User
	* `cost_center` - Identifies the name of a cost center.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Cost Center
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Cost Center]]
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `department` - Identifies the name of a department.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Department
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Department]]
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `division` - Identifies the name of a division.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Division
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Division]]
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `employee_number` - Numeric or alphanumeric identifier assigned to  a person, typically based on order of hire or association with an organization.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Employee Number
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Employee Number]]
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `manager` - The User's manager. A complex type that optionally allows Service Providers to represent organizational hierarchy by referencing the 'id' attribute of another User.

		**SCIM++ Properties:**
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Manager, deprecatedColumnHeaderName:Manager Name, mapsTo:manager.value]]
		* idcsPii: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display_name` - The displayName of the User's manager. OPTIONAL and READ-ONLY.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI of the SCIM resource representing the User's manager.  RECOMMENDED.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The id of the SCIM resource representing  the User's  manager.  RECOMMENDED.

			**SCIM++ Properties:**
			* idcsCsvAttributeName: Manager Name
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `organization` - Identifies the name of an organization.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Organization
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Organization Name, deprecatedColumnHeaderName:Organization]]
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextension_oci_tags` - Oracle Cloud Infrastructure Tags.
	* `defined_tags` - Oracle Cloud Infrastructure Defined Tags

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsCompositeKey: [namespace, key, value]
		* type: complex
		* idcsSearchable: true
		* required: false
		* mutability: readWrite
		* multiValued: true
		* returned: default
		* `key` - Oracle Cloud Infrastructure Tag key

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* type: string
			* required: true
			* mutability: readWrite
			* returned: default
			* idcsSearchable: true
			* uniqueness: none
		* `namespace` - Oracle Cloud Infrastructure Tag namespace

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* type: string
			* required: true
			* mutability: readWrite
			* returned: default
			* idcsSearchable: true
			* uniqueness: none
		* `value` - Oracle Cloud Infrastructure Tag value

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* required: true
			* idcsReturnEmptyWhenNull: true
			* mutability: readWrite
			* returned: default
			* type: string
			* idcsSearchable: true
			* uniqueness: none
	* `freeform_tags` - Oracle Cloud Infrastructure Freeform Tags

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsCompositeKey: [key, value]
		* idcsSearchable: true
		* type: complex
		* required: false
		* mutability: readWrite
		* returned: default
		* multiValued: true
		* `key` - Oracle Cloud Infrastructure Tag key

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* type: string
			* required: true
			* mutability: readWrite
			* returned: default
			* idcsSearchable: true
			* uniqueness: none
		* `value` - Oracle Cloud Infrastructure Tag value

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* required: true
			* idcsReturnEmptyWhenNull: true
			* mutability: readWrite
			* returned: default
			* type: string
			* idcsSearchable: true
			* uniqueness: none
	* `tag_slug` - Oracle Cloud Infrastructure Tag slug

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* type: binary
		* mutability: readOnly
		* returned: request
* `urnietfparamsscimschemasoracleidcsextensionadaptive_user` - This extension defines attributes to manage user's risk score.
	* `risk_level` - Risk Level

		**Added In:** 18.1.6

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `risk_scores` - The risk score pertaining to the user.

		**Added In:** 18.1.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `last_update_timestamp` - Last update timestamp for the risk score

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: dateTime
			* uniqueness: none
		* `ref` - Risk Provider Profile URI: URI that corresponds to risk source identifier.

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: reference
			* uniqueness: none
		* `risk_level` - Risk Level

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
		* `score` - Risk Score value

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: integer
			* uniqueness: none
			* idcsMaxValue: 100
			* idcsMinValue: 0
		* `source` - Risk Provider Profile Source

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `status` - Risk Provider Profile status

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `value` - Risk Provider Profile: Identifier for the provider service from which the risk score was received.

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensioncapabilities_user` - User's Capabilities
	* `can_use_api_keys` - Indicates whether a user can use API keys.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_auth_tokens` - Indicates whether a user can use Auth tokens.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_console` - Specifies whether user can access the Console.

		**Added In:** 2206280902

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_console_password` - Indicates whether a user can use Console passwords.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_customer_secret_keys` - Indicates whether a user can use customer secret keys.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_db_credentials` - Indicates whether a user can use database credentials.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_oauth2client_credentials` - Indicates whether a user can use OAuth2 client credentials.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `can_use_smtp_credentials` - Indicates whether a user can use SMTP credentials.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensiondb_credentials_user` - The database credentials user extension.
	* `db_login_attempts` - The number of failed login attempts. The value is reset to 0 after a successful login.

		**Added In:** 2102181953

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `db_user_name` - The database username.

		**Added In:** 2102181953

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* type: string
		* returned: request
		* caseExact: false
		* uniqueness: none
		* idcsSearchable: true
* `urnietfparamsscimschemasoracleidcsextensiondb_user_user` - DB User extension
	* `db_global_roles` - DB global roles to which the user is granted access.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsSensitive: none
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `domain_level_schema` - DB domain level schema to which the user is granted access.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsSensitive: none
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `instance_level_schema` - DB instance level schema to which the user is granted access.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsSensitive: none
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `is_db_user` - If true, indicates this is a database user.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `password_verifiers` - Password Verifiers for DB User.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsCompositeKey: [type]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `type` - Type of database password verifier (for example, MR-SHA512 or SSHA).

			**Added In:** 18.2.2

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - Hash value of database password verifier.

			**Added In:** 18.2.2

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* idcsSensitive: none
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionkerberos_user_user` - Kerberos User extension
	* `realm_users` - A list of kerberos realm users for an Oracle Identity Cloud Service User

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `principal_name` - Principal Name of the KerberosRealmUser associated with the Oracle Identity Cloud Service User.

			**SCIM++ Properties:**
			* idcsPii: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `realm_name` - Realm Name for the KerberosRealmUser associated with the Oracle Identity Cloud Service User.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI of the corresponding KerberosRealmUser resource associated with the Oracle Identity Cloud Service User.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - id of the KerberosRealmUser associated with the Oracle Identity Cloud Service User.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionmfa_user` - This extension defines attributes used to manage Multi-Factor Authentication within a service provider. The extension is typically applied to a User resource, but MAY be applied to other resources that use MFA.
	* `bypass_codes` - A list of bypass codes that belongs to the user.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `ref` - The URI of the corresponding BypassCode resource which belongs to user

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The user's bypass code identifier.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
	* `devices` - A list of devices enrolled by the user.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `authentication_method` - The authentication method.

			**Added In:** 2009232244

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `display` - A human readable name, primarily used for display purposes. READ-ONLY.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `factor_status` - The device authentication factor status.

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `factor_type` - The device authentication factor type.

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `last_sync_time` - The last sync time for device.

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: dateTime
			* uniqueness: none
		* `ref` - The URI of the corresponding Device resource which belongs to user.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `status` - The device's status.

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `third_party_vendor_name` - The third-party factor vendor name.

			**Added In:** 2009232244

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - The user's device identifier.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
	* `login_attempts` - The number of incorrect multi factor authentication sign in attempts made by this user. The user is  locked if this reaches the threshold specified in the maxIncorrectAttempts attribute in AuthenticationFactorSettings.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* idcsRequiresImmediateReadAfterWriteForAccessFlows: true
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `mfa_enabled_on` - The date when the user enrolled in multi factor authentication. This will be set to null, when the user resets their factors.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `mfa_ignored_apps` - User MFA Ignored Apps Identifiers

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `mfa_status` - The user opted for MFA.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `preferred_authentication_factor` - The preferred authentication factor type.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `preferred_authentication_method` - The preferred authentication method.

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `preferred_device` - The user's preferred device.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - The device display name.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI that corresponds to the device resource.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The user's preferred device identifier.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `preferred_third_party_vendor` - The preferred third-party vendor name.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `trusted_user_agents` - A list of trusted User Agents owned by this user. Multi-Factored Authentication uses Trusted User Agents to authenticate users.  A User Agent is software application that a user uses to issue requests. For example, a User Agent could be a particular browser (possibly one of several executing on a desktop or laptop) or a particular mobile application (again, oneof several executing on a particular mobile device). A User Agent is trusted once the Multi-Factor Authentication has verified it in some way.

		**Added In:** 18.3.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `display` - A human-readable identifier for this trusted user agent, used primarily for display purposes. READ-ONLY.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI of the corresponding trusted user agent resource.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The user's trusted user agent identifier.

			**Added In:** 18.3.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionpassword_state_user` - This extension defines attributes used to manage account passwords within a Service Provider. The extension is typically applied to a User resource, but MAY be applied to other resources that use passwords.
	* `applicable_password_policy` - Applicable Password Policy

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `display` - Password Policy Display Name

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `priority` - PasswordPolicy priority

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `ref` - The URI of the corresponding PasswordPolicy resource.

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The identifier of the password policy.

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
	* `cant_change` - Indicates that the current password MAY NOT be changed and all other password expiry settings SHALL be ignored

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `cant_expire` - Indicates that the password expiry policy will not be applied for the current Resource

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `expired` - Indicates that the password has expired

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `last_failed_validation_date` - A DateTime that specifies the date and time when last failed password validation was set

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `last_successful_set_date` - A DateTime that specifies the date and time when the current password was set

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `last_successful_validation_date` - A DateTime that specifies the date and time when last successful password validation was set

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `must_change` - Indicates that the subject password value MUST change on next login. If not changed, typically the account is locked. The value may be set indirectly when the subject's current password expires or directly set by an administrator.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionpasswordless_user` - This extension defines attributes used to manage Passwordless-Factor Authentication within a service provider. The extension is typically applied to a User resource, but MAY be applied to other resources that use MFA.
	* `factor_identifier` - Factor Identifier ID

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - Factor Identifier display name

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI that corresponds to the device resource

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The identifier of the User's preferred device

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `factor_method` - Authentication Factor Method

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `factor_type` - Authentication Factor Type

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionposix_user` - POSIX User extension
	* `gecos` - General information about the POSIX account such as their real name and phone number

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `gid_number` - Primary Group identifier of the POSIX user

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `home_directory` - The absolute path to the home directory of the POSIX account

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `login_shell` - The path to the login shell of the POSIX account

		**SCIM++ Properties:**
		* caseExact: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `uid_number` - Integer uniquely identifying a user in a POSIX administrative domain

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: integer
		* uniqueness: server
* `urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user` - This extension defines the attributes used to store the security questions of a user.
	* `sec_questions` - The schema used to mnage security question and answers provided by a user for account recovery and/or MFA. While setting up security questions, a user can also provide a hint for the answer.

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `answer` - The answer provided by a user for a security question.

			**SCIM++ Properties:**
			* idcsCsvAttributeName: Answer
			* idcsSearchable: false
			* idcsSensitive: hash
			* multiValued: false
			* mutability: writeOnly
			* required: true
			* returned: never
			* type: string
			* uniqueness: none
			* idcsPii: true
		* `hint_text` - The hint for an answer that's given by user when setting up a security question.

			**SCIM++ Properties:**
			* caseExact: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI of the corresponding Security Question resource.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The identifier of the question selected by the user when setting up a security question.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionself_change_user` - Controls whether a user can update themselves or not via User related APIs
	* `allow_self_change` - If true, allows requesting user to update themselves. If false, requesting user can't update themself (default).

		**Added In:** 2205182039

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: writeOnly
		* required: false
		* returned: never
		* type: boolean
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionself_registration_user` - This extension defines attributes used to manage self registration profile linked to the user.
	* `consent_granted` - A boolean value that indicates whether the consent is granted.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `self_registration_profile` - Self registration profile used when user is self registered.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: request
		* type: complex
		* uniqueness: none
		* `display` - A human readable name, primarily used for display purposes. READ-ONLY.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: request
			* type: string
			* uniqueness: none
		* `ref` - URI of the profile.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - Self Registration Profile Id

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
	* `user_token` - User token used for auto-login.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionsff_user` - SFF Auth Keys User extension
	* `sff_auth_keys` - SFF auth keys clob

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionsocial_account_user` - Social User extension
	* `social_accounts` - Description:

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* idcsPii: true
		* type: complex
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
		* `ref` - The URI of the corresponding SocialAccount resource linked with the user

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - 
* `urnietfparamsscimschemasoracleidcsextensionterms_of_use_user` - Terms Of Use extension
	* `terms_of_use_consents` - Description:

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none 
		* `ref` - The URI of the corresponding TermsOfUseConsent resource linked with the user

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - 
* `urnietfparamsscimschemasoracleidcsextensionuser_credentials_user` - User's credentials
	* `api_keys` - A list of API keys corresponding to user.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `key` - The user's API key value.

			**Added In:** 2106240046

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ocid` - The user's API key OCID.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - The URI of the corresponding ApiKey resource to which the user belongs.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The user's API key identifier.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `auth_tokens` - A list of Auth tokens corresponding to user.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `ocid` - The user's Auth token OCID.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - The URI of the corresponding AuthToken resource to which the user belongs.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The user's Auth token identifier.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `customer_secret_keys` - A list of customer secret keys corresponding to user.

		**Added In:** 2102181953

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `ocid` - The user's customer secret key OCID.

			**Added In:** 2102181953

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - The URI of the corresponding CustomerSecretKey resource to which the user belongs.

			**Added In:** 2102181953

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The user's customer secret key identifier.

			**Added In:** 2102181953

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `db_credentials` - A list of database credentials corresponding to user.

		**Added In:** 2102181953

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `ocid` - The user's database credential OCID.

			**Added In:** 2102181953

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - The URI of the corresponding UserDbCredential resource to which the user belongs.

			**Added In:** 2102181953

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The user's database credential identifier.

			**Added In:** 2102181953

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `o_auth2client_credentials` - A list of OAuth2 client credentials corresponding to a user.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `ocid` - The user's OAuth2 client credential OCID.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - The URI of the corresponding OAuth2ClientCredential resource to which the user belongs.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The user's OAuth2 client credential identifier.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `smtp_credentials` - A list of SMTP credentials corresponding to user.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `ocid` - The user's Auth token OCID.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - The URI of the corresponding SmtpCredential resource to which the user belongs.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The user's SMTP credential identifier.

			**Added In:** 2012271618

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionuser_state_user` - This extension defines the attributes used to manage account passwords within a service provider. The extension is typically applied to a User resource, but MAY be applied to other resources that use passwords.
	* `last_failed_login_date` - The last failed login date.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsAllowUpdatesInReadOnlyMode: true
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `last_successful_login_date` - The last successful login date.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* idcsAllowUpdatesInReadOnlyMode: true
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `locked` - A complex attribute that indicates an account is locked (blocking any new sessions).

		**SCIM++ Properties:**
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Locked, mapsTo:locked.on], [columnHeaderName:Locked Reason, mapsTo:locked.reason], [columnHeaderName:Locked Date, mapsTo:locked.lockDate]]
		* idcsSearchable: false
		* idcsAllowUpdatesInReadOnlyMode: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `expired` - Indicates whether the user password is expired. If this value is false, password expiry is still evaluated during user login.

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: request
			* type: boolean
			* uniqueness: none
		* `lock_date` - The date and time that the current resource was locked.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* idcsRequiresWriteForAccessFlows: true
			* required: false
			* returned: default
			* type: dateTime
			* uniqueness: none
		* `on` - Indicates that the account is locked.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* idcsRequiresWriteForAccessFlows: true
			* idcsRequiresImmediateReadAfterWriteForAccessFlows: true
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `reason` - Indicates the reason for locking the account. Valid values are: 0 - failed password login attempts, 1 - admin lock, 2 - failed reset password attempts, 3 - failed MFA login attempts, 4 - failed MFA login attempts for federated user, 5 - failed Database login attempts

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* idcsRequiresWriteForAccessFlows: true
			* idcsRequiresImmediateReadAfterWriteForAccessFlows: true
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
	* `login_attempts` - The number of failed login attempts. The value is reset to 0 after a successful login.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsAllowUpdatesInReadOnlyMode: true
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* idcsRequiresImmediateReadAfterWriteForAccessFlows: true
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `max_concurrent_sessions` - The maximum number of concurrent sessions for a user.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsMaxValue: 999
		* idcsMinValue: 1
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `previous_successful_login_date` - The previous successful login date.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `recovery_attempts` - The number of failed recovery attempts. The value is reset to 0 after a successful login.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `recovery_enroll_attempts` - The number of failed account recovery enrollment attempts.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `recovery_locked` - A complex attribute that indicates a password recovery is locked (blocking any new sessions).

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `lock_date` - The date and time that the current resource was locked.

			**Added In:** 19.1.4

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* idcsRequiresWriteForAccessFlows: true
			* required: false
			* returned: default
			* type: dateTime
			* uniqueness: none
		* `on` - Indicates that the recovery is locked.

			**Added In:** 19.1.4

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* idcsRequiresWriteForAccessFlows: true
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionuser_user` - Oracle Identity Cloud Service User
	* `user_provider` - Registration provider

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `account_recovery_required` - Boolean value to prompt user to setup account recovery during login.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `accounts` - Accounts assigned to this User. Each value of this attribute refers to an app-specific identity that is owned by this User. Therefore, this attribute is a convenience that allows one to see on each User the Apps to which that User has access.

		**SCIM++ Properties:**
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `active` - Status of the account

			**Added In:** 17.4.6

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `app_id` - The ID of the App to which this Account gives access.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `name` - Name of the account assigned to the User.

			**Added In:** 17.4.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI of the Account assigned to the User.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The Id of the Account assigned to the User.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `app_roles` - A list of all AppRoles to which this User belongs directly, indirectly or implicitly. The User could belong directly because the User is a member of the AppRole, could belong indirectly because the User is a member of a Group that is a member of the AppRole, or could belong implicitly because the AppRole is public.

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `admin_role` - If true, then the role provides administrative access privileges. READ-ONLY.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `app_id` - The ID of the App that defines this AppRole.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `app_name` - The name (Client ID) of the App that defines this AppRole.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `display` - The display name of the AppRole assigned to the User.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `legacy_group_name` - The name (if any) under which this AppRole should appear in this User's group-memberships for reasons of backward compatibility. Oracle Identity Cloud Service distinguishes between Groups and AppRoles, but some services still expect AppRoles appear as if they were service-instance-specific Groups.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI of the AppRole assigned to the User.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - The kind of membership this User has in the AppRole. A value of 'direct' indicates that the User is a member of the AppRole.  A value of  'indirect' indicates that the User is a member of a Group that is a member of the AppRole.  A value of 'implicit' indicates that the AppRole is public.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: request
			* type: string
			* uniqueness: none
		* `value` - The Id of the AppRole assigned to the User.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
	* `applicable_authentication_target_app` - The app against which the user will authenticate. The value is not persisted but rather calculated. If the user's delegatedAuthenticationTargetApp is set, that value is returned. Otherwise, the app returned by evaluating the user's applicable Delegated Authentication Policy is returned.

		**Added In:** 18.1.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `display` - App Display Name

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - App URI

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `target_request_timeout` - Timeout interval for Synchronization TargetAction in milliseconds

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `type` - A label that indicates whether this is an App or IdentitySource.

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - App identifier

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: true
			* multiValued: false
			* mutability: readOnly
			* returned: default
			* type: string
			* uniqueness: none
	* `bypass_notification` - A Boolean value indicating whether or not to send email notification after creating the user. This attribute is not used in update/replace operations.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeNameMappings: [[columnHeaderName:ByPass Notification]]
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: never
		* type: boolean
		* uniqueness: none
	* `creation_mechanism` - User creation mechanism

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeNameMappings: [[defaultValue:import]]
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `delegated_authentication_target_app` - If set, indicates the user's preferred authentication target app. If not set and the user's \"syncedFromApp\" is set and is enabled for delegated authentication, it is used. Otherwise, the user authenticates locally to Oracle Identity Cloud Service.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - App Display Name

			**Added In:** 17.4.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - App URI

			**Added In:** 17.4.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - A label that indicates whether this is an App or IdentitySource.

			**Added In:** 17.4.6

			**SCIM++ Properties:**
			* idcsDefaultValue: IdentitySource
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - App identifier

			**Added In:** 17.4.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `do_not_show_getting_started` - A Boolean value indicating whether or not to hide the getting started page

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `grants` - Grants to this User. Each value of this attribute refers to a Grant to this User of some App (and optionally of some entitlement). Therefore, this attribute is a convenience that allows one to see on each User all of the Grants to that User.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `app_id` - The ID of the App in this Grant.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `grant_mechanism` - Each value of grantMechanism indicates how (or by what component) some App (or App-Entitlement) was granted. A customer or the UI should use only grantMechanism values that start with 'ADMINISTRATOR':
			* 'ADMINISTRATOR_TO_USER' is for a direct grant to a specific User.
			* 'ADMINISTRATOR_TO_GROUP' is for a grant to a specific Group, which results in indirect grants to Users who are members of that Group.
			* 'ADMINISTRATOR_TO_APP' is for a grant to a specific App.  The grantee (client) App gains access to the granted (server) App.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `grantor_id` - Grantor identifier

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI of this Grant to this User.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The ID of this Grant to this User.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `group_membership_last_modified` - Specifies date time when a User's group membership was last modified.

		**Added In:** 2304270343

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `idcs_app_roles_limited_to_groups` - Description:

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsCompositeKey: [value, idcsAppRoleId]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex 
		* `display` - Group display name

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `idcs_app_role_id` - The id of the Oracle Identity Cloud Service AppRole grant limited to one or more Groups.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsCsvAttributeName: IDCS AppRole Name
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `ocid` - The ocid of a Group the AppRole Grant is limited to

			**Added In:** 2202230830

			**SCIM++ Properties:**
			* idcsCsvAttributeName: Group Ocid
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI of the SCIM resource representing the Group manager.  RECOMMENDED.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The id of a Group the AppRole Grant is limited to

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsCsvAttributeName: Group Name
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `is_account_recovery_enrolled` - A Boolean value indicating whether or not a user is enrolled for account recovery

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `is_authentication_delegated` - A Boolean value indicating whether or not authentication request by this user should be delegated to a remote app. This value should be true only when the User was originally synced from an app which is enabled for delegated authentication

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: never
		* type: boolean
		* uniqueness: none
	* `is_federated_user` - A Boolean value indicating whether or not the user is federated.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Federated
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Federated]]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* idcsRequiresWriteForAccessFlows: true
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_group_membership_normalized` - A Boolean value indicating whether or not group membership is normalized for this user.

		**Deprecated Since: 19.3.3**

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: never
		* type: boolean
		* uniqueness: none
	* `is_group_membership_synced_to_users_groups` - A Boolean value Indicates whether this User's group membership has been sync'ed from Group.members to UsersGroups.

		**Added In:** 19.3.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: never
		* type: boolean
		* uniqueness: none
	* `notification_email_template_id` - Specifies the EmailTemplate to be used when sending notification to the user this request is for. If specified, it overrides the default EmailTemplate for this event.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: writeOnly
		* required: false
		* returned: never
		* type: string
		* uniqueness: none
	* `preferred_ui_landing_page` - User's preferred landing page following login, logout and reset password.

		**Added In:** 2302092332

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `status` - A supplemental status indicating the reason why a user is disabled

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `support_accounts` - A list of Support Accounts corresponding to user.

		**Added In:** 2103141444

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `user_provider` - User Support Account Provider

			**Added In:** 2103141444

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ocid` - The OCID of the user's support account.

			**Added In:** 2103141444

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `ref` - The URI of the corresponding Support Account resource to which the user belongs

			**Added In:** 2103141444

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `user_id` - User Support User Id

			**Added In:** 2103141444

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
		* `value` - The identifier of the User's support Account.

			**Added In:** 2103141444

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `synced_from_app` - Managed App or an Identity Source from where the user is synced. If enabled, this Managed App or Identity Source can be used for performing delegated authentication.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - App Display Name

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - App URI

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - A label that indicates whether this is an App or IdentitySource.

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* idcsDefaultValue: IdentitySource
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - App identifier

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `user_flow_controlled_by_external_client` - A Boolean value indicating whether to bypass notification and return user token to be used by an external client to control the user flow.

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: never
		* type: boolean
		* uniqueness: none
	* `user_token` - User token returned if userFlowControlledByExternalClient is true

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `ref` - User Token URI

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - User Token identifier

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
* `user_name` - User name

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: User ID
	* idcsCsvAttributeNameMappings: [[columnHeaderName:User Name, deprecatedColumnHeaderName:User ID]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: always
	* type: string
	* uniqueness: global
* `user_type` - Used to identify the organization-to-user relationship

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: User Type
	* idcsCsvAttributeNameMappings: [[columnHeaderName:User Type]]
	* idcsPii: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `x509certificates` - A list of certificates issued to the User.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - A human readable name, primarily used for display purposes.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `primary` - A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `type` - A label indicating the attribute's function.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - The value of a X509 certificate.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: binary
		* uniqueness: none

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the User
	* `update` - (Defaults to 20 minutes), when updating the User
	* `delete` - (Defaults to 20 minutes), when destroying the User


## Import

Users can be imported using the `id`, e.g.

```
$ terraform import oci_identity_domains_user.test_user "idcsEndpoint/{idcsEndpoint}/users/{userId}" 
```

