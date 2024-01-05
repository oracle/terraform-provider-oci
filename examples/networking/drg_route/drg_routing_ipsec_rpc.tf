// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_drg" "test_drg" {
  // required
  compartment_id = var.compartment_ocid

  // optional
  display_name = "MyTestDrg"
}

resource "oci_core_drg_route_distribution" "test_drg_route_distribution" {
  // Required
  drg_id = oci_core_drg.test_drg.id
  distribution_type = "IMPORT"

  // Optional
  display_name = "MyTestDrgRouteDistribution"
}

resource "oci_core_drg_route_table" "test_drg_route_table" {
  // Required
  drg_id = oci_core_drg.test_drg.id

  // Optional
  import_drg_route_distribution_id = oci_core_drg_route_distribution.test_drg_route_distribution.id
  display_name = "MyTestDrgRouteTable"
}

// RPC
resource "oci_core_remote_peering_connection" "test_rpc" {

  compartment_id = var.compartment_ocid
  drg_id         = oci_core_drg.test_drg.id
  display_name   = "MyTestRemotePeeringConnection"
}


// IPSec Start

resource "oci_core_ipsec" "test_ip_sec_connection" {
  // Required
  compartment_id = var.compartment_ocid
  cpe_id = oci_core_cpe.test_cpe.id
  drg_id = oci_core_drg.test_drg.id
  static_routes = [
    "10.0.0.0/16"]

  // Optional
  cpe_local_identifier = "189.44.2.135"
  cpe_local_identifier_type = "IP_ADDRESS"
  display_name = "MyTestIPSecConnection"

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_core_cpe" "test_cpe" {
  compartment_id = var.compartment_ocid
  display_name = "test_cpe"
  ip_address = "189.44.2.135"
  cpe_device_shape_id = data.oci_core_cpe_device_shape.test_cpe_device_shape.id
}

data "oci_core_cpe_device_shape" "test_cpe_device_shape" {
  cpe_device_shape_id = data.oci_core_cpe_device_shapes.test_cpe_device_shapes.cpe_device_shapes[0].cpe_device_shape_id
}

data "oci_core_ipsec_connection_tunnels" "test_ip_sec_connection_tunnels" {
  ipsec_id = oci_core_ipsec.test_ip_sec_connection.id
}
data "oci_core_cpe_device_shapes" "test_cpe_device_shapes" {
}
//  IPsec end


// Get auto generated attachment for ipsec tunnel 1
resource "oci_core_drg_attachment_management" "test_drg_ipsec_attachment_tunnel_1" {
  // Required
  attachment_type = "IPSEC_TUNNEL"
  compartment_id = var.compartment_ocid
  drg_id = oci_core_drg.test_vcn_drg.id
  network_id = data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels[0].id

  // Optional
  display_name = "MyTestDrgAttachmentForTunnel1"
  drg_route_table_id = oci_core_drg_route_table.test_drg_route_table.id

}
// Get auto generated attachment for ipsec tunnel 2
resource "oci_core_drg_attachment_management" "test_drg_ipsec_attachment_tunnel_2" {
  // Required
  attachment_type = "IPSEC_TUNNEL"
  compartment_id = var.compartment_ocid
  drg_id = oci_core_drg.test_vcn_drg.id
  network_id = data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels[1].id

  // Optional
  display_name = "MyTestDrgAttachmentForTunnel2"
  drg_route_table_id = oci_core_drg_route_table.test_drg_route_table.id

}

// Get auto generated attachment for rpc by rpc.id
resource "oci_core_drg_attachment_management" "test_drg_rpc_attachment" {
  // Required
  attachment_type = "REMOTE_PEERING_CONNECTION"
  compartment_id = var.compartment_ocid
  network_id = oci_core_remote_peering_connection.test_rpc.id
  drg_id = oci_core_drg.test_vcn_drg.id

  // Optional
  display_name = "MyTestDrgAttachmentForRpc"
  drg_route_table_id = oci_core_drg_route_table.test_drg_route_table.id

}
//

// Add Static route
resource "oci_core_drg_route_table_route_rule" "test_drg_rpc_route_table_route_rule" {
  drg_route_table_id = oci_core_drg_route_table.test_drg_route_table.id
  destination = "10.0.0.0/16"
  destination_type = "CIDR_BLOCK"
  next_hop_drg_attachment_id = oci_core_drg_attachment_management.test_drg_rpc_attachment.id
}