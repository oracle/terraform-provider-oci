// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "batch_task_profile_defined_tags_value" {
  default = "value"
}

variable "batch_task_profile_description" {
  default = "description"
}

variable "batch_task_profile_display_name" {
  default = "displayName"
}

variable "batch_task_profile_extended_information_architecture" {
  default = "GENERIC_X86"
}

variable "batch_task_profile_extended_information_type" {
  default = "CPU_ARCHITECTURE_TASK_PROFILE_EXTENDED_INFORMATION"
}

variable "batch_task_profile_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "batch_task_profile_id" {
  default = "id"
}

variable "batch_task_profile_min_disk_size_in_gbs" {
  default = 10
}

variable "batch_task_profile_min_memory_in_gbs" {
  default = 10
}

variable "batch_task_profile_min_ocpus" {
  default = 10
}

variable "batch_task_profile_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_batch_batch_task_profile" "test_batch_task_profile" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  defined_tags = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.batch_task_profile_defined_tags_value)
  description  = var.batch_task_profile_description
  display_name = var.batch_task_profile_display_name
  extended_information {
    #Required
    type = var.batch_task_profile_extended_information_type

    #Optional
    architecture = var.batch_task_profile_extended_information_architecture
    shape_name   = oci_core_shape.test_shape.name
  }
  freeform_tags        = var.batch_task_profile_freeform_tags
  min_disk_size_in_gbs = var.batch_task_profile_min_disk_size_in_gbs
  min_memory_in_gbs    = var.batch_task_profile_min_memory_in_gbs
  min_ocpus            = var.batch_task_profile_min_ocpus
}

data "oci_batch_batch_task_profiles" "test_batch_task_profiles" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.batch_task_profile_display_name
  id             = var.batch_task_profile_id
  state          = var.batch_task_profile_state
}

