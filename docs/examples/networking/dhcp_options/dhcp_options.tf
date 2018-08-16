/*
 * This example demonstrates the various dhcp option configurations.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

resource "oci_core_virtual_network" "ExampleVCN" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_dhcp_options" "ExampleDhcpOptions1" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.ExampleVCN.id}"
  display_name   = "TFExampleDhcpOptions1"

  // required
  options {
    type        = "DomainNameServer"
    server_type = "VcnLocalPlusInternet"
  }

  // optional
  options {
    type                = "SearchDomain"
    search_domain_names = ["test.com"]
  }
}

resource "oci_core_dhcp_options" "ExampleDhcpOptions2" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.ExampleVCN.id}"
  display_name   = "TFExampleDhcpOptions2"

  // required
  options {
    type               = "DomainNameServer"
    server_type        = "CustomDnsServer"
    custom_dns_servers = ["8.8.4.4", "8.8.8.8"]
  }

  // optional
  options {
    type                = "SearchDomain"
    search_domain_names = ["test.com"]
  }
}
