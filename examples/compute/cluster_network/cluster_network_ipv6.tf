// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_vcn" "test_vcn_ipv6" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TestVcn"
  dns_label      = "dnslabel"
  is_ipv6enabled = true
}

resource "oci_core_route_table" "test_route_table_ipv6" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn_ipv6.id
  display_name   = "TestRouteTable"
}

resource "oci_core_subnet" "test_subnet_ipv6" {
  availability_domain = lower(
    data.oci_identity_availability_domains.test_availability_domains.availability_domains[1].name,
  )
  cidr_block        = "10.0.2.0/24"
  compartment_id    = var.compartment_ocid
  dhcp_options_id   = oci_core_vcn.test_vcn_ipv6.default_dhcp_options_id
  display_name      = "TestSubnet"
  dns_label         = "dnslabel"
  route_table_id    = oci_core_route_table.test_route_table_ipv6.id
  security_list_ids = [oci_core_vcn.test_vcn_ipv6.default_security_list_id]
  vcn_id            = oci_core_vcn.test_vcn_ipv6.id
  ipv6cidr_blocks   = ["${substr(oci_core_vcn.test_vcn_ipv6.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn_ipv6.ipv6cidr_blocks[0]) - 2)}${64}"]
}

resource "oci_core_instance_configuration" "test_instance_configuration_ipv6" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInstanceConfiguration"

  instance_details {
    instance_type = "compute"

    launch_details {
      availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[1].name
      compartment_id      = var.compartment_ocid

      create_vnic_details {
        assign_public_ip       = "false"
        display_name           = "backend-servers"
        hostname_label         = "hostnameLabel"
        private_ip             = "privateIp"
        skip_source_dest_check = "false"
        subnet_id              = oci_core_subnet.test_subnet_ipv6.id
        ipv6address_ipv6subnet_cidr_pair_details {
          ipv6subnet_cidr = oci_core_subnet.test_subnet_ipv6.ipv6cidr_blocks[0]
        }
      }

      display_name = "backend-servers"

      extended_metadata = {
        "extendedMetadata" = "extendedMetadata"
      }

      ipxe_script = "ipxeScript"

      metadata = {
        "metadata" = "metadata"
      }

      shape = "BM.HPC2.36"

      source_details {
        boot_volume_size_in_gbs = "55"
        image_id                = var.InstanceImageOCID[var.region]
        source_type             = "image"
      }
    }
    secondary_vnics {
      display_name = "TestClusterNetworkSecondaryVNIC"
      create_vnic_details {
        subnet_id = oci_core_subnet.test_subnet_ipv6.id
        assign_ipv6ip = true
        display_name = "TestClusterNetworkSecondaryVNIC"
        ipv6address_ipv6subnet_cidr_pair_details {
          ipv6subnet_cidr = oci_core_subnet.test_subnet_ipv6.ipv6cidr_blocks[0]
        }
      }
    }
  }

  source = "NONE"
}

resource "oci_core_cluster_network" "test_cluster_network_ipv6" {
  compartment_id = var.compartment_ocid
  display_name   = "hpc-cluster-network"

  instance_pools {
    display_name              = "hpc-cluster-network"
    instance_configuration_id = oci_core_instance_configuration.test_instance_configuration_ipv6.id
    size                      = "1"
  }

  placement_configuration {
    availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[1].name
    primary_vnic_subnets {
      subnet_id = oci_core_subnet.test_subnet_ipv6.id
      is_assign_ipv6ip = true
      ipv6address_ipv6subnet_cidr_pair_details {
        ipv6subnet_cidr = oci_core_subnet.test_subnet_ipv6.ipv6cidr_blocks[0]
      }
    }
    secondary_vnic_subnets {
      subnet_id = oci_core_subnet.test_subnet_ipv6.id
      is_assign_ipv6ip = true
      display_name = "TestClusterNetworkSecondaryVNIC"
      ipv6address_ipv6subnet_cidr_pair_details {
        ipv6subnet_cidr = oci_core_subnet.test_subnet_ipv6.ipv6cidr_blocks[0]
      }
    }
  }
}
