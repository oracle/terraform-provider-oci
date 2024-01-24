// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# Instance with IPV6 enabled

resource "oci_core_vcn" "test_vcn_with_ipv6" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "CompleteVCN"
  dns_label      = "examplevcn"
  is_ipv6enabled = true
}

resource "oci_core_subnet" "test_subnet_with_ipv6" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.0.1.0/24"
  display_name        = "TestSubnet"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_with_ipv6.id
  route_table_id      = oci_core_vcn.test_vcn_with_ipv6.default_route_table_id
  security_list_ids   = [oci_core_vcn.test_vcn_with_ipv6.default_security_list_id]
  dhcp_options_id     = oci_core_vcn.test_vcn_with_ipv6.default_dhcp_options_id
  dns_label           = "examplesubnet"
  ipv6cidr_blocks     = ["${substr(oci_core_vcn.test_vcn_with_ipv6.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn_with_ipv6.ipv6cidr_blocks[0]) - 2)}${64}"]
}

resource "oci_core_instance" "test_instance_with_ipv6" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstanceIpv6"
  shape               = var.instance_shape

  create_vnic_details {
    subnet_id                 = oci_core_subnet.test_subnet_with_ipv6.id
    display_name              = "Primaryvnic"
    assign_public_ip          = true
    assign_private_dns_record = true
    hostname_label            = "ipv6instance"
    assign_ipv6ip             = true

    ipv6address_ipv6subnet_cidr_pair_details {
      ipv6subnet_cidr = oci_core_subnet.test_subnet_with_ipv6.ipv6cidr_blocks[0]
      ipv6address = "${substr(oci_core_subnet.test_subnet_with_ipv6.ipv6cidr_blocks[0], 0, length(oci_core_subnet.test_subnet_with_ipv6.ipv6cidr_blocks[0]) - 7)}${1000}"
    }
  }

  shape_config {
    ocpus = var.instance_ocpus
    memory_in_gbs = var.instance_shape_config_memory_in_gbs
  }

  source_details {
    source_type = "image"
    source_id = var.flex_instance_image_ocid[var.region]
  }

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
  }

  timeouts {
    create = "60m"
  }
}
