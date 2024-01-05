// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "authentication_factor_setting_authorization" {
  default = "authorization"
}

variable "authentication_factor_setting_auto_enroll_email_factor_disabled" {
  default = false
}

variable "authentication_factor_setting_bypass_code_enabled" {
  default = false
}

variable "authentication_factor_setting_bypass_code_settings_help_desk_code_expiry_in_mins" {
  default = 10
}

variable "authentication_factor_setting_bypass_code_settings_help_desk_generation_enabled" {
  default = false
}

variable "authentication_factor_setting_bypass_code_settings_help_desk_max_usage" {
  default = 10
}

variable "authentication_factor_setting_bypass_code_settings_length" {
  default = 10
}

variable "authentication_factor_setting_bypass_code_settings_max_active" {
  default = "5"
}

variable "authentication_factor_setting_bypass_code_settings_self_service_generation_enabled" {
  default = false
}

variable "authentication_factor_setting_client_app_settings_device_protection_policy" {
  default = "APP_PIN"
}

variable "authentication_factor_setting_client_app_settings_initial_lockout_period_in_secs" {
  default = "60"
}

variable "authentication_factor_setting_client_app_settings_key_pair_length" {
  default = "32"
}

variable "authentication_factor_setting_client_app_settings_lockout_escalation_pattern" {
  default = "Linear"
}

variable "authentication_factor_setting_client_app_settings_max_failures_before_lockout" {
  default = "5"
}

variable "authentication_factor_setting_client_app_settings_max_failures_before_warning" {
  default = "0"
}

variable "authentication_factor_setting_client_app_settings_max_lockout_interval_in_secs" {
  default = "90"
}

variable "authentication_factor_setting_client_app_settings_min_pin_length" {
  default = "10"
}

variable "authentication_factor_setting_client_app_settings_policy_update_freq_in_days" {
  default = 10
}

variable "authentication_factor_setting_client_app_settings_request_signing_algo" {
  default = "SHA256withRSA"
}

variable "authentication_factor_setting_client_app_settings_shared_secret_encoding" {
  default = "Base32"
}

variable "authentication_factor_setting_client_app_settings_unlock_app_for_each_request_enabled" {
  default = "true"
}

variable "authentication_factor_setting_client_app_settings_unlock_app_interval_in_secs" {
  default = "0"
}

variable "authentication_factor_setting_client_app_settings_unlock_on_app_foreground_enabled" {
  default = "true"
}

variable "authentication_factor_setting_client_app_settings_unlock_on_app_start_enabled" {
  default = "true"
}

variable "authentication_factor_setting_email_enabled" {
  default = false
}

variable "authentication_factor_setting_email_settings_email_link_custom_url" {
  default = "emailLinkCustomUrl"
}

variable "authentication_factor_setting_email_settings_email_link_enabled" {
  default = false
}

variable "authentication_factor_setting_endpoint_restrictions_max_endpoint_trust_duration_in_days" {
  default = 10
}

variable "authentication_factor_setting_endpoint_restrictions_max_enrolled_devices" {
  default = 10
}

variable "authentication_factor_setting_endpoint_restrictions_max_incorrect_attempts" {
  default = 10
}

variable "authentication_factor_setting_endpoint_restrictions_max_trusted_endpoints" {
  default = 10
}

variable "authentication_factor_setting_endpoint_restrictions_trusted_endpoints_enabled" {
  default = false
}

variable "authentication_factor_setting_fido_authenticator_enabled" {
  default = false
}

variable "authentication_factor_setting_hide_backup_factor_enabled" {
  default = false
}

variable "authentication_factor_setting_identity_store_settings_mobile_number_enabled" {
  default = false
}

variable "authentication_factor_setting_identity_store_settings_mobile_number_update_enabled" {
  default = false
}

variable "authentication_factor_setting_mfa_enrollment_type" {
  default = "Optional"
}

variable "authentication_factor_setting_notification_settings_pull_enabled" {
  default = false
}

variable "authentication_factor_setting_phone_call_enabled" {
  default = "false"
}

variable "authentication_factor_setting_push_enabled" {
  default = false
}

variable "authentication_factor_setting_security_questions_enabled" {
  default = false
}

variable "authentication_factor_setting_sms_enabled" {
  default = "true"
}

variable "authentication_factor_setting_tags_key" {
  default = "key"
}

variable "authentication_factor_setting_tags_value" {
  default = "value"
}

variable "authentication_factor_setting_tenancy_ocid" {
  default = "tenancyOcid"
}

variable "authentication_factor_setting_third_party_factor_duo_security" {
  default = false
}

variable "authentication_factor_setting_totp_enabled" {
  default = false
}

variable "authentication_factor_setting_totp_settings_email_otp_validity_duration_in_mins" {
  default = 10
}

