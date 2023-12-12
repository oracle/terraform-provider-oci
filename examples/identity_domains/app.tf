// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "app_app_count" {
  default = 10
}

variable "app_app_filter" {
  default = "id ne \"IDCSAppId\""
}

variable "app_access_token_expiry" {
  default = 10
}

variable "app_accounts_active" {
  default = "false"
}

variable "app_accounts_name" {
  default = "name"
}

variable "app_accounts_ref" {
  default = "ref"
}

variable "app_accounts_value" {
  default = "value"
}

variable "app_active" {
  default = "false"
}

variable "app_admin_roles_description" {
  default = "description"
}

variable "app_admin_roles_display" {
  default = "display"
}

variable "app_admin_roles_ref" {
  default = "ref"
}

variable "app_admin_roles_value" {
  default = "value"
}

variable "app_alias_apps_description" {
  default = "description"
}

variable "app_alias_apps_display" {
  default = "display"
}

variable "app_alias_apps_ref" {
  default = "ref"
}

variable "app_alias_apps_value" {
  default = "value"
}

variable "app_all_url_schemes_allowed" {
  default = false
}

variable "app_allow_access_control" {
  default = false
}

variable "app_allow_offline" {
  default = false
}

variable "app_allowed_grants" {
  default = []
}

variable "app_allowed_operations" {
  default = []
}

variable "app_allowed_scopes_fqs" {
  default = "fqs"
}

variable "app_allowed_scopes_id_of_defining_app" {
  default = "idOfDefiningApp"
}

variable "app_allowed_scopes_read_only" {
  default = false
}

variable "app_allowed_tags_key" {
  default = "key"
}

variable "app_allowed_tags_read_only" {
  default = false
}

variable "app_allowed_tags_value" {
  default = "value"
}

variable "app_app_icon" {
  default = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII="
}

variable "app_app_signon_policy_ref" {
  default = "ref"
}

variable "app_app_signon_policy_value" {
  default = "value"
}

variable "app_app_thumbnail" {
  default = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII="
}

variable "app_apps_network_perimeters_ref" {
  default = "ref"
}

variable "app_apps_network_perimeters_value" {
  default = "value"
}

variable "app_as_opc_service_ref" {
  default = "ref"
}

variable "app_as_opc_service_value" {
  default = "value"
}

variable "app_attr_rendering_metadata_datatype" {
  default = "datatype"
}

variable "app_attr_rendering_metadata_helptext" {
  default = "helptext"
}

variable "app_attr_rendering_metadata_label" {
  default = "label"
}

variable "app_attr_rendering_metadata_max_length" {
  default = 10
}

variable "app_attr_rendering_metadata_max_size" {
  default = 10
}

variable "app_attr_rendering_metadata_min_length" {
  default = 10
}

variable "app_attr_rendering_metadata_min_size" {
  default = 10
}

variable "app_attr_rendering_metadata_name" {
  default = "name"
}

variable "app_attr_rendering_metadata_order" {
  default = 10
}

variable "app_attr_rendering_metadata_read_only" {
  default = false
}

variable "app_attr_rendering_metadata_regexp" {
  default = "regexp"
}

variable "app_attr_rendering_metadata_required" {
  default = false
}

variable "app_attr_rendering_metadata_section" {
  default = "saml"
}

variable "app_attr_rendering_metadata_visible" {
  default = false
}

variable "app_attr_rendering_metadata_widget" {
  default = "inputtext"
}

variable "app_attribute_sets" {
  default = []
}

variable "app_attributes" {
  default = "attributes"
}

variable "app_audience" {
  default = "audience"
}

variable "app_authorization" {
  default = "authorization"
}

variable "app_based_on_template_last_modified" {
  default = "lastModified"
}

variable "app_based_on_template_ref" {
  default = "ref"
}

variable "app_based_on_template_value" {
  default = "CustomWebAppTemplateId"
}

variable "app_bypass_consent" {
  default = false
}

variable "app_callback_service_url" {
  default = "callbackServiceUrl"
}

variable "app_certificates_cert_alias" {
  default = "certAlias"
}

variable "app_certificates_kid" {
  default = "kid"
}

variable "app_certificates_sha1thumbprint" {
  default = "sha1Thumbprint"
}

