// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_cpe" "test_cpe_ipsec_over_fc" {
  compartment_id      = var.compartment_ocid
  display_name        = "test_cpe_ipsec_over_fc"
  ip_address          = "10.1.6.7"
  cpe_device_shape_id = data.oci_core_cpe_device_shape.test_cpe_device_shape.id
  is_private          = true
}

resource "oci_core_drg" "test_drg_ipsec_over_fc" {
  compartment_id = var.compartment_ocid
  display_name   = "test_drg_ipsec_over_fc"
}

resource "oci_core_drg_route_table" "test_drg_ipsec_over_fc_route_table" {
  drg_id = oci_core_drg.test_drg_ipsec_over_fc.id
  display_name = "DrgRtForIpsecOverFC"
}

data "oci_core_cross_connect_locations" "cross_connect_locations" {
  #Required
  compartment_id = var.compartment_ocid
}

data "oci_core_cross_connect_port_speed_shapes" "cross_connect_port_speed_shapes" {
  #Required
  compartment_id = var.compartment_ocid
}

resource "oci_core_cross_connect" "test_ipsec_over_fc_cross_connect" {
  compartment_id        = var.compartment_ocid
  location_name         = data.oci_core_cross_connect_locations.cross_connect_locations.cross_connect_locations[0].name
  port_speed_shape_name = data.oci_core_cross_connect_port_speed_shapes.cross_connect_port_speed_shapes.cross_connect_port_speed_shapes[0].name
  display_name = "testIpsecOverFCCrossConnect"
  #Set Cross Connect to Active to provision (required to provision virtual circuits).
  #You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up
  is_active = true
}

resource "oci_core_virtual_circuit" "test_ipsec_over_fc_virtual_circuit" {
  compartment_id = var.compartment_ocid
  type           = "PRIVATE"
  bandwidth_shape_name = "100 Mbps"
  cross_connect_mappings {
    cross_connect_or_cross_connect_group_id = oci_core_cross_connect.test_ipsec_over_fc_cross_connect.id
    vlan = 101
    oracle_bgp_peering_ip = "10.0.1.21/30"
    customer_bgp_peering_ip = "10.0.1.22/30"
  }
  customer_asn = "64513"
  display_name = "testIpsecOverFCVirtualCircuit"
  gateway_id = oci_core_drg.test_drg_ipsec_over_fc.id
}

resource "oci_core_ipsec" "test_ipsec_over_fc_connection" {
  #Required
  compartment_id = var.compartment_ocid
  cpe_id         = oci_core_cpe.test_cpe_ipsec_over_fc.id
  drg_id         = oci_core_drg.test_drg_ipsec_over_fc.id
  static_routes  = ["10.0.0.0/16"]
  tunnel_configuration {
    oracle_tunnel_ip = "10.1.5.5"
    associated_virtual_circuits = [oci_core_virtual_circuit.test_ipsec_over_fc_virtual_circuit.id]
    drg_route_table_id = oci_core_drg_route_table.test_drg_ipsec_over_fc_route_table.id
  }
  tunnel_configuration {
    oracle_tunnel_ip = "10.1.7.7"
    associated_virtual_circuits = [oci_core_virtual_circuit.test_ipsec_over_fc_virtual_circuit.id]
    drg_route_table_id = oci_core_drg_route_table.test_drg_ipsec_over_fc_route_table.id
  }

  #Optional
  cpe_local_identifier      = "10.1.6.7"
  cpe_local_identifier_type = "IP_ADDRESS"
  defined_tags = {
    "${oci_identity_tag_namespace.tag_namespace1_fc.name}.${oci_identity_tag.tag1_fc.name}" = "value"
  }
  display_name = "MyIPSecConnectionOverFC"

  freeform_tags = {
    "Department" = "Finance"
  }
}

data "oci_core_ipsec_connections" "test_ip_sec_over_fc_connections" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  cpe_id = oci_core_cpe.test_cpe_ipsec_over_fc.id
  drg_id = oci_core_drg.test_drg_ipsec_over_fc.id
}

data "oci_core_ipsec_connection_tunnels" "test_ip_sec_connection_tunnels_over_fc" {
  ipsec_id = oci_core_ipsec.test_ipsec_over_fc_connection.id
}

data "oci_core_ipsec_connection_tunnel" "test_ipsec_connection_tunnel_over_fc" {
  ipsec_id  = oci_core_ipsec.test_ipsec_over_fc_connection.id
  tunnel_id = data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels_over_fc.ip_sec_connection_tunnels[0].id
}

resource "oci_core_ipsec_connection_tunnel_management" "test_ipsec_connection_tunnel_management_over_fc" {
  ipsec_id  = oci_core_ipsec.test_ipsec_over_fc_connection.id
  tunnel_id = data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels_over_fc.ip_sec_connection_tunnels[0].id

  #Optional
  bgp_session_info {
    customer_bgp_asn      = "1587232876"
    customer_interface_ip = "10.0.0.16/31"
    oracle_interface_ip   = "10.0.0.17/31"
    customer_interface_ipv6 = "2002:db2::6/64"
    oracle_interface_ipv6   = "2002:db2::7/64"
  }

  display_name  = "MyIPSecConnectionOverFCTunnelMgmt"
  routing       = "BGP"
  shared_secret = "sharedSecret"
  ike_version   = "V1"
}

resource "oci_core_ipsec_connection_tunnel_management" "test_ipsec_connection_second_tunnel_management_over_fc" {
  ipsec_id  = oci_core_ipsec.test_ipsec_over_fc_connection.id
  tunnel_id = data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels_over_fc.ip_sec_connection_tunnels[1].id

  #Optional
  display_name  = "MyIPSecConnectionOverFC-Tunnel2"
  routing       = "POLICY"
  shared_secret = "sharedSecret"
  ike_version   = "V1"

  nat_translation_enabled = "ENABLED"
  oracle_can_initiate = "RESPONDER_ONLY"

  encryption_domain_config {
    cpe_traffic_selector = ["10.0.0.16/31", "11.0.0.16/31"]
    oracle_traffic_selector = ["12.0.0.16/31"]
  }

  phase_one_details {
    is_custom_phase_one_config = false
    lifetime = 28600
  }

  phase_two_details{
    dh_group = "GROUP20"
    is_custom_phase_two_config = false
    is_pfs_enabled = true
    lifetime = 3602
  }
}

data "oci_core_ipsec_connection_tunnel_routes" "test_ipsec_connection_tunnel_routes_over_fc" {
  #Required
  ipsec_id  = oci_core_ipsec.test_ipsec_over_fc_connection.id
  tunnel_id = data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels_over_fc.ip_sec_connection_tunnels.0.id

  #Optional
  advertiser = var.ipsec_connection_tunnel_route_advertiser
}

resource "oci_identity_tag_namespace" "tag_namespace1_fc" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "Just a test"
  name           = "testexamples-tag-namespace"
}

resource "oci_identity_tag" "tag1_fc" {
  #Required
  description      = "tf example tag"
  name             = "tf-example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace1_fc.id
}

resource "oci_identity_tag" "tag2_fc" {
  #Required
  description      = "tf example tag 2"
  name             = "tf-example-tag-2"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace1_fc.id
}

