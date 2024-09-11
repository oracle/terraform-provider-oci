// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_vcn" "test_vcn_ipv6" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TestVcn"
  dns_label      = "testvcn"
  is_ipv6enabled = true
}

resource "oci_core_subnet" "test_subnet_ipv6" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "TestSubnet"
  dns_label           = "testsubnet"
  security_list_ids   = [oci_core_vcn.test_vcn_ipv6.default_security_list_id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_ipv6.id
  route_table_id      = oci_core_route_table.test_route_table_ipv6.id
  dhcp_options_id     = oci_core_vcn.test_vcn_ipv6.default_dhcp_options_id
  ipv6cidr_blocks     = ["${substr(oci_core_vcn.test_vcn_ipv6.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn_ipv6.ipv6cidr_blocks[0]) - 2)}${64}"]
}

resource "oci_core_internet_gateway" "test_internet_gateway_ipv6" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInternetGateway"
  vcn_id         = oci_core_vcn.test_vcn_ipv6.id
}

resource "oci_core_route_table" "test_route_table_ipv6" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn_ipv6.id
  display_name   = "TestRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway_ipv6.id
  }
}

resource "oci_load_balancer" "test_load_balancer_ipv6" {
  shape          = "100Mbps"
  compartment_id = var.compartment_ocid

  subnet_ids = [
    oci_core_subnet.test_subnet_ipv6.id,
  ]

  display_name = "TestLoadBalancer"
  is_private   = true
}

resource "oci_load_balancer_backend_set" "test_backend_set_ipv6" {
  name             = "lb-bes1"
  load_balancer_id = oci_load_balancer.test_load_balancer_ipv6.id
  policy           = "ROUND_ROBIN"

  health_checker {
    port                = "80"
    protocol            = "HTTP"
    response_body_regex = ".*"
    url_path            = "/"
  }
}

resource "oci_core_instance_configuration" "test_instance_configuration_ipv6" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInstanceConfiguration"

  instance_details {
    instance_type = "compute"

    /*
      Attach multiple block volumes
    */
    block_volumes {
      create_details {
        compartment_id      = var.compartment_ocid
        display_name        = "TestCreateVolumeDetails-1"
        availability_domain = data.oci_identity_availability_domain.ad.name
        size_in_gbs         = 50
        vpus_per_gb         = 20 // min vpus
      }

      attach_details {
        type                                = "paravirtualized"
        display_name                        = "TestAttachVolumeDetails-1"
        is_read_only                        = true
        is_shareable                        = true
      }
    }

    block_volumes {
      create_details {
        compartment_id      = var.compartment_ocid
        display_name        = "TestCreateVolumeDetails-2"
        availability_domain = data.oci_identity_availability_domain.ad.name
        size_in_gbs         = 50
        vpus_per_gb         = 20 // min vpus
      }

      attach_details {
        type                                = "paravirtualized"
        display_name                        = "TestAttachVolumeDetails-2"
        is_read_only                        = true
        is_shareable                        = true
      }
    }

    launch_details {
      compartment_id                      = var.compartment_ocid
      ipxe_script                         = "ipxeScript"
      shape                               = var.instance_shape
      display_name                        = "TestInstanceConfigurationLaunchDetails"
      is_pv_encryption_in_transit_enabled = false
      preferred_maintenance_action        = "LIVE_MIGRATE"
      launch_mode                         = "NATIVE"

      agent_config {
        is_management_disabled = false
        is_monitoring_disabled = false
      }

      launch_options {
        network_type = "PARAVIRTUALIZED"
      }

      instance_options {
        are_legacy_imds_endpoints_disabled = false
      }

      shape_config {
        ocpus = var.instance_ocpus
        memory_in_gbs = var.instance_shape_config_memory_in_gbs
      }

      create_vnic_details {
        assign_public_ip       = true
        display_name           = "TestInstanceConfigurationVNIC"
        skip_source_dest_check = false
        subnet_id                 = oci_core_subnet.test_subnet_ipv6.id
        assign_ipv6ip             = true
        ipv6address_ipv6subnet_cidr_pair_details {
          ipv6subnet_cidr = oci_core_subnet.test_subnet_ipv6.ipv6cidr_blocks[0]
        }
      }

      extended_metadata = {
        some_string   = "stringA"
        nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
      }

      source_details {
        source_type = "image"
        image_id    = var.flex_instance_image_ocid[var.region]
      }
    }

    secondary_vnics {
      display_name = "TestInstancePoolSecondaryVNIC"
      create_vnic_details {
        subnet_id = oci_core_subnet.test_subnet_ipv6.id
        assign_ipv6ip = true
        display_name = "TestInstancePoolSecondaryVNIC"
        ipv6address_ipv6subnet_cidr_pair_details {
          ipv6subnet_cidr = oci_core_subnet.test_subnet_ipv6.ipv6cidr_blocks[0]
        }
      }
    }
  }
}

resource "oci_core_instance_pool" "test_instance_pool_ipv6" {
  compartment_id = var.compartment_ocid
  instance_configuration_id = oci_core_instance_configuration.test_instance_configuration_ipv6.id
  size = 2
  state = "RUNNING"
  display_name = "TestInstancePool"

  placement_configurations {
    availability_domain = data.oci_identity_availability_domain.ad.name
    fault_domains = ["FAULT-DOMAIN-1"]
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
      display_name = "TestInstancePoolSecondaryVNIC"
      ipv6address_ipv6subnet_cidr_pair_details {
        ipv6subnet_cidr = oci_core_subnet.test_subnet_ipv6.ipv6cidr_blocks[0]
      }
    }
  }

  load_balancers {
    backend_set_name = oci_load_balancer_backend_set.test_backend_set_ipv6.name
    load_balancer_id = oci_load_balancer.test_load_balancer_ipv6.id
    port = 80
    vnic_selection = "primaryvnic"
  }

  lifecycle {
    ignore_changes = [size]
  }
}
