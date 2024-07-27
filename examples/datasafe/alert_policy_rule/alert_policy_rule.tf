// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "policy_id" {}

variable "alert_policy_rule_description" {
  default = "Check if remote login password file is exclusive"
}

variable "alert_policy_rule_display_name" {
  default = "displayName"
}

variable "alert_policy_rule_expression" {
  default = "operation eq \"GRANT\""
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

resource "oci_data_safe_alert_policy_rule" "test_alert_policy_rule" {
  #Required
  alert_policy_id = var.policy_id
  expression      = var.alert_policy_rule_expression

  #Optional
  description  = var.alert_policy_rule_description
  display_name = var.alert_policy_rule_display_name
}

data "oci_data_safe_alert_policy_rules" "test_alert_policy_rules" {
  #Required
  alert_policy_id = var.policy_id
}
