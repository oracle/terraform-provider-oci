---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_app"
sidebar_current: "docs-oci-resource-identity_domains-app"
description: |-
  Provides the App resource in Oracle Cloud Infrastructure Identity Domains service
---

# oci_identity_domains_app
This resource provides the App resource in Oracle Cloud Infrastructure Identity Domains service.

Create an App

## Example Usage

```hcl
resource "oci_identity_domains_app" "test_app" {
	#Required
	based_on_template {
		#Required
		value = var.app_based_on_template_value

		#Optional
		well_known_id = "CustomWebAppTemplateId"
	}
	display_name = var.app_display_name
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:App"]

	#Optional
	access_token_expiry = var.app_access_token_expiry
	active = var.app_active
	alias_apps {
		#Required
		value = var.app_alias_apps_value
	}
	all_url_schemes_allowed = var.app_all_url_schemes_allowed
	allow_access_control = var.app_allow_access_control
	allow_offline = var.app_allow_offline
	allowed_grants = var.app_allowed_grants
	allowed_operations = var.app_allowed_operations
	allowed_scopes {
		#Required
		fqs = var.app_allowed_scopes_fqs
	}
	allowed_tags {
		#Required
		key = var.app_allowed_tags_key
		value = var.app_allowed_tags_value
	}
	app_icon = var.app_app_icon
	app_signon_policy {
		#Required
		value = var.app_app_signon_policy_value
	}
	app_thumbnail = var.app_app_thumbnail
	apps_network_perimeters {
		#Required
		value = var.app_apps_network_perimeters_value
	}
	as_opc_service {
		#Required
		value = var.app_as_opc_service_value
	}
	attr_rendering_metadata {
		#Required
		name = var.app_attr_rendering_metadata_name

		#Optional
		datatype = var.app_attr_rendering_metadata_datatype
		helptext = var.app_attr_rendering_metadata_helptext
		label = var.app_attr_rendering_metadata_label
		max_length = var.app_attr_rendering_metadata_max_length
		max_size = var.app_attr_rendering_metadata_max_size
		min_length = var.app_attr_rendering_metadata_min_length
		min_size = var.app_attr_rendering_metadata_min_size
		order = var.app_attr_rendering_metadata_order
		read_only = var.app_attr_rendering_metadata_read_only
		regexp = var.app_attr_rendering_metadata_regexp
		required = var.app_attr_rendering_metadata_required
		section = var.app_attr_rendering_metadata_section
		visible = var.app_attr_rendering_metadata_visible
		widget = var.app_attr_rendering_metadata_widget
	}
	attribute_sets = ["all"]
	attributes = ""
	audience = var.app_audience
	authorization = var.app_authorization
	bypass_consent = var.app_bypass_consent
	certificates {
		#Required
		cert_alias = var.app_certificates_cert_alias
	}
	client_ip_checking = var.app_client_ip_checking
	client_type = var.app_client_type
	contact_email_address = var.app_contact_email_address
	delegated_service_names = var.app_delegated_service_names
	description = var.app_description
	disable_kmsi_token_authentication = var.app_disable_kmsi_token_authentication
	error_page_url = var.app_error_page_url
	home_page_url = var.app_home_page_url
	icon = var.app_icon
	id = var.app_id
	id_token_enc_algo = var.app_id_token_enc_algo
	identity_providers {
		#Required
		value = var.app_identity_providers_value
	}
	idp_policy {
		#Required
		value = var.app_idp_policy_value
	}
	is_alias_app = var.app_is_alias_app
	is_enterprise_app = var.app_is_enterprise_app
	is_form_fill = var.app_is_form_fill
	is_kerberos_realm = var.app_is_kerberos_realm
	is_login_target = var.app_is_login_target
	is_mobile_target = var.app_is_mobile_target
	is_multicloud_service_app = var.app_is_multicloud_service_app
	is_oauth_client = var.app_is_oauth_client
	is_oauth_resource = var.app_is_oauth_resource
	is_obligation_capable = var.app_is_obligation_capable
	is_radius_app = var.app_is_radius_app
	is_saml_service_provider = var.app_is_saml_service_provider
	is_unmanaged_app = var.app_is_unmanaged_app
	is_web_tier_policy = var.app_is_web_tier_policy
	landing_page_url = var.app_landing_page_url
	linking_callback_url = var.app_linking_callback_url
	login_mechanism = var.app_login_mechanism
	login_page_url = var.app_login_page_url
	logout_page_url = var.app_logout_page_url
	logout_uri = var.app_logout_uri
	name = var.app_name
	ocid = var.app_ocid
	post_logout_redirect_uris = var.app_post_logout_redirect_uris
	privacy_policy_url = var.app_privacy_policy_url
	product_logo_url = var.app_product_logo_url
	product_name = var.app_product_name
	protectable_secondary_audiences {
		#Required
		value = var.app_protectable_secondary_audiences_value
	}
	radius_policy {
		#Required
		value = var.app_radius_policy_value
	}
	redirect_uris = var.app_redirect_uris
	refresh_token_expiry = var.app_refresh_token_expiry
	resource_type_schema_version = var.app_resource_type_schema_version
	saml_service_provider {
		#Required
		value = var.app_saml_service_provider_value
	}
	scopes {
		#Required
		value = var.app_scopes_value

		#Optional
		description = var.app_scopes_description
		display_name = var.app_scopes_display_name
		requires_consent = var.app_scopes_requires_consent
	}
	secondary_audiences = ["secondaryAudiences"]
	service_params {
		#Required
		name = var.app_service_params_name

		#Optional
		value = var.app_service_params_value
	}
	service_type_urn = var.app_service_type_urn
	service_type_version = var.app_service_type_version
	show_in_my_apps = var.app_show_in_my_apps
	signon_policy {
		#Required
		value = var.app_signon_policy_value
	}
	tags {
		#Required
		key = var.app_tags_key
		value = var.app_tags_value
	}
	terms_of_service_url = var.app_terms_of_service_url
	terms_of_use {
		#Required
		value = var.app_terms_of_use_value
	}
	trust_policies {
		#Required
		value = var.app_trust_policies_value
	}
	trust_scope = var.app_trust_scope
	urnietfparamsscimschemasoracleidcsextension_oci_tags {

		#Optional
		defined_tags {
			#Required
			key = var.app_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_key
			namespace = var.app_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_namespace
			value = var.app_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_value
		}
		freeform_tags {
			#Required
			key = var.app_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_key
			value = var.app_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_value
		}
	}
	urnietfparamsscimschemasoracleidcsextensiondbcs_app {

		#Optional
		domain_app {
			#Required
			value = var.app_urnietfparamsscimschemasoracleidcsextensiondbcs_app_domain_app_value
		}
		domain_name = "domainName"
	}
	urnietfparamsscimschemasoracleidcsextensionenterprise_app_app {

		#Optional
		allow_authz_decision_ttl = var.app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_allow_authz_decision_ttl
		allow_authz_policy {
			#Required
			value = var.app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_allow_authz_policy_value
		}
		app_resources {
			#Required
			value = var.app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_app_resources_value
		}
		deny_authz_decision_ttl = var.app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_deny_authz_decision_ttl
		deny_authz_policy {
			#Required
			value = var.app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_deny_authz_policy_value
		}
	}
	urnietfparamsscimschemasoracleidcsextensionform_fill_app_app {

		#Optional
		configuration = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_configuration
		form_cred_method = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_cred_method
		form_credential_sharing_group_id = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_credential_sharing_group_id
		form_fill_url_match {
			#Required
			form_url = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_fill_url_match_form_url

			#Optional
			form_url_match_type = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_fill_url_match_form_url_match_type
		}
		form_type = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_type
		reveal_password_on_form = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_reveal_password_on_form
		user_name_form_expression = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_user_name_form_expression
		user_name_form_template = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_user_name_form_template
	}
	urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template {

		#Optional
		configuration = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_configuration
		form_cred_method = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_cred_method
		form_credential_sharing_group_id = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_credential_sharing_group_id
		form_fill_url_match {
			#Required
			form_url = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_fill_url_match_form_url

			#Optional
			form_url_match_type = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_fill_url_match_form_url_match_type
		}
		form_type = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_type
		reveal_password_on_form = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_reveal_password_on_form
		sync_from_template = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_sync_from_template
		user_name_form_expression = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_user_name_form_expression
		user_name_form_template = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_user_name_form_template
	}
	urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app {

		#Optional
		default_encryption_salt_type = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_default_encryption_salt_type
		master_key = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_master_key
		max_renewable_age = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_max_renewable_age
		max_ticket_life = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_max_ticket_life
		realm_name = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_realm_name
		supported_encryption_salt_types = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_supported_encryption_salt_types
		ticket_flags = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_ticket_flags
	}
	urnietfparamsscimschemasoracleidcsextensionmanagedapp_app {

		#Optional
		admin_consent_granted = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_admin_consent_granted
		bundle_configuration_properties {
			#Required
			icf_type = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_icf_type
			name = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_name
			required = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_required

			#Optional
			confidential = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_confidential
			display_name = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_display_name
			help_message = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_help_message
			order = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_order
			value = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_value
		}
		bundle_pool_configuration {

			#Optional
			max_idle = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_max_idle
			max_objects = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_max_objects
			max_wait = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_max_wait
			min_evictable_idle_time_millis = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_min_evictable_idle_time_millis
			min_idle = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_min_idle
		}
		connected = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_connected
		enable_auth_sync_new_user_notification = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_enable_auth_sync_new_user_notification
		enable_sync = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_enable_sync
		enable_sync_summary_report_notification = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_enable_sync_summary_report_notification
		flat_file_bundle_configuration_properties {
			#Required
			icf_type = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_icf_type
			name = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_name
			required = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_required

			#Optional
			confidential = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_confidential
			display_name = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_display_name
			help_message = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_help_message
			order = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_order
			value = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_value
		}
		flat_file_connector_bundle {
			#Required
			value = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_connector_bundle_value

			#Optional
			display = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_connector_bundle_display
			well_known_id = oci_identity_domains_well_known.test_well_known.id
		}
		is_authoritative = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_is_authoritative
		three_legged_oauth_credential {

			#Optional
			access_token = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_three_legged_oauth_credential_access_token
			access_token_expiry = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_three_legged_oauth_credential_access_token_expiry
			refresh_token = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_three_legged_oauth_credential_refresh_token
		}
	}
	urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app {
		#Required
		multicloud_service_type = var.app_urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app_multicloud_service_type

		#Optional
		multicloud_platform_url = var.app_urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app_multicloud_platform_url
	}
	urnietfparamsscimschemasoracleidcsextensionopc_service_app {

		#Optional
		service_instance_identifier = var.app_urnietfparamsscimschemasoracleidcsextensionopc_service_app_service_instance_identifier
	}
	urnietfparamsscimschemasoracleidcsextensionradius_app_app {
		#Required
		client_ip = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_client_ip
		include_group_in_response = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_include_group_in_response
		port = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_port
		secret_key = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_secret_key

		#Optional
		capture_client_ip = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_capture_client_ip
		country_code_response_attribute_id = "1"
		end_user_ip_attribute = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_end_user_ip_attribute
		group_membership_radius_attribute = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_group_membership_radius_attribute
		group_membership_to_return {
			#Required
			value = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_group_membership_to_return_value
		}
		group_name_format = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_group_name_format
		password_and_otp_together = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_password_and_otp_together
		radius_vendor_specific_id = "radiusVendorSpecificId"
		response_format = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_response_format
		response_format_delimiter = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_response_format_delimiter
		type_of_radius_app = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_type_of_radius_app
	}
	urnietfparamsscimschemasoracleidcsextensionrequestable_app {

		#Optional
		requestable = var.app_urnietfparamsscimschemasoracleidcsextensionrequestable_app_requestable
	}
	urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app {

		#Optional
		assertion_consumer_url = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_assertion_consumer_url
		encrypt_assertion = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_encrypt_assertion
		encryption_algorithm = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_encryption_algorithm
		encryption_certificate = "encryptionCertificate"
		federation_protocol = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_federation_protocol
		group_assertion_attributes {
			#Required
			name = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_group_assertion_attributes_name

			#Optional
			condition = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_group_assertion_attributes_condition
			format = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_group_assertion_attributes_format
			group_name = "groupName"
		}
		hok_acs_url = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_hok_acs_url
		hok_required = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_hok_required
		include_signing_cert_in_signature = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_include_signing_cert_in_signature
		key_encryption_algorithm = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_key_encryption_algorithm
		logout_binding = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_logout_binding
		logout_enabled = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_logout_enabled
		logout_request_url = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_logout_request_url
		logout_response_url = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_logout_response_url
		metadata = "metadata"
		name_id_format = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_name_id_format
		name_id_userstore_attribute = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_name_id_userstore_attribute
		partner_provider_id = "partnerProviderId"
		partner_provider_pattern = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_partner_provider_pattern
		sign_response_or_assertion = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_sign_response_or_assertion
		signature_hash_algorithm = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_signature_hash_algorithm
		signing_certificate = "signingCertificate"
		succinct_id = "succinctId"
		user_assertion_attributes {
			#Required
			name = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_user_assertion_attributes_name
			user_store_attribute_name = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_user_assertion_attributes_user_store_attribute_name

			#Optional
			format = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_user_assertion_attributes_format
		}
	}
	urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app {

		#Optional
		resource_ref = var.app_urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app_resource_ref
		web_tier_policy_az_control = var.app_urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app_web_tier_policy_az_control
		web_tier_policy_json = var.app_urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app_web_tier_policy_json
	}
}
```

