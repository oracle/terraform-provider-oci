variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

resource "oci_core_vcn" "vcn1" {
  cidr_block     = "10.0.1.0/24"
  dns_label      = "vcn1"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "vcn1"
}

resource "oci_core_vcn" "vcn2" {
  cidr_block     = "10.0.2.0/24"
  dns_label      = "vcn2"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "vcn2"
}

resource "oci_core_vcn" "vcn3" {
  cidr_block     = "10.0.3.0/24"
  dns_label      = "vcn3"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "vcn3"
}

# Peer vcn1 and vcn2 to vcn3. You need one peering gateway on each VCN per peering connection.

resource "oci_core_local_peering_gateway" "test_local_peering_gateway_1" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn1.id}"

  #Optional
  display_name = "localPeeringGateway1"
  peer_id      = "${oci_core_local_peering_gateway.test_local_peering_gateway_3_A.id}"
}

resource "oci_core_local_peering_gateway" "test_local_peering_gateway_2" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn2.id}"

  #Optional
  display_name = "localPeeringGateway2"
  peer_id      = "${oci_core_local_peering_gateway.test_local_peering_gateway_3_B.id}"
}

resource "oci_core_local_peering_gateway" "test_local_peering_gateway_3_A" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn3.id}"

  #Optional
  display_name = "localPeeringGateway3A"
}

resource "oci_core_local_peering_gateway" "test_local_peering_gateway_3_B" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn3.id}"

  #Optional
  display_name = "localPeeringGateway3B"
}

data "oci_core_local_peering_gateways" "test_local_peering_gateways" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn3.id}"
}
