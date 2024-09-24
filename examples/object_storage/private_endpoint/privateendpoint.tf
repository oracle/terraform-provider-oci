/*
 * This example shows how to manage a private endpoint resource
 */

variable "compartment_ocid" {
  default = "ocid1.tenancy.oc1..aaaaaaaaljp4td5xk4kxzgp2kvsreysajarqssyehkrdziff4fesqjhytl5q"
}

data "oci_identity_availability_domains" "ads1" {
  compartment_id = var.compartment_ocid
}

data "oci_objectstorage_namespace" "t1" {
  compartment_id = var.compartment_ocid
}

resource "oci_core_virtual_network" "test_vcn_1" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "network_name"
  dns_label      = "myvcn"
}

resource "oci_core_subnet" "test_subnet_1" {
  availability_domain = data.oci_identity_availability_domains.ads1.availability_domains.0.name
  cidr_block          = "10.0.1.0/24"
  display_name        = "-tf-subnet-1"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_virtual_network.test_vcn_1.id
  route_table_id      = oci_core_virtual_network.test_vcn_1.default_route_table_id
  dhcp_options_id     = oci_core_virtual_network.test_vcn_1.default_dhcp_options_id
  dns_label           = "testsubnet1"
}

resource "oci_objectstorage_private_endpoint" "testPe1" {
  compartment_id = var.compartment_ocid
  namespace = data.oci_objectstorage_namespace.t1.namespace
  name = "testPe1"
  subnet_id = oci_core_subnet.test_subnet_1.id
  prefix = "testPrefix1"
  access_targets  {
    namespace = "*"
    compartment_id = "*"
    bucket = "*"
  }
}

data "oci_objectstorage_private_endpoint_summaries" "testlist" {
  compartment_id = var.compartment_ocid
  namespace = data.oci_objectstorage_namespace.t1.namespace
  filter {
    name   = "name"
    values = [oci_objectstorage_private_endpoint.testPe1.name]
  }
}