## Argument Reference

The following arguments are supported:

* `access_token_expiry` - (Optional) (Updatable) Expiry-time in seconds for an Access Token. Any token that allows access to this App will expire after the specified duration.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `accounts` - (Optional) (Updatable) Accounts of App

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
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
	* `name` - (Optional) (Updatable) Name of the account

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
	* `owner_id` - (Optional) (Updatable) Owner identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) AccountMgmtInfo URI

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Optional) (Updatable) Account identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: always
		* type: string
		* uniqueness: none
* `active` - (Optional) (Updatable) If true, this App is able to participate in runtime services, such as automatic-login, OAuth, and SAML. If false, all runtime services are disabled for this App, and only administrative operations can be performed.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `admin_roles` - (Optional) (Updatable) A list of AppRoles defined by this UnmanagedApp. Membership in each of these AppRoles confers administrative privilege within this App.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* `description` - (Optional) (Updatable) The description of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `display` - (Optional) (Updatable) Display-name of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) URI of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) ID of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `alias_apps` - (Optional) (Updatable) Each value of this internal attribute refers to an Oracle Public Cloud infrastructure App on which this App depends.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `description` - (Optional) (Updatable) Description of the alias App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `display` - (Optional) (Updatable) Display name of the alias App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) URI of the alias App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) ID of the alias App.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `all_url_schemes_allowed` - (Optional) (Updatable) If true, indicates that the system should allow all URL-schemes within each value of the 'redirectUris' attribute.  Also indicates that the system should not attempt to confirm that each value of the 'redirectUris' attribute is a valid URI.  In particular, the system should not confirm that the domain component of the URI is a top-level domain and the system should not confirm that the hostname portion is a valid system that is reachable over the network.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `allow_access_control` - (Optional) (Updatable) If true, any managed App that is based on this template is checked for access control that is, access to this app is subject to successful authorization at SSO service, viz. app grants to start with.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `allow_offline` - (Optional) (Updatable) If true, indicates that the Refresh Token is allowed when this App acts as an OAuth Resource.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `allowed_grants` - (Optional) (Updatable) List of grant-types that this App is allowed to use when it acts as an OAuthClient.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `allowed_operations` - (Optional) (Updatable) OPTIONAL. Required only when this App acts as an OAuthClient. Supported values are 'introspect' and 'onBehalfOfUser'. The value 'introspect' allows the client to look inside the access-token. The value 'onBehalfOfUser' overrides how the client's privileges are combined with the privileges of the Subject User. Ordinarily, authorization calculates the set of effective privileges as the intersection of the client's privileges and the user's privileges. The value 'onBehalfOf' indicates that authorization should ignore the privileges of the client and use only the user's privileges to calculate the effective privileges.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `allowed_scopes` - (Optional) (Updatable) A list of scopes (exposed by this App or by other Apps) that this App is allowed to access when it acts as an OAuthClient.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsCompositeKey: [fqs]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `fqs` - (Required) (Updatable) A fully qualified scope that this App is allowed to access when it acts as an OAuthClient.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `id_of_defining_app` - (Optional) (Updatable) The ID of the App that defines this scope.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `read_only` - (Optional) (Updatable) If true, indicates that this value must be protected.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
* `allowed_tags` - (Optional) (Updatable) A list of tags, acting as an OAuthClient, this App is allowed to access.

	**Added In:** 17.4.6

	**SCIM++ Properties:**
	* idcsCompositeKey: [key, value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `key` - (Required) (Updatable) Key or name of the allowed tag.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `read_only` - (Optional) (Updatable) If true, indicates that this value must be protected.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `value` - (Required) (Updatable) Value of the allowed tag.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `app_icon` - (Optional) (Updatable) Application icon.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `app_signon_policy` - (Optional) (Updatable) App Sign-on Policy.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - (Optional) (Updatable) URI of the policy.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) Identifier of the Policy.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `app_thumbnail` - (Optional) (Updatable) Application thumbnail.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `apps_network_perimeters` - (Optional) (Updatable) Network Perimeter

	**Added In:** 2010242156

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - (Optional) (Updatable) URI of the Network Perimeter.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) List of identifier of Network Perimeters for App

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: always
		* type: string
		* uniqueness: none
* `as_opc_service` - (Optional) (Updatable) OPCService facet of the application.

	**Deprecated Since: 17.3.4**

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - (Optional) (Updatable) URI of the OPCService facet.

		**Deprecated Since: 17.3.4**

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) Identifier of the OPCService facet.

		**Deprecated Since: 17.3.4**

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `attr_rendering_metadata` - (Optional) (Updatable) Label for the attribute to be shown in the UI.

	**SCIM++ Properties:**
	* idcsCompositeKey: [name]
	* idcsSearchable: false
	* multiValued: true
	* mutability: immutable
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `datatype` - (Optional) (Updatable) Data type of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `helptext` - (Optional) (Updatable) Help text for the attribute. It can contain HTML tags.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `label` - (Optional) (Updatable) Label for the attribute to be shown in the UI.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `max_length` - (Optional) (Updatable) Maximum length of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_size` - (Optional) (Updatable) Maximum size of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `min_length` - (Optional) (Updatable) Minimum length of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `min_size` - (Optional) (Updatable) Minimum size of the attribute..

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `name` - (Required) (Updatable) Name of the attribute.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `order` - (Optional) (Updatable) Data type of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `read_only` - (Optional) (Updatable) Is the attribute readOnly.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `regexp` - (Optional) (Updatable) Regular expression of the attribute for validation.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `required` - (Optional) (Updatable) Attribute is required or optional.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `section` - (Optional) (Updatable) UI widget to use for the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `visible` - (Optional) (Updatable) Indicates whether the attribute is to be shown on the application creation UI.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `widget` - (Optional) (Updatable) UI widget to use for the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `attribute_sets` - (Optional) (Updatable) A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If 'attributes' query parameter is also available, union of the two is fetched. Valid values - all, always, never, request, default. Values are case-insensitive.
* `attributes` - (Optional) (Updatable) A comma-delimited string that specifies the names of resource attributes that should be returned in the response. By default, a response that contains resource attributes contains only attributes that are defined in the schema for that resource type as returned=always or returned=default. An attribute that is defined as returned=request is returned in a response only if the request specifies its name in the value of this query parameter. If a request specifies this query parameter, the response contains the attributes that this query parameter specifies, as well as any attribute that is defined as returned=always.
* `audience` - (Optional) (Updatable) The base URI for all of the scopes defined in this App. The value of 'audience' is combined with the 'value' of each scope to form an 'fqs' or fully qualified scope.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `authorization` - (Optional) (Updatable) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
* `based_on_template` - (Required) (Updatable) Application template on which the application is based.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: true
	* returned: default
	* type: complex
	* `last_modified` - (Optional) (Updatable) The most recent DateTime that the appTemplate on which the application based upon is updated. The attribute MUST be a DateTime.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
		* uniqueness: none
	* `ref` - (Optional) (Updatable) URI of the application template.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) Identifier of the application template.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: always
		* type: string
		* uniqueness: none
	* `well_known_id` - (Optional) (Updatable) Unique Well-known identifier used to reference app template.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `bypass_consent` - (Optional) (Updatable) If true, indicates that consent should be skipped for all scopes

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `callback_service_url` - (Optional) (Updatable) Callback Service URL

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `certificates` - (Optional) (Updatable) Each value of this attribute represent a certificate that this App uses when it acts as an OAuthClient.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [certAlias]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `cert_alias` - (Required) (Updatable) Certificate alias

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: always
		* type: string
		* uniqueness: none
	* `kid` - (Optional) (Updatable) Certificate kid

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `sha1thumbprint` - (Optional) (Updatable) sha1Thumbprint

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `x509base64certificate` - (Optional) (Updatable) Base-64-encoded certificate.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: binary
		* uniqueness: none
	* `x5t` - (Optional) (Updatable) Certificate x5t

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `client_ip_checking` - (Optional) (Updatable) Network Perimeters checking mode

	**Added In:** 2010242156

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `client_secret` - (Optional) (Updatable) This value is the credential of this App, which this App supplies as a password when this App authenticates to the Oracle Public Cloud infrastructure. This value is also the client secret of this App when it acts as an OAuthClient.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* idcsSensitive: none
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `client_type` - (Optional) (Updatable) Specifies the type of access that this App has when it acts as an OAuthClient.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `cloud_control_properties` - (Optional) (Updatable) A collection of arbitrary properties that scope the privileges of a cloud-control App.

	**Added In:** 18.4.2

	**SCIM++ Properties:**
	* idcsCompositeKey: [name]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `name` - (Required) (Updatable) The name of the property.

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `values` - (Required) (Updatable) The value(s) of the property.

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: true
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
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
* `contact_email_address` - (Optional) (Updatable) Contact Email Address

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `delegated_service_names` - (Optional) (Updatable) Service Names allow to use Oracle Cloud Infrastructure signature for client authentication instead of client credentials

	**Added In:** 2207040824

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
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
* `description` - (Optional) (Updatable) Description of the application.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `disable_kmsi_token_authentication` - (Optional) (Updatable) Indicates whether the application is allowed to be access using kmsi token.

	**Added In:** 2111190457

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: always
	* type: boolean
	* uniqueness: none
* `display_name` - (Required) (Updatable) Display name of the application. Display name is intended to be user-friendly, and an administrator can change the value at any time.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: always
	* type: string
	* uniqueness: server
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
* `editable_attributes` - (Optional) (Updatable) App attributes editable by subject

	**Added In:** 18.2.6

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [name]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `name` - (Required) (Updatable) Name of the attribute. The attribute name will be qualified by schema name if any extension schema defines the attribute. The attribute name will not be qualified by schema name if the base schema defines the attribute.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `error_page_url` - (Optional) (Updatable) This attribute specifies the URL of the page to which an application will redirect an end-user in case of error.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `granted_app_roles` - (Optional) (Updatable) A list of AppRoles that are granted to this App (and that are defined by other Apps). Within the Oracle Public Cloud infrastructure, this allows AppID-based association. Such an association allows this App to act as a consumer and thus to access resources of another App that acts as a producer.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `admin_role` - (Optional) (Updatable) If true, then this granted AppRole confers administrative privileges within the App that defines it. Otherwise, the granted AppRole confers only functional privileges.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `app_id` - (Optional) (Updatable) The id of the App that defines this AppRole, which is granted to this App. The App that defines the AppRole acts as the producer; the App to which the AppRole is granted acts as a consumer.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `app_name` - (Optional) (Updatable) The name of the App that defines this AppRole, which is granted to this App. The App that defines the AppRole acts as the producer; the App to which the AppRole is granted acts as a consumer.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `display` - (Optional) (Updatable) The display-name of an AppRole that is granted to this App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `legacy_group_name` - (Optional) (Updatable) The name of the legacy group associated with this AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `read_only` - (Optional) (Updatable) If true, indicates that this value must be protected.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `ref` - (Optional) (Updatable) The URI of an AppRole that is granted to this App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - (Optional) (Updatable) A label that indicates whether this AppRole was granted directly to the App (or indirectly through a Group). For an App, the value of this attribute will always be 'direct' (because an App cannot be a member of a Group).

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) The id of an AppRole that is granted to this App.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `grants` - (Optional) (Updatable) Grants assigned to the app

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
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
	* `grantee_id` - (Optional) (Updatable) Grantee identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `grantee_type` - (Optional) (Updatable) Grantee resource type. Allowed values are User and Group.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) Grant URI

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Optional) (Updatable) Grant identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `hashed_client_secret` - (Optional) (Updatable) Hashed Client Secret. This hash-value is used to verify the 'clientSecret' credential of this App

	**Added In:** 2106240046

	**SCIM++ Properties:**
	* idcsSearchable: false
	* idcsSensitive: hash_sc
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `home_page_url` - (Optional) (Updatable) Home Page URL

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `icon` - (Optional) (Updatable) URL of application icon.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: reference
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
* `id_token_enc_algo` - (Optional) (Updatable) Encryption Alogrithm to use for encrypting ID token.

	**Added In:** 2010242156

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
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
* `identity_providers` - (Optional) (Updatable) A list of IdentityProvider assigned to app. A user trying to access this app will be automatically redirected to configured IdP during the authentication phase, before being able to access App.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: request
	* type: complex
	* `display` - (Optional) (Updatable) Display-name of the IdentityProvider.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) URI of the IdentityProvider.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) ID of the IdentityProvider.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `idp_policy` - (Optional) (Updatable) IDP Policy.

	**Added In:** 18.1.2

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - (Optional) (Updatable) URI of the policy.

		**Added In:** 18.1.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) Identifier of the Policy.

		**Added In:** 18.1.2

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `infrastructure` - (Optional) (Updatable) If true, this App is an internal infrastructure App.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_alias_app` - (Optional) If true, this App is an AliasApp and it cannot be granted to an end-user directly.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: always
	* type: boolean
	* uniqueness: none
