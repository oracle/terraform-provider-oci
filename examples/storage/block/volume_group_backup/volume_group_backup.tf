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
  // Define the region the destination volume group backup copy will be created in
}

variable "source_region" {
  // Define the region the source volume group backup will be created in
}

variable "volume_group_id" {
}

variable "source_volume_group_backup_id" {
  // Define a source volume group backup Id to copy when using resource oci_core_volume_group_backup.test_volume_group_backup_cross_region_sourced
}

variable "volume_group_backup_defined_tags_value" {
  default = "value"
}

variable "volume_group_backup_display_name" {
  default = "displayName"
}

variable "volume_group_backup_copy_display_name" {
  default = "displayNameCopy"
}

variable "volume_group_backup_state" {
  default = "AVAILABLE"
}

variable "volume_group_backup_freeform_tags" {
  type = map(string)

  default = {
    Department = "Finance"
  }
}

variable "volume_group_backup_type" {
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

resource "oci_core_volume_group" "test_volume_group" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id = var.compartment_ocid

  source_details {
    #Required
    type = "volumeIds"

    // Mix of named volume and splatted multiple volumes
    volume_ids = concat([oci_core_volume.test_volume.id], oci_core_volume.test_volume.*.id)
  }

  #Optional
  display_name = "test-volume-group-from-vol-ids"
}

resource "oci_core_volume_group_backup" "test_volume_group_backup" {
  #Required
  volume_group_id = oci_core_volume_group.test_volume_group.id

  #Optional
  display_name  = var.volume_group_backup_display_name
  freeform_tags = var.volume_group_backup_freeform_tags
  type          = var.volume_group_backup_type
}

resource "oci_core_volume_group_backup" "test_volume_group_backup_cross_region_sourced" {
  #Required
  source_details {
    region           = var.source_region
    volume_group_backup_id = var.source_volume_group_backup_id
  }

  #Optional
  display_name = var.volume_group_backup_copy_display_name
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_core_volume_group_backups" "test_volume_group_backup" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.volume_group_backup_display_name
  volume_group_id    = oci_core_volume_group.test_volume_group.id
}

output "volumeGroupBackup" {
  value = data.oci_core_volume_group_backups.test_volume_group_backup.display_name
}