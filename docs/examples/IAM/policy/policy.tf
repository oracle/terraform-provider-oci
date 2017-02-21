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
  private_key = "${var.private_key}"
}

  resource "baremetal_identity_policy" "IAM_policy" {
    name = "a_TF_managed_policy"
    description = "TF managed policy"
    compartment_id = "${var.compartment_ocid}"
    statements = ["allow group network-admins to manage virtual-network-family on tenancy","allow group a_TF_managed_group to manage all-resources on tenancy"]
  }
