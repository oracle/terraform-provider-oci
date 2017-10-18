/*
 * This example demonstrates how to spin up a block volume
 *
 * See docs/examples/compute/instance/ for a real world scenario
 */
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

variable "DBSize" {
  default = "50" // size in GBs, min: 50, max 16384
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_core_volume" "t" {
  availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
  compartment_id = "${var.tenancy_ocid}"
  display_name = "-tf-volume"
  size_in_gbs = "${var.DBSize}"
}
