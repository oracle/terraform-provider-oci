/*
 * This example demonstrates how to read AD values from multiple regions and outputs them for display.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

provider "baremetal" {
  region = "us-phoenix-1"
  alias = "phx"
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

provider "baremetal" {
  region = "us-ashburn-1"
  alias = "iad"
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

data "baremetal_identity_availability_domains" "ad-phx" {
  compartment_id = "${var.tenancy_ocid}"
  provider = "baremetal.phx"
}

data "baremetal_identity_availability_domains" "ad-iad" {
  compartment_id = "${var.tenancy_ocid}"
  provider = "baremetal.iad"
}


output "ad-phx" {
  value = ["${data.baremetal_identity_availability_domains.ad-phx.availability_domains}"]
}

output "ad-iad" {
  value = ["${data.baremetal_identity_availability_domains.ad-iad.availability_domains}"]
}