variable "authentication_factor_setting_totp_settings_email_passcode_length" {
  default = 10
}

variable "authentication_factor_setting_totp_settings_hashing_algorithm" {
  default = "SHA1"
}

variable "authentication_factor_setting_totp_settings_jwt_validity_duration_in_secs" {
  default = "30"
}

variable "authentication_factor_setting_totp_settings_key_refresh_interval_in_days" {
  default = "30"
}

variable "authentication_factor_setting_totp_settings_passcode_length" {
  default = 10
}

variable "authentication_factor_setting_totp_settings_sms_otp_validity_duration_in_mins" {
  default = 10
}

variable "authentication_factor_setting_totp_settings_sms_passcode_length" {
  default = 10
}

variable "authentication_factor_setting_totp_settings_time_step_in_secs" {
  default = "300"
}

variable "authentication_factor_setting_totp_settings_time_step_tolerance" {
  default = "2"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings_attestation" {
  default = "DIRECT"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings_authenticator_selection_attachment" {
  default = "PLATFORM"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings_authenticator_selection_require_resident_key" {
  default = "true"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings_authenticator_selection_resident_key" {
  default = "REQUIRED"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings_authenticator_selection_user_verification" {
  default = "REQUIRED"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings_domain_validation_level" {
  default = "0"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings_exclude_credentials" {
  default = "true"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings_timeout" {
  default = "10000"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_api_hostname" {
  default = "apiHostname"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_attestation_key" {
  default = "attestationKey"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_integration_key" {
  default = "integrationKey"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_secret_key" {
  default = "secretKey"
}

variable "authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_user_mapping_attribute" {
  default = "primaryEmail"
}

variable "authentication_factor_setting_user_enrollment_disabled_factors" {
  default = []
}

variable "authentication_factor_setting_yubico_otp_enabled" {
  default = false
}


