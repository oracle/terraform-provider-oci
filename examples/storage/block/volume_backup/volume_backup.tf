// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "compartment_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
  // Define the region the destination volume backup copy will be created in

}

variable "source_region" {
  // Define the region the source volume backup will be created in
}

variable "source_volume_backup_id" {
  // Define a source volume backup Id to copy when using resource oci_core_volume_backup.test_volume_backup_cross_region_sourced
}

variable "volume_backup_defined_tags_value" {
  default = "value"
}

variable "volume_backup_display_name" {
  default = "displayName"
}

variable "volume_backup_copy_display_name" {
  default = "displayNameCopy"
}

variable "volume_backup_state" {
  default = "AVAILABLE"
}

variable "volume_backup_freeform_tags" {
  type = map(string)

  default = {
    Department = "Finance"
  }
}

variable "volume_backup_type" {
  default = "FULL"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_volume" "test_volume" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "-tf-volume"
}

resource "oci_core_volume_backup" "test_volume_backup" {
  #Required
  volume_id = oci_core_volume.test_volume.id

  #Optional
  display_name  = var.volume_backup_display_name
  freeform_tags = var.volume_backup_freeform_tags
  type          = var.volume_backup_type
}

resource "oci_core_volume_backup" "test_volume_backup_cross_region_sourced" {
  #Required
  source_details {
    region           = var.source_region
    volume_backup_id = var.source_volume_backup_id
  }

  #Optional
  display_name = var.volume_backup_copy_display_name
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_core_volume_backups" "test_volume_backup" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.volume_backup_display_name
  state        = var.volume_backup_state
  volume_id    = oci_core_volume.test_volume.id
}

data "oci_core_volume_backups" "test_volume_backup_from_source" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name            = var.volume_backup_copy_display_name
  source_volume_backup_id = var.source_volume_backup_id
}

output "volumeBackup" {
  value = data.oci_core_volume_backups.test_volume_backup.display_name
}

output "volumeBackupIdFromSource" {
  value = data.oci_core_volume_backups.test_volume_backup_from_source.source_volume_backup_id
}

