// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.
variable "ip_sec_connection_cpe_local_identifier" {
  default = "189.44.2.135"
}

variable "ip_sec_connection_cpe_local_identifier_type" {
  default = "IP_ADDRESS"
}

variable "ip_sec_connection_defined_tags_value" {
  default = "value"
}

variable "ip_sec_connection_display_name" {
  default = "MyIPSecConnection"
}

variable "ip_sec_connection_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "ip_sec_connection_static_routes" {
  default = ["10.0.0.0/16"]
}

variable "ip_sec_connection_tunnel_configuration_bgp_session_config_customer_bgp_asn" {
  default = "1587232876"
}

variable "ip_sec_connection_tunnel_configuration_bgp_session_config_customer_interface_ip" {
  default = "10.0.0.16/31"
}

variable "ip_sec_connection_tunnel_configuration_bgp_session_config_oracle_interface_ip" {
  default = "10.0.0.17/31"
}

variable "ip_sec_connection_tunnel_configuration_display_name" {
  default = "MyIPSecConnection"
}

variable "ip_sec_connection_tunnel_configuration_routing" {
  default = "BGP"
}

variable "ip_sec_connection_tunnel_configuration_shared_secret" {
  default = "sharedSecret"
}

resource oci_core_cpe "test_cpe" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "test_cpe"
  ip_address     = "189.44.2.135"
}

resource oci_core_drg "test_drg" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "test_drg"
}

resource "oci_core_ipsec" "test_ip_sec_connection" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  cpe_id         = "${oci_core_cpe.test_cpe.id}"
  drg_id         = "${oci_core_drg.test_drg.id}"
  static_routes  = "${var.ip_sec_connection_static_routes}"

  #Optional
  cpe_local_identifier      = "${var.ip_sec_connection_cpe_local_identifier}"
  cpe_local_identifier_type = "${var.ip_sec_connection_cpe_local_identifier_type}"
  defined_tags              = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.ip_sec_connection_defined_tags_value}")}"
  display_name              = "${var.ip_sec_connection_display_name}"
  freeform_tags             = "${var.ip_sec_connection_freeform_tags}"

  tunnel_configuration {
    #Optional
    bgp_session_config {
      #Optional
      customer_bgp_asn      = "${var.ip_sec_connection_tunnel_configuration_bgp_session_config_customer_bgp_asn}"
      customer_interface_ip = "${var.ip_sec_connection_tunnel_configuration_bgp_session_config_customer_interface_ip}"
      oracle_interface_ip   = "${var.ip_sec_connection_tunnel_configuration_bgp_session_config_oracle_interface_ip}"
    }

    display_name  = "${var.ip_sec_connection_tunnel_configuration_display_name}"
    routing       = "${var.ip_sec_connection_tunnel_configuration_routing}"
    shared_secret = "${var.ip_sec_connection_tunnel_configuration_shared_secret}"
  }
}

data "oci_core_ipsec_connections" "test_ip_sec_connections" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  cpe_id = "${oci_core_cpe.test_cpe.id}"
  drg_id = "${oci_core_drg.test_drg.id}"
}

data "oci_core_ipsec_connection_tunnels" "test_ip_sec_connection_tunnels" {
  ipsec_id = "${oci_core_ipsec.test_ip_sec_connection.id}"

  # tunnel_id = ""
}

data "oci_core_ipsec_connection_tunnel" "test_ipsec_connection_tunnel" {
  ipsec_id  = "${oci_core_ipsec.test_ip_sec_connection.id}"
  tunnel_id = "${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}"
}
