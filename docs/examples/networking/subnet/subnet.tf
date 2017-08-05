variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

variable "vcn_ocid" {}
variable "dhcp_options_ocid" {}
variable "route_table_ocid" {}
variable "security_list_ocid" {}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

/* 
Because you can specify multiple security lists/subnet the security_list_ids value must be specified as a list in []'s.
 See https://www.terraform.io/docs/configuration/syntax.html
   
Generally you wouldn't specify a subnet without first specifying a VCN. Once the VCN has been created you would get the vcn_id, route_table_id, and security_list_id(s) from that resource and use Terraform attributes below to populate those values.
 See https://www.terraform.io/docs/configuration/interpolation.html*/

resource "oci_core_subnet" "a_TF_managed_subnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block = "10.0.1.0/24"
  display_name = "subnet1"
  dns_label = "subnet1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${var.vcn_ocid}}"
  security_list_ids = ["${var.security_list_ocid}"]
  route_table_id = "${var.route_table_ocid}"
  dhcp_options_id = "${var.dhcp_options_ocid}"
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}