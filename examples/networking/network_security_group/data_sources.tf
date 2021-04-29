// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_core_network_security_groups" "test_network_security_groups" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = oci_core_network_security_group.test_network_security_group.display_name
  state        = oci_core_network_security_group.test_network_security_group.state
  vcn_id       = oci_core_vcn.test_vcn.id
}

data "oci_core_network_security_groups" "test_network_security_groups_vlan" {
  #Required
  vlan_id        = oci_core_vlan.test_vlan.id
}

data "oci_core_network_security_group_security_rules" "test_network_security_group_security_rules" {
  #Required
  network_security_group_id = oci_core_network_security_group.test_network_security_group.id

  #Optional
  direction = "EGRESS"
}

data "oci_core_network_security_group_vnics" "test_network_security_group_vnics" {
  #Required
  network_security_group_id = oci_core_network_security_group.test_network_security_group.id
}

data "oci_core_services" "test_services" {
  filter {
    name   = "name"
    values = ["OCI .* Object Storage"]
    regex  = true
  }
}

