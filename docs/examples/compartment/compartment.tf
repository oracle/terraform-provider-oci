variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  private_key = "${var.private_key}"
}

resource "baremetal_identity_compartment" "a_TF_managed_compartment" {
  compartment_id = "${var.tenancy_ocid}"
  name = "a_TF_managed_compartment"
  description = "My first Terraform Compartment"
}
