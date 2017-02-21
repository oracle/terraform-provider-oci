variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

resource "baremetal_identity_compartment" "a_TF_managed_compartment" {
  name = "a_TF_managed_compartment"
  description = "My first Terraform Compartment"
}
