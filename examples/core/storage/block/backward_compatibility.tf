// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "instance_shape" {
  default = "VM.Standard.E4.Flex"
}

variable "instance_ocpus" {
  default = 1
}

variable "instance_shape_config_memory_in_gbs" {
  default = 6
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_identity_availability_domain" "backcompat" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_core_images" "backcompat" {
  compartment_id           = var.compartment_ocid
  operating_system         = "Oracle Linux"
  operating_system_version = "8"
  shape                    = var.instance_shape
  sort_by                  = "TIMECREATED"
  sort_order               = "DESC"
}

resource "oci_core_vcn" "backcompat" {
  compartment_id = var.compartment_ocid
  cidr_block     = "10.10.0.0/16"
  display_name   = "tf-backcompat-block-vcn"
}

resource "oci_core_subnet" "backcompat" {
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.backcompat.id
  cidr_block          = "10.10.1.0/24"
  availability_domain = data.oci_identity_availability_domain.backcompat.name
  display_name        = "tf-backcompat-block-subnet"
}

resource "oci_core_instance" "backcompat" {
  availability_domain = data.oci_identity_availability_domain.backcompat.name
  compartment_id      = var.compartment_ocid
  display_name        = "tf-backcompat-block-instance"
  shape               = var.instance_shape

  shape_config {
    ocpus         = var.instance_ocpus
    memory_in_gbs = var.instance_shape_config_memory_in_gbs
  }

  create_vnic_details {
    assign_public_ip = false
    subnet_id        = oci_core_subnet.backcompat.id
  }

  source_details {
    source_type = "image"
    source_id   = lookup(data.oci_core_images.backcompat.images[0], "id")
  }

  instance_options {
    are_legacy_imds_endpoints_disabled = true
  }

  timeouts {
    create = "60m"
  }
}

resource "oci_core_boot_volume_backup" "backcompat" {
  boot_volume_id = oci_core_instance.backcompat.boot_volume_id
  display_name   = "tf-backcompat-boot-volume-backup"
}

data "oci_core_boot_volume_backup" "backcompat" {
  boot_volume_backup_id = oci_core_boot_volume_backup.backcompat.id
}

data "oci_core_boot_volume_backups" "backcompat" {
  compartment_id = var.compartment_ocid
  boot_volume_id = oci_core_instance.backcompat.boot_volume_id

  filter {
    name   = "id"
    values = [oci_core_boot_volume_backup.backcompat.id]
  }
}

resource "oci_core_volume" "backcompat" {
  availability_domain = data.oci_identity_availability_domain.backcompat.name
  compartment_id      = var.compartment_ocid
  display_name        = "tf-backcompat-volume"
  size_in_gbs         = "50"
}

resource "oci_core_volume_backup" "backcompat" {
  volume_id    = oci_core_volume.backcompat.id
  display_name = "tf-backcompat-volume-backup"
  type         = "FULL"
}

data "oci_core_volume_backups" "backcompat" {
  compartment_id = var.compartment_ocid
  volume_id      = oci_core_volume.backcompat.id

  filter {
    name   = "id"
    values = [oci_core_volume_backup.backcompat.id]
  }
}

resource "oci_core_volume_backup_policy" "backcompat" {
  compartment_id = var.compartment_ocid
  display_name   = "tf-backcompat-volume-backup-policy"

  schedules {
    backup_type       = "INCREMENTAL"
    period            = "ONE_DAY"
    retention_seconds = "604800"
  }
}

resource "oci_core_volume" "backcompat_group" {
  availability_domain = data.oci_identity_availability_domain.backcompat.name
  compartment_id      = var.compartment_ocid
  display_name        = "tf-backcompat-volume-group-volume"
  size_in_gbs         = "50"
}

resource "oci_core_volume_group" "backcompat" {
  availability_domain = data.oci_identity_availability_domain.backcompat.name
  compartment_id      = var.compartment_ocid
  display_name        = "tf-backcompat-volume-group"

  source_details {
    type       = "volumeIds"
    volume_ids = [oci_core_volume.backcompat_group.id]
  }
}

resource "oci_core_volume_group_backup" "backcompat" {
  volume_group_id = oci_core_volume_group.backcompat.id
  display_name    = "tf-backcompat-volume-group-backup"
  type            = "FULL"
}

data "oci_core_volume_group_backups" "backcompat" {
  compartment_id  = var.compartment_ocid
  volume_group_id = oci_core_volume_group.backcompat.id

  filter {
    name   = "id"
    values = [oci_core_volume_group_backup.backcompat.id]
  }
}

output "backcompat_block_backup_resources" {
  value = {
    boot_volume_backup        = data.oci_core_boot_volume_backup.backcompat.id
    boot_volume_backups_data  = data.oci_core_boot_volume_backups.backcompat.id
    volume_backup             = oci_core_volume_backup.backcompat.id
    volume_backups_data       = data.oci_core_volume_backups.backcompat.id
    volume_backup_policy      = oci_core_volume_backup_policy.backcompat.id
    volume_group_backup       = oci_core_volume_group_backup.backcompat.id
    volume_group_backups_data = data.oci_core_volume_group_backups.backcompat.id
  }
}