* `is_database_service` - (Optional) (Updatable) If true, this application acts as database service Application

	**Added In:** 18.2.2

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* type: boolean
* `is_enterprise_app` - (Optional) (Updatable) If true, this app acts as Enterprise app with Authentication and URL Authz policy.

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_form_fill` - (Optional) (Updatable) If true, this application acts as FormFill Application

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_kerberos_realm` - (Optional) (Updatable) If true, indicates that this App supports Kerberos Authentication

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_login_target` - (Optional) (Updatable) If true, this App allows runtime services to log end users into this App automatically.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_managed_app` - (Optional) (Updatable) If true, indicates that access to this App requires an account. That is, in order to log in to the App, a User must use an application-specific identity that is maintained in the remote identity-repository of that App.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_mobile_target` - (Optional) (Updatable) If true, indicates that the App should be visible in each end-user's mobile application.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_multicloud_service_app` - (Optional) (Updatable) If true, indicates the app is used for multicloud service integration.

	**Added In:** 2301202328

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_oauth_client` - (Optional) (Updatable) If true, this application acts as an OAuth Client

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_oauth_resource` - (Optional) (Updatable) If true, indicates that this application acts as an OAuth Resource.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_opc_service` - (Optional) (Updatable) If true, this application is an Oracle Public Cloud service-instance.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_obligation_capable` - (Optional) (Updatable) This flag indicates if the App is capable of validating obligations with the token for allowing access to the App.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_radius_app` - (Optional) (Updatable) If true, this application acts as an Radius App

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_saml_service_provider` - (Optional) (Updatable) If true, then this App acts as a SAML Service Provider.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_unmanaged_app` - (Optional) If true, indicates that this application accepts an Oracle Cloud Identity Service User as a login-identity (does not require an account) and relies for authorization on the User's memberships in AppRoles.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_web_tier_policy` - (Optional) (Updatable) If true, the webtier policy is active

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `landing_page_url` - (Optional) (Updatable) The URL of the landing page for this App, which is the first page that an end user should see if runtime services log that end user in to this App automatically.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `linking_callback_url` - (Optional) (Updatable) This attribute specifies the callback URL for the social linking operation.

	**Added In:** 18.2.4

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `login_mechanism` - (Optional) (Updatable) The protocol that runtime services will use to log end users in to this App automatically. If 'OIDC', then runtime services use the OpenID Connect protocol. If 'SAML', then runtime services use Security Assertion Markup Language protocol.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `login_page_url` - (Optional) (Updatable) This attribute specifies the URL of the page that the App uses when an end-user signs in to that App.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `logout_page_url` - (Optional) (Updatable) This attribute specifies the URL of the page that the App uses when an end-user signs out.

	**Added In:** 17.4.2

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `logout_uri` - (Optional) (Updatable) OAuth will use this URI to logout if this App wants to participate in SSO, and if this App's session gets cleared as part of global logout. Note: This attribute is used only if this App acts as an OAuthClient.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
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
* `meter_as_opc_service` - (Optional) (Updatable) Indicates whether the application is billed as an OPCService. If true, customer is not billed for runtime operations of the app.

	**Added In:** 18.4.2

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: always
	* type: boolean
	* uniqueness: none
* `migrated` - (Optional) (Updatable) If true, this App was migrated from an earlier version of Oracle Public Cloud infrastructure (and may therefore require special handling from runtime services such as OAuth or SAML). If false, this App requires no special handling from runtime services.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `name` - (Optional) Name of the application. Also serves as username if the application authenticates to Oracle Public Cloud infrastructure. This name may not be user-friendly and cannot be changed once an App is created.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: string
	* uniqueness: server
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
* `post_logout_redirect_uris` - (Optional) (Updatable) Each value of this attribute is the URI of a landing page within this App. It is used only when this App, acting as an OAuthClient, initiates the logout flow and wants to be redirected back to one of its landing pages.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `privacy_policy_url` - (Optional) (Updatable) Privacy Policy URL

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `product_logo_url` - (Optional) (Updatable) Application Logo URL

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `product_name` - (Optional) (Updatable) Product Name

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `protectable_secondary_audiences` - (Optional) (Updatable) A list of secondary audiences--additional URIs to be added automatically to any OAuth token that allows access to this App. Note: This attribute is used mainly for backward compatibility in certain Oracle Public Cloud Apps.

	**Added In:** 18.2.2

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `read_only` - (Optional) (Updatable) If true, indicates that this value must be protected.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `value` - (Required) (Updatable) The value of an secondary audience--additional URI to be added automatically to any OAuth token that allows access to this App.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `radius_policy` - (Optional) (Updatable) RADIUS Policy assigned to this application.

	**Added In:** 2209070044

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - (Optional) (Updatable) URI of the policy.

		**Added In:** 2209070044

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) Identifier of the Policy.

		**Added In:** 2209070044

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `ready_to_upgrade` - (Optional) (Updatable) If true, this App requires an upgrade and mandates attention from application administrator. The flag is used by UI to indicate this app is ready to upgrade.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `redirect_uris` - (Optional) (Updatable) OPTIONAL. Each value is a URI within this App. This attribute is required when this App acts as an OAuthClient and is involved in three-legged flows (authorization-code flows).

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `refresh_token_expiry` - (Optional) (Updatable) Expiry-time in seconds for a Refresh Token.  Any token that allows access to this App, once refreshed, will expire after the specified duration.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `resource_type_schema_version` - (Optional) (Updatable) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `saml_service_provider` - (Optional) (Updatable) An attribute that refers to the SAML Service Provider that runtime services will use to log an end user in to this App automatically. Note that this will be used only if the loginMechanism is 'SAML'.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `ref` - (Optional) (Updatable) The URI of the App that acts as a Service Provider.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) The id of the App that acts as a Service Provider.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: always
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
* `scopes` - (Optional) (Updatable) Scopes defined by this App. Used when this App acts as an OAuth Resource.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `description` - (Optional) (Updatable) OAuth scope description

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `display_name` - (Optional) (Updatable) OAuth scope display name

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `fqs` - (Optional) (Updatable) The fully qualified value of this scope within this App. A fully qualified scope combines the 'value' of each scope with the value of 'audience'. Each value of 'fqs' must be unique across the system. Used only when this App acts as an OAuth Resource.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: server
	* `read_only` - (Optional) (Updatable) If true, indicates that this value must be protected.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `requires_consent` - (Optional) (Updatable) If true, indicates that a user must provide consent to access this scope. Note: Used only when this App acts as an OAuth Resource.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `value` - (Required) (Updatable) OAuth scope.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `secondary_audiences` - (Optional) (Updatable) A list of secondary audiences--additional URIs to be added automatically to any OAuth token that allows access to this App. Note: This attribute is used mainly for backward compatibility in certain Oracle Public Cloud Apps.

	**Deprecated Since: 18.2.6**

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `service_params` - (Optional) (Updatable) Custom attribute that is required to compute other attribute values during app creation.

	**SCIM++ Properties:**
	* idcsCompositeKey: [name]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: always
	* type: complex
	* uniqueness: none
	* `name` - (Required) (Updatable) The name of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Optional) (Updatable) The value of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `service_type_urn` - (Optional) (Updatable) This Uniform Resource Name (URN) value identifies the type of Oracle Public Cloud service of which this app is an instance.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `service_type_version` - (Optional) (Updatable) This value specifies the version of the Oracle Public Cloud service of which this App is an instance

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `show_in_my_apps` - (Optional) (Updatable) If true, this app will be displayed in the MyApps page of each end-user who has access to the App.

	**Added In:** 18.1.2

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `signon_policy` - (Optional) (Updatable) Sign-on Policy.

	**Deprecated Since: 17.3.4**

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - (Optional) (Updatable) URI of the policy.

		**Deprecated Since: 17.3.4**

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) Identifier of the Policy.

		**Deprecated Since: 17.3.4**

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
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
* `terms_of_service_url` - (Optional) (Updatable) Terms of Service URL

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `terms_of_use` - (Optional) (Updatable) Terms Of Use.

	**Added In:** 18.2.6

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `name` - (Optional) (Updatable) Terms Of Use name

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) URI of the TermsOfUse.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) Identifier of the TermsOfUse

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
* `trust_policies` - (Optional) (Updatable) Trust Policies.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - (Optional) (Updatable) URI of the policy.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) Identifier of the Policy.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `trust_scope` - (Optional) (Updatable) Indicates the scope of trust for this App when acting as an OAuthClient. A value of 'Explicit' indicates that the App is allowed to access only the scopes of OAuthResources that are explicitly specified as 'allowedScopes'. A value of 'Account' indicates that the App is allowed implicitly to access any scope of any OAuthResource within the same Oracle Cloud Account. A value of 'Tags' indicates that the App is allowed to access any scope of any OAuthResource with a matching tag within the same Oracle Cloud Account. A value of 'Default' indicates that the Tenant default trust scope configured in the Tenant Settings is used.

	**Added In:** 17.4.2

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
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
* `urnietfparamsscimschemasoracleidcsextensiondbcs_app` - (Optional) (Updatable) This extension provides attributes for database service facet of an App
	* `domain_app` - (Optional) (Updatable) Description:

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none 
		* `display` - (Optional) (Updatable) DB Domain App display name

			**Added In:** 18.2.2

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: request
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) DB Domain App URI

			**Added In:** 18.2.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) DB Domain App identifier

			**Added In:** 18.2.2

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `domain_name` - (Optional) (Updatable) The name of the Enterprise Domain that contains any number of DBInstances. If specified, the value must be unique.  A non-null value indicates that App represents a DBDomain. A value of null indicates that the App represents an DB-instance.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: string
		* uniqueness: server
