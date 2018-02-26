/*
 * This example shows how the policy resource and datasource works.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

resource "oci_identity_compartment" "t" {
  name = "test-compartment"
  description = "automated test compartment"
}

resource "oci_identity_group" "t" {
  name = "TFExampleGroup"
  description = "automated test group"
}

resource "oci_identity_policy" "t" {
  name = "TFExamplePolicy"
  description = "automated test policy"
  compartment_id = "${oci_identity_compartment.t.id}"
  statements = ["Allow group ${oci_identity_group.t.name} to read instances in compartment ${oci_identity_compartment.t.name}"]
}

data "oci_identity_policies" "p" {
  compartment_id = "${oci_identity_compartment.t.id}"
}

output "policy" {
  value = ["${data.oci_identity_policies.p.policies}"]
}