variable "app_certificates_x509base64certificate" {
  default = "{\"dummyKey\": \"dummyValue\"}"
}

variable "app_certificates_x5t" {
  default = "x5t"
}

variable "app_client_ip_checking" {
  default = "anywhere"
}

variable "app_client_secret" {
  default = "clientSecret"
}

variable "app_client_type" {
  default = "confidential"
}

variable "app_cloud_control_properties_name" {
  default = "name"
}

variable "app_cloud_control_properties_values" {
  default = []
}

variable "app_contact_email_address" {
  default = "contact@email.com"
}

variable "app_delegated_service_names" {
  default = []
}

variable "app_delete_in_progress" {
  default = false
}

variable "app_description" {
  default = "description"
}

variable "app_disable_kmsi_token_authentication" {
  default = false
}

variable "app_display_name" {
  default = "displayName"
}

variable "app_editable_attributes_name" {
  default = "name"
}

variable "app_error_page_url" {
  default = "https://testurl.com"
}

variable "app_granted_app_roles_admin_role" {
  default = false
}

variable "app_granted_app_roles_app_name" {
  default = "appName"
}

variable "app_granted_app_roles_display" {
  default = "display"
}

variable "app_granted_app_roles_read_only" {
  default = false
}

variable "app_granted_app_roles_ref" {
  default = "ref"
}

variable "app_granted_app_roles_type" {
  default = "direct"
}

variable "app_granted_app_roles_value" {
  default = "value"
}

variable "app_grants_grant_mechanism" {
  default = "IMPORT_APPROLE_MEMBERS"
}

variable "app_grants_grantee_type" {
  default = "User"
}

variable "app_grants_ref" {
  default = "ref"
}

variable "app_grants_value" {
  default = "value"
}

variable "app_hashed_client_secret" {
  default = "hashedClientSecret"
}

variable "app_home_page_url" {
  default = "https://testurl.com"
}

variable "app_icon" {
  default = "icon"
}

variable "app_id_token_enc_algo" {
  default = "A128CBC-HS256"
}

variable "app_identity_providers_display" {
  default = "display"
}

variable "app_identity_providers_ref" {
  default = "ref"
}

variable "app_identity_providers_value" {
  default = "value"
}

variable "app_idp_policy_ref" {
  default = "ref"
}

variable "app_idp_policy_value" {
  default = "value"
}

variable "app_infrastructure" {
  default = false
}

variable "app_is_alias_app" {
  default = "false"
}

variable "app_is_database_service" {
  default = false
}

variable "app_is_enterprise_app" {
  default = false
}

variable "app_is_form_fill" {
  default = false
}

variable "app_is_kerberos_realm" {
  default = false
}

variable "app_is_login_target" {
  default = false
}

variable "app_is_managed_app" {
  default = false
}

variable "app_is_mobile_target" {
  default = false
}

variable "app_is_multicloud_service_app" {
  default = "false"
}

variable "app_is_oauth_client" {
  default = false
}

variable "app_is_oauth_resource" {
  default = false
}

variable "app_is_opc_service" {
  default = false
}

variable "app_is_obligation_capable" {
  default = false
}

variable "app_is_radius_app" {
  default = false
}

variable "app_is_saml_service_provider" {
  default = false
}

variable "app_is_unmanaged_app" {
  default = "false"
}

variable "app_is_web_tier_policy" {
  default = false
}

variable "app_landing_page_url" {
  default = "https://testurl.com"
}

variable "app_linking_callback_url" {
  default = "https://testurl.com"
}

variable "app_login_mechanism" {
  default = "OIDC"
}

variable "app_login_page_url" {
  default = "https://testurl.com"
}

variable "app_logout_page_url" {
  default = "https://testurl.com"
}

variable "app_logout_uri" {
  default = "logoutUri"
}

variable "app_meter_as_opc_service" {
  default = false
}

variable "app_migrated" {
  default = false
}

variable "app_name" {
  default = "name"
}

variable "app_post_logout_redirect_uris" {
  default = []
}

variable "app_privacy_policy_url" {
  default = "privacyPolicyUrl"
}

variable "app_product_logo_url" {
  default = "productLogoUrl"
}

variable "app_product_name" {
  default = "productName"
}

variable "app_protectable_secondary_audiences_read_only" {
  default = false
}

variable "app_protectable_secondary_audiences_value" {
  default = "secondaryAudiences"
}