* `urnietfparamsscimschemasoracleidcsextensionenterprise_app_app` - (Optional) (Updatable) This extension defines the Enterprise App related attributes.
	* `allow_authz_decision_ttl` - (Optional) (Updatable) Allow Authz policy decision expiry time in seconds.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsMaxValue: 3600
		* idcsMinValue: 0
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `allow_authz_policy` - (Optional) (Updatable) Allow Authz Policy.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* `ref` - (Optional) (Updatable) URI of the policy.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) Identifier of the Policy.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `app_resources` - (Optional) (Updatable) A list of AppResources of this App.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: true
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* `ref` - (Optional) (Updatable) The URI of an AppResource of this App.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) The id of an AppResource of this App.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `deny_authz_decision_ttl` - (Optional) (Updatable) Deny Authz policy decision expiry time in seconds.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsMaxValue: 3600
		* idcsMinValue: 0
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `deny_authz_policy` - (Optional) (Updatable) Deny Authz Policy.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* `ref` - (Optional) (Updatable) URI of the policy.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) Identifier of the Policy.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionform_fill_app_app` - (Optional) (Updatable) This extension provides attributes for Form-Fill facet of App
	* `configuration` - (Optional) (Updatable) FormFill Application Configuration CLOB which has to be maintained in Form-Fill APP for legacy code to do Form-Fill injection

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `form_cred_method` - (Optional) (Updatable) Indicates how FormFill obtains the username and password of the account that FormFill will use to sign into the target App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `form_credential_sharing_group_id` - (Optional) (Updatable) Credential Sharing Group to which this form-fill application belongs.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `form_fill_url_match` - (Optional) (Updatable) A list of application-formURLs that FormFill should match against any formUrl that the user-specifies when signing in to the target service.  Each item in the list also indicates how FormFill should interpret that formUrl.

		**SCIM++ Properties:**
		* idcsCompositeKey: [formUrl]
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `form_url` - (Required) (Updatable) An application formUrl that FormFill will match against any formUrl that a User enters in trying to access the target-service which this App represents.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `form_url_match_type` - (Optional) (Updatable) Indicates how to interpret the value of 'formUrl' when matching against a user-specified formUrl.  The system currently supports only 'Exact', which indicates that the value of 'formUrl' should be treated as a literal value.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `form_type` - (Optional) (Updatable) Type of the FormFill application like WebApplication, MainFrameApplication, WindowsApplication. Initially, we will support only WebApplication.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `reveal_password_on_form` - (Optional) (Updatable) If true, indicates that system is allowed to show the password in plain-text for this account after re-authentication.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `sync_from_template` - (Optional) (Updatable) If true, indicates that each of the Form-Fill-related attributes that can be inherited from the template actually will be inherited from the template. If false, indicates that the AppTemplate on which this App is based has disabled inheritance for these Form-Fill-related attributes.

		**Added In:** 17.4.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `user_name_form_expression` - (Optional) (Updatable) Indicates the custom expression, which can combine concat and substring operations with literals and with any attribute of the Oracle Identity Cloud Service User

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `user_name_form_template` - (Optional) (Updatable) Format for generating a username.  This value can be Username or Email Address; any other value will be treated as a custom expression.  A custom expression may combine 'concat' and 'substring' operations with literals and with any attribute of the Oracle Identity Cloud Service user.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template` - (Optional) (Updatable) This extension provides attributes for Form-Fill facet of AppTemplate
	* `configuration` - (Optional) (Updatable) FormFill Application Configuration CLOB which has to be maintained in Form-Fill APP for legacy code to do Form-Fill injection

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `form_cred_method` - (Optional) (Updatable) Indicates how FormFill obtains the username and password of the account that FormFill will use to sign into the target App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `form_credential_sharing_group_id` - (Optional) (Updatable) Credential Sharing Group to which this form-fill application belongs.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `form_fill_url_match` - (Optional) (Updatable) A list of application-formURLs that FormFill should match against any formUrl that the user-specifies when signing in to the target service.  Each item in the list also indicates how FormFill should interpret that formUrl.

		**SCIM++ Properties:**
		* idcsCompositeKey: [formUrl]
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `form_url` - (Required) (Updatable) An application formUrl that FormFill will match against any formUrl that a User enters in trying to access the target-service which this App represents.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `form_url_match_type` - (Optional) (Updatable) Indicates how to interpret the value of 'formUrl' when matching against a user-specified formUrl.  The system currently supports only 'Exact', which indicates that the value of 'formUrl' should be treated as a literal value.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `form_type` - (Optional) (Updatable) Type of the FormFill application like WebApplication, MainFrameApplication, WindowsApplication. Initially, we will support only WebApplication.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `reveal_password_on_form` - (Optional) (Updatable) If true, indicates that system is allowed to show the password in plain-text for this account after re-authentication.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `sync_from_template` - (Optional) (Updatable) If true, indicates that each of the Form-Fill-related attributes that can be inherited from the template actually will be inherited from the template. If false, indicates that the AppTemplate disabled inheritance for these Form-Fill-related attributes.

		**Added In:** 17.4.2

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `user_name_form_expression` - (Optional) (Updatable) Indicates the custom expression, which can combine concat and substring operations with literals and with any attribute of the Oracle Identity Cloud Service User

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `user_name_form_template` - (Optional) (Updatable) Format for generating a username.  This value can be Username or Email Address; any other value will be treated as a custom expression.  A custom expression may combine 'concat' and 'substring' operations with literals and with any attribute of the Oracle Identity Cloud Service user.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app` - (Optional) (Updatable) Kerberos Realm
	* `default_encryption_salt_type` - (Optional) (Updatable) The type of salt that the system will use to encrypt Kerberos-specific artifacts of this App unless another type of salt is specified.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `master_key` - (Optional) (Updatable) The primary key that the system should use to encrypt artifacts that are specific to this Kerberos realm -- for example, to encrypt the Principal Key in each KerberosRealmUser.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsSensitive: none
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `max_renewable_age` - (Optional) (Updatable) Max Renewable Age in seconds

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `max_ticket_life` - (Optional) (Updatable) Max Ticket Life in seconds

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `realm_name` - (Optional) (Updatable) The name of the Kerberos Realm that this App uses for authentication.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `supported_encryption_salt_types` - (Optional) (Updatable) The types of salt that are available for the system to use when encrypting Kerberos-specific artifacts for this App.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `ticket_flags` - (Optional) (Updatable) Ticket Flags

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionmanagedapp_app` - (Optional) (Updatable) Managed App
	* `account_form_visible` - (Optional) (Updatable) If true, then the account form will be displayed in the Oracle Identity Cloud Service UI to interactively create or update an account for this App. If a value is not specified for this attribute, a default value of \"false\" will be assumed as the value for this attribute.

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `admin_consent_granted` - (Optional) (Updatable) If true, admin has granted consent to perform managed app run-time operations.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `bundle_configuration_properties` - (Optional) (Updatable) ConnectorBundle configuration properties

		**SCIM++ Properties:**
		* idcsCompositeKey: [name]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `confidential` - (Optional) (Updatable) If true, this bundle configuration property value is confidential and will be encrypted in Oracle Identity Cloud Service. This attribute maps to \"isConfidential\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `display_name` - (Optional) (Updatable) Display name of the bundle configuration property. This attribute maps to \"displayName\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `help_message` - (Optional) (Updatable) Help message of the bundle configuration property. This attribute maps to \"helpMessage\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `icf_type` - (Required) (Updatable) ICF data type of the bundle configuration property. This attribute maps to \"type\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `name` - (Required) (Updatable) Name of the bundle configuration property. This attribute maps to \"name\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `order` - (Optional) (Updatable) Display sequence of the bundle configuration property.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `required` - (Required) (Updatable) If true, this bundle configuration property is required to connect to the target connected managed app. This attribute maps to \"isRequired\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: boolean
			* uniqueness: none
		* `value` - (Optional) (Updatable) Value of the bundle configuration property. This attribute maps to \"value\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* idcsSensitive: encrypt
			* multiValued: true
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `bundle_pool_configuration` - (Optional) (Updatable) Configurable options maintaining a pool of ICF connector instances. Values for sub attributes can be set only if the ConnectorBundle referenced in the App has connectorPoolingSupported set to true

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `max_idle` - (Optional) (Updatable) Maximum number of connector instances in the pool that are idle and active.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `max_objects` - (Optional) (Updatable) Maximum number of connector instances in the pool that are idle and active.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `max_wait` - (Optional) (Updatable) Maximum time (in milliseconds) to wait for a free connector instance to become available before failing.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `min_evictable_idle_time_millis` - (Optional) (Updatable) Minimum time (in milliseconds) to wait before evicting an idle conenctor instance from the pool.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `min_idle` - (Optional) (Updatable) Minimum number of idle connector instances in the pool.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
	* `can_be_authoritative` - (Optional) (Updatable) If true, the managed app can be authoritative.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `connected` - (Optional) (Updatable) If true, the accounts of the application are managed through an ICF connector bundle

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `connector_bundle` - (Optional) (Updatable) ConnectorBundle

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - (Optional) (Updatable) ConnectorBundle display name

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) ConnectorBundle URI

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - (Required) (Updatable) Connector Bundle type. Allowed values are ConnectorBundle, LocalConnectorBundle.

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: true
			* idcsDefaultValue: ConnectorBundle
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - (Required) (Updatable) ConnectorBundle identifier

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `well_known_id` - (Optional) (Updatable) Unique Well-known identifier used to reference connector bundle.

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `enable_auth_sync_new_user_notification` - (Optional) (Updatable) If true, send activation email to new users created from authoritative sync.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `enable_sync` - (Optional) (Updatable) If true, sync run-time operations are enabled for this App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `enable_sync_summary_report_notification` - (Optional) (Updatable) If true, send sync summary as notification upon job completion.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `flat_file_bundle_configuration_properties` - (Optional) (Updatable) Flat file connector bundle configuration properties

		**SCIM++ Properties:**
		* idcsCompositeKey: [name]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `confidential` - (Optional) (Updatable) If true, this flatfile bundle configuration property value is confidential and will be encrypted in Oracle Identity Cloud Service. This attribute maps to \"isConfidential\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `display_name` - (Optional) (Updatable) Display name of the flatfile bundle configuration property. This attribute maps to \"displayName\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `help_message` - (Optional) (Updatable) Help message of the flatfile bundle configuration property. This attribute maps to \"helpMessage\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `icf_type` - (Required) (Updatable) ICF data type of flatfile the bundle configuration property. This attribute maps to \"type\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `name` - (Required) (Updatable) Name of the flatfile bundle configuration property. This attribute maps to \"name\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `order` - (Optional) (Updatable) Display sequence of the bundle configuration property.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `required` - (Required) (Updatable) If true, this flatfile bundle configuration property is required to connect to the target connected managed app. This attribute maps to \"isRequired\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: boolean
			* uniqueness: none
		* `value` - (Optional) (Updatable) Value of the flatfile bundle configuration property. This attribute maps to \"value\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* idcsSensitive: encrypt
			* multiValued: true
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `flat_file_connector_bundle` - (Optional) (Updatable) Flat file connector bundle to sync from a flat file.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - (Optional) (Updatable) ConnectorBundle display name

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) ConnectorBundle URI

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) ConnectorBundle identifier

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `well_known_id` - (Optional) (Updatable) Unique well-known identifier used to reference connector bundle.

			**Added In:** 19.1.4

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `identity_bridges` - (Optional) (Updatable) IdentityBridges associated with this App

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `name` - (Optional) (Updatable) Name of the IdentityBridge associated with the App.

			**Added In:** 19.1.4

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the IdentityBridge associated with the App.

			**Added In:** 19.1.4

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Optional) (Updatable) The Id of the IdentityBridge associated with the App.

			**Added In:** 19.1.4

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `is_authoritative` - (Optional) (Updatable) If true, sync from the managed app will be performed as authoritative sync.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_directory` - (Optional) (Updatable) If true, the managed app is a directory.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_on_premise_app` - (Optional) (Updatable) If true, the managed app is an On-Premise app.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_schema_customization_supported` - (Optional) (Updatable) If true, the managed app supports schema customization.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_schema_discovery_supported` - (Optional) (Updatable) If true, the managed app supports schema discovery.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_three_legged_oauth_enabled` - (Optional) (Updatable) If true, the managed app requires 3-legged OAuth for authorization.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_two_legged_oauth_enabled` - (Optional) (Updatable) If true, indicates that Oracle Identity Cloud Service can use two-legged OAuth to connect to this ManagedApp.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `object_classes` - (Optional) (Updatable) Object classes

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - (Optional) (Updatable) Object class display name

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `is_account_object_class` - (Optional) (Updatable) If true, the object class represents an account. The isAccountObjectClass attribute value 'true' MUST appear no more than once.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `ref` - (Optional) (Updatable) Object class URI

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `resource_type` - (Optional) (Updatable) Object class resource type

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `type` - (Required) (Updatable) Object Class type. Allowed values are AccountObjectClass, ManagedObjectClass.

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsDefaultValue: AccountObjectClass
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - (Required) (Updatable) Object class template identifier

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `sync_config_last_modified` - (Optional) (Updatable) The most recent DateTime that the configuration of this App was updated. AppServices updates this timestamp whenever AppServices updates an App's configuration with respect to synchronization.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `three_legged_oauth_credential` - (Optional) (Updatable) The value of this attribute persists any OAuth access token that the system uses to connect to this ManagedApp. The system obtains this access token using an OAuth protocol flow that could be two-legged or three-legged. A two-legged flow involves only the requester and the server. A three-legged flow also requires the consent of a user -- in this case the consent of an administrator.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* `access_token` - (Optional) (Updatable) Access Token

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* idcsSensitive: encrypt
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `access_token_expiry` - (Optional) (Updatable) Access token expiry

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: dateTime
			* uniqueness: none
		* `refresh_token` - (Optional) (Updatable) Refresh Token

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* idcsSensitive: encrypt
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `three_legged_oauth_provider_name` - (Optional) (Updatable) Three legged OAuth provider name in Oracle Identity Cloud Service.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
* `urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app` - (Optional) (Updatable) This extension defines attributes specific to Apps that represent instances of Multicloud Service App
	* `multicloud_platform_url` - (Optional) (Updatable) The multicloud platform service URL which the application will invoke for runtime operations such as AWSCredentials api invocation

		**Added In:** 2301202328

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `multicloud_service_type` - (Required) (Updatable) Specifies the service type for which the application is configured for multicloud integration. For applicable external service types, app will invoke multicloud service for runtime operations

		**Added In:** 2301202328

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: request
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionopc_service_app` - (Optional) (Updatable) This extension defines attributes specific to Apps that represent instances of an Oracle Public Cloud (OPC) service.
	* `current_federation_mode` - (Optional) (Updatable) Current Federation Mode

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
	* `current_synchronization_mode` - (Optional) (Updatable) Current Synchronization Mode

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
	* `enabling_next_fed_sync_modes` - (Optional) (Updatable) If true, indicates that enablement is in progress started but not completed

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
	* `next_federation_mode` - (Optional) (Updatable) Next Federation Mode

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
	* `next_synchronization_mode` - (Optional) (Updatable) Next Synchronization Mode

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
	* `region` - (Optional) (Updatable) This value identifies the OPC region in which the service is running.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `service_instance_identifier` - (Optional) (Updatable) This value specifies the unique identifier assigned to an instance of an Oracle Public Cloud service app.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: string
		* uniqueness: server
* `urnietfparamsscimschemasoracleidcsextensionradius_app_app` - (Optional) (Updatable) This extension defines attributes specific to Apps that represent instances of Radius App.
	* `capture_client_ip` - (Optional) (Updatable) If true, capture the client IP address from the RADIUS request packet. IP Address is used for auditing, policy-evaluation and country-code calculation.

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
	* `client_ip` - (Required) (Updatable) This is the IP address of the RADIUS Client like Oracle Database server. It can be only IP address and not hostname.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `country_code_response_attribute_id` - (Optional) (Updatable) Vendor-specific identifier of the attribute in the RADIUS response that will contain the end-user's country code. This is an integer-value in the range 1 to 255

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `end_user_ip_attribute` - (Optional) (Updatable) The name of the attribute that contains the Internet Protocol address of the end-user.

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `group_membership_radius_attribute` - (Optional) (Updatable) RADIUS attribute that RADIUS-enabled system uses to pass the group membership

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `group_membership_to_return` - (Optional) (Updatable) In a successful authentication response, Oracle Identity Cloud Service will pass user's group information restricted to groups persisted in this attribute, in the specified RADIUS attribute.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `display` - (Optional) (Updatable) A human readable name, primarily used for display purposes. READ-ONLY.

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) The URI of the corresponding Group resource to which the user belongs

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) The identifier of the User's group.

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
	* `group_name_format` - (Optional) (Updatable) Configure the groupNameFormat based on vendor in order to pass it to RADIUS infra

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `include_group_in_response` - (Required) (Updatable) Indicates to include groups in RADIUS response

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
	* `password_and_otp_together` - (Optional) (Updatable) Indicates if password and OTP are passed in the same sign-in request or not.

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
	* `port` - (Required) (Updatable) This is the port of RADIUS Proxy which RADIUS client will connect to.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `radius_vendor_specific_id` - (Optional) (Updatable) ID used to identify a particular vendor.

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `response_format` - (Optional) (Updatable) Configure the responseFormat based on vendor in order to pass it to RADIUS infra

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `response_format_delimiter` - (Optional) (Updatable) The delimiter used if group membership responseFormat is a delimited list instead of repeating attributes

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `secret_key` - (Required) (Updatable) Secret key used to secure communication between RADIUS Proxy and RADIUS client

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `type_of_radius_app` - (Optional) (Updatable) Value consists of type of RADIUS App. Type can be Oracle Database, VPN etc

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
* `urnietfparamsscimschemasoracleidcsextensionrequestable_app` - (Optional) (Updatable) Requestable App
	* `requestable` - (Optional) (Updatable) Flag controlling whether resource can be request by user through self service console.

		**Added In:** 17.3.4

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app` - (Optional) (Updatable) This extension defines attributes related to the Service Providers configuration.
	* `assertion_consumer_url` - (Optional) (Updatable) The attribute represents the URL to which the SAML Assertions will be sent by the SAML IdP.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `encrypt_assertion` - (Optional) (Updatable) If true, indicates that the system must encrypt the Security Assertion Markup Language (SAML) assertion.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `encryption_algorithm` - (Optional) (Updatable) This attribute indicates the encryption algorithm used to encrypt the SAML assertion.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `encryption_certificate` - (Optional) (Updatable) This attribute represents the encryption certificate that an App uses to encrypt the Security Assertion Markup Language (SAML) assertion.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `federation_protocol` - (Optional) (Updatable) Specifies the preferred federation protocol (SAML2.0 or WS-Fed1.1).

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: true
		* idcsDefaultValue: SAML2.0
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `group_assertion_attributes` - (Optional) (Updatable) Each value of this attribute describes an attribute of Group that will be sent in a Security Assertion Markup Language (SAML) assertion.

		**Deprecated Since: 18.2.2**

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCompositeKey: [name]
		* idcsSearchable: false
		* idcsValuePersistedInOtherAttribute: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `condition` - (Optional) (Updatable) Indicates the filter types that are supported for the Group assertion attributes.

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `format` - (Optional) (Updatable) Indicates the format of the assertion attribute.

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `group_name` - (Optional) (Updatable) Indicates the group name that are supported for the group assertion attributes.

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `name` - (Required) (Updatable) The attribute represents the name of the attribute that will be used in the Security Assertion Markup Language (SAML) assertion

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `hok_acs_url` - (Optional) (Updatable) Hok Assertion Consumer Service Url

		**Added In:** 2101262133

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `hok_required` - (Optional) (Updatable) If enabled, then the SAML Service supports Hok for this App.

		**Added In:** 2101262133

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `include_signing_cert_in_signature` - (Optional) (Updatable) If true, then the signing certificate is included in the signature.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `key_encryption_algorithm` - (Optional) (Updatable) This attribute indicates the key encryption algorithm.

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `last_notification_sent_time` - (Optional) (Updatable) Records the notification timestamp for the SP whose signing certificate is about to expire.

		**Added In:** 2302092332

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
		* uniqueness: none
	* `logout_binding` - (Optional) (Updatable) This attribute represents the HTTP binding that would be used while logout.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `logout_enabled` - (Optional) (Updatable) If true, then the SAML Service supports logout for this App.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `logout_request_url` - (Optional) (Updatable) The URL to which the partner sends the logout request.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `logout_response_url` - (Optional) (Updatable) The URL to which the partner sends the logout response.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `metadata` - (Optional) (Updatable) This attribute represents the metadata of a Security Provider in the Security Assertion Markup Language protocol.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `name_id_format` - (Optional) (Updatable) This can be any string, but there are a set of standard nameIdFormats. If a nameIdFormat other than the standard list is chosen, it will be considered a custom nameidformat. The standard nameidformats include: saml-x509, saml-emailaddress, saml-windowsnamequalifier, saml-kerberos, saml-persistent, saml-transient, saml-unspecified, saml-none, and saml-persistent-opaque.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `name_id_userstore_attribute` - (Optional) (Updatable) **Deprecated Since: 18.2.2**

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* idcsValuePersistedInOtherAttribute: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none This property specifies which user attribute is used as the NameID value in the SAML assertion. This attribute can be constructed by using attributes from the Oracle Identity Cloud Service Core Users schema. 
	* `outbound_assertion_attributes` - (Optional) (Updatable) Use to construct the outgoing SAML attributes

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `direction` - (Optional) (Updatable) Mapped Attribute Direction

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - (Optional) (Updatable) Mapped Attribute URI

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - (Required) (Updatable) Mapped Attribute identifier

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `partner_provider_id` - (Optional) (Updatable) The ID of the Provider. This value corresponds to the entityID from the Service Provider metadata.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `partner_provider_pattern` - (Optional) (Updatable) The pattern of the Provider. This value corresponds to the entityID from the Service Provider metadata.

		**Added In:** 2202230830

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `sign_response_or_assertion` - (Optional) (Updatable) Indicates which part of the response should be signed.  A value of \"Assertion\" indicates that the Assertion should be signed.  A value of \"Response\" indicates that the SSO Response should be signed. A value of \"AssertionAndResponse\" indicates that both the Assertion and the SSO Response should be signed.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `signature_hash_algorithm` - (Optional) (Updatable) This attribute represents the algorithm used to hash the signature.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `signing_certificate` - (Optional) (Updatable) This attribute represents the signing certificate that an App uses to verify the signed authentication request.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `succinct_id` - (Optional) (Updatable) This attribute represents the Succinct ID.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: server
	* `tenant_provider_id` - (Optional) (Updatable) The alternate Provider ID to be used as the Oracle Identity Cloud Service providerID (instead of the one in SamlSettings) when interacting with this SP.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `user_assertion_attributes` - (Optional) (Updatable) Each value of this attribute describes an attribute of User that will be sent in a Security Assertion Markup Language (SAML) assertion.

		**Deprecated Since: 18.2.2**

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCompositeKey: [name]
		* idcsSearchable: false
		* idcsValuePersistedInOtherAttribute: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `format` - (Optional) (Updatable) Indicates the format of the assertion attribute.

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `name` - (Required) (Updatable) The attribute represents the name of the attribute that will be used in the Security Assertion Markup Language (SAML) assertion

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `user_store_attribute_name` - (Required) (Updatable) This attribute specifies which user attribute should be used to create the value of the SAML assertion attribute. The userstore attribute can be constructed by using attributes from the Oracle Identity Cloud Service Core Users schema. <br><b>Note</b>: Attributes from extensions to the Core User schema are not supported in v1.0.

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app` - (Optional) (Updatable) WebTier Policy
	* `resource_ref` - (Optional) (Updatable) If this Attribute is true, resource ref id and resource ref name attributes will we included in wtp json response.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `web_tier_policy_az_control` - (Optional) (Updatable) Webtier policy AZ Control

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `web_tier_policy_json` - (Optional) (Updatable) Store the web tier policy for an application as a string in Javascript Object Notification (JSON) format.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
* `user_roles` - (Optional) (Updatable) A list of AppRoles defined by this UnmanagedApp. Membership in each of these AppRoles confers end-user privilege within this App.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* `description` - (Optional) (Updatable) The description of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `display` - (Optional) (Updatable) Display-name of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) URI of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) ID of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_token_expiry` - Expiry-time in seconds for an Access Token. Any token that allows access to this App will expire after the specified duration.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `accounts` - Accounts of App

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
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
	* `name` - Name of the account

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
	* `owner_id` - Owner identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - AccountMgmtInfo URI

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Account identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: always
		* type: string
		* uniqueness: none
