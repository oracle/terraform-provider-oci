// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "setting_account_always_trust_scope" {
  default = false
}

variable "setting_allowed_forgot_password_flow_return_urls" {
  default = []
}

variable "setting_allowed_notification_redirect_urls" {
  default = []
}

variable "setting_audit_event_retention_period" {
  default = "30"
}

variable "setting_authorization" {
  default = "authorization"
}

variable "setting_certificate_validation_crl_check_on_ocsp_failure_enabled" {
  default = false
}

variable "setting_certificate_validation_crl_enabled" {
  default = false
}

variable "setting_certificate_validation_crl_location" {
  default = "crlLocation"
}

variable "setting_certificate_validation_crl_refresh_interval" {
  default = 10
}

variable "setting_certificate_validation_ocsp_enabled" {
  default = false
}

variable "setting_certificate_validation_ocsp_responder_url" {
  default = "ocspResponderURL"
}

variable "setting_certificate_validation_ocsp_settings_responder_url_preferred" {
  default = false
}

variable "setting_certificate_validation_ocsp_signing_certificate_alias" {
  default = "ocspSigningCertificateAlias"
}

variable "setting_certificate_validation_ocsp_timeout_duration" {
  default = 10
}

variable "setting_certificate_validation_ocsp_unknown_response_status_allowed" {
  default = false
}

variable "setting_cloud_account_name" {
  default = "cloudAccountName"
}

variable "setting_cloud_gate_cors_settings_cloud_gate_cors_allow_null_origin" {
  default = false
}

variable "setting_cloud_gate_cors_settings_cloud_gate_cors_enabled" {
  default = false
}

variable "setting_cloud_gate_cors_settings_cloud_gate_cors_exposed_headers" {
  default = []
}

variable "setting_cloud_gate_cors_settings_cloud_gate_cors_max_age" {
  default = 10
}

variable "setting_cloud_migration_custom_url" {
  default = "cloudMigrationCustomUrl"
}

variable "setting_cloud_migration_url_enabled" {
  default = false
}

variable "setting_company_names_locale" {
  default = "en"
}

variable "setting_company_names_value" {
  default = "value"
}

variable "setting_csr_access" {
  default = "readOnly"
}

variable "setting_custom_branding" {
  default = false
}

variable "setting_custom_css_location" {
  default = "customCssLocation"
}

variable "setting_custom_html_location" {
  default = "customHtmlLocation"
}

variable "setting_custom_translation" {
  default = "customTranslation"
}

variable "setting_default_trust_scope" {
  default = "Explicit"
}

variable "setting_diagnostic_level" {
  default = "0"
}

variable "setting_diagnostic_record_for_search_identifies_returned_resources" {
  default = false
}

variable "setting_enable_terms_of_use" {
  default = false
}

variable "setting_iam_upst_session_expiry" {
  default = "0"
}

variable "setting_images_display" {
  default = "display"
}

variable "setting_images_type" {
  default = "desktop logo"
}

variable "setting_images_value" {
  default = "https://idcs-guid.identity.oraclecloud.com/oracle-desktop-logo.gif"
}

variable "setting_is_hosted_page" {
  default = false
}

variable "setting_issuer" {
  default = "issuer"
}

variable "setting_locale" {
  default = "en"
}

variable "setting_login_texts_locale" {
  default = "en"
}

variable "setting_login_texts_value" {
  default = "value"
}

variable "setting_max_no_of_app_cmva_to_return" {
  default = 10
}

variable "setting_max_no_of_app_role_members_to_return" {
  default = 10
}

variable "setting_migration_status" {
  default = "migrationStatus"
}

variable "setting_on_premises_provisioning" {
  default = false
}

variable "setting_preferred_language" {
  default = "en"
}

variable "setting_prev_issuer" {
  default = "prevIssuer"
}

variable "setting_privacy_policy_url" {
  default = "privacyPolicyUrl"
}

variable "setting_purge_configs_retention_period" {
  default = "30"
}

variable "setting_re_auth_when_changing_my_authentication_factors" {
  default = false
}

variable "setting_service_admin_cannot_list_other_users" {
  default = false
}

variable "setting_signing_cert_public_access" {
  default = false
}

variable "setting_sub_mapping_attr" {
  default = "userName"
}

variable "setting_tags_key" {
  default = "key"
}

variable "setting_tags_value" {
  default = "value"
}

variable "setting_tenant_custom_claims_all_scopes" {
  default = false
}

variable "setting_tenant_custom_claims_expression" {
  default = false
}

variable "setting_tenant_custom_claims_mode" {
  default = "always"
}

variable "setting_tenant_custom_claims_name" {
  default = "customClaimName"
}

variable "setting_tenant_custom_claims_token_type" {
  default = "AT"
}

variable "setting_tenant_custom_claims_value" {
  default = "value"
}

variable "setting_terms_of_use_url" {
  default = "termsOfUseUrl"
}

variable "setting_timezone" {
  default = "America/Los_Angeles"
}