variable "app_radius_policy_ref" {
  default = "ref"
}

variable "app_radius_policy_value" {
  default = "value"
}

variable "app_ready_to_upgrade" {
  default = false
}

variable "app_redirect_uris" {
  default = []
}

variable "app_refresh_token_expiry" {
  default = 10
}

variable "app_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "app_saml_service_provider_ref" {
  default = "ref"
}

variable "app_saml_service_provider_value" {
  default = "value"
}

variable "app_schemas" {
  default = []
}

variable "app_scopes_description" {
  default = "description"
}

variable "app_scopes_display_name" {
  default = "displayName"
}

variable "app_scopes_fqs" {
  default = "fqs"
}

variable "app_scopes_read_only" {
  default = false
}

variable "app_scopes_requires_consent" {
  default = false
}

variable "app_scopes_value" {
  default = "value"
}

variable "app_secondary_audiences" {
  default = ["secondaryAudiences"]
}

variable "app_service_params_name" {
  default = "name"
}

variable "app_service_params_value" {
  default = "value"
}

variable "app_service_type_urn" {
  default = "serviceTypeURN"
}

variable "app_service_type_version" {
  default = "serviceTypeVersion"
}

variable "app_show_in_my_apps" {
  default = false
}

variable "app_signon_policy_ref" {
  default = "ref"
}

variable "app_signon_policy_value" {
  default = "value"
}

variable "app_start_index" {
  default = 1
}

variable "app_tags_key" {
  default = "key"
}

variable "app_tags_value" {
  default = "value"
}

variable "app_terms_of_service_url" {
  default = "https://testurl.com"
}

variable "app_terms_of_use_name" {
  default = "name"
}

variable "app_terms_of_use_ref" {
  default = "ref"
}

variable "app_terms_of_use_value" {
  default = "value"
}

variable "app_trust_policies_ref" {
  default = "ref"
}

variable "app_trust_policies_value" {
  default = "value"
}

variable "app_trust_scope" {
  default = "Explicit"
}

variable "app_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_key" {
  default = "key"
}

variable "app_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_namespace" {
  default = "namespace"
}

variable "app_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_value" {
  default = "value"
}

variable "app_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_key" {
  default = "freeformKey"
}

variable "app_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_value" {
  default = "freeformValue"
}

variable "app_urnietfparamsscimschemasoracleidcsextension_oci_tags_tag_slug" {
  default = "{\"dummyKey\": \"dummyValue\"}"
}

variable "app_urnietfparamsscimschemasoracleidcsextensiondbcs_app_domain_app_display" {
  default = "display"
}

variable "app_urnietfparamsscimschemasoracleidcsextensiondbcs_app_domain_app_ref" {
  default = "ref"
}

variable "app_urnietfparamsscimschemasoracleidcsextensiondbcs_app_domain_app_value" {
  default = "value"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_allow_authz_decision_ttl" {
  default = 10
}

variable "app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_allow_authz_policy_ref" {
  default = "ref"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_allow_authz_policy_value" {
  default = "value"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_app_resources_ref" {
  default = "ref"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_app_resources_value" {
  default = "value"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_deny_authz_decision_ttl" {
  default = 10
}

variable "app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_deny_authz_policy_ref" {
  default = "ref"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_deny_authz_policy_value" {
  default = "value"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_configuration" {
  default = "configuration"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_cred_method" {
  default = "ADMIN_SETS_CREDENTIALS"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_credential_sharing_group_id" {
  default = "formCredentialSharingGroupID"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_fill_url_match_form_url" {
  default = "formUrl"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_fill_url_match_form_url_match_type" {
  default = "exact"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_type" {
  default = "WebApplication"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_reveal_password_on_form" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_sync_from_template" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_user_name_form_expression" {
  default = "concat($user.firstname,\".\",$user.lastname)"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_user_name_form_template" {
  default = "username"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_configuration" {
  default = "configuration"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_cred_method" {
  default = "ADMIN_SETS_CREDENTIALS"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_credential_sharing_group_id" {
  default = "formCredentialSharingGroupID"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_fill_url_match_form_url" {
  default = "formUrl"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_fill_url_match_form_url_match_type" {
  default = "exact"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_type" {
  default = "WebApplication"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_reveal_password_on_form" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_sync_from_template" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_user_name_form_expression" {
  default = "concat($user.firstname,\".\",$user.lastname)"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_user_name_form_template" {
  default = "username"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_default_encryption_salt_type" {
  default = "defaultEncryptionSaltType"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_master_key" {
  default = "masterKey"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_max_renewable_age" {
  default = 10
}

