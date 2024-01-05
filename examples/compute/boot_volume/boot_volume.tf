// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# This example creates a new boot volume from an existing instance

# This demo connects to the running instance so you will need to supply public/private keys to create an ssh connection.
# NOTE: do not try to use your api keys, see [this doc](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/managingkeypairs.htm)
# for more info on configuring keys.

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

variable "ssh_public_key" {
}

variable "ssh_private_key" {
}

variable "source_region" {
}

variable "source_boot_volume_backup_id" {
}

variable "display_name" {
  default = "TestBootVolumeBackupCopy"
}

variable "instance_image_ocid" {
  type = map(string)

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"
    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
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

resource "oci_core_boot_volume_backup" "test_boot_volume_backup_from_source_boot_volume" {
  #Required
  boot_volume_id = oci_core_boot_volume.test_boot_volume_from_source_boot_volume.id
}

resource "oci_core_boot_volume_backup" "test_boot_volume_backup_cross_region_sourced" {
  #Required
  source_details {
    region                = var.source_region
    boot_volume_backup_id = var.source_boot_volume_backup_id
  }

  #Optional
  display_name = var.display_name
}

resource "oci_core_boot_volume" "test_boot_volume_from_source_boot_volume_backup" {
  availability_domain = oci_core_instance.test_instance.availability_domain
  compartment_id      = oci_core_instance.test_instance.compartment_id

  source_details {
    #Required
    id   = oci_core_boot_volume_backup.test_boot_volume_backup_from_source_boot_volume.id
    type = "bootVolumeBackup"
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
    source_id   = var.instance_image_ocid[var.region]
  }

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
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

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_core_boot_volume_backups" "test_boot_volume_backup_from_source_boot_volume_datasource" {
  compartment_id = oci_core_instance.test_instance.compartment_id

  filter {
    name   = "id"
    values = [oci_core_boot_volume_backup.test_boot_volume_backup_from_source_boot_volume.id]
  }
}

data "oci_core_boot_volume_backups" "test_boot_volume_backup_from_source" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name                 = var.display_name
  source_boot_volume_backup_id = var.source_boot_volume_backup_id
}

data "oci_core_boot_volumes" "test_boot_volume_from_source_boot_volume_datasource" {
  #Required
  availability_domain = oci_core_boot_volume.test_boot_volume_from_source_boot_volume.availability_domain
  compartment_id      = oci_core_boot_volume.test_boot_volume_from_source_boot_volume.compartment_id

  filter {
    name   = "id"
    values = [oci_core_boot_volume.test_boot_volume_from_source_boot_volume.id]
  }
}

data "oci_core_boot_volumes" "test_boot_volume_from_source_boot_volume_backup_datasource" {
  #Required
  availability_domain = oci_core_boot_volume.test_boot_volume_from_source_boot_volume_backup.availability_domain
  compartment_id      = oci_core_boot_volume.test_boot_volume_from_source_boot_volume_backup.compartment_id

  filter {
    name   = "id"
    values = [oci_core_boot_volume.test_boot_volume_from_source_boot_volume_backup.id]
  }
}

output "boot_volume_from_instance_outputs" {
  value = {
    boot_volume_from_instance                             = oci_core_instance.test_instance.boot_volume_id
    boot_volume_from_source_boot_volume_id                = oci_core_boot_volume.test_boot_volume_from_source_boot_volume.id
    boot_volume_backup_from_source_boot_volume_id         = oci_core_boot_volume_backup.test_boot_volume_backup_from_source_boot_volume.id
    boot_volume_from_source_boot_volume_backup_id         = oci_core_boot_volume.test_boot_volume_from_source_boot_volume_backup.id
    boot_volume_from_source_boot_volume_datasource        = data.oci_core_boot_volumes.test_boot_volume_from_source_boot_volume_datasource.boot_volumes
    boot_volume_backup_from_source_boot_volume_datasource = data.oci_core_boot_volume_backups.test_boot_volume_backup_from_source_boot_volume_datasource.boot_volume_backups
    boot_volume_from_source_boot_volume_backup_datasource = data.oci_core_boot_volumes.test_boot_volume_from_source_boot_volume_backup_datasource.boot_volumes
  }
}

