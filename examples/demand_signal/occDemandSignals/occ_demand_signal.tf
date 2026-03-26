// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "occ_demand_signal_defined_tags_value" {
  default = "value"
}

variable "occ_demand_signal_display_name" {
  default = "string"
}

variable "occ_demand_signal_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "occ_demand_signal_id" {
  default = "id"
}

variable "occ_demand_signal_is_active" {
  default = false
}

variable "occ_demand_signal_occ_demand_signals_resource_type" {
  default = "string"
}

variable "occ_demand_signal_occ_demand_signals_units" {
  default = "string"
}

variable "occ_demand_signal_occ_demand_signals_values_comments" {
  default = "string"
}

variable "occ_demand_signal_occ_demand_signals_values_time_expected" {
  default = "timeExpected"
}

variable "occ_demand_signal_occ_demand_signals_values_value" {
  default = 1.0
}

variable "occ_demand_signal_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_demand_signal_occ_demand_signal" "test_occ_demand_signal" {
  #Required
  compartment_id       = var.compartment_id
  is_active            = var.occ_demand_signal_is_active
  occ_demand_signal_id = var.occ_demand_signal_occ_demand_signal_id
  occ_demand_signals {
    #Required
    resource_type = var.occ_demand_signal_occ_demand_signals_resource_type
    units         = var.occ_demand_signal_occ_demand_signals_units
    values {
      #Required
      time_expected = var.occ_demand_signal_occ_demand_signals_values_time_expected
      value         = var.occ_demand_signal_occ_demand_signals_values_value

      #Optional
      comments = var.occ_demand_signal_occ_demand_signals_values_comments
    }
  }

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.occ_demand_signal_defined_tags_value)
  display_name  = var.occ_demand_signal_display_name
  freeform_tags = var.occ_demand_signal_freeform_tags
  patch_operations {
    #Required
    operation = var.occ_demand_signal_patch_operations_operation
    selection = var.occ_demand_signal_patch_operations_selection

    #Optional
    from          = var.occ_demand_signal_patch_operations_from
    position      = var.occ_demand_signal_patch_operations_position
    selected_item = var.occ_demand_signal_patch_operations_selected_item
    value         = var.occ_demand_signal_patch_operations_value
    values        = var.occ_demand_signal_patch_operations_values
  }
}

data "oci_demand_signal_occ_demand_signals" "test_occ_demand_signals" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.occ_demand_signal_display_name
  id             = var.occ_demand_signal_id
  state          = var.occ_demand_signal_state
}

