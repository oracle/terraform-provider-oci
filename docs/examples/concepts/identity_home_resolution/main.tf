/*
 * This example demonstrates how to use tenancy and region data sources to route identity calls to the correct
 * home region.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

variable "region" {
  default = "us-ashburn-1"
}

provider "oci" {
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

data "oci_identity_tenancy" "tenancy" {
  tenancy_id = "${var.tenancy_ocid}"
}

data "oci_identity_regions" "home-region" {
  filter {
    name   = "key"
    values = ["${data.oci_identity_tenancy.tenancy.home_region_key}"]
  }
}

provider "oci" {
  alias            = "home"
  region           = "${lookup(data.oci_identity_regions.home-region.regions[0], "name")}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

resource "oci_identity_user" "user1" {
  provider    = "oci.home"
  name        = "tf-example-user"
  description = "user created by terraform"
}
