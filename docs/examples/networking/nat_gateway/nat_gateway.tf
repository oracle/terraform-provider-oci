variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
    tenancy_ocid = "${var.tenancy_ocid}"
    user_ocid = "${var.user_ocid}"
    fingerprint = "${var.fingerprint}"
    private_key_path = "${var.private_key_path}"
    region = "${var.region}"
}

variable "vcn_cidr_block" { default = "10.0.0.0/16" }
variable "vcn_display_name" { default = "displayName" }
variable "vcn_dns_label" { default = "dnslabel" }

variable "nat_gateway_display_name" { default = "displayName" }
variable "nat_gateway_block_traffic" { default = false }

variable "ssh_public_key" {}
variable "InstanceShape" { default = "VM.Standard1.2" }

variable "BootStrapFile" { default = "./userdata/bootstrap" }

variable "InstanceImageOCID" {
    type = "map"
    default = {
        r1 = "ocid1.image.region1.sea.aaaaaaaam3qd34kzrkxqcmgaadb7meolnwlffkpz2iv5xiwopxgwkkhq5e6a"
        // See https://docs.us-phoenix-1.oraclecloud.com/Content/Resources/Assets/OracleProvidedImageOCIDs.pdf
        // Oracle-provided image "Oracle-Linux-7.4-2018.02.21-1"
        us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaupbfz5f5hdvejulmalhyb6goieolullgkpumorbvxlwkaowglslq"
        us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlw3xfie2t5t52uegyhiq2npx7bqyu4uvi2zyu3w3mqayc2bxmaa"
        eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt1.aaaaaaaa7d3fsb6272srnftyi4dphdgfjf6gurxqhmv6ileds7ba3m2gltxq"
        uk-london-1 = "ocid1.image.oc1.uk-london1.aaaaaaaaa6h6gj6v4n56mqrbgnosskq63blyv2752g36zerymy63cfkojiiq"
    }
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_core_vcn" "test_vcn" {
	#Required
	cidr_block = "${var.vcn_cidr_block}"
	compartment_id = "${var.compartment_ocid}"

	#Optional
	display_name = "${var.vcn_display_name}"
	dns_label = "${var.vcn_dns_label}"
}

resource "oci_core_nat_gateway" "test_nat_gateway" {
	#Required
	compartment_id = "${var.compartment_ocid}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.nat_gateway_display_name}"
	#block_traffic = "${var.nat_gateway_block_traffic}"
}

data "oci_core_nat_gateways" "test_nat_gateways" {
	#Required
	compartment_id = "${var.compartment_ocid}"

	#Optional
	#display_name = "${var.nat_gateway_display_name}"
	#state = "${var.nat_gateway_state}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_vcn.test_vcn.id}"
  display_name = "testRouteTable"
  route_rules {
    cidr_block = "0.0.0.0/0"
    network_entity_id = "${oci_core_nat_gateway.test_nat_gateway.id}"
  }
}

resource "oci_core_security_list" "test_security_list" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_vcn.test_vcn.id}"
  display_name = "natSecurityList"

  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol = "all"
  }
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block = "${cidrsubnet("${oci_core_vcn.test_vcn.cidr_block}", 4, 0)}"
  display_name = "testSubnet"
  dns_label = "testsubnet"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_vcn.test_vcn.id}"
  security_list_ids = ["${oci_core_security_list.test_security_list.id}"]
  route_table_id = "${oci_core_route_table.test_route_table.id}"
  dhcp_options_id = "${oci_core_vcn.test_vcn.default_dhcp_options_id}"
}

resource "oci_core_instance" "test_instance" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "testInstance"
  image = "${var.InstanceImageOCID[var.region]}"
  shape = "${var.InstanceShape}"

  create_vnic_details {
    subnet_id = "${oci_core_subnet.test_subnet.id}"
    display_name = "primaryvnic"
    assign_public_ip = true
    hostname_label = "testinstance"
  },

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data = "${base64encode(file(var.BootStrapFile))}"
  }
}

