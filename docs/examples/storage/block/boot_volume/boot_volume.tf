/*
 * This example demonstrates how to backup and clone boot volumes
 *
 */
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}

variable "AD" {
  default = "1"
}

variable "InstanceShape" {
  default = "VM.Standard1.2"
}

variable "InstanceImageOCID" {
  type = "map"
  default = {
    // See https://docs.us-phoenix-1.oraclecloud.com/images/
    // Oracle-provided image "Oracle-Linux-7.4-2018.02.21-1"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaupbfz5f5hdvejulmalhyb6goieolullgkpumorbvxlwkaowglslq"
    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlw3xfie2t5t52uegyhiq2npx7bqyu4uvi2zyu3w3mqayc2bxmaa"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaa7d3fsb6272srnftyi4dphdgfjf6gurxqhmv6ileds7ba3m2gltxq"
    uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaa6h6gj6v4n56mqrbgnosskq63blyv2752g36zerymy63cfkojiiq"
  }
}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_core_instance" "TFInstance" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "TFInstance"
  shape = "${var.InstanceShape}"

  create_vnic_details {
    subnet_id = "${oci_core_subnet.ExampleSubnet.id}"
    display_name = "primaryvnic"
    assign_public_ip = true
    hostname_label = "tfexampleinstance${count.index}"
  },

  source_details {
    source_type = "image"
    source_id = "${var.InstanceImageOCID[var.region]}"
  }

  timeouts {
    create = "60m"
  }
}

resource "oci_core_virtual_network" "ExampleVCN" {
  cidr_block = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name = "TFExampleVCN"
  dns_label = "tfexamplevcn"
}

resource "oci_core_subnet" "ExampleSubnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  cidr_block = "10.1.20.0/24"
  display_name = "TFExampleSubnet"
  dns_label = "tfexamplesubnet"
  security_list_ids = ["${oci_core_virtual_network.ExampleVCN.default_security_list_id}"]
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.ExampleVCN.id}"
  route_table_id = "${oci_core_route_table.ExampleRT.id}"
  dhcp_options_id = "${oci_core_virtual_network.ExampleVCN.default_dhcp_options_id}"
}

resource "oci_core_internet_gateway" "ExampleIG" {
  compartment_id = "${var.compartment_ocid}"
  display_name = "TFExampleIG"
  vcn_id = "${oci_core_virtual_network.ExampleVCN.id}"
}

resource "oci_core_route_table" "ExampleRT" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.ExampleVCN.id}"
  display_name = "TFExampleRouteTable"
  route_rules {
    cidr_block = "0.0.0.0/0"
    network_entity_id = "${oci_core_internet_gateway.ExampleIG.id}"
  }
}

resource "oci_core_boot_volume" "TFBootVolumeFromBootVolume" {
  #Required
  availability_domain = "${oci_core_instance.TFInstance.availability_domain}"
  compartment_id = "${var.compartment_ocid}"
  source_details {
    #Required
    type = "bootVolume"
    id = "${oci_core_instance.TFInstance.boot_volume_id}"
  }
}

resource "oci_core_boot_volume_backup" "TFBootVolumeBackup" {
  #Required
  boot_volume_id = "${oci_core_instance.TFInstance.boot_volume_id}"
}

resource "oci_core_boot_volume" "TFBootVolumeFromBootVolumeBackup" {
  #Required
  availability_domain = "${oci_core_instance.TFInstance.availability_domain}"
  compartment_id = "${var.compartment_ocid}"
  source_details {
    #Required
    type = "bootVolumeBackup"
    id = "${oci_core_boot_volume_backup.TFBootVolumeBackup.id}"
  }
}

output "SourceBootVolumeId" {
  value = "${oci_core_instance.TFInstance.boot_volume_id}"
}

output "ClonedBootVolumeIdFromSourceBootVolumeId" {
  value = "${oci_core_boot_volume.TFBootVolumeFromBootVolume.id}"
}

output "BootVolumeBackupId" {
  value = "${oci_core_boot_volume_backup.TFBootVolumeBackup.id}"
}

output "ClonedBootVolumeIdFromSourceBootVolumeBackupId" {
  value = "${oci_core_boot_volume.TFBootVolumeFromBootVolumeBackup.id}"
}

