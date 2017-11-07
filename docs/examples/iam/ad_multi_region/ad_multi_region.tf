/*
 * This example demonstrates how to read AD values from multiple regions and employ filters
 * to isolate a specific AD value.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

provider "oci" {
  region = "us-phoenix-1"
  alias = "phx"
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

provider "oci" {
  region = "us-ashburn-1"
  alias = "iad"
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

data "oci_identity_availability_domains" "ad-phx" {
  provider = "oci.phx"
  compartment_id = "${var.tenancy_ocid}"

  filter {
    name = "name"
    values = ["\\w*-AD-1"]
    regex = true
  }
}

data "oci_identity_availability_domains" "ad-iad" {
  provider = "oci.iad"
  compartment_id = "${var.tenancy_ocid}"

  filter {
    name = "name"
    values = ["\\w*-AD-1"]
    regex = true
  }
}


output "ad-phx" {
  value = ["${data.oci_identity_availability_domains.ad-phx.availability_domains}"]
}

output "ad-iad" {
  value = ["${data.oci_identity_availability_domains.ad-iad.availability_domains}"]
}
