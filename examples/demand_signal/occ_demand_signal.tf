// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}

variable "occ_demand_signal_display_name"{
  default = "displayName"
}

variable "occ_demand_signal_id" {
  default = "occ_demand_signal_id"
}

variable "occ_demand_signal_resource_type" {
  default = "Compute - Std Intel"
}

variable "occ_demand_signal_units"  {
  default = "(Cores)"
}

variable "occ_demand_signal_time_expected" {
  default = "2025-01-05T00:00:00.000Z"
}

variable "occ_demand_signal_value" {
  default = 100
}

variable "occ_demand_signal_comments" {
  default = "comments"
}

variable "occ_demand_signal_is_active" {
  default = false
}

variable "occ_demand_signal_state" {
  default = "ACTIVE"
}

variable "occ_demand_signal_defined_tags_value" {
  default = "value"
}

variable "occ_demand_signal_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "occ_demand_signal_lifecycle_details" {
  default = "lifecycleDetails"
}

variable "occ_demand_signal_test_occ_demand_signals" {
  default = [" "]
}

variable "occ_demand_signal_test_value" {
  default = [
    {
      time_expected = "2025-01-05T00:00:00.000Z"
      value         = "100"
    },
    {
      time_expected = "2025-02-05T00:00:00.000Z"
      value         = "100"
    },
    {
      time_expected = "2025-03-05T00:00:00.000Z"
      value         = "100"
    },
    {
      time_expected = "2025-04-05T00:00:00.000Z"
      value         = "100"
    },
    {
      time_expected = "2025-05-05T00:00:00.000Z"
      value         = "100"
    },
    {
      time_expected = "2025-06-05T00:00:00.000Z"
      value         = "100"
    },
    {
      time_expected = "2025-07-05T00:00:00.000Z"
      value         = "100"
    },
    {
      time_expected = "2025-08-05T00:00:00.000Z"
      value         = "100"
    },
    {
      time_expected = "2025-09-05T00:00:00.000Z"
      value         = "100"
    },
    {
      time_expected = "2025-10-05T00:00:00.000Z"
      value         = "100"
    },
    {
      time_expected = "2025-11-05T00:00:00.000Z"
      value         = "100"
    },
    {
      time_expected = "2025-12-05T00:00:00.000Z"
      value         = "100"
    }
  ]

}
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_demand_signal_occ_demand_signal" "test_occ_demand_signal" {
  #Required
  compartment_id    = var.compartment_id
  is_active         = var.occ_demand_signal_is_active
  occ_demand_signals {
    resource_type   = var.occ_demand_signal_resource_type
    units           = var.occ_demand_signal_units
    values {
      time_expected = var.occ_demand_signal_time_expected
      value         = var.occ_demand_signal_value
      #Optional
      comments      = var.occ_demand_signal_comments
    }
  }
  #Optional
  display_name      = var.occ_demand_signal_display_name
  id                = var.occ_demand_signal_id
  state             = var.occ_demand_signal_state
  freeform_tags     = var.occ_demand_signal_freeform_tags
  lifecycle_details = var.occ_demand_signal_lifecycle_details
  patch_operations {
    #Required
    from            = var.occ_demand_signal_patch_operations_from
    operation       = var.occ_demand_signal_patch_operations_operation
    value           = var.occ_demand_signal_patch_operations_value
    selection       = var.occ_demand_signal_patch_operations_selection
    #Optional
    position        = var.occ_demand_signal_patch_operations_position
    selected_item   = var.occ_demand_signal_patch_operations_selected_item

  }
}

data "oci_demand_signal_occ_demand_signals" "test_occ_demand_signals" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.occ_demand_signal_display_name
  id           = var.occ_demand_signal_id
  state       = var.occ_demand_signal_state
}