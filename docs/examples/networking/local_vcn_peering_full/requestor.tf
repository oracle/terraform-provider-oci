variable "user_requestor" {}
variable "compartment_ocid_requestor" {}
variable "compartment_name_requestor" {}
variable "fingerprint_requestor" {}
variable "private_key_path_requestor" {}

variable "requestor_cidr" {
  default = "10.0.0.0/16"
}

provider "oci" {
  alias            = "requestor"
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_requestor}"
  fingerprint      = "${var.fingerprint_requestor}"
  private_key_path = "${var.private_key_path_requestor}"
}

resource "oci_core_vcn" "vcn1" {
  depends_on     = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider       = "oci.requestor"
  display_name   = "vcn1"
  dns_label      = "vcn1"
  cidr_block     = "${var.requestor_cidr}"
  compartment_id = "${var.compartment_ocid_requestor}"
}

resource "oci_core_local_peering_gateway" "requestor" {
  depends_on     = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider       = "oci.requestor"
  compartment_id = "${var.compartment_ocid_requestor}"
  vcn_id         = "${oci_core_vcn.vcn1.id}"
  display_name   = "localPeeringGateway1"
  peer_id        = "${oci_core_local_peering_gateway.acceptor.id}"
}

resource "oci_core_internet_gateway" "requestorIG" {
  depends_on     = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider       = "oci.requestor"
  compartment_id = "${var.compartment_ocid_requestor}"
  display_name   = "requestorIG"
  vcn_id         = "${oci_core_vcn.vcn1.id}"
}

resource "oci_core_route_table" "requestor_route_table" {
  depends_on     = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider       = "oci.requestor"
  compartment_id = "${var.compartment_ocid_requestor}"
  vcn_id         = "${oci_core_vcn.vcn1.id}"
  display_name   = "requestorRouteTable"

  route_rules {
    destination       = "${var.acceptor_cidr}"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_local_peering_gateway.requestor.id}"
  }

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.requestorIG.id}"
  }
}

resource "oci_core_security_list" "requestor_security_list" {
  depends_on     = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider       = "oci.requestor"
  compartment_id = "${var.compartment_ocid_requestor}"
  vcn_id         = "${oci_core_vcn.vcn1.id}"
  display_name   = "RequestorSecurityList"

  egress_security_rules {
    destination = "${var.acceptor_cidr}"
    protocol    = "all"
  }

  ingress_security_rules {
    protocol = "all"
    source   = "${var.acceptor_cidr}"
  }

  ingress_security_rules {
    protocol = "${var.tcp_protocol}"
    source   = "0.0.0.0/0"

    tcp_options {
      max = "${var.ssh_port}"
      min = "${var.ssh_port}"
    }
  }
}

resource "oci_core_subnet" "requestor_subnet" {
  depends_on          = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider            = "oci.requestor"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain -1],"name")}"
  cidr_block          = "${cidrsubnet("${var.requestor_cidr}", 4, 0)}"
  display_name        = "RequestorSubnet"
  dns_label           = "requestorsubnet"
  compartment_id      = "${var.compartment_ocid_requestor}"
  vcn_id              = "${oci_core_vcn.vcn1.id}"
  security_list_ids   = ["${oci_core_security_list.requestor_security_list.id}"]
  route_table_id      = "${oci_core_route_table.requestor_route_table.id}"
  dhcp_options_id     = "${oci_core_vcn.vcn1.default_dhcp_options_id}"
}

resource "oci_core_instance" "requestor_instance" {
  depends_on          = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider            = "oci.requestor"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain -1],"name")}"
  compartment_id      = "${var.compartment_ocid_requestor}"
  display_name        = "requestorInstance"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.requestor_subnet.id}"
    display_name     = "primaryvnic"
    assign_public_ip = true
    hostname_label   = "requestorinstance"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }
}
