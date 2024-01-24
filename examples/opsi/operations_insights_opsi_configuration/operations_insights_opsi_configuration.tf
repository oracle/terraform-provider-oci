// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "opsi_configuration_config_item_custom_status" {
  default = ["customized"]
}

variable "opsi_configuration_config_item_field" {
  default = ["metadata", "name", "value", "defaultValue"]
}

variable "opsi_effective_configuration_config_item_field" {
  default = ["metadata", "name", "value", "defaultValue", "valueSourceConfig"]
}

variable "opsi_configuration_config_items_config_item_type" {
  default = "BASIC"
}

variable "opsi_configuration_config_items_name1" {
  default = "dbHighCpuThreshold"
}

variable "opsi_configuration_config_items_name2" {
  default = "dbHighMemoryThreshold"
}

variable "opsi_configuration_config_items_value1" {
  default = "83"
}

variable "opsi_configuration_config_items_value2" {
  default = "72"
}

variable "opsi_configuration_config_items_applicable_context" {
  default = ["DB_CAPACITY_PLANNING"]
}

variable "opsi_configuration_defined_tags_value" {
  default = "value"
}

variable "opsi_configuration_description" {
  default = "description"
}

variable "opsi_configuration_display_name" {
  default = "displayName"
}

variable "opsi_configuration_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "opsi_configuration_opsi_config_field" {
  default = ["configItems"]
}

variable "opsi_configuration_opsi_config_type_arr" {
  default = ["UX_CONFIGURATION"]
}

variable "opsi_configuration_opsi_config_type" {
  default = "UX_CONFIGURATION"
}

variable "opsi_configuration_state" {
  default = ["ACTIVE"]
}

variable "opsi_configuration_system_tags" {
  default = "value"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

// Create custom Opsi configuration
resource "oci_opsi_opsi_configuration" "test_opsi_configuration" {
  opsi_config_type          = var.opsi_configuration_opsi_config_type
  compartment_id            = var.compartment_ocid
  config_item_custom_status = var.opsi_configuration_config_item_custom_status
  config_item_field         = var.opsi_configuration_config_item_field
  config_items {
    config_item_type = var.opsi_configuration_config_items_config_item_type
    name             = var.opsi_configuration_config_items_name1
    value            = var.opsi_configuration_config_items_value1
  }
  config_items {
    config_item_type = var.opsi_configuration_config_items_config_item_type
    name             = var.opsi_configuration_config_items_name2
    value            = var.opsi_configuration_config_items_value2
  }
  config_items_applicable_context = var.opsi_configuration_config_items_applicable_context
  description                     = var.opsi_configuration_description
  display_name                    = var.opsi_configuration_display_name
  opsi_config_field               = var.opsi_configuration_opsi_config_field
}

// List opsi configuration in compartment
data "oci_opsi_opsi_configurations" "test_opsi_configurations" {
  compartment_id   = var.compartment_ocid
  display_name     = var.opsi_configuration_display_name
  opsi_config_type = var.opsi_configuration_opsi_config_type_arr
  state            = var.opsi_configuration_state
}

// Get opsi configuration by id
data "oci_opsi_opsi_configuration" "test_opsi_configuration" {
  opsi_configuration_id           = oci_opsi_opsi_configuration.test_opsi_configuration.id
  config_item_custom_status       = var.opsi_configuration_config_item_custom_status
  config_item_field               = var.opsi_configuration_config_item_field
  opsi_config_field               = var.opsi_configuration_opsi_config_field
  config_items_applicable_context = var.opsi_configuration_config_items_applicable_context
}

// Get effective opsi configuration for a compartment
data "oci_opsi_opsi_configuration_configuration_item" "test_opsi_configuration_configuration_item" {
  compartment_id                  = var.compartment_ocid
  config_item_field               = var.opsi_effective_configuration_config_item_field
  config_items_applicable_context = var.opsi_configuration_config_items_applicable_context
  opsi_config_type                = var.opsi_configuration_opsi_config_type
}

