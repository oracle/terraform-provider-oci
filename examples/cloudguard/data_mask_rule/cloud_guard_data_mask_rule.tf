// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "data_mask_rule_access_level" {
  default = "ACCESSIBLE"
}

variable "data_mask_rule_data_mask_categories" {
  default = ["PII"]
}

variable "data_mask_rule_data_mask_rule_status" {
  default = "ENABLED"
}

variable "data_mask_rule_defined_tags_value" {
  default = "value"
}

variable "data_mask_rule_description" {
  default = "description"
}

variable "data_mask_rule_display_name" {
  default = "displayName"
}

variable "data_mask_rule_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "data_mask_rule_state" {
  default = "ACTIVE"
}

variable "data_mask_rule_target_selected_kind" {
  default = "ALL"
}

variable "data_mask_rule_target_selected_values" {
  default = []
}

variable "data_mask_rule_target_type" {
  default = "COMPARTMENT"
}

variable "data_mask_rule_target_id" {
  default = "ocid.target.test1"
}

variable "data_mask_rule_iam_group_id" {
  default = "ocid.group.test1"
}



provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

resource "oci_cloud_guard_data_mask_rule" "test_data_mask_rule" {
  #Required
  compartment_id       = "${var.tenancy_ocid}"
  data_mask_categories = "${var.data_mask_rule_data_mask_categories}"
  display_name         = "${var.data_mask_rule_display_name}"
  iam_group_id         = "${var.data_mask_rule_iam_group_id}"
  target_selected {
    #Required
    kind = "${var.data_mask_rule_target_selected_kind}"

    #Optional
    values = "${var.data_mask_rule_target_selected_values}"
  }

  #Optional
  data_mask_rule_status = "${var.data_mask_rule_data_mask_rule_status}"
  description           = "${var.data_mask_rule_description}"
  state                 = "${var.data_mask_rule_state}"
}

data "oci_cloud_guard_data_mask_rules" "test_data_mask_rules" {
  #Required
  compartment_id            = "${var.tenancy_ocid}"

  #Optional
  access_level          = "${var.data_mask_rule_access_level}"
  data_mask_rule_status = "${var.data_mask_rule_data_mask_rule_status}"
  display_name          = "${var.data_mask_rule_display_name}"
  iam_group_id          = "${var.data_mask_rule_iam_group_id}"
  state                 = "${var.data_mask_rule_state}"
  target_id             = "${var.data_mask_rule_target_id}"
  target_type           = "${var.data_mask_rule_target_type}"
}

