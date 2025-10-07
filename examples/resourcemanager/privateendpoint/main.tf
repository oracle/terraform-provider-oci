variable "tenancy_ocid" {}
variable "compartment_ocid" {}
variable "region" {}
variable "auth" {}
variable "config_file_profile" {}
variable "user_ocid" {
}
variable "fingerprint" {
}
variable "private_key_path" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

# Locals and data
locals {
  vcn_cidr_block           = "10.12.0.0/16"
  default_shape_name       = "VM.Standard.E3.Flex"
  operating_system         = "Oracle Linux"
  operating_system_version = "8"
  tcp_protocol             = 6
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = local.vcn_cidr_block
  compartment_id = var.compartment_ocid
  display_name   = "test_vcn"
}

resource "oci_core_subnet" "test_subnet" {
  compartment_id             = var.compartment_ocid
  vcn_id                     = oci_core_vcn.test_vcn.id
  display_name               = "test_subnet"
  prohibit_public_ip_on_vnic = true
  cidr_block                 = cidrsubnet(local.vcn_cidr_block, 8, 1)
}

resource "oci_resourcemanager_private_endpoint" "my_private_endpoint" {
  compartment_id = var.compartment_ocid
  display_name   = "my_private_endpoint"
  description    = "Example"
  vcn_id         = oci_core_vcn.test_vcn.id
  subnet_id      = oci_core_subnet.test_subnet.id
  /*security_attributes = {
    "oracle-zpr.maxegresscount.mode"  = "enforce"
    "oracle-zpr.maxegresscount.value" = "42"
  }*/
}

data "oci_identity_availability_domains" "ads" { compartment_id = var.tenancy_ocid }

data "oci_core_images" "ol" {
  compartment_id           = var.compartment_ocid
  operating_system         = local.operating_system
  operating_system_version = local.operating_system_version
  shape                    = local.default_shape_name
}

resource "tls_private_key" "ssh" { algorithm = "RSA" }

resource "oci_core_instance" "test_instance" {
  compartment_id      = var.compartment_ocid
  display_name        = "test_instance"
  availability_domain = data.oci_identity_availability_domains.ads.availability_domains[0].name
  shape               = local.default_shape_name
  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    assign_public_ip = false
  }
  extended_metadata = {
    ssh_authorized_keys = tls_private_key.ssh.public_key_openssh
  }
  source_details {
    source_id   = data.oci_core_images.ol.images[0].id
    source_type = "image"
  }
  shape_config {
    memory_in_gbs = 4
    ocpus         = 1
  }
}

data "oci_resourcemanager_private_endpoint_reachable_ip" "instance_ip" {
  private_endpoint_id = oci_resourcemanager_private_endpoint.my_private_endpoint.id
  private_ip          = oci_core_instance.test_instance.private_ip
}

output "reachable_ip" {
  value = data.oci_resourcemanager_private_endpoint_reachable_ip.instance_ip.ip_address
}