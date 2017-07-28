variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

variable "vcn_ocid" {}


provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

# Protocols are specified as protocol numbers.
# http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml

resource "baremetal_core_security_list" "security_list1" {
  compartment_id = "${var.compartment_ocid}"
  display_name = "security_list1"
  vcn_id = "${var.vcn_ocid}}"
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol = "6"
  }
  ingress_security_rules {
    tcp_options {
      "max" = 22
      "min" = 22
    }
    protocol = "6"
    source = "0.0.0.0/0"
  }
}