resource "oci_identity_domains_setting" "test_setting" {
  #Required
  csr_access    = var.setting_csr_access
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:Settings"]
  setting_id    = "Settings"

  #Optional
  account_always_trust_scope               = var.setting_account_always_trust_scope
  allowed_domains                          = ["test.com"]
  allowed_forgot_password_flow_return_urls = var.setting_allowed_forgot_password_flow_return_urls
  allowed_notification_redirect_urls       = var.setting_allowed_notification_redirect_urls
  attribute_sets                           = ["all"]
  attributes                               = ""
  audit_event_retention_period             = var.setting_audit_event_retention_period
  authorization                            = var.setting_authorization
  certificate_validation {

    #Optional
    crl_check_on_ocsp_failure_enabled     = var.setting_certificate_validation_crl_check_on_ocsp_failure_enabled
    crl_enabled                           = var.setting_certificate_validation_crl_enabled
    crl_location                          = var.setting_certificate_validation_crl_location
    crl_refresh_interval                  = var.setting_certificate_validation_crl_refresh_interval
    ocsp_enabled                          = var.setting_certificate_validation_ocsp_enabled
    ocsp_responder_url                    = var.setting_certificate_validation_ocsp_responder_url
    ocsp_settings_responder_url_preferred = var.setting_certificate_validation_ocsp_settings_responder_url_preferred
    ocsp_signing_certificate_alias        = var.setting_certificate_validation_ocsp_signing_certificate_alias
    ocsp_timeout_duration                 = var.setting_certificate_validation_ocsp_timeout_duration
    ocsp_unknown_response_status_allowed  = var.setting_certificate_validation_ocsp_unknown_response_status_allowed
  }
  cloud_gate_cors_settings {

    #Optional
    cloud_gate_cors_allow_null_origin = var.setting_cloud_gate_cors_settings_cloud_gate_cors_allow_null_origin
    cloud_gate_cors_allowed_origins   = ["https://test.com"]
    cloud_gate_cors_enabled           = var.setting_cloud_gate_cors_settings_cloud_gate_cors_enabled
    cloud_gate_cors_exposed_headers   = var.setting_cloud_gate_cors_settings_cloud_gate_cors_exposed_headers
    cloud_gate_cors_max_age           = var.setting_cloud_gate_cors_settings_cloud_gate_cors_max_age
  }
  cloud_migration_custom_url  = var.setting_cloud_migration_custom_url
  cloud_migration_url_enabled = var.setting_cloud_migration_url_enabled
  company_names {
    #Required
    locale = var.setting_company_names_locale
    value  = var.setting_company_names_value
  }
  contact_emails                                             = ["contactEmails@test.com"]
  custom_branding                                            = var.setting_custom_branding
  custom_css_location                                        = var.setting_custom_css_location
  custom_html_location                                       = var.setting_custom_html_location
  custom_translation                                         = var.setting_custom_translation
  default_trust_scope                                        = var.setting_default_trust_scope
  diagnostic_level                                           = var.setting_diagnostic_level
  diagnostic_record_for_search_identifies_returned_resources = var.setting_diagnostic_record_for_search_identifies_returned_resources
  enable_terms_of_use                                        = var.setting_enable_terms_of_use
  external_id                                                = "externalId"
  iam_upst_session_expiry                                    = var.setting_iam_upst_session_expiry
  images {
    #Required
    type  = var.setting_images_type
    value = var.setting_images_value

    #Optional
    display = var.setting_images_display
  }
  is_hosted_page = var.setting_is_hosted_page
  issuer         = var.setting_issuer
  locale         = var.setting_locale
  login_texts {
    #Required
    locale = var.setting_login_texts_locale
    value  = var.setting_login_texts_value
  }
  max_no_of_app_cmva_to_return         = var.setting_max_no_of_app_cmva_to_return
  max_no_of_app_role_members_to_return = var.setting_max_no_of_app_role_members_to_return
  preferred_language                   = var.setting_preferred_language
  privacy_policy_url                   = var.setting_privacy_policy_url
  purge_configs {
    #Required
    resource_name    = "resourceName"
    retention_period = var.setting_purge_configs_retention_period
  }
  re_auth_factor                                  = ["password"]
  re_auth_when_changing_my_authentication_factors = var.setting_re_auth_when_changing_my_authentication_factors
  #use the latest version if not provided
  # resource_type_schema_version                    = var.setting_resource_type_schema_version
  service_admin_cannot_list_other_users           = var.setting_service_admin_cannot_list_other_users
  signing_cert_public_access                      = var.setting_signing_cert_public_access
  sub_mapping_attr                                = var.setting_sub_mapping_attr
  tags {
    #Required
    key   = var.setting_tags_key
    value = var.setting_tags_value
  }
  tenant_custom_claims {
    #Required
    all_scopes = var.setting_tenant_custom_claims_all_scopes
    expression = var.setting_tenant_custom_claims_expression
    mode       = var.setting_tenant_custom_claims_mode
    name       = var.setting_tenant_custom_claims_name
    token_type = var.setting_tenant_custom_claims_token_type
    value      = var.setting_tenant_custom_claims_value

    #Optional
    scopes = ["scopes"]
  }
  terms_of_use_url = var.setting_terms_of_use_url
  timezone         = var.setting_timezone
}

data "oci_identity_domains_settings" "test_settings" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.setting_authorization
  #use the latest version if not provided
  # resource_type_schema_version = var.setting_resource_type_schema_version
}

