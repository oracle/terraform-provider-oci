/*
 * This example demonstrates how to work with nested compartments.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

resource "oci_identity_compartment" "parent-compartment" {
  name           = "parent-compartment"
  description    = "compartment that holds a compartment"
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_identity_compartment" "child-compartment" {
  name           = "child-compartment"
  description    = "compartment inside another compartment"
  compartment_id = "${oci_identity_compartment.parent-compartment.id}"
}

data "oci_identity_compartments" "all-compartments" {
  compartment_id            = "${oci_identity_compartment.parent-compartment.compartment_id}"
  compartment_id_in_subtree = "true"
  access_level              = "ANY"

  filter {
    name   = "name"
    values = ["parent-compartment", "child-compartment"]
  }
}

data "oci_identity_compartment" "child-compartment" {
  id = "${oci_identity_compartment.child-compartment.id}"
}

output "print-child-compartment" {
  value = <<EOF

  id = ${data.oci_identity_compartment.child-compartment.id}
  compartment_id = ${data.oci_identity_compartment.child-compartment.compartment_id}
  name = ${data.oci_identity_compartment.child-compartment.name}
  description = ${data.oci_identity_compartment.child-compartment.description}
EOF
}

output "print-all-compartments" {
  value = "${data.oci_identity_compartments.all-compartments.compartments}"
}
