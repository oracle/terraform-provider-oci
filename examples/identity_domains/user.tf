// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "user_user_count" {
  default = 10
}

variable "user_user_filter" {
  default = ""
}

variable "user_active" {
  default = "true"
}

variable "user_addresses_country" {
  default = "us"
}

variable "user_addresses_formatted" {
  default = "formatted"
}

variable "user_addresses_locality" {
  default = "locality"
}

variable "user_addresses_postal_code" {
  default = "postalCode"
}

variable "user_addresses_primary" {
  default = false
}

variable "user_addresses_region" {
  default = "region"
}

variable "user_addresses_street_address" {
  default = "streetAddress"
}

variable "user_addresses_type" {
  default = "work"
}

variable "user_authorization" {
  default = "authorization"
}

variable "user_description" {
  default = "description"
}

variable "user_display_name" {
  default = "displayName"
}

variable "user_emails_primary" {
  default = true
}

variable "user_emails_secondary" {
  default = false
}

variable "user_emails_type" {
  default = "work"
}

variable "user_emails_value" {
  default = "value@email.com"
}

variable "user_emails_verified" {
  default = false
}

variable "user_entitlements_display" {
  default = "display"
}

variable "user_entitlements_primary" {
  default = false
}

variable "user_entitlements_type" {
  default = "type"
}

variable "user_entitlements_value" {
  default = "value"
}

variable "user_groups_date_added" {
  default = "dateAdded"
}

variable "user_groups_display" {
  default = "display"
}

variable "user_groups_membership_ocid" {
  default = "membershipOcid"
}

variable "user_groups_non_unique_display" {
  default = "nonUniqueDisplay"
}

variable "user_groups_type" {
  default = "direct"
}

variable "user_groups_value" {
  default = "value"
}

variable "user_ims_display" {
  default = "display"
}

variable "user_ims_primary" {
  default = false
}

variable "user_ims_type" {
  default = "aim"
}

variable "user_ims_value" {
  default = "value"
}

variable "user_locale" {
  default = "en"
}

variable "user_name_family_name" {
  default = "familyName"
}

variable "user_name_formatted" {
  default = "formatted"
}

variable "user_name_given_name" {
  default = "givenName"
}

variable "user_name_honorific_prefix" {
  default = "honorificPrefix"
}

variable "user_name_honorific_suffix" {
  default = "honorificSuffix"
}

variable "user_name_middle_name" {
  default = "middleName"
}

variable "user_nick_name" {
  default = "nickName"
}

variable "user_password" {
  default = "BEstrO0ng_#11"
}

variable "user_phone_numbers_display" {
  default = "display"
}

variable "user_phone_numbers_primary" {
  default = false
}

variable "user_phone_numbers_type" {
  default = "work"
}

variable "user_phone_numbers_value" {
  default = "1112223333"
}

variable "user_phone_numbers_verified" {
  default = false
}

variable "user_photos_display" {
  default = "display"
}

variable "user_photos_primary" {
  default = false
}

variable "user_photos_type" {
  default = "photo"
}

variable "user_photos_value" {
  default = "https://value.com"
}

variable "user_preferred_language" {
  default = "en"
}

variable "user_profile_url" {
  default = "https://profileUrl.com"
}

variable "user_roles_display" {
  default = "display"
}

variable "user_roles_primary" {
  default = false
}

variable "user_roles_type" {
  default = "type"
}

variable "user_roles_value" {
  default = "value"
}

variable "user_schemas" {
  default = []
}

variable "user_start_index" {
  default = 1
}

variable "user_tags_key" {
  default = "key"
}

variable "user_tags_value" {
  default = "value"
}

variable "user_timezone" {
  default = "America/Los_Angeles"
}

variable "user_title" {
  default = "title"
}

variable "user_urnietfparamsscimschemasextensionenterprise20user_cost_center" {
  default = "costCenter"
}

variable "user_urnietfparamsscimschemasextensionenterprise20user_department" {
  default = "department"
}

variable "user_urnietfparamsscimschemasextensionenterprise20user_division" {
  default = "division"
}

variable "user_urnietfparamsscimschemasextensionenterprise20user_employee_number" {
  default = "employeeNumber"
}

variable "user_urnietfparamsscimschemasextensionenterprise20user_manager_display_name" {
  default = "displayName"
}

variable "user_urnietfparamsscimschemasextensionenterprise20user_manager_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasextensionenterprise20user_organization" {
  default = "organization"
}

