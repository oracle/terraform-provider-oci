/*
 * This example demonstrates how to target multiple regions from one plan. It creates two vcns
 * in different regions.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

provider "oci" {
  region           = "us-phoenix-1"
  alias            = "phx"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

provider "oci" {
  region           = "us-ashburn-1"
  alias            = "iad"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

resource "oci_core_virtual_network" "vcn-phx" {
  provider       = "oci.phx"
  display_name   = "vcn-phx"
  dns_label      = "vcnwest"
  cidr_block     = "10.0.0.0/16"
  compartment_id = "${var.compartment_ocid}"
}

resource "oci_core_virtual_network" "vcn-iad" {
  provider       = "oci.iad"
  display_name   = "vcn-iad"
  dns_label      = "vcneast"
  cidr_block     = "10.0.0.0/16"
  compartment_id = "${var.compartment_ocid}"
}
