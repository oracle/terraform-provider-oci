variable "VPC-CIDR" {
  default = "10.0.0.0/16"
}

resource "oci_core_virtual_network" "cloudera_vcn" {
  cidr_block = "${var.VPC-CIDR}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "cloudera_vcn"
  dns_label = "cdhvcn"
}

resource "oci_core_internet_gateway" "cloudera_internet_gateway" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "cloudera_internet_gateway"
    vcn_id = "${oci_core_virtual_network.cloudera_vcn.id}"
}

resource "oci_core_route_table" "RouteForComplete" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.cloudera_vcn.id}"
    display_name = "RouteTableForComplete"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${oci_core_internet_gateway.cloudera_internet_gateway.id}"
    }
}

resource "oci_core_security_list" "PublicSubnet" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "Public Subnet"
    vcn_id = "${oci_core_virtual_network.cloudera_vcn.id}"
    egress_security_rules = [{
        destination = "0.0.0.0/0"
        protocol = "6"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 7180
            "min" = 7180
        }
        protocol = "6"
        source = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 80
            "min" = 80
        }
        protocol = "6"
        source = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 8888
            "min" = 8888
        }
        protocol = "6"
        source = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 22
            "min" = 22
        }
        protocol = "6"
        source = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        protocol = "6"
	source = "${var.VPC-CIDR}"
    }]


}

## Publicly Accessable Subnet Setup

resource "oci_core_subnet" "public" {
  count = "3"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[count.index],"name")}"
  cidr_block = "${cidrsubnet(var.VPC-CIDR, 8, count.index)}"
  display_name = "public_ad${count.index + 1}"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.cloudera_vcn.id}"
  route_table_id = "${oci_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${oci_core_security_list.PublicSubnet.id}"]
  dhcp_options_id = "${oci_core_virtual_network.cloudera_vcn.default_dhcp_options_id}"
  dns_label = "public${count.index + 1}"
}