variable "user_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_key" {
  default = "key"
}

variable "user_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_namespace" {
  default = "namespace"
}

variable "user_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_key" {
  default = "freeformKey"
}

variable "user_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_value" {
  default = "freeformValue"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_level" {
  default = "LOW"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_last_update_timestamp" {
  default = "lastUpdateTimestamp"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_risk_level" {
  default = "LOW"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_score" {
  default = 10
}

variable "user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_source" {
  default = "source"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_status" {
  default = "status"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionkerberos_user_user_realm_users_principal_name" {
  default = "principalName"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionkerberos_user_user_realm_users_realm_name" {
  default = "realmName"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionkerberos_user_user_realm_users_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_bypass_codes_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_authentication_method" {
  default = "authenticationMethod"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_display" {
  default = "display"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_factor_status" {
  default = "factorStatus"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_factor_type" {
  default = "factorType"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_last_sync_time" {
  default = "lastSyncTime"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_status" {
  default = "status"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_third_party_vendor_name" {
  default = "thirdPartyVendorName"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_login_attempts" {
  default = 10
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_mfa_enabled_on" {
  default = "mfaEnabledOn"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_mfa_ignored_apps" {
  default = []
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_mfa_status" {
  default = "ENROLLED"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_preferred_authentication_factor" {
  default = "EMAIL"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_preferred_authentication_method" {
  default = "preferredAuthenticationMethod"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_preferred_device_display" {
  default = "display"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_preferred_device_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_preferred_third_party_vendor" {
  default = "preferredThirdPartyVendor"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_trusted_user_agents_display" {
  default = "display"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionmfa_user_trusted_user_agents_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpassword_state_user_applicable_password_policy_display" {
  default = "display"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpassword_state_user_applicable_password_policy_priority" {
  default = 10
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpassword_state_user_applicable_password_policy_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpassword_state_user_cant_change" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpassword_state_user_cant_expire" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpassword_state_user_expired" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpassword_state_user_last_failed_validation_date" {
  default = "lastFailedValidationDate"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpassword_state_user_last_successful_set_date" {
  default = "lastSuccessfulSetDate"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpassword_state_user_last_successful_validation_date" {
  default = "lastSuccessfulValidationDate"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpassword_state_user_must_change" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpasswordless_user_factor_identifier_display" {
  default = "display"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpasswordless_user_factor_identifier_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpasswordless_user_factor_method" {
  default = "factorMethod"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionpasswordless_user_factor_type" {
  default = "EMAIL"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionposix_user_gecos" {
  default = "gecos"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionposix_user_gid_number" {
  default = 10
}

variable "user_urnietfparamsscimschemasoracleidcsextensionposix_user_home_directory" {
  default = "homeDirectory"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionposix_user_login_shell" {
  default = "loginShell"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionposix_user_uid_number" {
  default = 500
}

variable "user_urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user_sec_questions_answer" {
  default = "answer"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user_sec_questions_hint_text" {
  default = "hintText"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user_sec_questions_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionself_registration_user_consent_granted" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionself_registration_user_self_registration_profile_display" {
  default = "display"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionself_registration_user_self_registration_profile_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionself_registration_user_user_token" {
  default = "userToken"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionsff_user_sff_auth_keys" {
  default = "sffAuthKeys"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionsocial_account_user_social_accounts_display" {
  default = "display"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionsocial_account_user_social_accounts_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionterms_of_use_user_terms_of_use_consents_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_last_failed_login_date" {
  default = "lastFailedLoginDate"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_last_successful_login_date" {
  default = "lastSuccessfulLoginDate"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_locked_expired" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_locked_lock_date" {
  default = "lockDate"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_locked_on" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_locked_reason" {
  default = 10
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_login_attempts" {
  default = 10
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_max_concurrent_sessions" {
  default = 10
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_previous_successful_login_date" {
  default = "previousSuccessfulLoginDate"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_recovery_attempts" {
  default = 10
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_recovery_enroll_attempts" {
  default = 10
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_recovery_locked_lock_date" {
  default = "lockDate"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_recovery_locked_on" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_user_provider" {
  default = "facebook"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_account_recovery_required" {
  default = true
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_bypass_notification" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_creation_mechanism" {
  default = "api"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_delegated_authentication_target_app_display" {
  default = "display"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_delegated_authentication_target_app_type" {
  default = "App"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_delegated_authentication_target_app_value" {
  default = "value"
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_do_not_show_getting_started" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_is_authentication_delegated" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_is_federated_user" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_is_group_membership_normalized" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_is_group_membership_synced_to_users_groups" {
  default = false
}

variable "user_urnietfparamsscimschemasoracleidcsextensionuser_user_user_flow_controlled_by_external_client" {
  default = false
}

variable "user_user_type" {
  default = "Contractor"
}

variable "user_x509certificates_display" {
  default = "display"
}

variable "user_x509certificates_primary" {
  default = false
}

variable "user_x509certificates_type" {
  default = "type"
}

variable "user_x509certificates_value" {
  default = ""
}

// CREATE/UPDATE a user
resource "oci_identity_domains_user" "test_user" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  name {
    #Required
    family_name = var.user_name_family_name

    #Optional
    formatted        = var.user_name_formatted
    given_name       = var.user_name_given_name
    honorific_prefix = var.user_name_honorific_prefix
    honorific_suffix = var.user_name_honorific_suffix
    middle_name      = var.user_name_middle_name
  }
  schemas = [
    #This is the required schema for User resource
    "urn:ietf:params:scim:schemas:core:2.0:User",
    /* #The following schemas are not required, but they will be returned because we are setting values on the corresponding attributes.
    #If don't want to list them here, you can use "lifecycle" attritbute to "ignore_changes" on them, to avoid update when apply.
    #Please see comments in the "lifecycle" block below.
    "urn:ietf:params:scim:schemas:oracle:idcs:extension:passwordState:User",
    "urn:ietf:params:scim:schemas:oracle:idcs:extension:userState:User",
    "urn:ietf:params:scim:schemas:oracle:idcs:extension:passwordless:User",
    "urn:ietf:params:scim:schemas:oracle:idcs:extension:OCITags",
    "urn:ietf:params:scim:schemas:oracle:idcs:extension:sff:User",
    "urn:ietf:params:scim:schemas:oracle:idcs:extension:capabilities:User",
    "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
    "urn:ietf:params:scim:schemas:oracle:idcs:extension:adaptive:User",
    "urn:ietf:params:scim:schemas:oracle:idcs:extension:user:User",
    "urn:ietf:params:scim:schemas:oracle:idcs:extension:dbCredentials:User",
    "urn:ietf:params:scim:schemas:oracle:idcs:extension:posix:User",
    */
  ]
  user_name = "userName"
  /* Note: In most cases, a primary email is REQUIRED to create a user. Otherwise you might get a 400 error. Please see "emails" block below. */

  #Optional
  active = var.user_active
  addresses {
    #Required
    type = var.user_addresses_type

    #Optional
    country        = var.user_addresses_country
    formatted      = var.user_addresses_formatted
    locality       = var.user_addresses_locality
    postal_code    = var.user_addresses_postal_code
    primary        = var.user_addresses_primary
    region         = var.user_addresses_region
    street_address = var.user_addresses_street_address
  }
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.user_authorization
  description    = var.user_description
  display_name   = var.user_display_name

  /* One and only one "emails" block needs to have "primary" set to true */
  emails {
    #Required
    type  = var.user_emails_type
    value = var.user_emails_value

    #Optional
    primary   = var.user_emails_primary
    secondary = var.user_emails_secondary
    verified  = var.user_emails_verified
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
    type  = var.user_entitlements_type
    value = var.user_entitlements_value

    #Optional
    display = var.user_entitlements_display
    primary = var.user_entitlements_primary
  }
  external_id = "externalId"
  ims {
    #Required
    type  = var.user_ims_type
    value = var.user_ims_value

    #Optional
    display = var.user_ims_display
    primary = var.user_ims_primary
  }
  locale    = var.user_locale
  nick_name = var.user_nick_name
  password  = var.user_password
  phone_numbers {
    #Required
    type  = var.user_phone_numbers_type
    value = var.user_phone_numbers_value

    #Optional
    primary = var.user_phone_numbers_primary
  }
  photos {
    #Required
    type  = var.user_photos_type
    value = var.user_photos_value

    #Optional
    display = var.user_photos_display
    primary = var.user_photos_primary
  }
  preferred_language = var.user_preferred_language
  profile_url        = var.user_profile_url
  #use the latest if not provided
  # resource_type_schema_version = var.user_resource_type_schema_version
  roles {
    #Required
    type  = var.user_roles_type
    value = var.user_roles_value

    #Optional
    display = var.user_roles_display
    primary = var.user_roles_primary
  }
  tags {
    #Required
    key   = var.user_tags_key
    value = var.user_tags_value
  }
  timezone = var.user_timezone
  title    = var.user_title
  urnietfparamsscimschemasextensionenterprise20user {

    #Optional
    cost_center     = var.user_urnietfparamsscimschemasextensionenterprise20user_cost_center
    department      = var.user_urnietfparamsscimschemasextensionenterprise20user_department
    division        = var.user_urnietfparamsscimschemasextensionenterprise20user_division
    employee_number = var.user_urnietfparamsscimschemasextensionenterprise20user_employee_number
    /* #provider manager user's id
    manager {

      #Optional
      value = var.user_urnietfparamsscimschemasextensionenterprise20user_manager_value
    }
    */
    organization = var.user_urnietfparamsscimschemasextensionenterprise20user_organization
  }
  urnietfparamsscimschemasoracleidcsextension_oci_tags {

    #Optional
    /* #create tagNamespace to use defined tags
    defined_tags {
      #Required
      key       = var.user_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_key
      namespace = var.user_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_namespace
      value     = var.user_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_value
    }
    */
    freeform_tags {
      #Required
      key   = var.user_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_key
      value = var.user_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_value
    }
  }
  urnietfparamsscimschemasoracleidcsextensionadaptive_user {

    #Optional
    risk_level = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_level
    /* set it to a valid RiskProviderProfile
    risk_scores {
      #Reuired
      last_update_timestamp = "2030-01-01T00:00:00Z"
      risk_level            = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_risk_level
      score                 = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_score
      value                 = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_value

      #Optional
      source = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_source
      status = var.user_urnietfparamsscimschemasoracleidcsextensionadaptive_user_risk_scores_status
    }
    */
  }
  urnietfparamsscimschemasoracleidcsextensioncapabilities_user {

    #Optional
    can_use_api_keys                 = true
    can_use_auth_tokens              = true
    can_use_console_password         = true
    can_use_customer_secret_keys     = true
    can_use_db_credentials           = true
    can_use_oauth2client_credentials = true
    can_use_smtp_credentials         = true
  }
  urnietfparamsscimschemasoracleidcsextensiondb_credentials_user {

    #Optional
    db_user_name = "dbUserName"
  }
  /* set value to valid Kerberos realm users id
  urnietfparamsscimschemasoracleidcsextensionkerberos_user_user {

    #Optional
    realm_users {
      #Required
      value = var.user_urnietfparamsscimschemasoracleidcsextensionkerberos_user_user_realm_users_value

      #Optional
      principal_name = var.user_urnietfparamsscimschemasoracleidcsextensionkerberos_user_user_realm_users_principal_name
      realm_name     = var.user_urnietfparamsscimschemasoracleidcsextensionkerberos_user_user_realm_users_realm_name
    }
  }
  */
  /* set mfa
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
      authentication_method   = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_authentication_method
      display                 = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_display
      factor_status           = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_factor_status
      factor_type             = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_factor_type
      last_sync_time          = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_last_sync_time
      status                  = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_status
      third_party_vendor_name = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_devices_third_party_vendor_name
    }
    login_attempts                  = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_login_attempts
    mfa_enabled_on                  = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_mfa_enabled_on
    mfa_ignored_apps                = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_mfa_ignored_apps
    mfa_status                      = var.user_urnietfparamsscimschemasoracleidcsextensionmfa_user_mfa_status
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
  */
  urnietfparamsscimschemasoracleidcsextensionpasswordless_user {

    #Optional
    /* set value to factor id
    factor_identifier {
      #Required
      value = var.user_urnietfparamsscimschemasoracleidcsextensionpasswordless_user_factor_identifier_value

      #Optional
      display = var.user_urnietfparamsscimschemasoracleidcsextensionpasswordless_user_factor_identifier_display
    }
    */
    factor_method = var.user_urnietfparamsscimschemasoracleidcsextensionpasswordless_user_factor_method
    factor_type   = var.user_urnietfparamsscimschemasoracleidcsextensionpasswordless_user_factor_type
  }
  urnietfparamsscimschemasoracleidcsextensionposix_user {

    #Optional
    gecos = var.user_urnietfparamsscimschemasoracleidcsextensionposix_user_gecos
    #provide gid of a group
    # gid_number     = var.user_urnietfparamsscimschemasoracleidcsextensionposix_user_gid_number
    home_directory = var.user_urnietfparamsscimschemasoracleidcsextensionposix_user_home_directory
    login_shell    = var.user_urnietfparamsscimschemasoracleidcsextensionposix_user_login_shell
    uid_number     = var.user_urnietfparamsscimschemasoracleidcsextensionposix_user_uid_number
  }
  /* set security questions
  urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user {

    #Optional
    sec_questions {
      #Required
      answer = var.user_urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user_sec_questions_answer
      value  = var.user_urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user_sec_questions_value

      #Optional
      hint_text = var.user_urnietfparamsscimschemasoracleidcsextensionsecurity_questions_user_sec_questions_hint_text
    }
  }
  */
  urnietfparamsscimschemasoracleidcsextensionself_change_user {

    #Optional
    allow_self_change = var.user_urnietfparamsscimschemasoracleidcsextensionself_change_user_allow_self_change
  }
  /* set value to SelfRegistration id
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
    user_token      = var.user_urnietfparamsscimschemasoracleidcsextensionself_registration_user_user_token
  }
  */
  urnietfparamsscimschemasoracleidcsextensionsff_user {

    #Optional
    sff_auth_keys = var.user_urnietfparamsscimschemasoracleidcsextensionsff_user_sff_auth_keys
  }
  /* set value to SocialAccount id
  urnietfparamsscimschemasoracleidcsextensionsocial_account_user {

    #Optional
    social_accounts {
      #Required
      value = var.user_urnietfparamsscimschemasoracleidcsextensionsocial_account_user_social_accounts_value

      #Optional
      display = var.user_urnietfparamsscimschemasoracleidcsextensionsocial_account_user_social_accounts_display
    }
  }
  */
  /* set value to TermsOfUse id
  urnietfparamsscimschemasoracleidcsextensionterms_of_use_user {

    #Optional
    terms_of_use_consents {
      #Required
      value = var.user_urnietfparamsscimschemasoracleidcsextensionterms_of_use_user_terms_of_use_consents_value
    }
  }
  */
  urnietfparamsscimschemasoracleidcsextensionuser_state_user {

    #Optional
    /* set to lock/unlock user
    locked {

      #Optional
      expired   = var.user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_locked_expired
      lock_date = "2020-01-01T00:00:00Z"
      on        = true
      reason    = 1
    }
    */
    max_concurrent_sessions = var.user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_max_concurrent_sessions
    /* set to lock/unlock password recovery
    recovery_locked {

      #Optional
      lock_date = "2030-01-01T00:00:00Z"
      on        = var.user_urnietfparamsscimschemasoracleidcsextensionuser_state_user_recovery_locked_on
    }
    */
  }
  urnietfparamsscimschemasoracleidcsextensionuser_user {

    #Optional
    user_provider             = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_user_provider
    account_recovery_required = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_account_recovery_required
    bypass_notification       = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_bypass_notification
    creation_mechanism        = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_creation_mechanism
    /*set value to target app id
    delegated_authentication_target_app {
      #Required
      type  = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_delegated_authentication_target_app_type
      value = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_delegated_authentication_target_app_value

      #Optional
      display = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_delegated_authentication_target_app_display
    }
    */
    do_not_show_getting_started                = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_do_not_show_getting_started
    is_authentication_delegated                = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_is_authentication_delegated
    is_federated_user                          = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_is_federated_user
    is_group_membership_normalized             = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_is_group_membership_normalized
    is_group_membership_synced_to_users_groups = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_is_group_membership_synced_to_users_groups
    /*set up email template id
    notification_email_template_id             = oci_identity_domains_notification_email_template.test_notification_email_template.id
    */
    user_flow_controlled_by_external_client = var.user_urnietfparamsscimschemasoracleidcsextensionuser_user_user_flow_controlled_by_external_client
  }
  user_type = var.user_user_type

  /* #set value with the certificate
  x509certificates {
    #Required
    value = var.user_x509certificates_value

    #Optional
    display = var.user_x509certificates_display
    primary = var.user_x509certificates_primary
    type    = var.user_x509certificates_type
  }
  */

  lifecycle {
    ignore_changes = [
      #adding "schemas" here to ignore any additional schemas returned
      schemas,
      #the field is never returned
      urnietfparamsscimschemasoracleidcsextensionself_change_user,
    ]
  }
}

// GET list of users
data "oci_identity_domains_users" "test_users" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  user_count     = var.user_user_count
  user_filter    = var.user_user_filter
  attribute_sets = []
  attributes     = ""
  authorization  = var.user_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.user_resource_type_schema_version
  start_index = var.user_start_index
}

