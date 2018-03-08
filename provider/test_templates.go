// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

func testADs() string {
	return `
	data "oci_identity_availability_domains" "t" {
		compartment_id = "${var.compartment_id}"
	}`
}

func testVCN1() string {
	return `
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
		dns_label    = "vcndns"
	}`
}

func testSubnet1() string {
	return `
	resource "oci_core_subnet" "t" {
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.t.availability_domains[0],"name")}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		cidr_block          = "10.0.1.0/24"
		display_name        = "-tf-subnet"
		dns_label           = "subnetdns"
	}`
}

func testImage1() string {
	return `
	variable "InstanceImageOCID" {
	  type = "map"
	  default = {
		// Oracle-provided image "Oracle-Linux-7.4-2017.12.18-0"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaasc56hnpnx7swoyd2fw5gyvbn3kcdmqc2guiiuvnztl2erth62xnq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaxrqeombwty6jyqgk3fraczdd63bv66xgfsqka4ktr7c57awr3p5a"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaayxmzu6n5hsntq4wlffpb4h6qh6z3uskpbm5v3v4egqlqvwicfbyq"
	  }
	}`
}

func testInstance1() string {
	return `
	resource "oci_core_instance" "t" {
		availability_domain = "${data.oci_identity_availability_domains.t.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		subnet_id = "${oci_core_subnet.t.id}"
		image = "${var.InstanceImageOCID[var.region]}"
		shape = "VM.Standard1.1"
		metadata {}
		timeouts {
			create = "15m"
		}
	}`
}
