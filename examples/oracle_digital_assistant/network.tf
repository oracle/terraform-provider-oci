variable "vcn_cidr_block" {
  default = "10.0.0.0/25"
}

variable "vcn_display_name" {
  default = "vcn_name"
}

variable "vcn_dns_label" {
  default = "testvcn"
}

variable "ig_display_name" {
  default = "ig_name"
}

variable "rt_display_name" {
  default = "rt_name"
}

variable "subnet_cidr_block" {
  default = "10.0.0.16/28"
}

variable "subnet_display_name" {
  default = "subnet_name"
}

variable "subnet_dns_label" {
  default = "test_subnet"
}

variable "security_list_dns_label" {
  default = "test_security_list"
}

resource "oci_core_vcn" "pe_vcn" {
  cidr_block     = var.vcn_cidr_block
  compartment_id = var.compartment_ocid
  display_name   = var.vcn_display_name
  dns_label      = var.vcn_dns_label
}

resource "oci_core_internet_gateway" "pe_internet_gateway" {
  compartment_id = var.vcn_cidr_block
  display_name   = var.ig_display_name
  vcn_id         = oci_core_vcn.pe_vcn.id
}

resource "oci_core_route_table" "pe_route_table" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.pe_vcn.id
  display_name   = var.rt_display_name

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.pe_internet_gateway.id
  }
}

resource "oci_core_subnet" "pe_subnet" {
  cidr_block          = var.subnet_cidr_block
  display_name        = var.subnet_display_name
  dns_label           = var.subnet_dns_label
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.pe_vcn.id
  security_list_ids   = ["${oci_core_security_list.pe_security_list.id}"]
  route_table_id      = ["${oci_core_route_table.pe_route_table.id}"]
  prohibit_public_ip_on_vnic = true
}

resource "oci_core_security_list" "pe_security_list" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.pe_vcn.id
  display_name   = var.security_list_dns_label

  egress_security_rules {
    protocol    = "all"
    destination = "0.0.0.0/0"
    stateless = false
  }

  ingress_security_rules {
    protocol = "all"
    source   = "0.0.0.0/0"
    stateless = false
  }
}