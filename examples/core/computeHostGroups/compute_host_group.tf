// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {}
variable "config_file_profile" {}
variable "compartment_ocid" {}
variable "availability_domain" {}

provider "oci" {
  region              = var.region
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
}

resource "oci_core_compute_host_group" "quick_recycle_example" {
  availability_domain            = var.availability_domain
  compartment_id                 = var.compartment_ocid
  display_name                   = "terraform-quick-recycle-example"
  is_targeted_placement_required = false

  configurations {
    recycle_level = "SKIP_RECYCLE"
    target        = "BM.DenseIO.E4.128"

    quick_recycle_settings {
      nvme_wipe = true
    }
  }
}

data "oci_core_compute_host_group" "quick_recycle_example" {
  compute_host_group_id = oci_core_compute_host_group.quick_recycle_example.id
}
