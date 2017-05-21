variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

provider "baremetal" {
  region = "us-phoenix-1"
  alias = "phx"
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

provider "baremetal" {
  region = "us-ashburn-1"
  alias = "iad"
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

resource "baremetal_core_virtual_network" "vcn-phx" {
  provider = "baremetal.phx"
  display_name = "vcn-phx"
  dns_label = "vcnwest"
  cidr_block = "10.0.0.0/16"
  compartment_id = "${var.compartment_ocid}"
}

resource "baremetal_core_virtual_network" "vcn-iad" {
  provider = "baremetal.iad"
  display_name = "vcn-iad"
  dns_label = "vcneast"
  cidr_block = "10.0.0.0/16"
  compartment_id = "${var.compartment_ocid}"
}
