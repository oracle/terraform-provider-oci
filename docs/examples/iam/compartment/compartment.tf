variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "compartment_description" { default = "For network components" }
variable "compartment_name" { default = "Network" }


provider "oci" {
    tenancy_ocid = "${var.tenancy_ocid}"
    user_ocid = "${var.user_ocid}"
    fingerprint = "${var.fingerprint}"
    private_key_path = "${var.private_key_path}"
    region = "${var.region}"
}

resource "oci_identity_compartment" "test_compartment" {
	#Required
	compartment_id = "${var.compartment_id}"
	description = "${var.compartment_description}"
	name = "${var.compartment_name}"
}

data "oci_identity_compartments" "test_compartments" {
	#Required
	compartment_id = "${var.compartment_id}"
}
