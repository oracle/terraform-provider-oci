// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "fleet_id" {
  default = "id"
}

variable "fleet_display_name" {
  default = "Example Fleet"
}

variable "fleet_display_name_contains" {
  default = "Example Fleet"
}

variable "fleet_description" {
  default = "Example Fleet created by Terraform"
}

variable "fleet_is_advanced_features_enabled" {
  default = false
}

variable "fleet_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "fleet_defined_tags" {
  default  = { "example-tag-namespace-all.example-tag" = "value" }
}

variable "fleet_state" {
  default = "ACTIVE"
}

variable "fleet_log_group_ocid" {
  default = "example-log-group-id"
}

variable "fleet_inventory_log_ocid" {
  default = "example-inventory-log-id"
}

variable "fleet_operation_log_ocid" {
  default = "example-operation-log-id"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_jms_fleet" "example_fleet" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = var.fleet_display_name
  inventory_log {
    log_group_id = var.fleet_log_group_ocid
    log_id       = var.fleet_inventory_log_ocid
  }

  #Optional
  description                  = var.fleet_description
  freeform_tags                = var.fleet_freeform_tags
  operation_log {
    log_group_id = var.fleet_log_group_ocid
    log_id       = var.fleet_operation_log_ocid
  }

  # Create the Tag namespace in OCI before enabling
  # See user guide: https://docs.oracle.com/en-us/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm
  # defined_tags  = var.fleet_defined_tags
}

data "oci_jms_fleets" "example_fleets" {

  #Optional
  compartment_id        = var.compartment_ocid
  display_name          = var.fleet_display_name
  display_name_contains = var.fleet_display_name_contains
  id                    = var.fleet_id
  state                 = var.fleet_state
}

