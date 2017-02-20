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

resource "baremetal_core_route_table" "a_route_table" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "ocid1.vcn.oc1.phx.aaaaaa...j2jfxsq"
    display_name = "my_route_table"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "ocid1.internetgateway.oc1.phx.aaa...mz25a"
    }
}
