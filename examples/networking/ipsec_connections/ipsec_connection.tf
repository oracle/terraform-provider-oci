// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "config_file_profile" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "ipsec_connection_tunnel_route_advertiser" {
  default = "CUSTOMER"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  auth = "SecurityToken"
  config_file_profile = var.config_file_profile
  region           = var.region
}

resource "oci_core_cpe" "test_cpe" {
  compartment_id      = var.compartment_ocid
  display_name        = "test_cpe"
  ip_address          = "189.44.2.135"
  cpe_device_shape_id = data.oci_core_cpe_device_shape.test_cpe_device_shape.id
}

resource "oci_core_drg" "test_drg" {
  compartment_id = var.compartment_ocid
  display_name   = "test_drg"
}

data "oci_core_cpe_device_shapes" "test_cpe_device_shapes" {
}

data "oci_core_cpe_device_shape" "test_cpe_device_shape" {
  cpe_device_shape_id = data.oci_core_cpe_device_shapes.test_cpe_device_shapes.cpe_device_shapes[0].cpe_device_shape_id
}

resource "oci_core_ipsec" "test_ip_sec_connection" {
  #Required
  compartment_id = var.compartment_ocid
  cpe_id         = oci_core_cpe.test_cpe.id
  drg_id         = oci_core_drg.test_drg.id
  static_routes  = ["10.0.0.0/16"]

  #Optional
  cpe_local_identifier      = "189.44.2.135"
  cpe_local_identifier_type = "IP_ADDRESS"
  defined_tags = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  }
  display_name = "MyIPSecConnection"

  freeform_tags = {
    "Department" = "Finance"
  }
}

data "oci_core_ipsec_connections" "test_ip_sec_connections" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  cpe_id = oci_core_cpe.test_cpe.id
  drg_id = oci_core_drg.test_drg.id
}

data "oci_core_ipsec_connection_tunnels" "test_ip_sec_connection_tunnels" {
  ipsec_id = oci_core_ipsec.test_ip_sec_connection.id
}

data "oci_core_ipsec_connection_tunnel" "test_ipsec_connection_tunnel" {
  ipsec_id  = oci_core_ipsec.test_ip_sec_connection.id
  tunnel_id = data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels[0].id
}

resource "oci_core_ipsec_connection_tunnel_management" "test_ipsec_connection_tunnel_management" {
  ipsec_id  = oci_core_ipsec.test_ip_sec_connection.id
  tunnel_id = data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels[0].id

  #Optional
  bgp_session_info {
    customer_bgp_asn      = "1587232876"
    customer_interface_ip = "10.0.0.16/31"
    oracle_interface_ip   = "10.0.0.17/31"
    customer_interface_ipv6 = "2002:db2::6/64"
    oracle_interface_ipv6   = "2002:db2::7/64"
  }

  display_name  = "MyIPSecConnection"
  routing       = "BGP"
  shared_secret = "sharedSecret"
  ike_version   = "V1"
}

resource "oci_core_ipsec_connection_tunnel_management" "test_ipsec_connection_second_tunnel_management" {
  ipsec_id  = oci_core_ipsec.test_ip_sec_connection.id
  tunnel_id = data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels[1].id

  #Optional
  display_name  = "MyIPSecConnection-Tunnel2"
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

data "oci_core_ipsec_connection_tunnel_routes" "test_ipsec_connection_tunnel_routes" {
  #Required
  ipsec_id  = oci_core_ipsec.test_ip_sec_connection.id
  tunnel_id = data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id

  #Optional
  advertiser = var.ipsec_connection_tunnel_route_advertiser
}


resource "oci_identity_tag_namespace" "tag_namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "Just a test"
  name           = "testexamples-tag-namespace"
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "tf example tag"
  name             = "tf-example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace1.id
}

resource "oci_identity_tag" "tag2" {
  #Required
  description      = "tf example tag 2"
  name             = "tf-example-tag-2"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace1.id
}

