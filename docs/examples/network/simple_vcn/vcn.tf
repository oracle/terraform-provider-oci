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

resource "baremetal_core_virtual_network" "a_TF_managed_VCN" {
  cidr_block = "10.0.0.0/16"
  dns_label = "vnc1"
  compartment_id = "${var.compartment_ocid}"
  display_name = "a_TF_managed_VCN"
}
