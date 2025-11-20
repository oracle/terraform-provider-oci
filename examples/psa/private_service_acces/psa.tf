// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "private_service_access_defined_tags_value" {
  default = "value"
}

variable "private_service_access_description" {
  default = "description"
}

variable "private_service_access_display_name" {
  default = "displayName"
}

variable "private_service_access_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "private_service_access_ipv4ip" {
  default = "10.1.20.100"
}

variable "private_service_access_security_attributes" {
  default = { "oracle-zpr.sensitivity.value" = "low", "oracle-zpr.sensitivity.mode" = "enforce"}
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_vcn" "vcn1" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "vcn1"
  dns_label      = "vcn1"
  is_ipv6enabled =  true
}

resource "oci_core_subnet" "subnet1" {
  cidr_block          = "10.1.20.0/24"
  display_name        = "subnet1"
  dns_label           = "subnet1"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.vcn1.id
  dhcp_options_id     = oci_core_vcn.vcn1.default_dhcp_options_id

  provisioner "local-exec" {
    command = "sleep 5"
  }
}

resource "oci_core_network_security_group" "test_nsg" {
  compartment_id = var.compartment_ocid
  display_name   = "tfNsgForPSA"
  vcn_id         = oci_core_vcn.vcn1.id
}


data "oci_psa_psa_services" "test_psa_services" {
}

resource "oci_psa_private_service_access" "test_private_service_access" {
  #Required
  compartment_id = var.compartment_ocid
  service_id     = data.oci_psa_psa_services.test_psa_services.psa_service_collection.0.items.0.id
  subnet_id      = oci_core_subnet.subnet1.id

  #Optional
  defined_tags        = map("example-tag-namespace-all.example-tag", var.private_service_access_defined_tags_value)
  description         = var.private_service_access_description
  display_name        = var.private_service_access_display_name
  freeform_tags       = var.private_service_access_freeform_tags
  ipv4ip              = var.private_service_access_ipv4ip
  nsg_ids             = [oci_core_network_security_group.test_nsg.id]
  security_attributes = var.private_service_access_security_attributes
}

data "oci_psa_private_service_access" "test_private_service_access" {
  private_service_access_id = oci_psa_private_service_access.test_private_service_access.id
}

data "oci_psa_private_service_accesses" "test_private_service_accesses" {
  #Optional
  compartment_id = var.compartment_ocid
  depends_on = [oci_psa_private_service_access.test_private_service_access]
}
