// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

func testProvider1() string {
	return `
	variable "tenancy_ocid" { default = "` + getRequiredEnvSetting("tenancy_ocid") + `" }
	variable "user_ocid" {default = "` + getRequiredEnvSetting("user_ocid") + `"}
	variable "fingerprint" {default = "` + getRequiredEnvSetting("fingerprint") + `"}
	variable "private_key_path" {default = "` + getRequiredEnvSetting("private_key_path") + `"}
	variable "region" {default = "` + getRequiredEnvSetting("region") + `"}
	variable "compartment_ocid" {default = "` + getRequiredEnvSetting("compartment_id") + `"}
	
	provider "oci" {
		tenancy_ocid = "${var.tenancy_ocid}"
		user_ocid = "${var.user_ocid}"
		fingerprint = "${var.fingerprint}"
		private_key_path = "${var.private_key_path}"
		region = "${var.region}"
	}`
}

func testADs() string {
	return `
	data "oci_identity_availability_domains" "t" {
		compartment_id = "${var.compartment_ocid}"
	}`
}

func testVCN1() string {
	return `
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_ocid}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
		dns_label    = "vcndns"
	}`
}

func testSubnet1() string {
	return `
	resource "oci_core_subnet" "t" {
		compartment_id      = "${var.compartment_ocid}"
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
	data "oci_core_images" "t" {
		compartment_id = "${var.compartment_ocid}"
		operating_system = "Oracle Linux"
		operating_system_version = "7.4"
		limit = 1
	}`
}

func testInstance1() string {
	return `
	resource "oci_core_instance" "t" {
		availability_domain = "${data.oci_identity_availability_domains.t.availability_domains.0.name}"
		compartment_id = "${var.compartment_ocid}"
		subnet_id = "${oci_core_subnet.t.id}"
		image = "${data.oci_core_images.t.images.0.id}"
		shape = "VM.Standard1.1"
		metadata {}
		timeouts {
			create = "15m"
		}
	}`
}
