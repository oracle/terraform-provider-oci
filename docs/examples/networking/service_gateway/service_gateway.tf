variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

variable "vcn_cidr_block" {
  default = "10.0.0.0/16"
}

variable "vcn_display_name" {
  default = "displayName"
}

variable "vcn_dns_label" {
  default = "dnslabel"
}

variable "service_gateway_display_name" {
  default = "displayName2"
}

variable "service_gateway_state" {
  default = "AVAILABLE"
}

variable "tcp_protocol" {
  default = "6"
}

variable "ssh_port" {
  default = "22"
}

resource "oci_core_vcn" "test_vcn" {
  #Required
  cidr_block     = "${var.vcn_cidr_block}"
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "${var.vcn_display_name}"
  dns_label    = "${var.vcn_dns_label}"
}

data "oci_core_services" "test_services" {
  filter {
    name   = "name"
    values = [".*Object.*Storage"]
    regex  = true
  }
}

resource "oci_core_service_gateway" "test_service_gateway" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  services {
    service_id = "${lookup(data.oci_core_services.test_services.services[0], "id")}"
  }

  vcn_id = "${oci_core_vcn.test_vcn.id}"

  #Optional
  display_name = "${var.service_gateway_display_name}"
}

data "oci_core_service_gateways" "test_service_gateways" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  state  = "${var.service_gateway_state}"
  vcn_id = "${oci_core_vcn.test_vcn.id}"
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.test_vcn.id}"
  display_name   = "testRouteTable"

  route_rules {
    destination       = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"
    destination_type  = "SERVICE_CIDR_BLOCK"
    network_entity_id = "${oci_core_service_gateway.test_service_gateway.id}"
  }
}

resource "oci_core_security_list" "test_security_list" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.test_vcn.id}"
  display_name   = "natSecurityList"

  egress_security_rules {
    destination      = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"
    destination_type = "SERVICE_CIDR_BLOCK"
    protocol         = "all"
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
