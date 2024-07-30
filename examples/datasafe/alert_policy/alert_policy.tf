// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "alert_policy_access_level" {
  default = "RESTRICTED"
}

variable "alert_policy_alert_policy_rule_details_description" {
  default = "Check if remote login password file is exclusive and remote login is enabled "
}

variable "alert_policy_alert_policy_rule_details_display_name" {
  default = "Check remote login"
}

variable "alert_policy_alert_policy_rule_details_expression" {
  default = "operation eq \"abc\""
}

variable "alert_policy_alert_policy_type" {
  default = "AUDITING"
}

variable "alert_policy_compartment_id_in_subtree" {
  default = false
}

variable "alert_policy_defined_tags_value" {
  default = "value"
}

variable "alert_policy_description" {
  default = "Check if remote login password file is exclusive and remote login is enabled "
}

variable "alert_policy_display_name" {
  default = "Check remote login"
}

variable "alert_policy_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "alert_policy_is_user_defined" {
  default = true
}

variable "alert_policy_severity" {
  default = "CRITICAL"
}

variable "alert_policy_state" {
  default = "ACTIVE"
}

variable "alert_policy_time_created_less_than" {
  default = "2038-01-01T00:00:00.000Z"
}

variable "alert_policy_type" {
  default = "AUDITING"
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"

  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


resource "oci_data_safe_alert_policy" "test_alert_policy" {
  #Required
  alert_policy_type = var.alert_policy_alert_policy_type
  compartment_id    = var.compartment_ocid
  severity          = var.alert_policy_severity

  #Optional
  alert_policy_rule_details {
    #Required
    expression = var.alert_policy_alert_policy_rule_details_expression

    #Optional
    description  = var.alert_policy_alert_policy_rule_details_description
    display_name = var.alert_policy_alert_policy_rule_details_display_name
  }
  description   = var.alert_policy_description
  display_name  = var.alert_policy_display_name
  freeform_tags = var.alert_policy_freeform_tags
}

data "oci_data_safe_alert_policies" "test_alert_policies" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  access_level                          = var.alert_policy_access_level
  alert_policy_id                       = oci_data_safe_alert_policy.test_alert_policy.id
  compartment_id_in_subtree             = var.alert_policy_compartment_id_in_subtree
  is_user_defined                       = var.alert_policy_is_user_defined
  state                                 = var.alert_policy_state
  time_created_less_than                = var.alert_policy_time_created_less_than
}
