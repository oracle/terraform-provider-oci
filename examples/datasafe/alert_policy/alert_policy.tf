// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_alert_policy_ocid" {}

variable "alert_policy_is_user_defined" {
  default = false
}

variable "alert_policy_type" {
  default = "AUDITING"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_alert_policies" "test_alert_policies" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  alert_policy_id = var.data_safe_alert_policy_ocid
  is_user_defined = var.alert_policy_is_user_defined
  type = var.alert_policy_type
}

