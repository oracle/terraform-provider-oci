// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "batch_context_shape_availability_domain" {
  default = "availabilityDomain"
}

variable "batch_context_shape_shape_type" {
  default = "CPU"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_batch_batch_context_shapes" "test_batch_context_shapes" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  availability_domain = var.batch_context_shape_availability_domain
  shape_type          = var.batch_context_shape_shape_type
}

