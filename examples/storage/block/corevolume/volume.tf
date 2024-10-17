// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
variable "tenancy_ocid" {
}

variable "auth" {
}

variable "config_file_profile" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "destination_region" {
}

variable "vol_first_backup_ocid" {
}

variable "vol_second_backup_ocid" {
}

provider "oci" {
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_volume" "test_create_volume" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "tf-volume-1"
  size_in_gbs         = "50"
}

resource "oci_core_volume" "test_create_delta_restored_volume" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "tf-delta-restored-volume-1"
  size_in_gbs         = "50"
  source_details {
    first_backup_id   = var.vol_first_backup_ocid
    second_backup_id = var.vol_second_backup_ocid
    change_block_size_in_bytes = 4096
    type = "volumeBackupDelta"
  }
}