/*
 * This example shows how the policy resource and datasource works.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

variable "group_name" {}

provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

resource "baremetal_identity_compartment" "t" {
  name = "test-compartment"
  description = "automated test compartment"
}

resource "baremetal_identity_group" "t" {
  name = "-tf-group"
  description = "automated test group"
}

resource "baremetal_identity_policy" "p" {
  name = "-tf-policy"
  description = "automated test policy"
  compartment_id = "${baremetal_identity_compartment.t.id}"
  statements = ["Allow group ${baremetal_identity_group.t.name} to read instances in compartment ${baremetal_identity_compartment.t.name}"]
}

data "baremetal_identity_policies" "p" {
  compartment_id = "${baremetal_identity_compartment.t.id}"
}

output "policy" {
  value = ["${data.baremetal_identity_policies.p.policies}"]
}
