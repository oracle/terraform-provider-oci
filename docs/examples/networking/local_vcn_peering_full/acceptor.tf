variable "user_acceptor" {}
variable "compartment_ocid_acceptor" {}
variable "compartment_name_acceptor" {}
variable "fingerprint_acceptor" {}
variable "private_key_path_acceptor" {}

variable "acceptor_cidr" {
  default = "10.1.0.0/16"
}

provider "oci" {
  alias            = "acceptor"
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_acceptor}"
  fingerprint      = "${var.fingerprint_acceptor}"
  private_key_path = "${var.private_key_path_acceptor}"
}

resource "oci_core_vcn" "vcn2" {
  depends_on     = ["oci_identity_policy.acceptor_policy", "oci_identity_user_group_membership.acceptor_user_group_membership"]
  provider       = "oci.acceptor"
  display_name   = "vcn2"
  dns_label      = "vcn2"
  cidr_block     = "${var.acceptor_cidr}"
  compartment_id = "${var.compartment_ocid_acceptor}"
}

resource "oci_core_local_peering_gateway" "acceptor" {
  depends_on     = ["oci_identity_policy.acceptor_policy", "oci_identity_user_group_membership.acceptor_user_group_membership"]
  provider       = "oci.acceptor"
  compartment_id = "${var.compartment_ocid_acceptor}"
  vcn_id         = "${oci_core_vcn.vcn2.id}"
  display_name   = "localPeeringGateway2"
}

resource "oci_core_internet_gateway" "acceptorIG" {
  depends_on     = ["oci_identity_policy.acceptor_policy", "oci_identity_user_group_membership.acceptor_user_group_membership"]
  provider       = "oci.acceptor"
  compartment_id = "${var.compartment_ocid_acceptor}"
  display_name   = "acceptorIG"
  vcn_id         = "${oci_core_vcn.vcn2.id}"
}

resource "oci_core_route_table" "acceptor_route_table" {
  depends_on     = ["oci_identity_policy.acceptor_policy", "oci_identity_user_group_membership.acceptor_user_group_membership"]
  provider       = "oci.acceptor"
  compartment_id = "${var.compartment_ocid_acceptor}"
  vcn_id         = "${oci_core_vcn.vcn2.id}"
  display_name   = "acceptorRouteTable"

  route_rules {
    destination       = "${var.requestor_cidr}"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_local_peering_gateway.acceptor.id}"
  }

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.acceptorIG.id}"
  }
}

resource "oci_core_security_list" "acceptor_security_list" {
  depends_on     = ["oci_identity_policy.acceptor_policy", "oci_identity_user_group_membership.acceptor_user_group_membership"]
  provider       = "oci.acceptor"
  compartment_id = "${var.compartment_ocid_acceptor}"
  vcn_id         = "${oci_core_vcn.vcn2.id}"
  display_name   = "AcceptorSecurityList"

  egress_security_rules {
    destination = "${var.requestor_cidr}"
    protocol    = "all"
  }

  ingress_security_rules {
    protocol = "all"
    source   = "${var.requestor_cidr}"
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

resource "oci_core_subnet" "acceptor_subnet" {
  depends_on          = ["oci_identity_policy.acceptor_policy", "oci_identity_user_group_membership.acceptor_user_group_membership"]
  provider            = "oci.acceptor"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  cidr_block          = "${cidrsubnet("${var.acceptor_cidr}", 4, 0)}"
  display_name        = "AcceptorSubnet"
  dns_label           = "acceptorsubnet"
  compartment_id      = "${var.compartment_ocid_acceptor}"
  vcn_id              = "${oci_core_vcn.vcn2.id}"
  security_list_ids   = ["${oci_core_security_list.acceptor_security_list.id}"]
  route_table_id      = "${oci_core_route_table.acceptor_route_table.id}"
  dhcp_options_id     = "${oci_core_vcn.vcn2.default_dhcp_options_id}"
}

resource "oci_core_instance" "acceptor_instance" {
  depends_on          = ["oci_identity_policy.acceptor_policy", "oci_identity_user_group_membership.acceptor_user_group_membership"]
  provider            = "oci.acceptor"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid_acceptor}"
  display_name        = "acceptorInstance"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.acceptor_subnet.id}"
    display_name     = "primaryvnic"
    assign_public_ip = true
    hostname_label   = "acceptorinstance"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }
}
