variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "group_description" { default = "Group for network administrators" }
variable "group_name" { default = "NetworkAdmins" }


provider "oci" {
    tenancy_ocid = "${var.tenancy_ocid}"
    user_ocid = "${var.user_ocid}"
    fingerprint = "${var.fingerprint}"
    private_key_path = "${var.private_key_path}"
    region = "${var.region}"
}

resource "oci_identity_group" "test_group" {
	#Required
	compartment_id = "${var.compartment_id}"
	description = "${var.group_description}"
	name = "${var.group_name}"
}

data "oci_identity_groups" "test_groups" {
	#Required
	compartment_id = "${var.compartment_id}"
}
