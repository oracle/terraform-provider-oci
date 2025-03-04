// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_vcn" "my_vcn" {
  cidr_block     = var.my_vcn-cidr
  compartment_id = var.compartment_ocid
  display_name   = "myvcn"
  dns_label      = "myvcn"
}

resource "oci_core_internet_gateway" "my_internet_gateway" {
  compartment_id = var.compartment_ocid
  display_name   = "my internet gateway"
  vcn_id         = oci_core_vcn.my_vcn.id
}

resource "oci_core_route_table" "my_route_table" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.my_vcn.id
  display_name   = "my route table"
}

resource "oci_core_subnet" "my_subnet" {
  depends_on          = [oci_core_network_security_group.test_network_security_group]
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = var.my_subnet_cidr
  display_name        = "mysubnet"
  dns_label           = "mysubnet"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.my_vcn.id
  security_list_ids   = [oci_core_security_list.my_security_list.id]
  route_table_id      = oci_core_route_table.my_route_table.id
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.my_vcn.id
}

resource "oci_cluster_placement_groups_cluster_placement_group" "test_cpg" {
  #Required
  compartment_id = var.compartment_ocid
  description    = "cpg for lustre file system"
  name           = "test_cpg"
  availability_domain = data.oci_identity_availability_domain.ad.name
  cluster_placement_group_type = "STANDARD"

  #Optional
#   defined_tags = {
#     "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
#   }
#
#   freeform_tags = {
#     "Department" = "Finance"
#   }
}


