// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "byol_available_units_greater_than_or_equal_to" {
  default = 1.0
}

variable "byol_defined_tags_value" {
  default = "value"
}

variable "byol_description" {
  default = "description"
}

variable "byol_display_name" {
  default = "displayName"
}

variable "byol_entitlement_key" {
  default = "AAAAA-BBBBB-CCCCC-DDDDD-EEEEE"
}

variable "byol_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "byol_software_type" {
  default = "VCF"
}

variable "byol_state" {
  default = "ACTIVE"
}

variable "byol_time_term_end" {
  default = "2029-03-23T01:23:45.678Z"
}

variable "byol_time_term_start" {
  default = "2025-03-23T01:23:45.678Z"
}

variable "byol_total_units" {
  default = 10
}

variable "byol_allocation_allocated_units" {
  default = 10
}

variable "byol_allocation_available_units_greater_than_or_equal_to" {
  default = 1.0
}

variable "byol_allocation_defined_tags_value" {
  default = "value"
}

variable "byol_allocation_display_name" {
  default = "displayName"
}

variable "byol_allocation_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "byol_allocation_software_type" {
  default = "VCF"
}

variable "byol_allocation_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_ocvp_byol" "test_byol" {
  #Required
  compartment_id  = var.compartment_id
  display_name    = var.byol_display_name
  entitlement_key = var.byol_entitlement_key
  software_type   = var.byol_software_type
  time_term_end   = var.byol_time_term_end
  time_term_start = var.byol_time_term_start
  total_units     = var.byol_total_units

  #Optional
  description   = var.byol_description
  #defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.byol_defined_tags_value)
  #freeform_tags = var.byol_freeform_tags
}

data "oci_ocvp_byols" "test_byols" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  available_units_greater_than_or_equal_to = var.byol_available_units_greater_than_or_equal_to
  byol_id                                  = oci_ocvp_byol.test_byol.id
  display_name                             = var.byol_display_name
  software_type                            = var.byol_software_type
  state                                    = var.byol_state
}

resource "oci_ocvp_byol_allocation" "test_byol_allocation" {
  #Required
  allocated_units = var.byol_allocation_allocated_units
  byol_id         = oci_ocvp_byol.test_byol.id
  compartment_id  = var.compartment_id
  display_name    = var.byol_allocation_display_name

  #Optional
  #defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.byol_allocation_defined_tags_value)
  #freeform_tags = var.byol_allocation_freeform_tags
}

data "oci_ocvp_byol_allocations" "test_byol_allocations" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  available_units_greater_than_or_equal_to = var.byol_allocation_available_units_greater_than_or_equal_to
  byol_allocation_id                       = oci_ocvp_byol_allocation.test_byol_allocation.id
  byol_id                                  = oci_ocvp_byol.test_byol.id
  display_name                             = var.byol_allocation_display_name
  software_type                            = var.byol_allocation_software_type
  state                                    = var.byol_allocation_state
}

