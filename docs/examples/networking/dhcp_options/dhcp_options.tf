/*
 * This example demonstrates the various dhcp option configurations.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

variable "vcn_ocid" {}


provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}


resource "baremetal_core_dhcp_options" "dhcp-options1" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${var.vcn_ocid}"
  display_name = "dhcp-options1"

  // required
  options {
    type = "DomainNameServer"
    server_type = "VcnLocalPlusInternet"
  }

  // optional
  options {
    type = "SearchDomain"
    search_domain_names = [ "test.com" ]
  }
}


resource "baremetal_core_dhcp_options" "dhcp-options2" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${var.vcn_ocid}"
  display_name = "dhcp-options2"

  // required
  options {
    type = "DomainNameServer"
    server_type = "CustomDnsServer"
    custom_dns_servers = [  "8.8.4.4", "8.8.8.8" ]
  }

  // optional
  options {
    type = "SearchDomain"
    search_domain_names = [ "test.com" ]
  }
}
