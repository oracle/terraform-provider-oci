// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "batch_job_pool_defined_tags_value" {
  default = "value"
}

variable "batch_job_pool_description" {
  default = "description"
}

variable "batch_job_pool_display_name" {
  default = "displayName"
}

variable "batch_job_pool_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "batch_job_pool_id" {
  default = "id"
}

variable "batch_job_pool_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_batch_batch_job_pool" "test_batch_job_pool" {
  #Required
  batch_context_id = oci_batch_batch_context.test_batch_context.id
  compartment_id   = var.compartment_id

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.batch_job_pool_defined_tags_value)
  description   = var.batch_job_pool_description
  display_name  = var.batch_job_pool_display_name
  freeform_tags = var.batch_job_pool_freeform_tags
}

data "oci_batch_batch_job_pools" "test_batch_job_pools" {

  #Optional
  batch_context_id = oci_batch_batch_context.test_batch_context.id
  compartment_id   = var.compartment_id
  display_name     = var.batch_job_pool_display_name
  id               = var.batch_job_pool_id
  state            = var.batch_job_pool_state
}

