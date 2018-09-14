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

variable "availability_domain" {
  default = 3
}

resource "oci_core_virtual_network" "ExampleVCN" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleVCN"
  dns_label      = "tfexamplevcn"
}

/* 
Because you can specify multiple security lists/subnet the security_list_ids value must be specified as a list in []'s.
 See https://www.terraform.io/docs/configuration/syntax.html
   
Generally you wouldn't specify a subnet without first specifying a VCN. Once the VCN has been created you would get the vcn_id, route_table_id, and security_list_id(s) from that resource and use Terraform attributes below to populate those values.
 See https://www.terraform.io/docs/configuration/interpolation.html*/
resource "oci_core_subnet" "ExampleSubnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  cidr_block          = "10.1.1.0/24"
  display_name        = "TFExampleSubnet"
  dns_label           = "tfexamplesubnet"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.ExampleVCN.id}"
  security_list_ids   = ["${oci_core_virtual_network.ExampleVCN.default_security_list_id}"]
  route_table_id      = "${oci_core_virtual_network.ExampleVCN.default_route_table_id}"
  dhcp_options_id     = "${oci_core_virtual_network.ExampleVCN.default_dhcp_options_id}"
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}