resource "oci_identity_domains_authentication_factor_setting" "test_authentication_factor_setting" {
  #Required
  authentication_factor_setting_id = "AuthenticationFactorSettings"
  bypass_code_enabled              = var.authentication_factor_setting_bypass_code_enabled
  bypass_code_settings {
    #Required
    help_desk_code_expiry_in_mins   = var.authentication_factor_setting_bypass_code_settings_help_desk_code_expiry_in_mins
    help_desk_generation_enabled    = var.authentication_factor_setting_bypass_code_settings_help_desk_generation_enabled
    help_desk_max_usage             = var.authentication_factor_setting_bypass_code_settings_help_desk_max_usage
    length                          = var.authentication_factor_setting_bypass_code_settings_length
    max_active                      = "6"
    self_service_generation_enabled = var.authentication_factor_setting_bypass_code_settings_self_service_generation_enabled
  }
  client_app_settings {
    #Required
    device_protection_policy            = "NONE"
    initial_lockout_period_in_secs      = "30"
    key_pair_length                     = "2048"
    lockout_escalation_pattern          = "Constant"
    max_failures_before_lockout         = "10"
    max_failures_before_warning         = "5"
    max_lockout_interval_in_secs        = "86400"
    min_pin_length                      = "6"
    policy_update_freq_in_days          = var.authentication_factor_setting_client_app_settings_policy_update_freq_in_days
    request_signing_algo                = var.authentication_factor_setting_client_app_settings_request_signing_algo
    shared_secret_encoding              = var.authentication_factor_setting_client_app_settings_shared_secret_encoding
    unlock_app_for_each_request_enabled = "false"
    unlock_app_interval_in_secs         = "300"
    unlock_on_app_foreground_enabled    = "false"
    unlock_on_app_start_enabled         = "false"
  }
  compliance_policy {
    #Required
    action = "Allow"
    name   = "lockScreenRequired"
    value  = "false"
  }
  compliance_policy {
    #Required
    action = "Allow"
    name   = "lockScreenRequiredUnknown"
    value  = "false"
  }
  compliance_policy {
    #Required
    action = "Allow"
    name   = "jailBrokenDevice"
    value  = "false"
  }
  compliance_policy {
    #Required
    action = "Allow"
    name   = "jailBrokenDeviceUnknown"
    value  = "false"
  }
  compliance_policy {
    #Required
    action = "Allow"
    name   = "minWindowsVersion"
    value  = "8.1"
  }
  compliance_policy {
    #Required
    action = "Allow"
    name   = "minIosVersion"
    value  = "7.1"
  }
  compliance_policy {
    #Required
    action = "Allow"
    name   = "minAndroidVersion"
    value  = "4.1"
  }
  compliance_policy {
    #Required
    action = "Allow"
    name   = "minIosAppVersion"
    value  = "4.0"
  }
  compliance_policy {
    #Required
    action = "Allow"
    name   = "minAndroidAppVersion"
    value  = "8.0"
  }
  compliance_policy {
    #Required
    action = "Allow"
    name   = "minWindowsAppVersion"
    value  = "1.0"
  }

  endpoint_restrictions {
    #Required
    max_endpoint_trust_duration_in_days = "180"
    max_enrolled_devices                = var.authentication_factor_setting_endpoint_restrictions_max_enrolled_devices
    max_incorrect_attempts              = "20"
    max_trusted_endpoints               = "20"
    trusted_endpoints_enabled           = var.authentication_factor_setting_endpoint_restrictions_trusted_endpoints_enabled
  }
  idcs_endpoint       = data.oci_identity_domain.test_domain.url
  mfa_enrollment_type = var.authentication_factor_setting_mfa_enrollment_type
  notification_settings {
    #Required
    pull_enabled = var.authentication_factor_setting_notification_settings_pull_enabled
  }
  push_enabled               = var.authentication_factor_setting_push_enabled
  schemas                    = ["urn:ietf:params:scim:schemas:oracle:idcs:AuthenticationFactorSettings"]
  security_questions_enabled = var.authentication_factor_setting_security_questions_enabled
  sms_enabled                = var.authentication_factor_setting_sms_enabled
  totp_enabled               = var.authentication_factor_setting_totp_enabled
  totp_settings {
    #Required
    email_otp_validity_duration_in_mins = var.authentication_factor_setting_totp_settings_email_otp_validity_duration_in_mins
    email_passcode_length               = "6"
    hashing_algorithm                   = var.authentication_factor_setting_totp_settings_hashing_algorithm
    jwt_validity_duration_in_secs       = "300"
    key_refresh_interval_in_days        = "60"
    passcode_length                     = "6"
    sms_otp_validity_duration_in_mins   = "6"
    sms_passcode_length                 = "6"
    time_step_in_secs                   = "30"
    time_step_tolerance                 = "3"
  }

  #Optional
  attribute_sets                    = []
  attributes                        = ""
  authorization                     = var.authentication_factor_setting_authorization
  auto_enroll_email_factor_disabled = var.authentication_factor_setting_auto_enroll_email_factor_disabled
  email_enabled                     = var.authentication_factor_setting_email_enabled
  email_settings {
    #Required
    email_link_enabled = var.authentication_factor_setting_email_settings_email_link_enabled

    #Optional
    email_link_custom_url = var.authentication_factor_setting_email_settings_email_link_custom_url
  }
  fido_authenticator_enabled = var.authentication_factor_setting_fido_authenticator_enabled
  hide_backup_factor_enabled = var.authentication_factor_setting_hide_backup_factor_enabled
  identity_store_settings {

    #Optional
    mobile_number_enabled        = var.authentication_factor_setting_identity_store_settings_mobile_number_enabled
    mobile_number_update_enabled = var.authentication_factor_setting_identity_store_settings_mobile_number_update_enabled
  }
  phone_call_enabled = var.authentication_factor_setting_phone_call_enabled
  # tags {
  #   #Required
  #   key   = var.authentication_factor_setting_tags_key
  #   value = var.authentication_factor_setting_tags_value
  # }
  third_party_factor {
    #Required
    duo_security = var.authentication_factor_setting_third_party_factor_duo_security
  }
  urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings {
    #Required
    attestation                                  = "NONE"
    authenticator_selection_attachment           = "BOTH"
    authenticator_selection_require_resident_key = "false"
    authenticator_selection_resident_key         = "NONE"
    authenticator_selection_user_verification    = var.authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings_authenticator_selection_user_verification
    exclude_credentials                          = "false"
    public_key_types                             = ["RS1"]
    timeout                                      = "60000"

    #Optional
    domain_validation_level = "1"
  }
  urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings {

    #Optional
    duo_security_settings {
      #Required
      api_hostname           = var.authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_api_hostname
      integration_key        = var.authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_integration_key
      secret_key             = var.authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_secret_key
      user_mapping_attribute = var.authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_user_mapping_attribute

      #Optional
      #this field is never returned
      # attestation_key = var.authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_attestation_key
    }
  }
  user_enrollment_disabled_factors = var.authentication_factor_setting_user_enrollment_disabled_factors
  yubico_otp_enabled               = var.authentication_factor_setting_yubico_otp_enabled
  lifecycle {
    ignore_changes = [schemas]
  }
}

data "oci_identity_domains_authentication_factor_settings" "test_authentication_factor_settings" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  attribute_sets = []
  attributes     = ""
  authorization  = var.authentication_factor_setting_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.authentication_factor_setting_resource_type_schema_version
}
