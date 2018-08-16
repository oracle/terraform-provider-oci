variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "tenancy_ocid" {}
variable "user_ocid" {}

provider "oci" {
  alias            = "admin"
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

resource "oci_identity_group" "requestor_group" {
  provider    = "oci.admin"
  name        = "requestorGroup"
  description = "automated test group"
}

resource "oci_identity_user_group_membership" "requestor_user_group_membership" {
  provider = "oci.admin"
  group_id = "${oci_identity_group.requestor_group.id}"
  user_id  = "${var.user_requestor}"
}

resource "oci_identity_policy" "requestor_policy" {
  provider       = "oci.admin"
  name           = "requestorPolicy"
  description    = "automated test policy"
  compartment_id = "${var.tenancy_ocid}"

  statements = ["Allow group ${oci_identity_group.requestor_group.name} to manage virtual-network-family in compartment ${var.compartment_name_requestor}",
    "Allow group ${oci_identity_group.requestor_group.name} to manage instance-family in compartment ${var.compartment_name_requestor}",
  ]
}

resource "oci_identity_group" "acceptor_group" {
  provider    = "oci.admin"
  name        = "acceptorGroup"
  description = "automated test group"
}

resource "oci_identity_user_group_membership" "acceptor_user_group_membership" {
  provider = "oci.admin"
  group_id = "${oci_identity_group.acceptor_group.id}"
  user_id  = "${var.user_acceptor}"
}

resource "oci_identity_policy" "acceptor_policy" {
  provider       = "oci.admin"
  name           = "acceptorPolicy"
  description    = "automated test policy"
  compartment_id = "${var.tenancy_ocid}"

  statements = ["Allow group ${oci_identity_group.requestor_group.name} to manage local-peering-to in compartment ${var.compartment_name_acceptor}",
    "Allow group ${oci_identity_group.requestor_group.name} to inspect vcns in compartment ${var.compartment_name_acceptor}",
    "Allow group ${oci_identity_group.requestor_group.name} to inspect local-peering-gateways in compartment ${var.compartment_name_acceptor}",
    "Allow group ${oci_identity_group.acceptor_group.name} to manage virtual-network-family in compartment ${var.compartment_name_acceptor}",
    "Allow group ${oci_identity_group.acceptor_group.name} to manage instance-family in compartment ${var.compartment_name_acceptor}",
  ]
}
