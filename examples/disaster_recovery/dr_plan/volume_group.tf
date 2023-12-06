// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "volume_group_defined_tags_value" {
  default = "value"
}

variable "volume_group_display_name" {
  default = "example-volume-group"
}

variable "volume_group_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "volume_group_source_details_type" {
  default = "volumeIds"
}

variable "volume_group_source_details_volume_ids" {
  default = []
}

variable "volume_group_state" {
  default = "AVAILABLE"
}

variable "volume_group_volume_group_replicas_availability_domain" {
  default = "availabilityDomain"
}

variable "volume_group_volume_group_replicas_display_name" {
  default = "displayName"
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

data "oci_core_volume_groups" "test_volume_groups" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  display_name        = var.volume_group_display_name
  state               = var.volume_group_state
}