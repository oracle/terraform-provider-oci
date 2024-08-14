// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# This example creates a new boot volume from an existing instance

variable "boot_vol_first_backup_ocid" {
}

variable "boot_vol_second_backup_ocid" {
}

variable "instance_image_ocid" {
}

resource "oci_core_boot_volume" "test_boot_volume_from_source_boot_volume" {
  availability_domain = oci_core_instance.test_instance.availability_domain
  compartment_id      = oci_core_instance.test_instance.compartment_id

  source_details {
    #Required
    id   = oci_core_instance.test_instance.boot_volume_id
    type = "bootVolume"
  }
}

resource "oci_core_boot_volume" "test_create_delta_restored_boot_volume" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "tf-delta-restored-boot-volume-1"
  size_in_gbs         = "56"
  source_details {
    first_backup_id   = var.boot_vol_first_backup_ocid
    second_backup_id = var.boot_vol_second_backup_ocid
    change_block_size_in_bytes = 4096
    type = "bootVolumeBackupDelta"
  }
}

resource "oci_core_instance" "test_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstance"
  shape               = "VM.Standard2.1"

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
    hostname_label   = "testinstance"
  }

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid
  }

  timeouts {
    create = "60m"
  }
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TestVcn"
  dns_label      = "testvcn"
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "TestSubnet"
  dns_label           = "testsubnet"
  security_list_ids   = [oci_core_vcn.test_vcn.default_security_list_id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  route_table_id      = oci_core_vcn.test_vcn.default_route_table_id
  dhcp_options_id     = oci_core_vcn.test_vcn.default_dhcp_options_id
}

output "boot_volume_from_instance_outputs" {
  value = {
    boot_volume_from_instance                             = oci_core_instance.test_instance.boot_volume_id
    boot_volume_from_source_boot_volume_id                = oci_core_boot_volume.test_boot_volume_from_source_boot_volume.id
  }
}