* `active` - If true, this App is able to participate in runtime services, such as automatic-login, OAuth, and SAML. If false, all runtime services are disabled for this App, and only administrative operations can be performed.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `admin_roles` - A list of AppRoles defined by this UnmanagedApp. Membership in each of these AppRoles confers administrative privilege within this App.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* `description` - The description of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `display` - Display-name of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - URI of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - ID of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `alias_apps` - Each value of this internal attribute refers to an Oracle Public Cloud infrastructure App on which this App depends.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `description` - Description of the alias App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `display` - Display name of the alias App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - URI of the alias App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - ID of the alias App.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `all_url_schemes_allowed` - If true, indicates that the system should allow all URL-schemes within each value of the 'redirectUris' attribute.  Also indicates that the system should not attempt to confirm that each value of the 'redirectUris' attribute is a valid URI.  In particular, the system should not confirm that the domain component of the URI is a top-level domain and the system should not confirm that the hostname portion is a valid system that is reachable over the network.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `allow_access_control` - If true, any managed App that is based on this template is checked for access control that is, access to this app is subject to successful authorization at SSO service, viz. app grants to start with.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `allow_offline` - If true, indicates that the Refresh Token is allowed when this App acts as an OAuth Resource.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `allowed_grants` - List of grant-types that this App is allowed to use when it acts as an OAuthClient.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `allowed_operations` - OPTIONAL. Required only when this App acts as an OAuthClient. Supported values are 'introspect' and 'onBehalfOfUser'. The value 'introspect' allows the client to look inside the access-token. The value 'onBehalfOfUser' overrides how the client's privileges are combined with the privileges of the Subject User. Ordinarily, authorization calculates the set of effective privileges as the intersection of the client's privileges and the user's privileges. The value 'onBehalfOf' indicates that authorization should ignore the privileges of the client and use only the user's privileges to calculate the effective privileges.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `allowed_scopes` - A list of scopes (exposed by this App or by other Apps) that this App is allowed to access when it acts as an OAuthClient.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsCompositeKey: [fqs]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `fqs` - A fully qualified scope that this App is allowed to access when it acts as an OAuthClient.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `id_of_defining_app` - The ID of the App that defines this scope.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `read_only` - If true, indicates that this value must be protected.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
* `allowed_tags` - A list of tags, acting as an OAuthClient, this App is allowed to access.

	**Added In:** 17.4.6

	**SCIM++ Properties:**
	* idcsCompositeKey: [key, value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `key` - Key or name of the allowed tag.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `read_only` - If true, indicates that this value must be protected.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `value` - Value of the allowed tag.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `app_icon` - Application icon.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `app_signon_policy` - App Sign-on Policy.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - URI of the policy.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Identifier of the Policy.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `app_thumbnail` - Application thumbnail.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `apps_network_perimeters` - Network Perimeter

	**Added In:** 2010242156

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - URI of the Network Perimeter.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - List of identifier of Network Perimeters for App

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: always
		* type: string
		* uniqueness: none
* `as_opc_service` - OPCService facet of the application.

	**Deprecated Since: 17.3.4**

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - URI of the OPCService facet.

		**Deprecated Since: 17.3.4**

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Identifier of the OPCService facet.

		**Deprecated Since: 17.3.4**

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `attr_rendering_metadata` - Label for the attribute to be shown in the UI.

	**SCIM++ Properties:**
	* idcsCompositeKey: [name]
	* idcsSearchable: false
	* multiValued: true
	* mutability: immutable
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `datatype` - Data type of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `helptext` - Help text for the attribute. It can contain HTML tags.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `label` - Label for the attribute to be shown in the UI.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `max_length` - Maximum length of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_size` - Maximum size of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `min_length` - Minimum length of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `min_size` - Minimum size of the attribute..

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `name` - Name of the attribute.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `order` - Data type of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `read_only` - Is the attribute readOnly.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `regexp` - Regular expression of the attribute for validation.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `required` - Attribute is required or optional.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `section` - UI widget to use for the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `visible` - Indicates whether the attribute is to be shown on the application creation UI.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `widget` - UI widget to use for the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `audience` - The base URI for all of the scopes defined in this App. The value of 'audience' is combined with the 'value' of each scope to form an 'fqs' or fully qualified scope.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `based_on_template` - Application template on which the application is based.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: true
	* returned: default
	* type: complex
	* `last_modified` - The most recent DateTime that the appTemplate on which the application based upon is updated. The attribute MUST be a DateTime.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
		* uniqueness: none
	* `ref` - URI of the application template.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Identifier of the application template.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: always
		* type: string
		* uniqueness: none
	* `well_known_id` - Unique Well-known identifier used to reference app template.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `bypass_consent` - If true, indicates that consent should be skipped for all scopes

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `callback_service_url` - Callback Service URL

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `certificates` - Each value of this attribute represent a certificate that this App uses when it acts as an OAuthClient.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [certAlias]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `cert_alias` - Certificate alias

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: always
		* type: string
		* uniqueness: none
	* `kid` - Certificate kid

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `sha1thumbprint` - sha1Thumbprint

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `x509base64certificate` - Base-64-encoded certificate.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: binary
		* uniqueness: none
	* `x5t` - Certificate x5t

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `client_ip_checking` - Network Perimeters checking mode

	**Added In:** 2010242156

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `client_secret` - This value is the credential of this App, which this App supplies as a password when this App authenticates to the Oracle Public Cloud infrastructure. This value is also the client secret of this App when it acts as an OAuthClient.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* idcsSensitive: none
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `client_type` - Specifies the type of access that this App has when it acts as an OAuthClient.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `cloud_control_properties` - A collection of arbitrary properties that scope the privileges of a cloud-control App.

	**Added In:** 18.4.2

	**SCIM++ Properties:**
	* idcsCompositeKey: [name]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `name` - The name of the property.

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `values` - The value(s) of the property.

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: true
		* mutability: readOnly
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
* `contact_email_address` - Contact Email Address

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `delegated_service_names` - Service Names allow to use Oracle Cloud Infrastructure signature for client authentication instead of client credentials

	**Added In:** 2207040824

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
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
* `description` - Description of the application.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `disable_kmsi_token_authentication` - Indicates whether the application is allowed to be access using kmsi token.

	**Added In:** 2111190457

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: always
	* type: boolean
	* uniqueness: none
* `display_name` - Display name of the application. Display name is intended to be user-friendly, and an administrator can change the value at any time.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: always
	* type: string
	* uniqueness: server
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
* `editable_attributes` - App attributes editable by subject

	**Added In:** 18.2.6

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [name]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `name` - Name of the attribute. The attribute name will be qualified by schema name if any extension schema defines the attribute. The attribute name will not be qualified by schema name if the base schema defines the attribute.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `error_page_url` - This attribute specifies the URL of the page to which an application will redirect an end-user in case of error.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `granted_app_roles` - A list of AppRoles that are granted to this App (and that are defined by other Apps). Within the Oracle Public Cloud infrastructure, this allows AppID-based association. Such an association allows this App to act as a consumer and thus to access resources of another App that acts as a producer.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `admin_role` - If true, then this granted AppRole confers administrative privileges within the App that defines it. Otherwise, the granted AppRole confers only functional privileges.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `app_id` - The id of the App that defines this AppRole, which is granted to this App. The App that defines the AppRole acts as the producer; the App to which the AppRole is granted acts as a consumer.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `app_name` - The name of the App that defines this AppRole, which is granted to this App. The App that defines the AppRole acts as the producer; the App to which the AppRole is granted acts as a consumer.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `display` - The display-name of an AppRole that is granted to this App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `legacy_group_name` - The name of the legacy group associated with this AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `read_only` - If true, indicates that this value must be protected.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `ref` - The URI of an AppRole that is granted to this App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - A label that indicates whether this AppRole was granted directly to the App (or indirectly through a Group). For an App, the value of this attribute will always be 'direct' (because an App cannot be a member of a Group).

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - The id of an AppRole that is granted to this App.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `grants` - Grants assigned to the app

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
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
	* `grantee_id` - Grantee identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `grantee_type` - Grantee resource type. Allowed values are User and Group.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - Grant URI

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Grant identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `hashed_client_secret` - Hashed Client Secret. This hash-value is used to verify the 'clientSecret' credential of this App

	**Added In:** 2106240046

	**SCIM++ Properties:**
	* idcsSearchable: false
	* idcsSensitive: hash_sc
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `home_page_url` - Home Page URL

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `icon` - URL of application icon.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: reference
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
* `id_token_enc_algo` - Encryption Alogrithm to use for encrypting ID token.

	**Added In:** 2010242156

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
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
* `identity_providers` - A list of IdentityProvider assigned to app. A user trying to access this app will be automatically redirected to configured IdP during the authentication phase, before being able to access App.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: request
	* type: complex
	* `display` - Display-name of the IdentityProvider.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - URI of the IdentityProvider.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - ID of the IdentityProvider.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `idp_policy` - IDP Policy.

	**Added In:** 18.1.2

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - URI of the policy.

		**Added In:** 18.1.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Identifier of the Policy.

		**Added In:** 18.1.2

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `infrastructure` - If true, this App is an internal infrastructure App.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_alias_app` - If true, this App is an AliasApp and it cannot be granted to an end-user directly.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: always
	* type: boolean
	* uniqueness: none
* `is_database_service` - If true, this application acts as database service Application

	**Added In:** 18.2.2

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* type: boolean
* `is_enterprise_app` - If true, this app acts as Enterprise app with Authentication and URL Authz policy.

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_form_fill` - If true, this application acts as FormFill Application

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_kerberos_realm` - If true, indicates that this App supports Kerberos Authentication

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_login_target` - If true, this App allows runtime services to log end users into this App automatically.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_managed_app` - If true, indicates that access to this App requires an account. That is, in order to log in to the App, a User must use an application-specific identity that is maintained in the remote identity-repository of that App.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_mobile_target` - If true, indicates that the App should be visible in each end-user's mobile application.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_multicloud_service_app` - If true, indicates the app is used for multicloud service integration.

	**Added In:** 2301202328

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_oauth_client` - If true, this application acts as an OAuth Client

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_oauth_resource` - If true, indicates that this application acts as an OAuth Resource.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_opc_service` - If true, this application is an Oracle Public Cloud service-instance.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_obligation_capable` - This flag indicates if the App is capable of validating obligations with the token for allowing access to the App.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_radius_app` - If true, this application acts as an Radius App

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_saml_service_provider` - If true, then this App acts as a SAML Service Provider.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_unmanaged_app` - If true, indicates that this application accepts an Oracle Cloud Identity Service User as a login-identity (does not require an account) and relies for authorization on the User's memberships in AppRoles.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `is_web_tier_policy` - If true, the webtier policy is active

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `landing_page_url` - The URL of the landing page for this App, which is the first page that an end user should see if runtime services log that end user in to this App automatically.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `linking_callback_url` - This attribute specifies the callback URL for the social linking operation.

	**Added In:** 18.2.4

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `login_mechanism` - The protocol that runtime services will use to log end users in to this App automatically. If 'OIDC', then runtime services use the OpenID Connect protocol. If 'SAML', then runtime services use Security Assertion Markup Language protocol.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `login_page_url` - This attribute specifies the URL of the page that the App uses when an end-user signs in to that App.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `logout_page_url` - This attribute specifies the URL of the page that the App uses when an end-user signs out.

	**Added In:** 17.4.2

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `logout_uri` - OAuth will use this URI to logout if this App wants to participate in SSO, and if this App's session gets cleared as part of global logout. Note: This attribute is used only if this App acts as an OAuthClient.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
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
* `meter_as_opc_service` - Indicates whether the application is billed as an OPCService. If true, customer is not billed for runtime operations of the app.

	**Added In:** 18.4.2

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: always
	* type: boolean
	* uniqueness: none
* `migrated` - If true, this App was migrated from an earlier version of Oracle Public Cloud infrastructure (and may therefore require special handling from runtime services such as OAuth or SAML). If false, this App requires no special handling from runtime services.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `name` - Name of the application. Also serves as username if the application authenticates to Oracle Public Cloud infrastructure. This name may not be user-friendly and cannot be changed once an App is created.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: string
	* uniqueness: server
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
* `post_logout_redirect_uris` - Each value of this attribute is the URI of a landing page within this App. It is used only when this App, acting as an OAuthClient, initiates the logout flow and wants to be redirected back to one of its landing pages.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `privacy_policy_url` - Privacy Policy URL

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `product_logo_url` - Application Logo URL

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `product_name` - Product Name

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `protectable_secondary_audiences` - A list of secondary audiences--additional URIs to be added automatically to any OAuth token that allows access to this App. Note: This attribute is used mainly for backward compatibility in certain Oracle Public Cloud Apps.

	**Added In:** 18.2.2

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `read_only` - If true, indicates that this value must be protected.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `value` - The value of an secondary audience--additional URI to be added automatically to any OAuth token that allows access to this App.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `radius_policy` - RADIUS Policy assigned to this application.

	**Added In:** 2209070044

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - URI of the policy.

		**Added In:** 2209070044

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Identifier of the Policy.

		**Added In:** 2209070044

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `ready_to_upgrade` - If true, this App requires an upgrade and mandates attention from application administrator. The flag is used by UI to indicate this app is ready to upgrade.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `redirect_uris` - OPTIONAL. Each value is a URI within this App. This attribute is required when this App acts as an OAuthClient and is involved in three-legged flows (authorization-code flows).

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `refresh_token_expiry` - Expiry-time in seconds for a Refresh Token.  Any token that allows access to this App, once refreshed, will expire after the specified duration.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `saml_service_provider` - An attribute that refers to the SAML Service Provider that runtime services will use to log an end user in to this App automatically. Note that this will be used only if the loginMechanism is 'SAML'.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `ref` - The URI of the App that acts as a Service Provider.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - The id of the App that acts as a Service Provider.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: always
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
* `scopes` - Scopes defined by this App. Used when this App acts as an OAuth Resource.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `description` - OAuth scope description

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `display_name` - OAuth scope display name

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `fqs` - The fully qualified value of this scope within this App. A fully qualified scope combines the 'value' of each scope with the value of 'audience'. Each value of 'fqs' must be unique across the system. Used only when this App acts as an OAuth Resource.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: server
	* `read_only` - If true, indicates that this value must be protected.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `requires_consent` - If true, indicates that a user must provide consent to access this scope. Note: Used only when this App acts as an OAuth Resource.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `value` - OAuth scope.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `secondary_audiences` - A list of secondary audiences--additional URIs to be added automatically to any OAuth token that allows access to this App. Note: This attribute is used mainly for backward compatibility in certain Oracle Public Cloud Apps.

	**Deprecated Since: 18.2.6**

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `service_params` - Custom attribute that is required to compute other attribute values during app creation.

	**SCIM++ Properties:**
	* idcsCompositeKey: [name]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: always
	* type: complex
	* uniqueness: none
	* `name` - The name of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - The value of the attribute.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `service_type_urn` - This Uniform Resource Name (URN) value identifies the type of Oracle Public Cloud service of which this app is an instance.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `service_type_version` - This value specifies the version of the Oracle Public Cloud service of which this App is an instance

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `show_in_my_apps` - If true, this app will be displayed in the MyApps page of each end-user who has access to the App.

	**Added In:** 18.1.2

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `signon_policy` - Sign-on Policy.

	**Deprecated Since: 17.3.4**

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - URI of the policy.

		**Deprecated Since: 17.3.4**

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Identifier of the Policy.

		**Deprecated Since: 17.3.4**

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
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
* `terms_of_service_url` - Terms of Service URL

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `terms_of_use` - Terms Of Use.

	**Added In:** 18.2.6

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `name` - Terms Of Use name

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `ref` - URI of the TermsOfUse.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Identifier of the TermsOfUse

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
* `trust_policies` - Trust Policies.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `ref` - URI of the policy.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Identifier of the Policy.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `trust_scope` - Indicates the scope of trust for this App when acting as an OAuthClient. A value of 'Explicit' indicates that the App is allowed to access only the scopes of OAuthResources that are explicitly specified as 'allowedScopes'. A value of 'Account' indicates that the App is allowed implicitly to access any scope of any OAuthResource within the same Oracle Cloud Account. A value of 'Tags' indicates that the App is allowed to access any scope of any OAuthResource with a matching tag within the same Oracle Cloud Account. A value of 'Default' indicates that the Tenant default trust scope configured in the Tenant Settings is used.

	**Added In:** 17.4.2

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
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
* `urnietfparamsscimschemasoracleidcsextensiondbcs_app` - This extension provides attributes for database service facet of an App
	* `domain_app` - Description:

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none 
		* `display` - DB Domain App display name

			**Added In:** 18.2.2

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: request
			* type: string
			* uniqueness: none
		* `ref` - DB Domain App URI

			**Added In:** 18.2.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - DB Domain App identifier

			**Added In:** 18.2.2

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `domain_name` - The name of the Enterprise Domain that contains any number of DBInstances. If specified, the value must be unique.  A non-null value indicates that App represents a DBDomain. A value of null indicates that the App represents an DB-instance.

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: string
		* uniqueness: server
* `urnietfparamsscimschemasoracleidcsextensionenterprise_app_app` - This extension defines the Enterprise App related attributes.
	* `allow_authz_decision_ttl` - Allow Authz policy decision expiry time in seconds.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsMaxValue: 3600
		* idcsMinValue: 0
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `allow_authz_policy` - Allow Authz Policy.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* `ref` - URI of the policy.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - Identifier of the Policy.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `app_resources` - A list of AppResources of this App.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: true
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* `ref` - The URI of an AppResource of this App.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The id of an AppResource of this App.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `deny_authz_decision_ttl` - Deny Authz policy decision expiry time in seconds.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsMaxValue: 3600
		* idcsMinValue: 0
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `deny_authz_policy` - Deny Authz Policy.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* `ref` - URI of the policy.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - Identifier of the Policy.

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionform_fill_app_app` - This extension provides attributes for Form-Fill facet of App
	* `configuration` - FormFill Application Configuration CLOB which has to be maintained in Form-Fill APP for legacy code to do Form-Fill injection

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `form_cred_method` - Indicates how FormFill obtains the username and password of the account that FormFill will use to sign into the target App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `form_credential_sharing_group_id` - Credential Sharing Group to which this form-fill application belongs.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `form_fill_url_match` - A list of application-formURLs that FormFill should match against any formUrl that the user-specifies when signing in to the target service.  Each item in the list also indicates how FormFill should interpret that formUrl.

		**SCIM++ Properties:**
		* idcsCompositeKey: [formUrl]
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `form_url` - An application formUrl that FormFill will match against any formUrl that a User enters in trying to access the target-service which this App represents.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `form_url_match_type` - Indicates how to interpret the value of 'formUrl' when matching against a user-specified formUrl.  The system currently supports only 'Exact', which indicates that the value of 'formUrl' should be treated as a literal value.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `form_type` - Type of the FormFill application like WebApplication, MainFrameApplication, WindowsApplication. Initially, we will support only WebApplication.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `reveal_password_on_form` - If true, indicates that system is allowed to show the password in plain-text for this account after re-authentication.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `sync_from_template` - If true, indicates that each of the Form-Fill-related attributes that can be inherited from the template actually will be inherited from the template. If false, indicates that the AppTemplate on which this App is based has disabled inheritance for these Form-Fill-related attributes.

		**Added In:** 17.4.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `user_name_form_expression` - Indicates the custom expression, which can combine concat and substring operations with literals and with any attribute of the Oracle Identity Cloud Service User

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `user_name_form_template` - Format for generating a username.  This value can be Username or Email Address; any other value will be treated as a custom expression.  A custom expression may combine 'concat' and 'substring' operations with literals and with any attribute of the Oracle Identity Cloud Service user.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template` - This extension provides attributes for Form-Fill facet of AppTemplate
	* `configuration` - FormFill Application Configuration CLOB which has to be maintained in Form-Fill APP for legacy code to do Form-Fill injection

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `form_cred_method` - Indicates how FormFill obtains the username and password of the account that FormFill will use to sign into the target App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `form_credential_sharing_group_id` - Credential Sharing Group to which this form-fill application belongs.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `form_fill_url_match` - A list of application-formURLs that FormFill should match against any formUrl that the user-specifies when signing in to the target service.  Each item in the list also indicates how FormFill should interpret that formUrl.

		**SCIM++ Properties:**
		* idcsCompositeKey: [formUrl]
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `form_url` - An application formUrl that FormFill will match against any formUrl that a User enters in trying to access the target-service which this App represents.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `form_url_match_type` - Indicates how to interpret the value of 'formUrl' when matching against a user-specified formUrl.  The system currently supports only 'Exact', which indicates that the value of 'formUrl' should be treated as a literal value.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `form_type` - Type of the FormFill application like WebApplication, MainFrameApplication, WindowsApplication. Initially, we will support only WebApplication.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `reveal_password_on_form` - If true, indicates that system is allowed to show the password in plain-text for this account after re-authentication.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `sync_from_template` - If true, indicates that each of the Form-Fill-related attributes that can be inherited from the template actually will be inherited from the template. If false, indicates that the AppTemplate disabled inheritance for these Form-Fill-related attributes.

		**Added In:** 17.4.2

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `user_name_form_expression` - Indicates the custom expression, which can combine concat and substring operations with literals and with any attribute of the Oracle Identity Cloud Service User

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `user_name_form_template` - Format for generating a username.  This value can be Username or Email Address; any other value will be treated as a custom expression.  A custom expression may combine 'concat' and 'substring' operations with literals and with any attribute of the Oracle Identity Cloud Service user.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsPii: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app` - Kerberos Realm
	* `default_encryption_salt_type` - The type of salt that the system will use to encrypt Kerberos-specific artifacts of this App unless another type of salt is specified.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `master_key` - The primary key that the system should use to encrypt artifacts that are specific to this Kerberos realm -- for example, to encrypt the Principal Key in each KerberosRealmUser.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsSensitive: none
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `max_renewable_age` - Max Renewable Age in seconds

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `max_ticket_life` - Max Ticket Life in seconds

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
	* `realm_name` - The name of the Kerberos Realm that this App uses for authentication.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `supported_encryption_salt_types` - The types of salt that are available for the system to use when encrypting Kerberos-specific artifacts for this App.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `ticket_flags` - Ticket Flags

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: integer
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionmanagedapp_app` - Managed App
	* `account_form_visible` - If true, then the account form will be displayed in the Oracle Identity Cloud Service UI to interactively create or update an account for this App. If a value is not specified for this attribute, a default value of \"false\" will be assumed as the value for this attribute.

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `admin_consent_granted` - If true, admin has granted consent to perform managed app run-time operations.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `bundle_configuration_properties` - ConnectorBundle configuration properties

		**SCIM++ Properties:**
		* idcsCompositeKey: [name]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `confidential` - If true, this bundle configuration property value is confidential and will be encrypted in Oracle Identity Cloud Service. This attribute maps to \"isConfidential\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `display_name` - Display name of the bundle configuration property. This attribute maps to \"displayName\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `help_message` - Help message of the bundle configuration property. This attribute maps to \"helpMessage\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `icf_type` - ICF data type of the bundle configuration property. This attribute maps to \"type\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `name` - Name of the bundle configuration property. This attribute maps to \"name\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `order` - Display sequence of the bundle configuration property.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `required` - If true, this bundle configuration property is required to connect to the target connected managed app. This attribute maps to \"isRequired\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: boolean
			* uniqueness: none
		* `value` - Value of the bundle configuration property. This attribute maps to \"value\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* idcsSensitive: encrypt
			* multiValued: true
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `bundle_pool_configuration` - Configurable options maintaining a pool of ICF connector instances. Values for sub attributes can be set only if the ConnectorBundle referenced in the App has connectorPoolingSupported set to true

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `max_idle` - Maximum number of connector instances in the pool that are idle and active.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `max_objects` - Maximum number of connector instances in the pool that are idle and active.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `max_wait` - Maximum time (in milliseconds) to wait for a free connector instance to become available before failing.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `min_evictable_idle_time_millis` - Minimum time (in milliseconds) to wait before evicting an idle conenctor instance from the pool.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `min_idle` - Minimum number of idle connector instances in the pool.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
	* `can_be_authoritative` - If true, the managed app can be authoritative.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `connected` - If true, the accounts of the application are managed through an ICF connector bundle

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `connector_bundle` - ConnectorBundle

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - ConnectorBundle display name

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - ConnectorBundle URI

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - Connector Bundle type. Allowed values are ConnectorBundle, LocalConnectorBundle.

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: true
			* idcsDefaultValue: ConnectorBundle
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - ConnectorBundle identifier

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `well_known_id` - Unique Well-known identifier used to reference connector bundle.

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `enable_auth_sync_new_user_notification` - If true, send activation email to new users created from authoritative sync.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `enable_sync` - If true, sync run-time operations are enabled for this App.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `enable_sync_summary_report_notification` - If true, send sync summary as notification upon job completion.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `flat_file_bundle_configuration_properties` - Flat file connector bundle configuration properties

		**SCIM++ Properties:**
		* idcsCompositeKey: [name]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `confidential` - If true, this flatfile bundle configuration property value is confidential and will be encrypted in Oracle Identity Cloud Service. This attribute maps to \"isConfidential\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `display_name` - Display name of the flatfile bundle configuration property. This attribute maps to \"displayName\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `help_message` - Help message of the flatfile bundle configuration property. This attribute maps to \"helpMessage\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `icf_type` - ICF data type of flatfile the bundle configuration property. This attribute maps to \"type\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `name` - Name of the flatfile bundle configuration property. This attribute maps to \"name\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `order` - Display sequence of the bundle configuration property.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `required` - If true, this flatfile bundle configuration property is required to connect to the target connected managed app. This attribute maps to \"isRequired\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: boolean
			* uniqueness: none
		* `value` - Value of the flatfile bundle configuration property. This attribute maps to \"value\" attribute in \"ConfigurationProperty\" in ICF.

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* idcsSensitive: encrypt
			* multiValued: true
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `flat_file_connector_bundle` - Flat file connector bundle to sync from a flat file.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - ConnectorBundle display name

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - ConnectorBundle URI

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - ConnectorBundle identifier

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `well_known_id` - Unique well-known identifier used to reference connector bundle.

			**Added In:** 19.1.4

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `identity_bridges` - IdentityBridges associated with this App

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `name` - Name of the IdentityBridge associated with the App.

			**Added In:** 19.1.4

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI of the IdentityBridge associated with the App.

			**Added In:** 19.1.4

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The Id of the IdentityBridge associated with the App.

			**Added In:** 19.1.4

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: always
			* type: string
			* uniqueness: none
	* `is_authoritative` - If true, sync from the managed app will be performed as authoritative sync.

		**Added In:** 17.4.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_directory` - If true, the managed app is a directory.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_on_premise_app` - If true, the managed app is an On-Premise app.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_schema_customization_supported` - If true, the managed app supports schema customization.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_schema_discovery_supported` - If true, the managed app supports schema discovery.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_three_legged_oauth_enabled` - If true, the managed app requires 3-legged OAuth for authorization.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `is_two_legged_oauth_enabled` - If true, indicates that Oracle Identity Cloud Service can use two-legged OAuth to connect to this ManagedApp.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `object_classes` - Object classes

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - Object class display name

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `is_account_object_class` - If true, the object class represents an account. The isAccountObjectClass attribute value 'true' MUST appear no more than once.

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `ref` - Object class URI

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `resource_type` - Object class resource type

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `type` - Object Class type. Allowed values are AccountObjectClass, ManagedObjectClass.

			**Added In:** 18.1.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsDefaultValue: AccountObjectClass
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - Object class template identifier

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `sync_config_last_modified` - The most recent DateTime that the configuration of this App was updated. AppServices updates this timestamp whenever AppServices updates an App's configuration with respect to synchronization.

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: dateTime
		* uniqueness: none
	* `three_legged_oauth_credential` - The value of this attribute persists any OAuth access token that the system uses to connect to this ManagedApp. The system obtains this access token using an OAuth protocol flow that could be two-legged or three-legged. A two-legged flow involves only the requester and the server. A three-legged flow also requires the consent of a user -- in this case the consent of an administrator.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* `access_token` - Access Token

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* idcsSensitive: encrypt
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `access_token_expiry` - Access token expiry

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: dateTime
			* uniqueness: none
		* `refresh_token` - Refresh Token

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* idcsSensitive: encrypt
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `three_legged_oauth_provider_name` - Three legged OAuth provider name in Oracle Identity Cloud Service.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
* `urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app` - This extension defines attributes specific to Apps that represent instances of Multicloud Service App
	* `multicloud_platform_url` - The multicloud platform service URL which the application will invoke for runtime operations such as AWSCredentials api invocation

		**Added In:** 2301202328

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `multicloud_service_type` - Specifies the service type for which the application is configured for multicloud integration. For applicable external service types, app will invoke multicloud service for runtime operations

		**Added In:** 2301202328

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: request
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionopc_service_app` - This extension defines attributes specific to Apps that represent instances of an Oracle Public Cloud (OPC) service.
	* `current_federation_mode` - Current Federation Mode

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
	* `current_synchronization_mode` - Current Synchronization Mode

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
	* `enabling_next_fed_sync_modes` - If true, indicates that enablement is in progress started but not completed

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
	* `next_federation_mode` - Next Federation Mode

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
	* `next_synchronization_mode` - Next Synchronization Mode

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
	* `region` - This value identifies the OPC region in which the service is running.

		**Added In:** 19.1.4

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `service_instance_identifier` - This value specifies the unique identifier assigned to an instance of an Oracle Public Cloud service app.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: string
		* uniqueness: server
* `urnietfparamsscimschemasoracleidcsextensionradius_app_app` - This extension defines attributes specific to Apps that represent instances of Radius App.
	* `capture_client_ip` - If true, capture the client IP address from the RADIUS request packet. IP Address is used for auditing, policy-evaluation and country-code calculation.

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
	* `client_ip` - This is the IP address of the RADIUS Client like Oracle Database server. It can be only IP address and not hostname.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `country_code_response_attribute_id` - Vendor-specific identifier of the attribute in the RADIUS response that will contain the end-user's country code. This is an integer-value in the range 1 to 255

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `end_user_ip_attribute` - The name of the attribute that contains the Internet Protocol address of the end-user.

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `group_membership_radius_attribute` - RADIUS attribute that RADIUS-enabled system uses to pass the group membership

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `group_membership_to_return` - In a successful authentication response, Oracle Identity Cloud Service will pass user's group information restricted to groups persisted in this attribute, in the specified RADIUS attribute.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `display` - A human readable name, primarily used for display purposes. READ-ONLY.

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI of the corresponding Group resource to which the user belongs

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The identifier of the User's group.

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
	* `group_name_format` - Configure the groupNameFormat based on vendor in order to pass it to RADIUS infra

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `include_group_in_response` - Indicates to include groups in RADIUS response

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
	* `password_and_otp_together` - Indicates if password and OTP are passed in the same sign-in request or not.

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
	* `port` - This is the port of RADIUS Proxy which RADIUS client will connect to.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `radius_vendor_specific_id` - ID used to identify a particular vendor.

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `response_format` - Configure the responseFormat based on vendor in order to pass it to RADIUS infra

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `response_format_delimiter` - The delimiter used if group membership responseFormat is a delimited list instead of repeating attributes

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `secret_key` - Secret key used to secure communication between RADIUS Proxy and RADIUS client

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `type_of_radius_app` - Value consists of type of RADIUS App. Type can be Oracle Database, VPN etc

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
* `urnietfparamsscimschemasoracleidcsextensionrequestable_app` - Requestable App
	* `requestable` - Flag controlling whether resource can be request by user through self service console.

		**Added In:** 17.3.4

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app` - This extension defines attributes related to the Service Providers configuration.
	* `assertion_consumer_url` - The attribute represents the URL to which the SAML Assertions will be sent by the SAML IdP.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `encrypt_assertion` - If true, indicates that the system must encrypt the Security Assertion Markup Language (SAML) assertion.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `encryption_algorithm` - This attribute indicates the encryption algorithm used to encrypt the SAML assertion.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `encryption_certificate` - This attribute represents the encryption certificate that an App uses to encrypt the Security Assertion Markup Language (SAML) assertion.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `federation_protocol` - Specifies the preferred federation protocol (SAML2.0 or WS-Fed1.1).

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: true
		* idcsDefaultValue: SAML2.0
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `group_assertion_attributes` - Each value of this attribute describes an attribute of Group that will be sent in a Security Assertion Markup Language (SAML) assertion.

		**Deprecated Since: 18.2.2**

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCompositeKey: [name]
		* idcsSearchable: false
		* idcsValuePersistedInOtherAttribute: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `condition` - Indicates the filter types that are supported for the Group assertion attributes.

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `format` - Indicates the format of the assertion attribute.

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `group_name` - Indicates the group name that are supported for the group assertion attributes.

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `name` - The attribute represents the name of the attribute that will be used in the Security Assertion Markup Language (SAML) assertion

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `hok_acs_url` - Hok Assertion Consumer Service Url

		**Added In:** 2101262133

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `hok_required` - If enabled, then the SAML Service supports Hok for this App.

		**Added In:** 2101262133

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `include_signing_cert_in_signature` - If true, then the signing certificate is included in the signature.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `key_encryption_algorithm` - This attribute indicates the key encryption algorithm.

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `last_notification_sent_time` - Records the notification timestamp for the SP whose signing certificate is about to expire.

		**Added In:** 2302092332

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
		* uniqueness: none
	* `logout_binding` - This attribute represents the HTTP binding that would be used while logout.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `logout_enabled` - If true, then the SAML Service supports logout for this App.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `logout_request_url` - The URL to which the partner sends the logout request.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `logout_response_url` - The URL to which the partner sends the logout response.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `metadata` - This attribute represents the metadata of a Security Provider in the Security Assertion Markup Language protocol.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `name_id_format` - This can be any string, but there are a set of standard nameIdFormats. If a nameIdFormat other than the standard list is chosen, it will be considered a custom nameidformat. The standard nameidformats include: saml-x509, saml-emailaddress, saml-windowsnamequalifier, saml-kerberos, saml-persistent, saml-transient, saml-unspecified, saml-none, and saml-persistent-opaque.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `name_id_userstore_attribute` - **Deprecated Since: 18.2.2**

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* idcsValuePersistedInOtherAttribute: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none This property specifies which user attribute is used as the NameID value in the SAML assertion. This attribute can be constructed by using attributes from the Oracle Identity Cloud Service Core Users schema. 
	* `outbound_assertion_attributes` - Use to construct the outgoing SAML attributes

		**Added In:** 18.2.6

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `direction` - Mapped Attribute Direction

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - Mapped Attribute URI

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* idcsSearchable: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - Mapped Attribute identifier

			**Added In:** 18.2.6

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `partner_provider_id` - The ID of the Provider. This value corresponds to the entityID from the Service Provider metadata.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `partner_provider_pattern` - The pattern of the Provider. This value corresponds to the entityID from the Service Provider metadata.

		**Added In:** 2202230830

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `sign_response_or_assertion` - Indicates which part of the response should be signed.  A value of \"Assertion\" indicates that the Assertion should be signed.  A value of \"Response\" indicates that the SSO Response should be signed. A value of \"AssertionAndResponse\" indicates that both the Assertion and the SSO Response should be signed.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `signature_hash_algorithm` - This attribute represents the algorithm used to hash the signature.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `signing_certificate` - This attribute represents the signing certificate that an App uses to verify the signed authentication request.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `succinct_id` - This attribute represents the Succinct ID.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: server
	* `tenant_provider_id` - The alternate Provider ID to be used as the Oracle Identity Cloud Service providerID (instead of the one in SamlSettings) when interacting with this SP.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `user_assertion_attributes` - Each value of this attribute describes an attribute of User that will be sent in a Security Assertion Markup Language (SAML) assertion.

		**Deprecated Since: 18.2.2**

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCompositeKey: [name]
		* idcsSearchable: false
		* idcsValuePersistedInOtherAttribute: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `format` - Indicates the format of the assertion attribute.

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `name` - The attribute represents the name of the attribute that will be used in the Security Assertion Markup Language (SAML) assertion

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `user_store_attribute_name` - This attribute specifies which user attribute should be used to create the value of the SAML assertion attribute. The userstore attribute can be constructed by using attributes from the Oracle Identity Cloud Service Core Users schema. <br><b>Note</b>: Attributes from extensions to the Core User schema are not supported in v1.0.

			**Deprecated Since: 18.2.2**

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* idcsValuePersistedInOtherAttribute: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app` - WebTier Policy
	* `resource_ref` - If this Attribute is true, resource ref id and resource ref name attributes will we included in wtp json response.

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `web_tier_policy_az_control` - Webtier policy AZ Control

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `web_tier_policy_json` - Store the web tier policy for an application as a string in Javascript Object Notification (JSON) format.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
* `user_roles` - A list of AppRoles defined by this UnmanagedApp. Membership in each of these AppRoles confers end-user privilege within this App.

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* `description` - The description of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `display` - Display-name of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - URI of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - ID of the AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the App
	* `update` - (Defaults to 20 minutes), when updating the App
	* `delete` - (Defaults to 20 minutes), when destroying the App


## Import

Apps can be imported using the `id`, e.g.

```
$ terraform import oci_identity_domains_app.test_app "idcsEndpoint/{idcsEndpoint}/apps/{appId}" 
```

