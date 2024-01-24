// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "password_policy_password_policy_count" {
  default = 10
}

variable "password_policy_password_policy_filter" {
  default = ""
}

variable "password_policy_allowed_chars" {
  default = "allowedChars"
}

variable "password_policy_authorization" {
  default = "authorization"
}

variable "password_policy_configured_password_policy_rules_key" {
  default = "key"
}

variable "password_policy_configured_password_policy_rules_value" {
  default = "value"
}

variable "password_policy_delete_in_progress" {
  default = false
}

variable "password_policy_description" {
  default = "description"
}

variable "password_policy_dictionary_delimiter" {
  default = "dictionaryDelimiter"
}

variable "password_policy_dictionary_location" {
  default = "dictionaryLocation"
}

variable "password_policy_dictionary_word_disallowed" {
  default = false
}

variable "password_policy_disallowed_chars" {
  default = "a,b,c"
}

variable "password_policy_disallowed_substrings" {
  default = []
}

variable "password_policy_first_name_disallowed" {
  default = false
}

variable "password_policy_force_password_reset" {
  default = false
}

variable "password_policy_groups_display" {
  default = "display"
}

variable "password_policy_groups_ref" {
  default = "ref"
}

variable "password_policy_groups_value" {
  default = "value"
}

variable "password_policy_last_name_disallowed" {
  default = false
}

variable "password_policy_lockout_duration" {
  default = 10
}

variable "password_policy_max_incorrect_attempts" {
  default = 10
}

variable "password_policy_max_length" {
  default = 10
}

variable "password_policy_max_repeated_chars" {
  default = 100
}

variable "password_policy_max_special_chars" {
  default = 10
}

variable "password_policy_min_alpha_numerals" {
  default = 1
}

variable "password_policy_min_alphas" {
  default = 1
}

variable "password_policy_min_length" {
  default = 1
}

variable "password_policy_min_lower_case" {
  default = 1
}

variable "password_policy_min_numerals" {
  default = 1
}

variable "password_policy_min_password_age" {
  default = 10
}

variable "password_policy_min_special_chars" {
  default = 1
}

variable "password_policy_min_unique_chars" {
  default = 1
}

variable "password_policy_min_upper_case" {
  default = 1
}

variable "password_policy_name" {
  default = "name"
}

variable "password_policy_num_passwords_in_history" {
  default = 10
}

variable "password_policy_password_expire_warning" {
  default = 10
}

variable "password_policy_password_expires_after" {
  default = 10
}

variable "password_policy_password_strength" {
  default = "Custom"
}

variable "password_policy_priority" {
  default = 10
}

variable "password_policy_required_chars" {
  default = "x,y,z"
}

variable "password_policy_start_index" {
  default = 1
}

variable "password_policy_starts_with_alphabet" {
  default = false
}

variable "password_policy_tags_key" {
  default = "key"
}

variable "password_policy_tags_value" {
  default = "value"
}

variable "password_policy_user_name_disallowed" {
  default = false
}


resource "oci_identity_domains_password_policy" "test_password_policy" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  name          = var.password_policy_name
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:PasswordPolicy"]

  #Optional
  allowed_chars              = var.password_policy_allowed_chars
  attribute_sets             = ["all"]
  attributes                 = ""
  authorization              = var.password_policy_authorization
  description                = var.password_policy_description
  dictionary_delimiter       = var.password_policy_dictionary_delimiter
  dictionary_location        = var.password_policy_dictionary_location
  dictionary_word_disallowed = var.password_policy_dictionary_word_disallowed
  disallowed_chars           = var.password_policy_disallowed_chars
  disallowed_substrings      = var.password_policy_disallowed_substrings
  external_id                = "externalId"
  first_name_disallowed      = var.password_policy_first_name_disallowed
  force_password_reset       = var.password_policy_force_password_reset
  /* #provide group's id
  groups {
    #Required
    value = oci_identity_domains_group.test_group.id
  }
  */
  last_name_disallowed     = var.password_policy_last_name_disallowed
  lockout_duration         = var.password_policy_lockout_duration
  max_incorrect_attempts   = var.password_policy_max_incorrect_attempts
  max_length               = var.password_policy_max_length
  max_repeated_chars       = var.password_policy_max_repeated_chars
  max_special_chars        = var.password_policy_max_special_chars
  min_alpha_numerals       = var.password_policy_min_alpha_numerals
  min_alphas               = var.password_policy_min_alphas
  min_length               = var.password_policy_min_length
  min_lower_case           = var.password_policy_min_lower_case
  min_numerals             = var.password_policy_min_numerals
  min_password_age         = var.password_policy_min_password_age
  min_special_chars        = var.password_policy_min_special_chars
  min_unique_chars         = var.password_policy_min_unique_chars
  min_upper_case           = var.password_policy_min_upper_case
  num_passwords_in_history = var.password_policy_num_passwords_in_history
  password_expire_warning  = var.password_policy_password_expire_warning
  password_expires_after   = var.password_policy_password_expires_after
  password_strength        = var.password_policy_password_strength
  priority                 = var.password_policy_priority
  required_chars           = var.password_policy_required_chars
  #use the latest if not provided
  # resource_type_schema_version = var.password_policy_resource_type_schema_version
  starts_with_alphabet = var.password_policy_starts_with_alphabet
  tags {
    #Required
    key   = var.password_policy_tags_key
    value = var.password_policy_tags_value
  }
  user_name_disallowed = var.password_policy_user_name_disallowed
}

data "oci_identity_domains_password_policies" "test_password_policies" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  password_policy_count  = var.password_policy_password_policy_count
  password_policy_filter = var.password_policy_password_policy_filter
  attribute_sets         = []
  attributes             = ""
  authorization          = var.password_policy_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.password_policy_resource_type_schema_version
  start_index = var.password_policy_start_index
}

