variable "user_requestor" {}
variable "compartment_ocid_requestor" {}
variable "compartment_name_requestor" {}
variable "fingerprint_requestor" {}
variable "private_key_path_requestor" {}

variable "requestor_region" {
  default = "us-phoenix-1"
}

variable "requestor_cidr" {
  default = "10.0.0.0/16"
}

provider "oci" {
  alias            = "requestor"
  region           = "${var.requestor_region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_requestor}"
  fingerprint      = "${var.fingerprint_requestor}"
  private_key_path = "${var.private_key_path_requestor}"
}

resource "oci_core_vcn" "requestor_vcn" {
  depends_on     = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider       = "oci.requestor"
  display_name   = "requestor_vcn"
  dns_label      = "requestorvcn"
  cidr_block     = "${var.requestor_cidr}"
  compartment_id = "${var.compartment_ocid_requestor}"
}

resource "oci_core_drg" "requestor_drg" {
  depends_on     = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider       = "oci.requestor"
  compartment_id = "${var.compartment_ocid_requestor}"
}

resource "oci_core_drg_attachment" "requestor_drg_attachment" {
  depends_on = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider   = "oci.requestor"
  drg_id     = "${oci_core_drg.requestor_drg.id}"
  vcn_id     = "${oci_core_vcn.requestor_vcn.id}"
}

resource "oci_core_remote_peering_connection" "requestor" {
  depends_on       = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider         = "oci.requestor"
  compartment_id   = "${var.compartment_ocid_requestor}"
  drg_id           = "${oci_core_drg.requestor_drg.id}"
  display_name     = "remotePeeringConnectionRequestor"
  peer_id          = "${oci_core_remote_peering_connection.acceptor.id}"
  peer_region_name = "${var.acceptor_region}"
}

resource "oci_core_internet_gateway" "requestor_internet_gateway" {
  depends_on     = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider       = "oci.requestor"
  compartment_id = "${var.compartment_ocid_requestor}"
  display_name   = "requestor_internet_gateway"
  vcn_id         = "${oci_core_vcn.requestor_vcn.id}"
}

resource "oci_core_route_table" "requestor_route_table" {
  depends_on     = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider       = "oci.requestor"
  compartment_id = "${var.compartment_ocid_requestor}"
  vcn_id         = "${oci_core_vcn.requestor_vcn.id}"
  display_name   = "requestorRouteTable"

  route_rules {
    destination       = "${var.acceptor_cidr}"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_drg.requestor_drg.id}"
  }

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.requestor_internet_gateway.id}"
  }
}

resource "oci_core_security_list" "requestor_security_list" {
  depends_on     = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider       = "oci.requestor"
  compartment_id = "${var.compartment_ocid_requestor}"
  vcn_id         = "${oci_core_vcn.requestor_vcn.id}"
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

data "oci_identity_availability_domains" "requestor_ads" {
  provider       = "oci.requestor"
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_core_subnet" "requestor_subnet" {
  depends_on          = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider            = "oci.requestor"
  availability_domain = "${lookup(data.oci_identity_availability_domains.requestor_ads.availability_domains[var.availability_domain -1],"name")}"
  cidr_block          = "${cidrsubnet("${var.requestor_cidr}", 4, 0)}"
  display_name        = "RequestorSubnet"
  dns_label           = "requestorsubnet"
  compartment_id      = "${var.compartment_ocid_requestor}"
  vcn_id              = "${oci_core_vcn.requestor_vcn.id}"
  security_list_ids   = ["${oci_core_security_list.requestor_security_list.id}"]
  route_table_id      = "${oci_core_route_table.requestor_route_table.id}"
  dhcp_options_id     = "${oci_core_vcn.requestor_vcn.default_dhcp_options_id}"
}

resource "oci_core_instance" "requestor_instance" {
  depends_on          = ["oci_identity_policy.requestor_policy", "oci_identity_user_group_membership.requestor_user_group_membership"]
  provider            = "oci.requestor"
  availability_domain = "${lookup(data.oci_identity_availability_domains.requestor_ads.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid_requestor}"
  display_name        = "requestorInstance"

  #image = "${lookup(data.oci_core_images.requestor_images.images[0], "id")}"
  shape = "${var.instance_shape}"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.requestor_subnet.id}"
    display_name     = "primaryvnic"
    assign_public_ip = true
    hostname_label   = "requestorinstance"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.requestor_region]}"
  }

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }
}
