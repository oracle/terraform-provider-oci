// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "region" {
}

variable "compartment_ocid" {
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_vcn" "example_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "example_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "exampleSubnet"
  dns_label           = "tfexamplesubnet"
  security_list_ids   = [oci_core_vcn.example_vcn.default_security_list_id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.example_vcn.id
  route_table_id      = oci_core_vcn.example_vcn.default_route_table_id
  dhcp_options_id     = oci_core_vcn.example_vcn.default_dhcp_options_id
}

variable "instance_shape" {
  default = "VM.Standard.E4.Flex"
}

variable "instance_ocpus" { default = 1 }

variable "instance_shape_config_memory_in_gbs" { default = 6 }

# See https://docs.oracle.com/iaas/images/
data "oci_core_images" "test_images" {
  compartment_id           = var.compartment_ocid
  operating_system         = "Oracle Linux"
  operating_system_version = "8"
  sort_by                  = "TIMECREATED"
  sort_order               = "DESC"
}

# Create Instance
resource "oci_core_instance" "test_instance1" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "testInstance"
  shape               = var.instance_shape

  shape_config {
    ocpus = var.instance_ocpus
    memory_in_gbs = var.instance_shape_config_memory_in_gbs
  }

  create_vnic_details {
    subnet_id      = oci_core_subnet.example_subnet.id
    hostname_label = "instance"
  }

  source_details {
    source_type = "image"
    source_id = lookup(data.oci_core_images.test_images.images[0], "id")
  }
}

# Gets a list of VNIC attachments on the instance
data "oci_core_vnic_attachments" "instance_vnics" {
  compartment_id      = var.compartment_ocid
  availability_domain = data.oci_identity_availability_domain.ad.name
  instance_id         = oci_core_instance.test_instance1.id
}

# Gets the OCID of the first (default) VNIC
data "oci_core_vnic" "instance_vnic" {
  vnic_id = data.oci_core_vnic_attachments.instance_vnics.vnic_attachments[0]["vnic_id"]
}

# Create PrivateIP
resource "oci_core_private_ip" "private_ip" {
  vnic_id        = data.oci_core_vnic_attachments.instance_vnics.vnic_attachments[0]["vnic_id"]
  display_name   = "someDisplayName"
  hostname_label = "somehostnamelabel"
  route_table_id = oci_core_vcn.example_vcn.default_route_table_id
}

# List Private IPs
data "oci_core_private_ips" "private_ip_datasource" {
  depends_on = [oci_core_private_ip.private_ip]
  vnic_id    = oci_core_private_ip.private_ip.vnic_id
}

output "private_ips" {
  value = [data.oci_core_private_ips.private_ip_datasource.private_ips]
}