variable "app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_max_ticket_life" {
  default = 10
}

variable "app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_realm_name" {
  default = "realmName"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_supported_encryption_salt_types" {
  default = []
}

variable "app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_ticket_flags" {
  default = 10
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_account_form_visible" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_admin_consent_granted" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_confidential" {
  default = "false"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_display_name" {
  default = "displayName"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_help_message" {
  default = "helpMessage"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_icf_type" {
  default = "Long"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_name" {
  default = "name"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_order" {
  default = 10
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_required" {
  default = "false"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_value" {
  default = []
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_max_idle" {
  default = 10
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_max_objects" {
  default = 10
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_max_wait" {
  default = 10
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_min_evictable_idle_time_millis" {
  default = 10
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_min_idle" {
  default = 10
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_can_be_authoritative" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_connected" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_connector_bundle_display" {
  default = "display"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_connector_bundle_ref" {
  default = "ref"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_connector_bundle_type" {
  default = "ConnectorBundle"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_connector_bundle_value" {
  default = "value"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_enable_auth_sync_new_user_notification" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_enable_sync" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_enable_sync_summary_report_notification" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_confidential" {
  default = "false"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_display_name" {
  default = "displayName"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_help_message" {
  default = "helpMessage"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_icf_type" {
  default = "Long"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_name" {
  default = "name"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_order" {
  default = 10
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_required" {
  default = "false"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_value" {
  default = []
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_connector_bundle_display" {
  default = "display"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_connector_bundle_ref" {
  default = "ref"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_connector_bundle_value" {
  default = "value"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_identity_bridges_name" {
  default = "name"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_identity_bridges_ref" {
  default = "ref"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_identity_bridges_value" {
  default = "value"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_is_authoritative" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_is_directory" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_is_on_premise_app" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_is_schema_customization_supported" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_is_schema_discovery_supported" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_is_three_legged_oauth_enabled" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_is_two_legged_oauth_enabled" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_object_classes_display" {
  default = "display"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_object_classes_is_account_object_class" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_object_classes_ref" {
  default = "ref"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_object_classes_resource_type" {
  default = "resourceType"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_object_classes_type" {
  default = "AccountObjectClass"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_object_classes_value" {
  default = "value"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_sync_config_last_modified" {
  default = "syncConfigLastModified"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_three_legged_oauth_credential_access_token" {
  default = "accessToken"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_three_legged_oauth_credential_access_token_expiry" {
  default = "2032-01-01T00:00:00Z"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_three_legged_oauth_credential_refresh_token" {
  default = "refreshToken"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_three_legged_oauth_provider_name" {
  default = "threeLeggedOAuthProviderName"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app_multicloud_platform_url" {
  default = "multicloudPlatformUrl"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app_multicloud_service_type" {
  default = "AWSCognito"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionopc_service_app_current_federation_mode" {
  default = "None"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionopc_service_app_current_synchronization_mode" {
  default = "None"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionopc_service_app_enabling_next_fed_sync_modes" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionopc_service_app_next_federation_mode" {
  default = "None"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionopc_service_app_next_synchronization_mode" {
  default = "None"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionopc_service_app_region" {
  default = "region"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionopc_service_app_service_instance_identifier" {
  default = "serviceInstanceIdentifier"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_capture_client_ip" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_client_ip" {
  default = "clientIP"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_end_user_ip_attribute" {
  default = "31 Calling-Station-Id"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_group_membership_radius_attribute" {
  default = "groupMembershipRadiusAttribute"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_group_membership_to_return_display" {
  default = "display"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_group_membership_to_return_ref" {
  default = "ref"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_group_membership_to_return_value" {
  default = "value"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_group_name_format" {
  default = "groupNameFormat"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_include_group_in_response" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_password_and_otp_together" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_port" {
  default = "port"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_response_format" {
  default = "responseFormat"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_response_format_delimiter" {
  default = "responseFormatDelimiter"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_secret_key" {
  default = "secretKey"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_type_of_radius_app" {
  default = "Oracle Database"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionrequestable_app_requestable" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_assertion_consumer_url" {
  default = "https://testurl.com"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_encrypt_assertion" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_encryption_algorithm" {
  default = "3DES"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_federation_protocol" {
  default = "SAML2.0"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_group_assertion_attributes_condition" {
  default = "Starts With"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_group_assertion_attributes_format" {
  default = "Basic"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_group_assertion_attributes_name" {
  default = "groupName"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_hok_acs_url" {
  default = "hokAcsUrl"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_hok_required" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_include_signing_cert_in_signature" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_key_encryption_algorithm" {
  default = "RSA-v1.5"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_last_notification_sent_time" {
  default = "lastNotificationSentTime"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_logout_binding" {
  default = "Redirect"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_logout_enabled" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_logout_request_url" {
  default = "https://testurl.com"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_logout_response_url" {
  default = "https://testurl.com"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_name_id_format" {
  default = "nameIdFormat"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_name_id_userstore_attribute" {
  default = "emails.primary.value"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_outbound_assertion_attributes_direction" {
  default = "direction"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_outbound_assertion_attributes_ref" {
  default = "ref"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_outbound_assertion_attributes_value" {
  default = "value"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_partner_provider_pattern" {
  default = "partnerProviderPattern"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_sign_response_or_assertion" {
  default = "Assertion"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_signature_hash_algorithm" {
  default = "SHA-1"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_user_assertion_attributes_format" {
  default = "Basic"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_user_assertion_attributes_name" {
  default = "userName"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_user_assertion_attributes_user_store_attribute_name" {
  default = "emails.primary.value"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app_resource_ref" {
  default = false
}

variable "app_urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app_web_tier_policy_az_control" {
  default = "server"
}

variable "app_urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app_web_tier_policy_json" {
  default = "{\"cloudgatePolicy\":{\"version\":\"2.6\",\"disableAuthorize\":false,\"webtierPolicy\":[{\"policyName\":\"test\",\"resourceFilters\":[]}]}}"
}

variable "app_user_roles_description" {
  default = "description"
}

variable "app_user_roles_display" {
  default = "display"
}

variable "app_user_roles_ref" {
  default = "ref"
}

variable "app_user_roles_value" {
  default = "value"
}


resource "oci_identity_domains_app" "test_app" {
  #Required
  based_on_template {
    #Required
    value = var.app_based_on_template_value

    #Optional
    well_known_id = "CustomWebAppTemplateId"
  }
  display_name  = var.app_display_name
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:App"]

  #Optional
  access_token_expiry = var.app_access_token_expiry
  active              = var.app_active
  /* #provide App id of an alias app
  alias_apps {
    #Required
    value = var.app_alias_apps_value
  }
  */
  all_url_schemes_allowed = var.app_all_url_schemes_allowed
  allow_access_control    = var.app_allow_access_control
  allow_offline           = var.app_allow_offline
  allowed_grants          = var.app_allowed_grants
  allowed_operations      = var.app_allowed_operations
  /* provide scope defined by App
  allowed_scopes {
    #Required
    fqs = var.app_allowed_scopes_fqs
  }
  */
  allowed_tags {
    #Required
    key   = var.app_allowed_tags_key
    value = var.app_allowed_tags_value
  }
  app_icon = var.app_app_icon
  /* #provide Policy id
  app_signon_policy {
    #Required
    value = var.app_app_signon_policy_value
  }
  */
  app_thumbnail = var.app_app_thumbnail
  /* #provide NetworkPerimeter id
  apps_network_perimeters {
    #Required
    value = var.app_apps_network_perimeters_value
  }
  */
  attr_rendering_metadata {
    #Required
    name = var.app_attr_rendering_metadata_name

    #Optional
    datatype   = var.app_attr_rendering_metadata_datatype
    helptext   = var.app_attr_rendering_metadata_helptext
    label      = var.app_attr_rendering_metadata_label
    max_length = var.app_attr_rendering_metadata_max_length
    max_size   = var.app_attr_rendering_metadata_max_size
    min_length = var.app_attr_rendering_metadata_min_length
    min_size   = var.app_attr_rendering_metadata_min_size
    order      = var.app_attr_rendering_metadata_order
    read_only  = var.app_attr_rendering_metadata_read_only
    regexp     = var.app_attr_rendering_metadata_regexp
    required   = var.app_attr_rendering_metadata_required
    section    = var.app_attr_rendering_metadata_section
    visible    = var.app_attr_rendering_metadata_visible
    widget     = var.app_attr_rendering_metadata_widget
  }
  attribute_sets = ["all"]
  attributes     = ""
  audience       = var.app_audience
  authorization  = var.app_authorization
  bypass_consent = var.app_bypass_consent
  /* #provide Certificate alias
  certificates {
    #Required
    cert_alias = var.app_certificates_cert_alias
  }
  */
  client_ip_checking                = var.app_client_ip_checking
  client_type                       = var.app_client_type
  contact_email_address             = var.app_contact_email_address
  delegated_service_names           = var.app_delegated_service_names
  description                       = var.app_description
  disable_kmsi_token_authentication = var.app_disable_kmsi_token_authentication
  error_page_url                    = var.app_error_page_url
  home_page_url                     = var.app_home_page_url
  icon                              = var.app_icon
  id_token_enc_algo                 = var.app_id_token_enc_algo
  /* #provide IDP id
  identity_providers {
    #Required
    value = var.app_identity_providers_value
  }
  */
  /* #provide Policy id
  idp_policy {
    #Required
    value = var.app_idp_policy_value
  }
  */
  is_alias_app              = var.app_is_alias_app
  is_enterprise_app         = var.app_is_enterprise_app
  is_form_fill              = var.app_is_form_fill
  is_kerberos_realm         = var.app_is_kerberos_realm
  is_login_target           = var.app_is_login_target
  is_mobile_target          = var.app_is_mobile_target
  is_multicloud_service_app = var.app_is_multicloud_service_app
  is_oauth_client           = var.app_is_oauth_client
  is_oauth_resource         = var.app_is_oauth_resource
  is_obligation_capable     = var.app_is_obligation_capable
  is_radius_app             = var.app_is_radius_app
  is_saml_service_provider  = var.app_is_saml_service_provider
  is_unmanaged_app          = var.app_is_unmanaged_app
  is_web_tier_policy        = var.app_is_web_tier_policy
  landing_page_url          = var.app_landing_page_url
  linking_callback_url      = var.app_linking_callback_url
  login_mechanism           = var.app_login_mechanism
  login_page_url            = var.app_login_page_url
  logout_page_url           = var.app_logout_page_url
  logout_uri                = var.app_logout_uri
  name                      = var.app_name
  post_logout_redirect_uris = var.app_post_logout_redirect_uris
  privacy_policy_url        = var.app_privacy_policy_url
  product_logo_url          = var.app_product_logo_url
  product_name              = var.app_product_name
  protectable_secondary_audiences {
    #Required
    value = var.app_protectable_secondary_audiences_value
  }
  /* #provide Policy id
  radius_policy {
    #Required
    value = var.app_radius_policy_value
  }
  */
  redirect_uris                = var.app_redirect_uris
  refresh_token_expiry         = var.app_refresh_token_expiry
  #use the latest if not provided
  # resource_type_schema_version = var.app_resource_type_schema_version
  /* #provide SAML App id
  saml_service_provider {
    #Required
    value = var.app_saml_service_provider_value
  }
  */
  scopes {
    #Required
    value = var.app_scopes_value

    #Optional
    description      = var.app_scopes_description
    display_name     = var.app_scopes_display_name
    requires_consent = var.app_scopes_requires_consent
  }
  secondary_audiences = var.app_secondary_audiences
  service_params {
    #Required
    name = var.app_service_params_name

    #Optional
    value = var.app_service_params_value
  }
  service_type_urn     = var.app_service_type_urn
  service_type_version = var.app_service_type_version
  show_in_my_apps      = var.app_show_in_my_apps
  /* #provide Policy id
  signon_policy {
    #Required
    value = var.app_signon_policy_value
  }
  */
  tags {
    #Required
    key   = var.app_tags_key
    value = var.app_tags_value
  }
  terms_of_service_url = var.app_terms_of_service_url
  /* #provide TermsOfUse id
  terms_of_use {
    #Required
    value = var.app_terms_of_use_value
  }
  */
  /* #provide Policy id
  trust_policies {
    #Required
    value = var.app_trust_policies_value
  }
  */
  trust_scope = var.app_trust_scope
  urnietfparamsscimschemasoracleidcsextension_oci_tags {

    #Optional
    /* #create tagNamespace to use defined tags
    defined_tags {
      #Required
      key       = var.app_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_key
      namespace = var.app_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_namespace
      value     = var.app_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_value
    }
    */
    freeform_tags {
      #Required
      key   = var.app_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_key
      value = var.app_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_value
    }
  }
  /* #provide DB Domain info
  urnietfparamsscimschemasoracleidcsextensiondbcs_app {

    #Optional
    domain_app {
      #Required
      value = var.app_urnietfparamsscimschemasoracleidcsextensiondbcs_app_domain_app_value
    }
    domain_name = "domainName"
  }
  */
  urnietfparamsscimschemasoracleidcsextensionenterprise_app_app {

    #Optional
    allow_authz_decision_ttl = var.app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_allow_authz_decision_ttl
    /* #provide Policy id
    allow_authz_policy {
      #Required
      value = var.app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_allow_authz_policy_value
    }
    */
    /* #provide AppResource id
    app_resources {
      #Required
      value = var.app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_app_resources_value
    }
    */
    deny_authz_decision_ttl = var.app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_deny_authz_decision_ttl
    /* #provide Policy id
    deny_authz_policy {
      #Required
      value = var.app_urnietfparamsscimschemasoracleidcsextensionenterprise_app_app_deny_authz_policy_value
    }
    */
  }
  urnietfparamsscimschemasoracleidcsextensionform_fill_app_app {

    #Optional
    configuration                    = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_configuration
    form_cred_method                 = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_cred_method
    form_credential_sharing_group_id = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_credential_sharing_group_id
    form_fill_url_match {
      #Required
      form_url = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_fill_url_match_form_url

      #Optional
      form_url_match_type = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_fill_url_match_form_url_match_type
    }
    form_type                 = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_form_type
    reveal_password_on_form   = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_reveal_password_on_form
    user_name_form_expression = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_user_name_form_expression
    user_name_form_template   = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_app_user_name_form_template
  }
  urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template {

    #Optional
    configuration                    = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_configuration
    form_cred_method                 = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_cred_method
    form_credential_sharing_group_id = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_credential_sharing_group_id
    form_fill_url_match {
      #Required
      form_url = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_fill_url_match_form_url

      #Optional
      form_url_match_type = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_fill_url_match_form_url_match_type
    }
    form_type                 = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_form_type
    reveal_password_on_form   = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_reveal_password_on_form
    sync_from_template        = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_sync_from_template
    user_name_form_expression = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_user_name_form_expression
    user_name_form_template   = var.app_urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template_user_name_form_template
  }
  urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app {

    #Optional
    default_encryption_salt_type    = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_default_encryption_salt_type
    master_key                      = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_master_key
    max_renewable_age               = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_max_renewable_age
    max_ticket_life                 = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_max_ticket_life
    realm_name                      = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_realm_name
    supported_encryption_salt_types = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_supported_encryption_salt_types
    ticket_flags                    = var.app_urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app_ticket_flags
  }
  urnietfparamsscimschemasoracleidcsextensionmanagedapp_app {

    #Optional
    admin_consent_granted = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_admin_consent_granted
    bundle_configuration_properties {
      #Required
      icf_type = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_icf_type
      name     = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_name
      required = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_required

      #Optional
      confidential = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_confidential
      display_name = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_display_name
      help_message = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_help_message
      order        = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_order
      value        = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_configuration_properties_value
    }
    bundle_pool_configuration {

      #Optional
      max_idle                       = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_max_idle
      max_objects                    = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_max_objects
      max_wait                       = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_max_wait
      min_evictable_idle_time_millis = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_min_evictable_idle_time_millis
      min_idle                       = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_bundle_pool_configuration_min_idle
    }
    connected                               = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_connected
    enable_auth_sync_new_user_notification  = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_enable_auth_sync_new_user_notification
    enable_sync                             = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_enable_sync
    enable_sync_summary_report_notification = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_enable_sync_summary_report_notification
    flat_file_bundle_configuration_properties {
      #Required
      icf_type = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_icf_type
      name     = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_name
      required = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_required

      #Optional
      confidential = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_confidential
      display_name = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_display_name
      help_message = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_help_message
      order        = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_order
      value        = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_bundle_configuration_properties_value
    }
    /* #provide ConnectorBundle info
    flat_file_connector_bundle {
      #Required
      value = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_connector_bundle_value

      #Optional
      display       = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_flat_file_connector_bundle_display
      well_known_id = ""
    }
    */
    is_authoritative = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_is_authoritative
    three_legged_oauth_credential {

      #Optional
      access_token        = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_three_legged_oauth_credential_access_token
      access_token_expiry = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_three_legged_oauth_credential_access_token_expiry
      refresh_token       = var.app_urnietfparamsscimschemasoracleidcsextensionmanagedapp_app_three_legged_oauth_credential_refresh_token
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
    client_ip                 = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_client_ip
    include_group_in_response = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_include_group_in_response
    port                      = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_port
    secret_key                = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_secret_key

    #Optional
    capture_client_ip                  = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_capture_client_ip
    country_code_response_attribute_id = "1"
    end_user_ip_attribute              = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_end_user_ip_attribute
    group_membership_radius_attribute  = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_group_membership_radius_attribute
    /* #provide Group id
    group_membership_to_return {
      #Required
      value = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_group_membership_to_return_value
    }
    */
    group_name_format         = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_group_name_format
    password_and_otp_together = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_password_and_otp_together
    radius_vendor_specific_id = "radiusVendorSpecificId"
    response_format           = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_response_format
    response_format_delimiter = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_response_format_delimiter
    type_of_radius_app        = var.app_urnietfparamsscimschemasoracleidcsextensionradius_app_app_type_of_radius_app
  }
  urnietfparamsscimschemasoracleidcsextensionrequestable_app {

    #Optional
    requestable = var.app_urnietfparamsscimschemasoracleidcsextensionrequestable_app_requestable
  }
  urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app {

    #Optional
    assertion_consumer_url = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_assertion_consumer_url
    encrypt_assertion      = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_encrypt_assertion
    encryption_algorithm   = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_encryption_algorithm
    /* #provide certificate
    encryption_certificate = "encryptionCertificate"
    */
    federation_protocol    = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_federation_protocol
    group_assertion_attributes {
      #Required
      name = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_group_assertion_attributes_name

      #Optional
      condition  = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_group_assertion_attributes_condition
      format     = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_group_assertion_attributes_format
      group_name = "groupName"
    }
    /* provide SAML Holder-of-Key infos
    hok_acs_url                       = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_hok_acs_url
    hok_required                      = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_hok_required
    */
    include_signing_cert_in_signature = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_include_signing_cert_in_signature
    key_encryption_algorithm          = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_key_encryption_algorithm
    logout_binding                    = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_logout_binding
    logout_enabled                    = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_logout_enabled
    logout_request_url                = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_logout_request_url
    logout_response_url               = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_logout_response_url
    /* #provide metadata
    metadata                          = "metadata"
    */
    name_id_format                    = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_name_id_format
    name_id_userstore_attribute       = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_name_id_userstore_attribute
    partner_provider_id               = "partnerProviderId"
    partner_provider_pattern          = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_partner_provider_pattern
    sign_response_or_assertion        = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_sign_response_or_assertion
    signature_hash_algorithm          = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_signature_hash_algorithm
    /* #provide certificate
    signing_certificate               = "signingCertificate"
    */
    succinct_id                       = "succinctId"
    user_assertion_attributes {
      #Required
      name                      = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_user_assertion_attributes_name
      user_store_attribute_name = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_user_assertion_attributes_user_store_attribute_name

      #Optional
      format = var.app_urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app_user_assertion_attributes_format
    }
  }
  /* provide Web Tier Policy infos
  urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app {

    #Optional
    resource_ref               = var.app_urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app_resource_ref
    web_tier_policy_az_control = var.app_urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app_web_tier_policy_az_control
    web_tier_policy_json       = var.app_urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app_web_tier_policy_json
  }
  */

  lifecycle {
    ignore_changes = [schemas]
  }
}

data "oci_identity_domains_apps" "test_apps" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  app_count                    = var.app_app_count
  app_filter                   = var.app_app_filter
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.app_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.app_resource_type_schema_version
  start_index                  = var.app_start_index
}

