// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_core_vcn" "ExampleVCN" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "ExampleSubnet" {
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  cidr_block          = "10.1.20.0/24"
  display_name        = "TFExampleSubnet"
  dns_label           = "tfexamplesubnet"
  security_list_ids   = ["${oci_core_vcn.ExampleVCN.default_security_list_id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_vcn.ExampleVCN.id}"
  route_table_id      = "${oci_core_route_table.ExampleRT.id}"
  dhcp_options_id     = "${oci_core_vcn.ExampleVCN.default_dhcp_options_id}"
}

resource "oci_core_internet_gateway" "ExampleIG" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleIG"
  vcn_id         = "${oci_core_vcn.ExampleVCN.id}"
}

resource "oci_core_route_table" "ExampleRT" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.ExampleVCN.id}"
  display_name   = "TFExampleRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.ExampleIG.id}"
  }
}

resource "oci_load_balancer" "test_load_balancer" {
  shape          = "100Mbps"
  compartment_id = "${var.compartment_ocid}"

  subnet_ids = [
    "${oci_core_subnet.ExampleSubnet.id}",
  ]

  display_name = "test_load_balancer"
  is_private   = true
}

resource "oci_load_balancer_backend_set" "test_backend_set" {
  name             = "lb-bes1"
  load_balancer_id = "${oci_load_balancer.test_load_balancer.id}"
  policy           = "ROUND_ROBIN"

  health_checker {
    port                = "80"
    protocol            = "HTTP"
    response_body_regex = ".*"
    url_path            = "/"
  }
}
